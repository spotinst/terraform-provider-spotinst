package elastigroup_azure_strategy

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "strategy_"
)

const (
	Strategy           commons.FieldName = "strategy"
	SpotPercentage     commons.FieldName = "spot_percentage"
	OnDemandCount      commons.FieldName = "od_count"
	DrainingTimeout    commons.FieldName = "draining_timeout"
	FallbackToOnDemand commons.FieldName = "fallback_to_on_demand"
	RevertToSpot       commons.FieldName = "revert__to_spot"
	PerformAt          commons.FieldName = "perform_at"
)
