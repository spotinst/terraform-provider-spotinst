package administration_org_programmatic_user

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/administration"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.AdministrationOrgProgrammaticUser,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
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
		commons.AdministrationOrgProgrammaticUser,
		Description,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
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
		commons.AdministrationOrgProgrammaticUser,
		Policies,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			ForceNew: true,
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
			var value []*administration.ProgPolicy = nil
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
		commons.AdministrationOrgProgrammaticUser,
		Accounts,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			ForceNew: true,
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
			var value []*administration.Account = nil
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
}

func expandPolicies(data interface{}) ([]*administration.ProgPolicy, error) {
	list := data.(*schema.Set).List()

	if list != nil && list[0] != nil {
		ifaces := make([]*administration.ProgPolicy, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			iface := &administration.ProgPolicy{}

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

func flattenPolicies(policies []*administration.ProgPolicy) []interface{} {
	result := make([]interface{}, 0, len(policies))

	for _, policy := range policies {
		m := make(map[string]interface{})
		m[string(PolicyId)] = spotinst.StringValue(policy.PolicyId)
		m[string(PolicyAccountIds)] = policy.AccountIds
		result = append(result, m)
	}

	return result
}

func expandAccounts(data interface{}) ([]*administration.Account, error) {
	list := data.(*schema.Set).List()

	if list != nil && list[0] != nil {
		ifaces := make([]*administration.Account, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			iface := &administration.Account{}

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

func flattenAccounts(accounts []*administration.Account) []interface{} {
	result := make([]interface{}, 0, len(accounts))

	for _, account := range accounts {
		m := make(map[string]interface{})
		m[string(AccountId)] = spotinst.StringValue(account.Id)
		m[string(AccountRole)] = spotinst.StringValue(account.Role)
		result = append(result, m)
	}

	return result
}
