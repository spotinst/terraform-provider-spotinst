---
layout: "spotinst"
page_title: "Spotinst: ocean_ecs_launch_spec"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean ECS Launch Spec resource using AWS.
---

# spotinst\_ocean\_ecs\_launch\_spec

Provides a custom Spotinst Ocean ECS Launch Spec resource.

## Example Usage

```hcl
resource "spotinst_ocean_ecs_launch_spec" "example" {
  ocean_id  = "o-123456"
  image_id  = "ami-123456"
  user_data = "echo hello world"
  iam_instance_profile = "iam-profile"
  security_group_ids = ["awseb-12345"]
  
  attributes {
    key   = "fakeKey"
    value = "fakeValue"
  }
  
  autoscale_headrooms {
    num_of_units = 5
    cpu_per_unit = 1000
    memory_per_unit = 2048
  }

  tags {
     key   = "Env"
     value = "production"
  } 
}
```

## Argument Reference

The following arguments are supported:

* `ocean_id`  - (Required) The Ocean cluster ID .
* `name`      - (Required) The Ocean Launch Specification name. 
* `user_data` - (Optional) Base64-encoded MIME user data to make available to the instances.
* `image_id`  - (Optional) ID of the image used to launch the instances.
* `iam_instance_profile` - (Optional) The ARN or name of an IAM instance profile to associate with launched instances.
* `security_group_ids` - (Optional) One or more security group ids.
* `tags` - (Optional) A key/value mapping of tags to assign to the resource.

* `attributes` - (Optional) Optionally adds labels to instances launched in an Ocean cluster.
    * `key` - (Required) The label key.
    * `value` - (Required) The label value.
   
* `autoscale_headrooms` - (Optional) Set custom headroom per launch spec. provide list of headrooms object.
    * `num_of_units` - (Required) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.
    * `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in CPU units, where 1024 units = 1 vCPU.
    * `memory_per_unit` - (Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.
