package stateful_node_azure_extension

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Extensions              commons.FieldName = "extensions" //TODO - should this be of type list?
	APIVersion              commons.FieldName = "api_version"
	MinorVersionAutoUpgrade commons.FieldName = "minor_version_auto_upgrade"
	Name                    commons.FieldName = "name"
	Publisher               commons.FieldName = "publisher"
	Type                    commons.FieldName = "type"
	ProtectedSettings       commons.FieldName = "protected_settings"
	PublicSettings          commons.FieldName = "public_settings"
)
