package multai_target_set

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	BalancerID   commons.FieldName = "balancer_id"
	DeploymentID commons.FieldName = "deployment_id"
	Name         commons.FieldName = "name"
	Protocol     commons.FieldName = "protocol"
	Port         commons.FieldName = "port"
	Weight       commons.FieldName = "weight"
	HealthCheck  commons.FieldName = "health_check"
	Tags         commons.FieldName = "tags"

	// HealthCheck fields
	Path               commons.FieldName = "path"
	Interval           commons.FieldName = "interval"
	Timeout            commons.FieldName = "timeout"
	HealthyThreshold   commons.FieldName = "healthy_threshold"
	UnhealthyThreshold commons.FieldName = "unhealthy_threshold"

	// tag fields
	TagKey   commons.FieldName = "key"
	TagValue commons.FieldName = "value"
)
