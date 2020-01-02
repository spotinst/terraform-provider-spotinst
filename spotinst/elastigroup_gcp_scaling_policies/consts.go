package elastigroup_gcp_scaling_policies

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	ScalingDownPolicy commons.FieldName = "scaling_down_policy"
	ScalingUpPolicy   commons.FieldName = "scaling_up_policy"

	Cooldown   commons.FieldName = "cooldown"
	Dimensions commons.FieldName = "dimensions"
	MetricName commons.FieldName = "metric_name"
	Namespace  commons.FieldName = "namespace"
	PolicyName commons.FieldName = "policy_name"
	Source     commons.FieldName = "source"
	Statistic  commons.FieldName = "statistic"
	Unit       commons.FieldName = "unit"

	ActionType        commons.FieldName = "action_type"
	Adjustment        commons.FieldName = "adjustment"
	EvaluationPeriods commons.FieldName = "evaluation_periods"
	Operator          commons.FieldName = "operator"
	Period            commons.FieldName = "period"
	Threshold         commons.FieldName = "threshold"

	DimensionName  commons.FieldName = "name"
	DimensionValue commons.FieldName = "value"
)
