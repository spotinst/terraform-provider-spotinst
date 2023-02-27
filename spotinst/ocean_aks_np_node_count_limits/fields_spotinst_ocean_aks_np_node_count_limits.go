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
			}
			if err := resourceData.Set(string(MinCount), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MinCount), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(MinCount)).(int); ok && v > 0 {
				cluster.VirtualNodeGroupTemplate.NodeCountLimits.SetMinCount(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.NodeCountLimits.SetMinCount(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(MinCount)).(int); ok && v > 0 {
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
			}
			if err := resourceData.Set(string(MaxCount), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MaxCount), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(MaxCount)).(int); ok && v > 0 {
				cluster.VirtualNodeGroupTemplate.NodeCountLimits.SetMaxCount(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.NodeCountLimits.SetMaxCount(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(MaxCount)).(int); ok && v > 0 {
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
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(TagKey): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(TagValue): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value []interface{} = nil

			if cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.Tags != nil {
				value = flattenTags(cluster.VirtualNodeGroupTemplate.Tags)
			}
			if value != nil {
				if err := resourceData.Set(string(Tags), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tags), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOk(string(Tags)); ok {
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
			var value *azure_np.Tag = nil
			if v, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(v); err != nil {
					return err
				} else {
					value = tags
				}
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
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(LabelKey): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(LabelValue): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var result []interface{} = nil

			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.Labels != nil {
				result = flattenLabels(cluster.VirtualNodeGroupTemplate.Labels)
			}

			if err := resourceData.Set(string(Label), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Label), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.GetOk(string(Label)); ok && value != nil {
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
			var value *azure_np.Label = nil

			if v, ok := resourceData.GetOk(string(Label)); ok {
				if labels, err := expandLabels(v); err != nil {
					return err
				} else {
					value = labels
				}
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

func expandTags(data interface{}) (*azure_np.Tag, error) {
	if list := data.([]interface{}); len(list) > 0 {
		tags := &azure_np.Tag{}

		if list != nil || list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(TagKey)].(string); ok {
				tags.SetKey(spotinst.String(v))
			}

			if v, ok := m[string(TagValue)].(string); ok {
				tags.SetValue(spotinst.String(v))
			}
		}
		return tags, nil
	}
	return nil, nil
}

func flattenTags(tags *azure_np.Tag) []interface{} {
	tag := make(map[string]interface{})
	tag[string(TagKey)] = spotinst.StringValue(tags.Key)
	tag[string(TagValue)] = spotinst.StringValue(tags.Value)

	return []interface{}{tag}
}

func expandLabels(data interface{}) (*azure_np.Label, error) {
	if list := data.([]interface{}); len(list) > 0 {
		labels := &azure_np.Label{}

		if list != nil || list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(LabelKey)].(string); ok {
				labels.SetKey(spotinst.String(v))
			}

			if v, ok := m[string(LabelValue)].(string); ok {
				labels.SetValue(spotinst.String(v))
			}
		}
		return labels, nil
	}
	return nil, nil
}

func flattenLabels(labels *azure_np.Label) []interface{} {
	label := make(map[string]interface{})
	label[string(LabelKey)] = spotinst.StringValue(labels.Key)
	label[string(LabelValue)] = spotinst.StringValue(labels.Value)

	return []interface{}{label}
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
