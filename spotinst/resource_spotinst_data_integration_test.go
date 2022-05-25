package spotinst

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/dataintegration/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func createDataIntegrationName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.DataIntegrationResourceName), name)
}

func testDataIntegrationDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.DataIntegrationResourceName) {
			continue
		}
		input := &aws.ReadDataIntegrationInput{DataIntegrationId: spotinst.String(rs.Primary.ID)}
		resp, err := client.dataIntegration.CloudProviderAWS().ReadDataIntegration(context.Background(), input)
		if err == nil && resp != nil && resp.DataIntegration != nil {
			return fmt.Errorf("DataIntegration still exists")
		}
	}
	return nil
}

func testCheckDataIntegrationExists(di *aws.DataIntegration, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &aws.ReadDataIntegrationInput{DataIntegrationId: spotinst.String(rs.Primary.ID)}
		resp, err := client.dataIntegration.CloudProviderAWS().ReadDataIntegration(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.DataIntegration.ID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("DataIntegration not found: %+v,\n %+v\n", resp.DataIntegration, rs.Primary.Attributes)
		}
		*di = *resp.DataIntegration
		return nil
	}
}

type DataIntegrationMetadata struct {
	provider             string
	name                 string
	variables            string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createDataIntegrationTerraform(dim *DataIntegrationMetadata, formatToUse string) string {
	if dim == nil {
		return ""
	}

	if dim.provider == "" {
		dim.provider = "aws"
	}

	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`
	//format := formatToUse

	if dim.updateBaselineFields {
		format := testBaselineDataIntegrationConfig_Update
		template += fmt.Sprintf(format,
			dim.name,
			dim.provider,
			//dim.fieldsToAppend,
		)
	} else {
		format := testBaselineDataIntegrationConfig_Create
		template += fmt.Sprintf(format,
			dim.name,
			dim.provider,
			//dim.fieldsToAppend,
		)
	}

	if dim.variables != "" {
		template = dim.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", "data_integration_test", template)
	return template
}

// region DataIntegration: Baseline
func TestAccSpotinstDataIntegration_Baseline(t *testing.T) {
	name := "test-acc-data_integration_terraform_test"
	resourceName := createDataIntegrationName(name)

	var di aws.DataIntegration
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "aws") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testDataIntegrationDestroy,

		Steps: []resource.TestStep{
			{
				Config: createDataIntegrationTerraform(&DataIntegrationMetadata{
					name: name,
				}, testBaselineDataIntegrationConfig_Create),

				Check: resource.ComposeTestCheckFunc(
					testCheckDataIntegrationExists(&di, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "foo"),
					resource.TestCheckResourceAttr(resourceName, "status", "enabled"),
					resource.TestCheckResourceAttr(resourceName, "s3.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "s3.0.bucket_name", "terraform-test-do-not-delete"),
					resource.TestCheckResourceAttr(resourceName, "s3.0.subdir", "terraform-test-data-integration"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createDataIntegrationTerraform(&DataIntegrationMetadata{
					name:                 name,
					updateBaselineFields: true}, testBaselineDataIntegrationConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckDataIntegrationExists(&di, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "foo"),
					resource.TestCheckResourceAttr(resourceName, "status", "disabled"),
					resource.TestCheckResourceAttr(resourceName, "s3.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "s3.0.bucket_name", "terraform-test-do-not-delete-2"),
					resource.TestCheckResourceAttr(resourceName, "s3.0.subdir", "terraform-test-data-integration"),
				),
			},
		},
	})
}

const testBaselineDataIntegrationConfig_Create = `
resource "` + string(commons.DataIntegrationResourceName) + `" "%v" {
  provider = "%v"
  name  = "foo"
  status = "enabled"
  s3 {
  	bucket_name = "terraform-test-do-not-delete"
    subdir      = "terraform-test-data-integration"
  }
}
`

const testBaselineDataIntegrationConfig_Update = `
resource "` + string(commons.DataIntegrationResourceName) + `" "%v" {
  provider = "%v"
  name  = "foo"
  status = "disabled"
  s3 {
    bucket_name = "terraform-test-do-not-delete-2"
    subdir      = "terraform-test-data-integration"
  }
}
`

// endregion
