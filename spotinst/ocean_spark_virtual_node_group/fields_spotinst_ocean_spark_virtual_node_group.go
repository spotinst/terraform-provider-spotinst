package ocean_spark_virtual_node_group

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[OceanSparkClusterID] = commons.NewGenericField(
		commons.OceanSparkVirtualNodeGroup,
		OceanSparkClusterID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.SparkVirtualNodeGroupWrapper)
			vng := vngWrapper.GetVirtualNodeGroup()
			var value *string = nil
			if vng.OceanClusterID != nil {
				value = vng.OceanClusterID
			}
			if err := resourceData.Set(string(OceanSparkClusterID), value); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(OceanSparkClusterID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.SparkVirtualNodeGroupWrapper)
			vng := vngWrapper.GetVirtualNodeGroup()
			vng.OceanSparkClusterID = spotinst.String(resourceData.Get(string(OceanSparkClusterID)).(string))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[VngID] = commons.NewGenericField(
		commons.OceanSparkVirtualNodeGroup,
		VngID,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.SparkVirtualNodeGroupWrapper)
			vng := vngWrapper.GetVirtualNodeGroup()
			var value *string = nil
			if vng.VngID != nil {
				value = vng.VngID
			}
			if err := resourceData.Set(string(VngID), value); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(VngID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.SparkVirtualNodeGroupWrapper)
			vng := vngWrapper.GetVirtualNodeGroup()
			vng.VngID = spotinst.String(resourceData.Get(string(VngID)).(string))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)
}
