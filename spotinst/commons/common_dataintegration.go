package commons

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/dataintegration/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var vendorTypes = []string{"s3"}

const (
	DataIntegrationResourceName ResourceName = "spotinst_data_integration"
)

var DataIntegrationResource *DataIntegrationResourceTerraformResource

type DataIntegrationResourceTerraformResource struct {
	GenericResource
}

type DataIntegrationWrapper struct {
	DataIntegration *aws.DataIntegration
}

// NewDataIntegrationResource creates a new DataIntegration resource
func NewDataIntegrationResource(fieldMap map[FieldName]*GenericField) *DataIntegrationResourceTerraformResource {
	return &DataIntegrationResourceTerraformResource{
		GenericResource: GenericResource{
			resourceName: DataIntegrationResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

// OnCreate is called when creating a new resource block and returns a new DataIntegration resource or an error.
func (res *DataIntegrationResourceTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.DataIntegration, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	diWrapper := NewDataIntegrationWrapper()
	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(diWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return diWrapper.GetDataIntegration(), nil
}

// OnRead is called when reading an existing resource and throws an error if it is unable to do so.
func (res *DataIntegrationResourceTerraformResource) OnRead(
	extendedResourceDefinition *aws.DataIntegration,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	diWrapper := NewDataIntegrationWrapper()
	diWrapper.SetDataIntegration(extendedResourceDefinition)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(diWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

// OnUpdate is called when updating an existing resource and returns
// an DataIntegration with a bool indicating if had been updated, or an error.
func (res *DataIntegrationResourceTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *aws.DataIntegration, error) {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	diWrapper := NewDataIntegrationWrapper()
	hasChanged := false
	var vendor = ""
	for _, field := range res.fields.fieldsMap {
		if contains(vendorTypes, field.fieldNameStr) {
			vendor = field.fieldNameStr
		}
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(diWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}
	diWrapper.DataIntegration.SetVendor(spotinst.String(vendor))
	return hasChanged, diWrapper.GetDataIntegration(), nil
}

// Spotinst DataIntegration must have a wrapper struct.
// the wrapper struct is intended to help reflect the field states into the DataIntegration object properly.
func NewDataIntegrationWrapper() *DataIntegrationWrapper {
	return &DataIntegrationWrapper{
		DataIntegration: &aws.DataIntegration{},
	}
}

// GetDataIntegration returns a wrapped DataIntegration
func (diWrapper *DataIntegrationWrapper) GetDataIntegration() *aws.DataIntegration {
	return diWrapper.DataIntegration
}

// SetDataIntegration applies DataIntegration fields to the DataIntegration wrapper.
func (diWrapper *DataIntegrationWrapper) SetDataIntegration(di *aws.DataIntegration) {
	diWrapper.DataIntegration = di
}
