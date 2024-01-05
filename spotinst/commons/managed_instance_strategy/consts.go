package managed_instance_strategy

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	LifeCycle                commons.FieldName = "life_cycle"
	Orientation              commons.FieldName = "orientation"
	DrainingTimeout          commons.FieldName = "draining_timeout"
	FallbackToOd             commons.FieldName = "fallback_to_ondemand"
	UtilizeReservedInstances commons.FieldName = "utilize_reserved_instances"
	OptimizationWindows      commons.FieldName = "optimization_windows"
	RevertToSpot             commons.FieldName = "revert_to_spot"
	PerformAt                commons.FieldName = "perform_at"
	MinimumInstanceLifetime  commons.FieldName = "minimum_instance_lifetime"
)
