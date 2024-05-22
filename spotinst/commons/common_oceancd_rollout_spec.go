package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
)

const (
	OceanCDRolloutSpecResourceName ResourceName = "spotinst_oceancd_rollout_spec"
)

var OceanCDRolloutSpecResource *OceanCDRolloutSpecTerraformResource

type OceanCDRolloutSpecTerraformResource struct {
	GenericResource
}

type OceanCDRolloutSpecWrapper struct {
	rolloutSpec *oceancd.RolloutSpec
}

func NewOceanCDRolloutSpecResource(fieldsMap map[FieldName]*GenericField) *OceanCDRolloutSpecTerraformResource {
	return &OceanCDRolloutSpecTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanCDRolloutSpecResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanCDRolloutSpecTerraformResource) OnRead(
	rolloutSpec *oceancd.RolloutSpec,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	oceancdRolloutSpecWrapper := NewOceanCDRolloutSpecWrapper()
	oceancdRolloutSpecWrapper.SetRolloutSpec(rolloutSpec)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(oceancdRolloutSpecWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *OceanCDRolloutSpecTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*oceancd.RolloutSpec, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	oceancdRolloutSpecWrapper := NewOceanCDRolloutSpecWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(oceancdRolloutSpecWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return oceancdRolloutSpecWrapper.GetRolloutSpec(), nil
}

func (res *OceanCDRolloutSpecTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *oceancd.RolloutSpec, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	oceancdRolloutSpecWrapper := NewOceanCDRolloutSpecWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(oceancdRolloutSpecWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, oceancdRolloutSpecWrapper.GetRolloutSpec(), nil
}

func NewOceanCDRolloutSpecWrapper() *OceanCDRolloutSpecWrapper {
	return &OceanCDRolloutSpecWrapper{
		rolloutSpec: &oceancd.RolloutSpec{},
	}
}

func (oceancdRolloutSpecWrapper *OceanCDRolloutSpecWrapper) GetRolloutSpec() *oceancd.RolloutSpec {
	return oceancdRolloutSpecWrapper.rolloutSpec
}

func (oceancdRolloutSpecWrapper *OceanCDRolloutSpecWrapper) SetRolloutSpec(rolloutSpec *oceancd.RolloutSpec) {
	oceancdRolloutSpecWrapper.rolloutSpec = rolloutSpec
}
