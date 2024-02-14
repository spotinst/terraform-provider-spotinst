package ocean_aws_scheduling

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	ScheduledTask             commons.FieldName = "scheduled_task"
	ShutdownHours             commons.FieldName = "shutdown_hours"
	TimeWindows               commons.FieldName = "time_windows"
	ShutdownHoursIsEnabled    commons.FieldName = "is_enabled"
	Tasks                     commons.FieldName = "tasks"
	TasksIsEnabled            commons.FieldName = "is_enabled"
	CronExpression            commons.FieldName = "cron_expression"
	TaskType                  commons.FieldName = "task_type"
	Parameters                commons.FieldName = "parameters"
	ParametersClusterRoll     commons.FieldName = "parameters_cluster_roll"
	ApplyRoll                 commons.FieldName = "apply_roll"
	AmiAutoUpdate             commons.FieldName = "ami_auto_update"
	AmiAutoUpdateClusterRoll  commons.FieldName = "ami_auto_update_cluster_roll"
	MinorVersion              commons.FieldName = "minor_version"
	Patch                     commons.FieldName = "patch"
	BatchMinHealthyPercentage commons.FieldName = "batch_min_healthy_percentage"
	BatchSizePercentage       commons.FieldName = "batch_size_percentage"
	Comment                   commons.FieldName = "comment"
	RespectPdb                commons.FieldName = "respect_pdb"
)
