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

  user_data = ""

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
  
  // --- SCHEDULED TASK ------------------
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
 // -------------------------------------
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
* `user_data` - (Optional) Base64-encoded MIME user data to make available to the instances.
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