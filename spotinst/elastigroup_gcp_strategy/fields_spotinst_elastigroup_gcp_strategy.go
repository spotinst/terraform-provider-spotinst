package elastigroup_gcp_strategy

import (
	"fmt"

	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

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

	fieldsMap[ProvisioningModel] = commons.NewGenericField(
		commons.ElastigroupGCPStrategy,
		ProvisioningModel,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.ProvisioningModel != nil {
				value = elastigroup.Strategy.ProvisioningModel
			}
			if err := resourceData.Set(string(ProvisioningModel), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ProvisioningModel), err)
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(ProvisioningModel)); ok {
				value := v.(string)
				elastigroup.Strategy.SetProvisioningModel(spotinst.String(value))
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var pm *string = nil
			if v, ok := resourceData.GetOkExists(string(ProvisioningModel)); ok && v != nil {
				if value, ok := v.(string); ok && value != "" {
					pm = spotinst.String(value)
				}
			}
			elastigroup.Strategy.SetProvisioningModel(pm)
			return nil
		},
		nil,
	)
	fieldsMap[OptimizationWindows] = commons.NewGenericField(
		commons.ElastigroupGCPStrategy,
		OptimizationWindows,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []string
			if elastigroup.Strategy != nil && elastigroup.Strategy.OptimizationWindows != nil {
				result = append(result, elastigroup.Strategy.OptimizationWindows...)
			}
			if err := resourceData.Set(string(OptimizationWindows), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OptimizationWindows), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(OptimizationWindows)); ok {
				optimizationWindowList := v.([]interface{})
				optimizationWindow := make([]string, len(optimizationWindowList))
				for i, j := range optimizationWindowList {
					optimizationWindow[i] = j.(string)
				}
				elastigroup.Strategy.SetOptimizationWindows(optimizationWindow)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(OptimizationWindows)); ok {
				optimizationWindowList := v.([]interface{})
				optimizationWindow := make([]string, len(optimizationWindowList))
				for i, j := range optimizationWindowList {
					optimizationWindow[i] = j.(string)
				}
				elastigroup.Strategy.SetOptimizationWindows(optimizationWindow)
			} else {
				elastigroup.Strategy.SetOptimizationWindows(nil)
			}
			return nil
		},
		nil,
	)
	fieldsMap[RevertToPreemptible] = commons.NewGenericField(
		commons.ElastigroupGCPStrategy,
		RevertToPreemptible,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
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
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result interface{} = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.RevertToPreemptible != nil &&
				elastigroup.Strategy.RevertToPreemptible.PerformAt != nil {
				revertToPreemptible := elastigroup.Strategy.RevertToPreemptible
				result = flattenRevertToPreemptible(revertToPreemptible)
			}
			if result != nil {
				if err := resourceData.Set(string(RevertToPreemptible), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RevertToPreemptible), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(RevertToPreemptible)); ok {
				if revertToPreemptible, err := expandRevertToPreemptible(v); err != nil {
					return err
				} else {
					elastigroup.Strategy.SetRevertToPreemptible(revertToPreemptible)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result *gcp.RevertToPreemptible = nil
			if v, ok := resourceData.GetOk(string(RevertToPreemptible)); ok {
				if value, err := expandRevertToPreemptible(v); err != nil {
					return err
				} else {
					result = value
				}
			}
			elastigroup.Strategy.SetRevertToPreemptible(result)
			return nil
		},
		nil,
	)

	fieldsMap[ShouldUtilizeCommitments] = commons.NewGenericField(
		commons.ElastigroupGCPStrategy,
		ShouldUtilizeCommitments,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *bool = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.ShouldUtilizeCommitments != nil {
				value = elastigroup.Strategy.ShouldUtilizeCommitments
			}
			if err := resourceData.Set(string(ShouldUtilizeCommitments), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShouldUtilizeCommitments), err)
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(ShouldUtilizeCommitments)); ok {
				ftod := v.(bool)
				elastigroup.Strategy.SetShouldUtilizeCommitments(spotinst.Bool(ftod))
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var fallback *bool = nil
			if v, ok := resourceData.GetOkExists(string(ShouldUtilizeCommitments)); ok {
				ftod := v.(bool)
				fallback = spotinst.Bool(ftod)
			}
			elastigroup.Strategy.SetShouldUtilizeCommitments(fallback)
			return nil
		},
		nil,
	)
}
func flattenRevertToPreemptible(revertToPreemptible *gcp.RevertToPreemptible) []interface{} {
	result := make([]interface{}, 0, 1)
	m := make(map[string]string)
	m[string(PerformAt)] = spotinst.StringValue(revertToPreemptible.PerformAt)
	result = append(result, m)

	return result
}
func expandRevertToPreemptible(data interface{}) (*gcp.RevertToPreemptible, error) {
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		revertToPreemptibles := make([]*gcp.RevertToPreemptible, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			revertToPreemptible := &gcp.RevertToPreemptible{}

			if v, ok := m[string(PerformAt)].(string); ok && v != "" {
				revertToPreemptible.SetPerformAt(spotinst.String(v))
			}
			revertToPreemptibles = append(revertToPreemptibles, revertToPreemptible)
		}
		return revertToPreemptibles[0], nil
	}
	return nil, nil

}
