---
layout: "spotinst"
page_title: "Spotinst: ocean_aws"
sidebar_current: "docs-do-resource-ocean_aws"
description: |-
  Provides a Spotinst Ocean resource using AWS.
---

# spotinst\_ocean\_aws

Provides a Spotinst Ocean AWS resource.

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

  // --- LAUNCH CONFIGURATION --------------
  image_id             = "ami-123456"
  security_groups      = ["sg-987654321"]
  key_name             = "fake key"
  user_data            = "echo hello world"
  iam_instance_profile = "iam-profile"
  // ---------------------------------------

  // --- STRATEGY --------------------
  fallback_to_ondemand       = true
  spot_percentage            = 100
  utilize_reserved_instances = false
  // ---------------------------------

  // --- AUTOSCALER -----------------
  autoscaler = {
    autoscale_is_enabled     = false
    autoscale_is_auto_config = false
    autoscale_cooldown       = 300

    autoscale_headroom = {
      cpu_per_unit    = 1024
      memory_per_unit = 512
      num_of_units    = 2
    }

    autoscale_down = {
      evaluation_periods = 300
    }

    resource_limits = {
      max_vcpu       = 1024
      max_memory_gib = 20
    }
  }
  // --------------------------------
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

```hcl
  name = "demo"
  controller_id = "fakeClusterId"
  region = "us-west-2"

  max_size         = 2
  min_size         = 1
  desired_capacity = 2
```

* `whitelist` - (Optional) Instance types allowed in the Ocean cluster. Cannot be configured if `blacklist` is configured.
* `blacklist` - (Optional) Instance types not allowed in the Ocean cluster. Cannot be configured if `whitelist` is configured.

```hcl
whitelist = ["t1.micro", "m1.small"]
// blacklist = ["t1.micro", "m1.small"]
```

* `user_data` - (Optional) Base64-encoded MIME user data to make available to the instances.
* `image_id` - (Required) ID of the image used to launch the instances.
* `security_groups` - (Required) One or more security group ids.
* `key_name` - (Optional) The key pair to attach the instances.
* `iam_instance_profile` - (Optional) The instance profile iam role.

```hcl
  image_id             = "ami-79826301"
  security_groups      = ["sg-042d658b3ee907848"]
  key_name             = "fake key"
  user_data            = "echo hello world"
  iam_instance_profile = "iam-profile"
```

* `fallback_to_ondemand` - (Optional, Default: `true`) If not Spot instance markets are available, enable Ocean to launch On-Demand instances instead.
* `spot_percentage` - (Optional, Default: `100`) The percentage of Spot instances the cluster should maintain. Min 0, max 100.
* `utilize_reserved_instances` - (Optional, Default `false`) If Reserved instances exist, OCean will utilize them before launching Spot instances.

```hcl
  fallback_to_ondemand       = true
  spot_percentage            = 100
  utilize_reserved_instances = false
```

* `autoscaler` - (Optional) Describes the Ocean Kubernetes autoscaler.
* `autoscale_is_enabled` - (Optional, Default: `true`) Enable the Ocean Kubernetes autoscaler.
* `autoscale_is_auto_config` - (Optional, Default: `true`) Automatically configure and optimize headroom resources.
* `autoscale_cooldown` - (Optional, Default: `null`) Cooldown period between scaling actions.
* `autoscale_headroom` - (Optional) Spare resource capacity management enabling fast assignment of Pods without waiting for new resources to launch.
* `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
* `memory_per_unit` - (Optional) Optionally configure the amount of memory (MB) to allocate the headroom.
* `num_of_units` - (Optional) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.
* `autoscale_down` - (Optional) Auto Scaling scale down operations.
* `evaluation_periods` - (Optional, Default: `null`) The number of evaluation periods that should accumulate before a scale down action takes place.
* `resource_limits` - (Optional) Optionally set upper and lower bounds on the resource usage of the cluster.
* `max_vcpu` - (Optional) The maximum cpu in vCPU units that can be allocated to the cluster.
* `max_memory_gib` - (Optional) The maximum memory in GiB units that can be allocated to the cluster.

```hcl
  autoscaler = {
    autoscale_is_enabled     = false
    autoscale_is_auto_config = false
    autoscale_cooldown       = 300

    autoscale_headroom = {
      cpu_per_unit    = 1024
      memory_per_unit = 512
      num_of_units    = 2
    }

    autoscale_down = {
      evaluation_periods = 300
    }

    resource_limits = {
      max_vcpu       = 1024
      max_memory_gib = 20
    }
  }
```

* `tags` - (Optional) Optionally adds tags to instances launched in an Ocean cluster.
* `tag_key` - (Optional) The tag key.
* `tag_value` - (Optional) The tag value.

```hcl
tags = [{
  tag_key   = "fakeKey"
  tag_value = "fakeValue"
}]
```
