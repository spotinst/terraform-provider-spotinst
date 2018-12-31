---
layout: "spotinst"
page_title: "Spotinst: subscription"
sidebar_current: "docs-do-resource-subscription"
description: |-
  Provides a Spotinst subscription resource.
---

# spotinst\_subscription

Provides a Spotinst subscription resource.

## Example Usage

```hcl
# Create a Subscription
resource "spotinst_subscription" "default-subscription" {

  resource_id = "${spotinst_elastigroup_aws.my-eg.id}"
  event_type  = "AWS_EC2_INSTANCE_LAUNCH"
  protocol    = "http"
  endpoint    = "http://endpoint.com"
  
  format = {
    event         = "%event%"
    instance_id   = "%instance-id%"
    resource_id   = "%resource-id%"
    resource_name = "%resource-name%"
    tags          = "foo,baz,baz"
  } 
}
```

## Argument Reference

The following arguments are supported:

* `resource_id` - (Required) Spotinst Resource ID (Elastigroup ID).
* `event_type` - (Required) The event to send the notification when triggered. Valid values: `"AWS_EC2_INSTANCE_TERMINATE"`, `"AWS_EC2_INSTANCE_TERMINATED"`, `"AWS_EC2_INSTANCE_LAUNCH"`, `"AWS_EC2_INSTANCE_UNHEALTHY_IN_ELB"`, `"GROUP_ROLL_FAILED"`, `"GROUP_ROLL_FINISHED"`, `"CANT_SCALE_UP_GROUP_MAX_CAPACITY"`, `"GROUP_UPDATED"`, `"AWS_EC2_CANT_SPIN_OD"`, `"AWS_EMR_PROVISION_TIMEOUT"`, `"AWS_EC2_INSTANCE_READY_SIGNAL_TIMEOUT"`. 
* `protocol` - (Required) The protocol to send the notification. Valid values: `"http"`, `"https"`, `"email"`, `"email-json"`, `"aws-sns"`, `"web"`.
* `endpoint` - (Required) The endpoint the notification will be sent to: url in case of `"http"`/`"https"`, email address in case of `"email"`/`"email-json"`, sns-topic-arn in case of `"aws-sns"`.
* `format` - (Optional) The format of the notification content (JSON Format - Key+Value). Valid values: `"%instance-id%"`, `"%event%"`, `"%resource-id%"`, `"%resource-name%"`.
  
## Attributes Reference

The following attributes are exported:

* `id` - The subscription ID.
