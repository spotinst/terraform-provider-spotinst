package managed_instance_scheduling

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	ScheduledTask  commons.FieldName = "scheduled_task"
	Tasks          commons.FieldName = "tasks"
	IsEnabled      commons.FieldName = "is_enabled"
	Frequency      commons.FieldName = "frequency"
	StartTime      commons.FieldName = "start_time"
	CronExpression commons.FieldName = "cron_expression"
	TaskType       commons.FieldName = "task_type"
)
