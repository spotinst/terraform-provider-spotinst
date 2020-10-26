package ocean_ecs_autoscaler

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Autoscaler] = commons.NewGenericField(
		commons.OceanECSAutoScaler,
		Autoscaler,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(Cooldown): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(Down): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(MaxScaleDownPercentage): {
									Type:     schema.TypeFloat,
									Optional: true,
								},
							},
						},
					},

					string(Headroom): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(CpuPerUnit): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(MemoryPerUnit): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(NumOfUnits): {
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},

					string(IsAutoConfig): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(IsEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(ResourceLimits): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(MaxVCpu): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(MaxMemoryGib): {
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var result []interface{} = nil

			if cluster != nil && cluster.AutoScaler != nil {
				result = flattenAutoscaler(cluster.AutoScaler)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Autoscaler), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Autoscaler), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.GetOk(string(Autoscaler)); ok {
				if autoscaler, err := expandAutoscaler(v, false); err != nil {
					return err
				} else {
					cluster.SetAutoScaler(autoscaler)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value *aws.ECSAutoScaler = nil

			if v, ok := resourceData.GetOk(string(Autoscaler)); ok {
				if autoscaler, err := expandAutoscaler(v, true); err != nil {
					return err
				} else {
					value = autoscaler
				}
			}
			cluster.SetAutoScaler(value)
			return nil
		},

		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAutoscaler(data interface{}, nullify bool) (*aws.ECSAutoScaler, error) {
	autoscaler := &aws.ECSAutoScaler{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return autoscaler, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Cooldown)].(int); ok && v > 0 {
		autoscaler.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m[string(Down)]; ok {
		down, err := expandOceanAWSAutoScalerDown(v)
		if err != nil {
			return nil, err
		}
		if down != nil {
			autoscaler.SetDown(down)
		}
	}

	if v, ok := m[string(Headroom)]; ok {
		headroom, err := expandOceanAWSAutoScalerHeadroom(v)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			autoscaler.SetHeadroom(headroom)
		} else {
			autoscaler.Headroom = nil
		}
	}

	if v, ok := m[string(IsAutoConfig)].(bool); ok {
		autoscaler.SetIsAutoConfig(spotinst.Bool(v))
	}

	if v, ok := m[string(IsEnabled)].(bool); ok {
		autoscaler.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(ResourceLimits)]; ok {
		resLimits, err := expandOceanAWSAutoScalerResourceLimits(v)
		if err != nil {
			return nil, err
		}
		if resLimits != nil {
			autoscaler.SetResourceLimits(resLimits)
		} else {
			autoscaler.ResourceLimits = nil
		}
	}

	return autoscaler, nil
}

func expandOceanAWSAutoScalerHeadroom(data interface{}) (*aws.ECSAutoScalerHeadroom, error) {
	if list := data.([]interface{}); len(list) > 0 {
		headroom := &aws.ECSAutoScalerHeadroom{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(CpuPerUnit)].(int); ok && v >= 0 {
				headroom.SetCPUPerUnit(spotinst.Int(v))
			}

			if v, ok := m[string(MemoryPerUnit)].(int); ok && v >= 0 {
				headroom.SetMemoryPerUnit(spotinst.Int(v))
			}

			if v, ok := m[string(NumOfUnits)].(int); ok && v >= 0 {
				headroom.SetNumOfUnits(spotinst.Int(v))
			}
		}
		return headroom, nil
	}

	return nil, nil
}

func expandOceanAWSAutoScalerResourceLimits(data interface{}) (*aws.ECSAutoScalerResourceLimits, error) {
	if list := data.([]interface{}); len(list) > 0 {
		resLimits := &aws.ECSAutoScalerResourceLimits{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(MaxMemoryGib)].(int); ok && v > 0 {
				resLimits.SetMaxMemoryGiB(spotinst.Int(v))
			}

			if v, ok := m[string(MaxVCpu)].(int); ok && v > 0 {
				resLimits.SetMaxVCPU(spotinst.Int(v))
			}

		}
		return resLimits, nil
	}

	return nil, nil
}

func expandOceanAWSAutoScalerDown(data interface{}) (*aws.ECSAutoScalerDown, error) {
	if list := data.([]interface{}); len(list) > 0 {
		autoScaleDown := &aws.ECSAutoScalerDown{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(MaxScaleDownPercentage)].(float64); ok && v > 0 {
				autoScaleDown.SetMaxScaleDownPercentage(spotinst.Float64(v))
			}
		}
		return autoScaleDown, nil
	}

	return nil, nil
}

func flattenAutoscaler(autoScaler *aws.ECSAutoScaler) []interface{} {
	var out []interface{}

	if autoScaler != nil {
		result := make(map[string]interface{})

		result[string(IsEnabled)] = spotinst.BoolValue(autoScaler.IsEnabled)
		result[string(Cooldown)] = spotinst.IntValue(autoScaler.Cooldown)
		result[string(IsAutoConfig)] = spotinst.BoolValue(autoScaler.IsAutoConfig)

		if autoScaler.Headroom != nil {
			result[string(Headroom)] = flattenAutoScaleHeadroom(autoScaler.Headroom)
		}

		if autoScaler.Down != nil {
			result[string(Down)] = flattenAutoScaleDown(autoScaler.Down)
		}

		if autoScaler.ResourceLimits != nil {
			result[string(ResourceLimits)] = flattenAutoScaleResourceLimits(autoScaler.ResourceLimits)
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func flattenAutoScaleDown(autoScaleDown *aws.ECSAutoScalerDown) []interface{} {
	down := make(map[string]interface{})
	down[string(MaxScaleDownPercentage)] = spotinst.Float64Value(autoScaleDown.MaxScaleDownPercentage)

	return []interface{}{down}
}

func flattenAutoScaleHeadroom(autoScaleHeadroom *aws.ECSAutoScalerHeadroom) []interface{} {
	headRoom := make(map[string]interface{})
	headRoom[string(CpuPerUnit)] = spotinst.IntValue(autoScaleHeadroom.CPUPerUnit)
	headRoom[string(MemoryPerUnit)] = spotinst.IntValue(autoScaleHeadroom.MemoryPerUnit)
	headRoom[string(NumOfUnits)] = spotinst.IntValue(autoScaleHeadroom.NumOfUnits)

	return []interface{}{headRoom}
}

func flattenAutoScaleResourceLimits(autoScalerResourceLimits *aws.ECSAutoScalerResourceLimits) []interface{} {
	down := make(map[string]interface{})
	down[string(MaxVCpu)] = spotinst.IntValue(autoScalerResourceLimits.MaxVCPU)
	down[string(MaxMemoryGib)] = spotinst.IntValue(autoScalerResourceLimits.MaxMemoryGiB)
	return []interface{}{down}
}
