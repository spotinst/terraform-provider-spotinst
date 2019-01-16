package elastigroup_azure_strategy

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "strategy_"
)

const (
	Strategy              commons.FieldName = "strategy"
	LowPriorityPercentage commons.FieldName = "low_priority_percentage"
	OnDemandCount         commons.FieldName = "od_count"
	DrainingTimeout       commons.FieldName = "draining_timeout"
)
