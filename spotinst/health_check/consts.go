package health_check

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Name       commons.FieldName = "name"
	ResourceId commons.FieldName = "resource_id"
	ProxyAddr  commons.FieldName = "proxy_address"
	ProxyPort  commons.FieldName = "proxy_port"
	Check      commons.FieldName = "check"
	Protocol   commons.FieldName = "protocol"
	Port       commons.FieldName = "port"
	Endpoint   commons.FieldName = "endpoint"
	Interval   commons.FieldName = "interval"
	Timeout    commons.FieldName = "timeout"
	Unhealthy  commons.FieldName = "unhealthy"
	Healthy    commons.FieldName = "healthy"
)
