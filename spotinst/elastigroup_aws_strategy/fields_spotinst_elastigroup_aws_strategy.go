package elastigroup_aws_strategy

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[SpotPercentage] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		SpotPercentage,
		&schema.Schema{
			Type:          schema.TypeInt,
			Optional:      true,
			ConflictsWith: []string{string(OnDemandCount)},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			//Risk is configured as int in the API but float on Go-SDK (currently not aligning because of breaking changes effects)
			//There for value is of type float and cast as necessary
			var value *float64 = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.Risk != nil {
				value = elastigroup.Strategy.Risk
			}
			if value != nil {
				if err := resourceData.Set(string(SpotPercentage), spotinst.Int(int(*value))); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SpotPercentage), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(SpotPercentage)).(int); ok && v >= 0 {
				elastigroup.Strategy.SetRisk(spotinst.Float64(float64(v)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(SpotPercentage)).(int); ok && v >= 0 {
				elastigroup.Strategy.SetRisk(spotinst.Float64(float64(v)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[OnDemandCount] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		OnDemandCount,
		&schema.Schema{
			Type:          schema.TypeInt,
			Optional:      true,
			ConflictsWith: []string{string(SpotPercentage)},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(OnDemandCount)); ok && v != nil {
				value := v.(int)
				elastigroup.Strategy.SetOnDemandCount(spotinst.Int(value))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var count *int
			if v, ok := resourceData.GetOkExists(string(OnDemandCount)); ok && v != nil {
				if value, ok := v.(int); ok && value > 0 {
					count = spotinst.Int(value)
				}
			}
			elastigroup.Strategy.SetOnDemandCount(count)
			return nil
		},
		nil,
	)

	fieldsMap[Orientation] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		Orientation,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.AvailabilityVsCost != nil {
				value = elastigroup.Strategy.AvailabilityVsCost
			}
			if err := resourceData.Set(string(Orientation), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Orientation), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(Orientation)).(string); ok && v != "" {
				elastigroup.Strategy.SetAvailabilityVsCost(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(Orientation)).(string); ok && v != "" {
				elastigroup.Strategy.SetAvailabilityVsCost(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[LifetimePeriod] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		LifetimePeriod,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.LifetimePeriod != nil {
				value = elastigroup.Strategy.LifetimePeriod
			}
			if err := resourceData.Set(string(LifetimePeriod), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(LifetimePeriod), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(LifetimePeriod)).(string); ok && v != "" {
				period := spotinst.String(v)
				elastigroup.Strategy.SetLifetimePeriod(period)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var period *string = nil
			if v, ok := resourceData.Get(string(LifetimePeriod)).(string); ok && v != "" {
				period = spotinst.String(v)
			}
			elastigroup.Strategy.SetLifetimePeriod(period)
			return nil
		},
		nil,
	)

	fieldsMap[DrainingTimeout] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		DrainingTimeout,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(DrainingTimeout)).(int); ok && v > 0 {
				elastigroup.Strategy.SetDrainingTimeout(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(DrainingTimeout)).(int); ok && v > 0 {
				elastigroup.Strategy.SetDrainingTimeout(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[UtilizeReservedInstances] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		UtilizeReservedInstances,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *bool = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.UtilizeReservedInstances != nil {
				value = elastigroup.Strategy.UtilizeReservedInstances
			}
			if err := resourceData.Set(string(UtilizeReservedInstances), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UtilizeReservedInstances), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(UtilizeReservedInstances)).(bool); ok && v {
				ris := spotinst.Bool(v)
				elastigroup.Strategy.SetUtilizeReservedInstances(ris)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var ris *bool = nil
			if v, ok := resourceData.Get(string(UtilizeReservedInstances)).(bool); ok && v {
				ris = spotinst.Bool(v)
			}
			elastigroup.Strategy.SetUtilizeReservedInstances(ris)
			return nil
		},
		nil,
	)

	fieldsMap[FallbackToOnDemand] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		FallbackToOnDemand,
		&schema.Schema{
			Type:     schema.TypeBool,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(FallbackToOnDemand)); ok && v != nil {
				ftod := v.(bool)
				fallback := spotinst.Bool(ftod)
				elastigroup.Strategy.SetFallbackToOnDemand(fallback)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var fallback *bool = nil
			if v, ok := resourceData.GetOkExists(string(FallbackToOnDemand)); ok && v != nil {
				ftod := v.(bool)
				fallback = spotinst.Bool(ftod)
			}
			elastigroup.Strategy.SetFallbackToOnDemand(fallback)
			return nil
		},
		nil,
	)

	fieldsMap[ScalingStrategy] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		ScalingStrategy,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(TerminateAtEndOfBillingHour): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(TerminationPolicy): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{} = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.ScalingStrategy != nil {
				s := elastigroup.Strategy.ScalingStrategy
				value = flattenAWSGroupScalingStrategy(s)
			}
			if value != nil {
				if err := resourceData.Set(string(ScalingStrategy), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ScalingStrategy), err)
				}
			} else {
				if err := resourceData.Set(string(ScalingStrategy), []*aws.ScalingStrategy{}); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ScalingStrategy), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(ScalingStrategy)); ok {
				if s, err := expandAWSGroupScalingStrategy(v); err != nil {
					return err
				} else {
					elastigroup.Strategy.SetScalingStrategy(s)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.ScalingStrategy = nil
			if v, ok := resourceData.GetOk(string(ScalingStrategy)); ok {
				if s, err := expandAWSGroupScalingStrategy(v); err != nil {
					return err
				} else {
					value = s
				}
			}
			elastigroup.Strategy.SetScalingStrategy(value)
			return nil
		},
		nil,
	)

	fieldsMap[UtilizeCommitments] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		UtilizeCommitments,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *bool = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.UtilizeCommitments != nil {
				value = elastigroup.Strategy.UtilizeCommitments
			}
			if err := resourceData.Set(string(UtilizeCommitments), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UtilizeCommitments), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(UtilizeCommitments)); ok && v != nil {
				uc := v.(bool)
				utilizeCommitments := spotinst.Bool(uc)
				elastigroup.Strategy.SetUtilizeCommitments(utilizeCommitments)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var utilizeCommitments *bool = nil
			if v, ok := resourceData.GetOkExists(string(UtilizeCommitments)); ok && v != nil {
				uc := v.(bool)
				utilizeCommitments = spotinst.Bool(uc)
			}
			elastigroup.Strategy.SetUtilizeCommitments(utilizeCommitments)
			return nil
		},
		nil,
	)

	fieldsMap[MinimumInstanceLifetime] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		MinimumInstanceLifetime,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.MinimumInstanceLifetime != nil {
				value = elastigroup.Strategy.MinimumInstanceLifetime
			}
			if value != nil {
				if err := resourceData.Set(string(MinimumInstanceLifetime), spotinst.IntValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MinimumInstanceLifetime), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(MinimumInstanceLifetime)); ok && v != nil {
				value := v.(int)
				elastigroup.Strategy.SetMinimumInstanceLifetime(spotinst.Int(value))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var minimumInstanceLifetime *int
			if v, ok := resourceData.GetOkExists(string(MinimumInstanceLifetime)); ok && v != nil {
				if value, ok := v.(int); ok && value > 0 {
					minimumInstanceLifetime = spotinst.Int(value)
				}
			}
			elastigroup.Strategy.SetMinimumInstanceLifetime(minimumInstanceLifetime)
			return nil
		},
		nil,
	)

	fieldsMap[ConsiderODPricing] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		ConsiderODPricing,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *bool = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.ConsiderODPricing != nil {
				value = elastigroup.Strategy.ConsiderODPricing
			}
			if err := resourceData.Set(string(ConsiderODPricing), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ConsiderODPricing), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(ConsiderODPricing)); ok && v != nil {
				codp := v.(bool)
				cnsdrodprice := spotinst.Bool(codp)
				elastigroup.Strategy.SetConsiderODPricing(cnsdrodprice)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var cnsdrodprice *bool = nil
			if v, ok := resourceData.GetOkExists(string(ConsiderODPricing)); ok && v != nil {
				codp := v.(bool)
				cnsdrodprice = spotinst.Bool(codp)
			}
			elastigroup.Strategy.SetConsiderODPricing(cnsdrodprice)
			return nil
		},
		nil,
	)
	fieldsMap[ImmediateODRecoverThreshold] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		ImmediateODRecoverThreshold,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.ImmediateODRecoverThreshold != nil {
				value = elastigroup.Strategy.ImmediateODRecoverThreshold
			} else {
				value = spotinst.Int(-1)
			}
			if err := resourceData.Set(string(ImmediateODRecoverThreshold), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ImmediateODRecoverThreshold), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(ImmediateODRecoverThreshold)); ok && v != nil {
				temp := v.(int)
				if temp >= 0 {
					immediateODRecoverThreshold := spotinst.Int(temp)
					elastigroup.Strategy.SetImmediateODRecoverThreshold(immediateODRecoverThreshold)
				} else {
					elastigroup.Strategy.SetImmediateODRecoverThreshold(nil)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(ImmediateODRecoverThreshold)); ok && v != nil {
				temp := v.(int)
				if temp >= 0 {
					immediateODRecoverThreshold := spotinst.Int(temp)
					elastigroup.Strategy.SetImmediateODRecoverThreshold(immediateODRecoverThreshold)
				} else {
					elastigroup.Strategy.SetImmediateODRecoverThreshold(nil)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[RestrictSingleAz] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		RestrictSingleAz,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *bool = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.RestrictSingleAz != nil {
				value = elastigroup.Strategy.RestrictSingleAz
			}
			if err := resourceData.Set(string(RestrictSingleAz), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RestrictSingleAz), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(RestrictSingleAz)).(bool); ok && v {
				ris := spotinst.Bool(v)
				elastigroup.Strategy.SetRestrictSingleAz(ris)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var ris *bool = nil
			if v, ok := resourceData.Get(string(RestrictSingleAz)).(bool); ok && v {
				ris = spotinst.Bool(v)
			}
			elastigroup.Strategy.SetRestrictSingleAz(ris)
			return nil
		},
		nil,
	)

	fieldsMap[MaxReplacementsPercentage] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		MaxReplacementsPercentage,
		&schema.Schema{
			Type:         schema.TypeInt,
			Optional:     true,
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.MaxReplacementsPercentage != nil {
				value = elastigroup.Strategy.MaxReplacementsPercentage
			}
			if value != nil {
				if err := resourceData.Set(string(MaxReplacementsPercentage), spotinst.IntValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MaxReplacementsPercentage), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(MaxReplacementsPercentage)).(int); ok && v > 0 {
				elastigroup.Strategy.SetMaxReplacementsPercentage(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(MaxReplacementsPercentage)).(int); ok && v > -1 {
				elastigroup.Strategy.SetMaxReplacementsPercentage(spotinst.Int(v))
			} else {
				elastigroup.Strategy.SetMaxReplacementsPercentage(nil)
			}
			return nil
		},
		nil,
	)
}

func flattenAWSGroupScalingStrategy(strategy *aws.ScalingStrategy) []interface{} {
	result := make(map[string]interface{})
	result[string(TerminationPolicy)] = spotinst.StringValue(strategy.TerminationPolicy)
	result[string(TerminateAtEndOfBillingHour)] = spotinst.BoolValue(strategy.TerminateAtEndOfBillingHour)
	return []interface{}{result}
}

func expandAWSGroupScalingStrategy(data interface{}) (*aws.ScalingStrategy, error) {
	strategy := &aws.ScalingStrategy{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(TerminateAtEndOfBillingHour)].(bool); ok {
			strategy.SetTerminateAtEndOfBillingHour(spotinst.Bool(v))
		}

		if v, ok := m[string(TerminationPolicy)].(string); ok && v != "" {
			strategy.SetTerminationPolicy(spotinst.String(v))
		}
	}
	return strategy, nil
}
