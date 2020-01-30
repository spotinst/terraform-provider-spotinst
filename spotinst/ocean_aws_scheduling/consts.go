package ocean_aws_scheduling

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	ScheduledTask  commons.FieldName = "scheduled_task"
	ShutdownHours  commons.FieldName = "shutdown_hours"
	TimeWindows    commons.FieldName = "time_windows"
	IsEnabled      commons.FieldName = "is_enabled"
	tasks          commons.FieldName = "tasks"
	tasksIsEnabled commons.FieldName = "tasks_is_enabled"
	CronExpression commons.FieldName = "cron_expression"
	TaskType       commons.FieldName = "task_type"
)
