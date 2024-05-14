package oceancd_verification_template_args

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
}

func expandCanary(data interface{}) (*oceancd.Canary, error) {
	if list := data.([]interface{}); len(list) > 0 {
		canary := &oceancd.Canary{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(BackgroundVerification)]; ok && v != nil {

				backVerification, err := expandBackgroundVerification(v)
				if err != nil {
					return nil, err
				}
				if backVerification != nil {
					canary.SetBackgroundVerification(backVerification)
				} else {
					canary.SetBackgroundVerification(nil)
				}
			}

			if v, ok := m[string(Steps)]; ok {
				steps, err := expandSteps(v)
				if err != nil {
					return nil, err
				}
				if steps != nil {
					canary.SetSteps(steps)
				} else {
					canary.SetSteps(nil)
				}
			}
		}
		return canary, nil
	}
	return nil, nil
}

func expandBackgroundVerification(data interface{}) (*oceancd.BackgroundVerification, error) {
	backVerification := &oceancd.BackgroundVerification{}
	list := data.([]interface{})

	if list == nil || len(list) == 0 {
		return nil, nil
	}

	m := list[0].(map[string]interface{})

	if v, ok := m[string(TemplateNames)]; ok && v != nil {

		templateNames, err := expandTemplateNames(v)
		if err != nil {
			return nil, err
		}
		if templateNames != nil {
			backVerification.SetTemplateNames(templateNames)
		} else {
			backVerification.SetTemplateNames(nil)
		}
	}

	return backVerification, nil
}

func expandTemplateNames(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if templateNames, ok := v.(string); ok && templateNames != "" {
			result = append(result, templateNames)
		}
	}
	return result, nil
}

func expandSteps(data interface{}) ([]*oceancd.CanarySteps, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		steps := make([]*oceancd.CanarySteps, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			step := &oceancd.CanarySteps{}

			if v, ok := m[string(Name)].(string); ok && v != "" {
				step.SetName(spotinst.String(v))
			}

			if v, ok := m[string(Pause)]; ok {
				pause, err := expandPause(v)
				if err != nil {
					return nil, err
				}
				if pause != nil {
					step.SetPause(pause)
				} else {
					step.SetPause(nil)
				}
			}

			if v, ok := m[string(Verification)]; ok {
				verification, err := expandVerification(v)
				if err != nil {
					return nil, err
				}
				if verification != nil {
					step.SetVerification(verification)
				} else {
					step.SetVerification(nil)
				}
			}

			if v, ok := m[string(SetCanaryScale)]; ok {
				setCanaryScale, err := expandSetCanaryScale(v)
				if err != nil {
					return nil, err
				}
				if setCanaryScale != nil {
					step.SetSetCanaryScale(setCanaryScale)
				} else {
					step.SetSetCanaryScale(nil)
				}
			}

			if v, ok := m[string(SetHeaderRoute)]; ok {
				headerRoute, err := expandSetHeaderRoute(v)
				if err != nil {
					return nil, err
				}
				if headerRoute != nil {
					step.SetSetHeaderRoute(headerRoute)
				} else {
					step.SetSetHeaderRoute(nil)
				}
			}

			if v, ok := m[string(SetWeight)].(int); ok {
				if v == -1 {
					step.SetSetWeight(nil)
				} else {
					step.SetSetWeight(spotinst.Int(v))
				}
			}

			steps = append(steps, step)
		}
		return steps, nil
	}
	return nil, nil
}

func expandPause(data interface{}) (*oceancd.Pause, error) {

	pause := &oceancd.Pause{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return pause, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Duration)].(string); ok && v != "" {
		pause.SetDuration(spotinst.String(v))
	} else {
		pause.SetDuration(nil)
	}

	return pause, nil
}

func expandSetCanaryScale(data interface{}) (*oceancd.SetCanaryScale, error) {

	setCanaryScale := &oceancd.SetCanaryScale{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return setCanaryScale, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(MatchTrafficWeight)].(bool); ok {
		setCanaryScale.SetMatchTrafficWeight(spotinst.Bool(v))
	}

	if v, ok := m[string(Replicas)].(int); ok {
		if v == -1 {
			setCanaryScale.SetReplicas(nil)
		} else {
			setCanaryScale.SetReplicas(spotinst.Int(v))
		}
	}

	if v, ok := m[string(Weight)].(int); ok {
		if v == -1 {
			setCanaryScale.SetWeight(nil)
		} else {
			setCanaryScale.SetWeight(spotinst.Int(v))
		}
	}
	return setCanaryScale, nil
}

func expandVerification(data interface{}) (*oceancd.Verification, error) {
	verification := &oceancd.Verification{}
	list := data.([]interface{})
	if list == nil || len(list) == 0 {
		return nil, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(TemplateNames)]; ok && v != nil {
		templateNames, err := expandTemplateNames(v)
		if err != nil {
			return nil, err
		}
		if templateNames != nil {
			verification.SetTemplateNames(templateNames)
		} else {
			verification.SetTemplateNames(nil)
		}
	}
	return verification, nil
}

func expandSetHeaderRoute(data interface{}) (*oceancd.SetHeaderRoute, error) {

	setHeaderRoute := &oceancd.SetHeaderRoute{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return setHeaderRoute, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Name)].(string); ok && v != "" {
		setHeaderRoute.SetName(spotinst.String(v))
	} else {
		setHeaderRoute.SetName(nil)
	}

	if v, ok := m[string(Match)]; ok {
		match, err := expandMatch(v)
		if err != nil {
			return nil, err
		}
		if match != nil {
			setHeaderRoute.SetMatch(match)
		} else {
			setHeaderRoute.SetMatch(nil)
		}
	}
	return setHeaderRoute, nil
}

func expandMatch(data interface{}) ([]*oceancd.Match, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		matches := make([]*oceancd.Match, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			match := &oceancd.Match{}

			if v, ok := m[string(HeaderName)].(string); ok && v != "" {
				match.SetHeaderName(spotinst.String(v))
			}

			if v, ok := m[string(HeaderValue)]; ok {
				headerValue, err := expandHeaderValue(v)
				if err != nil {
					return nil, err
				}
				if headerValue != nil {
					match.SetHeaderValue(headerValue)
				} else {
					match.SetHeaderValue(nil)
				}
			}
			matches = append(matches, match)
		}
		return matches, nil
	}
	return nil, nil
}

func expandHeaderValue(data interface{}) (*oceancd.HeaderValue, error) {

	headerValue := &oceancd.HeaderValue{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return headerValue, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Exact)].(string); ok && v != "" {
		headerValue.SetExact(spotinst.String(v))
	} else {
		headerValue.SetExact(nil)
	}

	if v, ok := m[string(Regex)].(string); ok && v != "" {
		headerValue.SetRegex(spotinst.String(v))
	} else {
		headerValue.SetRegex(nil)
	}

	if v, ok := m[string(Prefix)].(string); ok && v != "" {
		headerValue.SetPrefix(spotinst.String(v))
	} else {
		headerValue.SetPrefix(nil)
	}
	return headerValue, nil
}

func flattenCanary(canary *oceancd.Canary) []interface{} {
	result := make(map[string]interface{})

	if canary.BackgroundVerification != nil {
		result[string(BackgroundVerification)] = flattenBackgroundVerification(canary.BackgroundVerification)
	}

	if canary.Steps != nil {
		result[string(Steps)] = flattenSteps(canary.Steps)
	}
	return []interface{}{result}
}

func flattenBackgroundVerification(backgroundVerification *oceancd.BackgroundVerification) []interface{} {
	result := make(map[string]interface{})

	if backgroundVerification.TemplateNames != nil {
		result[string(TemplateNames)] = spotinst.StringSlice(backgroundVerification.TemplateNames)
	}
	return []interface{}{result}
}

func flattenSteps(steps []*oceancd.CanarySteps) []interface{} {
	m := make([]interface{}, 0, len(steps))
	for _, step := range steps {
		result := make(map[string]interface{})
		value := spotinst.Int(-1)
		result[string(SetWeight)] = value

		if step.SetWeight != nil {
			result[string(SetWeight)] = spotinst.IntValue(step.SetWeight)
		}
		result[string(Name)] = spotinst.StringValue(step.Name)

		if step.Pause != nil {
			result[string(Pause)] = flattenPause(step.Pause)
		}

		if step.SetCanaryScale != nil {
			result[string(SetCanaryScale)] = flattenSetCanaryScale(step.SetCanaryScale)
		}

		if step.SetHeaderRoute != nil {
			result[string(SetHeaderRoute)] = flattenSetHeaderRoute(step.SetHeaderRoute)
		}

		if step.Verification != nil {
			result[string(Verification)] = flattenVerification(step.Verification)
		}
	}
	return []interface{}{m}
}

func flattenPause(pause *oceancd.Pause) []interface{} {
	result := make(map[string]interface{})
	result[string(Duration)] = spotinst.StringValue(pause.Duration)
	return []interface{}{result}
}

func flattenSetCanaryScale(scale *oceancd.SetCanaryScale) []interface{} {
	result := make(map[string]interface{})
	value := spotinst.Int(-1)
	result[string(Replicas)] = value
	result[string(Weight)] = value

	result[string(MatchTrafficWeight)] = spotinst.BoolValue(scale.MatchTrafficWeight)

	if scale.Replicas != nil {
		result[string(Replicas)] = spotinst.IntValue(scale.Replicas)
	}
	if scale.Weight != nil {
		result[string(Weight)] = spotinst.IntValue(scale.Weight)
	}
	return []interface{}{result}
}

func flattenVerification(verification *oceancd.Verification) []interface{} {
	result := make(map[string]interface{})

	if verification.TemplateNames != nil {
		result[string(TemplateNames)] = spotinst.StringSlice(verification.TemplateNames)
	}
	return []interface{}{result}
}

func flattenSetHeaderRoute(setHeaderRoute *oceancd.SetHeaderRoute) []interface{} {
	result := make(map[string]interface{})

	result[string(Name)] = spotinst.StringValue(setHeaderRoute.Name)

	if setHeaderRoute.Match != nil {
		result[string(Match)] = flattenMatch(setHeaderRoute.Match)
	}
	return []interface{}{result}
}

func flattenMatch(matches []*oceancd.Match) []interface{} {
	m := make([]interface{}, 0, len(matches))
	for _, match := range matches {
		result := make(map[string]interface{})
		result[string(HeaderName)] = spotinst.StringValue(match.HeaderName)

		if match.HeaderValue != nil {
			result[string(HeaderValue)] = flattenHeaderValue(match.HeaderValue)
		}
	}
	return []interface{}{m}
}

func flattenHeaderValue(headerValue *oceancd.HeaderValue) []interface{} {
	result := make(map[string]interface{})

	result[string(Exact)] = spotinst.StringValue(headerValue.Exact)
	result[string(Prefix)] = spotinst.StringValue(headerValue.Prefix)
	result[string(Regex)] = spotinst.StringValue(headerValue.Regex)
	return []interface{}{result}
}
