package ocean_spark_config

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IngressServiceAnnotations] = commons.NewGenericField(
		commons.OceanSparkConfig,
		IngressServiceAnnotations,
		&schema.Schema{
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
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster.Config != nil && cluster.Config.Ingress != nil &&
				cluster.Config.Ingress.ServiceAnnotations != nil {
				annotations := cluster.Config.Ingress.ServiceAnnotations
				result = flattenAnnotations(annotations)
			}
			if result != nil {
				if err := resourceData.Set(string(IngressServiceAnnotations), result); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(IngressServiceAnnotations), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if value, ok := resourceData.GetOk(string(IngressServiceAnnotations)); ok {
				if annotations, err := expandAnnotations(value); err != nil {
					return err
				} else {
					cluster.Config.Ingress.SetServiceAnnotations(annotations)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var annotationsToAdd map[string]string = nil
			if value, ok := resourceData.GetOk(string(IngressServiceAnnotations)); ok {
				if annotations, err := expandAnnotations(value); err != nil {
					return err
				} else {
					annotationsToAdd = annotations
				}
			}
			cluster.Config.Ingress.SetServiceAnnotations(annotationsToAdd)
			return nil
		},
		nil,
	)
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
