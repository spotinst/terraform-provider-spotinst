package commons

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/organization"
	"log"
)

const (
	OrgProgUserResourceName ResourceName = "spotinst_organization_programmatic_user"
)

var OrgProgUserResource *OrgProgUserTerraformResource

type OrgProgUserTerraformResource struct {
	GenericResource
}

type OrgProgUserWrapper struct {
	orgProgUser *organization.ProgrammaticUser
}

func NewOrgProgUserResource(fieldsMap map[FieldName]*GenericField) *OrgProgUserTerraformResource {
	return &OrgProgUserTerraformResource{
		GenericResource: GenericResource{
			resourceName: OrgProgUserResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OrgProgUserTerraformResource) OnRead(
	orgProgUser *organization.ProgrammaticUser,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	orgProgUserWrapper := NewOrgProgUserWrapper()
	orgProgUserWrapper.SetOrgProgUser(orgProgUser)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(orgProgUserWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *OrgProgUserTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*organization.ProgrammaticUser, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	orgProgUserWrapper := NewOrgProgUserWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(orgProgUserWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return orgProgUserWrapper.GetOrgProgUser(), nil
}

func (res *OrgProgUserTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *organization.ProgrammaticUser, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	orgProgUserWrapper := NewOrgProgUserWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(orgProgUserWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, orgProgUserWrapper.GetOrgProgUser(), nil
}

func NewOrgProgUserWrapper() *OrgProgUserWrapper {
	return &OrgProgUserWrapper{
		orgProgUser: &organization.ProgrammaticUser{},
	}
}

func (orgProgUserWrapper *OrgProgUserWrapper) GetOrgProgUser() *organization.ProgrammaticUser {
	return orgProgUserWrapper.orgProgUser
}

func (orgProgUserWrapper *OrgProgUserWrapper) SetOrgProgUser(orgProgUser *organization.ProgrammaticUser) {
	orgProgUserWrapper.orgProgUser = orgProgUser
}
