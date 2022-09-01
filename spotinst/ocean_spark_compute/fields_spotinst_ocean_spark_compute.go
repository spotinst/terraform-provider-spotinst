package ocean_spark_compute

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/spark"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Compute] = commons.NewGenericField(
		commons.OceanSparkCompute,
		Compute,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(UseTaints): {
						Type:     schema.TypeBool,
						Optional: true,
						Computed: true,
					},

					string(CreateVNGs): {
						Type:     schema.TypeBool,
						Optional: true,
						Computed: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster.Config != nil && cluster.Config.Compute != nil {
				result = flattenCompute(cluster.Config.Compute)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Compute), result); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(Compute), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if value, ok := resourceData.GetOk(string(Compute)); ok {
				if compute, err := expandCompute(value, false); err != nil {
					return err
				} else {
					if cluster.Config == nil {
						cluster.Config = &spark.Config{}
					}
					cluster.Config.SetCompute(compute)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *spark.ComputeConfig = nil
			if v, ok := resourceData.GetOk(string(Compute)); ok {
				if compute, err := expandCompute(v, true); err != nil {
					return err
				} else {
					value = compute
				}
			}
			if cluster.Config == nil {
				cluster.Config = &spark.Config{}
			}
			cluster.Config.SetCompute(value)
			return nil
		},
		nil,
	)
}

func flattenCompute(compute *spark.ComputeConfig) []interface{} {
	if compute == nil {
		return nil
	}
	result := make(map[string]interface{})
	result[string(UseTaints)] = spotinst.BoolValue(compute.UseTaints)
	result[string(CreateVNGs)] = spotinst.BoolValue(compute.CreateVngs)
	return []interface{}{result}
}

func expandCompute(data interface{}, nullify bool) (*spark.ComputeConfig, error) {
	compute := &spark.ComputeConfig{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return compute, nil
	}
	m := list[0].(map[string]interface{})

	// TODO Do I need to nullify the bools?
	if v, ok := m[string(UseTaints)].(bool); ok {
		compute.SetUseTaints(spotinst.Bool(v))
	}

	if v, ok := m[string(CreateVNGs)].(bool); ok {
		compute.SetCreateVNGs(spotinst.Bool(v))
	}

	return compute, nil
}
