---
layout: "spotinst"
page_title: "Spotinst: ocean_gke_launch_spec"
sidebar_current: "docs-do-resource-ocean_gke_launch_spec"
description: |-
  Provides a Spotinst Ocean Launch Spec resource using GKE.
---

# spotinst\_ocean\_gke\_launch\_spec

Provides a custom Spotinst Ocean GKE Launch Spec resource.

## Example Usage

```hcl
resource "spotinst_ocean_gke_launch_spec" "example" {
  ocean_id     = "o-123456"
  source_image = "image"
  
  metadata = [{
      key   = "gci-update-strategy"
      value = "update_disabled"
  }]
  
  labels = [{
   key   = "labelKey"
   value = "labelVal"
  }]
  
  taints = [{
   key    = "taintKey"
   value  = "taintVal"
   effect = "taintEffect"
  }]
}
```

## Argument Reference

The following arguments are supported:

* `ocean_id`       - (Required) The Ocean cluster ID required for launchSpec create. 
* `source_image`   - (Required) Image URL.
* `metadata`       - (Required) Cluster's metadata.
* `taints`         - (Optional) Cluster's taints.
* `labels`         - (Optional) Cluster's labels.