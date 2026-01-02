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
					string(InterruptionToleration): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Cooldown): {
									Type:     schema.TypeInt,
									Optional: true,
								},
								string(EvaluationPeriod): {
									Type:     schema.TypeInt,
									Optional: true,
								},
								string(IsEnabled): {
									Type:     schema.TypeBool,
									Optional: true,
								},
								string(Threshold): {
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},
					string(OptimizationWindows): {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					string(OdWindows): {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					string(AvailabilityVsCost): {
						Type:     schema.TypeInt,
						Optional: true,
					},
					string(VmAdmins): {
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
			var value *azure.Strategy = nil
			if v, ok := resourceData.GetOk(string(Strategy)); ok {
				if strategy, err := expandStatefulNodeAzureStrategy(v); err != nil {
					return err
				} else {
					value = strategy
				}
				statefulNode.SetStrategy(value)
			} else {
				statefulNode.SetStrategy(nil)
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

	if strategy.AvailabilityVsCost != nil {
		result[string(AvailabilityVsCost)] = spotinst.IntValue(strategy.AvailabilityVsCost)
	}

	if strategy.RevertToSpot != nil {
		result[string(RevertToSpot)] = flattenRevertToSpot(strategy.RevertToSpot)
	}

	if strategy.CapacityReservation != nil {
		result[string(CapacityReservation)] = flattenStatefulNodeAzureCapacityReservation(strategy.CapacityReservation)
	}

	if strategy.InterruptionToleration != nil {
		result[string(InterruptionToleration)] = flattenInterruptionToleration(strategy.InterruptionToleration)
	}

	if strategy.OptimizationWindows != nil {
		result[string(OptimizationWindows)] = spotinst.StringSlice(strategy.OptimizationWindows)
	}

	if strategy.OdWindows != nil {
		result[string(OdWindows)] = spotinst.StringSlice(strategy.OdWindows)
	}

	if strategy.VmAdmins != nil {
		result[string(VmAdmins)] = spotinst.StringSlice(strategy.VmAdmins)
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

		if v, ok := m[string(AvailabilityVsCost)].(int); ok && v > 0 {
			strategy.SetAvailabilityVsCost(spotinst.Int(v))
		} else {
			strategy.SetAvailabilityVsCost(nil)
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
			capacityReservation, err := expandStrategyCapacityReservation(v)
			if err != nil {
				return nil, err
			}

			strategy.SetCapacityReservation(capacityReservation)

		}

		if v, ok := m[string(InterruptionToleration)]; ok && v != nil {
			if interruptionToleration, err := expandInterruptionToleration(v); err != nil {
				return nil, err
			} else {
				if interruptionToleration != nil {
					strategy.SetInterruptionToleration(interruptionToleration)
				} else {
					strategy.SetInterruptionToleration(nil)
				}

			}
		} else {
			strategy.SetInterruptionToleration(nil)
		}

		if v, ok := m[string(OptimizationWindows)]; ok {
			optimizationWindows, err := expandStatefulNodeAzureStrategyList(v)
			if err != nil {
				return nil, err
			}

			if optimizationWindows != nil && len(optimizationWindows) > 0 {
				strategy.SetOptimizationWindows(optimizationWindows)
			} else {
				strategy.SetOptimizationWindows(nil)
			}
		}

		if v, ok := m[string(OdWindows)]; ok {
			odWindows, err := expandStatefulNodeAzureStrategyList(v)
			if err != nil {
				return nil, err
			}

			if odWindows != nil && len(odWindows) > 0 {
				strategy.SetOdWindows(odWindows)
			} else {
				strategy.SetOdWindows(nil)
			}
		}

		if v, ok := m[string(VmAdmins)]; ok {
			vmAdmins, err := expandStatefulNodeAzureStrategyList(v)
			if err != nil {
				return nil, err
			}

			if vmAdmins != nil && len(vmAdmins) > 0 {
				strategy.SetVmAdmins(vmAdmins)
			} else {
				strategy.SetVmAdmins(nil)
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

func expandStatefulNodeAzureStrategyList(data interface{}) ([]string, error) {
	list := data.([]interface{})
	if list != nil && len(list) > 0 {
		result := make([]string, 0, len(list))

		for _, v := range list {
			if value, ok := v.(string); ok && len(value) > 0 {
				result = append(result, value)
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

func expandStrategyCapacityReservation(data interface{}) (*azure.CapacityReservation, error) {

	list := data.(*schema.Set).List()
	capacityReservation := &azure.CapacityReservation{}

	if len(list) > 0 {
		item := list[0]
		attr := item.(map[string]interface{})
		if v, ok := attr[string(CapacityReservationGroups)]; ok {
			capacityReservationGroups, err := expandCapacityReservationGroups(v)
			if err != nil {
				return nil, err
			}
			capacityReservation.SetCapacityReservationGroups(capacityReservationGroups)
		} else {
			capacityReservation.CapacityReservationGroups = nil
		}

		if v, ok := attr[string(ShouldUtilize)].(bool); ok {
			capacityReservation.SetShouldUtilize(spotinst.Bool(v))
		}

		if v, ok := attr[string(UtilizationStrategy)].(string); ok {
			capacityReservation.SetUtilizationStrategy(spotinst.String(v))
		}

		return capacityReservation, nil
	}

	return nil, nil
}

func expandCapacityReservationGroups(data interface{}) ([]*azure.CapacityReservationGroup, error) {
	list := data.(*schema.Set).List()
	capacityReservationGroups := make([]*azure.CapacityReservationGroup, 0, len(list))

	if len(list) > 0 {
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
	return nil, nil
}

func expandInterruptionToleration(data interface{}) (*azure.InterruptionToleration, error) {
	if list := data.([]interface{}); len(list) > 0 {
		interruptionToleration := &azure.InterruptionToleration{}
		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(IsEnabled)]; ok && v != nil {
				if boolVal, isBool := v.(bool); isBool {
					interruptionToleration.SetIsEnabled(spotinst.Bool(boolVal))
				}
			} else {
				interruptionToleration.SetIsEnabled(nil)
			}

			if v, ok := m[string(Cooldown)].(int); ok && v > 0 {
				interruptionToleration.SetCooldown(spotinst.Int(v))
			} else {
				interruptionToleration.SetCooldown(nil)
			}

			if v, ok := m[string(EvaluationPeriod)].(int); ok && v > 0 {
				interruptionToleration.SetEvaluationPeriod(spotinst.Int(v))
			} else {
				interruptionToleration.SetEvaluationPeriod(nil)
			}

			if v, ok := m[string(Threshold)].(int); ok && v > 0 {
				interruptionToleration.SetThreshold(spotinst.Int(v))
			} else {
				interruptionToleration.SetThreshold(nil)
			}
		}
		return interruptionToleration, nil
	}
	return nil, nil
}

func flattenStatefulNodeAzureCapacityReservation(capacityReservation *azure.CapacityReservation) []interface{} {
	result := make(map[string]interface{})
	result[string(ShouldUtilize)] = spotinst.BoolValue(capacityReservation.ShouldUtilize)
	result[string(UtilizationStrategy)] = spotinst.StringValue(capacityReservation.UtilizationStrategy)
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

func flattenInterruptionToleration(interruptionToleration *azure.InterruptionToleration) []interface{} {

	var out []interface{}

	if interruptionToleration != nil {
		interTol := make(map[string]interface{})

		if interruptionToleration.IsEnabled != nil {
			interTol[string(IsEnabled)] = spotinst.BoolValue(interruptionToleration.IsEnabled)
		}

		if interruptionToleration.Cooldown != nil {
			interTol[string(Cooldown)] = spotinst.IntValue(interruptionToleration.Cooldown)
		}

		if interruptionToleration.EvaluationPeriod != nil {
			interTol[string(EvaluationPeriod)] = spotinst.IntValue(interruptionToleration.EvaluationPeriod)
		}

		if interruptionToleration.Threshold != nil {
			interTol[string(Threshold)] = spotinst.IntValue(interruptionToleration.Threshold)
		}

		if len(interTol) > 0 {
			out = append(out, interTol)
		}

		return []interface{}{interTol}
	}
	return out
}
