package managed_instance_healthcheck

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[HealthCheckType] = commons.NewGenericField(
		commons.ManagedInstanceAWSHealthCheck,
		HealthCheckType,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *string = nil
			if managedInstance.HealthCheck != nil && managedInstance.HealthCheck.HealthCheckType != nil {
				value = managedInstance.HealthCheck.HealthCheckType
			}
			if err := resourceData.Set(string(HealthCheckType), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(HealthCheckType), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(HealthCheckType)).(string); ok && v != "" {
				managedInstance.HealthCheck.SetHealthCheckType(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *string = nil
			if v, ok := resourceData.Get(string(HealthCheckType)).(string); ok && v != "" {
				value = spotinst.String(v)
			}
			managedInstance.HealthCheck.SetHealthCheckType(value)
			return nil
		},
		nil,
	)

	fieldsMap[AutoHealing] = commons.NewGenericField(
		commons.ManagedInstanceAWSHealthCheck,
		AutoHealing,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *bool = nil
			if managedInstance.HealthCheck != nil && managedInstance.HealthCheck.AutoHealing != nil {
				value = managedInstance.HealthCheck.AutoHealing
			}
			if err := resourceData.Set(string(AutoHealing), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AutoHealing), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(AutoHealing)).(bool); ok {
				managedInstance.HealthCheck.SetAutoHealing(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(AutoHealing)).(bool); ok {
				managedInstance.HealthCheck.SetAutoHealing(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[GracePeriod] = commons.NewGenericField(
		commons.ManagedInstanceAWSHealthCheck,
		GracePeriod,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *int = nil
			if managedInstance.HealthCheck != nil && managedInstance.HealthCheck.GracePeriod != nil {
				value = managedInstance.HealthCheck.GracePeriod
			}
			if err := resourceData.Set(string(GracePeriod), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(GracePeriod), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(GracePeriod)).(int); ok && v > 0 {
				managedInstance.HealthCheck.SetGracePeriod(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *int = nil
			if v, ok := resourceData.Get(string(GracePeriod)).(int); ok && v > 0 {
				value = spotinst.Int(v)
			}
			managedInstance.HealthCheck.SetGracePeriod(value)
			return nil
		},
		nil,
	)

	fieldsMap[UnhealthyDuration] = commons.NewGenericField(
		commons.ManagedInstanceAWS,
		UnhealthyDuration,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *int = nil
			if managedInstance.HealthCheck != nil && managedInstance.HealthCheck.UnhealthyDuration != nil {
				value = managedInstance.HealthCheck.UnhealthyDuration
			}
			if err := resourceData.Set(string(UnhealthyDuration), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UnhealthyDuration), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(UnhealthyDuration)).(int); ok && v > 0 {
				managedInstance.HealthCheck.SetUnhealthyDuration(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *int = nil
			if v, ok := resourceData.Get(string(UnhealthyDuration)).(int); ok && v > 0 {
				value = spotinst.Int(v)
			}
			managedInstance.HealthCheck.SetUnhealthyDuration(value)
			return nil
		},
		nil,
	)
}
