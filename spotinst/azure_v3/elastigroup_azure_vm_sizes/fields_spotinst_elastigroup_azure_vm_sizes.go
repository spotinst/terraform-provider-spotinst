package elastigroup_azure_vm_sizes

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
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
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup != nil && elastigroup.Compute != nil && elastigroup.Compute.VMSizes != nil {
				vmSizes := elastigroup.Compute.VMSizes
				result = flattenAzureGroupSizes(vmSizes)
			}

			if result != nil {
				if err := resourceData.Set(string(VmSizes), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(VmSizes), err)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(VmSizes)); ok {
				if vmSizes, err := expandAzureGroupSizes(v); err != nil {
					return err
				} else {
					elastigroup.Compute.SetVMSizes(vmSizes)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *azurev3.VMSizes = nil
			if v, ok := resourceData.GetOk(string(VmSizes)); ok {
				if vmSizes, err := expandAzureGroupSizes(v); err != nil {
					return err
				} else {
					value = vmSizes
				}

			}
			elastigroup.Compute.SetVMSizes(value)
			return nil
		},
		nil,
	)
}

func flattenAzureGroupSizes(vmSizes *azurev3.VMSizes) []interface{} {
	result := make(map[string]interface{})
	if len(vmSizes.OnDemandSizes) > 0 {
		result[string(OnDemandSizes)] = vmSizes.OnDemandSizes
	}
	if len(vmSizes.SpotSizes) > 0 {
		result[string(SpotSizes)] = vmSizes.SpotSizes
	}
	return []interface{}{result}
}

func expandAzureGroupSizes(data interface{}) (*azurev3.VMSizes, error) {
	vmSizes := &azurev3.VMSizes{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(OnDemandSizes)]; ok && v != nil {
			onDemandSizes, err := expandOnDemandSizes(v)
			if err != nil {
				return nil, err
			}
			if onDemandSizes != nil {
				vmSizes.SetOnDemandSizes(onDemandSizes)
			}
		}

		if v, ok := m[string(SpotSizes)]; ok && v != nil {
			spotSizes, err := expandSpotSizes(v)
			if err != nil {
				return nil, err
			}
			if spotSizes != nil {
				vmSizes.SetSpotSizes(spotSizes)
			}
		}

	}
	return vmSizes, nil
}

func expandOnDemandSizes(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if odSizes, ok := v.(string); ok {
			result = append(result, odSizes)
		}
	}
	return result, nil
}

func expandSpotSizes(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if spotSizes, ok := v.(string); ok {
			result = append(result, spotSizes)
		}
	}
	return result, nil
}
