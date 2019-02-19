---
layout: "spotinst"
page_title: "Spotinst: multai deployment"
sidebar_current: "docs-do-resource-multai_deployment"
description: |-
 Provides a Spotinst Multai Deployment.
---

# spotinst\_multai\_deployment

Provides a Spotinst Multai Deployment.

## Example Usage

```hcl
resource "multai_deployment" "my_deployment" {
  name = "foo"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The deployment name.