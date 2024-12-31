package ocean_ecs_strategy

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	DrainingTimeout          commons.FieldName = "draining_timeout"
	UtilizeReservedInstances commons.FieldName = "utilize_reserved_instances"
	UtilizeCommitments       commons.FieldName = "utilize_commitments"
	SpotPercentage           commons.FieldName = "spot_percentage"
	ClusterOrientation       commons.FieldName = "cluster_orientation"
	AvailabilityVsCost       commons.FieldName = "availability_vs_cost"
	FallbackToOnDemand       commons.FieldName = "fallback_to_ondemand"
)
