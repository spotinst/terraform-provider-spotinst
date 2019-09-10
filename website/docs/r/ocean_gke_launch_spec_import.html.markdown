---
layout: "spotinst"
page_title: "Spotinst: ocean_gke_launch_spec_import"
sidebar_current: "docs-do-resource-ocean_gke_launch_spec_import"
description: |-
  Provides a Spotinst Ocean Launch Spec Import resource using GKE.
---

# spotinst\_ocean\_gke\_launch\_spec_import

Provides a custom Spotinst Ocean GKE Launch Spec Import resource.

## Example Usage

```hcl
resource "spotinst_ocean_gke_launch_spec_import" "example" {
  ocean_id        = "o-123456"
  node_pool_name  = "default-pool"
}
```

## Argument Reference

The following arguments are supported:

* `ocean_id`       - (Required) The Ocean cluster ID required for launchSpec create. 
* `node_pool_name` - (Required) The node pool you wish to use in your launchSpec.