package multai_listener

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"

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

	fieldsMap[BalancerID] = commons.NewGenericField(
		commons.MultaiListener,
		BalancerID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			listenerWrapper := resourceObject.(*commons.MultaiListenerWrapper)
			listener := listenerWrapper.GetMultaiListener()
			var value *string = nil
			if listener.BalancerID != nil {
				value = listener.BalancerID
			}
			if err := resourceData.Set(string(BalancerID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BalancerID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			listenerWrapper := resourceObject.(*commons.MultaiListenerWrapper)
			listener := listenerWrapper.GetMultaiListener()
			listener.SetBalancerId(spotinst.String(resourceData.Get(string(BalancerID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			listenerWrapper := resourceObject.(*commons.MultaiListenerWrapper)
			listener := listenerWrapper.GetMultaiListener()
			listener.SetBalancerId(spotinst.String(resourceData.Get(string(BalancerID)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Port] = commons.NewGenericField(
		commons.MultaiListener,
		Port,
		&schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			listenerWrapper := resourceObject.(*commons.MultaiListenerWrapper)
			listener := listenerWrapper.GetMultaiListener()
			if v, ok := resourceData.GetOk(string(Port)); ok {
				listener.SetPort(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			listenerWrapper := resourceObject.(*commons.MultaiListenerWrapper)
			listener := listenerWrapper.GetMultaiListener()
			if v, ok := resourceData.GetOk(string(Port)); ok {
				listener.SetPort(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Protocol] = commons.NewGenericField(
		commons.MultaiListener,
		Protocol,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			StateFunc: func(v interface{}) string {
				value := v.(string)
				return strings.ToUpper(value)
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			listenerWrapper := resourceObject.(*commons.MultaiListenerWrapper)
			listener := listenerWrapper.GetMultaiListener()
			if v, ok := resourceData.GetOk(string(Protocol)); ok {
				listener.SetProtocol(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			listenerWrapper := resourceObject.(*commons.MultaiListenerWrapper)
			listener := listenerWrapper.GetMultaiListener()
			if v, ok := resourceData.GetOk(string(Protocol)); ok {
				listener.SetProtocol(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[TLSConfig] = commons.NewGenericField(
		commons.MultaiListener,
		TLSConfig,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(CertificateIDs): {
						Type:     schema.TypeList,
						Required: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(MinVersion): {
						Type:     schema.TypeString,
						Required: true,
						StateFunc: func(v interface{}) string {
							value := v.(string)
							return strings.ToUpper(value)
						},
					},

					string(MaxVersion): {
						Type:     schema.TypeString,
						Required: true,
						StateFunc: func(v interface{}) string {
							value := v.(string)
							return strings.ToUpper(value)
						},
					},

					string(SessionTicketsDisabled): {
						Type:     schema.TypeBool,
						Required: true,
					},

					string(PreferServerCipherSuites): {
						Type:     schema.TypeBool,
						Required: true,
					},

					string(CipherSuites): {
						Type:     schema.TypeList,
						Required: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
				},
			},
			Set: hashBalancerListenerTLSConfig,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			listenerWrapper := resourceObject.(*commons.MultaiListenerWrapper)
			listener := listenerWrapper.GetMultaiListener()
			var value interface{} = nil
			if listener.TLSConfig != nil {
				value = flattenListenerTLSConfig(listener.TLSConfig)
			}
			return resourceData.Set(string(TLSConfig), value)
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			listenerWrapper := resourceObject.(*commons.MultaiListenerWrapper)
			listener := listenerWrapper.GetMultaiListener()
			if v, ok := resourceData.GetOk(string(TLSConfig)); ok {
				if timeouts, err := expandListenerTLSConfig(v); err != nil {
					return err
				} else {
					listener.SetTLSConfig(timeouts)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			listenerWrapper := resourceObject.(*commons.MultaiListenerWrapper)
			listener := listenerWrapper.GetMultaiListener()
			var value *multai.TLSConfig = nil
			if v, ok := resourceData.GetOk(string(TLSConfig)); ok {
				if timeouts, err := expandListenerTLSConfig(v); err != nil {
					return err
				} else {
					value = timeouts
				}
			}
			listener.SetTLSConfig(value)
			return nil
		},
		nil,
	)

	fieldsMap[Tags] = commons.NewGenericField(
		commons.MultaiListener,
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
			listenerWrapper := resourceObject.(*commons.MultaiListenerWrapper)
			listener := listenerWrapper.GetMultaiListener()
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					listener.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			listenerWrapper := resourceObject.(*commons.MultaiListenerWrapper)
			listener := listenerWrapper.GetMultaiListener()
			var tagsToAdd []*multai.Tag = nil
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					tagsToAdd = tags
				}
			}
			listener.SetTags(tagsToAdd)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func expandListenerTLSConfig(data interface{}) (*multai.TLSConfig, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	tls := &multai.TLSConfig{}
	if v, ok := m[string(MaxVersion)].(string); ok {
		tls.MaxVersion = spotinst.String(v)
	}
	if v, ok := m[string(MinVersion)].(string); ok {
		tls.MinVersion = spotinst.String(v)
	}
	if v, ok := m[string(SessionTicketsDisabled)].(bool); ok {
		tls.SessionTicketsDisabled = spotinst.Bool(v)
	}
	if v, ok := m[string(PreferServerCipherSuites)].(bool); ok {
		tls.PreferServerCipherSuites = spotinst.Bool(v)
	}

	if v, ok := m[string(CertificateIDs)]; ok {
		ids := v.([]interface{})
		result := make([]string, len(ids))
		for i, j := range ids {
			result[i] = j.(string)
		}
		tls.CertificateIDs = result
	}

	if v, ok := m[string(CipherSuites)]; ok {
		ids := v.([]interface{})
		result := make([]string, len(ids))
		for i, j := range ids {
			result[i] = j.(string)
		}
		tls.CipherSuites = result
	}

	log.Printf("[DEBUG] TLSConfig configuration: %s", stringutil.Stringify(tls))
	return tls, nil
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

func flattenListenerTLSConfig(tls *multai.TLSConfig) []interface{} {
	out := make(map[string]interface{})
	out[string(MinVersion)] = spotinst.StringValue(tls.MinVersion)
	out[string(MaxVersion)] = spotinst.StringValue(tls.MaxVersion)
	out[string(SessionTicketsDisabled)] = spotinst.BoolValue(tls.SessionTicketsDisabled)
	out[string(PreferServerCipherSuites)] = spotinst.BoolValue(tls.PreferServerCipherSuites)
	out[string(CertificateIDs)] = tls.CertificateIDs
	out[string(CipherSuites)] = tls.CipherSuites
	return []interface{}{out}
}

func hashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(TagKey)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(TagValue)].(string)))
	return hashcode.String(buf.String())
}

func hashBalancerListenerTLSConfig(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m[string(MinVersion)].(string))))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m[string(MinVersion)].(string))))
	buf.WriteString(fmt.Sprintf("%t-", m[string(SessionTicketsDisabled)].(bool)))
	buf.WriteString(fmt.Sprintf("%t-", m[string(PreferServerCipherSuites)].(bool)))
	if ids, ok := m[string(CertificateIDs)].([]string); ok {
		for _, id := range ids {
			buf.WriteString(fmt.Sprintf("%s-", id))
		}
	}
	if css, ok := m[string(CipherSuites)].([]string); ok {
		for _, cs := range css {
			buf.WriteString(fmt.Sprintf("%s-", cs))
		}
	}
	return hashcode.String(buf.String())
}
