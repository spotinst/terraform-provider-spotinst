package commons

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/right_sizing"
	"log"
)

const (
	OceanRightSizingRuleResourceName ResourceName = "spotinst_ocean_right_sizing_rule"
)

var OceanRightSizingRuleResource *OceanRightSizingRuleTerraformResource

type OceanRightSizingRuleTerraformResource struct {
	GenericResource
}

type RightSizingRuleWrapper struct {
	rightSizingRule *right_sizing.RightsizingRule
}

// NewOceanRightSizingRuleResource creates a new OceanRightRuleSizing resource
func NewOceanRightSizingRuleResource(fieldMap map[FieldName]*GenericField) *OceanRightSizingRuleTerraformResource {
	return &OceanRightSizingRuleTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanRightSizingRuleResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

// OnCreate is called when creating a new resource block and returns a new OceanRightSizingRuleResourceDefinition or an error.
func (res *OceanRightSizingRuleTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*right_sizing.RightsizingRule, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	rsrWrapper := NewOceanRightSizingRuleWrapper()
	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(rsrWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return rsrWrapper.GetOceanRightSizingRule(), nil
}

// OnRead is called when reading an existing resource and throws an error if it is unable to do so.
func (res *OceanRightSizingRuleTerraformResource) OnRead(
	rightSizingRule *right_sizing.RightsizingRule,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	rsrWrapper := NewOceanRightSizingRuleWrapper()
	rsrWrapper.SetOceanRightSizingRule(rightSizingRule)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(rsrWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

// OnUpdate is called when updating an existing resource and returns
// an rightSizingResource with a bool indicating if had been updated, or an error.
func (res *OceanRightSizingRuleTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *right_sizing.RightsizingRule, error) {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	rsrWrapper := NewOceanRightSizingRuleWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(rsrWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, rsrWrapper.GetOceanRightSizingRule(), nil
}

// NewOceanRightSizingRuleWrapper Spotinst RightSizingRule must have a wrapper struct.
// the wrapper struct is intended to help reflect the field states into the RightSizingRule object properly.
func NewOceanRightSizingRuleWrapper() *RightSizingRuleWrapper {
	return &RightSizingRuleWrapper{
		rightSizingRule: &right_sizing.RightsizingRule{},
	}
}

// GetOceanRightSizingRule returns a wrapped OceanRightSizingRule
func (rsrWrapper *RightSizingRuleWrapper) GetOceanRightSizingRule() *right_sizing.RightsizingRule {
	return rsrWrapper.rightSizingRule
}

// SetOceanRightSizingRule  applies rightSizingRule fields to the rightSizingRule wrapper.
func (rsrWrapper *RightSizingRuleWrapper) SetOceanRightSizingRule(rsr *right_sizing.RightsizingRule) {
	rsrWrapper.rightSizingRule = rsr
}
