package managed_instance_aws

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Name        commons.FieldName = "name"
	Description commons.FieldName = "description"
	Region      commons.FieldName = "region"

	// - Instance Action ----------------------
	ManagedInstanceAction commons.FieldName = "managed_instance_action"
	ActionType            commons.FieldName = "type"
	// ----------------------------------------
)

const (
	Delete                               commons.FieldName = "delete"
	AmiBackupShouldDeleteImages          commons.FieldName = "ami_backup_should_delete_images"
	DeallocationConfigShouldDeleteImages commons.FieldName = "deallocation_config_should_delete_images"
	ShouldDeleteNetworkInterfaces        commons.FieldName = "should_delete_network_interfaces"
	ShouldDeleteSnapshots                commons.FieldName = "should_delete_snapshots"
	ShouldDeleteVolumes                  commons.FieldName = "should_delete_volumes"
	ShouldTerminateInstance              commons.FieldName = "should_terminate_instance"
)
