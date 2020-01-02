package managed_instance_persistence

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	PersistBlockDevices commons.FieldName = "persist_block_devices"
	PersistRootDevice   commons.FieldName = "persist_root_device"
	PersistPrivateIp    commons.FieldName = "persist_private_ip"
	BlockDevicesMode    commons.FieldName = "block_devices_mode"
)
