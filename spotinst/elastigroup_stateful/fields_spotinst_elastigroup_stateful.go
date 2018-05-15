package elastigroup_stateful

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[PersistRootDevice] = commons.NewGenericField(
		commons.ElastigroupStrategy,
		PersistRootDevice,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *bool = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.Persistence != nil &&
				elastigroup.Strategy.Persistence.ShouldPersistRootDevice != nil {
				value = elastigroup.Strategy.Persistence.ShouldPersistRootDevice
			}
			if err := resourceData.Set(string(PersistRootDevice), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PersistRootDevice), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.Get(string(PersistRootDevice)).(bool); ok {
				elastigroup.Strategy.Persistence.SetShouldPersistRootDevice(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *bool = nil
			if v, ok := resourceData.Get(string(PersistRootDevice)).(bool); ok {
				value = spotinst.Bool(v)
			}
			elastigroup.Strategy.Persistence.SetShouldPersistRootDevice(value)
			return nil
		},
		nil,
	)

	fieldsMap[PersistBlockDevices] = commons.NewGenericField(
		commons.ElastigroupStrategy,
		PersistBlockDevices,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *bool = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.Persistence != nil &&
				elastigroup.Strategy.Persistence.ShouldPersistBlockDevices != nil {
				value = elastigroup.Strategy.Persistence.ShouldPersistBlockDevices
			}
			if err := resourceData.Set(string(PersistBlockDevices), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PersistBlockDevices), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.Get(string(PersistBlockDevices)).(bool); ok {
				elastigroup.Strategy.Persistence.SetShouldPersistBlockDevices(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *bool = nil
			if v, ok := resourceData.Get(string(PersistBlockDevices)).(bool); ok {
				value = spotinst.Bool(v)
			}
			elastigroup.Strategy.Persistence.SetShouldPersistBlockDevices(value)
			return nil
		},
		nil,
	)

	fieldsMap[PersistPrivateIp] = commons.NewGenericField(
		commons.ElastigroupStrategy,
		PersistPrivateIp,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *bool = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.Persistence != nil &&
				elastigroup.Strategy.Persistence.ShouldPersistPrivateIP != nil {
				value = elastigroup.Strategy.Persistence.ShouldPersistPrivateIP
			}
			if err := resourceData.Set(string(PersistPrivateIp), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PersistPrivateIp), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.Get(string(PersistPrivateIp)).(bool); ok {
				elastigroup.Strategy.Persistence.SetShouldPersistPrivateIP(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *bool = nil
			if v, ok := resourceData.Get(string(PersistPrivateIp)).(bool); ok {
				value = spotinst.Bool(v)
			}
			elastigroup.Strategy.Persistence.SetShouldPersistPrivateIP(value)
			return nil
		},
		nil,
	)

	fieldsMap[BlockDevicesMode] = commons.NewGenericField(
		commons.ElastigroupStrategy,
		BlockDevicesMode,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *string = nil
			if elastigroup.Strategy != nil && elastigroup.Strategy.Persistence != nil &&
				elastigroup.Strategy.Persistence.BlockDevicesMode != nil {
				value = elastigroup.Strategy.Persistence.BlockDevicesMode
			}
			if err := resourceData.Set(string(BlockDevicesMode), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BlockDevicesMode), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.Get(string(BlockDevicesMode)).(string); ok && v != "" {
				elastigroup.Strategy.Persistence.SetBlockDevicesMode(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *string = nil
			if v, ok := resourceData.Get(string(BlockDevicesMode)).(string); ok && v != "" {
				value = spotinst.String(v)
			}
			elastigroup.Strategy.Persistence.SetBlockDevicesMode(value)
			return nil
		},
		nil,
	)

	fieldsMap[PrivateIps] = commons.NewGenericField(
		commons.ElastigroupStrategy,
		PrivateIps,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if elastigroup.Compute != nil && elastigroup.Compute.PrivateIPs != nil {
				value := elastigroup.Compute.PrivateIPs
				if err := resourceData.Set(string(PrivateIps), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PrivateIps), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.GetOk(string(PrivateIps)); ok {
				list := v.([]interface{})
				result := make([]string, 0, len(list))
				for _, v := range list {
					if privateIP, ok := v.(string); ok && privateIP != "" {
						result = append(result, privateIP)
					}
				}
				elastigroup.Compute.SetPrivateIPs(result)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value []string = nil
			if v, ok := resourceData.GetOk(string(PrivateIps)); ok {
				list := v.([]interface{})
				result := make([]string, 0, len(list))
				for _, v := range list {
					if privateIP, ok := v.(string); ok && privateIP != "" {
						result = append(result, privateIP)
					}
				}
				value = result
			}
			elastigroup.Compute.SetPrivateIPs(value)
			return nil
		},
		nil,
	)
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-