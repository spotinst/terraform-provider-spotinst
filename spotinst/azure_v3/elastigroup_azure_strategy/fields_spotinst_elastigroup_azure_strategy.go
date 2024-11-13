package elastigroup_azure_strategy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[SpotPercentage] = commons.NewGenericField(
		commons.ElastigroupAzureStrategy,
		SpotPercentage,
		&schema.Schema{
			Type:          schema.TypeInt,
			Optional:      true,
			ConflictsWith: []string{string(OnDemandCount)},
			//Force setting -1 as default value if it's not exists in initial creation,
			// to allow initialization of the field to 0
			Default:      -1,
			ValidateFunc: validation.IntAtLeast(-1),
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == "-1" && new == "null" {
					return true
				}
				return false
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.SpotPercentage != nil {
				value = elastigroup.Strategy.SpotPercentage
			}
			if value != nil {
				if err := resourceData.Set(string(SpotPercentage), spotinst.Int(int(*value))); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SpotPercentage), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(SpotPercentage)).(int); ok && v > -1 {
				elastigroup.Strategy.SetSpotPercentage(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(SpotPercentage)).(int); ok && v > -1 {
				elastigroup.Strategy.SetSpotPercentage(spotinst.Int(v))
			} else {
				elastigroup.Strategy.SetSpotPercentage(nil)
			}

			return nil
		},
		nil,
	)

	fieldsMap[OnDemandCount] = commons.NewGenericField(
		commons.ElastigroupAzureStrategy,
		OnDemandCount,
		&schema.Schema{
			Type:          schema.TypeInt,
			Optional:      true,
			ConflictsWith: []string{string(SpotPercentage)},
			//Force setting -1 as default value if it's not exists in initial creation,
			// to allow initialization of the field to 0
			Default:      -1,
			ValidateFunc: validation.IntAtLeast(-1),
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == "-1" && new == "null" {
					return true
				}
				return false
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.OnDemandCount != nil {
				value = elastigroup.Strategy.OnDemandCount
			}
			if value != nil {
				if err := resourceData.Set(string(OnDemandCount), spotinst.IntValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OnDemandCount), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(OnDemandCount)).(int); ok && v > -1 {
				elastigroup.Strategy.SetOnDemandCount(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(OnDemandCount)).(int); ok && v > -1 {
				elastigroup.Strategy.SetOnDemandCount(spotinst.Int(v))
			} else {
				elastigroup.Strategy.SetOnDemandCount(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[DrainingTimeout] = commons.NewGenericField(
		commons.ElastigroupAzureStrategy,
		DrainingTimeout,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.DrainingTimeout != nil {
				value = elastigroup.Strategy.DrainingTimeout
			}
			if err := resourceData.Set(string(DrainingTimeout), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DrainingTimeout), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(DrainingTimeout)).(int); ok && v > 0 {
				elastigroup.Strategy.SetDrainingTimeout(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(DrainingTimeout)).(int); ok && v > 0 {
				elastigroup.Strategy.SetDrainingTimeout(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[FallbackToOnDemand] = commons.NewGenericField(
		commons.ElastigroupAzureStrategy,
		FallbackToOnDemand,
		&schema.Schema{
			Type:     schema.TypeBool,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *bool = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.FallbackToOnDemand != nil {
				value = elastigroup.Strategy.FallbackToOnDemand
			}
			if err := resourceData.Set(string(FallbackToOnDemand), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(FallbackToOnDemand), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(FallbackToOnDemand)).(bool); ok {
				elastigroup.Strategy.SetFallbackToOnDemand(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(FallbackToOnDemand)).(bool); ok {
				elastigroup.Strategy.SetFallbackToOnDemand(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[AvailabilityVsCost] = commons.NewGenericField(
		commons.ElastigroupAzureStrategy,
		AvailabilityVsCost,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			//Force setting -1 as default value if it's not exists in initial creation,
			// to allow initialization of the field to 0
			Default:      -1,
			ValidateFunc: validation.IntAtLeast(-1),
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == "-1" && new == "null" {
					return true
				}
				return false
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.AvailabilityVsCost != nil {
				value = elastigroup.Strategy.AvailabilityVsCost
			}
			if value != nil {
				if err := resourceData.Set(string(AvailabilityVsCost), spotinst.Int(int(*value))); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AvailabilityVsCost), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(AvailabilityVsCost)).(int); ok && v > -1 {
				elastigroup.Strategy.SetAvailabilityVsCost(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(AvailabilityVsCost)).(int); ok && v > -1 {
				elastigroup.Strategy.SetAvailabilityVsCost(spotinst.Int(v))
			} else {
				elastigroup.Strategy.SetAvailabilityVsCost(nil)
			}

			return nil
		},
		nil,
	)

	fieldsMap[RevertToSpot] = commons.NewGenericField(
		commons.ElastigroupAzureStrategy,
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
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{} = nil

			if elastigroup.Strategy != nil && elastigroup.Strategy.RevertToSpot != nil {
				value = flattenRevertToSpot(elastigroup.Strategy.RevertToSpot)
			}
			if err := resourceData.Set(string(RevertToSpot), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RevertToSpot), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(RevertToSpot)); ok {
				if revertToSpot, err := expandRevertToSpot(v); err != nil {
					return err
				} else {
					elastigroup.Strategy.SetRevertToSpot(revertToSpot)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *azurev3.RevertToSpot = nil
			if v, ok := resourceData.GetOk(string(RevertToSpot)); ok {
				if revertToSpot, err := expandRevertToSpot(v); err != nil {
					return err
				} else {
					value = revertToSpot
				}
				elastigroup.Strategy.SetRevertToSpot(value)
			} else {
				elastigroup.Strategy.SetRevertToSpot(nil)
			}

			return nil
		},
		nil,
	)

	fieldsMap[CapacityReservation] = commons.NewGenericField(
		commons.ElastigroupAzureStrategy,
		CapacityReservation,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
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
						MaxItems: 1,
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
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{} = nil

			if elastigroup.Strategy != nil && elastigroup.Strategy.CapacityReservation != nil {
				value = flattenCapacityReservation(elastigroup.Strategy.CapacityReservation)
			}

			if len(value) > 0 {
				if err := resourceData.Set(string(CapacityReservation), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CapacityReservation), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(CapacityReservation)); ok {
				if capacityReservation, err := expandCapacityReservation(v); err != nil {
					return err
				} else {
					elastigroup.Strategy.SetCapacityReservation(capacityReservation)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *azurev3.CapacityReservation = nil
			if v, ok := resourceData.GetOk(string(CapacityReservation)); ok {
				if capacityReservation, err := expandCapacityReservation(v); err != nil {
					return err
				} else {
					value = capacityReservation
				}
				elastigroup.Strategy.SetCapacityReservation(value)
			} else {
				elastigroup.Strategy.SetCapacityReservation(nil)
			}

			return nil
		},
		nil,
	)

	fieldsMap[Signal] = commons.NewGenericField(
		commons.ElastigroupAzureStrategy,
		Signal,
		&schema.Schema{
			Type:     schema.TypeList,
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
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var signalsToAdd []interface{}

			if elastigroup.Strategy != nil && elastigroup.Strategy.Signals != nil {
				signalsToAdd = flattenSignals(elastigroup.Strategy.Signals)
			}
			if len(signalsToAdd) > 0 {
				if err := resourceData.Set(string(Signal), signalsToAdd); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Signal), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Signal)); ok {
				if signals, err := expandSignals(v); err != nil {
					return err
				} else {
					elastigroup.Strategy.SetSignals(signals)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var signalsToAdd []*azurev3.Signals = nil

			if v, ok := resourceData.GetOk(string(Signal)); ok {
				if signals, err := expandSignals(v); err != nil {
					return err
				} else {
					signalsToAdd = signals
				}
				elastigroup.Strategy.SetSignals(signalsToAdd)
			} else {
				elastigroup.Strategy.SetSignals(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[OptimizationWindows] = commons.NewGenericField(
		commons.ElastigroupAzureStrategy,
		OptimizationWindows,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []string = nil
			if elastigroup.Compute != nil && elastigroup.Strategy.OptimizationWindows != nil {
				result = elastigroup.Strategy.OptimizationWindows
			}
			if err := resourceData.Set(string(OptimizationWindows), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OptimizationWindows), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(OptimizationWindows)).([]interface{}); ok && v != nil {
				if ow, err := expandStrategyList(v); err != nil {
					return err
				} else {
					elastigroup.Strategy.SetOptimizationWindows(ow)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(OptimizationWindows)); ok {
				if ow, err := expandStrategyList(v); err != nil {
					return err
				} else {
					elastigroup.Strategy.SetOptimizationWindows(ow)
				}
			} else {
				elastigroup.Strategy.SetOptimizationWindows(nil)
			}
			return nil
		},
		nil,
	)
}

func flattenRevertToSpot(revertToSpot *azurev3.RevertToSpot) []interface{} {
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

func flattenSignals(signals []*azurev3.Signals) []interface{} {
	var result []interface{}

	for _, disk := range signals {
		m := make(map[string]interface{})
		m[string(Type)] = spotinst.StringValue(disk.Type)
		m[string(Timeout)] = spotinst.IntValue(disk.Timeout)
		result = append(result, m)
	}
	return result
}

func expandSignals(data interface{}) ([]*azurev3.Signals, error) {
	list := data.([]interface{})
	if list != nil && len(list) > 0 {
		signals := make([]*azurev3.Signals, 0, len(list))

		for _, item := range list {
			m, ok := item.(map[string]interface{})
			if !ok {
				continue
			}

			signal := &azurev3.Signals{}
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

func expandStrategyList(data interface{}) ([]string, error) {
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

func expandRevertToSpot(data interface{}) (*azurev3.RevertToSpot, error) {
	list := data.([]interface{})
	if list != nil && len(list) > 0 {
		if list[0] != nil {
			revertToSpot := &azurev3.RevertToSpot{}
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

func expandCapacityReservation(data interface{}) (*azurev3.CapacityReservation, error) {
	list := data.([]interface{})
	capacityReservation := &azurev3.CapacityReservation{}

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

func expandCapacityReservationGroups(data interface{}) ([]*azurev3.CapacityReservationGroups, error) {
	list := data.(*schema.Set).List()
	capacityReservationGroups := make([]*azurev3.CapacityReservationGroups, 0, len(list))

	if len(list) > 0 {
		for _, item := range list {
			attr := item.(map[string]interface{})

			capacityReservationGroup := &azurev3.CapacityReservationGroups{}

			if v, ok := attr[string(CRGName)].(string); ok && v != "" {
				capacityReservationGroup.SetName(spotinst.String(v))
			}
			if v, ok := attr[string(CRGResourceGroupName)].(string); ok && v != "" {
				capacityReservationGroup.SetResourceGroupName(spotinst.String(v))
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

func flattenCapacityReservation(capacityReservation *azurev3.CapacityReservation) []interface{} {
	result := make(map[string]interface{})
	result[string(ShouldUtilize)] = spotinst.BoolValue(capacityReservation.ShouldUtilize)
	result[string(UtilizationStrategy)] = spotinst.StringValue(capacityReservation.UtilizationStrategy)
	result[string(CapacityReservationGroups)] = flattenCapacityReservationGroups(capacityReservation.CapacityReservationGroups)
	return []interface{}{result}
}

func flattenCapacityReservationGroups(capacityReservationGroups []*azurev3.CapacityReservationGroups) []interface{} {
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
