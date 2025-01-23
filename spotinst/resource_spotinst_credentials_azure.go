package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/account/providers/azure"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/credentials_azure"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceSpotinstCredentialsAzure() *schema.Resource {
	setupCredentialsAzureResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstCredentialsAzureCreate,
		ReadContext:   resourceSpotinstCredentialsAzureRead,
		DeleteContext: schema.NoopContext,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: commons.CredentialsAzureResource.GetSchemaMap(),
	}
}

func setupCredentialsAzureResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	credentials_azure.Setup(fieldsMap)

	commons.CredentialsAzureResource = commons.NewCredentialsAzureResource(fieldsMap)
}

func resourceSpotinstCredentialsAzureCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.CredentialsAzureResource.GetName())

	credentials, err := commons.CredentialsAzureResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	resourceData.SetId(spotinst.StringValue(credentials.AccountId))

	err = createAzureCredentials(credentials, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Credentials set successfully: %s <===", resourceData.Id())
	return resourceSpotinstCredentialsAzureRead(ctx, resourceData, meta)

}

func createAzureCredentials(credentials *azure.Credentials, spotinstClient *Client) error {
	if json, err := commons.ToJson(credentials); err != nil {
		return err
	} else {
		log.Printf("===> Credentials configuration: %s", json)
	}
	input := &azure.SetCredentialsInput{Credentials: credentials}
	_, err := spotinstClient.account.CloudProviderAzure().SetCredentials(context.Background(), input)
	if err != nil {
		return fmt.Errorf("[ERROR] failed to set credential: %s", err)
	}
	return nil
}

func resourceSpotinstCredentialsAzureRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.CredentialsAzureResource.GetName(), id)
	input := &azure.ReadCredentialsInput{AccountId: spotinst.String(id)}
	resp, err := meta.(*Client).account.CloudProviderAzure().ReadCredentials(context.Background(), input)
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
		return diag.Errorf("failed to read credentials: %s", err)
	}
	// if nothing was found, return no state
	credentialsResponse := resp.Credentials
	if credentialsResponse == nil {
		resourceData.SetId("")
		return nil
	}
	if err := commons.CredentialsAzureResource.OnRead(credentialsResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> Credentials read successfully: %s <===", id)
	return nil
}
