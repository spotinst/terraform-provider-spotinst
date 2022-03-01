package ocean_aws_auto_scaling

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Autoscaler                       commons.FieldName = "autoscaler"
	AutoscaleCooldown                commons.FieldName = "autoscale_cooldown"
	AutoHeadroomPercentage           commons.FieldName = "auto_headroom_percentage"
	AutoscaleDown                    commons.FieldName = "autoscale_down"
	AutoscaleHeadroom                commons.FieldName = "autoscale_headroom"
	AutoscaleIsAutoConfig            commons.FieldName = "autoscale_is_auto_config"
	AutoscaleIsEnabled               commons.FieldName = "autoscale_is_enabled"
	EvaluationPeriods                commons.FieldName = "evaluation_periods"
	MaxScaleDownPercentage           commons.FieldName = "max_scale_down_percentage"
	CPUPerUnit                       commons.FieldName = "cpu_per_unit"
	GPUPerUnit                       commons.FieldName = "gpu_per_unit"
	MaxVCPU                          commons.FieldName = "max_vcpu"
	MaxMemoryGIB                     commons.FieldName = "max_memory_gib"
	MemoryPerUnit                    commons.FieldName = "memory_per_unit"
	NumOfUnits                       commons.FieldName = "num_of_units"
	ResourceLimits                   commons.FieldName = "resource_limits"
	EnableAutomaticAndManualHeadroom commons.FieldName = "enable_automatic_and_manual_headroom"
	ExtendedResourceDefinitions      commons.FieldName = "extended_resource_definitions"
)
