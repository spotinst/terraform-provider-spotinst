package ocean_aws_launch_spec

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

type LabelField string
type TaintField string
type IAMField string
type TagField string

const (
	LabelKey   LabelField = "key"
	LabelValue LabelField = "value"
)

const (
	TaintKey   TaintField = "key"
	TaintValue TaintField = "value"
	Effect     TaintField = "effect"
)

const (
	CPUPerUnit    commons.FieldName = "cpu_per_unit"
	GPUPerUnit    commons.FieldName = "gpu_per_unit"
	MemoryPerUnit commons.FieldName = "memory_per_unit"
	NumOfUnits    commons.FieldName = "num_of_units"
)

const (
	TagKey   TagField = "key"
	TagValue TagField = "value"
)

const (
	SecurityGroups           commons.FieldName = "security_groups"
	Name                     commons.FieldName = "name"
	OceanID                  commons.FieldName = "ocean_id"
	ImageID                  commons.FieldName = "image_id"
	UserData                 commons.FieldName = "user_data"
	IamInstanceProfile       commons.FieldName = "iam_instance_profile"
	Labels                   commons.FieldName = "labels"
	Taints                   commons.FieldName = "taints"
	AutoscaleHeadrooms       commons.FieldName = "autoscale_headrooms"
	SubnetIDs                commons.FieldName = "subnet_ids"
	InstanceTypes            commons.FieldName = "instance_types"
	PreferredSpotTypes       commons.FieldName = "preferred_spot_types"
	RootVolumeSize           commons.FieldName = "root_volume_size"
	Tags                     commons.FieldName = "tags"
	ElasticIpPool            commons.FieldName = "elastic_ip_pool"
	TagSelector              commons.FieldName = "tag_selector"
	TagSelectorKey           commons.FieldName = "tag_key"
	TagSelectorValue         commons.FieldName = "tag_value"
	ResourceLimits           commons.FieldName = "resource_limits"
	MaxInstanceCount         commons.FieldName = "max_instance_count"
	Strategy                 commons.FieldName = "strategy"
	AssociatePublicIPAddress commons.FieldName = "associate_public_ip_address"
	RestrictScaleDown        commons.FieldName = "restrict_scale_down"
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
	Throughput          commons.FieldName = "throughput"
)

const (
	SpotPercentage commons.FieldName = "spot_percentage"
)

const (
	CreateOptions commons.FieldName = "create_options"
	InitialNodes  commons.FieldName = "initial_nodes"
)
