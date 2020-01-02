package spotinst

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

func init() {
	resource.AddTestSweepers("spotinst_elastigroup_beanstalk", &resource.Sweeper{
		Name: "spotinst_elastigroup_beanstalk",
		F:    testSweepElastigroupBeanstalk,
	})
}

func testSweepElastigroupBeanstalk(region string) error {
	client, err := getProviderClient("aws")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).elastigroup.CloudProviderAWS()

	input := &aws.ListGroupsInput{}
	if resp, err := conn.List(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of groups to sweep")
	} else {
		if len(resp.Groups) == 0 {
			log.Printf("[INFO] No groups to sweep")
		}
		for _, group := range resp.Groups {
			if strings.Contains(spotinst.StringValue(group.Name), "test-acc-") {
				if _, err := conn.Delete(context.Background(), &aws.DeleteGroupInput{GroupID: group.ID}); err != nil {
					return fmt.Errorf("unable to delete group %v in sweep", spotinst.StringValue(group.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(group.ID))
				}
			}
		}
	}
	return nil
}

func createElastigroupAWSBeanstalkResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.ElastigroupAWSBeanstalkResourceName), name)
}

func testElastigroupAWSBeanstalkDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.ElastigroupAWSBeanstalkResourceName) {
			continue
		}
		input := &aws.ReadGroupInput{GroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.elastigroup.CloudProviderAWS().Read(context.Background(), input)
		if err == nil && resp != nil && resp.Group != nil {
			return fmt.Errorf("group still exists")
		}
	}
	return nil
}

func testCheckElastigroupAWSBeanstalkAttributes(group *aws.Group, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(group.Name) != expectedName {
			return fmt.Errorf("bad content: %v", group.Name)
		}
		return nil
	}
}

func testCheckElastigroupAWSBeanstalkExists(group *aws.Group, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &aws.ReadGroupInput{GroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.elastigroup.CloudProviderAWS().Read(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.Group.Name) != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Group not found: %+v,\n %+v\n", resp.Group, rs.Primary.Attributes)
		}
		*group = *resp.Group
		return nil
	}
}

type BeanstalkGroupConfigMetadata struct {
	groupName            string
	provider             string
	updateBaselineFields bool
}

func createElastigroupAWSBeanstalkTerraform(gcm *BeanstalkGroupConfigMetadata, update string, create string) string {
	if gcm == nil {
		return ""
	}

	if gcm.provider == "" {
		gcm.provider = "aws"
	}

	template :=
		`provider "aws" {
	token   = "fake"
	account = "fake"
	}
	`
	if gcm.updateBaselineFields {
		format := update

		template += fmt.Sprintf(format,
			gcm.groupName,
			gcm.provider,
			gcm.groupName,
		)
	} else {
		format := create
		template += fmt.Sprintf(format,
			gcm.groupName,
			gcm.provider,
			gcm.groupName,
		)
	}

	log.Printf("Terraform [%v] template:\n%v", gcm.groupName, template)
	return template
}

// region Beanstalk Elastigroup: Baseline
func TestAccSpotinstElastigroupAWSBeanstalk_Baseline(t *testing.T) {
	groupName := "test-acc-bs-baseline"
	resourceName := createElastigroupAWSBeanstalkResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAWSBeanstalkDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAWSBeanstalkTerraform(&BeanstalkGroupConfigMetadata{groupName: groupName}, testBaselineBeanstalkGroupConfig_Update, testBaselineBeanstalkGroupConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAWSBeanstalkExists(&group, resourceName),
					testCheckElastigroupAWSBeanstalkAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west-2"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "2"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "1"),
					resource.TestCheckResourceAttr(resourceName, "beanstalk_environment_name", "TerraformDoNotDelete-env"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.0", "t2.small"),
				),
			},
			{
				Config: createElastigroupAWSBeanstalkTerraform(&BeanstalkGroupConfigMetadata{groupName: groupName, updateBaselineFields: true}, testBaselineBeanstalkGroupConfig_Update, testBaselineBeanstalkGroupConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAWSBeanstalkExists(&group, resourceName),
					testCheckElastigroupAWSBeanstalkAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west-2"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "3"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "1"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "2"),
					resource.TestCheckResourceAttr(resourceName, "beanstalk_environment_name", "TerraformDoNotDelete-env"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.0", "t2.small"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.1", "t2.medium"),
				),
			},
		},
	})
}

const testBaselineBeanstalkGroupConfig_Create = `
resource "` + string(commons.ElastigroupAWSBeanstalkResourceName) + `" "%v" {
 provider = "%v"

 name 	 = "%v"
 product = "Linux/UNIX"
 region  = "us-west-2"

 max_size 		  = 2
 min_size 		  = 0
 desired_capacity = 1

 beanstalk_environment_name = "TerraformDoNotDelete-env"
 instance_types_spot        = ["t2.small"]

}

`

const testBaselineBeanstalkGroupConfig_Update = `
resource "` + string(commons.ElastigroupAWSBeanstalkResourceName) + `" "%v" {
 provider = "%v"

 name 	 = "%v"
 product = "Linux/UNIX"
 region  = "us-west-2"

 max_size 		  = 3
 min_size 		  = 1
 desired_capacity = 2

 beanstalk_environment_name = "TerraformDoNotDelete-env"
 instance_types_spot        = ["t2.small", "t2.medium"]

}

`

func TestAccSpotinstElastigroupAWSBeanstalk_Full(t *testing.T) {
	groupName := "test-acc-bs-baseline"
	resourceName := createElastigroupAWSBeanstalkResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAWSBeanstalkDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAWSBeanstalkTerraform(&BeanstalkGroupConfigMetadata{groupName: groupName}, testFullBeanstalkGroupConfig_Update, testFullBeanstalkGroupConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAWSBeanstalkExists(&group, resourceName),
					testCheckElastigroupAWSBeanstalkAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west-2"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "2"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "1"),
					resource.TestCheckResourceAttr(resourceName, "beanstalk_environment_id", "e-h3kze2cv9e"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.0", "t2.small"),
					resource.TestCheckResourceAttr(resourceName, "deployment_preferences.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_preferences.0.automatic_roll", "true"),
					resource.TestCheckResourceAttr(resourceName, "deployment_preferences.0.batch_size_percentage", "100"),
					resource.TestCheckResourceAttr(resourceName, "deployment_preferences.0.grace_period", "90"),
					resource.TestCheckResourceAttr(resourceName, "deployment_preferences.0.strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_preferences.0.strategy.0.action", "REPLACE_SERVER"),
					resource.TestCheckResourceAttr(resourceName, "deployment_preferences.0.strategy.0.should_drain_instances", "true"),
				),
			},
			{
				Config: createElastigroupAWSBeanstalkTerraform(&BeanstalkGroupConfigMetadata{groupName: groupName, updateBaselineFields: true}, testFullBeanstalkGroupConfig_Update, testFullBeanstalkGroupConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAWSBeanstalkExists(&group, resourceName),
					testCheckElastigroupAWSBeanstalkAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west-2"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "3"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "1"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "2"),
					resource.TestCheckResourceAttr(resourceName, "beanstalk_environment_id", "e-h3kze2cv9e"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.0", "t2.small"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.1", "t2.medium"),
					resource.TestCheckResourceAttr(resourceName, "deployment_preferences.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_preferences.0.automatic_roll", "true"),
					resource.TestCheckResourceAttr(resourceName, "deployment_preferences.0.batch_size_percentage", "100"),
					resource.TestCheckResourceAttr(resourceName, "deployment_preferences.0.grace_period", "90"),
					resource.TestCheckResourceAttr(resourceName, "deployment_preferences.0.strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deployment_preferences.0.strategy.0.action", "REPLACE_SERVER"),
					resource.TestCheckResourceAttr(resourceName, "deployment_preferences.0.strategy.0.should_drain_instances", "true"),
				),
			},
		},
	})
}

const testFullBeanstalkGroupConfig_Create = `
resource "` + string(commons.ElastigroupAWSBeanstalkResourceName) + `" "%v" {
 provider = "%v"

 name 	 = "%v"
 product = "Linux/UNIX"
 region  = "us-west-2"

 max_size 		  = 2
 min_size 		  = 0
 desired_capacity = 1

 beanstalk_environment_id = "e-h3kze2cv9e"
 instance_types_spot        = ["t2.small"]

 deployment_preferences {
  automatic_roll        = true
  batch_size_percentage = 100
  grace_period          = 90

  strategy {
   action                 = "REPLACE_SERVER"
   should_drain_instances = true
  }
 }
}

`

const testFullBeanstalkGroupConfig_Update = `
resource "` + string(commons.ElastigroupAWSBeanstalkResourceName) + `" "%v" {
 provider = "%v"

 name 	 = "%v"
 product = "Linux/UNIX"
 region  = "us-west-2"

 max_size 		  = 3
 min_size 		  = 1
 desired_capacity = 2

 beanstalk_environment_id = "e-h3kze2cv9e"
 instance_types_spot        = ["t2.small", "t2.medium"]

 deployment_preferences {
  automatic_roll        = true
  batch_size_percentage = 100
  grace_period          = 90

  strategy {
   action                 = "REPLACE_SERVER"
   should_drain_instances = true
  }
 }
}

`

// region Beanstalk Elastigroup: Scheduled Tasks
func TestAccSpotinstElastigroupAWSBeanstalk_ScheduledTask(t *testing.T) {
	groupName := "test-acc-bs-scheduled-task"
	resourceName := createElastigroupAWSBeanstalkResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAWSBeanstalkDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAWSBeanstalkTerraform(&BeanstalkGroupConfigMetadata{groupName: groupName}, testScheduledTaskBeanstalkGroupConfig_Update, testScheduledTaskBeanstalkGroupConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAWSBeanstalkExists(&group, resourceName),
					testCheckElastigroupAWSBeanstalkAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3981839301.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3981839301.task_type", "backup_ami"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3981839301.scale_min_capacity", "0"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3981839301.scale_max_capacity", "10"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3981839301.adjustment", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3981839301.frequency", "hourly"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3981839301.scale_target_capacity", "5"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3981839301.batch_size_percentage", "33"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3981839301.grace_period", "300"),
				),
			},
			{
				Config: createElastigroupAWSBeanstalkTerraform(&BeanstalkGroupConfigMetadata{groupName: groupName, updateBaselineFields: true}, testScheduledTaskBeanstalkGroupConfig_Update, testScheduledTaskBeanstalkGroupConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAWSBeanstalkExists(&group, resourceName),
					testCheckElastigroupAWSBeanstalkAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2572384164.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2572384164.task_type", "backup_ami"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2572384164.scale_min_capacity", "0"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2572384164.scale_max_capacity", "10"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2572384164.adjustment_percentage", "50"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2572384164.frequency", "hourly"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2572384164.scale_target_capacity", "5"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2572384164.batch_size_percentage", "33"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2572384164.grace_period", "300"),
				),
			},
			{
				Config: createElastigroupAWSBeanstalkTerraform(&BeanstalkGroupConfigMetadata{groupName: groupName, updateBaselineFields: true}, testScheduledTaskBeanstalkGroupConfig_Update2, testScheduledTaskBeanstalkGroupConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAWSBeanstalkExists(&group, resourceName),
					testCheckElastigroupAWSBeanstalkAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.1172895501.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.1172895501.task_type", "statefulUpdateCapacity"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.1172895501.target_capacity", "2"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.1172895501.min_capacity", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.1172895501.max_capacity", "3"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.1172895501.start_time", "2100-01-01T00:00:00Z"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.1172895501.cron_expression", "0 0 12 1/1 * ? *"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.1172895501.batch_size_percentage", "66"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.1172895501.grace_period", "150"),
				),
			},
			{
				Config: createElastigroupAWSBeanstalkTerraform(&BeanstalkGroupConfigMetadata{groupName: groupName, updateBaselineFields: true}, testScheduledTaskBeanstalkGroupConfig_EmptyFields, testScheduledTaskBeanstalkGroupConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAWSBeanstalkExists(&group, resourceName),
					testCheckElastigroupAWSBeanstalkAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "0"),
				),
			},
		},
	})
}

const testScheduledTaskBeanstalkGroupConfig_Create = `
 // --- SCHEDULED TASK ------------------
resource "` + string(commons.ElastigroupAWSBeanstalkResourceName) + `" "%v" {
 provider = "%v"

 name 	 = "%v"
 product = "Linux/UNIX"
 region  = "us-west-2"

 max_size 		  = 2
 min_size 		  = 0
 desired_capacity = 1

 beanstalk_environment_id = "e-mzp2rk47dw"
 instance_types_spot        = ["t2.small"]

  scheduled_task {
	is_enabled = false
    task_type = "backup_ami"
    scale_min_capacity = 0
    scale_max_capacity = 10
    adjustment = 1
    frequency = "hourly"
    scale_target_capacity = 5
    batch_size_percentage = 33
    grace_period = 300
  }
}
 // -------------------------------------

`

const testScheduledTaskBeanstalkGroupConfig_Update = `
 // --- SCHEDULED TASK ------------------
resource "` + string(commons.ElastigroupAWSBeanstalkResourceName) + `" "%v" {
 provider = "%v"

 name 	 = "%v"
 product = "Linux/UNIX"
 region  = "us-west-2"

 max_size 		  = 3
 min_size 		  = 1
 desired_capacity = 2

 beanstalk_environment_id = "e-mzp2rk47dw"
 instance_types_spot        = ["t2.small", "t2.medium"]

  scheduled_task {
	is_enabled = false
    task_type = "backup_ami"
    scale_min_capacity = 0
    scale_max_capacity = 10
    adjustment_percentage = 50
    frequency = "hourly"
    scale_target_capacity = 5
    batch_size_percentage = 33
    grace_period = 300
  }
}
 // -------------------------------------
`

const testScheduledTaskBeanstalkGroupConfig_Update2 = `
 // --- SCHEDULED TASK ------------------
resource "` + string(commons.ElastigroupAWSBeanstalkResourceName) + `" "%v" {
 provider = "%v"

 name 	 = "%v"
 product = "Linux/UNIX"
 region  = "us-west-2"

 max_size 		  = 3
 min_size 		  = 1
 desired_capacity = 2

 beanstalk_environment_id = "e-mzp2rk47dw"
 instance_types_spot        = ["t2.small", "t2.medium"]

  scheduled_task {
    is_enabled = true
    task_type = "statefulUpdateCapacity"
    target_capacity = 2
    min_capacity = 1
    max_capacity = 3
    start_time = "2100-01-01T00:00:00Z"
    cron_expression = "0 0 12 1/1 * ? *"
    batch_size_percentage = 66
    grace_period = 150
  }
}
 // -------------------------------------
`

const testScheduledTaskBeanstalkGroupConfig_EmptyFields = `
 // --- SCHEDULED TASK ------------------
resource "` + string(commons.ElastigroupAWSBeanstalkResourceName) + `" "%v" {
 provider = "%v"

 name 	 = "%v"
 product = "Linux/UNIX"
 region  = "us-west-2"

 max_size 		  = 3
 min_size 		  = 1
 desired_capacity = 2

 beanstalk_environment_id = "e-mzp2rk47dw"
 instance_types_spot        = ["t2.small", "t2.medium"]
}
 // -------------------------------------
`

//// endregion

// endregion
