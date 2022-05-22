package stateful_node_azure_health

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Health] = commons.NewGenericField(
		commons.StatefulNodeAzureHealth,
		Health,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(HealthCheckTypes): {
						Type: schema.TypeList,
						Elem: &schema.Schema{
							Type: schema.TypeString},
						Required: true,
					},
					string(GracePeriod): {
						Type:     schema.TypeInt,
						Optional: true,
						Computed: true,
					},
					string(UnhealthyDuration): {
						Type:     schema.TypeInt,
						Optional: true,
						Computed: true,
					},
					string(AutoHealing): {
						Type:     schema.TypeBool,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var result []interface{} = nil
			if statefulNode.Health != nil {
				health := statefulNode.Health
				result = flattenHealth(health)
			}
			if result != nil {
				if err := resourceData.Set(string(Health), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Health), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()

			if v, ok := resourceData.GetOk(string(Health)); ok {
				if health, err := expandHealth(v); err != nil {
					return err
				} else {
					statefulNode.SetHealth(health)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *azurev3.Health = nil

			if v, ok := resourceData.GetOk(string(Health)); ok {
				if health, err := expandHealth(v); err != nil {
					return err
				} else {
					value = health
				}
			}
			statefulNode.SetHealth(value)
			return nil
		},
		nil,
	)
}

func flattenHealth(health *azurev3.Health) []interface{} {
	var out []interface{}

	if health != nil {
		result := make(map[string]interface{})

		if health.HealthCheckTypes != nil {
			result[string(HealthCheckTypes)] = spotinst.StringSlice(health.HealthCheckTypes)
		}

		if health.GracePeriod != nil {
			result[string(GracePeriod)] = spotinst.IntValue(health.GracePeriod)
		}

		if health.UnhealthyDuration != nil {
			result[string(UnhealthyDuration)] = spotinst.IntValue(health.UnhealthyDuration)
		}

		if health.AutoHealing != nil {
			result[string(AutoHealing)] = spotinst.BoolValue(health.AutoHealing)
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

func expandHealth(data interface{}) (*azurev3.Health, error) {
	if list := data.([]interface{}); len(list) > 0 {
		health := &azurev3.Health{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(HealthCheckTypes)]; ok {
				htc, err := expandHealthCheckTypes(v)
				if err != nil {
					return nil, err
				}

				if htc != nil {
					health.SetHealthCheckTypes(htc)
				}
			} else {
				health.SetHealthCheckTypes(nil)
			}

			if v, ok := m[string(GracePeriod)].(int); ok && v >= 0 {
				health.SetGracePeriod(spotinst.Int(v))
			} else {
				health.SetGracePeriod(nil)
			}

			if v, ok := m[string(UnhealthyDuration)].(int); ok && v >= 0 {
				health.SetUnhealthyDuration(spotinst.Int(v))
			} else {
				health.SetUnhealthyDuration(nil)
			}

			if v, ok := m[string(AutoHealing)].(bool); ok {
				health.SetAutoHealing(spotinst.Bool(v))
			} else {
				health.SetAutoHealing(nil)
			}
		}

		return health, nil
	}

	return nil, nil
}

func expandHealthCheckTypes(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if healthCheckType, ok := v.(string); ok && healthCheckType != "" {
			result = append(result, healthCheckType)
		}
	}

	return result, nil
}
