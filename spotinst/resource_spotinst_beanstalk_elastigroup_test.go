package spotinst

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"log"
)

func createBeanstalkElastigroupResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.BeanstalkElastigroupResourceName), name)
}

func testBeanstalkElastigroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.BeanstalkElastigroupResourceName) {
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

func testCheckBeanstalkElastigroupAttributes(group *aws.Group, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(group.Name) != expectedName {
			return fmt.Errorf("bad content: %v", group.Name)
		}
		return nil
	}
}

func testCheckBeanstalkElastigroupExists(group *aws.Group, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProvider.Meta().(*Client)
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
	fieldsToAppend       string
	updateBaselineFields bool
}

func createBeanstalkElastigroupTerraform(gcm *BeanstalkGroupConfigMetadata) string {
	if gcm == nil {
		return ""
	}

	template := ""
	if gcm.updateBaselineFields {
		format := testBaselineBeanstalkGroupConfig_Update

		template = fmt.Sprintf(format,
			gcm.groupName,
			gcm.groupName,
		)
	} else {
		format := testBaselineBeanstalkGroupConfig_Create
		template = fmt.Sprintf(format,
			gcm.groupName,
			gcm.groupName,
		)
	}

	log.Printf("Terraform [%v] template:\n%v", gcm.groupName, template)
	return template
}

// region Beanstalk Elastigroup: Baseline
func TestAccSpotinstBeanstalkElastigroup_Baseline(t *testing.T) {
	groupName := "beanstalk-baseline"
	resourceName := createBeanstalkElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testBeanstalkElastigroupDestroy,

		Steps: []resource.TestStep{
			{
				Config: createBeanstalkElastigroupTerraform(&BeanstalkGroupConfigMetadata{groupName: groupName}),
				Check: resource.ComposeTestCheckFunc(
					testCheckBeanstalkElastigroupExists(&group, resourceName),
					testCheckBeanstalkElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west-2"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "2"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "1"),
					resource.TestCheckResourceAttr(resourceName, "beanstalk_environment_name", "TfBeanstalkAccTest-env"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.0", "t2.small"),
				),
			},
			{
				Config: createBeanstalkElastigroupTerraform(&BeanstalkGroupConfigMetadata{groupName: groupName, updateBaselineFields: true}),
				Check: resource.ComposeTestCheckFunc(
					testCheckBeanstalkElastigroupExists(&group, resourceName),
					testCheckBeanstalkElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west-2"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "3"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "1"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "2"),
					resource.TestCheckResourceAttr(resourceName, "beanstalk_environment_name", "TfBeanstalkAccTest-env"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.0", "t2.small"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.1", "t2.medium"),
				),
			},
		},
	})
}

const testBaselineBeanstalkGroupConfig_Create = `
resource "` + string(commons.BeanstalkElastigroupResourceName) + `" "%v" {

 name 	 = "%v"
 product = "Linux/UNIX"
 region  = "us-west-2"

 max_size 		  = 2
 min_size 		  = 0
 desired_capacity = 1

 beanstalk_environment_name = "TfBeanstalkAccTest-env"
 instance_types_spot        = ["t2.small"]

}

`

const testBaselineBeanstalkGroupConfig_Update = `
resource "` + string(commons.BeanstalkElastigroupResourceName) + `" "%v" {

 name 	 = "%v"
 product = "Linux/UNIX"
 region  = "us-west-2"

 max_size 		  = 3
 min_size 		  = 1
 desired_capacity = 2

 beanstalk_environment_name = "TfBeanstalkAccTest-env"
 instance_types_spot        = ["t2.small", "t2.medium"]

}

`

// endregion
