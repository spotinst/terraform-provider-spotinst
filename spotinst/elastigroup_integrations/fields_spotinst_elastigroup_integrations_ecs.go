package elastigroup_integrations

import (
	"fmt"
	"bytes"
	"errors"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupEcs(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationEcs] = commons.NewGenericField(
		commons.ElastigroupIntegrations,
		IntegrationEcs,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ClusterName): &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},

					string(ShouldScaleDownNonServiceTasks): &schema.Schema{
						Type:     schema.TypeBool,
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

					string(AutoscaleAttributes): &schema.Schema{
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Key): &schema.Schema{
									Type:      schema.TypeString,
									Required:  true,
								},

								string(Value): &schema.Schema{
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
						Set: attributeHashKV,
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
			if v, ok := resourceData.GetOk(string(IntegrationEcs)); ok {
				if integration, err := expandAWSGroupEC2ContainerServiceIntegration(v); err != nil {
					return err
				} else {
					elastigroup.Integration.SetEC2ContainerService(integration)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.EC2ContainerServiceIntegration = nil
			if v, ok := resourceData.GetOk(string(IntegrationEcs)); ok {
				if integration, err := expandAWSGroupEC2ContainerServiceIntegration(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetEC2ContainerService(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupEC2ContainerServiceIntegration(data interface{}) (*aws.EC2ContainerServiceIntegration, error) {
	integration := &aws.EC2ContainerServiceIntegration{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return integration, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(ClusterName)].(string); ok && v != "" {
		integration.SetClusterName(spotinst.String(v))
	}

	if v, ok := m[string(AutoscaleIsEnabled)].(bool); ok {
		if integration.AutoScaleECS == nil {
			integration.SetAutoScaleECS(&aws.AutoScaleECS{})
		}
		integration.AutoScaleECS.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(ShouldScaleDownNonServiceTasks)].(bool); ok {
		if integration.AutoScaleECS == nil {
			integration.SetAutoScaleECS(&aws.AutoScaleECS{})
		}
		integration.AutoScaleECS.SetShouldScaleDownNonServiceTasks(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoscaleCooldown)].(int); ok && v > 0 {
		if integration.AutoScaleECS == nil {
			integration.SetAutoScaleECS(&aws.AutoScaleECS{})
		}
		integration.AutoScaleECS.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m[string(AutoscaleHeadroom)]; ok {
		headroom, err := expandAWSGroupAutoScaleHeadroom(v)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			if integration.AutoScaleECS == nil {
				integration.SetAutoScaleECS(&aws.AutoScaleECS{})
			}
			integration.AutoScaleECS.SetHeadroom(headroom)
		}
	}

	if v, ok := m[string(AutoscaleDown)]; ok {
		down, err := expandAWSGroupAutoScaleDown(v)
		if err != nil {
			return nil, err
		}
		if down != nil {
			if integration.AutoScaleECS == nil {
				integration.SetAutoScaleECS(&aws.AutoScaleECS{})
			}
			integration.AutoScaleECS.SetDown(down)
		}
	}

	if v, ok := m[string(AutoscaleAttributes)]; ok {
		attributes, err := expandECSAutoScaleAttributes(v)
		if err != nil {
			return nil, err
		}
		if attributes != nil {
			if integration.AutoScaleECS == nil {
				integration.SetAutoScaleECS(&aws.AutoScaleECS{})
			}
			integration.AutoScaleECS.SetAttributes(attributes)
		}
	}
	return integration, nil
}

func expandECSAutoScaleAttributes(data interface{}) ([]*aws.AutoScaleAttributes, error) {
	list := data.(*schema.Set).List()
	out := make([]*aws.AutoScaleAttributes, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(Key)]; !ok {
			return nil, errors.New("invalid ECS attribute: key missing")
		}

		if _, ok := attr[string(Value)]; !ok {
			return nil, errors.New("invalid ECS attribute: value missing")
		}
		c := &aws.AutoScaleAttributes{
			Key:   spotinst.String(attr[string(Key)].(string)),
			Value: spotinst.String(attr[string(Value)].(string)),
		}
		out = append(out, c)
	}
	return out, nil
}

func attributeHashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(Key)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(Value)].(string)))
	return hashcode.String(buf.String())
}
