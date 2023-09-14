package credential_aws

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[IamRole] = commons.NewGenericField(
		commons.CredentialAWS,
		IamRole,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialWrapper := resourceObject.(*commons.AWSCredentialWrapper)
			credential := credentialWrapper.GetCredential()
			var value *string = nil
			if credential.IamRole != nil {
				value = credential.IamRole
			}
			if err := resourceData.Set(string(IamRole), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IamRole), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialWrapper := resourceObject.(*commons.AWSCredentialWrapper)
			credential := credentialWrapper.GetCredential()
			credential.SetIamRole(spotinst.String(resourceData.Get(string(IamRole)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialWrapper := resourceObject.(*commons.AWSCredentialWrapper)
			credential := credentialWrapper.GetCredential()
			credential.SetIamRole(spotinst.String(resourceData.Get(string(IamRole)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[AccountId] = commons.NewGenericField(
		commons.CredentialAWS,
		AccountId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialWrapper := resourceObject.(*commons.AWSCredentialWrapper)
			credential := credentialWrapper.GetCredential()
			var value *string = nil
			if credential.AccountId != nil {
				value = credential.AccountId
			}
			if err := resourceData.Set(string(AccountId), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AccountId), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialWrapper := resourceObject.(*commons.AWSCredentialWrapper)
			credential := credentialWrapper.GetCredential()
			credential.SetAccountId(spotinst.String(resourceData.Get(string(AccountId)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialWrapper := resourceObject.(*commons.AWSCredentialWrapper)
			credential := credentialWrapper.GetCredential()
			credential.SetAccountId(spotinst.String(resourceData.Get(string(AccountId)).(string)))
			return nil
		},
		nil,
	)

}
