package stateful_node_azure_launch_spec

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

// LaunchSpec
const (
	CustomData     commons.FieldName = "custom_data"
	ShutdownScript commons.FieldName = "shutdown_script"
)

// Tags
const (
	Tags     commons.FieldName = "tags" // TODO - TypeList?
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
	SizeGB commons.FieldName = "size_gb" // TODO - duplicate with data disk - might be an issue?
	Type   commons.FieldName = "type"
)

// Data disk
const (
	DataDisks commons.FieldName = "data_disks" // TODO - should this be of type list?
	LUN       commons.FieldName = "lun"
)

// Boot diagnostics
const (
	BootDiagnostics commons.FieldName = "boot_diagnostics"
	IsEnabled       commons.FieldName = "is_enabled"
	StorageURL      commons.FieldName = "storage_url"
)
