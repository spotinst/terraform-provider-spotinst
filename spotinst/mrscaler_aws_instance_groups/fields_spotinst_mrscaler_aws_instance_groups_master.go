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
func SetupMasterGroup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[MasterLifecycle] = commons.NewGenericField(
		commons.MRScalerAWSMasterGroup,
		MasterLifecycle,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			//ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Compute.InstanceGroups != nil && scaler.Compute.InstanceGroups.MasterGroup != nil &&
				scaler.Compute.InstanceGroups.MasterGroup.LifeCycle != nil {
				value = scaler.Compute.InstanceGroups.MasterGroup.LifeCycle
			}
			if err := resourceData.Set(string(MasterLifecycle), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MasterLifecycle), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.Get(string(MasterLifecycle)).(string); ok && v != "" {
				if scaler.Compute.InstanceGroups.MasterGroup == nil {
					scaler.Compute.InstanceGroups.SetMasterGroup(&mrscaler.InstanceGroup{Target: spotinst.Int(1)})
				}
				scaler.Compute.InstanceGroups.MasterGroup.SetLifeCycle(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(MasterLifecycle))
			return err
		},
		nil,
	)

	fieldsMap[MasterInstanceTypes] = commons.NewGenericField(
		commons.MRScalerAWSMasterGroup,
		MasterInstanceTypes,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			//ForceNew: true,
			Elem: &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var result []string
			if scaler.Compute.InstanceGroups != nil && scaler.Compute.InstanceGroups.MasterGroup != nil &&
				scaler.Compute.InstanceGroups.MasterGroup.InstanceTypes != nil {
				result = append(result, scaler.Compute.InstanceGroups.MasterGroup.InstanceTypes...)
			}
			if err := resourceData.Set(string(MasterInstanceTypes), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MasterInstanceTypes), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(MasterInstanceTypes)); ok {
				instances := expandInstanceTypesList(v)
				if scaler.Compute.InstanceGroups.MasterGroup == nil {
					scaler.Compute.InstanceGroups.SetMasterGroup(&mrscaler.InstanceGroup{Target: spotinst.Int(1)})
				}
				scaler.Compute.InstanceGroups.MasterGroup.SetInstanceTypes(instances)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(MasterInstanceTypes))
			return err
		},
		nil,
	)

	fieldsMap[MasterEBSOptimized] = commons.NewGenericField(
		commons.MRScalerAWSMasterGroup,
		MasterEBSOptimized,
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
			if v, ok := resourceData.GetOkExists(string(MasterEBSOptimized)); ok {
				if scaler.Compute.InstanceGroups.MasterGroup == nil {
					scaler.Compute.InstanceGroups.SetMasterGroup(&mrscaler.InstanceGroup{Target: spotinst.Int(1)})
				}
				if scaler.Compute.InstanceGroups.MasterGroup.EBSConfiguration == nil {
					scaler.Compute.InstanceGroups.MasterGroup.SetEBSConfiguration(&mrscaler.EBSConfiguration{})
				}
				scaler.Compute.InstanceGroups.MasterGroup.EBSConfiguration.SetOptimized(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			optimized := false
			if v, ok := resourceData.GetOkExists(string(MasterEBSOptimized)); ok {
				if scaler.Compute.InstanceGroups.MasterGroup == nil {
					scaler.Compute.InstanceGroups.SetMasterGroup(&mrscaler.InstanceGroup{})
				}
				optimized = v.(bool)
			}
			if optimized {
				if scaler.Compute.InstanceGroups.MasterGroup.EBSConfiguration == nil {
					scaler.Compute.InstanceGroups.MasterGroup.SetEBSConfiguration(&mrscaler.EBSConfiguration{})
				}
				scaler.Compute.InstanceGroups.MasterGroup.EBSConfiguration.SetOptimized(spotinst.Bool(optimized))
			}
			return nil
		},
		nil,
	)

	fieldsMap[MasterEBSBlockDevice] = commons.NewGenericField(
		commons.MRScalerAWSMasterGroup,
		MasterEBSBlockDevice,
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
			if scaler.Compute.InstanceGroups != nil && scaler.Compute.InstanceGroups.MasterGroup != nil &&
				scaler.Compute.InstanceGroups.MasterGroup.EBSConfiguration != nil &&
				scaler.Compute.InstanceGroups.MasterGroup.EBSConfiguration.BlockDeviceConfigs != nil {
				result = flattenMRscalerEBSBlockDevices(scaler.Compute.InstanceGroups.MasterGroup.EBSConfiguration.BlockDeviceConfigs)
			}
			if err := resourceData.Set(string(MasterEBSBlockDevice), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MasterEBSBlockDevice), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(MasterEBSBlockDevice)); ok {
				if scaler.Compute.InstanceGroups.MasterGroup == nil {
					scaler.Compute.InstanceGroups.SetMasterGroup(&mrscaler.InstanceGroup{Target: spotinst.Int(1)})
				}
				if scaler.Compute.InstanceGroups.MasterGroup.EBSConfiguration == nil {
					scaler.Compute.InstanceGroups.MasterGroup.SetEBSConfiguration(&mrscaler.EBSConfiguration{})
				}

				if devices, err := expandScalerAWSBlockDevices(v); err != nil {
					return err
				} else {
					scaler.Compute.InstanceGroups.MasterGroup.EBSConfiguration.SetBlockDeviceConfigs(devices)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(MasterInstanceTypes))
			return err
		},
		nil,
	)
}
