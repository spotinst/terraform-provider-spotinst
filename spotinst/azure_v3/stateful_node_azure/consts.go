package stateful_node_azure

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

//LaunchSpec
const (
	Name              commons.FieldName = "name"
	ResourceGroupName commons.FieldName = "resource_group_name"
	Region            commons.FieldName = "region"
	Description       commons.FieldName = "description"
)

//Tags
const (
	Tags     commons.FieldName = "tags"
	TagKey   commons.FieldName = "tag_key"
	TagValue commons.FieldName = "tag_value"
)

const (
	ManagedServiceIdentities commons.FieldName = "managed_service_identities"
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
	//Type     	 		commons.FieldName = "type"
	StorageURL commons.FieldName = "storage_url"
)
