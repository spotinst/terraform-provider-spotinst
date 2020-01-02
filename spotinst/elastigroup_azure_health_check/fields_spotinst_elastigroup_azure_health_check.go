package elastigroup_azure_health_check

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[HealthCheck] = commons.NewGenericField(
		commons.ElastigroupAzureHealthCheck,
		HealthCheck,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(AutoHealing): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(HealthCheckType): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(GracePeriod): {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{} = nil

			if elastigroup.Compute != nil && elastigroup.Compute.Health != nil {
				value = flattenAzureGroupHealthCheck(elastigroup.Compute.Health)
			}
			if err := resourceData.Set(string(HealthCheck), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(HealthCheck), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(HealthCheck)); ok {
				if healthcheck, err := expandAzureGroupHealthCheck(v); err != nil {
					return err
				} else {
					elastigroup.Compute.SetHealth(healthcheck)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *azure.Health = nil

			if v, ok := resourceData.GetOk(string(HealthCheck)); ok {
				if healthcheck, err := expandAzureGroupHealthCheck(v); err != nil {
					return err
				} else {
					value = healthcheck
				}
			}
			elastigroup.Compute.SetHealth(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utilities
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func flattenAzureGroupHealthCheck(healthcheck *azure.Health) []interface{} {
	result := make(map[string]interface{})
	result[string(AutoHealing)] = spotinst.BoolValue(healthcheck.AutoHealing)
	result[string(HealthCheckType)] = spotinst.StringValue(healthcheck.HealthCheckType)
	result[string(GracePeriod)] = spotinst.IntValue(healthcheck.GracePeriod)
	return []interface{}{result}
}

func expandAzureGroupHealthCheck(data interface{}) (*azure.Health, error) {
	healthcheck := &azure.Health{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(AutoHealing)].(bool); ok {
			healthcheck.SetAutoHealing(spotinst.Bool(v))
		}

		if v, ok := m[string(HealthCheckType)].(string); ok && v != "" {
			healthcheck.SetHealthCheckType(spotinst.String(v))
		}

		if v, ok := m[string(GracePeriod)].(int); ok && v >= 0 {
			healthcheck.SetGracePeriod(spotinst.Int(v))
		}
	}
	return healthcheck, nil
}
