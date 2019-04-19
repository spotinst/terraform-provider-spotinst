package elastigroup_azure_launch_configuration

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "azure_launch_configuration_"
)

const (
	UserData                 commons.FieldName = "user_data"
	ShutdownScript           commons.FieldName = "shutdown_script"
	CustomData               commons.FieldName = "custom_data"
	ManagedServiceIdentities commons.FieldName = "managed_service_identities"

	ResourceGroupName commons.FieldName = "resource_group_name"
	Name              commons.FieldName = "name"
)
