package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func init() {
	resource.AddTestSweepers("resource_spotinst_ocean_aks_np_import", &resource.Sweeper{
		Name: "resource_spotinst_ocean_aks_np",
		F:    testSweepOceanAKSNPCluster,
	})
}

func testSweepOceanAKSNPCluster(region string) error {
	client, err := getProviderClient("azure")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).ocean.CloudProviderAzureNP()

	if resp, err := conn.ListClusters(context.Background()); err != nil {
		return fmt.Errorf("error getting list of clusters to sweep")
	} else {
		if len(resp.Clusters) == 0 {
			log.Printf("[INFO] No clusters to sweep")
		}
		for _, cluster := range resp.Clusters {
			if strings.Contains(spotinst.StringValue(cluster.Name), "Terraform-Test-AKS-2-0-Cluster") {
				if _, err := conn.DeleteCluster(context.Background(), &azure_np.DeleteClusterInput{ClusterID: cluster.ID}); err != nil {
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
		input := &azure_np.ReadClusterInput{ClusterID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAzureNP().ReadCluster(context.Background(), input)
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
	autoScaler           string
	health               string
	scheduling           string
	taints               string
	headrooms            string
	filters              string
	variables            string
	updateBaselineFields bool
}

func createOceanAKSNPTerraform(clusterMeta *OceanAKSNPMetadata) string {
	if clusterMeta == nil {
		return ""
	}

	if clusterMeta.provider == "" {
		clusterMeta.provider = "azure"
	}

	if clusterMeta.autoScaler == "" {
		clusterMeta.autoScaler = testAutoScalerOceanAKSNPConfig_Create
	}

	if clusterMeta.health == "" {
		clusterMeta.health = testHealthOceanAKSNPConfig_Create
	}

	if clusterMeta.scheduling == "" {
		clusterMeta.scheduling = testSchedulingOceanAKSNPConfig_Create
	}

	if clusterMeta.taints == "" {
		clusterMeta.taints = testTaintsOceanAKSNPConfig_Create
	}

	if clusterMeta.headrooms == "" {
		clusterMeta.headrooms = testHeadroomsOceanAKSNPConfig_Create
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
			clusterMeta.autoScaler,
			clusterMeta.health,
			clusterMeta.scheduling,
			clusterMeta.taints,
			clusterMeta.headrooms,
			clusterMeta.filters,
		)
	} else {
		format := testBaselineOceanAKSNPConfig_Create
		template += fmt.Sprintf(format,
			clusterMeta.clusterName,
			clusterMeta.provider,
			clusterMeta.controllerClusterID,
			clusterMeta.autoScaler,
			clusterMeta.health,
			clusterMeta.scheduling,
			clusterMeta.taints,
			clusterMeta.headrooms,
			clusterMeta.filters,
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
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				//ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "name", "Terraform-Test-AKS-2-0-Cluster"),
					resource.TestCheckResourceAttr(resourceName, "controller_cluster_id", "terraform-aks-2-0-cluster"),
					resource.TestCheckResourceAttr(resourceName, "aks_cluster_name", "Terraform-Test-AKS-2-0-Do-Not-Delete"),
					resource.TestCheckResourceAttr(resourceName, "aks_region", "eastus"),
					resource.TestCheckResourceAttr(resourceName, "aks_resource_group_name", "AutomationResourceGroup"),
					resource.TestCheckResourceAttr(resourceName, "aks_infrastructure_resource_group_name", "MC_AutomationResourceGroup_Terraform-Test-AKS-2-0-Do-Not-Delete_eastus"),
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
					//resource.TestCheckResourceAttr(resourceName, "vnet_subnet_ids.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "vnet_subnet_ids.0", "/subscriptions/a9e813ad-f18b-4ad2-9dbc-5c6df28e9cb8/resourceGroups/AutomationResourceGroup/providers/Microsoft.Network/virtualNetworks/Automation-VirtualNetwork/subnets/default"),
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
				//ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Terraform-Test-AKS-2-0-Cluster-Updated"),
					resource.TestCheckResourceAttr(resourceName, "min_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "max_count", "150"),
					resource.TestCheckResourceAttr(resourceName, "max_pods_per_node", "50"),
					resource.TestCheckResourceAttr(resourceName, "enable_node_public_ip", "true"),
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

const testBaselineOceanAKSNPConfig_Create = `
resource "` + string(commons.OceanAKSNPResourceName) + `" "%v" {
  provider = "%v"
  name                  = "Terraform-Test-AKS-2-0-Cluster"
  controller_cluster_id = "%v"

  // --- AKS -----------------------------------------------------------
  
  aks_cluster_name                = "Terraform-Test-AKS-2-0-Do-Not-Delete"
  aks_resource_group_name = "AutomationResourceGroup"
  aks_region                = "eastus"
  aks_infrastructure_resource_group_name = "MC_AutomationResourceGroup_Terraform-Test-AKS-2-0-Do-Not-Delete_eastus"
  // -------------------------------------------------------------------

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
  //vnet_subnet_ids       = ["/subscriptions/a9e813ad-f18b-4ad2-9dbc-5c6df28e9cb8/resourceGroups/AutomationResourceGroup/providers/Microsoft.Network/virtualNetworks/Automation-VirtualNetwork/subnets/default"]
  // ----------------------------------------------------------------------

  // --- strategy ---------------------------------------------------------

  spot_percentage      = 50
  fallback_to_ondemand = false

  // ----------------------------------------------------------------------

  availability_zones = [
    "1",
    "2"
  ]

%v
%v
%v
%v
%v
%v
}
`

const testBaselineOceanAKSNPConfig_Update = `
resource "` + string(commons.OceanAKSNPResourceName) + `" "%v" {
  
  provider = "%v"
  name                  = "Terraform-Test-AKS-2-0-Cluster-Updated"
  controller_cluster_id = "%v"
  // --- AKS -----------------------------------------------------------
  
  aks_cluster_name                = "Terraform-Test-AKS-2-0-Do-Not-Delete"
  aks_resource_group_name = "AutomationResourceGroup"
  aks_region                = "eastus"
  aks_infrastructure_resource_group_name = "MC_AutomationResourceGroup_Terraform-Test-AKS-2-0-Do-Not-Delete_eastus"
  // -------------------------------------------------------------------

  // --- nodeCountLimits --------------------------------------------------

  min_count = 2
  max_count = 150

  // ----------------------------------------------------------------------

  // --- nodePoolProperties -----------------------------------------------

  max_pods_per_node     = 50
  enable_node_public_ip = true
  os_disk_size_gb       = 64
  os_disk_type          = "Managed"
  os_type               = "Linux"
  os_sku                = "Ubuntu"
  kubernetes_version    = "1.27"
  //vnet_subnet_ids       = ["/subscriptions/a9e813ad-f18b-4ad2-9dbc-5c6df28e9cb8/resourceGroups/AutomationResourceGroup/providers/Microsoft.Network/virtualNetworks/Automation-VirtualNetwork/subnets/default"]
  // ----------------------------------------------------------------------

  // --- strategy ---------------------------------------------------------

  spot_percentage      = 100
  fallback_to_ondemand = true

  // ----------------------------------------------------------------------

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
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				//ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.0.max_scale_down_percentage", "10"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.automatic.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.automatic.0.percentage", "10"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.automatic.0.is_enabled", "true"),
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
				//ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "autoscaler.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.0.max_scale_down_percentage", "50"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.automatic.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.automatic.0.percentage", "60"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.automatic.0.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_memory_gib", "80"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "2048"),
				),
			},
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					autoScaler:          testAutoScalerOceanAKSNPConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.#", "0"),
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
		  is_enabled = true
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
		  is_enabled = false
        }
      }
    }
    // -------------------------------------------------------------------
`

const testAutoScalerOceanAKSNPConfig_EmptyFields = `
// --- AutoScaler ---------------------------------------------------- 
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
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				//ExpectNonEmptyPlan: true,
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
				//ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "scheduling.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.shutdown_hours.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.shutdown_hours.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.shutdown_hours.0.time_windows.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.shutdown_hours.0.time_windows.0", "Sat:08:00-Sun:08:00"),
				),
			},
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					scheduling:          testSchedulingOceanAKSNPConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "scheduling.#", "0"),
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
      time_windows = ["Sat:08:00-Sun:08:00"]
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

const testSchedulingOceanAKSNPConfig_EmptyFields = `
// --- Scheduling ---------------------------------------------------- 
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
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				//ExpectNonEmptyPlan: true,
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
				//ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "health.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health.0.grace_period", "60"),
				),
			},
			/*{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					health:              testHealthOceanAKSNPConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "health.#", "0"),
				),
			},*/
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

/*const testHealthOceanAKSNPConfig_EmptyFields = `
// --- Health ----------------------------------------------------
// ---------------------------------------------------------------
`
*/
//endregion

// region Ocean AKS : Headrooms
func TestAccSpotinstOceanAKSNP_Headrooms(t *testing.T) {
	clusterName := "Terraform-Test-AKS-2-0-Cluster"
	controllerClusterID := "terraform-aks-2-0-cluster"
	resourceName := createOceanAKSNPResourceName(clusterName)

	var cluster azure_np.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				//ExpectNonEmptyPlan: true,
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
				//ExpectNonEmptyPlan: true,
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
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:          clusterName,
					controllerClusterID:  controllerClusterID,
					headrooms:            testHeadroomsOceanAKSNPConfig_EmptyFields,
					updateBaselineFields: true,
				}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "headrooms.#", "0"),
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

const testHeadroomsOceanAKSNPConfig_EmptyFields = `

// --- autoscale --------------------------------------------------------
//-----------------------------------------------------------------------
`

//endregion

// region Ocean AKS : Taints
func TestAccSpotinstOceanAKSNP_Taints(t *testing.T) {
	clusterName := "Terraform-Test-AKS-2-0-Cluster"
	controllerClusterID := "terraform-aks-2-0-cluster"
	resourceName := createOceanAKSNPResourceName(clusterName)

	var cluster azure_np.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				//ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "taints.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.key", "taintKey1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.value", "taintValue1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.effect", "NoExecute"),
				),
			},
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:          clusterName,
					controllerClusterID:  controllerClusterID,
					taints:               testTaintsOceanAKSNPConfig_Update,
					updateBaselineFields: true,
				}),
				//ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
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
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					taints:              testTaintsOceanAKSNPConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "taints.#", "0"),
				),
			},
		},
	})
}

const testTaintsOceanAKSNPConfig_Create = `
    
  taints {
    key    = "taintKey1"
    value  = "taintValue1"
    effect = "NoExecute"
  }
`
const testTaintsOceanAKSNPConfig_Update = `
  
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

`

const testTaintsOceanAKSNPConfig_EmptyFields = `
// --- Taints --------------------------------------------------------
//-----------------------------------------------------------------------
`

//endregion

// region Ocean AKS : Filters
func TestAccSpotinstOceanAKSNP_Filters(t *testing.T) {
	clusterName := "Terraform-Test-AKS-2-0-Cluster"
	controllerClusterID := "terraform-aks-2-0-cluster"
	resourceName := createOceanAKSNPResourceName(clusterName)

	var cluster azure_np.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAKSNPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				//ExpectNonEmptyPlan: true,
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
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:          clusterName,
					controllerClusterID:  controllerClusterID,
					filters:              testFiltersOceanAKSNPConfig_Update,
					updateBaselineFields: true,
				}),
				//ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
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
				Config: createOceanAKSNPTerraform(&OceanAKSNPMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					filters:             testFiltersOceanAKSNPConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSNPExists(&cluster, resourceName),
					testCheckOceanAKSNPAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "filters.#", "0"),
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

  // ----------------------------------------------------------------------
`

const testFiltersOceanAKSNPConfig_Update = `
  // --- vmSizes ----------------------------------------------------------

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

  // ----------------------------------------------------------------------
`

const testFiltersOceanAKSNPConfig_EmptyFields = `
// --- vmSizes --------------------------------------------------------
//-----------------------------------------------------------------------
`

//endregion
