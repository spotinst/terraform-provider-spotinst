package managed_instance_strategy

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	LifeCycle                commons.FieldName = "life_cycle"
	Orientation              commons.FieldName = "orientation"
	DrainingTimeout          commons.FieldName = "draining_timeout"
	FallbackToOd             commons.FieldName = "fall_back_to_od"
	UtilizeReservedInstances commons.FieldName = "utilize_reserved_instances"
	OptimizationWindows      commons.FieldName = "optimization_windows"
	RevertToSpot             commons.FieldName = "revert_to_spot"
	PerformAt                commons.FieldName = "perform_at"
)
