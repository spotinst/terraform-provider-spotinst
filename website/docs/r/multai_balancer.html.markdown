---
layout: "spotinst"
page_title: "Spotinst: multai balancer"
subcategory: "Multai"
description: |-
 Provides a Spotinst Multai Balancer.
---

# spotinst\_multai\_balancer

Provides a Spotinst Multai Balancer.

## Example Usage

```hcl
resource "spotinst_multai_balancer" "my_balancer" {
  name   = "foo"
  scheme = "internal"

  connection_timeouts {
    idle     = 10
    draining = 10
  }

  tags {
    key   = "env"
    value = "prod"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The balancer name. May contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
* `scheme` - (Optional)
* `dns_cname_aliases` - (Optional)

* `connection_timeouts` - (Optional)
* `idle` - (Optional) The idle timeout value, in seconds. (range: 1 - 3600).
* `draining` - (Optional) The time for the load balancer to keep connections alive before reporting the target as de-registered, in seconds (range: 1 - 3600).

* `tags` - (Optional) A list of key:value paired tags.
* `key` - (Required) The tag's key.
* `value` - (Required) The tag's value.
