---
layout: "spotinst"
page_title: "Spotinst: ocean_aks (Deprecated)"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean Virtual Node Group resource using AKS.
---

# spotinst\_ocean\_aks\_virtual\_node\_group (Deprecated)

Manages a Spotinst Ocean AKS Virtual Node Group resource.

## Example Usage

```hcl
resource "spotinst_ocean_aks_virtual_node_group" "example" {
   name = "vng_name"
   ocean_id = "o-12345"
   
   zones = ["1","2","3"]
 
   label {
     key = "label_key"
     value = "label_value"
   }
 
   taint {
     key = "taint_key"
     value = "taint_value"
     effect = "NoSchedule"
   }
 
   resource_limits {
     max_instance_count = 4
   }
 
   autoscale {
     auto_headroom_percentage = 5
     autoscale_headroom {
       cpu_per_unit = 4
       gpu_per_unit = 8
       memory_per_unit = 100
       num_of_units = 16
     }
   }
 
   launch_specification {
     os_disk {
       size_gb = 100
       type = "Standard_LRS"
       utilize_ephemeral_storage = false
     }
 
     tag {
       key = "label_key"
       value = "label_value"
     }
     
     max_pods = 30
   }
}
```

```
output "ocean_id" {
  value = spotinst_ocean_aks_.example.id
}
```

## Argument Reference

The following arguments are supported:

* `ocean_id` - (Required) The Ocean cluster ID.
* `name` - (Required) Set name for the virtual node group.
* `zones` - (Optional) An Array holding Availability Zones, this configures the availability zones the Ocean may launch instances in per VNG.
* `label` - (Optional) Additional labels for the virtual node group. Only custom user labels are allowed. Kubernetes built-in labels and Spot internal labels are not allowed.
    * `key` - (Required) The label key.
    * `value` - (Optional) The label value.
* `taint` - (Optional) Additional taints for the virtual node group. Only custom user labels are allowed. Kubernetes built-in labels and Spot internal labels are not allowed.
    * `key` - (Optional) The taint key.
    * `value` - (Optional) The taint value.
     * `effect` - (Optional) The effect of the taint. Valid values: `"NoSchedule"`, `"PreferNoSchedule"`, `"NoExecute"`, `"PreferNoExecute"`.
* `resource_limits` - (Optional).
    * `max_instance_count` - (Optional) Option to set a maximum number of instances per virtual node group. If set, value must be greater than or equal to 0.
* `autoscale` - (Optional).
    * `auto_headroom_percentage` - (Optional) Number between 0-200 to control the headroom % of the specific Virtual Node Group. Effective when `cluster.autoScaler.headroom.automatic.is_enabled` = true is set on the Ocean cluster.
    * `autoscale_headroom` - (Optional)
        * `cpu_per_unit` - (Optional) Configure the number of CPUs to allocate for the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
        * `gpu_per_unit` - (Optional) How many GPU cores should be allocated for headroom unit.
        * `memory_per_unit` - (Optional) Configure the amount of memory (MiB) to allocate the headroom.
        * `num_of_units` - (Required) The number of headroom units to maintain, where each unit has the defined CPU, memory and GPU.
* `launch_specification` - (Optional).
    * `os_disk` - (Optional) Specify OS disk specification other than default.
        * `size_gb` - (Required) The size of the OS disk in GB, Required if dataDisks is specified.
        * `type` - (Optional) The type of the OS disk. Valid values: `"Standard_LRS"`, `"Premium_LRS"`, `"StandardSSD_LRS"`.
        * `utilize_ephemeral_storage` - (Optional) Flag to enable/disable the Ephemeral OS Disk utilization.
    * `tag` - (Optional) Additional key-value pairs to be used to tag the VMs in the virtual node group.
        * `key` - (Optional) Tag Key for Vms in the cluster.
        * `value` - (Optional) Tag Value for VMs in the cluster.
  * `max_pods` - (Optional) The maximum number of pods per node in an AKS cluster.
<a id="update-policy"></a>
## Update Policy

* `update_policy` - (Optional)
  * `should_roll` - (Required) If set to true along with the vng update, roll will be triggered.
  * `conditioned_roll` - (Optional, Default: false) Spot will perform a cluster Roll in accordance with a relevant modification of the cluster’s settings. When set to true , only specific changes in the cluster’s configuration will trigger a cluster roll (such as availability_zones, max_pods_per_node, enable_node_public_ip, os_disk_size_gb, os_disk_type, os_sku, kubernetes_version, vnet_subnet_ids, pod_subnet_ids, labels, taints and tags).
  * `roll_config` - (Required) While used, you can control whether the group should perform a deployment after an update to the configuration.
    * `batch_min_healthy_percentage` - (Optional, Default: 50) Indicates the threshold of minimum healthy instances in single batch. If the amount of healthy instances in single batch is under the threshold, the cluster roll will fail. If exists, the parameter value will be in range of 1-100. In case of null as value, the default value in the backend will be 50%. Value of param should represent the number in percentage (%) of the batch.
    * `batch_size_percentage` - (Required) Value as a percent to set the size of a batch in a roll. Valid values are 0-100. In case of null as value, the default value in the backend will be 20%.
    * `comment` - (Optional) Add a comment description for the roll. The comment is limited to 256 chars and optional.
    * `respect_pdb` - (Optional, Default: true) During the roll, if the parameter is set to True we honor PDB during the instance replacement.
    * `respect_restrict_scale_down` - (Optional, Default: false) During the roll, if the parameter is set to true we honor Restrict Scale Down label during the nodes replacement.
    * `node_pool_names` - (Optional) List of node pools to be rolled. Each node pool name is a string. nodePoolNames can be null, and cannot be used together with nodeNames and vngIds. 
    * `node_names` - (Optional) List of node names to be rolled. Each identifier is a string. nodeNames can be null, and cannot be used together with nodePoolNames and vngIds.
    * `vng_ids` - (Optional) List of virtual node group identifiers to be rolled. Each identifier is a string. vngIds can be null, and cannot be used together with nodeNames and nodePoolNames.
```hcl
update_policy {
  should_roll = false
  conditioned_roll = true

  roll_config {
    vng_ids = ["ols-12345"]
    batch_size_percentage = 25
    batch_min_healthy_percentage = 100
    respect_pdb = true
  }
}
```
