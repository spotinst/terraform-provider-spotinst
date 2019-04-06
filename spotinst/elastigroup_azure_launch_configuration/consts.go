package elastigroup_azure_launch_configuration

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "azure_launch_configuration_"
)

const (
	UserData       commons.FieldName = "user_data"
	ShutdownScript commons.FieldName = "shutdown_script"
	CustomData     commons.FieldName = "custom_data"
)
