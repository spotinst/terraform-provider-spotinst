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
	MultaiRoutingRuleResourceName ResourceName = "spotinst_multai_routing_rule"
)

var MultaiRoutingRuleResource *MultaiRoutingRuleTerraformResource

type MultaiRoutingRuleTerraformResource struct {
	GenericResource // embedding
}

type MultaiRoutingRuleWrapper struct {
	routingRule *multai.RoutingRule
}

func NewMultaiRoutingRuleResource(fieldMap map[FieldName]*GenericField) *MultaiRoutingRuleTerraformResource {
	return &MultaiRoutingRuleTerraformResource{
		GenericResource: GenericResource{
			resourceName: MultaiRoutingRuleResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

func (res *MultaiRoutingRuleTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*multai.RoutingRule, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	mlbWrapper := NewMultaiRoutingRuleWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(mlbWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return mlbWrapper.GetMultaiRoutingRule(), nil
}

func (res *MultaiRoutingRuleTerraformResource) OnRead(
	routingRule *multai.RoutingRule,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	mlbWrapper := NewMultaiRoutingRuleWrapper()
	mlbWrapper.SetMultaiRoutingRule(routingRule)

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

func (res *MultaiRoutingRuleTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *multai.RoutingRule, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	mlbWrapper := NewMultaiRoutingRuleWrapper()
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

	return hasChanged, mlbWrapper.GetMultaiRoutingRule(), nil
}

func NewMultaiRoutingRuleWrapper() *MultaiRoutingRuleWrapper {
	return &MultaiRoutingRuleWrapper{
		routingRule: &multai.RoutingRule{},
	}
}

func (mlbWrapper *MultaiRoutingRuleWrapper) GetMultaiRoutingRule() *multai.RoutingRule {
	return mlbWrapper.routingRule
}

func (mlbWrapper *MultaiRoutingRuleWrapper) SetMultaiRoutingRule(routingRule *multai.RoutingRule) {
	mlbWrapper.routingRule = routingRule
}
