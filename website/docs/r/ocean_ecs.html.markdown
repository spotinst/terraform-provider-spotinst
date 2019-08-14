---
layout: "spotinst"
page_title: "Spotinst: ocean_ecs"
sidebar_current: "docs-do-resource-ocean_ecs"
description: |-
  Provides a Spotinst Ocean ECS resource using AWS.
---

# spotinst\_ocean\_ecs

Provides a Spotinst Ocean ECS resource.

## Example Usage

```hcl
resource "spotinst_ocean_ecs" "example" {
    region = "us-west-2"
    name = "terraform-ecs-cluster"
    cluster_name = "terraform-ecs-cluster"
  
    min_size         = "0"
    max_size         = "1"
    desired_capacity = "0"
  
    autoscaler {
      cooldown = 240
      headroom {
        cpu_per_unit = 512
        memory_per_unit = 1024
        num_of_units = 1
      }
      down {
        max_scale_down_percentage = 20
      }
      is_auto_config = false
      is_enabled = false
      resource_limits {
        max_vcpu = 1
        max_memory_gib = 2
      }
    }
  
    subnet_ids = ["subnet-12345"]
    whitelist = ["t3.medium"]
  
    security_group_ids = ["sg-12345"]
    image_id = "ami-12345"
    iam_instance_profile = "arn:aws:iam::12345:instance-profile/ecsInstanceProfile"
  
    key_pair = "KeyPair"
    user_data = "IyEvYmluL2Jhc2gKZWNobyB0ZXJyYWZvcm0tZWNzLWNsdXN0ZXIgPj4gL2V0Yy9lY3MvZWNzLmNvbmZpZw=="
    associate_public_ip_address = false
  
    update_policy {
      should_roll = true
      roll_config {
        batch_size_percentage = 100
      }
    }
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
* `subnet_ids` - (Optional) A comma-separated list of subnet identifiers for the Ocean cluster. Subnet IDs should be configured with auto assign public ip.

```hcl
  region = "us-west-2"
  name = "terraform-ecs-cluster"
  cluster_name = "terraform-ecs-cluster"

  max_size         = 2
  min_size         = 1
  desired_capacity = 2
  subnet_ids       = ["subnet-12345"]
```

* `whitelist` - (Optional) Instance types allowed in the Ocean cluster.

```hcl
whitelist = ["t1.micro", "m1.small"]
// blacklist = ["t1.micro", "m1.small"]
```

* `user_data` - (Optional) Base64-encoded MIME user data to make available to the instances.
* `image_id` - (Required) ID of the image used to launch the instances.
* `security_group_ids` - (Required) One or more security group ids.
* `key_pair` - (Optional) The key pair to attach the instances.
* `iam_instance_profile` - (Optional) The instance profile iam role.
* `associate_public_ip_address` - (Optional, Default: `false`) Configure public IP address allocation.
* `draining_timeout` - (Optional) The time in seconds, the instance is allowed to run while detached from the ELB. This is to allow the instance time to be drained from incoming TCP connections before terminating it, during a scale down operation.
* `monitoring` - (Optional) Enable detailed monitoring for cluster. Flag will enable Cloud Watch detailed detailed monitoring (one minute increments). Note: there are additional hourly costs for this service based on the region used.
* `ebs_optimized` - (Optional) Enable EBS optimized for cluster. Flag will enable optimized capacity for high bandwidth connectivity to the EB service for non EBS optimized instance types. For instances that are EBS optimized this flag will be ignored.


```hcl
  image_id                    = "ami-79826301"
  security_group_ids          = ["sg-042d658b3ee907848"]
  key_name                    = "fake key"
  user_data                   = "echo hello world"
  iam_instance_profile        = "iam-profile"
  associate_public_ip_address = true
  draining_timeout            = 120
  monitoring                  = true
  ebs_optimized               = true
```

* `autoscaler` - (Optional) Describes the Ocean Kubernetes autoscaler.
* `is_enabled` - (Optional, Default: `true`) Enable the Ocean Kubernetes autoscaler.
* `is_auto_config` - (Optional, Default: `true`) Automatically configure and optimize headroom resources.
* `cooldown` - (Optional, Default: `null`) Cooldown period between scaling actions.
* `headroom` - (Optional) Spare resource capacity management enabling fast assignment of Pods without waiting for new resources to launch.
* `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
* `memory_per_unit` - (Optional) Optionally configure the amount of memory (MB) to allocate the headroom.
* `num_of_units` - (Optional) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.
* `down` - (Optional) Auto Scaling scale down operations.
* `evaluation_periods` - (Optional, Default: `null`) The number of evaluation periods that should accumulate before a scale down action takes place.
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

    down = {
      max_scale_down_percentage = 20
    }

    resource_limits {
      max_vcpu       = 1024
      max_memory_gib = 20
    }
  }
```

* `tags` - (Optional) Optionally adds tags to instances launched in an Ocean cluster.
* `key` - (Optional) The tag key.
* `value` - (Optional) The tag value.

```hcl
tags = [{
  key   = "fakeKey"
  value = "fakeValue"
}]
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