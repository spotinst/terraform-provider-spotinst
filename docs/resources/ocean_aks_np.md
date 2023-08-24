---
layout: "spotinst"
page_title: "Spotinst: ocean_aks_np"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean resource using AKS.
---

# spotinst\_ocean\_aks\_np

Manages a Spotinst Ocean AKS resource.

## Prerequisites

Installation of the Ocean controller is required by this resource. You can accomplish this by using the [spotinst/ocean-controller](https://registry.terraform.io/modules/spotinst/ocean-controller/spotinst) module as follows:

```hcl
module "ocean-controller" {
  source = "spotinst/ocean-controller/spotinst"

  # Credentials.
  spotinst_token   = "redacted"
  spotinst_account = "redacted"

  # Configuration.
  cluster_identifier = "ocean-westus-dev-aks"
}
```

~> You must configure the same `cluster_identifier` for the Ocean controller and for the `spotinst_ocean_aks_np` resource.

## Basic Ocean Cluster Creation Example Usage

```hcl
resource "spotinst_ocean_aks_np" "example" {
  
  name = "test"
  
  // --- AKS ------------------------------------------------------------------------
  
  aks_region                             = "eastus"
  aks_cluster_name                       = "test-cluster"
  aks_infrastructure_resource_group_name = "MC_TestResourceGroup_test-cluster_eastus"
  aks_resource_group_name                = "TestResourceGroup"
  
  // --------------------------------------------------------------------------------

  controller_cluster_id = "test-123124"
  
  // --- virtualNodeGroupTemplate --------------------------------------
  
  availability_zones = [
    "1",
    "2",
    "3"
  ]
  
  // -------------------------------------------------------------------
}
```  
  
## Detailed Example Usage

```hcl  
resource "spotinst_ocean_aks_np" "example" {
  
  name = "test"
  
  // --- AKS -------------------------------------------------------------------------
  
  aks_region                             = "eastus"
  aks_cluster_name                       = "test-cluster"
  aks_infrastructure_resource_group_name = "MC_TestResourceGroup_test-cluster_eastus"
  aks_resource_group_name                = "TestResourceGroup"
  
  // ---------------------------------------------------------------------------------

  controller_cluster_id = "test-123124"

  // --- Auto Scaler ---------------------------------------------------
  
  autoscaler {
    autoscale_is_enabled = true
    resource_limits {
      max_vcpu       = 750
      max_memory_gib = 1500
    }
    autoscale_down {
      max_scale_down_percentage = 30
    }
    autoscale_headroom {
      automatic {
        percentage = 5
      }
    }
  }
  
  // ----------------------------------------------------------------------

  // --- Health -----------------------------------------------------------
  
  health {
    grace_period = 600
  }
  
  // ----------------------------------------------------------------------

  // --- Scheduling -------------------------------------------------------
  
  scheduling{
    shutdown_hours{
      is_enabled   = true
      time_windows = ["Sat:08:00-Sun:08:00"]
    }
  }
  
  // ----------------------------------------------------------------------

  // --- virtualNodeGroupTemplate -----------------------------------------

  // --- autoscale --------------------------------------------------------
  headrooms {
    cpu_per_unit    = 1024
    memory_per_unit = 512
    gpu_per_unit    = 0
    num_of_units    = 2
  }
  // ----------------------------------------------------------------------
  
  availability_zones = [
    "1",
    "2",
    "3"
  ]
  labels = {
    key   = "env"
    value = "test"
  }
  
  // --- nodeCountLimits --------------------------------------------------
  
  min_count = 1
  max_count = 100
  
  // ----------------------------------------------------------------------

  // --- nodePoolProperties -----------------------------------------------
  
  max_pods_per_node     = 30
  enable_node_public_ip = true
  os_disk_size_gb       = 30
  os_disk_type          = "Managed"
  os_type               = "Windows"
  os_sku                = "Windows2022"

  // ----------------------------------------------------------------------

  // --- strategy ---------------------------------------------------------
  
  spot_percentage      = 50
  fallback_to_ondemand = true

  // ----------------------------------------------------------------------

  taints {
    key    = "taintKey"
    value  = "taintValue"
    effect = "NoSchedule"
  }

  tags = {
    tagKey   = "env"
    tagValue = "staging"
  }
  // --- vmSizes ----------------------------------------------------------
  
  filters {
    min_vcpu       = 2
    max_vcpu       = 16
    min_memory_gib = 10
    max_memory_gib = 18
    architectures  = ["X86_64"]
    series         = ["D v3", "Dds_v4", "Dsv2"]
    exclude_series = ["Bs", "Da v4"]
  }
  
  // ----------------------------------------------------------------------
}
```

```
output "ocean_id" {
  value = spotinst_ocean_aks_np.example.id
}
```

## Argument Reference

The following arguments are supported:
* `aks` - (Required) AKS cluster configuration. Cannot be updated.
  * `aks_cluster_name` - (Required) The name of the AKS Cluster.
  * `aks_infrastructure_resource_group_name` - (Required) The name of the cluster's infrastructure resource group.
  * `aks_region` - (Required) The cluster's region.
  * `aks_resource_group_name` - (Required) The name of the cluster's resource group.
* `autoscaler` - (Optional) The Ocean Kubernetes Autoscaler object.
    * `autoscale_is_enabled` - (Optional, Default: `true`) Enable the Ocean Kubernetes Autoscaler.
    * `autoscale_down` - (Optional) Auto Scaling scale down operations.
        * `max_scale_down_percentage` - (Optional) The maximum percentage allowed to scale down in a single scaling action.
    * `resource_limits` - (Optional) Optionally set upper and lower bounds on the resource usage of the cluster.
        * `max_vcpu` - (Optional) The maximum cpu in vCpu units that can be allocated to the cluster.
        * `max_memory_gib` - (Optional) The maximum memory in GiB units that can be allocated to the cluster.
    * `autoscale_headroom` - (Optional) Spare resource capacity management enabling fast assignment of pods without waiting for new resources to launch.
        * `automatic` - (Optional) [Automatic headroom](https://docs.spot.io/ocean/features/headroom?id=automatic-headroom) configuration.
            * `percentage` - (Optional) Optionally set a number between 0-100 to control the percentage of total cluster resources dedicated to headroom.
* `controller_cluster_id` - (Required) Enter a unique Ocean cluster identifier. Cannot be updated. This needs to match with string that was used to install the controller in the cluster, typically clusterName + 8 digit string.
* `health` - (Optional) The Ocean AKS Health object.
    * `grace_period` - (Optional, Default: `600`) The amount of time to wait, in seconds, from the moment the instance has launched until monitoring of its health checks begins.
* `name` - (Required) Add a name for the Ocean cluster.
* `scheduling` - (Optional) An object used to specify times when the cluster will turn off. Once the shutdown time will be over, the cluster will return to its previous state.
    * `shutdown_hours` - (Optional) [Shutdown Hours](https://docs.spot.io/ocean/features/running-hours?id=shutdown-hours)An object used to specify times that the nodes in the cluster will be taken down.
        * `is_enabled` - (Optional) Flag to enable or disable the shutdown hours mechanism. When False, the mechanism is deactivated, and the cluster remains in its current state.
        * `time_windows` - (Optional) The times that the shutdown hours will apply.
* `headrooms` - (Optional) Specify the custom headroom per VNG. Provide a list of headroom objects.
  * `cpu_per_unit` - (Optional) Configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
  * `memory_per_unit` - (Optional) Configure the amount of memory (MiB) to allocate the headroom.
  * `gpu_per_unit` - (Optional) Amount of GPU to allocate for headroom unit.
  * `num_of_units` - (Optional) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.
* `availability_zones` - (Optional) An Array holding Availability Zones, this configures the availability zones the Ocean may launch instances in per VNG.
* `labels` - (Optional) An array of labels to add to the virtual node group. Only custom user labels are allowed, and not [Kubernetes well-known labels](https://kubernetes.io/docs/reference/labels-annotations-taints/) or [ Azure AKS labels](https://learn.microsoft.com/en-us/azure/aks/use-labels) or [Spot labels](https://docs.spot.io/ocean/features/labels-and-taints?id=spot-labels).
    * `key` - (Required) Set label key [spot labels](https://docs.spot.io/ocean/features/labels-and-taints?id=spotinstionode-lifecycle) and [Azure labels](https://learn.microsoft.com/en-us/azure/aks/use-labels). The following are not allowed: ["kubernetes.azure.com/agentpool","kubernetes.io/arch","kubernetes.io/os","node.kubernetes.io/instance-type", "topology.kubernetes.io/region", "topology.kubernetes.io/zone", "kubernetes.azure.com/cluster", "kubernetes.azure.com/mode", "kubernetes.azure.com/role", "kubernetes.azure.com/scalesetpriority", "kubernetes.io/hostname", "kubernetes.azure.com/storageprofile", "kubernetes.azure.com/storagetier", "kubernetes.azure.com/instance-sku", "kubernetes.azure.com/node-image-version", "kubernetes.azure.com/subnet", "kubernetes.azure.com/vnet", "kubernetes.azure.com/ppg", "kubernetes.azure.com/encrypted-set", "kubernetes.azure.com/accelerator", "kubernetes.azure.com/fips_enabled", "kubernetes.azure.com/os-sku"]
    * `value` - (Required) Set label value.
* `max_count` - (Optional) Maximum node count limit.
* `min_count` - (Optional) Minimum node count limit.
* `enable_node_public_ip` - (Optional) Enable node public IP.
* `max_pods_per_node` - (Optional) The maximum number of pods per node in the node pools.
* `os_disk_size_gb` - (Optional) The size of the OS disk in GB.
* `os_disk_type` - (Optional, Enum:`"Managed" ,"Ephemeral"`) The type of the OS disk.
* `os_type` - (Optional, Enum:`"Linux","Windows"`) The OS type of the OS disk.
* `fallback_to_ondemand` - (Optional, Default: `true`) If no spot VM markets are available, enable Ocean to launch regular (pay-as-you-go) nodes instead.
* `spot_percentage` - (Optional,Default: `100`) Percentage of spot VMs to maintain.
* `tag` - (Optional) A maximum of 10 unique key-value pairs for VM tags in the virtual node group.
    * `key` - (Optional) Tag key for VMs in the cluster.
    * `value` - (Optional) Tag value for VMs in the cluster.
* `taints` - (Optional) Add taints to a virtual node group. Only custom user taints are allowed, and not [Kubernetes well-known taints](https://kubernetes.io/docs/reference/labels-annotations-taints/) or Azure AKS [ScaleSetPrioirty (Spot VM) taint](https://learn.microsoft.com/en-us/azure/aks/spot-node-pool). For all Spot VMs, AKS injects a taint kubernetes.azure.com/scalesetpriority=spot:NoSchedule, to ensure that only workloads that can handle interruptions are scheduled on Spot nodes. To [schedule a pod to run on Spot node](https://learn.microsoft.com/en-us/azure/aks/spot-node-pool#schedule-a-pod-to-run-on-the-spot-node), add a toleration but dont include the nodeAffinity (not supported for Spot Ocean), this will prevent the pod from being scheduled using Spot Ocean.
    * `key` - (Optional) Set taint key. The following taint keys are not allowed: ["node.kubernetes.io/not-ready",  "node.kubernetes.io/unreachable", "node.kubernetes.io/unschedulable",  "node.kubernetes.io/memory-pressure",  "node.kubernetes.io/disk-pressure",  "node.kubernetes.io/network-unavailable",  "node.kubernetes.io/pid-pressure",  "node.kubernetes.io/out-of-service",  "node.cloudprovider.kubernetes.io/uninitialized",  "node.cloudprovider.kubernetes.io/shutdown", "kubernetes.azure.com/scalesetpriority"]
    * `value` - (Optional) Set taint value.
    * `effect` - (Optional, Enum: `"NoSchedule", "PreferNoSchedule", "NoExecute", "PreferNoExecute"`) Set taint effect.
* `filters` - (Optional) Filters for the VM sizes that can be launched from the virtual node group.
    * `architectures` - (Optional, Enum `"x86_64", "intel64", "amd64", "arm64"`) The filtered vm sizes will support at least one of the architectures from this list. x86_64 includes both intel64 and amd64.
    * `max_memory_gib` - (Optional) Maximum amount of Memory (GiB).
    * `max_vcpu` - (Optional) Maximum number of vcpus available.
    * `min_memory_gib` - (Optional) Minimum amount of Memory (GiB).
    * `min_vcpu` - (Optional) Minimum number of vcpus available.
    * `series` - (Optional) Vm sizes belonging to a series from the list will be available for scaling. We can specify include list and series can be specified with capital or small letters, with space, without space or with underscore '_' .  For example all of these "DSv2", "Ds v2", "ds_v2" refer to same DS_v2 series.
    * `exclude_series` - (Optional) Vm sizes belonging to a series from the list will not be available for scaling
