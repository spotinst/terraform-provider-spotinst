package ocean_aks_np_auto_scaling

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	AutoScaler         commons.FieldName = "autoscaler"
	AutoscaleIsEnabled commons.FieldName = "autoscale_is_enabled"
)

const (
	ResourceLimits commons.FieldName = "resource_limits"
	MaxVCPU        commons.FieldName = "max_vcpu"
	MaxMemoryGib   commons.FieldName = "max_memory_gib"
)

const (
	Down                   commons.FieldName = "autoscale_down"
	MaxScaleDownPercentage commons.FieldName = "max_scale_down_percentage"
)

const (
	Headroom   commons.FieldName = "autoscale_headroom"
	Automatic  commons.FieldName = "automatic"
	IsEnabled  commons.FieldName = "is_enabled"
	Percentage commons.FieldName = "percentage"
)
