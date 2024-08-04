package credentials_gcp

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[AccountId] = commons.NewGenericField(
		commons.CredentialsGCP,
		AccountId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
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
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetAccountId(spotinst.String(resourceData.Get(string(AccountId)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetAccountId(spotinst.String(resourceData.Get(string(AccountId)).(string)))
			return nil
		},
		nil,
	)
	fieldsMap[Type] = commons.NewGenericField(
		commons.CredentialsGCP,
		Type,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.Type != nil {
				value = credentials.Type
			}
			if err := resourceData.Set(string(Type), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Type), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetType(spotinst.String(resourceData.Get(string(Type)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetType(spotinst.String(resourceData.Get(string(Type)).(string)))
			return nil
		},
		nil,
	)
	fieldsMap[ProjectId] = commons.NewGenericField(
		commons.CredentialsGCP,
		ProjectId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.ProjectId != nil {
				value = credentials.ProjectId
			}
			if err := resourceData.Set(string(ProjectId), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ProjectId), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetProjectId(spotinst.String(resourceData.Get(string(ProjectId)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetProjectId(spotinst.String(resourceData.Get(string(ProjectId)).(string)))
			return nil
		},
		nil,
	)
	fieldsMap[PrivateKeyId] = commons.NewGenericField(
		commons.CredentialsGCP,
		PrivateKeyId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.PrivateKeyId != nil {
				value = credentials.PrivateKeyId
			}
			if err := resourceData.Set(string(PrivateKeyId), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PrivateKeyId), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetPrivateKeyId(spotinst.String(resourceData.Get(string(PrivateKeyId)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetPrivateKeyId(spotinst.String(resourceData.Get(string(PrivateKeyId)).(string)))
			return nil
		},
		nil,
	)
	fieldsMap[PrivateKey] = commons.NewGenericField(
		commons.CredentialsGCP,
		PrivateKey,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.PrivateKey != nil {
				value = credentials.PrivateKey
			}
			if err := resourceData.Set(string(PrivateKey), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PrivateKey), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetPrivateKey(spotinst.String(resourceData.Get(string(PrivateKey)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetPrivateKey(spotinst.String(resourceData.Get(string(PrivateKey)).(string)))
			return nil
		},
		nil,
	)
	fieldsMap[ClientEmail] = commons.NewGenericField(
		commons.CredentialsGCP,
		ClientEmail,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.ClientEmail != nil {
				value = credentials.ClientEmail
			}
			if err := resourceData.Set(string(ClientEmail), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ClientEmail), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetClientEmail(spotinst.String(resourceData.Get(string(ClientEmail)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetClientEmail(spotinst.String(resourceData.Get(string(ClientEmail)).(string)))
			return nil
		},
		nil,
	)
	fieldsMap[ClientId] = commons.NewGenericField(
		commons.CredentialsGCP,
		ClientId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.ClientId != nil {
				value = credentials.ClientId
			}
			if err := resourceData.Set(string(ClientId), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ClientId), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetClientId(spotinst.String(resourceData.Get(string(ClientId)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetClientId(spotinst.String(resourceData.Get(string(ClientId)).(string)))
			return nil
		},
		nil,
	)
	fieldsMap[AuthUri] = commons.NewGenericField(
		commons.CredentialsGCP,
		AuthUri,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.AuthUri != nil {
				value = credentials.AuthUri
			}
			if err := resourceData.Set(string(AuthUri), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AuthUri), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetAuthUri(spotinst.String(resourceData.Get(string(AuthUri)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetAuthUri(spotinst.String(resourceData.Get(string(AuthUri)).(string)))
			return nil
		},
		nil,
	)
	fieldsMap[TokenUri] = commons.NewGenericField(
		commons.CredentialsGCP,
		TokenUri,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.TokenUri != nil {
				value = credentials.TokenUri
			}
			if err := resourceData.Set(string(TokenUri), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(TokenUri), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetTokenUri(spotinst.String(resourceData.Get(string(TokenUri)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetTokenUri(spotinst.String(resourceData.Get(string(TokenUri)).(string)))
			return nil
		},
		nil,
	)
	fieldsMap[AuthProviderX509CertUrl] = commons.NewGenericField(
		commons.CredentialsGCP,
		AuthProviderX509CertUrl,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.AuthProviderX509CertUrl != nil {
				value = credentials.AuthProviderX509CertUrl
			}
			if err := resourceData.Set(string(AuthProviderX509CertUrl), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AuthProviderX509CertUrl), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetAuthProviderX509CertUrl(spotinst.String(resourceData.Get(string(AuthProviderX509CertUrl)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetAuthProviderX509CertUrl(spotinst.String(resourceData.Get(string(AuthProviderX509CertUrl)).(string)))
			return nil
		},
		nil,
	)
	fieldsMap[ClientX509CertUrl] = commons.NewGenericField(
		commons.CredentialsGCP,
		ClientX509CertUrl,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.ClientX509CertUrl != nil {
				value = credentials.ClientX509CertUrl
			}
			if err := resourceData.Set(string(ClientX509CertUrl), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ClientX509CertUrl), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetClientX509CertUrl(spotinst.String(resourceData.Get(string(ClientX509CertUrl)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.GCPCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetClientX509CertUrl(spotinst.String(resourceData.Get(string(ClientX509CertUrl)).(string)))
			return nil
		},
		nil,
	)

}
