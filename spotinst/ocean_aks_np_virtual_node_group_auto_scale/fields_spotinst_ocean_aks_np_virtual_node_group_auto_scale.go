package ocean_aks_np_virtual_node_group_auto_scale

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Headrooms] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroupAutoScale,
		Headrooms,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(CpuPerUnit): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(MemoryPerUnit): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(GpuPerUnit): {
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
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()

			var headroomsResults []interface{} = nil
			if virtualNodeGroup != nil && virtualNodeGroup.AutoScale != nil && virtualNodeGroup.AutoScale.Headrooms != nil {
				headrooms := virtualNodeGroup.AutoScale.Headrooms
				headroomsResults = flattenHeadrooms(headrooms)
			}

			if err := resourceData.Set(string(Headrooms), headroomsResults); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Headrooms), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()

			if value, ok := resourceData.GetOkExists(string(Headrooms)); ok {
				if headrooms, err := expandHeadrooms(value); err != nil {
					return err
				} else {
					virtualNodeGroup.AutoScale.SetHeadrooms(headrooms)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()

			var result []*azure_np.Headrooms = nil
			if value, ok := resourceData.GetOkExists(string(Headrooms)); ok {
				if headrooms, err := expandHeadrooms(value); err != nil {
					return err
				} else {
					result = headrooms
				}
			}

			if len(result) == 0 {
				virtualNodeGroup.AutoScale.SetHeadrooms(nil)
			} else {
				virtualNodeGroup.AutoScale.SetHeadrooms(result)
			}

			return nil
		},
		nil,
	)
}

func expandHeadrooms(data interface{}) ([]*azure_np.Headrooms, error) {
	list := data.(*schema.Set).List()
	headrooms := make([]*azure_np.Headrooms, 0, len(list))

	for _, v := range list {
		m := v.(map[string]interface{})
		headroom := &azure_np.Headrooms{}

		if v, ok := m[string(CpuPerUnit)].(int); ok {
			if v == -1 {
				headroom.SetCpuPerUnit(nil)
			} else {
				headroom.SetCpuPerUnit(spotinst.Int(v))
			}
		}

		if v, ok := m[string(GpuPerUnit)].(int); ok {
			if v == -1 {
				headroom.SetGpuPerUnit(nil)
			} else {
				headroom.SetGpuPerUnit(spotinst.Int(v))
			}
		}

		if v, ok := m[string(NumOfUnits)].(int); ok {
			if v == -1 {
				headroom.SetNumOfUnits(nil)
			} else {
				headroom.SetNumOfUnits(spotinst.Int(v))
			}
		}

		if v, ok := m[string(MemoryPerUnit)].(int); ok {
			if v == -1 {
				headroom.SetMemoryPerUnit(nil)
			} else {
				headroom.SetMemoryPerUnit(spotinst.Int(v))
			}
		}

		headrooms = append(headrooms, headroom)
	}
	return headrooms, nil
}

func flattenHeadrooms(autoScaleHeadrooms []*azure_np.Headrooms) []interface{} {
	result := make([]interface{}, 0, len(autoScaleHeadrooms))

	for _, headroom := range autoScaleHeadrooms {
		m := make(map[string]interface{})
		value := spotinst.Int(-1)
		m[string(CpuPerUnit)] = value
		m[string(GpuPerUnit)] = value
		m[string(MemoryPerUnit)] = value
		m[string(NumOfUnits)] = value

		if headroom.CpuPerUnit != nil {
			m[string(CpuPerUnit)] = spotinst.IntValue(headroom.CpuPerUnit)
		}
		if headroom.GpuPerUnit != nil {
			m[string(GpuPerUnit)] = spotinst.IntValue(headroom.GpuPerUnit)
		}
		if headroom.MemoryPerUnit != nil {
			m[string(MemoryPerUnit)] = spotinst.IntValue(headroom.MemoryPerUnit)
		}
		if headroom.NumberOfUnits != nil {
			m[string(NumOfUnits)] = spotinst.IntValue(headroom.NumberOfUnits)
		}

		result = append(result, m)
	}
	return result
}
