package elastigroup_azure_scaling_policies

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

type DimensionField string

const (
	ScalingUpPolicy   commons.FieldName = "scaling_up_policy"
	ScalingDownPolicy commons.FieldName = "scaling_down_policy"

	PolicyName commons.FieldName = "policy_name"
	MetricName commons.FieldName = "metric_name"
	Statistic  commons.FieldName = "statistic"
	Unit       commons.FieldName = "unit"
	Threshold  commons.FieldName = "threshold"
	Adjustment commons.FieldName = "adjustment"
	Namespace  commons.FieldName = "namespace"
	Period     commons.FieldName = "period"
	Cooldown   commons.FieldName = "cooldown"
	Operator   commons.FieldName = "operator"
	Dimensions commons.FieldName = "dimensions"

	EvaluationPeriods commons.FieldName = "evaluation_periods"
	MinTargetCapacity commons.FieldName = "min_target_capacity"
	MaxTargetCapacity commons.FieldName = "max_target_capacity"
	Minimum           commons.FieldName = "minimum"
	Maximum           commons.FieldName = "maximum"
	Target            commons.FieldName = "target"
	ActionType        commons.FieldName = "action_type"

	DimensionName  DimensionField = "name"
	DimensionValue DimensionField = "value"
)
