package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
)

const (
	OceanAWSRightSizingRuleResourceName ResourceName = "spotinst_ocean_aws_right_sizing_rule"
)

var OceanAWSRightSizingRuleResource *OceanAWSRightSizingRuleTerraformResource

type OceanAWSRightSizingRuleTerraformResource struct {
	GenericResource
}

type RightSizingRuleWrapper struct {
	rightSizingRule *aws.RightSizingRule
}

// NewOceanAWSRightSizingRuleResource creates a new OceanAWSRightRuleSizing resource
func NewOceanAWSRightSizingRuleResource(fieldMap map[FieldName]*GenericField) *OceanAWSRightSizingRuleTerraformResource {
	return &OceanAWSRightSizingRuleTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanAWSRightSizingRuleResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

// OnCreate is called when creating a new resource block and returns a new OceanAWSExtendedResourceDefinition or an error.
func (res *OceanAWSRightSizingRuleTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.RightSizingRule, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	rsrWrapper := NewOceanAWSRightSizingRuleWrapper()
	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(rsrWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return rsrWrapper.GetOceanAWSRightSizingRule(), nil
}

// OnRead is called when reading an existing resource and throws an error if it is unable to do so.
func (res *OceanAWSRightSizingRuleTerraformResource) OnRead(
	rightSizingRule *aws.RightSizingRule,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	rsrWrapper := NewOceanAWSRightSizingRuleWrapper()
	rsrWrapper.SetOceanAWSRightSizingRule(rightSizingRule)

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
// an extendedResourceDefinition with a bool indicating if had been updated, or an error.
func (res *OceanAWSRightSizingRuleTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *aws.RightSizingRule, error) {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	rsrWrapper := NewOceanAWSRightSizingRuleWrapper()
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

	return hasChanged, rsrWrapper.GetOceanAWSRightSizingRule(), nil
}

// Spotinst RightSizingRule must have a wrapper struct.
// the wrapper struct is intended to help reflect the field states into the RightSizingRule object properly.
func NewOceanAWSRightSizingRuleWrapper() *RightSizingRuleWrapper {
	return &RightSizingRuleWrapper{
		rightSizingRule: &aws.RightSizingRule{},
	}
}

// GetOceanAWSRightSizingRule returns a wrapped OceanAWSRightSizingRule
func (rsrWrapper *RightSizingRuleWrapper) GetOceanAWSRightSizingRule() *aws.RightSizingRule {
	return rsrWrapper.rightSizingRule
}

// SetOceanAWSExtendedResourceDefinition applies extendedResourceDefinition fields to the extendedResourceDefinition wrapper.
func (rsrWrapper *RightSizingRuleWrapper) SetOceanAWSRightSizingRule(rsr *aws.RightSizingRule) {
	rsrWrapper.rightSizingRule = rsr
}
