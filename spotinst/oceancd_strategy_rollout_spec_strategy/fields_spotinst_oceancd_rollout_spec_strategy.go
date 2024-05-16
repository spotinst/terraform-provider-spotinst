package oceancd_strategy_canary_strategy

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Strategy] = commons.NewGenericField(
		commons.OceanCDRolloutSpecStrategy,
		Strategy,
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
					string(Args): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(ArgName): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(ArgValue): {
									Type:     schema.TypeString,
									Optional: true,
								},

								string(ValueFrom): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(FieldRef): {
												Type:     schema.TypeList,
												Required: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														string(FieldPath): {
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
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			var result []interface{} = nil

			if rolloutSpec != nil && rolloutSpec.Strategy != nil {
				result = flattenStrategy(rolloutSpec.Strategy)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Strategy), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Strategy), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			var value *oceancd.RolloutSpecStrategy = nil

			if v, ok := resourceData.GetOkExists(string(Strategy)); ok {
				if strategy, err := expandStrategy(v); err != nil {
					return err
				} else {
					value = strategy
				}
			}
			rolloutSpec.SetStrategy(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			var value *oceancd.RolloutSpecStrategy = nil
			if v, ok := resourceData.GetOkExists(string(Strategy)); ok {
				if strategy, err := expandStrategy(v); err != nil {
					return err
				} else {
					value = strategy
				}
			}
			rolloutSpec.SetStrategy(value)
			return nil
		},
		nil,
	)
}

func expandStrategy(data interface{}) (*oceancd.RolloutSpecStrategy, error) {
	if list := data.([]interface{}); len(list) > 0 {
		strategy := &oceancd.RolloutSpecStrategy{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(Name)].(string); ok && v != "" {
				strategy.SetName(spotinst.String(v))
			} else {
				strategy.SetName(nil)
			}

			if v, ok := m[string(Args)]; ok && v != nil {

				args, err := expandArgs(v)
				if err != nil {
					return nil, err
				}
				if args != nil {
					strategy.SetArgs(args)
				} else {
					strategy.SetArgs(nil)
				}
			}
		}
		return strategy, nil
	}
	return nil, nil
}

func flattenStrategy(strategy *oceancd.RolloutSpecStrategy) []interface{} {
	var response []interface{}

	if strategy != nil {
		result := make(map[string]interface{})

		result[string(Name)] = spotinst.StringValue(strategy.Name)

		if strategy.Args != nil {
			result[string(Args)] = flattenArgs(strategy.Args)
		}

		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}

func expandArgs(data interface{}) ([]*oceancd.RolloutSpecArgs, error) {
	list := data.(*schema.Set).List()
	args := make([]*oceancd.RolloutSpecArgs, 0, len(list))

	for _, v := range list {
		m := v.(map[string]interface{})
		arg := &oceancd.RolloutSpecArgs{}

		if v, ok := m[string(ArgName)].(string); ok && v != "" {
			arg.SetName(spotinst.String(v))
		}
		if v, ok := m[string(ArgValue)].(string); ok && v != "" {
			arg.SetValue(spotinst.String(v))
		}

		if v, ok := m[string(ValueFrom)]; ok && v != nil {
			valueFrom, err := expandValueFrom(v)
			if err != nil {
				return nil, err
			}
			if valueFrom != nil {
				arg.SetValueFrom(valueFrom)
			} else {
				arg.SetValueFrom(nil)
			}
		}

		args = append(args, arg)
	}
	return args, nil
}

func expandValueFrom(data interface{}) (*oceancd.RolloutSpecValueFrom, error) {
	valueFrom := &oceancd.RolloutSpecValueFrom{}
	list := data.([]interface{})
	if list == nil || len(list) == 0 {
		return nil, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(FieldRef)]; ok && v != nil {
		fieldRef, err := expandFieldRef(v)
		if err != nil {
			return nil, err
		}
		if fieldRef != nil {
			valueFrom.SetFieldRef(fieldRef)
		} else {
			valueFrom.SetFieldRef(nil)
		}
	}
	return valueFrom, nil
}

func expandFieldRef(data interface{}) (*oceancd.FieldRef, error) {

	fieldRef := &oceancd.FieldRef{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return fieldRef, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(FieldPath)].(string); ok && v != "" {
		fieldRef.SetFieldPath(spotinst.String(v))
	}

	return fieldRef, nil
}

func flattenFieldRef(fieldRef *oceancd.FieldRef) []interface{} {
	result := make(map[string]interface{})
	result[string(FieldPath)] = spotinst.StringValue(fieldRef.FieldPath)

	return []interface{}{result}
}

func flattenValueFrom(valueFrom *oceancd.RolloutSpecValueFrom) []interface{} {
	result := make(map[string]interface{})

	if valueFrom.FieldRef != nil {
		result[string(FieldRef)] = flattenFieldRef(valueFrom.FieldRef)
	}
	return []interface{}{result}
}

func flattenArgs(args []*oceancd.RolloutSpecArgs) []interface{} {
	result := make([]interface{}, 0, len(args))

	for _, arg := range args {
		m := make(map[string]interface{})

		m[string(ArgName)] = spotinst.StringValue(arg.Name)
		m[string(ArgValue)] = spotinst.StringValue(arg.Value)

		if arg.ValueFrom != nil {
			m[string(ValueFrom)] = flattenValueFrom(arg.ValueFrom)
		}
		result = append(result, m)
	}
	return result
}
