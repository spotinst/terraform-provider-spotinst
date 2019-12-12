package mrscaler_aws_instance_groups

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/mrscaler"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	SetupCoreGroup(fieldsMap)
	SetupMasterGroup(fieldsMap)
	SetupTaskGroup(fieldsMap)
}

func expandScalerAWSBlockDevices(data interface{}) ([]*mrscaler.BlockDeviceConfig, error) {
	list := data.(*schema.Set).List()
	devices := make([]*mrscaler.BlockDeviceConfig, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		blockDevice := &mrscaler.BlockDeviceConfig{VolumeSpecification: &mrscaler.VolumeSpecification{}}

		if v, ok := m[string(VolumesPerInstance)].(int); ok {
			blockDevice.SetVolumesPerInstance(spotinst.Int(v))
		}

		if v, ok := m[string(VolumeType)].(string); ok {
			blockDevice.VolumeSpecification.SetVolumeType(spotinst.String(v))
		}

		if v, ok := m[string(SizeInGB)].(int); ok {
			blockDevice.VolumeSpecification.SetSizeInGB(spotinst.Int(v))
		}

		if v, ok := m[string(IOPS)].(int); ok && v > 0 {
			blockDevice.VolumeSpecification.SetIOPS(spotinst.Int(v))
		}

		devices = append(devices, blockDevice)
	}

	return devices, nil
}

func expandInstanceTypesList(instanceTypes interface{}) []string {
	list := instanceTypes.([]interface{})
	types := make([]string, 0, len(list))
	for _, str := range list {
		if typ, ok := str.(string); ok {
			types = append(types, typ)
		}
	}
	return types
}

func flattenMRscalerEBSBlockDevices(devices []*mrscaler.BlockDeviceConfig) []interface{} {
	result := make([]interface{}, 0, len(devices))
	for _, dev := range devices {
		m := make(map[string]interface{})

		m[string(VolumesPerInstance)] = spotinst.IntValue(dev.VolumesPerInstance)

		if dev.VolumeSpecification != nil {
			m[string(SizeInGB)] = spotinst.IntValue(dev.VolumeSpecification.SizeInGB)
			m[string(VolumeType)] = spotinst.StringValue(dev.VolumeSpecification.VolumeType)
			m[string(IOPS)] = spotinst.IntValue(dev.VolumeSpecification.IOPS)
		}
		result = append(result, m)
	}
	return result
}
