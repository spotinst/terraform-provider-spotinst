package health_check

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/healthcheck"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Name] = commons.NewGenericField(
		commons.HealthCheck,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			hcWrapper := resourceObject.(*commons.HealthCheckWrapper)
			healthCheck := hcWrapper.GetHealthCheck()
			var value *string = nil
			if healthCheck.Name != nil {
				value = healthCheck.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			hcWrapper := resourceObject.(*commons.HealthCheckWrapper)
			healthCheck := hcWrapper.GetHealthCheck()
			healthCheck.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			hcWrapper := resourceObject.(*commons.HealthCheckWrapper)
			healthCheck := hcWrapper.GetHealthCheck()
			healthCheck.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[ResourceId] = commons.NewGenericField(
		commons.HealthCheck,
		ResourceId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			hcWrapper := resourceObject.(*commons.HealthCheckWrapper)
			healthCheck := hcWrapper.GetHealthCheck()
			var value *string = nil
			if healthCheck.ResourceID != nil {
				value = healthCheck.ResourceID
			}
			if err := resourceData.Set(string(ResourceId), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ResourceId), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			hcWrapper := resourceObject.(*commons.HealthCheckWrapper)
			healthCheck := hcWrapper.GetHealthCheck()
			healthCheck.SetResourceId(spotinst.String(resourceData.Get(string(ResourceId)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			hcWrapper := resourceObject.(*commons.HealthCheckWrapper)
			healthCheck := hcWrapper.GetHealthCheck()
			healthCheck.SetResourceId(spotinst.String(resourceData.Get(string(ResourceId)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[ProxyAddr] = commons.NewGenericField(
		commons.HealthCheck,
		ProxyAddr,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			hcWrapper := resourceObject.(*commons.HealthCheckWrapper)
			healthCheck := hcWrapper.GetHealthCheck()
			var value *string = nil
			if healthCheck.ProxyAddr != nil {
				value = healthCheck.ProxyAddr
			}
			if err := resourceData.Set(string(ProxyAddr), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ProxyAddr), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			hcWrapper := resourceObject.(*commons.HealthCheckWrapper)
			healthCheck := hcWrapper.GetHealthCheck()
			healthCheck.SetProxyAddr(spotinst.String(resourceData.Get(string(ProxyAddr)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			hcWrapper := resourceObject.(*commons.HealthCheckWrapper)
			healthCheck := hcWrapper.GetHealthCheck()
			healthCheck.SetProxyAddr(spotinst.String(resourceData.Get(string(ProxyAddr)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[ProxyPort] = commons.NewGenericField(
		commons.HealthCheck,
		ProxyPort,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			hcWrapper := resourceObject.(*commons.HealthCheckWrapper)
			healthCheck := hcWrapper.GetHealthCheck()
			var value *int = nil
			if healthCheck.ProxyPort != nil {
				value = healthCheck.ProxyPort
			}
			if err := resourceData.Set(string(ProxyPort), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ProxyPort), err)
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			hcWrapper := resourceObject.(*commons.HealthCheckWrapper)
			healthCheck := hcWrapper.GetHealthCheck()
			if v, ok := resourceData.GetOkExists(string(ProxyPort)); ok {
				value := v.(int)
				healthCheck.SetProxyPort(spotinst.Int(value))
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			hcWrapper := resourceObject.(*commons.HealthCheckWrapper)
			healthCheck := hcWrapper.GetHealthCheck()
			if v, ok := resourceData.GetOkExists(string(ProxyPort)); ok {
				value := v.(int)
				healthCheck.SetProxyPort(spotinst.Int(value))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Check] = commons.NewGenericField(
		commons.HealthCheck,
		Check,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Protocol): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(Port): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(Endpoint): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(EndPoint): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(Interval): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(Timeout): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(TimeOut): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(Unhealthy): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(Healthy): {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			hcWrapper := resourceObject.(*commons.HealthCheckWrapper)
			healthCheck := hcWrapper.GetHealthCheck()
			if v, ok := resourceData.GetOk(string(Check)); ok {
				if check, err := expandCheck(v); err != nil {
					return err
				} else {
					healthCheck.SetCheck(check)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			hcWrapper := resourceObject.(*commons.HealthCheckWrapper)
			healthCheck := hcWrapper.GetHealthCheck()
			var value *healthcheck.Check = nil
			if v, ok := resourceData.GetOk(string(Check)); ok {
				if integration, err := expandCheck(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			healthCheck.SetCheck(value)
			return nil
		},
		nil,
	)

}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandCheck(data interface{}) (*healthcheck.Check, error) {
	check := &healthcheck.Check{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return check, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Protocol)].(string); ok && v != "" {
		check.SetProtocol(spotinst.String(v))
	}

	if v, ok := m[string(Port)].(int); ok && v > 0 {
		check.SetPort(spotinst.Int(v))
	} else {
		check.SetPort(nil)
	}

	if v, ok := m[string(Endpoint)].(string); ok && v != "" {
		check.SetEndpoint(spotinst.String(v))
	} else if v, ok := m[string(EndPoint)].(string); ok && v != "" {
		check.SetEndpoint(spotinst.String(v))
	} else {
		check.SetEndpoint(nil)
	}

	if v, ok := m[string(Interval)].(int); ok && v > 0 {
		check.SetInterval(spotinst.Int(v))
	} else {
		check.SetInterval(nil)
	}

	if v, ok := m[string(Timeout)].(int); ok && v > 0 {
		check.SetTimeout(spotinst.Int(v))
	} else if v, ok := m[string(TimeOut)].(int); ok && v > 0 {
		check.SetTimeout(spotinst.Int(v))
	} else {
		check.SetTimeout(nil)
	}

	if v, ok := m[string(Unhealthy)].(int); ok && v > 0 {
		check.SetUnhealthy(spotinst.Int(v))
	} else {
		check.SetUnhealthy(nil)
	}

	if v, ok := m[string(Healthy)].(int); ok && v > 0 {
		check.SetHealthy(spotinst.Int(v))
	} else {
		check.SetHealthy(nil)
	}

	return check, nil
}
