---
layout: "spotinst"
page_title: "Spotinst: multai listener"
sidebar_current: "docs-do-resource-multai_listener"
description: |-
 Provides a Spotinst Multai Listener.
---

# spotinst\_multai\_listener

Provides a Spotinst Multai Listener.

## Example Usage

```hcl
resource "multai_listener" "my_listener" {
  name        = "foo"
  balancer_id = "b-12345"
  protocol    = "http"
  port        = 1337

  tls_config = {
    min_version = "1.0"
    max_version = ""

  }

  tags = [{
    key   = "env"
    value = "prod"
  }]
}
```

## Argument Reference

The following arguments are supported:

* `balancer_id` - (Required) The ID of the balancer.
* `protocol` - (Required) The protocol to allow connections to the load balancer.
* `port` - (Required) The port on which the load balancer is listening.

* `tls_config` - (Optional) Describes the TLSConfig configuration.
* `min_version` - (Required) MinVersion contains the minimum SSL/TLS version that is acceptable (1.0 is the minimum)
* `max_version` - (Required) MaxVersion contains the maximum SSL/TLS version that is acceptable.
*

* `tags` - (Optional) A list of key:value paired tags.
* `key` - (Required) The tag's key.
* `value` - (Required) The tag's value.