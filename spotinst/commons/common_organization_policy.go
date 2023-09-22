package commons

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/organization"
	"log"
)

const (
	OrgPolicyResourceName ResourceName = "spotinst_organization_policy"
)

var OrgPolicyResource *OrgPolicyTerraformResource

type OrgPolicyTerraformResource struct {
	GenericResource
}

type OrgPolicyWrapper struct {
	OrgPolicy *organization.Policy
}

func NewOrgPolicyResource(fieldsMap map[FieldName]*GenericField) *OrgPolicyTerraformResource {
	return &OrgPolicyTerraformResource{
		GenericResource: GenericResource{
			resourceName: OrgPolicyResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OrgPolicyTerraformResource) OnRead(
	OrgPolicy *organization.Policy,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	OrgPolicyWrapper := NewOrgPolicyWrapper()
	OrgPolicyWrapper.SetOrgPolicy(OrgPolicy)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(OrgPolicyWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *OrgPolicyTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*organization.Policy, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	OrgPolicyWrapper := NewOrgPolicyWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(OrgPolicyWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return OrgPolicyWrapper.GetOrgPolicy(), nil
}

func (res *OrgPolicyTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *organization.Policy, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	OrgPolicyWrapper := NewOrgPolicyWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(OrgPolicyWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, OrgPolicyWrapper.GetOrgPolicy(), nil
}

func NewOrgPolicyWrapper() *OrgPolicyWrapper {
	return &OrgPolicyWrapper{
		OrgPolicy: &organization.Policy{},
	}
}

func (OrgPolicyWrapper *OrgPolicyWrapper) GetOrgPolicy() *organization.Policy {
	return OrgPolicyWrapper.OrgPolicy
}

func (OrgPolicyWrapper *OrgPolicyWrapper) SetOrgPolicy(OrgPolicy *organization.Policy) {
	OrgPolicyWrapper.OrgPolicy = OrgPolicy
}
