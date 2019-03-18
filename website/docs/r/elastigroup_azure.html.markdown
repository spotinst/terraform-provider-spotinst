---
layout: "spotinst"
page_title: "Spotinst: elastigroup_azure"
sidebar_current: "docs-do-resource-elastigroup_azure"
description: |-
 Provides a Spotinst elastigroup resource for Microsoft Azure.
---

# spotinst\_elastigroup\_azure

Provides a Spotinst elastigroup Azure resource.

## Example Usage

```hcl
resource "elastigroup_azure" "test_azure_group" {
  name                = "example_elastigroup_azure"
  resource_group_name = "spotinst-azure"
  region              = "eastus"
  product             = "Linux"

  user_data       = ""
  shutdown_script = ""

  // --- CAPACITY ------------------------------------------------------
  min_size         = 0
  max_size         = 1
  desired_capacity = 1
  // -------------------------------------------------------------------

  // --- INSTANCE TYPES ------------------------------------------------
  od_sizes           = ["standard_a1_v1", "standard_a1_v2"]
  low_priority_sizes = ["standard_a1_v1", "standard_a1_v2"]
  // -------------------------------------------------------------------

  // --- IMAGE ---------------------------------------------------------
  image = {
    marketplace = {
      publisher = "Canonical"
      offer     = "UbuntuServer"
      sku       = "16.04-LTS"
    }
  }
  // -------------------------------------------------------------------

  // --- STRATEGY ------------------------------------------------------
  strategy = {
    od_count          = 1
    draining_timeout  = 300
  }
  // -------------------------------------------------------------------

  // --- LOAD BALANCERS ------------------------------------------------
  load_balancers = [{
    type          = "MULTAI_TARGET_SET"
    balancer_id   = "lb-1ee2e3q"
    target_set_id = "ts-3eq"
    auto_weight   = true
  }]
  // -------------------------------------------------------------------

  // --- HEALTH-CHECKS -------------------------------------------------
  health_check = {
    health_check_type = "INSTANCE_STATE"
    grace_period      = 120
    auto_healing      = true
  }
  // -------------------------------------------------------------------

  // --- NETWORK -------------------------------------------------------
  network = {
    virtual_network_name = "vname"
    subnet_name          = "my-subnet-name"
    resource_group_name  = "subnetResourceGroup"
    assign_public_ip     = true
  }
  // -------------------------------------------------------------------

  // --- LOGIN ---------------------------------------------------------
  login = {
    user_name      = "admin"
    ssh_public_key = "33a2s1f3g5a1df5g1ad3f2g1adfg56dfg=="
  }
  // -------------------------------------------------------------------
  
  // --- SCHEDULED TASK ------------------------------------------------
  scheduled_task = [{
    is_enabled      = true
    cron_expression = "* * * * *"
    task_type       = "scale"
    
    scale_min_capacity = 5
    scale_max_capacity = 8
    adjustment         = 2
    
    adjustment_percentage = 50
    scale_target_capacity = 6
    batch_size_percentage = 33
    grace_period          = 300
  }]
 // -------------------------------------------------------------------
 
 // --- SCALING POLICIES ----------------------------------------------
   scaling_up_policy = [{
       policy_name = "policy-name"
       metric_name = "CPUUtilization"
       namespace   = "Microsoft.Compute"
       statistic   = "average"
       threshold   = 10
       unit        = "percent"
       cooldown    = 60
       
       dimensions = [
         {
           name  = "resourceName"
           value = "resource-name"
         },
         {
           name  = "resourceGroupName"
           value = "resource-group-name"
         },
       ]
       
       operator            = "gt"
       evaluation_periods  = "10"
       period              = "60"
       action_type         = "setMinTarget"
       min_target_capacity = 1
     }]
 
     scaling_down_policy = [{
       policy_name = "policy-name"
       metric_name = "CPUUtilization"
       namespace   = "Microsoft.Compute"
       statistic   = "average"
       threshold   = 10
       unit        = "percent"
       cooldown    = 60
       
       dimensions = {
           name  = "name-1"
           value = "value-1"
       }
       
       operator           = "gt"
       evaluation_periods = "10"
       period             = "60"
       action_type        = "adjustment"
       adjustment         = "MIN(5,10)"
     }]
 // -------------------------------------------------------------------
 
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The group name.
* `region` - (Required) The region your Azure group will be created in.
* `resource_group_name` - (Required) Name of the Resource Group for Elastigroup.
* `product` - (Required) Operation system type. Valid values: `"Linux"`, `"Windows"`.
* `max_size` - (Required) The maximum number of instances the group should have at any time.
* `min_size` - (Required) The minimum number of instances the group should have at any time.
* `desired_capacity` - (Required) The desired number of instances the group should have at any time.

* `od_sizes` - (Required) Available On-Demand sizes
* `low_priority_sizes` - (Required) Available Low-Priority sizes.

* `user_data` - (Optional) Base64-encoded MIME user data to make available to the instances.
* `shutdown_script` - (Optional) Shutdown script for the group. Value should be passed as a string encoded at Base64 only.

* `strategy` - (Required) Describes the deployment strategy.
* `low_priority_percentage` - (Optional, Default `100`) Percentage of Low Priority instances to maintain. Required if `od_count` is not specified.
* `od_count` - (Optional) Number of On-Demand instances to maintain. Required if low_priority_percentage is not specified.
* `draining_timeout` - (Optional, Default `120`) Time (seconds) to allow the instance to be drained from incoming TCP connections and detached from MLB before terminating it during a scale-down operation.

<a id="load-balancers"></a>
## Load Balancers

* `load_balancers` - (Required) Describes a set of one or more classic load balancer target groups and/or Multai load balancer target sets.
* `type` - (Required) The resource type. Valid values: CLASSIC, TARGET_GROUP, MULTAI_TARGET_SET.
* `balancer_id` - (Required) The balancer ID.
* `target_set_id` - (Required) The scale set ID associated with the load balancer.
* `auto_weight` - (Optional, Default: `false`)

```hcl
  load_balancers = [{
    type          = "MULTAI_TARGET_SET"
    balancer_id   = "lb-1ee2e3q"
    target_set_id = "ts-3eq"
    auto_weight   = true
  }]
```

<a id="image"></a>
## Image

* `image` - (Required) Image of a VM. An image is a template for creating new VMs. Choose from Azure image catalogue (marketplace) or use a custom image.
* `publisher` - (Optional) Image publisher. Required if resource_group_name is not specified.
* `offer` - (Optional) Name of the image to use. Required if publisher is specified.
* `sku` - (Optional) Image's Stock Keeping Unit, which is the specific version of the image. Required if publisher is specified.
* `resource_group_name` - (Optional) Name of Resource Group for custom image. Required if publisher not specified.
* `image_name` - (Optional) Name of the custom image. Required if resource_group_name is specified.

```hcl
  // market image
  image = {
    marketplace = {
      publisher = "Canonical"
      offer     = "UbuntuServer"
      sku       = "16.04-LTS"
    }
  }
  
  // custom image
  image = {
    custom = {
      image_name          = "customImage"
      resource_group_name = "resourceGroup"
    }
  } 
```

<a id="health-check"></a>
## Health Check

* `health_check` - (Optional) Describes the health check configuration.
* `health_check_type` - (Optional) Health check used to validate VM health. Valid values: “INSTANCE_STATE”.
* `grace_period` - (Optional) Period of time (seconds) to wait for VM to reach healthiness before monitoring for unhealthiness.
* `auto_healing` - (Optional) Enable auto-healing of unhealthy VMs.

```hcl
  health_check = {
    health_check_type = "INSTANCE_STATE"
    grace_period      = 120
    auto_healing      = true
  }
```

<a id="network"></a>
## Network

* `network` - (Required) Defines the Virtual Network and Subnet for your Elastigroup.
* `virtual_network_name` - (Required) Name of Vnet.
* `subnet_name` - (Required) ID of subnet.
* `resource_group_name` - (Required) Vnet Resource Group Name.
* `assign_public_up` - (Optional, Default: `false`) Assign a public IP to each VM in the Elastigroup.
* `additional_ip_configs` - (Optional) Array of additional IP configuration objects.
* `name` - (Required) The IP configuration name.
* `private_ip_version` - (Optional) Available from Azure Api-Version 2017-03-30 onwards, it represents whether the specific ipconfiguration is IPv4 or IPv6. Valid values: `IPv4`, `IPv6`.

```hcl
  network = {
    virtual_network_name = "vname"
    subnet_name          = "my-subnet-name"
    resource_group_name  = "subnetResourceGroup"
    assign_public_ip     = true
    
    additional_ip_configs = [{
      name = "test"
      private_ip_version = "IPv4"
    }]
  }
```

<a id="login"></a>
## Login

```hcl
  network = {
    virtual_network_name = "vname"
    subnet_name          = "my-subnet-name"
    resource_group_name  = "subnetResourceGroup"
    assign_public_ip     = true
  }
```

<a id="login"></a>
## Login

* `login` - (Required) Describes the login configuration.
* `user_name` - (Required) Set admin access for accessing your VMs.
* `ssh_public_key` - (Optional) SSH for admin access to Linux VMs. Required for Linux product types.
* `password` - (Optional) Password for admin access to Windows VMs. Required for Windows product types.

```hcl
  login = {
    user_name      = "admin"
    ssh_public_key = "33a2s1f3g5a1df5g1ad21651sag56dfg=="
  }
```

<a id="scaling-policy"></a>
## Scaling Policies

Each `scaling_*_policy` supports the following:

* `policy_name` - (Optional) The name of the policy.
* `metric_name` - (Required) Metric to monitor by Azure metric display name.
* `namespace` - (Optional, Default: `“Microsoft.Compute”`) The namespace for the alarm's associated metric. Valid values: 

```text
  Microsoft.AnalysisServices/servers 
  Microsoft.ApiManagement/service 
  Microsoft.Automation/automationAccounts 
  Microsoft.Batch/batchAccounts 
  Microsoft.Cache/redis 
  Microsoft.CognitiveServices/accounts 
  Microsoft.Compute 
  Microsoft.ContainerInstance/containerGroups 
  Microsoft.ContainerService/managedClusters 
  Microsoft.CustomerInsights/hubs 
  Microsoft.DataFactory/datafactories 
  Microsoft.DataFactory/factories 
  Microsoft.DataLakeAnalytics/accounts 
  Microsoft.DataLakeStore/accounts 
  Microsoft.DBforMariaDB/servers 
  Microsoft.DBforMySQL/servers 
  Microsoft.DBforPostgreSQL/servers 
  Microsoft.Devices/IotHubs 
  Microsoft.Devices/provisioningServices 
  Microsoft.DocumentDB/databaseAccounts 
  Microsoft.EventGrid/eventSubscriptions 
  Microsoft.EventGrid/extensionTopics 
  Microsoft.EventGrid/topics
  Microsoft.EventHub/clusters 
  Microsoft.EventHub/namespaces 
  Microsoft.HDInsight/clusters 
  Microsoft.Insights/AutoscaleSettings 
  Microsoft.Insights/Components 
  Microsoft.KeyVault/vaults 
  Microsoft.Kusto/Clusters 
  Microsoft.LocationBasedServices/accounts 
  Microsoft.Logic/workflows 
  Microsoft.NetApp/netAppAccounts/capacityPools/Volumes 
  Microsoft.NetApp/netAppAccounts/capacityPools 
  Microsoft.Network/applicationGateways 
  Microsoft.Network/dnszones 
  Microsoft.Network/connections 
  Microsoft.Network/expressRouteCircuits 
  Microsoft.Network/expressRouteCircuits/peerings 
  Microsoft.Network/frontdoors 
  Microsoft.Network/loadBalancers 
  Microsoft.Network/networkInterfaces 
  Microsoft.Network/networkWatchers/connectionMonitors 
  Microsoft.Network/publicIPAddresses 
  Microsoft.Network/trafficManagerProfiles 
  Microsoft.Network/virtualNetworkGateways 
  Microsoft.NotificationHubs/Namespaces/NotificationHubs 
  Microsoft.OperationalInsights/workspaces 
  Microsoft.PowerBIDedicated/capacities 
  Microsoft.Relay/namespaces 
  Microsoft.Search/searchServices 
  Microsoft.ServiceBus/namespaces 
  Microsoft.SignalRService/SignalR 
  Microsoft.Sql/managedInstances 
  Microsoft.Sql/servers/databases 
  Microsoft.Sql/servers/elasticPools 
  Microsoft.Storage/storageAccounts 
  Microsoft.Storage/storageAccounts/blobServices 
  Microsoft.Storage/storageAccounts/fileServices 
  Microsoft.Storage/storageAccounts/queueServices 
  Microsoft.Storage/storageAccounts/tableServices 
  Microsoft.StreamAnalytics/streamingjobs 
  Microsoft.TimeSeriesInsights/environments 
  Microsoft.TimeSeriesInsights/environments/eventsources 
  Microsoft.Web/hostingEnvironments/multiRolePools 
  Microsoft.Web/hostingEnvironments/workerPools 
  Microsoft.Web/serverfarms 
  Microsoft.Web/sites (excluding functions) 
  Microsoft.Web/sites (functions) 
  Microsoft.Web/sites/slots 
```
  
* `statistic` - (Optional) The metric statistics to return. Valid values: `average`.
* `threshold` - (Required) The value against which the specified statistic is compared.
* `unit` - (Required) The unit for the alarm's associated metric. Valid values: `"percent`, `"seconds"`, `"microseconds"`, `"milliseconds"`, `"bytes"`, `"kilobytes"`, `"megabytes"`, `"gigabytes"`, `"terabytes"`, `"bits"`, `"kilobits"`, `"megabits"`, `"gigabits"`, `"terabits"`, `"count"`, `"bytes/second"`, `"kilobytes/second"`, `"megabytes/second"`, `"gigabytes/second"`, `"terabytes/second"`, `"bits/second"`, `"kilobits/second"`, `"megabits/second"`, `"gigabits/second"`, `"terabits/second"`, `"count/second"`, `"none"`.  
* `cooldown` - (Optional, Default: `300`) The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start. If this parameter is not specified, the default cooldown period for the group applies.
* `operator` - (Optional, Scale Up Default: `gte`, Scale Down Default: `lte`) The operator to use in order to determine if the scaling policy is applicable. Valid values: `"gt"`, `"gte"`, `"lt"`, `"lte"`.
* `evaluation_periods` - (Optional, Default: `1`) The number of periods over which data is compared to the specified threshold.
* `period` - (Optional, Default: `300`) The granularity, in seconds, of the returned datapoints. Period must be at least 60 seconds and must be a multiple of 60.

* `dimensions` - (Optional) A list of dimensions describing qualities of the metric. Required when `namespace` is defined AND not `"Microsoft.Compute"`.
    * `name` - (Required) The dimension name.
    * `value` - (Optional) The dimension value.
    
When `namespace` is defined and is not `"Microsoft.Compute"` the list of dimensions must contain the following:

```hcl
  dimensions = [
    {
      name  = "resourceName"
      value = "example-resource-name"
    },
    {
      name  = "resourceGroupName"
      value = "example-resource-group-name"
    },
  ]
```

* `action_type` - (Optional; if not using `min_target_capacity` or `max_target_capacity`) The type of action to perform for scaling. Valid values: `"adjustment"`, `"percentageAdjustment"`, `"setMaxTarget"`, `"setMinTarget"`, `"updateCapacity"`.

If you do not specify an action type, you can only use – `adjustment`, `min_target_capacity`, `max_target_capacity`.
While using action_type, please also set the following:

When using `adjustment`           – set the field `adjustment`
When using `percentageAdjustment` - set the field `adjustment`
When using `setMaxTarget`         – set the field `max_target_capacity`
When using `setMinTarget`         – set the field `min_target_capacity`
When using `updateCapacity`       – set the fields `minimum`, `maximum`, and `target`

* `adjustment` - (Optional) Value to which the action type will be adjusted. Required if using `numeric` or `percentage_adjustment` action types.
* `min_target_capacity` - (Optional; if not using `adjustment`; available only for scale up). The number of the desired target (and minimum) capacity
* `max_target_capacity` - (Optional; if not using `adjustment`; available only for scale down). The number of the desired target (and maximum) capacity

* `minimum` - (Optional; if using `updateCapacity`) The minimal number of instances to have in the group.
* `maximum` - (Optional; if using `updateCapacity`) The maximal number of instances to have in the group.
* `target` - (Optional; if using `updateCapacity`) The target number of instances to have in the group.

Usage:

```hcl
// --- SCALE DOWN POLICY ------------------
  scaling_down_policy = [{
    policy_name = "policy-name"
    metric_name = "CPUUtilization"
    namespace   = "Microsoft.Compute"
    statistic   = "average"
    threshold   = 10
    unit        = "percent"
    cooldown    = 60
    
    dimensions = {
      name  = "name-1"
      value = "value-1"
    }
    
    operator           = "gt"
    evaluation_periods = "10"
    period             = "60"
    
    // === MIN TARGET ===================
    # action_type         = "setMinTarget"
    # min_target_capacity = 1
    // ==================================
    
    // === ADJUSTMENT ===================
    action_type   = "adjustment"
    # action_type = "percentageAdjustment"
    adjustment    = "MIN(5,10)"
    // ==================================
    
    // === UPDATE CAPACITY ==============
    # action_type = "updateCapacity"
    # minimum     = 0
    # maximum     = 10
    # target      = 5
    // ==================================
    
  }]
// ----------------------------------------

// --- SCALE DOWN POLICY ------------------
  scaling_down_policy = [{
    policy_name = "policy-name-update"
    metric_name = "CPUUtilization"
    namespace   = "Microsoft.Compute"
    statistic   = "sum"
    threshold   = 5
    unit        = "bytes"
    cooldown    = 120
    
    dimensions = {
        name  = "name-1-update"
        value = "value-1-update"
    }
    
    operator           = "lt"
    evaluation_periods = 5
    period             = 120
    
    //// === MIN TARGET ===================
    # action_type         = "setMinTarget"
    # min_target_capacity = 1
    //// ==================================
    
    // === ADJUSTMENT ===================
    # action_type = "percentageAdjustment"
    # action_type = "adjustment"
    # adjustment  = "MAX(5,10)"
    // ==================================
    
    // === UPDATE CAPACITY ==============
    action_type = "updateCapacity"
    minimum     = 0
    maximum     = 10
    target      = 5
    // ==================================
    
  }]
// ----------------------------------------
```

<a id="scheduling"></a>
## Scheduling

* `scheduled_task` - (Optional) Describes the configuration of one or more scheduled tasks.
* `is_enabled` - (Optional, Default: `true`) Describes whether the task is enabled. When true the task should run when false it should not run.
* `cron_expression` - (Required) A valid cron expression (`* * * * *`). The cron is running in UTC time zone and is in Unix cron format Cron Expression Validator Script.
* `task_type` - (Required) The task type to run. Valid Values: `backup_ami`, `scale`, `scaleUp`, `roll`, `statefulUpdateCapacity`, `statefulRecycle`.
* `scale_min_capacity` - (Optional) The min capacity of the group. Should be used when choosing ‘task_type' of ‘scale'.
* `scale_max_capacity` - (Optional) The max capacity of the group. Required when ‘task_type' is ‘scale'.
* `scale_target_capacity` - (Optional) The target capacity of the group. Should be used when choosing ‘task_type' of ‘scale'.
* `adjustment` - (Optional) The number of instances to add/remove to/from the target capacity when scale is needed.
* `adjustment_percentage` - (Optional) The percent of instances to add/remove to/from the target capacity when scale is needed.
* `batch_size_percentage` - (Optional) The percentage size of each batch in the scheduled deployment roll. Required when the 'task_type' is 'roll'.
* `grace_period` - (Optional) The time to allow instances to become healthy.

```hcl
  scheduled_task = [{
    is_enabled      = true
    cron_expression = "* * * * *"
    task_type       = "scale"
    
    scale_min_capacity = 5
    scale_max_capacity = 8
    adjustment         = 2
    
    adjustment_percentage = 50
    scale_target_capacity = 6
    batch_size_percentage = 33
    grace_period          = 300
  }]
```

<a id="update-policy"></a>
## Update Policy

* `update_policy` - (Optional)

    * `should_roll` - (Required) Sets the enablement of the roll option.
    * `roll_config` - (Required) While used, you can control whether the group should perform a deployment after an update to the configuration.
        * `batch_size_percentage` - (Required) Sets the percentage of the instances to deploy in each batch.
        * `health_check_type` - (Optional) Sets the health check type to use. Valid values: `"INSTANCE_STATE"`, `"NONE"`.
        * `grace_period` - (Optional) Sets the grace period for new instances to become healthy.
       
```hcl
  update_policy = {
    should_roll = false
    
    roll_config = {
      batch_size_percentage = 33
      health_check_type     = "INSTANCE_STATE"
      grace_period          = 300
    }
  }
```        

<a id="third-party-integrations"></a>
## Third-Party Integrations

* `integration_kubernetes` - (Optional) Describes the [Kubernetes](https://kubernetes.io/) integration.
    * `cluster_identifier` - (Required) The cluster ID.

Usage:

```hcl
  integration_kubernetes = {
    cluster_identifier = "k8s-cluster-id"
  }
```

* `integration_multai_runtime` - (Optional) Describes the [Multai Runtime](https://spotinst.com/) integration.
    * `deployment_id` - (Optional) The deployment id you want to get

Usage:

```hcl
  integration_multai_runtime = {
    deployment_id = ""
  }
```  