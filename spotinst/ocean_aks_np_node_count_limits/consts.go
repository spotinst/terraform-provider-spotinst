package ocean_aks_np_node_count_limits

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	MinCount commons.FieldName = "min_count"
	MaxCount commons.FieldName = "max_count"
)

const (
	Tag      commons.FieldName = "tag"
	TagKey   commons.FieldName = "key"
	TagValue commons.FieldName = "value"
)

const (
	Label      commons.FieldName = "label"
	LabelKey   commons.FieldName = "key"
	LabelValue commons.FieldName = "value"
)

const (
	Taint       commons.FieldName = "taint"
	TaintKey    commons.FieldName = "key"
	TaintValue  commons.FieldName = "value"
	TaintEffect commons.FieldName = "effect"
)
