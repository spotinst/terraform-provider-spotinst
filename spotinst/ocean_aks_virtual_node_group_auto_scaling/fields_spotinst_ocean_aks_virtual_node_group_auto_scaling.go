package ocean_aks_virtual_node_group_auto_scaling

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
						Type:     schema.TypeSet,
						Optional: true,
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
		if len(headroom) > 0 {
			autoscale.SetHeadrooms(headroom)
		} else {
			autoscale.SetHeadrooms(nil)
		}
	}

	if v, ok := m[string(AutoHeadroomPercentage)].(int); ok && v > -1 {
		autoscale.SetAutoHeadroomPercentage(spotinst.Int(v))
	} else if nullify {
		autoscale.SetAutoHeadroomPercentage(nil)
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

		value := spotinst.Int(-1)
		if autoScale.AutoHeadroomPercentage != nil {
			value = autoScale.AutoHeadroomPercentage
		}
		result[string(AutoHeadroomPercentage)] = spotinst.IntValue(value)

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func expandHeadrooms(data interface{}, nullify bool) ([]*azure.VirtualNodeGroupHeadroom, error) {
	list := data.(*schema.Set).List()
	headrooms := make([]*azure.VirtualNodeGroupHeadroom, 0, len(list))

	for _, v := range list {
		m, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		headroom := &azure.VirtualNodeGroupHeadroom{
			CPUPerUnit:    spotinst.Int(m[string(CPUPerUnit)].(int)),
			GPUPerUnit:    spotinst.Int(m[string(GPUPerUnit)].(int)),
			NumOfUnits:    spotinst.Int(m[string(NumOfUnits)].(int)),
			MemoryPerUnit: spotinst.Int(m[string(MemoryPerUnit)].(int)),
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
