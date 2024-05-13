package oceancd_strategy_canary

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Rolling] = commons.NewGenericField(
		commons.OceanCDStrategyCanary,
		Rolling,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Name): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(Steps): {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Name): {
									Type:     schema.TypeString,
									Optional: true,
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
								string(Verification): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(TemplateNames): {
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

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			strategyWrapper := resourceObject.(*commons.OceanCDStrategyWrapper)
			strategy := strategyWrapper.GetStrategy()
			var result []interface{} = nil

			if strategy != nil && strategy.Canary != nil {
				result = flattenRolling(strategy.Rolling)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Rolling), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Rolling), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			strategyWrapper := resourceObject.(*commons.OceanCDStrategyWrapper)
			strategy := strategyWrapper.GetStrategy()
			var value *oceancd.Rolling = nil

			if v, ok := resourceData.GetOkExists(string(Rolling)); ok {
				if rolling, err := expandRolling(v); err != nil {
					return err
				} else {
					value = rolling
				}
			}
			strategy.SetRolling(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			strategyWrapper := resourceObject.(*commons.OceanCDStrategyWrapper)
			strategy := strategyWrapper.GetStrategy()
			var value *oceancd.Rolling = nil
			if v, ok := resourceData.GetOkExists(string(Rolling)); ok {
				if rolling, err := expandRolling(v); err != nil {
					return err
				} else {
					value = rolling
				}
			}
			strategy.SetRolling(value)
			return nil
		},
		nil,
	)
}

func expandRolling(data interface{}) (*oceancd.Rolling, error) {
	if list := data.([]interface{}); len(list) > 0 {
		rolling := &oceancd.Rolling{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(Steps)]; ok {
				steps, err := expandSteps(v)
				if err != nil {
					return nil, err
				}
				if steps != nil {
					rolling.SetSteps(steps)
				} else {
					rolling.SetSteps(nil)
				}
			}
		}
		return rolling, nil
	}
	return nil, nil
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

func expandSteps(data interface{}) ([]*oceancd.RollingSteps, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		steps := make([]*oceancd.RollingSteps, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			step := &oceancd.RollingSteps{}

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

func flattenRolling(rolling *oceancd.Rolling) []interface{} {
	result := make(map[string]interface{})

	if rolling.Steps != nil {
		result[string(Steps)] = flattenSteps(rolling.Steps)
	}
	return []interface{}{result}
}

func flattenSteps(steps []*oceancd.RollingSteps) []interface{} {
	m := make([]interface{}, 0, len(steps))
	for _, step := range steps {
		result := make(map[string]interface{})
		result[string(Name)] = spotinst.StringValue(step.Name)

		if step.Pause != nil {
			result[string(Pause)] = flattenPause(step.Pause)
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

func flattenVerification(verification *oceancd.Verification) []interface{} {
	result := make(map[string]interface{})

	if verification.TemplateNames != nil {
		result[string(TemplateNames)] = spotinst.StringSlice(verification.TemplateNames)
	}
	return []interface{}{result}
}
