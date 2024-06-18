package stateful_node_azure_vm_sizes

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[VmSizes] = commons.NewGenericField(
		commons.ElastigroupAzureVMSizes,
		VmSizes,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(OnDemandSizes): {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},

					string(SpotSizes): {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},

					string(PreferredSpotSizes): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var result []interface{} = nil
			if statefulNode != nil && statefulNode.Compute != nil && statefulNode.Compute.VMSizes != nil {
				vmSizes := statefulNode.Compute.VMSizes
				result = flattenStatefulNodeAzureVmSizes(vmSizes)
			}

			if result != nil {
				if err := resourceData.Set(string(VmSizes), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(VmSizes), err)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(VmSizes)); ok {
				if vmSizes, err := expandStatefulNodeAzureVmSizes(v); err != nil {
					return err
				} else {
					statefulNode.Compute.SetVMSizes(vmSizes)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *azure.VMSizes = nil
			if v, ok := resourceData.GetOk(string(VmSizes)); ok {
				if vmSizes, err := expandStatefulNodeAzureVmSizes(v); err != nil {
					return err
				} else {
					value = vmSizes
				}

			}
			statefulNode.Compute.SetVMSizes(value)
			return nil
		},
		nil,
	)
}

func flattenStatefulNodeAzureVmSizes(vmSizes *azure.VMSizes) []interface{} {
	result := make(map[string]interface{})

	if vmSizes.OnDemandSizes != nil {
		result[string(OnDemandSizes)] = spotinst.StringSlice(vmSizes.OnDemandSizes)
	}

	if vmSizes.SpotSizes != nil {
		result[string(SpotSizes)] = spotinst.StringSlice(vmSizes.SpotSizes)
	}

	if vmSizes.PreferredSpotSizes != nil {
		result[string(PreferredSpotSizes)] = spotinst.StringSlice(vmSizes.PreferredSpotSizes)
	}

	return []interface{}{result}
}

func expandStatefulNodeAzureVmSizes(data interface{}) (*azure.VMSizes, error) {
	vmSizes := &azure.VMSizes{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(OnDemandSizes)]; ok && v != nil {
			onDemandSizes, err := expandStatefulNodeAzureSizes(v)
			if err != nil {
				return nil, err
			}
			if onDemandSizes != nil {
				vmSizes.SetOnDemandSizes(onDemandSizes)
			}
		}

		if v, ok := m[string(SpotSizes)]; ok && v != nil {
			spotSizes, err := expandStatefulNodeAzureSizes(v)
			if err != nil {
				return nil, err
			}
			if spotSizes != nil {
				vmSizes.SetSpotSizes(spotSizes)
			}
		}

		if v, ok := m[string(PreferredSpotSizes)]; ok {
			prefferedSpotSizes, err := expandStatefulNodeAzureSizes(v)
			if err != nil {
				return nil, err
			}

			if prefferedSpotSizes != nil && len(prefferedSpotSizes) > 0 {
				vmSizes.SetPreferredSpotSizes(prefferedSpotSizes)
			} else {
				vmSizes.SetPreferredSpotSizes(nil)
			}
		}
	}
	return vmSizes, nil
}

func expandStatefulNodeAzureSizes(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if sizes, ok := v.(string); ok {
			result = append(result, sizes)
		}
	}
	return result, nil
}
