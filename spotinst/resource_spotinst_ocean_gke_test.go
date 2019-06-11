package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"log"
	"strings"
	"testing"
)

func init() {
	resource.AddTestSweepers("spotinst_ocean_gke", &resource.Sweeper{
		Name: "spotinst_ocean_gke",
		F:    testSweepOceanGKE,
	})
}

func testSweepOceanGKE(region string) error {
	client, err := getProviderClient("gcp")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).ocean.CloudProviderGCP()
	input := &gcp.ListClustersInput{}
	if resp, err := conn.ListClusters(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of clusters to sweep")
	} else {
		if len(resp.Clusters) == 0 {
			log.Printf("[INFO] No clusters to sweep")
		}
		for _, cluster := range resp.Clusters {
			if strings.Contains(spotinst.StringValue(cluster.Name), "test-acc-") {
				if _, err := conn.DeleteCluster(context.Background(), &gcp.DeleteClusterInput{ClusterID: cluster.ID}); err != nil {
					return fmt.Errorf("unable to delete cluster %v in sweep", spotinst.StringValue(cluster.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(cluster.ID))
				}
			}
		}
	}
	return nil
}

func createOceanGKEResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanGKEResourceName), name)
}

func testOceanGKEDestroy(s *terraform.State) error {
	client := testAccProviderGCP.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OceanGKEResourceName) {
			continue
		}
		input := &gcp.ReadClusterInput{ClusterID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderGCP().ReadCluster(context.Background(), input)
		if err == nil && resp != nil && resp.Cluster != nil {
			return fmt.Errorf("cluster still exists")
		}
	}
	return nil
}

func testCheckOceanGKEAttributes(cluster *gcp.Cluster, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(cluster.Name) != expectedName {
			return fmt.Errorf("bad content: %v", cluster.Name)
		}
		return nil
	}
}

func testCheckOceanGKEExists(cluster *gcp.Cluster, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderGCP.Meta().(*Client)
		input := &gcp.ReadClusterInput{ClusterID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderGCP().ReadCluster(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.Cluster.Name) != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Cluster not found: %+v,\n %+v\n", resp.Cluster, rs.Primary.Attributes)
		}
		*cluster = *resp.Cluster
		return nil
	}
}

type GKEClusterConfigMetadata struct {
	provider             string
	clusterName          string
	controllerClusterID  string
	instanceWhitelist    string
	launchConfig         string
	networkInterface     string
	variables            string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createOceanGKETerraform(ccm *GKEClusterConfigMetadata) string {
	if ccm == nil {
		return ""
	}

	if ccm.provider == "" {
		ccm.provider = "gcp"
	}

	if ccm.launchConfig == "" {
		ccm.launchConfig = testLaunchConfigOceanGKE_Create
	}

	if ccm.networkInterface == "" {
		ccm.networkInterface = testNetworkInterfacesOceanGKEGroupConfig_Create
	}

	if ccm.instanceWhitelist == "" {
		ccm.instanceWhitelist = testInstanceTypesWhitelistGKEConfig_Create
	}

	template :=
		`provider "gcp" {
	 token   = "fake"
	 account = "fake"
	}
	`

	if ccm.updateBaselineFields {
		format := testBaselineGKEConfig_Update
		template += fmt.Sprintf(format,
			ccm.clusterName,
			ccm.provider,
			ccm.clusterName,
			ccm.controllerClusterID,
			ccm.instanceWhitelist,
			ccm.launchConfig,
			ccm.networkInterface,
			ccm.fieldsToAppend,
		)
	} else {
		format := testBaselineGKEConfig_Create
		template += fmt.Sprintf(format,
			ccm.clusterName,
			ccm.provider,
			ccm.clusterName,
			ccm.controllerClusterID,
			ccm.instanceWhitelist,
			ccm.launchConfig,
			ccm.networkInterface,
			ccm.fieldsToAppend,
		)
	}

	if ccm.variables != "" {
		template = ccm.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", ccm.clusterName, template)
	return template
}

// region OceanGKE: Baseline
func TestAccSpotinstOceanGKE_Baseline(t *testing.T) {
	clusterName := "test-acc-cluster-baseline"
	controllerClusterID := "baseline-controller-id"
	resourceName := createOceanGKEResourceName(clusterName)

	var cluster gcp.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "gcp") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanGKEDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanGKETerraform(&GKEClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEExists(&cluster, resourceName),
					testCheckOceanGKEAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "max_size", "1000"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "1"),
				),
			},
			{
				Config: createOceanGKETerraform(&GKEClusterConfigMetadata{
					clusterName:          clusterName,
					controllerClusterID:  controllerClusterID,
					updateBaselineFields: true}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEExists(&cluster, resourceName),
					testCheckOceanGKEAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "max_size", "10"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "2"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "2"),
				),
			},
		},
	})
}

const testBaselineGKEConfig_Create = `
resource "` + string(commons.OceanGKEResourceName) + `" "%v" {
  provider = "%v"  
  
  name = "%v"
  controller_id = "%v"
  cluster_name = "terraform-acc-test-cluster"

  availability_zones = ["us-central1-a"]

  //max_size         = 0
  //min_size         = 0
  //desired_capacity = 0

  subnet_name     = "tf-subnet-1"
  master_location = "us-central1-a"

 %v
 %v
 %v
 %v
}
`

const testBaselineGKEConfig_Update = `
resource "` + string(commons.OceanGKEResourceName) + `" "%v" {
  provider = "%v"

  name = "%v"
  controller_id = "%v"
  cluster_name = "terraform-acc-test-cluster"

  availability_zones = ["us-central1-a"]

  max_size         = 10
  min_size         = 2
  desired_capacity = 2

  subnet_name     = "tf-subnet-1"

 %v
 %v
 %v
 %v
}
`

// endregion

// region OceanGKE: Instance Types Whitelist
func TestAccSpotinstOceanGKE_InstanceTypesWhitelist(t *testing.T) {
	clusterName := "test-acc-cluster-instance-types-whitelist"
	controllerClusterID := "whitelist-controller-id"
	resourceName := createOceanGKEResourceName(clusterName)

	var cluster gcp.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "gcp") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanGKEDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanGKETerraform(&GKEClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					instanceWhitelist:   testInstanceTypesWhitelistGKEConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEExists(&cluster, resourceName),
					testCheckOceanGKEAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.0", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.1", "n1-standard-2"),
				),
			},
			{
				Config: createOceanGKETerraform(&GKEClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					instanceWhitelist:   testInstanceTypesWhitelistGKEConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEExists(&cluster, resourceName),
					testCheckOceanGKEAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.0", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.1", "n1-standard-2"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.2", "n1-standard-4"),
				),
			},
			{
				Config: createOceanGKETerraform(&GKEClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					instanceWhitelist:   testInstanceTypesWhitelistGKEConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEExists(&cluster, resourceName),
					testCheckOceanGKEAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "0"),
				),
			},
		},
	})
}

const testInstanceTypesWhitelistGKEConfig_Create = `
  whitelist = ["n1-standard-1", "n1-standard-2"]
`

const testInstanceTypesWhitelistGKEConfig_Update = `
  whitelist = ["n1-standard-1", "n1-standard-2", "n1-standard-4"]
`

const testInstanceTypesWhitelistGKEConfig_EmptyFields = `
`

// endregion

// region OceanGKE: Launch Configuration
func TestAccSpotinstOceanGKE_LaunchConfiguration(t *testing.T) {
	clusterName := "test-acc-luster-launch-configuration"
	controllerClusterID := "launch-config-cluster-id"
	resourceName := createOceanGKEResourceName(clusterName)

	var cluster gcp.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "gcp") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanGKEDestroy,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createOceanGKETerraform(&GKEClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					launchConfig:        testLaunchConfigOceanGKE_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEExists(&cluster, resourceName),
					testCheckOceanGKEAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "source_image", "https://www.googleapis.com/compute/v1/projects/spotinst-labs/global/images/test-image-1"),
					resource.TestCheckResourceAttr(resourceName, "root_volume_size_in_gb", "100"),
					resource.TestCheckResourceAttr(resourceName, "ip_forwarding", "true"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.0", "gke-test-native-vpc-5cb557f7-node"),
					resource.TestCheckResourceAttr(resourceName, "labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "labels.3733835725.key", "spotinst-gke-original-node-pool"),
					resource.TestCheckResourceAttr(resourceName, "labels.3733835725.value", "terraform-acc-test-cluster__default-pool"),
					resource.TestCheckResourceAttr(resourceName, "metadata.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "metadata.1434173804.key", "cluster-name"),
					resource.TestCheckResourceAttr(resourceName, "metadata.1434173804.value", "terraform-acc-test-cluster"),
					resource.TestCheckResourceAttr(resourceName, "metadata.1912256051.key", "metadata_key"),
					resource.TestCheckResourceAttr(resourceName, "metadata.1912256051.value", "metadata_value"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanGKETerraform(&GKEClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					launchConfig:        testLaunchConfigOceanGKE_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEExists(&cluster, resourceName),
					testCheckOceanGKEAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "source_image", "https://www.googleapis.com/compute/v1/projects/spotinst-labs/global/images/test-image-1"),
					resource.TestCheckResourceAttr(resourceName, "root_volume_size_in_gb", "101"),
					resource.TestCheckResourceAttr(resourceName, "ip_forwarding", "false"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.0", "gke-test-native-vpc-5cb557f7-node"),
					resource.TestCheckResourceAttr(resourceName, "labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "labels.3733835725.key", "spotinst-gke-original-node-pool"),
					resource.TestCheckResourceAttr(resourceName, "labels.3733835725.value", "terraform-acc-test-cluster__default-pool"),
					resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.1434173804.key", "cluster-name"),
					resource.TestCheckResourceAttr(resourceName, "metadata.1434173804.value", "terraform-acc-test-cluster"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanGKETerraform(&GKEClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					launchConfig:        testLaunchConfigOceanGKE_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEExists(&cluster, resourceName),
					testCheckOceanGKEAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "source_image", "https://www.googleapis.com/compute/v1/projects/spotinst-labs/global/images/test-image-1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.0", "gke-test-native-vpc-5cb557f7-node"),
					resource.TestCheckResourceAttr(resourceName, "labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "labels.3733835725.key", "spotinst-gke-original-node-pool"),
					resource.TestCheckResourceAttr(resourceName, "labels.3733835725.value", "terraform-acc-test-cluster__default-pool"),
					resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.1434173804.key", "cluster-name"),
					resource.TestCheckResourceAttr(resourceName, "metadata.1434173804.value", "terraform-acc-test-cluster"),
				),
			},
		},
	})
}

const testLaunchConfigOceanGKE_Create = `
 // --- LAUNCH CONFIGURATION --------------
  source_image = "https://www.googleapis.com/compute/v1/projects/spotinst-labs/global/images/test-image-1"
  service_account = "terraform-acc-test-account@spotinst-labs.iam.gserviceaccount.com"

  // required fields
  labels {
     key = "spotinst-gke-original-node-pool"
     value = "terraform-acc-test-cluster__default-pool" 
   }

  metadata {
       key = "metadata_key"
       value = "metadata_value"
    }

	metadata {
      key = "cluster-name"
      value = "terraform-acc-test-cluster"
    }

  tags = ["gke-test-native-vpc-5cb557f7-node"]

  // optional fields
  root_volume_size_in_gb = 100
  ip_forwarding = true

  backend_services {
    service_name  = "terraform-acc-test-backend-service"
    location_type = "global"
  
    named_ports {
      name = "http"
      ports = [80, 8080]
    }
  }

 // ---------------------------------------
`

const testLaunchConfigOceanGKE_Update = `
 // --- LAUNCH CONFIGURATION --------------
  source_image = "https://www.googleapis.com/compute/v1/projects/spotinst-labs/global/images/test-image-1"
  service_account = "terraform-acc-test-account@spotinst-labs.iam.gserviceaccount.com"

  labels {
      key = "spotinst-gke-original-node-pool"
      value = "terraform-acc-test-cluster__default-pool" 
    }

  metadata {
	   key = "cluster-name"
		 value = "terraform-acc-test-cluster"
	}

  tags = ["gke-test-native-vpc-5cb557f7-node"]

  // optional fields
  root_volume_size_in_gb = 101
  ip_forwarding = false

  backend_services {
    service_name  = "terraform-acc-test-backend-service"
    location_type = "global"
  
    named_ports {
      name = "http"
      ports = [80, 8080]
    }
  }
  
 // ---------------------------------------
`

const testLaunchConfigOceanGKE_EmptyFields = `
 // --- LAUNCH CONFIGURATION --------------
  source_image = "https://www.googleapis.com/compute/v1/projects/spotinst-labs/global/images/test-image-1"
  service_account = "terraform-acc-test-account@spotinst-labs.iam.gserviceaccount.com"

  labels {
      key = "spotinst-gke-original-node-pool"
      value = "terraform-acc-test-cluster__default-pool" 
    }

  metadata {
	   key = "cluster-name"
       value = "terraform-acc-test-cluster"
	}

  tags = ["gke-test-native-vpc-5cb557f7-node"]

  backend_services {
    service_name  = "terraform-acc-test-backend-service"
    location_type = "global"
  
    named_ports  {
      name = "http"
      ports = [80, 8080]
    }
  }
 // ---------------------------------------
`

// endregion

// region OceanGKE: Autoscaler
func TestAccSpotinstOceanGKE_Autoscaler(t *testing.T) {
	clusterName := "test-acc-cluster-autoscaler"
	controllerClusterID := "autoscaler-controller-id"
	resourceName := createOceanGKEResourceName(clusterName)

	var cluster gcp.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "gcp") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanGKEDestroy,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createOceanGKETerraform(&GKEClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					fieldsToAppend:      testGKEScalingConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEExists(&cluster, resourceName),
					testCheckOceanGKEAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_cooldown", "300"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.0.evaluation_periods", "300"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.gpu_per_unit", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.num_of_units", "2"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_auto_config", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_memory_gib", "20"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "1024"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanGKETerraform(&GKEClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					fieldsToAppend:      testGKEScalingConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEExists(&cluster, resourceName),
					testCheckOceanGKEAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_cooldown", "600"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.0.evaluation_periods", "600"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.cpu_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.gpu_per_unit", "2"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.memory_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.num_of_units", "4"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_auto_config", "true"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_memory_gib", "30"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "512"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanGKETerraform(&GKEClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					fieldsToAppend:      testGKEScalingConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEExists(&cluster, resourceName),
					testCheckOceanGKEAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_cooldown", "300"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.0.evaluation_periods", "300"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.num_of_units", "2"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_auto_config", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_memory_gib", "20"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "1024"),
				),
			},
		},
	})
}

const testGKEScalingConfig_Create = `
 // --- AUTOSCALER -----------------
 autoscaler {
    autoscale_is_enabled     = false
    autoscale_is_auto_config = false
    autoscale_cooldown       = 300

    autoscale_headroom {
      cpu_per_unit    = 1024
      gpu_per_unit    = 1
      memory_per_unit = 512
      num_of_units    = 2
    }

    autoscale_down {
      evaluation_periods = 300
    }

    resource_limits {
      max_vcpu       = 1024
      max_memory_gib = 20
    }
 }
 // --------------------------------

`

const testGKEScalingConfig_Update = `
 // --- AUTOSCALER -----------------
 autoscaler {
    autoscale_is_enabled     = true
    autoscale_is_auto_config = true
    autoscale_cooldown       = 600

    autoscale_headroom {
      cpu_per_unit    = 512
      gpu_per_unit    = 2
      memory_per_unit = 1024
      num_of_units    = 4
    }

    autoscale_down {
      evaluation_periods = 600
    }

    resource_limits {
      max_vcpu       = 512
      max_memory_gib = 30
    }
 }
 // --------------------------------
`

const testGKEScalingConfig_EmptyFields = `
 // --- AUTOSCALER -----------------
 autoscaler {
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

    resource_limits {
      max_vcpu   = 1024
      max_memory_gib = 20
    }
 }
 // --------------------------------
`

// endregion

// region OceanGKE: Network Interfaces

func TestAccSpotinstOceanGKE_NetworkInterfaces(t *testing.T) {
	clusterName := "test-acc-ocean-network-interfaces"
	controllerClusterID := "network-interface-controller-id"
	resourceName := createOceanGKEResourceName(clusterName)

	var cluster gcp.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "gcp") },
		Providers:     TestAccProviders,
		CheckDestroy:  testOceanGKEDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createOceanGKETerraform(&GKEClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					networkInterface:    testNetworkInterfacesOceanGKEGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEExists(&cluster, resourceName),
					testCheckOceanGKEAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "network_interface.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.network", "terraform-acc-test-vpc-network"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.1095511731.name", "config1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.1095511731.type", "ONE_TO_ONE_NAT"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.alias_ip_ranges.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.alias_ip_ranges.1202464105.subnetwork_range_name", "range-1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.alias_ip_ranges.1202464105.ip_cidr_range", "10.8.0.0/20"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanGKETerraform(&GKEClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					networkInterface:    testNetworkInterfacesOceanGKEGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEExists(&cluster, resourceName),
					testCheckOceanGKEAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "network_interface.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.network", "terraform-acc-test-vpc-network"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.2016707571.name", "config2"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.2016707571.type", "ONE_TO_ONE_NAT"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.alias_ip_ranges.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.alias_ip_ranges.1202464105.subnetwork_range_name", "range-1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.alias_ip_ranges.1202464105.ip_cidr_range", "10.8.0.0/20"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.1.network", "new-network"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.1.access_configs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.1.access_configs.1864165171.name", "config3"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.1.access_configs.1864165171.type", "ONE_TO_ONE_NAT"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanGKETerraform(&GKEClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					networkInterface:    testNetworkInterfacesOceanGKEGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEExists(&cluster, resourceName),
					testCheckOceanGKEAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "network_interface.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.network", "terraform-acc-test-vpc-network"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.2016707571.name", "config2"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.2016707571.type", "ONE_TO_ONE_NAT"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.alias_ip_ranges.#", "0"),
				),
			},
		},
	})
}

const testNetworkInterfacesOceanGKEGroupConfig_Create = `
 // --- NETWORK INTERFACE ------------------
  network_interface { 
    network = "terraform-acc-test-vpc-network"
	
   access_configs {
     name = "config1"
     type = "ONE_TO_ONE_NAT"
   }
  
   alias_ip_ranges {
     subnetwork_range_name = "range-1"
     ip_cidr_range = "10.8.0.0/20"
   }
  }
 // ----------------------------------------
`

const testNetworkInterfacesOceanGKEGroupConfig_Update = `
 // --- NETWORK INTERFACE ------------------
  	network_interface { 
      network = "terraform-acc-test-vpc-network"
	
      access_configs {
        name = "config2"
        type = "ONE_TO_ONE_NAT"
      }
  
      alias_ip_ranges {
        subnetwork_range_name = "range-1"
        ip_cidr_range = "10.8.0.0/20"
      }
    }
    network_interface { 
      network = "new-network"
	
      access_configs {
        name = "config3"
        type = "ONE_TO_ONE_NAT"
      }
    }
 // ----------------------------------------
`

const testNetworkInterfacesOceanGKEGroupConfig_EmptyFields = `
 // --- NETWORK INTERFACE ------------------
  network_interface {     
    network = "terraform-acc-test-vpc-network"

   access_configs {
     name = "config2"
     type = "ONE_TO_ONE_NAT"
   }
  }
 // ----------------------------------------
`

// endregion
