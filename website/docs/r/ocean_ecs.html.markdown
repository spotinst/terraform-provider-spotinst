---
layout: "spotinst"
page_title: "Spotinst: ocean_ecs"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean ECS resource using AWS.
---

# spotinst\_ocean\_ecs

Manages a Spotinst Ocean ECS resource.

## Example Usage

```hcl
resource "spotinst_ocean_ecs" "example" {
    region = "us-west-2"
    name = "terraform-ecs-cluster"
    cluster_name = "terraform-ecs-cluster"
  
    min_size         = "0"
    max_size         = "1"
    desired_capacity = "0" 

    subnet_ids = ["subnet-12345"]
    whitelist = ["t3.medium"]
  // blacklist = ["t1.micro", "m1.small"]

    security_group_ids = ["sg-12345"]
    image_id = "ami-12345"
    iam_instance_profile = "iam-profile"
  
    key_pair = "KeyPair"
    user_data = "echo hello world"
    associate_public_ip_address = false
    utilize_reserved_instances = false
    draining_timeout            = 120
    monitoring                  = true
    ebs_optimized               = true

    tags {
      key   = "fakeKey"
      value = "fakeValue"
    }
}
```
```
output "ocean_id" {
  value = spotinst_ocean_ecs.example.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The Ocean cluster name.
* `cluster_name` - (Required) The ocean cluster name.
* `region` - (Required) The region the cluster will run in.
* `max_size` - (Optional, Default: `1000`) The upper limit of instances the cluster can scale up to.
* `min_size` - (Optional) The lower limit of instances the cluster can scale down to.
* `desired_capacity` - (Optional) The number of instances to launch and maintain in the cluster.
* `subnet_ids` - (Required) A comma-separated list of subnet identifiers for the Ocean cluster. Subnet IDs should be configured with auto assign public ip.
* `tags` - (Optional) Optionally adds tags to instances launched in an Ocean cluster.
    * `key` - (Optional) The tag key.
    * `value` - (Optional) The tag value.
* `whitelist` - (Optional) Instance types allowed in the Ocean cluster, Cannot be configured if blacklist is configured.
* `blacklist` - (Optional) Instance types to avoid launching in the Ocean cluster. Cannot be configured if whitelist is configured.
* `user_data` - (Optional) Base64-encoded MIME user data to make available to the instances.
* `image_id` - (Required) ID of the image used to launch the instances.
* `security_group_ids` - (Required) One or more security group ids.
* `key_pair` - (Optional) The key pair to attach the instances.
* `iam_instance_profile` - (Optional) The instance profile iam role.
* `associate_public_ip_address` - (Optional, Default: `false`) Configure public IP address allocation.
* `utilize_reserved_instances` - (Optional, Default `true`) If Reserved instances exist, OCean will utilize them before launching Spot instances.
* `draining_timeout` - (Optional) The time in seconds, the instance is allowed to run while detached from the ELB. This is to allow the instance time to be drained from incoming TCP connections before terminating it, during a scale down operation.
* `monitoring` - (Optional) Enable detailed monitoring for cluster. Flag will enable Cloud Watch detailed detailed monitoring (one minute increments). Note: there are additional hourly costs for this service based on the region used.
* `ebs_optimized` - (Optional) Enable EBS optimized for cluster. Flag will enable optimized capacity for high bandwidth connectivity to the EB service for non EBS optimized instance types. For instances that are EBS optimized this flag will be ignored.

<a id="auto-scaler"></a>
## Auto Scaler
* `autoscaler` - (Optional) Describes the Ocean ECS autoscaler.
    * `is_enabled` - (Optional, Default: `true`) Enable the Ocean ECS autoscaler.
    * `is_auto_config` - (Optional, Default: `true`) Automatically configure and optimize headroom resources.
    * `cooldown` - (Optional, Default: `null`) Cooldown period between scaling actions.
    * `headroom` - (Optional) Spare resource capacity management enabling fast assignment of tasks without waiting for new resources to launch.
        * `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
        * `memory_per_unit` - (Optional) Optionally configure the amount of memory (MB) to allocate the headroom.
        * `num_of_units` - (Optional) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.
    * `down` - (Optional) Auto Scaling scale down operations.
        * `max_scale_down_percentage` - (Optional) Would represent the maximum % to scale-down. Number between 1-100
    * `resource_limits` - (Optional) Optionally set upper and lower bounds on the resource usage of the cluster.
        * `max_vcpu` - (Optional) The maximum cpu in vCPU units that can be allocated to the cluster.
        * `max_memory_gib` - (Optional) The maximum memory in GiB units that can be allocated to the cluster.

```hcl
  autoscaler {
    is_enabled     = false
    is_auto_config = false
    cooldown       = 300

    headroom {
      cpu_per_unit    = 1024
      memory_per_unit = 512
      num_of_units    = 2
    }

    down {
      max_scale_down_percentage = 20
    }

    resource_limits {
      max_vcpu       = 1024
      max_memory_gib = 20
    }
  }
```


<a id="update-policy"></a>
## Update Policy
* `update_policy` - (Optional) While used, you can control whether the group should perform a deployment after an update to the configuration.
    * `should_roll` - (Required) Enables the roll.
    * `roll_config` - (Required) 
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
* `scheduled_task` - (Optional) While used, you can control whether the group should perform a deployment after an update to the configuration.
    * `shutdown_hours` - (Optional) Set shutdown hours for cluster object.
        * `is_enabled` - (Optional)  Flag to enable / disable the shutdown hours.
                                     Example: True
        * `time_windows` - (Required) Set time windows for shutdown hours. specify a list of 'timeWindows' with at least one time window Each string is in the format of - ddd:hh:mm-ddd:hh:mm ddd = day of week = Sun | Mon | Tue | Wed | Thu | Fri | Sat hh = hour 24 = 0 -23 mm = minute = 0 - 59. Time windows should not overlap. required on cluster.scheduling.isEnabled = True. API Times are in UTC
                                      Example: Fri:15:30-Wed:14:30
    * `tasks` - (Optional) The scheduling tasks for the cluster.
        * `is_enabled` - (Required)  Describes whether the task is enabled. When true the task should run when false it should not run. Required for cluster.scheduling.tasks object.
        * `cron_expression` - (Required) A valid cron expression. For example : " * * * * * ".The cron is running in UTC time zone and is in Unix cron format Cron Expression Validator Script. Only one of ‘frequency’ or ‘cronExpression’ should be used at a time. Required for cluster.scheduling.tasks object
                                         Example: 0 1 * * *.
        * `task_type` - (Required) Valid values: "clusterRoll". Required for cluster.scheduling.tasks object
                                   Example: clusterRoll.
             
```hcl
  scheduled_task  {
    shutdown_hours  {
      is_enabled = false
      time_windows = ["Fri:15:30-Wed:13:30"]
    }
    tasks {
      is_enabled = false
      cron_expression = "* * * * *"
      task_type = "clusterRoll"
    }
  }
```

## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `id` - The Spotinst Ocean ID.