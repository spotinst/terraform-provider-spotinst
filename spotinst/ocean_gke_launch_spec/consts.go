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
	ResourceLimits            commons.FieldName = "resource_limits"
	MaxInstanceCount          commons.FieldName = "max_instance_count"
)

const (
	NodePoolName commons.FieldName = "node_pool_name"
)
