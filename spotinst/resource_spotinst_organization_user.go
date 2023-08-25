package spotinst

import (
	"context"
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/administration"
	administrationPackage "github.com/spotinst/terraform-provider-spotinst/spotinst/organization_user"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceOrgUser() *schema.Resource {
	setupOrgUser()
	return &schema.Resource{
		CreateContext: resourceOrgUserCreate,
		//UpdateContext: resourceOrgUserUpdate,
		ReadContext:   resourceOrgUserRead,
		DeleteContext: resourceOrgUserDelete,

		Schema: commons.OrgUserResource.GetSchemaMap(),
	}
}

func setupOrgUser() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	administrationPackage.Setup(fieldsMap)

	commons.OrgUserResource = commons.NewOrgUserResource(fieldsMap)
}

func resourceOrgUserDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OrgUserResource.GetName(), id)

	input := &administration.DeleteUserInput{UserID: spotinst.String(id)}
	if _, err := meta.(*Client).administration.DeleteUser(context.Background(), input); err != nil {
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
	input := &administration.ReadUserInput{UserID: spotinst.String(resourceData.Id())}
	userResponse, err := client.administration.ReadUser(context.Background(), input)
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
	if err != nil {
		return diag.FromErr(err)
	}

	userId, err := createUser(user, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(userId))
	log.Printf("===> User created successfully: %s <===", resourceData.Id())

	return resourceOrgUserRead(ctx, resourceData, meta)
}

func createUser(userObj *administration.User, spotinstClient *Client) (*string, error) {
	input := userObj
	resp, err := spotinstClient.administration.CreateUser(context.Background(), input, spotinst.Bool(true))
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create user: %s", err)
	}
	return resp.User.UserID, nil
}

func resourceOrgUserUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.OrgUserResource.GetName(), id)

	shouldUpdate, user, err := commons.OrgUserResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		user.UserID = spotinst.String(id)
		if err := updateUser(user, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("===> User updated successfully: %s <===", id)
	return resourceOrgUserRead(ctx, resourceData, meta)
}

func updateUser(user *administration.User, resourceData *schema.ResourceData, meta interface{}) error {
	/*input := user

	if json, err := commons.ToJson(user); err != nil {
		return err
	} else {
		log.Printf("===> user update configuration: %s", json)
	}

	if _, err := meta.(*Client).administration.Update(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] failed to update user %s: %s", resourceData.Id(), err)
	}*/
	return nil
}
