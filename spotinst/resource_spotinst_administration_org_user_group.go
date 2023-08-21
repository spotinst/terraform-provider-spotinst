package spotinst

import (
	"context"
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/administration"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

	administrationPackage "github.com/spotinst/terraform-provider-spotinst/spotinst/commons/administration_org_user_group"
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

	administrationPackage.Setup(fieldsMap)

	commons.OrgUserGroupResource = commons.NewOrgUserGroupResource(fieldsMap)
}

func resourceOrgUserGroupDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OrgUserGroupResource.GetName(), id)

	input := &administration.DeleteUserGroupInput{UserGroupID: spotinst.String(id)}
	if _, err := meta.(*Client).administration.DeleteUserGroup(context.Background(), input); err != nil {
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
	input := &administration.ReadUserGroupInput{UserGroupID: spotinst.String(resourceData.Id())}
	userGroupResponse, err := client.administration.ReadUserGroup(context.Background(), input)
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

func createUserGroup(userGroupObj *administration.UserGroup, spotinstClient *Client) (*string, error) {
	input := userGroupObj
	resp, err := spotinstClient.administration.CreateUserGroup(context.Background(), input)
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

	if shouldUpdate {
		userGroup.UserGroupId = spotinst.String(id)
		if err := updateUserGroup(userGroup, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("===> User Group updated successfully: %s <===", id)
	return resourceOrgUserGroupRead(ctx, resourceData, meta)
}

func updateUserGroup(userGroup *administration.UserGroup, resourceData *schema.ResourceData, meta interface{}) error {
	input := userGroup
	if json, err := commons.ToJson(userGroup); err != nil {
		return err
	} else {
		log.Printf("===> user group update configuration: %s", json)
	}

	if err := meta.(*Client).administration.UpdateUserGroup(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] failed to update user group %s: %s", resourceData.Id(), err)
	}
	return nil
}
