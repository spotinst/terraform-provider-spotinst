package ocean_ecs_autoscaler

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
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
									Type:     schema.TypeInt,
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

			if v, ok := m[string(MaxScaleDownPercentage)].(int); ok && v > 0 {
				autoScaleDown.SetMaxScaleDownPercentage(spotinst.Int(v))
			}
		}
		return autoScaleDown, nil
	}

	return nil, nil
}
