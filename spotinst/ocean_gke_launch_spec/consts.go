package ocean_gke_launch_spec

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

type LabelField string
type MetadataField string

const (
	LabelKey   LabelField = "key"
	LabelValue LabelField = "value"

	MetadataKey   MetadataField = "key"
	MetadataValue MetadataField = "value"

	TaintKey    MetadataField = "key"
	TaintValue  MetadataField = "value"
	TaintEffect MetadataField = "effect"
)

const (
	CPUPerUnit    commons.FieldName = "cpu_per_unit"
	GPUPerUnit    commons.FieldName = "gpu_per_unit"
	MemoryPerUnit commons.FieldName = "memory_per_unit"
	NumOfUnits    commons.FieldName = "num_of_units"
)

const (
	OceanId                   commons.FieldName = "ocean_id"
	Name                      commons.FieldName = "name"
	SourceImage               commons.FieldName = "source_image"
	Metadata                  commons.FieldName = "metadata"
	Labels                    commons.FieldName = "labels"
	Taints                    commons.FieldName = "taints"
	AutoscaleHeadrooms        commons.FieldName = "autoscale_headrooms"
	RestrictScaleDown         commons.FieldName = "restrict_scale_down"
	RootVolumeType            commons.FieldName = "root_volume_type"
	RootVolumeSizeInGB        commons.FieldName = "root_volume_size"
	InstanceTypes             commons.FieldName = "instance_types"
	ShieldedInstanceConfig    commons.FieldName = "shielded_instance_config"
	EnableSecureBoot          commons.FieldName = "enable_secure_boot"
	EnableIntegrityMonitoring commons.FieldName = "enable_integrity_monitoring"
	Storage                   commons.FieldName = "storage"
	LocalSSDCount             commons.FieldName = "local_ssd_count"
	ServiceAccount            commons.FieldName = "service_account"
	Tags                      commons.FieldName = "tags"
	ResourceLimits            commons.FieldName = "resource_limits"
	MaxInstanceCount          commons.FieldName = "max_instance_count"
	MinInstanceCount          commons.FieldName = "min_instance_count"
)

const (
	AutoscaleHeadroomsAutomatic commons.FieldName = "autoscale_headrooms_automatic"
	AutoHeadroomPercentage      commons.FieldName = "auto_headroom_percentage"
)

const (
	NodePoolName commons.FieldName = "node_pool_name"
)

const (
	UpdatePolicy commons.FieldName = "update_policy"
	ShouldRoll   commons.FieldName = "should_roll"

	RollConfig          commons.FieldName = "roll_config"
	BatchSizePercentage commons.FieldName = "batch_size_percentage"
)

const (
	NetworkInterfaces commons.FieldName = "network_interface"
	Network           commons.FieldName = "network"
	ProjectId         commons.FieldName = "project_id"

	LaunchSpecAccessConfigs     commons.FieldName = "access_configs"
	LaunchSpecAccessConfigsName commons.FieldName = "name"
	Type                        commons.FieldName = "type"

	LaunchSpecAliasIPRanges commons.FieldName = "alias_ip_ranges"
	IPCidrRange             commons.FieldName = "ip_cidr_range"
	SubnetworkRangeName     commons.FieldName = "subnetwork_range_name"
)
