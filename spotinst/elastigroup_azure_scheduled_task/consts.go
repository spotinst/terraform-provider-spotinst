package elastigroup_azure_scheduled_task

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	ScheduledTask        commons.FieldName = "scheduled_task"
	IsEnabled            commons.FieldName = "is_enabled"
	CronExpression       commons.FieldName = "cron_expression"
	TaskType             commons.FieldName = "task_type"
	ScaleTargetCapacity  commons.FieldName = "scale_target_capacity"
	ScaleMinCapacity     commons.FieldName = "scale_min_capacity"
	ScaleMaxCapacity     commons.FieldName = "scale_max_capacity"
	BatchSizePercentage  commons.FieldName = "batch_size_percentage"
	GracePeriod          commons.FieldName = "grace_period"
	Adjustment           commons.FieldName = "adjustment"
	AdjustmentPercentage commons.FieldName = "adjustment_percentage"
)
