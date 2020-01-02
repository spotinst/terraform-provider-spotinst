package ocean_aws_strategy

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
		commons.OceanAWSStrategy,
		DrainingTimeout,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *int = nil

			if cluster.Strategy != nil && cluster.Strategy.DrainingTimeout != nil {
				value = cluster.Strategy.DrainingTimeout
			}
			if value != nil {
				if err := resourceData.Set(string(DrainingTimeout), spotinst.IntValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DrainingTimeout), err)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(DrainingTimeout)); ok {
				cluster.Strategy.SetDrainingTimeout(spotinst.Int(v.(int)))
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var dt *int = nil

			if v, ok := resourceData.Get(string(DrainingTimeout)).(int); ok && v > 0 {
				dt = spotinst.Int(v)
			}

			cluster.Strategy.SetDrainingTimeout(dt)

			return nil
		},
		nil,
	)

	fieldsMap[SpotPercentage] = commons.NewGenericField(
		commons.OceanAWSStrategy,
		SpotPercentage,
		&schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(SpotPercentage)); ok {
				spotPct := v.(float64)
				cluster.Strategy.SetSpotPercentage(spotinst.Float64(spotPct))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var spotPct *float64 = nil
			if v, ok := resourceData.GetOk(string(SpotPercentage)); ok {
				spotPct = spotinst.Float64(v.(float64))
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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *bool = nil
			if cluster.Strategy != nil && cluster.Strategy.UtilizeReservedInstances != nil {
				value = cluster.Strategy.UtilizeReservedInstances
			}
			if value != nil {
				if err := resourceData.Set(string(UtilizeReservedInstances), spotinst.BoolValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UtilizeReservedInstances), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(UtilizeReservedInstances)); ok {
				cluster.Strategy.SetUtilizeReservedInstances(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var uri *bool = nil
			if v, ok := resourceData.GetOkExists(string(UtilizeReservedInstances)); ok {
				uri = spotinst.Bool(v.(bool))
			}
			cluster.Strategy.SetUtilizeReservedInstances(uri)
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
			Default:  true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *bool = nil
			if cluster.Strategy != nil && cluster.Strategy.FallbackToOnDemand != nil {
				value = cluster.Strategy.FallbackToOnDemand
			}
			if value != nil {
				if err := resourceData.Set(string(FallbackToOnDemand), spotinst.BoolValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(FallbackToOnDemand), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOkExists(string(FallbackToOnDemand)); ok && v != nil {
				ftod := v.(bool)
				fallback := spotinst.Bool(ftod)
				cluster.Strategy.SetFallbackToOnDemand(fallback)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
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
