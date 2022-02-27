package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/extendedresourcedefinition"
)

const (
	ExtendedResourceDefinitionResourceName ResourceName = "spotinst_extended_resource_definition"
)

var ExtendedResourceDefinitionResource *ExtendedResourceDefinitionTerraformResource

type ExtendedResourceDefinitionTerraformResource struct {
	GenericResource
}

type ExtendedResourceDefinitionWrapper struct {
	extendedResourceDefinition *extendedresourcedefinition.ExtendedResourceDefinition
}

// NewExtendedResourceDefinitionResource creates a new ExtendedResourceDefinition resource
func NewExtendedResourceDefinitionResource(fieldMap map[FieldName]*GenericField) *ExtendedResourceDefinitionTerraformResource {
	return &ExtendedResourceDefinitionTerraformResource{
		GenericResource: GenericResource{
			resourceName: ExtendedResourceDefinitionResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

// OnCreate is called when creating a new resource block and returns a new ExtendedResourceDefinition or an error.
func (res *ExtendedResourceDefinitionTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*extendedresourcedefinition.ExtendedResourceDefinition, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	erdWrapper := NewExtendedResourceDefinitionWrapper()
	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(erdWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return erdWrapper.GetExtendedResourceDefinition(), nil
}

// OnRead is called when reading an existing resource and throws an error if it is unable to do so.
func (res *ExtendedResourceDefinitionTerraformResource) OnRead(
	extendedResourceDefinition *extendedresourcedefinition.ExtendedResourceDefinition,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	erdWrapper := NewExtendedResourceDefinitionWrapper()
	erdWrapper.SetExtendedResourceDefinition(extendedResourceDefinition)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(erdWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

// OnUpdate is called when updating an existing resource and returns
// an extendedResourceDefinition with a bool indicating if had been updated, or an error.
func (res *ExtendedResourceDefinitionTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *extendedresourcedefinition.ExtendedResourceDefinition, error) {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	erdWrapper := NewExtendedResourceDefinitionWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(erdWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, erdWrapper.GetExtendedResourceDefinition(), nil
}

// Spotinst ExtendedResourceDefinition must have a wrapper struct.
// the wrapper struct is intended to help reflect the field states into the ExtendedResourceDefinition object properly.
func NewExtendedResourceDefinitionWrapper() *ExtendedResourceDefinitionWrapper {
	return &ExtendedResourceDefinitionWrapper{
		extendedResourceDefinition: &extendedresourcedefinition.ExtendedResourceDefinition{},
	}
}

// GetExtendedResourceDefinition returns a wrapped ExtendedResourceDefinition
func (erdWrapper *ExtendedResourceDefinitionWrapper) GetExtendedResourceDefinition() *extendedresourcedefinition.ExtendedResourceDefinition {
	return erdWrapper.extendedResourceDefinition
}

// SetExtendedResourceDefinition applies extendedResourceDefinition fields to the extendedResourceDefinition wrapper.
func (erdWrapper *ExtendedResourceDefinitionWrapper) SetExtendedResourceDefinition(erd *extendedresourcedefinition.ExtendedResourceDefinition) {
	erdWrapper.extendedResourceDefinition = erd
}
