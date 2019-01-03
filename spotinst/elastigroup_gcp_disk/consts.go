package elastigroup_gcp_disk

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Disk             commons.FieldName = "disk"
	AutoDelete       commons.FieldName = "auto_delete"
	Boot             commons.FieldName = "boot"
	DeviceName       commons.FieldName = "device_name"
	InitializeParams commons.FieldName = "initialize_params"
	Interface        commons.FieldName = "interface"
	Mode             commons.FieldName = "mode"
	Source           commons.FieldName = "source"
	Type             commons.FieldName = "type"

	DiskSizeGB  commons.FieldName = "disk_size_gb"
	DiskType    commons.FieldName = "disk_type"
	SourceImage commons.FieldName = "source_image"
)
