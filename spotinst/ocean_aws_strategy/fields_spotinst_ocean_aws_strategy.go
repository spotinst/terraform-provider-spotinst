package ocean_aws_strategy

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[GracePeriod] = commons.NewGenericField(
		commons.OceanAWSStrategy,
		GracePeriod,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *int = nil

			if cluster.Strategy != nil && cluster.Strategy.DrainingTimeout != nil {
				value = cluster.Strategy.GracePeriod
			}
			if value != nil {
				if err := resourceData.Set(string(GracePeriod), spotinst.IntValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(GracePeriod), err)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(GracePeriod)); ok {
				cluster.Strategy.SetGracePeriod(spotinst.Int(v.(int)))
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var dt *int = nil

			if v, ok := resourceData.Get(string(GracePeriod)).(int); ok && v > 0 {
				dt = spotinst.Int(v)
			}

			cluster.Strategy.SetGracePeriod(dt)

			return nil
		},
		nil,
	)

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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			//Force setting -1 as default value if it's not exists in initial creation,
			// to allow initialization of the field to 0
			//SpotPercentage is configured as int in the API but float on Go-SDK (currently not aligning because of breaking changes effects)
			//There for value is of type float and cast as necessary
			value := spotinst.Float64(-1)
			if cluster.Strategy != nil && cluster.Strategy.SpotPercentage != nil {
				value = cluster.Strategy.SpotPercentage
			}
			if err := resourceData.Set(string(SpotPercentage), spotinst.Int(int(*value))); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SpotPercentage), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v := resourceData.Get(string(SpotPercentage)).(int); v > -1 {
				cluster.Strategy.SetSpotPercentage(spotinst.Float64(float64(v)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var spotPercentage *float64 = nil
			if v := resourceData.Get(string(SpotPercentage)).(int); v > -1 {
				spotPercentage = spotinst.Float64(float64(v))
			}
			cluster.Strategy.SetSpotPercentage(spotPercentage)
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
			Default:  true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
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
			if v, ok := resourceData.GetOkExists(string(UtilizeReservedInstances)); ok {
				cluster.Strategy.SetUtilizeReservedInstances(spotinst.Bool(v.(bool)))
			}
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

	fieldsMap[UtilizeCommitments] = commons.NewGenericField(
		commons.OceanECSStrategy,
		UtilizeCommitments,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *bool = nil
			if cluster.Strategy != nil && cluster.Strategy.UtilizeCommitments != nil {
				value = cluster.Strategy.UtilizeCommitments
			}
			if err := resourceData.Set(string(UtilizeCommitments), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UtilizeCommitments), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOkExists(string(UtilizeCommitments)); ok && v != nil {
				uc := v.(bool)
				utilizeCommitments := spotinst.Bool(uc)
				cluster.Strategy.SetUtilizeCommitments(utilizeCommitments)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var utilizeCommitments *bool = nil
			if v, ok := resourceData.GetOkExists(string(UtilizeCommitments)); ok && v != nil {
				uc := v.(bool)
				utilizeCommitments = spotinst.Bool(uc)
			}
			cluster.Strategy.SetUtilizeCommitments(utilizeCommitments)
			return nil
		},
		nil,
	)

	fieldsMap[ClusterOrientation] = commons.NewGenericField(
		commons.OceanAWSStrategy,
		ClusterOrientation,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(AvailabilityVsCost): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value []interface{} = nil
			if cluster.Strategy != nil && cluster.Strategy.ClusterOrientation != nil {
				clusterOrientation := cluster.Strategy.ClusterOrientation
				value = flattenClusterOrientation(clusterOrientation)
			}
			if value != nil {
				if err := resourceData.Set(string(ClusterOrientation), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ClusterOrientation), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if value, ok := resourceData.GetOk(string(ClusterOrientation)); ok {
				if co, err := expandClusterOrientation(value); err != nil {
					return err
				} else {
					cluster.Strategy.SetClusterOrientation(co)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result *aws.ClusterOrientation = nil
			if v, ok := resourceData.GetOk(string(ClusterOrientation)); ok {
				if co, err := expandClusterOrientation(v); err != nil {
					return err
				} else {
					result = co
				}
			}
			cluster.Strategy.SetClusterOrientation(result)
			return nil
		},
		nil,
	)

}
func flattenClusterOrientation(clusterOrientation *aws.ClusterOrientation) []interface{} {
	var out []interface{}
	if clusterOrientation != nil {
		result := make(map[string]interface{})
		if clusterOrientation.AvailabilityVsCost != nil {
			result[string(AvailabilityVsCost)] = spotinst.StringValue(clusterOrientation.AvailabilityVsCost)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}
func expandClusterOrientation(co interface{}) (*aws.ClusterOrientation, error) {
	if list := co.([]interface{}); len(list) > 0 {
		clusterOrientation := &aws.ClusterOrientation{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})
			if v, ok := m[string(AvailabilityVsCost)].(string); ok {
				clusterOrientation.SetAvailabilityVsCost(spotinst.String(v))
			}
		}
		return clusterOrientation, nil
	}
	return nil, nil
}
