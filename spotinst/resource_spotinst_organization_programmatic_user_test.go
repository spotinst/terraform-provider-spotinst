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

func createOrganizationProgrammaticUserResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OrgProgrammaticUserResourceName), name)
}

func testOrganizationProgrammaticUserDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OrgProgrammaticUserResourceName) {
			continue
		}
		input := &organization.ReadUserInput{UserID: spotinst.String(rs.Primary.ID)}
		resp, err := client.organization.ReadProgUser(context.Background(), input)
		if err == nil && resp != nil && resp.ProgUser != nil {
			return fmt.Errorf("Organization programmatic user still exists")
		}
	}
	return nil
}

func testCheckOrganizationProgrammaticUserExists(user *organization.ProgrammaticUser, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &organization.ReadUserInput{UserID: spotinst.String(rs.Primary.ID)}
		resp, err := client.organization.ReadProgUser(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.ProgUser.Name) != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Organization programmatic user not found: %+v,\n %+v\n", resp.ProgUser, rs.Primary.Attributes)
		}
		*user = *resp.ProgUser
		return nil
	}
}

func createOrganizationProgrammaticUserTerraform(tfResource string, resourceName string) string {
	template := ""

	template += fmt.Sprintf(tfResource, resourceName)

	log.Printf("Terraform [%v] template:\n%v", resourceName, template)
	return template
}

// region Organization User
func TestAccSpotinstOrganization_ProgrammaticUser_WithAccounts(t *testing.T) {
	progUserName := "terraform-programmatic-user"
	progUserResourceName := createOrganizationProgrammaticUserResourceName(progUserName)

	var progUser organization.ProgrammaticUser
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOrganizationProgrammaticUserDestroy,

		Steps: []resource.TestStep{
			{
				Config:             createOrganizationProgrammaticUserTerraform(testOrganization_Programmatic_User_Create_With_Accounts, progUserName),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOrganizationProgrammaticUserExists(&progUser, progUserResourceName),
					//resource.TestCheckResourceAttr(progUserResourceName, "name", "terraform_programmatic_user"),
					resource.TestCheckResourceAttr(progUserResourceName, "description", "for terraform unit testing"),
					resource.TestCheckResourceAttr(progUserResourceName, "accounts.#", "2"),
					resource.TestCheckResourceAttr(progUserResourceName, "accounts.0.account_id", "act-75eb3ba3"),
					resource.TestCheckResourceAttr(progUserResourceName, "accounts.0.account_role", "viewer"),
					resource.TestCheckResourceAttr(progUserResourceName, "accounts.1.account_id", "act-e2be553a"),
					resource.TestCheckResourceAttr(progUserResourceName, "accounts.1.account_role", "viewer"),
				),
			},
		},
	})
}

const testOrganization_Programmatic_User_Create_With_Accounts = `
resource "` + string(commons.OrgProgrammaticUserResourceName) + `" "%v" {
  provider = "aws"
  name = "terraform_programmatic_user"
  description = "for terraform unit testing"

  accounts {
    account_id = "act-75eb3ba3"
    account_role = "viewer"
  }

  accounts {
    account_id = "act-e2be553a"
    account_role = "viewer"
  }
}
`

func TestAccSpotinstOrganization_ProgrammaticUser_WithPolicies(t *testing.T) {
	progUserName := "terraform-programmatic-user"
	progUserResourceName := createOrganizationProgrammaticUserResourceName(progUserName)

	var progUser organization.ProgrammaticUser
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOrganizationProgrammaticUserDestroy,

		Steps: []resource.TestStep{
			{
				Config:             createOrganizationProgrammaticUserTerraform(testOrganization_Programmatic_User_Create_With_Policies, progUserName),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOrganizationProgrammaticUserExists(&progUser, progUserResourceName),
					//resource.TestCheckResourceAttr(progUserResourceName, "name", "terraform_programmatic_user_with_policies"),
					resource.TestCheckResourceAttr(progUserResourceName, "description", "for terraform unit testing with policies"),
					resource.TestCheckResourceAttr(progUserResourceName, "policies.#", "1"),
					resource.TestCheckResourceAttr(progUserResourceName, "policies.0.policy_id", "pol-5479db5e"),
					resource.TestCheckResourceAttr(progUserResourceName, "policies.0.policy_account_ids.#", "1"),
					resource.TestCheckResourceAttr(progUserResourceName, "policies.0.policy_account_ids.0", "act-75eb3ba3"),
				),
			},
			{
				Config:             createOrganizationProgrammaticUserTerraform(testOrganization_Programmatic_User_Update_Policies, progUserName),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOrganizationProgrammaticUserExists(&progUser, progUserResourceName),
					resource.TestCheckResourceAttr(progUserResourceName, "policies.#", "2"),
					resource.TestCheckResourceAttr(progUserResourceName, "policies.0.policy_id", "pol-5479db5e"),
					resource.TestCheckResourceAttr(progUserResourceName, "policies.0.policy_account_ids.#", "1"),
					resource.TestCheckResourceAttr(progUserResourceName, "policies.0.policy_account_ids.0", "act-75eb3ba3"),
					resource.TestCheckResourceAttr(progUserResourceName, "policies.1.policy_id", "pol-467f634c"),
					resource.TestCheckResourceAttr(progUserResourceName, "policies.1.policy_account_ids.#", "1"),
					resource.TestCheckResourceAttr(progUserResourceName, "policies.1.policy_account_ids.0", "act-e2be553a"),
				),
			},
			{
				Config:             createOrganizationProgrammaticUserTerraform(testOrganization_Programmatic_User_Update_UserGroups, progUserName),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOrganizationProgrammaticUserExists(&progUser, progUserResourceName),
					resource.TestCheckResourceAttr(progUserResourceName, "user_group_ids.#", "1"),
					resource.TestCheckResourceAttr(progUserResourceName, "user_group_ids.0", "ugr-ef8935dd"),
				),
			},
		},
	})
}

const testOrganization_Programmatic_User_Create_With_Policies = `
resource "` + string(commons.OrgProgrammaticUserResourceName) + `" "%v" {
  provider = "aws"
  name = "terraform_programmatic_user_with_policies"
  description = "for terraform unit testing with policies"

  policies{
    policy_id = "pol-5479db5e"
    policy_account_ids = ["act-75eb3ba3"]
  }

}
`

const testOrganization_Programmatic_User_Update_Policies = `
resource "` + string(commons.OrgProgrammaticUserResourceName) + `" "%v" {
  provider = "aws"
  name = "terraform_programmatic_user_with_policies"
  description = "for terraform unit testing with policies and user groups"

  policies{
    policy_id = "pol-5479db5e"
    policy_account_ids = ["act-75eb3ba3"]
  }

  policies{
    policy_id = "pol-467f634c"
    policy_account_ids = ["act-e2be553a"]
  }

}
`

const testOrganization_Programmatic_User_Update_UserGroups = `
resource "` + string(commons.OrgProgrammaticUserResourceName) + `" "%v" {
  provider = "aws"
  name = "terraform_programmatic_user_with_policies"
  description = "for terraform unit testing with policies and user groups"

  user_group_ids = ["ugr-ef8935dd"]

  policies{
    policy_id = "pol-5479db5e"
    policy_account_ids = ["act-75eb3ba3"]
  }

  policies{
    policy_id = "pol-467f634c"
    policy_account_ids = ["act-e2be553a"]
  }

}
`
