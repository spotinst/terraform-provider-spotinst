package stateful_node_azure_strategy

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Strategy           commons.FieldName = "strategy"
	PreferredLifecycle commons.FieldName = "preferred_life_cycle"
	FallbackToOnDemand commons.FieldName = "fallback_to_on_demand"
	DrainingTimeout    commons.FieldName = "draining_timeout"
)

const (
	OptimizationWindows commons.FieldName = "optimization_windows"
)

const (
	Signal  commons.FieldName = "signal"
	Type    commons.FieldName = "type"
	Timeout commons.FieldName = "timeout"
)

const (
	RevertToSpot commons.FieldName = "revert_to_spot"
	PerformAt    commons.FieldName = "perform_at"
)
