package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/spark"
)

const (
	OceanSparkVirtualNodeGroupResourceName ResourceName = "spotinst_ocean_spark_virtual_node_group"
)

var OceanSparkVirtualNodeGroupResource *OceanSparkVirtualNodeGroupTerraformResource

type OceanSparkVirtualNodeGroupTerraformResource struct {
	GenericResource
}

type SparkVirtualNodeGroupWrapper struct {
	virtualNodeGroup *spark.DedicatedVirtualNodeGroup
}

func NewOceanSparkVirtualNodeGroupResource(fieldsMap map[FieldName]*GenericField) *OceanSparkVirtualNodeGroupTerraformResource {
	return &OceanSparkVirtualNodeGroupTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanSparkVirtualNodeGroupResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanSparkVirtualNodeGroupTerraformResource) OnRead(
	vng *spark.DedicatedVirtualNodeGroup,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	vngWrapper := NewSparkVirtualNodeGroupWrapper()
	vngWrapper.SetVirtualNodeGroup(vng)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(vngWrapper, resourceData, meta); err != nil {
			return err
		}
	}

	return nil
}

func (res *OceanSparkVirtualNodeGroupTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*spark.DedicatedVirtualNodeGroup, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	vngWrapper := NewSparkVirtualNodeGroupWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(vngWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return vngWrapper.GetVirtualNodeGroup(), nil
}

func NewSparkVirtualNodeGroupWrapper() *SparkVirtualNodeGroupWrapper {
	return &SparkVirtualNodeGroupWrapper{
		virtualNodeGroup: &spark.DedicatedVirtualNodeGroup{},
	}
}

func (virtualNodeGroupWrapper *SparkVirtualNodeGroupWrapper) GetVirtualNodeGroup() *spark.DedicatedVirtualNodeGroup {
	return virtualNodeGroupWrapper.virtualNodeGroup
}

func (virtualNodeGroupWrapper *SparkVirtualNodeGroupWrapper) SetVirtualNodeGroup(virtualNodeGroup *spark.DedicatedVirtualNodeGroup) {
	virtualNodeGroupWrapper.virtualNodeGroup = virtualNodeGroup
}
