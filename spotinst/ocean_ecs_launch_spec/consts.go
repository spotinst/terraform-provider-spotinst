package ocean_ecs_launch_spec

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

type LabelField string
type IAMField string
type TagField string

const (
	CPUPerUnit    commons.FieldName = "cpu_per_unit"
	MemoryPerUnit commons.FieldName = "memory_per_unit"
	NumOfUnits    commons.FieldName = "num_of_units"
)

const (
	OceanID            commons.FieldName = "ocean_id"
	ImageID            commons.FieldName = "image_id"
	IamInstanceProfile commons.FieldName = "iam_instance_profile"
	UserData           commons.FieldName = "user_data"
	SecurityGroupIds   commons.FieldName = "security_group_ids"
	Name               commons.FieldName = "name"
	Attributes         commons.FieldName = "attributes"
	AttributeKey       commons.FieldName = "key"
	AttributeValue     commons.FieldName = "value"
	AutoscaleHeadrooms commons.FieldName = "autoscale_headrooms"
	Tags               commons.FieldName = "tags"
	InstanceTypes      commons.FieldName = "instance_types"
)

const (
	TagKey   TagField = "key"
	TagValue TagField = "value"
)

const (
	LabelKey   LabelField = "key"
	LabelValue LabelField = "value"
)

const (
	BlockDeviceMappings commons.FieldName = "block_device_mappings"
	DeviceName          commons.FieldName = "device_name"
	Ebs                 commons.FieldName = "ebs"
	DeleteOnTermination commons.FieldName = "delete_on_termination"
	Encrypted           commons.FieldName = "encrypted"
	IOPS                commons.FieldName = "iops"
	KMSKeyID            commons.FieldName = "kms_key_id"
	SnapshotID          commons.FieldName = "snapshot_id"
	VolumeSize          commons.FieldName = "volume_size"
	DynamicVolumeSize   commons.FieldName = "dynamic_volume_size"
	BaseSize            commons.FieldName = "base_size"
	Resource            commons.FieldName = "resource"
	SizePerResourceUnit commons.FieldName = "size_per_resource_unit"
	VolumeType          commons.FieldName = "volume_type"
	NoDevice            commons.FieldName = "no_device"
	VirtualName         commons.FieldName = "virtual_name"
)
