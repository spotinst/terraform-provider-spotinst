package ocean_aks_virtual_node_group

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	OceanID commons.FieldName = "ocean_id"
	Name    commons.FieldName = "name"
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

const (
	ResourceLimits   commons.FieldName = "resource_limits"
	MaxInstanceCount commons.FieldName = "max_instance_count"
)
