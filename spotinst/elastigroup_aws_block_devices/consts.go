package elastigroup_aws_block_devices

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

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
	// --------------------------------------------------------------

	// - EPHEMERAL --------------------------------------------------
	EphemeralBlockDevice commons.FieldName = "ephemeral_block_device"
	VirtualName          commons.FieldName = "virtual_name"
	// --------------------------------------------------------------
)
