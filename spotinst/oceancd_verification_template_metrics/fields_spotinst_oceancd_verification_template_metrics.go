package oceancd_verification_template_metrics

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Metrics] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		Metrics,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ConsecutiveErrorLimit): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  4,
					},

					string(Count): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  1,
					},

					string(DryRun): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(FailureCondition): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(FailureLimit): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  0,
					},

					string(InitialDelay): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(Interval): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(MetricsName): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(SuccessCondition): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(BaseLine): {
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

					string(Provider): {
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
															Optional: true,
														},
														string(ID): {
															Type:     schema.TypeString,
															Required: true,
														},
														string(Label): {
															Type:     schema.TypeString,
															Optional: true,
														},
														string(Period): {
															Type:     schema.TypeInt,
															Optional: true,
															Default:  -1,
														},
														string(ReturnData): {
															Type:     schema.TypeBool,
															Optional: true,
														},
														string(MetricStat): {
															Type:     schema.TypeList,
															Optional: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	string(Stat): {
																		Type:     schema.TypeString,
																		Optional: true,
																	},
																	string(Unit): {
																		Type:     schema.TypeString,
																		Optional: true,
																	},
																	string(MetricPeriod): {
																		Type:     schema.TypeInt,
																		Optional: true,
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
																					Type:     schema.TypeSet,
																					Optional: true,
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
								string(Job): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Spec): {
												Type:     schema.TypeList,
												Required: true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														string(BackoffLimit): {
															Type:     schema.TypeString,
															Optional: true,
														},
														string(JobTemplate): {
															Type:     schema.TypeList,
															Required: true,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	string(TemplateSpec): {
																		Type:     schema.TypeList,
																		Required: true,
																		Elem: &schema.Resource{
																			Schema: map[string]*schema.Schema{
																				string(RestartPolicy): {
																					Type:     schema.TypeString,
																					Required: true,
																				},
																				string(Containers): {
																					Type:     schema.TypeSet,
																					Required: true,
																					Elem: &schema.Resource{
																						Schema: map[string]*schema.Schema{
																							string(Image): {
																								Type:     schema.TypeString,
																								Required: true,
																							},
																							string(ContinerName): {
																								Type:     schema.TypeString,
																								Required: true,
																							},
																							string(Command): {
																								Type:     schema.TypeList,
																								Required: true,
																								Elem:     &schema.Schema{Type: schema.TypeString},
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
												Optional: true,
											},
											string(JenkinsParameters): {
												Type:     schema.TypeSet,
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
												Optional: true,
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
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()

			var metricsResults []interface{} = nil
			if verificationTemplate != nil && verificationTemplate.Metrics != nil {
				metrics := verificationTemplate.Metrics
				metricsResults = flattenMetrics(metrics)
			}

			if err := resourceData.Set(string(Metrics), metricsResults); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Metrics), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if value, ok := resourceData.GetOkExists(string(Metrics)); ok {
				if metrics, err := expandMetrics(value); err != nil {
					return err
				} else {
					verificationTemplate.SetMetrics(metrics)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var result []*oceancd.Metrics = nil
			if value, ok := resourceData.GetOkExists(string(Metrics)); ok {
				if metrics, err := expandMetrics(value); err != nil {
					return err
				} else {
					result = metrics
				}
			}

			if len(result) == 0 {
				verificationTemplate.SetMetrics(result)
			} else {
				verificationTemplate.SetMetrics(result)
			}

			return nil
		}, nil,
	)
}

func expandMetrics(data interface{}) ([]*oceancd.Metrics, error) {
	list := data.(*schema.Set).List()
	metrics := make([]*oceancd.Metrics, 0, len(list))

	for _, v := range list {
		m := v.(map[string]interface{})
		metric := &oceancd.Metrics{}

		if v, ok := m[string(ConsecutiveErrorLimit)].(int); ok {
			if v == 4 {
				metric.SetConsecutiveErrorLimit(nil)
			} else {
				metric.SetConsecutiveErrorLimit(spotinst.Int(v))
			}
		}

		if v, ok := m[string(Count)].(int); ok {
			if v == 1 {
				metric.SetCount(nil)
			} else {
				metric.SetCount(spotinst.Int(v))
			}
		}

		if v, ok := m[string(FailureLimit)].(int); ok {
			if v == 1 {
				metric.SetCount(nil)
			} else {
				metric.SetCount(spotinst.Int(v))
			}
		}

		if v, ok := m[string(FailureCondition)].(string); ok && v != "" {
			metric.SetFailureCondition(spotinst.String(v))
		} else {
			metric.SetFailureCondition(nil)
		}

		if v, ok := m[string(InitialDelay)].(string); ok && v != "" {
			metric.SetInitialDelay(spotinst.String(v))
		} else {
			metric.SetInitialDelay(nil)
		}

		if v, ok := m[string(Interval)].(string); ok && v != "" {
			metric.SetInterval(spotinst.String(v))
		} else {
			metric.SetInterval(nil)
		}

		if v, ok := m[string(SuccessCondition)].(string); ok && v != "" {
			metric.SetSuccessCondition(spotinst.String(v))
		} else {
			metric.SetSuccessCondition(nil)
		}

		if v, ok := m[string(MetricsName)].(string); ok && v != "" {
			metric.SetName(spotinst.String(v))
		}

		if v, ok := m[string(DryRun)].(bool); ok {
			metric.SetDryRun(spotinst.Bool(v))
		} else {
			metric.SetDryRun(nil)
		}

		if v, ok := m[string(BaseLine)]; ok {
			baseline, err := expandBaseline(v)
			if err != nil {
				return nil, err
			}
			if baseline != nil {
				metric.SetBaseLine(baseline)
			} else {
				metric.SetBaseLine(nil)
			}
		}

		if v, ok := m[string(Provider)]; ok {
			provider, err := expandProvider(v)
			if err != nil {
				return nil, err
			}
			if provider != nil {
				metric.SetProvider(provider)
			} else {
				metric.SetProvider(nil)
			}
		}
		metrics = append(metrics, metric)
	}
	return metrics, nil
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
	if list == nil || len(list) == 0 {
		return nil, nil
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

	if v, ok := m[string(Job)]; ok {
		job, err := expandJob(v)
		if err != nil {
			return nil, err
		}
		if job != nil {
			provider.SetJob(job)
		} else {
			provider.SetJob(nil)
		}
	}
	return provider, nil
}

func expandDatadog(data interface{}) (*oceancd.DataDogProvider, error) {

	datadog := &oceancd.DataDogProvider{}
	list := data.([]interface{})
	if list == nil || len(list) == 0 {
		return nil, nil
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
	if list == nil || len(list) == 0 {
		return nil, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Profile)].(string); ok && v != "" {
		newRelic.SetProfile(spotinst.String(v))
	} else {
		newRelic.SetProfile(nil)
	}

	if v, ok := m[string(NewRelicQuery)].(string); ok && v != "" {
		newRelic.SetQuery(spotinst.String(v))
	}
	return newRelic, nil
}

func expandPrometheus(data interface{}) (*oceancd.PrometheusProvider, error) {

	prometheus := &oceancd.PrometheusProvider{}
	list := data.([]interface{})
	if list == nil || len(list) == 0 {
		return nil, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(PrometheusQuery)].(string); ok && v != "" {
		prometheus.SetQuery(spotinst.String(v))
	}

	return prometheus, nil
}

func expandProvider(data interface{}) (*oceancd.Provider, error) {

	provider := &oceancd.Provider{}
	list := data.([]interface{})
	if list == nil || len(list) == 0 {
		return nil, nil
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
	if list == nil || len(list) == 0 {
		return nil, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(PipelineName)].(string); ok && v != "" {
		jenkins.SetPipelineName(spotinst.String(v))
	}

	if v, ok := m[string(Timeout)].(string); ok && v != "" {
		jenkins.SetTimeout(spotinst.String(v))
	}

	if v, ok := m[string(JenkinsInterval)].(string); ok && v != "" {
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
	list := data.(*schema.Set).List()
	parameters := make([]*oceancd.Parameters, 0, len(list))

	for _, v := range list {
		m := v.(map[string]interface{})
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

func expandWeb(data interface{}) (*oceancd.Web, error) {
	web := &oceancd.Web{}
	list := data.([]interface{})
	if list == nil || len(list) == 0 {
		return nil, nil
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
	list := data.(*schema.Set).List()
	headers := make([]*oceancd.Headers, 0, len(list))

	for _, v := range list {
		m := v.(map[string]interface{})
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

func expandCloudWatch(data interface{}) (*oceancd.CloudWatchProvider, error) {
	cloudWatch := &oceancd.CloudWatchProvider{}
	list := data.([]interface{})
	if list == nil || len(list) == 0 {
		return nil, nil
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
	list := data.(*schema.Set).List()
	metricDataQueries := make([]*oceancd.MetricDataQueries, 0, len(list))

	for _, v := range list {
		m := v.(map[string]interface{})
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

func expandMetricStats(data interface{}) (*oceancd.MetricStat, error) {
	metricStat := &oceancd.MetricStat{}
	list := data.([]interface{})
	if list == nil || len(list) == 0 {
		return nil, nil
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
	if list == nil || len(list) == 0 {
		return nil, nil
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
	list := data.(*schema.Set).List()
	dimensions := make([]*oceancd.Dimensions, 0, len(list))

	for _, v := range list {
		m := v.(map[string]interface{})
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

func expandJob(data interface{}) (*oceancd.Job, error) {
	job := &oceancd.Job{}
	list := data.([]interface{})
	if list == nil || len(list) == 0 {
		return nil, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Spec)]; ok {
		spec, err := expandSpec(v)
		if err != nil {
			return nil, err
		}
		if spec != nil {
			job.SetSpec(spec)
		} else {
			job.SetSpec(nil)
		}
	}
	return job, nil
}

func expandSpec(data interface{}) (*oceancd.Spec, error) {
	spec := &oceancd.Spec{}
	list := data.([]interface{})
	if list == nil || len(list) == 0 {
		return nil, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(JobTemplate)]; ok {
		template, err := expandTemplate(v)
		if err != nil {
			return nil, err
		}
		if spec != nil {
			spec.SetTemplate(template)
		} else {
			spec.SetTemplate(nil)
		}
	}

	if v, ok := m[string(BackoffLimit)].(int); ok {
		if v == -1 {
			spec.SetBackoffLimit(nil)
		} else {
			spec.SetBackoffLimit(spotinst.Int(v))
		}
	}
	return spec, nil
}

func expandTemplate(data interface{}) (*oceancd.Template, error) {
	template := &oceancd.Template{}
	list := data.([]interface{})
	if list == nil || len(list) == 0 {
		return nil, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(JobTemplate)]; ok {
		templateSpec, err := expandTemplateSpec(v)
		if err != nil {
			return nil, err
		}
		if template != nil {
			template.SetSpec(templateSpec)
		} else {
			template.SetSpec(nil)
		}
	}
	return template, nil
}

func expandTemplateSpec(data interface{}) (*oceancd.TemplateSpec, error) {
	templateSpec := &oceancd.TemplateSpec{}
	list := data.([]interface{})
	if list == nil || len(list) == 0 {
		return nil, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(RestartPolicy)].(string); ok && v != "" {
		templateSpec.SetRestartPolicy(spotinst.String(v))
	}

	if v, ok := m[string(JobTemplate)]; ok {
		containers, err := expandContainers(v)
		if err != nil {
			return nil, err
		}
		if templateSpec != nil {
			templateSpec.SetContainers(containers)
		} else {
			templateSpec.SetContainers(nil)
		}
	}
	return templateSpec, nil
}

func expandContainers(data interface{}) ([]*oceancd.Containers, error) {
	list := data.(*schema.Set).List()
	containers := make([]*oceancd.Containers, 0, len(list))

	for _, v := range list {
		m := v.(map[string]interface{})
		container := &oceancd.Containers{}

		if v, ok := m[string(Image)].(string); ok && v != "" {
			container.SetImage(spotinst.String(v))
		}

		if v, ok := m[string(ContinerName)].(string); ok && v != "" {
			container.SetName(spotinst.String(v))
		}

		if v, ok := m[string(Command)]; ok && v != nil {
			command, err := expandCommand(v)
			if err != nil {
				return nil, err
			}
			if command != nil {
				container.SetCommand(command)
			} else {
				container.SetCommand(nil)
			}
		}

		containers = append(containers, container)
	}
	return containers, nil

}

func expandCommand(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if commands, ok := v.(string); ok && commands != "" {
			result = append(result, commands)
		}
	}
	return result, nil
}

func flattenMetrics(metrics []*oceancd.Metrics) []interface{} {
	m := make([]interface{}, 0, len(metrics))

	for _, metric := range metrics {
		result := make(map[string]interface{})
		failureLimitValue := spotinst.Int(0)
		result[string(FailureLimit)] = failureLimitValue

		consecutiveErrorLimitValue := spotinst.Int(4)
		result[string(FailureLimit)] = consecutiveErrorLimitValue

		countValue := spotinst.Int(0)
		result[string(FailureLimit)] = countValue

		result[string(FailureCondition)] = spotinst.StringValue(metric.FailureCondition)
		result[string(InitialDelay)] = spotinst.StringValue(metric.InitialDelay)
		result[string(Interval)] = spotinst.StringValue(metric.Interval)
		result[string(MetricsName)] = spotinst.StringValue(metric.Name)
		result[string(SuccessCondition)] = spotinst.StringValue(metric.SuccessCondition)

		result[string(DryRun)] = spotinst.BoolValue(metric.DryRun)

		if metric.Provider != nil {
			result[string(Provider)] = flattenProvider(metric.Provider)
		}

		if metric.Baseline != nil {
			result[string(BaseLine)] = flattenBaseline(metric.Baseline)
		}

		if metric.FailureLimit != nil {
			result[string(FailureLimit)] = spotinst.IntValue(metric.FailureLimit)
		}

		if metric.Count != nil {
			result[string(Count)] = spotinst.IntValue(metric.Count)
		}

		if metric.ConsecutiveErrorLimit != nil {
			result[string(ConsecutiveErrorLimit)] = spotinst.IntValue(metric.ConsecutiveErrorLimit)
		}

		m = append(m, result)

	}
	return m
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

	if provider.Job != nil {
		result[string(Job)] = flattenJob(provider.Job)
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
	result[string(JenkinsInterval)] = spotinst.StringValue(jenkins.Interval)
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

		m = append(m, result)
	}
	return m
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

		m = append(m, result)
	}
	return m
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
		m = append(m, result)
	}
	return m
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

		m = append(m, result)
	}
	return m
}

func flattenJob(job *oceancd.Job) []interface{} {
	result := make(map[string]interface{})

	if job.Spec != nil {
		result[string(Spec)] = flattenSpec(job.Spec)
	}

	return []interface{}{result}
}

func flattenSpec(spec *oceancd.Spec) []interface{} {
	result := make(map[string]interface{})
	value := spotinst.Int(-1)
	result[string(BackoffLimit)] = value

	if spec.Template != nil {
		result[string(JobTemplate)] = flattenTemplate(spec.Template)
	}

	if spec.BackoffLimit != nil {
		result[string(BackoffLimit)] = spotinst.IntValue(spec.BackoffLimit)
	}

	return []interface{}{result}
}

func flattenTemplate(template *oceancd.Template) []interface{} {
	result := make(map[string]interface{})

	if template.Spec != nil {
		result[string(TemplateSpec)] = flattenTemplateSpec(template.Spec)
	}

	return []interface{}{result}
}

func flattenTemplateSpec(templateSpec *oceancd.TemplateSpec) []interface{} {
	result := make(map[string]interface{})

	if templateSpec.Containers != nil {
		result[string(Containers)] = flattenContainers(templateSpec.Containers)
	}
	result[string(RestartPolicy)] = spotinst.StringValue(templateSpec.RestartPolicy)

	return []interface{}{result}
}

func flattenContainers(containers []*oceancd.Containers) []interface{} {
	m := make([]interface{}, 0, len(containers))

	for _, container := range containers {
		result := make(map[string]interface{})
		result[string(Image)] = spotinst.StringValue(container.Image)
		result[string(ContinerName)] = spotinst.StringValue(container.Name)

		if container.Command != nil {
			result[string(Command)] = spotinst.StringSlice(container.Command)
		}

		m = append(m, result)
	}
	return m
}
