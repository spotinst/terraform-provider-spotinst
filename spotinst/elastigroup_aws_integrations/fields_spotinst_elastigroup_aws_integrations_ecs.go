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
func SetupEcs(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationEcs] = commons.NewGenericField(
		commons.ElastigroupAWSIntegrations,
		IntegrationEcs,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ClusterName): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(ShouldScaleDownNonServiceTasks): {
						Type:     schema.TypeBool,
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

								string(MaxScaleDownPercentage): {
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},

					string(AutoscaleAttributes): {
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
						Set: attributeHashKV,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup != nil && elastigroup.Integration != nil && elastigroup.Integration.EC2ContainerService != nil {
				result = flattenECSIntegration(elastigroup.Integration.EC2ContainerService)
			}

			if result != nil {
				if err := resourceData.Set(string(IntegrationEcs), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationEcs), err)
				}
			}
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
		if integration.AutoScale == nil {
			integration.SetAutoScale(&aws.AutoScaleECS{})
		}
		integration.AutoScale.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(ShouldScaleDownNonServiceTasks)].(bool); ok {
		if integration.AutoScale == nil {
			integration.SetAutoScale(&aws.AutoScaleECS{})
		}
		integration.AutoScale.SetShouldScaleDownNonServiceTasks(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoscaleCooldown)].(int); ok && v > 0 {
		if integration.AutoScale == nil {
			integration.SetAutoScale(&aws.AutoScaleECS{})
		}
		integration.AutoScale.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m[string(AutoscaleIsAutoConfig)].(bool); ok {
		if integration.AutoScale == nil {
			integration.SetAutoScale(&aws.AutoScaleECS{})
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
				integration.SetAutoScale(&aws.AutoScaleECS{})
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
				integration.SetAutoScale(&aws.AutoScaleECS{})
			}
			integration.AutoScale.SetDown(down)
		}
	}

	if v, ok := m[string(AutoscaleAttributes)]; ok {
		attributes, err := expandECSAutoScaleAttributes(v)
		if err != nil {
			return nil, err
		}
		if attributes != nil {
			if integration.AutoScale == nil {
				integration.SetAutoScale(&aws.AutoScaleECS{})
			}
			integration.AutoScale.SetAttributes(attributes)
		}
	}
	return integration, nil
}

func flattenAutoScaleHeadroom(autoScaleHeadroom *aws.AutoScaleHeadroom) []interface{} {
	headRoom := make(map[string]interface{})
	headRoom[string(CpuPerUnit)] = spotinst.IntValue(autoScaleHeadroom.CPUPerUnit)
	headRoom[string(MemoryPerUnit)] = spotinst.IntValue(autoScaleHeadroom.MemoryPerUnit)
	headRoom[string(NumOfUnits)] = spotinst.IntValue(autoScaleHeadroom.NumOfUnits)
	return []interface{}{headRoom}
}

func flattenAutoScaleDown(autoScaleDown *aws.AutoScaleDown) []interface{} {
	down := make(map[string]interface{})
	down[string(EvaluationPeriods)] = spotinst.IntValue(autoScaleDown.EvaluationPeriods)
	down[string(MaxScaleDownPercentage)] = spotinst.IntValue(autoScaleDown.MaxScaleDownPercentage)
	return []interface{}{down}
}

func flattenAutoScale(autoScale *aws.AutoScale) []interface{} {
	result := make(map[string]interface{})
	result[string(AutoscaleIsEnabled)] = spotinst.BoolValue(autoScale.IsEnabled)
	result[string(AutoscaleIsEnabled)] = spotinst.BoolValue(autoScale.IsEnabled)
	result[string(AutoscaleIsAutoConfig)] = spotinst.BoolValue(autoScale.IsAutoConfig)
	result[string(AutoscaleCooldown)] = spotinst.IntValue(autoScale.Cooldown)
	if autoScale.Headroom != nil {
		result[string(AutoscaleHeadroom)] = flattenAutoScaleHeadroom(autoScale.Headroom)
	}
	if autoScale.Down != nil {
		result[string(AutoscaleDown)] = flattenAutoScaleDown(autoScale.Down)
	}
	return []interface{}{result}
}

func flattenECSIntegrationAutoScaleAttributes(attrs []*aws.AutoScaleAttributes) []interface{} {
	result := make(map[string]interface{})

	for _, attr := range attrs {
		result[string(Key)] = spotinst.StringValue(attr.Key)
		result[string(Value)] = spotinst.StringValue(attr.Value)
	}

	return []interface{}{result}
}

func flattenECSIntegration(ecs *aws.EC2ContainerServiceIntegration) []interface{} {
	result := make(map[string]interface{})
	result[string(ClusterName)] = spotinst.StringValue(ecs.ClusterName)

	if ecs.AutoScale != nil {
		if autoScale := flattenAutoScale(&ecs.AutoScale.AutoScale); len(autoScale) > 0 {
			for k, v := range autoScale[0].(map[string]interface{}) {
				result[k] = v
			}
		}
		if ecs.AutoScale.Attributes != nil {
			result[string(AutoscaleAttributes)] = flattenECSIntegrationAutoScaleAttributes(ecs.AutoScale.Attributes)
		}
	}

	return []interface{}{result}
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
