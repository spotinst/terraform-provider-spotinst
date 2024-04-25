package ocean_aws_right_sizing_rule

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/service/organization"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[RecommendationApplicationIntervals] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		RecommendationApplicationIntervals,
		&schema.Schema{
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(RepetitionBasis): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(WeeklyRepetitionBasis): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(IntervalDays): {
									Type:     schema.TypeList,
									Required: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								string(IntervalHoursStartTime): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(IntervalHoursEndTime): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},

					string(MonthlyRepetitionBasis): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(IntervalMonths): {
									Type:     schema.TypeList,
									Required: true,
									Elem:     &schema.Schema{Type: schema.TypeInt},
								},
								string(WeekOfTheMonth): {
									Type:     schema.TypeList,
									Required: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								string(MonthlyWeeklyRepetitionBasis): {
									Type:     schema.TypeSet,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(MonthlyWeeklyIntervalDays): {
												Type:     schema.TypeList,
												Required: true,
												Elem:     &schema.Schema{Type: schema.TypeString},
											},
											string(MonthlyWeeklyIntervalHoursStartTime): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(MonthlyWeeklyIntervalHoursEndTime): {
												Type:     schema.TypeString,
												Required: true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanAWSRightSizingRule()
			var result []interface{} = nil
			if rightSizingRule.RecommendationApplicationIntervals != nil {
				recommendationApplicationIntervals := rightSizingRule.RecommendationApplicationIntervals
				result = flattenRecommendationApplicationIntervals(recommendationApplicationIntervals)
			}
			if result != nil {
				if err := resourceData.Set(string(RecommendationApplicationIntervals), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RecommendationApplicationIntervals), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(NetworkInterfaces)); ok {
				if networks, err := expandLaunchSpecNetworkInterfaces(v); err != nil {
					return err
				} else {
					launchSpec.SetLaunchSpecNetworkInterfaces(networks)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value []*gcp.LaunchSpecNetworkInterfaces = nil
			if v, ok := resourceData.GetOk(string(NetworkInterfaces)); ok {
				if networks, err := expandLaunchSpecNetworkInterfaces(v); err != nil {
					return err
				} else {
					value = networks
				}
			}
			launchSpec.SetLaunchSpecNetworkInterfaces(value)
			return nil
		},
		nil,
	)

}

func flattenRecommendationApplicationIntervals(policyContent *organization.PolicyContent) []interface{} {
	result := make(map[string]interface{})
	result[string(Statements)] = flattenStatements(policyContent.Statements)
	return []interface{}{result}
}

func flattenStatements(statements []*organization.Statement) []interface{} {
	result := make([]interface{}, 0, len(statements))

	for _, statement := range statements {
		m := make(map[string]interface{})
		if statement.Actions != nil {
			m[string(Actions)] = statement.Actions
		}

		m[string(Effect)] = spotinst.StringValue(statement.Effect)

		if statement.Resources != nil {
			m[string(Resources)] = statement.Resources
		}
		result = append(result, m)
	}

	return result
}
