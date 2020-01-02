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
func SetupNomad(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationNomad] = commons.NewGenericField(
		commons.ElastigroupAWSIntegrations,
		IntegrationNomad,
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

					string(AclToken): {
						Type:     schema.TypeString,
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

					string(AutoscaleConstraints): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Key): {
									Type:      schema.TypeString,
									Required:  true,
									StateFunc: attrStateFunc,
								},

								string(Value): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
						Set: constraintHashKV,
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
			if v, ok := resourceData.GetOk(string(IntegrationNomad)); ok {
				if integration, err := expandAWSGroupNomadIntegration(v, false); err != nil {
					return err
				} else {
					elastigroup.Integration.SetNomad(integration)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.NomadIntegration = nil
			if v, ok := resourceData.GetOk(string(IntegrationNomad)); ok {
				if integration, err := expandAWSGroupNomadIntegration(v, true); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetNomad(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupNomadIntegration(data interface{}, nullify bool) (*aws.NomadIntegration, error) {
	integration := &aws.NomadIntegration{}
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

	if v, ok := m[string(AclToken)].(string); ok && v != "" {
		integration.SetAclToken(spotinst.String(v))
	} else if nullify {
		integration.SetAclToken(nil)
	}

	if v, ok := m[string(AutoscaleIsEnabled)].(bool); ok {
		if integration.AutoScale == nil {
			integration.SetAutoScale(&aws.AutoScaleNomad{})
		}
		integration.AutoScale.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoscaleCooldown)].(int); ok && v > 0 {
		if integration.AutoScale == nil {
			integration.SetAutoScale(&aws.AutoScaleNomad{})
		}
		integration.AutoScale.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m[string(AutoscaleHeadroom)]; ok {
		headroom, err := expandAWSGroupAutoScaleHeadroom(v)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			if integration.AutoScale == nil {
				integration.SetAutoScale(&aws.AutoScaleNomad{})
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
				integration.SetAutoScale(&aws.AutoScaleNomad{})
			}
			integration.AutoScale.SetDown(down)
		}
	}

	if v, ok := m[string(AutoscaleConstraints)]; ok {
		consts, err := expandNomadAutoScaleConstraints(v)
		if err != nil {
			return nil, err
		}
		if consts != nil {
			if integration.AutoScale == nil {
				integration.SetAutoScale(&aws.AutoScaleNomad{})
			}
			integration.AutoScale.SetConstraints(consts)
		}
	}
	return integration, nil
}

func expandNomadAutoScaleConstraints(data interface{}) ([]*aws.AutoScaleConstraint, error) {
	list := data.(*schema.Set).List()
	out := make([]*aws.AutoScaleConstraint, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(Key)]; !ok {
			return nil, errors.New("invalid Nomad constraint: key missing")
		}

		if _, ok := attr[string(Value)]; !ok {
			return nil, errors.New("invalid Nomad constraint: value missing")
		}
		c := &aws.AutoScaleConstraint{
			Key:   spotinst.String(fmt.Sprintf("${%s}", attr[string(Key)].(string))),
			Value: spotinst.String(attr[string(Value)].(string)),
		}
		out = append(out, c)
	}
	return out, nil
}

func attrStateFunc(v interface{}) string {
	switch s := v.(type) {
	case string:
		return fmt.Sprintf("${%s}", s)
	default:
		return ""
	}
}

func constraintHashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(Key)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(Value)].(string)))
	return hashcode.String(buf.String())
}
