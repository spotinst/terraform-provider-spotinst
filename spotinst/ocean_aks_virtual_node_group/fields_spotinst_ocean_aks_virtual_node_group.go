package ocean_aks_virtual_node_group

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[OceanID] = commons.NewGenericField(
		commons.OceanAKSVirtualNodeGroup,
		OceanID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var value *string = nil
			if virtualNodeGroup.OceanID != nil {
				value = virtualNodeGroup.OceanID
			}
			if err := resourceData.Set(string(OceanID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OceanID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			virtualNodeGroup.SetOceanId(spotinst.String(resourceData.Get(string(OceanID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			virtualNodeGroup.SetOceanId(spotinst.String(resourceData.Get(string(OceanID)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanAKSVirtualNodeGroup,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var value *string = nil
			if virtualNodeGroup.Name != nil {
				value = virtualNodeGroup.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if value, ok := resourceData.GetOk(string(Name)); ok && value != nil {
				virtualNodeGroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if value, ok := resourceData.GetOk(string(Name)); ok && value != nil {
				virtualNodeGroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Zones] = commons.NewGenericField(
		commons.OceanAKSVirtualNodeGroup,
		Zones,
		&schema.Schema{
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString},
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var value []string = nil
			if virtualNodeGroup.Zones != nil {
				value = virtualNodeGroup.Zones
			}
			if err := resourceData.Set(string(Zones), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Zones), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if value, ok := resourceData.GetOk(string(Zones)); ok && value != nil {
				if zones, err := expandZones(value); err != nil {
					return err
				} else {
					virtualNodeGroup.SetZones(zones)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if value, ok := resourceData.GetOk(string(Zones)); ok && value != nil {
				if zones, err := expandZones(value); err != nil {
					return err
				} else {
					virtualNodeGroup.SetZones(zones)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[Label] = commons.NewGenericField(
		commons.OceanAKSVirtualNodeGroup,
		Label,
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
						Optional: true,
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var result []interface{} = nil

			if virtualNodeGroup != nil && virtualNodeGroup.Labels != nil {
				result = flattenLabels(virtualNodeGroup.Labels)
			}

			if err := resourceData.Set(string(Label), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Label), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if value, ok := resourceData.GetOk(string(Label)); ok && value != nil {
				if labels, err := expandLabels(value); err != nil {
					return err
				} else {
					virtualNodeGroup.SetLabels(labels)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var value []*azure.Label = nil

			if v, ok := resourceData.GetOk(string(Label)); ok {
				if labels, err := expandLabels(v); err != nil {
					return err
				} else {
					value = labels
				}
			}
			virtualNodeGroup.SetLabels(value)
			return nil
		},
		nil,
	)

	fieldsMap[Taint] = commons.NewGenericField(
		commons.OceanAKSVirtualNodeGroup,
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
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var result []interface{} = nil
			if virtualNodeGroup.Labels != nil {
				taints := virtualNodeGroup.Taints
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
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if value, ok := resourceData.GetOk(string(Taint)); ok {
				if taints, err := expandTaints(value); err != nil {
					return err
				} else {
					virtualNodeGroup.SetTaints(taints)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var taintList []*azure.Taint = nil
			if value, ok := resourceData.GetOk(string(Taint)); ok {
				if taints, err := expandTaints(value); err != nil {
					return err
				} else {
					taintList = taints
				}
			}
			virtualNodeGroup.SetTaints(taintList)
			return nil
		},
		nil,
	)

	fieldsMap[ResourceLimits] = commons.NewGenericField(
		commons.OceanAKSVirtualNodeGroup,
		ResourceLimits,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MaxInstanceCount): {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var result []interface{} = nil
			if virtualNodeGroup.ResourceLimits != nil {
				resourceLimits := virtualNodeGroup.ResourceLimits
				result = flattenResourceLimits(resourceLimits)
			}
			if err := resourceData.Set(string(ResourceLimits), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ResourceLimits), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if value, ok := resourceData.GetOk(string(ResourceLimits)); ok {
				if resourceLimits, err := expandResourceLimits(value); err != nil {
					return err
				} else {
					virtualNodeGroup.SetResourceLimits(resourceLimits)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var resourceLimits *azure.VirtualNodeGroupResourceLimits = nil
			if value, ok := resourceData.GetOk(string(ResourceLimits)); ok {
				if rl, err := expandResourceLimits(value); err != nil {
					return err
				} else {
					resourceLimits = rl
				}
			}
			virtualNodeGroup.SetResourceLimits(resourceLimits)
			return nil
		},
		nil,
	)
}

func expandZones(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if zones, ok := v.(string); ok && zones != "" {
			result = append(result, zones)
		}
	}
	return result, nil
}

func expandLabels(data interface{}) ([]*azure.Label, error) {
	list := data.(*schema.Set).List()
	labels := make([]*azure.Label, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		labels = append(labels, &azure.Label{
			Key:   spotinst.String(attr[string(LabelKey)].(string)),
			Value: spotinst.String(attr[string(LabelValue)].(string)),
		})
	}

	return labels, nil
}

func flattenLabels(labels []*azure.Label) []interface{} {
	result := make([]interface{}, 0, len(labels))

	for _, label := range labels {
		m := make(map[string]interface{})
		m[string(LabelKey)] = spotinst.StringValue(label.Key)
		m[string(LabelValue)] = spotinst.StringValue(label.Value)
		result = append(result, m)
	}

	return result
}

func expandTaints(data interface{}) ([]*azure.Taint, error) {
	list := data.(*schema.Set).List()
	taints := make([]*azure.Taint, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		taints = append(taints, &azure.Taint{
			Key:    spotinst.String(attr[string(TaintKey)].(string)),
			Value:  spotinst.String(attr[string(TaintValue)].(string)),
			Effect: spotinst.String(attr[string(TaintEffect)].(string)),
		})
	}

	return taints, nil
}

func flattenTaints(taints []*azure.Taint) []interface{} {
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

func expandResourceLimits(data interface{}) (*azure.VirtualNodeGroupResourceLimits, error) {
	if list := data.([]interface{}); len(list) > 0 {
		resLimits := &azure.VirtualNodeGroupResourceLimits{}
		var maxInstanceCount *int = nil

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(MaxInstanceCount)].(int); ok && v > 0 {
				maxInstanceCount = spotinst.Int(v)
			}

			resLimits.SetMaxInstanceCount(maxInstanceCount)
		}

		return resLimits, nil
	}

	return nil, nil
}

func flattenResourceLimits(resourceLimits *azure.VirtualNodeGroupResourceLimits) []interface{} {
	result := make(map[string]interface{})
	result[string(MaxInstanceCount)] = spotinst.IntValue(resourceLimits.MaxInstanceCount)
	return []interface{}{result}
}
