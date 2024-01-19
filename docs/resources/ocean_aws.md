---
layout: "spotinst"
page_title: "Spotinst: ocean_aws"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean resource using AWS.
---

# spotinst\_ocean\_aws

Manages a Spotinst Ocean AWS resource.

## Prerequisites

Installation of the Ocean controller is required by this resource. You can accomplish this by using the [spotinst/ocean-controller](https://registry.terraform.io/modules/spotinst/ocean-controller/spotinst) module as follows:

```hcl
module "ocean-controller" {
  source = "spotinst/ocean-controller/spotinst"

  # Credentials.
  spotinst_token   = "redacted"
  spotinst_account = "redacted"

  # Configuration.
  cluster_identifier = "ocean-dev"
}
```

~> You must configure the same `cluster_identifier` both for the Ocean controller and for the `spotinst_ocean_aws` resource.

To learn more about how to integrate existing Kubernetes clusters into Ocean using Terraform, watch [this video](https://youtu.be/ffGmMlpPsPE).

## Example Usage

```hcl
resource "spotinst_ocean_aws" "example" {
  name          = "demo"
  controller_id = "ocean-dev"
  region        = "us-west-2"

  max_size         = 2
  min_size         = 1
  desired_capacity = 2

  subnet_ids = ["subnet-123456789"]

  // region INSTANCE-TYPES
  
  //whitelist  = ["t1.micro", "m1.small"]
  //blacklist = ["t1.micro", "m1.small"]
  filters {
      architectures             =   ["x86_64", "i386"]
      categories                =   ["Accelerated_computing", "Compute_optimized"]
      disk_types                =   ["EBS", "SSD"]
      exclude_families          =   ["m*"]
      exclude_metal             =   false
      hypervisor                =   ["xen"]
      include_families          =   ["c*", "t*"]
      is_ena_supported          =   false
      max_gpu                   =   4
      min_gpu                   =   0
      max_memory_gib            =   16
      max_network_performance   =   20
      max_vcpu                  =   16
      min_enis                  =   2
      min_memory_gib            =   8
      min_network_performance   =   2
      min_vcpu                  =   2
      root_device_types         =   ["ebs"]
      virtualization_types      =   ["hvm"] 
    }
  }
  
  

  // region LAUNCH CONFIGURATION
  image_id                    = "ami-123456"
  security_groups             = ["sg-987654321"]
  key_name                    = "fake key"
  user_data                   = "echo hello world"
  iam_instance_profile        = "iam-profile"
  root_volume_size            = 20
  monitoring                  = true
  ebs_optimized               = true
  associate_public_ip_address = true
  associate_ipv6_address      = true
  use_as_template_only        = true

  load_balancers {
    arn  = "arn:aws:elasticloadbalancing:us-west-2:fake-arn"
    type = "TARGET_GROUP"
  }
  load_balancers {
    name = "example"
    type = "CLASSIC"
  }
  
  resource_tag_specification {
    should_tag_volumes = true
  }
  // endregion

  // region STRATEGY 
  fallback_to_ondemand       = true
  draining_timeout           = 120
  utilize_reserved_instances = false
  grace_period               = 300
  spot_percentage            = 100
  utilize_commitments        = false
  spread_nodes_by            = "count"
  cluster_orientation{
    availability_vs_cost="balanced"
  }
  // endregion

  tags {
    key   = "fakeKey"
    value = "fakeValue"
  }
  
  instance_metadata_options {
    http_tokens = "required"
    http_put_response_hop_limit = 10
  }
  
  block_device_mappings {
    device_name = "/dev/xvda"
    ebs {
      delete_on_termination = "true"
      encrypted             = "false"
      volume_type           = "gp3"
      volume_size           = 50
      throughput            = 500
      dynamic_volume_size {
        base_size              = 50
        resource               = "CPU"
        size_per_resource_unit = 20
      }
      dynamic_iops {
        base_size              = 50
        resource               = "memory"
        size_per_resource_unit = 20
      }
    }
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
  value = spotinst_ocean_aws.example.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The cluster name.
* `controller_id` - (Required) A unique identifier used for connecting the Ocean SaaS platform and the Kubernetes cluster. Typically, the cluster name is used as its identifier.
* `region` - (Required) The region the cluster will run in.
* `max_size` - (Optional, Default: `1000`) The upper limit of instances the cluster can scale up to.
* `min_size` - (Optional) The lower limit of instances the cluster can scale down to.
* `desired_capacity` - (Optional) The number of instances to launch and maintain in the cluster.
* `subnet_ids` - (Required) A comma-separated list of subnet identifiers for the Ocean cluster. Subnet IDs should be configured with auto assign public IP.
* `instanceTypes` - (Optional) The type of instances that may or may not be a part of the Ocean cluster.
  * `whitelist` - (Optional) Instance types allowed in the Ocean cluster. Cannot be configured if `blacklist` is configured.
  * `blacklist` - (Optional) Instance types not allowed in the Ocean cluster. Cannot be configured if `whitelist` is configured.
  * `filters` - (Optional) List of filters. The Instance types that match with all filters compose the Ocean's whitelist parameter. Cannot be configured together with whitelist/blacklist.
    * `architectures` - (Optional) The filtered instance types will support at least one of the architectures from this list.
    * `categories` - (Optional) The filtered instance types will belong to one of the categories types from this list.
    * `disk_types` - (Optional) The filtered instance types will have one of the disk type from this list.
    * `exclude_families` - (Optional) Types belonging to a family from the ExcludeFamilies will not be available for scaling (asterisk wildcard is also supported). For example, C* will exclude instance types from these families: c5, c4, c4a, etc.
    * `exclude_metal` - (Optional, Default: false) In case excludeMetal is set to true, metal types will not be available for scaling.
    * `hypervisor` - (Optional) The filtered instance types will have a hypervisor type from this list.
    * `include_families` - (Optional) Types belonging to a family from the IncludeFamilies will be available for scaling (asterisk wildcard is also supported). For example, C* will include instance types from these families: c5, c4, c4a, etc.
    * `is_ena_supported` - (Optional) Ena is supported or not.
    * `max_gpu` - (Optional) Maximum total number of GPUs.
    * `max_memory_gib` - (Optional) Maximum amount of Memory (GiB).
    * `max_network_performance` - (Optional) Maximum Bandwidth in Gib/s of network performance.
    * `max_vcpu` - (Optional) Maximum number of vcpus available.
    * `min_enis` - (Optional) Minimum number of network interfaces (ENIs).
    * `min_gpu` - (Optional) Minimum total number of GPUs.
    * `min_memory_gib` - (Optional) Minimum amount of Memory (GiB).
    * `min_network_performance` - (Optional) Minimum Bandwidth in Gib/s of network performance.
    * `min_vcpu` - (Optional) Minimum number of vcpus available.
    * `root_device_types` - (Optional) The filtered instance types will have a root device types from this list.
    * `virtualization_types` - (Optional) The filtered instance types will support at least one of the virtualization types from this list.
* `user_data` - (Optional) Base64-encoded MIME user data to make available to the instances.
* `image_id` - (Required) ID of the image used to launch the instances.
* `security_groups` - (Required) One or more security group ids.
* `key_name` - (Optional) The key pair to attach the instances.
* `iam_instance_profile` - (Optional) The instance profile iam role.
* `associate_public_ip_address` - (Optional, Default: `false`) Configure public IP address allocation.
* `associate_ipv6_address` - (Optional, Default: `false`) Configure IPv6 address allocation.
* `root_volume_size` - (Optional) The size (in Gb) to allocate for the root volume. Minimum `20`.
* `monitoring` - (Optional) Enable detailed monitoring for cluster. Flag will enable Cloud Watch detailed monitoring (one minute increments). Note: there are additional hourly costs for this service based on the region used.
* `ebs_optimized` - (Optional) Enable EBS optimized for cluster. Flag will enable optimized capacity for high bandwidth connectivity to the EB service for non EBS optimized instance types. For instances that are EBS optimized this flag will be ignored.
* `use_as_template_only` - (Optional, Default: false) launch specification defined on the Ocean object will function only as a template for virtual node groups.
  When set to true, on Ocean resource creation please make sure your custom VNG has an initial_nodes parameter to create nodes for your VNG.
* `load_balancers` - (Optional) - Array of load balancer objects to add to ocean cluster
    * `arn` - (Optional) Required if type is set to `TARGET_GROUP`
    * `name` - (Optional) Required if type is set to `CLASSIC`
    * `type` - (Required) Can be set to `CLASSIC` or `TARGET_GROUP`
* `tags` - (Optional) Optionally adds tags to instances launched in an Ocean cluster.
    * `key` - (Optional) The tag key.
    * `value` - (Optional) The tag value.
* `fallback_to_ondemand` - (Optional, Default: `true`) If not Spot instance markets are available, enable Ocean to launch On-Demand instances instead.
* `utilize_reserved_instances` - (Optional, Default `true`) If Reserved instances exist, Ocean will utilize them before launching Spot instances.
* `draining_timeout` - (Optional) The time in seconds, the instance is allowed to run while detached from the ELB. This is to allow the instance time to be drained from incoming TCP connections before terminating it, during a scale down operation.
* `grace_period` - (Optional, Default: 300) The amount of time, in seconds, after the instance has launched to start checking its health.
* `spot_percentage` - (Optional) The desired percentage of Spot instances out of all running instances. Only available when the field is not set in any VNG directly (launchSpec.strategy.spotPercentage).
* `utilize_commitments` - (Optional, Default false) If savings plans exist, Ocean will utilize them before launching Spot instances.
* `spread_nodes_by` - (Optional, Default: `count`) Ocean will spread the nodes across markets by this value. Possible values: `vcpu` or `count`.
* `instance_metadata_options` - (Optional) Ocean instance metadata options object for IMDSv2.
    * `http_tokens` - (Required) Determines if a signed token is required or not. Valid values: `optional` or `required`.
    * `http_put_response_hop_limit` - (Optional) An integer from 1 through 64. The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further the instance metadata requests can travel.
* `block_device_mappings` - (Optional) Object. Array list of block devices that are exposed to the instance, specify either virtual devices and EBS volumes.
    * `device_name` - (Optional) String. Set device name. (Example: `/dev/xvda`).
    * `ebs`- (Optional) Object. Set Elastic Block Store properties .
        * `delete_on_termination` - (Optional) Boolean. Flag to delete the EBS on instance termination.
        * `encrypted` - (Optional) Boolean. Enables [EBS encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) on the volume.
        * `iops` - (Required for requests to create io1 volumes; it is not used in requests to create `gp2`, `st1`, `sc1`, or standard volumes) Int. The number of I/O operations per second (IOPS) that the volume supports.
        * `kms_key_id` - (Optional) String. Identifier (key ID, key alias, ID ARN, or alias ARN) for a customer managed CMK under which the EBS volume is encrypted.
        * `snapshot_id` - (Optional) (Optional) String. The Snapshot ID to mount by.
        * `volume_type` - (Optional, Default: `"standard"`) String. The type of the volume. (Example: `gp2`).
        * `volume_size` - (Optional) Int. The size, in GB of the volume.
        * `throughput` - (Optional) The amount of data transferred to or from a storage device per second, you can use this param just in a case that `volume_type` = `gp3`.
        * `dynamic_volume_size` - (Optional) Object. Set dynamic volume size properties. When using this object, you cannot use volumeSize. You must use one or the other.
            * `base_size`- (Required) Int. Initial size for volume. (Example: 50)
            * `resource`- (Required) String. Resource type to increase volume size dynamically by. (Valid values: `CPU`)
            * `size_per_resource_unit`- (Required) Int. Additional size (in GB) per resource unit. (Example: `baseSize=50`, `sizePerResourceUnit=20`, and instance with 2 CPU is launched; its total disk size will be: 90GB).
        * `iops` - (Optional) Must be greater than or equal to 0.
        * `dynamic_iops` - (Optional) Set dynamic IOPS properties. When using this object, you cannot use the `iops` attribute. You must use one or the other.
            * `base_size`- (Required) Initial size for IOPS.
            * `resource`- (Required, ENUM: `CPU`, `memory`)
            * `size_per_resource_unit`- (Required) Additional size per resource unit (in IOPS). (Example: `baseSize=50`, `sizePerResourceUnit=20`, and an instance with 2 CPU is launched; its IOPS size will be: 90).
* `cluster_orientation`
    * `availability_vs_cost` - (Optional, Default: `balanced`) You can control the approach that Ocean takes while launching nodes by configuring this value. Possible values: `costOriented`,`balanced`,`cheapest`.
* `logging` - (Optional) Logging configuration.
    * `export` - (Optional) Logging Export configuration.
        * `s3` - (Optional) Exports your cluster's logs to the S3 bucket and subdir configured on the S3 data integration given.
            * `id` - (Required) The identifier of The S3 data integration to export the logs to.
* `resource_tag_specification` - (Optional) Specify which resources should be tagged with Virtual Node Group tags or Ocean tags. If tags are set on the VNG, the resources will be tagged with the VNG tags; otherwise, they will be tagged with the Ocean tags.
    * `should_tag_volumes` - (Optional) Specify if Volume resources will be tagged with Virtual Node Group tags or Ocean tags.
## Auto Scaler
* `autoscaler` - (Optional) Describes the Ocean Kubernetes Auto Scaler.
    * `autoscale_is_enabled` - (Optional, Default: `true`) Enable the Ocean Kubernetes Auto Scaler.
    * `autoscale_is_auto_config` - (Optional, Default: `true`) Automatically configure and optimize headroom resources.
    * `autoscale_cooldown` - (Optional, Default: `null`) Cooldown period between scaling actions.
    * `auto_headroom_percentage` - (Optional) Set the auto headroom percentage (a number in the range [0, 200]) which controls the percentage of headroom from the cluster. Relevant only when `autoscale_is_auto_config` toggled on.
    * `enable_automatic_and_manual_headroom` - (Optional, Default: `false`) enables automatic and manual headroom to work in parallel. When set to false, automatic headroom overrides all other headroom definitions manually configured, whether they are at cluster or VNG level.
    * `autoscale_headroom` - (Optional) Spare resource capacity management enabling fast assignment of Pods without waiting for new resources to launch.
        * `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
        * `gpu_per_unit` - (Optional) Optionally configure the number of GPUs to allocate the headroom.
        * `memory_per_unit` - (Optional) Optionally configure the amount of memory (MB) to allocate the headroom.
        * `num_of_units` - (Optional) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.
    * `autoscale_down` - (Optional) Auto Scaling scale down operations.
        * `max_scale_down_percentage` - (Optional) Would represent the maximum % to scale-down. Number between 1-100.
    * `resource_limits` - (Optional) Optionally set upper and lower bounds on the resource usage of the cluster.
        * `max_vcpu` - (Optional) The maximum cpu in vCPU units that can be allocated to the cluster.
        * `max_memory_gib` - (Optional) The maximum memory in GiB units that can be allocated to the cluster.
    * `extended_resource_definitions` - (Optional) List of Ocean extended resource definitions to use in this cluster.

```hcl
autoscaler {
  autoscale_is_enabled     = true
  autoscale_is_auto_config = true
  auto_headroom_percentage = 100
  autoscale_cooldown       = 300
  enable_automatic_and_manual_headroom = false

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
  
  extended_resource_definitions = ["erd-abc123"]
}
```

<a id="update-policy"></a>
## Update Policy

* `update_policy` - (Optional)
    * `should_roll` - (Required) Enables the roll.
    * `conditioned_roll` - (Optional, Default: false) Spot will perform a cluster Roll in accordance with a relevant modification of the cluster’s settings. When set to true , only specific changes in the cluster’s configuration will trigger a cluster roll (such as AMI, Key Pair, user data, instance types, load balancers, etc).
    * `auto_apply_tags` - (Optional, Default: false) will update instance tags on the fly without rolling the cluster.
    * `roll_config` - (Required) While used, you can control whether the group should perform a deployment after an update to the configuration.
        * `batch_size_percentage` - (Required) Sets the percentage of the instances to deploy in each batch.
        * `launch_spec_ids` - (Optional) List of virtual node group identifiers to be rolled.
        * `batch_min_healthy_percentage` - (Optional) Default: 50. Indicates the threshold of minimum healthy instances in single batch. If the amount of healthy instances in single batch is under the threshold, the cluster roll will fail. If exists, the parameter value will be in range of 1-100. In case of null as value, the default value in the backend will be 50%. Value of param should represent the number in percentage (%) of the batch.
        * `respect_pdb` - (Optional, Default: false) During the roll, if the parameter is set to True we honor PDB during the instance replacement.
```hcl
update_policy {
  should_roll = false
  conditioned_roll = true
  auto_apply_tags = true

  roll_config {
    batch_size_percentage = 33
    launch_spec_ids = ["ols-1a2b3c4d"]
    batch_min_healthy_percentage = 20
    respect_pdb = true
  }
}
```

<a id="scheduled-task"></a>
## Scheduled Task
* `scheduled_task` - (Optional) Set scheduling object.
    * `shutdown_hours` - (Optional) Set shutdown hours for cluster object.
        * `is_enabled` - (Optional) Toggle the shutdown hours task. (Example: `true`).
        * `time_windows` - (Required) Set time windows for shutdown hours. Specify a list of `timeWindows` with at least one time window Each string is in the format of: `ddd:hh:mm-ddd:hh:mm` where `ddd` = day of week = Sun | Mon | Tue | Wed | Thu | Fri | Sat, `hh` = hour 24 = 0 -23, `mm` = minute = 0 - 59. Time windows should not overlap. Required if `cluster.scheduling.isEnabled` is `true`. (Example: `Fri:15:30-Wed:14:30`).
    * `tasks` - (Optional) The scheduling tasks for the cluster.
        * `is_enabled` - (Required)  Describes whether the task is enabled. When true the task should run when false it should not run. Required for `cluster.scheduling.tasks` object.
        * `cron_expression` - (Required) A valid cron expression. The cron is running in UTC time zone and is in Unix cron format Cron Expression Validator Script. Only one of `frequency` or `cronExpression` should be used at a time. Required for `cluster.scheduling.tasks` object. (Example: `0 1 * * *`).
        * `task_type` - (Required) Valid values: `clusterRoll`. Required for `cluster.scheduling.tasks` object. (Example: `clusterRoll`).
        * `parameters` - (Required) This filed will be compatible to the `task_type` field. If `task_type` is defined as `clusterRoll`, user cluster roll object in parameters.
            * `amiAutoUpdate` - (Optional) Set amiAutoUpdate object
                * `applyRoll` - (Optional) When the AMI is updated according to the configuration set, a cluster roll can be triggered
                * `clusterRoll` - (Optional) Set clusterRoll object
                    * `batchMinHealthyPercentage` - (Optional) Indicates the threshold of minimum healthy instances in single batch. If the amount of healthy instances in single batch is under the threshold, the cluster roll will fail. If exists, the parameter value will be in range of 1-100. In case of null as value, the default value in the backend will be 50%. Value of param should represent the number in percentage (%) of the batch.
                    * `batchSizePercentage` - (Optional) Value as a percent to set the size of a batch in a roll. Valid values are 0-100.
                    * `comment` - (Optional) Add a `comment` description for the roll. The `comment` is limited to 256 chars
                    * `respectPdb` - (Optional) During the roll, if the parameter is set to true we honor PDB during the instance replacement.
                * `minorVersion` - (Optional) When set to 'true', the auto-update process will update the VNGs’ AMI with the AMI to match the Kubernetes control plane version. either "patch" or "minorVersion" must be true.
                * `patch` - (Optional) When set to 'true', the auto-update process will update the VNGs’ images with the latest security patches. either "patch" or "minorVersion" must be true.
            * 
            * `clusterRoll` - (Optional) Set clusterRoll object
                * `batchMinHealthyPercentage` - (Optional) Indicates the threshold of minimum healthy instances in single batch. If the amount of healthy instances in single batch is under the threshold, the cluster roll will fail. If exists, the parameter value will be in range of 1-100. In case of null as value, the default value in the backend will be 50%. Value of param should represent the number in percentage (%) of the batch.
                * `batchSizePercentage` - (Optional) Value as a percent to set the size of a batch in a roll. Valid values are 0-100.
                * `comment` - (Optional) Add a `comment` description for the roll. The `comment` is limited to 256 chars
                * `respectPdb` - (Optional) During the roll, if the parameter is set to true we honor PDB during the instance replacement.

```hcl
scheduled_task {
  shutdown_hours {
    is_enabled   = true
    time_windows = [
      "Fri:15:30-Sat:13:30", 
      "Sun:15:30-Mon:13:30",
    ]
  }
  tasks {
    is_enabled      = false
    cron_expression = "* * * * *"
    task_type       = "amiAutoUpdate"
     parameters  {
        ami_auto_update  {
            apply_roll = false
            ami_auto_update_cluster_roll  {
                batch_min_healthy_percentage = 100
                batch_size_percentage = 20
                comment = "test comment"
                respect_pdb = true
            }
            minor_version = true
            patch = false
        }
     }
  }
}
```

<a id="attributes-reference"></a>
## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `id` - The Cluster ID.


<a id="import"></a>
## Import

Clusters can be imported using the Ocean `id`, e.g.,
```hcl
$ terraform import spotinst_ocean_aws.this o-12345678
```
