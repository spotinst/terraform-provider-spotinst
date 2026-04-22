package ocean_aks_np_strategy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[SpotPercentage] = commons.NewGenericField(
		commons.OceanAKSNPStrategy,
		SpotPercentage,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *int = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.Strategy != nil && cluster.VirtualNodeGroupTemplate.Strategy.SpotPercentage != nil {
				value = cluster.VirtualNodeGroupTemplate.Strategy.SpotPercentage
			} else {
				value = spotinst.Int(-1)
			}
			if err := resourceData.Set(string(SpotPercentage), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SpotPercentage), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(SpotPercentage)).(int); ok && v >= 0 {
				cluster.VirtualNodeGroupTemplate.Strategy.SetSpotPercentage(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.Strategy.SetSpotPercentage(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(SpotPercentage)).(int); ok && v >= 0 {
				cluster.VirtualNodeGroupTemplate.Strategy.SetSpotPercentage(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.Strategy.SetSpotPercentage(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[FallbackToOnDemand] = commons.NewGenericField(
		commons.OceanAKSNPStrategy,
		FallbackToOnDemand,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *bool = nil
			if cluster.VirtualNodeGroupTemplate.Strategy != nil && cluster.VirtualNodeGroupTemplate.Strategy.FallbackToOD != nil {
				value = cluster.VirtualNodeGroupTemplate.Strategy.FallbackToOD
			}
			if value != nil {
				if err := resourceData.Set(string(FallbackToOnDemand), spotinst.BoolValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(FallbackToOnDemand), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOkExists(string(FallbackToOnDemand)); ok && v != nil {
				ftod := v.(bool)
				fallback := spotinst.Bool(ftod)
				cluster.VirtualNodeGroupTemplate.Strategy.SetFallbackToOD(fallback)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var fallback *bool = nil
			if v, ok := resourceData.GetOkExists(string(FallbackToOnDemand)); ok && v != nil {
				ftod := v.(bool)
				fallback = spotinst.Bool(ftod)
			}
			cluster.VirtualNodeGroupTemplate.Strategy.SetFallbackToOD(fallback)
			return nil
		},
		nil,
	)

	fieldsMap[DrainingTimeout] = commons.NewGenericField(
		commons.OceanAKSNPStrategy,
		DrainingTimeout,
		&schema.Schema{
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(300, 3600),
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *int = nil
			if cluster.VirtualNodeGroupTemplate.Strategy != nil && cluster.VirtualNodeGroupTemplate.Strategy.DrainingTimeout != nil {
				value = cluster.VirtualNodeGroupTemplate.Strategy.DrainingTimeout
			}
			if value != nil {
				if err := resourceData.Set(string(DrainingTimeout), spotinst.IntValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DrainingTimeout), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOkExists(string(DrainingTimeout)); ok {
				cluster.VirtualNodeGroupTemplate.Strategy.SetDrainingTimeout(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var dt *int = nil
			if v, ok := resourceData.Get(string(DrainingTimeout)).(int); ok && v > 0 {
				dt = spotinst.Int(v)
			}
			cluster.VirtualNodeGroupTemplate.Strategy.SetDrainingTimeout(dt)
			return nil
		},
		nil,
	)

	fieldsMap[ShouldUtilizeCommitments] = commons.NewGenericField(
		commons.OceanAKSNPStrategy,
		ShouldUtilizeCommitments,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *bool = nil
			if cluster.VirtualNodeGroupTemplate.Strategy != nil && cluster.VirtualNodeGroupTemplate.Strategy.ShouldUtilizeCommitments != nil {
				value = cluster.VirtualNodeGroupTemplate.Strategy.ShouldUtilizeCommitments
			}
			if value != nil {
				if err := resourceData.Set(string(ShouldUtilizeCommitments), spotinst.BoolValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShouldUtilizeCommitments), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOkExists(string(ShouldUtilizeCommitments)); ok && v != nil {
				suc := v.(bool)
				commitments := spotinst.Bool(suc)
				cluster.VirtualNodeGroupTemplate.Strategy.SetShouldUtilizeCommitments(commitments)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var suc *bool = nil
			if v, ok := resourceData.GetOkExists(string(ShouldUtilizeCommitments)); ok && v != nil {
				shouldUse := v.(bool)
				suc = spotinst.Bool(shouldUse)
			}
			cluster.VirtualNodeGroupTemplate.Strategy.SetShouldUtilizeCommitments(suc)
			return nil
		},
		nil,
	)

}
