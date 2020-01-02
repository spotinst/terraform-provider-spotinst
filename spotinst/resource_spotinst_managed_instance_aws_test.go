package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws_launch_configuration"
)

func init() {
	resource.AddTestSweepers("spotinst_managed_instance_aws", &resource.Sweeper{
		Name: "spotinst_managed_instance_aws",
		F:    testSweepManagedInstance,
	})
}

func testSweepManagedInstance(region string) error {
	client, err := getProviderClient("aws")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).managedInstance.CloudProviderAWS()

	input := &aws.ListManagedInstancesInput{}
	if resp, err := conn.List(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of groups to sweep")
	} else {
		if len(resp.ManagedInstances) == 0 {
			log.Printf("[INFO] No groups to sweep")
		}
		for _, managedInstance := range resp.ManagedInstances {
			if strings.Contains(spotinst.StringValue(managedInstance.Name), "test-acc-") {
				if _, err := conn.Delete(context.Background(), &aws.DeleteManagedInstanceInput{ManagedInstanceID: managedInstance.ID}); err != nil {
					return fmt.Errorf("unable to delete managedInstance %v in sweep", spotinst.StringValue(managedInstance.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(managedInstance.ID))
				}
			}
		}
	}
	return nil
}

func createManagedInstanceAWSResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.ManagedInstanceAWSResourceName), name)
}

func testManagedInstanceAWSDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.ManagedInstanceAWSResourceName) {
			continue
		}
		input := &aws.ReadManagedInstanceInput{ManagedInstanceID: spotinst.String(rs.Primary.ID)}
		resp, err := client.managedInstance.CloudProviderAWS().Read(context.Background(), input)
		if err == nil && resp != nil && resp.ManagedInstance != nil {
			return fmt.Errorf("managedInstance still exists")
		}
	}
	return nil
}

func testCheckManagedInstanceAWSAttributes(managedInstance *aws.ManagedInstance, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(managedInstance.Name) != expectedName {
			return fmt.Errorf("bad content: %v", managedInstance.Name)
		}
		return nil
	}
}

func testCheckManagedInstanceAWSExists(managedInstance *aws.ManagedInstance, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &aws.ReadManagedInstanceInput{ManagedInstanceID: spotinst.String(rs.Primary.ID)}
		resp, err := client.managedInstance.CloudProviderAWS().Read(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.ManagedInstance.Name) != rs.Primary.Attributes["name"] {
			return fmt.Errorf("ManagedInstance not found: %+v,\n %+v\n", resp.ManagedInstance, rs.Primary.Attributes)
		}
		*managedInstance = *resp.ManagedInstance
		return nil
	}
}

type ManagedInstanceConfigMetadata struct {
	provider             string
	name                 string
	variables            string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createManagedInstanceTerraform(ccm *ManagedInstanceConfigMetadata) string {
	if ccm == nil {
		return ""
	}

	if ccm.provider == "" {
		ccm.provider = "aws"
	}

	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`

	if ccm.updateBaselineFields {
		format := testBaselineManagedInstanceConfig_Update
		template += fmt.Sprintf(format,
			ccm.name,
			ccm.provider,
			ccm.name,
			ccm.fieldsToAppend,
		)
	} else {
		format := testBaselineManagedInstanceConfig_Create
		template += fmt.Sprintf(format,
			ccm.name,
			ccm.provider,
			ccm.name,
			ccm.fieldsToAppend,
		)
	}

	if ccm.variables != "" {
		template = ccm.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", ccm.name, template)
	return template
}

// region managedInstance: Baseline
func TestAccSpotinstManagedInstanceBaseline(t *testing.T) {
	name := "test-acc-cluster-managed-instance"
	resourceName := createManagedInstanceAWSResourceName(name)

	var cluster aws.ManagedInstance
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testManagedInstanceAWSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name: name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west-2"),
					resource.TestCheckResourceAttr(resourceName, "persist_private_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "persist_block_devices", "true"),
					resource.TestCheckResourceAttr(resourceName, "persist_root_device", "true"),
					resource.TestCheckResourceAttr(resourceName, "block_devices_mode", "reattach"),
					resource.TestCheckResourceAttr(resourceName, "instance_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_types.0", "t3.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "preferred_type", "t3.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-082b5a644766e0e6f"),
					resource.TestCheckResourceAttr(resourceName, "product", "Linux/UNIX"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.0", "subnet-7f3fbf06"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.1", "subnet-79da021e"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.2", "subnet-03b7ed5b"),
					resource.TestCheckResourceAttr(resourceName, "vpc_id", "vpc-9dee6bfa"),
				),
			},
			{
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:                 name,
					updateBaselineFields: true}),
				Check: resource.ComposeTestCheckFunc(
					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "description", "description updated"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west-2"),
					resource.TestCheckResourceAttr(resourceName, "persist_private_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "persist_block_devices", "true"),
					resource.TestCheckResourceAttr(resourceName, "persist_root_device", "false"),
					resource.TestCheckResourceAttr(resourceName, "block_devices_mode", "reattach"),
					resource.TestCheckResourceAttr(resourceName, "instance_types.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "instance_types.0", "t3.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types.1", "t3.medium"),
					resource.TestCheckResourceAttr(resourceName, "preferred_type", "t3.medium"),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-e251209a"),
					resource.TestCheckResourceAttr(resourceName, "product", "Linux/UNIX"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.0", "subnet-03b7ed5b"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.1", "subnet-7f3fbf06"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.2", "subnet-79da021e"),
					resource.TestCheckResourceAttr(resourceName, "vpc_id", "vpc-0821b8599e5ea9d3c"),
				),
			},
		},
	})
}

const testBaselineManagedInstanceConfig_Create = `
resource "` + string(commons.ManagedInstanceAWSResourceName) + `" "%v" {
  provider = "%v"
  name = "%v"
  description = "description"
  region = "us-west-2"
  product = "Linux/UNIX"
  persist_private_ip = "false"
  persist_block_devices = "true"
  persist_root_device = "true"
  block_devices_mode = "reattach"
  subnet_ids = ["subnet-7f3fbf06", "subnet-79da021e", "subnet-03b7ed5b"]
  instance_types = ["t3.xlarge"]
  preferred_type = "t3.xlarge"
  image_id = "ami-082b5a644766e0e6f"
  vpc_id = "vpc-9dee6bfa"
 %v
}
`

const testBaselineManagedInstanceConfig_Update = `
resource "` + string(commons.ManagedInstanceAWSResourceName) + `" "%v" {
  provider = "%v"
  name = "%v"
  description = "description updated"
  region = "us-west-2"
  product = "Linux/UNIX"
  persist_private_ip = "true"
  persist_block_devices = "true"
  persist_root_device = "false"
  block_devices_mode = "reattach"
  subnet_ids = ["subnet-03b7ed5b","subnet-7f3fbf06", "subnet-79da021e"]  
  instance_types = [
    "t3.xlarge",
    "t3.medium",]
  preferred_type = "t3.medium"
  image_id = "ami-e251209a"
  vpc_id = "vpc-0821b8599e5ea9d3c"
  %v
}
`

// endregion

// region managedInstance: Strategy
func TestAccSpotinstManagedInstanceStrategy(t *testing.T) {
	name := "test-acc-cluster-managed-instance-strategy"
	resourceName := createManagedInstanceAWSResourceName(name)

	var cluster aws.ManagedInstance
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testManagedInstanceAWSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceStrategy_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "life_cycle", "on_demand"),
					resource.TestCheckResourceAttr(resourceName, "orientation", "balanced"),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "120"),
					resource.TestCheckResourceAttr(resourceName, "fall_back_to_od", "false"),
					resource.TestCheckResourceAttr(resourceName, "utilize_reserved_instances", "false"),
					resource.TestCheckResourceAttr(resourceName, "optimization_windows.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "optimization_windows.0", "Mon:03:00-Wed:02:20"),
					resource.TestCheckResourceAttr(resourceName, "revert_to_spot.0.perform_at", "never"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceStrategy_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "life_cycle", "spot"),
					resource.TestCheckResourceAttr(resourceName, "orientation", "cheapest"),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "240"),
					resource.TestCheckResourceAttr(resourceName, "fall_back_to_od", "true"),
					resource.TestCheckResourceAttr(resourceName, "utilize_reserved_instances", "true"),
					resource.TestCheckResourceAttr(resourceName, "optimization_windows.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "optimization_windows.0", "Mon:03:30-Wed:02:30"),
					resource.TestCheckResourceAttr(resourceName, "optimization_windows.1", "Mon:00:30-Wed:01:30"),
					resource.TestCheckResourceAttr(resourceName, "revert_to_spot.0.perform_at", "always"),
				),
			},
		},
	})
}

const managedInstanceStrategy_Create = `
 life_cycle = "on_demand"
 orientation = "balanced"
 draining_timeout = 120
 fall_back_to_od = "false"
 utilize_reserved_instances = "false"
 optimization_windows = ["Mon:03:00-Wed:02:20"]
 revert_to_spot {   
  perform_at = "never"
 }
`

const managedInstanceStrategy_Update = `
 life_cycle = "spot"
 orientation = "cheapest"
 draining_timeout = 240
 fall_back_to_od = "true"
 utilize_reserved_instances = "true"
 optimization_windows = ["Mon:03:30-Wed:02:30", "Mon:00:30-Wed:01:30"]
 revert_to_spot { 
 perform_at = "always"
}
`

// endregion

// region managedInstance: HealthCheck
func TestAccSpotinstManagedInstanceHealthCheck(t *testing.T) {
	name := "test-acc-cluster-managed-instance-healthCheck"
	resourceName := createManagedInstanceAWSResourceName(name)

	var cluster aws.ManagedInstance
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testManagedInstanceAWSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceHealthCheck_Create,
				}),
				Check: resource.ComposeTestCheckFunc(

					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "health_check_type", "EC2"),
					resource.TestCheckResourceAttr(resourceName, "auto_healing", "true"),
					resource.TestCheckResourceAttr(resourceName, "grace_period", "180"),
					resource.TestCheckResourceAttr(resourceName, "unhealthy_duration", "60"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceHealthCheck_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "health_check_type", "MULTAI_TARGET_SET"),
					resource.TestCheckResourceAttr(resourceName, "auto_healing", "false"),
					resource.TestCheckResourceAttr(resourceName, "grace_period", "100"),
					resource.TestCheckResourceAttr(resourceName, "unhealthy_duration", "120"),
				),
			},
		},
	})
}

const managedInstanceHealthCheck_Create = `
health_check_type = "EC2"
auto_healing = "true"
grace_period = "180"
unhealthy_duration = "60"
`

const managedInstanceHealthCheck_Update = `
health_check_type = "MULTAI_TARGET_SET"
auto_healing = "false"
grace_period = "100"
unhealthy_duration = "120"
`

// endregion

// region managedInstance: compute
func TestAccSpotinstManagedInstanceCompute(t *testing.T) {
	name := "test-acc-cluster-managed-instance-compute"
	resourceName := createManagedInstanceAWSResourceName(name)

	var cluster aws.ManagedInstance
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testManagedInstanceAWSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceCompute_Create,
				}),
				Check: resource.ComposeTestCheckFunc(

					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_monitoring", "false"),
					resource.TestCheckResourceAttr(resourceName, "placement_tenancy", "default"),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "ecsInstanceRole"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.0", "sg-1a29b065"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.1", "sg-5750fb2f"),
					resource.TestCheckResourceAttr(resourceName, "elastic_ip", "eipalloc-987654"),
					//resource.TestCheckResourceAttr(resourceName, "private_ip", "172.31.100.159"),
					resource.TestCheckResourceAttr(resourceName, "key_pair", "TamirKeyPair"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.2594194374.key", "explicit1"),
					resource.TestCheckResourceAttr(resourceName, "tags.2594194374.value", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.2281712832.key", "explicit2"),
					resource.TestCheckResourceAttr(resourceName, "tags.2281712832.value", "value2"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_aws_launch_configuration.Base64StateFunc("echo hello world")),
					resource.TestCheckResourceAttr(resourceName, "shutdown_script", elastigroup_aws_launch_configuration.Base64StateFunc("echo goodbye world")),
					resource.TestCheckResourceAttr(resourceName, "cpu_credits", "standard"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.1006920623.device_index", "0"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.1006920623.associate_public_ip_address", "false"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.1006920623.associate_ipv6_address", "false"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceCompute_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "true"),
					resource.TestCheckResourceAttr(resourceName, "enable_monitoring", "true"),
					resource.TestCheckResourceAttr(resourceName, "placement_tenancy", "dedicated"),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "ecsInstanceRole"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.0", "sg-1a29b065"),
					resource.TestCheckResourceAttr(resourceName, "elastic_ip", "eipalloc-123456"),
					//resource.TestCheckResourceAttr(resourceName, "private_ip", "2.2.2.2"),
					resource.TestCheckResourceAttr(resourceName, "key_pair", "my-key.ssh"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.2916442246.key", "explicit1-update"),
					resource.TestCheckResourceAttr(resourceName, "tags.2916442246.value", "value1-update"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_aws_launch_configuration.Base64StateFunc("echo hello world updated")),
					resource.TestCheckResourceAttr(resourceName, "shutdown_script", elastigroup_aws_launch_configuration.Base64StateFunc("echo goodbye world updated")),
					resource.TestCheckResourceAttr(resourceName, "cpu_credits", "unlimited"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.3418395336.device_index", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.3418395336.associate_public_ip_address", "true"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.3418395336.associate_ipv6_address", "true")),
			},
		},
	})
}

const managedInstanceCompute_Create = `
elastic_ip = "eipalloc-987654"
//private_ip = "pip"
//launchSpecification
ebs_optimized = "false"
enable_monitoring = "false"
placement_tenancy = "default"
iam_instance_profile = "ecsInstanceRole"
security_group_ids = ["sg-1a29b065","sg-5750fb2f"]
key_pair = "TamirKeyPair"

  tags {
    key = "explicit1"
    value = "value1"
  }

  tags {
    key = "explicit2"
    value = "value2"
  }
 
 user_data = "echo hello world"
 shutdown_script      = "echo goodbye world"
 cpu_credits          = "standard"

network_interface {
   device_index = 0
   associate_public_ip_address = "false"
   associate_ipv6_address = "false"
   }

`

const managedInstanceCompute_Update = `
elastic_ip = "eipalloc-123456"
//private_ip = "pip"
//launchSpecification
ebs_optimized = "true"
enable_monitoring = "true"
placement_tenancy = "dedicated"
iam_instance_profile = "ecsInstanceRole"
security_group_ids = ["sg-1a29b065"]
key_pair = "my-key.ssh" 

  tags {
     key = "explicit1-update"
     value = "value1-update"
   }
 
 user_data = "echo hello world updated"
 shutdown_script      = "echo goodbye world updated"
 cpu_credits          = "unlimited"

network_interface {
   device_index = 1
   associate_public_ip_address = "true"
   associate_ipv6_address = "true"
   }
 

`

// endregion

// region managedInstance: scheduling
func TestAccSpotinstManagedInstanceScheduling(t *testing.T) {
	name := "test-acc-cluster-managed-instance-scheduling"
	resourceName := createManagedInstanceAWSResourceName(name)

	var cluster aws.ManagedInstance
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testManagedInstanceAWSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceScheduling_Create,
				}),
				Check: resource.ComposeTestCheckFunc(

					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.404594403.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.404594403.task_type", "pause"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.404594403.frequency", "hourly"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceScheduling_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3199430203.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3199430203.task_type", "pause"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3199430203.cron_expression", "cron"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceScheduling_Update2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2406866099.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2406866099.task_type", "resume"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2406866099.frequency", "hourly"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2406866099.start_time", "2019-11-20T23:59:59Z"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceScheduling_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "0"),
				),
			},
		},
	})
}

const managedInstanceScheduling_Create = `
  scheduled_task {
    task_type             = "pause"
    frequency             = "hourly"
    is_enabled            = "true"
  }
`

const managedInstanceScheduling_Update = `

  scheduled_task {
    task_type             = "pause"
    cron_expression       = "cron"
    is_enabled            = "true"
  }
`

const managedInstanceScheduling_Update2 = `

  scheduled_task {
      task_type             = "resume"
      start_time = "2019-11-20T23:59:59Z"
       frequency             = "hourly"
      is_enabled            = "true"
    }
`
const managedInstanceScheduling_EmptyFields = `
 // --- SCHEDULED TASK ------------------
 // -------------------------------------
`

// endregion

// region managedInstance: integrations_route53
func TestAccSpotinstManagedInstanceIntegrationsRoute53(t *testing.T) {
	name := "test-acc-cluster-managed-instance-integrations_route53"
	resourceName := createManagedInstanceAWSResourceName(name)

	var cluster aws.ManagedInstance
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testManagedInstanceAWSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceIntegrations_route53_Create,
				}),
				Check: resource.ComposeTestCheckFunc(

					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.2768511080.hosted_zone_id", "id_create"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.2768511080.record_sets.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.2768511080.record_sets.3654964686.name", "test_create"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.2768511080.record_sets.3654964686.use_public_ip", "false"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceIntegrations_route53_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.522430925.hosted_zone_id", "id_update"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.522430925.record_sets.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.522430925.record_sets.2650004135.name", "test_update"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.522430925.record_sets.2650004135.use_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.522430925.record_sets.567353526.name", "test_update_two"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.522430925.record_sets.567353526.use_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.522430925.record_sets.241835256.name", "test_update_three"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.522430925.record_sets.241835256.use_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.3045402889.hosted_zone_id", "new_domain_on_update"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.3045402889.record_sets.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.3045402889.record_sets.2523873097.name", "new_set"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.3045402889.record_sets.2523873097.use_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.3045402889.record_sets.981666619.name", "test_update_default_ip"),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.0.domains.3045402889.record_sets.981666619.use_public_ip", "false"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceIntegrations_route53_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.#", "0"),
				),
			},
		},
	})
}

const managedInstanceIntegrations_route53_Create = `
// --- INTEGRATION: ROUTE53 ----------
integration_route53 {
	domains {
			hosted_zone_id   = "id_create"
            spotinst_acct_id = "act-123456"
			record_sets  {
				name = "test_create"
				use_public_ip = false
			}
		}
	
}
// ------------------------------------
`

const managedInstanceIntegrations_route53_Update = `
// --- INTEGRATION: ROUTE53 ----------
integration_route53 {
	domains {
		hosted_zone_id = "id_update"
		record_sets {
			name = "test_update"
			use_public_ip = true
		}
		record_sets {
			name = "test_update_two"
			use_public_ip = false
		}
		record_sets {
			name = "test_update_three"
			use_public_ip = false
		}
	}

	domains {
		hosted_zone_id = "new_domain_on_update"
		record_sets {
			name = "new_set"
			use_public_ip = true
		}
		record_sets {
			name = "test_update_default_ip"
		}
	}
}
// ------------------------------------
`
const managedInstanceIntegrations_route53_EmptyFields = `
// --- INTEGRATION: ROUTE53 ----------
// ------------------------------------
`

// endregion

// region managedInstance: integrations_route53
func TestAccSpotinstManagedInstanceIntegrationsLoadBalancers(t *testing.T) {
	name := "test-acc-cluster-managed-instance-integrations-load-balancers"
	resourceName := createManagedInstanceAWSResourceName(name)

	var cluster aws.ManagedInstance
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testManagedInstanceAWSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceIntegrations_Load_Balancers_Create,
				}),
				Check: resource.ComposeTestCheckFunc(

					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.217962633.arn", "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.217962633.name", "test_name"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.217962633.type", "TARGET_GROUP"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceIntegrations_Load_Balancers_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.3520850516.arn", "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.3520850516.balancer_id", "lb-1ee2e3q"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.3520850516.target_set_id", "ts-3eq"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.3520850516.type", "MULTAI_TARGET_SET"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceIntegrations_Load_Balancers_Update2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.3464633014.arn", "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.3464633014.balancer_id", "lb-1ee2e3q"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.3464633014.target_set_id", "ts-3eq"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.3464633014.type", "MULTAI_TARGET_SET"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.3464633014.auto_weight", "true"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.3464633014.az_awareness", "true"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createManagedInstanceTerraform(&ManagedInstanceConfigMetadata{
					name:           name,
					fieldsToAppend: managedInstanceIntegrations_Load_Balancers_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckManagedInstanceAWSExists(&cluster, resourceName),
					testCheckManagedInstanceAWSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "0"),
				),
			},
		},
	})
}

const managedInstanceIntegrations_Load_Balancers_Create = `
// --- INTEGRATION: Load_Balancers ----------
  load_balancers {
      arn  = "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456"
      name = "test_name"
      type = "TARGET_GROUP"
	}
// ------------------------------------
`

const managedInstanceIntegrations_Load_Balancers_Update = `
// --- INTEGRATION: load_balancers ----------
  load_balancers {
      arn  = "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456"
      type = "MULTAI_TARGET_SET"
      balancer_id   = "lb-1ee2e3q"
      target_set_id = "ts-3eq"
    }
// ------------------------------------
`
const managedInstanceIntegrations_Load_Balancers_Update2 = `
// --- INTEGRATION: load_balancers ----------
  load_balancers {
      arn  = "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456"
      type = "MULTAI_TARGET_SET"
      balancer_id   = "lb-1ee2e3q"
      target_set_id = "ts-3eq"
      auto_weight   = "true"
      az_awareness = "true"
    }
// ------------------------------------
`

const managedInstanceIntegrations_Load_Balancers_EmptyFields = `
// --- INTEGRATION: load_balancers ----------
// ------------------------------------
`

// endregion
