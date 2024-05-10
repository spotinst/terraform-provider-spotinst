package oceancd_verification_provider_new_relic

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[NewRelic] = commons.NewGenericField(
		commons.OceanCDVerificationProviderNewRelic,
		NewRelic,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(AccountId): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(BaseUrlNerdGraph): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(BaseUrlRest): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(PersonalApiKey): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(Region): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var result []interface{} = nil

			if verificationProvider != nil && verificationProvider.NewRelic != nil {
				result = flattenNewRelic(verificationProvider.NewRelic)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(NewRelic), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(NewRelic), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var value *oceancd.NewRelic = nil

			if v, ok := resourceData.GetOk(string(NewRelic)); ok {
				if newRelic, err := expandNewRelic(v); err != nil {
					return err
				} else {
					value = newRelic
				}
			}
			verificationProvider.SetNewRelic(value)
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var value *oceancd.NewRelic = nil

			if v, ok := resourceData.GetOk(string(NewRelic)); ok {
				if newRelic, err := expandNewRelic(v); err != nil {
					return err
				} else {
					value = newRelic
				}
			}
			verificationProvider.SetNewRelic(value)
			return nil
		},
		nil,
	)
}

func expandNewRelic(data interface{}) (*oceancd.NewRelic, error) {
	newRelic := &oceancd.NewRelic{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return newRelic, nil
	}
	result := list[0].(map[string]interface{})

	if v, ok := result[string(AccountId)].(string); ok && v != "" {
		newRelic.SetAccountId(spotinst.String(v))
	}

	if v, ok := result[string(BaseUrlNerdGraph)].(string); ok && v != "" {
		newRelic.SetBaseUrlNerdGraph(spotinst.String(v))
	} else {
		newRelic.SetBaseUrlNerdGraph(nil)
	}

	if v, ok := result[string(BaseUrlRest)].(string); ok && v != "" {
		newRelic.SetBaseUrlRest(spotinst.String(v))
	} else {
		newRelic.SetBaseUrlRest(nil)
	}

	if v, ok := result[string(PersonalApiKey)].(string); ok && v != "" {
		newRelic.SetPersonalApiKey(spotinst.String(v))
	}

	if v, ok := result[string(Region)].(string); ok && v != "" {
		newRelic.SetRegion(spotinst.String(v))
	} else {
		newRelic.SetRegion(nil)
	}

	return newRelic, nil
}

func flattenNewRelic(newRelic *oceancd.NewRelic) []interface{} {
	var out []interface{}

	if newRelic != nil {
		result := make(map[string]interface{})

		if newRelic.AccountId != nil {
			result[string(AccountId)] = spotinst.StringValue(newRelic.AccountId)
		}
		if newRelic.BaseUrlNerdGraph != nil {
			result[string(BaseUrlNerdGraph)] = spotinst.StringValue(newRelic.BaseUrlNerdGraph)
		}
		if newRelic.BaseUrlRest != nil {
			result[string(BaseUrlRest)] = spotinst.StringValue(newRelic.BaseUrlRest)
		}
		if newRelic.PersonalApiKey != nil {
			result[string(PersonalApiKey)] = spotinst.StringValue(newRelic.PersonalApiKey)
		}
		if newRelic.Region != nil {
			result[string(Region)] = spotinst.StringValue(newRelic.Region)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}
