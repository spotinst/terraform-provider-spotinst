package spotinst

import (
	"context"
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/administration"
	administrationPackage "github.com/spotinst/terraform-provider-spotinst/spotinst/organziation_programmatic_user"
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
		//UpdateContext: resourceOrgProgUserUpdate,
		ReadContext:   resourceOrgProgUserRead,
		DeleteContext: resourceOrgProgUserDelete,

		Schema: commons.OrgProgUserResource.GetSchemaMap(),
	}
}

func setupOrgProgUser() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	administrationPackage.Setup(fieldsMap)

	commons.OrgProgUserResource = commons.NewOrgProgUserResource(fieldsMap)
}

func resourceOrgProgUserDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OrgProgUserResource.GetName(), id)

	input := &administration.DeleteUserInput{UserID: spotinst.String(id)}
	if _, err := meta.(*Client).administration.DeleteUser(context.Background(), input); err != nil {
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
	input := &administration.ReadUserInput{UserID: spotinst.String(resourceData.Id())}
	userResponse, err := client.administration.ReadProgUser(context.Background(), input)
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

	userId, err := createProgUser(progUser, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(userId))
	log.Printf("===> User created successfully: %s <===", resourceData.Id())

	return resourceOrgProgUserRead(ctx, resourceData, meta)
}

func createProgUser(userObj *administration.ProgrammaticUser, spotinstClient *Client) (*string, error) {
	input := userObj
	resp, err := spotinstClient.administration.CreateProgUser(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create user: %s", err)
	}
	return resp.ProgrammaticUser.ProgUserId, nil
}

func resourceOrgProgUserUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceOrgProgUserRead(ctx, resourceData, meta)
}
