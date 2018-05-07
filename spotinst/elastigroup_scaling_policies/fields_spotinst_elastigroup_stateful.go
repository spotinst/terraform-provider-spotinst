package elastigroup_scaling_policies

import (
	"strconv"

	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ScalingUpPolicy] = commons.NewGenericField(
		commons.ElastigroupScalingPolicies,
		ScalingUpPolicy,
		upDownScalingPolicySchema(),
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			// Skip
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.GetOk(string(ScalingUpPolicy)); ok {
				if policies, err := expandAWSGroupScalingPolicies(v); err != nil {
					return err
				} else {
					elastigroup.Scaling.SetUp(policies)
				}
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value []*aws.ScalingPolicy = nil
			if v, ok := resourceData.GetOk(string(ScalingUpPolicy)); ok {
				if policies, err := expandAWSGroupScalingPolicies(v); err != nil {
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
		commons.ElastigroupScalingPolicies,
		ScalingDownPolicy,
		upDownScalingPolicySchema(),
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			// Skip
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.GetOk(string(ScalingDownPolicy)); ok {
				if policies, err := expandAWSGroupScalingPolicies(v); err != nil {
					return err
				} else {
					elastigroup.Scaling.SetDown(policies)
				}
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value []*aws.ScalingPolicy = nil
			if v, ok := resourceData.GetOk(string(ScalingDownPolicy)); ok {
				if policies, err := expandAWSGroupScalingPolicies(v); err != nil {
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

	fieldsMap[ScalingTargetPolicy] = commons.NewGenericField(
		commons.ElastigroupScalingPolicies,
		ScalingTargetPolicy,
		targetScalingPolicySchema(),
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			// Skip
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.GetOk(string(ScalingTargetPolicy)); ok {
				if policies, err := expandAWSGroupScalingPolicies(v); err != nil {
					return err
				} else {
					elastigroup.Scaling.SetTarget(policies)
				}
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value []*aws.ScalingPolicy = nil
			if v, ok := resourceData.GetOk(string(ScalingTargetPolicy)); ok {
				if policies, err := expandAWSGroupScalingPolicies(v); err != nil {
					return err
				} else {
					value = policies
				}
			}
			elastigroup.Scaling.SetTarget(value)
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
				string(PolicyName): &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},

				string(MetricName): &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},

				string(Namespace): &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},

				string(Source): &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},

				string(Statistic): &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},

				string(Unit): &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},

				string(Cooldown): &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Computed: true,
				},

				string(Dimensions): &schema.Schema{
					Type:     schema.TypeMap,
					Optional: true,
				},
			},
		},
	}
}

func upDownScalingPolicySchema() *schema.Schema {
	o := baseScalingPolicySchema()
	s := o.Elem.(*schema.Resource).Schema

	s[string(Threshold)] = &schema.Schema{
		Type:     schema.TypeFloat,
		Required: true,
	}

	s[string(Adjustment)] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}

	s[string(AdjustmentExpression)] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	s[string(MinTargetCapacity)] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}

	s[string(MaxTargetCapacity)] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}

	s[string(Operator)] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	}

	s[string(EvaluationPeriods)] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	}

	s[string(Period)] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	}

	s[string(Minimum)] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}

	s[string(Maximum)] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}

	s[string(Target)] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}

	s[string(ActionType)] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	return o
}

func targetScalingPolicySchema() *schema.Schema {
	o := baseScalingPolicySchema()
	s := o.Elem.(*schema.Resource).Schema

	s[string(Target)] = &schema.Schema{
		Type:     schema.TypeFloat,
		Required: true,
	}

	return o
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//             Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupScalingPolicies(data interface{}) ([]*aws.ScalingPolicy, error) {
	list := data.(*schema.Set).List()
	policies := make([]*aws.ScalingPolicy, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		policy := &aws.ScalingPolicy{}

		if v, ok := m[string(PolicyName)].(string); ok && v != "" {
			policy.SetPolicyName(spotinst.String(v))
		}

		if v, ok := m[string(MetricName)].(string); ok && v != "" {
			policy.SetMetricName(spotinst.String(v))
		}

		if v, ok := m[string(Namespace)].(string); ok && v != "" {
			policy.SetNamespace(spotinst.String(v))
		}

		if v, ok := m[string(Source)].(string); ok && v != "" {
			policy.SetSource(spotinst.String(v))
		}

		if v, ok := m[string(Statistic)].(string); ok && v != "" {
			policy.SetStatistic(spotinst.String(v))
		}

		if v, ok := m[string(Unit)].(string); ok && v != "" {
			policy.SetUnit(spotinst.String(v))
		}

		if v, ok := m[string(Threshold)].(float64); ok && v > 0 {
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
			dimensions := expandAWSGroupScalingPolicyDimensions(v.(map[string]interface{}))
			if len(dimensions) > 0 {
				policy.SetDimensions(dimensions)
			}
		}

		if v, ok := m[string(ActionType)].(string); ok && v != "" {
			action := &aws.Action{}
			action.SetType(spotinst.String(v))

			if v, ok := m[string(Adjustment)].(int); ok && v > 0 {
				action.SetAdjustment(spotinst.String(strconv.Itoa(v)))
			} else if v, ok := m[string(AdjustmentExpression)].(string); ok && v != "" {
				action.SetAdjustment(spotinst.String(v))
			}

			if v, ok := m[string(MinTargetCapacity)].(int); ok && v > 0 {
				action.SetMinTargetCapacity(spotinst.Int(v))
			}

			if v, ok := m[string(MaxTargetCapacity)].(int); ok && v > 0 {
				action.SetMaxTargetCapacity(spotinst.Int(v))
			}

			if v, ok := m[string(Minimum)].(int); ok && v > 0 {
				action.SetMinimum(spotinst.Int(v))
			}

			if v, ok := m[string(Maximum)].(int); ok && v > 0 {
				action.SetMaximum(spotinst.Int(v))
			}

			if v, ok := m[string(Target)].(int); ok && v > 0 {
				action.SetTarget(spotinst.Int(v))
			}

			policy.SetAction(action)
		} else {
			if v, ok := m[string(Adjustment)].(int); ok && v > 0 {
				policy.SetAdjustment(spotinst.Int(v))
			}

			if v, ok := m[string(MinTargetCapacity)].(int); ok && v > 0 {
				policy.SetMinTargetCapacity(spotinst.Int(v))
			}

			if v, ok := m[string(MaxTargetCapacity)].(int); ok && v > 0 {
				policy.SetMaxTargetCapacity(spotinst.Int(v))
			}
		}

		// Target scaling policy?
		if policy.Threshold == nil {
			if v, ok := m[string(Target)].(float64); ok && v >= 0 {
				policy.SetTarget(spotinst.Float64(v))
			}
		}

		if policy.Namespace != nil {
			policies = append(policies, policy)
		}
	}

	return policies, nil
}

func expandAWSGroupScalingPolicyDimensions(list map[string]interface{}) []*aws.Dimension {
	dimensions := make([]*aws.Dimension, 0, len(list))
	for name, val := range list {
		dimension := &aws.Dimension{}
		dimension.SetName(spotinst.String(name))
		dimension.SetValue(spotinst.String(val.(string)))
		dimensions = append(dimensions, dimension)
	}
	return dimensions
}