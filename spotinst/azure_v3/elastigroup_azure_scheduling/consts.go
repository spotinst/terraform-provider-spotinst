package elastigroup_azure_scheduling

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	SchedulingTask       commons.FieldName = "scheduling_task"
	IsEnabled            commons.FieldName = "is_enabled"
	CronExpression       commons.FieldName = "cron_expression"
	Type                 commons.FieldName = "type"
	ScaleMaxCapacity     commons.FieldName = "scale_max_capacity"
	ScaleMinCapacity     commons.FieldName = "scale_min_capacity"
	ScaleTargetCapacity  commons.FieldName = "scale_target_capacity"
	Adjustment           commons.FieldName = "adjustment"
	AdjustmentPercentage commons.FieldName = "adjustment_percentage"
	BatchSizePercentage  commons.FieldName = "batch_size_percentage"
	GracePeriod          commons.FieldName = "grace_period"
)
