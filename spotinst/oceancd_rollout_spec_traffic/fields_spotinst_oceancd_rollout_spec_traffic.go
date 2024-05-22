package oceancd_rollout_spec_traffic

import (
	"fmt"

	"github.com/spotinst/spotinst-sdk-go/service/oceancd"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Traffic] = commons.NewGenericField(
		commons.OceanCDRolloutSpecTraffic,
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
									Type:     schema.TypeInt,
									Required: true,
								},
								string(StickinessConfig): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(StickinessDuration): {
												Type:     schema.TypeInt,
												Optional: true,
												Default:  -1,
											},
											string(StickinessEnabled): {
												Type:     schema.TypeBool,
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
											string(CanarySubsetName): {
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
															Type:     schema.TypeList,
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
									Optional: true,
								},
								string(TrafficSplitName): {
									Type:     schema.TypeString,
									Optional: true,
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

			if rolloutSpec != nil && rolloutSpec.Traffic != nil {
				result = flattenTraffic(rolloutSpec.Traffic)
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
		alb := &oceancd.Alb{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(AlbAnnotationPrefix)].(string); ok && v != "" {
				alb.SetAnnotationPrefix(spotinst.String(v))
			} else {
				alb.SetAnnotationPrefix(nil)
			}

			if v, ok := m[string(AlbIngress)].(string); ok && v != "" {
				alb.SetIngress(spotinst.String(v))
			}

			if v, ok := m[string(AlbRootService)].(string); ok && v != "" {
				alb.SetRootService(spotinst.String(v))
			}

			if v, ok := m[string(ServicePort)].(int); ok {
				if v == -1 {
					alb.SetServicePort(nil)
				} else {
					alb.SetServicePort(spotinst.Int(v))
				}
			}

			if v, ok := m[string(StickinessConfig)]; ok && v != nil {

				stickinessConfig, err := expandStickinessConfig(v)
				if err != nil {
					return nil, err
				}
				if stickinessConfig != nil {
					alb.SetStickinessConfig(stickinessConfig)
				} else {
					alb.SetStickinessConfig(nil)
				}
			}
		}
		return alb, nil
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

func expandAmbassador(data interface{}) (*oceancd.Ambassador, error) {
	if list := data.([]interface{}); len(list) > 0 {
		ambassador := &oceancd.Ambassador{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(Mappings)]; ok && v != nil {
				mapping, err := expandMappings(v)
				if err != nil {
					return nil, err
				}
				if mapping != nil {
					ambassador.SetMappings(mapping)
				} else {
					ambassador.SetMappings(nil)
				}
			}

		}
		return ambassador, nil
	}
	return nil, nil
}

func expandMappings(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if mappings, ok := v.(string); ok && mappings != "" {
			result = append(result, mappings)
		}
	}
	return result, nil
}

func expandIstio(data interface{}) (*oceancd.Istio, error) {
	if list := data.([]interface{}); len(list) > 0 {
		istio := &oceancd.Istio{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(DestinationRule)]; ok && v != nil {

				destinationRule, err := expandDestinationRule(v)
				if err != nil {
					return nil, err
				}
				if destinationRule != nil {
					istio.SetDestinationRule(destinationRule)
				} else {
					istio.SetDestinationRule(nil)
				}
			}

			if v, ok := m[string(VirtualServices)]; ok && v != nil {

				virtualServices, err := expandVirtualServices(v)
				if err != nil {
					return nil, err
				}
				if virtualServices != nil {
					istio.SetVirtualServices(virtualServices)
				} else {
					istio.SetVirtualServices(nil)
				}
			}
		}
		return istio, nil
	}
	return nil, nil
}

func expandDestinationRule(data interface{}) (*oceancd.DestinationRule, error) {
	if list := data.([]interface{}); len(list) > 0 {
		destinationRule := &oceancd.DestinationRule{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(CanarySubsetName)].(string); ok && v != "" {
				destinationRule.SetCanarySubsetName(spotinst.String(v))
			}

			if v, ok := m[string(DestinationRuleName)].(string); ok && v != "" {
				destinationRule.SetName(spotinst.String(v))
			}

			if v, ok := m[string(StableSubsetName)].(string); ok && v != "" {
				destinationRule.SetStableSubsetName(spotinst.String(v))
			}
		}
		return destinationRule, nil
	}
	return nil, nil
}

func expandVirtualServices(data interface{}) ([]*oceancd.VirtualServices, error) {

	list := data.(*schema.Set).List()

	virtualServices := make([]*oceancd.VirtualServices, 0, len(list))

	for _, v := range list {

		m := v.(map[string]interface{})

		virtualService := &oceancd.VirtualServices{}

		if v, ok := m[string(VirtualServiceName)].(string); ok && v != "" {
			virtualService.SetName(spotinst.String(v))
		}

		if v, ok := m[string(VirtualServiceRoutes)]; ok && v != nil {
			routes, err := expandRoutes(v)
			if err != nil {
				return nil, err
			}
			if routes != nil {
				virtualService.SetRoutes(routes)
			} else {
				virtualService.SetRoutes(nil)
			}
		}

		if v, ok := m[string(TlsRoutes)]; ok && v != nil {

			tlsRoutes, err := expandTlsRoutes(v)
			if err != nil {
				return nil, err
			}
			if tlsRoutes != nil {
				virtualService.SetTlsRoutes(tlsRoutes)
			} else {
				virtualService.SetTlsRoutes(nil)
			}
		}

		virtualServices = append(virtualServices, virtualService)

	}
	return virtualServices, nil
}

func expandRoutes(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if mappings, ok := v.(string); ok && mappings != "" {
			result = append(result, mappings)
		}
	}
	return result, nil
}

func expandTlsRoutes(data interface{}) ([]*oceancd.TlsRoutes, error) {

	list := data.(*schema.Set).List()

	tlsRoutes := make([]*oceancd.TlsRoutes, 0, len(list))

	for _, v := range list {

		m := v.(map[string]interface{})

		tlsRoute := &oceancd.TlsRoutes{}

		if v, ok := m[string(Port)].(int); ok {
			if v == -1 {
				tlsRoute.SetPort(nil)
			} else {
				tlsRoute.SetPort(spotinst.Int(v))
			}
		}

		if v, ok := m[string(SniHosts)]; ok && v != nil {
			routes, err := expandSniHosts(v)
			if err != nil {
				return nil, err
			}
			if routes != nil {
				tlsRoute.SetSniHosts(routes)
			} else {
				tlsRoute.SetSniHosts(nil)
			}
		}
		tlsRoutes = append(tlsRoutes, tlsRoute)

	}
	return tlsRoutes, nil
}

func expandSniHosts(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if mappings, ok := v.(string); ok && mappings != "" {
			result = append(result, mappings)
		}
	}
	return result, nil
}

func expandNginx(data interface{}) (*oceancd.Nginx, error) {
	if list := data.([]interface{}); len(list) > 0 {
		nginx := &oceancd.Nginx{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(NginxAnnotationPrefix)].(string); ok && v != "" {
				nginx.SetAnnotationPreffix(spotinst.String(v))
			} else {
				nginx.SetAnnotationPreffix(nil)
			}

			if v, ok := m[string(StableIngress)].(string); ok && v != "" {
				nginx.SetStableIngress(spotinst.String(v))
			}

			if v, ok := m[string(AdditionalIngressAnnotation)]; ok && v != nil {

				additionalIngressAnnotation, err := expandAdditionalIngressAnnotation(v)
				if err != nil {
					return nil, err
				}
				if additionalIngressAnnotation != nil {
					nginx.SetAdditionalIngressAnnotation(additionalIngressAnnotation)
				} else {
					nginx.SetAdditionalIngressAnnotation(nil)
				}
			}
		}
		return nginx, nil
	}
	return nil, nil
}

func expandAdditionalIngressAnnotation(data interface{}) (*oceancd.AdditionalIngressAnnotations, error) {
	if list := data.([]interface{}); len(list) > 0 {
		additionalIngressAnnotation := &oceancd.AdditionalIngressAnnotations{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(CanaryByHeader)].(string); ok && v != "" {
				additionalIngressAnnotation.SetCanaryByHeader(spotinst.String(v))
			} else {
				additionalIngressAnnotation.SetCanaryByHeader(nil)
			}

			if v, ok := m[string(Key1)].(string); ok && v != "" {
				additionalIngressAnnotation.SetKey1(spotinst.String(v))
			} else {
				additionalIngressAnnotation.SetKey1(nil)
			}
		}
		return additionalIngressAnnotation, nil
	}
	return nil, nil
}

func expandPingPong(data interface{}) (*oceancd.PingPong, error) {
	if list := data.([]interface{}); len(list) > 0 {
		pingPong := &oceancd.PingPong{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(PingService)].(string); ok && v != "" {
				pingPong.SetPingService(spotinst.String(v))
			}

			if v, ok := m[string(PongService)].(string); ok && v != "" {
				pingPong.SetPongService(spotinst.String(v))
			}
		}
		return pingPong, nil
	}
	return nil, nil
}

func expandSmi(data interface{}) (*oceancd.Smi, error) {
	if list := data.([]interface{}); len(list) > 0 {
		smi := &oceancd.Smi{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(SmiRootService)].(string); ok && v != "" {
				smi.SetRootService(spotinst.String(v))
			} else {
				smi.SetRootService(nil)
			}

			if v, ok := m[string(TrafficSplitName)].(string); ok && v != "" {
				smi.SetTrafficSplitName(spotinst.String(v))
			} else {
				smi.SetTrafficSplitName(nil)
			}
		}
		return smi, nil
	}
	return nil, nil
}

func flattenTraffic(traffic *oceancd.Traffic) []interface{} {
	var response []interface{}

	if traffic != nil {
		result := make(map[string]interface{})

		result[string(StableService)] = spotinst.StringValue(traffic.StableService)

		result[string(CanaryService)] = spotinst.StringValue(traffic.CanaryService)

		if traffic.Alb != nil {
			result[string(Alb)] = flattenAlb(traffic.Alb)
		}

		if traffic.Ambassador != nil {
			result[string(Ambassador)] = flattenAmbassador(traffic.Ambassador)
		}

		if traffic.Istio != nil {
			result[string(Istio)] = flattenIstio(traffic.Istio)
		}

		if traffic.Nginx != nil {
			result[string(Nginx)] = flattenNginx(traffic.Nginx)
		}

		if traffic.PingPong != nil {
			result[string(PingPong)] = flattenPingPong(traffic.PingPong)
		}

		if traffic.Smi != nil {
			result[string(Smi)] = flattenSmi(traffic.Smi)
		}

		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}

func flattenAlb(alb *oceancd.Alb) []interface{} {
	var response []interface{}

	if alb != nil {
		result := make(map[string]interface{})
		value := spotinst.Int(-1)
		result[string(ServicePort)] = value

		result[string(AlbAnnotationPrefix)] = spotinst.StringValue(alb.AnnotationPrefix)

		result[string(AlbIngress)] = spotinst.StringValue(alb.Ingress)

		result[string(AlbRootService)] = spotinst.StringValue(alb.RootService)

		if alb.StickinessConfig != nil {
			result[string(StickinessConfig)] = flattenStickinessConfig(alb.StickinessConfig)
		}

		if alb.ServicePort != nil {
			result[string(ServicePort)] = spotinst.IntValue(alb.ServicePort)
		}

		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}

func flattenStickinessConfig(stickiness *oceancd.StickinessConfig) []interface{} {
	var response []interface{}

	if stickiness != nil {
		result := make(map[string]interface{})
		value := spotinst.Int(-1)
		result[string(StickinessDuration)] = value

		if stickiness.DurationSeconds != nil {
			result[string(StickinessDuration)] = spotinst.IntValue(stickiness.DurationSeconds)
		}

		result[string(StickinessEnabled)] = spotinst.BoolValue(stickiness.Enabled)

		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}

func flattenAmbassador(ambassador *oceancd.Ambassador) []interface{} {
	var response []interface{}

	if ambassador != nil {
		result := make(map[string]interface{})

		if ambassador.Mappings != nil {
			result[string(Mappings)] = spotinst.StringSlice(ambassador.Mappings)
		}

		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}

func flattenIstio(istio *oceancd.Istio) []interface{} {
	var response []interface{}

	if istio != nil {
		result := make(map[string]interface{})

		if istio.DestinationRule != nil {
			result[string(DestinationRule)] = flattenDestinationRule(istio.DestinationRule)
		}

		if istio.VirtualServices != nil {
			result[string(VirtualServices)] = flattenVirtualServices(istio.VirtualServices)
		}

		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}

func flattenDestinationRule(destinationRule *oceancd.DestinationRule) []interface{} {
	var response []interface{}

	if destinationRule != nil {
		result := make(map[string]interface{})

		result[string(CanarySubsetName)] = spotinst.StringValue(destinationRule.CanarySubsetName)

		result[string(DestinationRuleName)] = spotinst.StringValue(destinationRule.Name)

		result[string(StableSubsetName)] = spotinst.StringValue(destinationRule.StableSubsetName)

		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}

func flattenVirtualServices(virtualServices []*oceancd.VirtualServices) []interface{} {
	result := make([]interface{}, 0, len(virtualServices))

	for _, virtualService := range virtualServices {
		m := make(map[string]interface{})

		m[string(VirtualServiceName)] = spotinst.StringValue(virtualService.Name)

		if virtualService.Routes != nil {
			m[string(VirtualServiceRoutes)] = spotinst.StringSlice(virtualService.Routes)
		}

		if virtualService.TlsRoutes != nil {
			m[string(TlsRoutes)] = flattenTlsRoutes(virtualService.TlsRoutes)
		}
		result = append(result, m)
	}
	return result
}

func flattenTlsRoutes(tlsRoutes []*oceancd.TlsRoutes) []interface{} {
	result := make([]interface{}, 0, len(tlsRoutes))

	for _, tlsRoute := range tlsRoutes {
		m := make(map[string]interface{})
		value := spotinst.Int(-1)
		m[string(Port)] = value

		if tlsRoute.Port != nil {
			m[string(Port)] = spotinst.IntValue(tlsRoute.Port)
		}

		if tlsRoute.SniHosts != nil {
			m[string(SniHosts)] = spotinst.StringSlice(tlsRoute.SniHosts)
		}

		result = append(result, m)
	}
	return result
}

func flattenNginx(nginx *oceancd.Nginx) []interface{} {
	var response []interface{}

	if nginx != nil {
		result := make(map[string]interface{})

		result[string(NginxAnnotationPrefix)] = spotinst.StringValue(nginx.AnnotationPrefix)

		result[string(StableIngress)] = spotinst.StringValue(nginx.StableIngress)

		if nginx.AdditionalIngressAnnotations != nil {
			result[string(AdditionalIngressAnnotation)] = flattenAdditionalIngressAnnotations(nginx.AdditionalIngressAnnotations)
		}

		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}

func flattenAdditionalIngressAnnotations(additionalIngressAnnotation *oceancd.AdditionalIngressAnnotations) []interface{} {
	var response []interface{}

	if additionalIngressAnnotation != nil {
		result := make(map[string]interface{})

		result[string(CanaryByHeader)] = spotinst.StringValue(additionalIngressAnnotation.CanaryByHeader)

		result[string(Key1)] = spotinst.StringValue(additionalIngressAnnotation.Key1)

		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}

func flattenPingPong(pingPong *oceancd.PingPong) []interface{} {
	var response []interface{}

	if pingPong != nil {
		result := make(map[string]interface{})

		result[string(PingService)] = spotinst.StringValue(pingPong.PingService)

		result[string(PongService)] = spotinst.StringValue(pingPong.PongService)

		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}

func flattenSmi(smi *oceancd.Smi) []interface{} {
	var response []interface{}

	if smi != nil {
		result := make(map[string]interface{})

		result[string(SmiRootService)] = spotinst.StringValue(smi.RootService)

		result[string(TrafficSplitName)] = spotinst.StringValue(smi.TrafficSplitName)

		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}
