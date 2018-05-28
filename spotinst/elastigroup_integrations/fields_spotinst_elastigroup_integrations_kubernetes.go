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
func SetupKubernetes(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationKubernetes] = commons.NewGenericField(
		commons.ElastigroupIntegrations,
		IntegrationKubernetes,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(IntegrationMode): &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},

					string(ClusterIdentifier): &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},

					string(ApiServer): &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},

					string(Token): &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},

					string(AutoscaleIsEnabled): &schema.Schema{
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(AutoscaleCooldown): &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(AutoscaleIsAutoConfig): &schema.Schema{
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(AutoscaleHeadroom): &schema.Schema{
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(CpuPerUnit): &schema.Schema{
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(MemoryPerUnit): &schema.Schema{
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(NumOfUnits): &schema.Schema{
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},

					string(AutoscaleDown): &schema.Schema{
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(EvaluationPeriods): &schema.Schema{
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
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.GetOk(string(IntegrationKubernetes)); ok {
				if integration, err := expandAWSGroupKubernetesIntegration(v); err != nil {
					return err
				} else {
					elastigroup.Integration.SetKubernetes(integration)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *aws.KubernetesIntegration = nil
			if v, ok := resourceData.GetOk(string(IntegrationKubernetes)); ok {
				if integration, err := expandAWSGroupKubernetesIntegration(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetKubernetes(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupKubernetesIntegration(data interface{}) (*aws.KubernetesIntegration, error) {
	list := data.([]interface{})
	m := list[0].(map[string]interface{})
	i := &aws.KubernetesIntegration{}

	if v, ok := m[string(IntegrationMode)].(string); ok && v != "" {
		i.SetIntegrationMode(spotinst.String(v))
	}

	if v, ok := m[string(ClusterIdentifier)].(string); ok && v != "" {
		i.SetClusterIdentifier(spotinst.String(v))
	}

	if v, ok := m[string(ApiServer)].(string); ok && v != "" {
		i.SetServer(spotinst.String(v))
	}

	if v, ok := m[string(Token)].(string); ok && v != "" {
		i.SetToken(spotinst.String(v))
	}

	if v, ok := m[string(AutoscaleIsEnabled)].(bool); ok {
		if i.AutoScale == nil {
			i.SetAutoScale(&aws.AutoScale{})
		}
		i.AutoScale.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoscaleCooldown)].(int); ok && v > 0 {
		if i.AutoScale == nil {
			i.SetAutoScale(&aws.AutoScale{})
		}
		i.AutoScale.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m[string(AutoscaleIsAutoConfig)].(bool); ok {
		if i.AutoScale == nil {
			i.SetAutoScale(&aws.AutoScale{})
		}
		i.AutoScale.SetIsAutoConfig(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoscaleHeadroom)]; ok {
		headroom, err := expandAWSGroupAutoScaleHeadroom(v)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			if i.AutoScale == nil {
				i.SetAutoScale(&aws.AutoScale{})
			}
			i.AutoScale.SetHeadroom(headroom)
		}
	}

	if v, ok := m[string(AutoscaleDown)]; ok {
		down, err := expandAWSGroupAutoScaleDown(v)
		if err != nil {
			return nil, err
		}
		if down != nil {
			if i.AutoScale == nil {
				i.SetAutoScale(&aws.AutoScale{})
			}
			i.AutoScale.SetDown(down)
		}
	}
	return i, nil
}
