package elastigroup_strategy

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupSpotinstElastigroupStrategyResource() {
	fields := make(map[commons.FieldName]*commons.GenericField)
	var readFailurePattern = "elastigroup strategy failed reading field %s - %#v"

	fields[SpotPercentage] = commons.NewGenericField(
		SpotPercentage,
		&schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *float64 = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.Risk != nil {
				value = elastigroup.Strategy.Risk
			}
			if err := resourceData.Set(string(SpotPercentage), spotinst.Float64Value(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(SpotPercentage), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(SpotPercentage)).(float64); ok && v >= 0 {
				elastigroup.Strategy.SetRisk(spotinst.Float64(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(SpotPercentage)).(float64); ok && v >= 0 {
				elastigroup.Strategy.SetRisk(spotinst.Float64(v))
			}
			return nil
		},
		nil,
	)

	fields[OnDemandCount] = commons.NewGenericField(
		OnDemandCount,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *int = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.OnDemandCount != nil {
				value = elastigroup.Strategy.OnDemandCount
			}
			if err := resourceData.Set(string(OnDemandCount), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(OnDemandCount), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(OnDemandCount)).(int); ok && v > 0 {
				elastigroup.Strategy.SetOnDemandCount(spotinst.Int(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(OnDemandCount)).(int); ok && v > 0 {
				elastigroup.Strategy.SetOnDemandCount(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fields[Orientation] = commons.NewGenericField(
		Orientation,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *string = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.AvailabilityVsCost != nil {
				value = elastigroup.Strategy.AvailabilityVsCost
			}
			if err := resourceData.Set(string(Orientation), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(Orientation), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(Orientation)).(string); ok && v != "" {
				elastigroup.Strategy.SetAvailabilityVsCost(spotinst.String(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(Orientation)).(string); ok && v != "" {
				elastigroup.Strategy.SetAvailabilityVsCost(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fields[LifetimePeriod] = commons.NewGenericField(
		LifetimePeriod,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *string = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.LifetimePeriod != nil {
				value = elastigroup.Strategy.LifetimePeriod
			}
			if err := resourceData.Set(string(LifetimePeriod), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(LifetimePeriod), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(LifetimePeriod)).(string); ok && v != "" {
				period := spotinst.String(v)
				elastigroup.Strategy.SetLifetimePeriod(period)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var period *string = nil
			if v, ok := resourceData.Get(string(LifetimePeriod)).(string); ok && v != "" {
				period = spotinst.String(v)
			}
			elastigroup.Strategy.SetLifetimePeriod(period)
			return nil
		},
		nil,
	)

	fields[DrainingTimeout] = commons.NewGenericField(
		DrainingTimeout,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *int = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.DrainingTimeout != nil {
				value = elastigroup.Strategy.DrainingTimeout
			}
			if err := resourceData.Set(string(DrainingTimeout), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(DrainingTimeout), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(DrainingTimeout)).(int); ok && v > 0 {
				elastigroup.Strategy.SetDrainingTimeout(spotinst.Int(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(DrainingTimeout)).(int); ok && v > 0 {
				elastigroup.Strategy.SetDrainingTimeout(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fields[UtilizeReservedInstances] = commons.NewGenericField(
		UtilizeReservedInstances,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *bool = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.UtilizeReservedInstances != nil {
				value = elastigroup.Strategy.UtilizeReservedInstances
			}
			if err := resourceData.Set(string(UtilizeReservedInstances), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(UtilizeReservedInstances), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(UtilizeReservedInstances)).(bool); ok && v {
				ris := spotinst.Bool(v)
				elastigroup.Strategy.SetUtilizeReservedInstances(ris)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var ris *bool = nil
			if v, ok := resourceData.Get(string(UtilizeReservedInstances)).(bool); ok && v {
				ris = spotinst.Bool(v)
			}
			elastigroup.Strategy.SetUtilizeReservedInstances(ris)
			return nil
		},
		nil,
	)

	fields[FallbackToOnDemand] = commons.NewGenericField(
		FallbackToOnDemand,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *bool = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.FallbackToOnDemand != nil {
				value = elastigroup.Strategy.FallbackToOnDemand
			}
			if err := resourceData.Set(string(FallbackToOnDemand), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(FallbackToOnDemand), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(FallbackToOnDemand)).(bool); ok && v {
				fallback := spotinst.Bool(v)
				elastigroup.Strategy.SetFallbackToOnDemand(fallback)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var fallback *bool = nil
			if v, ok := resourceData.Get(string(FallbackToOnDemand)).(bool); ok && v {
				fallback = spotinst.Bool(v)
			}
			elastigroup.Strategy.SetFallbackToOnDemand(fallback)
			return nil
		},
		nil,
	)

	commons.ElastigroupStrategyResource = commons.NewGenericCachedResource(
		string(commons.ElastigroupStrategy),
		fields)
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-