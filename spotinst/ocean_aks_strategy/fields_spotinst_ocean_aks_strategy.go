package ocean_aks_strategy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Strategy] = commons.NewGenericField(
		commons.OceanAKSStrategy,
		Strategy,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(SpotPercentage): {
						Type:         schema.TypeInt,
						Optional:     true,
						Default:      -1,
						ValidateFunc: validation.IntAtLeast(-1),
					},
					string(FallbackToOnDemand): {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster.Strategy != nil {
				strategy := cluster.Strategy
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
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *azure.Strategy = nil

			if v, ok := resourceData.GetOk(string(Strategy)); ok {
				if strategy, err := expandStrategy(v); err != nil {
					return err
				} else {
					value = strategy
				}
			}
			cluster.SetStrategy(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *azure.Strategy = nil

			if v, ok := resourceData.GetOk(string(Strategy)); ok {
				if strategy, err := expandStrategy(v); err != nil {
					return err
				} else {
					value = strategy
				}
			}
			cluster.SetStrategy(value)
			return nil
		},
		nil,
	)
}

func expandStrategy(data interface{}) (*azure.Strategy, error) {
	if list := data.([]interface{}); len(list) > 0 {
		strategy := &azure.Strategy{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(SpotPercentage)].(int); ok && v > -1 {
				strategy.SetSpotPercentage(spotinst.Int(v))
			} else {
				strategy.SetSpotPercentage(nil)
			}

			if v, ok := m[string(FallbackToOnDemand)].(bool); ok {
				strategy.SetFallbackToOD(spotinst.Bool(v))
			}
		}

		return strategy, nil
	}

	return nil, nil
}

func flattenStrategy(strategy *azure.Strategy) []interface{} {
	var out []interface{}

	if strategy != nil {
		result := make(map[string]interface{})

		if strategy.SpotPercentage != nil {
			result[string(SpotPercentage)] = spotinst.IntValue(strategy.SpotPercentage)
		}

		if strategy.FallbackToOD != nil {
			result[string(FallbackToOnDemand)] = spotinst.BoolValue(strategy.FallbackToOD)
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}
