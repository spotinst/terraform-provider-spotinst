package ocean_spark

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[OceanClusterID] = commons.NewGenericField(
		commons.OceanSpark,
		OceanClusterID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if cluster.OceanClusterID != nil {
				value = cluster.OceanClusterID
			}
			if err := resourceData.Set(string(OceanClusterID), value); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(OceanClusterID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			cluster.OceanClusterID = spotinst.String(resourceData.Get(string(OceanClusterID)).(string))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern), string(OceanClusterID))
		},
		nil,
	)
}
