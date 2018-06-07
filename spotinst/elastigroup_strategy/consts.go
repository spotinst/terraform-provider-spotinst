package elastigroup_strategy

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "strategy_"
)

const (
	SpotPercentage           commons.FieldName = "spot_percentage"
	OnDemandCount            commons.FieldName = "ondemand_count"
	Orientation              commons.FieldName = "orientation"
	LifetimePeriod           commons.FieldName = "lifetime_period"
	DrainingTimeout          commons.FieldName = "draining_timeout"
	UtilizeReservedInstances commons.FieldName = "utilize_reserved_instances"
	FallbackToOnDemand       commons.FieldName = "fallback_to_ondemand"
)
