package ocean_ecs_strategy

import (
	"fmt"

	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[DrainingTimeout] = commons.NewGenericField(
		commons.OceanECSStrategy,
		DrainingTimeout,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
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
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()

			if v, ok := resourceData.GetOkExists(string(DrainingTimeout)); ok {
				cluster.Strategy.SetDrainingTimeout(spotinst.Int(v.(int)))
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var dt *int = nil

			if v, ok := resourceData.Get(string(DrainingTimeout)).(int); ok && v > 0 {
				dt = spotinst.Int(v)
			}

			cluster.Strategy.SetDrainingTimeout(dt)

			return nil
		},
		nil,
	)

	fieldsMap[UtilizeReservedInstances] = commons.NewGenericField(
		commons.OceanECSStrategy,
		UtilizeReservedInstances,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
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
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()

			if v, ok := resourceData.GetOkExists(string(UtilizeReservedInstances)); ok {
				cluster.Strategy.SetUtilizeReservedInstances(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var uri *bool = nil
			if v, ok := resourceData.GetOkExists(string(UtilizeReservedInstances)); ok {
				uri = spotinst.Bool(v.(bool))
			}
			cluster.Strategy.SetUtilizeReservedInstances(uri)
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
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
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
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.GetOkExists(string(UtilizeCommitments)); ok && v != nil {
				uc := v.(bool)
				utilizeCommitments := spotinst.Bool(uc)
				cluster.Strategy.SetUtilizeCommitments(utilizeCommitments)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
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

	fieldsMap[SpotPercentage] = commons.NewGenericField(
		commons.OceanECSStrategy,
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
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			//Force setting -1 as default value if it's not exists in initial creation,
			// to allow initialization of the field to 0
			value := spotinst.Int(-1)
			if cluster.Strategy != nil && cluster.Strategy.SpotPercentage != nil {
				value = cluster.Strategy.SpotPercentage
			}
			if err := resourceData.Set(string(SpotPercentage), spotinst.Int(*value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SpotPercentage), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v := resourceData.Get(string(SpotPercentage)).(int); v > -1 {
				cluster.Strategy.SetSpotPercentage(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var spotPercentage *int = nil
			if v := resourceData.Get(string(SpotPercentage)).(int); v > -1 {
				spotPercentage = spotinst.Int(v)
			}
			cluster.Strategy.SetSpotPercentage(spotPercentage)
			return nil
		},
		nil,
	)
	fieldsMap[ClusterOrientation] = commons.NewGenericField(
		commons.OceanECSStrategy,
		ClusterOrientation,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(AvailabilityVsCost): {
						Type:     schema.TypeString,
						Optional: true,
						Default:  "balanced",
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
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
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if value, ok := resourceData.GetOk(string(ClusterOrientation)); ok {
				if clusterOrientation, err := expandClusterOrientation(value); err != nil {
					return err
				} else {
					cluster.Strategy.SetECSClusterOrientation(clusterOrientation)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var result *aws.ECSClusterOrientation = nil
			if v, ok := resourceData.GetOk(string(ClusterOrientation)); ok {
				if clusterOrientation, err := expandClusterOrientation(v); err != nil {
					return err
				} else {
					result = clusterOrientation
				}
			}
			cluster.Strategy.SetECSClusterOrientation(result)
			return nil
		},
		nil,
	)
}
func flattenClusterOrientation(clusterOrientation *aws.ECSClusterOrientation) []interface{} {
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
func expandClusterOrientation(co interface{}) (*aws.ECSClusterOrientation, error) {
	if list := co.([]interface{}); len(list) > 0 {
		clusterOrientation := &aws.ECSClusterOrientation{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})
			if v, ok := m[string(AvailabilityVsCost)].(string); ok {
				clusterOrientation.SetECSAvailabilityVsCost(spotinst.String(v))
			}
		}
		return clusterOrientation, nil
	}
	return nil, nil
}
