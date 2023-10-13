---
layout: "spotinst"
page_title: "Spotinst: ocean_aks_np_virtual_node_group"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean Virtual Node Group resource using AKS.
---

# spotinst\_ocean\_aks\_np\_virtual_node_group

Manages a Spotinst Ocean AKS Virtual Node Groups resource.

## Example Usage

```hcl
resource "spotinst_ocean_aks_np_virtual_node_group" "example" {
  
  name  = "testVng"

  ocean_id = "o-134abcd"

  // --- autoscale ----------------------------------------------------------------
  headrooms {
    cpu_per_unit    = 1024
    memory_per_unit = 512
    gpu_per_unit    = 0
    num_of_units    = 2
  }
  // ----------------------------------------------------------------------------
  
  availability_zones = [
    "1",
    "2",
    "3"
  ]
  
  labels ={
    key   = "env"
    value = "test"
  }
  
  // --- nodeCountLimits ----------------------------------------------------
  
  min_count = 1
  max_count = 100
  
  // -------------------------------------------------------------------------

  // --- nodePoolProperties --------------------------------------------------
  
  max_pods_per_node     = 30
  enable_node_public_ip = true
  os_disk_size_gb       = 30
  os_disk_type          = "Managed"
  os_type               = "Linux"
  os_sku                = "Ubuntu"
  kubernetes_version    = "1.26"
  pod_subnet_ids       = ["/subscriptions/123456-1234-1234-1234-123456789/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/virtualNetworks/ExampleVirtualNetwork/subnets/default"]
  vnet_subnet_ids       = ["/subscriptions/123456-1234-1234-1234-123456789/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/virtualNetworks/ExampleVirtualNetwork/subnets/default"]

  // --------------------------------------------------------------------------

  // --- strategy -------------------------------------------------------------
  
  spot_percentage      = 50
  fallback_to_ondemand = true

  // ---------------------------------------------------------------------------

  taints {
    key    = "taintKey"
    value  = "taintValue"
    effect = "NoSchedule"
  }

  tags ={
    tagKey   = "env"
    tagValue   = "staging"
  }
  // --- vmSizes ---------------------------------------------------------------
  
  filters {
    min_vcpu = 2
    max_vcpu = 16
    min_memory_gib = 8
    max_memory_gib = 16
    architectures = ["X86_64"]
    series = ["D v3", "F", "E v4"]
    exclude_series = ["Bs", "Da v4"]
    accelerated_networking = "Enabled"
    disk_performance = "Premium"
    min_gpu = 1
    max_gpu = 2
    min_nics = 1
    vm_types = ["generalPurpose", "GPU"]
    min_data = 1
  }
  
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
  * `memory_per_unit` - (Optional) Configure the amount of memory (MiB) to allocate the headroom.
  * `gpu_per_unit` - (Optional) Amount of GPU to allocate for headroom unit.
  * `num_of_units` - (Optional) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.
* `availability_zones` - (Optional) An Array holding Availability Zones, this configures the availability zones the Ocean may launch instances in per VNG.
* `labels` - (Optional) An array of labels to add to the virtual node group.Only custom user labels are allowed, and not Kubernetes built-in labels or Spot internal labels.
  * `key` - (Required) Set label key. The following are not allowed: ["kubernetes.azure.com/agentpool", "kubernetes.io/arch", "kubernetes.io/os", "node.kubernetes.io/instance-type", "topology.kubernetes.io/region", "topology.kubernetes.io/zone", "kubernetes.azure.com/cluster", "kubernetes.azure.com/mode", "kubernetes.azure.com/role", "kubernetes.azure.com/scalesetpriority", "kubernetes.io/hostname", "kubernetes.azure.com/storageprofile", "kubernetes.azure.com/storagetier", "kubernetes.azure.com/instance-sku", "kubernetes.azure.com/node-image-version", "kubernetes.azure.com/subnet", "kubernetes.azure.com/vnet", "kubernetes.azure.com/ppg", "kubernetes.azure.com/encrypted-set", "kubernetes.azure.com/accelerator", "kubernetes.azure.com/fips_enabled", "kubernetes.azure.com/os-sku"]
  * `value` - (Required) Set label value.
* `max_count` - (Optional) Maximum node count limit.
* `min_count` - (Optional) Minimum node count limit.
* `enable_node_public_ip` - (Optional) Enable node public IP.
* `max_pods_per_node` - (Optional) The maximum number of pods per node in the node pools.
* `os_disk_size_gb` - (Optional) The size of the OS disk in GB.
* `os_disk_type` - (Optional, Enum:`"Managed" ,"Ephemeral"`) The type of the OS disk.
* `os_type` - (Optional) The OS type of the OS disk. Can't be modified once set.
* `os_sku` - (Optional, Enum: `"Ubuntu", "Windows2019", "Windows2022", "AzureLinux", "CBLMariner"`) The OS SKU of the OS type. Must correlate with the os type.
* `kubernetes_version` - (Optional) The desired Kubernetes version of the launched nodes. In case the value is null, the Kubernetes version of the control plane is used.
* `pod_subnet_ids` - (Optional) The IDs of subnets in an existing VNet into which to assign pods in the cluster (requires azure network-plugin).
* `vnet_subnet_ids` - (Optional) The IDs of subnets in an existing VNet into which to assign nodes in the cluster (requires azure network-plugin).
* `fallback_to_ondemand` - (Optional, Default: `true`) If no spot instance markets are available, enable Ocean to launch on-demand instances instead.
* `spot_percentage` - (Optional, Default: `100`) Percentage of spot VMs to maintain.
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
    * `exclude_series` - (Optional) Vm sizes belonging to a series from the list will not be available for scaling.
    * `accelerated_networking` - (Optional, Enum `"Enabled", "Disabled"`) In case acceleratedNetworking is set to Enabled, accelerated networking applies only to the VM that enables it.
    * `disk_performance` - (Optional, Enum `"Standard", "Premium"`) The filtered vm sizes will support at least one of the classes from this list.
    * `min_gpu` - (Optional) Minimum number of GPUs available.
    * `max_gpu` - (Optional) Maximum number of GPUs available.
    * `min_nics` - (Optional) Minimum number of network interfaces.
    * `min_data` - (Optional) Minimum number of data disks available.
    * `vm_types` - (Optional, Enum `"generalPurpose", "memoryOptimized", "computeOptimized", "highPerformanceCompute", "storageOptimized", "GPU"`) The filtered vm types will belong to one of the vm types from this list.
