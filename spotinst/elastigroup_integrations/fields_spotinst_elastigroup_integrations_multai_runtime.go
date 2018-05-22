package elastigroup_integrations

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupMultaiRuntime(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationMultaiRuntime] = commons.NewGenericField(
		commons.ElastigroupIntegrations,
		IntegrationMultaiRuntime,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(DeploymentId): &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value []interface{} = nil
			if elastigroup.Integration != nil && elastigroup.Integration.Multai != nil {
				value = flattenAWSGroupMultaiIntegration(elastigroup.Integration.Multai)
			}
			if value != nil {
				if err := resourceData.Set(string(IntegrationMultaiRuntime), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationMultaiRuntime), err)
				}
			} else {
				if err := resourceData.Set(string(IntegrationMultaiRuntime), []*aws.MultaiIntegration{}); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationMultaiRuntime), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.GetOk(string(IntegrationMultaiRuntime)); ok {
				if integration, err := expandAWSGroupMultaiIntegration(v); err != nil {
					return err
				} else {
					elastigroup.Integration.SetMultai(integration)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *aws.MultaiIntegration = nil
			if v, ok := resourceData.GetOk(string(IntegrationMultaiRuntime)); ok {
				if integration, err := expandAWSGroupMultaiIntegration(v); err != nil {
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
func expandAWSGroupMultaiIntegration(data interface{}) (*aws.MultaiIntegration, error) {
	list := data.([]interface{})
	m := list[0].(map[string]interface{})
	i := &aws.MultaiIntegration{}

	if v, ok := m[string(DeploymentId)].(string); ok && v != "" {
		i.SetDeploymentId(spotinst.String(v))
	}
	return i, nil
}

func flattenAWSGroupMultaiIntegration(integration *aws.MultaiIntegration) []interface{} {
	result := make(map[string]interface{})
	result[string(DeploymentId)] = spotinst.StringValue(integration.DeploymentID)
	return []interface{}{result}
}
