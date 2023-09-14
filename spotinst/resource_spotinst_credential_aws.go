package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
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
		ReadContext:   resourceSpotinstCredentialAWSRead,
		DeleteContext: schema.NoopContext,

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
	resourceData.SetId(spotinst.StringValue(credential.AccountId))

	err = createAWSCredential(credential, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Account created successfully: %s <===", resourceData.Id())
	return resourceSpotinstCredentialAWSRead(ctx, resourceData, meta)

}

func createAWSCredential(credential *aws.Credential, spotinstClient *Client) error {
	if json, err := commons.ToJson(credential); err != nil {
		return err
	} else {
		log.Printf("===> Set Credential configuration: %s", json)
	}
	input := &aws.SetCredentialInput{Credential: credential}
	_, err := spotinstClient.account.CloudProviderAWS().SetCredential(context.Background(), input)
	if err != nil {
		return fmt.Errorf("[ERROR] failed to set credential: %s", err)
	}
	return nil
}

func resourceSpotinstCredentialAWSRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.AccountAWSResource.GetName(), id)
	input := &aws.ReadCredentialInput{AccountId: spotinst.String(id)}
	resp, err := meta.(*Client).account.CloudProviderAWS().ReadCredential(context.Background(), input)
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
	credentialResponse := resp.Credential
	if credentialResponse == nil {
		resourceData.SetId("")
		return nil
	}
	if err := commons.CredentialAWSResource.OnRead(credentialResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> Credential read successfully: %s <===", id)
	return nil
}
