package ocean_aks_virtual_node_group_launch_specification

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[LaunchSpecification] = commons.NewGenericField(
		commons.OceanAKSVirtualNodeGroupLaunchSpecification,
		LaunchSpecification,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(OSDisk): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(SizeGB): {
									Type:     schema.TypeInt,
									Required: true,
								},
								string(Type): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(UtilizeEphemeralStorage): {
									Type:     schema.TypeBool,
									Optional: true,
								},
							},
						},
					},

					string(Tag): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(TagKey): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(TagValue): {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var result []interface{} = nil

			if virtualNodeGroup != nil && virtualNodeGroup.LaunchSpecification != nil {
				result = flattenLaunchSpecification(virtualNodeGroup.LaunchSpecification)
			}

			if len(result) > 0 {
				if err := resourceData.Set(string(LaunchSpecification), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OSDisk), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.GetOk(string(LaunchSpecification)); ok {
				if launchSpecification, err := expandLaunchSpecification(v); err != nil {
					return err
				} else {
					virtualNodeGroup.SetLaunchSpecification(launchSpecification)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.VirtualNodeGroupAKSWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVirtualNodeGroup()
			var value *azure.VirtualNodeGroupLaunchSpecification = nil

			if v, ok := resourceData.GetOk(string(OSDisk)); ok {
				if launchSpecification, err := expandLaunchSpecification(v); err != nil {
					return err
				} else {
					value = launchSpecification
				}
			}
			virtualNodeGroup.SetLaunchSpecification(value)
			return nil
		},

		nil,
	)
}

func expandLaunchSpecification(data interface{}) (*azure.VirtualNodeGroupLaunchSpecification, error) {
	launchSpecification := &azure.VirtualNodeGroupLaunchSpecification{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return launchSpecification, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(OSDisk)]; ok && v != nil {
		if osDisk, err := expandOSDisk(v); err != nil {
			return nil, err
		} else {
			if osDisk != nil {
				launchSpecification.SetOSDisk(osDisk)
			}
		}
	}

	if v, ok := m[string(Tag)]; ok && v != nil {
		if tags, err := expandTags(v); err != nil {
			return nil, err
		} else {
			if tags != nil {
				launchSpecification.SetTags(tags)
			}
		}
	}

	return launchSpecification, nil
}

func flattenLaunchSpecification(launchSpecification *azure.VirtualNodeGroupLaunchSpecification) []interface{} {
	var out []interface{}

	if launchSpecification != nil {
		result := make(map[string]interface{})

		if launchSpecification.OSDisk != nil {
			result[string(OSDisk)] = flattenOSDisk(launchSpecification.OSDisk)
		}

		if launchSpecification.Tags != nil {
			result[string(Tag)] = flattenTags(launchSpecification.Tags)
		}

		return []interface{}{result}
	}

	return out
}

func expandTags(data interface{}) ([]*azure.Tag, error) {
	list := data.(*schema.Set).List()
	tags := make([]*azure.Tag, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(TagKey)]; !ok {
			return nil, errors.New("invalid tag attributes: key missing")
		}

		if _, ok := attr[string(TagValue)]; !ok {
			return nil, errors.New("invalid tag attributes: value missing")
		}
		tag := &azure.Tag{
			Key:   spotinst.String(attr[string(TagKey)].(string)),
			Value: spotinst.String(attr[string(TagValue)].(string)),
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func flattenTags(tags []*azure.Tag) []interface{} {
	result := make([]interface{}, 0, len(tags))
	for _, tag := range tags {
		m := make(map[string]interface{})
		m[string(TagKey)] = spotinst.StringValue(tag.Key)
		m[string(TagValue)] = spotinst.StringValue(tag.Value)

		result = append(result, m)
	}
	return result
}

func flattenOSDisk(osd *azure.OSDisk) interface{} {
	osDisk := make(map[string]interface{})
	osDisk[string(SizeGB)] = spotinst.IntValue(osd.SizeGB)
	osDisk[string(Type)] = spotinst.StringValue(osd.Type)
	osDisk[string(UtilizeEphemeralStorage)] = spotinst.BoolValue(osd.UtilizeEphemeralStorage)
	return []interface{}{osDisk}
}

func expandOSDisk(data interface{}) (*azure.OSDisk, error) {
	if list := data.([]interface{}); len(list) > 0 {
		osDisk := &azure.OSDisk{}
		if list[0] != nil {
			m := list[0].(map[string]interface{})
			var sizeGB *int = nil
			var osType *string = nil
			var osUtilizeEphemeralStorage *bool = nil

			if v, ok := m[string(SizeGB)].(int); ok && v > 0 {
				sizeGB = spotinst.Int(v)
			}
			osDisk.SetSizeGB(sizeGB)

			if v, ok := m[string(Type)].(string); ok && v != "" {
				osType = spotinst.String(v)
			}
			osDisk.SetType(osType)

			if v, ok := m[string(UtilizeEphemeralStorage)].(bool); ok {
				osUtilizeEphemeralStorage = spotinst.Bool(v)
			}
			osDisk.SetUtilizeEphemeralStorage(osUtilizeEphemeralStorage)

		}
		return osDisk, nil
	}
	return nil, nil
}
