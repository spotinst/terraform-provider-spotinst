package ocean_aks_virtual_node_group_auto_scaling

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Autoscale] = commons.NewGenericField(
		commons.OceanAKSVirtualNodeGroupAutoScaling,
		Autoscale,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Headrooms): {
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
									Required: true,
								},
							},
						},
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			cluster := clusterWrapper.GetVirtualNodeGroup()
			var result []interface{} = nil
			if cluster != nil && cluster.AutoScale != nil {
				result = flattenAutoScale(cluster.AutoScale)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Autoscale), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Autoscale), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			cluster := clusterWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.GetOk(string(Autoscale)); ok {
				if autoscaler, err := expandAutoScale(v, false); err != nil {
					return err
				} else {
					cluster.SetAutoScale(autoscaler)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			cluster := clusterWrapper.GetVirtualNodeGroup()
			var value *azure.VirtualNodeGroupAutoScale = nil
			if v, ok := resourceData.GetOk(string(Autoscale)); ok {
				if autoscale, err := expandAutoScale(v, true); err != nil {
					return err
				} else {
					value = autoscale
				}
			}
			cluster.SetAutoScale(value)
			return nil
		},

		nil,
	)
}

func expandAutoScale(data interface{}, nullify bool) (*azure.VirtualNodeGroupAutoScale, error) {
	autoscale := &azure.VirtualNodeGroupAutoScale{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return autoscale, nil
	}

	m := list[0].(map[string]interface{})

	if v, ok := m[string(Headrooms)]; ok {
		headroom, err := expandHeadrooms(v, nullify)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			autoscale.SetHeadrooms(headroom)
		} else {
			autoscale.Headrooms = nil
		}
	}

	return autoscale, nil
}

func flattenAutoScale(autoScale *azure.VirtualNodeGroupAutoScale) []interface{} {
	var out []interface{}

	if autoScale != nil {
		result := make(map[string]interface{})
		if autoScale.Headrooms != nil {
			result[string(Headrooms)] = flattenHeadrooms(autoScale.Headrooms)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func expandHeadrooms(data interface{}, nullify bool) ([]*azure.VirtualNodeGroupHeadroom, error) {
	list := data.([]interface{})
	headrooms := make([]*azure.VirtualNodeGroupHeadroom, 0, len(list))

	for _, v := range list {
		m, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		headroom := &azure.VirtualNodeGroupHeadroom{}

		if v, ok := m[string(CPUPerUnit)].(int); ok && v > 0 {
			headroom.SetCPUPerUnit(spotinst.Int(v))
		} else if nullify {
			headroom.SetCPUPerUnit(nil)
		}

		if v, ok := m[string(MemoryPerUnit)].(int); ok && v > 0 {
			headroom.SetMemoryPerUnit(spotinst.Int(v))
		} else if nullify {
			headroom.SetMemoryPerUnit(nil)
		}

		if v, ok := m[string(GPUPerUnit)].(int); ok && v > 0 {
			headroom.SetGPUPerUnit(spotinst.Int(v))
		} else if nullify {
			headroom.SetGPUPerUnit(nil)
		}

		if v, ok := m[string(NumOfUnits)].(int); ok && v > 0 {
			headroom.SetNumOfUnits(spotinst.Int(v))
		} else if nullify {
			headroom.SetNumOfUnits(nil)
		}

		headrooms = append(headrooms, headroom)
	}

	return headrooms, nil

}

func flattenHeadrooms(autoScaleHeadrooms []*azure.VirtualNodeGroupHeadroom) []interface{} {
	result := make([]interface{}, 0, len(autoScaleHeadrooms))

	for _, headroom := range autoScaleHeadrooms {
		m := make(map[string]interface{})
		m[string(CPUPerUnit)] = spotinst.IntValue(headroom.CPUPerUnit)
		m[string(GPUPerUnit)] = spotinst.IntValue(headroom.GPUPerUnit)
		m[string(MemoryPerUnit)] = spotinst.IntValue(headroom.MemoryPerUnit)
		m[string(NumOfUnits)] = spotinst.IntValue(headroom.NumOfUnits)
		result = append(result, m)
	}
	return result

}
