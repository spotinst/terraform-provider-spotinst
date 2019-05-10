package elastigroup_gcp_scheduled_task

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	ScheduledTask  commons.FieldName = "scheduled_task"
	IsEnabled      commons.FieldName = "is_enabled"
	TaskType       commons.FieldName = "task_type"
	CronExpression commons.FieldName = "cron_expression"
	TargetCapacity commons.FieldName = "target_capacity"
	MinCapacity    commons.FieldName = "min_capacity"
	MaxCapacity    commons.FieldName = "max_capacity"
)
