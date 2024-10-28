package elastigroup_azure_launchspecification

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

type TagField string

const (
	TagKey   TagField = "key"
	TagValue TagField = "value"
)

const (
	CustomData     commons.FieldName = "custom_data"
	ShutdownScript commons.FieldName = "shutdown_script"
	UserData       commons.FieldName = "user_data"
	VmNamePrefix   commons.FieldName = "vm_name_prefix"
)

const (
	ManagedServiceIdentity                  commons.FieldName = "managed_service_identity"
	ManagedServiceIdentityResourceGroupName commons.FieldName = "resource_group_name"
	ManagedServiceIdentityName              commons.FieldName = "name"
	Tags                                    commons.FieldName = "tags"
)

const (
	DataDisk       commons.FieldName = "data_disk"
	DataDiskLUN    commons.FieldName = "lun"
	DataDiskSizeGB commons.FieldName = "size_gb"
	DataDiskType   commons.FieldName = "type"
)

const (
	OsDisk       commons.FieldName = "os_disk"
	OsDiskSizeGB commons.FieldName = "size_gb"
	OsDiskType   commons.FieldName = "type"
)

const (
	BootDiagnostics           commons.FieldName = "boot_diagnostics"
	BootDiagnosticsIsEnabled  commons.FieldName = "is_enabled"
	BootDiagnosticsStorageURL commons.FieldName = "storage_url"
	BootDiagnosticsType       commons.FieldName = "type"
)

const (
	Security                     commons.FieldName = "security"
	SecureBootEnabled            commons.FieldName = "secure_boot_enabled"
	SecurityType                 commons.FieldName = "security_type"
	VTpmEnabled                  commons.FieldName = "vtpm_enabled"
	ConfidentialOsDiskEncryption commons.FieldName = "confidential_os_disk_encryption"
)

const (
	ProximityPlacementGroups commons.FieldName = "proximity_placement_groups"
	PPGName                  commons.FieldName = "name"
	PPGResourceGroupName     commons.FieldName = "resource_group_name"
)
