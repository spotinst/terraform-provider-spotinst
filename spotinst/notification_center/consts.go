package notification_center

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Name         commons.FieldName = "name"
	Description  commons.FieldName = "description"
	PrivacyLevel commons.FieldName = "privacy_level"
	IsActive     commons.FieldName = "is_active"
)

const (
	RegisteredUsers   commons.FieldName = "registered_users"
	UserEmail         commons.FieldName = "user_email"
	SubscriptionTypes commons.FieldName = "subscription_types"
)

const (
	Subscriptions    commons.FieldName = "subscriptions"
	Endpoint         commons.FieldName = "endpoint"
	SubscriptionType commons.FieldName = "subscription_type"
)

const (
	ComputePolicyConfig       commons.FieldName = "compute_policy_config"
	ShouldIncludeAllResources commons.FieldName = "should_include_all_resources"
	ResourceIds               commons.FieldName = "resource_ids"
	Events                    commons.FieldName = "events"
	Event                     commons.FieldName = "event"
	EventType                 commons.FieldName = "event_type"
	DynamicRules              commons.FieldName = "dynamic_rules"
	FilterConditions          commons.FieldName = "filter_conditions"
	Identifier                commons.FieldName = "identifier"
	Operator                  commons.FieldName = "operator"
	Expression                commons.FieldName = "expression"
)
