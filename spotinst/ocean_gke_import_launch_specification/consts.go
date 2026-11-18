package ocean_gke_import_launch_specification

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	RootVolumeType    commons.FieldName = "root_volume_type"
	UseAsTemplateOnly commons.FieldName = "use_as_template_only"
)

const (
	ShieldedInstanceConfig    commons.FieldName = "shielded_instance_config"
	EnableSecureBoot          commons.FieldName = "enable_secure_boot"
	EnableIntegrityMonitoring commons.FieldName = "enable_integrity_monitoring"
)

const (
	DynamicVolumeSize   commons.FieldName = "dynamic_volume_size"
	BaseSize            commons.FieldName = "base_size"
	Resource            commons.FieldName = "resource"
	SizePerResourceUnit commons.FieldName = "size_per_resource_unit"
)
