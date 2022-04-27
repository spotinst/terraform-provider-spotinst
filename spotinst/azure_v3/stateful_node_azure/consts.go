package stateful_node_azure

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Name              commons.FieldName = "name"
	ResourceGroupName commons.FieldName = "resource_group_name"
	Region            commons.FieldName = "region"
	Description       commons.FieldName = "description"
)

const (
	OS            commons.FieldName = "os"
	Zones         commons.FieldName = "zones"
	PreferredZone commons.FieldName = "preferred_zone"
)

const (
	UpdateState       commons.FieldName = "update_state"
	UpdateStateConfig commons.FieldName = "update_state_config"
	ShouldUpdateState commons.FieldName = "should_update_state"
	State             commons.FieldName = "state"
)

const (
	AttachPrefix   = "attach_"
	DetachPrefix   = "detach_"
	ImportVMPrefix = "import_vm_"
)

const (
	AttachDataDisk                  commons.FieldName = AttachPrefix + "data_disk"
	AttachDataDiskConfig            commons.FieldName = AttachPrefix + "data_disk_config"
	AttachDataDiskName              commons.FieldName = AttachPrefix + "data_disk_name"
	AttachDataDiskResourceGroupName commons.FieldName = AttachPrefix + "data_disk_resource_group_name"
	AttachStorageAccountType        commons.FieldName = AttachPrefix + "storage_account_type"
	AttachSizeGB                    commons.FieldName = AttachPrefix + "size_gb"
	AttachLUN                       commons.FieldName = AttachPrefix + "lun"
	AttachZone                      commons.FieldName = AttachPrefix + "zone"

	ShouldAttachDataDisk commons.FieldName = "should_attach_data_disk"
)

const (
	DetachDataDisk                  commons.FieldName = DetachPrefix + "data_disk"
	DetachDataDiskConfig            commons.FieldName = DetachPrefix + "data_disk_config"
	DetachDataDiskName              commons.FieldName = DetachPrefix + "data_disk_name"
	DetachDataDiskResourceGroupName commons.FieldName = DetachPrefix + "data_disk_resource_group_name"
	DetachShouldDeallocate          commons.FieldName = DetachPrefix + "should_deallocate"

	ShouldDetachDataDisk commons.FieldName = "should_detach_data_disk"
)

const (
	ImportVMResourceGroupName     commons.FieldName = ImportVMPrefix + "resource_group_name"
	ImportVMOriginalVMName        commons.FieldName = ImportVMPrefix + "original_vm_name"
	ImportVMDrainingTimeout       commons.FieldName = ImportVMPrefix + "draining_timeout"
	ImportVMResourceRetentionTime commons.FieldName = ImportVMPrefix + "resource_retention_time"

	ShouldImportVM commons.FieldName = "should_import_vm"
)
