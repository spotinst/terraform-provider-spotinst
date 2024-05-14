package oceancd_verification_template_metrics

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[ConsecutiveErrorLimit] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		ConsecutiveErrorLimit,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *int = nil
			if verificationTemplate != nil && verificationTemplate.Metrics != nil && verificationTemplate.Metrics[].ConsecutiveErrorLimit != nil {
				value = verificationTemplate.Metrics[].ConsecutiveErrorLimit
			} else {
				value = spotinst.Int(4)
			}
			if err := resourceData.Set(string(ConsecutiveErrorLimit), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ConsecutiveErrorLimit), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(ConsecutiveErrorLimit)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetConsecutiveErrorLimit(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetConsecutiveErrorLimit(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(ConsecutiveErrorLimit)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetConsecutiveErrorLimit(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetConsecutiveErrorLimit(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[Count] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		Count,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *int = nil
			if verificationTemplate != nil && verificationTemplate.Metrics != nil && verificationTemplate.Metrics[].Count != nil {
				value = verificationTemplate.Metrics[].Count
			} else {
				value = spotinst.Int(1)
			}
			if err := resourceData.Set(string(Count), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Count), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(Count)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetCount(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetCount(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(Count)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetCount(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetCount(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[DryRun] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		DryRun,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *bool = nil
			if verificationTemplate.Metrics[] != nil && verificationTemplate.Metrics[].DryRun != nil {
				value = verificationTemplate.Metrics[].DryRun
			}
			if value != nil {
				if err := resourceData.Set(string(DryRun), spotinst.BoolValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DryRun), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(DryRun)); ok && v != nil {
				dryRuns := v.(bool)
				dryRun := spotinst.Bool(dryRuns)
				verificationTemplate.Metrics[].SetDryRun(dryRun)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var dryRun *bool = nil
			if v, ok := resourceData.GetOk(string(DryRun)); ok && v != nil {
				dryRuns := v.(bool)
				dryRun = spotinst.Bool(dryRuns)
			}
			verificationTemplate.Metrics[].SetDryRun(dryRun)
			return nil
		},
		nil,
	)

	fieldsMap[FailureCondition] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		FailureCondition,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if err := resourceData.Set(string(FailureCondition), spotinst.StringValue(verificationTemplate.Metrics[].FailureCondition)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(FailureCondition), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(FailureCondition)); ok {
				verificationTemplate.Metrics[].SetFailureCondition(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(FailureCondition)); ok {
				verificationTemplate.Metrics[].SetFailureCondition(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[FailureLimit] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		FailureLimit,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *int = nil
			if verificationTemplate != nil && verificationTemplate.Metrics != nil && verificationTemplate.Metrics[].FailureLimit != nil {
				value = verificationTemplate.Metrics[].FailureLimit
			} else {
				value = spotinst.Int(0)
			}
			if err := resourceData.Set(string(FailureLimit), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(FailureLimit), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(FailureLimit)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetFailureLimit(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetFailureLimit(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(FailureLimit)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetFailureLimit(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetFailureLimit(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[FailureLimit] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		FailureLimit,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *int = nil
			if verificationTemplate != nil && verificationTemplate.Metrics != nil && verificationTemplate.Metrics[].FailureLimit != nil {
				value = verificationTemplate.Metrics[].FailureLimit
			} else {
				value = spotinst.Int(0)
			}
			if err := resourceData.Set(string(FailureLimit), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(FailureLimit), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(FailureLimit)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetFailureLimit(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetFailureLimit(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(FailureLimit)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetFailureLimit(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetFailureLimit(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[InitialDelay] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		InitialDelay,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if err := resourceData.Set(string(InitialDelay), spotinst.StringValue(verificationTemplate.Metrics[].InitialDelay)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(InitialDelay), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(InitialDelay)); ok {
				verificationTemplate.Metrics[].SetInitialDelay(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(InitialDelay)); ok {
				verificationTemplate.Metrics[].SetInitialDelay(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Interval] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		Interval,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if err := resourceData.Set(string(Interval), spotinst.StringValue(verificationTemplate.Metrics[].Interval)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(Interval), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(Interval)); ok {
				verificationTemplate.Metrics[].SetInterval(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(Interval)); ok {
				verificationTemplate.Metrics[].SetInterval(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if err := resourceData.Set(string(Name), spotinst.StringValue(verificationTemplate.Metrics[].Name)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(Name)); ok {
				verificationTemplate.Metrics[].SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(Name)); ok {
				verificationTemplate.Metrics[].SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[SuccessCondition] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		SuccessCondition,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if err := resourceData.Set(string(SuccessCondition), spotinst.StringValue(verificationTemplate.Metrics[].SuccessCondition)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(SuccessCondition), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(SuccessCondition)); ok {
				verificationTemplate.Metrics[].SetSuccessCondition(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(SuccessCondition)); ok {
				verificationTemplate.Metrics[].SetSuccessCondition(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[BaseLine] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		BaseLine,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MaxRange): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
					string(MinRange): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
					string(Threshold): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(BaseLineProvider): {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Datadog): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Duration): {
												Type:     schema.TypeString,
												Optional: true,
											},
											string(DatadogQuery): {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},
								string(NewRelic): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Profile): {
												Type:     schema.TypeString,
												Optional: true,
											},
											string(NewRelicQuery): {
												Type:     schema.TypeString,
												Required: true,
											},
										},
									},
								},
								string(Prometheus): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(PrometheusQuery): {
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
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var result []interface{} = nil

			if verificationTemplate != nil && verificationTemplate.Metrics[].Baseline != nil {
				result = flattenBaseline(verificationTemplate.Metrics[].Baseline)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(BaseLine), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BaseLine), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *oceancd.Baseline = nil

			if v, ok := resourceData.GetOkExists(string(BaseLine)); ok {
				if baseline, err := expandBaseline(v); err != nil {
					return err
				} else {
					value = baseline
				}
			}
			verificationTemplate.Metrics[].SetBaseLine(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *oceancd.Baseline = nil
			if v, ok := resourceData.GetOkExists(string(BaseLine)); ok {
				if baseline, err := expandBaseline(v); err != nil {
					return err
				} else {
					value = baseline
				}
			}
			verificationTemplate.Metrics[].SetBaseLine(value)
			return nil
		},
		nil,
	)

	fieldsMap[Provider] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		Provider,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Datadog): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Duration): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(DatadogQuery): {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},
					string(NewRelic): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Profile): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(NewRelicQuery): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},
					string(Prometheus): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(PrometheusQuery): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},
					string(CloudWatch): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Duration): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(MetricDataQueries): {
									Type:     schema.TypeSet,
									Required: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Expression): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(ID): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(Label): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(Period): {
												Type:     schema.TypeInt,
												Required: true,
												Default:  -1,
											},
											string(ReturnData): {
												Type:     schema.TypeBool,
												Required: true,
											},
											string(MetricStat): {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														string(Stat): {
															Type:     schema.TypeString,
															Required: true,
														},
														string(Unit): {
															Type:     schema.TypeString,
															Required: true,
														},
														string(MetricPeriod): {
															Type:     schema.TypeInt,
															Required: true,
															Default:  -1,
														},
														string(Metric): {
															Type:     schema.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	string(MetricName): {
																		Type:     schema.TypeString,
																		Required: true,
																	},
																	string(Namespace): {
																		Type:     schema.TypeString,
																		Optional: true,
																	},
																	string(Dimensions): {
																		Type:     schema.TypeList,
																		Optional: true,
																		MaxItems: 1,
																		Elem: &schema.Resource{
																			Schema: map[string]*schema.Schema{
																				string(DimensionName): {
																					Type:     schema.TypeString,
																					Required: true,
																				},
																				string(DimensionValue): {
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
										},
									},
								},
							},
						},
					},
					string(Jenkins): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(JenkinsInterval): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(PipelineName): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(Timeout): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(TlsVerification): {
									Type:     schema.TypeBool,
									Required: true,
								},
								string(JenkinsParameters): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(ParameterKey): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(ParameterValue): {
												Type:     schema.TypeString,
												Required: true,
											},
										},
									},
								},
							},
						},
					},
					string(Web): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Body): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(Insecure): {
									Type:     schema.TypeBool,
									Optional: true,
								},
								string(JsonPath): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(Method): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(Url): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(TimeoutSeconds): {
									Type:     schema.TypeInt,
									Required: true,
									Default:  10,
								},
								string(WebHeader): {
									Type:     schema.TypeSet,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(WebHeaderKey): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(WebHeaderValue): {
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
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var result []interface{} = nil

			if verificationTemplate != nil && verificationTemplate.Metrics[].Provider != nil {
				result = flattenProvider(verificationTemplate.Metrics[].Provider)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Provider), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Provider), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *oceancd.Provider = nil

			if v, ok := resourceData.GetOkExists(string(Provider)); ok {
				if provider, err := expandProvider(v); err != nil {
					return err
				} else {
					value = provider
				}
			}
			verificationTemplate.Metrics[].SetProvider(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *oceancd.Provider = nil
			if v, ok := resourceData.GetOkExists(string(Provider)); ok {
				if provider, err := expandProvider(v); err != nil {
					return err
				} else {
					value = provider
				}
			}
			verificationTemplate.Metrics[].SetProvider(value)
			return nil
		},
		nil,
	)
}

func expandBaseline(data interface{}) (*oceancd.Baseline, error) {
	if list := data.([]interface{}); len(list) > 0 {
		baseline := &oceancd.Baseline{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(Threshold)].(string); ok && v != "" {
				baseline.SetThreshold(spotinst.String(v))
			}

			if v, ok := m[string(MaxRange)].(int); ok {
				if v == -1 {
					baseline.SetMaxRange(nil)
				} else {
					baseline.SetMaxRange(spotinst.Int(v))
				}
			}

			if v, ok := m[string(MinRange)].(int); ok {
				if v == -1 {
					baseline.SetMinRange(nil)
				} else {
					baseline.SetMinRange(spotinst.Int(v))
				}
			}

			if v, ok := m[string(BaseLineProvider)]; ok {
				provider, err := expandBaselineProvider(v)
				if err != nil {
					return nil, err
				}
				if provider != nil {
					baseline.SetProvider(provider)
				} else {
					baseline.SetProvider(nil)
				}
			}
		}
		return baseline, nil
	}
	return nil, nil
}

func expandBaselineProvider(data interface{}) (*oceancd.Provider, error) {

	provider := &oceancd.Provider{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return provider, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Datadog)]; ok {
		datadog, err := expandDatadog(v)
		if err != nil {
			return nil, err
		}
		if datadog != nil {
			provider.SetDataDog(datadog)
		} else {
			provider.SetDataDog(nil)
		}
	}

	if v, ok := m[string(NewRelic)]; ok {
		newRelic, err := expandNewRelic(v)
		if err != nil {
			return nil, err
		}
		if newRelic != nil {
			provider.SetNewRelic(newRelic)
		} else {
			provider.SetNewRelic(nil)
		}
	}

	if v, ok := m[string(Prometheus)]; ok {
		prometheus, err := expandPrometheus(v)
		if err != nil {
			return nil, err
		}
		if prometheus != nil {
			provider.SetPrometheus(prometheus)
		} else {
			provider.SetPrometheus(nil)
		}
	}
	return provider, nil
}

func expandDatadog(data interface{}) (*oceancd.DataDogProvider, error) {

	datadog := &oceancd.DataDogProvider{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return datadog, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Duration)].(string); ok && v != "" {
		datadog.SetDuration(spotinst.String(v))
	} else {
		datadog.SetDuration(nil)
	}

	if v, ok := m[string(DatadogQuery)].(string); ok && v != "" {
		datadog.SetQuery(spotinst.String(v))
	}

	return datadog, nil
}

func expandNewRelic(data interface{}) (*oceancd.NewRelicProvider, error) {
	newRelic := &oceancd.NewRelicProvider{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return newRelic, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Profile)].(string); ok && v != "" {
		newRelic.SetProfile(spotinst.String(v))
	} else {
		newRelic.SetProfile(nil)
	}

	if v, ok := m[string(NewRelicQuery)].(string); ok && v != "" {
		newRelic.SetProfile(spotinst.String(v))
	}
	return newRelic, nil
}

func expandPrometheus(data interface{}) (*oceancd.PrometheusProvider, error) {

	prometheus := &oceancd.PrometheusProvider{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return prometheus, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Duration)].(string); ok && v != "" {
		prometheus.SetQuery(spotinst.String(v))
	}
	return prometheus, nil
}

func expandProvider(data interface{}) (*oceancd.Provider, error) {

	provider := &oceancd.Provider{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return provider, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Datadog)]; ok {
		datadog, err := expandDatadog(v)
		if err != nil {
			return nil, err
		}
		if datadog != nil {
			provider.SetDataDog(datadog)
		} else {
			provider.SetDataDog(nil)
		}
	}

	if v, ok := m[string(NewRelic)]; ok {
		newRelic, err := expandNewRelic(v)
		if err != nil {
			return nil, err
		}
		if newRelic != nil {
			provider.SetNewRelic(newRelic)
		} else {
			provider.SetNewRelic(nil)
		}
	}

	if v, ok := m[string(Prometheus)]; ok {
		prometheus, err := expandPrometheus(v)
		if err != nil {
			return nil, err
		}
		if prometheus != nil {
			provider.SetPrometheus(prometheus)
		} else {
			provider.SetPrometheus(nil)
		}
	}

	if v, ok := m[string(Jenkins)]; ok {
		jenkins, err := expandJenkins(v)
		if err != nil {
			return nil, err
		}
		if jenkins != nil {
			provider.SetJenkins(jenkins)
		} else {
			provider.SetJenkins(nil)
		}
	}

	if v, ok := m[string(CloudWatch)]; ok {
		cloudWatch, err := expandCloudWatch(v)
		if err != nil {
			return nil, err
		}
		if cloudWatch != nil {
			provider.SetCloudWatch(cloudWatch)
		} else {
			provider.SetCloudWatch(nil)
		}
	}

	if v, ok := m[string(Web)]; ok {
		web, err := expandWeb(v)
		if err != nil {
			return nil, err
		}
		if web != nil {
			provider.SetWeb(web)
		} else {
			provider.SetWeb(nil)
		}
	}
	return provider, nil
}

func expandJenkins(data interface{}) (*oceancd.JenkinsProvider, error) {
	jenkins := &oceancd.JenkinsProvider{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return jenkins, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(PipelineName)].(string); ok && v != "" {
		jenkins.SetPipelineName(spotinst.String(v))
	}

	if v, ok := m[string(Timeout)].(string); ok && v != "" {
		jenkins.SetTimeout(spotinst.String(v))
	}

	if v, ok := m[string(Interval)].(string); ok && v != "" {
		jenkins.SetInterval(spotinst.String(v))
	}

	if v, ok := m[string(TlsVerification)].(bool); ok {
		jenkins.SetTLSVerification(spotinst.Bool(v))
	}

	if v, ok := m[string(JenkinsParameters)]; ok {
		parameters, err := expandParameters(v)
		if err != nil {
			return nil, err
		}
		if parameters != nil {
			jenkins.SetParameters(parameters)
		} else {
			jenkins.SetParameters(nil)
		}
	}
	return jenkins, nil
}

func expandParameters(data interface{}) ([]*oceancd.Parameters, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		parameters := make([]*oceancd.Parameters, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			parameter := &oceancd.Parameters{}

			if v, ok := m[string(ParameterKey)].(string); ok && v != "" {
				parameter.SetKey(spotinst.String(v))
			}
			if v, ok := m[string(ParameterValue)].(string); ok && v != "" {
				parameter.SetValue(spotinst.String(v))
			}

			parameters = append(parameters, parameter)
		}
		return parameters, nil
	}
	return nil, nil
}

func expandWeb(data interface{}) (*oceancd.Web, error) {
	web := &oceancd.Web{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return web, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Body)].(string); ok && v != "" {
		web.SetBody(spotinst.String(v))
	}

	if v, ok := m[string(JsonPath)].(string); ok && v != "" {
		web.SetJsonPath(spotinst.String(v))
	}

	if v, ok := m[string(Method)].(string); ok && v != "" {
		web.SetMethod(spotinst.String(v))
	}

	if v, ok := m[string(Insecure)].(bool); ok {
		web.SetInsecure(spotinst.Bool(v))
	}

	if v, ok := m[string(TimeoutSeconds)].(int); ok {
		if v == -1 {
			web.SetTimeoutSeconds(nil)
		} else {
			web.SetTimeoutSeconds(spotinst.Int(v))
		}
	}

	if v, ok := m[string(WebHeader)]; ok {
		headers, err := expandHeaders(v)
		if err != nil {
			return nil, err
		}
		if headers != nil {
			web.SetHeaders(headers)
		} else {
			web.SetHeaders(nil)
		}
	}
	return web, nil
}

func expandHeaders(data interface{}) ([]*oceancd.Headers, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		headers := make([]*oceancd.Headers, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			header := &oceancd.Headers{}

			if v, ok := m[string(WebHeaderKey)].(string); ok && v != "" {
				header.SetKey(spotinst.String(v))
			}
			if v, ok := m[string(WebHeaderValue)].(string); ok && v != "" {
				header.SetValue(spotinst.String(v))
			}

			headers = append(headers, header)
		}
		return headers, nil
	}
	return nil, nil
}

func expandCloudWatch(data interface{}) (*oceancd.CloudWatchProvider, error) {
	cloudWatch := &oceancd.CloudWatchProvider{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return cloudWatch, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(CloudWatchDuration)].(string); ok && v != "" {
		cloudWatch.SetDuration(spotinst.String(v))
	}

	if v, ok := m[string(MetricDataQueries)]; ok {
		metricDataQueries, err := expandMetricDataQueries(v)
		if err != nil {
			return nil, err
		}
		if metricDataQueries != nil {
			cloudWatch.SetMetricDataQueries(metricDataQueries)
		} else {
			cloudWatch.SetMetricDataQueries(nil)
		}
	}
	return cloudWatch, nil
}

func expandMetricDataQueries(data interface{}) ([]*oceancd.MetricDataQueries, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		metricDataQueries := make([]*oceancd.MetricDataQueries, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			metricDataQuery := &oceancd.MetricDataQueries{}

			if v, ok := m[string(ID)].(string); ok && v != "" {
				metricDataQuery.SetID(spotinst.String(v))
			}
			if v, ok := m[string(Label)].(string); ok && v != "" {
				metricDataQuery.SetLabel(spotinst.String(v))
			}
			if v, ok := m[string(Expression)].(string); ok && v != "" {
				metricDataQuery.SetExpression(spotinst.String(v))
			}

			if v, ok := m[string(ReturnData)].(bool); ok {
				metricDataQuery.SetReturnData(spotinst.Bool(v))
			}

			if v, ok := m[string(Period)].(int); ok {
				if v == -1 {
					metricDataQuery.SetPeriod(nil)
				} else {
					metricDataQuery.SetPeriod(spotinst.Int(v))
				}
			}

			if v, ok := m[string(MetricStat)]; ok {
				metricStats, err := expandMetricStats(v)
				if err != nil {
					return nil, err
				}
				if metricStats != nil {
					metricDataQuery.SetMetricStat(metricStats)
				} else {
					metricDataQuery.SetMetricStat(nil)
				}
			}

			metricDataQueries = append(metricDataQueries, metricDataQuery)
		}
		return metricDataQueries, nil
	}
	return nil, nil
}

func expandMetricStats(data interface{}) (*oceancd.MetricStat, error) {
	metricStat := &oceancd.MetricStat{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return metricStat, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Stat)].(string); ok && v != "" {
		metricStat.SetStat(spotinst.String(v))
	} else {
		metricStat.SetStat(nil)
	}

	if v, ok := m[string(Unit)].(string); ok && v != "" {
		metricStat.SetUnit(spotinst.String(v))
	} else {
		metricStat.SetUnit(nil)
	}

	if v, ok := m[string(Period)].(int); ok {
		if v == -1 {
			metricStat.SetPeriod(nil)
		} else {
			metricStat.SetPeriod(spotinst.Int(v))
		}
	}

	if v, ok := m[string(Metric)]; ok {
		metric, err := expandMetric(v)
		if err != nil {
			return nil, err
		}
		if metric != nil {
			metricStat.SetMetric(metric)
		} else {
			metricStat.SetMetric(nil)
		}
	}
	return metricStat, nil
}

func expandMetric(data interface{}) (*oceancd.Metric, error) {
	metric := &oceancd.Metric{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return metric, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Namespace)].(string); ok && v != "" {
		metric.SetNamespace(spotinst.String(v))
	} else {
		metric.SetNamespace(nil)
	}

	if v, ok := m[string(MetricName)].(string); ok && v != "" {
		metric.SetMetricName(spotinst.String(v))
	}

	if v, ok := m[string(Dimensions)]; ok {
		dimensions, err := expandDimensions(v)
		if err != nil {
			return nil, err
		}
		if dimensions != nil {
			metric.SetDimensions(dimensions)
		} else {
			metric.SetDimensions(nil)
		}
	}
	return metric, nil
}

func expandDimensions(data interface{}) ([]*oceancd.Dimensions, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		dimensions := make([]*oceancd.Dimensions, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			dimension := &oceancd.Dimensions{}

			if v, ok := m[string(DimensionName)].(string); ok && v != "" {
				dimension.SetName(spotinst.String(v))
			}
			if v, ok := m[string(DimensionValue)].(string); ok && v != "" {
				dimension.SetValue(spotinst.String(v))
			}

			dimensions = append(dimensions, dimension)
		}
		return dimensions, nil
	}
	return nil, nil
}

func flattenBaseline(baseline *oceancd.Baseline) []interface{} {
	result := make(map[string]interface{})
	value := spotinst.Int(-1)
	result[string(MaxRange)] = value
	result[string(MinRange)] = value

	if baseline.MaxRange != nil {
		result[string(MaxRange)] = spotinst.IntValue(baseline.MaxRange)
	}

	if baseline.MinRange != nil {
		result[string(MinRange)] = spotinst.IntValue(baseline.MinRange)
	}

	if baseline.Provider != nil {
		result[string(BaseLineProvider)] = flattenBaselineProvider(baseline.Provider)
	}

	result[string(Threshold)] = spotinst.StringValue(baseline.Threshold)

	return []interface{}{result}
}

func flattenProvider(provider *oceancd.Provider) []interface{} {
	result := make(map[string]interface{})

	if provider.Datadog != nil {
		result[string(Datadog)] = flattenDatadog(provider.Datadog)
	}

	if provider.NewRelic != nil {
		result[string(NewRelic)] = flattenNewRelic(provider.NewRelic)
	}

	if provider.Prometheus != nil {
		result[string(Prometheus)] = flattenPrometheus(provider.Prometheus)
	}

	if provider.CloudWatch != nil {
		result[string(CloudWatch)] = flattenCloudWatch(provider.CloudWatch)
	}

	if provider.Jenkins != nil {
		result[string(Jenkins)] = flattenJenkins(provider.Jenkins)
	}

	if provider.Web != nil {
		result[string(Web)] = flattenWeb(provider.Web)
	}
	return []interface{}{result}
}

func flattenBaselineProvider(provider *oceancd.Provider) []interface{} {
	result := make(map[string]interface{})

	if provider.Datadog != nil {
		result[string(Datadog)] = flattenDatadog(provider.Datadog)
	}

	if provider.NewRelic != nil {
		result[string(NewRelic)] = flattenNewRelic(provider.NewRelic)
	}

	if provider.Prometheus != nil {
		result[string(Prometheus)] = flattenPrometheus(provider.Prometheus)
	}
	return []interface{}{result}
}

func flattenDatadog(datadog *oceancd.DataDogProvider) []interface{} {
	result := make(map[string]interface{})
	result[string(Duration)] = spotinst.StringValue(datadog.Duration)
	result[string(DatadogQuery)] = spotinst.StringValue(datadog.Query)
	return []interface{}{result}
}

func flattenNewRelic(newRelic *oceancd.NewRelicProvider) []interface{} {
	result := make(map[string]interface{})
	result[string(Duration)] = spotinst.StringValue(newRelic.Profile)
	result[string(NewRelicQuery)] = spotinst.StringValue(newRelic.Query)
	return []interface{}{result}
}

func flattenPrometheus(prometheus *oceancd.PrometheusProvider) []interface{} {
	result := make(map[string]interface{})
	result[string(PrometheusQuery)] = spotinst.StringValue(prometheus.Query)
	return []interface{}{result}
}

func flattenJenkins(jenkins *oceancd.JenkinsProvider) []interface{} {
	result := make(map[string]interface{})
	result[string(Interval)] = spotinst.StringValue(jenkins.Interval)
	result[string(PipelineName)] = spotinst.StringValue(jenkins.PipelineName)
	result[string(Timeout)] = spotinst.StringValue(jenkins.Timeout)

	result[string(TlsVerification)] = spotinst.BoolValue(jenkins.TLSVerification)

	if jenkins.Parameters != nil {
		result[string(Prometheus)] = flattenParameters(jenkins.Parameters)
	}

	return []interface{}{result}
}

func flattenParameters(parameters []*oceancd.Parameters) []interface{} {
	m := make([]interface{}, 0, len(parameters))
	for _, parameter := range parameters {
		result := make(map[string]interface{})
		result[string(ParameterValue)] = spotinst.StringValue(parameter.Value)
		result[string(ParameterKey)] = spotinst.StringValue(parameter.Key)
	}
	return []interface{}{m}
}

func flattenWeb(web *oceancd.Web) []interface{} {
	result := make(map[string]interface{})
	value := spotinst.Int(-1)
	result[string(TimeoutSeconds)] = value

	result[string(Body)] = spotinst.StringValue(web.Body)
	result[string(JsonPath)] = spotinst.StringValue(web.JsonPath)
	result[string(Method)] = spotinst.StringValue(web.Method)

	result[string(Insecure)] = spotinst.BoolValue(web.Insecure)

	if web.Headers != nil {
		result[string(Prometheus)] = flattenHeaders(web.Headers)
	}

	if web.TimeoutSeconds != nil {
		result[string(TimeoutSeconds)] = spotinst.IntValue(web.TimeoutSeconds)
	}

	return []interface{}{result}
}

func flattenHeaders(headers []*oceancd.Headers) []interface{} {
	m := make([]interface{}, 0, len(headers))
	for _, header := range headers {
		result := make(map[string]interface{})
		result[string(WebHeaderKey)] = spotinst.StringValue(header.Key)
		result[string(WebHeaderValue)] = spotinst.StringValue(header.Value)
	}
	return []interface{}{m}
}

func flattenCloudWatch(cloudWatch *oceancd.CloudWatchProvider) []interface{} {
	result := make(map[string]interface{})
	result[string(MetricDataQueries)] = spotinst.StringValue(cloudWatch.Duration)

	if cloudWatch.MetricDataQueries != nil {
		result[string(MetricDataQueries)] = flattenMetricDataQueries(cloudWatch.MetricDataQueries)
	}

	return []interface{}{result}
}

func flattenMetricDataQueries(metricDataQueries []*oceancd.MetricDataQueries) []interface{} {
	m := make([]interface{}, 0, len(metricDataQueries))
	for _, metricDataQuery := range metricDataQueries {
		result := make(map[string]interface{})
		value := spotinst.Int(-1)
		result[string(Period)] = value

		result[string(Expression)] = spotinst.StringValue(metricDataQuery.Expression)
		result[string(ID)] = spotinst.StringValue(metricDataQuery.ID)
		result[string(Label)] = spotinst.StringValue(metricDataQuery.Label)

		result[string(ReturnData)] = spotinst.BoolValue(metricDataQuery.ReturnData)

		if metricDataQuery.MetricStat != nil {
			result[string(MetricStat)] = flattenMetricStat(metricDataQuery.MetricStat)
		}

		if metricDataQuery.Period != nil {
			result[string(Period)] = spotinst.IntValue(metricDataQuery.Period)
		}

	}
	return []interface{}{m}
}

func flattenMetricStat(metricStat *oceancd.MetricStat) []interface{} {
	result := make(map[string]interface{})
	value := spotinst.Int(-1)
	result[string(Period)] = value

	result[string(Stat)] = spotinst.StringValue(metricStat.Stat)
	result[string(Unit)] = spotinst.StringValue(metricStat.Unit)

	if metricStat.Metric != nil {
		result[string(Metric)] = flattenMetric(metricStat.Metric)
	}

	if metricStat.Period != nil {
		result[string(Period)] = spotinst.IntValue(metricStat.Period)
	}

	return []interface{}{result}
}

func flattenMetric(metric *oceancd.Metric) []interface{} {
	result := make(map[string]interface{})

	result[string(MetricName)] = spotinst.StringValue(metric.MetricName)
	result[string(Namespace)] = spotinst.StringValue(metric.Namespace)

	if metric.Dimensions != nil {
		result[string(Dimensions)] = flattenDimensions(metric.Dimensions)
	}

	return []interface{}{result}
}

func flattenDimensions(dimensions []*oceancd.Dimensions) []interface{} {
	m := make([]interface{}, 0, len(dimensions))
	for _, dimension := range dimensions {
		result := make(map[string]interface{})
		result[string(DimensionName)] = spotinst.StringValue(dimension.Name)
		result[string(DimensionValue)] = spotinst.StringValue(dimension.Value)
	}
	return []interface{}{m}
}
