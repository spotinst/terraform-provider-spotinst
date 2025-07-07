---
layout: "spotinst"
page_title: "Spotinst: notification_center"
subcategory: "Notification Center"
description: |-
  Provides a Spotinst Notification Center Service resource.
---

# spotinst\_notification\_center

Manages Spotinst Notification Center Service.

## Example Usage

```hcl
resource "spotinst_notification_center" "notifications" {
  name = "Notification-Center"
  description = "Creation of notification center policy through terraform"
  is_active = true
  privacy_level = "public"

  registered_users {
    user_email = "sample@xyz.com"
    subscription_types = ["email"]
  }

  registered_users {
    user_email = "terraforma@spot.com"
    subscription_types = ["email", "console"]
  }

  subscriptions {
    endpoint = "testing@xyz.com"
    subscription_type = "email"
  }
  subscriptions {
    endpoint = "https://webhook.si"
    subscription_type = "webhook"
  }

  compute_policy_config {
    events {
        event = "Beanstalk Missing Permissions"
        event_type = "ERROR"
    }
    events {
        event = "Maximum capacity reached"
        event_type = "WARN"
    }
    should_include_all_resources = false
    
    //resource_ids = ["sig-123456789", "sig-987654321"]
    
    dynamic_rules {
        filter_conditions{
            expression = "resourceId EQUALS 'sig-123456789'"
            identifier = "resource_id"
            operator = "equals"
        }
    }
 }
}
```

```
output "id" {
  value = spotinst_notification_center.notifications.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) The name of the notification policy.
* `description` - (Optional) A brief description of the notification policy. Can be null or not empty.
* `is_active` - (Optional) Notification Center Policy tobe active or not.
* `password` - (Optional) Password.
* `privacy_level` - (Required) Valid values: `"private"` `"public"`. The privacy level of the notification policy.
* `registered_users` - (Optional) Registsred Users to the notification policy.
    * `subscription_types` - (Optional) Valid values: `"email"` `"console"`. The subscription types for the registered user.
    * `user_email` - (Optional) User's email address. The user must be part of the organization.
* `subscriptions` - (Optional) Subscriptions to the notification policy.
    * `endpoint` - (Optional) The endpoint of the subscription.
    * `subscription_type` - (Optional) Valid values: `"email"` `"webhook"` `"slack"` `"sns"`. The type of subscription.
* `compute_policy_config` - (Required) Use only one of these parameters: a non-empty resourceIds, a non-empty dynamicRules, or shouldIncludeAllResources set to true.
    * `events` - (Required) A list of events to subscribe to.
      * `event` - (Optional) The event name.
      * `event_type` - (Optional) Valid values: `"ERROR"` `"WARN"` `"INFO"`. The type of event.
    * `should_include_all_resources` - (Optional) If true, all resources will be included in the policy.
    * `resource_ids` - (Optional) Manually specified resource IDs to include in the policy. Must be resources related to the account. Use this parameter only if `should_include_all_resources` is false.
    * `dynamic_rules` - (Optional) A list of dynamic rules to apply to the policy.
      * `filter_conditions` - (Optional) A list of filter conditions to apply to the dynamic rule.
        * `expression` - (Optional) The expression to filter resources.
        * `identifier` - (Optional) The identifier of the resource to filter. Valid values: `"resource_name"` `"resource_id"` `"region"` `"image"` `"tag"` `"load_balancer"`, `"availability_zones"`, `"security_groups"`.
        * `operator` - (Optional) The operator to use for filtering. Valid values: `"equals"` `"not_equals"` `"contains"` `"not_contains"` `"start_with"` `"end_with"`.