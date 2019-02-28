package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_azure_launch_configuration"
	"log"
	"testing"
)

func createElastigroupAzureResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.ElastigroupAzureResourceName), name)
}

func testElastigroupAzureDestroy(s *terraform.State) error {
	client := testAccProviderAzure.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.ElastigroupAzureResourceName) {
			continue
		}
		input := &azure.ReadGroupInput{GroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.elastigroup.CloudProviderAzure().Read(context.Background(), input)
		if err == nil && resp != nil && resp.Group != nil {
			return fmt.Errorf("group still exists")
		}
	}
	return nil
}

func testCheckElastigroupAzureAttributes(group *azure.Group, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(group.Name) != expectedName {
			return fmt.Errorf("bad content: %v", group.Name)
		}
		return nil
	}
}

func testCheckElastigroupAzureExists(group *azure.Group, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAzure.Meta().(*Client)
		input := &azure.ReadGroupInput{GroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.elastigroup.CloudProviderAzure().Read(context.Background(), input)
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

type AzureGroupConfigMetadata struct {
	variables            string
	provider             string
	groupName            string
	vmSizes              string
	launchConfig         string
	strategy             string
	image                string
	loadBalancers        string
	network              string
	login                string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createElastigroupAzureTerraform(gcm *AzureGroupConfigMetadata) string {
	if gcm == nil {
		return ""
	}

	if gcm.provider == "" {
		gcm.provider = "azure"
	}

	if gcm.vmSizes == "" {
		gcm.vmSizes = testAzureVMSizesGroupConfig_Create
	}

	if gcm.launchConfig == "" {
		gcm.launchConfig = testAzureLaunchConfigurationGroupConfig_Create
	}

	if gcm.strategy == "" {
		gcm.strategy = testAzureStrategyGroupConfig_Create
	}

	if gcm.image == "" {
		gcm.image = testAzureImageGroupConfig_Create
	}

	if gcm.loadBalancers == "" {
		gcm.loadBalancers = testAzureLoadBalancersGroupConfig_Create
	}

	if gcm.network == "" {
		gcm.network = testAzureNetworkGroupConfig_Create
	}

	if gcm.login == "" {
		gcm.login = testAzureLoginGroupConfig_Create
	}

	template := `provider "azure" {
 account = "fake"
 token = "fake"
}`
	if gcm.updateBaselineFields {
		format := testBaselineAzureGroupConfig_Update
		template += fmt.Sprintf(format,
			gcm.groupName,
			gcm.provider,
			gcm.groupName,
			gcm.vmSizes,
			gcm.launchConfig,
			gcm.strategy,
			gcm.image,
			gcm.loadBalancers,
			gcm.network,
			gcm.login,
			gcm.fieldsToAppend,
		)
	} else {
		format := testBaselineAzureGroupConfig_Create
		template += fmt.Sprintf(format,
			gcm.groupName,
			gcm.provider,
			gcm.groupName,
			gcm.vmSizes,
			gcm.launchConfig,
			gcm.strategy,
			gcm.image,
			gcm.loadBalancers,
			gcm.network,
			gcm.login,
			gcm.fieldsToAppend,
		)
	}

	if gcm.variables != "" {
		template = gcm.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", gcm.groupName, template)
	return template
}

// region Elastigroup Azure: Baseline
func TestAccSpotinstElastigroupAzure_Baseline(t *testing.T) {
	groupName := "eg-azure-baseline"
	resourceName := createElastigroupAzureResourceName(groupName)

	var group azure.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAzureDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{groupName: groupName}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "eastus"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "0"),
				),
			},
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{groupName: groupName, updateBaselineFields: true}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "eastus"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "0"),
				),
			},
		},
	})
}

const testBaselineAzureGroupConfig_Create = `
resource "` + string(commons.ElastigroupAzureResourceName) + `" "%v" {
 provider = "%v"

 name 				 = "%v"
 product 			 = "Linux"
 region              = "eastus"
 resource_group_name = "alex-test"

 // --- CAPACITY ------------
 max_size 		  = 0
 min_size 		  = 0
 desired_capacity = 0
 // -------------------------
 
 %v
 %v
 %v
 %v
 %v
 %v
 %v
 %v
}

`

const testBaselineAzureGroupConfig_Update = `
resource "` + string(commons.ElastigroupAzureResourceName) + `" "%v" {
 provider = "%v"

 name 				 = "%v"
 product 			 = "Linux"
 region              = "eastus"
 resource_group_name = "alex-test"

 // --- CAPACITY ------------
 max_size 		  = 0
 min_size 		  = 0
 desired_capacity = 0
 // -------------------------
 
 %v
 %v
 %v
 %v
 %v
 %v
 %v
 %v
}

`

// endregion

// region Azure Elastigroup: Health Checks
func TestAccSpotinstElastigroupAzure_HealthChecks(t *testing.T) {
	groupName := "eg-azure-health-checks"
	resourceName := createElastigroupAzureResourceName(groupName)

	var group azure.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "azure") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupAzureDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureHealthChecksGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "health_check.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_check.0.health_check_type", "INSTANCE_STATE"),
					resource.TestCheckResourceAttr(resourceName, "health_check.0.auto_healing", "true"),
					resource.TestCheckResourceAttr(resourceName, "health_check.0.grace_period", "180"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureHealthChecksGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "health_check.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_check.0.health_check_type", "INSTANCE_STATE"),
					resource.TestCheckResourceAttr(resourceName, "health_check.0.auto_healing", "false"),
					resource.TestCheckResourceAttr(resourceName, "health_check.0.grace_period", "240"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureHealthChecksGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "health_check.#", "0"),
				),
			},
		},
	})
}

const testAzureHealthChecksGroupConfig_Create = `
 // --- HEALTH-CHECKS ------------------------------------
 health_check = {
   health_check_type = "INSTANCE_STATE"
   auto_healing      = true
   grace_period      = 180	 
 }
 // ------------------------------------------------------
`

const testAzureHealthChecksGroupConfig_Update = `
 // --- HEALTH-CHECKS ------------------------------------
 health_check = {
   health_check_type = "INSTANCE_STATE"
   auto_healing      = false
   grace_period      = 240	 
 }
 // ------------------------------------------------------
`

const testAzureHealthChecksGroupConfig_EmptyFields = `
 // --- HEALTH-CHECKS ------------------------------------
 // ------------------------------------------------------
`

// endregion

// region Azure Elastigroup: Image
func TestAccSpotinstElastigroupAzure_Image(t *testing.T) {
	groupName := "eg-azure-image"
	resourceName := createElastigroupAzureResourceName(groupName)

	var group azure.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAzureDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName: groupName,
					image:     testAzureImageGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "image.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.offer", "UbuntuServer"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.publisher", "Canonical"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.sku", "16.04-LTS"),
				),
			},
		},
	})
}

const testAzureImageGroupConfig_Create = `
// --- IMAGES --------------------------------
  image = {
    marketplace = {
      publisher = "Canonical"
      offer = "UbuntuServer"
      sku = "16.04-LTS"
    }
  }
// ---------------------------------------------
`

// endregion

// region Azure Elastigroup: Launch Configuration
func TestAccSpotinstElastigroupAzure_LaunchConfiguration(t *testing.T) {
	groupName := "eg-azure-launch-configuration"
	resourceName := createElastigroupAzureResourceName(groupName)

	var group azure.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAzureDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName:    groupName,
					launchConfig: testAzureLaunchConfigurationGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_azure_launch_configuration.HexStateFunc("hello world"))),
			},
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName:    groupName,
					launchConfig: testAzureLaunchConfigurationGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_azure_launch_configuration.HexStateFunc("hello world"))),
			},
		},
	})
}

const testAzureLaunchConfigurationGroupConfig_Create = `
// --- LAUNCH CONFIGURATION --------------------
user_data = "hello world"
// ---------------------------------------------
`

const testAzureLaunchConfigurationGroupConfig_Update = `
// --- LAUNCH CONFIGURATION --------------------
user_data = "hello world"
// ---------------------------------------------
`

// endregion

// region Azure Elastigroup: Load Balancers
func TestAccSpotinstElastigroupAzure_LoadBalancers(t *testing.T) {
	groupName := "eg-azure-load-balancers"
	resourceName := createElastigroupAzureResourceName(groupName)

	var group azure.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAzureDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName:     groupName,
					loadBalancers: testAzureLoadBalancersGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.116057219.auto_weight", "true"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.116057219.balancer_id", "lb-0be85d6aa269"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.116057219.target_set_id", "ts-ae9c9603c365"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.116057219.type", "MULTAI_TARGET_SET"),
				),
			},
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName:     groupName,
					loadBalancers: testAzureLoadBalancersGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.9777277.auto_weight", "false"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.9777277.balancer_id", "lb-0be85d6aa269"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.9777277.target_set_id", "ts-ae9c9603c365"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.9777277.type", "MULTAI_TARGET_SET"),
				),
			},
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName:     groupName,
					loadBalancers: testAzureLoadBalancersGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "0"),
				),
			},
		},
	})
}

const testAzureLoadBalancersGroupConfig_Create = `
// --- LOAD BALANCERS --------------------------
  load_balancers = [
    {
      type = "MULTAI_TARGET_SET"
      balancer_id = "lb-0be85d6aa269"
      target_set_id = "ts-ae9c9603c365"
      auto_weight = true
    }
  ]
// ---------------------------------------------
`

const testAzureLoadBalancersGroupConfig_Update = `
// --- LOAD BALANCERS --------------------------
  load_balancers = [
    {
      type = "MULTAI_TARGET_SET"
      balancer_id = "lb-0be85d6aa269"
      target_set_id = "ts-ae9c9603c365"
      auto_weight = false
    }
  ]
// ---------------------------------------------
`

const testAzureLoadBalancersGroupConfig_EmptyFields = `
// --- LOAD BALANCERS --------------------------
// ---------------------------------------------
`

// endregion

// region Azure Elastigroup: Login
func TestAccSpotinstElastigroupAzure_Login(t *testing.T) {
	groupName := "eg-azure-login"
	resourceName := createElastigroupAzureResourceName(groupName)

	var group azure.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAzureDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName: groupName,
					login:     testAzureLoginGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "login.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "login.0.user_name", "alex-test"),
					resource.TestCheckResourceAttr(resourceName, "login.0.ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDLn7RIjgivW2nWoh56XV2wpDKjjWFk92UgfTsqL8qYI0lGCJuoGeMlje1VWyAemMGZZsK5et8j3/caZsVd1Hypui3xV+tRAmtnyqVCjDGYsBQIMKoDzLrrZz7/s2WNKbMegOgQ+8YxXLhxuS5YGKhNjvxC2kJCe1HkPAPvx03kzNGmxxv6pt5TaQPXUqVxfWoeoaLRDcL8Ns2kikZC6v2cfY/PcmwoYd7XlVuILLTMNF6ujOUsX9kHt/910dEW66iZpc+PjHnKuAu/5238lssiUZULTHWbjE09MG8kHIiZq9Z9HgmAS++YLUc2G9InBqiLXMbie4S9qMcp+crl1oG/"),
				),
			},
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName: groupName,
					login:     testAzureLoginGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "login.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "login.0.user_name", "alex-test"),
					resource.TestCheckResourceAttr(resourceName, "login.0.ssh_public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDLn7RIjgivW2nWoh56XV2wpDKjjWFk92UgfTsqL8qYI0lGCJuoGeMlje1VWyAemMGZZsK5et8j3/caZsVd1Hypui3xV+tRAmtnyqVCjDGYsBQIMKoDzLrrZz7/s2WNKbMegOgQ+8YxXLhxuS5YGKhNjvxC2kJCe1HkPAPvx03kzNGmxxv6pt5TaQPXUqVxfWoeoaLRDcL8Ns2kikZC6v2cfY/PcmwoYd7XlVuILLTMNF6ujOUsX9kHt/910dEW66iZpc+PjHnKuAu/5238lssiUZULTHWbjE09MG8kHIiZq9Z9HgmAS++YLUc2G9InBqiLXMbie4S9qMcp+crl1oG/"),
				),
			},
		},
	})
}

const testAzureLoginGroupConfig_Create = `
// --- LOGIN ---------------------------------
  login = {
    user_name = "alex-test"
    ssh_public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDLn7RIjgivW2nWoh56XV2wpDKjjWFk92UgfTsqL8qYI0lGCJuoGeMlje1VWyAemMGZZsK5et8j3/caZsVd1Hypui3xV+tRAmtnyqVCjDGYsBQIMKoDzLrrZz7/s2WNKbMegOgQ+8YxXLhxuS5YGKhNjvxC2kJCe1HkPAPvx03kzNGmxxv6pt5TaQPXUqVxfWoeoaLRDcL8Ns2kikZC6v2cfY/PcmwoYd7XlVuILLTMNF6ujOUsX9kHt/910dEW66iZpc+PjHnKuAu/5238lssiUZULTHWbjE09MG8kHIiZq9Z9HgmAS++YLUc2G9InBqiLXMbie4S9qMcp+crl1oG/"
  }
// -------------------------------------------
`

const testAzureLoginGroupConfig_Update = `
// --- LOGIN ---------------------------------
  login = {
    user_name = "alex-test"
    ssh_public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDLn7RIjgivW2nWoh56XV2wpDKjjWFk92UgfTsqL8qYI0lGCJuoGeMlje1VWyAemMGZZsK5et8j3/caZsVd1Hypui3xV+tRAmtnyqVCjDGYsBQIMKoDzLrrZz7/s2WNKbMegOgQ+8YxXLhxuS5YGKhNjvxC2kJCe1HkPAPvx03kzNGmxxv6pt5TaQPXUqVxfWoeoaLRDcL8Ns2kikZC6v2cfY/PcmwoYd7XlVuILLTMNF6ujOUsX9kHt/910dEW66iZpc+PjHnKuAu/5238lssiUZULTHWbjE09MG8kHIiZq9Z9HgmAS++YLUc2G9InBqiLXMbie4S9qMcp+crl1oG/"
  }
// -------------------------------------------
`

// region Azure Elastigroup: Network
func TestAccSpotinstElastigroupAzure_Network(t *testing.T) {
	groupName := "eg-azure-network"
	resourceName := createElastigroupAzureResourceName(groupName)

	var group azure.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAzureDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName: groupName,
					network:   testAzureNetworkGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "network.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "network.0.resource_group_name", "alex-test"),
					resource.TestCheckResourceAttr(resourceName, "network.0.subnet_name", "alex-test-subnet"),
					resource.TestCheckResourceAttr(resourceName, "network.0.virtual_network_name", "alex-test-netwrk"),
				),
			},
		},
	})
}

const testAzureNetworkGroupConfig_Create = `
// --- NETWORK ---------------------------------
  network = {
    virtual_network_name = "alex-test-netwrk"
    subnet_name = "alex-test-subnet"                 
    resource_group_name = "alex-test"         
    assign_public_ip = true                
  }
// ---------------------------------------------
`

// endregion

// region Azure Elastigroup: Strategy
func TestAccSpotinstElastigroupAzure_Strategy(t *testing.T) {
	groupName := "eg-azure-strategy"
	resourceName := createElastigroupAzureResourceName(groupName)

	var group azure.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAzureDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName: groupName,
					strategy:  testAzureStrategyGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.low_priority_percentage", "100"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.draining_timeout", "180"),
				),
			},
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName: groupName,
					strategy:  testAzureStrategyGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.od_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.draining_timeout", "240"),
				),
			},
		},
	})
}

const testAzureStrategyGroupConfig_Create = `
// --- STRATEGY --------------------------------
  strategy = {
    low_priority_percentage = 100
    draining_timeout = 180
  }
// ---------------------------------------------
`

const testAzureStrategyGroupConfig_Update = `
// --- STRATEGY --------------------------------
  strategy = {
    od_count = 1
    draining_timeout = 240
  }
// ---------------------------------------------
`

// endregion

// region Azure Elastigroup: VM Sizes
func TestAccSpotinstElastigroupAzure_VMSizes(t *testing.T) {
	groupName := "eg-azure-vm-sizes"
	resourceName := createElastigroupAzureResourceName(groupName)

	var group azure.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAzureDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName: groupName,
					vmSizes:   testAzureVMSizesGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "od_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "od_sizes.0", "basic_a1"),
					resource.TestCheckResourceAttr(resourceName, "low_priority_sizes.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "low_priority_sizes.0", "basic_a2"),
					resource.TestCheckResourceAttr(resourceName, "low_priority_sizes.1", "basic_a1"),
					resource.TestCheckResourceAttr(resourceName, "low_priority_sizes.2", "basic_a3"),
				),
			},
			{
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName: groupName,
					vmSizes:   testAzureVMSizesGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "od_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "od_sizes.0", "basic_a2"),
					resource.TestCheckResourceAttr(resourceName, "low_priority_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "low_priority_sizes.0", "basic_a2"),
				),
			},
		},
	})
}

const testAzureVMSizesGroupConfig_Create = `
// --- VM SIZES --------------------------------------------
 od_sizes           = ["basic_a1"]
 low_priority_sizes = ["basic_a2", "basic_a1", "basic_a3"]
// ---------------------------------------------------------
`

const testAzureVMSizesGroupConfig_Update = `
// --- VM SIZES --------------------------------
 od_sizes           = ["basic_a2"]
 low_priority_sizes = ["basic_a2"]
// ---------------------------------------------
`

// endregion

// region Azure Elastigroup: Scheduled Task
func TestAccSpotinstElastigroupAzure_ScheduledTask(t *testing.T) {
	groupName := "eg-azure-scheduled-task"
	resourceName := createElastigroupAzureResourceName(groupName)

	var group azure.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "azure") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupAzureDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureScheduledTaskGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3930008834.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3930008834.task_type", "scale"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3930008834.scale_min_capacity", "5"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3930008834.scale_max_capacity", "8"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3930008834.adjustment", "2"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3930008834.scale_target_capacity", "6"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3930008834.batch_size_percentage", "33"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.3930008834.grace_period", "300"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureScheduledTaskGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.705557774.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.705557774.task_type", "scale"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.705557774.scale_min_capacity", "0"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.705557774.scale_max_capacity", "10"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.705557774.adjustment_percentage", "50"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.705557774.scale_target_capacity", "5"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.705557774.batch_size_percentage", "50"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.705557774.grace_period", "360"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureScheduledTaskGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "0"),
				),
			},
		},
	})
}

const testAzureScheduledTaskGroupConfig_Create = `
 // --- SCHEDULED TASK ------------------
  scheduled_task = [{
    is_enabled = true
    cron_expression = "* * * * *"
    task_type = "scale"
    scale_min_capacity = 5
    scale_max_capacity = 8
    adjustment = 2
    adjustment_percentage = 50
    scale_target_capacity = 6
    batch_size_percentage = 33
    grace_period = 300
  }]
 // -------------------------------------
`

const testAzureScheduledTaskGroupConfig_Update = `
 // --- SCHEDULED TASK ------------------
  scheduled_task = [{
    is_enabled = false
    cron_expression = "* * * * *"
    task_type = "scale"
    scale_min_capacity = 0
    scale_max_capacity = 10
    adjustment_percentage = 50
    scale_target_capacity = 5
    batch_size_percentage = 50
    grace_period = 360
  }]
 // -------------------------------------
`

const testAzureScheduledTaskGroupConfig_EmptyFields = `
 // --- SCHEDULED TASK ------------------
 // -------------------------------------
`

// endregion

// region Elastigroup: Update Policy
func TestAccSpotinstElastigroupAzure_UpdatePolicy(t *testing.T) {
	groupName := "eg-azure-update-policy"
	resourceName := createElastigroupAzureResourceName(groupName)

	var group azure.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "azure") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupAzureDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureUpdatePolicyGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "update_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.should_roll", "false"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_size_percentage", "33"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.grace_period", "300"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.health_check_type", "NONE"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureUpdatePolicyGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "update_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.should_roll", "true"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_size_percentage", "66"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.grace_period", "600"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.health_check_type", "INSTANCE_STATE"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureUpdatePolicyGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureExists(&group, resourceName),
					testCheckElastigroupAzureAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "update_policy.#", "0"),
				),
			},
		},
	})
}

const testAzureUpdatePolicyGroupConfig_Create = `
 // --- UPDATE POLICY ----------------
  update_policy = {
    should_roll = false
    roll_config = {
      batch_size_percentage = 33
      grace_period = 300
      health_check_type = "NONE"
    }
  }
 // ----------------------------------
`

const testAzureUpdatePolicyGroupConfig_Update = `
 // --- UPDATE POLICY ----------------
  update_policy = {
    should_roll = true
    roll_config = {
      batch_size_percentage = 66
      grace_period = 600
      health_check_type = "INSTANCE_STATE"
    }
  }
 // ----------------------------------
`

const testAzureUpdatePolicyGroupConfig_EmptyFields = `
 // --- UPDATE POLICY ----------------
 // ----------------------------------
`

// endregion
