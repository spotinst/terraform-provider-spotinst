package administration_org_policy

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/administration"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.AdministrationOrgPolicy,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgPolicyWrapper := resourceObject.(*commons.OrgPolicyWrapper)
			orgPolicy := orgPolicyWrapper.GetOrgPolicy()
			var value *string = nil
			if orgPolicy.Name != nil {
				value = orgPolicy.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgPolicyWrapper := resourceObject.(*commons.OrgPolicyWrapper)
			orgPolicy := orgPolicyWrapper.GetOrgPolicy()
			if v, ok := resourceData.GetOk(string(Name)); ok && v != "" {
				orgPolicy.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgPolicyWrapper := resourceObject.(*commons.OrgPolicyWrapper)
			orgPolicy := orgPolicyWrapper.GetOrgPolicy()
			if v, ok := resourceData.GetOk(string(Name)); ok && v != "" {
				orgPolicy.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Description] = commons.NewGenericField(
		commons.AdministrationOrgPolicy,
		Description,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgPolicyWrapper := resourceObject.(*commons.OrgPolicyWrapper)
			orgPolicy := orgPolicyWrapper.GetOrgPolicy()
			var value *string = nil
			if orgPolicy.Description != nil {
				value = orgPolicy.Description
			}
			if err := resourceData.Set(string(Description), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Description), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgPolicyWrapper := resourceObject.(*commons.OrgPolicyWrapper)
			orgPolicy := orgPolicyWrapper.GetOrgPolicy()
			if v, ok := resourceData.GetOk(string(Description)); ok && v != "" {
				orgPolicy.SetDescription(spotinst.String(resourceData.Get(string(Description)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgPolicyWrapper := resourceObject.(*commons.OrgPolicyWrapper)
			orgPolicy := orgPolicyWrapper.GetOrgPolicy()
			if v, ok := resourceData.GetOk(string(Description)); ok && v != "" {
				orgPolicy.SetDescription(spotinst.String(resourceData.Get(string(Description)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[PolicyContent] = commons.NewGenericField(
		commons.AdministrationOrgPolicy,
		PolicyContent,
		&schema.Schema{
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Statements): {
						Type:     schema.TypeSet,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Actions): {
									Type:     schema.TypeList,
									Required: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},

								string(Effect): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(Resources): {
									Type:     schema.TypeList,
									Required: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgPolicyWrapper := resourceObject.(*commons.OrgPolicyWrapper)
			orgPolicy := orgPolicyWrapper.GetOrgPolicy()
			var result []interface{} = nil
			if orgPolicy.PolicyContent != nil {
				policyContent := orgPolicy.PolicyContent
				result = flattenPolicyContent(policyContent)
			}
			if result != nil {
				if err := resourceData.Set(string(PolicyContent), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PolicyContent), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgPolicyWrapper := resourceObject.(*commons.OrgPolicyWrapper)
			orgPolicy := orgPolicyWrapper.GetOrgPolicy()
			if v, ok := resourceData.GetOk(string(PolicyContent)); ok {
				if policyContent, err := expandPolicyContent(v); err != nil {
					return err
				} else {
					orgPolicy.SetPolicyContent(policyContent)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			orgPolicyWrapper := resourceObject.(*commons.OrgPolicyWrapper)
			orgPolicy := orgPolicyWrapper.GetOrgPolicy()
			var value *administration.PolicyContent = nil
			if v, ok := resourceData.GetOk(string(PolicyContent)); ok {
				if policyContent, err := expandPolicyContent(v); err != nil {
					return err
				} else {
					value = policyContent
				}
			}
			orgPolicy.SetPolicyContent(value)
			return nil
		},
		nil,
	)
}

func expandPolicyContent(data interface{}) (*administration.PolicyContent, error) {
	list := data.(*schema.Set).List()

	if list != nil && list[0] != nil {
		ifaces := make([]*administration.PolicyContent, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			iface := &administration.PolicyContent{}

			if v, ok := m[string(Statements)]; ok {
				statements, err := expandStatements(v)
				if err != nil {
					return nil, err
				}

				if statements != nil {
					iface.SetStatements(statements)
				}
			} else {
				iface.Statements = nil
			}

			ifaces = append(ifaces, iface)
		}
		return ifaces[0], nil
	}
	return nil, nil
}

// expandStatements sets the values from the plan as objects
func expandStatements(data interface{}) ([]*administration.Statement, error) {
	list := data.(*schema.Set).List()
	statements := make([]*administration.Statement, 0, len(list))

	for _, item := range list {
		attr := item.(map[string]interface{})

		statement := &administration.Statement{}

		if v, ok := attr[string(Actions)]; ok {
			actionsList := v.([]interface{})
			actions := make([]string, len(actionsList))
			for i, j := range actionsList {
				actions[i] = j.(string)
			}
			statement.SetActions(actions)
		}

		if v, ok := attr[string(Effect)].(string); ok && v != "" {
			statement.SetEffect(spotinst.String(v))
		}

		if v, ok := attr[string(Resources)]; ok {
			resourcesList := v.([]interface{})
			resources := make([]string, len(resourcesList))
			for i, j := range resourcesList {
				resources[i] = j.(string)
			}
			statement.SetResources(resources)
		}

		statements = append(statements, statement)
	}
	return statements, nil
}

func flattenPolicyContent(policyContent *administration.PolicyContent) []interface{} {
	result := make(map[string]interface{})
	result[string(Statements)] = flattenStatements(policyContent.Statements)
	return []interface{}{result}
}

func flattenStatements(statements []*administration.Statement) []interface{} {
	result := make([]interface{}, 0, len(statements))

	for _, statement := range statements {
		m := make(map[string]interface{})
		m[string(Actions)] = statement.Actions
		m[string(Effect)] = spotinst.StringValue(statement.Effect)
		m[string(Resources)] = statement.Resources
		result = append(result, m)
	}

	return result
}
