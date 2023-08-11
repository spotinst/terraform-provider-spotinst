package stateful_node_azure_strategy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
						Computed: true,
					},
					string(FallbackToOnDemand): {
						Type:     schema.TypeBool,
						Required: true,
					},
					string(RevertToSpot): {
						Type:     schema.TypeList,
						Optional: true,
						Computed: true,
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
					string(CapacityReservation): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(ShouldUtilize): {
									Type:     schema.TypeBool,
									Required: true,
								},
								string(UtilizationStrategy): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(CapacityReservationGroups): {
									Type:     schema.TypeSet,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(CRGName): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(CRGResourceGroupName): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(CRGShouldPrioritize): {
												Type:     schema.TypeBool,
												Optional: true,
											},
										},
									},
								},
							},
						},
					},
					string(OptimizationWindows): {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
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
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
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
				signalsToAdd = flattenSignals(signals)
			}
			if signalsToAdd != nil {
				if err := resourceData.Set(string(Signal), signalsToAdd); err != nil {
					return fmt.Errorf("failed to set signals configuration: %#v", err)
				}
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

			statefulNode.Strategy.SetSignals(signalsToAdd)
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

func flattenSignals(signals []*azure.Signal) []interface{} {
	var result []interface{}

	for _, disk := range signals {
		m := make(map[string]interface{})
		m[string(Type)] = spotinst.StringValue(disk.Type)
		m[string(Timeout)] = spotinst.IntValue(disk.Timeout)
		result = append(result, m)
	}
	return result
}

func flattenStatefulNodeAzureStrategy(strategy *azure.Strategy) []interface{} {
	result := make(map[string]interface{})
	result[string(FallbackToOnDemand)] = spotinst.BoolValue(strategy.FallbackToOnDemand)
	if strategy.PreferredLifecycle != nil {
		result[string(PreferredLifecycle)] = spotinst.StringValue(strategy.PreferredLifecycle)
	}

	if strategy.DrainingTimeout != nil {
		result[string(DrainingTimeout)] = spotinst.IntValue(strategy.DrainingTimeout)
	}

	if strategy.RevertToSpot != nil {
		result[string(RevertToSpot)] = flattenRevertToSpot(strategy.RevertToSpot)
	}

	if strategy.CapacityReservation != nil {
		result[string(CapacityReservation)] = flattenStatefulNodeAzureCapacityReservation(strategy.CapacityReservation)
	}

	if strategy.OptimizationWindows != nil {
		result[string(OptimizationWindows)] = spotinst.StringSlice(strategy.OptimizationWindows)
	}

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

		if v, ok := m[string(DrainingTimeout)].(int); ok && v > 0 {
			strategy.SetDrainingTimeout(spotinst.Int(v))
		}

		if v, ok := m[string(FallbackToOnDemand)].(bool); ok {
			strategy.SetFallbackToOnDemand(spotinst.Bool(v))
		}

		if v, ok := m[string(RevertToSpot)]; ok {
			revertToSpot, err := expandStatefulNodeAzureStrategyRevertToSpot(v)
			if err != nil {
				return nil, err
			}

			if revertToSpot != nil {
				strategy.SetRevertToSpot(revertToSpot)
			}
		}

		if v, ok := m[string(CapacityReservation)]; ok {
			capacityReservation, err := expandStatefulNodeAzureStrategyCapacityReservation(v)
			if err != nil {
				return nil, err
			}

			if capacityReservation != nil {
				strategy.SetCapacityReservation(capacityReservation)
			}
		}

		if v, ok := m[string(OptimizationWindows)]; ok {
			optimizationWindows, err := expandStatefulNodeAzureStrategyOptimizationWindows(v)
			if err != nil {
				return nil, err
			}

			if optimizationWindows != nil {
				strategy.SetOptimizationWindows(optimizationWindows)
			}
		}

	}

	return strategy, nil
}

func expandStatefulNodeAzureStrategySignals(data interface{}) ([]*azure.Signal, error) {
	list := data.([]interface{})
	if list != nil && len(list) > 0 {
		signals := make([]*azure.Signal, 0, len(list))

		for _, item := range list {
			m := item.(map[string]interface{})
			signal := &azure.Signal{}

			if v, ok := m[string(Type)].(string); ok && v != "" {
				signal.SetType(spotinst.String(v))
			} else {
				signal.SetType(nil)
			}

			if v, ok := m[string(Timeout)].(int); ok && v > 0 {
				signal.SetTimeout(spotinst.Int(v))
			} else {
				signal.SetTimeout(nil)
			}
			signals = append(signals, signal)
		}

		return signals, nil
	}

	return nil, nil
}

func expandStatefulNodeAzureStrategyOptimizationWindows(data interface{}) ([]string, error) {
	list := data.([]interface{})
	if list != nil && len(list) > 0 {
		result := make([]string, 0, len(list))

		for _, v := range list {
			if optimizationWindow, ok := v.(string); ok && optimizationWindow != "" {
				result = append(result, optimizationWindow)
			}
		}

		return result, nil
	}

	return nil, nil
}

func expandStatefulNodeAzureStrategyRevertToSpot(data interface{}) (*azure.RevertToSpot, error) {
	list := data.([]interface{})
	if list != nil && len(list) > 0 {
		if list[0] != nil {
			revertToSpot := &azure.RevertToSpot{}
			m := list[0].(map[string]interface{})
			var performAt *string = nil
			if v, ok := m[string(PerformAt)].(string); ok {
				performAt = spotinst.String(v)
			}

			revertToSpot.SetPerformAt(performAt)
			return revertToSpot, nil
		}
	}

	return nil, nil
}

func expandStatefulNodeAzureStrategyCapacityReservation(data interface{}) (*azure.CapacityReservation, error) {
	list := data.(*schema.Set).List()
	capacityReservations := make([]*azure.CapacityReservation, 0, len(list))

	for _, item := range list {
		attr := item.(map[string]interface{})

		capacityReservation := &azure.CapacityReservation{}

		if v, ok := attr[string(CapacityReservationGroups)]; ok {
			capacityReservationGroups, err := expandCapacityReservationGroups(v)
			if err != nil {
				return nil, err
			}

			if capacityReservationGroups != nil {
				capacityReservation.SetCapacityReservationGroups(capacityReservationGroups)
			}
		} else {
			capacityReservation.CapacityReservationGroups = nil
		}

		if v, ok := attr[string(ShouldUtilize)].(bool); ok {
			capacityReservation.SetShouldUtilize(spotinst.Bool(v))
		}

		if v, ok := attr[string(UtilizationStrategy)].(string); ok {
			capacityReservation.SetUtilizationStrategy(spotinst.String(v))
		}

		capacityReservations = append(capacityReservations, capacityReservation)
	}
	return capacityReservations[0], nil
}

func expandCapacityReservationGroups(data interface{}) ([]*azure.CapacityReservationGroup, error) {
	list := data.(*schema.Set).List()
	capacityReservationGroups := make([]*azure.CapacityReservationGroup, 0, len(list))

	for _, item := range list {
		attr := item.(map[string]interface{})

		capacityReservationGroup := &azure.CapacityReservationGroup{}

		if v, ok := attr[string(CRGName)].(string); ok && v != "" {
			capacityReservationGroup.SetCRGName(spotinst.String(v))
		}

		if v, ok := attr[string(CRGResourceGroupName)].(string); ok && v != "" {
			capacityReservationGroup.SetCRGResourceGroupName(spotinst.String(v))
		}

		if v, ok := attr[string(CRGShouldPrioritize)].(bool); ok {
			capacityReservationGroup.SetShouldPrioritize(spotinst.Bool(v))
		}

		capacityReservationGroups = append(capacityReservationGroups, capacityReservationGroup)
	}
	return capacityReservationGroups, nil
}

func flattenStatefulNodeAzureCapacityReservation(capacityReservation *azure.CapacityReservation) []interface{} {
	result := make(map[string]interface{})
	result[string(CapacityReservationGroups)] = flattenCapacityReservationGroups(capacityReservation.CapacityReservationGroups)
	return []interface{}{result}
}

func flattenCapacityReservationGroups(capacityReservationGroups []*azure.CapacityReservationGroup) []interface{} {
	result := make([]interface{}, 0, len(capacityReservationGroups))

	for _, capacityReservationGroup := range capacityReservationGroups {
		m := make(map[string]interface{})
		m[string(CRGName)] = spotinst.StringValue(capacityReservationGroup.Name)
		m[string(CRGResourceGroupName)] = spotinst.StringValue(capacityReservationGroup.ResourceGroupName)
		m[string(CRGShouldPrioritize)] = spotinst.BoolValue(capacityReservationGroup.ShouldPrioritize)
		result = append(result, m)
	}

	return result
}
