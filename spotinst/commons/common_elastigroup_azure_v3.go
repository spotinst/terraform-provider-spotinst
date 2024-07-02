package commons

import (
	"fmt"
	"log"

	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ElastigroupAzureV3ResourceName ResourceName = "spotinst_elastigroup_azure_v3"
)

var ElastigroupAzureV3Resource *ElastigroupAzureV3TerraformResource

type ElastigroupAzureV3TerraformResource struct {
	GenericResource
}

type ElastigroupAzureV3Wrapper struct {
	elastigroup *azurev3.Group
}

func NewElastigroupAzureV3Resource(fieldsMap map[FieldName]*GenericField) *ElastigroupAzureV3TerraformResource {
	return &ElastigroupAzureV3TerraformResource{
		GenericResource: GenericResource{
			resourceName: ElastigroupAzureV3ResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *ElastigroupAzureV3TerraformResource) OnRead(
	elastigroup *azurev3.Group,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	egWrapper := NewElastigroupAzureV3Wrapper()
	egWrapper.SetElastigroup(elastigroup)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(egWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *ElastigroupAzureV3TerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*azurev3.Group, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	egWrapper := NewElastigroupAzureV3Wrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(egWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return egWrapper.GetElastigroup(), nil
}

func (res *ElastigroupAzureV3TerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *azurev3.Group, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	egWrapper := NewElastigroupAzureV3Wrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(egWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, egWrapper.GetElastigroup(), nil
}

// Spotinst elastigroup must have a wrapper struct.
// Reason is that there are multiple fields who share the same elastigroup API object
// e.g. LoadBalancersConfig fields and BlockDeviceMapping fields
// Wrapper struct intended to help reflecting these fields state properly into the elastigroup object.
func NewElastigroupAzureV3Wrapper() *ElastigroupAzureV3Wrapper {
	return &ElastigroupAzureV3Wrapper{
		elastigroup: &azurev3.Group{
			Compute: &azurev3.Compute{
				LaunchSpecification: &azurev3.LaunchSpecification{
					//LoadBalancersConfig: &v3.LoadBalancersConfig{},
				},
			},
			Capacity: &azurev3.Capacity{},
			Strategy: &azurev3.Strategy{},
		},
	}
}

func (egWrapper *ElastigroupAzureV3Wrapper) GetElastigroup() *azurev3.Group {
	return egWrapper.elastigroup
}

func (egWrapper *ElastigroupAzureV3Wrapper) SetElastigroup(elastigroup *azurev3.Group) {
	egWrapper.elastigroup = elastigroup
}
