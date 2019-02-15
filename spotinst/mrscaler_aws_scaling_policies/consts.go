package mrscaler_aws_scaling_policies

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	TaskScalingUpPolicy   commons.FieldName = "task_scaling_up_policy"
	TaskScalingDownPolicy commons.FieldName = "task_scaling_down_policy"
	CoreScalingUpPolicy   commons.FieldName = "core_scaling_up_policy"
	CoreScalingDownPolicy commons.FieldName = "core_scaling_down_policy"

	//ScalingDownPolicy commons.FieldName = "scaling_down_policy"
	//ScalingUpPolicy   commons.FieldName = "scaling_up_policy"

	Cooldown   commons.FieldName = "cooldown"
	Dimensions commons.FieldName = "dimensions"
	MetricName commons.FieldName = "metric_name"
	Namespace  commons.FieldName = "namespace"
	PolicyName commons.FieldName = "policy_name"
	Statistic  commons.FieldName = "statistic"
	Unit       commons.FieldName = "unit"

	ActionType        commons.FieldName = "action_type"
	Adjustment        commons.FieldName = "adjustment"
	EvaluationPeriods commons.FieldName = "evaluation_periods"
	Operator          commons.FieldName = "operator"
	Period            commons.FieldName = "period"
	Threshold         commons.FieldName = "threshold"
	Minimum           commons.FieldName = "minimum"
	Maximum           commons.FieldName = "maximum"
	Target            commons.FieldName = "target"
	MinTargetCapacity commons.FieldName = "min_target_capacity"
	MaxTargetCapacity commons.FieldName = "max_target_capacity"
)
