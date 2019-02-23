package elastigroup_gcp_integrations

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	SetupDockerSwarm(fieldsMap)
}
