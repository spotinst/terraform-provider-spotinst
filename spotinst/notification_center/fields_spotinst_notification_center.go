package notification_center

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/notificationcenter"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.NotificationCenter,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			if err := resourceData.Set(string(Name), nc.Name); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			if v, ok := resourceData.GetOk(string(Name)); ok {
				nc.SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			if v, ok := resourceData.GetOk(string(Name)); ok {
				nc.SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Description] = commons.NewGenericField(
		commons.NotificationCenter,
		Description,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			if err := resourceData.Set(string(Description), nc.Description); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Description), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			if v, ok := resourceData.GetOk(string(Description)); ok {
				nc.SetDescription(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			if v, ok := resourceData.GetOk(string(Description)); ok {
				nc.SetDescription(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[PrivacyLevel] = commons.NewGenericField(
		commons.NotificationCenter,
		PrivacyLevel,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			if err := resourceData.Set(string(PrivacyLevel), nc.PrivacyLevel); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PrivacyLevel), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			if v, ok := resourceData.GetOk(string(PrivacyLevel)); ok {
				nc.SetPrivacyLevel(spotinst.String(v.(string)))
			}
			return nil
		},
		/*func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			if v, ok := resourceData.GetOk(string(PrivacyLevel)); ok {
				nc.SetPrivacyLevel(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,*/

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[IsActive] = commons.NewGenericField(
		commons.NotificationCenter,
		IsActive,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			var value *bool = nil
			if nc != nil {
				value = nc.IsActive
			}
			if value != nil {
				if err := resourceData.Set(string(IsActive), nc.IsActive); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IsActive), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			if v, ok := resourceData.GetOkExists(string(IsActive)); ok && v != nil {
				ia := v.(bool)
				isActive := spotinst.Bool(ia)
				nc.SetIsActive(isActive)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			var isActive *bool = nil
			if v, ok := resourceData.GetOkExists(string(IsActive)); ok && v != nil {
				ia := v.(bool)
				isActive = spotinst.Bool(ia)
			}
			nc.SetIsActive(isActive)
			return nil
		},
		nil,
	)

	fieldsMap[RegisteredUsers] = commons.NewGenericField(
		commons.NotificationCenter,
		RegisteredUsers,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(UserEmail): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(SubscriptionTypes): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			var result []interface{} = nil
			if nc != nil && nc.RegisteredUsers != nil {
				result = flattenRegisteredUsers(nc.RegisteredUsers)
			}

			if len(result) > 0 {
				if err := resourceData.Set(string(RegisteredUsers), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RegisteredUsers), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			var value []*notificationcenter.RegisteredUsers = nil
			if v, ok := resourceData.GetOkExists(string(RegisteredUsers)); ok {
				if registeredUsers, err := expandRegisteredUsers(v); err != nil {
					return err
				} else {
					value = registeredUsers
				}
			}
			nc.SetRegisteredUsers(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			var value []*notificationcenter.RegisteredUsers = nil
			if v, ok := resourceData.GetOkExists(string(RegisteredUsers)); ok {
				if registeredUsers, err := expandRegisteredUsers(v); err != nil {
					return err
				} else {
					value = registeredUsers
				}
			}
			nc.SetRegisteredUsers(value)
			return nil
		},
		nil,
	)

	fieldsMap[Subscriptions] = commons.NewGenericField(
		commons.NotificationCenter,
		Subscriptions,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Endpoint): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(SubscriptionType): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			var result []interface{} = nil

			if nc != nil && nc.Subscriptions != nil {
				result = flattenSubscriptions(nc.Subscriptions)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Subscriptions), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Subscriptions), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			if v, ok := resourceData.GetOk(string(Subscriptions)); ok {
				if v, err := expandSubscriptions(v); err != nil {
					return err
				} else {
					nc.SetSubscriptions(v)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			var value []*notificationcenter.Subscriptions = nil
			if v, ok := resourceData.GetOk(string(Subscriptions)); ok {
				if subscriptions, err := expandSubscriptions(v); err != nil {
					return err
				} else {
					value = subscriptions
				}
			}
			nc.SetSubscriptions(value)
			return nil
		},
		nil,
	)

	fieldsMap[ComputePolicyConfig] = commons.NewGenericField(
		commons.NotificationCenter,
		ComputePolicyConfig,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Events): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Event): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(EventType): {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},
					string(ShouldIncludeAllResources): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(ResourceIds): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					string(DynamicRules): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(FilterConditions): {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Identifier): {
												Type:     schema.TypeString,
												Optional: true,
											},
											string(Operator): {
												Type:     schema.TypeString,
												Optional: true,
											},
											string(Expression): {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			var result []interface{} = nil

			if nc != nil && nc.ComputePolicyConfig != nil {
				result = flattenComputePolicyConfig(nc.ComputePolicyConfig)
			}

			if len(result) > 0 {
				if err := resourceData.Set(string(ComputePolicyConfig), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ComputePolicyConfig), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			if v, ok := resourceData.GetOk(string(ComputePolicyConfig)); ok {
				if cpc, err := expandComputePolicyConfig(v); err != nil {
					return err
				} else {
					nc.SetComputePolicyConfig(cpc)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ncWrapper := resourceObject.(*commons.NotificationCenterWrapper)
			nc := ncWrapper.GetNotificationCenter()
			var value *notificationcenter.ComputePolicyConfig = nil

			if v, ok := resourceData.GetOk(string(ComputePolicyConfig)); ok {
				if computePolicyConfig, err := expandComputePolicyConfig(v); err != nil {
					return err
				} else {
					value = computePolicyConfig
				}
			}
			nc.SetComputePolicyConfig(value)
			return nil
		},
		nil,
	)
}

func expandRegisteredUsers(data interface{}) ([]*notificationcenter.RegisteredUsers, error) {
	list := data.([]interface{})
	users := make([]*notificationcenter.RegisteredUsers, 0, len(list))
	for _, user := range list {
		attr, ok := user.(map[string]interface{})
		if !ok {
			continue
		}

		ru := &notificationcenter.RegisteredUsers{}

		if v, ok := attr[string(UserEmail)].(string); ok && v != "" {
			ru.SetUserEmail(spotinst.String(v))
		}

		if v, ok := attr[string(SubscriptionTypes)]; ok {
			if st, err := expandNotificationList(v); err != nil {
				return nil, err
			} else {
				ru.SetSubscriptionTypes(st)
			}
		}
		users = append(users, ru)
	}
	return users, nil
}

func expandNotificationList(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if ncList, ok := v.(string); ok && ncList != "" {
			result = append(result, ncList)
		}
	}
	return result, nil
}

func expandSubscriptions(data interface{}) ([]*notificationcenter.Subscriptions, error) {
	list := data.([]interface{})
	subscriptions := make([]*notificationcenter.Subscriptions, 0, len(list))
	for _, user := range list {
		attr, ok := user.(map[string]interface{})
		if !ok {
			continue
		}

		sub := &notificationcenter.Subscriptions{}

		if v, ok := attr[string(Endpoint)].(string); ok && v != "" {
			sub.SetEndpoint(spotinst.String(v))
		}
		if v, ok := attr[string(SubscriptionType)].(string); ok && v != "" {
			sub.SetType(spotinst.String(v))
		}

		subscriptions = append(subscriptions, sub)
	}
	return subscriptions, nil
}

func expandComputePolicyConfig(data interface{}) (*notificationcenter.ComputePolicyConfig, error) {
	computePolicyConfig := &notificationcenter.ComputePolicyConfig{}
	if list := data.([]interface{}); len(list) > 0 {
		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(Events)]; ok {
				events, err := expandEvents(v)
				if err != nil {
					return nil, err
				}
				if events != nil {
					computePolicyConfig.SetEvents(events)
				} else {
					computePolicyConfig.SetEvents(nil)
				}
			}
			if v, ok := m[string(ResourceIds)]; ok {
				resourceIds, err := expandNotificationList(v)
				if err != nil {
					return nil, err
				}
				if resourceIds != nil {
					computePolicyConfig.SetResourceIds(resourceIds)
				} else {
					computePolicyConfig.SetResourceIds(nil)
				}
			}
			if v, ok := m[string(DynamicRules)]; ok {
				dynamicRules, err := expandDynamicRules(v)
				if err != nil {
					return nil, err
				}
				if dynamicRules != nil {
					computePolicyConfig.SetDynamicRules(dynamicRules)
				} else {
					computePolicyConfig.SetDynamicRules(nil)
				}
			}
			var shouldInclude = spotinst.Bool(false)
			if v, ok := m[string(ShouldIncludeAllResources)].(bool); ok {
				shouldInclude = spotinst.Bool(v)
			}
			computePolicyConfig.SetShouldIncludeAllResources(shouldInclude)
		}
		return computePolicyConfig, nil
	}
	return nil, nil
}

func expandEvents(data interface{}) ([]*notificationcenter.Events, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		events := make([]*notificationcenter.Events, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			event := &notificationcenter.Events{}

			if v, ok := m[string(Event)].(string); ok && v != "" {
				event.SetEvent(spotinst.String(v))
			}
			if v, ok := m[string(EventType)].(string); ok && v != "" {
				event.SetType(spotinst.String(v))
			}

			events = append(events, event)
		}
		return events, nil
	}
	return nil, nil
}

func expandDynamicRules(data interface{}) ([]*notificationcenter.DynamicRules, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		dynamicRules := make([]*notificationcenter.DynamicRules, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			rule := &notificationcenter.DynamicRules{}

			if v, ok := m[string(FilterConditions)]; ok {
				filterConditions, err := expandFilterConditions(v)
				if err != nil {
					return nil, err
				}
				if filterConditions != nil {
					rule.SetFilterConditions(filterConditions)
				} else {
					rule.SetFilterConditions(nil)
				}
			}
			dynamicRules = append(dynamicRules, rule)
		}
		return dynamicRules, nil
	}
	return nil, nil
}

func expandFilterConditions(data interface{}) ([]*notificationcenter.FilterConditions, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		filterConditions := make([]*notificationcenter.FilterConditions, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			filterCondition := &notificationcenter.FilterConditions{}

			if v, ok := m[string(Identifier)].(string); ok && v != "" {
				filterCondition.SetIdentifier(spotinst.String(v))
			}
			if v, ok := m[string(Operator)].(string); ok && v != "" {
				filterCondition.SetOperator(spotinst.String(v))
			}
			if v, ok := m[string(Expression)].(string); ok && v != "" {
				filterCondition.SetExpression(spotinst.String(v))
			}
			filterConditions = append(filterConditions, filterCondition)
		}
		return filterConditions, nil
	}
	return nil, nil
}

func flattenRegisteredUsers(registeredUsers []*notificationcenter.RegisteredUsers) []interface{} {
	result := make([]interface{}, 0, len(registeredUsers))

	for _, user := range registeredUsers {
		m := make(map[string]interface{})
		m[string(UserEmail)] = spotinst.StringValue(user.UserEmail)
		if len(user.SubscriptionTypes) > 0 {
			m[string(SubscriptionTypes)] = user.SubscriptionTypes
		}
		result = append(result, m)
	}
	return result
}

func flattenSubscriptions(subscriptions []*notificationcenter.Subscriptions) []interface{} {
	result := make([]interface{}, 0, len(subscriptions))
	for _, subscription := range subscriptions {
		m := make(map[string]interface{})
		m[string(Endpoint)] = spotinst.StringValue(subscription.Endpoint)
		m[string(SubscriptionType)] = spotinst.StringValue(subscription.Type)

		result = append(result, m)
	}
	return result
}

func flattenComputePolicyConfig(computePolicy *notificationcenter.ComputePolicyConfig) []interface{} {
	var out []interface{}
	if computePolicy != nil {
		result := make(map[string]interface{})
		if computePolicy.Events != nil {
			result[string(Events)] = flattenEvents(computePolicy.Events)
		}
		if computePolicy.DynamicRules != nil {
			result[string(DynamicRules)] = flattenDynamicRules(computePolicy.DynamicRules)
		}
		result[string(ShouldIncludeAllResources)] = spotinst.BoolValue(computePolicy.ShouldIncludeAllResources)
		if len(computePolicy.ResourceIds) > 0 {
			result[string(ResourceIds)] = computePolicy.ResourceIds
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

func flattenEvents(events []*notificationcenter.Events) []interface{} {
	result := make([]interface{}, 0, len(events))
	for _, event := range events {
		m := make(map[string]interface{})
		m[string(Event)] = spotinst.StringValue(event.Event)
		m[string(EventType)] = spotinst.StringValue(event.Type)

		result = append(result, m)
	}
	return result
}

func flattenDynamicRules(dynamicRules []*notificationcenter.DynamicRules) []interface{} {
	result := make([]interface{}, 0, len(dynamicRules))
	for _, dynamicRule := range dynamicRules {
		m := make(map[string]interface{})
		if dynamicRule.FilterConditions != nil {
			m[string(FilterConditions)] = flattenFilterConditions(dynamicRule.FilterConditions)
		}
		result = append(result, m)
	}
	return result
}

func flattenFilterConditions(filterConditions []*notificationcenter.FilterConditions) []interface{} {
	result := make([]interface{}, 0, len(filterConditions))
	for _, filterCondition := range filterConditions {
		m := make(map[string]interface{})
		m[string(Identifier)] = spotinst.StringValue(filterCondition.Identifier)
		m[string(Operator)] = spotinst.StringValue(filterCondition.Operator)
		m[string(Expression)] = spotinst.StringValue(filterCondition.Expression)

		result = append(result, m)
	}
	return result
}
