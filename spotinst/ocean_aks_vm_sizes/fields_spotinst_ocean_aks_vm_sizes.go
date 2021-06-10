package ocean_aks_vm_sizes

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Whitelist] = commons.NewGenericField(
		commons.OceanAKSVMSizes,
		VMSizes,

		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Whitelist): {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result interface{}

			if cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.VMSizes != nil {
				vmSizes := cluster.VirtualNodeGroupTemplate.VMSizes
				result = flattenVMSizes(vmSizes)
			}
			if result != nil {
				if err := resourceData.Set(string(VMSizes), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(VMSizes), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *azure.VMSizes = nil

			if v, ok := resourceData.GetOk(string(VMSizes)); ok {
				if vmSizes, err := expandVMSizes(v); err != nil {
					return err
				} else {
					value = vmSizes
				}
			}
			if cluster.VirtualNodeGroupTemplate != nil {
				cluster.VirtualNodeGroupTemplate.SetVMSizes(value)

			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *azure.VMSizes = nil

			if v, ok := resourceData.GetOk(string(VMSizes)); ok {
				if vmSizes, err := expandVMSizes(v); err != nil {
					return err
				} else {
					value = vmSizes
				}
			}
			if cluster.VirtualNodeGroupTemplate != nil {
				cluster.VirtualNodeGroupTemplate.SetVMSizes(value)

			}
			return nil
		},
		nil,
	)
}

func expandVMSizes(data interface{}) (*azure.VMSizes, error) {
	if list := data.([]interface{}); len(list) > 0 {
		vmSizes := &azure.VMSizes{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(Whitelist)]; ok {
				if whitelist, err := expandWhitelist(v); err != nil {
					return nil, err
				} else {
					vmSizes.SetWhitelist(whitelist)
				}
			}
		}

		return vmSizes, nil
	}

	return nil, nil
}

func expandWhitelist(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if whitelist, ok := v.(string); ok && whitelist != "" {
			result = append(result, whitelist)
		}
	}

	return result, nil
}

func flattenVMSizes(vmSizes *azure.VMSizes) []interface{} {
	var out []interface{}

	if vmSizes != nil {
		result := make(map[string]interface{})

		if vmSizes.Whitelist != nil {
			result[string(Whitelist)] = spotinst.StringSlice(vmSizes.Whitelist)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}
