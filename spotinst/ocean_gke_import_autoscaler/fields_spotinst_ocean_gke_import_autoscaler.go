package ocean_gke_import_autoscaler

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Autoscaler] = commons.NewGenericField(
		commons.OceanGKEImportAutoScaler,
		Autoscaler,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(IsEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  true,
					},

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
								string(EvaluationPeriods): {
									Type:     schema.TypeInt,
									Optional: true,
								},
								string(IsAggressiveScaleDownEnabled): {
									Type:     schema.TypeBool,
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

								string(GPUPerUnit): {
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
						Default:  true,
					},

					string(AutoHeadroomPercentage): {
						Type:     schema.TypeInt,
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

					string(EnableAutomaticAndManualHeadroom): {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
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
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(Autoscaler)); ok {
				if autoscaler, err := expandAutoscaler(v); err != nil {
					return err
				} else {
					cluster.SetAutoScaler(autoscaler)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *gcp.AutoScaler = nil

			if v, ok := resourceData.GetOk(string(Autoscaler)); ok {
				if autoscaler, err := expandAutoscaler(v); err != nil {
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

func expandAutoscaler(data interface{}) (*gcp.AutoScaler, error) {
	autoscaler := &gcp.AutoScaler{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return autoscaler, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(IsEnabled)].(bool); ok {
		autoscaler.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(Cooldown)].(int); ok && v > 0 {
		autoscaler.SetCooldown(spotinst.Int(v))
	} else {
		autoscaler.SetCooldown(nil)
	}

	if v, ok := m[string(IsAutoConfig)].(bool); ok {
		autoscaler.SetIsAutoConfig(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoHeadroomPercentage)].(int); ok && v > 0 {
		autoscaler.SetAutoHeadroomPercentage(spotinst.Int(v))
	} else {
		autoscaler.SetAutoHeadroomPercentage(nil)
	}

	if v, ok := m[string(Down)]; ok {
		down, err := expandOceanGCPAutoScalerDown(v)
		if err != nil {
			return nil, err
		}
		if down != nil {
			autoscaler.SetDown(down)
		} else {
			autoscaler.SetDown(nil)
		}
	}

	if v, ok := m[string(Headroom)]; ok {
		headroom, err := expandOceanGCPAutoScalerHeadroom(v)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			autoscaler.SetHeadroom(headroom)
		} else {
			autoscaler.SetHeadroom(nil)
		}
	}

	if v, ok := m[string(ResourceLimits)]; ok {
		resLimits, err := expandOceanAWSAutoScalerResourceLimits(v)
		if err != nil {
			return nil, err
		}
		if resLimits != nil {
			autoscaler.SetResourceLimits(resLimits)
		} else {
			autoscaler.SetResourceLimits(nil)
		}
	}

	if v, ok := m[string(EnableAutomaticAndManualHeadroom)].(bool); ok {
		autoscaler.SetEnableAutomaticAndManualHeadroom(spotinst.Bool(v))
	}

	return autoscaler, nil
}

func expandOceanGCPAutoScalerDown(data interface{}) (*gcp.AutoScalerDown, error) {
	if list := data.([]interface{}); len(list) > 0 {
		autoScaleDown := &gcp.AutoScalerDown{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(EvaluationPeriods)].(int); ok && v > 0 {
				autoScaleDown.SetEvaluationPeriods(spotinst.Int(v))
			}

			if v, ok := m[string(MaxScaleDownPercentage)].(float64); ok && v > 0 {
				autoScaleDown.SetMaxScaleDownPercentage(spotinst.Float64(v))
			} else {
				autoScaleDown.SetMaxScaleDownPercentage(nil)
			}

			if v, ok := m[string(IsAggressiveScaleDownEnabled)].(bool); ok {
				aggressiveScaleDown := &gcp.AggressiveScaleDown{}
				autoScaleDown.SetAggressiveScaleDown(aggressiveScaleDown)
				autoScaleDown.AggressiveScaleDown.SetIsEnabled(spotinst.Bool(v))
			}
		}
		return autoScaleDown, nil
	}

	return nil, nil
}

func expandOceanGCPAutoScalerHeadroom(data interface{}) (*gcp.AutoScalerHeadroom, error) {
	if list := data.([]interface{}); len(list) > 0 {
		headroom := &gcp.AutoScalerHeadroom{}
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

			if v, ok := m[string(GPUPerUnit)].(int); ok && v >= 0 {
				headroom.SetGPUPerUnit(spotinst.Int(v))
			}
		}
		return headroom, nil
	}

	return nil, nil
}

func expandOceanAWSAutoScalerResourceLimits(data interface{}) (*gcp.AutoScalerResourceLimits, error) {
	if list := data.([]interface{}); len(list) > 0 {
		resLimits := &gcp.AutoScalerResourceLimits{}
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

func flattenAutoscaler(autoScaler *gcp.AutoScaler) []interface{} {
	var out []interface{}

	if autoScaler != nil {
		result := make(map[string]interface{})

		result[string(IsEnabled)] = spotinst.BoolValue(autoScaler.IsEnabled)
		result[string(Cooldown)] = spotinst.IntValue(autoScaler.Cooldown)
		result[string(IsAutoConfig)] = spotinst.BoolValue(autoScaler.IsAutoConfig)
		result[string(AutoHeadroomPercentage)] = spotinst.IntValue(autoScaler.AutoHeadroomPercentage)
		result[string(EnableAutomaticAndManualHeadroom)] = spotinst.BoolValue(autoScaler.EnableAutomaticAndManualHeadroom)

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

func flattenAutoScaleHeadroom(autoScaleHeadroom *gcp.AutoScalerHeadroom) []interface{} {
	headRoom := make(map[string]interface{})
	headRoom[string(CpuPerUnit)] = spotinst.IntValue(autoScaleHeadroom.CPUPerUnit)
	headRoom[string(MemoryPerUnit)] = spotinst.IntValue(autoScaleHeadroom.MemoryPerUnit)
	headRoom[string(NumOfUnits)] = spotinst.IntValue(autoScaleHeadroom.NumOfUnits)
	headRoom[string(GPUPerUnit)] = spotinst.IntValue(autoScaleHeadroom.GPUPerUnit)

	return []interface{}{headRoom}
}

func flattenAutoScaleDown(autoScaleDown *gcp.AutoScalerDown) []interface{} {
	down := make(map[string]interface{})
	down[string(EvaluationPeriods)] = spotinst.IntValue(autoScaleDown.EvaluationPeriods)
	down[string(MaxScaleDownPercentage)] = spotinst.Float64Value(autoScaleDown.MaxScaleDownPercentage)
	if autoScaleDown != nil && autoScaleDown.AggressiveScaleDown != nil {
		down[string(IsAggressiveScaleDownEnabled)] = spotinst.BoolValue(autoScaleDown.AggressiveScaleDown.IsEnabled)
	}

	return []interface{}{down}
}

func flattenAutoScaleResourceLimits(autoScalerResourceLimits *gcp.AutoScalerResourceLimits) []interface{} {
	down := make(map[string]interface{})
	down[string(MaxVCpu)] = spotinst.IntValue(autoScalerResourceLimits.MaxVCPU)
	down[string(MaxMemoryGib)] = spotinst.IntValue(autoScalerResourceLimits.MaxMemoryGiB)
	return []interface{}{down}
}
