package spotinst

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/extendedresourcedefinition"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func createExtendedResourceDefinitionResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.ExtendedResourceDefinitionResourceName), name)
}

func testExtendedResourceDefinitionDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.ExtendedResourceDefinitionResourceName) {
			continue
		}
		input := &extendedresourcedefinition.ReadExtendedResourceDefinitionInput{ExtendedResourceDefinitionID: spotinst.String(rs.Primary.ID)}
		resp, err := client.extendedResourceDefinition.Read(context.Background(), input)
		if err == nil && resp != nil && resp.ExtendedResourceDefinition != nil {
			return fmt.Errorf("extendedResourceDefinition still exists")
		}
	}
	return nil
}

func testCheckExtendedResourceDefinitionExists(erd *extendedresourcedefinition.ExtendedResourceDefinition, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &extendedresourcedefinition.ReadExtendedResourceDefinitionInput{ExtendedResourceDefinitionID: spotinst.String(rs.Primary.ID)}
		resp, err := client.extendedResourceDefinition.Read(context.Background(), input)
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

func createExtendedResourceDefinitionTerraform(ccm *ExtendedResourceDefinitionMetadata, formatToUse string) string {
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

	log.Printf("Terraform [%v] template:\n%v", "extended_resource_definition_test", template)
	return template
}

// region ExtendedResourceDefinition: Baseline
func TestAccSpotinstExtendedResourceDefinition_Baseline(t *testing.T) {
	name := "test-acc-extended_resource_definition_terraform_test"
	resourceName := createExtendedResourceDefinitionResourceName(name)

	var erd extendedresourcedefinition.ExtendedResourceDefinition
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testExtendedResourceDefinitionDestroy,

		Steps: []resource.TestStep{
			{
				Config: createExtendedResourceDefinitionTerraform(&ExtendedResourceDefinitionMetadata{
					name: name,
				}, testBaselineExtendedResourceDefinitionConfig_Create),

				Check: resource.ComposeTestCheckFunc(
					testCheckExtendedResourceDefinitionExists(&erd, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "example.com/terraform-test-baseline"),
					resource.TestCheckResourceAttr(resourceName, "resource_mapping.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "resource_mapping.c3.large", "2Ki"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createExtendedResourceDefinitionTerraform(&ExtendedResourceDefinitionMetadata{
					name:                 name,
					updateBaselineFields: true}, testBaselineExtendedResourceDefinitionConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckExtendedResourceDefinitionExists(&erd, resourceName),
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
resource "` + string(commons.ExtendedResourceDefinitionResourceName) + `" "%v" {
  provider = "%v"
  name  = "example.com/terraform-test-baseline"
  resource_mapping = {
    "c3.large"  = "2Ki"
  }
}
`

const testBaselineExtendedResourceDefinitionConfig_Update = `
resource "` + string(commons.ExtendedResourceDefinitionResourceName) + `" "%v" {
  provider = "%v"
  name  = "example.com/terraform-test-baseline"
  resource_mapping = {
    "c3.large"  = "2Ki"
    "c3.xlarge" = "4Ki"
  }
}
`

// endregion
