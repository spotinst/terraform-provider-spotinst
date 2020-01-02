package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

func init() {
	resource.AddTestSweepers("spotinst_multai_deployment", &resource.Sweeper{
		Name: "spotinst_multai_deployment",
		F:    testSweepMultaiDeployment,
	})
}

func testSweepMultaiDeployment(region string) error {
	client, err := getProviderClient("aws")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).multai
	input := &multai.ListDeploymentsInput{}
	if resp, err := conn.ListDeployments(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of deployments to sweep")
	} else {
		if len(resp.Deployments) == 0 {
			log.Printf("[INFO] No deployments to sweep")
		}
		for _, depl := range resp.Deployments {
			if strings.Contains(spotinst.StringValue(depl.Name), "test-acc-") {
				if _, err := conn.DeleteDeployment(context.Background(), &multai.DeleteDeploymentInput{DeploymentID: depl.ID}); err != nil {
					return fmt.Errorf("unable to delete deployment %v in sweep", spotinst.StringValue(depl.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(depl.ID))
				}
			}
		}
	}
	return nil
}

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
	deployName := "test-acc-mlb-baseline"
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
					resource.TestCheckResourceAttr(resourceName, "name", "test-acc-mlb-baseline"),
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
					resource.TestCheckResourceAttr(resourceName, "name", "test-acc-mlb-baseline"),
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
