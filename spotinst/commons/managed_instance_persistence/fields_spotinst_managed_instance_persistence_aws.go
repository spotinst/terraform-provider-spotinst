package managed_instance_persistence

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[BlockDevicesMode] = commons.NewGenericField(
		commons.ManagedInstanceAWSPersistence,
		BlockDevicesMode,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *string = nil
			if managedInstance.Persistence != nil && managedInstance.Persistence.BlockDevicesMode != nil {
				value = managedInstance.Persistence.BlockDevicesMode
			}
			if err := resourceData.Set(string(BlockDevicesMode), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BlockDevicesMode), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(BlockDevicesMode)).(string); ok && v != "" {
				initPersistenceIfNeeded(managedInstance)
				managedInstance.Persistence.SetBlockDevicesMode(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *string = nil
			if v, ok := resourceData.Get(string(BlockDevicesMode)).(string); ok && v != "" {
				initPersistenceIfNeeded(managedInstance)
				value = spotinst.String(v)
			}
			managedInstance.Persistence.SetBlockDevicesMode(value)
			return nil
		},
		nil,
	)

	fieldsMap[PersistPrivateIp] = commons.NewGenericField(
		commons.ManagedInstanceAWSPersistence,
		PersistPrivateIp,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *bool = nil
			if managedInstance.Persistence != nil && managedInstance.Persistence.PersistPrivateIP != nil {
				value = managedInstance.Persistence.PersistPrivateIP
			}
			if err := resourceData.Set(string(PersistPrivateIp), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PersistPrivateIp), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(PersistPrivateIp)).(bool); ok {
				initPersistenceIfNeeded(managedInstance)
				managedInstance.Persistence.SetPersistPrivateIP(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(PersistPrivateIp)).(bool); ok {
				initPersistenceIfNeeded(managedInstance)
				managedInstance.Persistence.SetPersistPrivateIP(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[PersistRootDevice] = commons.NewGenericField(
		commons.ManagedInstanceAWSPersistence,
		PersistRootDevice,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *bool = nil
			if managedInstance.Persistence != nil && managedInstance.Persistence.PersistRootDevice != nil {
				value = managedInstance.Persistence.PersistRootDevice
			}
			if err := resourceData.Set(string(PersistRootDevice), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PersistRootDevice), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(PersistRootDevice)).(bool); ok {
				initPersistenceIfNeeded(managedInstance)
				managedInstance.Persistence.SetShouldPersistRootDevice(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *bool = nil
			if v, ok := resourceData.Get(string(PersistRootDevice)).(bool); ok {
				initPersistenceIfNeeded(managedInstance)
				value = spotinst.Bool(v)
			}
			managedInstance.Persistence.SetShouldPersistRootDevice(value)
			return nil
		},
		nil,
	)

	fieldsMap[PersistBlockDevices] = commons.NewGenericField(
		commons.ManagedInstanceAWSPersistence,
		PersistBlockDevices,
		&schema.Schema{
			Type:     schema.TypeBool,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *bool = nil
			if managedInstance.Persistence != nil && managedInstance.Persistence.PersistBlockDevices != nil {
				value = managedInstance.Persistence.PersistBlockDevices
			}
			if err := resourceData.Set(string(PersistBlockDevices), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PersistBlockDevices), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(PersistBlockDevices)).(bool); ok {
				initPersistenceIfNeeded(managedInstance)
				managedInstance.Persistence.SetShouldPersistBlockDevices(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *bool = nil
			if v, ok := resourceData.Get(string(PersistBlockDevices)).(bool); ok {
				initPersistenceIfNeeded(managedInstance)
				value = spotinst.Bool(v)
			}
			managedInstance.Persistence.SetShouldPersistBlockDevices(value)
			return nil
		},
		nil,
	)
}

func initPersistenceIfNeeded(managedInstance *aws.ManagedInstance) {
	if managedInstance.Persistence == nil {
		managedInstance.Persistence = new(aws.Persistence)
	}
}
