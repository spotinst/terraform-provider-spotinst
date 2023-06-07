package ocean_spark_log_collection

import (
	"fmt"
	cty2 "github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/spark"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[LogCollection] = commons.NewGenericField(
		commons.OceanSparkLogCollection,
		LogCollection,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(CollectDriverLogs): {
						Type:     schema.TypeBool,
						Optional: true,
						Computed: true,
					},
					string(CollectAppLogs): {
						Type:     schema.TypeBool,
						Optional: true,
						Computed: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster.Config != nil && cluster.Config.LogCollection != nil {
				result = flattenLogCollection(cluster.Config.LogCollection)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(LogCollection), result); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(LogCollection), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if _, ok := resourceData.GetOk(string(LogCollection)); ok {
				v := resourceData.GetRawConfig().GetAttr(string(LogCollection))
				if logCollection, err := expandLogCollection(v); err != nil {
					return err
				} else {
					if cluster.Config == nil {
						cluster.Config = &spark.Config{}
					}
					cluster.Config.SetLogCollection(logCollection)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *spark.LogCollectionConfig = nil
			if _, ok := resourceData.GetOk(string(LogCollection)); ok {
				v := resourceData.GetRawConfig().GetAttr(string(LogCollection))
				if logCollection, err := expandLogCollection(v); err != nil {
					return err
				} else {
					value = logCollection
				}
			}
			if cluster.Config == nil {
				cluster.Config = &spark.Config{}
			}
			cluster.Config.SetLogCollection(value)
			return nil
		},
		nil,
	)
}

func flattenLogCollection(logCollection *spark.LogCollectionConfig) []interface{} {
	if logCollection == nil {
		return nil
	}
	result := make(map[string]interface{})
	result[string(CollectDriverLogs)] = spotinst.BoolValue(logCollection.CollectDriverLogs)
	result[string(CollectAppLogs)] = spotinst.BoolValue(logCollection.CollectAppLogs)
	return []interface{}{result}
}

func expandLogCollection(rawData cty2.Value) (*spark.LogCollectionConfig, error) {
	logCollection := &spark.LogCollectionConfig{}
	if !rawData.AsValueSlice()[0].AsValueMap()[string(CollectDriverLogs)].IsNull() {
		v := rawData.AsValueSlice()[0].AsValueMap()[string(CollectDriverLogs)].True()
		logCollection.SetCollectDriverLogs(spotinst.Bool(v))
	}
	if !rawData.AsValueSlice()[0].AsValueMap()[string(CollectAppLogs)].IsNull() {
		v := rawData.AsValueSlice()[0].AsValueMap()[string(CollectAppLogs)].True()
		logCollection.SetCollectAppLogs(spotinst.Bool(v))
	}

	return logCollection, nil
}
