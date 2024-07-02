package oceancd_rollout_spec

import (
	"fmt"

	"github.com/spotinst/spotinst-sdk-go/service/oceancd"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanCDRolloutSpec,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			var value *string = nil
			if rolloutSpec.Name != nil {
				value = rolloutSpec.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			if value, ok := resourceData.Get(string(Name)).(string); ok && value != "" {
				rolloutSpec.SetName(spotinst.String(value))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[FailurePolicy] = commons.NewGenericField(
		commons.OceanCDStrategyCanary,
		FailurePolicy,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Action): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			var result []interface{} = nil

			if rolloutSpec != nil && rolloutSpec.FailurePolicy != nil {
				result = flattenFailurePolicy(rolloutSpec.FailurePolicy)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(FailurePolicy), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(FailurePolicy), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			var value *oceancd.FailurePolicy = nil

			if v, ok := resourceData.GetOkExists(string(FailurePolicy)); ok {
				if rolloutSpec, err := expandFailurePolicy(v); err != nil {
					return err
				} else {
					value = rolloutSpec
				}
			}
			rolloutSpec.SetFailurePolicy(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			var value *oceancd.FailurePolicy = nil
			if v, ok := resourceData.GetOkExists(string(FailurePolicy)); ok {
				if canary, err := expandFailurePolicy(v); err != nil {
					return err
				} else {
					value = canary
				}
			}
			rolloutSpec.SetFailurePolicy(value)
			return nil
		},
		nil,
	)
}

func expandFailurePolicy(data interface{}) (*oceancd.FailurePolicy, error) {
	if list := data.([]interface{}); len(list) > 0 {
		rolloutSpec := &oceancd.FailurePolicy{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})
			if v, ok := m[string(Action)].(string); ok && v != "" {
				rolloutSpec.SetAction(spotinst.String(v))
			}
		}
		return rolloutSpec, nil
	}
	return nil, nil
}

func flattenFailurePolicy(failurePolicy *oceancd.FailurePolicy) []interface{} {
	var response []interface{}

	if failurePolicy != nil {
		result := make(map[string]interface{})

		result[string(Action)] = spotinst.StringValue(failurePolicy.Action)

		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}
