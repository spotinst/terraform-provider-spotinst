package stateful_node_azure_launch_spec

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

// LaunchSpec
const (
	CustomData     commons.FieldName = "custom_data"
	ShutdownScript commons.FieldName = "shutdown_script"
)

// Tags
const (
	Tags     commons.FieldName = "tags"
	TagKey   commons.FieldName = "tag_key"
	TagValue commons.FieldName = "tag_value"
)

// Managed service identities
const (
	ManagedServiceIdentities commons.FieldName = "managed_service_identities"
	Name                     commons.FieldName = "name"
	ResourceGroupName        commons.FieldName = "resource_group_name"
)

// OS disk
const (
	OSDisk commons.FieldName = "os_disk"
	SizeGB commons.FieldName = "size_gb"
	Type   commons.FieldName = "type"
)

// Data disk
const (
	DataDisk commons.FieldName = "data_disk"
	LUN      commons.FieldName = "lun"
)

// Boot diagnostics
const (
	BootDiagnostics commons.FieldName = "boot_diagnostics"
	IsEnabled       commons.FieldName = "is_enabled"
	StorageURL      commons.FieldName = "storage_url"
)
