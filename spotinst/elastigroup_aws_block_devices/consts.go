package elastigroup_aws_block_devices

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "ebs_block_device"
)

const (
	// - COMMON -----------------------------------------------------
	DeviceName commons.FieldName = "device_name"
	// --------------------------------------------------------------

	// - EBS --------------------------------------------------------
	EbsBlockDevice      commons.FieldName = "ebs_block_device"
	SnapshotId          commons.FieldName = "snapshot_id"
	VolumeType          commons.FieldName = "volume_type"
	VolumeSize          commons.FieldName = "volume_size"
	Iops                commons.FieldName = "iops"
	DeleteOnTermination commons.FieldName = "delete_on_termination"
	Encrypted           commons.FieldName = "encrypted"
	KmsKeyId            commons.FieldName = "kms_key_id"
	NoDevice            commons.FieldName = "nodevice"
	Throughput          commons.FieldName = "throughput"

	EphemeralBlockDevice commons.FieldName = "ephemeral_block_device"
	VirtualName          commons.FieldName = "virtual_name"

	DynamicVolumeSize   commons.FieldName = "dynamic_volume_size"
	BaseSize            commons.FieldName = "base_size"
	Resource            commons.FieldName = "resource"
	SizePerResourceUnit commons.FieldName = "size_per_resource_unit"
)
const (
	DynamicIops             commons.FieldName = "dynamic_iops"
	IopsBaseSize            commons.FieldName = "base_size"
	IopsResource            commons.FieldName = "resource"
	IopsSizePerResourceUnit commons.FieldName = "size_per_resource_unit"
)
