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
func TestAccSpotinstElastigroup_Baseline(t *testing.T) {
	groupName := "eg-baseline"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupDestroy,

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

// endregion

// region Elastigroup: Instance Types
func TestAccSpotinstElastigroup_InstanceTypes(t *testing.T) {
	groupName := "eg-instance-types"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:     groupName,
					instanceTypes: testInstanceTypesGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "instance_types_ondemand", "m4.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.0", "m4.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.1", "m4.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_weights.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_weights.1650831227.instance_type", "m4.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_weights.1650831227.weight", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_weights.2214348274.instance_type", "m4.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_weights.2214348274.weight", "2"),
				),
			},
			{
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:     groupName,
					instanceTypes: testInstanceTypesGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "instance_types_ondemand", "c4.4xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.0", "c4.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.1", "c4.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_spot.2", "c4.4xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_weights.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_weights.3291405167.instance_type", "c4.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_weights.3291405167.weight", "3"),
				),
			},
		},
	})
}

const testInstanceTypesGroupConfig_Create = `
 // --- INSTANCE TYPES --------------------------------
 instance_types_ondemand = "m4.2xlarge"
 instance_types_spot 	 = ["m4.xlarge", "m4.2xlarge"]
 instance_types_weights  = [
  {
    instance_type = "m4.xlarge"
    weight        = 1
  },
  {
    instance_type = "m4.2xlarge"
    weight        = 2
  }]
 // ---------------------------------------------------
`

const testInstanceTypesGroupConfig_Update = `
 // --- INSTANCE TYPES --------------------------------
 instance_types_ondemand = "c4.4xlarge"
 instance_types_spot 	 = ["c4.xlarge", "c4.2xlarge", "c4.4xlarge"]
 instance_types_weights  = [
  {
    instance_type = "c4.xlarge"
    weight        = 3
  }]
 // ---------------------------------------------------
`

// endregion

// region Elastigroup: Launch Configuration
func TestAccSpotinstElastigroup_LaunchConfiguration(t *testing.T) {
	groupName := "eg-launch-configuration"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:    groupName,
					launchConfig: testLaunchConfigurationGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-e251209a"),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "iam-profile"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "my-key.ssh"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-123456"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_launch_configuration.HexStateFunc("echo hello world")),
					resource.TestCheckResourceAttr(resourceName, "enable_monitoring", "false"),
					resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "false"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:    groupName,
					launchConfig: testLaunchConfigurationGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-31394949"),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "iam-profile updated"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "my-key-updated.ssh"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-123456"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.1", "sg-987654"),
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
 image_id             = "ami-e251209a"
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
 image_id             = "ami-31394949"
 iam_instance_profile = "iam-profile updated"
 key_name             = "my-key-updated.ssh"
 security_groups      = ["sg-123456", "sg-987654"]
 user_data            = "echo hello world updated"
 enable_monitoring    = true
 ebs_optimized        = false
 placement_tenancy    = "default"
 // ---------------------------------------
`

// endregion

// region Elastigroup: Strategy
func TestAccSpotinstElastigroup_Strategy(t *testing.T) {
	groupName := "eg-strategy"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName: groupName,
					strategy:  testStrategyGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "orientation", "balanced"),
					resource.TestCheckResourceAttr(resourceName, "fallback_to_ondemand", "true"),
					resource.TestCheckResourceAttr(resourceName, "spot_percentage", "100"),
					resource.TestCheckResourceAttr(resourceName, "lifetime_period", ""),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "300"),
					resource.TestCheckResourceAttr(resourceName, "utilize_reserved_instances", "true"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName: groupName,
					strategy:  testStrategyGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "orientation", "costOriented"),
					resource.TestCheckResourceAttr(resourceName, "fallback_to_ondemand", "false"),
					resource.TestCheckResourceAttr(resourceName, "ondemand_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "lifetime_period", ""),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "600"),
					resource.TestCheckResourceAttr(resourceName, "utilize_reserved_instances", "false"),
				),
			},
		},
	})
}

const testStrategyGroupConfig_Create = `
 // --- STRATEGY ---------------------
 orientation = "balanced"
 fallback_to_ondemand = true
 spot_percentage = 100
 lifetime_period = ""
 draining_timeout = 300
 utilize_reserved_instances = true
 // ---------------------------------
`

const testStrategyGroupConfig_Update = `
 // --- STRATEGY ---------------------
 orientation = "costOriented"
 fallback_to_ondemand = false
 ondemand_count = 1
 lifetime_period = ""
 draining_timeout = 600
 utilize_reserved_instances = false
 // ---------------------------------
`

// endregion

// region Elastigroup: Load Balancers
func TestAccSpotinstElastigroup_LoadBalancers(t *testing.T) {
	groupName := "eg-load-balancers"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testLoadBalancersGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "elastic_load_balancers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "elastic_load_balancers.0", "bal1"),
					resource.TestCheckResourceAttr(resourceName, "elastic_load_balancers.1", "bal2"),
					resource.TestCheckResourceAttr(resourceName, "target_group_arns.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target_group_arns.0", "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456"),
					resource.TestCheckResourceAttr(resourceName, "multai_target_sets.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "multai_target_sets.2753680074.target_set_id", "ts-123"),
					resource.TestCheckResourceAttr(resourceName, "multai_target_sets.2753680074.balancer_id", "bal-123"),
					resource.TestCheckResourceAttr(resourceName, "multai_target_sets.979814926.target_set_id", "ts-234"),
					resource.TestCheckResourceAttr(resourceName, "multai_target_sets.979814926.balancer_id", "bal-234"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testLoadBalancersGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "elastic_load_balancers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "elastic_load_balancers.0", "bal1"),
					resource.TestCheckResourceAttr(resourceName, "elastic_load_balancers.1", "bal3"),
					resource.TestCheckResourceAttr(resourceName, "target_group_arns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "target_group_arns.0", "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456"),
					resource.TestCheckResourceAttr(resourceName, "target_group_arns.1", "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testNewTargetGroup/1234567890123456"),
					resource.TestCheckResourceAttr(resourceName, "multai_target_sets.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "multai_target_sets.2753680074.target_set_id", "ts-123"),
					resource.TestCheckResourceAttr(resourceName, "multai_target_sets.2753680074.balancer_id", "bal-123"),
				),
			},
		},
	})
}

const testLoadBalancersGroupConfig_Create = `
 // --- LOAD BALANCERS --------------------
 elastic_load_balancers = ["bal1", "bal2"]
 target_group_arns = ["arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456"]
 multai_target_sets = [{
    target_set_id = "ts-123",
    balancer_id = "bal-123"
  },
  {
    target_set_id = "ts-234",
    balancer_id = "bal-234"
  }]
 // ---------------------------------------
`

const testLoadBalancersGroupConfig_Update = `
 // --- LOAD BALANCERS --------------------
 elastic_load_balancers = ["bal1", "bal3"]
 target_group_arns = ["arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456", "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testNewTargetGroup/1234567890123456"]
 multai_target_sets = [{
    target_set_id = "ts-123",
    balancer_id = "bal-123"
  }]
 // ---------------------------------------
`

// endregion

// region Elastigroup: Health Checks
func TestAccSpotinstElastigroup_HealthChecks(t *testing.T) {
	groupName := "eg-health-checks"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testHealthChecksGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "health_check_type", "ELB"),
					resource.TestCheckResourceAttr(resourceName, "health_check_grace_period", "100"),
					resource.TestCheckResourceAttr(resourceName, "health_check_unhealthy_duration_before_replacement", "60"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testHealthChecksGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "health_check_type", "TARGET_GROUP"),
					resource.TestCheckResourceAttr(resourceName, "health_check_grace_period", "50"),
					resource.TestCheckResourceAttr(resourceName, "health_check_unhealthy_duration_before_replacement", "120"),
				),
			},
		},
	})
}

const testHealthChecksGroupConfig_Create = `
 // --- HEALTH-CHECKS ------------------------------------
 health_check_type = "ELB" 
 health_check_grace_period = 100
 health_check_unhealthy_duration_before_replacement = 60
 // ------------------------------------------------------
`

const testHealthChecksGroupConfig_Update = `
 // --- HEALTH-CHECKS ------------------------------------
 health_check_type = "TARGET_GROUP" 
 health_check_grace_period = 50
 health_check_unhealthy_duration_before_replacement = 120
 // ------------------------------------------------------
`

// endregion

// region Elastigroup: Elastic IPs
func TestAccSpotinstElastigroup_ElasticIPs(t *testing.T) {
	groupName := "eg-elastic-ips"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testElasticIPsGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "elastic_ips.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "elastic_ips.0", "eipalloc-123456"),
					resource.TestCheckResourceAttr(resourceName, "elastic_ips.1", "eipalloc-987654"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testElasticIPsGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "elastic_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "elastic_ips.0", "eipalloc-123456"),
				),
			},
		},
	})
}

const testElasticIPsGroupConfig_Create = `
 // --- ELASTIC IPs --------------------------------------
  elastic_ips = ["eipalloc-123456", "eipalloc-987654"]
  // ------------------------------------------------------
`

const testElasticIPsGroupConfig_Update = `
 // --- ELASTIC IPs --------------------------------------
  elastic_ips = ["eipalloc-123456"]
  // ------------------------------------------------------
`

// endregion

// region Elastigroup: Subnet Ids
func TestAccSpotinstElastigroup_SubnetIDs(t *testing.T) {
	groupName := "eg-subnet-ids"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testSubnetIDsGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west-1"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.0", "sub-123"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.1", "sub-456"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testSubnetIDsGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west-2"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.0", "sub-123"),
				),
			},
		},
	})
}

const testSubnetIDsGroupConfig_Create = `
 // --- SUBNET IDS -------------------
 region      = "us-west-1"
 subnet_ids  = ["sub-123", "sub-456"]
 // ----------------------------------
`

const testSubnetIDsGroupConfig_Update = `
 // --- SUBNET IDS -------------------
 region      = "us-west-2"
 subnet_ids  = ["sub-123"]
 // ----------------------------------
`

// endregion

// region Elastigroup: Signals
func TestAccSpotinstElastigroup_Signals(t *testing.T) {
	groupName := "eg-signals"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testSignalsGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "signal.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "signal.1191208186.name", "INSTANCE_READY_TO_SHUTDOWN"),
					resource.TestCheckResourceAttr(resourceName, "signal.1191208186.timeout", "100"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testSignalsGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "signal.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "signal.1191208186.name", "INSTANCE_READY_TO_SHUTDOWN"),
					resource.TestCheckResourceAttr(resourceName, "signal.1191208186.timeout", "100"),
					resource.TestCheckResourceAttr(resourceName, "signal.1735739494.name", "INSTANCE_READY"),
					resource.TestCheckResourceAttr(resourceName, "signal.1735739494t.timeout", "200"),
				),
			},
		},
	})
}

const testSignalsGroupConfig_Create = `
 // --- SIGNAL -----
  signal = {
    name = "INSTANCE_READY_TO_SHUTDOWN"
    timeout = 100
  }
  // ----------------
`

const testSignalsGroupConfig_Update = `
 // --- SIGNAL -----
  signal = [{
    name = "INSTANCE_READY_TO_SHUTDOWN"
    timeout = 100
  },
  {
    name = "INSTANCE_READY"
    timeout = 200
  }]
  // ----------------
`

// endregion

// region Elastigroup: Revert To Spot (Maintenance Window)
func TestAccSpotinstElastigroup_RevertToSpot(t *testing.T) {
	groupName := "eg-revert-to-spot"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testRevertToSpotGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "revert_to_spot.0.perform_at", "always"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testRevertToSpotGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "revert_to_spot.0.perform_at", "timeWindow"),
					resource.TestCheckResourceAttr(resourceName, "revert_to_spot.0.time_windows.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "revert_to_spot.0.time_windows.0", "Mon:12:00-Tue:12:00"),
					resource.TestCheckResourceAttr(resourceName, "revert_to_spot.0.time_windows.1", "Fri:12:00-Sat:12:00"),
				),
			},
		},
	})
}

const testRevertToSpotGroupConfig_Create = `
 // --- REVERT TO SPOT -------------------------------------------
 revert_to_spot {
  perform_at    = "always"
 }
 // -------------------------------------------------------------
`

const testRevertToSpotGroupConfig_Update = `
 // --- REVERT TO SPOT -------------------------------------------
 revert_to_spot {
  perform_at    = "timeWindow"
  time_windows  = ["Mon:12:00-Tue:12:00", "Fri:12:00-Sat:12:00"]
 }
 // -------------------------------------------------------------
`

// endregion

// region Elastigroup: Network Interfaces
func TestAccSpotinstElastigroup_NetworkInterfaces(t *testing.T) {
	groupName := "eg-network-interfaces"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testNetworkInterfacesGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "network_interface.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.1760224316.associate_public_ip_address", "false"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.1760224316.delete_on_termination", "false"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.1760224316.description", "network interface description"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.1760224316.device_index", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.1760224316.network_interface_id", "n-123456"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.1760224316.private_ip_address", "1.1.1.1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.1760224316.secondary_private_ip_address_count", "1"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testNetworkInterfacesGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "network_interface.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.2833641110.associate_public_ip_address", "true"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.2833641110.delete_on_termination", "true"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.2833641110.description", "network interface description updated"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.2833641110.device_index", "2"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.2833641110.network_interface_id", "n-987654"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.2833641110.private_ip_address", "2.2.2.2"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.2833641110.secondary_private_ip_address_count", "2"),
				),
			},
		},
	})
}

const testNetworkInterfacesGroupConfig_Create = `
 // --- NETWORK INTERFACE ------------------
 network_interface = [{ 
    description = "network interface description"
    device_index = 1
    secondary_private_ip_address_count = 1
    associate_public_ip_address = false
    delete_on_termination = false
    network_interface_id = "n-123456"
    private_ip_address = "1.1.1.1"
  }]
 // ----------------------------------------
`

const testNetworkInterfacesGroupConfig_Update = `
 // --- NETWORK INTERFACE ------------------
 network_interface = [{ 
    description = "network interface description updated"
    device_index = 2
    secondary_private_ip_address_count = 2
    associate_public_ip_address = true
    delete_on_termination = true
    network_interface_id = "n-987654"
    private_ip_address = "2.2.2.2"
  }]
 // ----------------------------------------
`

// endregion

// region Elastigroup: Scaling Policies
func TestAccSpotinstElastigroup_ScaleUpPolicy(t *testing.T) {
	groupName := "eg-scale-up-policy"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScaleUpPolicyGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.policy_name", "policy-name"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.metric_name", "CPUUtilization"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.namespace", "AWS/EC2"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.source", "cloudWatch"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.statistic", "average"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.unit", "percent"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.cooldown", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.dimensions.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.dimensions.name", "name-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.dimensions.value", "value-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.threshold", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.operator", "gt"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.evaluation_periods", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.period", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.action_type", "setMinTarget"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.min_target_capacity", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.adjustment", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.max_target_capacity", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.maximum", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.minimum", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.37737847.target", ""),
				),
			},
			//{
			//	ResourceName: resourceName,
			//	Config: createElastigroupTerraform(&GroupConfigMetadata{
			//		groupName:      groupName,
			//		fieldsToAppend: testScaleUpPolicyGroupConfig_Update,
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testCheckElastigroupExists(&group, resourceName),
			//		testCheckElastigroupAttributes(&group, groupName),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.#", "1"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.policy_name", "policy-name-update"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.metric_name", "CPUUtilization"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.namespace", "AWS/EC2"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.source", "spectrum"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.statistic", "sum"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.unit", "bytes"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.cooldown", "120"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.dimensions.%", "2"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.dimensions.name", "name-1-update"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.dimensions.value", "value-1-update"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.threshold", "5"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.operator", "lt"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.evaluation_periods", "5"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.period", "120"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.action_type", "adjustment"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.adjustment", "MAX(5,10)"),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.min_target_capacity", ""),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.max_target_capacity", ""),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.maximum", ""),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.minimum", ""),
			//		resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.1565231540.target", ""),
			//	),
			//},
		},
	})
}

const testScaleUpPolicyGroupConfig_Create = `
 // --- SCALE UP POLICY ------------------
 scaling_up_policy = [{
  policy_name = "policy-name"
  metric_name = "CPUUtilization"
  namespace = "AWS/EC2"
  source = "cloudWatch"
  statistic = "average"
  unit = "percent"
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
  policy_name = "policy-name-update"
  metric_name = "CPUUtilization"
  namespace = "AWS/EC2"
  source = "spectrum"
  statistic = "sum"
  unit = "bytes"
  cooldown = 120
  dimensions = {
      name = "name-1-update"
      value = "value-1-update"
  }
  threshold = 5

  operator = "lt"
  evaluation_periods = 5
  period = 120

  //// === MIN TARGET ===================
  //action_type = "setMinTarget"
  //min_target_capacity = 1
  //// ==================================

  // === ADJUSTMENT ===================
  // action_type = "percentageAdjustment"
  action_type = "adjustment"
  adjustment = "MAX(5,10)"
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
func TestAccSpotinstElastigroup_ScheduledTask(t *testing.T) {
	groupName := "eg-scheduled-task"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScheduledTaskGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3463887611.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3463887611.task_type", "backup_ami"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3463887611.scale_min_capacity", "0"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3463887611.scale_max_capacity", "10"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3463887611.frequency", "hourly"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3463887611.scale_target_capacity", "5"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3463887611.batch_size_percentage", "33"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3463887611.grace_period", "300"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScheduledTaskGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2674842669.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2674842669.task_type", "statefulUpdateCapacity"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2674842669.target_capacity", "2"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2674842669.min_capacity", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2674842669.max_capacity", "3"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2674842669.start_time", "2100-01-01T00:00:00Z"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2674842669.cron_expression", "0 0 12 1/1 * ? *"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2674842669.batch_size_percentage", "66"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.2674842669.grace_period", "150"),
				),
			},
		},
	})
}

const testScheduledTaskGroupConfig_Create = `
 // --- SCHEDULED TASK ------------------
  scheduled_task = [{
	is_enabled = false
    task_type = "backup_ami"
    scale_min_capacity = 0
    scale_max_capacity = 10
    frequency = "hourly"
    scale_target_capacity = 5
    batch_size_percentage = 33
    grace_period = 300
  }]
 // -------------------------------------
`

const testScheduledTaskGroupConfig_Update = `
 // --- SCHEDULED TASK ------------------
  scheduled_task = [{
    is_enabled = true
    task_type = "statefulUpdateCapacity"
    target_capacity = 2
    min_capacity = 1
    max_capacity = 3
    start_time = "2100-01-01T00:00:00Z"
    cron_expression = "0 0 12 1/1 * ? *"
    batch_size_percentage = 66
    grace_period = 150
  }]
 // -------------------------------------
`

// endregion

// region Elastigroup: Stateful
func TestAccSpotinstElastigroup_Stateful(t *testing.T) {
	groupName := "eg-stateful"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testStatefulGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "persist_root_device", "true"),
					resource.TestCheckResourceAttr(resourceName, "persist_block_devices", "true"),
					resource.TestCheckResourceAttr(resourceName, "persist_private_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "block_devices_mode", "reattach"),
					resource.TestCheckResourceAttr(resourceName, "private_ips.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "private_ips.0", "1.1.1.1"),
					resource.TestCheckResourceAttr(resourceName, "private_ips.1", "2.2.2.2"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testStatefulGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "persist_root_device", "false"),
					resource.TestCheckResourceAttr(resourceName, "persist_block_devices", "false"),
					resource.TestCheckResourceAttr(resourceName, "persist_private_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "block_devices_mode", "onLaunch"),
					resource.TestCheckResourceAttr(resourceName, "private_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "private_ips.0", "3.3.3.3"),
				),
			},
		},
	})
}

const testStatefulGroupConfig_Create = `
 // --- STATEFUL ----------------------
 persist_root_device = true
 persist_block_devices = true
 persist_private_ip = true
 block_devices_mode = "reattach"
 private_ips = ["1.1.1.1", "2.2.2.2"]
 // -----------------------------------
`

const testStatefulGroupConfig_Update = `
// --- STATEFUL ----------------------
 persist_root_device = false
 persist_block_devices = false
 persist_private_ip = false
 block_devices_mode = "onLaunch"
 private_ips = ["3.3.3.3"]
 // -----------------------------------
`

// endregion

// region Elastigroup: Block Devices
func TestAccSpotinstElastigroup_BlockDevices(t *testing.T) {
	groupName := "eg-block-devices"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testElastigroupBlockDevices_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.272394590.delete_on_termination", "true"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.272394590.device_name", "/dev/sdb"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.272394590.encrypted", "false"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.272394590.iops", "1"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.272394590.snapshot_id", ""),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.272394590.volume_size", "12"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.272394590.volume_type", "gp2"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.3570307215.delete_on_termination", "false"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.3570307215.device_name", "/dev/sda"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.3570307215.encrypted", "false"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.3570307215.iops", "1"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.3570307215.snapshot_id", ""),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.3570307215.volume_size", "8"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.3570307215.volume_type", "io1"),
					resource.TestCheckResourceAttr(resourceName, "ephemeral_block_device.3796236554.device_name", "/dev/xvdc"),
					resource.TestCheckResourceAttr(resourceName, "ephemeral_block_device.3796236554.virtual_name", "ephemeral0"),
				),
			},
			{
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testElastigroupBlockDevices_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.66039894.delete_on_termination", "true"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.66039894.device_name", "/dev/sda"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.66039894.encrypted", "true"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.66039894.iops", "1"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.66039894.snapshot_id", ""),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.66039894.volume_size", "10"),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.66039894.volume_type", "sc1"),
					resource.TestCheckResourceAttr(resourceName, "ephemeral_block_device.4217292875.device_name", "/dev/xvdc"),
					resource.TestCheckResourceAttr(resourceName, "ephemeral_block_device.4217292875.virtual_name", "ephemeral1"),
				),
			},
		},
	})
}

const testElastigroupBlockDevices_Create = `
 // --- EBS BLOCK DEVICE -----------------
 ebs_block_device = [{
   device_name 			    = "/dev/sdb"
   snapshot_id 				= ""
   volume_type 				= "gp2"
   volume_size 				= 12
   iops 					= 1
   delete_on_termination 	= true
   encrypted 				= false
 },
 {
   device_name 			    = "/dev/sda"
   snapshot_id 				= ""
   volume_type 				= "io1"
   volume_size 				= 8
   iops 					= 1
   delete_on_termination 	= false
   encrypted 				= false
 }]
 // --------------------------------------

 // --- EPHEMERAL BLOCK DEVICE ----
 ephemeral_block_device = [{
  device_name  = "/dev/xvdc"
  virtual_name = "ephemeral0"
 }]
 // -------------------------------
`

const testElastigroupBlockDevices_Update = `
 // --- EBS BLOCK DEVICE -----------------
 ebs_block_device = [{
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
 ephemeral_block_device = [{
  device_name  = "/dev/xvdc"
  virtual_name = "ephemeral1"
 }]
 // -------------------------------
`

// endregion

// region Elastigroup: Tags
func TestAccSpotinstElastigroup_Tags(t *testing.T) {
	groupName := "eg-tags"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testTagsGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.2594194374.key", "explicit1"),
					resource.TestCheckResourceAttr(resourceName, "tags.2594194374.value", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.2281712832.key", "explicit2"),
					resource.TestCheckResourceAttr(resourceName, "tags.2281712832.value", "value2"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testTagsGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.2916442246.key", "explicit1-update"),
					resource.TestCheckResourceAttr(resourceName, "tags.2916442246.value", "value1-update"),
				),
			},
		},
	})
}

const testTagsGroupConfig_Create = `
 // --- TAGS ---------
  tags = [
   {
     key = "explicit1"
     value = "value1"
   },
   {
     key = "explicit2"
     value = "value2"
   }
 ]
 // ------------------
`

const testTagsGroupConfig_Update = `
 // --- TAGS ---------
  tags = [
   {
     key = "explicit1-update"
     value = "value1-update"
   }
 ]
 // ------------------
`

// endregion

// region Elastigroup: Rancher Integration
func TestAccSpotinstElastigroup_IntegrationRancher(t *testing.T) {
	groupName := "eg-integration-rancher"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationRancherGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.0.master_host", "master-host"),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.0.access_key", "access-key"),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.0.secret_key", "secret-key"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationRancherGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.0.master_host", "master-host-update"),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.0.access_key", "access-key-update"),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.0.secret_key", "secret-key-update"),
				),
			},
		},
	})
}

const testIntegrationRancherGroupConfig_Create = `
// --- INTEGRATION: RANCHER ---
integration_rancher = {
   master_host = "master-host"
   access_key = "access-key"
   secret_key = "secret-key"
}
// ----------------------------
`

const testIntegrationRancherGroupConfig_Update = `
 // --- INTEGRATION: RANCHER ---
 integration_rancher = {
    master_host = "master-host-update"
    access_key = "access-key-update"
    secret_key = "secret-key-update"
 }
 // ----------------------------
`

// endregion

// region Elastigroup: Code Deploy Integration
func TestAccSpotinstElastigroup_IntegrationCodeDeploy(t *testing.T) {
	groupName := "eg-integration-code-deploy"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationCodeDeployGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_codedeploy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_codedeploy.0.cleanup_on_failure", "false"),
					resource.TestCheckResourceAttr(resourceName, "integration_codedeploy.0.terminate_instance_on_failure", "false"),
					resource.TestCheckResourceAttr(resourceName, "integration_codedeploy.0.deployment_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_codedeploy.0.deployment_groups.699338831.application_name", "code-deploy-application"),
					resource.TestCheckResourceAttr(resourceName, "integration_codedeploy.0.deployment_groups.699338831.deployment_group_name", "code-deploy-deployment"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationCodeDeployGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_codedeploy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_codedeploy.0.cleanup_on_failure", "true"),
					resource.TestCheckResourceAttr(resourceName, "integration_codedeploy.0.terminate_instance_on_failure", "true"),
					resource.TestCheckResourceAttr(resourceName, "integration_codedeploy.0.deployment_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_codedeploy.0.deployment_groups.2845984724.application_name", "code-deploy-application-update"),
					resource.TestCheckResourceAttr(resourceName, "integration_codedeploy.0.deployment_groups.2845984724.deployment_group_name", "code-deploy-deployment-update"),
				),
			},
		},
	})
}

const testIntegrationCodeDeployGroupConfig_Create = `
 // --- INTEGRATION: CODE-DEPLOY ----------
 integration_codedeploy = {
    cleanup_on_failure = false
    terminate_instance_on_failure = false
    deployment_groups = {
      application_name = "code-deploy-application"
      deployment_group_name = "code-deploy-deployment"
    }
  }
 // ---------------------------------------
`

const testIntegrationCodeDeployGroupConfig_Update = `
 // --- INTEGRATION: CODE-DEPLOY ----------
 integration_codedeploy = {
    cleanup_on_failure = true
    terminate_instance_on_failure = true
    deployment_groups = {
      application_name = "code-deploy-application-update"
      deployment_group_name = "code-deploy-deployment-update"
    }
  }
 // ---------------------------------------
`

// endregion

// region Elastigroup: ECS Integration
func TestAccSpotinstElastigroup_IntegrationECS(t *testing.T) {
	groupName := "eg-integration-ecs"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationECSGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.cluster_name", "ecs-cluster-name"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_cooldown", "300"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.0.num_of_units", "2"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_down.0.evaluation_periods", "300"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationECSGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.cluster_name", "ecs-cluster-name-update"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_cooldown", "180"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.0.cpu_per_unit", "2048"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.0.memory_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.0.num_of_units", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_down.0.evaluation_periods", "150"),
				),
			},
		},
	})
}

const testIntegrationECSGroupConfig_Create = `
 // --- INTEGRATION: ECS ------------
 integration_ecs = { 
    cluster_name = "ecs-cluster-name"
    autoscale_is_enabled = false
    autoscale_cooldown = 300

    autoscale_headroom = {
      cpu_per_unit = 1024
      memory_per_unit = 512
      num_of_units = 2
    }

    autoscale_down = {
      evaluation_periods = 300
    }
  }
 // --------------------------------
`

const testIntegrationECSGroupConfig_Update = `
 // --- INTEGRATION: ECS ------------
 integration_ecs = { 
    cluster_name = "ecs-cluster-name-update"
    autoscale_is_enabled = true
    autoscale_cooldown = 180

    autoscale_headroom = {
      cpu_per_unit = 2048
      memory_per_unit = 1024
      num_of_units = 1
    }

    autoscale_down = {
      evaluation_periods = 150
    }
  }
 // --------------------------------
`

// endregion

// region Elastigroup: Kubernetes Integration
func TestAccSpotinstElastigroup_IntegrationKubernetes(t *testing.T) {
	groupName := "eg-integration-kubernetes"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationKubernetesGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.integration_mode", "pod"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.cluster_identifier", "k8s-cluster-id"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_is_auto_config", "false"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_cooldown", "300"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_headroom.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_headroom.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_headroom.0.num_of_units", "2"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_down.0.evaluation_periods", "300"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationKubernetesGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.integration_mode", "saas"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.api_server", "k8s-server"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.token", "k8s-token"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_is_auto_config", "true"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_cooldown", "180"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_headroom.0.cpu_per_unit", "2048"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_headroom.0.memory_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_headroom.0.num_of_units", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_down.0.evaluation_periods", "150"),
				),
			},
		},
	})
}

const testIntegrationKubernetesGroupConfig_Create = `
 // --- INTEGRATION: KUBERNETES --------------
 integration_kubernetes = {
    integration_mode = "pod"
    cluster_identifier = "k8s-cluster-id"

    autoscale_is_enabled = false
    autoscale_is_auto_config = false
    autoscale_cooldown = 300

    autoscale_headroom = {
      cpu_per_unit = 1024
      memory_per_unit = 512
      num_of_units = 2
    }

    autoscale_down = {
      evaluation_periods = 300
    }
  }
 // ------------------------------------------
`

const testIntegrationKubernetesGroupConfig_Update = `
 // --- INTEGRATION: KUBERNETES --------------
 integration_kubernetes = {
	integration_mode = "saas"
    api_server = "k8s-server"
    token = "k8s-token"

    autoscale_is_enabled = true
    autoscale_is_auto_config = true
    autoscale_cooldown = 180

    autoscale_headroom = {
      cpu_per_unit = 2048
      memory_per_unit = 1024
      num_of_units = 1
    }

    autoscale_down = {
      evaluation_periods = 150
    }
  }
 // ------------------------------------------
`

// endregion

// region Elastigroup: Nomad Integration
func TestAccSpotinstElastigroup_IntegrationNomad(t *testing.T) {
	groupName := "eg-integration-nomad"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationNomadGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.master_host", "nomad-master-host"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.master_port", "8000"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_cooldown", "300"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_headroom.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_headroom.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_headroom.0.num_of_units", "2"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_down.0.evaluation_periods", "300"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationNomadGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.master_host", "nomad-master-host-update"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.master_port", "9000"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_cooldown", "180"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_headroom.0.cpu_per_unit", "2048"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_headroom.0.memory_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_headroom.0.num_of_units", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_down.0.evaluation_periods", "150"),
				),
			},
		},
	})
}

const testIntegrationNomadGroupConfig_Create = `
 // --- INTEGRATION: NOMAD --------------
 integration_nomad = {
    master_host = "nomad-master-host"
    master_port = 8000

    autoscale_is_enabled = false
    autoscale_cooldown = 300

    autoscale_headroom = {
      cpu_per_unit = 1024
      memory_per_unit = 512
      num_of_units = 2
    }

    autoscale_down = {
      evaluation_periods = 300
    }
  }
 // ------------------------------------------
`

const testIntegrationNomadGroupConfig_Update = `
 // --- INTEGRATION: NOMAD --------------
 integration_nomad = {
	master_host = "nomad-master-host-update"
    master_port = 9000

    autoscale_is_enabled = true
    autoscale_cooldown = 180

    autoscale_headroom = {
      cpu_per_unit = 2048
      memory_per_unit = 1024
      num_of_units = 1
    }

    autoscale_down = {
      evaluation_periods = 150
    }
  }
 // ------------------------------------------
`

// endregion

// region Elastigroup: Mesosphere Integration
func TestAccSpotinstElastigroup_IntegrationMesosphere(t *testing.T) {
	groupName := "eg-integration-mesosphere"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationMesosphereGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_mesosphere.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_mesosphere.0.api_server", "mesosphere-api-server"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationMesosphereGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_mesosphere.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_mesosphere.0.api_server", "mesosphere-api-server-update"),
				),
			},
		},
	})
}

const testIntegrationMesosphereGroupConfig_Create = `
 // --- INTEGRATION: MESOSPHERE --------------
 integration_mesosphere = {
    api_server = "mesosphere-api-server"
 }
 // --------------------------------------------
`

const testIntegrationMesosphereGroupConfig_Update = `
 // --- INTEGRATION: MESOSPHERE --------------
 integration_mesosphere = {
	api_server = "mesosphere-api-server-update"
 }
 // --------------------------------------------
`

// endregion

// region Elastigroup: Multai Runtime Integration
func TestAccSpotinstElastigroup_IntegrationMultaiRuntime(t *testing.T) {
	groupName := "eg-integration-multai-runtime"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationMultaiRuntimeGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_multai_runtime.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_multai_runtime.0.deployment_id", "multai-deployment-id"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationMultaiRuntimeGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_multai_runtime.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_multai_runtime.0.deployment_id", "multai-deployment-id-update"),
				),
			},
		},
	})
}

const testIntegrationMultaiRuntimeGroupConfig_Create = `
 // --- INTEGRATION: MULTAI-RUNTIME ------
 integration_multai_runtime = {
    deployment_id = "multai-deployment-id"
  }
 // --------------------------------------
`

const testIntegrationMultaiRuntimeGroupConfig_Update = `
 // --- INTEGRATION: MULTAI-RUNTIME ------
 integration_multai_runtime = {
    deployment_id = "multai-deployment-id-update"
  }
 // --------------------------------------
`

// endregion

// region Elastigroup: Update Policy
func TestAccSpotinstElastigroup_UpdatePolicy(t *testing.T) {
	groupName := "eg-update-policy"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testUpdatePolicyGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "update_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.should_resume_stateful", "false"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.should_roll", "false"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_size_percentage", "33"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.grace_period", "300"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.health_check_type", "ELB"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testUpdatePolicyGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "update_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.should_resume_stateful", "true"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.should_roll", "true"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_size_percentage", "66"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.grace_period", "600"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.health_check_type", "TARGET_GROUP"),
				),
			},
		},
	})
}

const testUpdatePolicyGroupConfig_Create = `
 // --- UPDATE POLICY -----------------
  update_policy = {
    should_resume_stateful = false
    should_roll = false
    roll_config = {
      batch_size_percentage = 33
      grace_period = 300
      health_check_type = "ELB"
    }
  }
  // ----------------------------------
`

const testUpdatePolicyGroupConfig_Update = `
 // --- UPDATE POLICY -----------------
  update_policy = {
    should_resume_stateful = true
    should_roll = true
    roll_config = {
      batch_size_percentage = 66
      grace_period = 600
      health_check_type = "TARGET_GROUP"
    }
  }
  // ----------------------------------
`

// endregion
