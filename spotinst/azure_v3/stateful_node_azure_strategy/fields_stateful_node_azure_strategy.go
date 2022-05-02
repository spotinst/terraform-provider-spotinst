package stateful_node_azure_strategy

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Strategy] = commons.NewGenericField(
		commons.StatefulNodeAzureStrategy,
		Strategy,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(PreferredLifecycle): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					string(DrainingTimeout): {
						Type:     schema.TypeInt,
						Optional: true,
					},
					string(FallbackToOnDemand): {
						Type:     schema.TypeBool,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value []interface{} = nil

			if statefulNode.Strategy != nil {
				value = flattenStatefulNodeAzureStrategy(statefulNode.Strategy)
			}
			if err := resourceData.Set(string(Strategy), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Strategy), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(Strategy)); ok {
				if strategy, err := expandStatefulNodeAzureStrategy(v); err != nil {
					return err
				} else {
					statefulNode.SetStrategy(strategy)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(Strategy)); ok {
				if strategy, err := expandStatefulNodeAzureStrategy(v); err != nil {
					return err
				} else {
					statefulNode.SetStrategy(strategy)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[Signal] = commons.NewGenericField(
		commons.StatefulNodeAzureStrategy,
		Signal,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Type): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(Timeout): {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var signalsToAdd []interface{} = nil
			if statefulNode.Strategy != nil && statefulNode.Strategy.Signals != nil {
				signals := statefulNode.Strategy.Signals
				signalsToAdd = make([]interface{}, 0, len(signals))
				for _, s := range signals {
					m := make(map[string]interface{})
					m[string(Type)] = spotinst.StringValue(s.Type)
					m[string(Timeout)] = spotinst.IntValue(s.Timeout)
					signalsToAdd = append(signalsToAdd, m)
				}
			}
			if err := resourceData.Set(string(Signal), signalsToAdd); err != nil {
				return fmt.Errorf("failed to set signals configuration: %#v", err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(Signal)); ok {
				if signals, err := expandStatefulNodeAzureStrategySignals(v); err != nil {
					return err
				} else {
					statefulNode.Strategy.SetSignals(signals)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var signalsToAdd []*azure.Signal = nil
			if v, ok := resourceData.GetOk(string(Signal)); ok {
				if signals, err := expandStatefulNodeAzureStrategySignals(v); err != nil {
					return err
				} else {
					signalsToAdd = signals
				}
			}
			if statefulNode.Strategy == nil {
				statefulNode.SetStrategy(&azure.Strategy{})
			}
			statefulNode.Strategy.SetSignals(signalsToAdd)
			return nil
		},
		nil,
	)

	fieldsMap[OptimizationWindows] = commons.NewGenericField(
		commons.StatefulNodeAzureStrategy,
		OptimizationWindows,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value []string = nil
			if statefulNode.Strategy != nil && statefulNode.Strategy.OptimizationWindows != nil {
				value = statefulNode.Strategy.OptimizationWindows
			}
			if err := resourceData.Set(string(OptimizationWindows), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OptimizationWindows), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if value, ok := resourceData.GetOk(string(OptimizationWindows)); ok && value != nil {
				if subnetIds, err := expandStatefulNodeAzureStrategyOptimizationWindows(value); err != nil {
					return err
				} else {
					statefulNode.Strategy.SetOptimizationWindows(subnetIds)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if value, ok := resourceData.GetOk(string(OptimizationWindows)); ok && value != nil {
				if subnetIds, err := expandStatefulNodeAzureStrategyOptimizationWindows(value); err != nil {
					return err
				} else {
					statefulNode.Strategy.SetOptimizationWindows(subnetIds)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[RevertToSpot] = commons.NewGenericField(
		commons.StatefulNodeAzureStrategy,
		RevertToSpot,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(PerformAt): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var result []interface{} = nil
			if statefulNode.Strategy != nil && statefulNode.Strategy.RevertToSpot != nil {
				rts := statefulNode.Strategy.RevertToSpot
				result = flattenRevertToSpot(rts)
			}
			if result != nil {
				if err := resourceData.Set(string(RevertToSpot), result); err != nil {
					return fmt.Errorf("failed to set revertToSpot configuration: %#v", err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(RevertToSpot)); ok {
				if revertToSpot, err := expandStatefulNodeAzureStrategyRevertToSpot(v); err != nil {
					return err
				} else {
					statefulNode.Strategy.SetRevertToSpot(revertToSpot)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var revertToSpot *azure.RevertToSpot = nil
			if v, ok := resourceData.GetOk(string(RevertToSpot)); ok {
				if rts, err := expandStatefulNodeAzureStrategyRevertToSpot(v); err != nil {
					return err
				} else {
					revertToSpot = rts
				}
			}
			statefulNode.Strategy.SetRevertToSpot(revertToSpot)
			return nil
		},
		nil,
	)
}

func flattenRevertToSpot(revertToSpot *azure.RevertToSpot) []interface{} {
	var out []interface{}

	if revertToSpot != nil {
		result := make(map[string]interface{})

		if revertToSpot.PerformAt != nil {
			result[string(PerformAt)] = spotinst.StringValue(revertToSpot.PerformAt)
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

func flattenStatefulNodeAzureStrategy(strategy *azure.Strategy) []interface{} {
	result := make(map[string]interface{})

	result[string(PreferredLifecycle)] = spotinst.StringValue(strategy.PreferredLifecycle)
	result[string(DrainingTimeout)] = spotinst.IntValue(strategy.DrainingTimeout)
	result[string(FallbackToOnDemand)] = spotinst.BoolValue(strategy.FallbackToOnDemand)

	return []interface{}{result}
}

func expandStatefulNodeAzureStrategy(data interface{}) (*azure.Strategy, error) {
	strategy := &azure.Strategy{}
	list := data.([]interface{})

	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})
		if v, ok := m[string(PreferredLifecycle)].(string); ok && v != "" {
			strategy.SetPreferredLifecycle(spotinst.String(v))
		}
		if v, ok := m[string(DrainingTimeout)].(int); ok && v >= 0 {
			strategy.SetDrainingTimeout(spotinst.Int(v))
		}
		if v, ok := m[string(FallbackToOnDemand)].(bool); ok {
			strategy.SetFallbackToOnDemand(spotinst.Bool(v))
		}
	}

	return strategy, nil
}

func expandStatefulNodeAzureStrategySignals(data interface{}) ([]*azure.Signal, error) {
	list := data.(*schema.Set).List()
	signals := make([]*azure.Signal, 0, len(list))

	for _, item := range list {
		m := item.(map[string]interface{})
		signal := &azure.Signal{}

		if v, ok := m[string(Type)].(string); ok && v != "" {
			signal.SetType(spotinst.String(v))
		}

		if v, ok := m[string(Timeout)].(int); ok && v > 0 {
			signal.SetTimeout(spotinst.Int(v))
		}
		signals = append(signals, signal)
	}

	return signals, nil
}

func expandStatefulNodeAzureStrategyOptimizationWindows(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if optimizationWindow, ok := v.(string); ok && optimizationWindow != "" {
			result = append(result, optimizationWindow)
		}
	}

	return result, nil
}

func expandStatefulNodeAzureStrategyRevertToSpot(data interface{}) (*azure.RevertToSpot, error) {
	if list := data.([]interface{}); len(list) > 0 {
		revertToSpot := &azure.RevertToSpot{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(PerformAt)].(string); ok && v != "" {
				revertToSpot.SetPerformAt(spotinst.String(v))
			} else {
				revertToSpot.SetPerformAt(nil)
			}
		}
		return revertToSpot, nil
	}
	return nil, nil
}
