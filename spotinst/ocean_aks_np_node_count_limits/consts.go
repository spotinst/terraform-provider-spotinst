package ocean_aks_np_node_count_limits

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	MinCount commons.FieldName = "min_count"
	MaxCount commons.FieldName = "max_count"
)

type TagField string

const (
	Tags     commons.FieldName = "tags"
	TagKey   TagField          = "key"
	TagValue TagField          = "value"
)

const (
	Label      commons.FieldName = "labels"
	LabelKey   commons.FieldName = "key"
	LabelValue commons.FieldName = "value"
)

const (
	Taint       commons.FieldName = "taints"
	TaintKey    commons.FieldName = "key"
	TaintValue  commons.FieldName = "value"
	TaintEffect commons.FieldName = "effect"
)
