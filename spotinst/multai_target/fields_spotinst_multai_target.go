package multai_target

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[BalancerID] = commons.NewGenericField(
		commons.MultaiTarget,
		BalancerID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			var value *string = nil
			if target.BalancerID != nil {
				value = target.BalancerID
			}
			if err := resourceData.Set(string(BalancerID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BalancerID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			target.SetBalancerId(spotinst.String(resourceData.Get(string(BalancerID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			target.SetBalancerId(spotinst.String(resourceData.Get(string(BalancerID)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[TargetSetID] = commons.NewGenericField(
		commons.MultaiTarget,
		TargetSetID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			var value *string = nil
			if target.TargetSetID != nil {
				value = target.TargetSetID
			}
			if err := resourceData.Set(string(TargetSetID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(TargetSetID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			target.SetTargetSetId(spotinst.String(resourceData.Get(string(TargetSetID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			target.SetTargetSetId(spotinst.String(resourceData.Get(string(TargetSetID)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Port] = commons.NewGenericField(
		commons.MultaiTarget,
		Port,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			var value *int = nil
			if target.Port != nil {
				value = target.Port
			}
			if err := resourceData.Set(string(Port), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Port), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			if v, ok := resourceData.GetOk(string(Port)); ok {
				target.SetPort(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			if v, ok := resourceData.GetOk(string(Port)); ok {
				target.SetPort(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Host] = commons.NewGenericField(
		commons.MultaiTarget,
		Host,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			if v, ok := resourceData.GetOk(string(Host)); ok {
				target.SetHost(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			if v, ok := resourceData.GetOk(string(Host)); ok {
				target.SetHost(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Name] = commons.NewGenericField(
		commons.MultaiTarget,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			if v, ok := resourceData.GetOk(string(Name)); ok {
				target.SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			if v, ok := resourceData.GetOk(string(Name)); ok {
				target.SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Weight] = commons.NewGenericField(
		commons.MultaiTarget,
		Weight,
		&schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			if v, ok := resourceData.GetOk(string(Weight)); ok {
				target.SetWeight(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			if v, ok := resourceData.GetOk(string(Weight)); ok {
				target.SetWeight(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Tags] = commons.NewGenericField(
		commons.MultaiTarget,
		Tags,
		&schema.Schema{
			Type:     schema.TypeSet,
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
			Set: hashKV,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					target.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetWrapper := resourceObject.(*commons.MultaiTargetWrapper)
			target := targetWrapper.GetMultaiTarget()
			var tagsToAdd []*multai.Tag = nil
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					tagsToAdd = tags
				}
			}
			target.SetTags(tagsToAdd)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func expandTags(data interface{}) ([]*multai.Tag, error) {
	list := data.(*schema.Set).List()
	tags := make([]*multai.Tag, 0, len(list))
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
		tag := &multai.Tag{
			Key:   spotinst.String(attr[string(TagKey)].(string)),
			Value: spotinst.String(attr[string(TagValue)].(string)),
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func flattenTags(tags []*multai.Tag) []interface{} {
	result := make([]interface{}, 0, len(tags))
	for _, tag := range tags {
		m := make(map[string]interface{})
		m[string(TagKey)] = spotinst.StringValue(tag.Key)
		m[string(TagValue)] = spotinst.StringValue(tag.Value)

		result = append(result, m)
	}
	return result
}

func hashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(TagKey)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(TagValue)].(string)))
	return hashcode.String(buf.String())
}
