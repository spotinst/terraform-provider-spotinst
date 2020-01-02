package ocean_gke_import

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
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
			if v, ok := resourceData.GetOk(string(DesiredCapacity)); ok {
				if cluster.Capacity == nil {
					cluster.Capacity = &gcp.Capacity{}
				}

				cluster.Capacity.SetTarget(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(DesiredCapacity)); ok {
				cluster.Capacity.SetTarget(spotinst.Int(v.(int)))
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
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
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
