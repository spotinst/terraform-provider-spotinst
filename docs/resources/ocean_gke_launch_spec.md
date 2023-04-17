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
  instance_types = ["n1-standard-1, n1-standard-2"]
  tags = ["tag1", "tag2"]
  
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
  
  autoscale_headrooms_automatic {
    auto_headroom_percentage = 5
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
  
  scheduling_task {
    is_enabled = true
    cron_expression = "0 1 * * *"
    task_type = "manualHeadroomUpdate"
    task_headroom {
        num_of_units    = 5
        cpu_per_unit     = 1000
        gpu_per_unit    = 0
        memory_per_unit = 2048
    }
  }
  
  network_interface {
    network = "test-vng-network"
    project_id = "test-vng-network-project"
    access_configs {
      access_configs_name = "external-nat-vng"
      type     = "ONE_TO_ONE_NAT"
    }
    alias_ip_ranges {
      ip_cidr_range         = "/25"
      subnetwork_range_name = "gke-test-native-vpc-pods-5cb557f7-vng"
    }
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
* `metadata` - (Required only if `node_pool_name` is not set) Cluster's metadata.
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
* `tags` - (Optional) Every node launched from this configuration will be tagged with those tags. Note: during creation some tags are automatically imported to the state file, it is required to manually add it to the template configuration
* `autoscale_headrooms_automatic` - (Optional) Set automatic headroom per launch spec.
  * `auto_headroom_percentage` - (Optional) Number between 0-200 to control the headroom % of the specific Virtual Node Group. Effective when cluster.autoScaler.headroom.automatic.`is_enabled` = true is set on the Ocean cluster.
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
* `scheduling_task` - (Optional) Used to define scheduled tasks such as a manual headroom update.
  * `is_enabled` - (Required) Describes whether the task is enabled. When True, the task runs. When False, it does not run.
  * `cron_expression` - (Required) A valid cron expression. For example : " * * * * * ". The cron job runs in UTC time and is in Unix cron format.
  * `task_type` - (Required) The activity that you are scheduling. Valid values: "manualHeadroomUpdate".
  * `task_headroom` - (Optional) The config of this scheduled task. Depends on the value of taskType.
    * `num_of_units` - (Required) The number of units to retain as headroom, where each unit has the defined headroom CPU, memory and GPU.
    * `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
    * `gpu_per_unit` - (Optional) Optionally configure the number of GPUS to allocate for each headroom unit.
    * `memory_per_unit` - (Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.
* `network_interface` - (Optional) Settings for network interfaces.
  * `network` - (Required) The name of the network.
  * `project_id` - (Optional) Use a network resource from a different project. Set the project identifier to use its network resource. This parameter is relevant only if the network resource is in a different project.
  * `access_configs` - (Optional) The network protocol of the VNG.
    * `access_configs_name` - (Optional) The name of the access configuration.
    * `type` - (Optional) The type of the access configuration.
  * `alias_ip_ranges` - (Optional) use the imported node pool’s associated aliasIpRange to assign secondary IP addresses to the nodes. Cannot be changed after VNG creation.
    * `ip_cidr_range` - (Required) specify the IP address range in CIDR notation that can be used for the alias IP addresses associated with the imported node pool.
    * `subnetwork_range_name` - (Required) specify the IP address range for the subnet secondary IP range.

<a id="update-policy"></a>
## Update Policy

* `update_policy` - (Optional)
  * `should_roll` - (Required) Enables the roll.
  * `roll_config` - (Required) Holds the roll configuration.
    * `batch_size_percentage` - (Required) Sets the percentage of the instances to deploy in each batch.

```hcl
update_policy {
  should_roll = false

  roll_config {
    batch_size_percentage = 33
  }
}
```

## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `id` - The Spotinst LaunchSpec ID.
