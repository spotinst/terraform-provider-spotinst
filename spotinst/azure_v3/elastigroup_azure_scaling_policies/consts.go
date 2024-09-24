package elastigroup_azure_scaling_policies

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	ScalingUpPolicy   commons.FieldName = "scaling_up_policy"
	ScalingDownPolicy commons.FieldName = "scaling_down_policy"

	PolicyName        commons.FieldName = "policy_name"
	MetricName        commons.FieldName = "metric_name"
	Namespace         commons.FieldName = "namespace"
	Statistic         commons.FieldName = "statistic"
	Unit              commons.FieldName = "unit"
	Cooldown          commons.FieldName = "cooldown"
	Dimensions        commons.FieldName = "dimensions"
	Action            commons.FieldName = "action"
	Threshold         commons.FieldName = "threshold"
	Operator          commons.FieldName = "operator"
	EvaluationPeriods commons.FieldName = "evaluation_periods"
	Period            commons.FieldName = "period"
	Source            commons.FieldName = "source"
	IsEnabled         commons.FieldName = "is_enabled"

	Adjustment commons.FieldName = "adjustment"
	Minimum    commons.FieldName = "minimum"
	Maximum    commons.FieldName = "maximum"
	Target     commons.FieldName = "target"
	Type       commons.FieldName = "type"

	DimensionName  commons.FieldName = "name"
	DimensionValue commons.FieldName = "value"
)
