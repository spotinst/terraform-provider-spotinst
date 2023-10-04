package commons

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/organization"
	"log"
)

const (
	OrgUserResourceName ResourceName = "spotinst_organization_user"
)

var OrgUserResource *OrgUserTerraformResource

type OrgUserTerraformResource struct {
	GenericResource
}

type OrgUserWrapper struct {
	orgUser *organization.User
}

func NewOrgUserResource(fieldsMap map[FieldName]*GenericField) *OrgUserTerraformResource {
	return &OrgUserTerraformResource{
		GenericResource: GenericResource{
			resourceName: OrgUserResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OrgUserTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*organization.User, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	orgUserWrapper := NewOrgUserWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(orgUserWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return orgUserWrapper.GetOrgUser(), nil
}

func (res *OrgUserTerraformResource) OnRead(
	orgUser *organization.User,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	orgUserWrapper := NewOrgUserWrapper()
	orgUserWrapper.SetOrgUser(orgUser)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(orgUserWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *OrgUserTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *organization.User, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	orgUserWrapper := NewOrgUserWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(orgUserWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, orgUserWrapper.GetOrgUser(), nil
}

func NewOrgUserWrapper() *OrgUserWrapper {
	return &OrgUserWrapper{
		orgUser: &organization.User{},
	}
}

func (orgUserWrapper *OrgUserWrapper) GetOrgUser() *organization.User {
	return orgUserWrapper.orgUser
}

func (orgUserWrapper *OrgUserWrapper) SetOrgUser(orgUser *organization.User) {
	orgUserWrapper.orgUser = orgUser
}
