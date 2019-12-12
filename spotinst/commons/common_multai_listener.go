package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	MultaiListenerResourceName ResourceName = "spotinst_multai_listener"
)

var MultaiListenerResource *MultaiListenerTerraformResource

type MultaiListenerTerraformResource struct {
	GenericResource // embedding
}

type MultaiListenerWrapper struct {
	listener *multai.Listener
}

func NewMultaiListenerResource(fieldMap map[FieldName]*GenericField) *MultaiListenerTerraformResource {
	return &MultaiListenerTerraformResource{
		GenericResource: GenericResource{
			resourceName: MultaiListenerResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

func (res *MultaiListenerTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*multai.Listener, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	mlbWrapper := NewMultaiListenerWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(mlbWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return mlbWrapper.GetMultaiListener(), nil
}

func (res *MultaiListenerTerraformResource) OnRead(
	listener *multai.Listener,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	mlbWrapper := NewMultaiListenerWrapper()
	mlbWrapper.SetMultaiListener(listener)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(mlbWrapper, resourceData, meta); err != nil {
			return err
		}
	}

	return nil
}

func (res *MultaiListenerTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *multai.Listener, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	mlbWrapper := NewMultaiListenerWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(mlbWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, mlbWrapper.GetMultaiListener(), nil
}

func NewMultaiListenerWrapper() *MultaiListenerWrapper {
	return &MultaiListenerWrapper{
		listener: &multai.Listener{},
	}
}

func (mlbWrapper *MultaiListenerWrapper) GetMultaiListener() *multai.Listener {
	return mlbWrapper.listener
}

func (mlbWrapper *MultaiListenerWrapper) SetMultaiListener(listener *multai.Listener) {
	mlbWrapper.listener = listener
}
