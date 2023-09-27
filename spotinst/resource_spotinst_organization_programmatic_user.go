package spotinst

import (
	"context"
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/organization"
	organizationPackage "github.com/spotinst/terraform-provider-spotinst/spotinst/organization_programmatic_user"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceOrgProgUser() *schema.Resource {
	setupOrgProgUser()
	return &schema.Resource{
		CreateContext: resourceOrgProgUserCreate,
		UpdateContext: resourceOrgProgUserUpdate,
		ReadContext:   resourceOrgProgUserRead,
		DeleteContext: resourceOrgProgUserDelete,

		Schema: commons.OrgProgUserResource.GetSchemaMap(),
	}
}

func setupOrgProgUser() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	organizationPackage.Setup(fieldsMap)

	commons.OrgProgUserResource = commons.NewOrgProgUserResource(fieldsMap)
}

func resourceOrgProgUserDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OrgProgUserResource.GetName(), id)

	input := &organization.DeleteUserInput{UserID: spotinst.String(id)}
	if _, err := meta.(*Client).organization.DeleteUser(context.Background(), input); err != nil {
		return diag.Errorf("[ERROR] Failed to delete user: %s", err)
	}

	resourceData.SetId("")
	return nil
}

func resourceOrgProgUserRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.OrgProgUserResource.GetName(), id)

	client := meta.(*Client)
	input := &organization.ReadUserInput{UserID: spotinst.String(resourceData.Id())}
	userResponse, err := client.organization.ReadProgUser(context.Background(), input)
	if err != nil {
		return diag.Errorf("[ERROR] Failed to read user: %s", err)
	}

	// If nothing was found, then return no state.
	progUser := userResponse.ProgUser
	if progUser == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OrgProgUserResource.OnRead(progUser, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> User read successfully: %s <===", id)
	return nil
}

func resourceOrgProgUserCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OrgProgUserResource.GetName())

	progUser, err := commons.OrgProgUserResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	var userGroupIds = progUser.UserGroupIds

	progUser.UserGroupIds = nil

	if err != nil {
		return diag.FromErr(err)
	}

	userId, err := createProgUser(progUser, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	var updateErr error = nil

	if userGroupIds != nil {
		updateErr = updateUserGroupMapping(userGroupIds, userId, meta.(*Client))
	}

	if updateErr != nil {
		return diag.FromErr(updateErr)
	}

	resourceData.SetId(spotinst.StringValue(userId))
	log.Printf("===> User created successfully: %s <===", resourceData.Id())

	return resourceOrgProgUserRead(ctx, resourceData, meta)
}

func createProgUser(userObj *organization.ProgrammaticUser, spotinstClient *Client) (*string, error) {
	input := userObj
	resp, err := spotinstClient.organization.CreateProgUser(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create user: %s", err)
	}
	return resp.ProgrammaticUser.ProgUserId, nil
}

func resourceOrgProgUserUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.OrgProgUserResource.GetName(), id)

	shouldUpdate, user, err := commons.OrgProgUserResource.OnUpdate(resourceData, meta)

	if err != nil {
		return diag.FromErr(err)
	}

	var policies []*organization.ProgPolicy = user.Policies
	var userGroupIds []string = user.UserGroupIds

	if shouldUpdate {
		user.ProgUserId = spotinst.String(id)
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
			var accountViewerPolicy organization.ProgPolicy
			accountViewerPolicy.PolicyId = spotinst.String("3")
			accountViewerPolicy.AccountIds = accountIds
			policies = append(policies, &accountViewerPolicy)
			var policiesToUpdate []*organization.UserPolicy
			var policyToUpdate *organization.UserPolicy
			for _, policy := range policies {
				policyToUpdate = &organization.UserPolicy{
					PolicyId:   policy.PolicyId,
					AccountIds: policy.AccountIds,
				}
				policiesToUpdate = append(policiesToUpdate, policyToUpdate)
			}
			if err := updatePolicyMapping(policiesToUpdate, &id, meta.(*Client)); err != nil {
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
	return resourceOrgProgUserRead(ctx, resourceData, meta)
}
