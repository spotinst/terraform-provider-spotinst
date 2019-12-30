package spotinst

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws_launch_configuration"
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

func createOceanAWSLaunchSpecTerraform(lscm *LaunchSpecConfigMetadata, formatToUse string) string {
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

	format := formatToUse

	if lscm.updateBaselineFields {
		template += fmt.Sprintf(format,
			lscm.oceanID,
			lscm.provider,
			lscm.oceanID,
			lscm.fieldsToAppend,
		)
	} else {
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
				}, testBaselineOceanAWSLaunchSpecConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-79826301"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-0041bd3fd6aa2ee3c"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_aws_launch_configuration.Base64StateFunc("hello world")),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "labels.72815409.key", "label key"),
					resource.TestCheckResourceAttr(resourceName, "labels.72815409.value", "label value"),
					resource.TestCheckResourceAttr(resourceName, "taints.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "taints.1785420166.key", "taint key"),
					resource.TestCheckResourceAttr(resourceName, "taints.1785420166.value", "taint value"),
					resource.TestCheckResourceAttr(resourceName, "taints.1785420166.effect", "NoSchedule"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.0", "subnet-7f3fbf06"),
					resource.TestCheckResourceAttr(resourceName, "root_volume_size", "20"),
				),
			},
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}, testBaselineOceanAWSLaunchSpecConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-79826301"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-0041bd3fd6aa2ee3c"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.1", "sg-0195f2ac3a6014a15"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_aws_launch_configuration.Base64StateFunc("hello world updated")),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "updated"),
					resource.TestCheckResourceAttr(resourceName, "labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "labels.3686834679.key", "label key updated"),
					resource.TestCheckResourceAttr(resourceName, "labels.3686834679.value", "label value updated"),
					resource.TestCheckResourceAttr(resourceName, "taints.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "taints.4133802144.key", "taint key updated"),
					resource.TestCheckResourceAttr(resourceName, "taints.4133802144.value", "taint value updated"),
					resource.TestCheckResourceAttr(resourceName, "taints.4133802144.effect", "NoExecute"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.0", "subnet-7f3fbf06"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.1", "subnet-03b7ed5b"),
					resource.TestCheckResourceAttr(resourceName, "root_volume_size", "30"),
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
  security_groups = ["sg-0041bd3fd6aa2ee3c"]
  user_data = "hello world"
  iam_instance_profile = "test"
  subnet_ids = ["subnet-7f3fbf06"]
  root_volume_size = 20
  
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
  subnet_ids = ["subnet-7f3fbf06", "subnet-03b7ed5b"]
  security_groups = ["sg-0041bd3fd6aa2ee3c","sg-0195f2ac3a6014a15" ]
  root_volume_size = 30

  
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

// region OceanAWSLaunchSpec: AutoScale
func TestAccSpotinstOceanAWSLaunchSpec_AutoScale(t *testing.T) {
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
				}, testAutoScaleOceanAWSLaunchSpecConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
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
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}, testAutoScaleOceanAWSLaunchSpecConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.3279616137.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.3279616137.gpu_per_unit", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.3279616137.num_of_units", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.3279616137.memory_per_unit", "512"),
				),
			},
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}, testAutoScaleOceanAWSLaunchSpecConfig_Delete),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.#", "0"),
				),
			},
		},
	})
}

const testAutoScaleOceanAWSLaunchSpecConfig_Create = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"  
 ocean_id = "%v"

 image_id = "ami-79826301"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"

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
%v
}

`

const testAutoScaleOceanAWSLaunchSpecConfig_Update = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"  
 ocean_id = "%v"
 
 image_id = "ami-79826301"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"

 autoscale_headrooms {
   cpu_per_unit = 1024
   gpu_per_unit = 1
   memory_per_unit = 512
   num_of_units = 1
 }
%v
}

`

const testAutoScaleOceanAWSLaunchSpecConfig_Delete = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"
 ocean_id = "%v"
 
 image_id = "ami-79826301"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
%v
}

`

//endregion
