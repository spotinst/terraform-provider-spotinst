package elastigroup_integrations

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupGitlab(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationGitlab] = commons.NewGenericField(
		commons.ElastigroupIntegrations,
		IntegrationGitlab,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(GitlabIsEnabled): &schema.Schema{
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value []interface{} = nil
			if elastigroup.Integration != nil && elastigroup.Integration.Gitlab != nil {
				value = flattenAWSGroupGitlabIntegration(elastigroup.Integration.Gitlab)
			}
			if value != nil {
				if err := resourceData.Set(string(IntegrationGitlab), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationGitlab), err)
				}
			} else {
				if err := resourceData.Set(string(IntegrationGitlab), []*aws.GitlabIntegration{}); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationGitlab), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.GetOk(string(IntegrationGitlab)); ok {
				if integration, err := expandAWSGroupGitlabIntegration(v); err != nil {
					return err
				} else {
					elastigroup.Integration.SetGitlab(integration)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *aws.GitlabIntegration = nil
			if v, ok := resourceData.GetOk(string(IntegrationGitlab)); ok {
				if integration, err := expandAWSGroupGitlabIntegration(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetGitlab(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func flattenAWSGroupGitlabIntegration(integration *aws.GitlabIntegration) []interface{} {
	result := make(map[string]interface{})
	result[string(GitlabIsEnabled)] = spotinst.BoolValue(integration.IsEnabled)
	return []interface{}{result}
}

func expandAWSGroupGitlabIntegration(data interface{}) (*aws.GitlabIntegration, error) {
	integration := &aws.GitlabIntegration{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(GitlabIsEnabled)].(bool); ok {
			integration.SetIsEnabled(spotinst.Bool(v))
		}
	}
	return integration, nil
}
