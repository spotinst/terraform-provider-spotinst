package ocean_aks_np_virtual_node_group_vm_sizes

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Filters] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroupVmSizes,
		Filters,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MinVcpu): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(MaxVcpu): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(MinMemoryGiB): {
						Type:     schema.TypeFloat,
						Optional: true,
						Default:  -1,
					},

					string(MaxMemoryGiB): {
						Type:     schema.TypeFloat,
						Optional: true,
						Default:  -1,
					},

					string(Series): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(Architectures): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(ExcludeSeries): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(AcceleratedNetworking): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(DiskPerformance): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(MinGpu): {
						Type:     schema.TypeFloat,
						Optional: true,
						Default:  -1,
					},

					string(MaxGpu): {
						Type:     schema.TypeFloat,
						Optional: true,
						Default:  -1,
					},

					string(MinNICs): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(MinDisk): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(VmTypes): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(GpuTypes): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var result []interface{} = nil

			if virtualNodeGroup != nil && virtualNodeGroup.VmSizes != nil &&
				virtualNodeGroup.VmSizes.Filters != nil {
				result = flattenFilters(virtualNodeGroup.VmSizes.Filters)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Filters), result); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(Filters), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.GetOk(string(Filters)); ok {
				if filters, err := expandFilters(v, false); err != nil {
					return err
				} else {
					virtualNodeGroup.VmSizes.SetFilters(filters)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var value *azure_np.Filters = nil

			if v, ok := resourceData.GetOk(string(Filters)); ok {
				if filters, err := expandFilters(v, true); err != nil {
					return err
				} else {
					value = filters
				}
			}
			if virtualNodeGroup.VmSizes == nil {
				virtualNodeGroup.VmSizes = &azure_np.VmSizes{}
			}
			virtualNodeGroup.VmSizes.SetFilters(value)
			return nil
		},
		nil,
	)
}

func expandFilters(data interface{}, nullify bool) (*azure_np.Filters, error) {
	filters := &azure_np.Filters{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return filters, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Architectures)]; ok {
		architectures, err := expandVmSizesFiltersList(v)
		if err != nil {
			return nil, err
		}
		if architectures != nil && len(architectures) > 0 {
			filters.SetArchitectures(architectures)
		} else {
			if nullify {
				filters.SetArchitectures(nil)
			}
		}
	}

	if v, ok := m[string(Series)]; ok {
		series, err := expandVmSizesFiltersList(v)
		if err != nil {
			return nil, err
		}
		if series != nil && len(series) > 0 {
			filters.SetSeries(series)
		} else {
			if nullify {
				filters.SetSeries(nil)
			}
		}
	}

	if v, ok := m[string(ExcludeSeries)]; ok {
		excludeSeries, err := expandVmSizesFiltersList(v)
		if err != nil {
			return nil, err
		}
		if excludeSeries != nil && len(excludeSeries) > 0 {
			filters.SetExcludeSeries(excludeSeries)
		} else {
			if nullify {
				filters.SetExcludeSeries(nil)
			}
		}
	}

	if v, ok := m[string(MaxMemoryGiB)].(float64); ok {
		if v == -1 {
			filters.SetMaxMemoryGiB(nil)
		} else {
			filters.SetMaxMemoryGiB(spotinst.Float64(v))
		}
	}

	if v, ok := m[string(MaxVcpu)].(int); ok {
		if v == -1 {
			filters.SetMaxVcpu(nil)
		} else {
			filters.SetMaxVcpu(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MinMemoryGiB)].(float64); ok {
		if v == -1 {
			filters.SetMinMemoryGiB(nil)
		} else {
			filters.SetMinMemoryGiB(spotinst.Float64(v))
		}
	}

	if v, ok := m[string(MinVcpu)].(int); ok {
		if v == -1 {
			filters.SetMinVcpu(nil)
		} else {
			filters.SetMinVcpu(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MinGpu)].(float64); ok {
		if v == -1 {
			filters.SetMinGpu(nil)
		} else {
			filters.SetMinGpu(spotinst.Float64(v))
		}
	}

	if v, ok := m[string(MaxGpu)].(float64); ok {
		if v == -1 {
			filters.SetMaxGpu(nil)
		} else {
			filters.SetMaxGpu(spotinst.Float64(v))
		}
	}

	if v, ok := m[string(MinNICs)].(int); ok {
		if v == -1 {
			filters.SetMinNICs(nil)
		} else {
			filters.SetMinNICs(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MinDisk)].(int); ok {
		if v == -1 {
			filters.SetMinDisk(nil)
		} else {
			filters.SetMinDisk(spotinst.Int(v))
		}
	}

	if v, ok := m[string(AcceleratedNetworking)].(string); ok && v != "" {
		filters.SetAcceleratedNetworking(spotinst.String(v))
	} else {
		filters.SetAcceleratedNetworking(nil)
	}

	if v, ok := m[string(DiskPerformance)].(string); ok && v != "" {
		filters.SetDiskPerformance(spotinst.String(v))
	} else {
		filters.SetDiskPerformance(nil)
	}

	if v, ok := m[string(VmTypes)]; ok {
		vmTypes, err := expandVmSizesFiltersList(v)
		if err != nil {
			return nil, err
		}
		if vmTypes != nil && len(vmTypes) > 0 {
			filters.SetVmTypes(vmTypes)
		} else {
			if nullify {
				filters.SetVmTypes(nil)
			}
		}
	}

	if v, ok := m[string(GpuTypes)]; ok {
		gpuTypes, err := expandVmSizesFiltersList(v)
		if err != nil {
			return nil, err
		}
		if gpuTypes != nil && len(gpuTypes) > 0 {
			filters.SetGpuTypes(gpuTypes)
		} else {
			if nullify {
				filters.SetGpuTypes(nil)
			}
		}
	}

	return filters, nil
}

func expandVmSizesFiltersList(data interface{}) ([]string, error) {
	list := data.(*schema.Set).List()
	result := make([]string, 0, len(list))

	for _, v := range list {
		if vmSizeList, ok := v.(string); ok && vmSizeList != "" {
			result = append(result, vmSizeList)
		}
	}
	return result, nil
}

func flattenFilters(filters *azure_np.Filters) []interface{} {
	var out []interface{}

	if filters != nil {
		result := make(map[string]interface{})
		value := spotinst.Int(-1)
		result[string(MinVcpu)] = value
		result[string(MaxVcpu)] = value
		result[string(MinMemoryGiB)] = value
		result[string(MaxMemoryGiB)] = value
		result[string(MinGpu)] = value
		result[string(MaxGpu)] = value
		result[string(MinNICs)] = value
		result[string(MinDisk)] = value

		if filters.MinVcpu != nil {
			result[string(MinVcpu)] = spotinst.IntValue(filters.MinVcpu)
		}
		if filters.MaxVcpu != nil {
			result[string(MaxVcpu)] = spotinst.IntValue(filters.MaxVcpu)
		}
		if filters.MinMemoryGiB != nil {
			result[string(MinMemoryGiB)] = spotinst.Float64Value(filters.MinMemoryGiB)
		}
		if filters.MaxMemoryGiB != nil {
			result[string(MaxMemoryGiB)] = spotinst.Float64Value(filters.MaxMemoryGiB)
		}
		if filters.MinGpu != nil {
			result[string(MinGpu)] = spotinst.Float64Value(filters.MinGpu)
		}
		if filters.MaxGpu != nil {
			result[string(MaxGpu)] = spotinst.Float64Value(filters.MaxGpu)
		}
		if filters.MinNICs != nil {
			result[string(MinNICs)] = spotinst.IntValue(filters.MinNICs)
		}
		if filters.MinDisk != nil {
			result[string(MinDisk)] = spotinst.IntValue(filters.MinDisk)
		}
		if filters.AcceleratedNetworking != nil {
			result[string(AcceleratedNetworking)] = spotinst.StringValue(filters.AcceleratedNetworking)
		}
		if filters.DiskPerformance != nil {
			result[string(DiskPerformance)] = spotinst.StringValue(filters.DiskPerformance)
		}

		if filters.Architectures != nil {
			result[string(Architectures)] = filters.Architectures
		}

		if filters.Series != nil {
			result[string(Series)] = filters.Series
		}

		if filters.ExcludeSeries != nil {
			result[string(ExcludeSeries)] = filters.ExcludeSeries
		}

		if filters.VmTypes != nil {
			result[string(VmTypes)] = filters.VmTypes
		}

		if filters.GpuTypes != nil {
			result[string(GpuTypes)] = filters.GpuTypes
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}
