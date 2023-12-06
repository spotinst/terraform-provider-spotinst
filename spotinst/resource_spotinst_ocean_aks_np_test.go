package spotinst

import (
	"context"
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func init() {
	resource.AddTestSweepers("resource_spotinst_ocean_aks_np_import", &resource.Sweeper{
		Name: "resource_spotinst_ocean_aks",
		F:    testSweepOceanAKSNPCluster,
	})
}

func testSweepOceanAKSNPCluster(region string) error {
	client, err := getProviderClient("azure")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).ocean.CloudProviderAzure()

	if resp, err := conn.ListClusters(context.Background()); err != nil {
		return fmt.Errorf("error getting list of clusters to sweep")
	} else {
		if len(resp.Clusters) == 0 {
			log.Printf("[INFO] No clusters to sweep")
		}
		for _, cluster := range resp.Clusters {
			if strings.Contains(spotinst.StringValue(cluster.Name), "terraform-acc-tests-") {
				if _, err := conn.DeleteCluster(context.Background(), &azure.DeleteClusterInput{ClusterID: cluster.ID}); err != nil {
					return fmt.Errorf("unable to delete group %v in sweep", spotinst.StringValue(cluster.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(cluster.ID))
				}
			}
		}
	}
	return nil
}

func createOceanAKSNPResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanAKSNPResourceName), name)
}

func testOceanAKSNPDestroy(s *terraform.State) error {
	client := testAccProviderAzure.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OceanAKSNPResourceName) {
			continue
		}
		input := &azure.ReadClusterInput{ClusterID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAzure().ReadCluster(context.Background(), input)
		if err == nil && resp != nil && resp.Cluster != nil {
			return fmt.Errorf("cluster still exists")
		}
	}
	return nil
}

func testCheckOceanAKSNPAttributes(cluster *azure_np.Cluster, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(cluster.Name) != expectedName {
			return fmt.Errorf("bad content: %v", cluster.Name)
		}
		return nil
	}
}

func testCheckOceanAKSNPExists(cluster *azure_np.Cluster, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAzure.Meta().(*Client)
		input := &azure_np.ReadClusterInput{ClusterID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAzureNP().ReadCluster(context.Background(), input)
		if err != nil {
			return err
		}
		*cluster = *resp.Cluster
		return nil
	}
}

type OceanAKSNPMetadata struct {
	clusterName          string
	controllerClusterID  string
	provider             string
	launchSpecification  string
	scheduling           string
	autoScaler           string
	health               string
	vmSizes              string
	osDisk               string
	image                string
	taints               string
	headrooms            string
	extensions           string
	filters              string
	variables            string
	updateBaselineFields bool
}

func createOceanAKSNPTerraform(clusterMeta *OceanAKSNPMetadata) string {
	if clusterMeta == nil {
		return ""
	}

	if clusterMeta.provider == "" {
		clusterMeta.provider = "azure_np"
	}

	if clusterMeta.scheduling == "" {
		clusterMeta.scheduling = testSchedulingOceanAKSNPConfig_Create
	}

	if clusterMeta.autoScaler == "" {
		clusterMeta.autoScaler = testAutoScalerOceanAKSNPConfig_Create
	}

	if clusterMeta.health == "" {
		clusterMeta.health = testHealthOceanAKSNPConfig_Create
	}

	if clusterMeta.headrooms == "" {
		clusterMeta.headrooms = testHeadroomsOceanAKSNPConfig_Create
	}

	if clusterMeta.taints == "" {
		clusterMeta.taints = testTaintsOceanAKSNPConfig_Create
	}

	if clusterMeta.filters == "" {
		clusterMeta.filters = testFiltersOceanAKSNPConfig_Create
	}

	template :=
		`provider "azure" {
	token   = "fake"
	account = "fake"
	}
	`
	if clusterMeta.updateBaselineFields {
		format := testBaselineOceanAKSNPConfig_Update
		template += fmt.Sprintf(format,
			clusterMeta.clusterName,
			clusterMeta.provider,
			clusterMeta.controllerClusterID,
			clusterMeta.launchSpecification,
			clusterMeta.osDisk,
			clusterMeta.scheduling,
			clusterMeta.image,
			clusterMeta.autoScaler,
			clusterMeta.filters,
			clusterMeta.headrooms,
			clusterMeta.health,
			clusterMeta.taints,
			clusterMeta.vmSizes,
		)
	} else {
		format := testBaselineOceanAKSNPConfig_Create
		template += fmt.Sprintf(format,
			clusterMeta.clusterName,
			clusterMeta.provider,
			clusterMeta.controllerClusterID,
			clusterMeta.launchSpecification,
			clusterMeta.osDisk,
			clusterMeta.scheduling,
			clusterMeta.image,
			clusterMeta.autoScaler,
			clusterMeta.filters,
			clusterMeta.headrooms,
			clusterMeta.health,
			clusterMeta.taints,
			clusterMeta.vmSizes,
		)

	}

	if clusterMeta.variables != "" {
		template = clusterMeta.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", clusterMeta.clusterName, template)
	return template
}

// region Ocean AKS : Baseline
func TestAccSpotinstOceanAKSNP_Baseline(t *testing.T) {
	clusterName := "Terraform-Test-AKS-2-0-Cluster"
	controllerClusterID := "terraform-aks-2-0-cluster"
	resourceName := createOceanAKSNPResourceName(clusterName)

	var cluster azure_np.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure_np") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "name", "Terraform-Test-AKS-2-0-Do-Not-Delete"),
					resource.TestCheckResourceAttr(resourceName, "controller_cluster_id", "terraform-aks-2-0-cluster"),
					resource.TestCheckResourceAttr(resourceName, "aks_name", "Terraform-Test-AKS-2-0-Do-Not-Delete"),
					resource.TestCheckResourceAttr(resourceName, "aks_region", "eastus"),
					resource.TestCheckResourceAttr(resourceName, "aks_resource_group_name", "AutomationResourceGroup"),
					resource.TestCheckResourceAttr(resourceName, "aks_infrastructure_resource_group_name", "MC_AutomationResourceGroup_Terraform-Test-AKS-2-0-Do-Not-Delete_eastus"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.0", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.1", "2"),
				),
			},
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:          clusterName,
					controllerClusterID:  controllerClusterID,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Terraform-Test-AKS-2-0-Do-Not-Delete-Updated"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.0", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.1", "2"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.2", "3"),
				),
			},
		},
	})
}

const testBaselineOceanAKSNPConfig_Create = `
resource "` + string(commons.OceanAKSResourceName) + `" "%v" {
  provider = "%v"
  name                  = "Terraform-Test-AKS-2-0-Do-Not-Delete"
  controller_cluster_id = "%v"

  // --- AKS -----------------------------------------------------------
  aks_name                = "Terraform-Test-AKS-2-0-Do-Not-Delete"
  aks_resource_group_name = "AutomationResourceGroup"
  aks_region                = "Terraform-Test-AKS-2-0-Do-Not-Delete"
  aks_infrastructure_resource_group_name = "MC_AutomationResourceGroup_Terraform-Test-AKS-2-0-Do-Not-Delete_eastus"
  // -------------------------------------------------------------------

  availability_zones = [
    "1",
    "2"
  ]

  // --- nodeCountLimits --------------------------------------------------

  min_count = 1
  max_count = 100

  // ----------------------------------------------------------------------

  // --- nodePoolProperties -----------------------------------------------

  max_pods_per_node     = 30
  enable_node_public_ip = true
  os_disk_size_gb       = 32
  os_disk_type          = "Managed"
  os_type               = "Linux"
  os_sku                = "Ubuntu"
  kubernetes_version    = "1.26"
  //pod_subnet_ids      = ["/subscriptions/123456-1234-1234-1234-123456789/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/virtualNetworks/ExampleVirtualNetwork/subnets/default"]
  vnet_subnet_ids       = ["/subscriptions/a9e813ad-f18b-4ad2-9dbc-5c6df28e9cb8/resourceGroups/AutomationResourceGroup/providers/Microsoft.Network/virtualNetworks/Automation-VirtualNetwork/subnets/default"]
  // ----------------------------------------------------------------------

  // --- strategy ---------------------------------------------------------

  spot_percentage      = 50
  fallback_to_ondemand = true

  // ----------------------------------------------------------------------

%v
%v
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

const testBaselineOceanAKSNPConfig_Update = `
resource "` + string(commons.OceanAKSResourceName) + `" "%v" {
  
  provider = "%v"
  name                  = "Terraform-Test-AKS-2-0-Do-Not-Delete-Updated"
  controller_cluster_id = "%v"
  // --- AKS -----------------------------------------------------------
  aks_name                = "Terraform-Test-AKS-2-0-Do-Not-Delete"
  aks_resource_group_name = "AutomationResourceGroup"
  aks_region                = "Terraform-Test-AKS-2-0-Do-Not-Delete"
  aks_infrastructure_resource_group_name = "MC_AutomationResourceGroup_Terraform-Test-AKS-2-0-Do-Not-Delete_eastus"
  // -------------------------------------------------------------------

  availability_zones = [
    "1",
    "2",
	"3"
  ]

%v
%v
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

//endregion

// region Ocean AKS : AutoScaler
func TestAccSpotinstOceanAKSNP_AutoScaler(t *testing.T) {
	clusterName := "Terraform-Test-AKS-2-0-Cluster"
	controllerClusterID := "terraform-aks-2-0-cluster"
	resourceName := createOceanAKSNPResourceName(clusterName)

	var cluster azure_np.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure_np") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.0.max_scale_down_percentage", "10"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.automatic.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.automatic.0.percentage", "10"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_memory_gib", "40"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "1024"),
				),
			},
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:          clusterName,
					controllerClusterID:  controllerClusterID,
					autoScaler:           testAutoScalerOceanAKSNPConfig_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "autoscaler.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.0.max_scale_down_percentage", "50"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.automatic.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.automatic.0.percentage", "60"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_memory_gib", "80"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "2048"),
				),
			},
		},
	})
}

const testAutoScalerOceanAKSNPConfig_Create = `
    // --- AutoScaler ----------------------------------------------------
    autoscaler {
      autoscale_is_enabled = true

      autoscale_down {
        max_scale_down_percentage = 10
      }

      resource_limits {
        max_vcpu = 1024
        max_memory_gib = 40
      }

      autoscale_headroom {
        automatic {
          percentage = 10
        }
      }
    }
    // -------------------------------------------------------------------
`

const testAutoScalerOceanAKSNPConfig_Update = `
    // --- AutoScaler ----------------------------------------------------
    autoscaler {
      autoscale_is_enabled = false

      autoscale_down {
        max_scale_down_percentage = 50
      }

      resource_limits {
        max_vcpu = 2048
        max_memory_gib = 80
      }

      autoscale_headroom {
        automatic {
          percentage = 60
        }
      }
    }
    // -------------------------------------------------------------------
`

//endregion

// region Ocean AKS : Scheduling
func TestAccSpotinstOceanAKSNP_Scheduling(t *testing.T) {
	clusterName := "Terraform-Test-AKS-2-0-Cluster"
	controllerClusterID := "terraform-aks-2-0-cluster"
	resourceName := createOceanAKSNPResourceName(clusterName)

	var cluster azure_np.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure_np") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "scheduling.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.shutdown_hours.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.shutdown_hours.0.is_enabled", "false"),
				),
			},
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:          clusterName,
					controllerClusterID:  controllerClusterID,
					scheduling:           testSchedulingOceanAKSNPConfig_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "scheduling.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.shutdown_hours.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.shutdown_hours.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.shutdown_hours.0.time_windows.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.shutdown_hours.0.time_windows.0", "Sat:08:00-Sun:08:00"),
				),
			},
		},
	})
}

const testSchedulingOceanAKSNPConfig_Create = `
  // --- Scheduling ------------------------------------------------------
  scheduling {
    shutdown_hours{
      is_enabled   = false
    }
  }
  // -------------------------------------------------------------------
`

const testSchedulingOceanAKSNPConfig_Update = `
  // --- Scheduling ------------------------------------------------------
  scheduling {
    shutdown_hours{
      is_enabled   = true
      time_windows = ["Sat:08:00-Sun:08:00"]
    }
  }
  // -------------------------------------------------------------------
`

//endregion

// region Ocean AKS : Health
func TestAccSpotinstOceanAKSNP_Health(t *testing.T) {
	clusterName := "Terraform-Test-AKS-2-0-Cluster"
	controllerClusterID := "terraform-aks-2-0-cluster"
	resourceName := createOceanAKSNPResourceName(clusterName)

	var cluster azure_np.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure_np") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "health.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health.0.grace_period", "10"),
				),
			},
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:          clusterName,
					controllerClusterID:  controllerClusterID,
					health:               testHealthOceanAKSNPConfig_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "health.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health.0.grace_period", "60"),
				),
			},
		},
	})
}

const testHealthOceanAKSNPConfig_Create = `
  // --- Health --------------------------------------------------------
  health {
    grace_period = 10
  }
  // -------------------------------------------------------------------
`

const testHealthOceanAKSNPConfig_Update = `
  // --- Health --------------------------------------------------------
  health {
    grace_period = 60
  }
  // -------------------------------------------------------------------
`

//endregion

// region Ocean AKS : Headrooms
func TestAccSpotinstOceanAKSNP_Headrooms(t *testing.T) {
	clusterName := "Terraform-Test-AKS-2-0-Cluster"
	controllerClusterID := "terraform-aks-2-0-cluster"
	resourceName := createOceanAKSNPResourceName(clusterName)

	var cluster azure_np.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure_np") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "headrooms.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.0.gpu_per_unit", "0"),
					resource.TestCheckResourceAttr(resourceName, "headrooms.0.num_of_units", "2"),
				),
			},
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:          clusterName,
					controllerClusterID:  controllerClusterID,
					headrooms:            testHeadroomsOceanAKSNPConfig_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
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
		},
	})
}

const testHeadroomsOceanAKSNPConfig_Create = `
  // --- autoscale --------------------------------------------------------
  headrooms {
    cpu_per_unit    = 1024
    memory_per_unit = 512
    gpu_per_unit    = 0
    num_of_units    = 2
  }
// ----------------------------------------------------------------------

`
const testHeadroomsOceanAKSNPConfig_Update = `
  // --- autoscale --------------------------------------------------------
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
// ----------------------------------------------------------------------

`

//endregion

// region Ocean AKS : Taints
func TestAccSpotinstOceanAKSNP_Taints(t *testing.T) {
	clusterName := "Terraform-Test-AKS-2-0-Cluster"
	controllerClusterID := "terraform-aks-2-0-cluster"
	resourceName := createOceanAKSNPResourceName(clusterName)

	var cluster azure_np.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure_np") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "taints.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.key", "taintKey1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.value", "taintValue1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.effect", "NoSchedule"),
				),
			},
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:          clusterName,
					controllerClusterID:  controllerClusterID,
					taints:               testTaintsOceanAKSNPConfig_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "taints.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.key", "taintKey1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.value", "taintValue1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.effect", "NoSchedule"),
					resource.TestCheckResourceAttr(resourceName, "taints.1.key", "taintKey2"),
					resource.TestCheckResourceAttr(resourceName, "taints.1.value", "taintValue2"),
					resource.TestCheckResourceAttr(resourceName, "taints.1.effect", "NoExecute"),
				),
			},
		},
	})
}

const testTaintsOceanAKSNPConfig_Create = `
    
  taints {
    key    = "taintKey"
    value  = "taintValue"
    effect = "NoSchedule"
  }
`
const testTaintsOceanAKSNPConfig_Update = `
  
  taints {
    key    = "taintKey1"
    value  = "taintValue1"
    effect = "NoSchedule"
  }

  taints {
    key    = "taintKey2"
    value  = "taintValue2"
    effect = "NoExecute"
  }

`

//endregion

// region Ocean AKS : Filters
func TestAccSpotinstOceanAKSNP_Filters(t *testing.T) {
	clusterName := "Terraform-Test-AKS-2-0-Cluster"
	controllerClusterID := "terraform-aks-2-0-cluster"
	resourceName := createOceanAKSNPResourceName(clusterName)

	var cluster azure_np.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure_np") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "filters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.min_vcpu", "2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.max_vcpu", "16"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.min_memory_gib", "8"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.max_memory_gib", "16"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.architectures.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.architectures.0", "X86_64"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.#", "5"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.0", "D v3"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.1", "Dds_v4"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.2", "Dsv2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.3", "A"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.4", "A v2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.exclude_series.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.exclude_series.0", "E v3"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.exclude_series.1", "Esv3"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.exclude_series.2", "Eas_v5"),
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
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:          clusterName,
					controllerClusterID:  controllerClusterID,
					filters:              testFiltersOceanAKSNPConfig_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "filters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.min_vcpu", "2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.max_vcpu", "16"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.min_memory_gib", "8"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.max_memory_gib", "16"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.architectures.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.architectures.0", "X86_64"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.#", "5"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.0", "D v3"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.1", "Dds_v4"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.2", "Dsv2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.3", "A"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.series.4", "A v2"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.exclude_series.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.exclude_series.0", "E v3"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.exclude_series.1", "Esv3"),
					resource.TestCheckResourceAttr(resourceName, "filters.0.exclude_series.2", "Eas_v5"),
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
		},
	})
}

const testFiltersOceanAKSNPConfig_Create = `
  // --- vmSizes ----------------------------------------------------------

  filters {
    min_vcpu               = 2
    max_vcpu               = 16
    min_memory_gib         = 8
    max_memory_gib         = 16
    architectures          = ["x86_64"]
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

  // ----------------------------------------------------------------------
`

const testFiltersOceanAKSNPConfig_Update = `
  // --- vmSizes ----------------------------------------------------------

  filters {
    min_vcpu               = 2
    max_vcpu               = 16
    min_memory_gib         = 8
    max_memory_gib         = 16
    architectures          = ["x86_64"]
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

  // ----------------------------------------------------------------------
`

//endregion
