package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	ElastigroupGCPResourceName ResourceName = "spotinst_elastigroup_gcp"
)

var ElastigroupGCPResource *ElastigroupGCPTerraformResource

type ElastigroupGCPTerraformResource struct {
	GenericResource // embedding
}

type ElastigroupGCPWrapper struct {
	elastigroup *gcp.Group
}

// NewElastigroupGCPResource creates a new GCP resource
func NewElastigroupGCPResource(fieldMap map[FieldName]*GenericField) *ElastigroupGCPTerraformResource {
	return &ElastigroupGCPTerraformResource{
		GenericResource: GenericResource{
			resourceName: ElastigroupGCPResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

// OnCreate is called when creating a new resource block and returns a new elastigroup or an error.
func (res *ElastigroupGCPTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*gcp.Group, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	egWrapper := NewElastigroupGCPWrapper()

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

// OnRead is called when reading an existing resource and throws an error if it is unable to do so.
func (res *ElastigroupGCPTerraformResource) OnRead(
	elastigroup *gcp.Group,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	egWrapper := NewElastigroupGCPWrapper()
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

// OnUpdate is called when updating an existing resource and returns
// an elastigroup with a bool indicating if had been updated, or an error.
func (res *ElastigroupGCPTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *gcp.Group, error) {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	egWrapper := NewElastigroupGCPWrapper()
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

// NewElastigroupGCPWrapper avoids parameter collisions and returns a GCP Elastigroup.
// Spotinst elastigroup must have a wrapper struct.
// The reason is that there are multiple fields that share the same elastigroup API object.
// the wrapper struct is intended to help reflect the field states into the elastigroup object properly.
func NewElastigroupGCPWrapper() *ElastigroupGCPWrapper {
	return &ElastigroupGCPWrapper{
		elastigroup: &gcp.Group{
			Capacity: &gcp.Capacity{},
			Compute: &gcp.Compute{
				LaunchSpecification: &gcp.LaunchSpecification{},
				InstanceTypes:       &gcp.InstanceTypes{},
			},
			Integration: &gcp.Integration{},
			Scaling:     &gcp.Scaling{},
			Scheduling:  &gcp.Scheduling{},
			Strategy:    &gcp.Strategy{},
		},
	}
}

// GetElastigroup returns a wrapped elastigroup
func (egWrapper *ElastigroupGCPWrapper) GetElastigroup() *gcp.Group {
	return egWrapper.elastigroup
}

// SetElastigroup applies elastigroup fields to the elastigroup wrapper.
func (egWrapper *ElastigroupGCPWrapper) SetElastigroup(elastigroup *gcp.Group) {
	egWrapper.elastigroup = elastigroup
}
