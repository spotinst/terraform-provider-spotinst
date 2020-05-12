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
        
* `autoscale_headrooms` - (Optional) Set custom headroom per launch spec. provide list of headrooms object.
    * `num_of_units` - (Required) The number of units to retain as headroom, where each unit has the defined headroom CPU, memory and GPU.
    * `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
    * `gpu_per_unit` - (Optional) Optionally configure the number of GPUS to allocate for each headroom unit.
    * `memory_per_unit` - (Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.

