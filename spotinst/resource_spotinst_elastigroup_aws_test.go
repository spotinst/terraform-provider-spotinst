package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws_launch_configuration"
)

func init() {
	resource.AddTestSweepers("spotinst_elastigroup_aws", &resource.Sweeper{
		Name: "spotinst_elastigroup_aws",
		F:    testSweepElastigroupAWS,
	})
}

func testSweepElastigroupAWS(region string) error {
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

func createElastigroupResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.ElastigroupAWSResourceName), name)
}

func testElastigroupDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.ElastigroupAWSResourceName) {
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

type GroupConfigMetadata struct {
	variables            string
	provider             string
	groupName            string
	instanceTypes        string
	launchConfig         string
	strategy             string
	fieldsToAppend       string
	updateBaselineFields bool
	useSubnetIDs         bool
}

func createElastigroupTerraform(gcm *GroupConfigMetadata) string {
	if gcm == nil {
		return ""
	}

	if gcm.provider == "" {
		gcm.provider = "aws"
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

	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`
	if gcm.updateBaselineFields {
		format := testBaselineGroupConfig_Update
		if gcm.useSubnetIDs {
			format = testBaselineSubnetIdsGroupConfig_Update
		}
		template += fmt.Sprintf(format,
			gcm.groupName,
			gcm.provider,
			gcm.groupName,
			gcm.instanceTypes,
			gcm.launchConfig,
			gcm.strategy,
			gcm.fieldsToAppend,
		)
	} else {
		format := testBaselineGroupConfig_Create
		if gcm.useSubnetIDs {
			format = testBaselineSubnetIdsGroupConfig_Create
		}
		template += fmt.Sprintf(format,
			gcm.groupName,
			gcm.provider,
			gcm.groupName,
			gcm.instanceTypes,
			gcm.launchConfig,
			gcm.strategy,
			gcm.fieldsToAppend,
		)
	}

	if gcm.variables != "" {
		template = gcm.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", gcm.groupName, template)
	return template
}

// region Elastigroup: Baseline
func TestAccSpotinstElastigroupAWS_Baseline(t *testing.T) {
	groupName := "test-acc-eg-baseline"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName: groupName,
				}),
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
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:            groupName,
					updateBaselineFields: true,
				}),
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
resource "` + string(commons.ElastigroupAWSResourceName) + `" "%v" {
	provider = "%v"
	
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
resource "` + string(commons.ElastigroupAWSResourceName) + `" "%v" {
	provider = "%v"
	
	name 							= "%v"
	description  			= "created by Terraform"
	product 						= "Linux/UNIX"
	availability_zones = ["us-west-2a"]
	
	// --- CAPACITY ------------
	max_size 		    = 0
	min_size 		    = 0
	desired_capacity = 0
	capacity_unit 	  = "weight"
	// -------------------------
	
	%v
	%v
	%v
	%v
}

`

const testBaselineSubnetIdsGroupConfig_Create = `
resource "` + string(commons.ElastigroupAWSResourceName) + `" "%v" {
	provider = "%v"
	
	name 		   = "%v"
	description = "created by Terraform"
	product 		 = "Linux/UNIX"
	
	// --- SUBNET IDS -------------------
	region      = "us-west-2"
	subnet_ids  = ["subnet-79da021e", "subnet-03b7ed5b"]
	// ----------------------------------
	
	// --- CAPACITY ------------
	max_size 		    = 0
	min_size 		    = 0
	desired_capacity = 0
	capacity_unit 	  = "weight"
	// -------------------------
	
	%v
	%v
	%v
	%v
}

`

const testBaselineSubnetIdsGroupConfig_Update = `
resource "` + string(commons.ElastigroupAWSResourceName) + `" "%v" {
	provider = "%v"
	
	name 			 = "%v"
	description = "created by Terraform"
	product     = "Linux/UNIX"
	
	// --- SUBNET IDS -------------------
	region      = "us-west-2"
	subnet_ids  = ["subnet-79da021e"]
	// ----------------------------------
	
	// --- CAPACITY ------------
	max_size 		    = 0
	min_size 		    = 0
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
func TestAccSpotinstElastigroupAWS_InstanceTypes(t *testing.T) {
	groupName := "test-acc-eg-instance-types"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
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
	instance_types_spot 	  = ["m4.xlarge", "m4.2xlarge"]

	instance_types_weights {
		instance_type = "m4.xlarge"
		weight = 1
	}

	instance_types_weights {
		instance_type = "m4.2xlarge"
		weight = 2
	}
	// ---------------------------------------------------
`

const testInstanceTypesGroupConfig_Update = `
 	// --- INSTANCE TYPES --------------------------------
	instance_types_ondemand = "c4.4xlarge"
	instance_types_spot 	 = ["c4.xlarge", "c4.2xlarge", "c4.4xlarge"]

	instance_types_weights {
		instance_type = "c4.xlarge"
		weight = 3
	}
	// ---------------------------------------------------
`

// endregion

// region Elastigroup: Launch Configuration
func TestAccSpotinstElastigroupAWS_LaunchConfiguration(t *testing.T) {
	groupName := "test-acc-eg-launch-configuration"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
					//resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "iam-profile"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "my-key.ssh"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-123456"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_aws_launch_configuration.Base64StateFunc("echo hello world")),
					resource.TestCheckResourceAttr(resourceName, "shutdown_script", elastigroup_aws_launch_configuration.Base64StateFunc("echo goodbye world")),
					resource.TestCheckResourceAttr(resourceName, "enable_monitoring", "false"),
					resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "false"),
					resource.TestCheckResourceAttr(resourceName, "cpu_credits", "standard"),
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
					//resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "iam-profile updated"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "my-key-updated.ssh"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-123456"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.1", "sg-987654"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_aws_launch_configuration.Base64StateFunc("echo hello world updated")),
					resource.TestCheckResourceAttr(resourceName, "shutdown_script", elastigroup_aws_launch_configuration.Base64StateFunc("echo goodbye world updated")),
					resource.TestCheckResourceAttr(resourceName, "enable_monitoring", "true"),
					resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "true"),
					resource.TestCheckResourceAttr(resourceName, "cpu_credits", "unlimited"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:    groupName,
					launchConfig: testLaunchConfigurationGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-31394949"),
					//resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "iam-profile updated"),
					resource.TestCheckResourceAttr(resourceName, "key_name", "cannot set empty key name"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-123456"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_aws_launch_configuration.Base64StateFunc("cannot set empty user data")),
					resource.TestCheckResourceAttr(resourceName, "shutdown_script", elastigroup_aws_launch_configuration.Base64StateFunc("cannot set empty shutdown script")),
					resource.TestCheckResourceAttr(resourceName, "enable_monitoring", "false"),
					resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "true"),
				),
			},
		},
	})
}

const testLaunchConfigurationGroupConfig_Create = `
 // --- LAUNCH CONFIGURATION --------------
 image_id             = "ami-e251209a"
 //iam_instance_profile = "iam-profile"
 key_name             = "my-key.ssh"
 security_groups      = ["sg-123456"]
 user_data            = "ZWNobyBoZWxsbyB3b3JsZA=="
 shutdown_script      = "echo goodbye world"
 enable_monitoring    = false
 ebs_optimized        = false
 placement_tenancy    = "default"
 cpu_credits          = "standard"
 // ---------------------------------------
`

const testLaunchConfigurationGroupConfig_Update = `
 // --- LAUNCH CONFIGURATION --------------
 image_id             = "ami-31394949"
 //iam_instance_profile = "iam-profile updated"
 key_name             = "my-key-updated.ssh"
 security_groups      = ["sg-123456", "sg-987654"]
 user_data            = "echo hello world updated"
 shutdown_script      = "echo goodbye world updated"
 enable_monitoring    = true
 ebs_optimized        = true
 placement_tenancy    = "default"
 cpu_credits          = "unlimited"
 // ---------------------------------------
`

const testLaunchConfigurationGroupConfig_EmptyFields = `
 // --- LAUNCH CONFIGURATION --------------
 image_id        = "ami-31394949"
 user_data       = "cannot set empty user data"
 shutdown_script = "cannot set empty shutdown script"
 key_name        = "cannot set empty key name"
 security_groups = ["sg-123456"]
 // ---------------------------------------
`

// endregion

// region Elastigroup: Preferred Spot
func TestAccSpotinstElastigroupAWS_PreferredSpot(t *testing.T) {
	groupName := "test-acc-eg-preferred-spot"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					instanceTypes:  testInstanceTypesGroupConfig_Create,
					fieldsToAppend: testPreferredSpotGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preferred_spot.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preferred_spot.0", "m4.xlarge"),
				),
			},
			{
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					instanceTypes:  testInstanceTypesGroupConfig_Update,
					fieldsToAppend: testPreferredSpotGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preferred_spot.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preferred_spot.0", "c4.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preferred_spot.1", "c4.2xlarge"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					instanceTypes:  testInstanceTypesGroupConfig_Update,
					fieldsToAppend: testPreferredSpotGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preferred_spot.#", "0"),
				),
			},
		},
	})
}

const testPreferredSpotGroupConfig_Create = `
 // --- PREFERRED SPOT --------------------------------
 instance_types_preferred_spot = ["m4.xlarge"]
 // ---------------------------------------------------
`

const testPreferredSpotGroupConfig_Update = `
 // --- PREFERRED SPOT --------------------------------
 instance_types_preferred_spot = ["c4.xlarge", "c4.2xlarge"]
 // ---------------------------------------------------
`

const testPreferredSpotGroupConfig_EmptyFields = `
 // --- PREFERRED SPOT --------------------
 // ---------------------------------------
`

// endregion

// region Elastigroup: Strategy
func TestAccSpotinstElastigroupAWS_Strategy(t *testing.T) {
	groupName := "test-acc-eg-strategy"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
					resource.TestCheckResourceAttr(resourceName, "utilize_reserved_instances", "false"),
					resource.TestCheckResourceAttr(resourceName, "scaling_strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_strategy.0.terminate_at_end_of_billing_hour", "true"),
					resource.TestCheckResourceAttr(resourceName, "scaling_strategy.0.termination_policy", "default"),
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
					resource.TestCheckResourceAttr(resourceName, "utilize_reserved_instances", "true"),
					resource.TestCheckResourceAttr(resourceName, "scaling_strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_strategy.0.terminate_at_end_of_billing_hour", "false"),
					resource.TestCheckResourceAttr(resourceName, "scaling_strategy.0.termination_policy", "newestInstance"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName: groupName,
					strategy:  testStrategyGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "orientation", "costOriented"),
					resource.TestCheckResourceAttr(resourceName, "fallback_to_ondemand", "true"),
					resource.TestCheckResourceAttr(resourceName, "ondemand_count", "0"),
					resource.TestCheckResourceAttr(resourceName, "lifetime_period", ""),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "600"),
					resource.TestCheckResourceAttr(resourceName, "utilize_reserved_instances", "false"),
					resource.TestCheckResourceAttr(resourceName, "scaling_strategy.#", "0"),
				),
			},
		},
	})
}

const testStrategyGroupConfig_Create = `
	// --- STRATEGY --------------------
	orientation								 = "balanced"
	fallback_to_ondemand 			 = true
	spot_percentage 		 			 = 100
	lifetime_period	 		 			 = ""
	draining_timeout 		 			 = 300
	utilize_reserved_instances = false

	scaling_strategy {
	 terminate_at_end_of_billing_hour = true
	 termination_policy 						  = "default"
	}
 // ---------------------------------
`

const testStrategyGroupConfig_Update = `
	// --- STRATEGY --------------------
	orientation 							= "costOriented"
	fallback_to_ondemand 		  = false
	ondemand_count 		        = 1
	lifetime_period 		     	= ""
	draining_timeout 					= 600
	utilize_reserved_instances = true
	
	scaling_strategy {
	 terminate_at_end_of_billing_hour = false
	 termination_policy = "newestInstance"
	}
	// ---------------------------------
`

const testStrategyGroupConfig_EmptyFields = `
	// --- STRATEGY ---------------------
	fallback_to_ondemand = true
	orientation 		     = "costOriented"
	draining_timeout 	   = 600
	// ---------------------------------
`

// endregion

// region Elastigroup: Subnet IDs
func TestAccSpotinstElastigroupAWS_SubnetIDs(t *testing.T) {
	groupName := "test-acc-eg-subnet-ids"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:    groupName,
					useSubnetIDs: true,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west-2"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.0", "subnet-79da021e"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.1", "subnet-03b7ed5b"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:            groupName,
					useSubnetIDs:         true,
					updateBaselineFields: true,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "us-west-2"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.0", "subnet-79da021e"),
				),
			},
		},
	})
}

// endregion

// region Elastigroup: Preferred Availability Zones
func TestAccSpotinstElastigroupAWS_PreferredAvailabilityZones(t *testing.T) {
	groupName := "test-acc-eg-preferred-availability-zones"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testPreferredAvailabilityZonesGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "preferred_availability_zones.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "preferred_availability_zones.0", "us-west-2b"),
					resource.TestCheckResourceAttr(resourceName, "preferred_availability_zones.1", "us-west-2c"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testPreferredAvailabilityZonesGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "preferred_availability_zones.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "preferred_availability_zones.0", "us-west-2b"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testPreferredAvailabilityZonesGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "preferred_availability_zones.#", "0"),
				),
			},
		},
	})
}

const testPreferredAvailabilityZonesGroupConfig_Create = `
	// --- PREFERRED AVAILABILITY ZONES -------------------------
	preferred_availability_zones = ["us-west-2b", "us-west-2c"]
	// ----------------------------------------------------------
`

const testPreferredAvailabilityZonesGroupConfig_Update = `
	// --- PREFERRED AVAILABILITY ZONES -------------------------
	preferred_availability_zones = ["us-west-2b"]
	// ----------------------------------------------------------
`

const testPreferredAvailabilityZonesGroupConfig_EmptyFields = `
	// --- PREFERRED AVAILABILITY ZONES -------------------------
	// ----------------------------------------------------------
`

// endregion

// region Elastigroup: Load Balancers
func TestAccSpotinstElastigroupAWS_LoadBalancers(t *testing.T) {
	groupName := "test-acc-eg-load-balancers"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testLoadBalancersGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "elastic_load_balancers.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "target_group_arns.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "multai_target_sets.#", "0"),
				),
			},
		},
	})
}

const testLoadBalancersGroupConfig_Create = `
 // --- LOAD BALANCERS --------------------
 elastic_load_balancers = ["bal1", "bal2"]
 target_group_arns 					= ["arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456"]

	multai_target_sets {
				target_set_id = "ts-123"
				balancer_id = "bal-123"
	}

	multai_target_sets {
			target_set_id = "ts-234"
			balancer_id = "bal-234"
	}
 // ---------------------------------------
`

const testLoadBalancersGroupConfig_Update = `
 // --- LOAD BALANCERS --------------------
 elastic_load_balancers = ["bal1", "bal3"]
 target_group_arns = ["arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456", "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testNewTargetGroup/1234567890123456"]

	multai_target_sets {
		target_set_id = "ts-123"
		balancer_id = "bal-123"
	}

 // ---------------------------------------
`

const testLoadBalancersGroupConfig_EmptyFields = `
 // --- LOAD BALANCERS --------------------
 // ---------------------------------------
`

// endregion

// region Elastigroup: Health Checks
func TestAccSpotinstElastigroupAWS_HealthChecks(t *testing.T) {
	groupName := "test-acc-eg-health-checks"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
					//resource.TestCheckResourceAttr(resourceName, "health_check_type", "ELB"),
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
					//resource.TestCheckResourceAttr(resourceName, "health_check_type", "TARGET_GROUP"),
					resource.TestCheckResourceAttr(resourceName, "health_check_grace_period", "50"),
					resource.TestCheckResourceAttr(resourceName, "health_check_unhealthy_duration_before_replacement", "120"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testHealthChecksGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "health_check_type", ""),
					resource.TestCheckResourceAttr(resourceName, "health_check_grace_period", "0"),
					resource.TestCheckResourceAttr(resourceName, "health_check_unhealthy_duration_before_replacement", "0"),
				),
			},
		},
	})
}

const testHealthChecksGroupConfig_Create = `
 // --- HEALTH-CHECKS ------------------------------------
 //health_check_type = "ELB" 
 health_check_grace_period = 100
 health_check_unhealthy_duration_before_replacement = 60
 // ------------------------------------------------------
`

const testHealthChecksGroupConfig_Update = `
 // --- HEALTH-CHECKS ------------------------------------
 //health_check_type = "TARGET_GROUP" 
 health_check_grace_period = 50
 health_check_unhealthy_duration_before_replacement = 120
 // ------------------------------------------------------
`

const testHealthChecksGroupConfig_EmptyFields = `
 // --- HEALTH-CHECKS ------------------------------------
 // ------------------------------------------------------
`

// endregion

// region Elastigroup: Elastic IPs
func TestAccSpotinstElastigroupAWS_ElasticIPs(t *testing.T) {
	groupName := "test-acc-eg-elastic-ips"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testElasticIPsGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "elastic_ips.#", "0"),
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

const testElasticIPsGroupConfig_EmptyFields = `
	// --- ELASTIC IPs --------------------------------------
	// ------------------------------------------------------
`

// endregion

// region Elastigroup: Elastic IPs with Terraform Count Parallelism
func TestAccSpotinstElastigroupAWS_ElasticIPs_Count_Parallelism(t *testing.T) {
	groupName := "test-acc-eg-elastic-ips-tf-count-parallelism"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupDestroy,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					variables:      testElasticIPsGroupConfig_Count_Variables,
					groupName:      groupName,
					fieldsToAppend: testElasticIPsGroupConfig_Count_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName+".0"),
					testCheckElastigroupExists(&group, resourceName+".1"),
					testCheckElastigroupExists(&group, resourceName+".2"),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName+".0", "elastic_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName+".0", "elastic_ips.0", "eipalloc-123"),
					resource.TestCheckResourceAttr(resourceName+".1", "elastic_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName+".1", "elastic_ips.0", "eipalloc-456"),
					resource.TestCheckResourceAttr(resourceName+".2", "elastic_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName+".2", "elastic_ips.0", "eipalloc-789"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					variables:      testElasticIPsGroupConfig_Count_Variables,
					groupName:      groupName,
					fieldsToAppend: testElasticIPsGroupConfig_Count_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName+".0"),
					testCheckElastigroupExists(&group, resourceName+".1"),
					testCheckElastigroupExists(&group, resourceName+".2"),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName+".0", "elastic_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName+".0", "elastic_ips.0", "eipalloc-111"),
					resource.TestCheckResourceAttr(resourceName+".1", "elastic_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName+".1", "elastic_ips.0", "eipalloc-444"),
					resource.TestCheckResourceAttr(resourceName+".2", "elastic_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName+".2", "elastic_ips.0", "eipalloc-777"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					variables:      testElasticIPsGroupConfig_Count_Variables,
					groupName:      groupName,
					fieldsToAppend: testElasticIPsGroupConfig_Count_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName+".0"),
					testCheckElastigroupExists(&group, resourceName+".1"),
					testCheckElastigroupExists(&group, resourceName+".2"),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName+".0", "elastic_ips.#", "0"),
					resource.TestCheckResourceAttr(resourceName+".1", "elastic_ips.#", "0"),
					resource.TestCheckResourceAttr(resourceName+".2", "elastic_ips.#", "0"),
				),
			},
		},
	})
}

const testElasticIPsGroupConfig_Count_Variables = `
// --- VARIABLES --------------------------------------------
variable "elastic_ips" {
  description = "List with the Elastic IPs for elastigroups"
  type        = "list"
  default     = ["eipalloc-123", "eipalloc-456", "eipalloc-789"]
}

variable "elastic_ips_update" {
  description = "List with the Elastic IPs for elastigroups"
  type        = "list"
  default     = ["eipalloc-111", "eipalloc-444", "eipalloc-777"]
}
// ----------------------------------------------------------
`

const testElasticIPsGroupConfig_Count_Create = `
 // --- ELASTIC IPs --------------------------------------
 count = 3
 elastic_ips = ["${element(var.elastic_ips, count.index)}"]
 // ------------------------------------------------------
`

const testElasticIPsGroupConfig_Count_Update = `
 // --- ELASTIC IPs --------------------------------------
 count = 3
 elastic_ips = ["${element(var.elastic_ips_update, count.index)}"]
 // ------------------------------------------------------
`

const testElasticIPsGroupConfig_Count_EmptyFields = `
 // --- ELASTIC IPs --------------------------------------
 count = 3
 // ------------------------------------------------------
`

// endregion

// region Elastigroup: Signals
func TestAccSpotinstElastigroupAWS_Signals(t *testing.T) {
	groupName := "test-acc-eg-signals"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
					resource.TestCheckResourceAttr(resourceName, "signal.1735739494.timeout", "200"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testSignalsGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "signal.#", "0"),
				),
			},
		},
	})
}

const testSignalsGroupConfig_Create = `
 // --- SIGNAL -----
  signal {
    name = "INSTANCE_READY_TO_SHUTDOWN"
    timeout = 100
  }
 // ----------------
`

const testSignalsGroupConfig_Update = `
 // --- SIGNAL -----
  signal {
    name = "INSTANCE_READY_TO_SHUTDOWN"
    timeout = 100
  }
  signal {
    name = "INSTANCE_READY"
    timeout = 200
  }
 // ----------------
`

const testSignalsGroupConfig_EmptyFields = `
 // --- SIGNAL -----
 // ----------------
`

// endregion

// region Elastigroup: Revert To Spot (Maintenance Window)
func TestAccSpotinstElastigroupAWS_RevertToSpot(t *testing.T) {
	groupName := "test-acc-eg-revert-to-spot"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testRevertToSpotGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "revert_to_spot.#", "0"),
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

const testRevertToSpotGroupConfig_EmptyFields = `
 // --- REVERT TO SPOT -------------------------------------------
 // -------------------------------------------------------------
`

// endregion

// region Elastigroup: Network Interfaces
func TestAccSpotinstElastigroupAWS_NetworkInterfaces(t *testing.T) {
	groupName := "test-acc-eg-network-interfaces"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
					resource.TestCheckResourceAttr(resourceName, "network_interface.4092817214.associate_public_ip_address", "false"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.4092817214.associate_ipv6_address", "false"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.4092817214.delete_on_termination", "false"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.4092817214.description", "network interface description"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.4092817214.device_index", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.4092817214.network_interface_id", "n-123456"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.4092817214.private_ip_address", "1.1.1.1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.4092817214.secondary_private_ip_address_count", "1"),
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
					resource.TestCheckResourceAttr(resourceName, "network_interface.3760037182.associate_public_ip_address", "true"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.3760037182.associate_ipv6_address", "true"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.3760037182.delete_on_termination", "true"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.3760037182.description", "network interface description updated"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.3760037182.device_index", "2"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.3760037182.private_ip_address", "2.2.2.2"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.3760037182.secondary_private_ip_address_count", "2"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationNetworkInterfacesGroupConfig_ShouldFail,
				}),
				ExpectError: regexp.MustCompile("invalid Network interface: associate_public_ip_address must be undefined when using network_interface_id"),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testNetworkInterfacesGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "network_interface.#", "0"),
				),
			},
		},
	})
}

const testNetworkInterfacesGroupConfig_Create = `
 // --- NETWORK INTERFACE ------------------
 network_interface { 
    description = "network interface description"
    device_index = 1
    secondary_private_ip_address_count = 1
    associate_public_ip_address = false
    associate_ipv6_address = false
    delete_on_termination = false
    network_interface_id = "n-123456"
    private_ip_address = "1.1.1.1"
  }
 // ----------------------------------------
`

const testNetworkInterfacesGroupConfig_Update = `
 // --- NETWORK INTERFACE ------------------
 network_interface { 
    description = "network interface description updated"
    device_index = 2
    secondary_private_ip_address_count = 2
    associate_public_ip_address = true
    associate_ipv6_address = true
    delete_on_termination = true
    private_ip_address = "2.2.2.2"
  }
 // ----------------------------------------
`

const testIntegrationNetworkInterfacesGroupConfig_ShouldFail = `
 // --- NETWORK INTERFACE ------------------
 network_interface { 
    description = "network interface description updated"
    device_index = 2
    secondary_private_ip_address_count = 2
    associate_public_ip_address = true
    network_interface_id = "n-123456"
    delete_on_termination = true
    private_ip_address = "2.2.2.2"
  }
 // ----------------------------------------
`

const testNetworkInterfacesGroupConfig_EmptyFields = `
 // --- NETWORK INTERFACE ------------------
 // ----------------------------------------
`

// endregion

// region Elastigroup: Scaling Up Policies
func TestAccSpotinstElastigroupAWS_ScalingUpPolicies(t *testing.T) {
	groupName := "test-acc-eg-scaling-up-policy"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingUpPolicyGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.policy_name", "policy-name"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.metric_name", "CPUUtilization"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.namespace", "AWS/EC2"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.source", "cloudWatch"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.statistic", "average"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.unit", "percent"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.cooldown", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.dimensions.0.name", "name-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.dimensions.0.value", "value-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.threshold", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.operator", "gt"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.evaluation_periods", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.period", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.action_type", "setMinTarget"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.min_target_capacity", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.adjustment", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.max_target_capacity", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.maximum", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.minimum", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.190076102.target", ""),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingUpPolicyGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.policy_name", "policy-name-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.metric_name", "CPUUtilization"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.namespace", "AWS/EC2"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.source", "spectrum"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.statistic", "sum"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.unit", "bytes"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.cooldown", "120"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.dimensions.0.name", "name-1-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.dimensions.0.value", "value-1-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.threshold", "5"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.operator", "lt"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.evaluation_periods", "5"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.period", "120"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.action_type", "adjustment"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.adjustment", "MAX(5,10)"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.min_target_capacity", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.max_target_capacity", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.maximum", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.minimum", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.3752994292.target", ""),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingUpPolicyGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.#", "0"),
				),
			},
		},
	})
}

const testScalingUpPolicyGroupConfig_Create = `
 // --- SCALE UP POLICY ------------------
 scaling_up_policy {
  policy_name = "policy-name"
  metric_name = "CPUUtilization"
  namespace = "AWS/EC2"
  source = "cloudWatch"
  statistic = "average"
  unit = "percent"
  cooldown = 60
  is_enabled = false
  dimensions {
      name = "name-1"
      value = "value-1"
  }
  threshold = 10

  operator = "gt"
  evaluation_periods = "10"
  period = "60"

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

  }
 // ----------------------------------------
`

const testScalingUpPolicyGroupConfig_Update = `
 // --- SCALE UP POLICY ------------------
 scaling_up_policy {
  policy_name = "policy-name-update"
  metric_name = "CPUUtilization"
  namespace = "AWS/EC2"
  source = "spectrum"
  statistic = "sum"
  unit = "bytes"
  is_enabled = true
  cooldown = 120
  dimensions {
      name = "name-1-update"
      value = "value-1-update"
  }
  threshold = 5

  operator = "lt"
  evaluation_periods = 5
  period = 120

  //// === MIN TARGET ===================
  # action_type = "setMinTarget"
  # min_target_capacity = 1
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

  }
 // ----------------------------------------
`

const testScalingUpPolicyGroupConfig_EmptyFields = `
 // --- SCALE UP POLICY ------------------
 // ----------------------------------------
`

// endregion

// region Elastigroup: Scaling Down Policies
func TestAccSpotinstElastigroupAWS_ScalingDownPolicies(t *testing.T) {
	groupName := "test-acc-eg-scaling-down-policy"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingDownPolicyGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.policy_name", "policy-name"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.metric_name", "CPUUtilization"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.namespace", "AWS/EC2"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.source", "cloudWatch"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.statistic", "average"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.unit", "percent"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.cooldown", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.dimensions.0.name", "name-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.dimensions.0.value", "value-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.threshold", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.operator", "lt"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.evaluation_periods", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.period", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.action_type", "adjustment"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.adjustment", "MIN(5,10)"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.max_target_capacity", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.min_target_capacity", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.maximum", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.minimum", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.2789098273.target", ""),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingDownPolicyGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.policy_name", "policy-name-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.metric_name", "CPUUtilization"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.namespace", "AWS/EC2"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.source", "spectrum"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.statistic", "sum"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.unit", "bytes"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.cooldown", "120"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.dimensions.0.name", "name-1-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.dimensions.0.value", "value-1-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.threshold", "5"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.operator", "lt"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.evaluation_periods", "5"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.period", "120"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.action_type", "updateCapacity"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.minimum", "0"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.maximum", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.target", "5"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.max_target_capacity", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.3634262214.min_target_capacity", ""),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingDownPolicyGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.#", "0"),
				),
			},
		},
	})
}

const testScalingDownPolicyGroupConfig_Create = `
 // --- SCALE DOWN POLICY ------------------
 scaling_down_policy {
  policy_name = "policy-name"
  metric_name = "CPUUtilization"
  namespace = "AWS/EC2"
  source = "cloudWatch"
  statistic = "average"
  unit = "percent"
  is_enabled = false
  cooldown = 60
  dimensions {
      name = "name-1"
      value = "value-1"
  }
  threshold = 10

  operator = "lt"
  evaluation_periods = 10
  period = 60

  // === MIN TARGET ===================
  # action_type = "setMinTarget"
  # min_target_capacity = 1
  // ==================================

  // === ADJUSTMENT ===================
  # action_type = "percentageAdjustment"
  action_type = "adjustment"
  adjustment = "MIN(5,10)"
  // ==================================

  // === UPDATE CAPACITY ==============
  # action_type = "updateCapacity"
  # minimum = 0
  # maximum = 10
  # target = 5
  // ==================================

  }
 // ----------------------------------------
`

const testScalingDownPolicyGroupConfig_Update = `
 // --- SCALE DOWN POLICY ------------------
 scaling_down_policy {
  policy_name = "policy-name-update"
  metric_name = "CPUUtilization"
  namespace = "AWS/EC2"
  source = "spectrum"
  statistic = "sum"
  unit = "bytes"
  is_enabled = true
  cooldown = 120
  dimensions {
      name = "name-1-update"
      value = "value-1-update"
  }
  threshold = 5

  operator = "lt"
  evaluation_periods = 5
  period = 120

  // === MIN TARGET ===================
  # action_type = "setMinTarget"
  # min_target_capacity = 1
  // ==================================

  // === ADJUSTMENT ===================
  # action_type = "percentageAdjustment"
  # action_type = "adjustment"
  # adjustment = "MAX(5,10)"
  // ==================================

  // === UPDATE CAPACITY ==============
  action_type = "updateCapacity"
  minimum = 0
  maximum = 10
  target = 5
  // ==================================

  }
 // ----------------------------------------
`

const testScalingDownPolicyGroupConfig_EmptyFields = `
 // --- SCALE DOWN POLICY ------------------
 // ----------------------------------------
`

// endregion

// region Elastigroup: Scaling Target Policies
func TestAccSpotinstElastigroupAWS_ScalingTargetPolicies(t *testing.T) {
	groupName := "test-acc-eg-scaling-target-policy"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingTargetPolicyGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.2949818563.policy_name", "policy-name"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.2949818563.metric_name", "CPUUtilization"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.2949818563.namespace", "AWS/EC2"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.2949818563.source", "cloudWatch"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.2949818563.statistic", "average"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.2949818563.unit", "percent"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.2949818563.cooldown", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.2949818563.predictive_mode", "FORECAST_AND_SCALE"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.2949818563.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.2949818563.dimensions.0.name", "name-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.2949818563.dimensions.0.value", "value-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.2949818563.target", "1.1"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingTargetPolicyGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.948979961.policy_name", "policy-name-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.948979961.metric_name", "CPUUtilization"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.948979961.namespace", "AWS/EC2"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.948979961.source", "spectrum"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.948979961.statistic", "sum"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.948979961.unit", "bytes"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.948979961.cooldown", "120"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.948979961.predictive_mode", "FORECAST_ONLY"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.948979961.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.948979961.dimensions.0.name", "name-1-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.948979961.dimensions.0.value", "value-1-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.948979961.target", "2.2"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingTargetPolicyGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_target_policy.#", "0"),
				),
			},
		},
	})
}

const testScalingTargetPolicyGroupConfig_Create = `
 // --- SCALE TARGET POLICY ----------------
 scaling_target_policy {
  policy_name = "policy-name"
  metric_name = "CPUUtilization"
  namespace = "AWS/EC2"
  source = "cloudWatch"
  statistic = "average"
  unit = "percent"
  cooldown = 60
  predictive_mode = "FORECAST_AND_SCALE"
  dimensions {
      name = "name-1"
      value = "value-1"
  }
  target=1.1
  }
 // ----------------------------------------
`

const testScalingTargetPolicyGroupConfig_Update = `
 // --- SCALE TARGET POLICY ----------------
 scaling_target_policy {
  policy_name = "policy-name-update"
  metric_name = "CPUUtilization"
  namespace = "AWS/EC2"
  source = "spectrum"
  statistic = "sum"
  unit = "bytes"
  cooldown = 120
  predictive_mode = "FORECAST_ONLY"
  dimensions {
      name = "name-1-update"
      value = "value-1-update"
  }
  target=2.2
  }
 // ----------------------------------------
`

const testScalingTargetPolicyGroupConfig_EmptyFields = `
 // --- SCALE TARGET POLICY ----------------
 // ----------------------------------------
`

// endregion

// region Elastigroup: Scheduled Tasks
func TestAccSpotinstElastigroupAWS_ScheduledTask(t *testing.T) {
	groupName := "test-acc-eg-scheduled-task"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScheduledTaskGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
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
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScheduledTaskGroupConfig_Update2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
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
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScheduledTaskGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "0"),
				),
			},
		},
	})
}

const testScheduledTaskGroupConfig_Create = `
 // --- SCHEDULED TASK ------------------
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
 // -------------------------------------
`

const testScheduledTaskGroupConfig_Update = `
 // --- SCHEDULED TASK ------------------
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
 // -------------------------------------
`

const testScheduledTaskGroupConfig_Update2 = `
 // --- SCHEDULED TASK ------------------
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
 // -------------------------------------
`

const testScheduledTaskGroupConfig_EmptyFields = `
 // --- SCHEDULED TASK ------------------
 // -------------------------------------
`

// endregion

// region Elastigroup: Stateful
func TestAccSpotinstElastigroupAWS_Stateful(t *testing.T) {
	groupName := "test-acc-eg-stateful"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testStatefulGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "persist_root_device", "false"),
					resource.TestCheckResourceAttr(resourceName, "persist_block_devices", "false"),
					resource.TestCheckResourceAttr(resourceName, "persist_private_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "block_devices_mode", ""),
					resource.TestCheckResourceAttr(resourceName, "private_ips.#", "0"),
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

const testStatefulGroupConfig_EmptyFields = `
 // --- STATEFUL ----------------------
 // -----------------------------------
`

func TestAccSpotinstElastigroupAWS_DeallocationStateful_DeleteNetworkInterfacesAndSnapshots(t *testing.T) {
	groupName := "test-acc-eg-stateful-deallocation-network-interfaces-snapshots"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testDeallocationStatefulGroupConfig_DeleteNetworkInterfacesAndSnapshots,
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
					resource.TestCheckResourceAttr(resourceName, "stateful_deallocation.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "stateful_deallocation.0.should_delete_images", "false"),
					resource.TestCheckResourceAttr(resourceName, "stateful_deallocation.0.should_delete_network_interfaces", "true"),
					resource.TestCheckResourceAttr(resourceName, "stateful_deallocation.0.should_delete_volumes", "false"),
					resource.TestCheckResourceAttr(resourceName, "stateful_deallocation.0.should_delete_snapshots", "true"),
				),
			},
		},
	})
}

const testDeallocationStatefulGroupConfig_DeleteNetworkInterfacesAndSnapshots = `
 // --- DEALLOCATION STATEFUL ---------
 persist_root_device = true
 persist_block_devices = true
 persist_private_ip = true
 block_devices_mode = "reattach"
 private_ips = ["1.1.1.1", "2.2.2.2"]
 stateful_deallocation {
   should_delete_images              = false
   should_delete_network_interfaces  = true
   should_delete_volumes             = false
   should_delete_snapshots           = true
 }
 // -----------------------------------
`

func TestAccSpotinstElastigroupAWS_DeallocationStateful_DeleteVolumesAndImages(t *testing.T) {
	groupName := "test-acc-eg-stateful-deallocation-volumes-images"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testDeallocationStatefulGroupConfig_DeleteVolumesAndImages,
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
					resource.TestCheckResourceAttr(resourceName, "stateful_deallocation.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "stateful_deallocation.0.should_delete_images", "true"),
					resource.TestCheckResourceAttr(resourceName, "stateful_deallocation.0.should_delete_network_interfaces", "false"),
					resource.TestCheckResourceAttr(resourceName, "stateful_deallocation.0.should_delete_volumes", "true"),
					resource.TestCheckResourceAttr(resourceName, "stateful_deallocation.0.should_delete_snapshots", "false"),
				),
			},
		},
	})
}

const testDeallocationStatefulGroupConfig_DeleteVolumesAndImages = `
 // --- DEALLOCATION STATEFUL ---------
 persist_root_device = true
 persist_block_devices = true
 persist_private_ip = true
 block_devices_mode = "reattach"
 private_ips = ["1.1.1.1", "2.2.2.2"]
 stateful_deallocation {
   should_delete_images              = true
   should_delete_network_interfaces  = false
   should_delete_volumes             = true
   should_delete_snapshots           = false
 }
 // -----------------------------------
`

func TestAccSpotinstElastigroupAWS_DeallocationStateful_DeleteWithoutStatefulResources(t *testing.T) {
	groupName := "test-acc-eg-stateful-deallocation-empty"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testDeallocationStatefulGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "persist_root_device", "false"),
					resource.TestCheckResourceAttr(resourceName, "persist_block_devices", "false"),
					resource.TestCheckResourceAttr(resourceName, "persist_private_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "block_devices_mode", ""),
				),
			},
		},
	})
}

const testDeallocationStatefulGroupConfig_EmptyFields = `
 // --- DEALLOCATION STATEFUL ---------
 // -----------------------------------
`

// endregion

// region Elastigroup: Block Devices
func TestAccSpotinstElastigroupAWS_BlockDevices(t *testing.T) {
	groupName := "test-acc-eg-block-devices"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
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
			{
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testElastigroupBlockDevices_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "ebs_block_device.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "ephemeral_block_device.#", "0"),
				),
			},
		},
	})
}

const testElastigroupBlockDevices_Create = `
	// --- EBS BLOCK DEVICE -----------------
	ebs_block_device {
	 device_name 			     = "/dev/sdb"
	 snapshot_id 				   = ""
	 volume_type 				   = "gp2"
	 volume_size 				   = 12
	 iops 					       = 1
	 delete_on_termination = true
	 encrypted             = false
	}
	ebs_block_device {
	 device_name 			     = "/dev/sda"
	 snapshot_id 				   = ""
	 volume_type 				   = "io1"
	 volume_size 				   = 8
	 iops 					       = 1
	 delete_on_termination = false
	 encrypted 				     = false
	}
	// --------------------------------------

	// --- EPHEMERAL BLOCK DEVICE ----
	ephemeral_block_device {
		device_name  = "/dev/xvdc"
		virtual_name = "ephemeral0"
	}
	// -------------------------------
`

const testElastigroupBlockDevices_Update = `
 // --- EBS BLOCK DEVICE -----------------
 ebs_block_device {
   device_name 				   = "/dev/sda"
   snapshot_id 				   = ""
   volume_type 				   = "sc1"
   volume_size 				   = 10
   iops 					       = 1
   delete_on_termination = true
   encrypted 				     = true
   kms_key_id 				   = "acceptance-kms-key"
 }
 // --------------------------------------

 // --- EPHEMERAL BLOCK DEVICE ----
 ephemeral_block_device {
  device_name  = "/dev/xvdc"
  virtual_name = "ephemeral1"
 }
 // -------------------------------
`

const testElastigroupBlockDevices_EmptyFields = `
 // --- EBS BLOCK DEVICE -----------------
 // --------------------------------------

 // --- EPHEMERAL BLOCK DEVICE ----
 // -------------------------------
`

// endregion

// region Elastigroup: Tags
func TestAccSpotinstElastigroupAWS_Tags(t *testing.T) {
	groupName := "test-acc-eg-tags"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testTagsGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "0"),
				),
			},
		},
	})
}

const testTagsGroupConfig_Create = `
 // --- TAGS ---------
  tags {
     key = "explicit1"
     value = "value1"
	}

	tags {
     key = "explicit2"
     value = "value2"
   }
 // ------------------
`

const testTagsGroupConfig_Update = `
 // --- TAGS ---------
  tags {
     key = "explicit1-update"
     value = "value1-update"
   }
 // ------------------
`

const testTagsGroupConfig_EmptyFields = `
 // --- TAGS ---------
 // ------------------
`

// endregion

// region Elastigroup: Rancher Integration
func TestAccSpotinstElastigroupAWS_IntegrationRancher(t *testing.T) {
	groupName := "test-acc-eg-integration-rancher"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.0.master_host", "master-host"),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.0.access_key", "access-key"),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.0.secret_key", "secret-key"),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.0.version", "1"),
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
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.0.master_host", "master-host-update"),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.0.access_key", "access-key-update"),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.0.secret_key", "secret-key-update"),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.0.version", "2"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationRancherGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_rancher.#", "0"),
				),
			},
		},
	})
}

const testIntegrationRancherGroupConfig_Create = `
// --- INTEGRATION: RANCHER ---
integration_rancher {
   master_host = "master-host"
   access_key = "access-key"
   secret_key = "secret-key"
   version = "1"
}
// ----------------------------
`

const testIntegrationRancherGroupConfig_Update = `
 // --- INTEGRATION: RANCHER ---
 integration_rancher {
    master_host = "master-host-update"
    access_key = "access-key-update"
    secret_key = "secret-key-update"
    version = "2"
 }

integration_kubernetes {
	autoscale_is_enabled = true
 }
 // ----------------------------
`

const testIntegrationRancherGroupConfig_EmptyFields = `
 // --- INTEGRATION: RANCHER ---
 // ----------------------------
`

// endregion

// region Elastigroup: Code Deploy Integration
func TestAccSpotinstElastigroupAWS_IntegrationCodeDeploy(t *testing.T) {
	groupName := "test-acc-eg-integration-code-deploy"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationCodeDeployGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_codedeploy.#", "0"),
				),
			},
		},
	})
}

const testIntegrationCodeDeployGroupConfig_Create = `
 // --- INTEGRATION: CODE-DEPLOY ----------
 integration_codedeploy {
    cleanup_on_failure = false
    terminate_instance_on_failure = false

    deployment_groups {
      application_name = "code-deploy-application"
      deployment_group_name = "code-deploy-deployment"
    }
  }
 // ---------------------------------------
`

const testIntegrationCodeDeployGroupConfig_Update = `
 // --- INTEGRATION: CODE-DEPLOY ----------
 integration_codedeploy {
    cleanup_on_failure = true
    terminate_instance_on_failure = true

    deployment_groups {
      application_name = "code-deploy-application-update"
      deployment_group_name = "code-deploy-deployment-update"
    }
  }
 // ---------------------------------------
`

const testIntegrationCodeDeployGroupConfig_EmptyFields = `
 // --- INTEGRATION: CODE-DEPLOY ----------
 // ---------------------------------------
`

// endregion

// region Elastigroup: Route53 integration
func TestAccSpotinstElastigroupAWS_IntegrationRoute53(t *testing.T) {
	groupName := "test-acc-eg-integration-route53"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationRoute53GroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
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
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationRoute53GroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
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
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationRoute53GroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_route53.#", "0"),
				),
			},
		},
	})
}

const testIntegrationRoute53GroupConfig_Create = `
// --- INTEGRATION: ROUTE53 ----------
integration_route53 {
	domains {
		hosted_zone_id   = "id_create"
    spotinst_acct_id = "act-123456"
		record_sets {
			name = "test_create"
			use_public_ip = false
		}
	}
}
// ------------------------------------
`
const testIntegrationRoute53GroupConfig_Update = `
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
const testIntegrationRoute53GroupConfig_EmptyFields = `
// --- INTEGRATION: ROUTE53 ----------
// ------------------------------------
`

// endregion

// region Elastigroup: ECS Integration
func TestAccSpotinstElastigroupAWS_IntegrationECS(t *testing.T) {
	groupName := "test-acc-eg-integration-ecs"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_is_auto_config", "false"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_cooldown", "300"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.0.num_of_units", "2"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_down.0.evaluation_periods", "300"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_down.0.max_scale_down_percentage", "70"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_scale_down_non_service_tasks", "false"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_attributes.2966515502.key", "test.key.ecs"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_attributes.2966515502.value", "test.value.ecs"),
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
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_is_auto_config", "true"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_cooldown", "180"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.0.cpu_per_unit", "2048"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.0.memory_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_headroom.0.num_of_units", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_down.0.evaluation_periods", "150"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_down.0.max_scale_down_percentage", "50"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_scale_down_non_service_tasks", "true"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_attributes.2266469793.key", "test.key.ecs.update"),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.0.autoscale_attributes.2266469793.value", "test.value.ecs.update"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationECSGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_ecs.#", "0"),
				),
			},
		},
	})
}

const testIntegrationECSGroupConfig_Create = `
 // --- INTEGRATION: ECS -----------
 integration_ecs { 
    cluster_name = "ecs-cluster-name"
    autoscale_is_enabled = false
		autoscale_is_auto_config = false
    autoscale_cooldown = 300
    autoscale_scale_down_non_service_tasks = false

    autoscale_headroom {
      cpu_per_unit = 1024
      memory_per_unit = 512
      num_of_units = 2
    }

    autoscale_down {
      evaluation_periods = 300
      max_scale_down_percentage = 70
    }

    autoscale_attributes {
      key   = "test.key.ecs"
      value = "test.value.ecs"
    }
  }
 // --------------------------------
`

const testIntegrationECSGroupConfig_Update = `
 // --- INTEGRATION: ECS -----------
 integration_ecs { 
    cluster_name = "ecs-cluster-name-update"
    autoscale_is_enabled = true
	  autoscale_is_auto_config = true
    autoscale_cooldown = 180
    autoscale_scale_down_non_service_tasks = true

    autoscale_headroom {
      cpu_per_unit = 2048
      memory_per_unit = 1024
      num_of_units = 1
    }

    autoscale_down {
      evaluation_periods = 150
      max_scale_down_percentage = 50
    }

    autoscale_attributes {
      key   = "test.key.ecs.update"
      value = "test.value.ecs.update"
    }
  }
 // --------------------------------
`

const testIntegrationECSGroupConfig_EmptyFields = `
 // --- INTEGRATION: ECS -----------
 // --------------------------------
`

// endregion

// region Elastigroup: Kubernetes Integration
func TestAccSpotinstElastigroupAWS_IntegrationKubernetes(t *testing.T) {
	groupName := "test-acc-eg-integration-kubernetes"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_labels.3445854412.key", "test.key.k8s"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_labels.3445854412.value", "test.value.k8s"),
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
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_labels.208193869.key", "test.key.k8s.update"),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.0.autoscale_labels.208193869.value", "test.value.k8s.update"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationKubernetesGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_kubernetes.#", "0"),
				),
			},
		},
	})
}

const testIntegrationKubernetesGroupConfig_Create = `
 // --- INTEGRATION: KUBERNETES --------------
 integration_kubernetes {
    integration_mode = "pod"
    cluster_identifier = "k8s-cluster-id"

    autoscale_is_enabled = false
    autoscale_is_auto_config = false
    autoscale_cooldown = 300

    autoscale_headroom {
      cpu_per_unit = 1024
      memory_per_unit = 512
      num_of_units = 2
    }

    autoscale_down {
      evaluation_periods = 300
    }

    autoscale_labels {
      key   = "test.key.k8s"
      value = "test.value.k8s"
    }
  }
 // ------------------------------------------
`

const testIntegrationKubernetesGroupConfig_Update = `
 // --- INTEGRATION: KUBERNETES --------------
 integration_kubernetes {
	integration_mode = "saas"
    api_server = "k8s-server"
    token = "k8s-token"

    autoscale_is_enabled = true
    autoscale_is_auto_config = true
    autoscale_cooldown = 180

    autoscale_headroom {
      cpu_per_unit = 2048
      memory_per_unit = 1024
      num_of_units = 1
    }

    autoscale_down {
      evaluation_periods = 150
    }

    autoscale_labels {
      key   = "test.key.k8s.update"
      value = "test.value.k8s.update"
    }
  }
 // ------------------------------------------
`

const testIntegrationKubernetesGroupConfig_EmptyFields = `
 // --- INTEGRATION: KUBERNETES --------------
 // ------------------------------------------
`

// endregion

// region Elasitgroup: Beanstalk Integration
func TestAccSpotinstElastigroupAWS_IntegrationBeanstalk(t *testing.T) {
	groupName := "test-acc-eg-integration-beanstalk"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationBeanstalkGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.environment_id", "e-h3kze2cv9e"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.deployment_preferences.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.deployment_preferences.0.automatic_roll", "true"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.deployment_preferences.0.batch_size_percentage", "100"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.deployment_preferences.0.grace_period", "90"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.deployment_preferences.0.strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.deployment_preferences.0.strategy.0.action", "REPLACE_SERVER"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.deployment_preferences.0.strategy.0.should_drain_instances", "true"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationBeanstalkGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.environment_id", "e-h3kze2cv9e"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.deployment_preferences.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.deployment_preferences.0.automatic_roll", "false"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.deployment_preferences.0.batch_size_percentage", "10"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.deployment_preferences.0.grace_period", "900"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.deployment_preferences.0.strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.deployment_preferences.0.strategy.0.action", "REPLACE_SERVER"),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.0.deployment_preferences.0.strategy.0.should_drain_instances", "false"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationBeanstalkGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_beanstalk.#", "0"),
				),
			},
		},
	})
}

const testIntegrationBeanstalkGroupConfig_Create = `
 // --- INTEGRATION: BEANSTALK  --------------
 integration_beanstalk {
  environment_id = "e-h3kze2cv9e"

  deployment_preferences {
    automatic_roll        = true
    batch_size_percentage = 100
    grace_period          = 90

    strategy {
      action                = "REPLACE_SERVER"
      should_drain_instances = true
    }
  }
 }
 // ------------------------------------------
`

const testIntegrationBeanstalkGroupConfig_Update = `
 // --- INTEGRATION: BEANSTALK  --------------
 integration_beanstalk {
  environment_id = "e-h3kze2cv9e"

  deployment_preferences {
    automatic_roll        = false
    batch_size_percentage = 10
    grace_period          = 900

    strategy {
      action                 = "REPLACE_SERVER"
      should_drain_instances = false
    }
  }
 }
 // ------------------------------------------
`

const testIntegrationBeanstalkGroupConfig_EmptyFields = `
 // --- INTEGRATION: BEANSTALK  --------------
 // ------------------------------------------
`

// endregion

// region Elastigroup: Nomad Integration
func TestAccSpotinstElastigroupAWS_IntegrationNomad(t *testing.T) {
	groupName := "test-acc-eg-integration-nomad"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
					//resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_constraints.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_constraints.0.key", "test.key.nomad"),
					//resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_constraints.0.value", "test.value.nomad"),
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
					//resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_constraints.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_constraints.0.key", "test.key.nomad.update"),
					//resource.TestCheckResourceAttr(resourceName, "integration_nomad.0.autoscale_constraints.0.value", "test.value.nomad.update"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationNomadGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_nomad.#", "0"),
				),
			},
		},
	})
}

const testIntegrationNomadGroupConfig_Create = `
 // --- INTEGRATION: NOMAD --------------
 integration_nomad {
    master_host = "nomad-master-host"
    master_port = 8000

    autoscale_is_enabled = false
    autoscale_cooldown = 300

    autoscale_headroom {
      cpu_per_unit = 1024
      memory_per_unit = 512
      num_of_units = 2
    }

    autoscale_down {
      evaluation_periods = 300
    }

    //autoscale_constraints {
    //  key   = "test.key.nomad"
    //  value = "test.value.nomad"
    //}
  }
 // --------------------------------------
`

const testIntegrationNomadGroupConfig_Update = `
 // --- INTEGRATION: NOMAD --------------
 integration_nomad {
		master_host = "nomad-master-host-update"
    master_port = 9000

    autoscale_is_enabled = true
    autoscale_cooldown = 180

    autoscale_headroom {
      cpu_per_unit = 2048
      memory_per_unit = 1024
      num_of_units = 1
    }

    autoscale_down {
      evaluation_periods = 150
    }

    //autoscale_constraints {
    //  key   = "test.key.nomad.update"
    //  value = "test.value.nomad.update"
    //}
  }
 // --------------------------------------
`

const testIntegrationNomadGroupConfig_EmptyFields = `
 // --- INTEGRATION: NOMAD --------------
 // -------------------------------------
`

// endregion

// region Docker Swarm integration

func TestAccSpotinstElastigroupAWS_IntegrationDockerSwarm(t *testing.T) {
	groupName := "test-acc-eg-integration-docker-swarm"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationDockerSwarmGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.master_host", "docker-swarm-master-host"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.master_port", "8000"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.autoscale_is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.autoscale_cooldown", "300"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.autoscale_headroom.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.autoscale_headroom.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.autoscale_headroom.0.num_of_units", "2"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.autoscale_down.0.evaluation_periods", "300"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationDockerSwarmGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.master_host", "docker-swarm-master-host-update"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.master_port", "9000"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.autoscale_is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.autoscale_cooldown", "180"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.autoscale_headroom.0.cpu_per_unit", "2048"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.autoscale_headroom.0.memory_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.autoscale_headroom.0.num_of_units", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.autoscale_down.0.evaluation_periods", "150"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationDockerSwarmGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.#", "0"),
				),
			},
		},
	})
}

const testIntegrationDockerSwarmGroupConfig_Create = `
 // --- INTEGRATION: DOCKER SWARM -------
 integration_docker_swarm {
    master_host = "docker-swarm-master-host"
    master_port = 8000

    autoscale_is_enabled = false
    autoscale_cooldown = 300

    autoscale_headroom {
      cpu_per_unit = 1024
      memory_per_unit = 512
      num_of_units = 2
    }

    autoscale_down {
      evaluation_periods = 300
    }
 }
 // -------------------------------------
`

const testIntegrationDockerSwarmGroupConfig_Update = `
 // --- INTEGRATION: DOCKER SWARM -------
 integration_docker_swarm {
		master_host = "docker-swarm-master-host-update"
    master_port = 9000

    autoscale_is_enabled = true
    autoscale_cooldown = 180

    autoscale_headroom {
      cpu_per_unit = 2048
      memory_per_unit = 1024
      num_of_units = 1
    }

    autoscale_down {
      evaluation_periods = 150
    }
  }
 // -------------------------------------
`

const testIntegrationDockerSwarmGroupConfig_EmptyFields = `
 // --- INTEGRATION: DOCKER SWARM -------
 // -------------------------------------
`

// endregion

// region Elastigroup: Gitlab Integration
func TestAccSpotinstElastigroupAWS_IntegrationGitlab(t *testing.T) {
	groupName := "test-acc-eg-integration-gitlab"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationGitlabGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_gitlab.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_gitlab.0.runner.0.is_enabled", "true"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationGitlabGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_gitlab.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_gitlab.0.runner.0.is_enabled", "false"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationGitlabGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_gitlab.#", "0"),
				),
			},
		},
	})
}

const testIntegrationGitlabGroupConfig_Create = `
 // --- INTEGRATION: GITLAB ---
 integration_gitlab {
	runner {
		is_enabled = true
	}
 }
 // ----------------------------
`

const testIntegrationGitlabGroupConfig_Update = `
 // --- INTEGRATION: GITLAB ---
 integration_gitlab {
    runner {
		is_enabled = false
	}
 }
 // ----------------------------
`

const testIntegrationGitlabGroupConfig_EmptyFields = `
 // --- INTEGRATION: GITLAB --------------
 // -------------------------------------
`

// endregion

// region Elastigroup: Mesosphere Integration
func TestAccSpotinstElastigroupAWS_IntegrationMesosphere(t *testing.T) {
	groupName := "test-acc-eg-integration-mesosphere"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationMesosphereGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_mesosphere.#", "0"),
				),
			},
		},
	})
}

const testIntegrationMesosphereGroupConfig_Create = `
 // --- INTEGRATION: MESOSPHERE --------------
 integration_mesosphere {
    api_server = "mesosphere-api-server"
 }
 // ------------------------------------------
`

const testIntegrationMesosphereGroupConfig_Update = `
 // --- INTEGRATION: MESOSPHERE --------------
 integration_mesosphere {
		api_server = "mesosphere-api-server-update"
 }
 // ------------------------------------------
`

const testIntegrationMesosphereGroupConfig_EmptyFields = `
 // --- INTEGRATION: MESOSPHERE --------------
 // ------------------------------------------
`

// endregion

// region Elastigroup: Multai Runtime Integration
func TestAccSpotinstElastigroupAWS_IntegrationMultaiRuntime(t *testing.T) {
	groupName := "test-acc-eg-integration-multai-runtime"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testIntegrationMultaiRuntimeGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_multai_runtime.#", "0"),
				),
			},
		},
	})
}

const testIntegrationMultaiRuntimeGroupConfig_Create = `
 // --- INTEGRATION: MULTAI-RUNTIME ------
 integration_multai_runtime {
    deployment_id = "multai-deployment-id"
  }
 // --------------------------------------
`

const testIntegrationMultaiRuntimeGroupConfig_Update = `
 // --- INTEGRATION: MULTAI-RUNTIME ------
 integration_multai_runtime {
    deployment_id = "multai-deployment-id-update"
  }
 // --------------------------------------
`

const testIntegrationMultaiRuntimeGroupConfig_EmptyFields = `
 // --- INTEGRATION: MULTAI-RUNTIME ------
 // --------------------------------------
`

// endregion

// region Elastigroup: Update Policy
func TestAccSpotinstElastigroupAWS_UpdatePolicy(t *testing.T) {
	groupName := "test-acc-eg-update-policy"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
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
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.auto_apply_tags", "false"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.should_roll", "false"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_size_percentage", "33"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.grace_period", "300"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.health_check_type", "ELB"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.strategy.0.action", "REPLACE_SERVER"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.strategy.0.batch_min_healthy_percentage", "50"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.strategy.0.should_drain_instances", "false"),
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
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.auto_apply_tags", "true"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.should_roll", "true"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_size_percentage", "66"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.grace_period", "600"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.health_check_type", "EC2"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.strategy.0.action", "RESTART_SERVER"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.strategy.0.batch_min_healthy_percentage", "20"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.strategy.0.should_drain_instances", "true"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testUpdatePolicyGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "update_policy.#", "0"),
				),
			},
		},
	})
}

const testUpdatePolicyGroupConfig_Create = `
 // --- UPDATE POLICY ----------------
  update_policy {
    should_resume_stateful = false
    auto_apply_tags = false
    should_roll = false

    roll_config {
      batch_size_percentage = 33
      grace_period = 300
      health_check_type = "ELB"

      strategy {
        action = "REPLACE_SERVER"
        should_drain_instances = false
        //batch_min_healthy_percentage = 10
      }
    }
  }
 // ----------------------------------
`

const testUpdatePolicyGroupConfig_Update = `
 // --- UPDATE POLICY ----------------
  update_policy {
    should_resume_stateful = true
    auto_apply_tags = true
    should_roll = true

    roll_config {
      batch_size_percentage = 66
      grace_period = 600
      health_check_type = "EC2"

      strategy {
        action = "RESTART_SERVER"
        should_drain_instances = true
        batch_min_healthy_percentage = 20
      }
    }
  }
 // ----------------------------------
`

const testUpdatePolicyGroupConfig_EmptyFields = `
 // --- UPDATE POLICY ----------------
 // ----------------------------------
`

// endregion

// region Wait for Capacity

func TestAccSpotinstElastigroupAWS_WaitForCapacity(t *testing.T) {
	groupName := "test-acc-eg-update-policy"
	resourceName := createElastigroupResourceName(groupName)

	var group aws.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAwaitCapacity_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "wait_for_capacity", "0"),
					resource.TestCheckResourceAttr(resourceName, "wait_for_capacity_timeout", "30"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAwaitCapacity_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "wait_for_capacity", "0"),
					resource.TestCheckResourceAttr(resourceName, "wait_for_capacity_timeout", "0"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupTerraform(&GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAwaitCapacity_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, resourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "wait_for_capacity", "0"),
					resource.TestCheckResourceAttr(resourceName, "wait_for_capacity_timeout", "0"),
				),
			},
		},
	})
}

const testAwaitCapacity_Create = `
	wait_for_capacity = 0
	wait_for_capacity_timeout = 30
`

const testAwaitCapacity_Update = `
	wait_for_capacity = 0
	wait_for_capacity_timeout = 0
`

const testAwaitCapacity_EmptyFields = `

`

// endregion
