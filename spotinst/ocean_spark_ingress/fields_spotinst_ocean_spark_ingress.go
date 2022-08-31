package ocean_spark_ingress

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/spark"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Ingress] = commons.NewGenericField(
		commons.OceanSparkIngress,
		Ingress,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(ServiceAnnotations): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(AnnotationKey): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(AnnotationValue): {
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
	return []interface{}{result}
}

func expandIngress(data interface{}, nullify bool) (*spark.IngressConfig, error) {
	ingress := &spark.IngressConfig{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return ingress, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(ServiceAnnotations)]; ok {
		annotations, err := expandAnnotations(v)
		if err != nil {
			return nil, err
		}
		if annotations != nil && len(annotations) > 0 {
			ingress.SetServiceAnnotations(annotations)
		} else {
			if nullify {
				ingress.SetServiceAnnotations(nil)
			}
		}
	}

	return ingress, nil
}

func flattenAnnotations(annotations map[string]string) []interface{} {
	result := make([]interface{}, 0, len(annotations))
	for k, v := range annotations {
		m := make(map[string]interface{})
		m[string(AnnotationKey)] = k
		m[string(AnnotationValue)] = v
		result = append(result, m)
	}
	return result
}

func expandAnnotations(data interface{}) (map[string]string, error) {
	list := data.(*schema.Set).List()
	annotations := make(map[string]string, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(AnnotationKey)]; !ok {
			return nil, errors.New("invalid annotation: key missing")
		}
		if _, ok := attr[string(AnnotationValue)]; !ok {
			return nil, errors.New("invalid annotation: value missing")
		}
		key := attr[string(AnnotationKey)].(string)
		value := attr[string(AnnotationValue)].(string)
		annotations[key] = value
	}
	return annotations, nil
}
