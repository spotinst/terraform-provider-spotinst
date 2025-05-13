package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func init() {
	resource.AddTestSweepers("spotinst_elastigroup_azure_v3", &resource.Sweeper{
		Name: "spotinst_elastigroup_azure_v3",
		F:    testSweepElastigroupAzureV3,
	})
}

func testSweepElastigroupAzureV3(region string) error {
	client, err := getProviderClient("azure")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).elastigroup.CloudProviderAzureV3()

	input := &azurev3.ListGroupsInput{}
	if resp, err := conn.List(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of groups to sweep")
	} else {
		if len(resp.Groups) == 0 {
			log.Printf("[INFO] No groups to sweep")
		}
		for _, group := range resp.Groups {
			if strings.Contains(spotinst.StringValue(group.Name), "test-acc-") {
				if _, err := conn.Delete(context.Background(), &azurev3.DeleteGroupInput{GroupID: group.ID}); err != nil {
					return fmt.Errorf("unable to delete group %v in sweep", spotinst.StringValue(group.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(group.ID))
				}
			}
		}
	}
	return nil
}

func createElastigroupAzureV3ResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.ElastigroupAzureV3ResourceName), name)
}

func testElastigroupAzureV3Destroy(s *terraform.State) error {
	client := testAccProviderAzure.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.ElastigroupAzureV3ResourceName) {
			continue
		}
		input := &azurev3.ReadGroupInput{GroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.elastigroup.CloudProviderAzureV3().Read(context.Background(), input)
		if err == nil && resp != nil && resp.Group != nil {
			return fmt.Errorf("group still exists")
		}
	}
	return nil
}

func testCheckElastigroupAzureV3Attributes(group *azurev3.Group, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(group.Name) != expectedName {
			return fmt.Errorf("bad content: %v", group.Name)
		}
		return nil
	}
}

func testCheckElastigroupAzureV3Exists(group *azurev3.Group, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAzure.Meta().(*Client)
		input := &azurev3.ReadGroupInput{GroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.elastigroup.CloudProviderAzureV3().Read(context.Background(), input)
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

type AzureV3GroupConfigMetadata struct {
	variables            string
	provider             string
	groupName            string
	vmSizes              string
	strategy             string
	image                string
	network              string
	login                string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createElastigroupAzureV3Terraform(gcm *AzureV3GroupConfigMetadata) string {
	if gcm == nil {
		return ""
	}

	if gcm.provider == "" {
		gcm.provider = "azure"
	}

	if gcm.vmSizes == "" {
		gcm.vmSizes = testAzureV3VMSizesGroupConfig_Create
	}

	if gcm.strategy == "" {
		gcm.strategy = testAzureV3StrategyGroupConfig_Create
	}

	if gcm.image == "" {
		gcm.image = testAzureV3ImageGroupConfig_Create
	}

	if gcm.network == "" {
		gcm.network = testAzureV3NetworkGroupConfig_Create
	}

	if gcm.login == "" {
		gcm.login = testAzureV3LoginGroupConfig_Create
	}

	template := `provider "azure" {
 account = "fake"
 token = "fake"
}`
	if gcm.updateBaselineFields {
		format := testBaselineAzureV3GroupConfig_Update
		template += fmt.Sprintf(format,
			gcm.groupName,
			gcm.provider,
			gcm.groupName,
			gcm.vmSizes,
			gcm.strategy,
			gcm.image,
			gcm.network,
			gcm.login,
			gcm.fieldsToAppend,
		)
	} else {
		format := testBaselineAzureV3GroupConfig_Create
		template += fmt.Sprintf(format,
			gcm.groupName,
			gcm.provider,
			gcm.groupName,
			gcm.vmSizes,
			gcm.strategy,
			gcm.image,
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
func TestAccSpotinstElastigroupAzureV3_Baseline(t *testing.T) {
	groupName := "test-acc-eg-azure-v3-baseline"
	resourceName := createElastigroupAzureV3ResourceName(groupName)

	var group azurev3.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config:             createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{groupName: groupName}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "eastus"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "0"),
					resource.TestCheckResourceAttr(resourceName, "os", "Linux"),
					resource.TestCheckResourceAttr(resourceName, "custom_data", "IyEvY=IyEvYmluL2Jhc2gKZWNobyAidGVzdCI="),
					resource.TestCheckResourceAttr(resourceName, "managed_service_identity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_service_identity.0.resource_group_name", "AutomationResourceGroup"),
					resource.TestCheckResourceAttr(resourceName, "managed_service_identity.0.name", "AutomationResourceIdentity"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key1"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.1.key", "key2"),
					resource.TestCheckResourceAttr(resourceName, "tags.1.value", "value2"),
				),
			},
			{
				Config:             createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{groupName: groupName, updateBaselineFields: true}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "eastus"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "5"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "0"),
					resource.TestCheckResourceAttr(resourceName, "os", "Linux"),
					resource.TestCheckResourceAttr(resourceName, "managed_service_identity.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "managed_service_identity.0.resource_group_name", "AutomationResourceGroup"),
					resource.TestCheckResourceAttr(resourceName, "managed_service_identity.0.name", "AutomationResourceIdentity"),
					resource.TestCheckResourceAttr(resourceName, "managed_service_identity.1.resource_group_name", "AutomationResourceGroup"),
					resource.TestCheckResourceAttr(resourceName, "managed_service_identity.1.name", "AutomationResourceIdentity2"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key1"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.1.key", "key2"),
					resource.TestCheckResourceAttr(resourceName, "tags.1.value", "value2"),
					resource.TestCheckResourceAttr(resourceName, "tags.2.key", "key3"),
					resource.TestCheckResourceAttr(resourceName, "tags.2.value", "value3"),
				),
			},
		},
	})
}

const testBaselineAzureV3GroupConfig_Create = `
resource "` + string(commons.ElastigroupAzureV3ResourceName) + `" "%v" {
 provider = "%v"

 name 				 = "%v"
 os 			     = "Linux"
 region              = "eastus"
 resource_group_name = "AutomationResourceGroup"

 // --- CAPACITY ------------
 max_size 		  = 0
 min_size 		  = 0
 desired_capacity = 0
 // -------------------------
 

 // --- LAUNCHSPEC ----------
 custom_data = "IyEvY=IyEvYmluL2Jhc2gKZWNobyAidGVzdCI="

 managed_service_identity {
    resource_group_name = "AutomationResourceGroup"
    name                = "AutomationResourceIdentity"
  }

 tags {
    key = "key1"
    value = "value1"
  }

  tags {
    key = "key2"
    value = "value2"
  }
 // -------------------------

 %v
 %v
 %v
 %v
 %v
 %v
}

`

const testBaselineAzureV3GroupConfig_Update = `
resource "` + string(commons.ElastigroupAzureV3ResourceName) + `" "%v" {
 provider = "%v"

 name 				 = "%v"
 os 			     = "Linux"
 region              = "eastus"
 resource_group_name = "AutomationResourceGroup"

 // --- CAPACITY ------------
 max_size 		  = 5
 min_size 		  = 0
 desired_capacity = 0
 // -------------------------

 managed_service_identity {
    resource_group_name = "AutomationResourceGroup"
    name                = "AutomationResourceIdentity"
  }

 managed_service_identity {
    resource_group_name = "AutomationResourceGroup"
    name                = "AutomationResourceIdentity2"
  }

 tags {
    key = "key1"
    value = "value1"
  }

  tags {
    key = "key2"
    value = "value2"
  }

 tags {
    key = "key3"
    value = "value3"
  }
 
 %v
 %v
 %v
 %v
 %v
 %v
}

`

// endregion

// region Azure Elastigroup: Image
func TestAccSpotinstElastigroupAzureV3_Image(t *testing.T) {
	groupName := "test-acc-eg-azure-v3-image"
	resourceName := createElastigroupAzureV3ResourceName(groupName)

	var group azurev3.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
					groupName: groupName,
					image:     testAzureV3ImageGroupConfig_Create,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "image.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.offer", "UbuntuServer"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.publisher", "Canonical"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.sku", "18.04-LTS"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.version", "latest"),
				),
			},
		},
	})
}

const testAzureV3ImageGroupConfig_Create = `
// --- IMAGES --------------------------------
  image {
    marketplace {
      publisher = "Canonical"
      offer = "UbuntuServer"
      sku = "18.04-LTS"
      version = "latest"
    }
  }
// ---------------------------------------------
`

// endregion

// region Azure Elastigroup: Network
func TestAccSpotinstElastigroupAzureV3_Network(t *testing.T) {
	groupName := "test-acc-eg-azure-v3-network"
	resourceName := createElastigroupAzureV3ResourceName(groupName)

	var group azurev3.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "azure") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupAzureV3Destroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
					groupName: groupName,
					network:   testAzureV3NetworkGroupConfig_Create,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "network.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.virtual_network_name", "Automation-VirtualNetwork"),
					resource.TestCheckResourceAttr(resourceName, "network.0.resource_group_name", "AutomationResourceGroup"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interfaces.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interfaces.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interfaces.0.subnet_name", "Automation-PrivateSubnet"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interfaces.0.is_primary", "true"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interfaces.0.additional_ip_configs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interfaces.0.additional_ip_configs.0.name", "terraformTestSecondaryIpConfig"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interfaces.0.application_security_group.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interfaces.0.application_security_group.0.resource_group_name", "AutomationResourceGroup"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interfaces.0.application_security_group.0.name", "Terraform-Testing-ASG"),
				),
			},
		},
	})
}

const testAzureV3NetworkGroupConfig_Create = `
// --- NETWORK ---------------------------------
  network {
    virtual_network_name = "Automation-VirtualNetwork"
    resource_group_name = "AutomationResourceGroup"         
    network_interfaces {
      subnet_name = "Automation-PrivateSubnet"
      assign_public_ip = false
      is_primary = true

      additional_ip_configs {
        name = "terraformTestSecondaryIpConfig"
      }

      application_security_group {
        name = "Terraform-Testing-ASG"
        resource_group_name = "AutomationResourceGroup"
      }

  	}
  }
// ---------------------------------------------
`

// endregion

// region Azure Elastigroup: Strategy
func TestAccSpotinstElastigroupAzureV3_Strategy(t *testing.T) {
	groupName := "test-acc-eg-azure-v3-strategy"
	resourceName := createElastigroupAzureV3ResourceName(groupName)

	var group azurev3.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
					groupName: groupName,
					strategy:  testAzureV3StrategyGroupConfig_Create,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "spot_percentage", "65"),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "30"),
					resource.TestCheckResourceAttr(resourceName, "fallback_to_on_demand", "true"),
				),
			},
			{
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
					groupName: groupName,
					strategy:  testAzureV3StrategyGroupConfig_Update,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "300"),
					resource.TestCheckResourceAttr(resourceName, "on_demand_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "fallback_to_on_demand", "false"),
				),
			},
		},
	})
}

const testAzureV3StrategyGroupConfig_Create = `
// --- STRATEGY --------------------------------
    spot_percentage = 65
    draining_timeout = 30
    fallback_to_on_demand = true
// ---------------------------------------------
`

const testAzureV3StrategyGroupConfig_Update = `
// --- STRATEGY --------------------------------
    draining_timeout = 300
    fallback_to_on_demand = false
	on_demand_count  = 2
// ---------------------------------------------
`

// endregion

// region Azure Elastigroup: VM Sizes
func TestAccSpotinstElastigroupAzureV3_VMSizes(t *testing.T) {
	groupName := "test-acc-eg-azure-v3-vm-sizes"
	resourceName := createElastigroupAzureV3ResourceName(groupName)

	var group azurev3.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
					groupName: groupName,
					vmSizes:   testAzureV3VMSizesGroupConfig_Create,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.od_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.od_sizes.0", "standard_a1_v2"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.spot_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.spot_sizes.0", "standard_a1_v2"),
				),
			},
			{
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
					groupName: groupName,
					vmSizes:   testAzureV3VMSizesGroupConfig_Update,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.od_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.od_sizes.0", "standard_a1_v2"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.spot_size_attributes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.spot_size_attributes.0.min_cpu", "2"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.spot_size_attributes.0.max_cpu", "100"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.spot_size_attributes.0.min_memory", "8"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.spot_size_attributes.0.max_memory", "64"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.spot_size_attributes.0.min_storage", "32"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.spot_size_attributes.0.max_storage", "1024"),
				),
			},
		},
	})
}

const testAzureV3VMSizesGroupConfig_Create = `
// --- VM SIZES --------------------------------------------
   vm_sizes{  
     od_sizes = ["standard_a1_v2"]
     spot_sizes = ["standard_a1_v2"]
   }
// ---------------------------------------------------------
`

const testAzureV3VMSizesGroupConfig_Update = `
// --- VM SIZES --------------------------------
   vm_sizes{  
     od_sizes = ["standard_a1_v2"]
     spot_size_attributes {
		min_cpu = "2"
		max_cpu = "100"
		min_memory = "8"
		max_memory = "64"
		min_storage = "32"
		max_storage = "1024"
	}
   }
// ---------------------------------------------
`

// endregion

// region Azure Elastigroup: Login
func TestAccSpotinstElastigroupAzureV3_Login(t *testing.T) {
	groupName := "test-acc-eg-azure-v3-login"
	resourceName := createElastigroupAzureV3ResourceName(groupName)

	var group azurev3.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
					groupName: groupName,
					login:     testAzureV3LoginGroupConfig_Create,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "login.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "login.0.user_name", "azure_v3_terraform"),
					resource.TestCheckResourceAttr(resourceName, "login.0.password", ""),
				),
			},
		},
	})
}

const testAzureV3LoginGroupConfig_Create = `
// --- LOGIN --------------------------------
  login {
    user_name = "azure_v3_terraform"
	password  = "terraform-password" 
  }
// ---------------------------------------------
`

// endregion

func TestAccSpotinstElastigroupAzureV3_ScalingUpPolicies(t *testing.T) {
	groupName := "test-acc-eg-scaling-up-policy"
	resourceName := createElastigroupAzureV3ResourceName(groupName)

	var group azurev3.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "azure") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupAzureV3Destroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureV3ScalingUpPolicyGroupConfig_Create,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.policy_name", "policy-name"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.metric_name", "Disk Read Bytes"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.namespace", "Microsoft.Compute"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.source", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.statistic", "average"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.unit", "bytes"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.cooldown", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.dimensions.0.name", "name-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.dimensions.0.value", "value-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.threshold", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.operator", "gt"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.evaluation_periods", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.period", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.action.0.type", "adjustment"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.action.0.adjustment", "1"),
				),
			},
			{
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureV3ScalingUpPolicyGroupConfig_Update,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.policy_name", "policy-name-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.metric_name", "Disk Read Bytes"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.namespace", "Microsoft.Compute"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.source", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.statistic", "average"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.unit", "bytes"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.cooldown", "300"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.dimensions.0.name", "name-1-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.dimensions.0.value", "value-1-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.threshold", "5"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.operator", "lt"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.evaluation_periods", "5"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.period", "300"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.action.0.type", "updateCapacity"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.action.0.maximum", "6"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.action.0.minimum", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.action.0.target", "2"),
				),
			},
			{
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureV3ScalingUpPolicyGroupConfig_EmptyFields,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.#", "0"),
				),
			},
		},
	})
}

const testAzureV3ScalingUpPolicyGroupConfig_Create = `
 // --- SCALE UP POLICY ------------------
 scaling_up_policy {
  policy_name = "policy-name"
  metric_name = "Disk Read Bytes"
  namespace = "Microsoft.Compute"
  source = ""
  statistic = "average"
  unit = "bytes"
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

  action {
	type = "adjustment"
    adjustment = "1"
  }

  }
 // ----------------------------------------
`

const testAzureV3ScalingUpPolicyGroupConfig_Update = `
 // --- SCALE UP POLICY ------------------
 scaling_up_policy {
  policy_name = "policy-name-update"
  metric_name = "Disk Read Bytes"
  namespace = "Microsoft.Compute"
  source = ""
  statistic = "average"
  unit = "bytes"
  is_enabled = true
  cooldown = 300
  dimensions {
      name = "name-1-update"
      value = "value-1-update"
  }
  threshold = 5

  operator = "lt"
  evaluation_periods = 5
  period = 300

  action {
	type = "updateCapacity"
    minimum = 1
    maximum = 6
    target = 2
  }
  }
 // ----------------------------------------
`

const testAzureV3ScalingUpPolicyGroupConfig_EmptyFields = `
 // --- SCALE UP POLICY ------------------
 // ----------------------------------------
`

// endregion

// region Elastigroup: Scaling Down Policies
func TestAccSpotinstElastigroupAzureV3_ScalingDownPolicies(t *testing.T) {
	groupName := "test-acc-eg-scaling-down-policy"
	resourceName := createElastigroupAzureV3ResourceName(groupName)

	var group azurev3.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "aws") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupAzureV3Destroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureV3ScalingDownPolicyGroupConfig_Create,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.policy_name", "policy-name"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.metric_name", "Disk Read Bytes"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.namespace", "Microsoft.Compute"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.source", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.statistic", "average"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.unit", "bytes"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.cooldown", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.dimensions.0.name", "name-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.dimensions.0.value", "value-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.threshold", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.operator", "lt"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.evaluation_periods", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.period", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.action.0.type", "adjustment"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.action.0.adjustment", "1"),
				),
			},
			{
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureV3ScalingDownPolicyGroupConfig_Update,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.policy_name", "policy-name-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.metric_name", "Disk Read Bytes"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.namespace", "Microsoft.Compute"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.source", ""),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.statistic", "average"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.unit", "bytes"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.cooldown", "300"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.dimensions.0.name", "name-1-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.dimensions.0.value", "value-1-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.threshold", "5"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.operator", "lt"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.evaluation_periods", "5"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.period", "300"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.action.0.type", "updateCapacity"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.action.0.minimum", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.action.0.maximum", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.action.0.target", "5"),
				),
			},
			{
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testAzureV3ScalingDownPolicyGroupConfig_EmptyFields,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.#", "0"),
				),
			},
		},
	})
}

const testAzureV3ScalingDownPolicyGroupConfig_Create = `
 // --- SCALE DOWN POLICY ------------------
 scaling_down_policy {
  policy_name = "policy-name"
  metric_name = "Disk Read Bytes"
  namespace = "Microsoft.Compute"
  source = ""
  statistic = "average"
  unit = "bytes"
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
  action {
    type = "adjustment"
    adjustment = "1"
  }

  }
 // ----------------------------------------
`

const testAzureV3ScalingDownPolicyGroupConfig_Update = `
 // --- SCALE DOWN POLICY ------------------
 scaling_down_policy {
  policy_name = "policy-name-update"
  metric_name = "Disk Read Bytes"
  namespace = "Microsoft.Compute"
  source = ""
  statistic = "average"
  unit = "bytes"
  is_enabled = true
  cooldown = 300
  dimensions {
      name = "name-1-update"
      value = "value-1-update"
  }
  threshold = 5

  operator = "lt"
  evaluation_periods = 5
  period = 300

  action {
  	type = "updateCapacity"
  	minimum = 1
  	maximum = 10
	target = 5
 }

  }
 // ----------------------------------------
`

const testAzureV3ScalingDownPolicyGroupConfig_EmptyFields = `
 // --- SCALE DOWN POLICY ------------------
 // ----------------------------------------
`

// endregion
