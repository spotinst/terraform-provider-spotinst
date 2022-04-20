package stateful_node_azure

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Name              commons.FieldName = "name"
	ResourceGroupName commons.FieldName = "resource_group_name"
	Region            commons.FieldName = "region"
	Description       commons.FieldName = "description"
)

const (
	OS            commons.FieldName = "os"
	Zones         commons.FieldName = "zones"
	PreferredZone commons.FieldName = "preferred_zone"
)
