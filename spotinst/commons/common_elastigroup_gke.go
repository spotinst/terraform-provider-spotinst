package commons

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"log"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	ElastigroupGKEResourceName ResourceName = "spotinst_elastigroup_gke"
)

var ElastigroupGKEResource *ElastigroupGKETerraformResource

type ElastigroupGKETerraformResource struct {
	GenericResource // embedding
}

type ElastigroupGKEWrapper struct {
	elastigroup     *gcp.Group
	ClusterID       string
	ClusterZoneName string
}

type ImportGKEWrapper struct {
	elastigroup     *gcp.ImportGKEGroup
	ClusterID       string
	ClusterZoneName string
}

// NewElastigroupGKEResource creates a new GKE resource
func NewElastigroupGKEResource(fieldMap map[FieldName]*GenericField) *ElastigroupGKETerraformResource {
	return &ElastigroupGKETerraformResource{
		GenericResource: GenericResource{
			resourceName: ElastigroupGKEResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

// OnCreate is called when creating a new resource block and returns a new elastigroup or an error.
func (res *ElastigroupGKETerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*gcp.ImportGKEGroup, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	gkeGroupImport := NewImportGKEWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)

		if err := field.onCreate(gkeGroupImport, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return gkeGroupImport.GetImport(), nil
}

// OnRead is called when reading an existing resource and throws an error if it is unable to do so.
func (res *ElastigroupGKETerraformResource) OnRead(
	elastigroup *gcp.Group,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	gkeGroupWrapper := NewElastigroupGKEWrapper()
	gkeGroupWrapper.SetElastigroup(elastigroup)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(gkeGroupWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

// OnUpdate is called when updating an existing resource and returns
// an elastigroup with a bool indicating if had been updated, or an error.
func (res *ElastigroupGKETerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *gcp.Group, error) {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	egWrapper := NewElastigroupGKEWrapper()
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

// NewElastigroupGKEWrapper avoids parameter collisions and returns a GKE Elastigroup.
// Spotinst elastigroup must have a wrapper struct.
// The reason is that there are multiple fields that share the same elastigroup API object.
// the wrapper struct is intended to help reflect the field states into the elastigroup object properly.
func NewElastigroupGKEWrapper() *ElastigroupGKEWrapper {
	return &ElastigroupGKEWrapper{
		elastigroup: &gcp.Group{
			Capacity: &gcp.Capacity{},
			Compute: &gcp.Compute{
				LaunchSpecification: &gcp.LaunchSpecification{},
				InstanceTypes:       &gcp.InstanceTypes{},
			},
			Integration: &gcp.Integration{
				GKE: &gcp.GKEIntegration{},
			},
			Scaling:  &gcp.Scaling{},
			Strategy: &gcp.Strategy{},
		},
	}
}

func NewImportGKEWrapper() *ImportGKEWrapper {
	return &ImportGKEWrapper{
		elastigroup: &gcp.ImportGKEGroup{
			Capacity: &gcp.CapacityGKE{},
		},
	}
}

func (egWrapper *ImportGKEWrapper) GetImport() *gcp.ImportGKEGroup {
	return egWrapper.elastigroup
}

func (egWrapper *ImportGKEWrapper) SetImport(elastigroup *gcp.ImportGKEGroup) {
	egWrapper.elastigroup = elastigroup
}

// GetElastigroup returns a wrapped elastigroup
func (egWrapper *ElastigroupGKEWrapper) GetElastigroup() *gcp.Group {
	return egWrapper.elastigroup
}

// SetElastigroup applies elastigroup fields to the elastigroup wrapper.
func (egWrapper *ElastigroupGKEWrapper) SetElastigroup(elastigroup *gcp.Group) {
	egWrapper.elastigroup = elastigroup
}
