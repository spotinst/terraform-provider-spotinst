package elastigroup_instance_types

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	OnDemand commons.FieldName = "ondemand"
	Spot     commons.FieldName = "spot"

	InstanceTypeWeights commons.FieldName = "instance_type_weights"
	InstanceType        commons.FieldName = "instance_type"
	Weight              commons.FieldName = "weight"
)
