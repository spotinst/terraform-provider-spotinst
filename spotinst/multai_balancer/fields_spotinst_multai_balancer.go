package multai_balancer

import (
	"bytes"
	"errors"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.MultaiBalancer,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mlbWrapper := resourceObject.(*commons.MultaiBalancerWrapper)
			balancer := mlbWrapper.GetMultaiBalancer()
			var value *string = nil
			if balancer.Name != nil {
				value = balancer.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mlbWrapper := resourceObject.(*commons.MultaiBalancerWrapper)
			balancer := mlbWrapper.GetMultaiBalancer()
			balancer.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mlbWrapper := resourceObject.(*commons.MultaiBalancerWrapper)
			balancer := mlbWrapper.GetMultaiBalancer()
			balancer.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Scheme] = commons.NewGenericField(
		commons.MultaiBalancer,
		Scheme,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mlbWrapper := resourceObject.(*commons.MultaiBalancerWrapper)
			balancer := mlbWrapper.GetMultaiBalancer()
			if v, ok := resourceData.GetOk(string(Scheme)); ok {
				balancer.SetScheme(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mlbWrapper := resourceObject.(*commons.MultaiBalancerWrapper)
			balancer := mlbWrapper.GetMultaiBalancer()
			if v, ok := resourceData.GetOk(string(Scheme)); ok {
				balancer.SetScheme(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[DNSCnameAliases] = commons.NewGenericField(
		commons.MultaiBalancer,
		DNSCnameAliases,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mlbWrapper := resourceObject.(*commons.MultaiBalancerWrapper)
			balancer := mlbWrapper.GetMultaiBalancer()
			var value []string = nil
			if balancer.DNSCNAMEAliases != nil {
				value = balancer.DNSCNAMEAliases
			}
			if err := resourceData.Set(string(DNSCnameAliases), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DNSCnameAliases), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mlbWrapper := resourceObject.(*commons.MultaiBalancerWrapper)
			balancer := mlbWrapper.GetMultaiBalancer()
			if value, ok := resourceData.GetOk(string(DNSCnameAliases)); ok && value != nil {
				if aliases, err := expandCnameAliases(value); err != nil {
					return err
				} else {
					balancer.SetDNSCNAMEAliases(aliases)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mlbWrapper := resourceObject.(*commons.MultaiBalancerWrapper)
			balancer := mlbWrapper.GetMultaiBalancer()
			if value, ok := resourceData.GetOk(string(DNSCnameAliases)); ok && value != nil {
				if aliases, err := expandCnameAliases(value); err != nil {
					return err
				} else {
					balancer.SetDNSCNAMEAliases(aliases)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[ConnectionTimeouts] = commons.NewGenericField(
		commons.MultaiBalancer,
		ConnectionTimeouts,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Idle): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(Draining): {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mlbWrapper := resourceObject.(*commons.MultaiBalancerWrapper)
			balancer := mlbWrapper.GetMultaiBalancer()
			var value interface{} = nil
			if balancer.Timeouts != nil {
				value = flattenBalancerTimeouts(balancer.Timeouts)
			}
			return resourceData.Set(string(ConnectionTimeouts), value)
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mlbWrapper := resourceObject.(*commons.MultaiBalancerWrapper)
			balancer := mlbWrapper.GetMultaiBalancer()
			if v, ok := resourceData.GetOk(string(ConnectionTimeouts)); ok {
				if timeouts, err := expandBalancerTimeouts(v); err != nil {
					return err
				} else {
					balancer.SetTimeouts(timeouts)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mlbWrapper := resourceObject.(*commons.MultaiBalancerWrapper)
			balancer := mlbWrapper.GetMultaiBalancer()
			var value *multai.Timeouts = nil
			if v, ok := resourceData.GetOk(string(ConnectionTimeouts)); ok {
				if timeouts, err := expandBalancerTimeouts(v); err != nil {
					return err
				} else {
					value = timeouts
				}
			}
			balancer.SetTimeouts(value)
			return nil
		},
		nil,
	)

	fieldsMap[Tags] = commons.NewGenericField(
		commons.MultaiBalancer,
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
			mlbWrapper := resourceObject.(*commons.MultaiBalancerWrapper)
			balancer := mlbWrapper.GetMultaiBalancer()
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					balancer.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mlbWrapper := resourceObject.(*commons.MultaiBalancerWrapper)
			balancer := mlbWrapper.GetMultaiBalancer()
			var tagsToAdd []*multai.Tag = nil
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					tagsToAdd = tags
				}
			}
			balancer.SetTags(tagsToAdd)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func expandCnameAliases(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if alias, ok := v.(string); ok && alias != "" {
			result = append(result, alias)
		}
	}
	return result, nil
}

func expandBalancerTimeouts(data interface{}) (*multai.Timeouts, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	timeouts := &multai.Timeouts{}
	if v, ok := m["idle"].(int); ok {
		timeouts.Idle = spotinst.Int(v)
	}
	if v, ok := m["draining"].(int); ok {
		timeouts.Draining = spotinst.Int(v)
	}
	log.Printf("[DEBUG] Timeouts configuration: %s", stringutil.Stringify(timeouts))
	return timeouts, nil
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

func flattenBalancerTimeouts(timeouts *multai.Timeouts) []interface{} {
	out := make(map[string]interface{})
	out["idle"] = spotinst.IntValue(timeouts.Idle)
	out["draining"] = spotinst.IntValue(timeouts.Draining)
	return []interface{}{out}
}

func hashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(TagKey)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(TagValue)].(string)))
	return hashcode.String(buf.String())
}
