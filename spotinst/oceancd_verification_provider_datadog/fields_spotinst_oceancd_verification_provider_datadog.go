package oceancd_verification_provider_datadog

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Datadog] = commons.NewGenericField(
		commons.OceanCDVerificationProviderDataDog,
		Datadog,
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
				if err := resourceData.Set(string(Datadog), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Datadog), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var value *oceancd.DataDog = nil

			if v, ok := resourceData.GetOk(string(Datadog)); ok {
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

			if v, ok := resourceData.GetOk(string(Datadog)); ok {
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
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return datadog, nil
	}
	result := list[0].(map[string]interface{})

	if v, ok := result[string(Address)].(string); ok && v != "" {
		datadog.SetAddress(spotinst.String(v))
	}

	if v, ok := result[string(ApiKey)].(string); ok && v != "" {
		datadog.SetApiKey(spotinst.String(v))
	}

	if v, ok := result[string(AppKey)].(string); ok && v != "" {
		datadog.SetAppKey(spotinst.String(v))
	}

	return datadog, nil
}

func flattenDataDog(datadogvp *oceancd.DataDog) []interface{} {
	var out []interface{}

	if datadogvp != nil {
		result := make(map[string]interface{})

		if datadogvp.Address != nil {
			result[string(Address)] = spotinst.StringValue(datadogvp.Address)
		}
		if datadogvp.ApiKey != nil {
			result[string(ApiKey)] = spotinst.StringValue(datadogvp.ApiKey)
		}
		if datadogvp.AppKey != nil {
			result[string(AppKey)] = spotinst.StringValue(datadogvp.AppKey)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}
