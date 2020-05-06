package mrscaler_aws_terminationPolicies

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	TerminationPolicies commons.FieldName = "termination_policies"
	Statements          commons.FieldName = "statements"
	Namespace           commons.FieldName = "namespace"
	MetricName          commons.FieldName = "metric_name"
	Statistic           commons.FieldName = "statistic"
	Unit                commons.FieldName = "unit"
	Threshold           commons.FieldName = "threshold"
	Period              commons.FieldName = "period"
	EvaluationPeriods   commons.FieldName = "evaluation_periods"
	Operator            commons.FieldName = "operator"
)
