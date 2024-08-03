package commons

import (
	"fmt"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/account/providers/gcp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	CredentialsGCPResourceName ResourceName = "spotinst_credentials_gcp"
)

var CredentialsGCPResource *CredentialsGCPTerraformResource

type CredentialsGCPTerraformResource struct {
	GenericResource
}

type GCPCredentialsWrapper struct {
	credentials *gcp.ServiceAccounts
}

func NewCredentialsGCPResource(fieldsMap map[FieldName]*GenericField) *CredentialsGCPTerraformResource {
	return &CredentialsGCPTerraformResource{
		GenericResource: GenericResource{
			resourceName: CredentialsGCPResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *CredentialsGCPTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*gcp.ServiceAccounts, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	credentialsWrapper := NewCredentialsGCPWrapper()

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

func (res *CredentialsGCPTerraformResource) OnRead(
	credentials *gcp.ServiceAccounts,
	resourceData *schema.ResourceData,
	meta interface{}) error {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}
	credentialsWrapper := NewCredentialsGCPWrapper()
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

func NewCredentialsGCPWrapper() *GCPCredentialsWrapper {
	return &GCPCredentialsWrapper{
		credentials: &gcp.ServiceAccounts{},
	}
}

func (credentialsWrapper *GCPCredentialsWrapper) GetCredentials() *gcp.ServiceAccounts {
	return credentialsWrapper.credentials
}

func (credentialsWrapper *GCPCredentialsWrapper) SetCredentials(credentials *gcp.ServiceAccounts) {
	credentialsWrapper.credentials = credentials
}
