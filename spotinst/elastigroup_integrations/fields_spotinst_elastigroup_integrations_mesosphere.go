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
func SetupMesosphere(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationMesosphere] = commons.NewGenericField(
		commons.ElastigroupIntegrations,
		IntegrationMesosphere,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ApiServer): &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value []interface{} = nil
			if elastigroup.Integration != nil && elastigroup.Integration.Mesosphere != nil {
				value = flattenAWSGroupMesosphereIntegration(elastigroup.Integration.Mesosphere)
			}
			if value != nil {
				if err := resourceData.Set(string(IntegrationMesosphere), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationMesosphere), err)
				}
			} else {
				if err := resourceData.Set(string(IntegrationMesosphere), []*aws.MesosphereIntegration{}); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationMesosphere), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.GetOk(string(IntegrationMesosphere)); ok {
				if integration, err := expandAWSGroupMesosphereIntegration(v); err != nil {
					return err
				} else {
					elastigroup.Integration.SetMesosphere(integration)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *aws.MesosphereIntegration = nil
			if v, ok := resourceData.GetOk(string(IntegrationMesosphere)); ok {
				if integration, err := expandAWSGroupMesosphereIntegration(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetMesosphere(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupMesosphereIntegration(data interface{}) (*aws.MesosphereIntegration, error) {
	list := data.([]interface{})
	m := list[0].(map[string]interface{})
	i := &aws.MesosphereIntegration{}

	if v, ok := m[string(ApiServer)].(string); ok && v != "" {
		i.SetServer(spotinst.String(v))
	}
	return i, nil
}

func flattenAWSGroupMesosphereIntegration(integration *aws.MesosphereIntegration) []interface{} {
	result := make(map[string]interface{})
	result[string(ApiServer)] = spotinst.StringValue(integration.Server)
	return []interface{}{result}
}