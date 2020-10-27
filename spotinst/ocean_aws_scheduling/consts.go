package ocean_aws_scheduling

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	ScheduledTask          commons.FieldName = "scheduled_task"
	ShutdownHours          commons.FieldName = "shutdown_hours"
	TimeWindows            commons.FieldName = "time_windows"
	ShutdownHoursIsEnabled commons.FieldName = "is_enabled"
	Tasks                  commons.FieldName = "tasks"
	TasksIsEnabled         commons.FieldName = "is_enabled"
	CronExpression         commons.FieldName = "cron_expression"
	TaskType               commons.FieldName = "task_type"
)
