---
layout: "spotinst"
page_title: "Spotinst: ocean_right_sizing_cluster_config"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean Right Sizing Cluster Configuration resource.
---

# spotinst\_ocean\_right\_sizing\_cluster\_config

Manages a Spotinst Ocean right sizing cluster configuration resource. This resource allows you to configure cluster-level right sizing settings.

**NOTE:**
* This resource supports `POST`, `PUT` and `READ` operations only. `Delete` operation is not supported through Terraform.
* `POST` and `PUT` operations can take a few minutes to complete.


## Example Usage

```hcl
resource "spotinst_ocean_right_sizing_cluster_config" "example" {
  ocean_id           = "o-abcd1234"
  cluster_identifier = "dev-cluster"

  config {
    adjust_limit_on_downsize           = true
    downside_only                      = false
    recommendations_cpu_percentile     = 85
    recommendations_memory_percentile  = 85
  }
}
```

## Argument Reference

The following arguments are supported:

* `ocean_id` - (Required) Identifier of the Ocean cluster. Example: `o-abcd1234`
* `cluster_identifier` - (Required) Identifier of the cluster. Example: `dev-cluster`
* `config` - (Required) The Ocean right sizing cluster configuration. The `config` block supports:
    * `adjust_limit_on_downsize` - (Optional, Default: `false`) When set to `true`, the limit will be adjusted when downscale recommendations are applied.
    * `downside_only` - (Optional, Default: `false`) When set to `true`, only downscale recommendations will be applied.
    * `recommendations_cpu_percentile` - (Optional) Change the CPU percentile that the right-sizing recommendations calculation will take into account. Valid values: `85`, `90`, `95`, `99`.
    * `recommendations_memory_percentile` - (Optional) Change the memory percentile that the right-sizing recommendations calculation will take into account. Valid values: `85`, `90`, `95`, `100`.
