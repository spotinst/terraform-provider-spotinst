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
resource "spotinst_subscription" "foo" {
	resource_id = "sig-foo"
	event_type = "aws_ec2_instance_launch"
	protocol = "http"
	endpoint = "http://endpoint.com"
	format = {
		instance_id = "%instance-id%"
		tags = "foo,baz,baz"
	}
}
```

## Argument Reference

The following arguments are supported:

* `resource_id` - (Required) The resource to subscribe to
* `event_type` - (Required) The events to subscribe to
* `protocol` - (Required) The protocol to use to connect with the instance. Valid values: http, https
* `endpoint` - (Required) The destination for the request
* `format` - (Optional) The structure of the payload.

## Attributes Reference

The following attributes are exported:

* `id` - The subscription ID.
