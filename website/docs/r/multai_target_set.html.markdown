---
layout: "spotinst"
page_title: "Spotinst: multai target set"
sidebar_current: "docs-do-resource-multai_target_set"
description: |-
 Provides a Spotinst Multai Target Set.
---

# spotinst\_multai\_target\_set

Provides a Spotinst Multai Target Set.

## Example Usage

```hcl
resource "multai_target_set" "my_target_set" {
  balancer_id   = "b-12345"
  deployment_id = "dp-12345"
  name          = "foo"
  protocol      = "http"
  port          = 1338
  weight        = 2

  health_check = {
    protocol = "http"
    path     = "/"
    port     = 3001
    interval = 20
    timeout  = 5

    healthy_threshold   = 3
    unhealthy_threshold = 3
  }

  tags = [{
   key   = "env"
   value = "prod"
  }]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Target Set. Must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
* `balancer_id` - (Required) The id of the balancer.
* `deployment_id` - (Required) The id of the deployment.
* `protocol` - (Required) The protocol to allow connections to the target.
* `port`
* `weight` - (Required) Defines how traffic is distributed between the Target Set.

* `health_check`
* `protocol` - (Required) The protocol to allow connections to the target for the health check.
* `path` - (Required) The path to perform the health check.
* `port` - (Required) The port on which the load balancer is listening.
* `interval` - (Required) The interval for the health check.
* `timeout` - (Required) The time out for the health check.

* `healthy_threshold` - (Required) Total number of allowed healthy Targets.
* `unhealthy_threshold` - (Required) Total number of allowed unhealthy Targets.

* `tags` - (Optional) A list of key:value paired tags.
* `key` - (Required) The tag's key.
* `value` - (Required) The tag's value.