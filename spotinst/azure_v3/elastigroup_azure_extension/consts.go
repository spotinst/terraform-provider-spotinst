package elastigroup_azure_extension

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Extensions              commons.FieldName = "extensions"
	APIVersion              commons.FieldName = "api_version"
	MinorVersionAutoUpgrade commons.FieldName = "minor_version_auto_upgrade"
	EnableAutomaticUpgrade  commons.FieldName = "enable_automatic_upgrade"
	Name                    commons.FieldName = "name"
	Publisher               commons.FieldName = "publisher"
	Type                    commons.FieldName = "type"
	ProtectedSettings       commons.FieldName = "protected_settings"
	PublicSettings          commons.FieldName = "public_settings"

	ProtectedSettingsFromKeyVault commons.FieldName = "protected_settings_from_key_vault"
	SecretUrl                     commons.FieldName = "secret_url"
	SourceVault                   commons.FieldName = "source_vault"
)
