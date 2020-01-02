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

// import constructs a specialized group and makes an import call
// it returns a standard gcp group with pre-filled fields
// this group is passed to OnCreate
// OnCreate reconciles the generated fields with user-defined fields in the plan
//  - what do we do if field is undefined?
// OnCreate sends the reconciled group to the create API endpoint

func (res *ElastigroupGKETerraformResource) OnImport(
	templateGroup *gcp.Group,
	resourceData *schema.ResourceData,
	meta interface{}) (*gcp.ImportGKEGroup, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}
	gkeGroupImportWrapper := NewImportGKEWrapper()

	gkeGroupImportWrapper.SetImport(&gcp.ImportGKEGroup{
		AvailabilityZones: templateGroup.Compute.AvailabilityZones,
		Capacity: &gcp.CapacityGKE{
			Capacity: gcp.Capacity{
				Minimum: templateGroup.Capacity.Minimum,
				Maximum: templateGroup.Capacity.Maximum,
				Target:  templateGroup.Capacity.Target,
			},
		},
		Name: templateGroup.Name,
		InstanceTypes: &gcp.InstanceTypesGKE{
			OnDemand:    templateGroup.Compute.InstanceTypes.OnDemand,
			Preemptible: templateGroup.Compute.InstanceTypes.Preemptible,
		},
		PreemptiblePercentage: templateGroup.Strategy.PreemptiblePercentage,
		NodeImage:             templateGroup.NodeImage,
	})

	return gkeGroupImportWrapper.GetImport(), nil
}

// OnCreate is called when creating a new resource block and returns a new elastigroup or an error.
func (res *ElastigroupGKETerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*gcp.Group, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	egWrapper := NewElastigroupGKEWrapper()

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

func (res *ElastigroupGKETerraformResource) OnMerge(
	importedGroup *gcp.Group,
	resourceData *schema.ResourceData,
	meta interface{}) (*gcp.Group, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	egWrapper := NewElastigroupGKEWrapper()
	egWrapper.SetElastigroup(importedGroup)

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnMerge), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(egWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}

	return egWrapper.GetElastigroup(), nil
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
func NewElastigroupGKEWrapper() *ElastigroupGCPWrapper {
	return &ElastigroupGCPWrapper{
		elastigroup: &gcp.Group{
			Capacity: &gcp.Capacity{},
			Compute: &gcp.Compute{
				LaunchSpecification: &gcp.LaunchSpecification{},
				InstanceTypes:       &gcp.InstanceTypes{},
			},
			Integration: &gcp.Integration{},
			Scaling:     &gcp.Scaling{},
			Strategy:    &gcp.Strategy{},
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

// SuppressIfImportedFromGKE suppresses a large diff between the state and user's template as a result of the import.
// users may update imported fields, but may not set them to null
func SuppressIfImportedFromGKE(k, old, new string, d *schema.ResourceData) bool {
	if _, ok := d.GetOk(string("integration_gke")); ok {
		if new == "" || new == "0" {
			return true
		}
		if old == "true" && new == "false" {
			return true
		}
	}
	return false
}
