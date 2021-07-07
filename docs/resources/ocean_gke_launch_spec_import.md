---
layout: "spotinst"
page_title: "Spotinst: ocean_gke_launch_spec_import"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean Launch Spec Import resource using GKE.
---

# spotinst\_ocean\_gke\_launch\_spec_import (legacy)

Manages a custom Spotinst Ocean GKE Launch Spec Import resource.

## Example Usage

```hcl
resource "spotinst_ocean_gke_launch_spec_import" "example" {
  ocean_id        = "o-123456"
  node_pool_name  = "default-pool"
}
```
```
output "ocean_launchspec_id" {
  value = spotinst_ocean_gke_launch_spec_import.example.id
}
```

## Argument Reference

The following arguments are supported:

* `ocean_id`       - (Required) The Ocean cluster ID required for launchSpec create. 
* `node_pool_name` - (Required) The node pool you wish to use in your launchSpec.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `id` - The Spotinst LaunchSpec ID.