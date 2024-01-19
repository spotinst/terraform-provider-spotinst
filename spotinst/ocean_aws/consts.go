package ocean_aws

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

type TagField string

const (
	TagKey   TagField = "key"
	TagValue TagField = "value"
)

const (
	Name                commons.FieldName = "name"
	ControllerClusterID commons.FieldName = "controller_id"

	MaxSize         commons.FieldName = "max_size"
	MinSize         commons.FieldName = "min_size"
	DesiredCapacity commons.FieldName = "desired_capacity"

	Region    commons.FieldName = "region"
	SubnetIDs commons.FieldName = "subnet_ids"

	Tags commons.FieldName = "tags"

	UpdatePolicy          commons.FieldName = "update_policy"
	ShouldRoll            commons.FieldName = "should_roll"
	ConditionedRoll       commons.FieldName = "conditioned_roll"
	AutoApplyTags         commons.FieldName = "auto_apply_tags"
	ConditionedRollParams commons.FieldName = "conditioned_roll_params"

	RollConfig                commons.FieldName = "roll_config"
	BatchSizePercentage       commons.FieldName = "batch_size_percentage"
	LaunchSpecIDs             commons.FieldName = "launch_spec_ids"
	BatchMinHealthyPercentage commons.FieldName = "batch_min_healthy_percentage"
	RespectPDB                commons.FieldName = "respect_pdb"
)
