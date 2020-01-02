package spotinst

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//func init() {
//	resource.AddTestSweepers("spotinst_ocean_gke_launch_spec", &resource.Sweeper{
//		Name: "spotinst_ocean_gke_launch_spec",
//		F:    testSweepOceanGKELaunchSpec,
//	})
//}
//
//func testSweepOceanGKELaunchSpec(region string) error {
//	client, err := getProviderClient("gcp")
//	if err != nil {
//		return fmt.Errorf("error getting client: %v", err)
//	}
//
//	conn := client.(*Client).ocean.CloudProviderGCP()
//	input := &gcp.ListLaunchSpecsInput{}
//	if resp, err := conn.ListLaunchSpecs(context.Background(), input); err != nil {
//		return fmt.Errorf("error getting list of launch specs to sweep")
//	} else {
//		if len(resp.LaunchSpecs) == 0 {
//			log.Printf("[INFO] No launch specs to sweep")
//		}
//		for _, launchSpec := range resp.LaunchSpecs {
//			if strings.Contains(spotinst.StringValue(launchSpec.<WHAT>), "test-acc-") {
//				if _, err := conn.DeleteLaunchSpec(context.Background(), &gcp.DeleteLaunchSpecInput{LaunchSpecID: launchSpec.ID}); err != nil {
//					return fmt.Errorf("unable to delete launch spec %v in sweep", spotinst.StringValue(launchSpec.ID))
//				} else {
//					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(launchSpec.ID))
//				}
//			}
//		}
//	}
//	return nil
//}

func createOceanGKELaunchSpecResource(oceanID string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanGKELaunchSpecResourceName), oceanID)
}

func testOceanGKELaunchSpecDestroy(s *terraform.State) error {
	client := testAccProviderGCP.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OceanGKELaunchSpecResourceName) {
			continue
		}

		input := &gcp.ReadLaunchSpecInput{LaunchSpecID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderGCP().ReadLaunchSpec(context.Background(), input)

		if err == nil && resp != nil && resp.LaunchSpec != nil {
			return fmt.Errorf("launchSpec still exists")
		}
	}

	return nil
}

func testCheckOceanGKELaunchSpecAttributes(launchSpec *gcp.LaunchSpec, expectedID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(launchSpec.OceanID) != expectedID {
			return fmt.Errorf("bad content: %v", launchSpec.OceanID)
		}
		return nil
	}
}

func testCheckOceanGKELaunchSpecExists(launchSpec *gcp.LaunchSpec, resourceName string) resource.TestCheckFunc {
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

type GKELaunchSpecConfigMetadata struct {
	provider             string
	oceanID              string
	updateBaselineFields bool
}

func createOceanGKELaunchSpecTerraform(lscm *GKELaunchSpecConfigMetadata, formatToUse string) string {
	if lscm == nil {
		return ""
	}

	if lscm.provider == "" {
		lscm.provider = "gcp"
	}

	template :=
		`provider "gcp" {
	 token   = "fake"
	 account = "fake"
	}
	`

	format := formatToUse

	if lscm.updateBaselineFields {
		template += fmt.Sprintf(format,
			lscm.oceanID,
			lscm.provider,
			lscm.oceanID,
		)
	} else {
		template += fmt.Sprintf(format,
			lscm.oceanID,
			lscm.provider,
			lscm.oceanID,
		)
	}

	log.Printf("Terraform LaunchSpec template:\n%v", template)

	return template
}

// region OceanGKELaunchSpec: Baseline
func TestAccSpotinstOceanGKELaunchSpec_Baseline(t *testing.T) {
	oceanID := "o-431d5b4a"
	resourceName := createOceanGKELaunchSpecResource(oceanID)

	var launchSpec gcp.LaunchSpec
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "gcp") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanGKELaunchSpecDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanGKELaunchSpecTerraform(&GKELaunchSpecConfigMetadata{oceanID: oceanID, updateBaselineFields: true}, testBaselineOceanGKELaunchSpecConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKELaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanGKELaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "source_image", "https://www.googleapis.com/compute/v1/projects/gke-node-images/global/images/gke-1118-gke6-cos-69-10895-138-0-v190330-pre"),
					resource.TestCheckResourceAttr(resourceName, "metadata.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "metadata.1537634279.key", "gci-update-strategy"),
					resource.TestCheckResourceAttr(resourceName, "metadata.1537634279.value", "update_disabled"),
					resource.TestCheckResourceAttr(resourceName, "metadata.2322445456.key", "gci-ensure-gke-docker"),
					resource.TestCheckResourceAttr(resourceName, "metadata.2322445456.value", "true"),
				),
			},
			{
				Config: createOceanGKELaunchSpecTerraform(&GKELaunchSpecConfigMetadata{oceanID: oceanID, updateBaselineFields: true}, testBaselineOceanGKELaunchSpecConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKELaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanGKELaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "source_image", "https://www.googleapis.com/compute/v1/projects/gke-node-images/global/images/gke-1118-gke6-cos-69-10895-138-0-v190330-pre"),
					resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.1537634279.key", "gci-update-strategy"),
					resource.TestCheckResourceAttr(resourceName, "metadata.1537634279.value", "update_disabled"),
				),
			},
		},
	})
}

const testBaselineOceanGKELaunchSpecConfig_Create = `
resource "` + string(commons.OceanGKELaunchSpecResourceName) + `" "%v" {
 provider = "%v"  

 ocean_id = "%v"
 source_image = "https://www.googleapis.com/compute/v1/projects/gke-node-images/global/images/gke-1118-gke6-cos-69-10895-138-0-v190330-pre"

 metadata {
     key = "gci-update-strategy"
     value = "update_disabled"
   }
 
 metadata {
     key = "gci-ensure-gke-docker"
     value = "true"
   }
 
 labels {
     key = "testKey2"
     value = "testVal2"
   }
 

 taints {
     key = "testTaintKey"
     value = "testTaintVal"
     effect = "NoSchedule"
   }
}

`

const testBaselineOceanGKELaunchSpecConfig_Update = `
resource "` + string(commons.OceanGKELaunchSpecResourceName) + `" "%v" {
 provider = "%v"

 ocean_id = "%v"
 source_image = "https://www.googleapis.com/compute/v1/projects/gke-node-images/global/images/gke-1118-gke6-cos-69-10895-138-0-v190330-pre"

 metadata {
     key = "gci-update-strategy"
     value = "update_disabled"
   }

 taints {
     key = "testTaintKey"
     value = "testTaintVal"
     effect = "NoSchedule"
   }

 taints {
     key = "testTaintKey2"
     value = "testTaintVal2"
     effect = "NoSchedule"
   }
}

`

// endregion

// region OceanGKELaunchSpec: Labels
func TestAccSpotinstOceanGKELaunchSpec_Labels(t *testing.T) {
	oceanID := "o-431d5b4a"
	resourceName := createOceanGKELaunchSpecResource(oceanID)

	var launchSpec gcp.LaunchSpec
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "gcp") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanGKELaunchSpecDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanGKELaunchSpecTerraform(&GKELaunchSpecConfigMetadata{oceanID: oceanID}, testLabelsOceanGKELaunchSpecConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKELaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanGKELaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "labels.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "labels.2480521734.key", "testKey2"),
					resource.TestCheckResourceAttr(resourceName, "labels.2480521734.value", "testVal2"),
					resource.TestCheckResourceAttr(resourceName, "labels.940728818.key", "testKey"),
					resource.TestCheckResourceAttr(resourceName, "labels.940728818.value", "testVal"),
				),
			},
			{
				Config: createOceanGKELaunchSpecTerraform(&GKELaunchSpecConfigMetadata{oceanID: oceanID, updateBaselineFields: true}, testLabelsOceanGKELaunchSpecConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKELaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanGKELaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "labels.940728818.key", "testKey"),
					resource.TestCheckResourceAttr(resourceName, "labels.940728818.value", "testVal"),
				),
			},
		},
	})
}

const testLabelsOceanGKELaunchSpecConfig_Create = `
resource "` + string(commons.OceanGKELaunchSpecResourceName) + `" "%v" {
 provider = "%v"  

 ocean_id = "%v"
 source_image = "https://www.googleapis.com/compute/v1/projects/gke-node-images/global/images/gke-1118-gke6-cos-69-10895-138-0-v190330-pre"

 metadata {
     key = "gci-update-strategy"
     value = "update_disabled"
   }

 labels {
     key = "testKey"
     value = "testVal"
   }

 labels {
     key = "testKey2"
     value = "testVal2"
   }
}

`

const testLabelsOceanGKELaunchSpecConfig_Update = `
resource "` + string(commons.OceanGKELaunchSpecResourceName) + `" "%v" {
 provider = "%v"

 ocean_id = "%v"
 source_image = "https://www.googleapis.com/compute/v1/projects/gke-node-images/global/images/gke-1118-gke6-cos-69-10895-138-0-v190330-pre"

 metadata {
     key = "gci-update-strategy"
     value = "update_disabled"
   }

 metadata {
     key = "gci-ensure-gke-docker"
     value = "true"
   }

 labels {
     key = "testKey"
     value = "testVal"
   }
}

`

//endregion

// region OceanGKELaunchSpec: Taints
func TestAccSpotinstOceanGKELaunchSpec_Taints(t *testing.T) {
	oceanID := "o-431d5b4a"
	resourceName := createOceanGKELaunchSpecResource(oceanID)

	var launchSpec gcp.LaunchSpec
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "gcp") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanGKELaunchSpecDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanGKELaunchSpecTerraform(&GKELaunchSpecConfigMetadata{oceanID: oceanID}, testTaintsOceanGKELaunchSpecConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKELaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanGKELaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "taints.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "taints.1137791537.key", "testTaintKey"),
					resource.TestCheckResourceAttr(resourceName, "taints.1137791537.value", "testTaintVal"),
					resource.TestCheckResourceAttr(resourceName, "taints.1137791537.effect", "NoSchedule"),
				),
			},
			{
				Config: createOceanGKELaunchSpecTerraform(&GKELaunchSpecConfigMetadata{oceanID: oceanID, updateBaselineFields: true}, testTaintsOceanGKELaunchSpecConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKELaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanGKELaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "taints.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "taints.1137791537.key", "testTaintKey"),
					resource.TestCheckResourceAttr(resourceName, "taints.1137791537.value", "testTaintVal"),
					resource.TestCheckResourceAttr(resourceName, "taints.1137791537.effect", "NoSchedule"),
					resource.TestCheckResourceAttr(resourceName, "taints.1372190553.key", "testTaintKey2"),
					resource.TestCheckResourceAttr(resourceName, "taints.1372190553.value", "testTaintVal2"),
					resource.TestCheckResourceAttr(resourceName, "taints.1372190553.effect", "NoSchedule"),
				),
			},
		},
	})
}

const testTaintsOceanGKELaunchSpecConfig_Create = `
resource "` + string(commons.OceanGKELaunchSpecResourceName) + `" "%v" {
 provider = "%v"  

 ocean_id = "%v"
 source_image = "https://www.googleapis.com/compute/v1/projects/gke-node-images/global/images/gke-1118-gke6-cos-69-10895-138-0-v190330-pre"

 metadata {
     key = "gci-update-strategy"
     value = "update_disabled"
   }

 taints {
     key = "testTaintKey"
     value = "testTaintVal"
     effect = "NoSchedule"
   }
}

`

const testTaintsOceanGKELaunchSpecConfig_Update = `
resource "` + string(commons.OceanGKELaunchSpecResourceName) + `" "%v" {
 provider = "%v"

 ocean_id = "%v"
 source_image = "https://www.googleapis.com/compute/v1/projects/gke-node-images/global/images/gke-1118-gke6-cos-69-10895-138-0-v190330-pre"

 metadata {
     key = "gci-update-strategy"
     value = "update_disabled"
   }

 metadata {
     key = "gci-ensure-gke-docker"
     value = "true"
   }

 taints {
     key = "testTaintKey"
     value = "testTaintVal"
     effect = "NoSchedule"
   }

 taints {
     key = "testTaintKey2"
     value = "testTaintVal2"
     effect = "NoSchedule"
   }
}

`

//endregion

// region OceanGKELaunchSpec: AutoScale
func TestAccSpotinstOceanGKELaunchSpec_AutoScale(t *testing.T) {
	oceanID := "o-431d5b4a"
	resourceName := createOceanGKELaunchSpecResource(oceanID)

	var launchSpec gcp.LaunchSpec
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "gcp") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanGKELaunchSpecDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanGKELaunchSpecTerraform(&GKELaunchSpecConfigMetadata{oceanID: oceanID}, testAutoScaleOceanGKELaunchSpecConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKELaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanGKELaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.3279616137.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.3279616137.gpu_per_unit", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.3279616137.num_of_units", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.3279616137.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.4058284811.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.4058284811.gpu_per_unit", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.4058284811.num_of_units", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.4058284811.memory_per_unit", "256"),
				),
			},
			{
				Config: createOceanGKELaunchSpecTerraform(&GKELaunchSpecConfigMetadata{oceanID: oceanID, updateBaselineFields: true}, testAutoScaleOceanGKELaunchSpecConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKELaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanGKELaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.3279616137.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.3279616137.gpu_per_unit", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.3279616137.num_of_units", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.3279616137.memory_per_unit", "512"),
				),
			},
			{
				Config: createOceanGKELaunchSpecTerraform(&GKELaunchSpecConfigMetadata{oceanID: oceanID, updateBaselineFields: true}, testAutoScaleOceanGKELaunchSpecConfig_Delete),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanGKELaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanGKELaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.#", "0"),
				),
			},
		},
	})
}

const testAutoScaleOceanGKELaunchSpecConfig_Create = `
resource "` + string(commons.OceanGKELaunchSpecResourceName) + `" "%v" {
 provider = "%v"  

 ocean_id = "%v"
 source_image = "https://www.googleapis.com/compute/v1/projects/gke-node-images/global/images/gke-1118-gke6-cos-69-10895-138-0-v190330-pre"

 metadata {
   key = "gci-update-strategy"
   value = "update_disabled"
 }

 metadata {
   key = "gci-ensure-gke-docker"
   value = "true"
 }

 autoscale_headrooms {
   cpu_per_unit = 1024
   gpu_per_unit = 1
   memory_per_unit = 512
   num_of_units = 1
 }
   
 autoscale_headrooms {
   cpu_per_unit = 1024
   gpu_per_unit = 1
   memory_per_unit = 256
   num_of_units = 1
 }
}

`

const testAutoScaleOceanGKELaunchSpecConfig_Update = `
resource "` + string(commons.OceanGKELaunchSpecResourceName) + `" "%v" {
 provider = "%v"

 ocean_id = "%v"
 source_image = "https://www.googleapis.com/compute/v1/projects/gke-node-images/global/images/gke-1118-gke6-cos-69-10895-138-0-v190330-pre"

 metadata {
   key = "gci-update-strategy"
   value = "update_disabled"
 }

 metadata {
   key = "gci-ensure-gke-docker"
   value = "true"
 }

 autoscale_headrooms {
   cpu_per_unit = 1024
   gpu_per_unit = 1
   memory_per_unit = 512
   num_of_units = 1
 }
}

`

const testAutoScaleOceanGKELaunchSpecConfig_Delete = `
resource "` + string(commons.OceanGKELaunchSpecResourceName) + `" "%v" {
 provider = "%v"

 ocean_id = "%v"
 source_image = "https://www.googleapis.com/compute/v1/projects/gke-node-images/global/images/gke-1118-gke6-cos-69-10895-138-0-v190330-pre"

 metadata {
   key = "gci-update-strategy"
   value = "update_disabled"
 }

 metadata {
   key = "gci-ensure-gke-docker"
   value = "true"
 }
}

`

//endregion
