package spotinst

import (
	"context"
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/organization"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func createOrganizationPolicyResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OrgPolicyResourceName), name)
}

func testOrganizationPolicyDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OrgPolicyResourceName) {
			continue
		}
		input := &organization.ReadPolicyInput{PolicyID: spotinst.String(rs.Primary.ID)}
		resp, err := client.organization.ReadPolicy(context.Background(), input)
		if err == nil && resp != nil && resp.Policy != nil {
			return fmt.Errorf("policy still exists")
		}
	}
	return nil
}

func testCheckOrganizationPolicyExists(policy *organization.Policy, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &organization.ReadPolicyInput{PolicyID: spotinst.String(rs.Primary.ID)}
		resp, err := client.organization.ReadPolicy(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.Policy.Name) != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Policy not found: %+v,\n %+v\n", resp.Policy, rs.Primary.Attributes)
		}
		*policy = *resp.Policy
		return nil
	}
}

func createOrganizationPolicyTerraform(tfResource string, resourceName string) string {
	template := ""

	template += fmt.Sprintf(tfResource, resourceName)

	log.Printf("Terraform [%v] template:\n%v", resourceName, template)
	return template
}

// region Organziation: Policy
func TestAccSpotinstOrganization_Policy(t *testing.T) {
	policyName := "terraform-policy"
	policyResourceName := createOrganizationPolicyResourceName(policyName)

	var policy organization.Policy
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOrganizationPolicyDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOrganizationPolicyTerraform(testOrganization_Policy_Create, policyName),
				Check: resource.ComposeTestCheckFunc(
					testCheckOrganizationPolicyExists(&policy, policyResourceName),
					resource.TestCheckResourceAttr(policyResourceName, "name", "test-terraform-policy"),
					resource.TestCheckResourceAttr(policyResourceName, "description", "test-terraform-policy-create"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.#", "1"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.#", "1"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.0.actions.#", "1"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.0.actions.0", "ocean:createCluster"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.0.effect", "ALLOW"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.0.resources.#", "1"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.0.resources.0", "*"),
				),
			},
			{
				Config: createOrganizationPolicyTerraform(testOrganization_Policy_Update, policyName),
				Check: resource.ComposeTestCheckFunc(
					testCheckOrganizationPolicyExists(&policy, policyResourceName),
					resource.TestCheckResourceAttr(policyResourceName, "name", "test-terraform-policy-updated"),
					resource.TestCheckResourceAttr(policyResourceName, "description", "test-terraform-policy-update"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.#", "1"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.#", "2"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.0.actions.#", "1"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.0.actions.0", "ocean:createCluster"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.0.effect", "ALLOW"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.0.resources.#", "1"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.0.resources.0", "*"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.1.actions.#", "1"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.1.actions.0", "ocean:deleteCluster"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.1.effect", "DENY"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.1.resources.#", "1"),
					resource.TestCheckResourceAttr(policyResourceName, "policy_content.0.statements.1.resources.0", "*"),
				),
			},
		},
	})
}

const testOrganization_Policy_Create = `
resource "` + string(commons.OrgPolicyResourceName) + `" "%v" {
  provider = "aws"
  name = "test-terraform-policy"
  description = "test-terraform-policy-create"
  policy_content{
    statements{
      actions = ["ocean:createCluster"]
      effect = "ALLOW"
      resources = ["*"]
    }
  }
}
`

const testOrganization_Policy_Update = `
resource "` + string(commons.OrgPolicyResourceName) + `" "%v" {
  provider = "aws"
  name = "test-terraform-policy-updated"
  description = "test-terraform-policy-update"
  policy_content{
    statements{
      actions = ["ocean:createCluster"]
      effect = "ALLOW"
      resources = ["*"]
    }
    statements{
      actions = ["ocean:deleteCluster"]
      effect = "DENY"
      resources = ["*"]
    }
  }
}
`

// endregion
