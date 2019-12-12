package elastigroup_azure_strategy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
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
					string(LowPriorityPercentage): {
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
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
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
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
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
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
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
func flattenAzureGroupStrategy(strategy *azure.Strategy) []interface{} {
	result := make(map[string]interface{})
	result[string(LowPriorityPercentage)] = spotinst.IntValue(strategy.LowPriorityPercentage)
	result[string(OnDemandCount)] = spotinst.IntValue(strategy.OnDemandCount)
	result[string(DrainingTimeout)] = spotinst.IntValue(strategy.DrainingTimeout)
	return []interface{}{result}
}

func expandAzureGroupStrategy(data interface{}) (*azure.Strategy, error) {
	strategy := &azure.Strategy{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(LowPriorityPercentage)].(int); ok && v >= 0 {
			strategy.SetLowPriorityPercentage(spotinst.Int(v))
		}

		if v, ok := m[string(OnDemandCount)].(int); ok && v >= 0 {
			strategy.SetOnDemandCount(spotinst.Int(v))
		}

		if v, ok := m[string(DrainingTimeout)].(int); ok && v >= 0 {
			strategy.SetDrainingTimeout(spotinst.Int(v))
		}
	}
	return strategy, nil
}
