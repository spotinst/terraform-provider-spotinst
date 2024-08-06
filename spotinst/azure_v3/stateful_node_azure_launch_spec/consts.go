package stateful_node_azure_launch_spec

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

// LaunchSpec
const (
	CustomData     commons.FieldName = "custom_data"
	UserData       commons.FieldName = "user_data"
	ShutdownScript commons.FieldName = "shutdown_script"
	VMName         commons.FieldName = "vm_name"
	VMNamePrefix   commons.FieldName = "vm_name_prefix"
	LicenseType    commons.FieldName = "license_type"
)

// Tags
const (
	Tag      commons.FieldName = "tag"
	TagKey   commons.FieldName = "tag_key"
	TagValue commons.FieldName = "tag_value"
)

// Managed service identities
const (
	ManagedServiceIdentities commons.FieldName = "managed_service_identities"
	Name                     commons.FieldName = "name"
	ResourceGroupName        commons.FieldName = "resource_group_name"
)

// Proximity Placement Groups
const (
	ProximityPlacementGroups commons.FieldName = "proximity_placement_groups"
	PPGName                  commons.FieldName = "name"
	PPGResourceGroupName     commons.FieldName = "resource_group_name"
)

// OS disk
const (
	OSDisk       commons.FieldName = "os_disk"
	OSDiskSizeGB commons.FieldName = "size_gb"
	OSDiskType   commons.FieldName = "type"
	OSCaching    commons.FieldName = "caching"
)

// Data disk
const (
	DataDisk       commons.FieldName = "data_disk"
	DataDiskSizeGB commons.FieldName = "size_gb"
	DataDiskType   commons.FieldName = "type"
	DataDiskLUN    commons.FieldName = "lun"
)

// Boot diagnostics
const (
	BootDiagnostics           commons.FieldName = "boot_diagnostics"
	BootDiagnosticsIsEnabled  commons.FieldName = "is_enabled"
	BootDiagnosticsStorageURL commons.FieldName = "storage_url"
	BootDiagnosticsType       commons.FieldName = "type"
)

// Security
const (
	Security                     commons.FieldName = "security"
	SecureBootEnabled            commons.FieldName = "secure_boot_enabled"
	SecurityType                 commons.FieldName = "security_type"
	VTpmEnabled                  commons.FieldName = "vtpm_enabled"
	EncryptionAtHost             commons.FieldName = "encryption_at_host"
	ConfidentialOsDiskEncryption commons.FieldName = "confidential_os_disk_encryption"
)
