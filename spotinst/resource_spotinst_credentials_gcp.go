package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/account/providers/gcp"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/credentials_gcp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceSpotinstCredentialsGCP() *schema.Resource {
	setupCredentialsGCPResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstCredentialsGCPCreate,
		ReadContext:   resourceSpotinstCredentialsGCPRead,
		DeleteContext: schema.NoopContext,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: commons.CredentialsGCPResource.GetSchemaMap(),
	}
}

func setupCredentialsGCPResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	credentials_gcp.Setup(fieldsMap)

	commons.CredentialsGCPResource = commons.NewCredentialsGCPResource(fieldsMap)
}

func resourceSpotinstCredentialsGCPCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.CredentialsGCPResource.GetName())

	credentials, err := commons.CredentialsGCPResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	resourceData.SetId(spotinst.StringValue(credentials.AccountId))

	err = createGCPCredentials(credentials, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Account credentials set successfully: %s <===", resourceData.Id())
	return resourceSpotinstCredentialsGCPRead(ctx, resourceData, meta)

}

func createGCPCredentials(credentials *gcp.ServiceAccounts, spotinstClient *Client) error {
	if json, err := commons.ToJson(credentials); err != nil {
		return err
	} else {
		log.Printf("===> Credentials configuration: %s", json)
	}
	input := &gcp.SetServiceAccountsInput{ServiceAccounts: credentials}
	_, err := spotinstClient.account.CloudProviderGCP().SetServiceAccount(context.Background(), input)
	if err != nil {
		return fmt.Errorf("[ERROR] failed to set credential: %s", err)
	}
	return nil
}

const ErrCodeServiceAccountNotFound = "ServiceAccount_DOESNT_EXIST"

func resourceSpotinstCredentialsGCPRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.CredentialsGCPResource.GetName(), id)
	input := &gcp.ReadServiceAccountsInput{AccountId: spotinst.String(id)}
	resp, err := meta.(*Client).account.CloudProviderGCP().ReadServiceAccount(context.Background(), input)
	if err != nil {
		// If the serviceAccount was not found, return nil so that we can show
		// that the credential was not set.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeServiceAccountNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}
		// Some other error, report it.
		return diag.Errorf("failed to read credential: %s", err)
	}
	// if nothing was found, return no state
	credentialsResponse := resp.ServiceAccounts
	if credentialsResponse == nil {
		resourceData.SetId("")
		return nil
	}
	if err := commons.CredentialsGCPResource.OnRead(credentialsResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> Credentials read successfully: %s <===", id)
	return nil
}
