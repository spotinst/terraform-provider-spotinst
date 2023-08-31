package organization_user

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/administration"
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
			Optional: true,
			Computed: true,
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
			Optional: true,
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

	fieldsMap[Policies] = commons.NewGenericField(
		commons.AdministrationOrgProgrammaticUser,
		Policies,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(PolicyAccountIds): {
						Type:     schema.TypeList,
						Required: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					string(PolicyId): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			var result []interface{} = nil
			if orgUser.Policies != nil {
				policies := orgUser.Policies
				result = flattenPolicies(policies)
			}
			if result != nil {
				if err := resourceData.Set(string(Policies), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Policies), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgUser := orgUserWrapper.GetOrgUser()
			if v, ok := resourceData.GetOk(string(Policies)); ok {
				if policies, err := expandPolicies(v); err != nil {
					return err
				} else {
					orgUser.SetUserPolicies(policies)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgUserWrapper)
			orgProgUser := orgUserWrapper.GetOrgUser()
			var value []*administration.UserPolicy = nil
			if v, ok := resourceData.GetOk(string(Policies)); ok {
				if policies, err := expandPolicies(v); err != nil {
					return err
				} else {
					value = policies
				}
			}
			orgProgUser.SetUserPolicies(value)
			return nil
		},
		nil,
	)

}

func expandPolicies(data interface{}) ([]*administration.UserPolicy, error) {
	list := data.(*schema.Set).List()

	if list != nil && list[0] != nil {
		ifaces := make([]*administration.UserPolicy, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			iface := &administration.UserPolicy{}

			if v, ok := m[string(PolicyAccountIds)]; ok && v != nil {
				accountIdsList := v.([]interface{})
				accountIds := make([]string, len(accountIdsList))
				for i, j := range accountIdsList {
					accountIds[i] = j.(string)
				}

				if accountIds != nil {
					iface.SetUserPolicyAccountIds(accountIds)
				}
			}

			if v, ok := m[string(PolicyId)].(string); ok && v != "" {
				iface.SetUserPolicyId(spotinst.String(v))
			}

			ifaces = append(ifaces, iface)
		}
		return ifaces, nil
	}
	return nil, nil
}

func flattenPolicies(policies []*administration.UserPolicy) []interface{} {
	result := make([]interface{}, 0, len(policies))

	for _, policy := range policies {
		m := make(map[string]interface{})
		m[string(PolicyId)] = spotinst.StringValue(policy.PolicyId)
		m[string(PolicyAccountIds)] = policy.AccountIds
		result = append(result, m)
	}

	return result
}
