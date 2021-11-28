package ocean_gke_launch_spec_scheduling

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	SchedulingTask commons.FieldName = "scheduling_task"
)

const (
	IsEnabled      commons.FieldName = "is_enabled"
	CronExpression commons.FieldName = "cron_expression"
	TaskType       commons.FieldName = "task_type"
	TaskHeadroom   commons.FieldName = "task_headroom"
)

const (
	CPUPerUnit    commons.FieldName = "cpu_per_unit"
	GPUPerUnit    commons.FieldName = "gpu_per_unit"
	MemoryPerUnit commons.FieldName = "memory_per_unit"
	NumOfUnits    commons.FieldName = "num_of_units"
)
