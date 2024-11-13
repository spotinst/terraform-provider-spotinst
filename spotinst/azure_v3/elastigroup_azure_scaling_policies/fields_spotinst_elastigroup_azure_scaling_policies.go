package elastigroup_azure_scaling_policies

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ScalingUpPolicy] = commons.NewGenericField(
		commons.ElastigroupAzureScalingPolicies,
		ScalingUpPolicy,
		baseScalingPolicySchema(),
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var policiesResult []interface{} = nil
			if elastigroup.Scaling != nil && elastigroup.Scaling.Up != nil {
				scaleUpPolicies := elastigroup.Scaling.Up
				policiesResult = flattenAzureGroupScalingPolicy(scaleUpPolicies)
			}
			if err := resourceData.Set(string(ScalingUpPolicy), policiesResult); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ScalingUpPolicy), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(ScalingUpPolicy)); ok {
				if policies, err := expandAzureGroupScalingPolicies(v); err != nil {
					return err
				} else {
					elastigroup.Scaling.SetUp(policies)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []*azurev3.ScalingPolicy = nil
			if v, ok := resourceData.GetOk(string(ScalingUpPolicy)); ok && v != nil {
				if policies, err := expandAzureGroupScalingPolicies(v); err != nil {
					return err
				} else {
					value = policies
				}
			}
			elastigroup.Scaling.SetUp(value)
			return nil
		},
		nil,
	)

	fieldsMap[ScalingDownPolicy] = commons.NewGenericField(
		commons.ElastigroupAzureScalingPolicies,
		ScalingDownPolicy,
		baseScalingPolicySchema(),
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var policiesResult []interface{} = nil
			if elastigroup.Scaling != nil && elastigroup.Scaling.Down != nil {
				scaleDownPolicies := elastigroup.Scaling.Down
				policiesResult = flattenAzureGroupScalingPolicy(scaleDownPolicies)
			}
			if err := resourceData.Set(string(ScalingDownPolicy), policiesResult); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ScalingDownPolicy), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(ScalingDownPolicy)); ok {
				if policies, err := expandAzureGroupScalingPolicies(v); err != nil {
					return err
				} else {
					elastigroup.Scaling.SetDown(policies)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []*azurev3.ScalingPolicy = nil
			if v, ok := resourceData.GetOk(string(ScalingDownPolicy)); ok && v != nil {
				if policies, err := expandAzureGroupScalingPolicies(v); err != nil {
					return err
				} else {
					value = policies
				}
			}
			elastigroup.Scaling.SetDown(value)
			return nil
		},
		nil,
	)
}

func baseScalingPolicySchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				string(PolicyName): {
					Type:     schema.TypeString,
					Required: true,
				},

				string(MetricName): {
					Type:     schema.TypeString,
					Required: true,
				},

				string(Namespace): {
					Type:     schema.TypeString,
					Required: true,
				},

				string(Statistic): {
					Type:     schema.TypeString,
					Required: true,
				},

				string(Unit): {
					Type:     schema.TypeString,
					Optional: true,
				},

				string(Cooldown): {
					Type:     schema.TypeInt,
					Required: true,
				},

				string(Source): {
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},

				string(Dimensions): {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							string(DimensionName): {
								Type:     schema.TypeString,
								Optional: true,
							},

							string(DimensionValue): {
								Type:     schema.TypeString,
								Optional: true,
							},
						},
					},
				},

				string(Threshold): {
					Type:     schema.TypeFloat,
					Required: true,
				},

				string(Operator): {
					Type:     schema.TypeString,
					Required: true,
				},

				string(EvaluationPeriods): {
					Type:     schema.TypeInt,
					Required: true,
				},

				string(Period): {
					Type:     schema.TypeInt,
					Required: true,
				},

				string(IsEnabled): {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  true,
				},

				string(Action): {
					Type:     schema.TypeList,
					Required: true,
					MaxItems: 1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							string(Minimum): {
								Type:     schema.TypeString,
								Optional: true,
							},

							string(Maximum): {
								Type:     schema.TypeString,
								Optional: true,
							},

							string(Target): {
								Type:     schema.TypeString,
								Optional: true,
							},

							string(Type): {
								Type:     schema.TypeString,
								Optional: true,
							},

							string(Adjustment): {
								Type:     schema.TypeString,
								Optional: true,
							},
						},
					},
				},
			},
		},
	}
}

func expandAzureGroupScalingPolicies(data interface{}) ([]*azurev3.ScalingPolicy, error) {
	if list := data.([]interface{}); len(list) > 0 {
		policies := make([]*azurev3.ScalingPolicy, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			policy := &azurev3.ScalingPolicy{}

			if v, ok := m[string(PolicyName)].(string); ok && v != "" {
				policy.SetPolicyName(spotinst.String(v))
			}

			if v, ok := m[string(MetricName)].(string); ok && v != "" {
				policy.SetMetricName(spotinst.String(v))
			}

			if v, ok := m[string(Namespace)].(string); ok && v != "" {
				policy.SetNamespace(spotinst.String(v))
			}

			if v, ok := m[string(Statistic)].(string); ok && v != "" {
				policy.SetStatistic(spotinst.String(v))
			}

			if v, ok := m[string(Unit)].(string); ok && v != "" {
				policy.SetUnit(spotinst.String(v))
			}

			if v, ok := m[string(Threshold)].(float64); ok && v > -1 {
				policy.SetThreshold(spotinst.Float64(v))
			}

			if v, ok := m[string(Operator)].(string); ok && v != "" {
				policy.SetOperator(spotinst.String(v))
			}

			if v, ok := m[string(Period)].(int); ok && v > 0 {
				policy.SetPeriod(spotinst.Int(v))
			}

			if v, ok := m[string(EvaluationPeriods)].(int); ok && v > 0 {
				policy.SetEvaluationPeriods(spotinst.Int(v))
			}

			if v, ok := m[string(Cooldown)].(int); ok && v > 0 {
				policy.SetCooldown(spotinst.Int(v))
			}

			if v, ok := m[string(Dimensions)]; ok {
				dimensions := expandAzureGroupScalingPolicyDimensions(v.(interface{}))
				if len(dimensions) > 0 {
					policy.SetDimensions(dimensions)
				}
			}

			if v, ok := m[string(Action)]; ok && v != nil {
				action, err := expandAzureGroupScalingPolicyAction(v)
				if err != nil {
					return nil, err
				}
				if action != nil {
					policy.SetAction(action)
				} else {
					policy.SetAction(nil)
				}
			}

			if v, ok := m[string(IsEnabled)].(bool); ok {
				policy.SetIsEnabled(spotinst.Bool(v))
			}

			if v, ok := m[string(Source)].(string); ok && v != "" {
				policy.SetSource(spotinst.String(v))
			}

			if policy.Namespace != nil {
				policies = append(policies, policy)
			}
		}

		return policies, nil
	}
	return nil, nil
}

func expandAzureGroupScalingPolicyDimensions(data interface{}) []*azurev3.Dimensions {
	list := data.([]interface{})
	dimensions := make([]*azurev3.Dimensions, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(DimensionName)]; !ok {
			continue
		}

		if _, ok := attr[string(DimensionValue)]; !ok {
			continue
		}
		dimension := &azurev3.Dimensions{
			Name:  spotinst.String(attr[string(DimensionName)].(string)),
			Value: spotinst.String(attr[string(DimensionValue)].(string)),
		}
		if (dimension.Name != nil) && (dimension.Value != nil) {
			dimensions = append(dimensions, dimension)
		}
	}
	return dimensions
}

func expandAzureGroupScalingPolicyAction(data interface{}) (*azurev3.Action, error) {
	list := data.([]interface{})
	action := &azurev3.Action{}

	if list == nil || len(list) == 0 {
		return nil, nil
	}

	m := list[0].(map[string]interface{})
	if v, ok := m[string(Type)].(string); ok && v != "" {
		action.SetType(spotinst.String(v))
	}

	if v, ok := m[string(Adjustment)].(string); ok && v != "" {
		action.SetAdjustment(spotinst.String(v))
	}

	if v, ok := m[string(Minimum)].(string); ok && v != "" {
		action.SetMinimum(spotinst.String(v))
	}

	if v, ok := m[string(Maximum)].(string); ok && v != "" {
		action.SetMaximum(spotinst.String(v))
	}

	if v, ok := m[string(Target)].(string); ok && v != "" {
		action.SetTarget(spotinst.String(v))
	}
	return action, nil
}

func flattenAzureGroupScalingPolicy(policies []*azurev3.ScalingPolicy) []interface{} {
	result := make([]interface{}, 0, len(policies))
	for _, policy := range policies {
		m := make(map[string]interface{})
		m[string(PolicyName)] = spotinst.StringValue(policy.PolicyName)
		m[string(MetricName)] = spotinst.StringValue(policy.MetricName)
		m[string(Namespace)] = spotinst.StringValue(policy.Namespace)
		m[string(Statistic)] = spotinst.StringValue(policy.Statistic)
		m[string(Unit)] = spotinst.StringValue(policy.Unit)
		m[string(Cooldown)] = spotinst.IntValue(policy.Cooldown)
		m[string(EvaluationPeriods)] = spotinst.IntValue(policy.EvaluationPeriods)
		m[string(Period)] = spotinst.IntValue(policy.Period)
		m[string(Operator)] = spotinst.StringValue(policy.Operator)
		m[string(Threshold)] = spotinst.Float64Value(policy.Threshold)
		m[string(IsEnabled)] = spotinst.BoolValue(policy.IsEnabled)
		m[string(Source)] = spotinst.StringValue(policy.Source)

		if policy.Dimensions != nil && len(policy.Dimensions) > 0 {
			dimMap := make([]interface{}, 0, len(policy.Dimensions))
			for _, dimension := range policy.Dimensions {
				d := make(map[string]interface{})
				d[string(DimensionName)] = spotinst.StringValue(dimension.Name)
				d[string(DimensionValue)] = spotinst.StringValue(dimension.Value)

				if (d[string(DimensionName)] != nil) && (d[string(DimensionValue)] != nil) {
					dimMap = append(dimMap, d)
				}
			}
			m[string(Dimensions)] = dimMap
		}

		if policy.Action != nil {
			m[string(Action)] = flattenAction(policy.Action)
		}
		result = append(result, m)
	}
	return result
}

func flattenAction(action *azurev3.Action) []interface{} {
	result := make(map[string]interface{})
	if action != nil {
		result[string(Adjustment)] = spotinst.StringValue(action.Adjustment)
		result[string(Type)] = spotinst.StringValue(action.Type)
		result[string(Maximum)] = spotinst.StringValue(action.Maximum)
		result[string(Minimum)] = spotinst.StringValue(action.Minimum)
		result[string(Target)] = spotinst.StringValue(action.Target)
	}
	return []interface{}{result}
}
