package ocean_gke_instance_types

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Whitelist] = commons.NewGenericField(
		commons.OceanGKEInstanceTypes,
		Whitelist,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []string
			if cluster.Compute != nil && cluster.Compute.InstanceTypes != nil &&
				cluster.Compute.InstanceTypes.Whitelist != nil {
				result = append(result, cluster.Compute.InstanceTypes.Whitelist...)
			}
			if err := resourceData.Set(string(Whitelist), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Whitelist), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(Whitelist)); ok {
				instances := v.([]interface{})
				instanceTypes := make([]string, len(instances))
				for i, j := range instances {
					instanceTypes[i] = j.(string)
				}
				cluster.Compute.InstanceTypes.SetWhitelist(instanceTypes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var instanceTypes []string = nil
			if v, ok := resourceData.GetOk(string(Whitelist)); ok {
				instances := v.([]interface{})
				instanceTypes = make([]string, len(instances))
				for i, v := range instances {
					instanceTypes[i] = v.(string)
				}
			}
			cluster.Compute.InstanceTypes.SetWhitelist(instanceTypes)
			return nil
		},
		nil,
	)

	fieldsMap[Blacklist] = commons.NewGenericField(
		commons.OceanGKEInstanceTypes,
		Blacklist,
		&schema.Schema{
			Type:          schema.TypeList,
			Optional:      true,
			Elem:          &schema.Schema{Type: schema.TypeString},
			ConflictsWith: []string{string(Whitelist)},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []string
			if cluster.Compute != nil && cluster.Compute.InstanceTypes != nil &&
				cluster.Compute.InstanceTypes.Blacklist != nil {
				result = cluster.Compute.InstanceTypes.Blacklist
			}
			if err := resourceData.Set(string(Blacklist), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Blacklist), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(Blacklist)); ok {
				if blacklist, err := expandInstanceTypeList(v); err != nil {
					return err
				} else {
					cluster.Compute.InstanceTypes.SetBlacklist(blacklist)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(Blacklist)); ok {
				if blacklist, err := expandInstanceTypeList(v); err != nil {
					return err
				} else {
					cluster.Compute.InstanceTypes.SetBlacklist(blacklist)
				}
			} else {
				cluster.Compute.InstanceTypes.SetBlacklist(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[PreferredTypes] = commons.NewGenericField(
		commons.OceanGKEInstanceTypes,
		PreferredTypes,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []string
			if cluster.Compute != nil && cluster.Compute.InstanceTypes != nil &&
				cluster.Compute.InstanceTypes.PreferredTypes != nil {
				result = cluster.Compute.InstanceTypes.PreferredTypes
			}
			if err := resourceData.Set(string(PreferredTypes), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PreferredTypes), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(PreferredTypes)); ok {
				if preferredTypes, err := expandInstanceTypeList(v); err != nil {
					return err
				} else {
					cluster.Compute.InstanceTypes.SetPreferredTypes(preferredTypes)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(PreferredTypes)); ok {
				if preferredTypes, err := expandInstanceTypeList(v); err != nil {
					return err
				} else {
					cluster.Compute.InstanceTypes.SetPreferredTypes(preferredTypes)
				}
			} else {
				cluster.Compute.InstanceTypes.SetPreferredTypes(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[Filters] = commons.NewGenericField(
		commons.OceanGKEInstanceTypes,
		Filters,
		&schema.Schema{
			Type:          schema.TypeList,
			Optional:      true,
			MaxItems:      1,
			ConflictsWith: []string{string(Blacklist), string(Whitelist)},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(ExcludeFamilies): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(IncludeFamilies): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(MaxMemoryGiB): {
						Type:     schema.TypeFloat,
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

					string(MinVcpu): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil

			if cluster != nil && cluster.Compute != nil && cluster.Compute.InstanceTypes != nil &&
				cluster.Compute.InstanceTypes.Filters != nil {
				result = flattenFilters(cluster.Compute.InstanceTypes.Filters)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Filters), result); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(Filters), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(Filters)); ok {
				if filters, err := expandFilters(v, false); err != nil {
					return err
				} else {
					cluster.Compute.InstanceTypes.SetFilters(filters)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *gcp.Filters = nil

			if v, ok := resourceData.GetOk(string(Filters)); ok {
				if filters, err := expandFilters(v, true); err != nil {
					return err
				} else {
					value = filters
				}
			}
			if cluster.Compute.InstanceTypes == nil {
				cluster.Compute.InstanceTypes = &gcp.InstanceTypes{}
			}
			cluster.Compute.InstanceTypes.SetFilters(value)
			return nil
		},
		nil,
	)

}

func expandFilters(data interface{}, nullify bool) (*gcp.Filters, error) {
	filters := &gcp.Filters{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return filters, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(ExcludeFamilies)]; ok {
		excludeFamilies, err := expandInstanceTypeFiltersList(v)
		if err != nil {
			return nil, err
		}
		if excludeFamilies != nil && len(excludeFamilies) > 0 {
			filters.SetExcludeFamilies(excludeFamilies)
		} else {
			if nullify {
				filters.SetExcludeFamilies(nil)
			}
		}
	}

	if v, ok := m[string(IncludeFamilies)]; ok {
		includeFamilies, err := expandInstanceTypeFiltersList(v)
		if err != nil {
			return nil, err
		}
		if includeFamilies != nil && len(includeFamilies) > 0 {
			filters.SetIncludeFamilies(includeFamilies)
		} else {
			if nullify {
				filters.SetIncludeFamilies(nil)
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

	return filters, nil
}

func expandInstanceTypeList(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if instanceTypeList, ok := v.(string); ok && instanceTypeList != "" {
			result = append(result, instanceTypeList)
		}
	}
	return result, nil
}

func expandInstanceTypeFiltersList(data interface{}) ([]string, error) {
	list := data.(*schema.Set).List()
	result := make([]string, 0, len(list))

	for _, v := range list {
		if instanceTypeList, ok := v.(string); ok && instanceTypeList != "" {
			result = append(result, instanceTypeList)
		}
	}
	return result, nil
}

func flattenFilters(filters *gcp.Filters) []interface{} {
	var out []interface{}

	if filters != nil {
		result := make(map[string]interface{})
		value := spotinst.Int(-1)
		result[string(MaxMemoryGiB)] = value
		result[string(MinMemoryGiB)] = value
		result[string(MaxVcpu)] = value
		result[string(MinVcpu)] = value

		if filters.MaxMemoryGiB != nil {
			result[string(MaxMemoryGiB)] = spotinst.Float64Value(filters.MaxMemoryGiB)
		}
		if filters.MinMemoryGiB != nil {
			result[string(MinMemoryGiB)] = spotinst.Float64Value(filters.MinMemoryGiB)
		}
		if filters.MaxVcpu != nil {
			result[string(MaxVcpu)] = spotinst.IntValue(filters.MaxVcpu)
		}
		if filters.MinVcpu != nil {
			result[string(MinVcpu)] = spotinst.IntValue(filters.MinVcpu)
		}
		if filters.ExcludeFamilies != nil {
			result[string(ExcludeFamilies)] = filters.ExcludeFamilies
		}
		if filters.IncludeFamilies != nil {
			result[string(IncludeFamilies)] = filters.IncludeFamilies
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}
