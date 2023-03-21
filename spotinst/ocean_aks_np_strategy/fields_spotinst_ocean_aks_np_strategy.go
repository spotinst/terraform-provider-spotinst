package ocean_aks_np_strategy

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			Default:  100,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *int = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.Strategy != nil && cluster.VirtualNodeGroupTemplate.Strategy.SpotPercentage != nil {
				value = cluster.VirtualNodeGroupTemplate.Strategy.SpotPercentage
			}
			if err := resourceData.Set(string(SpotPercentage), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SpotPercentage), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(SpotPercentage)).(int); ok && v > 0 {
				cluster.VirtualNodeGroupTemplate.Strategy.SetSpotPercentage(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.Strategy.SetSpotPercentage(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(SpotPercentage)).(int); ok && v > 0 {
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
			Default:  true,
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
			if v, ok := resourceData.GetOk(string(FallbackToOnDemand)); ok && v != nil {
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
			if v, ok := resourceData.GetOk(string(FallbackToOnDemand)); ok && v != nil {
				ftod := v.(bool)
				fallback = spotinst.Bool(ftod)
			}
			cluster.VirtualNodeGroupTemplate.Strategy.SetFallbackToOD(fallback)
			return nil
		},
		nil,
	)
}
