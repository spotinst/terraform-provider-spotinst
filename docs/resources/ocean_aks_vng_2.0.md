---
layout: "spotinst"
page_title: "Spotinst: ocean_aks_np"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean Virtual Node Group resource using AKS.
---

# spotinst\_ocean\_aks\_np\_virtual_node_group

Manages a Spotinst Ocean AKS Virtual Node Groups resource.

## Example Usage

```hcl
resource "spotinst_ocean_aks_np_virtual_node_group" "example" {
  
  name  = "oceanVNGName"

  ocean_id = "o-12345"

  // --- autoscale ----------------------------------------------------------------
  headrooms {
    cpu_per_unit    = 6
    memory_per_unit = 10
    gpu_per_unit    = 4
    num_of_units    = 10
  }
  // ----------------------------------------------------------------------------
  
  availability_zones = [1]
  labels ={
    key1   = "label1"
    key2 = "label2"
  }
  
  // --- nodeCountLimits ----------------------------------------------------
  
  min_count = 11
  max_count = 100
  
  // -------------------------------------------------------------------------

  // --- nodePoolProperties --------------------------------------------------
  
  max_pods_per_node     = 110
  enable_node_public_ip = false
  os_disk_size_gb       = 128
  os_disk_type         = "Managed"
  os_type             = "Linux"

  // --------------------------------------------------------------------------

  // --- strategy -------------------------------------------------------------
  
  spot_percentage      = 100
  fallback_to_ondemand = true

  // ---------------------------------------------------------------------------

  taints {
    key    = "key"
    value  = "value"
    effect = "NoSchedule"
  }

  tags ={
    key1   = "value1"
    key2   = "value2"
  }
  // --- vmSizes ---------------------------------------------------------------
  
  filters {
    min_vcpu = 2
    max_vcpu = 16
    min_memory_gib = 10
    max_memory_gib = 18
    architectures = ["X86_64"]
    series = ["D v3"]
  }
  
  // ----------------------------------------------------------------------------
  
  // ----------------------------------------------------------------------------
}
```

```
output "vng_id" {
  value = spotinst_ocean_aks_np_virtual_node_group.example.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Enter a name for the virtual node group.
* `ocean_id` - (Required) The Ocean cluster identifier. Required for Launch Spec creation.
* `headrooms` - (Optional) Specify the custom headroom per VNG. Provide a list of headroom objects.
  * `cpu_per_unit` - (Optional) Configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
  * `memory_per_unit` - (Optional) Amont of GPU to allocate for headroom unit.
  * `gpu_per_unit` - (Optional) Configure the amount of memory (MiB) to allocate the headroom.
  * `num_of_units` - (Optional) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.
* `availability_zones` - (Optional) An Array holding Availability Zones, this configures the availability zones the Ocean may launch instances in per VNG.
* `labels` - (Optional) An array of labels to add to the virtual node group.Only custom user labels are allowed, and not Kubernetes built-in labels or Spot internal labels.
    * `key` - (Required) Set label key.
    * The following are not allowed:
    * "kubernetes.azure.com/agentpool"
"kubernetes.io/arch"
"kubernetes.io/os"
"node.kubernetes.io/instance-type"
"topology.kubernetes.io/region"
"topology.kubernetes.io/zone"
"kubernetes.azure.com/cluster"
"kubernetes.azure.com/mode"
"kubernetes.azure.com/role"
"kubernetes.azure.com/scalesetpriority"
"kubernetes.io/hostname"
"kubernetes.azure.com/storageprofile"
"kubernetes.azure.com/storagetier"
"kubernetes.azure.com/instance-sku"
"kubernetes.azure.com/node-image-version"
"kubernetes.azure.com/subnet"
"kubernetes.azure.com/vnet"
"kubernetes.azure.com/ppg"
"kubernetes.azure.com/encrypted-set"
"kubernetes.azure.com/accelerator"
"kubernetes.azure.com/fips_enabled"
"kubernetes.azure.com/os-sku"
    * `value` - (Required) Set label value.
* `max_count` - (Optional) Maximum node count limit.
* `min_count` - (Optional) Minimum node count limit.
* `enable_node_public_ip` - (Optional) Enable node public IP.
* `max_pods_per_node` - (Optional) The maximum number of pods per node in the node pools.
* `os_disk_size_gb` - (Optional) The size of the OS disk in GB.
* `os_disk_type` - (Optional) The type of the OS disk.
* `os_type` - (Optional) The OS type of the OS disk.
* `fallback_to_ondemand` - (Optional, Default: `true`) If no spot instance markets are available, enable Ocean to launch on-demand instances instead.
* `spot_percentage` - (Optional,Default: `100`) Percentage of spot VMs to maintain.
* `tag` - (Optional) A maximum of 10 unique key-value pairs for VM tags in the virtual node group.
    * `key` - (Optional) Tag key for VMs in the cluster.
    * `value` - (Optional) Tag value for VMs in the cluster.
* `taints` - (Optional) Add taints to a virtual node group.
    * `key` - (Optional) Set taint key. The following are not allowed: "kubernetes.azure.com/scalesetpriority".
    * `value` - (Optional) Set taint value.
    * `effect` - (Optional, Enum: `"NoSchedule", "PreferNoSchedule", "NoExecute", "PreferNoExecute"`) Set taint effect.
* `filters` - (Optional) Filters for the VM sizes that can be launched from the virtual node group.
    * `architectures` - (Optional, Enum `"x86_64", "intel64", "amd64", "arm64"`) The filtered vm sizes will support at least one of the architectures from this list. x86_64 includes both intel64 and amd64.
    * `max_memory_gib` - (Optional) Maximum amount of Memory (GiB).
    * `max_vcpu` - (Optional) Maximum number of vcpus available.
    * `min_memory_gib` - (Optional) Minimum amount of Memory (GiB).
    * `min_vcpu` - (Optional) Minimum number of vcpus available.
    * `series` - (Optional) Vm sizes belonging to a series from the list will be available for scaling.
