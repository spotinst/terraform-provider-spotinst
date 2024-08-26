package ocean_right_sizing_rule

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/right_sizing"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[RuleName] = commons.NewGenericField(
		commons.OceanRightSizingRule,
		RuleName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			var value *string = nil
			if rightSizingRule.RuleName != nil {
				value = rightSizingRule.RuleName
			}
			if err := resourceData.Set(string(RuleName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RuleName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			if v, ok := resourceData.GetOk(string(RuleName)); ok && v != "" {
				rightSizingRule.SetRuleName(spotinst.String(resourceData.Get(string(RuleName)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			if v, ok := resourceData.GetOk(string(RuleName)); ok && v != "" {
				rightSizingRule.SetRuleName(spotinst.String(resourceData.Get(string(RuleName)).(string)))
			}
			return nil
		},
		nil,
	)
	fieldsMap[RestartReplicas] = commons.NewGenericField(
		commons.OceanRightSizingRule,
		RestartReplicas,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			var value *string = nil
			if rightSizingRule.RestartReplicas != nil {
				value = rightSizingRule.RestartReplicas
			}
			if err := resourceData.Set(string(RestartReplicas), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RestartReplicas), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			if v, ok := resourceData.GetOk(string(RestartReplicas)); ok && v != "" {
				rightSizingRule.SetRestartReplicas(spotinst.String(resourceData.Get(string(RestartReplicas)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			if v, ok := resourceData.GetOk(string(RestartReplicas)); ok && v != "" {
				rightSizingRule.SetRuleName(spotinst.String(resourceData.Get(string(RestartReplicas)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[ExcludePreliminaryRecommendations] = commons.NewGenericField(
		commons.OceanRightSizingRule,
		ExcludePreliminaryRecommendations,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			var value *bool = nil
			if rightSizingRule.ExcludePreliminaryRecommendations != nil {
				value = rightSizingRule.ExcludePreliminaryRecommendations
			}
			if err := resourceData.Set(string(ExcludePreliminaryRecommendations), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ExcludePreliminaryRecommendations), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			if v, ok := resourceData.GetOk(string(ExcludePreliminaryRecommendations)); ok && v != "" {
				rightSizingRule.SetExcludePreliminaryRecommendations(spotinst.Bool(resourceData.Get(string(ExcludePreliminaryRecommendations)).(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			if v, ok := resourceData.GetOk(string(ExcludePreliminaryRecommendations)); ok && v != "" {
				rightSizingRule.SetExcludePreliminaryRecommendations(spotinst.Bool(resourceData.Get(string(ExcludePreliminaryRecommendations)).(bool)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[OceanId] = commons.NewGenericField(
		commons.OrganizationPolicy,
		OceanId,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			if v, ok := resourceData.GetOk(string(OceanId)); ok && v != "" {
				rightSizingRule.SetOceanId(spotinst.String(resourceData.Get(string(OceanId)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[RecommendationApplicationIntervals] = commons.NewGenericField(
		commons.OceanRightSizingRule,
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
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
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
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			if v, ok := resourceData.GetOk(string(RecommendationApplicationIntervals)); ok {
				if recommendationApplicationIntervals, err := expandRecommendationApplicationIntervals(v); err != nil {
					return err
				} else {
					rightSizingRule.SetRecommendationApplicationIntervals(recommendationApplicationIntervals)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			var value []*right_sizing.RecommendationApplicationIntervals = nil
			if v, ok := resourceData.GetOk(string(RecommendationApplicationIntervals)); ok {
				if recommendationApplicationIntervals, err := expandRecommendationApplicationIntervals(v); err != nil {
					return err
				} else {
					value = recommendationApplicationIntervals
				}
			}
			rightSizingRule.SetRecommendationApplicationIntervals(value)
			return nil
		},
		nil,
	)

	fieldsMap[RecommendationApplicationMinThreshold] = commons.NewGenericField(
		commons.OceanRightSizingRule,
		RecommendationApplicationMinThreshold,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(CpuPercentage): {
						Type:     schema.TypeFloat,
						Optional: true,
						Default:  -1,
					},
					string(MemoryPercentage): {
						Type:     schema.TypeFloat,
						Optional: true,
						Default:  -1,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			var result []interface{} = nil
			if rightSizingRule.RecommendationApplicationMinThreshold != nil {
				recommendationApplicationMinThreshold := rightSizingRule.RecommendationApplicationMinThreshold
				result = flattenRecommendationApplicationMinThreshold(recommendationApplicationMinThreshold)
			}
			if result != nil {
				if err := resourceData.Set(string(RecommendationApplicationMinThreshold), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RecommendationApplicationMinThreshold), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			if v, ok := resourceData.GetOk(string(RecommendationApplicationMinThreshold)); ok {
				if recommendationApplicationMinThreshold, err := expandRecommendationApplicationMinThreshold(v); err != nil {
					return err
				} else {
					rightSizingRule.SetRecommendationApplicationMinThreshold(recommendationApplicationMinThreshold)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			var value *right_sizing.RecommendationApplicationMinThreshold = nil
			if v, ok := resourceData.GetOk(string(RecommendationApplicationMinThreshold)); ok {
				if recommendationApplicationMinThreshold, err := expandRecommendationApplicationMinThreshold(v); err != nil {
					return err
				} else {
					value = recommendationApplicationMinThreshold
				}
			}
			rightSizingRule.SetRecommendationApplicationMinThreshold(value)
			return nil
		},
		nil,
	)

	fieldsMap[RecommendationApplicationOverheadValues] = commons.NewGenericField(
		commons.OceanRightSizingRule,
		RecommendationApplicationOverheadValues,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(OverheadCpuPercentage): {
						Type:     schema.TypeFloat,
						Optional: true,
						Default:  -1,
					},
					string(OverheadMemoryPercentage): {
						Type:     schema.TypeFloat,
						Optional: true,
						Default:  -1,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			var result []interface{} = nil
			if rightSizingRule.RecommendationApplicationOverheadValues != nil {
				recommendationApplicationOverheadValues := rightSizingRule.RecommendationApplicationOverheadValues
				result = flattenRecommendationApplicationOverheadValues(recommendationApplicationOverheadValues)
			}
			if result != nil {
				if err := resourceData.Set(string(RecommendationApplicationOverheadValues), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RecommendationApplicationOverheadValues), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			if v, ok := resourceData.GetOk(string(RecommendationApplicationOverheadValues)); ok {
				if recommendationApplicationOverheadValues, err := expandRecommendationApplicationOverheadValues(v); err != nil {
					return err
				} else {
					rightSizingRule.SetRecommendationApplicationOverheadValues(recommendationApplicationOverheadValues)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			var value *right_sizing.RecommendationApplicationOverheadValues = nil
			if v, ok := resourceData.GetOk(string(RecommendationApplicationOverheadValues)); ok {
				if recommendationApplicationOverheadValues, err := expandRecommendationApplicationOverheadValues(v); err != nil {
					return err
				} else {
					value = recommendationApplicationOverheadValues
				}
			}
			rightSizingRule.SetRecommendationApplicationOverheadValues(value)
			return nil
		},
		nil,
	)

	fieldsMap[RecommendationApplicationBoundaries] = commons.NewGenericField(
		commons.OceanRightSizingRule,
		RecommendationApplicationBoundaries,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(CpuMin): {
						Type:     schema.TypeFloat,
						Optional: true,
						Default:  -1,
					},
					string(CpuMax): {
						Type:     schema.TypeFloat,
						Optional: true,
						Default:  -1,
					},
					string(MemoryMin): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
					string(MemoryMax): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			var result []interface{} = nil
			if rightSizingRule.RecommendationApplicationBoundaries != nil {
				recommendationApplicationBoundaries := rightSizingRule.RecommendationApplicationBoundaries
				result = flattenRecommendationApplicationBoundaries(recommendationApplicationBoundaries)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(RecommendationApplicationBoundaries), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RecommendationApplicationBoundaries), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			if v, ok := resourceData.GetOk(string(RecommendationApplicationBoundaries)); ok {
				if recommendationApplicationBoundaries, err := expandRecommendationApplicationBoundaries(v); err != nil {
					return err
				} else {
					rightSizingRule.SetRecommendationApplicationBoundaries(recommendationApplicationBoundaries)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			var value *right_sizing.RecommendationApplicationBoundaries = nil
			if v, ok := resourceData.GetOk(string(RecommendationApplicationBoundaries)); ok {
				if recommendationApplicationBoundaries, err := expandRecommendationApplicationBoundaries(v); err != nil {
					return err
				} else {
					value = recommendationApplicationBoundaries
				}
			}
			rightSizingRule.SetRecommendationApplicationBoundaries(value)
			return nil
		},
		nil,
	)

	fieldsMap[AttachRightSizingRule] = commons.NewGenericField(
		commons.OceanRightSizingRule,
		AttachRightSizingRule,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Namespaces): {
						Type:     schema.TypeSet,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(NamespaceName): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(Workloads): {
									Type:     schema.TypeSet,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(WorkloadName): {
												Type:     schema.TypeString,
												Optional: true,
											},
											string(WorkloadType): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(RegexName): {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},
								string(Labels): {
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
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[DetachRightSizingRule] = commons.NewGenericField(
		commons.OceanRightSizingRule,
		DetachRightSizingRule,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(NamespacesForDetach): {
						Type:     schema.TypeSet,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(NamespaceNameForDetach): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(WorkloadsForDetach): {
									Type:     schema.TypeSet,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(WorkloadNameForDetach): {
												Type:     schema.TypeString,
												Optional: true,
											},
											string(WorkloadTypeForDetach): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(RegexNameForDetach): {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},
								string(LabelsForDetach): {
									Type:     schema.TypeSet,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(KeyForDetach): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(ValueForDetach): {
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
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[RecommendationApplicationHPA] = commons.NewGenericField(
		commons.OceanRightSizingRule,
		RecommendationApplicationHPA,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(AllowHPARecommendation): {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			var result []interface{} = nil
			if rightSizingRule.RecommendationApplicationHPA != nil {
				recommendationApplicationHPA := rightSizingRule.RecommendationApplicationHPA
				result = flattenRecommendationApplicationHPA(recommendationApplicationHPA)
			}
			if result != nil {
				if err := resourceData.Set(string(RecommendationApplicationOverheadValues), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RecommendationApplicationOverheadValues), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			if v, ok := resourceData.GetOk(string(RecommendationApplicationHPA)); ok {
				if recommendationApplicationHPA, err := expandRecommendationApplicationHPA(v); err != nil {
					return err
				} else {
					rightSizingRule.SetRecommendationApplicationHPA(recommendationApplicationHPA)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanRightSizingRule()
			var value *right_sizing.RecommendationApplicationHPA = nil
			if v, ok := resourceData.GetOk(string(RecommendationApplicationHPA)); ok {
				if recommendationApplicationHPA, err := expandRecommendationApplicationHPA(v); err != nil {
					return err
				} else {
					value = recommendationApplicationHPA
				}
			}
			rightSizingRule.SetRecommendationApplicationHPA(value)
			return nil
		},
		nil,
	)

}

func flattenRecommendationApplicationIntervals(recommendationApplicationIntervals []*right_sizing.RecommendationApplicationIntervals) []interface{} {
	result := make([]interface{}, 0, len(recommendationApplicationIntervals))

	for _, recommendationApplicationInterval := range recommendationApplicationIntervals {
		m := make(map[string]interface{})

		m[string(RepetitionBasis)] = spotinst.StringValue(recommendationApplicationInterval.RepetitionBasis)

		if recommendationApplicationInterval.WeeklyRepetitionBasis != nil {
			m[string(WeeklyRepetitionBasis)] = flattenWeeklyRepetitionBasis(recommendationApplicationInterval.WeeklyRepetitionBasis)
		}

		if recommendationApplicationInterval.MonthlyRepetitionBasis != nil {
			m[string(MonthlyRepetitionBasis)] = flattenMonthlyRepetitionBasis(recommendationApplicationInterval.MonthlyRepetitionBasis)
		}

		result = append(result, m)
	}

	return result
}

func flattenWeeklyRepetitionBasis(weeklyRepetitionBasis *right_sizing.WeeklyRepetitionBasis) []interface{} {
	result := make(map[string]interface{})
	if weeklyRepetitionBasis.IntervalDays != nil {
		result[string(IntervalDays)] = weeklyRepetitionBasis.IntervalDays
	}
	if weeklyRepetitionBasis.IntervalHours.StartTime != nil {
		result[string(IntervalHoursStartTime)] = spotinst.StringValue(weeklyRepetitionBasis.IntervalHours.StartTime)
	}
	if weeklyRepetitionBasis.IntervalHours.EndTime != nil {
		result[string(IntervalHoursEndTime)] = spotinst.StringValue(weeklyRepetitionBasis.IntervalHours.EndTime)
	}
	return []interface{}{result}
}

func flattenMonthlyRepetitionBasis(monthlyRepetitionBasis *right_sizing.MonthlyRepetitionBasis) []interface{} {
	result := make(map[string]interface{})
	if monthlyRepetitionBasis.IntervalMonths != nil {
		result[string(IntervalMonths)] = monthlyRepetitionBasis.IntervalMonths
	}
	if monthlyRepetitionBasis.WeekOfTheMonth != nil {
		result[string(WeekOfTheMonth)] = monthlyRepetitionBasis.WeekOfTheMonth
	}
	if monthlyRepetitionBasis.WeeklyRepetitionBasis != nil {
		result[string(MonthlyWeeklyRepetitionBasis)] = flattenMonthlyWeeklyRepetitionBasis(monthlyRepetitionBasis.WeeklyRepetitionBasis)
	}
	return []interface{}{result}
}

func flattenMonthlyWeeklyRepetitionBasis(weeklyRepetitionBasis *right_sizing.WeeklyRepetitionBasis) []interface{} {
	result := make(map[string]interface{})
	if weeklyRepetitionBasis.IntervalDays != nil {
		result[string(MonthlyWeeklyIntervalDays)] = weeklyRepetitionBasis.IntervalDays
	}
	if weeklyRepetitionBasis.IntervalHours.StartTime != nil {
		result[string(MonthlyWeeklyIntervalHoursStartTime)] = spotinst.StringValue(weeklyRepetitionBasis.IntervalHours.StartTime)
	}
	if weeklyRepetitionBasis.IntervalHours.EndTime != nil {
		result[string(MonthlyWeeklyIntervalHoursEndTime)] = spotinst.StringValue(weeklyRepetitionBasis.IntervalHours.EndTime)
	}
	return []interface{}{result}
}

// expandRecommendationApplicationIntervals sets the values from the plan as objects
func expandRecommendationApplicationIntervals(data interface{}) ([]*right_sizing.RecommendationApplicationIntervals, error) {
	list := data.(*schema.Set).List()

	if list != nil && list[0] != nil {
		ifaces := make([]*right_sizing.RecommendationApplicationIntervals, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			iface := &right_sizing.RecommendationApplicationIntervals{}

			if v, ok := m[string(RepetitionBasis)].(string); ok && v != "" {
				iface.SetRepetitionBasis(spotinst.String(v))
			}

			if v, ok := m[string(WeeklyRepetitionBasis)]; ok {
				weeklyRepetitionBasis, err := expandWeeklyRepetitionBasis(v)
				if err != nil {
					return nil, err
				}

				if weeklyRepetitionBasis != nil {
					iface.SetWeeklyRepetitionBasis(weeklyRepetitionBasis)
				}
			} else {
				iface.WeeklyRepetitionBasis = nil
			}

			if v, ok := m[string(MonthlyRepetitionBasis)]; ok {
				monthlyRepetitionBasis, err := expandMonthlyRepetitionBasis(v)
				if err != nil {
					return nil, err
				}

				if monthlyRepetitionBasis != nil {
					iface.SetMonthlyRepetitionBasis(monthlyRepetitionBasis)
				}
			} else {
				iface.MonthlyRepetitionBasis = nil
			}

			ifaces = append(ifaces, iface)
		}
		return ifaces, nil
	}
	return nil, nil
}

func expandWeeklyRepetitionBasis(data interface{}) (*right_sizing.WeeklyRepetitionBasis, error) {
	list := data.(*schema.Set).List()
	weeklyRepetitionBasis := &right_sizing.WeeklyRepetitionBasis{}
	intervalHours := &right_sizing.IntervalHours{}

	if len(list) > 0 {
		item := list[0]
		m := item.(map[string]interface{})

		if v, ok := m[string(IntervalHoursStartTime)].(string); ok && v != "" {
			intervalHours.SetStartTime(spotinst.String(v))
		} else {
			intervalHours.SetStartTime(nil)
		}

		if v, ok := m[string(IntervalHoursEndTime)].(string); ok && v != "" {
			intervalHours.SetEndTime(spotinst.String(v))
		} else {
			intervalHours.SetEndTime(nil)
		}

		weeklyRepetitionBasis.SetIntervalHours(intervalHours)

		if v, ok := m[string(IntervalDays)]; ok {
			intervalDays, err := expandIntervalDaysList(v)
			if err != nil {
				return nil, err
			}
			if intervalDays != nil && len(intervalDays) > 0 {
				weeklyRepetitionBasis.SetIntervalDays(intervalDays)
			} else {
				weeklyRepetitionBasis.SetIntervalDays(nil)
			}
		}

		return weeklyRepetitionBasis, nil

	}
	return nil, nil
}

func expandMonthlyRepetitionBasis(data interface{}) (*right_sizing.MonthlyRepetitionBasis, error) {
	list := data.(*schema.Set).List()
	monthlyRepetitionBasis := &right_sizing.MonthlyRepetitionBasis{}

	if len(list) > 0 {
		item := list[0]
		m := item.(map[string]interface{})

		if v, ok := m[string(WeekOfTheMonth)]; ok {
			weekOfTheMonth, err := expandIntervalDaysList(v)
			if err != nil {
				return nil, err
			}
			if weekOfTheMonth != nil && len(weekOfTheMonth) > 0 {
				monthlyRepetitionBasis.SetWeekOfTheMonth(weekOfTheMonth)
			} else {
				monthlyRepetitionBasis.SetWeekOfTheMonth(nil)
			}
		}

		if v, ok := m[string(IntervalMonths)]; ok {
			intervalMonths, err := expandIntervalMonthsList(v)
			if err != nil {
				return nil, err
			}
			if intervalMonths != nil && len(intervalMonths) > 0 {
				monthlyRepetitionBasis.SetIntervalMonths(intervalMonths)
			} else {
				monthlyRepetitionBasis.SetIntervalMonths(intervalMonths)
			}
		}

		if v, ok := m[string(MonthlyWeeklyRepetitionBasis)]; ok {
			monthlyWeeklyRepetitionBasis, err := expandMonthlyWeeklyRepetitionBasis(v)
			if err != nil {
				return nil, err
			}
			if monthlyWeeklyRepetitionBasis != nil {
				monthlyRepetitionBasis.SetMonthlyWeeklyRepetitionBasis(monthlyWeeklyRepetitionBasis)
			}
		} else {
			monthlyRepetitionBasis.SetMonthlyWeeklyRepetitionBasis(nil)
		}

		return monthlyRepetitionBasis, nil

	}
	return nil, nil
}

func expandMonthlyWeeklyRepetitionBasis(data interface{}) (*right_sizing.WeeklyRepetitionBasis, error) {
	list := data.(*schema.Set).List()
	weeklyRepetitionBasis := &right_sizing.WeeklyRepetitionBasis{}
	intervalHours := &right_sizing.IntervalHours{}

	if len(list) > 0 {
		item := list[0]
		m := item.(map[string]interface{})

		if v, ok := m[string(MonthlyWeeklyIntervalHoursStartTime)].(string); ok && v != "" {
			intervalHours.SetStartTime(spotinst.String(v))
		} else {
			intervalHours.SetStartTime(nil)
		}

		if v, ok := m[string(MonthlyWeeklyIntervalHoursEndTime)].(string); ok && v != "" {
			intervalHours.SetEndTime(spotinst.String(v))
		} else {
			intervalHours.SetEndTime(nil)
		}

		weeklyRepetitionBasis.SetIntervalHours(intervalHours)

		if v, ok := m[string(MonthlyWeeklyIntervalDays)]; ok {
			intervalDays, err := expandIntervalDaysList(v)
			if err != nil {
				return nil, err
			}
			if intervalDays != nil && len(intervalDays) > 0 {
				weeklyRepetitionBasis.SetIntervalDays(intervalDays)
			} else {
				weeklyRepetitionBasis.SetIntervalDays(nil)
			}
		}

		return weeklyRepetitionBasis, nil

	}
	return nil, nil
}

func expandIntervalDaysList(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if intervalDays, ok := v.(string); ok && intervalDays != "" {
			result = append(result, intervalDays)
		}
	}
	return result, nil
}

func expandIntervalMonthsList(data interface{}) ([]int, error) {
	list := data.([]interface{})
	result := make([]int, 0, len(list))

	for _, v := range list {
		if intervalMonths, ok := v.(int); ok && intervalMonths != 0 {
			result = append(result, intervalMonths)
		}
	}
	return result, nil
}

func flattenRecommendationApplicationMinThreshold(recommendationApplicationMinThreshold *right_sizing.RecommendationApplicationMinThreshold) []interface{} {
	result := make(map[string]interface{})

	if recommendationApplicationMinThreshold.CpuPercentage != nil {
		result[string(CpuPercentage)] = spotinst.Float64Value(recommendationApplicationMinThreshold.CpuPercentage)
	}
	if recommendationApplicationMinThreshold.MemoryPercentage != nil {
		result[string(MemoryPercentage)] = spotinst.Float64Value(recommendationApplicationMinThreshold.MemoryPercentage)
	}
	return []interface{}{result}
}

func expandRecommendationApplicationMinThreshold(data interface{}) (*right_sizing.RecommendationApplicationMinThreshold, error) {
	list := data.(*schema.Set).List()
	recommendationApplicationMinThreshold := &right_sizing.RecommendationApplicationMinThreshold{}

	if len(list) > 0 {
		item := list[0]
		m := item.(map[string]interface{})

		if v, ok := m[string(CpuPercentage)].(float64); ok {
			if v == -1 {
				recommendationApplicationMinThreshold.SetCpuPercentage(nil)
			} else {
				recommendationApplicationMinThreshold.SetCpuPercentage(spotinst.Float64(v))
			}
		}

		if v, ok := m[string(MemoryPercentage)].(float64); ok {
			if v == -1 {
				recommendationApplicationMinThreshold.SetMemoryPercentage(nil)
			} else {
				recommendationApplicationMinThreshold.SetMemoryPercentage(spotinst.Float64(v))
			}
		}

		return recommendationApplicationMinThreshold, nil

	}
	return nil, nil
}

func flattenRecommendationApplicationBoundaries(recommendationApplicationBoundaries *right_sizing.RecommendationApplicationBoundaries) []interface{} {
	var out []interface{}

	if recommendationApplicationBoundaries != nil {
		result := make(map[string]interface{})
		/*value := spotinst.Float64(-1)
		result[string(CpuMin)] = value
		result[string(CpuMax)] = value
		result[string(MemoryMin)] = value
		result[string(MemoryMax)] = value*/

		if recommendationApplicationBoundaries.Cpu.Min != nil {
			result[string(CpuMin)] = spotinst.Float64Value(recommendationApplicationBoundaries.Cpu.Min)
		}
		if recommendationApplicationBoundaries.Cpu.Max != nil {
			result[string(CpuMax)] = spotinst.Float64Value(recommendationApplicationBoundaries.Cpu.Max)
		}
		if recommendationApplicationBoundaries.Memory.Min != nil {
			result[string(MemoryMin)] = spotinst.IntValue(recommendationApplicationBoundaries.Memory.Min)
		}
		if recommendationApplicationBoundaries.Memory.Max != nil {
			result[string(MemoryMax)] = spotinst.IntValue(recommendationApplicationBoundaries.Memory.Max)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

func expandRecommendationApplicationBoundaries(data interface{}) (*right_sizing.RecommendationApplicationBoundaries, error) {
	recommendationApplicationBoundaries := &right_sizing.RecommendationApplicationBoundaries{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return recommendationApplicationBoundaries, nil
	}
	m := list[0].(map[string]interface{})

	cpu := &right_sizing.Cpu{}
	recommendationApplicationBoundaries.SetCpu(cpu)
	if v, ok := m[string(CpuMin)].(float64); ok {
		if v == -1 {
			cpu.SetMin(nil)
		} else {
			cpu.SetMin(spotinst.Float64(v))
		}
	}

	if v, ok := m[string(CpuMax)].(float64); ok {
		if v == -1 {
			cpu.SetMax(nil)
		} else {
			cpu.SetMax(spotinst.Float64(v))
		}
	}

	memory := &right_sizing.Memory{}
	recommendationApplicationBoundaries.SetMemory(memory)
	if v, ok := m[string(MemoryMin)].(int); ok {
		if v == -1 {
			memory.SetMin(nil)
		} else {
			memory.SetMin(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MemoryMax)].(int); ok {
		if v == -1 {
			memory.SetMax(nil)
		} else {
			memory.SetMax(spotinst.Int(v))
		}
	}

	return recommendationApplicationBoundaries, nil
}

func flattenRecommendationApplicationOverheadValues(recommendationApplicationOverheadValues *right_sizing.RecommendationApplicationOverheadValues) []interface{} {
	result := make(map[string]interface{})

	if recommendationApplicationOverheadValues.CpuPercentage != nil {
		result[string(OverheadCpuPercentage)] = spotinst.Float64Value(recommendationApplicationOverheadValues.CpuPercentage)
	}
	if recommendationApplicationOverheadValues.MemoryPercentage != nil {
		result[string(OverheadMemoryPercentage)] = spotinst.Float64Value(recommendationApplicationOverheadValues.MemoryPercentage)
	}
	return []interface{}{result}
}

func expandRecommendationApplicationOverheadValues(data interface{}) (*right_sizing.RecommendationApplicationOverheadValues, error) {
	list := data.(*schema.Set).List()
	recommendationApplicationOverheadValues := &right_sizing.RecommendationApplicationOverheadValues{}

	if len(list) > 0 {
		item := list[0]
		m := item.(map[string]interface{})

		if v, ok := m[string(OverheadCpuPercentage)].(float64); ok {
			if v == -1 {
				recommendationApplicationOverheadValues.SetOverheadCpuPercentage(nil)
			} else {
				recommendationApplicationOverheadValues.SetOverheadCpuPercentage(spotinst.Float64(v))
			}
		}

		if v, ok := m[string(OverheadMemoryPercentage)].(float64); ok {
			if v == -1 {
				recommendationApplicationOverheadValues.SetOverheadMemoryPercentage(nil)
			} else {
				recommendationApplicationOverheadValues.SetOverheadMemoryPercentage(spotinst.Float64(v))
			}
		}

		return recommendationApplicationOverheadValues, nil

	}
	return nil, nil
}

func flattenRecommendationApplicationHPA(recommendationApplicationHPA *right_sizing.RecommendationApplicationHPA) []interface{} {
	result := make(map[string]interface{})

	if recommendationApplicationHPA.AllowHPARecommendations != nil {
		result[string(AllowHPARecommendation)] = spotinst.BoolValue(recommendationApplicationHPA.AllowHPARecommendations)
	}
	return []interface{}{result}
}

func expandRecommendationApplicationHPA(data interface{}) (*right_sizing.RecommendationApplicationHPA, error) {
	list := data.(*schema.Set).List()
	recommendationApplicationHPA := &right_sizing.RecommendationApplicationHPA{}

	if len(list) > 0 {
		item := list[0]
		m := item.(map[string]interface{})

		if v, ok := m[string(AllowHPARecommendation)].(bool); ok {
			recommendationApplicationHPA.SetAllowHPARecommendations(spotinst.Bool(v))
		}

		return recommendationApplicationHPA, nil

	}
	return nil, nil
}
