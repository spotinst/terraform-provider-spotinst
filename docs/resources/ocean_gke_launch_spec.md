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
  name = "specialty.nodes.spotk8s.com"
  source_image = "image"
  restrict_scale_down = true
  root_volume_size = 10
  root_volume_type = "pd-standard"
  instance_types: ["n1-standard-1, n1-standard-2"]
  
  shielded_instance_config {
    enable_secure_boot = false
    enable_integrity_monitoring = true
  }

  storage {
    local_ssd_count = 5
  }

  resource_limits {
    max_instance_count = 3
    min_instance_count = 0
  }
  
  service_account = "default"

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
* `node_pool_name` - (Optional) The node pool you wish to use in your Launch Spec.
* `name` - (Optional) The launch specification name.
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
* `root_volume_type` - (Optional) Root volume disk type. Valid values: `"pd-standard"`, `"pd-ssd"`.
* `root_volume_size` - (Optional) Root volume size (in GB).
* `instance_types` - (Optional) List of supported machine types for the Launch Spec.
* `autoscale_headrooms` - (Optional) Set custom headroom per launch spec. provide list of headrooms object.
    * `num_of_units` - (Required) The number of units to retain as headroom, where each unit has the defined headroom CPU, memory and GPU.
    * `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
    * `gpu_per_unit` - (Optional) Optionally configure the number of GPUS to allocate for each headroom unit.
    * `memory_per_unit` - (Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.
* `strategy` - (Optional) The Ocean Launch Spec Strategy object.
    * `preemptible_percentage` - (Optional) Defines the desired preemptible percentage for this launch specification.
* `shielded_instance_config` - (Optional) The Ocean shielded instance configuration object.
  * `enable_integrity_monitoring` - (Optional) Boolean. Enable the integrity monitoring parameter on the GCP instances.
  * `enable_secure_boot` - (Optional) Boolean. Enable the secure boot parameter on the GCP instances.
* `storage` - (Optional) The Ocean virtual node group storage object.
  * `local_ssd_count` - (Optional) Defines the number of local SSDs to be attached per node for this VNG.
* `resource_limits` - (Optional) The Ocean virtual node group resource limits object.
  * `max_instance_count` - (Optional) Option to set a maximum number of instances per virtual node group. Can be null. If set, the value must be greater than or equal to 0.
  * `min_instance_count` - (Optional) Option to set a minimum number of instances per virtual node group. Can be null. If set, the value must be greater than or equal to 0.
* `service_account` - (Optional) The account used by applications running on the VM to call GCP APIs.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `id` - The Spotinst LaunchSpec ID.
