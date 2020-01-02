package spotinst

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

var GcpClusterName = "terraform-acc-tests-do-not-delete"

func init() {
	resource.AddTestSweepers("resource_spotinst_ocean_gke_import", &resource.Sweeper{
		Name: "resource_spotinst_ocean_gke_import",
		F:    testSweepOceanGKEImportCluster,
	})
}

func testSweepOceanGKEImportCluster(region string) error {
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
			if strings.Contains(spotinst.StringValue(cluster.Name), "terraform-acc-tests-") {
				if _, err := conn.DeleteCluster(context.Background(), &gcp.DeleteClusterInput{ClusterID: cluster.ID}); err != nil {
					return fmt.Errorf("unable to delete group %v in sweep", spotinst.StringValue(cluster.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(cluster.ID))
				}
			}
		}
	}
	return nil
}

func createOceanGKEImportResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanGKEImportResourceName), name)
}

func testOceanGKEImportDestroy(s *terraform.State) error {
	client := testAccProviderGCP.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OceanGKEImportResourceName) {
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

func testCheckOceanGKEImportAttributes(cluster *gcp.Cluster, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(cluster.Name) != expectedName {
			return fmt.Errorf("bad content: %v", cluster.Name)
		}
		return nil
	}
}

func testCheckOceanGKEImportExists(cluster *gcp.Cluster, resourceName string) resource.TestCheckFunc {
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
		//
		// During import, Spotinst API sets the 'name' field to be as the GCP cluster name
		// GCP cluster name cannot be changed after creation while resource have dynamic names per test
		//
		//if spotinst.StringValue(resp.Cluster.Name) != rs.Primary.Attributes["name"] {
		//	return fmt.Errorf("Cluster not found: %+v,\n %+v\n", resp.Cluster, rs.Primary.Attributes)
		//}
		*cluster = *resp.Cluster
		return nil
	}
}

type OceanGKEImportMetadata struct {
	clusterName          string
	provider             string
	updateBaselineFields bool
}

func createOceanGKEImportTerraform(clusterMeta *OceanGKEImportMetadata, update string, create string) string {
	if clusterMeta == nil {
		return ""
	}

	if clusterMeta.provider == "" {
		clusterMeta.provider = "gcp"
	}

	template :=
		`provider "gcp" {
	token   = "fake"
	account = "fake"
	}
	`
	if clusterMeta.updateBaselineFields {
		format := update

		template += fmt.Sprintf(format,
			clusterMeta.clusterName,
			clusterMeta.provider,
		)
	} else {
		format := create
		template += fmt.Sprintf(format,
			clusterMeta.clusterName,
			clusterMeta.provider,
		)
	}

	log.Printf("Terraform [%v] template:\n%v", clusterMeta.clusterName, template)
	return template
}

// region Ocean GKE Import: Baseline
func TestAccSpotinstOceanGKEImport_Baseline(t *testing.T) {
	spotClusterName := "terraform-acc-tests-ocean-gke-import-baseline"
	resourceName := createOceanGKEImportResourceName(spotClusterName)

	var cluster gcp.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "gcp") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanGKEImportDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanGKEImportTerraform(&OceanGKEImportMetadata{clusterName: spotClusterName}, testBaselineOceanGKEImportConfig_Update, testBaselineOceanGKEImportConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEImportExists(&cluster, resourceName),
					testCheckOceanGKEImportAttributes(&cluster, GcpClusterName),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.0", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.1", "n1-standard-2"),
				),
			},
			{
				Config: createOceanGKEImportTerraform(&OceanGKEImportMetadata{clusterName: spotClusterName, updateBaselineFields: true}, testBaselineOceanGKEImportConfig_Update, testBaselineOceanGKEImportConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEImportExists(&cluster, resourceName),
					testCheckOceanGKEImportAttributes(&cluster, GcpClusterName),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.0", "n1-standard-1"),
				),
			},
		},
	})
}

const testBaselineOceanGKEImportConfig_Create = `
resource "` + string(commons.OceanGKEImportResourceName) + `" "%v" {
 provider = "%v"

 //min_size 			= 0
 //max_size 			= 1
 //desired_capacity 	= 0
 cluster_name = "terraform-acc-tests-do-not-delete"
 location     = "us-central1-a"

 whitelist = ["n1-standard-1", "n1-standard-2"]
}

`

const testBaselineOceanGKEImportConfig_Update = `
resource "` + string(commons.OceanGKEImportResourceName) + `" "%v" {
 provider = "%v"

 //min_size = 0
 //max_size = 1
 //desired_capacity = 0
 cluster_name = "terraform-acc-tests-do-not-delete"
 location     = "us-central1-a"

 whitelist = ["n1-standard-1"]
}

`

// endregion

// region Ocean GKE Import: BackendServices
func TestAccSpotinstOceanGKEImport_BackendServices(t *testing.T) {
	spotClusterName := "terraform-acc-tests-ocean-gke-import-be-services"
	resourceName := createOceanGKEImportResourceName(spotClusterName)

	var cluster gcp.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "gcp") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanGKEImportDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanGKEImportTerraform(&OceanGKEImportMetadata{clusterName: spotClusterName}, testBackendServicesOceanGKEImportConfig_Update, testBackendServicesOceanGKEImportConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEImportExists(&cluster, resourceName),
					testCheckOceanGKEImportAttributes(&cluster, GcpClusterName),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.0", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.1", "n1-standard-2"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.1781664423.location_type", "global"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.1781664423.service_name", "terraform-acc-test-backend-service"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.1781664423.named_ports.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.1781664423.named_ports.571950593.name", "http"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.1781664423.named_ports.571950593.ports.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.1781664423.named_ports.571950593.ports.0", "80"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.1781664423.named_ports.571950593.ports.1", "8080"),
				),
			},
			{
				Config: createOceanGKEImportTerraform(&OceanGKEImportMetadata{clusterName: spotClusterName, updateBaselineFields: true}, testBackendServicesOceanGKEImportConfig_Update, testBackendServicesOceanGKEImportConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKEImportExists(&cluster, resourceName),
					testCheckOceanGKEImportAttributes(&cluster, GcpClusterName),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.0", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.1", "n1-standard-2"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.3984833389.location_type", "global"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.3984833389.service_name", "terraform-acc-test-backend-service"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.3984833389.named_ports.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.3984833389.named_ports.2171153412.name", "https"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.3984833389.named_ports.2171153412.ports.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backend_services.3984833389.named_ports.2171153412.ports.0", "443"),
				),
			},
		},
	})
}

const testBackendServicesOceanGKEImportConfig_Create = `
resource "` + string(commons.OceanGKEImportResourceName) + `" "%v" {
 provider = "%v"

 cluster_name = "terraform-acc-tests-do-not-delete"
 location     = "us-central1-a"

 whitelist = ["n1-standard-1", "n1-standard-2"]

  backend_services {
      service_name = "terraform-acc-test-backend-service"
      location_type = "global"

      named_ports {
        name = "http"
        ports = [
          80,
          8080]
      }
    }
}

`

const testBackendServicesOceanGKEImportConfig_Update = `
resource "` + string(commons.OceanGKEImportResourceName) + `" "%v" {
 provider = "%v"

 cluster_name = "terraform-acc-tests-do-not-delete"
 location     = "us-central1-a"

 whitelist = ["n1-standard-1", "n1-standard-2"]

  backend_services {
      service_name = "terraform-acc-test-backend-service"
      location_type = "global"

      named_ports {
        name = "https"
        ports = [443]
      }
    }
}

`

// endregion
