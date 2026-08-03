package spotinst

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/right_sizing_cluster_config"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func createOceanRightSizingClusterConfigResource(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanRightSizingClusterConfigResourceName), name)
}

func testCheckOceanRightSizingClusterConfigExists(config *right_sizing_cluster_config.RightsizingClusterConfiguration, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}

		oceanID := rs.Primary.Attributes["ocean_id"]
		clusterIdentifier := rs.Primary.Attributes["cluster_identifier"]

		client := testAccProviderAWS.Meta().(*Client)
		input := &right_sizing_cluster_config.ReadRightsizingClusterConfigurationInput{
			OceanId:           spotinst.String(oceanID),
			ClusterIdentifier: spotinst.String(clusterIdentifier),
		}

		resp, err := client.ocean.RightSizingClusterConfig().ReadRightSizingClusterConfiguration(context.Background(), input)
		if err != nil {
			return err
		}

		if resp.ClusterConfiguration == nil {
			return fmt.Errorf("cluster right sizing configuration not found")
		}

		*config = *resp.ClusterConfiguration
		return nil
	}
}

type RightSizingClusterConfigMetadata struct {
	provider             string
	oceanID              string
	clusterIdentifier    string
	config               string
	updateBaselineFields bool
}

func createOceanRightSizingClusterConfigTerraform(ccm *RightSizingClusterConfigMetadata) string {
	if ccm == nil {
		return ""
	}

	if ccm.provider == "" {
		ccm.provider = "aws"
	}

	if ccm.config == "" {
		ccm.config = testRightSizingClusterConfigBaseline_Create
	}

	template :=
		`provider "aws" {
     token   = "fake"
     account = "fake"
    }
    `

	var format string
	if ccm.updateBaselineFields {
		format = testBaselineRightSizingClusterConfig_Update
	} else {
		format = testBaselineRightSizingClusterConfig_Create
	}

	template += fmt.Sprintf(format,
		ccm.oceanID,
		ccm.provider,
		ccm.clusterIdentifier,
		ccm.config,
	)

	log.Printf("Terraform [%v] template:\n%v", ccm.oceanID, template)
	return template
}

const testBaselineRightSizingClusterConfig_Create = `
resource "` + string(commons.OceanRightSizingClusterConfigResourceName) + `" "test" {
  ocean_id           = "%v"
  provider           = "%v"
  cluster_identifier = "%v"

  %v
}
`

const testBaselineRightSizingClusterConfig_Update = `
resource "` + string(commons.OceanRightSizingClusterConfigResourceName) + `" "test" {
  ocean_id           = "%v"
  provider           = "%v"
  cluster_identifier = "%v"

  %v
}
`

const testRightSizingClusterConfigBaseline_Create = `
  config {
    adjust_limit_on_downsize           = true
    downside_only                      = true
    recommendations_cpu_percentile     = 85
    recommendations_memory_percentile  = 85
  }
`

const testRightSizingClusterConfigBaseline_Update = `
  config {
    adjust_limit_on_downsize           = false
    downside_only                      = false
    recommendations_cpu_percentile     = 99
    recommendations_memory_percentile  = 100
  }
`

// TestAccSpotinstOceanRightSizingClusterConfig tests the resource lifecycle (Create, Read, Update, Read)
func TestAccSpotinstOceanRightSizingClusterConfig(t *testing.T) {
	resourceName := createOceanRightSizingClusterConfigResource("test")
	var config right_sizing_cluster_config.RightsizingClusterConfiguration

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t, "aws") },
		Providers: TestAccProviders,

		Steps: []resource.TestStep{
			{
				Config: createOceanRightSizingClusterConfigTerraform(&RightSizingClusterConfigMetadata{
					oceanID:           "o-8b34732f",
					clusterIdentifier: "ocean.k8s",
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanRightSizingClusterConfigExists(&config, resourceName),
					resource.TestCheckResourceAttr(resourceName, "ocean_id", "o-8b34732f"),
					resource.TestCheckResourceAttr(resourceName, "cluster_identifier", "ocean.k8s"),
					resource.TestCheckResourceAttr(resourceName, "config.0.adjust_limit_on_downsize", "true"),
					resource.TestCheckResourceAttr(resourceName, "config.0.downside_only", "true"),
					resource.TestCheckResourceAttr(resourceName, "config.0.recommendations_cpu_percentile", "85"),
					resource.TestCheckResourceAttr(resourceName, "config.0.recommendations_memory_percentile", "85"),
				),
			},
			{
				PreConfig: func() {
					log.Printf("Waiting 8 minutes before update step...")
					time.Sleep(8 * time.Minute)
				},
				Config: createOceanRightSizingClusterConfigTerraform(&RightSizingClusterConfigMetadata{
					oceanID:              "o-8b34732f",
					clusterIdentifier:    "ocean.k8s",
					config:               testRightSizingClusterConfigBaseline_Update,
					updateBaselineFields: true,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanRightSizingClusterConfigExists(&config, resourceName),
					resource.TestCheckResourceAttr(resourceName, "ocean_id", "o-8b34732f"),
					resource.TestCheckResourceAttr(resourceName, "cluster_identifier", "ocean.k8s"),
					resource.TestCheckResourceAttr(resourceName, "config.0.adjust_limit_on_downsize", "false"),
					resource.TestCheckResourceAttr(resourceName, "config.0.downside_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "config.0.recommendations_cpu_percentile", "99"),
					resource.TestCheckResourceAttr(resourceName, "config.0.recommendations_memory_percentile", "100"),
				),
			},
		},
	})
}
