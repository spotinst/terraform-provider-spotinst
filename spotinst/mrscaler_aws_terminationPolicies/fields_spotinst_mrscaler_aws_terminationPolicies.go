package mrscaler_aws_terminationPolicies

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/mrscaler"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[TerminationPolicies] = commons.NewGenericField(
		commons.MRScalerAWSTerminationPolicies,
		TerminationPolicies,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(Statements): {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Namespace): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(MetricName): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(Statistic): {
									Type:     schema.TypeString,
									Optional: true,
									Default:  "sum",
								},

								string(Unit): {
									Type:     schema.TypeString,
									Optional: true,
									Default:  "count",
								},

								string(Threshold): {
									Type:     schema.TypeFloat,
									Required: true,
								},

								string(Period): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  300,
								},

								string(EvaluationPeriods): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  1,
								},

								string(Operator): {
									Type:     schema.TypeString,
									Optional: true,
									Default:  "gte",
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value []interface{} = nil
			if scaler.TerminationPolicies != nil {
				value = flattenTerminationPolicies(scaler.TerminationPolicies)
			}
			if err := resourceData.Set(string(TerminationPolicies), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(TerminationPolicies), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.Get(string(TerminationPolicies)).([]interface{}); ok {
				if terminationPolicies, err := expandTerminationPolicies(v); err != nil {
					return err
				} else {
					scaler.SetTerminationPolicies(terminationPolicies)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value []*mrscaler.TerminationPolicy = nil
			if v, ok := resourceData.Get(string(TerminationPolicies)).([]interface{}); ok {
				if terminationPolicies, err := expandTerminationPolicies(v); err != nil {
					return err
				} else if len(terminationPolicies) > 0 {
					value = terminationPolicies
				}
			}
			scaler.SetTerminationPolicies(value)
			return nil
		},
		nil,
	)

}

func expandTerminationPolicies(data interface{}) ([]*mrscaler.TerminationPolicy, error) {
	list := data.([]interface{})
	terminationPolicies := make([]*mrscaler.TerminationPolicy, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		terminationPolicie := &mrscaler.TerminationPolicy{}
		if v, ok := m[string(Statements)]; ok {
			statements := expandStatements(v.(interface{}))
			if len(statements) > 0 {
				terminationPolicie.SetStatements(statements)
			}
		}
		terminationPolicies = append(terminationPolicies, terminationPolicie)
	}

	return terminationPolicies, nil
}

func expandStatements(data interface{}) []*mrscaler.Statement {
	list := data.([]interface{})
	statements := make([]*mrscaler.Statement, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		statement := &mrscaler.Statement{}

		if v, ok := m[string(Namespace)].(string); ok && v != "" {
			statement.SetNamespace(spotinst.String(v))
		}

		if v, ok := m[string(MetricName)].(string); ok && v != "" {
			statement.SetMetricName(spotinst.String(v))
		}

		if v, ok := m[string(Statistic)].(string); ok && v != "" {
			statement.SetStatistic(spotinst.String(v))
		}

		if v, ok := m[string(Unit)].(string); ok && v != "" {
			statement.SetUnit(spotinst.String(v))
		}

		if v, ok := m[string(Threshold)].(float64); ok && v >= 0 {
			statement.SetThreshold(spotinst.Float64(v))
		}

		if v, ok := m[string(Period)].(int); ok && v > 0 {
			statement.SetPeriod(spotinst.Int(v))
		}

		if v, ok := m[string(EvaluationPeriods)].(int); ok && v > 0 {
			statement.SetEvaluationPeriods(spotinst.Int(v))
		}

		if v, ok := m[string(Operator)].(string); ok && v != "" {
			statement.SetOperator(spotinst.String(v))
		}

		statements = append(statements, statement)
	}
	return statements
}

func flattenTerminationPolicies(policies []*mrscaler.TerminationPolicy) []interface{} {
	result := make([]interface{}, 0, len(policies))
	for _, policy := range policies {
		m := make(map[string]interface{})

		if policy.Statements != nil {
			m[string(Statements)] = flattenStatements(policy.Statements)
		}

		result = append(result, m)
	}
	return result
}

func flattenStatements(statements []*mrscaler.Statement) []interface{} {
	result := make([]interface{}, 0, len(statements))
	for _, statement := range statements {
		m := make(map[string]interface{})

		m[string(Namespace)] = spotinst.StringValue(statement.Namespace)
		m[string(MetricName)] = spotinst.StringValue(statement.MetricName)
		m[string(Statistic)] = spotinst.StringValue(statement.Statistic)
		m[string(Unit)] = spotinst.StringValue(statement.Unit)
		m[string(Threshold)] = spotinst.Float64Value(statement.Threshold)
		m[string(Period)] = spotinst.IntValue(statement.Period)
		m[string(EvaluationPeriods)] = spotinst.IntValue(statement.EvaluationPeriods)
		m[string(Operator)] = spotinst.StringValue(statement.Operator)

		result = append(result, m)
	}
	return result
}
