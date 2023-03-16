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
)
