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
func SetupTaskGroup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[TaskMin] = commons.NewGenericField(
		commons.MRScalerAWSTaskGroup,
		TaskMin,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.Get(string(TaskMin)).(int); ok && v >= 0 {
				if scaler.Compute.InstanceGroups.TaskGroup.Capacity == nil {
					scaler.Compute.InstanceGroups.TaskGroup.SetCapacity(&mrscaler.InstanceGroupCapacity{})
				}
				scaler.Compute.InstanceGroups.TaskGroup.Capacity.SetMinimum(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.Get(string(TaskMin)).(int); ok && v >= 0 {
				if scaler.Compute.InstanceGroups.TaskGroup.Capacity == nil {
					scaler.Compute.InstanceGroups.TaskGroup.SetCapacity(&mrscaler.InstanceGroupCapacity{})
				}
				scaler.Compute.InstanceGroups.TaskGroup.Capacity.SetMinimum(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[TaskMax] = commons.NewGenericField(
		commons.MRScalerAWSTaskGroup,
		TaskMax,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.Get(string(TaskMax)).(int); ok && v >= 0 {
				if scaler.Compute.InstanceGroups.TaskGroup.Capacity == nil {
					scaler.Compute.InstanceGroups.TaskGroup.SetCapacity(&mrscaler.InstanceGroupCapacity{})
				}
				scaler.Compute.InstanceGroups.TaskGroup.Capacity.SetMaximum(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.Get(string(TaskMax)).(int); ok && v >= 0 {
				if scaler.Compute.InstanceGroups.TaskGroup.Capacity == nil {
					scaler.Compute.InstanceGroups.TaskGroup.SetCapacity(&mrscaler.InstanceGroupCapacity{})
				}
				scaler.Compute.InstanceGroups.TaskGroup.Capacity.SetMaximum(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[TaskTarget] = commons.NewGenericField(
		commons.MRScalerAWSTaskGroup,
		TaskTarget,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.Get(string(TaskTarget)).(int); ok && v >= 0 {
				if scaler.Compute.InstanceGroups.TaskGroup.Capacity == nil {
					scaler.Compute.InstanceGroups.TaskGroup.SetCapacity(&mrscaler.InstanceGroupCapacity{})
				}
				scaler.Compute.InstanceGroups.TaskGroup.Capacity.SetTarget(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.Get(string(TaskTarget)).(int); ok && v >= 0 {
				if scaler.Compute.InstanceGroups.TaskGroup.Capacity == nil {
					scaler.Compute.InstanceGroups.TaskGroup.SetCapacity(&mrscaler.InstanceGroupCapacity{})
				}
				scaler.Compute.InstanceGroups.TaskGroup.Capacity.SetTarget(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[TaskLifecycle] = commons.NewGenericField(
		commons.MRScalerAWSTaskGroup,
		TaskLifecycle,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.Get(string(TaskLifecycle)).(string); ok && v != "" {
				scaler.Compute.InstanceGroups.TaskGroup.SetLifeCycle(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(TaskLifecycle))
			return err
		},
		nil,
	)

	fieldsMap[TaskInstanceTypes] = commons.NewGenericField(
		commons.MRScalerAWSTaskGroup,
		TaskInstanceTypes,
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
			if v, ok := resourceData.GetOk(string(TaskInstanceTypes)); ok {
				instances := expandInstanceTypesList(v)
				scaler.Compute.InstanceGroups.TaskGroup.SetInstanceTypes(instances)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(TaskInstanceTypes))
			return err
		},
		nil,
	)

	fieldsMap[TaskEBSOptimized] = commons.NewGenericField(
		commons.MRScalerAWSTaskGroup,
		TaskEBSOptimized,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOkExists(string(TaskEBSOptimized)); ok {
				if scaler.Compute.InstanceGroups.TaskGroup == nil {
					scaler.Compute.InstanceGroups.SetTaskGroup(&mrscaler.InstanceGroup{})
				}
				if scaler.Compute.InstanceGroups.TaskGroup.EBSConfiguration == nil {
					scaler.Compute.InstanceGroups.TaskGroup.SetEBSConfiguration(&mrscaler.EBSConfiguration{})
				}
				scaler.Compute.InstanceGroups.TaskGroup.EBSConfiguration.SetOptimized(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			optimized := false
			if v, ok := resourceData.GetOkExists(string(TaskEBSOptimized)); ok {
				if scaler.Compute.InstanceGroups.TaskGroup == nil {
					scaler.Compute.InstanceGroups.SetTaskGroup(&mrscaler.InstanceGroup{})
				}
				optimized = v.(bool)
			}
			if optimized {
				if scaler.Compute.InstanceGroups.TaskGroup.EBSConfiguration == nil {
					scaler.Compute.InstanceGroups.TaskGroup.SetEBSConfiguration(&mrscaler.EBSConfiguration{})
				}
				scaler.Compute.InstanceGroups.TaskGroup.EBSConfiguration.SetOptimized(spotinst.Bool(optimized))
			}
			return nil
		},
		nil,
	)

	fieldsMap[TaskEBSBlockDevice] = commons.NewGenericField(
		commons.MRScalerAWSTaskGroup,
		TaskEBSBlockDevice,
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
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(TaskEBSBlockDevice)); ok {
				if scaler.Compute.InstanceGroups.TaskGroup == nil {
					scaler.Compute.InstanceGroups.SetTaskGroup(&mrscaler.InstanceGroup{})
				}
				if scaler.Compute.InstanceGroups.TaskGroup.EBSConfiguration == nil {
					scaler.Compute.InstanceGroups.TaskGroup.SetEBSConfiguration(&mrscaler.EBSConfiguration{})
				}

				if devices, err := expandScalerAWSBlockDevices(v); err != nil {
					return err
				} else {
					scaler.Compute.InstanceGroups.TaskGroup.EBSConfiguration.SetBlockDeviceConfigs(devices)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			value := []*mrscaler.BlockDeviceConfig{}
			if v, ok := resourceData.GetOk(string(TaskEBSBlockDevice)); ok {
				if scaler.Compute.InstanceGroups.TaskGroup == nil {
					scaler.Compute.InstanceGroups.SetTaskGroup(&mrscaler.InstanceGroup{})
				}
				if scaler.Compute.InstanceGroups.TaskGroup.EBSConfiguration == nil {
					scaler.Compute.InstanceGroups.TaskGroup.SetEBSConfiguration(&mrscaler.EBSConfiguration{})
				}

				if devices, err := expandScalerAWSBlockDevices(v); err != nil {
					return err
				} else {
					value = devices
				}
			}
			if scaler.Compute.InstanceGroups.TaskGroup.EBSConfiguration != nil {
				scaler.Compute.InstanceGroups.TaskGroup.EBSConfiguration.SetBlockDeviceConfigs(value)
			}
			return nil
		},
		nil,
	)
}
