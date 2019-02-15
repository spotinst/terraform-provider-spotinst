package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"log"
	"testing"
)

func createMultaiDeploymentResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.MultaiDeploymentResourceName), name)
}

func testAccCheckSpotinstMultaiDeploymentDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "spotinst_multai_deployment" {
			continue
		}
		input := &multai.ReadDeploymentInput{
			DeploymentID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.multai.ReadDeployment(context.Background(), input)
		if err == nil && resp != nil && resp.Deployment != nil {
			return fmt.Errorf("deployment still exists")
		}
	}
	return nil
}

func testAccCheckSpotinstMultaiDeploymentAttributes(deployment *multai.Deployment, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if p := spotinst.StringValue(deployment.Name); p != expectedName {
			return fmt.Errorf("bad content: %s", p)
		}
		return nil
	}
}

func testAccCheckSpotinstMultaiDeploymentExists(deployment *multai.Deployment, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &multai.ReadDeploymentInput{
			DeploymentID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.multai.ReadDeployment(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.Deployment.ID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("deployment not found: %+v,\n %+v\n", resp.Deployment, rs.Primary.Attributes)
		}
		*deployment = *resp.Deployment
		return nil
	}
}

type DeploymentConfigMetadata struct {
	provider             string
	name                 string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createDeploymentTerraform(bcm *DeploymentConfigMetadata) string {
	if bcm == nil {
		return ""
	}

	if bcm.provider == "" {
		bcm.provider = "aws"
	}

	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`

	if bcm.updateBaselineFields {
		format := testBaselineDeploymentConfig_Update
		template += fmt.Sprintf(format,
			bcm.name,
			bcm.provider,
			bcm.name,
		)
	} else {
		format := testBaselineDeploymentConfig_Create

		template += fmt.Sprintf(format,
			bcm.name,
			bcm.provider,
			bcm.name,
		)
	}

	log.Printf("Terraform [%v] template:\n%v", bcm.name, template)
	return template
}

func TestAccSpotinstMultaiDeployment_Baseline(t *testing.T) {
	deployName := "mlb-baseline"
	resourceName := createMultaiDeploymentResourceName(deployName)

	var deployment multai.Deployment
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckSpotinstMultaiDeploymentDestroy,

		Steps: []resource.TestStep{
			{
				Config: createDeploymentTerraform(&DeploymentConfigMetadata{
					name: deployName,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiDeploymentExists(&deployment, resourceName),
					testAccCheckSpotinstMultaiDeploymentAttributes(&deployment, deployName),
					resource.TestCheckResourceAttr(resourceName, "name", "mlb-baseline"),
				),
			},
			{
				Config: createDeploymentTerraform(&DeploymentConfigMetadata{
					name:                 deployName,
					updateBaselineFields: true,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiDeploymentExists(&deployment, resourceName),
					testAccCheckSpotinstMultaiDeploymentAttributes(&deployment, deployName),
					resource.TestCheckResourceAttr(resourceName, "name", "mlb-baseline"),
				),
			},
		},
	})
}

const testBaselineDeploymentConfig_Create = `
resource "` + string(commons.MultaiDeploymentResourceName) + `" "%v" {
  provider = "%v"
  name = "%v"
}`

const testBaselineDeploymentConfig_Update = `
resource "` + string(commons.MultaiDeploymentResourceName) + `" "%v" {
  provider = "%v"
  name = "%v"
}`
