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
	UpdateState commons.FieldName = "update_state"
	State       commons.FieldName = "state"
)

const (
	Delete                   commons.FieldName = "delete"
	ShouldTerminateVm        commons.FieldName = "should_terminate_vm"
	NetworkShouldDeallocate  commons.FieldName = NetworkPrefix + "should_deallocate"
	NetworkTTLInHours        commons.FieldName = NetworkPrefix + "ttl_in_hours"
	DiskShouldDeallocate     commons.FieldName = DiskPrefix + "should_deallocate"
	DiskTTLInHours           commons.FieldName = DiskPrefix + "ttl_in_hours"
	SnapshotShouldDeallocate commons.FieldName = SnapshotPrefix + "should_deallocate"
	SnapshotTTLInHours       commons.FieldName = SnapshotPrefix + "ttl_in_hours"
	PublicIPShouldDeallocate commons.FieldName = PublicIPDeallocation + "should_deallocate"
	PublicIPTTLInHours       commons.FieldName = PublicIPDeallocation + "ttl_in_hours"
)

const (
	AttachDataDisk                  commons.FieldName = "attach_data_disk"
	AttachDataDiskName              commons.FieldName = "data_disk_name"
	AttachDataDiskResourceGroupName commons.FieldName = "data_disk_resource_group_name"
	AttachStorageAccountType        commons.FieldName = "storage_account_type"
	AttachSizeGB                    commons.FieldName = "size_gb"
	AttachLUN                       commons.FieldName = "lun"
	AttachZone                      commons.FieldName = "zone"
)

const (
	DetachDataDisk                  commons.FieldName = "detach_data_disk"
	DetachDataDiskName              commons.FieldName = "data_disk_name"
	DetachDataDiskResourceGroupName commons.FieldName = "data_disk_resource_group_name"
	DetachShouldDeallocate          commons.FieldName = "should_deallocate"
	DetachTTLInHours                commons.FieldName = "ttl_in_hours"
)

const (
	ImportVM                       commons.FieldName = "import_vm"
	ImportVMResourceGroupName      commons.FieldName = "resource_group_name"
	ImportVMOriginalVMName         commons.FieldName = "original_vm_name"
	ImportVMDrainingTimeout        commons.FieldName = "draining_timeout"
	ImportVMResourcesRetentionTime commons.FieldName = "resources_retention_time"
)

const (
	NetworkPrefix        = "network_"
	DiskPrefix           = "disk_"
	SnapshotPrefix       = "snapshot_"
	PublicIPDeallocation = "public_ip_"
)
