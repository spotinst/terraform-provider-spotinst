package commons

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/aws"
	"log"
)

const (
	CredentialsAWSResourceName ResourceName = "spotinst_credentials_aws"
)

var CredentialsAWSResource *CredentialsAWSTerraformResource

type CredentialsAWSTerraformResource struct {
	GenericResource
}

type AWSCredentialsWrapper struct {
	credentials *aws.Credentials
}

func NewCredentialsAWSResource(fieldsMap map[FieldName]*GenericField) *CredentialsAWSTerraformResource {
	return &CredentialsAWSTerraformResource{
		GenericResource: GenericResource{
			resourceName: CredentialsAWSResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *CredentialsAWSTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.Credentials, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	credentialsWrapper := NewCredentialsWrapper()

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

func (res *CredentialsAWSTerraformResource) OnRead(
	credentials *aws.Credentials,
	resourceData *schema.ResourceData,
	meta interface{}) error {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}
	credentialsWrapper := NewCredentialsWrapper()
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

func NewCredentialsWrapper() *AWSCredentialsWrapper {
	return &AWSCredentialsWrapper{
		credentials: &aws.Credentials{},
	}
}

func (credentialsWrapper *AWSCredentialsWrapper) GetCredentials() *aws.Credentials {
	return credentialsWrapper.credentials
}

func (credentialsWrapper *AWSCredentialsWrapper) SetCredentials(credentials *aws.Credentials) {
	credentialsWrapper.credentials = credentials
}
