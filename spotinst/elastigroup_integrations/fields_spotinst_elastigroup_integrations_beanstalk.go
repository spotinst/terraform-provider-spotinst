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
func SetupElasticBeanstalk(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationElasticBeanstalk] = commons.NewGenericField(
		commons.ElastigroupIntegrations,
		IntegrationElasticBeanstalk,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(EnvironmentId): &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value []interface{} = nil
			if elastigroup.Integration != nil && elastigroup.Integration.ElasticBeanstalk != nil {
				value = flattenAWSGroupElasticBeanstalkIntegration(elastigroup.Integration.ElasticBeanstalk)
			}
			if value != nil {
				if err := resourceData.Set(string(IntegrationElasticBeanstalk), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationElasticBeanstalk), err)
				}
			} else {
				if err := resourceData.Set(string(IntegrationElasticBeanstalk), []*aws.ElasticBeanstalkIntegration{}); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationElasticBeanstalk), err)
				}
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.GetOk(string(IntegrationElasticBeanstalk)); ok {
				if integration, err := expandAWSGroupElasticBeanstalkIntegration(v); err != nil {
					return err
				} else {
					elastigroup.Integration.SetElasticBeanstalk(integration)
				}
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *aws.ElasticBeanstalkIntegration = nil
			if v, ok := resourceData.GetOk(string(IntegrationElasticBeanstalk)); ok {
				if integration, err := expandAWSGroupElasticBeanstalkIntegration(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetElasticBeanstalk(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func flattenAWSGroupElasticBeanstalkIntegration(integration *aws.ElasticBeanstalkIntegration) []interface{} {
	result := make(map[string]interface{})
	result[string(EnvironmentId)] = spotinst.StringValue(integration.EnvironmentID)
	return []interface{}{result}
}

func expandAWSGroupElasticBeanstalkIntegration(data interface{}) (*aws.ElasticBeanstalkIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &aws.ElasticBeanstalkIntegration{}

	if v, ok := m[string(EnvironmentId)].(string); ok && v != "" {
		i.SetEnvironmentId(spotinst.String(v))
	}
	return i, nil
}