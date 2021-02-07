package elastigroup_azure_strategy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Strategy] = commons.NewGenericField(
		commons.ElastigroupAzureStrategy,
		Strategy,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(SpotPercentage): {
						Type:     schema.TypeInt,
						Optional: true,
					},
					string(OnDemandCount): {
						Type:     schema.TypeInt,
						Optional: true,
					},
					string(DrainingTimeout): {
						Type:     schema.TypeInt,
						Optional: true,
					},
					string(FallbackToOnDemand): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					//string(RevertToSpot): {
					//	Type:     schema.TypeList,
					//	Optional: true,
					//	MaxItems: 1,
					//	Elem: &schema.Resource{
					//		Schema: map[string]*schema.Schema{
					//			string(PerformAt): {
					//				Type:     schema.TypeString,
					//				Required: true,
					//			},
					//		},
					//	},
					//},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{} = nil

			if elastigroup.Strategy != nil {
				value = flattenAzureGroupStrategy(elastigroup.Strategy)
			}
			if err := resourceData.Set(string(Strategy), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Strategy), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Strategy)); ok {
				if strategy, err := expandAzureGroupStrategy(v); err != nil {
					return err
				} else {
					elastigroup.SetStrategy(strategy)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Strategy)); ok {
				if strategy, err := expandAzureGroupStrategy(v); err != nil {
					return err
				} else {
					elastigroup.SetStrategy(strategy)
				}
			}
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func flattenAzureGroupStrategy(strategy *v3.Strategy) []interface{} {
	result := make(map[string]interface{})

	result[string(SpotPercentage)] = spotinst.IntValue(strategy.SpotPercentage)
	result[string(OnDemandCount)] = spotinst.IntValue(strategy.OnDemandCount)
	result[string(DrainingTimeout)] = spotinst.IntValue(strategy.DrainingTimeout)
	result[string(FallbackToOnDemand)] = spotinst.BoolValue(strategy.FallbackToOnDemand)
	//if strategy.RevertToSpot != nil {
	//	result[string(RevertToSpot)] = flattenAzureGroupRevertToSpot(strategy.RevertToSpot)
	//}

	return []interface{}{result}
}

//func flattenAzureGroupRevertToSpot(revertToSpot *v3.RevertToSpot) []interface{} {
//	result := make(map[string]interface{})
//	result[string(PerformAt)] = spotinst.StringValue(revertToSpot.PerformAt)
//	return []interface{}{result}
//}

func expandAzureGroupStrategy(data interface{}) (*v3.Strategy, error) {
	strategy := &v3.Strategy{}
	list := data.([]interface{})

	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})
		if v, ok := m[string(SpotPercentage)].(int); ok && v > 0 {
			strategy.SetSpotPercentage(spotinst.Int(v))
		}
		if v, ok := m[string(OnDemandCount)].(int); ok && v > 0 {
			strategy.SetOnDemandCount(spotinst.Int(v))
		}
		if v, ok := m[string(DrainingTimeout)].(int); ok && v >= 0 {
			strategy.SetDrainingTimeout(spotinst.Int(v))
		}
		if v, ok := m[string(FallbackToOnDemand)].(bool); ok {
			strategy.SetFallbackToOnDemand(spotinst.Bool(v))
		}
		//if v, ok := m[string(RevertToSpot)]; ok {
		//	if revertToSpot, err := expandAzureGroupRevertToSpot(v); err != nil {
		//		return nil, err
		//	} else {
		//		strategy.SetRevertToSpot(revertToSpot)
		//	}
		//}
	}

	return strategy, nil
}

//func expandAzureGroupRevertToSpot(data interface{}) (*v3.RevertToSpot, error) {
//	if list := data.([]interface{}); len(list) > 0 {
//		revertToSpot := &v3.RevertToSpot{}
//		if list != nil && list[0] != nil {
//			m := list[0].(map[string]interface{})
//			if v, ok := m[string(PerformAt)].(string); ok && v != "" {
//				revertToSpot.SetPerformAt(spotinst.String(v))
//			}
//		}
//		return revertToSpot, nil
//	}
//	return nil, nil
//}
