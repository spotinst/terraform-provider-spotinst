---
layout: "spotinst"
page_title: "Spotinst: multai deployment"
subcategory: "Multai"
description: |-
 Provides a Spotinst Multai Deployment.
---

# spotinst\_multai\_deployment

Provides a Spotinst Multai Deployment.

## Example Usage

```hcl
resource "spotinst_multai_deployment" "my_deployment" {
  name = "foo"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The deployment name.
