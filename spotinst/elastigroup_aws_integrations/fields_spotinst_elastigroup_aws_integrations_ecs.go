package elastigroup_aws_integrations

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

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
						Default:  300,
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
									Type:     schema.TypeFloat,
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
					},

					string(Batch): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(JobQueueNames): {
									Type:     schema.TypeList,
									Required: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
							},
						},
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
		down, err := expandAWSGroupAutoScaleDown(v, true)
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

	if v, ok := m[string(Batch)]; ok {
		batch, err := expandECSBatch(v)
		if err != nil {
			return nil, err
		}
		integration.SetBatch(batch)
	}
	return integration, nil
}

func expandECSBatch(data interface{}) (*aws.Batch, error) {
	if list := data.([]interface{}); len(list) > 0 {
		batch := &aws.Batch{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})
			var jobQueueNames []string = nil
			if v, ok := m[string(JobQueueNames)].([]interface{}); ok && len(v) > 0 {
				jobQueueNamesList := make([]string, 0, len(v))
				for _, jobQueueName := range v {
					if v, ok := jobQueueName.(string); ok {
						jobQueueNamesList = append(jobQueueNamesList, v)
					}
				}
				jobQueueNames = jobQueueNamesList
			}
			batch.SetJobQueueNames(jobQueueNames)
		}
		return batch, nil
	}
	return nil, nil
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
	down[string(MaxScaleDownPercentage)] = spotinst.Float64Value(autoScaleDown.MaxScaleDownPercentage)
	return []interface{}{down}
}

func flattenAutoScale(autoScale *aws.AutoScale) []interface{} {
	result := make(map[string]interface{})
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
	result := make([]interface{}, 0, len(attrs))

	for _, attr := range attrs {
		m := make(map[string]interface{})
		m[string(Key)] = spotinst.StringValue(attr.Key)
		m[string(Value)] = spotinst.StringValue(attr.Value)

		result = append(result, m)
	}

	return result
}

func flattenECSIntegration(ecs *aws.EC2ContainerServiceIntegration) []interface{} {
	result := make(map[string]interface{})
	result[string(ClusterName)] = spotinst.StringValue(ecs.ClusterName)
	result[string(ShouldScaleDownNonServiceTasks)] = spotinst.BoolValue(ecs.AutoScale.ShouldScaleDownNonServiceTasks)

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

	if ecs.Batch != nil {
		result[string(Batch)] = flattenECSBatch(ecs.Batch)
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

func flattenECSBatch(batch *aws.Batch) []interface{} {
	result := make(map[string]interface{})
	if len(batch.JobQueueNames) > 0 {
		result[string(JobQueueNames)] = batch.JobQueueNames
	}
	return []interface{}{result}
}
