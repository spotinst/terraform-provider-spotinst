package multai_target_set

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
		commons.MultaiTargetSet,
		BalancerID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			var value *string = nil
			if targetSet.BalancerID != nil {
				value = targetSet.BalancerID
			}
			if err := resourceData.Set(string(BalancerID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BalancerID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			targetSet.SetBalancerId(spotinst.String(resourceData.Get(string(BalancerID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			targetSet.SetBalancerId(spotinst.String(resourceData.Get(string(BalancerID)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[DeploymentID] = commons.NewGenericField(
		commons.MultaiTargetSet,
		DeploymentID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			var value *string = nil
			if targetSet.DeploymentID != nil {
				value = targetSet.DeploymentID
			}
			if err := resourceData.Set(string(DeploymentID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DeploymentID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			targetSet.SetDeploymentId(spotinst.String(resourceData.Get(string(DeploymentID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			targetSet.SetDeploymentId(spotinst.String(resourceData.Get(string(DeploymentID)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Port] = commons.NewGenericField(
		commons.MultaiTargetSet,
		Port,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			var value *int = nil
			if targetSet.Port != nil {
				value = targetSet.Port
			}
			if err := resourceData.Set(string(Port), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Port), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			if v, ok := resourceData.GetOk(string(Port)); ok {
				targetSet.SetPort(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			if v, ok := resourceData.GetOk(string(Port)); ok {
				targetSet.SetPort(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Protocol] = commons.NewGenericField(
		commons.MultaiTargetSet,
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
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			var value *string = nil
			if targetSet.Protocol != nil {
				value = targetSet.Protocol
			}
			if err := resourceData.Set(string(Protocol), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Protocol), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			if v, ok := resourceData.GetOk(string(Protocol)); ok {
				targetSet.SetProtocol(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			if v, ok := resourceData.GetOk(string(Protocol)); ok {
				targetSet.SetProtocol(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Name] = commons.NewGenericField(
		commons.MultaiTargetSet,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			var value *string = nil
			if targetSet.Name != nil {
				value = targetSet.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			if v, ok := resourceData.GetOk(string(Name)); ok {
				targetSet.SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			if v, ok := resourceData.GetOk(string(Name)); ok {
				targetSet.SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Weight] = commons.NewGenericField(
		commons.MultaiTargetSet,
		Weight,
		&schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			var value *int = nil
			if targetSet.Weight != nil {
				value = targetSet.Weight
			}
			if err := resourceData.Set(string(Weight), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Weight), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			if v, ok := resourceData.GetOk(string(Weight)); ok {
				targetSet.SetWeight(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			if v, ok := resourceData.GetOk(string(Weight)); ok {
				targetSet.SetWeight(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[HealthCheck] = commons.NewGenericField(
		commons.MultaiTargetSet,
		HealthCheck,
		&schema.Schema{
			Type:     schema.TypeSet,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Protocol): {
						Type:     schema.TypeString,
						Required: true,
						StateFunc: func(v interface{}) string {
							value := v.(string)
							return strings.ToUpper(value)
						},
					},

					string(Path): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(Port): {
						Type:     schema.TypeInt,
						Optional: true,
						Computed: true,
					},

					string(Interval): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(Timeout): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(HealthyThreshold): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(UnhealthyThreshold): {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
			Set: hashBalancerTargetSetHealthCheck,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			var value interface{} = nil
			if targetSet.HealthCheck != nil {
				value = flattenTargetSetHealthCheck(targetSet.HealthCheck)
			}
			return resourceData.Set(string(HealthCheck), value)
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			if v, ok := resourceData.GetOk(string(HealthCheck)); ok {
				if timeouts, err := expandTargetSetHealthCheck(v); err != nil {
					return err
				} else {
					targetSet.SetHealthCheck(timeouts)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			var value *multai.TargetSetHealthCheck = nil
			if v, ok := resourceData.GetOk(string(HealthCheck)); ok {
				if timeouts, err := expandTargetSetHealthCheck(v); err != nil {
					return err
				} else {
					value = timeouts
				}
			}
			targetSet.SetHealthCheck(value)
			return nil
		},
		nil,
	)

	fieldsMap[Tags] = commons.NewGenericField(
		commons.MultaiTargetSet,
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
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					targetSet.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			targetSetWrapper := resourceObject.(*commons.MultaiTargetSetWrapper)
			targetSet := targetSetWrapper.GetMultaiTargetSet()
			var tagsToAdd []*multai.Tag = nil
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					tagsToAdd = tags
				}
			}
			targetSet.SetTags(tagsToAdd)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func expandTargetSetHealthCheck(data interface{}) (*multai.TargetSetHealthCheck, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	tshc := &multai.TargetSetHealthCheck{}

	if v, ok := m[string(Path)].(string); ok {
		tshc.Path = spotinst.String(v)
	}

	if v, ok := m[string(Interval)].(int); ok {
		tshc.Interval = spotinst.Int(v)
	}

	if v, ok := m[string(Timeout)].(int); ok {
		tshc.Timeout = spotinst.Int(v)
	}

	if v, ok := m[string(Port)].(int); ok {
		tshc.Port = spotinst.Int(v)
	}

	if v, ok := m[string(Protocol)].(string); ok {
		tshc.Protocol = spotinst.String(v)
	}

	if v, ok := m[string(HealthyThreshold)].(int); ok {
		tshc.HealthyThresholdCount = spotinst.Int(v)
	}

	if v, ok := m[string(UnhealthyThreshold)].(int); ok {
		tshc.UnhealthyThresholdCount = spotinst.Int(v)
	}

	log.Printf("[DEBUG] TargetSet configuration: %s", stringutil.Stringify(tshc))
	return tshc, nil
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

func flattenTargetSetHealthCheck(tshc *multai.TargetSetHealthCheck) []interface{} {
	out := make(map[string]interface{})
	out[string(Path)] = spotinst.StringValue(tshc.Path)
	out[string(Protocol)] = spotinst.StringValue(tshc.Protocol)
	out[string(Port)] = spotinst.IntValue(tshc.Port)
	out[string(Interval)] = spotinst.IntValue(tshc.Interval)
	out[string(Timeout)] = spotinst.IntValue(tshc.Timeout)
	out[string(HealthyThreshold)] = spotinst.IntValue(tshc.HealthyThresholdCount)
	out[string(UnhealthyThreshold)] = spotinst.IntValue(tshc.UnhealthyThresholdCount)
	return []interface{}{out}
}

func hashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(TagKey)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(TagValue)].(string)))
	return hashcode.String(buf.String())
}

func hashBalancerTargetSetHealthCheck(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m[string(Protocol)].(string))))
	buf.WriteString(fmt.Sprintf("%s-", m[string(Path)].(string)))
	buf.WriteString(fmt.Sprintf("%d-", m[string(Port)].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m[string(Interval)].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m[string(Timeout)].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m[string(HealthyThreshold)].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m[string(UnhealthyThreshold)].(int)))
	return hashcode.String(buf.String())
}
