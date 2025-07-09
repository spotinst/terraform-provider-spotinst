---
layout: "spotinst"
page_title: "Spotinst: ocean_aws_launch_spec"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Virtual Node Group resource using AWS.
---

# spotinst\_ocean\_aws\_launch\_spec

Manages a Spotinst Ocean AWS [Virtual Node Group](https://docs.spot.io/ocean/features/launch-specifications) resource.

## Example Usage

```hcl
resource "spotinst_ocean_aws_launch_spec" "example" {
  ocean_id                    = "o-123456"
  name                        = "example"
  image_id                    = "ami-123456"
  user_data                   = "echo Hello, world!"
  iam_instance_profile        = "iam-profile"
  security_groups             = ["sg-987654321"]
  subnet_ids                  = ["subnet-1234"]
  restrict_scale_down         = true
  root_volume_size            = 30
  associate_public_ip_address = true
  reserved_enis               = 1
  
  images {
    image_id = "ami-id1"
    }
  images {
    image_id = "ami-id2"
    }

  instance_types = [
    "m4.large",
    "m4.xlarge",
    "m4.2xlarge",
    "m4.4xlarge",
  ]

  preferred_spot_types = [
    "m4.large",
    "m4.xlarge"
  ]
  
  preferred_od_types = [
    "c3.large",
    "m4.large"
  ]
  
  labels {
    key   = "key1"
    value = "value1"
  }

  taints {
    key    = "key1"
    value  = "value1"
    effect = "NoExecute"
  }

  startup_taints {
    key    = "another-key"
    value  = "another-value"
    effect = "NoSchedule"
  }

  autoscale_headrooms_automatic {
    auto_headroom_percentage = 5
  }
  
  autoscale_headrooms {
    num_of_units    = 5
    cpu_per_unit     = 1000
    gpu_per_unit    = 0
    memory_per_unit = 2048
  }
  
  autoscale_down {
    max_scale_down_percentage    = 20
  }
  
  elastic_ip_pool {
    tag_selector {
      tag_key   = "key"
      tag_value = "value"
    }
  }

  block_device_mappings {
    device_name = "/dev/xvda"
    ebs {
      delete_on_termination = "true"
      encrypted             = "false"
      volume_type           = "gp2"
      volume_size           = 50
      throughput            = 500
      dynamic_volume_size {
        base_size              = 50
        resource               = "CPU"
        size_per_resource_unit = 20
      }
    }
  }
  
  ephemeral_storage{
    ephemeral_storage_device_name = "/dev/xvda"
  }

  resource_limits {
    max_instance_count = 4
    min_instance_count = 0
  }

  tags {
    key   = "Env"
    value = "production"
  }

  strategy {
    spot_percentage = 70
    draining_timeout = 300
    utilize_commitments= false
    utilize_reserved_instances= true
    cluster_orientation {
        availability_vs_cost = "balanced"
    }
  }
  
  create_options {
    initial_nodes = 1
  }
  
  delete_options {
    force_delete = true
    delete_nodes = true
  }
  
  instance_metadata_options {
    http_tokens = "required"
    http_put_response_hop_limit = 10
  }
  
  scheduling_task {
    is_enabled = true
    cron_expression = "0 1 * * *"
    task_type = "manualHeadroomUpdate"
    task_headroom {
        num_of_units    = 5
        cpu_per_unit     = 1000
        gpu_per_unit    = 0
        memory_per_unit = 2048
    }
  }
  
  scheduling_shutdown_hours {
    is_enabled = true
    time_windows = ["Sat:08:00-Sat:08:30", "Sun:08:00-Sun:08:30"]
  }
}

  instance_types_filters {
    categories                =   ["Accelerated_computing", "Compute_optimized"]
    disk_types                =   ["NVMe", "EBS"]
    exclude_families          =   ["t2","R4*"]
    exclude_metal             =   true
    hypervisor                =   ["nitro"]
    include_families          =   ["c5*", "g5"]
    is_ena_supported          =   true
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
```
```
output "ocean_launchspec_id" {
  value = spotinst_ocean_aws_launch_spec.example.id
}
```

## Argument Reference

The following arguments are supported:

* `ocean_id` - (Required) The ID of the Ocean cluster. 
* `name` - (Optional) The name of the Virtual Node Group.
* `user_data` - (Optional) Base64-encoded MIME user data to make available to the instances.
* `image_id` - (Optional) ID of the image used to launch the instances.
* `images` - Array of objects (Image object, containing the id of the image used to launch instances.) You can configure VNG with either the imageId or images objects, but not both simultaneously.
              For each architecture type (amd64, arm64) only one AMI is allowed. Valid values: null, or an array with at least one element.
    * `image_id` - Identifier of the image in AWS. Valid values: any string which is not empty or null.
* `iam_instance_profile` - (Optional) The ARN or name of an IAM instance profile to associate with launched instances.
* `security_groups` - (Optional) Optionally adds security group IDs.
* `subnet_ids` - (Optional) A list of subnet IDs.
* `instance_types` - (Optional) A list of instance types allowed to be provisioned for pods pending under the specified launch specification. The list overrides the list defined for the cluster.
* `preferred_spot_types` - (Optional) A list of instance types. Takes the preferred types into consideration while maintaining a variety of machine types running for optimized distribution.
* `preferred_od_types` - (Optional) A list of instance types. Takes the preferred types into consideration while maintaining a variety of machine types running for optimized distribution.
* `root_volume_size` - (Optional) Set root volume size (in GB).
* `tags` - (Optional) A key/value mapping of tags to assign to the resource.
* `associate_public_ip_address` - (Optional, Default: `false`) Configure public IP address allocation.
* `restrict_scale_down`- (Optional) Boolean. When set to `true`, nodes will be treated as if all pods running have the restrict-scale-down label. Therefore, Ocean will not scale nodes down unless empty.
* `reserved_enis` - (Optional, Default: `0`) Specifies the count of ENIs to reserve per instance type for scaling purposes.
* `labels` - (Optional) Optionally adds labels to instances launched in the cluster.
    * `key` - (Required) The label key.
    * `value` - (Required) The label value.
* `taints` - (Optional) Optionally adds labels to instances launched in the cluster.
    * `key` - (Required) The taint key.
    * `value` - (Required) The taint value.
    * `effect` - (Required) The effect of the taint. Valid values: `"NoSchedule"`, `"PreferNoSchedule"`, `"NoExecute"`.
* `startup_taints` - (Optional) Temporary taints applied to a node during its initialization phase. For a startup taint to work, it must also be set as a regular taint in the userData for the cluster.
    * `key` - (Optional) Set startup taint key.
    * `value` - (Optional) Set startup taint value.
    * `effect` - (Optional) Set startup taint effect.
* `elastic_ip_pool` - (Optional) Assign an Elastic IP to the instances spun by the Virtual Node Group. Can be null.
    * `tag_selector` - (Optional) A key-value pair, which defines an Elastic IP from the customer pool. Can be null.
        * `tag_key` - (Required) Elastic IP tag key. The Virtual Node Group will consider all Elastic IPs tagged with this tag as a part of the Elastic IP pool to use.
        * `tag_value` - (Optional) Elastic IP tag value. Can be null.    
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
            * `size_per_resource_unit`- (Required) Int. Additional size (in GB) per resource unit. (Example: `baseSize=50`, `sizePerResourceUnit=20`, and instance with 2 CPU is launched; its total disk size will be: 90GB)
        * `no_device` - (Optional) String. Suppresses the specified device included in the block device mapping of the AMI. Default value is set to `unset` intentionally, which will appear in the terminal during a terraform plan if this field is not configured or removed. This prevents confusion, as Terraform otherwise considers empty string as null.
* `autoscale_headrooms_automatic` - (Optional) Set automatic headroom per launch spec.
    * `auto_headroom_percentage` - (Optional) Number between 0-200 to control the headroom % of the specific Virtual Node Group. Effective when cluster.autoScaler.headroom.automatic.`is_enabled` = true is set on the Ocean cluster.
* `autoscale_headrooms` - (Optional) Set custom headroom per Virtual Node Group. Provide a list of headrooms object.
    * `num_of_units` - (Required) The number of units to retain as headroom, where each unit has the defined headroom CPU, memory and GPU.
    * `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
    * `gpu_per_unit` - (Optional) Optionally configure the number of GPUS to allocate for each headroom unit.
    * `memory_per_unit` - (Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.
* `autoscale_down` - (Optional) Auto Scaling scale down operations.
    * `max_scale_down_percentage` - (Optional) The maximum percentage allowed to scale down in a single scaling action on the nodes running in a specific VNG. Allowed only if maxScaleDownPercentage is set to null at the cluster level. Number between [0.1-100].
* `resource_limits` - (Optional) 
    * `max_instance_count` - (Optional) Set a maximum number of instances per Virtual Node Group. Can be null. If set, value must be greater than or equal to 0.
    * `min_instance_count` - (Optional) Set a minimum number of instances per Virtual Node Group. Can be null. If set, value must be greater than or equal to 0.
* `strategy` - (Optional) 
    * `spot_percentage` - (Optional) The desired percentage of the Spot instances out of all running instances for this VNG. Only available when the field is not set in the cluster directly (cluster.strategy.spotPercentage).
    * `draining_timeout` - (Optional, >=300) The configurable amount of time that Ocean will wait for the draining process to complete before terminating an instance. If you have not defined a draining timeout, the default of 300 seconds will be used.
    * `utilize_commitments` - (Optional, Default: `false`) When set as ‘true’, if savings plans commitments have available capacity, Ocean will utilize them alongside RIs (if exist) to maximize cost efficiency. If the value is set as 'null', it will automatically be inherited from the cluster level.
    * `utilize_reserved_instances` - (Optional, Default: `true`) When set as ‘true’, if reserved instances exist, Ocean will utilize them before launching spot instances. If the value is set as 'null', it will automatically be inherited from the cluster level.
    * `orientation` - (Optional) Vng orientation configuration.
        * `availability_vs_cost` - (Optional, Default: `"balanced"`) Set this value to control the approach that Ocean takes while launching nodes. Valid values: `"costOriented"`, `"cheapest"`, `"balanced"`.
* `create_options` - (Optional)
    * `initial_nodes` - (Optional) When set to an integer greater than 0, a corresponding amount of nodes will be launched from the created Virtual Node Group. The parameter is recommended in case the use_as_template_only (in spotinst_ocean_aws resource) is set to true during Ocean resource creation.
* `delete_options` - (Optional)
    * `force_delete` - (Optional) When set to `true`, delete even if it is the last Virtual Node Group (also, the default Virtual Node Group must be configured with `useAsTemlateOnly = true`). Should be set at creation or update, but will be used only at deletion.
    * `delete_nodes` - (Optional) When set to "true", all instances belonging to the deleted launch specification will be drained, detached, and terminated.
* `scheduling_task` - (Optional) Used to define scheduled tasks such as a manual headroom update.
    * `is_enabled` - (Required) Describes whether the task is enabled. When `true`, the task runs. When `false`, it does not run.
    * `cron_expression` - (Required) A valid cron expression. For example : " * * * * * ". The cron job runs in UTC time and is in Unix cron format.
    * `task_type` - (Required) The activity that you are scheduling. Valid values: "manualHeadroomUpdate".
    * `task_headroom` - (Optional) The config of this scheduled task. Depends on the value of taskType.
      * `num_of_units` - (Required) The number of units to retain as headroom, where each unit has the defined headroom CPU, memory and GPU.
      * `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
      * `gpu_per_unit` - (Optional) Optionally configure the number of GPUS to allocate for each headroom unit.
      * `memory_per_unit` - (Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.
* `scheduling_shutdown_hours` - (Optional) Used to specify times that the nodes in the virtual node group will be taken down.
    * `time_windows` - (Required ) The times that the shutdown hours will apply.
    * `is_enabled` - (Optional) Flag to enable or disable the shutdown hours mechanism. When `false`, the mechanism is deactivated, and the virtual node group remains in its current state.
* `instance_metadata_options` - (Optional) Ocean instance metadata options object for IMDSv2.
    * `http_tokens` - (Required) Determines if a signed token is required or not. Valid values: `optional` or `required`.
    * `http_put_response_hop_limit` - (Optional) An integer from 1 through 64. The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further the instance metadata requests can travel.
* `instance_store_policy` - (Optional) Determines the utilization of instance store volumes. If not defined, instance store volumes will not be used.
    * `instance_store_policy_type` - (Optional) Value: `"RAID0"` The method for using the instance store volumes (must also be defined in the userData).
* `instance_types_filters` - The instance types that match with all filters compose the Virtual Node Group's instanceTypes parameter.
The architectures that come from the Virtual Node Group's images will be taken into account when using this parameter. Cannot be configured together with Virtual Node Group's instanceTypes and with the Cluster's whitelist/blacklist/filters.
    * `categories` - (Optional) The filtered instance types will belong to one of the categories types from this list. Valid values: `Accelerated_computing`, `Compute_optimized`, `General_purpose`, `Memory_optimized`, `Storage_optimized`.
    * `disk_types` - (Optional) The filtered instance types will have one of the disk type from this list. Valid values: `NVMe`, `EBS`, `SSD`, `HDD`.
    * `exclude_families` - (Optional) Types belonging to a family from the ExcludeFamilies will not be available for scaling (asterisk wildcard is also supported). For example, C* will exclude instance types from these families: c5, c4, c4a, etc.
    * `exclude_metal` - (Optional, Default: `false`) In case excludeMetal is set to true, metal types will not be available for scaling.
    * `hypervisor` - (Optional) The filtered instance types will have a hypervisor type from this list. Valid values: `nitro`, `xen`.
    * `include_families` - (Optional) Types belonging to a family from the IncludeFamilies will be available for scaling (asterisk wildcard is also supported). For example, C* will include instance types from these families: c5, c4, c4a, etc.
    * `is_ena_supported` - (Optional) Ena is supported or not.
    * `max_gpu` - (Optional, >=1) Maximum total number of GPUs.
    * `max_memory_gib` - (Optional, >=0) Maximum amount of Memory (GiB).
    * `max_network_performance` - (Optional, >=1) Maximum Bandwidth in Gib/s of network performance.
    * `max_vcpu` - (Optional, >=1) Maximum number of vcpus available.
    * `min_enis` - (Optional, >=0) Minimum number of network interfaces (ENIs).
    * `min_gpu` - (Optional, >=0) Minimum total number of GPUs.
    * `min_memory_gib` - (Optional, >=0) Minimum amount of Memory (GiB).
    * `min_network_performance` - (Optional, >=0) Minimum Bandwidth in Gib/s of network performance.
    * `min_vcpu` - (Optional, >=0) Minimum number of vcpus available.
    * `root_device_types` - (Optional) The filtered instance types will have a root device types from this list. Valid values: `ebs`, or `instance-store`.
    * `virtualization_types` - (Optional) The filtered instance types will support at least one of the virtualization types from this list. Valid values: `hvm`, `paravirtual`.
* `ephemeral_storage` - (Optional)
    * `ephemeral_storage_device_name` - (Optional) Specify an alternative device name from which ephemeral storage calculations should be derived. This parameter is used when the ephemeral storage should not utilize the root device. Provide the device name configured in the VNG's BDM or AMI's BDM that differs from the default root device.


<a id="update-policy"></a>
## Update Policy

* `update_policy` - (Optional)
    * `should_roll` - (Required) Enables the roll.
    * `roll_config` - (Required) Holds the roll configuration.
        * `batch_size_percentage` - (Required) Sets the percentage of the instances to deploy in each batch.
        * `respect_pdb` - (Optional, Default: false) During the roll, if the parameter is set to `true` we honor PDB during the instance replacement.


```hcl
update_policy {
  should_roll = false

  roll_config {
    batch_size_percentage = 33
    respect_pdb = true
  }
}
```

## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `id` - The Virtual Node Group ID.


<a id="import"></a>
## Import

Launch_Specs can be imported using the Launch_Spec `id`, e.g.,
```hcl
$ terraform import spotinst_ocean_aws_launch_spec.nameOfTheResource ols-1a2b576
```
