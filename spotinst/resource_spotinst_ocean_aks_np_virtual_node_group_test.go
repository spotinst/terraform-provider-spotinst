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
  vnet_subnet_ids       = ["/subscriptions/123456-1234-1234-1234-123456789/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/virtualNetworks/ExampleVirtualNetwork/subnets/default"]

  // --- strategy -------------------------------------------------------------

  spot_percentage      = 100
  fallback_to_ondemand = true

  // ---------------------------------------------------------------------------

}

`

// endregion

// region OceanAKSNPVirtualNodeGroup: Headrooms
func TestAccSpotinstOceanAKSNPVirtualNodeGroup_Headrooms(t *testing.T) {
	vngResourceName := "test-aks-vng"
	resourceName := createOceanAKSNPVirtualNodeGroupResource(vngResourceName)

	var virtualNodeGroup azure_np.VirtualNodeGroup
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPVirtualNodeGroupDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPVirtualNodeGroupTerraform(&AKSNPVirtualNodeGroupConfigMetadata{vngResourceName: vngResourceName, updateBaselineFields: true}, testHeadroomsOceanAKSNPVirtualNodeGroup_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPVirtualNodeGroupExists(&virtualNodeGroup, resourceName),
					resource.TestCheckResourceAttr(resourceName, "headrooms.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.0.gpu_per_unit", "0"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.0.num_of_units", "2"),
				),
			},
			{
				Config: createOceanAKSNPVirtualNodeGroupTerraform(&AKSNPVirtualNodeGroupConfigMetadata{vngResourceName: vngResourceName, updateBaselineFields: true}, testHeadroomsOceanAKSNPVirtualNodeGroup_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPVirtualNodeGroupExists(&virtualNodeGroup, resourceName),
					resource.TestCheckResourceAttr(resourceName, "headrooms.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.0.gpu_per_unit", "0"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.0.num_of_units", "2"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.1.cpu_per_unit", "2048"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.1.memory_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.1.gpu_per_unit", "2"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.1.num_of_units", "4"),
				),
			},
			{
				Config: createOceanAKSNPVirtualNodeGroupTerraform(&AKSNPVirtualNodeGroupConfigMetadata{
					vngResourceName:      vngResourceName,
					updateBaselineFields: true,
				}, testHeadroomsOceanAKSNPVirtualNodeGroup_EmptyFields),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPVirtualNodeGroupExists(&virtualNodeGroup, resourceName),
					resource.TestCheckResourceAttr(resourceName, "headrooms.#", "0"),
				),
			},
		},
	})
}

const testHeadroomsOceanAKSNPVirtualNodeGroup_Create = `
resource "` + string(commons.OceanAKSNPVirtualNodeGroupResourceName) + `" "%v" {
  provider = "%v"  

  name  = "testVng"

  ocean_id = "o-751eaa33"

  headrooms {
    cpu_per_unit    = 1024
    memory_per_unit = 512
    gpu_per_unit    = 0
    num_of_units    = 2
  }

}

`

const testHeadroomsOceanAKSNPVirtualNodeGroup_Update = `
resource "` + string(commons.OceanAKSNPVirtualNodeGroupResourceName) + `" "%v" {
  provider = "%v"  

  name  = "testVng"

  ocean_id = "o-751eaa33"

  headrooms {
    cpu_per_unit    = 1024
    memory_per_unit = 512
    gpu_per_unit    = 0
    num_of_units    = 2
  }

  headrooms {
    cpu_per_unit    = 2048
    memory_per_unit = 1024
    gpu_per_unit    = 2
    num_of_units    = 4
  }

}

`

const testHeadroomsOceanAKSNPVirtualNodeGroup_EmptyFields = `
resource "` + string(commons.OceanAKSNPVirtualNodeGroupResourceName) + `" "%v" {
  provider = "%v"  

  name  = "testVng"

  ocean_id = "o-751eaa33"

}
`

// endregion

// region OceanAKSNPVirtualNodeGroup: Taints
func TestAccSpotinstOceanAKSNPVirtualNodeGroup_Taints(t *testing.T) {
	vngResourceName := "test-aks-vng"
	resourceName := createOceanAKSNPVirtualNodeGroupResource(vngResourceName)

	var virtualNodeGroup azure_np.VirtualNodeGroup
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPVirtualNodeGroupDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPVirtualNodeGroupTerraform(&AKSNPVirtualNodeGroupConfigMetadata{vngResourceName: vngResourceName, updateBaselineFields: true}, testTaintsOceanAKSNPVirtualNodeGroup_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPVirtualNodeGroupExists(&virtualNodeGroup, resourceName),
					resource.TestCheckResourceAttr(resourceName, "taints.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.key", "taintKey1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.value", "taintValue1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.effect", "NoExecute"),
				),
			},
			{
				Config: createOceanAKSNPVirtualNodeGroupTerraform(&AKSNPVirtualNodeGroupConfigMetadata{vngResourceName: vngResourceName, updateBaselineFields: true}, testTaintsOceanAKSNPVirtualNodeGroup_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPVirtualNodeGroupExists(&virtualNodeGroup, resourceName),
					resource.TestCheckResourceAttr(resourceName, "taints.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.key", "taintKey1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.value", "taintValue1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.effect", "NoExecute"),
					resource.TestCheckResourceAttr(resourceName, "taints.1.key", "taintKey2"),
					resource.TestCheckResourceAttr(resourceName, "taints.1.value", "taintValue2"),
					resource.TestCheckResourceAttr(resourceName, "taints.1.effect", "NoSchedule"),
				),
			},
			{
				Config: createOceanAKSNPVirtualNodeGroupTerraform(&AKSNPVirtualNodeGroupConfigMetadata{
					vngResourceName:      vngResourceName,
					updateBaselineFields: true,
				}, testTaintsOceanAKSNPVirtualNodeGroup_EmptyFields),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPVirtualNodeGroupExists(&virtualNodeGroup, resourceName),
					resource.TestCheckResourceAttr(resourceName, "taints.#", "0"),
				),
			},
		},
	})
}

const testTaintsOceanAKSNPVirtualNodeGroup_Create = `
resource "` + string(commons.OceanAKSNPVirtualNodeGroupResourceName) + `" "%v" {
  provider = "%v"  

  name  = "testVng"

  ocean_id = "o-751eaa33"

  taints {
    key    = "taintKey1"
    value  = "taintValue1"
    effect = "NoExecute"
  }

}

`

const testTaintsOceanAKSNPVirtualNodeGroup_Update = `
resource "` + string(commons.OceanAKSNPVirtualNodeGroupResourceName) + `" "%v" {
  provider = "%v"  

  name  = "testVng"

  ocean_id = "o-751eaa33"

  taints {
    key    = "taintKey1"
    value  = "taintValue1"
    effect = "NoExecute"
  }

  taints {
    key    = "taintKey2"
    value  = "taintValue2"
    effect = "NoSchedule"
  }

}

`

const testTaintsOceanAKSNPVirtualNodeGroup_EmptyFields = `
resource "` + string(commons.OceanAKSNPVirtualNodeGroupResourceName) + `" "%v" {
  provider = "%v"  

  name  = "testVng"

  ocean_id = "o-751eaa33"

}
`

// endregion

// region OceanAKSNPVirtualNodeGroup: Filters
func TestAccSpotinstOceanAKSNPVirtualNodeGroup_Filters(t *testing.T) {
	vngResourceName := "test-aks-vng"
	resourceName := createOceanAKSNPVirtualNodeGroupResource(vngResourceName)

	var virtualNodeGroup azure_np.VirtualNodeGroup
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPVirtualNodeGroupDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPVirtualNodeGroupTerraform(&AKSNPVirtualNodeGroupConfigMetadata{vngResourceName: vngResourceName, updateBaselineFields: true}, testFiltersOceanAKSNPVirtualNodeGroup_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPVirtualNodeGroupExists(&virtualNodeGroup, resourceName),
					resource.TestCheckResourceAttr(resourceName, "filters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.min_vcpu", "2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.max_vcpu", "16"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.min_memory_gib", "8"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.max_memory_gib", "16"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.architectures.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.architectures.0", "X86_64"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.#", "5"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.0", "A"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.1", "A v2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.2", "D v3"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.3", "Dds_v4"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.4", "Dsv2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.exclude_series.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.exclude_series.0", "E v3"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.exclude_series.1", "Eas_v5"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.exclude_series.2", "Esv3"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.accelerated_networking", "Disabled"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.disk_performance", "Standard"),
					//resource.TestCheckResourceAttr(resourceName, "filters.0.min_gpu", "1"),
					//resource.TestCheckResourceAttr(resourceName, "filters.0.max_gpu", "2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.min_nics", "1"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.vm_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.vm_types.0", "generalPurpose"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.min_disk", "1"),
				),
			},
			{
				Config: createOceanAKSNPVirtualNodeGroupTerraform(&AKSNPVirtualNodeGroupConfigMetadata{vngResourceName: vngResourceName, updateBaselineFields: true}, testFiltersOceanAKSNPVirtualNodeGroup_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPVirtualNodeGroupExists(&virtualNodeGroup, resourceName),
					resource.TestCheckResourceAttr(resourceName, "filters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.min_vcpu", "4"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.max_vcpu", "32"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.min_memory_gib", "4"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.max_memory_gib", "32"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.architectures.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.architectures.0", "AMD64"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.architectures.1", "X86_64"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.accelerated_networking", "Enabled"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.disk_performance", "Premium"),
					//resource.TestCheckResourceAttr(resourceName, "filters.0.min_gpu", "1"),
					//resource.TestCheckResourceAttr(resourceName, "filters.0.max_gpu", "2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.min_nics", "2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.vm_types.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.vm_types.0", "computeOptimized"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.vm_types.1", "generalPurpose"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.min_disk", "2"),
				),
			},
			{
				Config: createOceanAKSNPVirtualNodeGroupTerraform(&AKSNPVirtualNodeGroupConfigMetadata{
					vngResourceName:      vngResourceName,
					updateBaselineFields: true,
				}, testFiltersOceanAKSNPVirtualNodeGroup_EmptyFields),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPVirtualNodeGroupExists(&virtualNodeGroup, resourceName),
					resource.TestCheckResourceAttr(resourceName, "filters.#", "0"),
				),
			},
		},
	})
}

const testFiltersOceanAKSNPVirtualNodeGroup_Create = `
resource "` + string(commons.OceanAKSNPVirtualNodeGroupResourceName) + `" "%v" {
  provider = "%v"  

  name  = "testVng"

  ocean_id = "o-751eaa33"

  filters {
    min_vcpu               = 2
    max_vcpu               = 16
    min_memory_gib         = 8
    max_memory_gib         = 16
    architectures          = ["X86_64"]
    series                 = ["D v3", "Dds_v4", "Dsv2", "A", "A v2"]
    exclude_series         = ["E v3", "Esv3", "Eas_v5"]
    accelerated_networking = "Disabled"
    disk_performance       = "Standard"
    //min_gpu                = 1
    //max_gpu                = 2
    min_nics               = 1
    vm_types               = ["generalPurpose"]
    min_disk               = 1
  }

}

`

const testFiltersOceanAKSNPVirtualNodeGroup_Update = `
resource "` + string(commons.OceanAKSNPVirtualNodeGroupResourceName) + `" "%v" {
  provider = "%v"  

  name  = "testVng"

  ocean_id = "o-751eaa33"

  filters {
    min_vcpu               = 4
    max_vcpu               = 32
    min_memory_gib         = 4
    max_memory_gib         = 32
    architectures          = ["X86_64","AMD64"]
    accelerated_networking = "Enabled"
    disk_performance       = "Premium"
    //min_gpu                = 1
    //max_gpu                = 2
    min_nics               = 2
    vm_types               = ["generalPurpose","computeOptimized"]
    min_disk               = 2
  }

}

`

const testFiltersOceanAKSNPVirtualNodeGroup_EmptyFields = `
resource "` + string(commons.OceanAKSNPVirtualNodeGroupResourceName) + `" "%v" {
  provider = "%v"  

  name  = "testVng"

  ocean_id = "o-751eaa33"

}
`

// endregion
