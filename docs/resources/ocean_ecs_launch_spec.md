---
layout: "spotinst"
page_title: "Spotinst: ocean_ecs_launch_spec"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean ECS Launch Spec resource using AWS.
---

# spotinst\_ocean\_ecs\_launch\_spec

Manages a custom Spotinst Ocean ECS Launch Spec resource.

## Example Usage

```hcl
resource "spotinst_ocean_ecs_launch_spec" "example" {
  ocean_id  = "o-123456"
  image_id  = "ami-123456"
  user_data = "echo hello world"
  iam_instance_profile = "iam-profile"
  subnet_ids = ["subnet-12345"]
  security_group_ids = ["awseb-12345"]
  restrict_scale_down = true
  instance_types = ["m3.large", "m3.xlarge", "m3.2xlarge", "m4.large", "m4.xlarge",
      "m4.4xlarge", "m4.2xlarge", "m4.10xlarge", "m4.16xlarge", "m5.large", 
      "m5.xlarge", "m5.2xlarge", "m5.4xlarge", "m5.12xlarge", "m5.24xlarge"
    ]
  preferred_spot_types = ["m3.large", "m3.xlarge", "m3.2xlarge", "m4.large", "m4.xlarge"]
  
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

  attributes {
    key   = "fakeKey"
    value = "fakeValue"
  }
  
  instance_metadata_options {
    http_tokens = "required"
    http_put_response_hop_limit = 10
  }
  
  autoscale_headrooms {
    num_of_units = 5
    cpu_per_unit = 1000
    memory_per_unit = 2048
  }
  
  strategy {
    spot_percentage = 50
  }

  tags {
     key   = "Env"
     value = "production"
  }
  
  //images object to support multi ami
  images {
    image_id = "ami-12345"
  }
  images {
    image_id = "ami-67890"
  }
  
    scheduling_task {
    is_enabled = true
    cron_expression = "0 1 * * *"
    task_type = "manualHeadroomUpdate"
    task_headroom {
        num_of_units    = 5
        cpu_per_unit     = 1000
        memory_per_unit = 2048
    }
  } 
}
```
```
output "ocean_launchspec_id" {
  value = spotinst_ocean_ecs_launch_spec.example.id
}
```

## Argument Reference

The following arguments are supported:

* `ocean_id`  - (Required) The Ocean cluster ID .
* `name`      - (Required) The Ocean Launch Specification name. 
* `user_data` - (Optional) Base64-encoded MIME user data to make available to the instances.
* `image_id`  - (Optional) ID of the image used to launch the instances.
* `images`    - (Optional) You can configure VNG with either the imageId or images objects, but not both simultaneously. For each architecture type (amd64, arm64) only one AMI is allowed. Valid values: null, or an array with at least one element
    * `image_id` - Identifier of the image in AWS. Valid values: any string which is not empty or null.
* `iam_instance_profile` - (Optional) The ARN or name of an IAM instance profile to associate with launched instances.
* `security_group_ids` - (Optional) One or more security group ids.
* `tags` - (Optional) A key/value mapping of tags to assign to the resource.
* `instance_types` - (Optional) A list of instance types allowed to be provisioned for pods pending under the specified launch specification. The list overrides the list defined for the Ocean cluster.
* `preferred_spot_types` - (Optional) When Ocean scales up instances, it takes your preferred types into consideration while maintaining a variety of machine types running for optimized distribution.
* `restrict_scale_down`- (Optional) Boolean. When set to `true`, VNG nodes will be treated as if all pods running have the restrict-scale-down label. Therefore, Ocean will not scale nodes down unless empty.
* `subnet_ids` - (Optional) Set subnets in launchSpec. Each element in the array should be a subnet ID.
* `instance_metadata_options` - (Optional) Ocean instance metadata options object for IMDSv2.
    * `http_tokens` - (Required) Determines if a signed token is required or not. Valid values: `optional` or `required`.
    * `http_put_response_hop_limit` - (Optional) An integer from 1 through 64. The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further the instance metadata requests can travel.

* `attributes` - (Optional) Optionally adds labels to instances launched in an Ocean cluster.
    * `key` - (Required) The label key.
    * `value` - (Required) The label value.

* `autoscale_headrooms` - (Optional) Set custom headroom per launch spec. provide list of headrooms object.
    * `num_of_units` - (Required) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.
    * `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in CPU units, where 1024 units = 1 vCPU.
    * `memory_per_unit` - (Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.

* `scheduling_task` - (Optional) Used to define scheduled tasks such as a manual headroom update.
    * `is_enabled` - (Required) Describes whether the task is enabled. When `true`, the task runs. When `false`, it does not run.
    * `cron_expression` - (Required) A valid cron expression. For example : " * * * * * ". The cron job runs in UTC time and is in Unix cron format.
    * `task_type` - (Required) The activity that you are scheduling. Valid values: "manualHeadroomUpdate".
    * `task_headroom` - (Optional) The config of this scheduled task. Depends on the value of taskType.
        * `num_of_units` - (Required) The number of units to retain as headroom, where each unit has the defined headroom CPU, memory and GPU.
        * `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
        * `memory_per_unit` - (Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.
*  `strategy` - (Optional) Similar to a strategy for an Ocean cluster, but applying only to a virtual node group.
    * `spot_percentage` - (Optional- if not using `spot_percentege` under `ocean strategy`) When set, Ocean will proactively try to maintain as close as possible to the percentage of Spot instances out of all the Virtual Node Group instances.

<a id="block-devices"></a>
## Block Devices
* `block_device_mappings`- (Optional) Object. Array list of block devices that are exposed to the instance, specify either virtual devices and EBS volumes.   
    * `device_name` - (Optional) String. Set device name. (Example: "/dev/xvda1").
    * `ebs`- (Optional) Object. Set Elastic Block Store properties .
        * `delete_on_termination`- (Optional) Boolean. Flag to delete the EBS on instance termination. 
        * `encrypted`- (Optional) Boolean. Enables [EBS encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) on the volume.
        * `iops`- (Required for requests to create io1 volumes; it is not used in requests to create gp2, st1, sc1, or standard volumes) Int. The number of I/O operations per second (IOPS) that the volume supports.
        * `kms_key_id`- (Optional) String. Identifier (key ID, key alias, ID ARN, or alias ARN) for a customer managed CMK under which the EBS volume is encrypted.
        * `snapshot_id`- (Optional) (Optional) String. The Snapshot ID to mount by. 
        * `volume_type`- (Optional, Default: `"standard"`) String. The type of the volume (example: "gp2").
        * `volume_size`- (Optional) Int. The size, in GB of the volume.
        * `throughput`- (Optional) The amount of data transferred to or from a storage device per second, you can use this param just in a case that `volume_type` = gp3.
        * `dynamic_volume_size`- (Optional) Object. Set dynamic volume size properties. When using this object, you cannot use volumeSize. You must use one or the other.
            * `base_size`- (Required) Int. Initial size for volume. (Example: 50)
            * `resource`- (Required) String. Resource type to increase volume size dynamically by. (valid values: "CPU")
            * `size_per_resource_unit`- (Required) Int. Additional size (in GB) per resource unit. (Example: baseSize= 50, sizePerResourceUnit=20, and instance with 2 CPU is launched - its total disk size will be: 90GB)
        * `no_device`- (Optional) String. suppresses the specified device included in the block device mapping of the AMI.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `id` - The Spotinst LaunchSpec ID.


<a id="import"></a>
## Import

Launch_Specs can be imported using the Launch_Spec `id`, e.g.,
```hcl
$ terraform import spotinst_ocean_ecs_launch_spec.nameOfTheResource ols-1a2345
```