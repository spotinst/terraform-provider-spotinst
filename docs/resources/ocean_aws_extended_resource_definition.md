---
layout: "spotinst"
page_title: "Spotinst: ocean_aws_extended_resource_definition"
subcategory: "Ocean"
description: |-
  Manages an Ocean extended resource definition resource.
---

# spotinst\_ocean\_aws\_extended\_resource\_definition

Provides a Spotinst Ocean AWS Extended Resource Definition resource.

## Example Usage

```hcl
resource "spotinst_ocean_aws_extended_resource_definition" "example" {
  name  = "terraform_extended_resource_definition"
  resource_mapping = {
    "c3.large"  = "2Ki"
    "c3.xlarge" = "4Ki"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The extended resource name as should be requested by your pods and registered to the nodes. Cannot be updated.
  The name should be a valid Kubernetes extended resource name.
* `resource_mapping` - (Required) A mapping between AWS instanceType or * as default and its value for the given extended resource.

  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `id` - The Spotinst Extended Resource Definition ID.
