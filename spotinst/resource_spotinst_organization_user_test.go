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

func createOrganizationUserResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OrgUserResourceName), name)
}

func testOrganizationUserDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OrgUserResourceName) {
			continue
		}
		input := &organization.ReadUserInput{UserID: spotinst.String(rs.Primary.ID)}
		resp, err := client.organization.ReadUser(context.Background(), input)
		if err == nil && resp != nil && resp.User != nil {
			return fmt.Errorf("organization user still exists")
		}
	}
	return nil
}

func testCheckOrganizationUserExists(user *organization.User, resourceName string) resource.TestCheckFunc {
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
		resp, err := client.organization.ReadUser(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.User.Email) != rs.Primary.Attributes["email"] {
			return fmt.Errorf("Organization User not found: %+v,\n %+v\n", resp.User, rs.Primary.Attributes)
		}
		*user = *resp.User
		return nil
	}
}

func createOrganizationUserTerraform(tfResource string, resourceName string) string {
	template := ""

	template += fmt.Sprintf(tfResource, resourceName)

	log.Printf("Terraform [%v] template:\n%v", resourceName, template)
	return template
}

// region Organization User
func TestAccSpotinstOrganization_User(t *testing.T) {
	userName := "terraform-user"
	userResourceName := createOrganizationUserResourceName(userName)

	var user organization.User
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOrganizationUserDestroy,

		Steps: []resource.TestStep{
			{
				Config:             createOrganizationUserTerraform(testOrganization_User_Create, userName),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOrganizationUserExists(&user, userResourceName),
					resource.TestCheckResourceAttr(userResourceName, "email", "terraformUser3@netapp.com"),
					resource.TestCheckResourceAttr(userResourceName, "first_name", "terraform"),
					resource.TestCheckResourceAttr(userResourceName, "last_name", "user"),
					resource.TestCheckResourceAttr(userResourceName, "role", "viewer"),
					resource.TestCheckResourceAttr(userResourceName, "policies.#", "1"),
					resource.TestCheckResourceAttr(userResourceName, "policies.0.policy_id", "pol-5479db5e"),
					resource.TestCheckResourceAttr(userResourceName, "policies.0.policy_account_ids.#", "1"),
					resource.TestCheckResourceAttr(userResourceName, "policies.0.policy_account_ids.0", "act-75eb3ba3"),
				),
			},
			{
				Config:             createOrganizationUserTerraform(testOrganization_User_Update_Policy, userName),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(userResourceName, "email", "terraformUpdatedUser3@netapp.com"),
					resource.TestCheckResourceAttr(userResourceName, "first_name", "terraform_updated_with_second_policy"),
					resource.TestCheckResourceAttr(userResourceName, "last_name", "user_updated_with_second_policy"),
					resource.TestCheckResourceAttr(userResourceName, "role", "viewer"),
					resource.TestCheckResourceAttr(userResourceName, "policies.#", "2"),
					resource.TestCheckResourceAttr(userResourceName, "policies.0.policy_id", "pol-5479db5e"),
					resource.TestCheckResourceAttr(userResourceName, "policies.0.policy_account_ids.#", "1"),
					resource.TestCheckResourceAttr(userResourceName, "policies.0.policy_account_ids.0", "act-75eb3ba3"),
					resource.TestCheckResourceAttr(userResourceName, "policies.1.policy_id", "pol-467f634c"),
					resource.TestCheckResourceAttr(userResourceName, "policies.1.policy_account_ids.#", "1"),
					resource.TestCheckResourceAttr(userResourceName, "policies.1.policy_account_ids.0", "act-e2be553a"),
				),
			},
			{
				Config:             createOrganizationUserTerraform(testOrganization_User_Update_Policy, userName),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(userResourceName, "email", "terraformUpdatedUser3@netapp.com"),
					resource.TestCheckResourceAttr(userResourceName, "first_name", "terraform_updated_with_second_policy"),
					resource.TestCheckResourceAttr(userResourceName, "last_name", "user_updated_with_second_policy"),
					resource.TestCheckResourceAttr(userResourceName, "role", "viewer"),
					resource.TestCheckResourceAttr(userResourceName, "policies.#", "2"),
					resource.TestCheckResourceAttr(userResourceName, "policies.0.policy_id", "pol-5479db5e"),
					resource.TestCheckResourceAttr(userResourceName, "policies.0.policy_account_ids.#", "1"),
					resource.TestCheckResourceAttr(userResourceName, "policies.0.policy_account_ids.0", "act-75eb3ba3"),
					resource.TestCheckResourceAttr(userResourceName, "policies.1.policy_id", "pol-467f634c"),
					resource.TestCheckResourceAttr(userResourceName, "policies.1.policy_account_ids.#", "1"),
					resource.TestCheckResourceAttr(userResourceName, "policies.1.policy_account_ids.0", "act-e2be553a"),
				),
			},
			{
				Config:             createOrganizationUserTerraform(testOrganization_User_Update_User_Group, userName),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(userResourceName, "user_group_ids.#", "1"),
					resource.TestCheckResourceAttr(userResourceName, "user_group_ids.0", "ugr-ef8935dd"),
				),
			},
		},
	})
}

const testOrganization_User_Create = `
resource "` + string(commons.OrgUserResourceName) + `" "%v" {
  provider = "aws"
  email = "terraformUser3@netapp.com"
  first_name = "terraform"
  last_name = "user"
  password = "terraformPwd@108"
  role = "viewer"
  
  policies{
    policy_id = "pol-5479db5e"
    policy_account_ids = ["act-75eb3ba3"]
  }
}
`

const testOrganization_User_Update_Policy = `
resource "` + string(commons.OrgUserResourceName) + `" "%v" {
  provider = "aws"
  email = "terraformUpdatedUser3@netapp.com"
  first_name = "terraform_updated_with_second_policy"
  last_name = "user_updated_with_second_policy"
  password = "terraformPwd@108"
  role = "viewer"
  
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

const testOrganization_User_Update_User_Group = `
resource "` + string(commons.OrgUserResourceName) + `" "%v" {
  provider = "aws"
  email = "terraformUpdatedUser3@netapp.com"
  first_name = "terraform_updated_with_second_policy"
  last_name = "user_updated_with_second_policy"
  password = "terraformPwd@108"
  role = "viewer"

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

// endregion
