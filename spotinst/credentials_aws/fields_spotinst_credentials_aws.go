package credentials_aws

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[IamRole] = commons.NewGenericField(
		commons.CredentialsAWS,
		IamRole,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AWSCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.IamRole != nil {
				value = credentials.IamRole
			}
			if err := resourceData.Set(string(IamRole), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IamRole), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AWSCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetIamRole(spotinst.String(resourceData.Get(string(IamRole)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AWSCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetIamRole(spotinst.String(resourceData.Get(string(IamRole)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[AccountId] = commons.NewGenericField(
		commons.CredentialsAWS,
		AccountId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AWSCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.AccountId != nil {
				value = credentials.AccountId
			}
			if err := resourceData.Set(string(AccountId), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AccountId), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AWSCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetAccountId(spotinst.String(resourceData.Get(string(AccountId)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AWSCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetAccountId(spotinst.String(resourceData.Get(string(AccountId)).(string)))
			return nil
		},
		nil,
	)

}
