package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
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

	input := &v3.ListGroupsInput{}
	if resp, err := conn.List(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of groups to sweep")
	} else {
		if len(resp.Groups) == 0 {
			log.Printf("[INFO] No groups to sweep")
		}
		for _, group := range resp.Groups {
			if strings.Contains(spotinst.StringValue(group.Name), "test-acc-") {
				if _, err := conn.Delete(context.Background(), &v3.DeleteGroupInput{GroupID: group.ID}); err != nil {
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
		input := &v3.ReadGroupInput{GroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.elastigroup.CloudProviderAzureV3().Read(context.Background(), input)
		if err == nil && resp != nil && resp.Group != nil {
			return fmt.Errorf("group still exists")
		}
	}
	return nil
}

func testCheckElastigroupAzureV3Attributes(group *v3.Group, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(group.Name) != expectedName {
			return fmt.Errorf("bad content: %v", group.Name)
		}
		return nil
	}
}

func testCheckElastigroupAzureV3Exists(group *v3.Group, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAzure.Meta().(*Client)
		input := &v3.ReadGroupInput{GroupID: spotinst.String(rs.Primary.ID)}
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
			gcm.fieldsToAppend,
		)
	}

	if gcm.variables != "" {
		template = gcm.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", gcm.groupName, template)
	return template
}

//region Elastigroup Azure: Baseline
func TestAccSpotinstElastigroupAzureV3_Baseline(t *testing.T) {
	groupName := "test-acc-eg-azure-v3-baseline"
	resourceName := createElastigroupAzureV3ResourceName(groupName)

	var group v3.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{groupName: groupName}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "eastus"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "0"),
				),
			},
			{
				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{groupName: groupName, updateBaselineFields: true}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupAzureV3Exists(&group, resourceName),
					testCheckElastigroupAzureV3Attributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "region", "eastus"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "0"),
				),
			},
		},
	})
}

const testBaselineAzureV3GroupConfig_Create = `
resource "` + string(commons.ElastigroupAzureV3ResourceName) + `" "%v" {
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

const testBaselineAzureV3GroupConfig_Update = `
resource "` + string(commons.ElastigroupAzureV3ResourceName) + `" "%v" {
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

// region Azure Elastigroup: Image
//func TestAccSpotinstElastigroupAzure_Image(t *testing.T) {
//	groupName := "test-acc-eg-azure-image"
//	resourceName := createElastigroupAzureResourceName(groupName)
//
//	var group azure.Group
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t, "azure") },
//		Providers:    TestAccProviders,
//		CheckDestroy: testElastigroupAzureDestroy,
//
//		Steps: []resource.TestStep{
//			{
//				Config: createElastigroupAzureTerraform(&AzureGroupConfigMetadata{
//					groupName: groupName,
//					image:     testAzureImageGroupConfig_Create,
//				}),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckElastigroupAzureExists(&group, resourceName),
//					testCheckElastigroupAzureAttributes(&group, groupName),
//					resource.TestCheckResourceAttr(resourceName, "image.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.offer", "UbuntuServer"),
//					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.publisher", "Canonical"),
//					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.sku", "16.04-LTS"),
//				),
//			},
//		},
//	})
//}

const testAzureV3ImageGroupConfig_Create = `
// --- IMAGES --------------------------------
  image {
    marketplace {
      publisher = "Canonical"
      offer = "UbuntuServer"
      sku = "16.04-LTS"
    }
  }
// ---------------------------------------------
`

// endregion

// region Azure Elastigroup: Network
//func TestAccSpotinstElastigroupAzureV3_Network(t *testing.T) {
//	groupName := "test-acc-eg-azure-v3-network"
//	resourceName := createElastigroupAzureV3ResourceName(groupName)
//
//	var group azure.Group
//	resource.Test(t, resource.TestCase{
//		PreCheck:      func() { testAccPreCheck(t, "azure") },
//		Providers:     TestAccProviders,
//		CheckDestroy:  testElastigroupAzureV3Destroy,
//		IDRefreshName: resourceName,
//
//		Steps: []resource.TestStep{
//			{
//				ResourceName: resourceName,
//				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
//					groupName: groupName,
//					network:   testAzureV3NetworkGroupConfig_Create,
//				}),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckElastigroupAzureV3Exists(&group, resourceName),
//					testCheckElastigroupAzureV3Attributes(&group, groupName),
//					resource.TestCheckResourceAttr(resourceName, "network.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "network.0.assign_public_ip", "true"),
//					resource.TestCheckResourceAttr(resourceName, "network.0.resource_group_name", "alex-test"),
//					resource.TestCheckResourceAttr(resourceName, "network.0.subnet_name", "alex-test-subnet"),
//					resource.TestCheckResourceAttr(resourceName, "network.0.virtual_network_name", "alex-test-netwrk"),
//					resource.TestCheckResourceAttr(resourceName, "network.0.additional_ip_configs.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "network.0.additional_ip_configs.0.name", "test"),
//					resource.TestCheckResourceAttr(resourceName, "network.0.additional_ip_configs.0.private_ip_version", "IPV4"),
//				),
//			},
//		},
//	})
//}

const testAzureV3NetworkGroupConfig_Create = `
// --- NETWORK ---------------------------------
  network {
    virtual_network_name = "alex-test-netwrk"
    subnet_name = "alex-test-subnet"                 
    resource_group_name = "alex-test"         
    assign_public_ip = true

    additional_ip_configs {
      name = "test"
      private_ip_version = "IPv4"
    }
  }
// ---------------------------------------------
`

// endregion

// region Azure Elastigroup: Strategy
//func TestAccSpotinstElastigroupAzureV3_Strategy(t *testing.T) {
//	groupName := "test-acc-eg-azure-v3-strategy"
//	resourceName := createElastigroupAzureV3ResourceName(groupName)
//
//	var group v3.Group
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t, "azure") },
//		Providers:    TestAccProviders,
//		CheckDestroy: testElastigroupAzureV3Destroy,
//
//		Steps: []resource.TestStep{
//			{
//				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
//					groupName: groupName,
//					strategy:  testAzureV3StrategyGroupConfig_Create,
//				}),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckElastigroupAzureV3Exists(&group, resourceName),
//					testCheckElastigroupAzureV3Attributes(&group, groupName),
//					resource.TestCheckResourceAttr(resourceName, "strategy.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "strategy.0.low_priority_percentage", "100"),
//					resource.TestCheckResourceAttr(resourceName, "strategy.0.draining_timeout", "180"),
//				),
//			},
//			{
//				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
//					groupName: groupName,
//					strategy:  testAzureV3StrategyGroupConfig_Update,
//				}),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckElastigroupAzureV3Exists(&group, resourceName),
//					testCheckElastigroupAzureV3Attributes(&group, groupName),
//					resource.TestCheckResourceAttr(resourceName, "strategy.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "strategy.0.od_count", "1"),
//					resource.TestCheckResourceAttr(resourceName, "strategy.0.draining_timeout", "240"),
//				),
//			},
//		},
//	})
//}

const testAzureV3StrategyGroupConfig_Create = `
// --- STRATEGY --------------------------------
  strategy {
    low_priority_percentage = 100
    draining_timeout = 180
  }
// ---------------------------------------------
`

//const testAzureV3StrategyGroupConfig_Update = `
//// --- STRATEGY --------------------------------
//  strategy {
//    od_count = 1
//    draining_timeout = 240
//  }
//// ---------------------------------------------
//`

// endregion

// region Azure Elastigroup: VM Sizes
//func TestAccSpotinstElastigroupAzureV3_VMSizes(t *testing.T) {
//	groupName := "test-acc-eg-azure-v3-vm-sizes"
//	resourceName := createElastigroupAzureV3ResourceName(groupName)
//
//	var group v3.Group
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t, "azure") },
//		Providers:    TestAccProviders,
//		CheckDestroy: testElastigroupAzureV3Destroy,
//
//		Steps: []resource.TestStep{
//			{
//				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
//					groupName: groupName,
//					vmSizes:   testAzureV3VMSizesGroupConfig_Create,
//				}),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckElastigroupAzureV3Exists(&group, resourceName),
//					testCheckElastigroupAzureV3Attributes(&group, groupName),
//					resource.TestCheckResourceAttr(resourceName, "od_sizes.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "od_sizes.0", "basic_a1"),
//					resource.TestCheckResourceAttr(resourceName, "low_priority_sizes.#", "3"),
//					resource.TestCheckResourceAttr(resourceName, "low_priority_sizes.0", "basic_a2"),
//					resource.TestCheckResourceAttr(resourceName, "low_priority_sizes.1", "basic_a1"),
//					resource.TestCheckResourceAttr(resourceName, "low_priority_sizes.2", "basic_a3"),
//				),
//			},
//			{
//				Config: createElastigroupAzureV3Terraform(&AzureV3GroupConfigMetadata{
//					groupName: groupName,
//					vmSizes:   testAzureV3MSizesGroupConfig_Update,
//				}),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckElastigroupAzureV3Exists(&group, resourceName),
//					testCheckElastigroupAzureV3Attributes(&group, groupName),
//					resource.TestCheckResourceAttr(resourceName, "od_sizes.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "od_sizes.0", "basic_a2"),
//					resource.TestCheckResourceAttr(resourceName, "low_priority_sizes.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "low_priority_sizes.0", "basic_a2"),
//				),
//			},
//		},
//	})
//}

const testAzureV3VMSizesGroupConfig_Create = `
// --- VM SIZES --------------------------------------------
 od_sizes           = ["basic_a1"]
 low_priority_sizes = ["basic_a2", "basic_a1", "basic_a3"]
// ---------------------------------------------------------
`

//const testAzureV3VMSizesGroupConfig_Update = `
//// --- VM SIZES --------------------------------
// od_sizes           = ["basic_a2"]
// low_priority_sizes = ["basic_a2"]
//// ---------------------------------------------
//`

// endregion
