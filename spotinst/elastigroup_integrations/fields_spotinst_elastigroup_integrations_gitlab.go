package elastigroup_integrations

import (
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
					string(GitlabRunner): &schema.Schema{
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(GitlabRunnerIsEnabled): &schema.Schema{
									Type:     schema.TypeBool,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
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
func expandAWSGroupGitlabIntegration(data interface{}) (*aws.GitlabIntegration, error) {
	integration := &aws.GitlabIntegration{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(GitlabRunner)]; ok {
			runner, err := expandAWSGroupGitlabRunner(v)
			if err != nil {
				return nil, err
			}
			if runner != nil {
				integration.SetRunner(runner)
			}
		}
	}
	return integration, nil
}

func expandAWSGroupGitlabRunner(data interface{}) (*aws.GitlabRunner, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		runner := &aws.GitlabRunner{}
		m := list[0].(map[string]interface{})

		if v, ok := m[string(GitlabRunnerIsEnabled)].(bool); ok {
			runner.SetIsEnabled(spotinst.Bool(v))
		}
		return runner, nil
	}

	return nil, nil
}
