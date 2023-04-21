package spotinst

import (
	"context"
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/account_aws"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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

	accountID, err := createAWSAccount(resourceData, account, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(accountID))

	log.Printf("===> Account created successfully: %s <===", resourceData.Id())
	return resourceSpotinstClusterAWSRead(ctx, resourceData, meta)
}

func createAWSAccount(resourceData *schema.ResourceData, account *aws.Account, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(account); err != nil {
		return nil, err
	} else {
		log.Printf("===> Account create configuration: %s", json)
	}

	var resp *aws.CreateAccountOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &aws.CreateAccountInput{Account: account}
		r, err := spotinstClient.account.CloudProviderAWS().CreateAccount(context.Background(), input)
		if err != nil {
			// Checks whether we should retry cluster creation.
			if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
				for _, err := range errs {
					if err.Code == "InvalidParamterValue" &&
						strings.Contains(err.Message, "Invalid IAM Instance Profile") {
						return resource.NonRetryableError(err)
					}
				}
			}
			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create cluster: %s", err)
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
		// If the cluster was not found, return nil so that we can show
		// that the cluster does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeAccountNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return diag.Errorf("failed to read cluster: %s", err)
	}

	// if nothing was found, return no state
	clusterResponse := resp.Account
	if clusterResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.AccountAWSResource.OnRead(clusterResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> Cluster read successfully: %s <===", id)
	return nil
}

func resourceSpotinstAccountAWSDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.AccountAWSResource.GetName(), id)

	if err := deleteAWSAccount(resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Cluster deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteAWSAccount(resourceData *schema.ResourceData, meta interface{}) error {
	clusterID := resourceData.Id()
	input := &aws.DeleteAccountInput{
		AccountID: spotinst.String(clusterID),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Account delete configuration: %s", json)
	}

	if _, err := meta.(*Client).account.CloudProviderAWS().DeleteAccount(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete cluster: %s", err)
	}
	return nil
}
func (s *ServiceOp) DeleteAccount(ctx context.Context, input *DeleteAccountInput) (*DeleteAccountOutput, error) {
	path, err := uritemplates.Expand("/setup/account/{accountId}", uritemplates.Values{
		"AccountId": spotinst.StringValue(input.AccountID),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodDelete, path)

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteAccountOutput{}, nil
}
