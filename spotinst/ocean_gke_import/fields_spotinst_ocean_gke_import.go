package ocean_gke_import

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Location] = commons.NewGenericField(
		commons.OceanGKEImport,
		Location,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(Location))
			return err
		},
		nil,
	)

	fieldsMap[ClusterName] = commons.NewGenericField(
		commons.OceanGKEImport,
		ClusterName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(ClusterName))
			return err
		},
		nil,
	)

	fieldsMap[Whitelist] = commons.NewGenericField(
		commons.OceanGKEImport,
		Whitelist,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
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
		commons.OceanGKEImport,
		Blacklist,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []string
			if cluster.Compute != nil && cluster.Compute.InstanceTypes != nil &&
				cluster.Compute.InstanceTypes.Blacklist != nil {
				result = append(result, cluster.Compute.InstanceTypes.Blacklist...)
			}
			if err := resourceData.Set(string(Blacklist), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Blacklist), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var instanceTypes []string = nil
			if v, ok := resourceData.GetOk(string(Blacklist)); ok {
				instances := v.([]interface{})
				instanceTypes = make([]string, len(instances))
				for i, v := range instances {
					instanceTypes[i] = v.(string)
				}
				cluster.Compute.InstanceTypes.SetBlacklist(instanceTypes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var instanceTypes []string = nil
			if v, ok := resourceData.GetOk(string(Blacklist)); ok {
				instances := v.([]interface{})
				instanceTypes = make([]string, len(instances))
				for i, v := range instances {
					instanceTypes[i] = v.(string)
				}
			}
			cluster.Compute.InstanceTypes.SetBlacklist(instanceTypes)
			return nil
		},
		nil,
	)

	fieldsMap[BackendServices] = commons.NewGenericField(
		commons.OceanGKEImport,
		BackendServices,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ServiceName): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(LocationType): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(Scheme): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(NamedPorts): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Name): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(Ports): {
									Type:     schema.TypeList,
									Required: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(BackendServices)); ok {
				if services, err := expandServices(v); err != nil {
					return err
				} else {
					cluster.Compute.SetBackendServices(services)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value []*gcp.BackendService = nil
			if v, ok := resourceData.GetOk(string(BackendServices)); ok {
				if services, err := expandServices(v); err != nil {
					return err
				} else {
					value = services
				}
			}
			cluster.Compute.SetBackendServices(value)
			return nil
		},
		nil,
	)

	fieldsMap[MaxSize] = commons.NewGenericField(
		commons.OceanGKEImport,
		MaxSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *int = nil
			if cluster.Capacity != nil && cluster.Capacity.Maximum != nil {
				value = cluster.Capacity.Maximum
			}
			if err := resourceData.Set(string(MaxSize), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MaxSize), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(MaxSize)); ok {
				if cluster.Capacity == nil {
					cluster.Capacity = &gcp.Capacity{}
				}

				cluster.Capacity.SetMaximum(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(MaxSize)); ok {
				cluster.Capacity.SetMaximum(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[MinSize] = commons.NewGenericField(
		commons.OceanGKEImport,
		MinSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *int = nil
			if cluster.Capacity != nil && cluster.Capacity.Minimum != nil {
				value = cluster.Capacity.Minimum
			}
			if err := resourceData.Set(string(MinSize), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MinSize), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(MinSize)); ok {
				if cluster.Capacity == nil {
					cluster.Capacity = &gcp.Capacity{}
				}

				cluster.Capacity.SetMinimum(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(MinSize)); ok {
				cluster.Capacity.SetMinimum(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[DesiredCapacity] = commons.NewGenericField(
		commons.OceanGKEImport,
		DesiredCapacity,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *int = nil
			if cluster.Capacity != nil && cluster.Capacity.Target != nil {
				value = cluster.Capacity.Target
			}
			if err := resourceData.Set(string(DesiredCapacity), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DesiredCapacity), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(DesiredCapacity)).(int); ok {
				if cluster.Capacity == nil {
					cluster.Capacity = &gcp.Capacity{}
				}

				if v >= 0 {
					cluster.Capacity.SetTarget(spotinst.Int(v))
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(DesiredCapacity)).(int); ok && v >= 0 {
				cluster.Capacity.SetTarget(spotinst.Int(v))
			}

			return nil
		},
		nil,
	)

	fieldsMap[ClusterControllerID] = commons.NewGenericField(
		commons.OceanGKEImport,
		ClusterControllerID,
		&schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		nil,
		nil,
		nil,
		nil,
	)

	fieldsMap[ControllerClusterID] = commons.NewGenericField(
		commons.OceanGKEImport,
		ControllerClusterID,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if err := resourceData.Set(string(ControllerClusterID), spotinst.StringValue(cluster.ControllerClusterID)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(ControllerClusterID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(ControllerClusterID)); ok {
				cluster.SetControllerClusterId(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
		nil,
	)

	fieldsMap[UpdatePolicy] = commons.NewGenericField(
		commons.OceanGKEImport,
		UpdatePolicy,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ShouldRoll): {
						Type:     schema.TypeBool,
						Required: true,
					},
					string(ConditionedRoll): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(RollConfig): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(BatchSizePercentage): {
									Type:     schema.TypeInt,
									Required: true,
								},
								string(LaunchSpecIDs): {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								string(BatchMinHealthyPercentage): {
									Type:     schema.TypeInt,
									Optional: true,
								},
								string(RespectPdb): {
									Type:     schema.TypeBool,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		nil, nil, nil, nil,
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
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
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

func expandServices(data interface{}) ([]*gcp.BackendService, error) {
	list := data.(*schema.Set).List()
	out := make([]*gcp.BackendService, 0, len(list))

	for _, v := range list {
		elem := &gcp.BackendService{}
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		if v, ok := attr[string(ServiceName)]; ok {
			elem.SetBackendServiceName(spotinst.String(v.(string)))
		}

		if v, ok := attr[string(Scheme)].(string); ok && v != "" {
			elem.SetScheme(spotinst.String(v))
		}

		if v, ok := attr[string(LocationType)].(string); ok && v != "" {
			elem.SetLocationType(spotinst.String(v))

			if v != "regional" {
				if v, ok := attr[string(NamedPorts)]; ok {
					namedPorts, err := expandNamedPorts(v)
					if err != nil {
						return nil, err
					}
					if namedPorts != nil {
						elem.SetNamedPorts(namedPorts)
					}
				}
			}
		}
		out = append(out, elem)
	}
	return out, nil
}

func expandNamedPorts(data interface{}) (*gcp.NamedPorts, error) {
	list := data.(*schema.Set).List()
	namedPorts := &gcp.NamedPorts{}

	for _, item := range list {
		m := item.(map[string]interface{})
		if v, ok := m[string(Name)].(string); ok && v != "" {
			namedPorts.SetName(spotinst.String(v))
		}

		if v, ok := m[string(Ports)]; ok && v != nil {
			portsList := v.([]interface{})
			result := make([]int, len(portsList))
			for i, j := range portsList {
				if intVal, err := strconv.Atoi(j.(string)); err != nil {
					return nil, err
				} else {
					result[i] = intVal
				}
			}
			namedPorts.SetPorts(result)
		}
	}
	return namedPorts, nil
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
