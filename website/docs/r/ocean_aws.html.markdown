---
layout: "spotinst"
page_title: "Spotinst: ocean_aws"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean resource using AWS.
---

# spotinst\_ocean\_aws

Manages a Spotinst Ocean AWS resource.

## Example Usage

```hcl
resource "spotinst_ocean_aws" "example" {
  name = "demo"
  controller_id = "fakeClusterId"
  region = "us-west-2"

  max_size         = 2
  min_size         = 1
  desired_capacity = 2

  subnet_ids = ["subnet-123456789"]
  whitelist  = ["t1.micro", "m1.small"]
// blacklist = ["t1.micro", "m1.small"]

  // --- LAUNCH CONFIGURATION --------------
  image_id             = "ami-123456"
  security_groups      = ["sg-987654321"]
  key_name             = "fake key"
  user_data            = "echo hello world"
  iam_instance_profile = "iam-profile"
  root_volume_size     = 20
  monitoring           = true
  ebs_optimized        = true
  associate_public_ip_address = true
  
  load_balancers {
    arn = "arn:aws:elasticloadbalancing:us-west-2:fake-arn"
    type = "TARGET_GROUP"
  }
  load_balancers {
    name = "AntonK"
    type = "CLASSIC"
  }
  // ---------------------------------------

  // --- STRATEGY --------------------
  fallback_to_ondemand       = true
  draining_timeout           = 120
  utilize_reserved_instances = false
  grace_period               = 600
  spot_percentage            = 100
  // ---------------------------------

  tags {
    key   = "fakeKey"
    value = "fakeValue"
  }
}
```
```
output "ocean_id" {
  value = spotinst_ocean_aws.example.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The cluster name.
* `controller_id` - (Required) The ocean cluster identifier. Example: `ocean.k8s`
* `region` - (Required) The region the cluster will run in.
* `max_size` - (Optional, Default: `1000`) The upper limit of instances the cluster can scale up to.
* `min_size` - (Optional) The lower limit of instances the cluster can scale down to.
* `desired_capacity` - (Optional) The number of instances to launch and maintain in the cluster.
* `subnet_ids` - (Required) A comma-separated list of subnet identifiers for the Ocean cluster. Subnet IDs should be configured with auto assign public ip.
* `whitelist` - (Optional) Instance types allowed in the Ocean cluster. Cannot be configured if `blacklist` is configured.
* `blacklist` - (Optional) Instance types not allowed in the Ocean cluster. Cannot be configured if `whitelist` is configured.
* `user_data` - (Optional) Base64-encoded MIME user data to make available to the instances.
* `image_id` - (Required) ID of the image used to launch the instances.
* `security_groups` - (Required) One or more security group ids.
* `key_name` - (Optional) The key pair to attach the instances.
* `iam_instance_profile` - (Optional) The instance profile iam role.
* `associate_public_ip_address` - (Optional, Default: `false`) Configure public IP address allocation.
* `root_volume_size` - (Optional) The size (in Gb) to allocate for the root volume. Minimum `20`.
* `monitoring` - (Optional) Enable detailed monitoring for cluster. Flag will enable Cloud Watch detailed detailed monitoring (one minute increments). Note: there are additional hourly costs for this service based on the region used.
* `ebs_optimized` - (Optional) Enable EBS optimized for cluster. Flag will enable optimized capacity for high bandwidth connectivity to the EB service for non EBS optimized instance types. For instances that are EBS optimized this flag will be ignored.
* `load_balancers` - (Optional) - Array of load balancer objects to add to ocean cluster
    * `arn` - (Optional) Required if type is set to TARGET_GROUP
    * `name` - (Optional) Required if type is set to CLASSIC
    * `type` - (Required) Can be set to CLASSIC or TARGET_GROUP
* `tags` - (Optional) Optionally adds tags to instances launched in an Ocean cluster.
* `key` - (Optional) The tag key.
* `value` - (Optional) The tag value.
* `fallback_to_ondemand` - (Optional, Default: `true`) If not Spot instance markets are available, enable Ocean to launch On-Demand instances instead.
* `utilize_reserved_instances` - (Optional, Default `true`) If Reserved instances exist, Ocean will utilize them before launching Spot instances.
* `draining_timeout` - (Optional) The time in seconds, the instance is allowed to run while detached from the ELB. This is to allow the instance time to be drained from incoming TCP connections before terminating it, during a scale down operation.
* `grace_period` - (Optional, Default: 600) The amount of time, in seconds, after the instance has launched to start checking its health.
* `spot_percentage` - (Optional; Required if not using `ondemand_count`) The percentage of Spot instances that would spin up from the `desired_capacity` number.


<a id="auto-scaler"></a>
## Auto Scaler
* `autoscaler` - (Optional) Describes the Ocean Kubernetes autoscaler.
* `autoscale_is_enabled` - (Optional, Default: `true`) Enable the Ocean Kubernetes autoscaler.
* `autoscale_is_auto_config` - (Optional, Default: `true`) Automatically configure and optimize headroom resources.
* `autoscale_cooldown` - (Optional, Default: `null`) Cooldown period between scaling actions.
* `auto_headroom_percentage` - (Optional) Set the auto headroom percentage (a number in the range [0, 200]) which controls the percentage of headroom from the cluster. Relevant only when `autoscale_is_auto_config` toggled on.
* `autoscale_headroom` - (Optional) Spare resource capacity management enabling fast assignment of Pods without waiting for new resources to launch.
* `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
* `gpu_per_unit` - (Optional) Optionally configure the number of GPUS to allocate the headroom.
* `memory_per_unit` - (Optional) Optionally configure the amount of memory (MB) to allocate the headroom.
* `num_of_units` - (Optional) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.
* `autoscale_down` - (Optional) Auto Scaling scale down operations.
* `max_scale_down_percentage` - (Optional) Would represent the maximum % to scale-down. Number between 1-100.
* `resource_limits` - (Optional) Optionally set upper and lower bounds on the resource usage of the cluster.
* `max_vcpu` - (Optional) The maximum cpu in vCPU units that can be allocated to the cluster.
* `max_memory_gib` - (Optional) The maximum memory in GiB units that can be allocated to the cluster.

```hcl
  autoscaler {
    autoscale_is_enabled     = true
    autoscale_is_auto_config = true
    auto_headroom_percentage = 100
    autoscale_cooldown       = 300

    autoscale_headroom {
      cpu_per_unit    = 1024
      gpu_per_unit    = 0
      memory_per_unit = 512
      num_of_units    = 2
    }

    autoscale_down {
      max_scale_down_percentage = 60
    }

    resource_limits {
      max_vcpu       = 1024
      max_memory_gib = 1500
    }
  }
```

<a id="update-policy"></a>
## Update Policy

* `update_policy` - (Optional)
    * `should_roll` - (Required) Enables the roll.
    * `roll_config` - (Required) While used, you can control whether the group should perform a deployment after an update to the configuration.
        * `batch_size_percentage` - (Required) Sets the percentage of the instances to deploy in each batch.

```hcl
  update_policy {
    should_roll = false
    
    roll_config {
      batch_size_percentage = 33
    }
  }
```

<a id="scheduled-task"></a>
## scheduled task
* `scheduled_task` - (Optional) Set scheduling object.
    * `shutdown_hours` - (Optional) Set shutdown hours for cluster object.
        * `is_enabled` - (Optional)  Flag to enable / disable the shutdown hours.
                                     Example: True
        * `time_windows` - (Required) Set time windows for shutdown hours. specify a list of 'timeWindows' with at least one time window Each string is in the format of - ddd:hh:mm-ddd:hh:mm ddd = day of week = Sun | Mon | Tue | Wed | Thu | Fri | Sat hh = hour 24 = 0 -23 mm = minute = 0 - 59. Time windows should not overlap. required on cluster.scheduling.isEnabled = True. API Times are in UTC
                                      Example: Fri:15:30-Wed:14:30
    * `tasks` - (Optional) The scheduling tasks for the cluster.
        * `is_enabled` - (Required)  Describes whether the task is enabled. When true the task should run when false it should not run. Required for cluster.scheduling.tasks object.
        * `cron_expression` - (Required) A valid cron expression. For example : " * * * * * ".The cron is running in UTC time zone and is in Unix cron format Cron Expression Validator Script. Only one of ‘frequency’ or ‘cronExpression’ should be used at a time. Required for cluster.scheduling.tasks object
                                         Example: 0 1 * * *
        * `task_type` - (Required) Valid values: "clusterRoll". Required for cluster.scheduling.tasks object
                                   Example: clusterRoll
             
```hcl
  scheduled_task  {
    shutdown_hours  {
      is_enabled = true
      time_windows = ["Fri:15:30-Sat:13:30","Sun:15:30-Mon:13:30"]
    }
    tasks  {
      is_enabled = false
      cron_expression = "* * * * *"
      task_type = "clusterRoll"
    }
  }
```

## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `id` - The Spotinst Ocean ID.