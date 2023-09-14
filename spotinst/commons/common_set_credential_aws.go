package commons

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/aws"
	"log"
)

const (
	SetCredentialAWSResourceName ResourceName = "spotinst_set_credential_aws"
)

var CredentialAWSResource *CredentialAWSTerraformResource

type CredentialAWSTerraformResource struct {
	GenericResource
}

type AWSCredentialWrapper struct {
	credential *aws.Credential
}

func NewCredentialAWSResource(fieldsMap map[FieldName]*GenericField) *CredentialAWSTerraformResource {
	return &CredentialAWSTerraformResource{
		GenericResource: GenericResource{
			resourceName: SetCredentialAWSResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *CredentialAWSTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.Credential, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	credentialWrapper := NewCredentialWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(credentialWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return credentialWrapper.GetCredential(), nil
}

func NewCredentialWrapper() *AWSCredentialWrapper {
	return &AWSCredentialWrapper{
		credential: &aws.Credential{},
	}
}

func (credentialWrapper *AWSCredentialWrapper) GetCredential() *aws.Credential {
	return credentialWrapper.credential
}

func (credentialWrapper *AWSCredentialWrapper) SetCredential(credential *aws.Credential) {
	credentialWrapper.credential = credential
}

func (res *CredentialAWSTerraformResource) OnRead(
	credential *aws.Credential,
	resourceData *schema.ResourceData,
	meta interface{}) error {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}
	credentialWrapper := NewCredentialWrapper()
	credentialWrapper.SetCredential(credential)
	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(credentialWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}
