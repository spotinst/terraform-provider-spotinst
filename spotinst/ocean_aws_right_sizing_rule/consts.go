package ocean_aws_right_sizing_rule

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	OceanId commons.FieldName = "ocean_id"
	Name    commons.FieldName = "name"

	RecommendationApplicationIntervals  commons.FieldName = "recommendation_application_intervals"
	RepetitionBasis                     commons.FieldName = "repetition_basis"
	WeeklyRepetitionBasis               commons.FieldName = "weekly_repetition_basis"
	IntervalDays                        commons.FieldName = "interval_days"
	IntervalHoursStartTime              commons.FieldName = "interval_hours_start_time"
	IntervalHoursEndTime                commons.FieldName = "interval_hours_end_time"
	MonthlyRepetitionBasis              commons.FieldName = "monthly_repetition_basis"
	IntervalMonths                      commons.FieldName = "interval_months"
	WeekOfTheMonth                      commons.FieldName = "week_of_the_month"
	MonthlyWeeklyRepetitionBasis        commons.FieldName = "weekly_repetition_basis"
	MonthlyWeeklyIntervalDays           commons.FieldName = "interval_days"
	MonthlyWeeklyIntervalHoursStartTime commons.FieldName = "interval_hours_start_time"
	MonthlyWeeklyIntervalHoursEndTime   commons.FieldName = "interval_hours_end_time"

	RecommendationApplicationBoundaries commons.FieldName = "RecommendationApplicationBoundaries"
	CpuMin                              commons.FieldName = "cpu_min"
	CpuMax                              commons.FieldName = "cpu_max"
	MemoryMin                           commons.FieldName = "memory_min"
	MemoryMax                           commons.FieldName = "memory_max"

	RecommendationApplicationMinThreshold commons.FieldName = "recommendation_application_min_threshold"
	CpuPercentage                         commons.FieldName = "cpu_percentage"
	MemoryPercentage                      commons.FieldName = "memory_percentage"
)
