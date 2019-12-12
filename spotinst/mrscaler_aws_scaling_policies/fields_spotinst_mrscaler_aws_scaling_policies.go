package mrscaler_aws_scaling_policies

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/mrscaler"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	SetupTaskScalingPolicies(fieldsMap)
	SetupCoreScalingPolicies(fieldsMap)
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
					Type:     schema.TypeMap,
					Optional: true,
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
		Type:     schema.TypeString,
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

	s[string(MinTargetCapacity)] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	s[string(MaxTargetCapacity)] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	s[string(Minimum)] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	s[string(Maximum)] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	s[string(Target)] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	return o
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//             Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandMRScalerAWSScalingPolicies(data interface{}) ([]*mrscaler.ScalingPolicy, error) {
	list := data.(*schema.Set).List()
	policies := make([]*mrscaler.ScalingPolicy, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		policy := &mrscaler.ScalingPolicy{}

		if v, ok := m[string(Cooldown)].(int); ok && v > 0 {
			policy.SetCooldown(spotinst.Int(v))
		}

		if v, ok := m[string(Dimensions)]; ok {
			dimensions := expandMRScalerAWSScalingPolicyDimensions(v.(map[string]interface{}))
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

		if v, ok := m[string(ActionType)].(string); ok && v != "" {
			action := &mrscaler.Action{}
			action.SetType(spotinst.String(v))

			if v, ok := m[string(Adjustment)].(string); ok && v != "" {
				action.SetAdjustment(spotinst.String(v))
			}

			if v, ok := m[string(MinTargetCapacity)].(string); ok && v != "" {
				action.SetMinTargetCapacity(spotinst.String(v))
			}

			if v, ok := m[string(MaxTargetCapacity)].(string); ok && v != "" {
				action.SetMaxTargetCapacity(spotinst.String(v))
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

			policy.SetAction(action)
		}

		if policy.Namespace != nil {
			policies = append(policies, policy)
		}
	}

	return policies, nil
}

func expandMRScalerAWSScalingPolicyDimensions(list map[string]interface{}) []*mrscaler.Dimension {
	dimensions := make([]*mrscaler.Dimension, 0, len(list))
	for name, val := range list {
		dimension := &mrscaler.Dimension{}
		dimension.SetName(spotinst.String(name))
		dimension.SetValue(spotinst.String(val.(string)))
		dimensions = append(dimensions, dimension)
	}
	return dimensions
}

func flattenMRScalerAWSScalingPolicy(policies []*mrscaler.ScalingPolicy) []interface{} {
	result := make([]interface{}, 0, len(policies))
	for _, policy := range policies {
		m := make(map[string]interface{})
		m[string(Cooldown)] = spotinst.IntValue(policy.Cooldown)
		m[string(MetricName)] = spotinst.StringValue(policy.MetricName)
		m[string(Namespace)] = spotinst.StringValue(policy.Namespace)
		m[string(PolicyName)] = spotinst.StringValue(policy.PolicyName)
		m[string(Statistic)] = spotinst.StringValue(policy.Statistic)
		m[string(Unit)] = spotinst.StringValue(policy.Unit)

		if policy.Dimensions != nil && len(policy.Dimensions) > 0 {
			dimMap := make(map[string]interface{})
			for _, dimension := range policy.Dimensions {
				dimMap[spotinst.StringValue(dimension.Name)] = spotinst.StringValue(dimension.Value)
			}
			m[string(Dimensions)] = dimMap
		}

		if policy.Action != nil && policy.Action.Type != nil {
			m[string(ActionType)] = spotinst.StringValue(policy.Action.Type)
			m[string(Adjustment)] = spotinst.StringValue(policy.Action.Adjustment)
			m[string(EvaluationPeriods)] = spotinst.IntValue(policy.EvaluationPeriods)
			m[string(Maximum)] = spotinst.StringValue(policy.Action.Maximum)
			m[string(MaxTargetCapacity)] = spotinst.StringValue(policy.Action.MaxTargetCapacity)
			m[string(Minimum)] = spotinst.StringValue(policy.Action.Minimum)
			m[string(MinTargetCapacity)] = spotinst.StringValue(policy.Action.MinTargetCapacity)
			m[string(Operator)] = spotinst.StringValue(policy.Operator)
			m[string(Period)] = spotinst.IntValue(policy.Period)
			m[string(Target)] = spotinst.StringValue(policy.Action.Target)
			m[string(Threshold)] = spotinst.Float64Value(policy.Threshold)
		}

		result = append(result, m)
	}
	return result
}
