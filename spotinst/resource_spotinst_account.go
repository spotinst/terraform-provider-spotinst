package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/account/providers/common"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/account"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceSpotinstAccount() *schema.Resource {
	setupAccountResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstAccountCreate,
		ReadContext:   resourceSpotinstAccountRead,
		UpdateContext: schema.NoopContext,
		DeleteContext: resourceSpotinstAccountDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: commons.AccountResource.GetSchemaMap(),
	}
}

func setupAccountResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	account.Setup(fieldsMap)

	commons.AccountResource = commons.NewAccountResource(fieldsMap)
}

func resourceSpotinstAccountCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.AccountResource.GetName())

	account, err := commons.AccountResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	accountID, err := createAccount(account, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(accountID))

	log.Printf("===> Account created successfully: %s <===", resourceData.Id())
	return resourceSpotinstAccountRead(ctx, resourceData, meta)
}

func createAccount(account *common.Account, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(account); err != nil {
		return nil, err
	} else {
		log.Printf("===> Account create configuration: %s", json)
	}

	var output *common.CreateAccountOutput = nil
	input := &common.CreateAccountInput{Account: account}
	output, err := spotinstClient.account.CloudProviderCommon().CreateAccount(context.Background(), input)

	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create account: %s", err)
	}
	return output.Account.ID, nil
}

const ErrAccountNotFound = "Account_DOESNT_EXIST"

func resourceSpotinstAccountRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.AccountResource.GetName(), id)

	input := &common.ReadAccountInput{AccountID: spotinst.String(id)}
	output, err := meta.(*Client).account.CloudProviderCommon().ReadAccount(context.Background(), input)

	if err != nil {
		// If the account was not found, return nil so that we can show
		// that the account  does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrAccountNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return diag.Errorf("failed to read account : %s", err)
	}

	// if nothing was found, return no state
	accountResponse := output.Account
	if accountResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.AccountResource.OnRead(accountResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> Account read successfully: %s <===", id)
	return nil
}

func resourceSpotinstAccountDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.AccountResource.GetName(), id)

	if err := deleteAccount(resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Account deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteAccount(resourceData *schema.ResourceData, meta interface{}) error {
	accountID := resourceData.Id()
	input := &common.DeleteAccountInput{
		AccountID: spotinst.String(accountID),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Account delete configuration: %s", json)
	}

	if _, err := meta.(*Client).account.CloudProviderCommon().DeleteAccount(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete account: %s", err)
	}
	return nil
}
