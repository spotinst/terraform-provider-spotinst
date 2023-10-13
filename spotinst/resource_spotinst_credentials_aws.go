package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/credentials_aws"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/aws"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceSpotinstCredentialsAWS() *schema.Resource {
	setupCredentialsAWSResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstCredentialsAWSCreate,
		ReadContext:   resourceSpotinstCredentialsAWSRead,
		DeleteContext: schema.NoopContext,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: commons.CredentialsAWSResource.GetSchemaMap(),
	}
}

func setupCredentialsAWSResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	credentials_aws.Setup(fieldsMap)

	commons.CredentialsAWSResource = commons.NewCredentialsAWSResource(fieldsMap)
}

func resourceSpotinstCredentialsAWSCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.CredentialsAWSResource.GetName())

	credentials, err := commons.CredentialsAWSResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	resourceData.SetId(spotinst.StringValue(credentials.AccountId))

	err = createAWSCredentials(credentials, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Account created successfully: %s <===", resourceData.Id())
	return resourceSpotinstCredentialsAWSRead(ctx, resourceData, meta)

}

func createAWSCredentials(credentials *aws.Credentials, spotinstClient *Client) error {
	if json, err := commons.ToJson(credentials); err != nil {
		return err
	} else {
		log.Printf("===> Credentials configuration: %s", json)
	}
	input := &aws.SetCredentialsInput{Credentials: credentials}
	_, err := spotinstClient.account.CloudProviderAWS().Credentials(context.Background(), input)
	if err != nil {
		return fmt.Errorf("[ERROR] failed to set credential: %s", err)
	}
	return nil
}

func resourceSpotinstCredentialsAWSRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.AccountAWSResource.GetName(), id)
	input := &aws.ReadCredentialsInput{AccountId: spotinst.String(id)}
	resp, err := meta.(*Client).account.CloudProviderAWS().ReadCredentials(context.Background(), input)
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
	credentialsResponse := resp.Credentials
	if credentialsResponse == nil {
		resourceData.SetId("")
		return nil
	}
	if err := commons.CredentialsAWSResource.OnRead(credentialsResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> Credentials read successfully: %s <===", id)
	return nil
}
