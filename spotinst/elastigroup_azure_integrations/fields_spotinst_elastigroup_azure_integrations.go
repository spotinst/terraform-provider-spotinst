package elastigroup_azure_integrations

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	SetupKubernetes(fieldsMap)
	SetupMultaiRuntime(fieldsMap)
}
