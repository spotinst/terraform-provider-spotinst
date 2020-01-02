package mrscaler_aws_instance_groups

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/mrscaler"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupCoreGroup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[CoreMin] = commons.NewGenericField(
		commons.MRScalerAWSCoreGroup,
		CoreMin,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *int = nil
			if scaler.Compute != nil && scaler.Compute.InstanceGroups != nil &&
				scaler.Compute.InstanceGroups.CoreGroup != nil &&
				scaler.Compute.InstanceGroups.CoreGroup.Capacity != nil &&
				scaler.Compute.InstanceGroups.CoreGroup.Capacity.Minimum != nil {
				value = scaler.Compute.InstanceGroups.CoreGroup.Capacity.Minimum
			}
			if err := resourceData.Set(string(CoreMin), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CoreMin), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(CoreMin)); ok {
				if scaler.Compute.InstanceGroups.CoreGroup == nil {
					scaler.Compute.InstanceGroups.SetCoreGroup(&mrscaler.InstanceGroup{})
				}
				if scaler.Compute.InstanceGroups.CoreGroup.Capacity == nil {
					scaler.Compute.InstanceGroups.CoreGroup.SetCapacity(&mrscaler.InstanceGroupCapacity{})
				}
				scaler.Compute.InstanceGroups.CoreGroup.Capacity.SetMinimum(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(CoreMin)); ok {
				if scaler.Compute.InstanceGroups.CoreGroup == nil {
					scaler.Compute.InstanceGroups.SetCoreGroup(&mrscaler.InstanceGroup{})
				}
				if scaler.Compute.InstanceGroups.CoreGroup.Capacity == nil {
					scaler.Compute.InstanceGroups.CoreGroup.SetCapacity(&mrscaler.InstanceGroupCapacity{})
				}
				scaler.Compute.InstanceGroups.CoreGroup.Capacity.SetMinimum(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[CoreMax] = commons.NewGenericField(
		commons.MRScalerAWSCoreGroup,
		CoreMax,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *int = nil
			if scaler.Compute != nil && scaler.Compute.InstanceGroups != nil &&
				scaler.Compute.InstanceGroups.CoreGroup != nil &&
				scaler.Compute.InstanceGroups.CoreGroup.Capacity != nil &&
				scaler.Compute.InstanceGroups.CoreGroup.Capacity.Minimum != nil {
				value = scaler.Compute.InstanceGroups.CoreGroup.Capacity.Maximum
			}
			if err := resourceData.Set(string(CoreMax), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CoreMax), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(CoreMax)); ok {
				if scaler.Compute.InstanceGroups.CoreGroup == nil {
					scaler.Compute.InstanceGroups.SetCoreGroup(&mrscaler.InstanceGroup{})
				}
				if scaler.Compute.InstanceGroups.CoreGroup.Capacity == nil {
					scaler.Compute.InstanceGroups.CoreGroup.SetCapacity(&mrscaler.InstanceGroupCapacity{})
				}
				scaler.Compute.InstanceGroups.CoreGroup.Capacity.SetMaximum(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(CoreMax)); ok {
				if scaler.Compute.InstanceGroups.CoreGroup == nil {
					scaler.Compute.InstanceGroups.SetCoreGroup(&mrscaler.InstanceGroup{})
				}
				if scaler.Compute.InstanceGroups.CoreGroup.Capacity == nil {
					scaler.Compute.InstanceGroups.CoreGroup.SetCapacity(&mrscaler.InstanceGroupCapacity{})
				}
				scaler.Compute.InstanceGroups.CoreGroup.Capacity.SetMaximum(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[CoreTarget] = commons.NewGenericField(
		commons.MRScalerAWSCoreGroup,
		CoreTarget,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *int = nil
			if scaler.Compute != nil && scaler.Compute.InstanceGroups != nil &&
				scaler.Compute.InstanceGroups.CoreGroup != nil &&
				scaler.Compute.InstanceGroups.CoreGroup.Capacity != nil &&
				scaler.Compute.InstanceGroups.CoreGroup.Capacity.Minimum != nil {
				value = scaler.Compute.InstanceGroups.CoreGroup.Capacity.Minimum
			}
			if err := resourceData.Set(string(CoreTarget), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CoreTarget), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(CoreTarget)); ok {
				if scaler.Compute.InstanceGroups.CoreGroup == nil {
					scaler.Compute.InstanceGroups.SetCoreGroup(&mrscaler.InstanceGroup{})
				}
				if scaler.Compute.InstanceGroups.CoreGroup.Capacity == nil {
					scaler.Compute.InstanceGroups.CoreGroup.SetCapacity(&mrscaler.InstanceGroupCapacity{})
				}
				scaler.Compute.InstanceGroups.CoreGroup.Capacity.SetTarget(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(CoreTarget)); ok {
				if scaler.Compute.InstanceGroups.CoreGroup == nil {
					scaler.Compute.InstanceGroups.SetCoreGroup(&mrscaler.InstanceGroup{})
				}
				if scaler.Compute.InstanceGroups.CoreGroup.Capacity == nil {
					scaler.Compute.InstanceGroups.CoreGroup.SetCapacity(&mrscaler.InstanceGroupCapacity{})
				}
				scaler.Compute.InstanceGroups.CoreGroup.Capacity.SetTarget(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[CoreLifecycle] = commons.NewGenericField(
		commons.MRScalerAWSCoreGroup,
		CoreLifecycle,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Compute != nil && scaler.Compute.InstanceGroups != nil && scaler.Compute.InstanceGroups.CoreGroup != nil &&
				scaler.Compute.InstanceGroups.CoreGroup.LifeCycle != nil {
				value = scaler.Compute.InstanceGroups.CoreGroup.LifeCycle
			}
			if value != nil {
				if err := resourceData.Set(string(CoreLifecycle), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CoreLifecycle), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.Get(string(CoreLifecycle)).(string); ok && v != "" {
				if scaler.Compute.InstanceGroups.CoreGroup == nil {
					scaler.Compute.InstanceGroups.SetCoreGroup(&mrscaler.InstanceGroup{})
				}
				scaler.Compute.InstanceGroups.CoreGroup.SetLifeCycle(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(CoreLifecycle))
			return err
		},
		nil,
	)

	fieldsMap[CoreInstanceTypes] = commons.NewGenericField(
		commons.MRScalerAWSCoreGroup,
		CoreInstanceTypes,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(CoreInstanceTypes)); ok {
				instances := expandInstanceTypesList(v)
				if scaler.Compute.InstanceGroups.CoreGroup == nil {
					scaler.Compute.InstanceGroups.SetCoreGroup(&mrscaler.InstanceGroup{})
				}
				scaler.Compute.InstanceGroups.CoreGroup.SetInstanceTypes(instances)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(CoreInstanceTypes))
			return err
		},
		nil,
	)

	fieldsMap[CoreEBSOptimized] = commons.NewGenericField(
		commons.MRScalerAWSCoreGroup,
		CoreEBSOptimized,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOkExists(string(CoreEBSOptimized)); ok {
				if scaler.Compute.InstanceGroups.CoreGroup == nil {
					scaler.Compute.InstanceGroups.SetCoreGroup(&mrscaler.InstanceGroup{})
				}
				if scaler.Compute.InstanceGroups.CoreGroup.EBSConfiguration == nil {
					scaler.Compute.InstanceGroups.CoreGroup.SetEBSConfiguration(&mrscaler.EBSConfiguration{})
				}
				scaler.Compute.InstanceGroups.CoreGroup.EBSConfiguration.SetOptimized(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			optimized := false
			if v, ok := resourceData.GetOkExists(string(CoreEBSOptimized)); ok {
				if scaler.Compute.InstanceGroups.CoreGroup == nil {
					scaler.Compute.InstanceGroups.SetCoreGroup(&mrscaler.InstanceGroup{})
				}
				optimized = v.(bool)
			}
			if optimized {
				if scaler.Compute.InstanceGroups.CoreGroup.EBSConfiguration == nil {
					scaler.Compute.InstanceGroups.CoreGroup.SetEBSConfiguration(&mrscaler.EBSConfiguration{})
				}
				scaler.Compute.InstanceGroups.CoreGroup.EBSConfiguration.SetOptimized(spotinst.Bool(optimized))
			}
			return nil
		},
		nil,
	)

	fieldsMap[CoreEBSBlockDevice] = commons.NewGenericField(
		commons.MRScalerAWSCoreGroup,
		CoreEBSBlockDevice,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(VolumesPerInstance): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(VolumeType): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(SizeInGB): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(IOPS): {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var result []interface{}
			if scaler.Compute.InstanceGroups != nil && scaler.Compute.InstanceGroups.CoreGroup != nil &&
				scaler.Compute.InstanceGroups.CoreGroup.EBSConfiguration != nil &&
				scaler.Compute.InstanceGroups.CoreGroup.EBSConfiguration.BlockDeviceConfigs != nil {
				result = flattenMRscalerEBSBlockDevices(scaler.Compute.InstanceGroups.CoreGroup.EBSConfiguration.BlockDeviceConfigs)
			}
			if err := resourceData.Set(string(CoreEBSBlockDevice), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CoreEBSBlockDevice), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(CoreEBSBlockDevice)); ok {
				if scaler.Compute.InstanceGroups.CoreGroup == nil {
					scaler.Compute.InstanceGroups.SetCoreGroup(&mrscaler.InstanceGroup{})
				}
				if scaler.Compute.InstanceGroups.CoreGroup.EBSConfiguration == nil {
					scaler.Compute.InstanceGroups.CoreGroup.SetEBSConfiguration(&mrscaler.EBSConfiguration{})
				}

				if devices, err := expandScalerAWSBlockDevices(v); err != nil {
					return err
				} else {
					scaler.Compute.InstanceGroups.CoreGroup.EBSConfiguration.SetBlockDeviceConfigs(devices)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			value := []*mrscaler.BlockDeviceConfig{}
			if v, ok := resourceData.GetOk(string(CoreEBSBlockDevice)); ok {
				if scaler.Compute.InstanceGroups.CoreGroup == nil {
					scaler.Compute.InstanceGroups.SetCoreGroup(&mrscaler.InstanceGroup{})
				}
				if scaler.Compute.InstanceGroups.CoreGroup.EBSConfiguration == nil {
					scaler.Compute.InstanceGroups.CoreGroup.SetEBSConfiguration(&mrscaler.EBSConfiguration{})
				}

				if devices, err := expandScalerAWSBlockDevices(v); err != nil {
					return err
				} else {
					value = devices
				}
			}
			if scaler.Compute.InstanceGroups.CoreGroup.EBSConfiguration != nil {
				scaler.Compute.InstanceGroups.CoreGroup.EBSConfiguration.SetBlockDeviceConfigs(value)
			}
			return nil
		},
		nil,
	)
}
