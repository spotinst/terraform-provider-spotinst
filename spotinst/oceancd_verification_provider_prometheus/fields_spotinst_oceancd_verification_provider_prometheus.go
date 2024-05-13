package oceancd_verification_provider_prometheus

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Prometheus] = commons.NewGenericField(
		commons.OceanCDVerificationProviderPrometheus,
		Prometheus,
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
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var result []interface{} = nil

			if verificationProvider != nil && verificationProvider.Prometheus != nil {
				result = flattenPrometheus(verificationProvider.Prometheus)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Prometheus), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Prometheus), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var value *oceancd.Prometheus = nil

			if v, ok := resourceData.GetOk(string(Prometheus)); ok {
				if prometheus, err := expandPrometheus(v); err != nil {
					return err
				} else {
					value = prometheus
				}
			}
			verificationProvider.SetPrometheus(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var value *oceancd.Prometheus = nil

			if v, ok := resourceData.GetOk(string(Prometheus)); ok {
				if prometheus, err := expandPrometheus(v); err != nil {
					return err
				} else {
					value = prometheus
				}
			}
			verificationProvider.SetPrometheus(value)
			return nil
		},
		nil,
	)
}

func expandPrometheus(data interface{}) (*oceancd.Prometheus, error) {

	prometheus := &oceancd.Prometheus{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return prometheus, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Address)].(string); ok && v != "" {
		prometheus.SetAddress(spotinst.String(v))
	}

	return prometheus, nil
}

func flattenPrometheus(prometheus *oceancd.Prometheus) []interface{} {
	var out []interface{}

	if prometheus != nil {
		result := make(map[string]interface{})

		if prometheus.Address != nil {
			result[string(Address)] = spotinst.StringValue(prometheus.Address)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}
