package elastigroup_azure_strategy

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			if v, ok := resourceData.GetOk(string(FallbackToOnDemand)); ok && v != nil {
				ftod := v.(bool)
				fallback := spotinst.Bool(ftod)
				elastigroup.Strategy.SetFallbackToOnDemand(fallback)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var fallback *bool = nil
			if v, ok := resourceData.GetOk(string(FallbackToOnDemand)); ok && v != nil {
				result := v.(bool)
				fallback = spotinst.Bool(result)
			}
			elastigroup.Strategy.SetFallbackToOnDemand(fallback)
			return nil
		},
		nil,
	)
}
