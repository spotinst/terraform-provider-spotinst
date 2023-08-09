package administration_org_user

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Email] = commons.NewGenericField(
		commons.AdministrationOrgUser,
		Email,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			var value *string = nil
			if orgUser.Email != nil {
				value = orgUser.Email
			}
			if err := resourceData.Set(string(Email), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Email), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			if v, ok := resourceData.GetOk(string(Email)); ok && v != "" {
				orgUser.SetEmail(spotinst.String(resourceData.Get(string(Email)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			if v, ok := resourceData.GetOk(string(Email)); ok && v != "" {
				orgUser.SetEmail(spotinst.String(resourceData.Get(string(Email)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[FirstName] = commons.NewGenericField(
		commons.AdministrationOrgUser,
		FirstName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			var value *string = nil
			if orgUser.FirstName != nil {
				value = orgUser.FirstName
			}
			if err := resourceData.Set(string(FirstName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(FirstName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			if v, ok := resourceData.GetOk(string(FirstName)); ok && v != "" {
				orgUser.SetFirstName(spotinst.String(resourceData.Get(string(FirstName)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			if v, ok := resourceData.GetOk(string(FirstName)); ok && v != "" {
				orgUser.SetFirstName(spotinst.String(resourceData.Get(string(FirstName)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[LastName] = commons.NewGenericField(
		commons.AdministrationOrgUser,
		LastName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			var value *string = nil
			if orgUser.LastName != nil {
				value = orgUser.LastName
			}
			if err := resourceData.Set(string(LastName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(LastName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			if v, ok := resourceData.GetOk(string(LastName)); ok && v != "" {
				orgUser.SetLastName(spotinst.String(resourceData.Get(string(LastName)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			if v, ok := resourceData.GetOk(string(LastName)); ok && v != "" {
				orgUser.SetLastName(spotinst.String(resourceData.Get(string(LastName)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Password] = commons.NewGenericField(
		commons.AdministrationOrgUser,
		Password,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			var value *string = nil
			if orgUser.Password != nil {
				value = orgUser.Password
			}
			if err := resourceData.Set(string(Password), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Password), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			if v, ok := resourceData.GetOk(string(Password)); ok && v != "" {
				orgUser.SetFirstName(spotinst.String(resourceData.Get(string(Password)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			if v, ok := resourceData.GetOk(string(Password)); ok && v != "" {
				orgUser.SetFirstName(spotinst.String(resourceData.Get(string(Password)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Role] = commons.NewGenericField(
		commons.AdministrationOrgUser,
		Role,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			var value *string = nil
			if orgUser.Role != nil {
				value = orgUser.Role
			}
			if err := resourceData.Set(string(Role), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Role), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			if v, ok := resourceData.GetOk(string(Role)); ok && v != "" {
				orgUser.SetFirstName(spotinst.String(resourceData.Get(string(Role)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			if v, ok := resourceData.GetOk(string(Role)); ok && v != "" {
				orgUser.SetFirstName(spotinst.String(resourceData.Get(string(Role)).(string)))
			}
			return nil
		},
		nil,
	)

}
