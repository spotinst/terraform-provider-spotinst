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
	MultaiTargetResourceName ResourceName = "spotinst_multai_target"
)

var MultaiTargetResource *MultaiTargetTerraformResource

type MultaiTargetTerraformResource struct {
	GenericResource // embedding
}

type MultaiTargetWrapper struct {
	target *multai.Target
}

func NewMultaiTargetResource(fieldMap map[FieldName]*GenericField) *MultaiTargetTerraformResource {
	return &MultaiTargetTerraformResource{
		GenericResource: GenericResource{
			resourceName: MultaiTargetResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

func (res *MultaiTargetTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*multai.Target, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	targetWrapper := NewMultaiTargetWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(targetWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return targetWrapper.GetMultaiTarget(), nil
}

func (res *MultaiTargetTerraformResource) OnRead(
	target *multai.Target,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	targetWrapper := NewMultaiTargetWrapper()
	targetWrapper.SetMultaiTarget(target)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(targetWrapper, resourceData, meta); err != nil {
			return err
		}
	}

	return nil
}

func (res *MultaiTargetTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *multai.Target, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	targetWrapper := NewMultaiTargetWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(targetWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, targetWrapper.GetMultaiTarget(), nil
}

func NewMultaiTargetWrapper() *MultaiTargetWrapper {
	return &MultaiTargetWrapper{
		target: &multai.Target{},
	}
}

func (targetWrapper *MultaiTargetWrapper) GetMultaiTarget() *multai.Target {
	return targetWrapper.target
}

func (targetWrapper *MultaiTargetWrapper) SetMultaiTarget(target *multai.Target) {
	targetWrapper.target = target
}
