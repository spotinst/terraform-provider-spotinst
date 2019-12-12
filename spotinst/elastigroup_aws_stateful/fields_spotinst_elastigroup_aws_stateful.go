package elastigroup_aws_stateful

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

	fieldsMap[PersistRootDevice] = commons.NewGenericField(
		commons.ElastigroupAWSStateful,
		PersistRootDevice,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(PersistRootDevice)).(bool); ok {
				elastigroup.Strategy.Persistence.SetShouldPersistRootDevice(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
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
		commons.ElastigroupAWSStateful,
		PersistBlockDevices,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(PersistBlockDevices)).(bool); ok {
				elastigroup.Strategy.Persistence.SetShouldPersistBlockDevices(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
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
		commons.ElastigroupAWSStateful,
		PersistPrivateIp,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(PersistPrivateIp)).(bool); ok {
				elastigroup.Strategy.Persistence.SetShouldPersistPrivateIP(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
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
		commons.ElastigroupAWSStateful,
		BlockDevicesMode,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(BlockDevicesMode)).(string); ok && v != "" {
				elastigroup.Strategy.Persistence.SetBlockDevicesMode(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
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
		commons.ElastigroupAWSStateful,
		PrivateIps,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if elastigroup.Compute != nil && elastigroup.Compute.PrivateIPs != nil {
				value := elastigroup.Compute.PrivateIPs
				if err := resourceData.Set(string(PrivateIps), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PrivateIps), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(PrivateIps)); ok {
				if eips, err := expandAWSGroupPrivateIPs(v); err != nil {
					return err
				} else {
					elastigroup.Compute.SetPrivateIPs(eips)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []string = nil
			if v, ok := resourceData.GetOk(string(PrivateIps)); ok {
				if eips, err := expandAWSGroupPrivateIPs(v); err != nil {
					return err
				} else {
					result = eips
				}
			}
			elastigroup.Compute.SetPrivateIPs(result)
			return nil
		},
		nil,
	)

	fieldsMap[StatefulDeallocation] = commons.NewGenericField(
		commons.ElastigroupAWSStateful,
		StatefulDeallocation,
		&schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ShouldDeleteImages): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(ShouldDeleteNetworkInterfaces): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(ShouldDeleteVolumes): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(ShouldDeleteSnapshots): {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		nil, nil, nil, nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupPrivateIPs(data interface{}) ([]string, error) {
	list := data.([]interface{})
	pips := make([]string, 0, len(list))
	for _, str := range list {
		if pip, ok := str.(string); ok {
			pips = append(pips, pip)
		}
	}
	return pips, nil
}
