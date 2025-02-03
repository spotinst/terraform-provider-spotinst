package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/azure"
)

const (
	CredentialsAzureResourceName ResourceName = "spotinst_credentials_azure"
)

var CredentialsAzureResource *CredentialsAzureTerraformResource

type CredentialsAzureTerraformResource struct {
	GenericResource
}

type AzureCredentialsWrapper struct {
	credentials *azure.Credentials
}

func NewCredentialsAzureResource(fieldsMap map[FieldName]*GenericField) *CredentialsAzureTerraformResource {
	return &CredentialsAzureTerraformResource{
		GenericResource: GenericResource{
			resourceName: CredentialsAzureResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *CredentialsAzureTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*azure.Credentials, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	credentialsWrapper := NewCredentialsAzureWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(credentialsWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return credentialsWrapper.GetCredentials(), nil
}

func (res *CredentialsAzureTerraformResource) OnRead(
	credentials *azure.Credentials,
	resourceData *schema.ResourceData,
	meta interface{}) error {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}
	credentialsWrapper := NewCredentialsAzureWrapper()
	credentialsWrapper.SetCredentials(credentials)
	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(credentialsWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func NewCredentialsAzureWrapper() *AzureCredentialsWrapper {
	return &AzureCredentialsWrapper{
		credentials: &azure.Credentials{},
	}
}

func (credentialsWrapper *AzureCredentialsWrapper) GetCredentials() *azure.Credentials {
	return credentialsWrapper.credentials
}

func (credentialsWrapper *AzureCredentialsWrapper) SetCredentials(credentials *azure.Credentials) {
	credentialsWrapper.credentials = credentials
}
