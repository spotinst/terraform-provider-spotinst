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
	testOceanSparkClusterID = "osc-d552c5b5"
	testOceanDedicatedVngID = "ols-b66e444f"
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
	oceanDedicatedVngID := os.Getenv("TEST_ACC_OCEAN_DEDICATED_VNG_ID")
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
	input := &spark.ListVngsInput{}
	input.ClusterID = spotinst.String(oceanSparkClusterID)
	if resp, err := conn.ListVirtualNodeGroups(context.Background(), input); err != nil {
		return fmt.Errorf("error getting VNGs to sweep")
	} else {
		if len(resp.VirtualNodeGroups) == 0 {
			log.Printf("[INFO] No VNGs to sweep")
		}
		for _, vng := range resp.VirtualNodeGroups {
			// Our test clusters should have a controller cluster ID starting with "tf-test-acc-"
			if strings.Compare(spotinst.StringValue(vng.VngID), oceanSparkVngID) == 0 {
				if _, err := conn.DetachVirtualNodeGroup(context.Background(), &spark.DetachVngInput{ClusterID: spotinst.String(oceanSparkClusterID), VngID: vng.VngID}); err != nil {
					return fmt.Errorf("unable to detach VNG %v in sweep", spotinst.StringValue(vng.VngID))
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
		input := &spark.ListVngsInput{ClusterID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.Spark().ListVirtualNodeGroups(context.Background(), input)
		if err == nil && resp != nil && resp.VirtualNodeGroups != nil {
			return fmt.Errorf("VNG still attached")
		}
	}
	return nil
}

func testCheckOceanSparkVngAttached(cluster *spark.Cluster, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &spark.ReadClusterInput{ClusterID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.Spark().ReadCluster(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.Cluster.OceanClusterID) != rs.Primary.Attributes["ocean_cluster_id"] {
			return fmt.Errorf("Cluster not found: %+v,\n %+v\n", resp.Cluster, rs.Primary.Attributes)
		}
		*cluster = *resp.Cluster
		return nil
	}
}

type SparkClusterVngAttachMetadata struct {
	oceanSparkVngID     string
	oceanSparkClusterID string
}

func createOceanSparkVngResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanSparkResourceName), name)
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

	var cluster spark.Cluster
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
					testCheckOceanSparkVngAttached(&cluster, resourceName),
					resource.TestCheckResourceAttr(resourceName, "webhook.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "webhook.0.ocean_spark_cluster_id", testOceanSparkClusterID),
					resource.TestCheckResourceAttr(resourceName, "webhook.0.virtual_node_group_id", testOceanDedicatedVngID),
				),
			},
		},
	})
}

const testAttachVngConfig = `
resource "` + string(commons.OceanSparkVirtualNodeGroupResourceName) + `" "%v" {
  virtual_node_group_id = "%v"
  ocean_spark_cluster_id = "%v"
}
`
