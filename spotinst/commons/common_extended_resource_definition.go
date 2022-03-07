package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
)

const (
	OceanAWSExtendedResourceDefinitionResourceName ResourceName = "spotinst_ocean_aws_extended_resource_definition"
)

var OceanAWSExtendedResourceDefinitionResource *OceanAWSExtendedResourceDefinitionTerraformResource

type OceanAWSExtendedResourceDefinitionTerraformResource struct {
	GenericResource
}

type ExtendedResourceDefinitionWrapper struct {
	extendedResourceDefinition *aws.ExtendedResourceDefinition
}

// NewOceanAWSExtendedResourceDefinitionResource creates a new OceanAWSExtendedResourceDefinition resource
func NewOceanAWSExtendedResourceDefinitionResource(fieldMap map[FieldName]*GenericField) *OceanAWSExtendedResourceDefinitionTerraformResource {
	return &OceanAWSExtendedResourceDefinitionTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanAWSExtendedResourceDefinitionResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

// OnCreate is called when creating a new resource block and returns a new OceanAWSExtendedResourceDefinition or an error.
func (res *OceanAWSExtendedResourceDefinitionTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.ExtendedResourceDefinition, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	erdWrapper := NewOceanAWSExtendedResourceDefinitionWrapper()
	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(erdWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return erdWrapper.GetOceanAWSExtendedResourceDefinition(), nil
}

// OnRead is called when reading an existing resource and throws an error if it is unable to do so.
func (res *OceanAWSExtendedResourceDefinitionTerraformResource) OnRead(
	extendedResourceDefinition *aws.ExtendedResourceDefinition,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	erdWrapper := NewOceanAWSExtendedResourceDefinitionWrapper()
	erdWrapper.SetOceanAWSExtendedResourceDefinition(extendedResourceDefinition)

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
func (res *OceanAWSExtendedResourceDefinitionTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *aws.ExtendedResourceDefinition, error) {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	erdWrapper := NewOceanAWSExtendedResourceDefinitionWrapper()
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

	return hasChanged, erdWrapper.GetOceanAWSExtendedResourceDefinition(), nil
}

// Spotinst ExtendedResourceDefinition must have a wrapper struct.
// the wrapper struct is intended to help reflect the field states into the ExtendedResourceDefinition object properly.
func NewOceanAWSExtendedResourceDefinitionWrapper() *ExtendedResourceDefinitionWrapper {
	return &ExtendedResourceDefinitionWrapper{
		extendedResourceDefinition: &aws.ExtendedResourceDefinition{},
	}
}

// GetOceanAWSExtendedResourceDefinition returns a wrapped OceanAWSExtendedResourceDefinition
func (erdWrapper *ExtendedResourceDefinitionWrapper) GetOceanAWSExtendedResourceDefinition() *aws.ExtendedResourceDefinition {
	return erdWrapper.extendedResourceDefinition
}

// SetOceanAWSExtendedResourceDefinition applies extendedResourceDefinition fields to the extendedResourceDefinition wrapper.
func (erdWrapper *ExtendedResourceDefinitionWrapper) SetOceanAWSExtendedResourceDefinition(erd *aws.ExtendedResourceDefinition) {
	erdWrapper.extendedResourceDefinition = erd
}
