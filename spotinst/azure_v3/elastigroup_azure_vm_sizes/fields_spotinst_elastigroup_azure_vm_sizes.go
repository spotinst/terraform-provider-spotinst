package elastigroup_azure_vm_sizes

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

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

					string(PreferredSpotSizes): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},

					string(ExcludedVmSizes): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},

					string(SpotSizeAttributes): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{

								string(CPUArchitecture): {
									Type:     schema.TypeString,
									Optional: true,
								},

								string(MaxMemory): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  -1,
								},
								string(MaxCpu): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  -1,
								},
								string(MaxStorage): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  -1,
								},
								string(MinMemory): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  -1,
								},
								string(MinCpu): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  -1,
								},
								string(MinStorage): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  -1,
								},
							},
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
	if len(vmSizes.PreferredSpotSizes) > 0 {
		result[string(PreferredSpotSizes)] = vmSizes.PreferredSpotSizes
	}

	if len(vmSizes.ExcludedVmSizes) > 0 {
		result[string(ExcludedVmSizes)] = vmSizes.ExcludedVmSizes
	}

	if vmSizes.SpotSizeAttributes != nil {
		result[string(SpotSizeAttributes)] = flattenSpotSizeAttributes(vmSizes.SpotSizeAttributes)
	}

	return []interface{}{result}
}

func expandAzureGroupSizes(data interface{}) (*azurev3.VMSizes, error) {
	vmSizes := &azurev3.VMSizes{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(OnDemandSizes)]; ok && v != nil {
			onDemandSizes, err := expandSizes(v)
			if err != nil {
				return nil, err
			}
			if onDemandSizes != nil {
				vmSizes.SetOnDemandSizes(onDemandSizes)
			}
		}

		if v, ok := m[string(SpotSizes)]; ok && v != nil {
			spotSizes, err := expandSizes(v)
			if err != nil {
				return nil, err
			}
			if spotSizes != nil {
				vmSizes.SetSpotSizes(spotSizes)
			}
		}

		if v, ok := m[string(PreferredSpotSizes)]; ok && v != nil {
			preferredSpotSizes, err := expandSizes(v)
			if err != nil {
				return nil, err
			}
			if preferredSpotSizes != nil {
				vmSizes.SetPreferredSpotSizes(preferredSpotSizes)
			}
		}

		if v, ok := m[string(ExcludedVmSizes)]; ok {
			excludedVmSizes, err := expandSizes(v)
			if err != nil {
				return nil, err
			}

			if excludedVmSizes != nil && len(excludedVmSizes) > 0 {
				vmSizes.SetExcludedVmSizes(excludedVmSizes)
			} else {
				vmSizes.SetExcludedVmSizes(nil)
			}
		}

		if v, ok := m[string(SpotSizeAttributes)]; ok && v != nil {

			spotSizeAttributes, err := expandSpotSizeAttributes(v)
			if err != nil {
				return nil, err
			}
			if spotSizeAttributes != nil {
				vmSizes.SetSpotSizeAttributes(spotSizeAttributes)
			} else {
				vmSizes.SetSpotSizeAttributes(nil)
			}
		}

	}
	return vmSizes, nil
}

func expandSizes(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if odSizes, ok := v.(string); ok {
			result = append(result, odSizes)
		}
	}
	return result, nil
}

func expandSpotSizeAttributes(data interface{}) (*azurev3.SpotSizeAttributes, error) {
	spotSizeAttributes := &azurev3.SpotSizeAttributes{}
	list := data.([]interface{})

	if list == nil || len(list) == 0 {
		return nil, nil
	}

	m := list[0].(map[string]interface{})

	if v, ok := m[string(MaxMemory)].(int); ok {
		// here -1 is used to set MaxMemoryGib field to null when the customer doesn't want to set this param,
		//as terraform set it 0 for integer type param by default.
		if v == -1 {
			spotSizeAttributes.SetMaxMemory(nil)
		} else {
			spotSizeAttributes.SetMaxMemory(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MaxCpu)].(int); ok {
		//Here -1 is used to set MaxVCPU field to null when the customer doesn't want to set this param,
		//as terraform set it 0 for integer type param by default.
		if v == -1 {
			spotSizeAttributes.SetMaxCpu(nil)
		} else {
			spotSizeAttributes.SetMaxCpu(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MaxStorage)].(int); ok {
		//Here -1 is used to set MaxVCPU field to null when the customer doesn't want to set this param,
		//as terraform set it 0 for integer type param by default.
		if v == -1 {
			spotSizeAttributes.SetMaxStorage(nil)
		} else {
			spotSizeAttributes.SetMaxStorage(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MinMemory)].(int); ok {
		// here -1 is used to set MaxMemoryGib field to null when the customer doesn't want to set this param,
		//as terraform set it 0 for integer type param by default.
		if v == -1 {
			spotSizeAttributes.SetMinMemory(nil)
		} else {
			spotSizeAttributes.SetMinMemory(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MinCpu)].(int); ok {
		//Here -1 is used to set MaxVCPU field to null when the customer doesn't want to set this param,
		//as terraform set it 0 for integer type param by default.
		if v == -1 {
			spotSizeAttributes.SetMinCpu(nil)
		} else {
			spotSizeAttributes.SetMinCpu(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MinStorage)].(int); ok {
		//Here -1 is used to set MaxVCPU field to null when the customer doesn't want to set this param,
		//as terraform set it 0 for integer type param by default.
		if v == -1 {
			spotSizeAttributes.SetMinStorage(nil)
		} else {
			spotSizeAttributes.SetMinStorage(spotinst.Int(v))
		}
	}

	if v, ok := m[string(CPUArchitecture)].(string); ok && v != "" {
		spotSizeAttributes.SetCPUArchitecture(spotinst.String(v))
	}
	return spotSizeAttributes, nil
}

func flattenSpotSizeAttributes(spotSizeAttributes *azurev3.SpotSizeAttributes) []interface{} {
	spotAttributes := make(map[string]interface{})
	if spotSizeAttributes != nil {
		value := spotinst.Int(-1)
		spotAttributes[string(MaxCpu)] = value
		spotAttributes[string(MaxMemory)] = value
		spotAttributes[string(MaxStorage)] = value
		spotAttributes[string(MinCpu)] = value
		spotAttributes[string(MinMemory)] = value
		spotAttributes[string(MinStorage)] = value

		if spotSizeAttributes.MaxCpu != nil {
			spotAttributes[string(MaxCpu)] = spotinst.IntValue(spotSizeAttributes.MaxCpu)
		}
		if spotSizeAttributes.MaxMemory != nil {
			spotAttributes[string(MaxMemory)] = spotinst.IntValue(spotSizeAttributes.MaxMemory)
		}
		if spotSizeAttributes.MaxStorage != nil {
			spotAttributes[string(MaxStorage)] = spotinst.IntValue(spotSizeAttributes.MaxStorage)
		}
		if spotSizeAttributes.MinCpu != nil {
			spotAttributes[string(MinCpu)] = spotinst.IntValue(spotSizeAttributes.MinCpu)
		}
		if spotSizeAttributes.MinMemory != nil {
			spotAttributes[string(MinMemory)] = spotinst.IntValue(spotSizeAttributes.MinMemory)
		}
		if spotSizeAttributes.MinStorage != nil {
			spotAttributes[string(MinStorage)] = spotinst.IntValue(spotSizeAttributes.MinStorage)
		}
		if spotSizeAttributes.CPUArchitecture != nil {
			spotAttributes[string(CPUArchitecture)] = spotinst.StringValue(spotSizeAttributes.CPUArchitecture)
		}
	}

	return []interface{}{spotAttributes}
}
