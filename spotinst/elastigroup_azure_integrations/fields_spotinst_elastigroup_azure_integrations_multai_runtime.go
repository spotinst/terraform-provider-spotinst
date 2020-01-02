package elastigroup_azure_integrations

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupMultaiRuntime(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationMultaiRuntime] = commons.NewGenericField(
		commons.ElastigroupAzureIntegrations,
		IntegrationMultaiRuntime,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(DeploymentId): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{} = nil
			if elastigroup.Integration != nil && elastigroup.Integration.Multai != nil {
				value = flattenAureGroupMultaiIntegration(elastigroup.Integration.Multai)
			}
			if value != nil {
				if err := resourceData.Set(string(IntegrationMultaiRuntime), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationMultaiRuntime), err)
				}
			} else {
				if err := resourceData.Set(string(IntegrationMultaiRuntime), []*azure.MultaiIntegration{}); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationMultaiRuntime), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(IntegrationMultaiRuntime)); ok {
				if integration, err := expandAzureGroupMultaiIntegration(v); err != nil {
					return err
				} else {
					elastigroup.Integration.SetMultai(integration)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *azure.MultaiIntegration = nil
			if v, ok := resourceData.GetOk(string(IntegrationMultaiRuntime)); ok {
				if integration, err := expandAzureGroupMultaiIntegration(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetMultai(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAzureGroupMultaiIntegration(data interface{}) (*azure.MultaiIntegration, error) {
	integration := &azure.MultaiIntegration{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(DeploymentId)].(string); ok && v != "" {
			integration.SetDeploymentId(spotinst.String(v))
		}
	}
	return integration, nil
}

func flattenAureGroupMultaiIntegration(integration *azure.MultaiIntegration) []interface{} {
	result := make(map[string]interface{})
	result[string(DeploymentId)] = spotinst.StringValue(integration.DeploymentID)
	return []interface{}{result}
}
