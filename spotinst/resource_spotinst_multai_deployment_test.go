package spotinst

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

func TestAccSpotinstMultaiDeployment_Basic(t *testing.T) {
	var deployment multai.Deployment
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSpotinstMultaiDeploymentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSpotinstMultaiDeploymentConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiDeploymentExists("spotinst_multai_deployment.foo", &deployment),
					testAccCheckSpotinstMultaiDeploymentAttributes(&deployment),
					resource.TestCheckResourceAttr("spotinst_multai_deployment.foo", "name", "foo"),
				),
			},
		},
	})
}

func TestAccSpotinstMultaiDeployment_Updated(t *testing.T) {
	var deployment multai.Deployment
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSpotinstMultaiDeploymentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSpotinstMultaiDeploymentConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiDeploymentExists("spotinst_multai_deployment.foo", &deployment),
					testAccCheckSpotinstMultaiDeploymentAttributes(&deployment),
					resource.TestCheckResourceAttr("spotinst_multai_deployment.foo", "name", "foo"),
				),
			},
			{
				Config: testAccCheckSpotinstMultaiDeploymentConfigNewValue,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiDeploymentExists("spotinst_multai_deployment.foo", &deployment),
					testAccCheckSpotinstMultaiDeploymentAttributesUpdated(&deployment),
					resource.TestCheckResourceAttr("spotinst_multai_deployment.foo", "name", "bar"),
				),
			},
		},
	})
}

func testAccCheckSpotinstMultaiDeploymentDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Client)
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

func testAccCheckSpotinstMultaiDeploymentAttributes(deployment *multai.Deployment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if p := spotinst.StringValue(deployment.Name); p != "foo" {
			return fmt.Errorf("bad content: %s", p)
		}
		return nil
	}
}

func testAccCheckSpotinstMultaiDeploymentAttributesUpdated(deployment *multai.Deployment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if p := spotinst.StringValue(deployment.Name); p != "bar" {
			return fmt.Errorf("bad content: %s", p)
		}
		return nil
	}
}

func testAccCheckSpotinstMultaiDeploymentExists(n string, deployment *multai.Deployment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProvider.Meta().(*Client)
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

const testAccCheckSpotinstMultaiDeploymentConfigBasic = `
resource "spotinst_multai_deployment" "foo" {
  name = "foo"

  tags {
    env = "prod"
    app = "web"
  }
}`

const testAccCheckSpotinstMultaiDeploymentConfigNewValue = `
resource "spotinst_multai_deployment" "foo" {
  name = "bar"

  tags {
    env = "prod"
    app = "web"
  }
}`
