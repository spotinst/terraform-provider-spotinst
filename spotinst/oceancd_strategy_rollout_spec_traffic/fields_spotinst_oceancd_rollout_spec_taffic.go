package oceancd_strategy_canary_traffic

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Traffic] = commons.NewGenericField(
		commons.OceanCDRolloutSpecTraffio,
		Traffic,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(CanaryService): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(StableService): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(Alb): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(AlbAnnotationPrefix): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(AlbIngress): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(AlbRootService): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(ServicePort): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(StickinessConfig): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(StickinessDuration): {
												Type:     schema.TypeString,
												Optional: true,
											},
											string(StickinessEnabled): {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},
							},
						},
					},
					string(Ambassador): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Mappings): {
									Type:     schema.TypeList,
									Required: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
							},
						},
					},
					string(Istio): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(DestinationRule): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(CanarySubsetRule): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(DestinationRuleName): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(StableSubsetName): {
												Type:     schema.TypeString,
												Required: true,
											},
										},
									},
								},
								string(VirtualServices): {
									Type:     schema.TypeSet,
									Required: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(VirtualServiceName): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(VirtualServiceRoutes): {
												Type:     schema.TypeList,
												Optional: true,
												Elem:     &schema.Schema{Type: schema.TypeString},
											},
											string(TlsRoutes): {
												Type:     schema.TypeSet,
												Optional: true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														string(Port): {
															Type:     schema.TypeInt,
															Optional: true,
															Default:  -1,
														},
														string(SniHosts): {
															Type:     schema.TypeInt,
															Optional: true,
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
					string(Nginx): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(NginxAnnotationPrefix): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(StableIngress): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(AdditionalIngressAnnotation): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(CanaryByHeader): {
												Type:     schema.TypeString,
												Optional: true,
											},
											string(Key1): {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},
							},
						},
					},
					string(PingPong): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(PingService): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(PongService): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},
					string(Smi): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(SmiRootService): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(TrafficSplitName): {
									Type:     schema.TypeString,
									Required: true,
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
				if err := resourceData.Set(string(Traffic), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Traffic), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			var value *oceancd.Traffic = nil

			if v, ok := resourceData.GetOkExists(string(Traffic)); ok {
				if traffic, err := expandTraffic(v); err != nil {
					return err
				} else {
					value = traffic
				}
			}
			rolloutSpec.SetTraffic(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			var value *oceancd.Traffic = nil
			if v, ok := resourceData.GetOkExists(string(Traffic)); ok {
				if traffic, err := expandTraffic(v); err != nil {
					return err
				} else {
					value = traffic
				}
			}
			rolloutSpec.SetTraffic(value)
			return nil
		},
		nil,
	)
}

func expandTraffic(data interface{}) (*oceancd.Traffic, error) {
	if list := data.([]interface{}); len(list) > 0 {
		traffic := &oceancd.Traffic{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(StableService)].(string); ok && v != "" {
				traffic.SetStableService(spotinst.String(v))
			} else {
				traffic.SetStableService(nil)
			}

			if v, ok := m[string(CanaryService)].(string); ok && v != "" {
				traffic.SetCanaryService(spotinst.String(v))
			} else {
				traffic.SetCanaryService(nil)
			}

			if v, ok := m[string(Alb)]; ok && v != nil {

				alb, err := expandAlb(v)
				if err != nil {
					return nil, err
				}
				if alb != nil {
					traffic.SetAlb(alb)
				} else {
					traffic.SetAlb(nil)
				}
			}

			if v, ok := m[string(Ambassador)]; ok && v != nil {

				args, err := expandAmbassador(v)
				if err != nil {
					return nil, err
				}
				if args != nil {
					traffic.SetAmbassador(args)
				} else {
					traffic.SetAmbassador(nil)
				}
			}

			if v, ok := m[string(Istio)]; ok && v != nil {

				args, err := expandIstio(v)
				if err != nil {
					return nil, err
				}
				if args != nil {
					traffic.SetIstio(args)
				} else {
					traffic.SetIstio(nil)
				}
			}

			if v, ok := m[string(Nginx)]; ok && v != nil {

				args, err := expandNginx(v)
				if err != nil {
					return nil, err
				}
				if args != nil {
					traffic.SetNginx(args)
				} else {
					traffic.SetNginx(nil)
				}
			}

			if v, ok := m[string(PingPong)]; ok && v != nil {

				args, err := expandPingPong(v)
				if err != nil {
					return nil, err
				}
				if args != nil {
					traffic.SetPingPong(args)
				} else {
					traffic.SetPingPong(nil)
				}
			}

			if v, ok := m[string(Smi)]; ok && v != nil {

				args, err := expandSmi(v)
				if err != nil {
					return nil, err
				}
				if args != nil {
					traffic.SetSmi(args)
				} else {
					traffic.SetSmi(nil)
				}
			}

		}
		return traffic, nil
	}
	return nil, nil
}

func expandAlb(data interface{}) (*oceancd.Alb, error) {
	if list := data.([]interface{}); len(list) > 0 {
		traffic := &oceancd.Alb{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(AlbAnnotationPrefix)].(string); ok && v != "" {
				traffic.SetAnnotationPrefix(spotinst.String(v))
			} else {
				traffic.SetAnnotationPrefix(nil)
			}

			if v, ok := m[string(AlbIngress)].(string); ok && v != "" {
				traffic.SetIngress(spotinst.String(v))
			}

			if v, ok := m[string(AlbRootService)].(string); ok && v != "" {
				traffic.SetRootService(spotinst.String(v))
			}

			if v, ok := m[string(ServicePort)].(int); ok {
				if v == -1 {
					traffic.SetServicePort(nil)
				} else {
					traffic.SetServicePort(spotinst.Int(v))
				}
			}

			if v, ok := m[string(StickinessConfig)]; ok && v != nil {

				stickinessConfig, err := expandStickinessConfig(v)
				if err != nil {
					return nil, err
				}
				if stickinessConfig != nil {
					traffic.SetStickinessConfig(stickinessConfig)
				} else {
					traffic.SetStickinessConfig(nil)
				}
			}
		}
		return traffic, nil
	}
	return nil, nil
}

func expandAmbassador(data interface{}) (*oceancd.Ambassador, error) {
	if list := data.([]interface{}); len(list) > 0 {
		ambassador := &oceancd.Ambassador{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

		}
		return ambassador, nil
	}
	return nil, nil
}

func expandStickinessConfig(data interface{}) (*oceancd.StickinessConfig, error) {
	if list := data.([]interface{}); len(list) > 0 {
		stickiness := &oceancd.StickinessConfig{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(StickinessDuration)].(int); ok {
				if v == -1 {
					stickiness.SetDurationSeconds(nil)
				} else {
					stickiness.SetDurationSeconds(spotinst.Int(v))
				}
			}

			if v, ok := m[string(StickinessEnabled)].(bool); ok {
				stickiness.SetEnabled(spotinst.Bool(v))
			}

		}
		return stickiness, nil
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
