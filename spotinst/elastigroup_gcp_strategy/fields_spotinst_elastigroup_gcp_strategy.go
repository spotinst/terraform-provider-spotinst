package elastigroup_gcp_strategy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[DrainingTimeout] = commons.NewGenericField(
		commons.ElastigroupGCPStrategy,
		DrainingTimeout,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
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
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(DrainingTimeout)); ok {
				value := v.(int)
				elastigroup.Strategy.SetDrainingTimeout(spotinst.Int(value))
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(DrainingTimeout)); ok {
				value := v.(int)
				elastigroup.Strategy.SetDrainingTimeout(spotinst.Int(value))
			}
			return nil
		},
		nil,
	)

	fieldsMap[FallbackToOnDemand] = commons.NewGenericField(
		commons.ElastigroupGCPStrategy,
		FallbackToOnDemand,
		&schema.Schema{
			Type:             schema.TypeBool,
			Optional:         true,
			DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
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
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(FallbackToOnDemand)); ok {
				ftod := v.(bool)
				elastigroup.Strategy.SetFallbackToOnDemand(spotinst.Bool(ftod))
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var fallback *bool = nil
			if v, ok := resourceData.GetOkExists(string(FallbackToOnDemand)); ok {
				ftod := v.(bool)
				fallback = spotinst.Bool(ftod)
			}
			elastigroup.Strategy.SetFallbackToOnDemand(fallback)
			return nil
		},
		nil,
	)

	fieldsMap[OnDemandCount] = commons.NewGenericField(
		commons.ElastigroupGCPStrategy,
		OnDemandCount,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.OnDemandCount != nil {
				value = elastigroup.Strategy.OnDemandCount
			}
			if err := resourceData.Set(string(OnDemandCount), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OnDemandCount), err)
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(OnDemandCount)); ok {
				value := v.(int)
				elastigroup.Strategy.SetOnDemandCount(spotinst.Int(value))
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(OnDemandCount)); ok {
				value := v.(int)
				elastigroup.Strategy.SetOnDemandCount(spotinst.Int(value))
			}
			return nil
		},
		nil,
	)

	fieldsMap[PreemptiblePercentage] = commons.NewGenericField(
		commons.ElastigroupGCPStrategy,
		PreemptiblePercentage,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.PreemptiblePercentage != nil {
				value = elastigroup.Strategy.PreemptiblePercentage
			}
			if err := resourceData.Set(string(PreemptiblePercentage), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PreemptiblePercentage), err)
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(PreemptiblePercentage)); ok {
				value := v.(int)
				elastigroup.Strategy.SetPreemptiblePercentage(spotinst.Int(value))
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(PreemptiblePercentage)); ok {
				value := v.(int)
				elastigroup.Strategy.SetPreemptiblePercentage(spotinst.Int(value))
			}
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
