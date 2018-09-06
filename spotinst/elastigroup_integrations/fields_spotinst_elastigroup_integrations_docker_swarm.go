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
func SetupDockerSwarm(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[IntegrationDockerSwarm] = commons.NewGenericField(
		commons.ElastigroupIntegrations,
		IntegrationDockerSwarm,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MasterHost): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(MasterPort): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(AutoscaleIsEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(AutoscaleCooldown): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(AutoscaleHeadroom): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(CpuPerUnit): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(MemoryPerUnit): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(NumOfUnits): {
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},

					string(AutoscaleDown): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(EvaluationPeriods): {
									Type:     schema.TypeInt,
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
			if v, ok := resourceData.GetOk(string(IntegrationDockerSwarm)); ok {
				if integration, err := expandAWSGroupDockerSwarmIntegration(v, false); err != nil {
					return err
				} else {
					elastigroup.Integration.SetDockerSwarm(integration)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.DockerSwarmIntegration = nil

			if v, ok := resourceData.GetOk(string(IntegrationDockerSwarm)); ok {
				if integration, err := expandAWSGroupDockerSwarmIntegration(v, true); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetDockerSwarm(value)
			return nil
		},

		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupDockerSwarmIntegration(data interface{}, nullify bool) (*aws.DockerSwarmIntegration, error) {
	integration := &aws.DockerSwarmIntegration{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return integration, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(MasterHost)].(string); ok && v != "" {
		integration.SetMasterHost(spotinst.String(v))
	}

	if v, ok := m[string(MasterPort)].(int); ok && v > 0 {
		integration.SetMasterPort(spotinst.Int(v))
	}

	if v, ok := m[string(AutoscaleIsEnabled)].(bool); ok {
		if integration.AutoScaleDockerSwarm == nil {
			integration.SetAutoScaleDockerSwarm(&aws.AutoScaleDockerSwarm{})
		}
		integration.AutoScaleDockerSwarm.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoscaleCooldown)].(int); ok && v > 0 {
		if integration.AutoScaleDockerSwarm == nil {
			integration.SetAutoScaleDockerSwarm(&aws.AutoScaleDockerSwarm{})
		}
		integration.AutoScaleDockerSwarm.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m[string(AutoscaleHeadroom)]; ok {
		headroom, err := expandAWSGroupAutoScaleHeadroom(v)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			if integration.AutoScaleDockerSwarm == nil {
				integration.SetAutoScaleDockerSwarm(&aws.AutoScaleDockerSwarm{})
			}
			integration.AutoScaleDockerSwarm.SetHeadroom(headroom)
		} else {
			integration.AutoScaleDockerSwarm.Headroom = nil
		}
	}

	if v, ok := m[string(AutoscaleDown)]; ok {
		down, err := expandAWSGroupAutoScaleDown(v)
		if err != nil {
			return nil, err
		}
		if down != nil {
			if integration.AutoScaleDockerSwarm == nil {
				integration.SetAutoScaleDockerSwarm(&aws.AutoScaleDockerSwarm{})
			}
			integration.AutoScaleDockerSwarm.SetDown(down)
		}
	}

	return integration, nil
}
