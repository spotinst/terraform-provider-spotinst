package ocean_aws_right_sizing_rule

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanAWSRightSizingRule,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanAWSRightSizingRule()
			var value *string = nil
			if rightSizingRule.Name != nil {
				value = rightSizingRule.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanAWSRightSizingRule()
			if v, ok := resourceData.GetOk(string(Name)); ok && v != "" {
				rightSizingRule.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanAWSRightSizingRule()
			if v, ok := resourceData.GetOk(string(Name)); ok && v != "" {
				rightSizingRule.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
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
			rightSizingRule := rightSizingRuleWrapper.GetOceanAWSRightSizingRule()
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
		commons.OceanAWSRightSizingRule,
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
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanAWSRightSizingRule()
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
			rightSizingRule := rightSizingRuleWrapper.GetOceanAWSRightSizingRule()
			var value []*aws.RecommendationApplicationInterval = nil
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
		commons.OceanAWSRightSizingRule,
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
			rightSizingRule := rightSizingRuleWrapper.GetOceanAWSRightSizingRule()
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
			rightSizingRule := rightSizingRuleWrapper.GetOceanAWSRightSizingRule()
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
			rightSizingRule := rightSizingRuleWrapper.GetOceanAWSRightSizingRule()
			var value *aws.RecommendationApplicationMinThreshold = nil
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

	fieldsMap[RecommendationApplicationBoundaries] = commons.NewGenericField(
		commons.OceanAWSRightSizingRule,
		RecommendationApplicationBoundaries,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(CpuMin): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
					string(CpuMax): {
						Type:     schema.TypeInt,
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
			rightSizingRule := rightSizingRuleWrapper.GetOceanAWSRightSizingRule()
			var result []interface{} = nil
			if rightSizingRule.RecommendationApplicationBoundaries != nil {
				recommendationApplicationBoundaries := rightSizingRule.RecommendationApplicationBoundaries
				result = flattenRecommendationApplicationBoundaries(recommendationApplicationBoundaries)
			}
			if result != nil {
				if err := resourceData.Set(string(RecommendationApplicationBoundaries), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RecommendationApplicationBoundaries), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rightSizingRuleWrapper := resourceObject.(*commons.RightSizingRuleWrapper)
			rightSizingRule := rightSizingRuleWrapper.GetOceanAWSRightSizingRule()
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
			rightSizingRule := rightSizingRuleWrapper.GetOceanAWSRightSizingRule()
			var value *aws.RecommendationApplicationBoundaries = nil
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

	fieldsMap[AttachWorkloads] = commons.NewGenericField(
		commons.OceanAWSRightSizingRule,
		AttachWorkloads,
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
									Required: true,
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
									Required: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Key): {
												Type:     schema.TypeString,
												Optional: true,
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

}

func flattenRecommendationApplicationIntervals(recommendationApplicationIntervals []*aws.RecommendationApplicationInterval) []interface{} {
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

func flattenWeeklyRepetitionBasis(weeklyRepetitionBasis *aws.WeeklyRepetitionBasis) []interface{} {
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

func flattenMonthlyRepetitionBasis(monthlyRepetitionBasis *aws.MonthlyRepetitionBasis) []interface{} {
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

func flattenMonthlyWeeklyRepetitionBasis(weeklyRepetitionBasis *aws.WeeklyRepetitionBasis) []interface{} {
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

// expandVNGNetworkInterface sets the values from the plan as objects
func expandRecommendationApplicationIntervals(data interface{}) ([]*aws.RecommendationApplicationInterval, error) {
	list := data.(*schema.Set).List()

	if list != nil && list[0] != nil {
		ifaces := make([]*aws.RecommendationApplicationInterval, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			iface := &aws.RecommendationApplicationInterval{}

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

func expandWeeklyRepetitionBasis(data interface{}) (*aws.WeeklyRepetitionBasis, error) {
	list := data.(*schema.Set).List()
	weeklyRepetitionBasis := &aws.WeeklyRepetitionBasis{}
	intervalHours := &aws.IntervalHours{}

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

func expandMonthlyRepetitionBasis(data interface{}) (*aws.MonthlyRepetitionBasis, error) {
	list := data.(*schema.Set).List()
	monthlyRepetitionBasis := &aws.MonthlyRepetitionBasis{}

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

func expandMonthlyWeeklyRepetitionBasis(data interface{}) (*aws.WeeklyRepetitionBasis, error) {
	list := data.(*schema.Set).List()
	weeklyRepetitionBasis := &aws.WeeklyRepetitionBasis{}
	intervalHours := &aws.IntervalHours{}

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

func flattenRecommendationApplicationMinThreshold(recommendationApplicationMinThreshold *aws.RecommendationApplicationMinThreshold) []interface{} {
	result := make(map[string]interface{})

	if recommendationApplicationMinThreshold.CpuPercentage != nil {
		result[string(CpuPercentage)] = spotinst.Float64Value(recommendationApplicationMinThreshold.CpuPercentage)
	}
	if recommendationApplicationMinThreshold.MemoryPercentage != nil {
		result[string(MemoryPercentage)] = spotinst.Float64Value(recommendationApplicationMinThreshold.CpuPercentage)
	}
	return []interface{}{result}
}

func expandRecommendationApplicationMinThreshold(data interface{}) (*aws.RecommendationApplicationMinThreshold, error) {
	list := data.(*schema.Set).List()
	recommendationApplicationMinThreshold := &aws.RecommendationApplicationMinThreshold{}

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

func flattenRecommendationApplicationBoundaries(recommendationApplicationBoundaries *aws.RecommendationApplicationBoundaries) []interface{} {
	result := make(map[string]interface{})

	if recommendationApplicationBoundaries.Cpu.Min != nil {
		result[string(CpuMin)] = spotinst.IntValue(recommendationApplicationBoundaries.Cpu.Min)
	}
	if recommendationApplicationBoundaries.Cpu.Max != nil {
		result[string(CpuMax)] = spotinst.IntValue(recommendationApplicationBoundaries.Cpu.Max)
	}
	if recommendationApplicationBoundaries.Memory.Min != nil {
		result[string(MemoryMin)] = spotinst.IntValue(recommendationApplicationBoundaries.Memory.Min)
	}
	if recommendationApplicationBoundaries.Memory.Max != nil {
		result[string(MemoryMax)] = spotinst.IntValue(recommendationApplicationBoundaries.Memory.Max)
	}
	return []interface{}{result}
}

func expandRecommendationApplicationBoundaries(data interface{}) (*aws.RecommendationApplicationBoundaries, error) {
	list := data.(*schema.Set).List()
	recommendationApplicationBoundaries := &aws.RecommendationApplicationBoundaries{}
	cpu := &aws.Cpu{}
	memory := &aws.Memory{}

	if len(list) > 0 {
		item := list[0]
		m := item.(map[string]interface{})

		if v, ok := m[string(CpuMin)].(int); ok {
			if v == -1 {
				cpu.SetMin(nil)
			} else {
				cpu.SetMin(spotinst.Int(v))
			}
		}

		if v, ok := m[string(CpuMax)].(int); ok {
			if v == -1 {
				cpu.SetMax(nil)
			} else {
				cpu.SetMax(spotinst.Int(v))
			}
		}

		recommendationApplicationBoundaries.SetCpu(cpu)

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

		recommendationApplicationBoundaries.SetMemory(memory)

		return recommendationApplicationBoundaries, nil

	}
	return nil, nil
}
