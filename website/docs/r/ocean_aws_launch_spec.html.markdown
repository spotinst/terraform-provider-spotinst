---
layout: "spotinst"
page_title: "Spotinst: ocean_aws_launch_spec"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean Launch Spec resource using AWS.
---

# spotinst\_ocean\_aws\_launch\_spec

Provides a custom Spotinst Ocean AWS Launch Spec resource.

## Example Usage

```hcl
resource "spotinst_ocean_aws_launch_spec" "example" {
  ocean_id  = "o-123456"
  name = "launch spec name test"
  image_id  = "ami-123456"
  user_data = "echo hello world"
  iam_instance_profile = "iam-profile"
  security_groups = ["sg-987654321"]
  subnet_ids = ["subnet-1234"]
  intance_types = ["m3.large", "m3.xlarge", "m3.2xlarge", "m4.large", "m4.xlarge",
      "m4.4xlarge", "m4.2xlarge", "m4.10xlarge", "m4.16xlarge", "m5.large", 
      "m5.xlarge", "m5.2xlarge", "m5.4xlarge", "m5.12xlarge", "m5.24xlarge"
    ]
  root_volume_size = 30

  labels {
    key   = "fakeKey"
    value = "fakeValue"
  }
  
  taints {
    key    = "taint key updated"
    value  = "taint value updated"
    effect = "NoExecute"
  }
  
  autoscale_headrooms {
    num_of_units = 5
    cpu_per_unit = 1000
    gpu_per_unit = 0
    memory_per_unit = 2048
  }

  elastic_ip_pool  {
    tag_selector  {
      tag_key = "key"
      tag_value = "value"
    }
  }

  block_device_mappings {
      device_name = "/dev/xvda1"
      ebs {
        delete_on_termination = "true"
        encrypted = "false"
        volume_type = "gp2"
        volume_size = 50
        dynamic_volume_size {
          base_size = 50
          resource = "CPU"
          size_per_resource_unit = 20
        }
      }
   }

  resource_limits {
    max_instance_count = 4
  }

 tags {
     key   = "Env"
     value = "production"
  } 
}
```

## Argument Reference

The following arguments are supported:

* `ocean_id` - (Required) The ocean cluster you wish to 
* `name` - (Optional) Set Launch Specification name 
* `user_data` - (Optional) Base64-encoded MIME user data to make available to the instances.
* `image_id` - (Optional) ID of the image used to launch the instances.
* `iam_instance_profile` - (Optional) The ARN or name of an IAM instance profile to associate with launched instances.
* `security_groups` - (Optional) Optionally adds security group IDs.
* `subnet_ids` - (Optional) Set subnets in launchSpec. Each element in array should be subnet ID.
* `intance_types` - (Optional) A list of instance types allowed to be provisioned for pods pending under the specified launch specification. The list overrides the list defined for the Ocean cluster. 
* `root_volume_size` - (Optional) Set root volume size (in GB).
* `tags` - (Optional) A key/value mapping of tags to assign to the resource.

* `labels` - (Optional) Optionally adds labels to instances launched in an Ocean cluster.
    * `key` - (Required) The tag key.
    * `value` - (Required) The tag value.
    
* `taints` - (Optional) Optionally adds labels to instances launched in an Ocean cluster.
    * `key` - (Required) The tag key.
    * `value` - (Required) The tag value.
    * `effect` - (Required) The effect of the taint. Valid values: `"NoSchedule"`, `"PreferNoSchedule"`, `"NoExecute"`.

* `elastic_ip_pool` - (Optional) Assign an Elastic IP to the instances spun by the launch spec. Can be null.
    * `tag_selector` - (Optional) Key-value object, which defines an Elastic IP from the customer pool. Can be null.
        * `tag_key` - (Required) Elastic IP tag key. The launch spec will consider all elastic IPs tagged with this tag as a part of the elastic IP pool to use.
        * `tag_value` - (Optional) Elastic IP tag value. Can be null.
        
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
        * `dynamic_volume_size`- (Optional) Object. Set dynamic volume size properties. When using this object, you cannot use volumeSize. You must use one or the other.
            * `base_size`- (Required) Int. Initial size for volume. (Example: 50)
            * `resource`- (Required) String. Resource type to increase volume size dynamically by. (valid values: "CPU")
            * `size_per_resource_unit`- (Required) Int. Additional size (in GB) per resource unit. (Example: baseSize= 50, sizePerResourceUnit=20, and instance with 2 CPU is launched - its total disk size will be: 90GB)
        * `no_device`- (Optional) String. suppresses the specified device included in the block device mapping of the AMI.
   
* `autoscale_headrooms` - (Optional) Set custom headroom per launch spec. provide list of headrooms object.
    * `num_of_units` - (Required) The number of units to retain as headroom, where each unit has the defined headroom CPU, memory and GPU.
    * `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
    * `gpu_per_unit` - (Optional) Optionally configure the number of GPUS to allocate for each headroom unit.
    * `memory_per_unit` - (Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.

* `resource_limits` - (Optional) 
    * `max_instance_count` - (Optional) set a maximum number of instances per launch specification. Can be null. If set, value must be greater than or equal to 0.
    

