package spotinst

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func createOceanAWSExtendedResourceDefinitionResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanAWSExtendedResourceDefinitionResourceName), name)
}

func testOceanAWSExtendedResourceDefinitionDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OceanAWSExtendedResourceDefinitionResourceName) {
			continue
		}
		input := &aws.ReadExtendedResourceDefinitionInput{ExtendedResourceDefinitionID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAWS().ReadExtendedResourceDefinition(context.Background(), input)
		if err == nil && resp != nil && resp.ExtendedResourceDefinition != nil {
			return fmt.Errorf("extendedResourceDefinition still exists")
		}
	}
	return nil
}

func testCheckOceanAWSExtendedResourceDefinitionExists(erd *aws.ExtendedResourceDefinition, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &aws.ReadExtendedResourceDefinitionInput{ExtendedResourceDefinitionID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAWS().ReadExtendedResourceDefinition(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.ExtendedResourceDefinition.ID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("extendedResourceDefinition not found: %+v,\n %+v\n", resp.ExtendedResourceDefinition, rs.Primary.Attributes)
		}
		*erd = *resp.ExtendedResourceDefinition
		return nil
	}
}

type ExtendedResourceDefinitionMetadata struct {
	provider             string
	name                 string
	variables            string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createOceanAWSExtendedResourceDefinitionTerraform(ccm *ExtendedResourceDefinitionMetadata, formatToUse string) string {
	if ccm == nil {
		return ""
	}

	if ccm.provider == "" {
		ccm.provider = "aws"
	}

	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`
	//format := formatToUse

	if ccm.updateBaselineFields {
		format := testBaselineExtendedResourceDefinitionConfig_Update
		template += fmt.Sprintf(format,
			ccm.name,
			ccm.provider,
			//ccm.fieldsToAppend,
		)
	} else {
		format := testBaselineExtendedResourceDefinitionConfig_Create
		template += fmt.Sprintf(format,
			ccm.name,
			ccm.provider,
			//ccm.fieldsToAppend,
		)
	}

	if ccm.variables != "" {
		template = ccm.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", "ocean_aws_extended_resource_definition_test", template)
	return template
}

// region ExtendedResourceDefinition: Baseline
func TestAccSpotinstOceanAWSExtendedResourceDefinition_Baseline(t *testing.T) {
	name := "test-acc-ocean_aws_extended_resource_definition_terraform_test"
	resourceName := createOceanAWSExtendedResourceDefinitionResourceName(name)

	var erd aws.ExtendedResourceDefinition
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "aws") },
		ProviderFactories: TestAccProvidersV2,
		CheckDestroy:      testOceanAWSExtendedResourceDefinitionDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAWSExtendedResourceDefinitionTerraform(&ExtendedResourceDefinitionMetadata{
					name: name,
				}, testBaselineExtendedResourceDefinitionConfig_Create),

				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExtendedResourceDefinitionExists(&erd, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "example.com/terraform-test-baseline"),
					resource.TestCheckResourceAttr(resourceName, "resource_mapping.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "resource_mapping.c3.large", "2Ki"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanAWSExtendedResourceDefinitionTerraform(&ExtendedResourceDefinitionMetadata{
					name:                 name,
					updateBaselineFields: true}, testBaselineExtendedResourceDefinitionConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExtendedResourceDefinitionExists(&erd, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "example.com/terraform-test-baseline"),
					resource.TestCheckResourceAttr(resourceName, "resource_mapping.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "resource_mapping.c3.large", "2Ki"),
					resource.TestCheckResourceAttr(resourceName, "resource_mapping.c3.xlarge", "4Ki"),
				),
			},
		},
	})
}

const testBaselineExtendedResourceDefinitionConfig_Create = `
resource "` + string(commons.OceanAWSExtendedResourceDefinitionResourceName) + `" "%v" {
  provider = "%v"
  name  = "example.com/terraform-test-baseline"
  resource_mapping = {
    "c3.large"  = "2Ki"
  }
}
`

const testBaselineExtendedResourceDefinitionConfig_Update = `
resource "` + string(commons.OceanAWSExtendedResourceDefinitionResourceName) + `" "%v" {
  provider = "%v"
  name  = "example.com/terraform-test-baseline"
  resource_mapping = {
    "c3.large"  = "2Ki"
    "c3.xlarge" = "4Ki"
  }
}
`

// endregion
