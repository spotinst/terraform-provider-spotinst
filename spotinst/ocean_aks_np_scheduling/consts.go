package ocean_aks_np_scheduling

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Scheduling          commons.FieldName = "scheduling"
	ShutdownHours       commons.FieldName = "shutdown_hours"
	TimeWindows         commons.FieldName = "time_windows"
	SchedulingIsEnabled commons.FieldName = "is_enabled"
	SuspensionHours     commons.FieldName = "suspension_hours"
)
const (
	Tasks                     commons.FieldName = "tasks"
	TasksIsEnabled            commons.FieldName = "is_enabled"
	Parameters                commons.FieldName = "parameters"
	ParametersClusterRoll     commons.FieldName = "parameters_cluster_roll"
	TaskType                  commons.FieldName = "task_type"
	CronExpression            commons.FieldName = "cron_expression"
	BatchMinHealthyPercentage commons.FieldName = "batch_min_healthy_percentage"
	BatchSizePercentage       commons.FieldName = "batch_size_percentage"
	Comment                   commons.FieldName = "comment"
	RespectPdb                commons.FieldName = "respect_pdb"
	RespectRestrictScaleDown  commons.FieldName = "respect_restrict_scale_down"
	VngIDs                    commons.FieldName = "vng_ids"
)
const (
	ParametersUpgradeConfig commons.FieldName = "parameters_upgrade_config"
	ApplyRoll               commons.FieldName = "apply_roll"
	ScopeVersion            commons.FieldName = "scope_version"
	RollParameters          commons.FieldName = "roll_parameters"
)
