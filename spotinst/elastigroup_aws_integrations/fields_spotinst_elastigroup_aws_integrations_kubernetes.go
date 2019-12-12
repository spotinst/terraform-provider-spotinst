package elastigroup_aws_integrations

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupKubernetes(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationKubernetes] = commons.NewGenericField(
		commons.ElastigroupAWSIntegrations,
		IntegrationKubernetes,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(IntegrationMode): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(ClusterIdentifier): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(ApiServer): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(Token): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(AutoscaleIsEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(AutoscaleCooldown): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(AutoscaleIsAutoConfig): {
						Type:     schema.TypeBool,
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

					string(AutoscaleLabels): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Key): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(Value): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
						Set: labelHashKV,
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
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
	integration := &aws.KubernetesIntegration{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return integration, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(IntegrationMode)].(string); ok && v != "" {
		integration.SetIntegrationMode(spotinst.String(v))
	}

	if v, ok := m[string(ClusterIdentifier)].(string); ok && v != "" {
		integration.SetClusterIdentifier(spotinst.String(v))
	}

	if v, ok := m[string(ApiServer)].(string); ok && v != "" {
		integration.SetServer(spotinst.String(v))
	}

	if v, ok := m[string(Token)].(string); ok && v != "" {
		integration.SetToken(spotinst.String(v))
	}

	if v, ok := m[string(AutoscaleIsEnabled)].(bool); ok {
		if integration.AutoScale == nil {
			integration.SetAutoScale(&aws.AutoScaleKubernetes{})
		}
		integration.AutoScale.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoscaleCooldown)].(int); ok && v > 0 {
		if integration.AutoScale == nil {
			integration.SetAutoScale(&aws.AutoScaleKubernetes{})
		}
		integration.AutoScale.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m[string(AutoscaleIsAutoConfig)].(bool); ok {
		if integration.AutoScale == nil {
			integration.SetAutoScale(&aws.AutoScaleKubernetes{})
		}
		integration.AutoScale.SetIsAutoConfig(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoscaleHeadroom)]; ok {
		headroom, err := expandAWSGroupAutoScaleHeadroom(v)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			if integration.AutoScale == nil {
				integration.SetAutoScale(&aws.AutoScaleKubernetes{})
			}
			integration.AutoScale.SetHeadroom(headroom)
		}
	}

	if v, ok := m[string(AutoscaleDown)]; ok {
		down, err := expandAWSGroupAutoScaleDown(v)
		if err != nil {
			return nil, err
		}
		if down != nil {
			if integration.AutoScale == nil {
				integration.SetAutoScale(&aws.AutoScaleKubernetes{})
			}
			integration.AutoScale.SetDown(down)
		}
	}

	if v, ok := m[string(AutoscaleLabels)]; ok {
		labels, err := expandKubernetesAutoScaleLabels(v)
		if err != nil {
			return nil, err
		}
		if labels != nil {
			if integration.AutoScale == nil {
				integration.SetAutoScale(&aws.AutoScaleKubernetes{})
			}
			integration.AutoScale.SetLabels(labels)
		}
	}
	return integration, nil
}

func expandKubernetesAutoScaleLabels(data interface{}) ([]*aws.AutoScaleLabel, error) {
	list := data.(*schema.Set).List()
	out := make([]*aws.AutoScaleLabel, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(Key)]; !ok {
			return nil, errors.New("invalid Kubernetes label: key missing")
		}

		if _, ok := attr[string(Value)]; !ok {
			return nil, errors.New("invalid Kubernetes label: value missing")
		}
		c := &aws.AutoScaleLabel{
			Key:   spotinst.String(attr[string(Key)].(string)),
			Value: spotinst.String(attr[string(Value)].(string)),
		}
		out = append(out, c)
	}
	return out, nil
}

func labelHashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(Key)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(Value)].(string)))
	return hashcode.String(buf.String())
}
