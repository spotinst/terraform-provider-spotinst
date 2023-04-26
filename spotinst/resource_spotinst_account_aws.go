package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/account_aws"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceSpotinstAccountAWS() *schema.Resource {
	setupAccountAWSResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstAccountAWSCreate,
		ReadContext:   resourceSpotinstAccountAWSRead,
		UpdateContext: resourceSpotinstAccountAWSUpdate,
		DeleteContext: resourceSpotinstAccountAWSDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: commons.AccountAWSResource.GetSchemaMap(),
	}
}

func setupAccountAWSResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	account_aws.Setup(fieldsMap)

	commons.AccountAWSResource = commons.NewAccountAWSResource(fieldsMap)
}

func resourceSpotinstAccountAWSCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.AccountAWSResource.GetName())

	account, err := commons.AccountAWSResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	accountID, err := createAWSAccount(account, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(accountID))

	log.Printf("===> Account created successfully: %s <===", resourceData.Id())
	return resourceSpotinstAccountAWSRead(ctx, resourceData, meta)
}

func createAWSAccount(account *aws.Account, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(account); err != nil {
		return nil, err
	} else {
		log.Printf("===> Account create configuration: %s", json)
	}
	var resp *aws.CreateAccountOutput = nil
	input := &aws.CreateAccountInput{Account: account}
	println(input)
	resp, err := spotinstClient.account.CloudProviderAWS().CreateAccount(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create account: %s", err)
	}
	return resp.Account.ID, nil
}

const ErrCodeAccountNotFound = "Account_DOESNT_EXIST"

func resourceSpotinstAccountAWSRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.AccountAWSResource.GetName(), id)

	input := &aws.ReadAccountInput{AccountID: spotinst.String(id)}
	resp, err := meta.(*Client).account.CloudProviderAWS().ReadAccount(context.Background(), input)

	if err != nil {
		// If the account was not found, return nil so that we can show
		// that the account does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeAccountNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return diag.Errorf("failed to read account: %s", err)
	}

	// if nothing was found, return no state
	accountResponse := resp.Account
	if accountResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.AccountAWSResource.OnRead(accountResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> Account read successfully: %s <===", id)
	return nil
}

func resourceSpotinstAccountAWSDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.AccountAWSResource.GetName(), id)

	if err := deleteAWSAccount(resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Account deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteAWSAccount(resourceData *schema.ResourceData, meta interface{}) error {
	accountID := resourceData.Id()
	input := &aws.DeleteAccountInput{
		AccountID: spotinst.String(accountID),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Account delete configuration: %s", json)
	}

	if _, err := meta.(*Client).account.CloudProviderAWS().DeleteAccount(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete account: %s", err)
	}
	return nil
}

func resourceSpotinstAccountAWSUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.AccountAWSResource.GetName(), id)

	shouldUpdate, changesRequiredRoll, tagsChanged, account, err := commons.AccountAWSResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		account.SetId(spotinst.String(id))
		if err := updateAWSAccount(account, resourceData, meta, changesRequiredRoll, tagsChanged); err != nil {
			return diag.FromErr(err)
		}
	}
	log.Printf("===> Account updated successfully: %s <===", id)
	return resourceSpotinstAccountAWSRead(ctx, resourceData, meta)
}

func updateAWSAccount(account *aws.Account, resourceData *schema.ResourceData, meta interface{}, changesRequiredRoll bool, tagsChanged bool) error {
	var input = &aws.UpdateAccountInput{
		Account: account,
	}
	accountID := resourceData.Id()

	if json, err := commons.ToJson(account); err != nil {
		return err
	} else {
		log.Printf("===> Account update configuration: %s", json)
	}

	if _, err := meta.(*Client).account.CloudProviderAWS().UpdateAccount(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update account [%v]: %v", accountID, err)
	}

	return nil
}
