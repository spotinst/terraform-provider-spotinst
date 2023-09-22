package organization_programmatic_user

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/organization"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.OrganizationProgrammaticUser,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgProgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgProgUserWrapper.GetOrgProgUser()
			var value *string = nil
			if orgProgUser.Name != nil {
				value = orgProgUser.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgProgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgProgUserWrapper.GetOrgProgUser()
			if v, ok := resourceData.GetOk(string(Name)); ok && v != "" {
				orgProgUser.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgProgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgProgUserWrapper.GetOrgProgUser()
			if v, ok := resourceData.GetOk(string(Name)); ok && v != "" {
				orgProgUser.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Description] = commons.NewGenericField(
		commons.OrganizationProgrammaticUser,
		Description,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgProgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgProgUserWrapper.GetOrgProgUser()
			var value *string = nil
			if orgProgUser.Description != nil {
				value = orgProgUser.Description
			}
			if err := resourceData.Set(string(Description), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Description), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgProgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgProgUserWrapper.GetOrgProgUser()
			if v, ok := resourceData.GetOk(string(Description)); ok && v != "" {
				orgProgUser.SetDescription(spotinst.String(resourceData.Get(string(Description)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgProgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgProgUserWrapper.GetOrgProgUser()
			if v, ok := resourceData.GetOk(string(Description)); ok && v != "" {
				orgProgUser.SetDescription(spotinst.String(resourceData.Get(string(Description)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Policies] = commons.NewGenericField(
		commons.OrganizationProgrammaticUser,
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
			orgProgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgProgUserWrapper.GetOrgProgUser()
			var result []interface{} = nil
			if orgProgUser.Policies != nil {
				policies := orgProgUser.Policies
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
			orgProgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgProgUserWrapper.GetOrgProgUser()
			if v, ok := resourceData.GetOk(string(Policies)); ok {
				if policies, err := expandPolicies(v); err != nil {
					return err
				} else {
					orgProgUser.SetPolicies(policies)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgProgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgProgUserWrapper.GetOrgProgUser()
			var value []*organization.ProgPolicy = nil
			if v, ok := resourceData.GetOk(string(Policies)); ok {
				if policies, err := expandPolicies(v); err != nil {
					return err
				} else {
					value = policies
				}
			}
			orgProgUser.SetPolicies(value)
			return nil
		},
		nil,
	)

	fieldsMap[Accounts] = commons.NewGenericField(
		commons.OrganizationProgrammaticUser,
		Accounts,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(AccountId): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(AccountRole): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgProgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgProgUserWrapper.GetOrgProgUser()
			var result []interface{} = nil
			if orgProgUser.Accounts != nil {
				accounts := orgProgUser.Accounts
				result = flattenAccounts(accounts)
			}
			if result != nil {
				if err := resourceData.Set(string(Accounts), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Accounts), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgProgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgProgUserWrapper.GetOrgProgUser()
			if v, ok := resourceData.GetOk(string(Accounts)); ok {
				if accounts, err := expandAccounts(v); err != nil {
					return err
				} else {
					orgProgUser.SetAccounts(accounts)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgProgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgProgUserWrapper.GetOrgProgUser()
			var value []*organization.Account = nil
			if v, ok := resourceData.GetOk(string(Accounts)); ok {
				if accounts, err := expandAccounts(v); err != nil {
					return err
				} else {
					value = accounts
				}
			}
			orgProgUser.SetAccounts(value)
			return nil
		},
		nil,
	)

	fieldsMap[UserGroupIds] = commons.NewGenericField(
		commons.OrganizationUser,
		UserGroupIds,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgUserWrapper.GetOrgProgUser()
			var value []string = nil
			if orgProgUser.UserGroupIds != nil {
				value = orgProgUser.UserGroupIds
			}
			if value != nil {
				if err := resourceData.Set(string(UserGroupIds), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UserGroupIds), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgProgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgProgUserWrapper.GetOrgProgUser()
			if value, ok := resourceData.GetOk(string(UserGroupIds)); ok && value != nil {
				if userGroupIds, err := expandUserGroupIds(value); err != nil {
					return err
				} else {
					orgProgUser.SetProgUserGroupIds(userGroupIds)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgProgUserWrapper := resourceObject.(*commons.OrgProgUserWrapper)
			orgProgUser := orgProgUserWrapper.GetOrgProgUser()
			if value, ok := resourceData.GetOk(string(UserGroupIds)); ok && value != nil {
				if userGroupIds, err := expandUserGroupIds(value); err != nil {
					return err
				} else {
					orgProgUser.SetProgUserGroupIds(userGroupIds)
				}
			}
			return nil
		},
		nil,
	)
}

func expandPolicies(data interface{}) ([]*organization.ProgPolicy, error) {
	list := data.(*schema.Set).List()

	if list != nil && list[0] != nil {
		ifaces := make([]*organization.ProgPolicy, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			iface := &organization.ProgPolicy{}

			if v, ok := m[string(PolicyAccountIds)]; ok && v != nil {
				accountIdsList := v.([]interface{})
				accountIds := make([]string, len(accountIdsList))
				for i, j := range accountIdsList {
					accountIds[i] = j.(string)
				}
				if accountIds != nil {
					iface.SetAccountIds(accountIds)
				}
			}

			if v, ok := m[string(PolicyId)].(string); ok && v != "" {
				iface.SetPolicyId(spotinst.String(v))
			}

			ifaces = append(ifaces, iface)
		}
		return ifaces, nil
	}
	return nil, nil
}

func flattenPolicies(policies []*organization.ProgPolicy) []interface{} {
	result := make([]interface{}, 0, len(policies))

	for _, policy := range policies {
		m := make(map[string]interface{})
		m[string(PolicyId)] = spotinst.StringValue(policy.PolicyId)
		m[string(PolicyAccountIds)] = policy.AccountIds
		result = append(result, m)
	}

	return result
}

func expandAccounts(data interface{}) ([]*organization.Account, error) {
	list := data.(*schema.Set).List()

	if list != nil && list[0] != nil {
		ifaces := make([]*organization.Account, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			iface := &organization.Account{}

			if v, ok := m[string(AccountId)].(string); ok && v != "" {
				iface.SetAccountId(spotinst.String(v))
			}

			if v, ok := m[string(AccountRole)].(string); ok && v != "" {
				iface.SetRole(spotinst.String(v))
			}

			ifaces = append(ifaces, iface)
		}
		return ifaces, nil
	}
	return nil, nil
}

func flattenAccounts(accounts []*organization.Account) []interface{} {
	result := make([]interface{}, 0, len(accounts))

	for _, account := range accounts {
		m := make(map[string]interface{})
		m[string(AccountId)] = spotinst.StringValue(account.Id)
		m[string(AccountRole)] = spotinst.StringValue(account.Role)
		result = append(result, m)
	}

	return result
}

func expandUserGroupIds(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if userGroupId, ok := v.(string); ok && userGroupId != "" {
			result = append(result, userGroupId)
		}
	}
	return result, nil
}
