package ocean_aks_np_virtual_node_group

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[OceanID] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroup,
		OceanID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
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
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			virtualNodeGroup.SetOceanId(spotinst.String(resourceData.Get(string(OceanID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			virtualNodeGroup.SetOceanId(spotinst.String(resourceData.Get(string(OceanID)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroup,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
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
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if value, ok := resourceData.GetOk(string(Name)); ok && value != nil {
				virtualNodeGroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if value, ok := resourceData.GetOk(string(Name)); ok && value != nil {
				virtualNodeGroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[AvailabilityZones] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroup,
		AvailabilityZones,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var value []string = nil
			if virtualNodeGroup.AvailabilityZones != nil {
				value = virtualNodeGroup.AvailabilityZones
			}
			if err := resourceData.Set(string(AvailabilityZones), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AvailabilityZones), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if value, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				if zones, err := expandAvailaiblityZones(value); err != nil {
					return err
				} else {
					virtualNodeGroup.SetAvailabilityZones(zones)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if value, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				if zones, err := expandAvailaiblityZones(value); err != nil {
					return err
				} else {
					virtualNodeGroup.SetAvailabilityZones(zones)
				}
			} else {
				virtualNodeGroup.SetAvailabilityZones(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[Taints] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroup,
		Taints,
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
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var result []interface{} = nil
			if virtualNodeGroup.Taints != nil {
				taints := virtualNodeGroup.Taints
				result = flattenTaints(taints)
			}
			if result != nil {
				if err := resourceData.Set(string(Taints), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Taints), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if value, ok := resourceData.GetOk(string(Taints)); ok {
				if taints, err := expandTaints(value); err != nil {
					return err
				} else {
					virtualNodeGroup.SetTaints(taints)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var taintList []*azure_np.Taint = nil
			if value, ok := resourceData.GetOk(string(Taints)); ok {
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

	fieldsMap[Tags] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroup,
		Tags,
		&schema.Schema{
			Type:     schema.TypeMap,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			result := make(map[string]string)
			if virtualNodeGroup.Tags != nil {
				result = flattenTags(*virtualNodeGroup.Tags)
			}

			if err := resourceData.Set(string(Tags), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tags), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.Get(string(Tags)).(interface{}); ok {
				if tags, err := expandTags(v); err != nil {
					return err
				} else {
					virtualNodeGroup.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var value *map[string]string = nil
			if v, ok := resourceData.GetOk(string(Tags)); ok {
				if tag, err := expandTags(v); err != nil {
					return err
				} else {
					value = tag
				}
			}
			if virtualNodeGroup.Tags == nil {
				virtualNodeGroup.Tags = &map[string]string{}
			}
			virtualNodeGroup.SetTags(value)
			return nil
		},
		nil,
	)

	fieldsMap[Labels] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroup,
		Labels,
		&schema.Schema{
			Type:     schema.TypeMap,
			Optional: true,
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			result := make(map[string]string)
			if virtualNodeGroup.Tags != nil {
				result = flattenLabels(*virtualNodeGroup.Labels)
			}
			if err := resourceData.Set(string(Labels), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Labels), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if value, ok := resourceData.Get(string(Labels)).(interface{}); ok {
				if labels, err := expandLabels(value); err != nil {
					return err
				} else {
					virtualNodeGroup.SetLabels(labels)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var value *map[string]string = nil
			if v, ok := resourceData.GetOk(string(Labels)); ok {
				if label, err := expandLabels(v); err != nil {
					return err
				} else {
					value = label
				}
			}
			if virtualNodeGroup.Labels == nil {
				virtualNodeGroup.Labels = &map[string]string{}
			}
			virtualNodeGroup.SetLabels(value)
			return nil
		},
		nil,
	)

	fieldsMap[UpdatePolicy] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroup,
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
					string(ConditionedRollParams): {
						Type:     schema.TypeList,
						Elem:     &schema.Schema{Type: schema.TypeString},
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
								string(VngIDs): {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								string(BatchMinHealthyPercentage): {
									Type:     schema.TypeInt,
									Optional: true,
								},
								string(RespectPDB): {
									Type:     schema.TypeBool,
									Optional: true,
								},
								string(Comment): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(NodePoolNames): {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								string(RespectRestrictScaleDown): {
									Type:     schema.TypeBool,
									Optional: true,
								},
								string(NodeNames): {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
							},
						},
					},
				},
			},
		},
		nil, nil, nil, nil,
	)
}

func expandAvailaiblityZones(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if zones, ok := v.(string); ok && zones != "" {
			result = append(result, zones)
		}
	}
	return result, nil
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

func flattenTags(tags map[string]string) map[string]string {
	result := make(map[string]string, len(tags))
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

func flattenLabels(labels map[string]string) map[string]string {
	result := make(map[string]string, len(labels))
	for k, v := range labels {
		result[k] = v
	}
	return result
}
