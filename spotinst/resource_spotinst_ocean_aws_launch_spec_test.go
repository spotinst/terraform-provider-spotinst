package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws_launch_configuration"
	"log"
	"testing"
)

//func init() {
//	resource.AddTestSweepers("spotinst_ocean_aws_launch_spec", &resource.Sweeper{
//		Name: "spotinst_ocean_aws_launch_spec",
//		F:    testSweepOceanAWSLaunchSpec,
//	})
//}
//
//func testSweepOceanAWSLaunchSpec(region string) error {
//	client, err := getProviderClient("aws")
//	if err != nil {
//		return fmt.Errorf("error getting client: %v", err)
//	}
//
//	conn := client.(*Client).ocean.CloudProviderAWS()
//	input := &aws.ListLaunchSpecsInput{}
//	if resp, err := conn.ListLaunchSpecs(context.Background(), input); err != nil {
//		return fmt.Errorf("error getting list of launch specs to sweep")
//	} else {
//		if len(resp.LaunchSpecs) == 0 {
//			log.Printf("[INFO] No launch specs to sweep")
//		}
//		for _, launchSpec := range resp.LaunchSpecs {
//			if strings.Contains(spotinst.StringValue(launchSpec.<WHAT>), "test-acc-") {
//				if _, err := conn.DeleteLaunchSpec(context.Background(), &aws.DeleteLaunchSpecInput{LaunchSpecID: launchSpec.ID}); err != nil {
//					return fmt.Errorf("unable to delete launch spec %v in sweep", spotinst.StringValue(launchSpec.ID))
//				} else {
//					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(launchSpec.ID))
//				}
//			}
//		}
//	}
//	return nil
//}

func createOceanAWSLaunchSpecResourceOceanID(oceanID string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanAWSLaunchSpecResourceName), oceanID)
}

func testOceanAWSLaunchSpecDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OceanAWSLaunchSpecResourceName) {
			continue
		}
		input := &aws.ReadLaunchSpecInput{LaunchSpecID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAWS().ReadLaunchSpec(context.Background(), input)
		if err == nil && resp != nil && resp.LaunchSpec != nil {
			return fmt.Errorf("launchSpec still exists")
		}
	}
	return nil
}

func testCheckOceanAWSLaunchSpecAttributes(launchSpec *aws.LaunchSpec, expectedID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(launchSpec.OceanID) != expectedID {
			return fmt.Errorf("bad content: %v", launchSpec.OceanID)
		}
		return nil
	}
}

func testCheckOceanAWSLaunchSpecExists(launchSpec *aws.LaunchSpec, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &aws.ReadLaunchSpecInput{LaunchSpecID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAWS().ReadLaunchSpec(context.Background(), input)
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

type LaunchSpecConfigMetadata struct {
	provider             string
	oceanID              string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createOceanAWSLaunchSpecTerraform(lscm *LaunchSpecConfigMetadata) string {
	if lscm == nil {
		return ""
	}

	if lscm.provider == "" {
		lscm.provider = "aws"
	}

	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`

	if lscm.updateBaselineFields {
		format := testBaselineOceanAWSLaunchSpecConfig_Update
		template += fmt.Sprintf(format,
			lscm.oceanID,
			lscm.provider,
			lscm.oceanID,
			lscm.fieldsToAppend,
		)
	} else {
		format := testBaselineOceanAWSLaunchSpecConfig_Create
		template += fmt.Sprintf(format,
			lscm.oceanID,
			lscm.provider,
			lscm.oceanID,
			lscm.fieldsToAppend,
		)
	}

	log.Printf("Terraform LaunchSpec template:\n%v", template)
	return template
}

// region OceanAWSLaunchSpec: Baseline
func TestAccSpotinstOceanAWSLaunchSpec_Baseline(t *testing.T) {
	oceanID := "o-4bc9b7d9"
	resourceName := createOceanAWSLaunchSpecResourceOceanID(oceanID)

	var launchSpec aws.LaunchSpec
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSLaunchSpecDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID: oceanID,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-79826301"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_aws_launch_configuration.Base64StateFunc("hello world")),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "labels.72815409.key", "label key"),
					resource.TestCheckResourceAttr(resourceName, "labels.72815409.value", "label value"),
					resource.TestCheckResourceAttr(resourceName, "taints.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "taints.1785420166.key", "taint key"),
					resource.TestCheckResourceAttr(resourceName, "taints.1785420166.value", "taint value"),
					resource.TestCheckResourceAttr(resourceName, "taints.1785420166.effect", "NoSchedule"),
				),
			},
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-79826301"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_aws_launch_configuration.Base64StateFunc("hello world updated")),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "updated"),
					resource.TestCheckResourceAttr(resourceName, "labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "labels.3686834679.key", "label key updated"),
					resource.TestCheckResourceAttr(resourceName, "labels.3686834679.value", "label value updated"),
					resource.TestCheckResourceAttr(resourceName, "taints.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "taints.4133802144.key", "taint key updated"),
					resource.TestCheckResourceAttr(resourceName, "taints.4133802144.value", "taint value updated"),
					resource.TestCheckResourceAttr(resourceName, "taints.4133802144.effect", "NoExecute"),
				),
			},
		},
	})
}

const testBaselineOceanAWSLaunchSpecConfig_Create = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
  provider = "%v"  

  ocean_id = "%v"
  image_id = "ami-79826301"
  user_data = "hello world"
  iam_instance_profile = "test"

  labels {
    key = "label key"
    value = "label value"
  }

  taints {
    key = "taint key"
    value = "taint value"
    effect = "NoSchedule"
  }

 %v
}
`

const testBaselineOceanAWSLaunchSpecConfig_Update = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
  provider = "%v"

  ocean_id = "%v"
  image_id = "ami-79826301"
  user_data = "hello world updated"
  iam_instance_profile = "updated"
  
  labels {
    key = "label key updated"
    value = "label value updated"
  }

  taints {
    key = "taint key updated"
    value = "taint value updated"
    effect = "NoExecute"
  }

%v
}
`

// endregion
