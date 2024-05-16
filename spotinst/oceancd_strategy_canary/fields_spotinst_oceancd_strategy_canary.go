package oceancd_strategy_canary

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanCDStrategy,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			strategyWrapper := resourceObject.(*commons.OceanCDStrategyWrapper)
			strategy := strategyWrapper.GetStrategy()
			var value *string = nil
			if strategy.Name != nil {
				value = strategy.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			strategyWrapper := resourceObject.(*commons.OceanCDStrategyWrapper)
			strategy := strategyWrapper.GetStrategy()
			if value, ok := resourceData.Get(string(Name)).(string); ok && value != "" {
				strategy.SetName(spotinst.String(value))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[Canary] = commons.NewGenericField(
		commons.OceanCDStrategyCanary,
		Canary,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(BackgroundVerification): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(BGTemplateNames): {
									Type:     schema.TypeList,
									Required: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
							},
						},
					},
					string(Steps): {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(StepName): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(SetWeight): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  -1,
								},
								string(Pause): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Duration): {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},
								string(SetCanaryScale): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(MatchTrafficWeight): {
												Type:     schema.TypeBool,
												Optional: true,
											},
											string(Replicas): {
												Type:     schema.TypeInt,
												Optional: true,
												Default:  -1,
											},
											string(Weight): {
												Type:     schema.TypeInt,
												Optional: true,
												Default:  -1,
											},
										},
									},
								},
								string(SetHeaderRoute): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(HeaderRouteName): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(Match): {
												Type:     schema.TypeList,
												Required: true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														string(HeaderName): {
															Type:     schema.TypeString,
															Required: true,
														},
														string(HeaderValue): {
															Type:     schema.TypeList,
															Required: true,
															MaxItems: 1,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	string(Exact): {
																		Type:     schema.TypeString,
																		Optional: true,
																	},
																	string(Prefix): {
																		Type:     schema.TypeString,
																		Optional: true,
																	},
																	string(Regex): {
																		Type:     schema.TypeString,
																		Optional: true,
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
								string(Verification): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(TemplateNames): {
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

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			strategyWrapper := resourceObject.(*commons.OceanCDStrategyWrapper)
			strategy := strategyWrapper.GetStrategy()
			var result []interface{} = nil

			if strategy != nil && strategy.Canary != nil {
				result = flattenCanary(strategy.Canary)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Canary), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Canary), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			strategyWrapper := resourceObject.(*commons.OceanCDStrategyWrapper)
			strategy := strategyWrapper.GetStrategy()
			var value *oceancd.Canary = nil

			if v, ok := resourceData.GetOkExists(string(Canary)); ok {
				if canary, err := expandCanary(v); err != nil {
					return err
				} else {
					value = canary
				}
			}
			strategy.SetCanary(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			strategyWrapper := resourceObject.(*commons.OceanCDStrategyWrapper)
			strategy := strategyWrapper.GetStrategy()
			var value *oceancd.Canary = nil
			if v, ok := resourceData.GetOkExists(string(Canary)); ok {
				if canary, err := expandCanary(v); err != nil {
					return err
				} else {
					value = canary
				}
			}
			strategy.SetCanary(value)
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

	if v, ok := m[string(BGTemplateNames)]; ok && v != nil {

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

			if v, ok := m[string(StepName)].(string); ok && v != "" {
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

	if v, ok := m[string(HeaderRouteName)].(string); ok && v != "" {
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
	var response []interface{}

	if canary != nil {
		result := make(map[string]interface{})

		if canary.BackgroundVerification != nil {
			result[string(BackgroundVerification)] = flattenBackgroundVerification(canary.BackgroundVerification)
		}

		if canary.Steps != nil {
			result[string(Steps)] = flattenSteps(canary.Steps)
		}
		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}

func flattenBackgroundVerification(backgroundVerification *oceancd.BackgroundVerification) []interface{} {
	result := make(map[string]interface{})

	if backgroundVerification.TemplateNames != nil {
		result[string(BGTemplateNames)] = spotinst.StringSlice(backgroundVerification.TemplateNames)
	}
	return []interface{}{result}
}

func flattenSteps(steps []*oceancd.CanarySteps) []interface{} {
	response := make([]interface{}, 0, len(steps))
	for _, step := range steps {
		result := make(map[string]interface{})
		value := spotinst.Int(-1)
		result[string(SetWeight)] = value

		if step.SetWeight != nil {
			result[string(SetWeight)] = spotinst.IntValue(step.SetWeight)
		}
		result[string(StepName)] = spotinst.StringValue(step.Name)

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
		response = append(response, result)
	}
	return response
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

	result[string(HeaderRouteName)] = spotinst.StringValue(setHeaderRoute.Name)

	if setHeaderRoute.Match != nil {
		result[string(Match)] = flattenMatch(setHeaderRoute.Match)
	}
	return []interface{}{result}
}

func flattenMatch(matches []*oceancd.Match) []interface{} {
	response := make([]interface{}, 0, len(matches))
	for _, match := range matches {
		result := make(map[string]interface{})
		result[string(HeaderName)] = spotinst.StringValue(match.HeaderName)

		if match.HeaderValue != nil {
			result[string(HeaderValue)] = flattenHeaderValue(match.HeaderValue)
		}
		response = append(response, result)
	}
	return response
}

func flattenHeaderValue(headerValue *oceancd.HeaderValue) []interface{} {
	result := make(map[string]interface{})

	result[string(Exact)] = spotinst.StringValue(headerValue.Exact)
	result[string(Prefix)] = spotinst.StringValue(headerValue.Prefix)
	result[string(Regex)] = spotinst.StringValue(headerValue.Regex)
	return []interface{}{result}
}
