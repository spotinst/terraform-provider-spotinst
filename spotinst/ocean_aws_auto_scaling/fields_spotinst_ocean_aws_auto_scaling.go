package ocean_aws_auto_scaling

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Autoscaler] = commons.NewGenericField(
		commons.OceanAWSAutoScaling,
		Autoscaler,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(AutoscaleCooldown): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(AutoHeadroomPercentage): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
						DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
							if old == "-1" && new == "null" {
								return true
							}
							return false
						},
					},

					string(AutoscaleDown): {
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

					string(AutoscaleHeadroom): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(CPUPerUnit): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  -1,
								},

								string(GPUPerUnit): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  -1,
								},
								string(MemoryPerUnit): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  -1,
								},

								string(NumOfUnits): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  -1,
								},
							},
						},
					},

					string(AutoscaleIsAutoConfig): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(AutoscaleIsEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  true,
					},

					string(ResourceLimits): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(MaxVCPU): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(MaxMemoryGIB): {
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
					string(ExtendedResourceDefinitions): {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *aws.AutoScaler = nil

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

func expandAutoscaler(data interface{}, nullify bool) (*aws.AutoScaler, error) {
	autoscaler := &aws.AutoScaler{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return autoscaler, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(AutoscaleCooldown)].(int); ok && v > 0 {
		autoscaler.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m[string(AutoscaleDown)]; ok {
		down, err := expandOceanAWSAutoScalerDown(v)
		if err != nil {
			return nil, err
		}
		if down != nil {
			autoscaler.SetDown(down)
		}
	}

	if v, ok := m[string(AutoscaleHeadroom)]; ok {
		headroom, err := expandOceanAWSAutoScalerHeadroom(v)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			autoscaler.SetHeadroom(headroom)
		} else {
			autoscaler.SetHeadroom(nil)
		}
	}

	if v, ok := m[string(AutoscaleIsAutoConfig)].(bool); ok {
		autoscaler.SetIsAutoConfig(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoHeadroomPercentage)].(int); ok && v > -1 {
		autoscaler.SetAutoHeadroomPercentage(spotinst.Int(v))
	} else if nullify {
		autoscaler.SetAutoHeadroomPercentage(nil)
	}

	if v, ok := m[string(AutoscaleIsEnabled)].(bool); ok {
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

	if v, ok := m[string(EnableAutomaticAndManualHeadroom)].(bool); ok {
		autoscaler.SetEnableAutomaticAndManualHeadroom(spotinst.Bool(v))
	}

	if v, ok := m[string(ExtendedResourceDefinitions)]; ok {
		extendedResourceDefinitions, err := expandExtendedResourceDefinitions(v)
		if err != nil {
			return nil, err
		}
		if extendedResourceDefinitions != nil && len(extendedResourceDefinitions) > 0 {
			autoscaler.SetExtendedResourceDefinitions(extendedResourceDefinitions)
		} else {
			if nullify {
				autoscaler.SetExtendedResourceDefinitions(nil)
			}
		}
	}

	return autoscaler, nil
}

func expandExtendedResourceDefinitions(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))
	for _, v := range list {
		if resourceDefinition, ok := v.(string); ok && resourceDefinition != "" {
			result = append(result, resourceDefinition)
		}
	}

	return result, nil
}

func expandOceanAWSAutoScalerHeadroom(data interface{}) (*aws.AutoScalerHeadroom, error) {
	if list := data.([]interface{}); len(list) > 0 {
		headroom := &aws.AutoScalerHeadroom{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(CPUPerUnit)].(int); ok && v >= 0 {
				headroom.SetCPUPerUnit(spotinst.Int(v))
			} else {
				headroom.SetCPUPerUnit(nil)
			}

			if v, ok := m[string(MemoryPerUnit)].(int); ok && v >= 0 {
				headroom.SetMemoryPerUnit(spotinst.Int(v))
			} else {
				headroom.SetMemoryPerUnit(nil)
			}

			if v, ok := m[string(NumOfUnits)].(int); ok && v >= 0 {
				headroom.SetNumOfUnits(spotinst.Int(v))
			} else {
				headroom.SetNumOfUnits(nil)
			}

			if v, ok := m[string(GPUPerUnit)].(int); ok && v >= 0 {
				headroom.SetGPUPerUnit(spotinst.Int(v))
			} else {
				headroom.SetGPUPerUnit(nil)
			}
			if headroom.MemoryPerUnit == nil && headroom.GPUPerUnit == nil && headroom.CPUPerUnit == nil && headroom.NumOfUnits == nil {
				headroom = nil
			}
		}
		return headroom, nil
	}

	return nil, nil
}

func expandOceanAWSAutoScalerResourceLimits(data interface{}) (*aws.AutoScalerResourceLimits, error) {
	if list := data.([]interface{}); len(list) > 0 {
		resLimits := &aws.AutoScalerResourceLimits{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(MaxMemoryGIB)].(int); ok {
				if v == 0 {
					resLimits.SetMaxMemoryGiB(nil)
				} else {
					resLimits.SetMaxMemoryGiB(spotinst.Int(v))
				}
			}

			if v, ok := m[string(MaxVCPU)].(int); ok {
				if v == 0 {
					resLimits.SetMaxVCPU(nil)
				} else {
					resLimits.SetMaxVCPU(spotinst.Int(v))
				}
			}
		}
		return resLimits, nil
	}

	return nil, nil
}

func expandOceanAWSAutoScalerDown(data interface{}) (*aws.AutoScalerDown, error) {
	if list := data.([]interface{}); len(list) > 0 {
		autoScaleDown := &aws.AutoScalerDown{}
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
				aggressiveScaleDown := &aws.AggressiveScaleDown{}
				autoScaleDown.SetAggressiveScaleDown(aggressiveScaleDown)
				autoScaleDown.AggressiveScaleDown.SetIsEnabled(spotinst.Bool(v))
			}
		}
		return autoScaleDown, nil
	}

	return nil, nil
}

func flattenAutoscaler(autoScaler *aws.AutoScaler) []interface{} {
	var out []interface{}

	if autoScaler != nil {
		result := make(map[string]interface{})

		result[string(AutoscaleIsEnabled)] = spotinst.BoolValue(autoScaler.IsEnabled)
		result[string(AutoscaleCooldown)] = spotinst.IntValue(autoScaler.Cooldown)
		result[string(AutoscaleIsAutoConfig)] = spotinst.BoolValue(autoScaler.IsAutoConfig)
		result[string(EnableAutomaticAndManualHeadroom)] = spotinst.BoolValue(autoScaler.EnableAutomaticAndManualHeadroom)

		value := spotinst.Int(-1)
		if autoScaler.AutoHeadroomPercentage != nil {
			value = autoScaler.AutoHeadroomPercentage
		}
		result[string(AutoHeadroomPercentage)] = spotinst.IntValue(value)

		if autoScaler.Headroom != nil {
			result[string(AutoscaleHeadroom)] = flattenAutoScaleHeadroom(autoScaler.Headroom)
		}

		if autoScaler.Down != nil {
			result[string(AutoscaleDown)] = flattenAutoScaleDown(autoScaler.Down)
		}

		if autoScaler.ResourceLimits != nil {
			result[string(ResourceLimits)] = flattenAutoScaleResourceLimits(autoScaler.ResourceLimits)
		}

		if autoScaler.ExtendedResourceDefinitions != nil {
			result[string(ExtendedResourceDefinitions)] = autoScaler.ExtendedResourceDefinitions
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func flattenAutoScaleHeadroom(autoScaleHeadroom *aws.AutoScalerHeadroom) []interface{} {
	headRoom := make(map[string]interface{})
	headRoom[string(CPUPerUnit)] = spotinst.IntValue(autoScaleHeadroom.CPUPerUnit)
	headRoom[string(MemoryPerUnit)] = spotinst.IntValue(autoScaleHeadroom.MemoryPerUnit)
	headRoom[string(NumOfUnits)] = spotinst.IntValue(autoScaleHeadroom.NumOfUnits)
	headRoom[string(GPUPerUnit)] = spotinst.IntValue(autoScaleHeadroom.GPUPerUnit)

	return []interface{}{headRoom}
}

func flattenAutoScaleDown(autoScaleDown *aws.AutoScalerDown) []interface{} {
	down := make(map[string]interface{})
	down[string(EvaluationPeriods)] = spotinst.IntValue(autoScaleDown.EvaluationPeriods)
	down[string(MaxScaleDownPercentage)] = spotinst.Float64Value(autoScaleDown.MaxScaleDownPercentage)

	if autoScaleDown != nil && autoScaleDown.AggressiveScaleDown != nil {
		down[string(IsAggressiveScaleDownEnabled)] = spotinst.BoolValue(autoScaleDown.AggressiveScaleDown.IsEnabled)
	}

	return []interface{}{down}
}

func flattenAutoScaleResourceLimits(autoScalerResourceLimits *aws.AutoScalerResourceLimits) []interface{} {
	down := make(map[string]interface{})
	value := spotinst.Int(0)
	down[string(MaxVCPU)] = value
	down[string(MaxMemoryGIB)] = value

	if autoScalerResourceLimits.MaxVCPU != nil {
		down[string(MaxVCPU)] = spotinst.IntValue(autoScalerResourceLimits.MaxVCPU)
	}
	if autoScalerResourceLimits.MaxMemoryGiB != nil {
		down[string(MaxMemoryGIB)] = spotinst.IntValue(autoScalerResourceLimits.MaxMemoryGiB)
	}
	return []interface{}{down}
}
