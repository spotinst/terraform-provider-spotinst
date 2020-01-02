package spotinst

import (
	"context"
	"fmt"
	"testing"

	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//func init() {
//	resource.AddTestSweepers("resource_spotinst_ocean_gke_launch_spec_import", &resource.Sweeper{
//		Name: "resource_spotinst_ocean_gke_import",
//		F:    testSweepOceanGKELaunchSpecImport,
//	})
//}
//
//func testSweepOceanGKELaunchSpecImport(region string) error {
//	client, err := getProviderClient("gcp")
//	if err != nil {
//		return fmt.Errorf("error getting client: %v", err)
//	}
//
//	conn := client.(*Client).ocean.CloudProviderGCP()
//
//	input := &gcp.ListLaunchSpecsInput{}
//	if resp, err := conn.ListLaunchSpecs(context.Background(), input); err != nil {
//		return fmt.Errorf("error getting list of clusters to sweep")
//	} else {
//		if len(resp.LaunchSpecs) == 0 {
//			log.Printf("[INFO] No clusters to sweep")
//		}
//
//		for _, launchSpec := range resp.LaunchSpecs {
//			if strings.Contains(spotinst.StringValue(launchSpec.ID), "terraform-acc-tests-") {
//				if _, err := conn.DeleteLaunchSpec(context.Background(), &gcp.DeleteLaunchSpecInput{LaunchSpecID: launchSpec.ID}); err != nil {
//					return fmt.Errorf("unable to delete group %v in sweep", spotinst.StringValue(launchSpec.ID))
//				} else {
//					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(launchSpec.ID))
//				}
//			}
//		}
//	}
//	return nil
//}

func createOceanGKELaunchSpecImportResource(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanGKELaunchSpecImportResourceName), name)
}

func testOceanGKELaunchSpecImportDestroy(s *terraform.State) error {
	client := testAccProviderGCP.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OceanGKELaunchSpecImportResourceName) {
			continue
		}
		input := &gcp.ReadLaunchSpecInput{LaunchSpecID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderGCP().ReadLaunchSpec(context.Background(), input)
		if err == nil && resp != nil && resp.LaunchSpec != nil {
			return fmt.Errorf("launch spec still exists")
		}
	}
	return nil
}

func testCheckOceanGKELaunchSpecImportAttributes(launchSpec *gcp.LaunchSpec, expectedID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(launchSpec.OceanID) != expectedID {
			return fmt.Errorf("bad content: %v", launchSpec.OceanID)
		}
		return nil
	}
}

func testCheckOceanGKELaunchSpecImportExists(launchSpec *gcp.LaunchSpec, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderGCP.Meta().(*Client)
		input := &gcp.ReadLaunchSpecInput{LaunchSpecID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderGCP().ReadLaunchSpec(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.LaunchSpec.ID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("LaunchSpec not found: %+v,\n %+v\n", resp.LaunchSpec, rs.Primary.Attributes)
		}
		*launchSpec = *resp.LaunchSpec
		return nil
	}
}

type OceanGKELaunchSpecImportMetadata struct {
	provider             string
	oceanID              string
	updateBaselineFields bool
}

func createOceanGKELaunchSpecImportTerraform(launchSpecMeta *OceanGKELaunchSpecImportMetadata, update string, create string) string {
	if launchSpecMeta == nil {
		return ""
	}

	if launchSpecMeta.provider == "" {
		launchSpecMeta.provider = "gcp"
	}

	template :=
		`provider "gcp" {
	token   = "fake"
	account = "fake"
	}
	`
	if launchSpecMeta.updateBaselineFields {
		format := update

		template += fmt.Sprintf(format,
			launchSpecMeta.oceanID,
			launchSpecMeta.provider,
			launchSpecMeta.oceanID,
		)
	} else {
		format := create
		template += fmt.Sprintf(format,
			launchSpecMeta.oceanID,
			launchSpecMeta.provider,
			launchSpecMeta.oceanID,
		)
	}

	log.Printf("Terraform LaunchSpec template:\n%v", template)

	return template
}

// region Ocean GKE Import: Baseline
func TestAccSpotinstOceanGKELaunchSpecImport_Baseline(t *testing.T) {
	oceanID := "o-c290e75c"
	resourceName := createOceanGKELaunchSpecImportResource(oceanID)

	var launchSpec gcp.LaunchSpec
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "gcp") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanGKELaunchSpecImportDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanGKELaunchSpecImportTerraform(&OceanGKELaunchSpecImportMetadata{oceanID: oceanID}, testBaselineOceanGKELaunchSpecImportConfig_Create, testBaselineOceanGKELaunchSpecImportConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKELaunchSpecImportExists(&launchSpec, resourceName),
					testCheckOceanGKELaunchSpecImportAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "ocean_id", oceanID),
					resource.TestCheckResourceAttr(resourceName, "node_pool_name", "default-pool"),
				),
			},
		},
	})
}

const testBaselineOceanGKELaunchSpecImportConfig_Create = `
resource "` + string(commons.OceanGKELaunchSpecImportResourceName) + `" "%v" {
 provider = "%v"
 ocean_id = "%v"
 node_pool_name = "default-pool"
}

`

// endregion
