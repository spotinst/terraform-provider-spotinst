package commons

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
)

const (
	OceanECSResourceName ResourceName = "spotinst_ocean_ecs"
)

var OceanECSResource *OceanECSTerraformResource

type OceanECSTerraformResource struct {
	GenericResource
}

type ECSClusterWrapper struct {
	cluster *aws.ECSCluster
}

func NewOceanECSResource(fieldsMap map[FieldName]*GenericField) *OceanECSTerraformResource {
	return &OceanECSTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanECSResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanECSTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.ECSCluster, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	clusterWrapper := NewECSClusterWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(clusterWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return clusterWrapper.GetECSCluster(), nil
}

func (res *OceanECSTerraformResource) OnRead(
	cluster *aws.ECSCluster,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	clusterWrapper := NewECSClusterWrapper()
	clusterWrapper.SetECSCluster(cluster)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(clusterWrapper, resourceData, meta); err != nil {
			return err
		}
	}

	return nil
}

func (res *OceanECSTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, bool, bool, *aws.ECSCluster, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, false, false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	clusterWrapper := NewECSClusterWrapper()
	hasChanged := false
	changesRequiredRoll := false
	tagsChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			if contains(conditionedRollFieldsECS, field.fieldNameStr) {
				changesRequiredRoll = true
			}

			if strings.Compare(field.fieldNameStr, "tags") == 0 {
				tagsChanged = true
			}
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(clusterWrapper, resourceData, meta); err != nil {
				return false, false, false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, changesRequiredRoll, tagsChanged, clusterWrapper.GetECSCluster(), nil
}

func NewECSClusterWrapper() *ECSClusterWrapper {
	return &ECSClusterWrapper{
		cluster: &aws.ECSCluster{
			Capacity: &aws.ECSCapacity{},
			Compute: &aws.ECSCompute{
				LaunchSpecification: &aws.ECSLaunchSpecification{},
				InstanceTypes:       &aws.ECSInstanceTypes{},
			},
			Strategy:   &aws.ECSStrategy{},
			Scheduling: &aws.ECSScheduling{},
			Logging:    &aws.ECSLogging{},
		},
	}
}

func (clusterWrapper *ECSClusterWrapper) GetECSCluster() *aws.ECSCluster {
	return clusterWrapper.cluster
}

func (clusterWrapper *ECSClusterWrapper) SetECSCluster(cluster *aws.ECSCluster) {
	clusterWrapper.cluster = cluster
}
