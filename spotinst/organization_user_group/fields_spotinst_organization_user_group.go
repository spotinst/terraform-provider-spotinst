package organization_user_group

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/organization"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.OrganizationUserGroup,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserGroupWrapper := resourceObject.(*commons.OrgUserGroupWrapper)
			orgUserGroup := orgUserGroupWrapper.GetOrgUserGroup()
			var value *string = nil
			if orgUserGroup.Name != nil {
				value = orgUserGroup.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserGroupWrapper := resourceObject.(*commons.OrgUserGroupWrapper)
			orgUserGroup := orgUserGroupWrapper.GetOrgUserGroup()
			if v, ok := resourceData.GetOk(string(Name)); ok && v != "" {
				orgUserGroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserGroupWrapper := resourceObject.(*commons.OrgUserGroupWrapper)
			orgUserGroup := orgUserGroupWrapper.GetOrgUserGroup()
			if v, ok := resourceData.GetOk(string(Name)); ok && v != "" {
				orgUserGroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Description] = commons.NewGenericField(
		commons.OrganizationUserGroup,
		Description,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserGroupWrapper := resourceObject.(*commons.OrgUserGroupWrapper)
			orgUserGroup := orgUserGroupWrapper.GetOrgUserGroup()
			var value *string = nil
			if orgUserGroup.Description != nil {
				value = orgUserGroup.Description
			}
			if err := resourceData.Set(string(Description), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Description), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserGroupWrapper := resourceObject.(*commons.OrgUserGroupWrapper)
			orgUserGroup := orgUserGroupWrapper.GetOrgUserGroup()
			if v, ok := resourceData.GetOk(string(Description)); ok && v != "" {
				orgUserGroup.SetDescription(spotinst.String(resourceData.Get(string(Description)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserGroupWrapper := resourceObject.(*commons.OrgUserGroupWrapper)
			orgUserGroup := orgUserGroupWrapper.GetOrgUserGroup()
			if v, ok := resourceData.GetOk(string(Description)); ok && v != "" {
				orgUserGroup.SetDescription(spotinst.String(resourceData.Get(string(Description)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[UserIds] = commons.NewGenericField(
		commons.OrganizationUserGroup,
		UserIds,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserGroupWrapper := resourceObject.(*commons.OrgUserGroupWrapper)
			orgUserGroup := orgUserGroupWrapper.GetOrgUserGroup()
			var value []string = nil
			if orgUserGroup.UserIds != nil {
				value = orgUserGroup.UserIds
			}
			if value != nil {
				if err := resourceData.Set(string(UserIds), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UserIds), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserGroupWrapper := resourceObject.(*commons.OrgUserGroupWrapper)
			orgUserGroup := orgUserGroupWrapper.GetOrgUserGroup()
			if value, ok := resourceData.GetOk(string(UserIds)); ok && value != nil {
				if userIds, err := expandUserIds(value); err != nil {
					return err
				} else {
					orgUserGroup.SetUserIds(userIds)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserGroupWrapper := resourceObject.(*commons.OrgUserGroupWrapper)
			orgUserGroup := orgUserGroupWrapper.GetOrgUserGroup()
			if value, ok := resourceData.GetOk(string(UserIds)); ok && value != nil {
				if userIds, err := expandUserIds(value); err != nil {
					return err
				} else {
					orgUserGroup.SetUserIds(userIds)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[Policies] = commons.NewGenericField(
		commons.OrganizationUserGroup,
		Policies,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(AccountIds): {
						Type:     schema.TypeSet,
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
			orgUserGroupWrapper := resourceObject.(*commons.OrgUserGroupWrapper)
			orgUserGroup := orgUserGroupWrapper.GetOrgUserGroup()
			var result []interface{} = nil
			if orgUserGroup.Policies != nil {
				policies := orgUserGroup.Policies
				result = flattenPolicies(policies)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Policies), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Policies), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserGroupWrapper := resourceObject.(*commons.OrgUserGroupWrapper)
			orgUserGroup := orgUserGroupWrapper.GetOrgUserGroup()
			if v, ok := resourceData.GetOk(string(Policies)); ok {
				if policies, err := expandPolicies(v); err != nil {
					return err
				} else {
					orgUserGroup.SetPolicies(policies)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgUserGroupWrapper := resourceObject.(*commons.OrgUserGroupWrapper)
			orgUserGroup := orgUserGroupWrapper.GetOrgUserGroup()
			var value []*organization.UserGroupPolicy = nil
			if v, ok := resourceData.GetOk(string(Policies)); ok {
				if policies, err := expandPolicies(v); err != nil {
					return err
				} else {
					value = policies
				}
			}
			orgUserGroup.SetPolicies(value)
			return nil
		},
		nil,
	)
}

func expandPolicies(data interface{}) ([]*organization.UserGroupPolicy, error) {

	if list := data.([]interface{}); len(list) > 0 {
		ifaces := make([]*organization.UserGroupPolicy, 0, len(list))
		if list != nil && list[0] != nil {
			for _, item := range list {
				m := item.(map[string]interface{})
				iface := &organization.UserGroupPolicy{}

				if v, ok := m[string(AccountIds)]; ok {
					accounts, err := expandAccountIds(v)
					if err != nil {
						return nil, err
					}

					if accounts != nil {
						iface.SetAccountIds(accounts)
					} else {
						iface.SetAccountIds(nil)
					}
				}

				if v, ok := m[string(PolicyId)].(string); ok && v != "" {
					iface.SetPolicyId(spotinst.String(v))
				}
				ifaces = append(ifaces, iface)
			}
		}
		return ifaces, nil
	}
	return nil, nil
}

func expandAccountIds(data interface{}) ([]string, error) {
	list := data.(*schema.Set).List()
	result := make([]string, 0, len(list))

	for _, v := range list {
		if accountId, ok := v.(string); ok && accountId != "" {
			result = append(result, accountId)
		}
	}
	return result, nil
}

func expandUserIds(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if userId, ok := v.(string); ok && userId != "" {
			result = append(result, userId)
		}
	}
	return result, nil
}

func flattenPolicies(policies []*organization.UserGroupPolicy) []interface{} {
	result := make([]interface{}, 0, len(policies))

	for _, policy := range policies {
		m := make(map[string]interface{})

		if policy.AccountIds != nil {
			m[string(AccountIds)] = policy.AccountIds
		}
		m[string(PolicyId)] = spotinst.StringValue(policy.PolicyId)

		result = append(result, m)
	}
	return result
}
