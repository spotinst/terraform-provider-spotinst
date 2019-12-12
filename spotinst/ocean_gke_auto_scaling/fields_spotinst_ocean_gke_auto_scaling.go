package ocean_gke_auto_scaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Autoscaler] = commons.NewGenericField(
		commons.OceanGKEAutoScaling,
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

					string(AutoscaleDown): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(EvaluationPeriods): {
									Type:     schema.TypeInt,
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
								},

								string(GPUPerUnit): {
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

					string(AutoscaleIsAutoConfig): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(AutoscaleIsEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
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
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *gcp.AutoScaler = nil

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
func expandAutoscaler(data interface{}, nullify bool) (*gcp.AutoScaler, error) {
	autoscaler := &gcp.AutoScaler{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return autoscaler, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(AutoscaleCooldown)].(int); ok && v > 0 {
		autoscaler.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m[string(AutoscaleDown)]; ok {
		down, err := expandOceanGKEAutoScalerDown(v)
		if err != nil {
			return nil, err
		}
		if down != nil {
			autoscaler.SetDown(down)
		}
	}

	if v, ok := m[string(AutoscaleHeadroom)]; ok {
		headroom, err := expandOceanGKEAutoScalerHeadroom(v)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			autoscaler.SetHeadroom(headroom)
		} else {
			autoscaler.Headroom = nil
		}
	}

	if v, ok := m[string(AutoscaleIsAutoConfig)].(bool); ok {
		autoscaler.SetIsAutoConfig(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoscaleIsEnabled)].(bool); ok {
		autoscaler.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(ResourceLimits)]; ok {
		resLimits, err := expandOceanGKEAutoScalerResourceLimits(v)
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

func expandOceanGKEAutoScalerHeadroom(data interface{}) (*gcp.AutoScalerHeadroom, error) {
	if list := data.([]interface{}); len(list) > 0 {
		headroom := &gcp.AutoScalerHeadroom{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(CPUPerUnit)].(int); ok && v >= 0 {
				headroom.SetCPUPerUnit(spotinst.Int(v))
			}

			if v, ok := m[string(MemoryPerUnit)].(int); ok && v >= 0 {
				headroom.SetMemoryPerUnit(spotinst.Int(v))
			}

			if v, ok := m[string(NumOfUnits)].(int); ok && v >= 0 {
				headroom.SetNumOfUnits(spotinst.Int(v))
			}

			if v, ok := m[string(GPUPerUnit)].(int); ok && v >= 0 {
				headroom.SetNumOfUnits(spotinst.Int(v))
			}
		}
		return headroom, nil
	}

	return nil, nil
}

func expandOceanGKEAutoScalerResourceLimits(data interface{}) (*gcp.AutoScalerResourceLimits, error) {
	if list := data.([]interface{}); len(list) > 0 {
		resLimits := &gcp.AutoScalerResourceLimits{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(MaxMemoryGIB)].(int); ok && v > 0 {
				resLimits.SetMaxMemoryGiB(spotinst.Int(v))
			}

			if v, ok := m[string(MaxVCPU)].(int); ok && v > 0 {
				resLimits.SetMaxVCPU(spotinst.Int(v))
			}

		}
		return resLimits, nil
	}

	return nil, nil
}

func expandOceanGKEAutoScalerDown(data interface{}) (*gcp.AutoScalerDown, error) {
	if list := data.([]interface{}); len(list) > 0 {
		autoScaleDown := &gcp.AutoScalerDown{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(EvaluationPeriods)].(int); ok && v > 0 {
				autoScaleDown.SetEvaluationPeriods(spotinst.Int(v))
			}
		}
		return autoScaleDown, nil
	}

	return nil, nil
}
