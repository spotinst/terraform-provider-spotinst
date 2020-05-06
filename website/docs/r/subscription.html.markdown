---
layout: "spotinst"
page_title: "Spotinst: subscription"
subcategory: "Subscription"
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
  
  format {
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

* `resource_id` - (Required) Spotinst Resource id (Elastigroup or Ocean ID).
* `event_type` - (Required) The event to send the notification when triggered. Valid values: `"AWS_EC2_INSTANCE_TERMINATE"`, `"AWS_EC2_INSTANCE_TERMINATED"`, `"AWS_EC2_INSTANCE_LAUNCH"`, `"AWS_EC2_INSTANCE_READY_SIGNAL_TIMEOUT"`, `"AWS_EC2_CANT_SPIN_OD"`, `"AWS_EC2_INSTANCE_UNHEALTHY_IN_ELB"`, `"GROUP_ROLL_FAILED"`, `"GROUP_ROLL_FINISHED"`,
                            `"CANT_SCALE_UP_GROUP_MAX_CAPACITY"`,
                            `"GROUP_UPDATED"`,
                            `"AWS_EMR_PROVISION_TIMEOUT"`,
                            `"GROUP_BEANSTALK_INIT_READY"`,
                            `"AZURE_VM_TERMINATED"`,
                            `"AZURE_VM_TERMINATE"`,
                            `"AWS_EC2_MANAGED_INSTANCE_PAUSING"`,
                            `"AWS_EC2_MANAGED_INSTANCE_RESUMING"`,
                            `"AWS_EC2_MANAGED_INSTANCE_RECYCLING"`,`"AWS_EC2_MANAGED_INSTANCE_DELETING"`.
                            Ocean Events:`"CLUSTER_ROLL_FINISHED"`,`"GROUP_ROLL_FAILED"`. 
* `protocol` - (Required) The protocol to send the notification. Valid values: `"email"`, `"email-json"`, `"aws-sns"`, `"web"`. 
                          The following values are deprecated: `"http"` , `"https"`
                          You can use the generic `"web"` protocol instead.
                          `"aws-sns"` is only supported with AWS provider
* `endpoint` - (Required) The endpoint the notification will be sent to. url in case of `"http"`/`"https"`/`"web"`, email address in case of `"email"`/`"email-json"` and sns-topic-arn in case of `"aws-sns"`.
* `format` - (Optional) The format of the notification content (JSON Format - Key+Value). Valid Values : `"instance-id"`, `"event"`, `"resource-id"`, `"resource-name"`, `"subnet-id"`, `"availability-zone"`, `"reason"`, `"private-ip"`, `"launchspec-id"`
                        Example: {"event": `"event"`, `"resourceId"`: `"resource-id"`, `"resourceName"`: `"resource-name"`", `"myCustomKey"`: `"My content is set here"` }
                        Default: {`"event"`: `"<event>"`, `"instanceId"`: `"<instance-id>"`, `"resourceId"`: `"<resource-id>"`, `"resourceName"`: `"<resource-name>"` }.
  
## Attributes Reference

The following attributes are exported:

* `id` - The subscription ID.
