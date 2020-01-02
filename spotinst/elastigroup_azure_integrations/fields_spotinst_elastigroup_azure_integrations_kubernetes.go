package elastigroup_azure_integrations

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupKubernetes(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationKubernetes] = commons.NewGenericField(
		commons.ElastigroupAzureIntegrations,
		IntegrationKubernetes,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(ClusterIdentifier): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(IntegrationKubernetes)); ok {
				if integration, err := expandAzureGroupKubernetesIntegration(v); err != nil {
					return err
				} else {
					elastigroup.Integration.SetKubernetes(integration)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *azure.KubernetesIntegration = nil
			if v, ok := resourceData.GetOk(string(IntegrationKubernetes)); ok {
				if integration, err := expandAzureGroupKubernetesIntegration(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetKubernetes(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAzureGroupKubernetesIntegration(data interface{}) (*azure.KubernetesIntegration, error) {
	integration := &azure.KubernetesIntegration{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return integration, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(ClusterIdentifier)].(string); ok && v != "" {
		integration.SetClusterIdentifier(spotinst.String(v))
	}

	return integration, nil
}
