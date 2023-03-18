package ocean_aks_np_auto_scale

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Headrooms] = commons.NewGenericField(
		commons.OceanAKSNPGroupAutoScale,
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
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()

			var headroomsResults []interface{} = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate.AutoScale != nil && cluster.VirtualNodeGroupTemplate.AutoScale.Headrooms != nil {
				headrooms := cluster.VirtualNodeGroupTemplate.AutoScale.Headrooms
				headroomsResults = flattenHeadrooms(headrooms)
			}

			if err := resourceData.Set(string(Headrooms), headroomsResults); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Headrooms), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()

			if value, ok := resourceData.GetOkExists(string(Headrooms)); ok {
				if headrooms, err := expandHeadrooms(value); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.AutoScale.SetHeadrooms(headrooms)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()

			var result []*azure_np.Headrooms = nil
			if value, ok := resourceData.GetOkExists(string(Headrooms)); ok {
				if headrooms, err := expandHeadrooms(value); err != nil {
					return err
				} else {
					result = headrooms
				}
			}

			cluster.VirtualNodeGroupTemplate.AutoScale.SetHeadrooms(result)
			return nil
		},
		nil,
	)
}

func expandHeadrooms(headroom interface{}) ([]*azure_np.Headrooms, error) {
	list := headroom.(*schema.Set).List()
	headrooms := make([]*azure_np.Headrooms, 0, len(list))

	for _, v := range list {
		m, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		headrooms = append(headrooms, &azure_np.Headrooms{
			CpuPerUnit:    spotinst.Int(m[string(CpuPerUnit)].(int)),
			GpuPerUnit:    spotinst.Int(m[string(GpuPerUnit)].(int)),
			NumberOfUnits: spotinst.Int(m[string(NumOfUnits)].(int)),
			MemoryPerUnit: spotinst.Int(m[string(MemoryPerUnit)].(int)),
		})
	}
	return headrooms, nil
}

func flattenHeadrooms(autoScaleHeadrooms []*azure_np.Headrooms) []interface{} {
	result := make([]interface{}, 0, len(autoScaleHeadrooms))

	for _, headroom := range autoScaleHeadrooms {
		m := make(map[string]interface{})
		m[string(CpuPerUnit)] = spotinst.IntValue(headroom.CpuPerUnit)
		m[string(GpuPerUnit)] = spotinst.IntValue(headroom.GpuPerUnit)
		m[string(MemoryPerUnit)] = spotinst.IntValue(headroom.MemoryPerUnit)
		m[string(NumOfUnits)] = spotinst.IntValue(headroom.NumberOfUnits)
		result = append(result, m)
	}
	return result
}
