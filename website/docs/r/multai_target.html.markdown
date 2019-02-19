---
layout: "spotinst"
page_title: "Spotinst: multai target"
sidebar_current: "docs-do-resource-multai_target"
description: |-
 Provides a Spotinst Multai target.
---

# spotinst\_multai\_target

Provides a Spotinst Multai Target.

## Example Usage

```hcl
resource "multai_target" "my_target" {
  balancer_id   = "b-12345"
  target_set_id = "l-98765"

  name   = "foo"
  port   = 1338
  host   = "host"
  weight = 1

  tags = [{
    key   = "env"
    value = "prod"
  }]
}
```

## Argument Reference

The following arguments are supported:

* `balancer_id` - (Required) The ID of the balancer.
* `target_set_id` - (Required) The ID of the target set.
* `name` - (Required) The name of the Target . Must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
* `port` - (Required) The port the target will register to.
* `host` - (Required) The address (IP or URL) of the targets to register
* `weight` - (Required) Defines how traffic is distributed between targets.

* `tags` - (Optional) A list of key:value paired tags.
* `key` - (Required) The tag's key.
* `value` - (Required) The tag's value.