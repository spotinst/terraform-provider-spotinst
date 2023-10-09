package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/organization"
)

const (
	OrgProgrammaticUserResourceName ResourceName = "spotinst_organization_programmatic_user"
)

var OrgProgrammaticUserResource *OrgProgrammaticUserTerraformResource

type OrgProgrammaticUserTerraformResource struct {
	GenericResource
}

type OrgProgrammaticUserWrapper struct {
	orgProgrammaticUser *organization.ProgrammaticUser
}

func NewOrgProgrammaticUserResource(fieldsMap map[FieldName]*GenericField) *OrgProgrammaticUserTerraformResource {
	return &OrgProgrammaticUserTerraformResource{
		GenericResource: GenericResource{
			resourceName: OrgProgrammaticUserResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OrgProgrammaticUserTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*organization.ProgrammaticUser, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	orgProgrammaticUserWrapper := NewOrgProgrammaticUserWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(orgProgrammaticUserWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return orgProgrammaticUserWrapper.GetOrgProgrammaticUser(), nil
}

func (res *OrgProgrammaticUserTerraformResource) OnRead(
	orgProgrammaticUser *organization.ProgrammaticUser,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	orgProgrammaticUserWrapper := NewOrgProgrammaticUserWrapper()
	orgProgrammaticUserWrapper.SetOrgProgrammaticUser(orgProgrammaticUser)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(orgProgrammaticUserWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *OrgProgrammaticUserTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *organization.ProgrammaticUser, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	orgProgrammaticUserWrapper := NewOrgProgrammaticUserWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(orgProgrammaticUserWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, orgProgrammaticUserWrapper.GetOrgProgrammaticUser(), nil
}

func NewOrgProgrammaticUserWrapper() *OrgProgrammaticUserWrapper {
	return &OrgProgrammaticUserWrapper{
		orgProgrammaticUser: &organization.ProgrammaticUser{},
	}
}

func (orgProgrammaticUserWrapper *OrgProgrammaticUserWrapper) GetOrgProgrammaticUser() *organization.ProgrammaticUser {
	return orgProgrammaticUserWrapper.orgProgrammaticUser
}

func (orgProgrammaticUserWrapper *OrgProgrammaticUserWrapper) SetOrgProgrammaticUser(orgProgrammaticUser *organization.ProgrammaticUser) {
	orgProgrammaticUserWrapper.orgProgrammaticUser = orgProgrammaticUser
}
