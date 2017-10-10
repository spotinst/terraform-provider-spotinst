---
layout: "spotinst"
page_title: "Spotinst: Mr Scaler"
sidebar_current: "docs-do-resource-mr_scaler"
description: |-
  Provides a Spotinst MrScaler group resource.
---

# spotinst\_mrscaler

Provides a Spotinst AWS MrScaler resource.

## Example Usage - Clone Strategy

```hcl
# Create a Mr Scaler with Clone strategy and Task scaling

resource "spotinst_mrscaler" "example-scaler-1" {
  name        = "spotinst-mr-scaler-1"
  description = "created by Terraform"
  region = "us-west-2"
  strategy = "clone"
  cluster_id = "j-27UVDEHXL4OQM"

  master_instance_types = ["c3.xlarge","c4.xlarge"]
  master_target = 1
  master_lifecycle = "ON_DEMAND"
  master_ebs_block_device {
    volumes_per_instance = 2
    volume_type = "gp2"
    size_in_gb = 30
  }

  core_instance_types = ["c3.xlarge","c4.xlarge"]
  core_target = 2
  core_maximum = 8
  core_minimum = 1
  core_lifecycle = "SPOT"
  core_ebs_block_device {
    volume_type = "gp2"
    size_in_gb = 20
  }


  task_instance_types = ["c3.xlarge","c4.xlarge"]
  task_target = 3
  task_minimum = 0
  task_maximum = 4
  task_lifecycle = "SPOT"
  task_ebs_block_device {
    volume_type = "gp2"
    size_in_gb = 20
  }

  task_scaling_up_policy {
    policy_name        = "scaling policy up 1"
    metric_name        = "CPUUtilization"
    statistic          = "average"
    unit               = "percent"
    threshold          = 30
    namespace          = "AWS/EC2"
    operator           = "gte"
    evaluation_periods = 5
    period             = 60
    cooldown           = 1200
    action_type        = "adjustment"
    adjustment         = 1
  }

  tags = [
    {
     key="Creator"
     value= "Spotinst"
    }
  ]

  availability_zones = ["us-west-2a:subnet-79da021e"]


}
```

## Example Usage - Wrap Strategy

```hcl
# Create a Mr Scaler with Wrap strategy

resource "spotinst_mrscaler" "example-scaler-2" {
  name        = "spotinst-mr-scaler-2"
  description = "created by Terraform"
  region = "us-west-2"
  strategy = "wrap"
  cluster_id = "j-27UVDEHXL4OQM"
  task_instance_types = ["c3.xlarge","c4.xlarge"]
  task_target = 2
  task_minimum = 0
  task_maximum = 4
  task_lifecycle = "SPOT"
  task_ebs_block_device {
    volumes_per_instance = 1
    volume_type = "gp2"
    size_in_gb = 20
  }
}

```


## Argument Reference

The following arguments are supported:

* `name` - (Required) The MrScaler name.
* `description` - (Optional) The MrScaler description.
* `region` - (Required) The MrScaler region.
* `strategy` - (Required) The MrScaler strategy. Allowed values are 'clone' and 'wrap'.
* `cluster_id` - (Required) The MrScaler cluster id.

<a id="task-group"></a>
## Task Group (Wrap and Clone strategies)
* `task_instance_types` - (Required) The MrScaler instance types for the task nodes.
* `task_target` - (Required) amount of instances in task group.
* `task_maximum` - (Optional) maximal amount of instances in task group.
* `task_minimum` - (Optional) The minimal amount of instances in task group.
* `task_lifecycle` - (Required) The MrScaler lifecycle for instances in task group. Allowed values are 'SPOT' and 'ON_DEMAND'.
* `task_ebs_block_device` - (Required) This determines the ebs configuration for your task group instances. Only a single block is allowed.
    * `volumes_per_instance` - (Optional; Default 1) Amount of volumes per instance in the task group.
    * `volume_type` - (Required) volume type. Allowed values are 'gp2', 'io1' and others.
    * `size_in_gb` - (Required) Size of the volume, in GBs.
    * `iops` - (Optional) IOPS for the volume. Required in some volume types, such as io1.

<a id="core-group"></a>
## Core Group (Clone strategy only)
* `core_instance_types` - (Required) The MrScaler instance types for the core nodes.
* `core_target` - (Required) amount of instances in core group.
* `core_maximum` - (Optional) maximal amount of instances in core group.
* `core_minimum` - (Optional) The minimal amount of instances in core group.
* `core_lifecycle` - (Required) The MrScaler lifecycle for instances in core group. Allowed values are 'SPOT' and 'ON_DEMAND'.
* `core_ebs_block_device` - (Required) This determines the ebs configuration for your core group instances. Only a single block is allowed.
    * `volumes_per_instance` - (Optional; Default 1) Amount of volumes per instance in the core group.
    * `volume_type` - (Required) volume type. Allowed values are 'gp2', 'io1' and others.
    * `size_in_gb` - (Required) Size of the volume, in GBs.
    * `iops` - (Optional) IOPS for the volume. Required in some volume types, such as io1.

<a id="master-group"></a>
## Master Group (Clone strategy only)
* `master_instance_types` - (Required) The MrScaler instance types for the master nodes.
* `master_target` - (Required) amount of instances in master group.
* `master_lifecycle` - (Required) The MrScaler lifecycle for instances in master group. Allowed values are 'SPOT' and 'ON_DEMAND'.
* `master_ebs_block_device` - (Required) This determines the ebs configuration for your master group instances. Only a single block is allowed.
    * `volumes_per_instance` - (Optional; Default 1) Amount of volumes per instance in the master group.
    * `volume_type` - (Required) volume type. Allowed values are 'gp2', 'io1' and others.
    * `size_in_gb` - (Required) Size of the volume, in GBs.
    * `iops` - (Optional) IOPS for the volume. Required in some volume types, such as io1.


<a id="tags"></a>
## Tags (Clone strategy only)
* `tags` - (Optional) A list of tags to assign to the resource. You may define multiple tags.
    * `key` - (Required) Tag key.
    * `value` - (Required) Tag value.

<a id="availability-zone"></a>
## Availability Zones (Clone strategy only)

* `availability_zones` - (Required in Clone) List of AZs and their subnet Ids. See example above for usage.

<a id="scaling-policy"></a>
## Scaling Policies

Possible task group scaling policies (Wrap and Clone strategies):
* `task_scaling_up_policy`
* `task_scaling_down_policy`

Possible core group scaling policies (Clone strategy only):
* `core_scaling_up_policy`
* `core_scaling_down_policy`

Each `*_scaling_*_policy` supports the following:

* `namespace` - (Required) The namespace for the alarm's associated metric.
* `metric_name` - (Required) The name of the metric, with or without spaces.
* `threshold` - (Required) The value against which the specified statistic is compared.
* `policy_name` - (Optional) The name of the policy.
* `statistic` - (Optional) The metric statistics to return. For information about specific statistics go to [Statistics](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/index.html?CHAP_TerminologyandKeyConcepts.html#Statistic) in the Amazon CloudWatch Developer Guide.
* `unit` - (Optional) The unit for the alarm's associated metric.
* `adjustment` - (Required) The number of instances to add/remove to/from the target capacity when scale is needed.
* `action_type` - (Required) The number of instances to add/remove to/from the target capacity when scale is needed.
* `period` - (Optional) The granularity, in seconds, of the returned datapoints. Period must be at least 60 seconds and must be a multiple of 60.
* `evaluation_periods` - (Optional) The number of periods over which data is compared to the specified threshold.
* `cooldown` - (Optional) The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start. If this parameter is not specified, the default cooldown period for the group applies.
* `dimensions` - (Optional) A mapping of dimensions describing qualities of the metric.


## Attributes Reference

The following attributes are exported:

* `id` - The scaler ID.
