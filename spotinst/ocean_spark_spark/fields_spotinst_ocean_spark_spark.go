package ocean_spark_spark

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/spark"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

const (
	// defaultAppNamespace is the default spark app namespace
	defaultAppNamespace = "spark-apps"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Spark] = commons.NewGenericField(
		commons.OceanSparkSpark,
		Spark,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(AdditionalAppNamespaces): {
						Type:             schema.TypeSet, // We don't care about ordering, so we use a set here
						Optional:         true,
						Computed:         true,
						Elem:             &schema.Schema{Type: schema.TypeString},
						DiffSuppressFunc: SuppressDiffDefaultAppNamespace,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster.Config != nil && cluster.Config.Spark != nil {
				result = flattenSparkConfig(cluster.Config.Spark)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Spark), result); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(Spark), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if value, ok := resourceData.GetOk(string(Spark)); ok {
				if sparkConfig, err := expandSparkConfig(value, false); err != nil {
					return err
				} else {
					if cluster.Config == nil {
						cluster.Config = &spark.Config{}
					}
					cluster.Config.SetSpark(sparkConfig)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *spark.SparkConfig = nil
			if v, ok := resourceData.GetOk(string(Spark)); ok {
				if sparkConfig, err := expandSparkConfig(v, true); err != nil {
					return err
				} else {
					value = sparkConfig
				}
			}
			if cluster.Config == nil {
				cluster.Config = &spark.Config{}
			}
			cluster.Config.SetSpark(value)
			return nil
		},
		nil,
	)
}

func flattenSparkConfig(sparkConfig *spark.SparkConfig) []interface{} {
	if sparkConfig == nil {
		return nil
	}
	result := make(map[string]interface{})
	if sparkConfig.AppNamespaces != nil {
		namespaces := make([]string, 0)
		for i := range sparkConfig.AppNamespaces {
			namespace := spotinst.StringValue(sparkConfig.AppNamespaces[i])
			if namespace == defaultAppNamespace {
				// We filter out the default app namespace which is always present in the appNamespaces list given by the API.
				// To make this work well with terraform, we call the field additional_app_namespaces,
				// and enforce behind the scenes that the spark-apps namespace is filtered out.
				continue
			}
			namespaces = append(namespaces, namespace)
		}
		result[string(AdditionalAppNamespaces)] = namespaces
	}
	return []interface{}{result}
}

func expandSparkConfig(data interface{}, nullify bool) (*spark.SparkConfig, error) {
	sparkConfig := &spark.SparkConfig{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return sparkConfig, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(AdditionalAppNamespaces)]; ok {
		namespaces, err := expandAppNamespaces(v)
		if err != nil {
			return nil, err
		}
		if len(namespaces) > 0 {
			sparkConfig.SetAppNamespaces(namespaces)
		} else {
			if nullify {
				sparkConfig.SetAppNamespaces(nil)
			}
		}
	}

	return sparkConfig, nil
}

func expandAppNamespaces(data interface{}) ([]*string, error) {
	list := data.(*schema.Set).List()
	result := make([]*string, 0)
	for _, v := range list {
		if namespace, ok := v.(string); ok && namespace != "" {
			result = append(result, spotinst.String(namespace))
		}
	}

	return result, nil
}

func SuppressDiffDefaultAppNamespace(_, old, new string, _ *schema.ResourceData) bool {
	if old == "" && new == defaultAppNamespace {
		return true
	}
	if old == defaultAppNamespace && new == "" {
		return true
	}
	return false
}
