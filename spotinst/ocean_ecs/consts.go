package ocean_ecs

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

type TagField string

const (
	Region              commons.FieldName = "region"
	Name                commons.FieldName = "name"
	ClusterName         commons.FieldName = "cluster_name"
	MaxSize             commons.FieldName = "max_size"
	MinSize             commons.FieldName = "min_size"
	DesiredCapacity     commons.FieldName = "desired_capacity"
	SubnetIDs           commons.FieldName = "subnet_ids"
	UpdatePolicy        commons.FieldName = "update_policy"
	ShouldRoll          commons.FieldName = "should_roll"
	RollConfig          commons.FieldName = "roll_config"
	BatchSizePercentage commons.FieldName = "batch_size_percentage"
	Tags                commons.FieldName = "tags"
	TagKey              TagField          = "key"
	TagValue            TagField          = "value"
)
