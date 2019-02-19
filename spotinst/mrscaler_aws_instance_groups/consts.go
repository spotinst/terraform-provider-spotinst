package mrscaler_aws_instance_groups

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	TaskInstanceTypes  commons.FieldName = "task_instance_types"
	TaskMin            commons.FieldName = "task_min_size"
	TaskMax            commons.FieldName = "task_max_size"
	TaskTarget         commons.FieldName = "task_desired_capacity"
	TaskLifecycle      commons.FieldName = "task_lifecycle"
	TaskEBSBlockDevice commons.FieldName = "task_ebs_block_device"
	TaskEBSOptimized   commons.FieldName = "task_ebs_optimized"

	MasterInstanceTypes  commons.FieldName = "master_instance_types"
	MasterLifecycle      commons.FieldName = "master_lifecycle"
	MasterEBSBlockDevice commons.FieldName = "master_ebs_block_device"
	MasterEBSOptimized   commons.FieldName = "master_ebs_optimized"

	CoreInstanceTypes  commons.FieldName = "core_instance_types"
	CoreMin            commons.FieldName = "core_min_size"
	CoreMax            commons.FieldName = "core_max_size"
	CoreTarget         commons.FieldName = "core_desired_capacity"
	CoreLifecycle      commons.FieldName = "core_lifecycle"
	CoreEBSBlockDevice commons.FieldName = "core_ebs_block_device"
	CoreEBSOptimized   commons.FieldName = "core_ebs_optimized"

	VolumesPerInstance commons.FieldName = "volumes_per_instance"
	VolumeType         commons.FieldName = "volume_type"
	SizeInGB           commons.FieldName = "size_in_gb"
	IOPS               commons.FieldName = "iops"
)
