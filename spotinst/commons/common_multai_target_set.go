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
	MultaiTargetSetResourceName ResourceName = "spotinst_multai_target_set"
)

var MultaiTargetSetResource *MultaiTargetSetTerraformResource

type MultaiTargetSetTerraformResource struct {
	GenericResource // embedding
}

type MultaiTargetSetWrapper struct {
	targetSet *multai.TargetSet
}

func NewMultaiTargetSetResource(fieldMap map[FieldName]*GenericField) *MultaiTargetSetTerraformResource {
	return &MultaiTargetSetTerraformResource{
		GenericResource: GenericResource{
			resourceName: MultaiTargetSetResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

func (res *MultaiTargetSetTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*multai.TargetSet, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	targetSetWrapper := NewMultaiTargetSetWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(targetSetWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return targetSetWrapper.GetMultaiTargetSet(), nil
}

func (res *MultaiTargetSetTerraformResource) OnRead(
	targetSet *multai.TargetSet,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	targetSetWrapper := NewMultaiTargetSetWrapper()
	targetSetWrapper.SetMultaiTargetSet(targetSet)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(targetSetWrapper, resourceData, meta); err != nil {
			return err
		}
	}

	return nil
}

func (res *MultaiTargetSetTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *multai.TargetSet, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	targetSetWrapper := NewMultaiTargetSetWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(targetSetWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, targetSetWrapper.GetMultaiTargetSet(), nil
}

func NewMultaiTargetSetWrapper() *MultaiTargetSetWrapper {
	return &MultaiTargetSetWrapper{
		targetSet: &multai.TargetSet{},
	}
}

func (targetSetWrapper *MultaiTargetSetWrapper) GetMultaiTargetSet() *multai.TargetSet {
	return targetSetWrapper.targetSet
}

func (targetSetWrapper *MultaiTargetSetWrapper) SetMultaiTargetSet(targetSet *multai.TargetSet) {
	targetSetWrapper.targetSet = targetSet
}
