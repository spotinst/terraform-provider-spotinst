package ocean_aks_np_virtual_node_group

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	OceanID           commons.FieldName = "ocean_id"
	Name              commons.FieldName = "name"
	AvailabilityZones commons.FieldName = "availability_zones"
	Tags              commons.FieldName = "tags"
	Labels            commons.FieldName = "labels"
)

const (
	Taints      commons.FieldName = "taints"
	TaintKey    commons.FieldName = "key"
	TaintValue  commons.FieldName = "value"
	TaintEffect commons.FieldName = "effect"

	UpdatePolicy          commons.FieldName = "update_policy"
	ShouldRoll            commons.FieldName = "should_roll"
	ConditionedRoll       commons.FieldName = "conditioned_roll"
	ConditionedRollParams commons.FieldName = "conditioned_roll_params"

	RollConfig                commons.FieldName = "roll_config"
	BatchSizePercentage       commons.FieldName = "batch_size_percentage"
	VngIDs                    commons.FieldName = "vng_ids"
	BatchMinHealthyPercentage commons.FieldName = "batch_min_healthy_percentage"
	RespectPDB                commons.FieldName = "respect_pdb"
	Comment                   commons.FieldName = "comment"
	NodePoolNames             commons.FieldName = "node_pool_names"
	RespectRestrictScaleDown  commons.FieldName = "respect_restrict_scale_down"
	NodeNames                 commons.FieldName = "node_names"
)
