package ocean_gke_import_scheduling

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
	BatchSizePercentage       commons.FieldName = "batch_size_percentage"
	TaskParameters            commons.FieldName = "task_parameters"
	BatchMinHealthyPercentage commons.FieldName = "batch_min_healthy_percentage"
	Comment                   commons.FieldName = "comment"
	RespectPdb                commons.FieldName = "respect_pdb"
	ClusterRoll               commons.FieldName = "cluster_roll"
)
