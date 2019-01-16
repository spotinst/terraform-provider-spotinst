package elastigroup_azure_login

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "azure_login_"
)

const (
	Login        commons.FieldName = "login"
	UserName     commons.FieldName = "user_name"
	SSHPublicKey commons.FieldName = "ssh_public_key"
	Password     commons.FieldName = "password"
)
