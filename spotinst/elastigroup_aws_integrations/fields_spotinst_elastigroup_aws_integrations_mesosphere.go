package elastigroup_aws_integrations

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupMesosphere(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationMesosphere] = commons.NewGenericField(
		commons.ElastigroupAWSIntegrations,
		IntegrationMesosphere,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ApiServer): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
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
	integration := &aws.MesosphereIntegration{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(ApiServer)].(string); ok && v != "" {
			integration.SetServer(spotinst.String(v))
		}
	}
	return integration, nil
}

func flattenAWSGroupMesosphereIntegration(integration *aws.MesosphereIntegration) []interface{} {
	result := make(map[string]interface{})
	result[string(ApiServer)] = spotinst.StringValue(integration.Server)
	return []interface{}{result}
}
