package elastigroup_aws_integrations

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupCodeDeploy(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationCodeDeploy] = commons.NewGenericField(
		commons.ElastigroupAWSIntegrations,
		IntegrationCodeDeploy,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(CleanupOnFailure): {
						Type:     schema.TypeBool,
						Required: true,
					},

					string(TerminateInstanceOnFailure): {
						Type:     schema.TypeBool,
						Required: true,
					},

					string(DeploymentGroups): {
						Type:     schema.TypeSet,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(ApplicationName): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(DeploymentGroupName): {
									Type:     schema.TypeString,
									Required: true,
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
			if v, ok := resourceData.GetOk(string(IntegrationCodeDeploy)); ok {
				if integration, err := expandAWSGroupCodeDeployIntegration(v); err != nil {
					return err
				} else {
					elastigroup.Integration.SetCodeDeploy(integration)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.CodeDeployIntegration = nil
			if v, ok := resourceData.GetOk(string(IntegrationCodeDeploy)); ok {
				if integration, err := expandAWSGroupCodeDeployIntegration(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetCodeDeploy(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupCodeDeployIntegration(data interface{}) (*aws.CodeDeployIntegration, error) {
	integration := &aws.CodeDeployIntegration{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(CleanupOnFailure)].(bool); ok {
			integration.SetCleanUpOnFailure(spotinst.Bool(v))
		}

		if v, ok := m[string(TerminateInstanceOnFailure)].(bool); ok {
			integration.SetTerminateInstanceOnFailure(spotinst.Bool(v))
		}

		if v, ok := m[string(DeploymentGroups)]; ok {
			deploymentGroups, err := expandAWSGroupCodeDeployIntegrationDeploymentGroups(v)
			if err != nil {
				return nil, err
			}
			integration.SetDeploymentGroups(deploymentGroups)
		}
	}
	return integration, nil
}

func expandAWSGroupCodeDeployIntegrationDeploymentGroups(data interface{}) ([]*aws.DeploymentGroup, error) {
	list := data.(*schema.Set).List()
	deploymentGroups := make([]*aws.DeploymentGroup, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(ApplicationName)]; !ok {
			return nil, errors.New("invalid deployment group attributes: application_name missing")
		}

		if _, ok := attr[string(DeploymentGroupName)]; !ok {
			return nil, errors.New("invalid deployment group attributes: deployment_group_name missing")
		}
		deploymentGroup := &aws.DeploymentGroup{
			ApplicationName:     spotinst.String(attr[string(ApplicationName)].(string)),
			DeploymentGroupName: spotinst.String(attr[string(DeploymentGroupName)].(string)),
		}
		deploymentGroups = append(deploymentGroups, deploymentGroup)
	}
	return deploymentGroups, nil
}
