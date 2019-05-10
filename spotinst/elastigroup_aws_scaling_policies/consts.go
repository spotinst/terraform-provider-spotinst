package elastigroup_aws_scaling_policies

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

type DimensionField string

const (
	ScalingUpPolicy     commons.FieldName = "scaling_up_policy"
	ScalingDownPolicy   commons.FieldName = "scaling_down_policy"
	ScalingTargetPolicy commons.FieldName = "scaling_target_policy"

	PolicyName commons.FieldName = "policy_name"
	MetricName commons.FieldName = "metric_name"
	Namespace  commons.FieldName = "namespace"
	Source     commons.FieldName = "source"
	Statistic  commons.FieldName = "statistic"
	Unit       commons.FieldName = "unit"
	Cooldown   commons.FieldName = "cooldown"
	Dimensions commons.FieldName = "dimensions"

	Threshold         commons.FieldName = "threshold"
	Adjustment        commons.FieldName = "adjustment"
	MinTargetCapacity commons.FieldName = "min_target_capacity"
	MaxTargetCapacity commons.FieldName = "max_target_capacity"
	Operator          commons.FieldName = "operator"
	EvaluationPeriods commons.FieldName = "evaluation_periods"
	Period            commons.FieldName = "period"
	Minimum           commons.FieldName = "minimum"
	Maximum           commons.FieldName = "maximum"
	Target            commons.FieldName = "target"
	ActionType        commons.FieldName = "action_type"
	IsEnabled         commons.FieldName = "is_enabled"
	PredictiveMode    commons.FieldName = "predictive_mode"

	DimensionName  DimensionField = "name"
	DimensionValue DimensionField = "value"
)
