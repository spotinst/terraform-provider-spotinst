package elastigroup_azure_health

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Health            commons.FieldName = "health"
	HealthCheckTypes  commons.FieldName = "health_check_types"
	GracePeriod       commons.FieldName = "grace_period"
	UnhealthyDuration commons.FieldName = "unhealthy_duration"
	AutoHealing       commons.FieldName = "auto_healing"
)
