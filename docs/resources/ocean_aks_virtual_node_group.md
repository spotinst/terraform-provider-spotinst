---
layout: "spotinst"
page_title: "Spotinst: ocean_aks"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean Virtual Node Group resource using AKS.
---

# spotinst\_ocean\_aks\_virtual\_node\_group

Manages a Spotinst Ocean AKS Virtual Node Group resource.

## Example Usage

```hcl
resource "spotinst_ocean_aks_virtual_node_group" "example" {
   name = "vng_name"
   ocean_id = "o-12345"
 
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
     }
 
     tag {
       key = "label_key"
       value = "label_value"
     }
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
    * `autoscale_headroom` - (Optional)
        * `cpu_per_unit` - (Optional) Configure the number of CPUs to allocate for the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
        * `gpu_per_unit` - (Optional) How many GPU cores should be allocated for headroom unit.
        * `memory_per_unit` - (Optional) Configure the amount of memory (MiB) to allocate the headroom.
        * `num_of_units` - (Required) The number of headroom units to maintain, where each unit has the defined CPU, memory and GPU.
* `launch_specification` - (Optional).
    * `os_disk` - (Optional) Specify OS disk specification other than default.
        * `size_gb` - (Required) The size of the OS disk in GB, Required if dataDisks is specified.
        * `type` - (Optional) The type of the OS disk. Valid values: `"Standard_LRS"`, `"Premium_LRS"`, `"StandardSSD_LRS"`.
    * `tag` - (Optional) Additional key-value pairs to be used to tag the VMs in the virtual node group.
        * `key` - (Optional) Tag Key for Vms in the cluster.
        * `value` - (Optional) Tag Value for VMs in the cluster.
