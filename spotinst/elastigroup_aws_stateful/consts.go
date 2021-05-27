package elastigroup_aws_stateful

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "stateful_"
)

const (
	PersistRootDevice   commons.FieldName = "persist_root_device"
	PersistBlockDevices commons.FieldName = "persist_block_devices"
	PersistPrivateIp    commons.FieldName = "persist_private_ip"
	BlockDevicesMode    commons.FieldName = "block_devices_mode"
	PrivateIps          commons.FieldName = "private_ips"

	// - DEALLOCATION -------------------------
	StatefulDeallocation          commons.FieldName = "stateful_deallocation"
	ShouldDeleteImages            commons.FieldName = "should_delete_images"
	ShouldDeleteNetworkInterfaces commons.FieldName = "should_delete_network_interfaces"
	ShouldDeleteVolumes           commons.FieldName = "should_delete_volumes"
	ShouldDeleteSnapshots         commons.FieldName = "should_delete_snapshots"
	// ----------------------------------------

	// - Stateful Instance Actions -------------------------
	StatefulInstancesActions         commons.FieldName = "stateful_instances_actions"
	StatefulInstancePauseAction      commons.FieldName = "stateful_instance_pause_action"
	StatefulInstanceResumeAction     commons.FieldName = "stateful_instance_resume_action"
	StatefulInstanceRecycleAction    commons.FieldName = "stateful_instance_recycle_action"
	StatefulInstanceDeAllocateAction commons.FieldName = "stateful_instance_de_allocate_action"
	StatefulInstanceID               commons.FieldName = "stateful_instance_id"
	PauseStatefulInstance            commons.FieldName = "pause_stateful_instance"
	ResumeStatefulInstance           commons.FieldName = "resume_stateful_instance"
	RecycleStatefulInstance          commons.FieldName = "recycle_stateful_instance"
	DeAllocateStatefulInstance       commons.FieldName = "de_allocate_stateful_instance"
	// ----------------------------------------
)
