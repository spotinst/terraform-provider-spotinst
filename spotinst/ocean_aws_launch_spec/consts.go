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
	AutoscaleHeadroomsAutomatic commons.FieldName = "autoscale_headrooms_automatic"
	AutoHeadroomPercentage      commons.FieldName = "auto_headroom_percentage"
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
	MinInstanceCount         commons.FieldName = "min_instance_count"
	Strategy                 commons.FieldName = "strategy"
	AssociatePublicIPAddress commons.FieldName = "associate_public_ip_address"
	RestrictScaleDown        commons.FieldName = "restrict_scale_down"
	SchedulingTask           commons.FieldName = "scheduling_task"
	SchedulingShutdownHours  commons.FieldName = "scheduling_shutdown_hours"
	AutoscaleDown            commons.FieldName = "autoscale_down"
	MaxScaleDownPercentage   commons.FieldName = "max_scale_down_percentage"
	Images                   commons.FieldName = "images"
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

const (
	UpdatePolicy commons.FieldName = "update_policy"
	ShouldRoll   commons.FieldName = "should_roll"

	RollConfig          commons.FieldName = "roll_config"
	BatchSizePercentage commons.FieldName = "batch_size_percentage"
)

const (
	DeleteOptions commons.FieldName = "delete_options"
	ForceDelete   commons.FieldName = "force_delete"
	DeleteNodes   commons.FieldName = "delete_nodes"
)

const (
	IsEnabled      commons.FieldName = "is_enabled"
	CronExpression commons.FieldName = "cron_expression"
	TaskType       commons.FieldName = "task_type"
	TaskHeadroom   commons.FieldName = "task_headroom"
)

const (
	TimeWindows commons.FieldName = "time_windows"
)

const (
	ImageId commons.FieldName = "image_id"
)

const (
	InstanceMetadataOptions commons.FieldName = "instance_metadata_options"
	HTTPTokens              commons.FieldName = "http_tokens"
	HTTPPutResponseHopLimit commons.FieldName = "http_put_response_hop_limit"
)
const (
	InstanceTypesFilters  commons.FieldName = "instance_types_filters"
	Categories            commons.FieldName = "categories"
	DiskTypes             commons.FieldName = "disk_types"
	ExcludeFamilies       commons.FieldName = "exclude_families"
	ExcludeMetal          commons.FieldName = "exclude_metal"
	Hypervisor            commons.FieldName = "hypervisor"
	IncludeFamilies       commons.FieldName = "include_families"
	IsEnaSupported        commons.FieldName = "is_ena_supported"
	MaxGpu                commons.FieldName = "max_gpu"
	MaxMemoryGiB          commons.FieldName = "max_memory_gib"
	MaxNetworkPerformance commons.FieldName = "max_network_performance"
	MaxVcpu               commons.FieldName = "max_vcpu"
	MinEnis               commons.FieldName = "min_enis"
	MinGpu                commons.FieldName = "min_gpu"
	MinMemoryGiB          commons.FieldName = "min_memory_gib"
	MinNetworkPerformance commons.FieldName = "min_network_performance"
	MinVcpu               commons.FieldName = "min_vcpu"
	RootDeviceTypes       commons.FieldName = "root_device_types"
	VirtualizationTypes   commons.FieldName = "virtualization_types"
)
const (
	EphemeralStorage           commons.FieldName = "ephemeral_storage"
	EphemeralStorageDeviceName commons.FieldName = "ephemeral_storage_device_name"
)
