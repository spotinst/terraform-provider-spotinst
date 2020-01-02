package ocean_gke

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanGKE,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if cluster.Name != nil {
				value = cluster.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			cluster.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			cluster.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[MasterLocation] = commons.NewGenericField(
		commons.OceanGKE,
		MasterLocation,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(MasterLocation)); ok && v != nil {
				if cluster.GKE == nil {
					cluster.SetGKE(&gcp.GKE{})
				}
				cluster.GKE.SetMasterLocation(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[ClusterName] = commons.NewGenericField(
		commons.OceanGKE,
		ClusterName,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if cluster.GKE != nil && cluster.GKE.ClusterName != nil {
				value = cluster.GKE.ClusterName
			}
			if err := resourceData.Set(string(ClusterName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ClusterName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(ClusterName)); ok && v != nil {
				if cluster.GKE == nil {
					cluster.SetGKE(&gcp.GKE{})
				}
				cluster.GKE.SetClusterName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if v, ok := resourceData.GetOk(string(ClusterName)); ok && v != nil {
				if cluster.GKE == nil {
					cluster.SetGKE(&gcp.GKE{})
				}
				value = spotinst.String(v.(string))
			}
			cluster.GKE.SetClusterName(value)
			return nil
		},
		nil,
	)

	fieldsMap[ControllerClusterID] = commons.NewGenericField(
		commons.OceanGKE,
		ControllerClusterID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if cluster.ControllerClusterID != nil {
				value = cluster.ControllerClusterID
			}
			if err := resourceData.Set(string(ControllerClusterID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ControllerClusterID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			cluster.SetControllerClusterId(spotinst.String(resourceData.Get(string(ControllerClusterID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			cluster.SetControllerClusterId(spotinst.String(resourceData.Get(string(ControllerClusterID)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[MaxSize] = commons.NewGenericField(
		commons.OceanGKE,
		MaxSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(MaxSize)); ok {
				cluster.Capacity.SetMaximum(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(MaxSize)); ok {
				cluster.Capacity.SetMaximum(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[MinSize] = commons.NewGenericField(
		commons.OceanGKE,
		MinSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(MinSize)); ok {
				cluster.Capacity.SetMinimum(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(MinSize)); ok {
				cluster.Capacity.SetMinimum(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[DesiredCapacity] = commons.NewGenericField(
		commons.OceanGKE,
		DesiredCapacity,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(DesiredCapacity)); ok {
				cluster.Capacity.SetTarget(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(DesiredCapacity)); ok {
				cluster.Capacity.SetTarget(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[SubnetName] = commons.NewGenericField(
		commons.OceanGKE,
		SubnetName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if cluster.Compute != nil && cluster.Compute.SubnetName != nil {
				value = cluster.Compute.SubnetName
			}
			if err := resourceData.Set(string(SubnetName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SubnetName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			cluster.Compute.SetSubnetName(spotinst.String(resourceData.Get(string(SubnetName)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			cluster.Compute.SetSubnetName(spotinst.String(resourceData.Get(string(SubnetName)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[AvailabilityZones] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		AvailabilityZones,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			//DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []string = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.AvailabilityZones != nil {
				result = append(result, cluster.Compute.AvailabilityZones...)
			}
			if err := resourceData.Set(string(AvailabilityZones), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AvailabilityZones), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []string
			if v, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				azList := v.([]interface{})
				result = make([]string, len(azList))
				for i, j := range azList {
					result[i] = j.(string)
				}
				cluster.Compute.SetAvailabilityZones(result)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []string
			if v, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				azList := v.([]interface{})
				result = make([]string, len(azList))
				for i, j := range azList {
					result[i] = j.(string)
				}
				cluster.Compute.SetAvailabilityZones(result)
			}
			return nil
		},
		nil,
	)

	fieldsMap[BackendServices] = commons.NewGenericField(
		commons.ElastigroupGCPLaunchConfiguration,
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
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
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

	fieldsMap[SourceImage] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		SourceImage,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.GKEClusterWrapper)
			ls := lsWrapper.GetCluster().Compute.LaunchSpecification
			var value *string = nil
			if ls != nil && ls.SourceImage != nil {
				value = ls.SourceImage
			}
			if err := resourceData.Set(string(SourceImage), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SourceImage), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.GKEClusterWrapper)
			ls := lsWrapper.GetCluster().Compute.LaunchSpecification
			if v, ok := resourceData.Get(string(SourceImage)).(string); ok && v != "" {
				ls.SetSourceImage(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.GKEClusterWrapper)
			ls := lsWrapper.GetCluster().Compute.LaunchSpecification
			if v, ok := resourceData.Get(string(SourceImage)).(string); ok && v != "" {
				ls.SetSourceImage(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Metadata] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		Metadata,
		&schema.Schema{
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MetadataKey): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(MetadataValue): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			Set: hashKV,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.GKEClusterWrapper)
			ls := lsWrapper.GetCluster().Compute.LaunchSpecification
			var result []interface{} = nil
			if ls != nil && ls.Metadata != nil {
				metadata := ls.Metadata
				result = flattenMetadata(metadata)
			}
			if result != nil {
				if err := resourceData.Set(string(Metadata), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Metadata), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.GKEClusterWrapper)
			ls := lsWrapper.GetCluster().Compute.LaunchSpecification
			if value, ok := resourceData.GetOk(string(Metadata)); ok {
				if metadata, err := expandMetadata(value); err != nil {
					return err
				} else {
					ls.SetMetadata(metadata)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.GKEClusterWrapper)
			ls := lsWrapper.GetCluster().Compute.LaunchSpecification
			var metadataList []*gcp.Metadata
			if value, ok := resourceData.GetOk(string(Metadata)); ok {
				if metadata, err := expandMetadata(value); err != nil {
					return err
				} else {
					metadataList = metadata
				}
			}
			ls.SetMetadata(metadataList)
			return nil
		},
		nil,
	)

	fieldsMap[Labels] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		Labels,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(LabelKey): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(LabelValue): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			Set: hashKV,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.GKEClusterWrapper)
			ls := lsWrapper.GetCluster().Compute.LaunchSpecification
			var result []interface{} = nil
			if ls != nil && ls.Labels != nil {
				labels := ls.Labels
				result = flattenLabels(labels)
			}
			if result != nil {
				if err := resourceData.Set(string(Labels), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Labels), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.GKEClusterWrapper)
			ls := lsWrapper.GetCluster().Compute.LaunchSpecification
			if value, ok := resourceData.GetOk(string(Labels)); ok {
				if labels, err := expandLabels(value); err != nil {
					return err
				} else {
					ls.SetLabels(labels)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.GKEClusterWrapper)
			ls := lsWrapper.GetCluster().Compute.LaunchSpecification
			var labelList []*gcp.Label = nil
			if value, ok := resourceData.GetOk(string(Labels)); ok {
				if labels, err := expandLabels(value); err != nil {
					return err
				} else {
					labelList = labels
				}
			}
			ls.SetLabels(labelList)
			return nil
		},
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

func hashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(LabelKey)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(LabelValue)].(string)))
	return hashcode.String(buf.String())
}

func expandLabels(data interface{}) ([]*gcp.Label, error) {
	list := data.(*schema.Set).List()
	labels := make([]*gcp.Label, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(LabelKey)]; !ok {
			return nil, errors.New("invalid label attributes: key missing")
		}

		if _, ok := attr[string(LabelValue)]; !ok {
			return nil, errors.New("invalid label attributes: value missing")
		}
		label := &gcp.Label{
			Key:   spotinst.String(attr[string(LabelKey)].(string)),
			Value: spotinst.String(attr[string(LabelValue)].(string)),
		}
		labels = append(labels, label)
	}
	return labels, nil
}

func expandMetadata(data interface{}) ([]*gcp.Metadata, error) {
	list := data.(*schema.Set).List()
	metadata := make([]*gcp.Metadata, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(MetadataKey)]; !ok {
			return nil, errors.New("invalid metadata attributes: key missing")
		}

		if _, ok := attr[string(MetadataValue)]; !ok {
			return nil, errors.New("invalid metadata attributes: value missing")
		}
		metaObject := &gcp.Metadata{
			Key:   spotinst.String(attr[string(MetadataKey)].(string)),
			Value: spotinst.String(attr[string(MetadataValue)].(string)),
		}
		metadata = append(metadata, metaObject)
	}
	return metadata, nil
}

func flattenLabels(labels []*gcp.Label) []interface{} {
	result := make([]interface{}, 0, len(labels))
	for _, label := range labels {
		m := make(map[string]interface{})
		m[string(LabelKey)] = spotinst.StringValue(label.Key)
		m[string(LabelValue)] = spotinst.StringValue(label.Value)

		result = append(result, m)
	}
	return result
}

func flattenMetadata(metadata []*gcp.Metadata) []interface{} {
	result := make([]interface{}, 0, len(metadata))
	for _, metaObject := range metadata {
		m := make(map[string]interface{})
		m[string(MetadataKey)] = spotinst.StringValue(metaObject.Key)
		m[string(MetadataValue)] = spotinst.StringValue(metaObject.Value)

		result = append(result, m)
	}
	return result
}
