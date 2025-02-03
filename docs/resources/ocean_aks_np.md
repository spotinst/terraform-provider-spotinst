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

Installation of the Ocean controller is required by this resource. You can accomplish this by using the [spotinst/terraform-ocean-kubernetes-controller ](https://registry.terraform.io/modules/spotinst/kubernetes-controller/ocean) module as follows:

```hcl
module "kubernetes-controller" {
  source = "spotinst/kubernetes-controller/ocean"

  # Credentials.
  spotinst_token   = "redacted"
  spotinst_account = "redacted"

  # Configuration.
  cluster_identifier = "ocean-aks"
}
```

~> You must configure the same `cluster_identifier` for the Ocean controller and for the `spotinst_ocean_aks_np` resource.

## Basic Ocean Cluster Creation Usage Example - using minimum configuration with only required parameters

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
  
## Detailed Ocean Cluster Creation Usage Example - using all available parameters with sample values

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
        is_enabled = true
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
  
  // ---- Logging ---------------------------------------------------------
  
  logging {
    export {
      azure_blob {
        id = "di-abcd123"
      }
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
  kubernetes_version    = "1.26"
  pod_subnet_ids       = ["/subscriptions/123456-1234-1234-1234-123456789/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/virtualNetworks/ExampleVirtualNetwork/subnets/default"]
  vnet_subnet_ids       = ["/subscriptions/123456-1234-1234-1234-123456789/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/virtualNetworks/ExampleVirtualNetwork/subnets/default"]
  linux_os_config {
    sysctls {
      vm_max_map_count = 79550
    }
  }
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
  
  vng_template_scheduling {
    vng_template_shutdown_hours {
      is_enabled   = true
      time_windows = [
        "Fri:15:30-Sat:13:30", 
        "Sun:15:30-Mon:13:30",
      ]
    }
  }

  tags = {
    tagKey   = "env"
    tagValue = "staging"
  }
  // --- vmSizes ----------------------------------------------------------
  
  filters {
    min_vcpu               = 2
    max_vcpu               = 16
    min_memory_gib         = 8
    max_memory_gib         = 128
    architectures          = ["x86_64", "arm64"]
    series                 = ["D v3", "Dds_v4", "Dsv2"]
    exclude_series         = ["Av2", "A", "Bs", "D", "E"]
    accelerated_networking = "Enabled"
    disk_performance       = "Premium"
    min_gpu                = 1
    max_gpu                = 2
    min_nics               = 1
    vm_types               = ["generalPurpose", "GPU"]
    min_disk               = 1
    gpu_types              = ["nvidia-tesla-t4"]
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
  * `autoscale_is_enabled` - (Optional) Enable the Ocean Kubernetes Autoscaler.
  * `autoscale_down` - (Optional) Auto Scaling scale down operations.
    * `max_scale_down_percentage` - (Optional) The maximum percentage allowed to scale down in a single scaling action.
  * `resource_limits` - (Optional) Optionally set upper and lower bounds on the resource usage of the cluster.
    * `max_vcpu` - (Optional) The maximum cpu in vCpu units that can be allocated to the cluster.
    * `max_memory_gib` - (Optional) The maximum memory in GiB units that can be allocated to the cluster.
  * `autoscale_headroom` - (Optional) Spare resource capacity management enabling fast assignment of pods without waiting for new resources to launch.
    * `automatic` - (Optional) [Automatic headroom](https://docs.spot.io/ocean/features/headroom?id=automatic-headroom) configuration.
      * `is_enabled` - (Optional, Default - false) Enable automatic headroom. When set to `true`, Ocean configures and optimizes headroom automatically.
      * `percentage` - (Optional) Optionally set a number between 0-100 to control the percentage of total cluster resources dedicated to headroom.
* `controller_cluster_id` - (Required) Enter a unique Ocean cluster identifier. Cannot be updated. This needs to match with string that was used to install the controller in the cluster, typically clusterName + 8 digit string.
* `health` - (Optional) The Ocean AKS Health object.
  * `grace_period` - (Optional, Default: `600`) The amount of time to wait, in seconds, from the moment the instance has launched until monitoring of its health checks begins.
* `name` - (Required) Add a name for the Ocean cluster.
* `headrooms` - (Optional) Specify the custom headroom per VNG. Provide a list of headroom objects.
  * `cpu_per_unit` - (Optional) Configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
  * `memory_per_unit` - (Optional) Configure the amount of memory (MiB) to allocate the headroom.
  * `gpu_per_unit` - (Optional) Amount of GPU to allocate for headroom unit.
  * `num_of_units` - (Optional) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.
* `availability_zones` - (Required) An Array holding Availability Zones, this configures the availability zones the Ocean may launch instances in per VNG.
* `labels` - (Optional) An array of labels to add to the virtual node group. Only custom user labels are allowed, and not [Kubernetes well-known labels](https://kubernetes.io/docs/reference/labels-annotations-taints/) or [ Azure AKS labels](https://learn.microsoft.com/en-us/azure/aks/use-labels) or [Spot labels](https://docs.spot.io/ocean/features/labels-and-taints?id=spot-labels).
  * `key` - (Required) Set label key [spot labels](https://docs.spot.io/ocean/features/labels-and-taints?id=spotinstionode-lifecycle) and [Azure labels](https://learn.microsoft.com/en-us/azure/aks/use-labels). The following are not allowed: ["kubernetes.azure.com/agentpool","kubernetes.io/arch","kubernetes.io/os","node.kubernetes.io/instance-type", "topology.kubernetes.io/region", "topology.kubernetes.io/zone", "kubernetes.azure.com/cluster", "kubernetes.azure.com/mode", "kubernetes.azure.com/role", "kubernetes.azure.com/scalesetpriority", "kubernetes.io/hostname", "kubernetes.azure.com/storageprofile", "kubernetes.azure.com/storagetier", "kubernetes.azure.com/instance-sku", "kubernetes.azure.com/node-image-version", "kubernetes.azure.com/subnet", "kubernetes.azure.com/vnet", "kubernetes.azure.com/ppg", "kubernetes.azure.com/encrypted-set", "kubernetes.azure.com/accelerator", "kubernetes.azure.com/fips_enabled", "kubernetes.azure.com/os-sku"]
  * `value` - (Required) Set label value.
* `max_count` - (Optional, Default: 1000) Maximum node count limit.
* `min_count` - (Optional, Default: 0) Minimum node count limit.
* `enable_node_public_ip` - (Optional) Enable node public IP.
* `max_pods_per_node` - (Optional) The maximum number of pods per node in the node pools.
* `os_disk_size_gb` - (Optional) The size of the OS disk in GB.
* `os_disk_type` - (Optional, Enum:`"Managed" ,"Ephemeral"`) The type of the OS disk.
* `os_type` - (Optional, Enum:`"Linux","Windows"`) The OS type of the OS disk. Can't be modified once set.
* `os_sku` - (Optional, Enum: `"Ubuntu", "Windows2019", "Windows2022", "AzureLinux", "CBLMariner"`) The OS SKU of the OS type. Must correlate with the os type.
* `pod_subnet_ids` - (Optional) The IDs of subnets in an existing VNet into which to assign pods in the cluster (requires azure network-plugin).
* `vnet_subnet_ids` - (Optional) The IDs of subnets in an existing VNet into which to assign nodes in the cluster (requires azure network-plugin).
* `linux_os_config` - (Optional) Custom Linux OS configuration.
  * `sysctls` - (Optional) System Controls
    * `vm_max_map_count` - (Optional) Maximum number of memory map areas a process may have. Can be configured only if OS type is Linux.
* `kubernetes_version` - (Optional) The desired Kubernetes version of the launched nodes. In case the value is null, the Kubernetes version of the control plane is used.
* `fallback_to_ondemand` - (Optional) If no spot VM markets are available, enable Ocean to launch regular (pay-as-you-go) nodes instead.
* `spot_percentage` - (Optional) Percentage of spot VMs to maintain.
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
  * `accelerated_networking` - (Optional, Enum `"Enabled", "Disabled"`) In case acceleratedNetworking is set to Enabled, accelerated networking applies only to the VM that enables it.
  * `disk_performance` - (Optional, Enum `"Standard", "Premium"`) The filtered vm sizes will support at least one of the classes from this list.
  * `min_gpu` - (Optional) Minimum number of GPUs available.
  * `max_gpu` - (Optional) Maximum number of GPUs available.
  * `min_nics` - (Optional) Minimum number of network interfaces.
  * `min_disk` - (Optional) Minimum number of data disks available.
  * `vm_types` - (Optional, Enum `"generalPurpose", "memoryOptimized", "computeOptimized", "highPerformanceCompute", "storageOptimized", "GPU"`) The filtered vm types will belong to one of the vm types from this list.
  * `gpu_types` - (Optional, Enum `"nvidia-tesla-v100", "amd-radeon-instinct-mi25", "nvidia-a10", "nvidia-tesla-a100", "nvidia-tesla-k80", "nvidia-tesla-m60", "nvidia-tesla-p100", "nvidia-tesla-p40", "nvidia-tesla-t4", "nvidia-tesla-h100"`) The filtered gpu types will belong to one of the gpu types from this list.
* `logging` - (Optional) The Ocean AKS Logging Object.
  * `export` - The Ocean AKS Logging Export object.
    * `azure_blob` -  Exports your cluster's logs to the storage account and container configured on the storage account [data integration](https://docs.spot.io/ocean/features/log-integration-with-azure-blob?id=log-integration-with-azure-blob) given. Each file contains logs of 3 minutes where each log is separated by a new line and saved as a JSON. The file formats are `container`/`accountId``oceanId``oceanName`_`startTime`.log
      * `id` - (Required) The identifier of The Azure Blob data integration to export the logs to.
* `vng_template_scheduling` - (Optional) An object used to specify times when the virtual node group will turn off all its node pools. Once the shutdown time will be over, the virtual node group will return to its previous state.
  * `shutdown_hours` - (Optional) An object used to specify times that the nodes in the virtual node group will be stopped.
    * `is_enabled` - (Optional) Flag to enable or disable the shutdown hours mechanism. When False, the mechanism is deactivated, and the virtual node gorup remains in its current state.
    * `time_windows` - (Optional) The times that the shutdown hours will apply. Required if `is_enabled` is true.

<a id="update-policy"></a>
## Update Policy

* `update_policy` - (Optional)
  * `should_roll` - (Required) If set to true along with the cluster update, roll will be triggered.
  * `conditioned_roll` - (Optional, Default: false) Spot will perform a cluster Roll in accordance with a relevant modification of the cluster’s settings. When set to true , only specific changes in the cluster’s configuration will trigger a cluster roll (such as availability_zones, max_pods_per_node, enable_node_public_ip, os_disk_size_gb, os_disk_type, os_sku, kubernetes_version, vnet_subnet_ids, pod_subnet_ids, labels, taints and tags).
  * `roll_config` - (Optional) While used, you can control whether the group should perform a deployment after an update to the configuration.
    * `batch_min_healthy_percentage` - (Optional, Default: 50) Indicates the threshold of minimum healthy nodes in single batch. If the amount of healthy nodes in single batch is under the threshold, the roll will fail. If exists, the parameter value will be in range of 1-100. In case of null as value, the default value in the backend will be 50%. Value of param should represent the number in percentage (%) of the batch.
    * `batch_size_percentage` - (Optional) Value as a percent to set the size of a batch in a roll. Valid values are 0-100. In case of null as value, the default value in the backend will be 20%.
    * `comment` - (Optional) Add a comment description for the roll. The comment is limited to 256 chars and optional.
    * `respect_pdb` - (Optional, Default: true) During the roll, if the parameter is set to true we honor PDB during the nodes replacement.
    * `respect_restrict_scale_down` - (Optional, Default: false) During the roll, if the parameter is set to true we honor Restrict Scale Down label during the nodes replacement.
    * `node_pool_names` - (Optional) List of node pools to be rolled. Each node pool name is a string. nodePoolNames can be null, and cannot be used together with nodeNames and vngIds. 
    * `node_names` - (Optional) List of node names to be rolled. Each identifier is a string. nodeNames can be null, and cannot be used together with nodePoolNames and vngIds.
    * `vng_ids` - (Optional) List of virtual node group identifiers to be rolled. Each identifier is a string. vngIds can be null, and cannot be used together with nodeNames and nodePoolNames.
```hcl
update_policy {
  should_roll = false
  conditioned_roll = true

  roll_config {
    batch_size_percentage = 25
    batch_min_healthy_percentage = 100
    respect_pdb = true
    node_names = ["aks-omnp123456-7890-vmss000001"]
  }
}
```
<a id="Scheduling"></a>
## Scheduling
* `scheduling` - (Optional) An object used to specify times when the cluster will turn off. Once the shutdown time will be over, the cluster will return to its previous state.
  * `shutdown_hours` - (Optional) An object used to specify times that the nodes in the cluster will be taken down.
    * `is_enabled` - (Optional) Flag to enable or disable the shutdown hours mechanism. When `false`, the mechanism is deactivated, and the cluster remains in its current state.
    * `time_windows` - (Optional) The times that the shutdown hours will apply. Required if isEnabled is true.
  * `suspension_hours` - (Optional) An object used to specify times that the cluster should be exempted from Ocean's scaling-down activities to ensure uninterrupted operations during critical periods.
    * `is_enabled` - (Optional) Flag to enable or disable the suspension hours mechanism. When `false`, the mechanism is deactivated, and the cluster remains in its current state.
    * `time_windows` - (Optional) The times that the suspension hours will apply. Required if isEnabled is true.
  * `tasks` - (Optional) A list of scheduling tasks to preform on the cluster at a specific cron time.
    * `is_enabled` - (Required)  Describes whether the task is enabled. When true the task should run when false it should not run. Required for `cluster.scheduling.tasks` object.
    * `cron_expression` - (Required) A valid cron expression. The cron is running in UTC time zone and is in Unix cron format Cron Expression Validator Script. Only one of `frequency` or `cronExpression` should be used at a time. Required for `cluster.scheduling.tasks` object. (Example: `0 1 * * *`).
    * `task_type` - (Required) The type of the scheduling task. Valid values: "`clusterRoll`,`autoUpgradeVersion`".
    * `parameters` - (Optional) The parameters of the scheduling task. Each task type will have properties relevant only to it.
      * `parameters_cluster_roll` - (Optional) The parameters of the cluster roll scheduling task.
        * `batch_min_healthy_percentage` - (Optional) The minimum percentage of the scaled nodes that should be healthy at each batch. Valid values are 1-100.
        * `batch_size_percentage` - (Optional) The percentage of the cluster that will be rolled at each batch. Valid values are 1-100.
        * `comment` - (Optional) A comment to be added to the cluster roll.
        * `respect_pdb` - (Optional) During the roll, if the parameter is set to true we honor PDB during the instance replacement.
        * `respect_restrict_scale_down` - (Optional) During the roll, if the parameter is set to true we honor Restrict Scale Down label during the nodes replacement.
        * `vng_ids` - (Optional) List of Virtual Node Group IDs to be rolled. If not set or set to null, cluster roll will be applied.
      * `parameters_upgrade_config` - (Optional) The parameters of the upgrade config scheduling task.
        * `apply_roll` - (Optional) - When set to True, a cluster roll will be initiated if a new version is available to upgrade in the dedicated virtual node groups.
        * `roll_parameters` - (Optional) - The parameters of the cluster roll that will be initiated.
          * `batch_min_healthy_percentage` - (Optional) The minimum percentage of the scaled nodes that should be healthy at each batch. Valid values are 1-100.
          * `batch_size_percentage` - (Optional) The percentage of the cluster that will be rolled at each batch. Valid values are 1-100.
          * `comment` - (Optional) A comment to be added to the cluster roll.
          * `respect_pdb` - (Optional) During the roll, if the parameter is set to true we honor PDB during the instance replacement.
          * `respect_restrict_scale_down` - (Optional) During the roll, if the parameter is set to true we honor Restrict Scale Down label during the nodes replacement.

```hcl
scheduling {
  shutdown_hours {
    is_enabled   = true
    time_windows = [
      "Fri:15:30-Sat:13:30", 
      "Sun:15:30-Mon:13:30",
    ]
  }
   suspension_hours {
    is_enabled   = true
    time_windows = [
      "Fri:15:30-Sun:13:30", 
      "Mon:15:30-Tue:13:30",
    ]
  }
  #task for clusterRoll
  tasks {
    is_enabled      = true
    cron_expression = "* 1 * * *"
    task_type       = "clusterRoll"
    parameters  {
      parameters_cluster_roll {
        batch_min_healthy_percentage = 50
        batch_size_percentage = 20
        comment = "Scheduled cluster roll"
        respect_pdb = true
        respect_restrict_scale_down=true
        vng_ids=["vng123","vng456"]
      }
    }
  }
  #task for autoUpgradeVersion
  tasks {
      is_enabled      = true
      cron_expression = "* 10 * * *"
      task_type       = "autoUpgradeVersion"
      parameters  {
        parameters_upgrade_config {
          apply_roll = true
          scope_version = "patch"
          roll_parameters {
            batch_min_healthy_percentage = 75
            batch_size_percentage        = 50
            comment                      = "Scheduled upgrade roll"
            respect_pdb                  = false
            respect_restrict_scale_down  = false
          }
        }
      }
    }
}
```