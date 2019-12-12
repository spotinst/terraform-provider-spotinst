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
func SetupElasticBeanstalk(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[IntegrationBeanstalk] = commons.NewGenericField(
		commons.ElastigroupAWSIntegrations,
		IntegrationBeanstalk,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(EnvironmentId): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(DeploymentPreferences): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(AutomaticRoll): {
									Type:     schema.TypeBool,
									Optional: true,
								},

								string(BatchSizePercentage): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(GracePeriod): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(Strategy): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Action): {
												Type:     schema.TypeString,
												Optional: true,
											},

											string(ShouldDrainInstances): {
												Type:     schema.TypeBool,
												Optional: true,
											},
										},
									},
								},
							},
						},
					},

					string(ManagedActions): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(PlatformUpdate): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(PerformAt): {
												Type:     schema.TypeString,
												Optional: true,
											},

											string(TimeWindow): {
												Type:     schema.TypeString,
												Optional: true,
											},

											string(UpdateLevel): {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
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
			if v, ok := resourceData.GetOk(string(IntegrationBeanstalk)); ok {
				if integration, err := expandAWSGroupElasticBeanstalkIntegration(v, false); err != nil {
					return err
				} else {
					elastigroup.Integration.SetElasticBeanstalk(integration)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.ElasticBeanstalkIntegration = nil

			if v, ok := resourceData.GetOk(string(IntegrationBeanstalk)); ok {
				if integration, err := expandAWSGroupElasticBeanstalkIntegration(v, true); err != nil {
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
func expandAWSGroupElasticBeanstalkIntegration(data interface{}, nullify bool) (*aws.ElasticBeanstalkIntegration, error) {
	integration := &aws.ElasticBeanstalkIntegration{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return integration, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(EnvironmentId)].(string); ok {
		integration.SetEnvironmentID(spotinst.String(v))
	}

	if v, ok := m[string(DeploymentPreferences)]; ok {
		integration.SetDeploymentPreferences(&aws.BeanstalkDeploymentPreferences{})

		list := v.([]interface{})
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(AutomaticRoll)].(bool); ok {
				integration.DeploymentPreferences.SetAutomaticRoll(spotinst.Bool(v))
			}

			if v, ok := m[string(BatchSizePercentage)].(int); ok {
				integration.DeploymentPreferences.SetBatchSizePercentage(spotinst.Int(v))
			}

			if v, ok := m[string(GracePeriod)].(int); ok {
				integration.DeploymentPreferences.SetGracePeriod(spotinst.Int(v))
			}

			if v, ok := m[string(Strategy)]; ok {
				integration.DeploymentPreferences.SetStrategy(&aws.BeanstalkDeploymentStrategy{})

				list := v.([]interface{})
				if list != nil && list[0] != nil {
					m := list[0].(map[string]interface{})

					if v, ok := m[string(Action)].(string); ok {
						integration.DeploymentPreferences.Strategy.SetAction(spotinst.String(v))
					}

					if v, ok := m[string(ShouldDrainInstances)].(bool); ok {
						integration.DeploymentPreferences.Strategy.SetShouldDrainInstances(spotinst.Bool(v))
					}
				}
			}
		}
	}

	if v, ok := m[string(ManagedActions)]; ok && len(v.([]interface{})) > 0 {
		integration.SetManagedActions(&aws.BeanstalkManagedActions{})

		list := v.([]interface{})

		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(PlatformUpdate)]; ok {
				integration.ManagedActions.SetPlatformUpdate(&aws.BeanstalkPlatformUpdate{})

				list := v.([]interface{})
				if list != nil && list[0] != nil {
					m := list[0].(map[string]interface{})

					if v, ok := m[string(PerformAt)].(string); ok {
						integration.ManagedActions.PlatformUpdate.SetPerformAt(spotinst.String(v))
					}

					if v, ok := m[string(TimeWindow)].(string); ok {
						integration.ManagedActions.PlatformUpdate.SetTimeWindow(spotinst.String(v))
					}

					if v, ok := m[string(UpdateLevel)].(string); ok {
						integration.ManagedActions.PlatformUpdate.SetUpdateLevel(spotinst.String(v))
					}
				}
			}
		}
	}

	return integration, nil
}
