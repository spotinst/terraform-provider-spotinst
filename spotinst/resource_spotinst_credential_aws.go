package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/credential_aws"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/aws"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceSpotinstCredentialAWS() *schema.Resource {
	setupCredentialAWSResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstCredentialAWSCreate,
		ReadContext:   resourceSpotinstAccountAWSRead,
		//UpdateContext: resourceSpotinstAccountAWSUpdate,
		DeleteContext: resourceSpotinstAccountAWSDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: commons.CredentialAWSResource.GetSchemaMap(),
	}
}

func setupCredentialAWSResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	credential_aws.Setup(fieldsMap)

	commons.CredentialAWSResource = commons.NewCredentialAWSResource(fieldsMap)
}

func resourceSpotinstCredentialAWSCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.CredentialAWSResource.GetName())

	credential, err := commons.CredentialAWSResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	err = createAWSCredential(credential, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	//resourceData.SetId(spotinst.StringValue(accountID))

	log.Printf("===> Credential set successfully successfully: %s <===", resourceData.Id())
	//return diag.FromErr(err)
	//return resourceSpotinstAccountAWSRead(ctx, resourceData, meta)
	return nil
}

func createAWSCredential(credential *aws.Credentials, spotinstClient *Client) error {
	if json, err := commons.ToJson(credential); err != nil {
		return err
	} else {
		log.Printf("===> Set credential configuration: %s", json)
	}
	//var resp *aws.CreateCredentialOutput = nil
	input := &aws.SetCredentialInput{Credential: credential}
	println(input)
	//resp, err := spotinstClient.account.CloudProviderAWS().CreateAccount(context.Background(), input)
	err := spotinstClient.account.CloudProviderAWS().SetCredential(context.Background(), input)
	if err != nil {
		return fmt.Errorf("[ERROR] failed to set credential: %s", err)
	}
	//return resp.Account.ID, nil
	return nil
}

//const ErrCodeAccountNotFound = "Account_DOESNT_EXIST"

/*func resourceSpotinstAccountAWSRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
}*/

/*func resourceSpotinstAccountAWSDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
}*/

/*func resourceSpotinstAccountAWSUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
}*/
/*
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
}*/
