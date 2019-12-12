package multai_routing_rule

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
		commons.MultaiRoutingRule,
		BalancerID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			var value *string = nil
			if routing.BalancerID != nil {
				value = routing.BalancerID
			}
			if err := resourceData.Set(string(BalancerID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BalancerID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			routing.SetBalancerId(spotinst.String(resourceData.Get(string(BalancerID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			routing.SetBalancerId(spotinst.String(resourceData.Get(string(BalancerID)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[ListenerID] = commons.NewGenericField(
		commons.MultaiRoutingRule,
		ListenerID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			var value *string = nil
			if routing.ListenerID != nil {
				value = routing.ListenerID
			}
			if err := resourceData.Set(string(ListenerID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ListenerID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			routing.SetListenerId(spotinst.String(resourceData.Get(string(ListenerID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			routing.SetListenerId(spotinst.String(resourceData.Get(string(ListenerID)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Route] = commons.NewGenericField(
		commons.MultaiRoutingRule,
		Route,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			var value *string = nil
			if routing.Route != nil {
				value = routing.Route
			}
			if err := resourceData.Set(string(Route), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Route), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			routing.SetRoute(spotinst.String(resourceData.Get(string(Route)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			routing.SetRoute(spotinst.String(resourceData.Get(string(Route)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Priority] = commons.NewGenericField(
		commons.MultaiRoutingRule,
		Priority,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			var value *int = nil
			if routing.Priority != nil {
				value = routing.Priority
			}
			if err := resourceData.Set(string(Priority), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Priority), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			routing.SetPriority(spotinst.Int(resourceData.Get(string(Priority)).(int)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			routing.SetPriority(spotinst.Int(resourceData.Get(string(Priority)).(int)))
			return nil
		},
		nil,
	)

	fieldsMap[Strategy] = commons.NewGenericField(
		commons.MultaiRoutingRule,
		Strategy,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  multai.StrategyRoundRobin.String(),
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			var value *string = nil
			if routing.Strategy != nil {
				value = routing.Strategy
			}
			if err := resourceData.Set(string(Strategy), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Strategy), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			routing.SetStrategy(spotinst.String(resourceData.Get(string(Strategy)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			routing.SetStrategy(spotinst.String(resourceData.Get(string(Strategy)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[MiddlewareIDs] = commons.NewGenericField(
		commons.MultaiRoutingRule,
		MiddlewareIDs,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			var value []string = nil
			if routing.MiddlewareIDs != nil {
				value = routing.MiddlewareIDs
			}
			if err := resourceData.Set(string(MiddlewareIDs), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MiddlewareIDs), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			if value, ok := resourceData.GetOk(string(MiddlewareIDs)); ok && value != nil {
				if ids, err := expandMiddlewareIDs(value); err != nil {
					return err
				} else {
					routing.SetMiddlewareIDs(ids)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			if value, ok := resourceData.GetOk(string(MiddlewareIDs)); ok && value != nil {
				if ids, err := expandMiddlewareIDs(value); err != nil {
					return err
				} else {
					routing.SetMiddlewareIDs(ids)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[TargetSetIDs] = commons.NewGenericField(
		commons.MultaiRoutingRule,
		TargetSetIDs,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			var value []string = nil
			if routing.TargetSetIDs != nil {
				value = routing.TargetSetIDs
			}
			if err := resourceData.Set(string(TargetSetIDs), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(TargetSetIDs), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			if value, ok := resourceData.GetOk(string(TargetSetIDs)); ok && value != nil {
				if ids, err := expandTargetSetIDs(value); err != nil {
					return err
				} else {
					routing.SetTargetSetIDs(ids)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			if value, ok := resourceData.GetOk(string(TargetSetIDs)); ok && value != nil {
				if ids, err := expandTargetSetIDs(value); err != nil {
					return err
				} else {
					routing.SetTargetSetIDs(ids)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[Tags] = commons.NewGenericField(
		commons.MultaiRoutingRule,
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
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					routing.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			routingWrapper := resourceObject.(*commons.MultaiRoutingRuleWrapper)
			routing := routingWrapper.GetMultaiRoutingRule()
			var tagsToAdd []*multai.Tag = nil
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					tagsToAdd = tags
				}
			}
			routing.SetTags(tagsToAdd)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func expandMiddlewareIDs(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if item, ok := v.(string); ok && item != "" {
			result = append(result, item)
		}
	}
	return result, nil
}

func expandTargetSetIDs(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if item, ok := v.(string); ok && item != "" {
			result = append(result, item)
		}
	}
	return result, nil
}

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

func hashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(TagKey)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(TagValue)].(string)))
	return hashcode.String(buf.String())
}
