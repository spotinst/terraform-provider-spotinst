package oceancd_verification_template_args

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	BaseLine         commons.FieldName = "baseline"
	MaxRange         commons.FieldName = "max_range"
	MinRange         commons.FieldName = "min_range"
	BaseLineProvider commons.FieldName = "provider"
	Threshold        commons.FieldName = "threshold"
)

const (
	Datadog      commons.FieldName = "datadog"
	Duration     commons.FieldName = "duration"
	DatadogQuery commons.FieldName = "datadog_query"

	NewRelic      commons.FieldName = "new_relic"
	Profile       commons.FieldName = "Profile"
	NewRelicQuery commons.FieldName = "new_relic_query"

	Prometheus      commons.FieldName = "prometheus"
	PrometheusQuery commons.FieldName = "prometheus_query"
)

const (
	ConsecutiveErrorLimit commons.FieldName = "consecutive_error_limit"
	Count                 commons.FieldName = "count"
	DryRun                commons.FieldName = "dry_run"
	FailureCondition      commons.FieldName = "failure_condition"
	FailureLimit          commons.FieldName = "failure_limit"
	InitialDelay          commons.FieldName = "initial_delay"
	Interval              commons.FieldName = "interval"
	Name                  commons.FieldName = "name"
	SuccessCondition      commons.FieldName = "success_condition"
)
