package elastigroup_integrations

import (
	"fmt"
	"errors"

	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupCodeDeploy(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationCodeDeploy] = commons.NewGenericField(
		commons.ElastigroupIntegrations,
		IntegrationCodeDeploy,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(CleanupOnFailure): &schema.Schema{
						Type:     schema.TypeBool,
						Required: true,
					},

					string(TerminateInstanceOnFailure): &schema.Schema{
						Type:     schema.TypeBool,
						Required: true,
					},

					string(DeploymentGroups): &schema.Schema{
						Type:     schema.TypeSet,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(ApplicationName): &schema.Schema{
									Type:     schema.TypeString,
									Required: true,
								},

								string(DeploymentGroupName): &schema.Schema{
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
			elastigroup := resourceObject.(*aws.Group)
			var value []interface{} = nil
			if elastigroup.Integration != nil && elastigroup.Integration.CodeDeploy != nil {
				value = flattenAWSGroupCodeDeployIntegration(elastigroup.Integration.CodeDeploy)
			}
			if value != nil {
				if err := resourceData.Set(string(IntegrationCodeDeploy), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationCodeDeploy), err)
				}
			} else {
				if err := resourceData.Set(string(IntegrationCodeDeploy), []*aws.CodeDeployIntegration{}); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationCodeDeploy), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
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
			elastigroup := resourceObject.(*aws.Group)
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
	list := data.([]interface{})
	m := list[0].(map[string]interface{})
	i := &aws.CodeDeployIntegration{}

	if v, ok := m[string(CleanupOnFailure)].(bool); ok {
		i.SetCleanUpOnFailure(spotinst.Bool(v))
	}

	if v, ok := m[string(TerminateInstanceOnFailure)].(bool); ok {
		i.SetTerminateInstanceOnFailure(spotinst.Bool(v))
	}

	if v, ok := m[string(DeploymentGroups)]; ok {
		deploymentGroups, err := expandAWSGroupCodeDeployIntegrationDeploymentGroups(v)
		if err != nil {
			return nil, err
		}
		i.SetDeploymentGroups(deploymentGroups)
	}
	return i, nil
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

func flattenAWSGroupCodeDeployIntegration(integration *aws.CodeDeployIntegration) []interface{} {
	result := make(map[string]interface{})
	result[string(CleanupOnFailure)] = spotinst.BoolValue(integration.CleanUpOnFailure)
	result[string(TerminateInstanceOnFailure)] = spotinst.BoolValue(integration.TerminateInstanceOnFailure)

	deploymentGroups := make([]interface{}, len(integration.DeploymentGroups))
	for i, dg := range integration.DeploymentGroups {
		m := make(map[string]interface{})
		m[string(ApplicationName)] = spotinst.StringValue(dg.ApplicationName)
		m[string(DeploymentGroupName)] = spotinst.StringValue(dg.DeploymentGroupName)
		deploymentGroups[i] = m
	}

	return []interface{}{result}
}

