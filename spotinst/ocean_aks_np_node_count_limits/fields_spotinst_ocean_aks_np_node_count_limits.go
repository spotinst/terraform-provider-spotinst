package ocean_aks_np_node_count_limits

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[MinCount] = commons.NewGenericField(
		commons.OceanAKSNPNodeCountLimits,
		MinCount,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *int = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.NodeCountLimits != nil && cluster.VirtualNodeGroupTemplate.NodeCountLimits.MinCount != nil {
				value = cluster.VirtualNodeGroupTemplate.NodeCountLimits.MinCount
			} else {
				value = spotinst.Int(-1)
			}

			if err := resourceData.Set(string(MinCount), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MinCount), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(MinCount)).(int); ok && v >= 0 {
				cluster.VirtualNodeGroupTemplate.NodeCountLimits.SetMinCount(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.NodeCountLimits.SetMinCount(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(MinCount)).(int); ok && v >= 0 {
				cluster.VirtualNodeGroupTemplate.NodeCountLimits.SetMinCount(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.NodeCountLimits.SetMinCount(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[MaxCount] = commons.NewGenericField(
		commons.OceanAKSNPNodeCountLimits,
		MaxCount,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *int = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.NodeCountLimits != nil && cluster.VirtualNodeGroupTemplate.NodeCountLimits.MaxCount != nil {
				value = cluster.VirtualNodeGroupTemplate.NodeCountLimits.MaxCount
			} else {
				value = spotinst.Int(-1)
			}
			if err := resourceData.Set(string(MaxCount), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MaxCount), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(MaxCount)).(int); ok && v >= 0 {
				cluster.VirtualNodeGroupTemplate.NodeCountLimits.SetMaxCount(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.NodeCountLimits.SetMaxCount(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(MaxCount)).(int); ok && v >= 0 {
				cluster.VirtualNodeGroupTemplate.NodeCountLimits.SetMaxCount(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.NodeCountLimits.SetMaxCount(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[Tags] = commons.NewGenericField(
		commons.OceanAKSNPNodeCountLimits,
		Tags,
		&schema.Schema{
			Type:     schema.TypeMap,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()

			if err := resourceData.Set(string(Tags), cluster.VirtualNodeGroupTemplate.Tags); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tags), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(Tags)).(interface{}); ok {
				if tags, err := expandTags(v); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *map[string]string = nil
			if v, ok := resourceData.GetOk(string(Tags)); ok {
				if tag, err := expandTags(v); err != nil {
					return err
				} else {
					value = tag
				}
			}
			if cluster.VirtualNodeGroupTemplate.Tags == nil {
				cluster.VirtualNodeGroupTemplate.Tags = &map[string]string{}
			}
			cluster.VirtualNodeGroupTemplate.SetTags(value)
			return nil
		},
		nil,
	)

	fieldsMap[Label] = commons.NewGenericField(
		commons.OceanAKSNPNodeCountLimits,
		Label,
		&schema.Schema{
			Type:     schema.TypeMap,
			Optional: true,
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()

			if err := resourceData.Set(string(Label), cluster.VirtualNodeGroupTemplate.Labels); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Label), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.Get(string(Label)).(interface{}); ok {
				if labels, err := expandLabels(value); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.SetLabels(labels)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()

			var value *map[string]string = nil
			if v, ok := resourceData.GetOk(string(Label)); ok {
				if label, err := expandLabels(v); err != nil {
					return err
				} else {
					value = label
				}
			}
			if cluster.VirtualNodeGroupTemplate.Labels == nil {
				cluster.VirtualNodeGroupTemplate.Labels = &map[string]string{}
			}
			cluster.VirtualNodeGroupTemplate.SetLabels(value)
			return nil
		},
		nil,
	)

	fieldsMap[Taint] = commons.NewGenericField(
		commons.OceanAKSNPNodeCountLimits,
		Taint,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(TaintKey): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(TaintValue): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(TaintEffect): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var result []interface{} = nil
			if cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.Taints != nil {
				taints := cluster.VirtualNodeGroupTemplate.Taints
				result = flattenTaints(taints)
			}
			if result != nil {
				if err := resourceData.Set(string(Taint), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Taint), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.GetOk(string(Taint)); ok {
				if taints, err := expandTaints(value); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.SetTaints(taints)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var taintList []*azure_np.Taint = nil
			if value, ok := resourceData.GetOk(string(Taint)); ok {
				if taints, err := expandTaints(value); err != nil {
					return err
				} else {
					taintList = taints
				}
			}
			cluster.VirtualNodeGroupTemplate.SetTaints(taintList)
			return nil
		},
		nil,
	)
}

func expandTags(data interface{}) (*map[string]string, error) {
	m, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("could not cast tags")
	}
	result := make(map[string]string, len(m))
	for k, v := range m {
		val, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("could not cast tags value to string")
		}
		result[k] = val
	}
	return &result, nil
}

func flattenTags(tags map[string]string) map[string]interface{} {
	result := make(map[string]interface{}, len(tags))
	for k, v := range tags {
		result[k] = v
	}
	return result
}

func expandLabels(data interface{}) (*map[string]string, error) {
	m, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("could not cast labels")
	}
	result := make(map[string]string, len(m))
	for k, v := range m {
		val, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("could not cast labels value to string")
		}
		result[k] = val
	}
	return &result, nil
}

func flattenLabels(labels map[string]string) map[string]interface{} {
	result := make(map[string]interface{}, len(labels))
	for k, v := range labels {
		result[k] = v
	}
	return result
}

func expandTaints(data interface{}) ([]*azure_np.Taint, error) {
	list := data.(*schema.Set).List()
	taints := make([]*azure_np.Taint, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		taints = append(taints, &azure_np.Taint{
			Key:    spotinst.String(attr[string(TaintKey)].(string)),
			Value:  spotinst.String(attr[string(TaintValue)].(string)),
			Effect: spotinst.String(attr[string(TaintEffect)].(string)),
		})
	}
	return taints, nil
}

func flattenTaints(taints []*azure_np.Taint) []interface{} {
	result := make([]interface{}, 0, len(taints))

	for _, taint := range taints {
		m := make(map[string]interface{})
		m[string(TaintKey)] = spotinst.StringValue(taint.Key)
		m[string(TaintValue)] = spotinst.StringValue(taint.Value)
		m[string(TaintEffect)] = spotinst.StringValue(taint.Effect)
		result = append(result, m)
	}
	return result
}
