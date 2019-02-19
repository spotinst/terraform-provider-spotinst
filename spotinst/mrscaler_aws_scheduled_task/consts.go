package mrscaler_aws_scheduled_task

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	ScheduledTask     commons.FieldName = "scheduled_task"
	IsEnabled         commons.FieldName = "is_enabled"
	TaskType          commons.FieldName = "task_type"
	InstanceGroupType commons.FieldName = "instance_group_type"
	CronExpression    commons.FieldName = "cron"
	TargetCapacity    commons.FieldName = "desired_capacity"
	MinCapacity       commons.FieldName = "min_capacity"
	MaxCapacity       commons.FieldName = "max_capacity"
)
