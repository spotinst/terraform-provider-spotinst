package spotinst

import (
	"context"
	"fmt"
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
	resource.AddTestSweepers("resource_spotinst_ocean_aks_import", &resource.Sweeper{
		Name: "resource_spotinst_ocean_aks",
		F:    testSweepOceanAKSCluster,
	})
}

func testSweepOceanAKSCluster(region string) error {
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

func createOceanAKSResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanAKSResourceName), name)
}

func testOceanAKSDestroy(s *terraform.State) error {
	client := testAccProviderAzure.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OceanAKSResourceName) {
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

func testCheckOceanAKSAttributes(cluster *azure.Cluster, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(cluster.Name) != expectedName {
			return fmt.Errorf("bad content: %v", cluster.Name)
		}
		return nil
	}
}

func testCheckOceanAKSExists(cluster *azure.Cluster, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAzure.Meta().(*Client)
		input := &azure.ReadClusterInput{ClusterID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAzure().ReadCluster(context.Background(), input)
		if err != nil {
			return err
		}
		*cluster = *resp.Cluster
		return nil
	}
}

type OceanAKSMetadata struct {
	clusterName          string
	acdIdentifier        string
	controllerClusterID  string
	provider             string
	launchSpecification  string
	strategy             string
	autoScaler           string
	health               string
	vmSizes              string
	osDisk               string
	image                string
	loadBalancers        string
	network              string
	extensions           string
	login                string
	variables            string
	updateBaselineFields bool
}

func createOceanAKSTerraform(clusterMeta *OceanAKSMetadata) string {
	if clusterMeta == nil {
		return ""
	}

	if clusterMeta.provider == "" {
		clusterMeta.provider = "azure"
	}

	if clusterMeta.launchSpecification == "" {
		clusterMeta.launchSpecification = testLaunchSpecificationOceanAKSConfig_Create
	}

	if clusterMeta.strategy == "" {
		clusterMeta.strategy = testStrategyOceanAKSConfig_Create
	}

	if clusterMeta.autoScaler == "" {
		clusterMeta.autoScaler = testAutoScalerOceanAKSConfig_Create
	}

	if clusterMeta.health == "" {
		clusterMeta.health = testHealthOceanAKSConfig_Create
	}

	if clusterMeta.vmSizes == "" {
		clusterMeta.vmSizes = testVMSizesOceanAKSConfig_Create
	}

	if clusterMeta.osDisk == "" {
		clusterMeta.osDisk = testOSDiskOceanAKSConfig_Create
	}

	if clusterMeta.image == "" {
		clusterMeta.image = testImageOceanAKSConfig_Create
	}

	if clusterMeta.loadBalancers == "" {
		clusterMeta.loadBalancers = testLoadBalancersOceanAKSConfig_Create
	}

	if clusterMeta.network == "" {
		clusterMeta.network = testNetworkOceanAKSConfig_Create
	}

	if clusterMeta.extensions == "" {
		clusterMeta.extensions = testExtensionsOceanAKSConfig_Create
	}

	if clusterMeta.login == "" {
		clusterMeta.login = testLoginOceanAKSConfig_Create
	}

	template :=
		`provider "azure" {
	token   = "fake"
	account = "fake"
	}
	`
	if clusterMeta.updateBaselineFields {
		format := testBaselineOceanAKSConfig_Update
		template += fmt.Sprintf(format,
			clusterMeta.clusterName,
			clusterMeta.provider,
			clusterMeta.acdIdentifier,
			clusterMeta.controllerClusterID,
			clusterMeta.login,
			clusterMeta.launchSpecification,
			clusterMeta.osDisk,
			clusterMeta.strategy,
			clusterMeta.image,
			clusterMeta.autoScaler,
			clusterMeta.extensions,
			clusterMeta.network,
			clusterMeta.health,
			clusterMeta.loadBalancers,
			clusterMeta.vmSizes,
		)
	} else {
		format := testBaselineOceanAKSConfig_Create
		template += fmt.Sprintf(format,
			clusterMeta.clusterName,
			clusterMeta.provider,
			clusterMeta.clusterName,
			clusterMeta.acdIdentifier,
			clusterMeta.controllerClusterID,
			clusterMeta.login,
			clusterMeta.launchSpecification,
			clusterMeta.osDisk,
			clusterMeta.strategy,
			clusterMeta.image,
			clusterMeta.autoScaler,
			clusterMeta.extensions,
			clusterMeta.network,
			clusterMeta.health,
			clusterMeta.loadBalancers,
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
func TestAccSpotinstOceanAKS_Baseline(t *testing.T) {
	clusterName := "terraform-Kubernetes-cluster"
	acdIdentifier := "acd-aa5c6795"
	controllerClusterID := "terraform-Kubernetes-cluster"
	resourceName := createOceanAKSResourceName(clusterName)

	var cluster azure.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "azure") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testOceanAKSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:         clusterName,
					acdIdentifier:       acdIdentifier,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSExists(&cluster, resourceName),
					testCheckOceanAKSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "acd_identifier", "acd-aa5c6795"),
					resource.TestCheckResourceAttr(resourceName, "name", "terraform-Kubernetes-cluster"),
					resource.TestCheckResourceAttr(resourceName, "controller_cluster_id", "terraform-Kubernetes-cluster"),
					resource.TestCheckResourceAttr(resourceName, "aks_name", "terraform-Kubernetes-cluster"),
					resource.TestCheckResourceAttr(resourceName, "aks_resource_group_name", "terraform-resource-group-DO-NOT-DELETE"),
				),
			},
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:          clusterName,
					acdIdentifier:        acdIdentifier,
					controllerClusterID:  controllerClusterID,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "new-ocean-cluster-name"),
				),
			},
		},
	})
}

const testBaselineOceanAKSConfig_Create = `
resource "` + string(commons.OceanAKSResourceName) + `" "%v" {
  provider = "%v"
  name                  = "%v"
  acd_identifier        = "%v"
  controller_cluster_id = "%v"

  // --- AKS -----------------------------------------------------------
  aks_name                = "terraform-Kubernetes-cluster"
  aks_resource_group_name = "terraform-resource-group-DO-NOT-DELETE"
  // -------------------------------------------------------------------

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
%v
}
`

const testBaselineOceanAKSConfig_Update = `
resource "` + string(commons.OceanAKSResourceName) + `" "%v" {
  
  provider = "%v"
  name                  = "new-ocean-cluster-name"
  acd_identifier        = "%v"
  controller_cluster_id = "%v"
  // --- AKS -----------------------------------------------------------
  aks_name                = "terraform-Kubernetes-cluster"
  aks_resource_group_name = "terraform-resource-group-DO-NOT-DELETE"
  // -------------------------------------------------------------------

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
%v
}
`

//endregion

// region Ocean AKS : Login
func TestAccSpotinstOceanAKS_Login(t *testing.T) {
	clusterName := "terraform-tests-do-not-delete"
	acdIdentifier := "acd-aa5c6795"
	controllerClusterID := "terraform-Kubernetes-cluster"
	resourceName := createOceanAKSResourceName(clusterName)

	var cluster azure.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "azure") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testOceanAKSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:         clusterName,
					acdIdentifier:       acdIdentifier,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSExists(&cluster, resourceName),
					testCheckOceanAKSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "ssh_public_key", ""),
					resource.TestCheckResourceAttr(resourceName, "user_name", "terraform-user"),
				),
			},
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:          clusterName,
					acdIdentifier:        acdIdentifier,
					controllerClusterID:  controllerClusterID,
					login:                testLoginOceanAKSConfig_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ssh_public_key", ""),
					resource.TestCheckResourceAttr(resourceName, "user_name", "new-terraform-user"),
				),
			},
		},
	})
}

const testLoginOceanAKSConfig_Create = `
  // --- LOGIN ---------------------------------------------------------
  ssh_public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDrN+SKCc648JwaroAemjmTxUqKgDtmYZ8UE7lhTdPaHG/Yybaghci1KYCnVuGHkyN4sAqWpjxBwCVFWWN+rUEIwitJgXBj9Q4D5LPWr6kopbpvU/HFG0pF8OkcJzvKG5Xo2kmJXG0f8dwe8e4P0S5rRgIiESLC66CJCDtvhDNIKCzvmsuvp88fCzjhia8AZnVWszJXgyn8bXt6c5f29BMDTq8VXX1RcdMS2ZWOqpjlvKxI/bsnqVy20G3W0veMNJiNjUKY9f2mqTAQoaQnCy8LOZMa/tCGOtwKZ2Yr28V8vDWXDF4EDpmjcVbStX/n40Ag6ZFxcaz+2byLp8HXV1PIfry+F3x6+sHSDV3CQvgju/W5+U0sBRIlYhSlbzQQV2qLvt2sXqyWgvfBu5GQ/xb6DH0oskp2QS8SCrry9kLb8IReT0f5INFKGEcTCgT/SEvGFeUHt2wZ4HsVSwsTHp1Z/T0AmDtCEGdZuxx0llEPFFSRKfnckdK1pRHrn6GRI4U= generated-by-azure"
  user_name      = "terraform-user"
  // -------------------------------------------------------------------
`

const testLoginOceanAKSConfig_Update = `
  // --- LOGIN ---------------------------------------------------------
  ssh_public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDrN+SKCc648JwaroAemjmTxUqKgDtmYZ8UE7lhTdPaHG/Yybaghci1KYCnVuGHkyN4sAqWpjxBwCVFWWN+rUEIwitJgXBj9Q4D5LPWr6kopbpvU/HFG0pF8OkcJzvKG5Xo2kmJXG0f8dwe8e4P0S5rRgIiESLC66CJCDtvhDNIKCzvmsuvp88fCzjhia8AZnVWszJXgyn8bXt6c5f29BMDTq8VXX1RcdMS2ZWOqpjlvKxI/bsnqVy20G3W0veMNJiNjUKY9f2mqTAQoaQnCy8LOZMa/tCGOtwKZ2Yr28V8vDWXDF4EDpmjcVbStX/n40Ag6ZFxcaz+2byLp8HXV1PIfry+F3x6+sHSDV3CQvgju/W5+U0sBRIlYhSlbzQQV2qLvt2sXqyWgvfBu5GQ/xb6DH0oskp2QS8SCrry9kLb8IReT0f5INFKGEcTCgT/SEvGFeUHt2wZ4HsVSwsTHp1Z/T0AmDtCEGdZuxx0llEPFFSRKfnckdK1pRHrn6GRI4U= generated-by-azure"
	user_name      = "new-terraform-user"
  // -------------------------------------------------------------------
`

//endregion

// region Ocean AKS : Launch Specification
func TestAccSpotinstOceanAKS_LaunchSpecification(t *testing.T) {
	clusterName := "terraform-tests-do-not-delete"
	acdIdentifier := "acd-aa5c6795"
	controllerClusterID := "terraform-Kubernetes-cluster"
	resourceName := createOceanAKSResourceName(clusterName)

	var cluster azure.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "azure") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testOceanAKSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:         clusterName,
					acdIdentifier:       acdIdentifier,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSExists(&cluster, resourceName),
					testCheckOceanAKSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "resource_group_name", "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"),
					resource.TestCheckResourceAttr(resourceName, "tag.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tag.0.key", "tag-key"),
					resource.TestCheckResourceAttr(resourceName, "tag.0.value", "tag-value")),
			},
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:         clusterName,
					acdIdentifier:       acdIdentifier,
					controllerClusterID: controllerClusterID,
					launchSpecification: testLaunchSpecificationOceanAKSConfig_Update,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSExists(&cluster, resourceName),
					testCheckOceanAKSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "resource_group_name", "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"),
					resource.TestCheckResourceAttr(resourceName, "tag.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tag.0.key", "first-tag-key"),
					resource.TestCheckResourceAttr(resourceName, "tag.0.value", "first-tag-value"),
					resource.TestCheckResourceAttr(resourceName, "tag.1.key", "second-tag-key"),
					resource.TestCheckResourceAttr(resourceName, "tag.1.value", "second-tag-value"),
				),
			},
		},
	})
}

const testLaunchSpecificationOceanAKSConfig_Create = `
  // --- Launch Specification ------------------------------------------------
    resource_group_name = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"
    
	tag {
      key = "tag-key"
      value = "tag-value"
    }

	managed_service_identity {
		resource_group_name = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"
    	name                = "terraform-Kubernetes-cluster-agentpool"
  	}
  // -------------------------------------------------------------------
`

const testLaunchSpecificationOceanAKSConfig_Update = `
  // --- Launch Specification ------------------------------------------------
    resource_group_name = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"
    
	tag {
      key = "first-tag-key"
      value = "first-tag-value"
    }

	tag {
      key = "second-tag-key"
      value = "second-tag-value"
    }
  // -------------------------------------------------------------------
`

//endregion

// region Ocean AKS : AutoScaler
func TestAccSpotinstOceanAKS_AutoScaler(t *testing.T) {
	clusterName := "terraform-tests-do-not-delete"
	acdIdentifier := "acd-aa5c6795"
	controllerClusterID := "terraform-Kubernetes-cluster"
	resourceName := createOceanAKSResourceName(clusterName)

	var cluster azure.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "azure") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testOceanAKSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:         clusterName,
					acdIdentifier:       acdIdentifier,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSExists(&cluster, resourceName),
					testCheckOceanAKSAttributes(&cluster, clusterName),
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
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:          clusterName,
					acdIdentifier:        acdIdentifier,
					controllerClusterID:  controllerClusterID,
					autoScaler:           testAutoScalerOceanAKSConfig_Update,
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

const testAutoScalerOceanAKSConfig_Create = `
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

const testAutoScalerOceanAKSConfig_Update = `
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

// region Ocean AKS : Strategy
func TestAccSpotinstOceanAKS_Strategy(t *testing.T) {
	clusterName := "terraform-tests-do-not-delete"
	acdIdentifier := "acd-aa5c6795"
	controllerClusterID := "terraform-Kubernetes-cluster"
	resourceName := createOceanAKSResourceName(clusterName)

	var cluster azure.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "azure") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testOceanAKSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:         clusterName,
					acdIdentifier:       acdIdentifier,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSExists(&cluster, resourceName),
					testCheckOceanAKSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.fallback_to_ondemand", "true"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.spot_percentage", "40"),
				),
			},
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:          clusterName,
					acdIdentifier:        acdIdentifier,
					controllerClusterID:  controllerClusterID,
					strategy:             testStrategyOceanAKSConfig_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.fallback_to_ondemand", "false"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.spot_percentage", "100"),
				),
			},
		},
	})
}

const testStrategyOceanAKSConfig_Create = `
  // --- Strategy ------------------------------------------------------
  strategy {
    fallback_to_ondemand = true
    spot_percentage = 40
  }
  // -------------------------------------------------------------------
`

const testStrategyOceanAKSConfig_Update = `
  // --- Strategy ------------------------------------------------------
  strategy {
    fallback_to_ondemand = false
    spot_percentage = 100
  }
  // -------------------------------------------------------------------
`

//endregion

// region Ocean AKS : Health
func TestAccSpotinstOceanAKS_Health(t *testing.T) {
	clusterName := "terraform-tests-do-not-delete"
	acdIdentifier := "acd-aa5c6795"
	controllerClusterID := "terraform-Kubernetes-cluster"
	resourceName := createOceanAKSResourceName(clusterName)

	var cluster azure.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "azure") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testOceanAKSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:         clusterName,
					acdIdentifier:       acdIdentifier,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSExists(&cluster, resourceName),
					testCheckOceanAKSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "health.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health.0.grace_period", "10"),
				),
			},
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:          clusterName,
					acdIdentifier:        acdIdentifier,
					controllerClusterID:  controllerClusterID,
					health:               testHealthOceanAKSConfig_Update,
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

const testHealthOceanAKSConfig_Create = `
  // --- Health --------------------------------------------------------
  health {
    grace_period = 10
  }
  // -------------------------------------------------------------------
`

const testHealthOceanAKSConfig_Update = `
  // --- Health --------------------------------------------------------
  health {
    grace_period = 60
  }
  // -------------------------------------------------------------------
`

//endregion

// region Ocean AKS : VMSizes
func TestAccSpotinstOceanAKS_VMSizes(t *testing.T) {
	clusterName := "terraform-tests-do-not-delete"
	acdIdentifier := "acd-aa5c6795"
	controllerClusterID := "terraform-Kubernetes-cluster"
	resourceName := createOceanAKSResourceName(clusterName)

	var cluster azure.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "azure") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testOceanAKSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:         clusterName,
					acdIdentifier:       acdIdentifier,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSExists(&cluster, resourceName),
					testCheckOceanAKSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.whitelist.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.whitelist.0", "standard_ds2_v2"),
				),
			},
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:          clusterName,
					acdIdentifier:        acdIdentifier,
					controllerClusterID:  controllerClusterID,
					vmSizes:              testVMSizesOceanAKSConfig_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.whitelist.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.whitelist.0", "standard_ds2_v2"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.whitelist.1", "standard_ds3_v2"),
				),
			},
		},
	})
}

const testVMSizesOceanAKSConfig_Create = `
  // --- VMSizes -------------------------------------------------------
  vm_sizes {
    whitelist = [
      "standard_ds2_v2"]
  }
  // -------------------------------------------------------------------
`

const testVMSizesOceanAKSConfig_Update = `
  // --- VMSizes -------------------------------------------------------
  vm_sizes {
    whitelist = [
      "standard_ds2_v2",
       "standard_ds3_v2"]
  }
  // -------------------------------------------------------------------
`

//endregion

// region Ocean AKS : OSDisk
func TestAccSpotinstOceanAKS_OSDisk(t *testing.T) {
	clusterName := "terraform-tests-do-not-delete"
	acdIdentifier := "acd-aa5c6795"
	controllerClusterID := "terraform-Kubernetes-cluster"
	resourceName := createOceanAKSResourceName(clusterName)

	var cluster azure.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "azure") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testOceanAKSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:         clusterName,
					acdIdentifier:       acdIdentifier,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSExists(&cluster, resourceName),
					testCheckOceanAKSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "os_disk.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "os_disk.0.size_gb", "130"),
					resource.TestCheckResourceAttr(resourceName, "os_disk.0.type", "Standard_LRS"),
				),
			},
		},
	})
}

const testOSDiskOceanAKSConfig_Create = `
  // --- OSDisk -------------------------------------------------------
  os_disk {
    size_gb = 130
    type = "Standard_LRS"
  }
  // -------------------------------------------------------------------
`

//endregion

// region Ocean AKS : Image
func TestAccSpotinstOceanAKS_Image(t *testing.T) {
	clusterName := "terraform-tests-do-not-delete"
	acdIdentifier := "acd-aa5c6795"
	controllerClusterID := "terraform-Kubernetes-cluster"
	resourceName := createOceanAKSResourceName(clusterName)

	var cluster azure.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "azure") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testOceanAKSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:         clusterName,
					acdIdentifier:       acdIdentifier,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSExists(&cluster, resourceName),
					testCheckOceanAKSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "image.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.offer", "UbuntuServer"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.publisher", "Canonical"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.sku", "18.04-LTS"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.version", "latest"),
				),
			},
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:          clusterName,
					acdIdentifier:        acdIdentifier,
					controllerClusterID:  controllerClusterID,
					image:                testImageOceanAKSConfig_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "image.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.offer", "aks"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.publisher", "microsoft-aks"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.sku", "aks-ubuntu-1804-gen2-2021-q2"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.version", "2021.05.19"),
				),
			},
		},
	})
}

const testImageOceanAKSConfig_Create = `
  // --- Image ---------------------------------------------------------
  image {
    marketplace {
      publisher = "Canonical"
      offer = "UbuntuServer"
      sku = "18.04-LTS"
      version = "latest"
    }
  }
  // ---------------------------------------------------------------------
`

const testImageOceanAKSConfig_Update = `
  // --- Image ---------------------------------------------------------
 image {
    marketplace {
      publisher = "microsoft-aks"
      offer = "aks"
      sku = "aks-ubuntu-1804-gen2-2021-q2"
      version = "2021.05.19"
    }
  }
  // ---------------------------------------------------------------------
`

//endregion

// region Ocean AKS : Load Balancers
func TestAccSpotinstOceanAKS_LoadBalancers(t *testing.T) {
	clusterName := "terraform-tests-do-not-delete"
	acdIdentifier := "acd-aa5c6795"
	controllerClusterID := "terraform-Kubernetes-cluster"
	resourceName := createOceanAKSResourceName(clusterName)

	var cluster azure.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "azure") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testOceanAKSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:         clusterName,
					acdIdentifier:       acdIdentifier,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSExists(&cluster, resourceName),
					testCheckOceanAKSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.0.backend_pool_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.0.backend_pool_names.0", "kubernetes"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.0.load_balancer_sku", "Standard"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.0.name", "kubernetes"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.0.resource_group_name", "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.0.type", "loadBalancer"),
				),
			},
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:          clusterName,
					acdIdentifier:        acdIdentifier,
					controllerClusterID:  controllerClusterID,
					loadBalancers:        testLoadBalancersOceanAKSConfig_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "load_balancer.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.0.backend_pool_names.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.0.backend_pool_names.0", "aksOutboundBackendPool"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.0.backend_pool_names.1", "kubernetes"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.0.load_balancer_sku", "Standard"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.0.name", "kubernetes"),
				),
			},
		},
	})
}

const testLoadBalancersOceanAKSConfig_Create = `
  // --- Load Balancers --------------------------------------------------
  load_balancer {
    backend_pool_names = [
      "kubernetes"
    ]
    load_balancer_sku = "Standard"
    name = "kubernetes"
    resource_group_name = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"
    type = "loadBalancer"
  }
 // -------------------------------------------------------------------
`

const testLoadBalancersOceanAKSConfig_Update = `
  // --- Load Balancers --------------------------------------------------
  load_balancer {
    backend_pool_names = [
      "aksOutboundBackendPool",
       "kubernetes"
    ]
    load_balancer_sku = "Standard"
    name = "kubernetes"
    resource_group_name = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"
    type = "loadBalancer"
  }
 // -------------------------------------------------------------------
`

//endregion

// region Ocean AKS : Network
func TestAccSpotinstOceanAKS_Network(t *testing.T) {
	clusterName := "terraform-tests-do-not-delete"
	acdIdentifier := "acd-aa5c6795"
	controllerClusterID := "terraform-Kubernetes-cluster"
	resourceName := createOceanAKSResourceName(clusterName)

	var cluster azure.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "azure") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testOceanAKSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:         clusterName,
					acdIdentifier:       acdIdentifier,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSExists(&cluster, resourceName),
					testCheckOceanAKSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "network.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.additional_ip_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.additional_ip_config.0.name", "ip-config-name"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.additional_ip_config.0.private_ip_version", "IPv4"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.is_primary", "false"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.security_group.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.security_group.0.name", "aks-agentpool-48068046-nsg"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.security_group.0.resource_group_name", "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.subnet_name", "terraform-subnet-DO-NOT-DELETE"),
					resource.TestCheckResourceAttr(resourceName, "network.0.resource_group_name", "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"),
					resource.TestCheckResourceAttr(resourceName, "network.0.virtual_network_name", "aks-vnet-48068046"),
				),
			},
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:          clusterName,
					acdIdentifier:        acdIdentifier,
					controllerClusterID:  controllerClusterID,
					network:              testNetworkOceanAKSConfig_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "network.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.additional_ip_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.additional_ip_config.0.name", "new-ip-config-name"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.additional_ip_config.0.private_ip_version", "IPv6"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.is_primary", "true"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.security_group.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.security_group.0.name", "aks-agentpool-48068046-nsg"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.security_group.0.resource_group_name", "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.subnet_name", "aks-subnet"),
					resource.TestCheckResourceAttr(resourceName, "network.0.resource_group_name", "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"),
					resource.TestCheckResourceAttr(resourceName, "network.0.virtual_network_name", "aks-vnet-48068046"),
				),
			},
		},
	})
}

const testNetworkOceanAKSConfig_Create = `
  //  // --- NETWORK -------------------------------------------------------
  network {
    virtual_network_name = "aks-vnet-48068046"
    resource_group_name = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"

    network_interface {
      subnet_name = "terraform-subnet-DO-NOT-DELETE"
      assign_public_ip = false
      is_primary = false

      additional_ip_config {
        name = "ip-config-name"
        private_ip_version = "IPv4"
      }

      security_group{
        name = "aks-agentpool-48068046-nsg"
        resource_group_name = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"
      }
    }
  }
  // -------------------------------------------------------------------
`

const testNetworkOceanAKSConfig_Update = `
  //  // --- NETWORK -------------------------------------------------------
  network {
    virtual_network_name = "aks-vnet-48068046"
    resource_group_name = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"

    network_interface {
      subnet_name = "aks-subnet"
      assign_public_ip = false
      is_primary = true

      additional_ip_config {
        name = "new-ip-config-name"
        private_ip_version = "IPv6"
      }

      security_group{
        name = "aks-agentpool-48068046-nsg"
        resource_group_name = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"
      }
    }
  }
  // -------------------------------------------------------------------
`

//endregion

// region Ocean AKS : Extensions
func TestAccSpotinstOceanAKS_Extensions(t *testing.T) {
	clusterName := "terraform-tests-do-not-delete"
	acdIdentifier := "acd-aa5c6795"
	controllerClusterID := "terraform-Kubernetes-cluster"
	resourceName := createOceanAKSResourceName(clusterName)

	var cluster azure.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "azure") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testOceanAKSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:         clusterName,
					acdIdentifier:       acdIdentifier,
					controllerClusterID: controllerClusterID,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAKSExists(&cluster, resourceName),
					testCheckOceanAKSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "extension.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "extension.0.api_version", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "extension.0.minor_version_auto_upgrade", "true"),
					resource.TestCheckResourceAttr(resourceName, "extension.0.name", "terraform-extension"),
					resource.TestCheckResourceAttr(resourceName, "extension.0.publisher", "Microsoft.Azure.Extensions"),
					resource.TestCheckResourceAttr(resourceName, "extension.0.type", "Linux"),
				),
			},
			{
				Config: createOceanAKSTerraform(&OceanAKSMetadata{
					clusterName:          clusterName,
					acdIdentifier:        acdIdentifier,
					controllerClusterID:  controllerClusterID,
					extensions:           testExtensionsOceanAKSConfig_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "extension.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "extension.0.api_version", "2.0"),
					resource.TestCheckResourceAttr(resourceName, "extension.0.minor_version_auto_upgrade", "false"),
					resource.TestCheckResourceAttr(resourceName, "extension.0.name", "OceanAKS"),
					resource.TestCheckResourceAttr(resourceName, "extension.0.publisher", "Microsoft.Azure.Extensions"),
					resource.TestCheckResourceAttr(resourceName, "extension.0.type", "customScript"),
				),
			},
		},
	})
}

const testExtensionsOceanAKSConfig_Create = `
 // --- Extensions ----------------------------------------------------
    extension {
      api_version = "1.0"
      minor_version_auto_upgrade = true
      name = "terraform-extension"
      publisher = "Microsoft.Azure.Extensions"
      type = "Linux"
    }
`
const testExtensionsOceanAKSConfig_Update = `
 // --- Extensions ----------------------------------------------------
    extension {
      api_version = "2.0"
      minor_version_auto_upgrade = false
      name = "OceanAKS"
      publisher = "Microsoft.Azure.Extensions"
      type = "customScript"
    }
`

//endregion
