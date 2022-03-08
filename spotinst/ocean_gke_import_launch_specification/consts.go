package ocean_gke_import_launch_specification

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	RootVolumeType commons.FieldName = "root_volume_type"
)

const (
	ShieldedInstanceConfig    commons.FieldName = "shielded_instance_config"
	EnableSecureBoot          commons.FieldName = "enable_secure_boot"
	EnableIntegrityMonitoring commons.FieldName = "enable_integrity_monitoring"
)
