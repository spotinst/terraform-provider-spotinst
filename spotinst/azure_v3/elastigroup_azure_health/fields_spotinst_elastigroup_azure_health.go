package elastigroup_azure_health

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Health] = commons.NewGenericField(
		commons.ElastigroupAzureHealth,
		Health,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(HealthCheckTypes): {
						Type: schema.TypeList,
						Elem: &schema.Schema{
							Type: schema.TypeString},
						Optional: true,
					},
					string(GracePeriod): {
						Type:     schema.TypeInt,
						Optional: true,
					},
					string(UnhealthyDuration): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
					string(AutoHealing): {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup.Health != nil {
				result = flattenHealth(elastigroup.Health)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Health), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Health), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()

			if v, ok := resourceData.GetOk(string(Health)); ok {
				if health, err := expandHealth(v); err != nil {
					return err
				} else {
					elastigroup.SetHealth(health)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *azurev3.Health = nil

			if v, ok := resourceData.GetOk(string(Health)); ok {
				if health, err := expandHealth(v); err != nil {
					return err
				} else {
					value = health
				}
			}
			elastigroup.SetHealth(value)
			return nil
		},
		nil,
	)
}

func flattenHealth(health *azurev3.Health) []interface{} {
	var out []interface{}

	if health != nil {
		result := make(map[string]interface{})
		value := spotinst.Int(-1)
		result[string(UnhealthyDuration)] = value
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
	health := &azurev3.Health{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return health, nil
	}
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

	if v, ok := m[string(UnhealthyDuration)].(int); ok {
		if v == -1 {
			health.SetUnhealthyDuration(nil)
		} else {
			health.SetUnhealthyDuration(spotinst.Int(v))
		}
	}

	if v, ok := m[string(AutoHealing)].(bool); ok {
		health.SetAutoHealing(spotinst.Bool(v))
	} else {
		health.SetAutoHealing(nil)
	}
	return health, nil
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
