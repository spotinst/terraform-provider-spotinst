package ocean_gke_launch_spec_strategy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Strategy] = commons.NewGenericField(
		commons.OceanGKELaunchSpecStrategy,
		Strategy,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(PreemptiblePercentage): {
						Type:         schema.TypeInt,
						Optional:     true,
						Default:      -1,
						ValidateFunc: validation.IntAtLeast(-1),
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var result interface{} = nil
			if launchSpec.Strategy != nil {
				strategy := launchSpec.Strategy
				result = flattenStrategy(strategy)
			}
			if result != nil {
				if err := resourceData.Set(string(Strategy), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Strategy), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(Strategy)); ok {
				if strategy, err := expandStrategy(value); err != nil {
					return err
				} else {
					launchSpec.SetStrategy(strategy)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var value *gcp.LaunchSpecStrategy = nil
			if v, ok := resourceData.GetOk(string(Strategy)); ok {
				if strategy, err := expandStrategy(v); err != nil {
					return err
				} else {
					value = strategy
				}
			}
			launchSpec.SetStrategy(value)
			return nil
		},
		nil,
	)
}

func expandStrategy(data interface{}) (*gcp.LaunchSpecStrategy, error) {
	strategy := &gcp.LaunchSpecStrategy{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return strategy, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(PreemptiblePercentage)].(int); ok && v > -1 {
		strategy.SetPreemptiblePercentage(spotinst.Int(v))
	}

	return strategy, nil
}

func flattenStrategy(ebs *gcp.LaunchSpecStrategy) []interface{} {
	strategy := make(map[string]interface{})
	if spotinst.IntValue(ebs.PreemptiblePercentage) > -1 {
		strategy[string(PreemptiblePercentage)] = spotinst.IntValue(ebs.PreemptiblePercentage)
	} else {
		strategy[string(PreemptiblePercentage)] = nil
	}
	return []interface{}{strategy}
}
