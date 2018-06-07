package elastigroup_stateful

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "stateful_"
)

const (
	PersistRootDevice   commons.FieldName = "persist_root_device"
	PersistBlockDevices commons.FieldName = "persist_block_devices"
	PersistPrivateIp    commons.FieldName = "persist_private_ip"
	BlockDevicesMode    commons.FieldName = "block_devices_mode"
	PrivateIps          commons.FieldName = "private_ips"
)
