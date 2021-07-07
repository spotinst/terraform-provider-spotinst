---
layout: "spotinst"
page_title: "Spotinst: ocean_gke_launch_spec"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean Launch Spec resource using GKE.
---

# spotinst\_ocean\_gke\_launch\_spec

Manages a custom Spotinst Ocean GKE Launch Spec resource.

-> This resource can be imported from GKE node pool or not. If you want to import the node pool and create the VNG from it, please provide `node_pool_name`.

## Example Usage

```hcl
resource "spotinst_ocean_gke_launch_spec" "example" {
  ocean_id     = "o-123456"
  node_pool_name  = "default-pool"
  source_image = "image"
  restrict_scale_down = true
  root_volume_size = 10
  root_volume_type = "pd-standard"
  instance_types: ["n1-standard-1, n1-standard-2"]

  metadata {
    key   = "gci-update-strategy"
    value = "update_disabled"
  }
  
  labels {
    key   = "labelKey"
    value = "labelVal"
  }
  
  taints {
    key    = "taintKey"
    value  = "taintVal"
    effect = "taintEffect"
  }
  
  autoscale_headrooms {
    num_of_units = 5
    cpu_per_unit = 1000
    gpu_per_unit = 0
    memory_per_unit = 2048
  }

  strategy {
    preemptible_percentage = 30
  }
}
```
```
output "ocean_launchspec_id" {
  value = spotinst_ocean_gke_launch_spec.example.id
}
```

## Argument Reference

The following arguments are supported:

* `ocean_id` - (Required) The Ocean cluster ID.
* `node_pool_name` - (Optional) The node pool you wish to use in your launchSpec.
* `source_image` - (Required) Image URL.
* `metadata` - (Required) Cluster's metadata.
    * `key` - (Required) The metadata key.
    * `value` - (Required) The metadata value.
* `taints` - (Optional) Optionally adds labels to instances launched in an Ocean cluster.
    * `key` - (Required) The taint key.
    * `value` - (Required) The taint value.
    * `effect` - (Required) The effect of the taint. Valid values: `"NoSchedule"`, `"PreferNoSchedule"`, `"NoExecute"`.
* `labels` - (Optional) Optionally adds labels to instances launched in an Ocean cluster.
    * `key` - (Required) The label key.
    * `value` - (Required) The label value.
* `restrict_scale_down` - (Optional) Boolean. When set to `true`, VNG nodes will be treated as if all pods running have the restrict-scale-down label. Therefore, Ocean will not scale nodes down unless empty.
* `root_volume_type` - (Optional) Set the root volume disk type. Valid values: `"pd-standard"`, `"pd-ssd"`.
* `root_volume_size` - (Optional) Set root volume size (in GB).
* `instance_types` - (Optional) A array of supported machine types for the Launch Spec.
* `autoscale_headrooms` - (Optional) Set custom headroom per launch spec. provide list of headrooms object.
    * `num_of_units` - (Required) The number of units to retain as headroom, where each unit has the defined headroom CPU, memory and GPU.
    * `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
    * `gpu_per_unit` - (Optional) Optionally configure the number of GPUS to allocate for each headroom unit.
    * `memory_per_unit` - (Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.
* `strategy` - (Optional) The Ocean Launch Spec Strategy object.
    * `preemptible_percentage` - (Optional) Defines the desired preemptible percentage for this launch specification.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `id` - The Spotinst LaunchSpec ID.