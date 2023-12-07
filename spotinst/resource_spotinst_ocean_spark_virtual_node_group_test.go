package spotinst

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/spark"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

const (
	testOceanSparkClusterID = "osc-3dda7ea4"
	testOceanDedicatedVngID = "ols-24b985c5"
)

var oceanSparkClusterID = getOceanSparkClusterID() // NOTE: This needs to be an existing ocean cluster
var oceanSparkVngID = getOceanSparkVngID()         // NOTE: This needs to be an existing ocean VNG

func init() {
	resource.AddTestSweepers("spotinst_ocean_spark_virtual_node_group", &resource.Sweeper{
		Name: "spotinst_ocean_spark_virtual_node_group",
		F:    testSweepOceanSparkVng,
	})
}

func getOceanSparkClusterID() string {
	// Prefer environment variable
	oceanSparkClusterID := os.Getenv("TEST_ACC_OCEAN_SPARK_CLUSTER_ID")
	if oceanSparkClusterID == "" {
		// Default to hardcoded ID
		oceanSparkClusterID = testOceanSparkClusterID
	}

	return oceanSparkClusterID
}

func getOceanSparkVngID() string {
	// Prefer environment variable
	oceanDedicatedVngID := os.Getenv("TEST_ACC_OCEAN_SPARK_DEDICATED_VNG_ID")
	if oceanDedicatedVngID == "" {
		// Default to hardcoded ID
		oceanDedicatedVngID = testOceanDedicatedVngID
	}

	return oceanDedicatedVngID
}

func testSweepOceanSparkVng(_ string) error {
	client, err := getProviderClient("aws")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).ocean.Spark()
	input := &spark.ListVngsInput{ClusterID: spotinst.String(oceanSparkClusterID)}
	if resp, err := conn.ListVirtualNodeGroups(context.Background(), input); err != nil {
		return fmt.Errorf("error getting VNGs to sweep")
	} else {
		if len(resp.VirtualNodeGroups) == 0 {
			log.Printf("[INFO] No VNGs to sweep")
		}
		for _, vng := range resp.VirtualNodeGroups {
			if strings.Compare(spotinst.StringValue(vng.VngID), oceanSparkVngID) == 0 {
				if _, err := conn.DetachVirtualNodeGroup(context.Background(), &spark.DetachVngInput{ClusterID: spotinst.String(oceanSparkClusterID), VngID: vng.VngID}); err != nil {
					return fmt.Errorf("unable to detach VNG %v in sweep, %w", spotinst.StringValue(vng.VngID), err)
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(vng.VngID))
				}
			}
		}
	}
	return nil
}

func testOceanSparkVngDetach(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OceanSparkVirtualNodeGroupResourceName) {
			continue
		}
		input := &spark.ListVngsInput{ClusterID: spotinst.String(rs.Primary.Attributes["ocean_spark_cluster_id"])}
		resp, err := client.ocean.Spark().ListVirtualNodeGroups(context.Background(), input)
		if err == nil && resp != nil && resp.VirtualNodeGroups != nil {
			for i := range resp.VirtualNodeGroups {
				if spotinst.StringValue(resp.VirtualNodeGroups[i].VngID) == rs.Primary.ID {
					return fmt.Errorf("VNG still attached")
				}
			}
		}
	}
	return nil
}

func testCheckOceanSparkVngAttached(vng *spark.DedicatedVirtualNodeGroup, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &spark.ListVngsInput{ClusterID: spotinst.String(rs.Primary.Attributes["ocean_spark_cluster_id"])}
		resp, err := client.ocean.Spark().ListVirtualNodeGroups(context.Background(), input)
		if err != nil {
			return err
		}

		for i := range resp.VirtualNodeGroups {
			if spotinst.StringValue(resp.VirtualNodeGroups[i].VngID) == rs.Primary.ID {
				*vng = *resp.VirtualNodeGroups[i]
				return nil
			}
		}

		return fmt.Errorf("VNG not found: %+v,\n %+v\n", resp.VirtualNodeGroups, rs.Primary.Attributes)
	}
}

type SparkClusterVngAttachMetadata struct {
	oceanSparkVngID     string
	oceanSparkClusterID string
}

func createOceanSparkVngResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanSparkVirtualNodeGroupResourceName), name)
}

func attachOceanSparkVngTerraform(scvam *SparkClusterVngAttachMetadata) string {
	if scvam == nil {
		return ""
	}

	format := testAttachVngConfig
	template := fmt.Sprintf(format,
		scvam.oceanSparkVngID,
		scvam.oceanSparkVngID,
		scvam.oceanSparkClusterID,
	)

	log.Printf("Terraform [%v] template:\n%v", scvam.oceanSparkVngID, template)
	return template
}

func TestAccSpotinstOceanSparkVng_attach(t *testing.T) {
	resourceName := createOceanSparkVngResourceName(oceanSparkVngID)

	var vng spark.DedicatedVirtualNodeGroup
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanSparkVngDetach,

		Steps: []resource.TestStep{
			{
				Config: attachOceanSparkVngTerraform(&SparkClusterVngAttachMetadata{
					oceanSparkVngID:     oceanSparkVngID,
					oceanSparkClusterID: oceanSparkClusterID,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanSparkVngAttached(&vng, resourceName),
				),
			},
		},
	})
}

const testAttachVngConfig = `
resource "` + string(commons.OceanSparkVirtualNodeGroupResourceName) + `" "%v" {
  provider = "aws"

  virtual_node_group_id = "%v"
  ocean_spark_cluster_id = "%v"
}
`
