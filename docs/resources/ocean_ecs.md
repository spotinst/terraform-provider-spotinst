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

    security_group_ids = ["sg-12345"]
    image_id = "ami-12345"
    iam_instance_profile = "iam-profile"
  
    key_pair = "KeyPair"
    user_data = "echo hello world"
    associate_public_ip_address = false
    utilize_reserved_instances  = false
    draining_timeout            = 120
    monitoring                  = true
    ebs_optimized               = true
    use_as_template_only        = true

    spot_percentage     = 100
    utilize_commitments = false

  instance_metadata_options {
    http_tokens                 = "required"
    http_put_response_hop_limit = 10
  }

  block_device_mappings {
      device_name = "/dev/xvda1"
      ebs {
        delete_on_termination = "true"
        encrypted = "false"
        volume_type = "gp2"
        volume_size = 50
        throughput = 500
        dynamic_volume_size {
          base_size = 50
          resource = "CPU"
          size_per_resource_unit = 20
        }
      }
   }

  optimize_images {
    perform_at              = "timeWindow"
    time_windows            = ["Sun:02:00-Sun:12:00","Sun:05:00-Sun:16:00"]
    should_optimize_ecs_ami = true
  }

   tags {
     key   = "fakeKey"
     value = "fakeValue"
    }
    
   logging {
     export {
       s3 {
         id = "di-abcd123"
          }
        }
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
* `whitelist` - (Optional) Instance types allowed in the Ocean cluster.
* `user_data` - (Optional) Base64-encoded MIME user data to make available to the instances.
* `image_id` - (Required) ID of the image used to launch the instances.
* `security_group_ids` - (Required) One or more security group ids.
* `key_pair` - (Optional) The key pair to attach the instances.
* `iam_instance_profile` - (Optional) The instance profile iam role.
* `associate_public_ip_address` - (Optional, Default: `false`) Configure public IP address allocation.
* `utilize_reserved_instances` - (Optional, Default `true`) If Reserved instances exist, Ocean will utilize them before launching Spot instances.
* `draining_timeout` - (Optional) The time in seconds, the instance is allowed to run while detached from the ELB. This is to allow the instance time to be drained from incoming TCP connections before terminating it, during a scale down operation.
* `monitoring` - (Optional) Enable detailed monitoring for cluster. Flag will enable Cloud Watch detailed monitoring (one minute increments). Note: there are additional hourly costs for this service based on the region used.
* `ebs_optimized` - (Optional) Enable EBS optimized for cluster. Flag will enable optimized capacity for high bandwidth connectivity to the EB service for non EBS optimized instance types. For instances that are EBS optimized this flag will be ignored.
* `use_as_template_only` - (Optional, Default: false) launch specification defined on the Ocean object will function only as a template for virtual node groups.
* `spot_percentage` - (Optional) The percentage of Spot instances that would spin up from the `desired_capacity` number.
* `utilize_commitments` - (Optional, Default false) If savings plans exist, Ocean will utilize them before launching Spot instances.
* `instance_metadata_options` - (Optional) Ocean instance metadata options object for IMDSv2.
    * `http_tokens` - (Required) Determines if a signed token is required or not. Valid values: `optional` or `required`.
    * `http_put_response_hop_limit` - (Optional) An integer from 1 through 64. The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further the instance metadata requests can travel.
* `logging` - (Optional) Logging configuration.
    * `export` - (Optional) Logging Export configuration.
        * `s3` - (Optional) Exports your cluster's logs to the S3 bucket and subdir configured on the S3 data integration given.
            * `id` - (Required) The identifier of The S3 data integration to export the logs to.


<a id="block-devices"></a>
## Block Devices
* `block_device_mappings` - (Optional) Object. List of block devices that are exposed to the instance, specify either virtual devices and EBS volumes.   
    * `device_name` - (Optional) String. Set device name. Example: `/dev/xvda1`.
    * `ebs` - (Optional) Object. Set Elastic Block Store properties.
        * `delete_on_termination` - (Optional) Boolean. Toggles EBS deletion upon instance termination. 
        * `encrypted` - (Optional) Boolean. Enables [EBS encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) on the volume.
        * `iops` - (Required for requests to create `io1` volumes; it is not used in requests to create `gp2`, `st1`, `sc1`, or standard volumes) Int. The number of I/O operations per second (IOPS) that the volume supports.
        * `kms_key_id` - (Optional) String. Identifier (key ID, key alias, ID ARN, or alias ARN) for a customer managed CMK under which the EBS volume is encrypted.
        * `snapshot_id` - (Optional) (Optional) String. The snapshot ID to mount by. 
        * `volume_type` - (Optional, Default: `standard`) String. The type of the volume. Example: `gp2`.
        * `volume_size` - (Optional) Int. The size (in GB) of the volume.
        * `throughput`- (Optional) The amount of data transferred to or from a storage device per second, you can use this param just in a case that `volume_type` = gp3.
        * `dynamic_volume_size` - (Optional) Object. Set dynamic volume size properties. When using this object, you cannot use volumeSize. You must use one or the other.
            * `base_size` - (Required) Int. Initial size for volume. Example: `50`.
            * `resource` - (Required) String. Resource type to increase volume size dynamically by. Valid values: `CPU`.
            * `size_per_resource_unit` - (Required) Int. Additional size (in GB) per resource unit. Example: When the `baseSize=50`, `sizePerResourceUnit=20`, and instance with two CPUs is launched, its total disk size will be: 90GB.
    * `no_device` - (Optional) String. Suppresses the specified device included in the block device mapping of the AMI.
* `optimize_images` - (Optional) Object. Set auto image update settings.
    * `perform_at` - (Required) String. Valid values: "always" "never" "timeWindow".
    * `time_windows` - (Optional; Required if not using `perform_at` = timeWindow) Array of strings. Set time windows for image update, at least one time window. Each string is in the format of ddd:hh:mm-ddd:hh:mm ddd. Time windows should not overlap.
    * `should_optimize_ecs_ami` - (Required) Boolean. Enable auto image (AMI) update for the ECS container instances. The auto update applies for ECS-Optimized AMIs.
    

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
        * `max_scale_down_percentage` - (Optional) Would represent the maximum % to scale-down. Number between 1-100.
    * `resource_limits` - (Optional) Optionally set upper and lower bounds on the resource usage of the cluster.
        * `max_vcpu` - (Optional) The maximum cpu in vCPU units that can be allocated to the cluster.
        * `max_memory_gib` - (Optional) The maximum memory in GiB units that can be allocated to the cluster.
    * `auto_headroom_percentage` - (Optional) The auto-headroom percentage. Set a number between 0-200 to control the headroom % of the cluster. Relevant when `isAutoConfig`= true.

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
   
    auto_headroom_percentage = 10
  }
```


<a id="update-policy"></a>
## Update Policy
* `update_policy` - (Optional) While used, you can control whether the group should perform a deployment after an update to the configuration.
    * `should_roll` - (Required) Enables the roll.
    * `conditioned_roll` - (Optional, Default: false) Spot will perform a cluster Roll in accordance with a relevant modification of the cluster’s settings. When set to true , only specific changes in the cluster’s configuration will trigger a cluster roll (such as AMI, Key Pair, user data, instance types, load balancers, etc).
    * `auto_apply_tags` - (Optional, Default: false) will update instance tags on the fly without rolling the cluster.
    * `roll_config` - (Required) 
        * `batch_size_percentage` - (Required) Sets the percentage of the instances to deploy in each batch.
        * `batch_min_healthy_percentage` - (Optional) Default: 50. Indicates the threshold of minimum healthy instances in single batch. If the amount of healthy instances in single batch is under the threshold, the cluster roll will fail. If exists, the parameter value will be in range of 1-100. In case of null as value, the default value in the backend will be 50%. Value of param should represent the number in percentage (%) of the batch.

```hcl
  update_policy {
    should_roll = false
    conditioned_roll = true
    auto_apply_tags = true
    
    roll_config {
      batch_size_percentage = 33
      batch_min_healthy_percentage = 20
    }
  }
```


<a id="scheduled-tasks"></a>
## Scheduled Tasks
* `scheduled_task` - (Optional) While used, you can control whether the group should perform a deployment after an update to the configuration.
    * `shutdown_hours` - (Optional) Set shutdown hours for cluster object.
        * `is_enabled` - (Optional)  Flag to enable / disable the shutdown hours.
        * `time_windows` - (Required) Set time windows for shutdown hours. Specify a list of `timeWindows` with at least one time window Each string is in the format of `ddd:hh:mm-ddd:hh:mm` (ddd = day of week = Sun | Mon | Tue | Wed | Thu | Fri | Sat hh = hour 24 = 0 -23 mm = minute = 0 - 59). Time windows should not overlap. Required when `cluster.scheduling.isEnabled` is true. API Times are in UTC. Example: `Fri:15:30-Wed:14:30`.
    * `tasks` - (Optional) The scheduling tasks for the cluster.
        * `is_enabled` - (Required) Describes whether the task is enabled. When true the task should run when false it should not run. Required for `cluster.scheduling.tasks` object.
        * `cron_expression` - (Required) A valid cron expression. The cron is running in UTC time zone and is in Unix cron format Cron Expression Validator Script. Only one of `frequency` or `cronExpression` should be used at a time. Required for `cluster.scheduling.tasks` object. Example: `0 1 * * *`.
        * `task_type` - (Required) Valid values: "clusterRoll". Required for `cluster.scheduling.tasks object`. Example: `clusterRoll`.
             
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
