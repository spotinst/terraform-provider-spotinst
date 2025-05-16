package credentials_azure

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[AccountId] = commons.NewGenericField(
		commons.CredentialsAzure,
		AccountId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
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
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetAccountId(spotinst.String(resourceData.Get(string(AccountId)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetAccountId(spotinst.String(resourceData.Get(string(AccountId)).(string)))
			return nil
		},
		nil,
	)
	fieldsMap[ClientId] = commons.NewGenericField(
		commons.CredentialsAzure,
		ClientId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
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
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetClientId(spotinst.String(resourceData.Get(string(ClientId)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetClientId(spotinst.String(resourceData.Get(string(ClientId)).(string)))
			return nil
		},
		nil,
	)
	fieldsMap[ClientSecret] = commons.NewGenericField(
		commons.CredentialsAzure,
		ClientSecret,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.ClientSecret != nil {
				value = credentials.ClientSecret
			}
			if err := resourceData.Set(string(ClientSecret), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ClientSecret), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetClientSecret(spotinst.String(resourceData.Get(string(ClientSecret)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetClientSecret(spotinst.String(resourceData.Get(string(ClientSecret)).(string)))
			return nil
		},
		nil,
	)
	fieldsMap[TenantId] = commons.NewGenericField(
		commons.CredentialsAzure,
		TenantId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.TenantId != nil {
				value = credentials.TenantId
			}
			if err := resourceData.Set(string(TenantId), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(TenantId), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetTenantId(spotinst.String(resourceData.Get(string(TenantId)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetTenantId(spotinst.String(resourceData.Get(string(TenantId)).(string)))
			return nil
		},
		nil,
	)
	fieldsMap[SubscriptionId] = commons.NewGenericField(
		commons.CredentialsAzure,
		SubscriptionId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.SubscriptionId != nil {
				value = credentials.SubscriptionId
			}
			if err := resourceData.Set(string(SubscriptionId), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SubscriptionId), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetSubscriptionId(spotinst.String(resourceData.Get(string(SubscriptionId)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetSubscriptionId(spotinst.String(resourceData.Get(string(SubscriptionId)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[ExpirationDate] = commons.NewGenericField(
		commons.CredentialsAzure,
		ExpirationDate,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
			Default:  -1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			var value *string = nil
			if credentials.ExpirationDate != nil {
				value = credentials.ExpirationDate
			}
			if err := resourceData.Set(string(ExpirationDate), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ExpirationDate), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetExpirationDate(spotinst.String(resourceData.Get(string(ExpirationDate)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			credentialsWrapper := resourceObject.(*commons.AzureCredentialsWrapper)
			credentials := credentialsWrapper.GetCredentials()
			credentials.SetExpirationDate(spotinst.String(resourceData.Get(string(ExpirationDate)).(string)))
			return nil
		},
		nil,
	)

}
