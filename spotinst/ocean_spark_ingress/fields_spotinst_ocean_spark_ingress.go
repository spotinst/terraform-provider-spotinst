package ocean_spark_ingress

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/spark"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Ingress] = commons.NewGenericField(
		commons.OceanSparkIngress,
		Ingress,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ServiceAnnotations): {
						Type:     schema.TypeMap,
						Optional: true,
						Computed: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					string(Controller): {
						Type:     schema.TypeList,
						Optional: true,
						Computed: false,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Managed): {
									Type:     schema.TypeBool,
									Optional: true,
									Computed: false,
									Default:  true,
								},
							},
						},
					},
					string(LoadBalancer): {
						Type:     schema.TypeList,
						Optional: true,
						Computed: false,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(ServiceAnnotations): {
									Type:     schema.TypeMap,
									Optional: true,
									Computed: false,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								string(Managed): {
									Type:     schema.TypeBool,
									Optional: true,
									Computed: false,
									Default:  true,
								},
								string(TargetGroupARN): {
									Type:     schema.TypeString,
									Optional: true,
									Computed: false,
								},
							},
						},
					},
					string(CustomEndpoint): {
						Type:     schema.TypeList,
						Optional: true,
						Computed: false,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Enabled): {
									Type:     schema.TypeBool,
									Optional: true,
									Computed: false,
									Default:  false,
								},
								string(Address): {
									Type:     schema.TypeString,
									Optional: true,
									Computed: false,
								},
							},
						},
					},
					string(PrivateLink): {
						Type:     schema.TypeList,
						Optional: true,
						Computed: false,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Enabled): {
									Type:     schema.TypeBool,
									Optional: true,
									Computed: false,
									Default:  false,
								},
								string(VPCEndpointService): {
									Type:     schema.TypeString,
									Optional: true,
									Computed: false,
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster.Config != nil && cluster.Config.Ingress != nil {
				result = flattenIngress(cluster.Config.Ingress)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Ingress), result); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(Ingress), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if value, ok := resourceData.GetOk(string(Ingress)); ok {
				if ingress, err := expandIngress(value, false); err != nil {
					return err
				} else {
					if cluster.Config == nil {
						cluster.Config = &spark.Config{}
					}
					cluster.Config.SetIngress(ingress)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *spark.IngressConfig = nil
			if v, ok := resourceData.GetOk(string(Ingress)); ok {
				if ingress, err := expandIngress(v, true); err != nil {
					return err
				} else {
					value = ingress
				}
			}
			if cluster.Config == nil {
				cluster.Config = &spark.Config{}
			}
			cluster.Config.SetIngress(value)
			return nil
		},
		nil,
	)
}

func flattenIngress(ingress *spark.IngressConfig) []interface{} {
	if ingress == nil {
		return nil
	}
	result := make(map[string]interface{})
	result[string(ServiceAnnotations)] = flattenAnnotations(ingress.ServiceAnnotations)
	result[string(Controller)] = flattenController(ingress.Controller)
	result[string(LoadBalancer)] = flattenLoadBalancer(ingress.LoadBalancer)
	result[string(CustomEndpoint)] = flattenCustomEndpoint(ingress.CustomEndpoint)
	result[string(PrivateLink)] = flattenPrivateLink(ingress.PrivateLink)

	return []interface{}{result}
}

func expandIngress(data interface{}, nullify bool) (*spark.IngressConfig, error) {
	ingress := &spark.IngressConfig{}
	list := data.([]interface{})
	if len(list) == 0 || list[0] == nil {
		return ingress, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(ServiceAnnotations)]; ok {
		annotations, err := expandAnnotations(v)
		if err != nil {
			return nil, err
		}
		if len(annotations) > 0 {
			ingress.SetServiceAnnotations(annotations)
		} else {
			if nullify {
				ingress.SetServiceAnnotations(nil)
			}
		}
	}

	if v, ok := m[string(Controller)]; ok {
		controller, err := expandController(v)
		if err != nil {
			return nil, err
		}
		ingress.SetController(controller)
	}

	if v, ok := m[string(LoadBalancer)]; ok {
		loadBalancer, err := expandLoadBalancer(v, nullify)
		if err != nil {
			return nil, err
		}
		ingress.SetLoadBalancer(loadBalancer)
	}

	if v, ok := m[string(CustomEndpoint)]; ok {
		customEndpoint, err := expandCustomEndpoint(v)
		if err != nil {
			return nil, err
		}
		ingress.SetCustomEndpoint(customEndpoint)
	}

	if v, ok := m[string(PrivateLink)]; ok {
		privateLink, err := expandPrivateLink(v)
		if err != nil {
			return nil, err
		}
		ingress.SetPrivateLink(privateLink)
	}

	return ingress, nil
}

func flattenController(controller *spark.IngressConfigController) []interface{} {
	if controller == nil {
		return nil
	}
	result := make(map[string]interface{})
	if controller.Managed != nil {
		result[string(Managed)] = spotinst.BoolValue(controller.Managed)
	}
	return []interface{}{result}
}

func flattenLoadBalancer(loadBalancer *spark.IngressConfigLoadBalancer) []interface{} {
	if loadBalancer == nil {
		return nil
	}
	result := make(map[string]interface{})

	if loadBalancer.ServiceAnnotations != nil {
		result[string(ServiceAnnotations)] = flattenAnnotations(loadBalancer.ServiceAnnotations)
	}

	if loadBalancer.Managed != nil {
		result[string(Managed)] = spotinst.BoolValue(loadBalancer.Managed)
	}

	if loadBalancer.TargetGroupARN != nil {
		result[string(TargetGroupARN)] = spotinst.StringValue(loadBalancer.TargetGroupARN)
	}

	return []interface{}{result}
}

func flattenCustomEndpoint(customEndpoint *spark.IngressConfigCustomEndpoint) []interface{} {
	if customEndpoint == nil {
		return nil
	}
	result := make(map[string]interface{})

	if customEndpoint.Enabled != nil {
		result[string(Enabled)] = spotinst.BoolValue(customEndpoint.Enabled)
	}

	if customEndpoint.Address != nil {
		result[string(Address)] = spotinst.StringValue(customEndpoint.Address)
	}

	return []interface{}{result}
}

func flattenPrivateLink(privateLink *spark.IngressConfigPrivateLink) []interface{} {
	if privateLink == nil {
		return nil
	}
	result := make(map[string]interface{})

	if privateLink.Enabled != nil {
		result[string(Enabled)] = spotinst.BoolValue(privateLink.Enabled)
	}

	if privateLink.VPCEndpointService != nil {
		result[string(VPCEndpointService)] = spotinst.StringValue(privateLink.VPCEndpointService)
	}

	return []interface{}{result}
}

func flattenAnnotations(annotations map[string]string) map[string]interface{} {
	result := make(map[string]interface{}, len(annotations))
	for k, v := range annotations {
		result[k] = v
	}
	return result
}

func expandController(data interface{}) (*spark.IngressConfigController, error) {
	list := data.([]interface{})
	if len(list) == 0 || list[0] == nil {
		return nil, nil
	}
	controller := &spark.IngressConfigController{}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Managed)].(bool); ok {
		controller.SetManaged(spotinst.Bool(v))
	}

	return controller, nil
}

func expandLoadBalancer(data interface{}, nullify bool) (*spark.IngressConfigLoadBalancer, error) {
	list := data.([]interface{})
	if len(list) == 0 || list[0] == nil {
		return nil, nil
	}
	loadBalancer := &spark.IngressConfigLoadBalancer{}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Managed)].(bool); ok {
		loadBalancer.SetManaged(spotinst.Bool(v))
	}

	if v, ok := m[string(TargetGroupARN)].(string); ok {
		loadBalancer.SetTargetGroupARN(spotinst.String(v))
	}

	if v, ok := m[string(ServiceAnnotations)]; ok {
		annotations, err := expandAnnotations(v)
		if err != nil {
			return nil, err
		}
		if len(annotations) > 0 {
			loadBalancer.SetServiceAnnotations(annotations)
		} else {
			if nullify {
				loadBalancer.SetServiceAnnotations(nil)
			}
		}
	}

	return loadBalancer, nil
}

func expandCustomEndpoint(data interface{}) (*spark.IngressConfigCustomEndpoint, error) {
	list := data.([]interface{})
	if len(list) == 0 || list[0] == nil {
		return nil, nil
	}
	customEndpoint := &spark.IngressConfigCustomEndpoint{}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Enabled)].(bool); ok {
		customEndpoint.SetEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(Address)].(string); ok {
		customEndpoint.SetAddress(spotinst.String(v))
	}

	return customEndpoint, nil
}

func expandPrivateLink(data interface{}) (*spark.IngressConfigPrivateLink, error) {
	list := data.([]interface{})
	if len(list) == 0 || list[0] == nil {
		return nil, nil
	}
	privateLink := &spark.IngressConfigPrivateLink{}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Enabled)].(bool); ok {
		privateLink.SetEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(VPCEndpointService)].(string); ok {
		privateLink.SetVPCEndpointService(spotinst.String(v))
	}

	return privateLink, nil
}

func expandAnnotations(data interface{}) (map[string]string, error) {
	m, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("could not cast annotations")
	}
	result := make(map[string]string, len(m))
	for k, v := range m {
		val, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("could not cast annotation value to string")
		}
		result[k] = val
	}
	return result, nil
}
