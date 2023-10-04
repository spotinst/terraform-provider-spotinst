package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/spotinst/spotinst-sdk-go/service/organization"
	organizationPackage "github.com/spotinst/terraform-provider-spotinst/spotinst/organization_user"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceOrgUser() *schema.Resource {
	setupOrgUser()
	return &schema.Resource{
		CreateContext: resourceOrgUserCreate,
		UpdateContext: resourceOrgUserUpdate,
		ReadContext:   resourceOrgUserRead,
		DeleteContext: resourceOrgUserDelete,

		Schema: commons.OrgUserResource.GetSchemaMap(),
	}
}

func setupOrgUser() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	organizationPackage.Setup(fieldsMap)

	commons.OrgUserResource = commons.NewOrgUserResource(fieldsMap)
}

func resourceOrgUserDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OrgUserResource.GetName(), id)

	input := &organization.DeleteUserInput{UserID: spotinst.String(id)}
	if _, err := meta.(*Client).organization.DeleteUser(context.Background(), input); err != nil {
		return diag.Errorf("[ERROR] Failed to delete user: %s", err)
	}

	resourceData.SetId("")
	return nil
}

func resourceOrgUserRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.OrgUserResource.GetName(), id)

	client := meta.(*Client)
	input := &organization.ReadUserInput{UserID: spotinst.String(resourceData.Id())}
	userResponse, err := client.organization.ReadUser(context.Background(), input)
	if err != nil {
		return diag.Errorf("[ERROR] Failed to read user: %s", err)
	}

	// If nothing was found, then return no state.
	user := userResponse.User
	if user == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OrgUserResource.OnRead(user, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> User read successfully: %s <===", id)
	return nil
}

func resourceOrgUserCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OrgUserResource.GetName())

	user, err := commons.OrgUserResource.OnCreate(resourceData, meta)
	var policies = user.Policies
	var userGroupIds = user.UserGroupIds

	user.Policies = nil
	user.UserGroupIds = nil

	if err != nil {
		return diag.FromErr(err)
	}

	userId, err := createUser(user, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	var updateErr error = nil

	if policies != nil {
		updateErr = updatePolicyMapping(policies, userId, meta.(*Client))
	}

	if updateErr != nil {
		return diag.FromErr(updateErr)
	}

	if userGroupIds != nil {
		updateErr = updateUserGroupMapping(userGroupIds, userId, meta.(*Client))
	}

	if updateErr != nil {
		return diag.FromErr(updateErr)
	}

	resourceData.SetId(spotinst.StringValue(userId))
	log.Printf("===> User created successfully: %s <===", resourceData.Id())

	return resourceOrgUserRead(ctx, resourceData, meta)
}

func createUser(userObj *organization.User, spotinstClient *Client) (*string, error) {
	input := userObj
	resp, err := spotinstClient.organization.CreateUser(context.Background(), input, spotinst.Bool(true))
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create user: %s", err)
	}
	return resp.User.UserID, nil
}

func updatePolicyMapping(policies []*organization.UserPolicy, userId *string, spotinstClient *Client) error {
	err := spotinstClient.organization.UpdatePolicyMappingOfUser(context.Background(), &organization.UpdatePolicyMappingOfUserInput{
		UserID:   userId,
		Policies: policies,
	})
	if err != nil {
		return fmt.Errorf("[ERROR] failed to update policy mapping for user: %s", err)
	}
	return nil
}

func updateUserGroupMapping(userGroupIds []string, userId *string, spotinstClient *Client) error {
	err := spotinstClient.organization.UpdateUserGroupMappingOfUser(context.Background(), &organization.UpdateUserGroupMappingOfUserInput{
		UserID:       userId,
		UserGroupIds: userGroupIds,
	})
	if err != nil {
		return fmt.Errorf("[ERROR] failed to update policy mapping for user: %s", err)
	}
	return nil
}

func resourceOrgUserUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.OrgUserResource.GetName(), id)

	shouldUpdate, user, err := commons.OrgUserResource.OnUpdate(resourceData, meta)

	if err != nil {
		return diag.FromErr(err)
	}

	var policies []*organization.UserPolicy = user.Policies
	var userGroupIds []string = user.UserGroupIds

	if shouldUpdate {
		user.UserID = spotinst.String(id)
		var accountIds []string
		if policies != nil {
			for _, policy := range policies {
				for _, actId := range policy.AccountIds {
					if spotinst.StringValue(policy.PolicyId) != "3" {
						for i := 0; i < len(policy.AccountIds); i++ {
							if accountIds != nil {
								if accountIds[i] == actId {
									break
								}
							}
							accountIds = append(accountIds, actId)
						}
					}
				}
			}
			var accountViewerPolicy organization.UserPolicy
			accountViewerPolicy.PolicyId = spotinst.String("3")
			accountViewerPolicy.AccountIds = accountIds
			policies = append(policies, &accountViewerPolicy)
			if err := updatePolicyMapping(policies, &id, meta.(*Client)); err != nil {
				return diag.FromErr(err)
			}
		}

		if policies == nil {
			var accountViewerPolicy organization.UserPolicy
			accountViewerPolicy.PolicyId = spotinst.String("3")
			accountViewerPolicy.AccountIds = append(accountIds, os.Getenv("SPOTINST_ACCOUNT_AWS"))
			policies = append(policies, &accountViewerPolicy)
			if err := updatePolicyMapping(policies, &id, meta.(*Client)); err != nil {
				return diag.FromErr(err)
			}
		}

		if userGroupIds != nil {
			if err := updateUserGroupMapping(userGroupIds, &id, meta.(*Client)); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	log.Printf("===> User mapping updated successfully: %s <===", id)
	return resourceOrgUserRead(ctx, resourceData, meta)
}
