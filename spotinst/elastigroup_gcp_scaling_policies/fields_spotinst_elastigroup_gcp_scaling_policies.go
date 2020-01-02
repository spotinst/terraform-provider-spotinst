package elastigroup_gcp_scaling_policies

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ScalingUpPolicy] = commons.NewGenericField(
		commons.ElastigroupGCPScalingPolicies,
		ScalingUpPolicy,
		upDownScalingPolicySchema(),
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var policiesResult []interface{} = nil
			if elastigroup.Scaling != nil && elastigroup.Scaling.Up != nil {
				scaleUpPolicies := elastigroup.Scaling.Up
				policiesResult = flattenGCPGroupScalingPolicy(scaleUpPolicies)
			}
			if err := resourceData.Set(string(ScalingUpPolicy), policiesResult); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ScalingUpPolicy), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(ScalingUpPolicy)); ok {
				if policies, err := expandGCPGroupScalingPolicies(v); err != nil {
					return err
				} else {
					elastigroup.Scaling.SetUp(policies)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []*gcp.ScalingPolicy = nil
			if v, ok := resourceData.GetOk(string(ScalingUpPolicy)); ok && v != nil {
				if policies, err := expandGCPGroupScalingPolicies(v); err != nil {
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
		commons.ElastigroupGCPScalingPolicies,
		ScalingDownPolicy,
		upDownScalingPolicySchema(),
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var policiesResult []interface{} = nil
			if elastigroup.Scaling != nil && elastigroup.Scaling.Down != nil {
				scaleDownPolicies := elastigroup.Scaling.Down
				policiesResult = flattenGCPGroupScalingPolicy(scaleDownPolicies)
			}
			if err := resourceData.Set(string(ScalingDownPolicy), policiesResult); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ScalingDownPolicy), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(ScalingDownPolicy)); ok {
				if policies, err := expandGCPGroupScalingPolicies(v); err != nil {
					return err
				} else {
					elastigroup.Scaling.SetDown(policies)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []*gcp.ScalingPolicy = nil
			if v, ok := resourceData.GetOk(string(ScalingDownPolicy)); ok && v != nil {
				if policies, err := expandGCPGroupScalingPolicies(v); err != nil {
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

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//             Schema
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func baseScalingPolicySchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				string(Cooldown): {
					Type:     schema.TypeInt,
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
								Required: true,
							},

							string(DimensionValue): {
								Type:     schema.TypeString,
								Optional: true,
							},
						},
					},
				},

				string(MetricName): {
					Type:     schema.TypeString,
					Required: true,
				},

				string(Namespace): {
					Type:     schema.TypeString,
					Required: true,
				},

				string(PolicyName): {
					Type:     schema.TypeString,
					Required: true,
				},

				string(Source): {
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},

				string(Statistic): {
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},

				string(Unit): {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}

func upDownScalingPolicySchema() *schema.Schema {
	o := baseScalingPolicySchema()
	s := o.Elem.(*schema.Resource).Schema

	s[string(ActionType)] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	s[string(Adjustment)] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}

	s[string(EvaluationPeriods)] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	}

	s[string(Operator)] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	}

	s[string(Period)] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	}

	s[string(Threshold)] = &schema.Schema{
		Type:     schema.TypeFloat,
		Required: true,
	}

	return o
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//             Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandGCPGroupScalingPolicies(data interface{}) ([]*gcp.ScalingPolicy, error) {
	list := data.(*schema.Set).List()
	policies := make([]*gcp.ScalingPolicy, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		policy := &gcp.ScalingPolicy{}

		if v, ok := m[string(ActionType)].(string); ok && v != "" {
			action := &gcp.Action{}
			action.SetType(spotinst.String(v))

			if v, ok := m[string(Adjustment)].(int); ok && v >= 0 {
				action.SetAdjustment(spotinst.Int(v))
			}

			policy.SetAction(action)
		}

		if v, ok := m[string(Cooldown)].(int); ok && v > 0 {
			policy.SetCooldown(spotinst.Int(v))
		}

		if v, ok := m[string(Dimensions)]; ok {
			dimensions := expandGCPGroupScalingPolicyDimensions(v.(interface{}))
			if len(dimensions) > 0 {
				policy.SetDimensions(dimensions)
			}
		}

		if v, ok := m[string(EvaluationPeriods)].(int); ok && v > 0 {
			policy.SetEvaluationPeriods(spotinst.Int(v))
		}

		if v, ok := m[string(MetricName)].(string); ok && v != "" {
			policy.SetMetricName(spotinst.String(v))
		}

		if v, ok := m[string(Namespace)].(string); ok && v != "" {
			policy.SetNamespace(spotinst.String(v))
		}

		if v, ok := m[string(Operator)].(string); ok && v != "" {
			policy.SetOperator(spotinst.String(v))
		}

		if v, ok := m[string(Period)].(int); ok && v > 0 {
			policy.SetPeriod(spotinst.Int(v))
		}

		if v, ok := m[string(Source)].(string); ok && v != "" {
			policy.SetSource(spotinst.String(v))
		}

		if v, ok := m[string(Statistic)].(string); ok && v != "" {
			policy.SetStatistic(spotinst.String(v))
		}

		if v, ok := m[string(Threshold)].(float64); ok && v > 0 {
			policy.SetThreshold(spotinst.Float64(v))
		}

		if v, ok := m[string(Unit)].(string); ok && v != "" {
			policy.SetUnit(spotinst.String(v))
		}

		if v, ok := m[string(PolicyName)].(string); ok && v != "" {
			policy.SetPolicyName(spotinst.String(v))
		}

		if policy.Namespace != nil {
			policies = append(policies, policy)
		}
	}

	return policies, nil
}

func expandGCPGroupScalingPolicyDimensions(data interface{}) []*gcp.Dimension {
	list := data.([]interface{})
	dimensions := make([]*gcp.Dimension, 0, len(list))
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
		dimension := &gcp.Dimension{
			Name:  spotinst.String(attr[string(DimensionName)].(string)),
			Value: spotinst.String(attr[string(DimensionValue)].(string)),
		}
		if (dimension.Name != nil) && (dimension.Value != nil) {
			dimensions = append(dimensions, dimension)
		}
	}
	return dimensions
}

func flattenGCPGroupScalingPolicy(policies []*gcp.ScalingPolicy) []interface{} {
	result := make([]interface{}, 0, len(policies))
	for _, policy := range policies {
		m := make(map[string]interface{})
		m[string(Cooldown)] = spotinst.IntValue(policy.Cooldown)
		m[string(MetricName)] = spotinst.StringValue(policy.MetricName)
		m[string(Namespace)] = spotinst.StringValue(policy.Namespace)
		m[string(PolicyName)] = spotinst.StringValue(policy.PolicyName)
		m[string(Source)] = spotinst.StringValue(policy.Source)
		m[string(Statistic)] = spotinst.StringValue(policy.Statistic)
		m[string(Unit)] = spotinst.StringValue(policy.Unit)

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

		if policy.Action != nil && policy.Action.Type != nil {
			m[string(ActionType)] = spotinst.StringValue(policy.Action.Type)
			m[string(Adjustment)] = spotinst.IntValue(policy.Action.Adjustment)
			m[string(EvaluationPeriods)] = spotinst.IntValue(policy.EvaluationPeriods)
			m[string(Operator)] = spotinst.StringValue(policy.Operator)
			m[string(Period)] = spotinst.IntValue(policy.Period)
			m[string(Threshold)] = spotinst.Float64Value(policy.Threshold)
		}

		result = append(result, m)
	}
	return result
}
