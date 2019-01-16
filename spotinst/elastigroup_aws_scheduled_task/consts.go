package elastigroup_aws_scheduled_task

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "scheduled_task"
)

const TaskTypeStatefulUpdateCapacity = "statefulUpdateCapacity"

const (
	ScheduledTask        commons.FieldName = "scheduled_task"
	IsEnabled            commons.FieldName = "is_enabled"
	TaskType             commons.FieldName = "task_type"
	Frequency            commons.FieldName = "frequency"
	CronExpression       commons.FieldName = "cron_expression"
	StartTime            commons.FieldName = "start_time"
	ScaleTargetCapacity  commons.FieldName = "scale_target_capacity"
	ScaleMinCapacity     commons.FieldName = "scale_min_capacity"
	ScaleMaxCapacity     commons.FieldName = "scale_max_capacity"
	BatchSizePercentage  commons.FieldName = "batch_size_percentage"
	GracePeriod          commons.FieldName = "grace_period"
	TargetCapacity       commons.FieldName = "target_capacity"
	MinCapacity          commons.FieldName = "min_capacity"
	MaxCapacity          commons.FieldName = "max_capacity"
	Adjustment           commons.FieldName = "adjustment"
	AdjustmentPercentage commons.FieldName = "adjustment_percentage"
)
