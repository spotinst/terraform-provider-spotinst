package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
)

const (
	OceanCDStrategyResourceName ResourceName = "spotinst_oceancd_strategy"
)

var OceanCDStrategyResource *OceanCDStrategyTerraformResource

type OceanCDStrategyTerraformResource struct {
	GenericResource
}

type OceanCDStrategyWrapper struct {
	Strategy *oceancd.Strategy
}

func NewOceanCDStrategyResource(fieldsMap map[FieldName]*GenericField) *OceanCDStrategyTerraformResource {
	return &OceanCDStrategyTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanCDStrategyResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanCDStrategyTerraformResource) OnRead(
	Strategy *oceancd.Strategy,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	oceancdStrategyWrapper := NewOceanCDStrategyWrapper()
	oceancdStrategyWrapper.SetStrategy(Strategy)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(oceancdStrategyWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *OceanCDStrategyTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*oceancd.Strategy, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	oceancdStrategyWrapper := NewOceanCDStrategyWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(oceancdStrategyWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return oceancdStrategyWrapper.GetStrategy(), nil
}

func (res *OceanCDStrategyTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *oceancd.Strategy, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	oceancdStrategyWrapper := NewOceanCDStrategyWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(oceancdStrategyWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, oceancdStrategyWrapper.GetStrategy(), nil
}

func NewOceanCDStrategyWrapper() *OceanCDStrategyWrapper {
	return &OceanCDStrategyWrapper{
		Strategy: &oceancd.Strategy{},
	}
}

func (oceancdStrategyWrapper *OceanCDStrategyWrapper) GetStrategy() *oceancd.Strategy {
	return oceancdStrategyWrapper.Strategy
}

func (oceancdStrategyWrapper *OceanCDStrategyWrapper) SetStrategy(Strategy *oceancd.Strategy) {
	oceancdStrategyWrapper.Strategy = Strategy
}
