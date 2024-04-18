package oceancd_verification_provider_datadog

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[DataDog] = commons.NewGenericField(
		commons.OceanCDVerificationProviderDataDog,
		DataDog,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Address): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(ApiKey): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(AppKey): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var result []interface{} = nil

			if verificationProvider != nil && verificationProvider.DataDog != nil {
				result = flattenDataDog(verificationProvider.DataDog)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(DataDog), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DataDog), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var value *oceancd.DataDog = nil

			if v, ok := resourceData.GetOk(string(DataDog)); ok {
				if datadog, err := expandDataDog(v); err != nil {
					return err
				} else {
					value = datadog
				}
			}
			verificationProvider.SetDataDog(value)
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var value *oceancd.DataDog = nil

			if v, ok := resourceData.GetOk(string(DataDog)); ok {
				if datadog, err := expandDataDog(v); err != nil {
					return err
				} else {
					value = datadog
				}
			}
			verificationProvider.SetDataDog(value)
			return nil
		},
		nil,
	)
}

func expandDataDog(data interface{}) (*oceancd.DataDog, error) {
	datadog := &oceancd.DataDog{}
	list := data.(*schema.Set).List()
	if len(list) > 0 {
		if list != nil && list[0] != nil {
			result := list[0].(map[string]interface{})

			if v, ok := result[string(Address)].(string); ok && v != "" {
				datadog.SetAddress(spotinst.String(v))
			} else {
				datadog.SetAddress(nil)
			}

			if v, ok := result[string(ApiKey)].(string); ok && v != "" {
				datadog.SetApiKey(spotinst.String(v))
			} else {
				datadog.SetApiKey(nil)
			}

			if v, ok := result[string(AppKey)].(string); ok && v != "" {
				datadog.SetAppKey(spotinst.String(v))
			} else {
				datadog.SetAppKey(nil)
			}
		}
		return datadog, nil
	}
	return nil, nil
}

func flattenDataDog(datadogvp *oceancd.DataDog) []interface{} {
	datadog := make(map[string]interface{})
	datadog[string(Address)] = spotinst.StringValue(datadogvp.Address)
	datadog[string(ApiKey)] = spotinst.StringValue(datadogvp.ApiKey)
	datadog[string(AppKey)] = spotinst.StringValue(datadogvp.AppKey)
	return []interface{}{datadog}
}
