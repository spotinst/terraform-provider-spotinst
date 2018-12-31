package ocean_cluster_aws_strategy

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[SpotPercentage] = commons.NewGenericField(
		commons.OceanAWSStrategy,
		SpotPercentage,
		&schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *float64 = nil
			if cluster.Strategy != nil && cluster.Strategy.SpotPercentage != nil {
				value = cluster.Strategy.SpotPercentage
			}
			if err := resourceData.Set(string(SpotPercentage), spotinst.Float64Value(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SpotPercentage), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(SpotPercentage)).(float64); ok && v >= 0 {
				cluster.Strategy.SetSpotPercentage(spotinst.Float64(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var spotPct *float64 = nil
			if v, ok := resourceData.Get(string(SpotPercentage)).(float64); ok && v >= 0 {
				spotPct = spotinst.Float64(v)
			}
			cluster.Strategy.SetSpotPercentage(spotPct)
			return nil
		},
		nil,
	)

	fieldsMap[UtilizeReservedInstances] = commons.NewGenericField(
		commons.OceanAWSStrategy,
		UtilizeReservedInstances,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *bool = nil
			if cluster.Strategy != nil && cluster.Strategy.UtilizeReservedInstances != nil {
				value = cluster.Strategy.UtilizeReservedInstances
			}
			if err := resourceData.Set(string(UtilizeReservedInstances), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UtilizeReservedInstances), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(UtilizeReservedInstances)).(bool); ok && v {
				ris := spotinst.Bool(v)
				cluster.Strategy.SetUtilizeReservedInstances(ris)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var ris *bool = nil
			if v, ok := resourceData.Get(string(UtilizeReservedInstances)).(bool); ok && v {
				ris = spotinst.Bool(v)
			}
			cluster.Strategy.SetUtilizeReservedInstances(ris)
			return nil
		},
		nil,
	)

	fieldsMap[FallbackToOnDemand] = commons.NewGenericField(
		commons.OceanAWSStrategy,
		FallbackToOnDemand,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *bool = nil
			if cluster.Strategy != nil && cluster.Strategy.FallbackToOnDemand != nil {
				value = cluster.Strategy.FallbackToOnDemand
			}
			if err := resourceData.Set(string(FallbackToOnDemand), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(FallbackToOnDemand), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOkExists(string(FallbackToOnDemand)); ok && v != nil {
				ftod := v.(bool)
				fallback := spotinst.Bool(ftod)
				cluster.Strategy.SetFallbackToOnDemand(fallback)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var fallback *bool = nil
			if v, ok := resourceData.GetOkExists(string(FallbackToOnDemand)); ok && v != nil {
				ftod := v.(bool)
				fallback = spotinst.Bool(ftod)
			}
			cluster.Strategy.SetFallbackToOnDemand(fallback)
			return nil
		},
		nil,
	)
}
