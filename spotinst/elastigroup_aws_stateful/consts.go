package elastigroup_aws_stateful

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	PersistRootDevice   commons.FieldName = "persist_root_device"
	PersistBlockDevices commons.FieldName = "persist_block_devices"
	PersistPrivateIp    commons.FieldName = "persist_private_ip"
	BlockDevicesMode    commons.FieldName = "block_devices_mode"
	PrivateIps          commons.FieldName = "private_ips"

	// - Deallocation -------------------------
	StatefulDeallocation          commons.FieldName = "stateful_deallocation"
	ShouldDeleteImages            commons.FieldName = "should_delete_images"
	ShouldDeleteNetworkInterfaces commons.FieldName = "should_delete_network_interfaces"
	ShouldDeleteVolumes           commons.FieldName = "should_delete_volumes"
	ShouldDeleteSnapshots         commons.FieldName = "should_delete_snapshots"
	// ----------------------------------------

	// - Instance Action ----------------------
	StatefulInstanceAction commons.FieldName = "stateful_instance_action"
	StatefulInstanceID     commons.FieldName = "stateful_instance_id"
	ActionType             commons.FieldName = "type"
	// ----------------------------------------
)
