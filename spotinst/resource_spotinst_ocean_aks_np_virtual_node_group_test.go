package spotinst

import (
	"context"
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

//func init() {
//	resource.AddTestSweepers("spotinst_ocean_gke_launch_spec", &resource.Sweeper{
//		Name: "spotinst_ocean_gke_launch_spec",
//		F:    testSweepOceanGKELaunchSpec,
//	})
//}
//
//func testSweepOceanGKELaunchSpec(region string) error {
//	client, err := getProviderClient("gcp")
//	if err != nil {
//		return fmt.Errorf("error getting client: %v", err)
//	}
//
//	conn := client.(*Client).ocean.CloudProviderGCP()
//	input := &gcp.ListLaunchSpecsInput{}
//	if resp, err := conn.ListLaunchSpecs(context.Background(), input); err != nil {
//		return fmt.Errorf("error getting list of launch specs to sweep")
//	} else {
//		if len(resp.LaunchSpecs) == 0 {
//			log.Printf("[INFO] No launch specs to sweep")
//		}
//		for _, launchSpec := range resp.LaunchSpecs {
//			if strings.Contains(spotinst.StringValue(launchSpec.<WHAT>), "test-acc-") {
//				if _, err := conn.DeleteLaunchSpec(context.Background(), &gcp.DeleteLaunchSpecInput{LaunchSpecID: launchSpec.ID}); err != nil {
//					return fmt.Errorf("unable to delete launch spec %v in sweep", spotinst.StringValue(launchSpec.ID))
//				} else {
//					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(launchSpec.ID))
//				}
//			}
//		}
//	}
//	return nil
//}

func createOceanAKSNPVirtualNodeGroupResource(oceanID string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanAKSNPVirtualNodeGroupResourceName), oceanID)
}

func testOceanAKSNPVirtualNodeGroupDestroy(s *terraform.State) error {
	client := testAccProviderAzure.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OceanAKSNPVirtualNodeGroupResourceName) {
			continue
		}

		input := &azure_np.ReadVirtualNodeGroupInput{VirtualNodeGroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAzureNP().ReadVirtualNodeGroup(context.Background(), input)

		if err == nil && resp != nil && resp.VirtualNodeGroup != nil {
			return fmt.Errorf("Virtual Node Group still exists")
		}
	}

	return nil
}

func testCheckOceanAKSNPVirtualNodeGroupAttributes(virtualNodeGroup *azure_np.VirtualNodeGroup, expectedID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(virtualNodeGroup.OceanID) != expectedID {
			return fmt.Errorf("bad content: %v", virtualNodeGroup.OceanID)
		}
		return nil
	}
}

func testCheckOceanAKSNPVirtualNodeGroupExists(virtualNodeGroup *azure_np.VirtualNodeGroup, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAzure.Meta().(*Client)
		input := &azure_np.ReadVirtualNodeGroupInput{VirtualNodeGroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAzureNP().ReadVirtualNodeGroup(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.VirtualNodeGroup.ID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("Virtual Node Group not found: %+v,\n %+v\n", resp.VirtualNodeGroup, rs.Primary.Attributes)
		}
		*virtualNodeGroup = *resp.VirtualNodeGroup
		return nil
	}
}

type AKSNPVirtualNodeGroupConfigMetadata struct {
	provider             string
	vngResourceName      string
	updateBaselineFields bool
}

func createOceanAKSNPVirtualNodeGroupTerraform(vngcm *AKSNPVirtualNodeGroupConfigMetadata, formatToUse string) string {
	if vngcm == nil {
		return ""
	}

	if vngcm.provider == "" {
		vngcm.provider = "azure"
	}

	template :=
		`provider "azure" {
	 token   = "fake"
	 account = "fake"
	}
	`

	format := formatToUse

	if vngcm.updateBaselineFields {
		template += fmt.Sprintf(format,
			vngcm.vngResourceName,
			vngcm.provider,
		)
	} else {
		template += fmt.Sprintf(format,
			vngcm.vngResourceName,
			vngcm.provider,
		)
	}

	log.Printf("Terraform LaunchSpec template:\n%v", template)

	return template
}

// region OceanAKSNPVirtualNodeGroup: Baseline
func TestAccSpotinstOceanAKSNPVirtualNodeGroup_Baseline(t *testing.T) {
	vngResourceName := "test-aks-vng"
	resourceName := createOceanAKSNPVirtualNodeGroupResource(vngResourceName)

	var virtualNodeGroup azure_np.VirtualNodeGroup
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPVirtualNodeGroupDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPVirtualNodeGroupTerraform(&AKSNPVirtualNodeGroupConfigMetadata{vngResourceName: vngResourceName, updateBaselineFields: true}, testBaselineOceanAKSNPVirtualNodeGroup_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPVirtualNodeGroupExists(&virtualNodeGroup, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "testVng"),
					resource.TestCheckResourceAttr(resourceName, "ocean_id", "o-751eaa33"),
					resource.TestCheckResourceAttr(resourceName, "min_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "max_count", "100"),
					resource.TestCheckResourceAttr(resourceName, "max_pods_per_node", "30"),
					resource.TestCheckResourceAttr(resourceName, "enable_node_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "os_disk_size_gb", "32"),
					resource.TestCheckResourceAttr(resourceName, "os_disk_type", "Managed"),
					resource.TestCheckResourceAttr(resourceName, "os_type", "Linux"),
					resource.TestCheckResourceAttr(resourceName, "os_sku", "Ubuntu"),
					resource.TestCheckResourceAttr(resourceName, "kubernetes_version", "1.26"),
					resource.TestCheckResourceAttr(resourceName, "spot_percentage", "50"),
					resource.TestCheckResourceAttr(resourceName, "fallback_to_ondemand", "false"),
					//resource.TestCheckResourceAttr(resourceName, "pod_subnet_ids.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "pod_subnet_ids.0", "/subscriptions/123456-1234-1234-1234-123456789/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/virtualNetworks/ExampleVirtualNetwork/subnets/default"),
					resource.TestCheckResourceAttr(resourceName, "vnet_subnet_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vnet_subnet_ids.0", "/subscriptions/a9e813ad-f18b-4ad2-9dbc-5c6df28e9cb8/resourceGroups/AutomationResourceGroup/providers/Microsoft.Network/virtualNetworks/Automation-VirtualNetwork/subnets/default"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.0", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.1", "2"),
				),
			},
			{
				Config: createOceanAKSNPVirtualNodeGroupTerraform(&AKSNPVirtualNodeGroupConfigMetadata{vngResourceName: vngResourceName, updateBaselineFields: true}, testBaselineOceanAKSNPVirtualNodeGroup_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPVirtualNodeGroupExists(&virtualNodeGroup, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "testVngUpdated"),
					resource.TestCheckResourceAttr(resourceName, "min_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "max_count", "150"),
					resource.TestCheckResourceAttr(resourceName, "max_pods_per_node", "50"),
					resource.TestCheckResourceAttr(resourceName, "enable_node_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "os_disk_size_gb", "64"),
					resource.TestCheckResourceAttr(resourceName, "kubernetes_version", "1.27"),
					resource.TestCheckResourceAttr(resourceName, "spot_percentage", "100"),
					resource.TestCheckResourceAttr(resourceName, "fallback_to_ondemand", "true"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.0", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.1", "2"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.2", "3"),
				),
			},
		},
	})
}

const testBaselineOceanAKSNPVirtualNodeGroup_Create = `
resource "` + string(commons.OceanAKSNPVirtualNodeGroupResourceName) + `" "%v" {
  provider = "%v"  

  name  = "testVng"

  ocean_id = "o-751eaa33"

  availability_zones = [
    "1",
    "2"
  ]

  // --- nodeCountLimits ----------------------------------------------------

  min_count = 1
  max_count = 100

  // --- nodePoolProperties --------------------------------------------------

  max_pods_per_node     = 30
  enable_node_public_ip = true
  os_disk_size_gb       = 32
  os_disk_type          = "Managed"
  os_type               = "Linux"
  os_sku                = "Ubuntu"
  kubernetes_version    = "1.26"
  //pod_subnet_ids       = ["/subscriptions/123456-1234-1234-1234-123456789/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/virtualNetworks/ExampleVirtualNetwork/subnets/default"]
  vnet_subnet_ids       = ["/subscriptions/a9e813ad-f18b-4ad2-9dbc-5c6df28e9cb8/resourceGroups/AutomationResourceGroup/providers/Microsoft.Network/virtualNetworks/Automation-VirtualNetwork/subnets/default"]

  // --- strategy -------------------------------------------------------------

  spot_percentage      = 50
  fallback_to_ondemand = false

  // ---------------------------------------------------------------------------
 
}

`

const testBaselineOceanAKSNPVirtualNodeGroup_Update = `
resource "` + string(commons.OceanAKSNPVirtualNodeGroupResourceName) + `" "%v" {
  provider = "%v"  

  name  = "testVngUpdated"

  ocean_id = "o-751eaa33"

  availability_zones = [
    "1",
    "2",
    "3"
  ]

  // --- nodeCountLimits ----------------------------------------------------

  min_count = 2
  max_count = 150

  // --- nodePoolProperties --------------------------------------------------

  max_pods_per_node     = 50
  enable_node_public_ip = false
  os_disk_size_gb       = 64
  os_disk_type          = "Managed"
  os_type               = "Linux"
  os_sku                = "Ubuntu"
  kubernetes_version    = "1.27"
  //pod_subnet_ids       = ["/subscriptions/123456-1234-1234-1234-123456789/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/virtualNetworks/ExampleVirtualNetwork/subnets/default"]
  vnet_subnet_ids       = ["/subscriptions/123456-1234-1234-1234-123456789/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/virtualNetworks/ExampleVirtualNetwork/subnets/default"]

  // --- strategy -------------------------------------------------------------

  spot_percentage      = 100
  fallback_to_ondemand = true

  // ---------------------------------------------------------------------------

}

`

// endregion
