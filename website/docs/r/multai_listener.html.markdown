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
    certificate_ids             = ["ce-12345"]
    min_version                 = "TLS10"
    max_version                 = "TLS12"
    cipher_suites               = [""]
    prefer_server_cipher_suites = true
    session_tickets_disabled    = false
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
    * `min_version` - (Required) MinVersion contains the minimum SSL/TLS version that is acceptable (1.0 is the minimum).
    * `max_version` - (Required) MaxVersion contains the maximum SSL/TLS version that is acceptable.
    * `certificate_ids` - (Optional) Contains one or more certificate chains to present to the other side of the connection.
    * `cipher_suites` - (Optional) List of supported cipher suites. If cipherSuites is nil, TLS uses a list of suites supported by the implementation.
    * `prefer_server_cipher_suites` - (Optional) Controls whether the server selects the client’s most preferred ciphersuite, or the server’s most preferred ciphersuite.
    * `session_tickets_disabled` - (Optional) May be set to true to disable session ticket (resumption) support.

* `tags` - (Optional) A list of key:value paired tags.
    * `key` - (Required) The tag's key.
    * `value` - (Required) The tag's value.