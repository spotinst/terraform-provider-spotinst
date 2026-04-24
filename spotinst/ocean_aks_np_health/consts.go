package ocean_aks_np_health

import (
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

const (
	Health                                        commons.FieldName = "health"
	GracePeriod                                   commons.FieldName = "grace_period"
	ShouldReplaceUnhealthyInstances               commons.FieldName = "should_replace_unhealthy_instances"
	HealthCheckUnhealthyDurationBeforeReplacement commons.FieldName = "health_check_unhealthy_duration_before_replacement"
)
