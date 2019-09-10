package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_ecs_launch_specification"
	"log"
	"strings"
	"testing"
)

func init() {
	resource.AddTestSweepers("spotinst_ocean_ecs_launch_spec", &resource.Sweeper{
		Name: "spotinst_ocean_ecs_launch_spec",
		F:    testSweepOceanECSLaunchSpec,
	})
}

func testSweepOceanECSLaunchSpec(region string) error {
	client, err := getProviderClient("aws")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).ocean.CloudProviderAWS()
	input := &aws.ListECSLaunchSpecsInput{}
	if resp, err := conn.ListECSLaunchSpecs(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of launch specs to sweep")
	} else {
		if len(resp.LaunchSpecs) == 0 {
			log.Printf("[INFO] No launch specs to sweep")
		}
		for _, launchSpec := range resp.LaunchSpecs {
			if strings.Contains(spotinst.StringValue(launchSpec.Name), "test-acc-") {
				if _, err := conn.DeleteECSLaunchSpec(context.Background(), &aws.DeleteECSLaunchSpecInput{LaunchSpecID: launchSpec.ID}); err != nil {
					return fmt.Errorf("unable to delete launch spec %v in sweep", spotinst.StringValue(launchSpec.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(launchSpec.ID))
				}
			}
		}
	}
	return nil
}

func createOceanECSLaunchSpecResourceOceanName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanECSLaunchSpecResourceName), name)
}

func testOceanECSLaunchSpecDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OceanECSLaunchSpecResourceName) {
			continue
		}
		input := &aws.ReadECSLaunchSpecInput{LaunchSpecID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAWS().ReadECSLaunchSpec(context.Background(), input)
		if err == nil && resp != nil && resp.LaunchSpec != nil {
			return fmt.Errorf("launchSpec still exists")
		}
	}
	return nil
}

func testCheckOceanECSLaunchSpecAttributes(launchSpec *aws.ECSLaunchSpec, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(launchSpec.Name) != name {
			return fmt.Errorf("bad content: %v", launchSpec.Name)
		}
		return nil
	}
}

func testCheckOceanECSLaunchSpecExists(launchSpec *aws.ECSLaunchSpec, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &aws.ReadECSLaunchSpecInput{LaunchSpecID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAWS().ReadECSLaunchSpec(context.Background(), input)
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

type ECSLaunchSpecConfigMetadata struct {
	provider             string
	oceanID              string
	name                 string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createOceanECSLaunchSpecTerraform(lscm *ECSLaunchSpecConfigMetadata) string {
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
		format := testBaselineOceanECSLaunchSpecConfig_Update
		template += fmt.Sprintf(format,
			lscm.name,
			lscm.provider,
			lscm.name,
			lscm.oceanID,
			lscm.fieldsToAppend,
		)
	} else {
		format := testBaselineOceanECSLaunchSpecConfig_Create
		template += fmt.Sprintf(format,
			lscm.name,
			lscm.provider,
			lscm.name,
			lscm.oceanID,
			lscm.fieldsToAppend,
		)
	}

	log.Printf("Terraform LaunchSpec template:\n%v", template)
	return template
}

// region OceanECSLaunchSpec: Baseline
func TestAccSpotinstOceanECSLaunchSpec_Baseline(t *testing.T) {
	oceanID := "o-92189543"
	launchSpecName := "test-acc-ocean-ecs-launch-spec"
	resourceName := createOceanECSLaunchSpecResourceOceanName(launchSpecName)

	var launchSpec aws.ECSLaunchSpec
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanECSLaunchSpecDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanECSLaunchSpecTerraform(&ECSLaunchSpecConfigMetadata{
					oceanID: oceanID,
					name:    launchSpecName,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanECSLaunchSpecAttributes(&launchSpec, launchSpecName),
					resource.TestCheckResourceAttr(resourceName, "name", launchSpecName),
					resource.TestCheckResourceAttr(resourceName, "user_data", ocean_ecs_launch_specification.Base64StateFunc("hello world")),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-082b5a644766e0e6f"),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "ecsInstanceRole"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.0", "awseb-e-sznmxim22e-stack-AWSEBSecurityGroup-10FZKNGB09G1W"),
					resource.TestCheckResourceAttr(resourceName, "attributes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "attributes.3334082635.key", "key"),
					resource.TestCheckResourceAttr(resourceName, "attributes.3334082635.value", "value"),
				),
			},
			{
				Config: createOceanECSLaunchSpecTerraform(&ECSLaunchSpecConfigMetadata{
					oceanID:              oceanID,
					name:                 launchSpecName,
					updateBaselineFields: true}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanECSLaunchSpecAttributes(&launchSpec, launchSpecName),
					resource.TestCheckResourceAttr(resourceName, "name", launchSpecName),
					resource.TestCheckResourceAttr(resourceName, "user_data", ocean_ecs_launch_specification.Base64StateFunc("hello world updated")),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-082b5a644766e0e6f"),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "ecsInstanceRole"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "attributes.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "attributes.3334082635.key", "key"),
					resource.TestCheckResourceAttr(resourceName, "attributes.3334082635.value", "value"),
					resource.TestCheckResourceAttr(resourceName, "attributes.3254712145.key", "key2"),
					resource.TestCheckResourceAttr(resourceName, "attributes.3254712145.value", "value2"),
				),
			},
		},
	})
}

const testBaselineOceanECSLaunchSpecConfig_Create = `
resource "` + string(commons.OceanECSLaunchSpecResourceName) + `" "%v" {
  provider = "%v"  
  name = "%v"
  ocean_id = "%v"

  user_data = "hello world"
  image_id = "ami-082b5a644766e0e6f"
  security_group_ids = ["awseb-e-sznmxim22e-stack-AWSEBSecurityGroup-10FZKNGB09G1W"]
  iam_instance_profile = "ecsInstanceRole"

  attributes {
    key = "key"
    value = "value"
  }
 %v
}
`

const testBaselineOceanECSLaunchSpecConfig_Update = `
resource "` + string(commons.OceanECSLaunchSpecResourceName) + `" "%v" {
  provider = "%v"  
  name = "%v"
  ocean_id = "%v"
  
  user_data = "hello world updated"
  image_id = "ami-082b5a644766e0e6f"
  iam_instance_profile = "ecsInstanceRole"

  attributes {
    key = "key"
    value = "value"
  }
  
  attributes {
  	key = "key2"
  	value = "value2"
  }

%v
}
`

// endregion
