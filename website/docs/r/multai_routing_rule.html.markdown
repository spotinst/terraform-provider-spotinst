---
layout: "spotinst"
page_title: "Spotinst: multai routing_rule"
sidebar_current: "docs-do-resource-multai_routing_rule"
description: |-
 Provides a Spotinst Multai Balancer.
---

# spotinst\_multai\_routing\_rule

Provides a Spotinst Multai Routing Rule.

## Example Usage

```hcl
resource "multai_routing_rule" "my_routing_rule" {
  balancer_id = "b-12345"
  listener_id = "l-98765"
  route       = "Path(\x60/bar\x60)"
  strategy    = "LEASTCONN"

  tags = [{
    key   = "env"
    value = "prod"
  }]
}
```

## Argument Reference

The following arguments are supported:

* `balancer_id` - (Required) The ID of the balancer.
* `listener_id` - (Required) The ID of the listener.
* `route` - (Required) Route defines a simple language for matching HTTP requests and route the traffic accordingly. Route provides series of matchers that follow the syntax: Path matcher: — Path("/foo/bar") // trie-based PathRegexp(“/foo/.*”) // regexp-based Method matcher: — Method(“GET”) // trie-based MethodRegexp(“POST|PUT”) // regexp based Header matcher: — Header(“Content-Type”, “application/json”) // trie-based HeaderRegexp(“Content-Type”, “application/.*”) // regexp based Matchers can be combined using && operator: — Method(“POST”) && Path("/v1")
* `strategy` - (Optional) Balancing strategy. Valid values: `ROUNDROBIN`, `RANDOM`, `LEASTCONN`, `IPHASH`.

* `tags` - (Optional) A list of key:value paired tags.
* `key` - (Required) The tag's key.
* `value` - (Required) The tag's value.