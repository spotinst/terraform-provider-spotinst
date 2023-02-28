package ocean_aks_np_node_count_limits

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

type TaintField string

const (
	MinCount commons.FieldName = "min_count"
	MaxCount commons.FieldName = "max_count"
)

const (
	Tags commons.FieldName = "tags"
)

const (
	Label commons.FieldName = "labels"
)

const (
	Taint       commons.FieldName = "taints"
	TaintKey    TaintField        = "key"
	TaintValue  TaintField        = "value"
	TaintEffect TaintField        = "effect"
)
