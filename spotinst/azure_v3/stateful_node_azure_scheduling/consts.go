package stateful_node_azure_scheduling

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	SchedulingTask commons.FieldName = "scheduling_task"
	IsEnabled      commons.FieldName = "is_enabled"
	CronExpression commons.FieldName = "cron_expression"
	Type           commons.FieldName = "type"
)
