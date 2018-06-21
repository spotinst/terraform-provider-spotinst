package subscription

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "subscription_"
)

const (
	ResourceId commons.FieldName = "resource_id"
	EventType  commons.FieldName = "event_type"
	Protocol   commons.FieldName = "protocol"
	Endpoint   commons.FieldName = "endpoint"
	Format     commons.FieldName = "format"
)
