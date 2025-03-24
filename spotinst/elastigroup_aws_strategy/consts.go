package elastigroup_aws_strategy

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "strategy_"
)

const (
	SpotPercentage              commons.FieldName = "spot_percentage"
	OnDemandCount               commons.FieldName = "ondemand_count"
	Orientation                 commons.FieldName = "orientation"
	LifetimePeriod              commons.FieldName = "lifetime_period"
	DrainingTimeout             commons.FieldName = "draining_timeout"
	UtilizeReservedInstances    commons.FieldName = "utilize_reserved_instances"
	FallbackToOnDemand          commons.FieldName = "fallback_to_ondemand"
	ScalingStrategy             commons.FieldName = "scaling_strategy"
	TerminateAtEndOfBillingHour commons.FieldName = "terminate_at_end_of_billing_hour"
	TerminationPolicy           commons.FieldName = "termination_policy"
	UtilizeCommitments          commons.FieldName = "utilize_commitments"
	MinimumInstanceLifetime     commons.FieldName = "minimum_instance_lifetime"
	ConsiderODPricing           commons.FieldName = "consider_od_pricing"
	ImmediateODRecoverThreshold commons.FieldName = "immediate_od_recover_threshold"
	RestrictSingleAz            commons.FieldName = "restrict_single_az"
	MaxReplacementsPercentage   commons.FieldName = "max_replacements_percentage"
)
