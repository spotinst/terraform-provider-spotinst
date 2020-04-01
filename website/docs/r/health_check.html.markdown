---
layout: "spotinst"
page_title: "Spotinst: health_check"
subcategory: "Elastigroup"
description: |-
  Provides a Spotinst Health Check resource.
---

# spotinst\_health\_check

Provides a Spotinst Health Check resource.

## Example Usage

```hcl 
resource "spotinst_health_check" "http_check" {
  name        = "terraform_healt_cheack"
  resource_id = "sig-123"

  check {
    protocol = "http"
    endpoint = "http://endpoint.com"
    port     = 1337
    interval = 10
    timeout  = 10
  }

  threshold {
    healthy   = 1
    unhealthy = 1
  }

  proxy {
    addr = "http://proxy.com"
    port = 80
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) The name of the health check.
* `resource_id` - (Required) The ID of the resource to check.
* `check` - (Required) Describes the check to execute.

    * `protocol` - (Required) The protocol to use to connect with the instance. Valid values: http, https.
    * `endpoint` - (Required) The destination for the request.
    * `port` - (Required) The port to use to connect with the instance.
    * `interval` - (Required) The amount of time (in seconds) between each health check (minimum: 10).
    * `timeout` - (Required) the amount of time (in seconds) to wait when receiving a response from the health check.

* `threshold` - (Required)

  * `healthy` - (Required) The number of consecutive successful health checks that must occur before declaring an instance healthy.
  * `unhealthy` - (Required) The number of consecutive failed health checks that must occur before declaring an instance unhealthy.

* `proxy` - (Required)

  * `addr` - (Required) The public hostname / IP where you installed the Spotinst HCS.
  * `port` - (Required) The port of the Spotinst HCS (default: 80).

## Attributes Reference

The following attributes are exported:

* `id` - The Health Check ID.
