package elastigroup_aws_instance_types

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "instance_types_"
)

const (
	OnDemand      commons.FieldName = Prefix + "ondemand"
	Spot          commons.FieldName = Prefix + "spot"
	PreferredSpot commons.FieldName = Prefix + "preferred_spot"

	InstanceTypeWeights commons.FieldName = Prefix + "weights"
	InstanceType        commons.FieldName = "instance_type"
	Weight              commons.FieldName = "weight"
)
