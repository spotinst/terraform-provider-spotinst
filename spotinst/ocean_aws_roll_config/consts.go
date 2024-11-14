package ocean_aws_roll_config

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Roll                         commons.FieldName = "roll"
	BatchSizePercentage          commons.FieldName = "batch_size_percentage"
	LaunchSpecIDs                commons.FieldName = "launch_spec_ids"
	BatchMinHealthyPercentage    commons.FieldName = "batch_min_healthy_percentage"
	RespectPDB                   commons.FieldName = "respect_pdb"
	Comment                      commons.FieldName = "comment"
	InstanceIds                  commons.FieldName = "instance_ids"
	DisableLaunchSpecAutoScaling commons.FieldName = "disable_launch_spec_auto_scaling"
	OceanClusterId               commons.FieldName = "ocean_cluster_id"
)
