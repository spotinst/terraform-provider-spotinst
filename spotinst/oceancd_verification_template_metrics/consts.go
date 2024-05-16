package oceancd_verification_template_metrics

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Metrics          commons.FieldName = "metrics"
	BaseLine         commons.FieldName = "baseline"
	MaxRange         commons.FieldName = "max_range"
	MinRange         commons.FieldName = "min_range"
	BaseLineProvider commons.FieldName = "baseline_provider"
	Threshold        commons.FieldName = "threshold"
)

const (
	Datadog      commons.FieldName = "datadog"
	Duration     commons.FieldName = "duration"
	DatadogQuery commons.FieldName = "datadog_query"

	NewRelic      commons.FieldName = "new_relic"
	Profile       commons.FieldName = "profile"
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
	MetricsName           commons.FieldName = "metrics_name"
	SuccessCondition      commons.FieldName = "success_condition"
)

const (
	Provider   commons.FieldName = "template_provider"
	Job        commons.FieldName = "job"
	Jenkins    commons.FieldName = "jenkins"
	CloudWatch commons.FieldName = "cloud_watch"
)

const (
	CloudWatchDuration commons.FieldName = "cloud_watch_duration"
	Expression         commons.FieldName = "expression"
	MetricDataQueries  commons.FieldName = "metric_data_queries"
	ID                 commons.FieldName = "id"
	Label              commons.FieldName = "label"
	MetricStat         commons.FieldName = "metric_stat"
	Period             commons.FieldName = "period"
	ReturnData         commons.FieldName = "return_data"
	Metric             commons.FieldName = "metric"
	Stat               commons.FieldName = "stat"
	Unit               commons.FieldName = "unit"
	MetricPeriod       commons.FieldName = "metric_period"
	Dimensions         commons.FieldName = "dimensions"
	MetricName         commons.FieldName = "metric_name"
	Namespace          commons.FieldName = "namespace"
	DimensionName      commons.FieldName = "dimension_name"
	DimensionValue     commons.FieldName = "dimension_value"
)

const (
	JenkinsInterval   commons.FieldName = "jenkins_interval"
	JenkinsParameters commons.FieldName = "jenkins_parameters"
	ParameterKey      commons.FieldName = "parameter_key"
	ParameterValue    commons.FieldName = "parameter_value"
	PipelineName      commons.FieldName = "pipeline_name"
	Timeout           commons.FieldName = "timeout"
	TlsVerification   commons.FieldName = "tls_verification"
)

const (
	Web            commons.FieldName = "web"
	Body           commons.FieldName = "body"
	WebHeader      commons.FieldName = "web_header"
	Insecure       commons.FieldName = "insecure"
	JsonPath       commons.FieldName = "json_path"
	Method         commons.FieldName = "method"
	TimeoutSeconds commons.FieldName = "timeout_seconds"
	Url            commons.FieldName = "url"
	WebHeaderKey   commons.FieldName = "web_header_key"
	WebHeaderValue commons.FieldName = "web_header_value"
)

const (
	Spec          commons.FieldName = "spec"
	BackoffLimit  commons.FieldName = "backoff_limit"
	JobTemplate   commons.FieldName = "job_template"
	TemplateSpec  commons.FieldName = "template_spec"
	Containers    commons.FieldName = "containers"
	Command       commons.FieldName = "command"
	Image         commons.FieldName = "image"
	ContinerName  commons.FieldName = "container_name"
	RestartPolicy commons.FieldName = "restart_policy"
)
