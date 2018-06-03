package spotinst

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_launch_configuration"
)

func createElastigroupResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.ElastigroupAwsResourceName), name)
}

func testElastigroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.ElastigroupAwsResourceName) {
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

func testCheckElastigroupAttributes(group *aws.Group, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(group.Name) != expectedName {
			return fmt.Errorf("bad content: %v", group.Name)
		}
		return nil
	}
}

func testCheckElastigroupExists(group *aws.Group, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
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

type GroupConfigMetadata struct {
	groupName            string
	instanceTypes        string
	launchConfig         string
	strategy             string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createElastigroupTerraform(gcm *GroupConfigMetadata) string {
	if gcm == nil {
		return ""
	}

	if gcm.instanceTypes == "" {
		gcm.instanceTypes = testInstanceTypesGroupConfig_Create
	}

	if gcm.launchConfig == "" {
		gcm.launchConfig = testLaunchConfigurationGroupConfig_Create
	}

	if gcm.strategy == "" {
		gcm.strategy = testStrategyGroupConfig_Create
	}

	template := ""
	if gcm.updateBaselineFields {
		template = fmt.Sprintf(testBaselineGroupConfig_Update,
			gcm.groupName,
			gcm.groupName,
			gcm.instanceTypes,
			gcm.launchConfig,
			gcm.strategy,
			gcm.fieldsToAppend,
		)
	} else {
		template = fmt.Sprintf(testBaselineGroupConfig_Create,
			gcm.groupName,
			gcm.groupName,
			gcm.instanceTypes,
			gcm.launchConfig,
			gcm.strategy,
			gcm.fieldsToAppend,
		)
	}

	log.Printf("Terraform [%v] template:\n%v", gcm.groupName, template)
	return template
}

// region Elastigroup: Baseline
func TestElastigroupBaseline(t *testing.T) {
	groupName := "eg-baseline"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { TestAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupTerraform(&GroupConfigMetadata{groupName: groupName}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "0"),
					resource.TestCheckResourceAttr(resourceName, "capacity_unit", "weight"),
				),
			},
			{
				Config: createElastigroupTerraform(&GroupConfigMetadata{groupName: groupName, updateBaselineFields: true}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "0"),
					resource.TestCheckResourceAttr(resourceName, "capacity_unit", "weight"),
				),
			},
		},
	})
}

const testBaselineGroupConfig_Create = `
resource "` + string(commons.ElastigroupAwsResourceName) + `" "%v" {

 name 				= "%v"
 description 		= "created by Terraform"
 product 			= "Linux/UNIX"
 availability_zones = ["us-west-2b", "us-west-2c"]

 // --- CAPACITY ------------
 max_size 		  = 0
 min_size 		  = 0
 desired_capacity = 0
 capacity_unit 	  = "weight"
 // -------------------------
 
 %v
 %v
 %v
 %v

}
`

const testBaselineGroupConfig_Update = `
resource "` + string(commons.ElastigroupAwsResourceName) + `" "%v" {

 name 				= "%v"
 description 		= "created by Terraform"
 product 			= "Linux/UNIX"
 availability_zones = ["us-west-2a"]

 //// --- CAPACITY ------------
 //max_size 		  = 0
 //min_size 		  = 0
 //desired_capacity   = 0
 //capacity_unit 	  = "weight"
 //// -------------------------
 
 %v
 %v
 %v
 %v

}
`

// endregion

// region Elastigroup: Instance Types
func TestElastigroupInstanceTypes(t *testing.T) {
	groupName := "eg-instance-types"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { TestAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testInstanceTypesGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "instance_types_ondemand", "m3.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.0", "m3.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.1", "m3.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_weights.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_weights.1229520056.instance_type", "m4.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_weights.1229520056.weight", "2"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_weights.2520866583.instance_type = m3.large", "m3.large"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_weights.2520866583.weight", "1"),
					// Add more asserts
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testInstanceTypesGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "instance_types_ondemand", "c3.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_weights.#", "1"),
					// Add more asserts
				),
			},
		},
	})
}

const testInstanceTypesGroupConfig_Create = `
 // --- INSTANCE TYPES --------------------------------
 instance_types_ondemand = "m3.2xlarge"
 instance_types_spot 	 = ["m3.xlarge", "m3.2xlarge"]
 instance_types_weights  = [
  {
    instance_type = "m3.large"
    weight        = 1
  },
  {
    instance_type = "m4.xlarge"
    weight        = 2
  }]
 // ---------------------------------------------------
`

const testInstanceTypesGroupConfig_Update = `
 // --- INSTANCE TYPES --------------------------------
 instance_types_ondemand = "c3.2xlarge"
 instance_types_spot 	 = ["c2.2xlarge", "c3.xlarge", "c4.2xlarge"]
 instance_types_weights  = [
  {
    instance_type = "c3.xlarge"
    weight        = 3
  }]
 // ---------------------------------------------------
`

// endregion

// region Elastigroup: Launch Configuration
func TestElastigroupLaunchConfiguration(t *testing.T) {
	groupName := "eg-launch-configuration"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { TestAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testLaunchConfigurationGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "ami-a27d8fda"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "my-key.ssh"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "security_groups.1231423", "sg-123456"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_launch_configuration.HexStateFunc("echo hello world")),
					resource.TestCheckResourceAttr(resourceName, "enable_monitoring", "false"),
					resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "false"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testLaunchConfigurationGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "ami-a27d8fda"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "my-key-updated.ssh"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "2"),
					//resource.TestCheckResourceAttr(resourceName, "security_groups.1231423", "sg-123456"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_launch_configuration.HexStateFunc("echo hello world updated")),
					resource.TestCheckResourceAttr(resourceName, "enable_monitoring", "true"),
					resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "false"),
				),
			},
		},
	})
}

const testLaunchConfigurationGroupConfig_Create = `
 // --- LAUNCH CONFIGURATION --------------
 image_id             = "ami-a27d8fda"
 iam_instance_profile = "iam-profile"
 key_name             = "my-key.ssh"
 security_groups      = ["sg-123456"]
 user_data            = "echo hello world"
 enable_monitoring    = false
 ebs_optimized        = false
 placement_tenancy    = "default"
 // ---------------------------------------
`

const testLaunchConfigurationGroupConfig_Update = `
 // --- LAUNCH CONFIGURATION --------------
 image_id             = "ami-a27d8fda"
 iam_instance_profile = "iam-profile updated"
 key_name             = "my-key-updated.ssh"
 security_groups      = ["sg-987654", "sg-123456"]
 user_data            = "echo hello world updated"
 enable_monitoring    = true
 ebs_optimized        = false
 placement_tenancy    = "default"
 // ---------------------------------------
`

// endregion

// region Elastigroup: Strategy
func createElastigroupStrategy(resourceName string, groupName string, fields string) string {
	return fmt.Sprintf(testBaselineGroupConfig_Create, resourceName, groupName, fields)
}

func TestElastigroupStrategy(t *testing.T) {
	groupName := "eg-instance-types"
	resourceName := createElastigroupResourceName(groupName)
	description := "created by Terraform"

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { TestAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       createElastigroupTerraform(&GroupConfigMetadata{groupName: groupName}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "product", "Linux/UNIX"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_launch_configuration.HexStateFunc("echo hello world")),
					resource.TestCheckResourceAttr(resourceName, "orientation", "balanced"),
					resource.TestCheckResourceAttr(resourceName, "fallback_to_ondemand", "false"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_ondemand", "m3.2xlarge"),
				),
			},
			{
				ResourceName: resourceName,
				Config:       createElastigroupTerraform(&GroupConfigMetadata{groupName: groupName, updateBaselineFields: true}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "product", "Linux/UNIX"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_launch_configuration.HexStateFunc("echo hello world updated")),
					resource.TestCheckResourceAttr(resourceName, "orientation", "costOriented"),
					resource.TestCheckResourceAttr(resourceName, "fallback_to_ondemand", "true"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_ondemand", "c3.2xlarge"),
				),
			},
		},
	})
}

const testStrategyGroupConfig_Create = `
 // --- STRATEGY ---------------------
 orientation = "balanced"
 fallback_to_ondemand = false
 spot_percentage = 100
 lifetime_period = ""
 draining_timeout = 50
 utilize_reserved_instances = true
 // ---------------------------------
`

const testStrategyGroupConfig_Update = `
 // --- STRATEGY ---------------------
 orientation = "balanced"
 fallback_to_ondemand = false
 //spot_percentage = 100
 ondemand_count = 1
 lifetime_period = ""
 draining_timeout = 50
 utilize_reserved_instances = true
 // ---------------------------------
`

// endregion

// region Elastigroup: Network Interfaces
const testNetworkInterfacesGroupConfig_Create = `
 // --- NETWORK INTERFACE ------------------
 network_interface = [{ 
    description = ""
    device_index = 1
    secondary_private_ip_address_count = 1
    associate_public_ip_address = true
    delete_on_termination = false
    network_interface_id = ""
    private_ip_address = "1.1.1.1"
  }]
 // ----------------------------------------
`

const testNetworkInterfacesGroupConfig_Update = `
 // --- NETWORK INTERFACE ------------------
 network_interface = [{ 
    description = ""
    device_index = 1
    secondary_private_ip_address_count = 1
    associate_public_ip_address = true
    delete_on_termination = false
    network_interface_id = ""
    private_ip_address = "1.1.1.1"
  }]
 // ----------------------------------------
`

// endregion

// region Elastigroup: Scaling Policies
func createElastigroupScaleUpPolicy(resourceName string, groupName string, fields string) string {
	return fmt.Sprintf(testBaselineGroupConfig_Create, resourceName, groupName, fields)
}

const testScaleUpPolicyGroupConfig_Create = `
 // --- SCALE UP POLICY ------------------
 scaling_up_policy = [{
  policy_name = "policy-name"
  metric_name = "CPUUtilization"
  namespace = "AWS/EC2"
  source = ""
  statistic = "average"
  unit = ""
  cooldown = 60
  dimensions = {
      name = "name-1"
      value = "value-1"
  }
  threshold = 10

  operator = "gt"
  evaluation_periods = 10
  period = 60

  // === MIN TARGET ===================
  action_type = "setMinTarget"
  min_target_capacity = 1
  // ==================================

  // === ADJUSTMENT ===================
  # action_type = "adjustment"
  # action_type = "percentageAdjustment"
  # adjustment = "MAX(5,10)"
  // ==================================

  // === UPDATE CAPACITY ==============
  # action_type = "updateCapacity"
  # minimum = 0
  # maximum = 10
  # target = 5
  // ==================================

  }]
 // ----------------------------------------
`

const testScaleUpPolicyGroupConfig_Update = `
 // --- SCALE UP POLICY ------------------
 scaling_up_policy = [{
  policy_name = "policy-name"
  metric_name = "CPUUtilization"
  namespace = "AWS/EC2"
  source = ""
  statistic = "average"
  unit = ""
  cooldown = 60
  dimensions = {
      name = "name-1"
      value = "value-1"
  }
  threshold = 10

  operator = "gt"
  evaluation_periods = 10
  period = 60

  // === MIN TARGET ===================
  action_type = "setMinTarget"
  min_target_capacity = 1
  // ==================================

  // === ADJUSTMENT ===================
  # action_type = "adjustment"
  # action_type = "percentageAdjustment"
  # adjustment = "MAX(5,10)"
  // ==================================

  // === UPDATE CAPACITY ==============
  # action_type = "updateCapacity"
  # minimum = 0
  # maximum = 10
  # target = 5
  // ==================================

  }]
 // ----------------------------------------
`

// endregion

// region Elastigroup: Scheduled Tasks
func createElastigroupScheduledTask(resourceName string, groupName string, fields string) string {
	return fmt.Sprintf(testBaselineGroupConfig_Create, resourceName, groupName, fields)
}

const testScheduledTaskGroupConfig_Create = `
 // --- SCHEDULED TASK ------------------
  scheduled_task = [{
    is_enabled = false
    start_time = "2018-05-29T02:00:00Z"
    task_type = "backup_ami"
    frequency = "hourly"
    cron_expression = ""
    scale_target_capacity = 5
    scale_min_capacity = 0
    scale_max_capacity = 10
    batch_size_percentage = 33
    grace_period = 300
    target_capacity = 5
    min_capacity = 0
    max_capacity = 10
  }]
 // -------------------------------------
`

const testScheduledTaskGroupConfig_Update = `
 // --- SCHEDULED TASK ------------------
  scheduled_task = [{
    is_enabled = false
    start_time = "2018-05-29T02:00:00Z"
    task_type = "backup_ami"
    frequency = "hourly"
    cron_expression = ""
    scale_target_capacity = 5
    scale_min_capacity = 0
    scale_max_capacity = 10
    batch_size_percentage = 33
    grace_period = 300
    target_capacity = 5
    min_capacity = 0
    max_capacity = 10
  }]
 // -------------------------------------
`

// endregion

// region Elastigroup: Stateful
const testStatefulGroupConfig_Create = `
 // --- STATEFUL ----------------------
 persist_root_device = false
 persist_block_devices = false
 persist_private_ip = true
 block_devices_mode = "onLaunch"
 # block_devices_mode = "reattach"
 private_ips = ["1.1.1.1", "2.2.2.2"]
 // -----------------------------------
`

const testStatefulGroupConfig_Update = `
// --- STATEFUL ----------------------
 persist_root_device = false
 persist_block_devices = false
 persist_private_ip = true
 block_devices_mode = "onLaunch"
 # block_devices_mode = "reattach"
 private_ips = ["1.1.1.1", "2.2.2.2"]
 // -----------------------------------
`

// endregion

// region Elastigroup: Block Devices
//func TestElastigroupBlockDevices(t *testing.T) {
//	blockDeviceResource := "elastigroup-block-devices"
//	groupName := createElastigroupResourceName(blockDeviceResource)
//	name := "TF-Tests-Block-Devices"
//	description := "created by Terraform"
//
//	var group aws.Group
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { TestAccPreCheck(t) },
//		Providers:    TestAccProviders,
//		CheckDestroy: testElastigroupDestroy,
//
//		Steps: []resource.TestStep{
//			{
//				Config: createElastigroupBlockDevicesConfig(blockDeviceResource, name, testElastigroupBlockDevices_Create),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckElastigroupExists(groupName, &group),
//					testCheckElastigroupAttributes(&group, name),
//					resource.TestCheckResourceAttr(groupName, "name", name),
//					resource.TestCheckResourceAttr(groupName, "description", description),
//					resource.TestCheckResourceAttr(groupName, "ebs_block_device.272394590.delete_on_termination", "true"),
//					resource.TestCheckResourceAttr(groupName, "ebs_block_device.272394590.device_name", "/dev/sdb"),
//					resource.TestCheckResourceAttr(groupName, "ebs_block_device.272394590.encrypted", "false"),
//					resource.TestCheckResourceAttr(groupName, "ebs_block_device.272394590.iops", "1"),
//					resource.TestCheckResourceAttr(groupName, "ebs_block_device.272394590.snapshot_id", ""),
//					resource.TestCheckResourceAttr(groupName, "ebs_block_device.272394590.volume_size", "12"),
//					resource.TestCheckResourceAttr(groupName, "ebs_block_device.272394590.volume_type", "gp2"),
//					resource.TestCheckResourceAttr(groupName, "ephemeral_block_device.3796236554.device_name", "/dev/xvdc"),
//					resource.TestCheckResourceAttr(groupName, "ephemeral_block_device.3796236554.virtual_name", "ephemeral0"),
//				),
//			},
//			{
//				Config: createElastigroupBlockDevicesConfig(blockDeviceResource, name, testElastigroupBlockDevices_Update),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckElastigroupExists(groupName, &group),
//					testCheckElastigroupAttributes(&group, name),
//					resource.TestCheckResourceAttr(groupName, "name", name),
//					resource.TestCheckResourceAttr(groupName, "description", description),
//					resource.TestCheckResourceAttr(groupName, "ebs_block_device.66039894.delete_on_termination", "true"),
//					resource.TestCheckResourceAttr(groupName, "ebs_block_device.66039894.device_name", "/dev/sda"),
//					resource.TestCheckResourceAttr(groupName, "ebs_block_device.66039894.encrypted", "true"),
//					resource.TestCheckResourceAttr(groupName, "ebs_block_device.66039894.iops", "1"),
//					resource.TestCheckResourceAttr(groupName, "ebs_block_device.66039894.snapshot_id", ""),
//					resource.TestCheckResourceAttr(groupName, "ebs_block_device.66039894.volume_size", "10"),
//					resource.TestCheckResourceAttr(groupName, "ebs_block_device.66039894.volume_type", "sc1"),
//					resource.TestCheckResourceAttr(groupName, "ephemeral_block_device.4217292875.device_name", "/dev/xvdc"),
//					resource.TestCheckResourceAttr(groupName, "ephemeral_block_device.4217292875.virtual_name", "ephemeral1"),
//				),
//			},
//		},
//	})
//}

func createElastigroupBlockDevicesConfig(resourceName string, groupName string, fields string) string {
	return fmt.Sprintf(testBaselineGroupConfig_Create, resourceName, groupName, fields)
}

const testElastigroupBlockDevices_Create = `
// --- EBS BLOCK DEVICE -----------------
ebs_block_device [{
   device_name 			    = "/dev/sdb"
   snapshot_id 				= ""
   volume_type 				= "gp2"
   volume_size 				= 12
   iops 					= 1
   delete_on_termination 	= true
   encrypted 				= false
}]
// --------------------------------------

// --- EPHEMERAL BLOCK DEVICE ----
ephemeral_block_device [{
  device_name  = "/dev/xvdc"
  virtual_name = "ephemeral0"
}]
// -------------------------------
`

const testElastigroupBlockDevices_Update = `
// --- EBS BLOCK DEVICE -----------------
ebs_block_device [{
   device_name 				= "/dev/sda"
   snapshot_id 				= ""
   volume_type 				= "sc1"
   volume_size 				= 10
   iops 					= 1
   delete_on_termination 	= true
   encrypted 				= true
}]
// --------------------------------------

// --- EPHEMERAL BLOCK DEVICE ----
ephemeral_block_device [{
  device_name  = "/dev/xvdc"
  virtual_name = "ephemeral1"
}]
// -------------------------------
`

// endregion

//testCheckElastigroupBlockDevices(&group.Compute.LaunchSpecification.BlockDeviceMappings,
//"ebs_block_device.272394590",
//map[string]interface{}{
//"delete_on_termination": "true",
//"device_name":           "/dev/sdb",
//"encrypted": "false",
//"iops": "1",
//"snapshot_id": "",
//"volume_size": "12",
//"volume_type": "gp2",
//}),
