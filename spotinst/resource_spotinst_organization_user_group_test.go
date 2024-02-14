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

func createOrganizationUserGroupResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OrgUserGroupResourceName), name)
}

func testOrganizationUserGroupDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OrgUserGroupResourceName) {
			continue
		}
		input := &organization.ReadUserGroupInput{UserGroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.organization.ReadUserGroup(context.Background(), input)
		if err == nil && resp != nil && resp.UserGroup != nil {
			return fmt.Errorf("policy still exists")
		}
	}
	return nil
}

func testCheckOrganizationUserGroupExists(userGroup *organization.UserGroup, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &organization.ReadUserGroupInput{UserGroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.organization.ReadUserGroup(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.UserGroup.Name) != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Policy not found: %+v,\n %+v\n", resp.UserGroup, rs.Primary.Attributes)
		}
		*userGroup = *resp.UserGroup
		return nil
	}
}

func createOrganizationUserGroupTerraform(tfResource string, resourceName string) string {
	template := ""

	template += fmt.Sprintf(tfResource, resourceName)

	log.Printf("Terraform [%v] template:\n%v", resourceName, template)
	return template
}

// region Organziation: User Group
func TestAccSpotinstOrganization_UserGroup(t *testing.T) {
	userGroupName := "terraform-user-group"
	userGroupResourceName := createOrganizationUserGroupResourceName(userGroupName)

	var userGroup organization.UserGroup
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOrganizationUserGroupDestroy,

		Steps: []resource.TestStep{
			{
				Config:             createOrganizationUserGroupTerraform(testOrganization_User_Group_Create, userGroupName),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckOrganizationUserGroupExists(&userGroup, userGroupResourceName),
					resource.TestCheckResourceAttr(userGroupResourceName, "name", "terraform_user_group"),
					resource.TestCheckResourceAttr(userGroupResourceName, "description", "user group by terraform"),
					resource.TestCheckResourceAttr(userGroupResourceName, "policies.#", "1"),
					resource.TestCheckResourceAttr(userGroupResourceName, "policies.0.policy_id", "pol-5479db5e"),
					resource.TestCheckResourceAttr(userGroupResourceName, "policies.0.account_ids.#", "1"),
					resource.TestCheckResourceAttr(userGroupResourceName, "policies.0.account_ids.0", "act-75eb3ba3"),
				),
			},
			{
				Config:             createOrganizationUserGroupTerraform(testOrganization_User_Group_Update, userGroupName),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(userGroupResourceName, "name", "terraform_user_group_updated"),
					resource.TestCheckResourceAttr(userGroupResourceName, "description", "user group by terraform updated"),
				),
			},
			{
				Config:             createOrganizationUserGroupTerraform(testOrganization_User_Group_Update_Policy, userGroupName),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(userGroupResourceName, "policies.#", "2"),
					resource.TestCheckResourceAttr(userGroupResourceName, "policies.0.policy_id", "pol-5479db5e"),
					resource.TestCheckResourceAttr(userGroupResourceName, "policies.0.account_ids.#", "1"),
					resource.TestCheckResourceAttr(userGroupResourceName, "policies.0.account_ids.0", "act-75eb3ba3"),
					resource.TestCheckResourceAttr(userGroupResourceName, "policies.1.policy_id", "pol-467f634c"),
					resource.TestCheckResourceAttr(userGroupResourceName, "policies.1.account_ids.#", "1"),
					resource.TestCheckResourceAttr(userGroupResourceName, "policies.1.account_ids.0", "act-e2be553a"),
				),
			},
			{
				Config:             createOrganizationUserGroupTerraform(testOrganization_User_Group_Update_User, userGroupName),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(userGroupResourceName, "user_ids.#", "1"),
					resource.TestCheckResourceAttr(userGroupResourceName, "user_ids.0", "u-429562d5"),
				),
			},
		},
	})
}

const testOrganization_User_Group_Create = `
resource "` + string(commons.OrgUserGroupResourceName) + `" "%v" {
  provider = "aws"
  name = "terraform_user_group"
  description = "user group by terraform"
  
  policies {
    account_ids = ["act-75eb3ba3"]
    policy_id = "pol-5479db5e"
  }
}
`

const testOrganization_User_Group_Update = `
resource "` + string(commons.OrgUserGroupResourceName) + `" "%v" {
  provider = "aws"
  name = "terraform_user_group_updated"
  description = "user group by terraform updated"
  
  policies {
    account_ids = ["act-75eb3ba3"]
    policy_id = "pol-5479db5e"
  }
}
`

const testOrganization_User_Group_Update_Policy = `
resource "` + string(commons.OrgUserGroupResourceName) + `" "%v" {
  provider = "aws"
  name = "terraform_user_group_updated"
  description = "user group by terraform updated"
  
  policies {
    account_ids = ["act-75eb3ba3"]
    policy_id = "pol-5479db5e"
  }

  policies {
    account_ids = ["act-e2be553a"]
    policy_id = "pol-467f634c"
  }
}
`

const testOrganization_User_Group_Update_User = `
resource "` + string(commons.OrgUserGroupResourceName) + `" "%v" {
  provider = "aws"
  name = "terraform_user_group_updated"
  description = "user group by terraform updated"
  
  user_ids = ["u-429562d5"]  

  policies {
    account_ids = ["act-75eb3ba3"]
    policy_id = "pol-5479db5e"
  }

  policies {
    account_ids = ["act-e2be553a"]
    policy_id = "pol-467f634c"
  }
}
`

// endregion
