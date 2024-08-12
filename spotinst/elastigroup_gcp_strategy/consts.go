package elastigroup_gcp_strategy

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	PreemptiblePercentage commons.FieldName = "preemptible_percentage"
	OnDemandCount         commons.FieldName = "ondemand_count"
	DrainingTimeout       commons.FieldName = "draining_timeout"
	FallbackToOnDemand    commons.FieldName = "fallback_to_ondemand"
	ProvisioningModel     commons.FieldName = "provisioning_model"
	OptimizationWindow    commons.FieldName = "optimization_window"
	RevertToPreemptible   commons.FieldName = "revert_to_preemptible"
	PerformAt             commons.FieldName = "perform_at"
)
