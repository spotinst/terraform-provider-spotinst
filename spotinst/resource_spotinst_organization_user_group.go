package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/organization"
	organizationPackage "github.com/spotinst/terraform-provider-spotinst/spotinst/organization_user_group"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceOrgUserGroup() *schema.Resource {
	setupOrgUserGroup()
	return &schema.Resource{
		CreateContext: resourceOrgUserGroupCreate,
		UpdateContext: resourceOrgUserGroupUpdate,
		ReadContext:   resourceOrgUserGroupRead,
		DeleteContext: resourceOrgUserGroupDelete,

		Schema: commons.OrgUserGroupResource.GetSchemaMap(),
	}
}

func setupOrgUserGroup() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	organizationPackage.Setup(fieldsMap)

	commons.OrgUserGroupResource = commons.NewOrgUserGroupResource(fieldsMap)
}

func resourceOrgUserGroupDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OrgUserGroupResource.GetName(), id)

	input := &organization.DeleteUserGroupInput{UserGroupID: spotinst.String(id)}
	if _, err := meta.(*Client).organization.DeleteUserGroup(context.Background(), input); err != nil {
		return diag.Errorf("[ERROR] Failed to delete User Group: %s", err)
	}

	resourceData.SetId("")
	return nil
}

func resourceOrgUserGroupRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.OrgUserGroupResource.GetName(), id)

	client := meta.(*Client)
	input := &organization.ReadUserGroupInput{UserGroupID: spotinst.String(resourceData.Id())}
	userGroupResponse, err := client.organization.ReadUserGroup(context.Background(), input)
	if err != nil {
		return diag.Errorf("[ERROR] Failed to read User Group: %s", err)
	}

	// If nothing was found, then return no state.
	userGroup := userGroupResponse.UserGroup
	if userGroup == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OrgUserGroupResource.OnRead(userGroup, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> User Group read successfully: %s <===", id)
	return nil
}

func resourceOrgUserGroupCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OrgUserGroupResource.GetName())

	userGroup, err := commons.OrgUserGroupResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	userGroupId, err := createUserGroup(userGroup, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(userGroupId))
	log.Printf("===> User Group created successfully: %s <===", resourceData.Id())

	return resourceOrgUserGroupRead(ctx, resourceData, meta)
}

func createUserGroup(userGroupObj *organization.UserGroup, spotinstClient *Client) (*string, error) {
	input := userGroupObj
	resp, err := spotinstClient.organization.CreateUserGroup(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create user group: %s", err)
	}
	return resp.UserGroup.UserGroupId, nil
}

func resourceOrgUserGroupUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.OrgUserGroupResource.GetName(), id)

	shouldUpdate, userGroup, err := commons.OrgUserGroupResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	var policies []*organization.UserGroupPolicy = userGroup.Policies
	userGroup.UserGroupId = spotinst.String(id)

	if shouldUpdate {
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
			var accountViewerPolicy organization.UserGroupPolicy
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
			if err := updatePolicyMappingOfUserGroup(policiesToUpdate, &id, meta.(*Client)); err != nil {
				return diag.FromErr(err)
			}
		}

		if userGroup.UserIds != nil {
			userIds := userGroup.UserIds
			updateUserIdsMapping(userIds, &id, meta.(*Client))
		}

		userGroup.UserIds = nil
		userGroup.Policies = nil

		if userGroup.Name != nil || userGroup.Description != nil {
			if err := updateUserGroup(userGroup, resourceData, meta); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	log.Printf("===> User Group updated successfully: %s <===", id)
	return resourceOrgUserGroupRead(ctx, resourceData, meta)
}

func updatePolicyMappingOfUserGroup(policies []*organization.UserPolicy, userGroupId *string, spotinstClient *Client) error {
	err := spotinstClient.organization.UpdatePolicyMappingOfUserGroup(context.Background(), &organization.UpdatePolicyMappingOfUserGroupInput{
		UserGroupId: userGroupId,
		Policies:    policies,
	})
	if err != nil {
		return fmt.Errorf("[ERROR] failed to update policy mapping for user: %s", err)
	}
	return nil
}

func updateUserIdsMapping(userIds []string, userGroupId *string, spotinstClient *Client) error {
	err := spotinstClient.organization.UpdateUserMappingOfUserGroup(context.Background(), &organization.UpdateUserMappingOfUserGroupInput{
		UserGroupId: userGroupId,
		UserIds:     userIds,
	})
	if err != nil {
		return fmt.Errorf("[ERROR] failed to update policy mapping for user: %s", err)
	}
	return nil
}

func updateUserGroup(userGroup *organization.UserGroup, resourceData *schema.ResourceData, meta interface{}) error {
	input := userGroup
	if json, err := commons.ToJson(userGroup); err != nil {
		return err
	} else {
		log.Printf("===> user group update configuration: %s", json)
	}

	if err := meta.(*Client).organization.UpdateUserGroup(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] failed to update user group %s: %s", resourceData.Id(), err)
	}
	return nil
}
