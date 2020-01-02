package elastigroup_aws_integrations

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupGitlab(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationGitlab] = commons.NewGenericField(
		commons.ElastigroupAWSIntegrations,
		IntegrationGitlab,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(GitlabRunner): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(GitlabRunnerIsEnabled): {
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
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

		var runnerResult *aws.GitlabRunner = nil
		if v, ok := m[string(GitlabRunner)]; ok {
			runner, err := expandAWSGroupGitlabRunner(v)
			if err != nil {
				return nil, err
			}
			if runner != nil {
				runnerResult = runner
			}
		}

		integration.SetRunner(runnerResult)
	}
	return integration, nil
}

func expandAWSGroupGitlabRunner(data interface{}) (*aws.GitlabRunner, error) {
	if list := data.([]interface{}); len(list) > 0 && list[0] != nil {
		runner := &aws.GitlabRunner{}
		m := list[0].(map[string]interface{})

		var isEnabled = spotinst.Bool(false)
		if v, ok := m[string(GitlabRunnerIsEnabled)].(bool); ok {
			isEnabled = spotinst.Bool(v)
		}

		runner.SetIsEnabled(isEnabled)
		return runner, nil
	}

	return nil, nil
}
