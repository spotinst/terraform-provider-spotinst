package ocean_gke_import_autoscaler

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Autoscaler                       commons.FieldName = "autoscaler"
	IsAutoConfig                     commons.FieldName = "is_auto_config"
	IsEnabled                        commons.FieldName = "is_enabled"
	Cooldown                         commons.FieldName = "cooldown"
	Headroom                         commons.FieldName = "headroom"
	CpuPerUnit                       commons.FieldName = "cpu_per_unit"
	MemoryPerUnit                    commons.FieldName = "memory_per_unit"
	NumOfUnits                       commons.FieldName = "num_of_units"
	ResourceLimits                   commons.FieldName = "resource_limits"
	MaxVCpu                          commons.FieldName = "max_vcpu"
	MaxMemoryGib                     commons.FieldName = "max_memory_gib"
	Down                             commons.FieldName = "down"
	MaxScaleDownPercentage           commons.FieldName = "max_scale_down_percentage"
	EvaluationPeriods                commons.FieldName = "evaluation_periods"
	GPUPerUnit                       commons.FieldName = "gpu_per_unit"
	AutoHeadroomPercentage           commons.FieldName = "auto_headroom_percentage"
	EnableAutomaticAndManualHeadroom commons.FieldName = "enable_automatic_and_manual_headroom"
	IsAggressiveScaleDownEnabled     commons.FieldName = "is_aggressive_scale_down_enabled"
)
