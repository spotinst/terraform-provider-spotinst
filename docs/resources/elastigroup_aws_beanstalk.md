---
layout: "spotinst"
page_title: "Spotinst: elastigroup_aws_beanstalk"
subcategory: "Elastigroup"
description: |-
 Provides a Spotinst AWS group resource using Elastic Beanstalk.
---

# spotinst\_elastigroup\_aws\_beanstalk

Provides a Spotinst AWS group resource using Elastic Beanstalk.

## Example Usage

```hcl
resource "spotinst_elastigroup_aws_beanstalk" "elastigoup-aws-beanstalk" {

 name    = "example-elastigroup-beanstalk"
 region  = "us-west-2"
 product = "Linux/UNIX"

 min_size         = 0
 max_size         = 1
 desired_capacity = 0

 beanstalk_environment_name = "example-env"
 beanstalk_environment_id   = "e-example"
 instance_types_spot        = ["t2.micro", "t2.medium", "t2.large"]

 deployment_preferences {
  automatic_roll        = true
  batch_size_percentage = 100
  grace_period          = 90
    strategy {
      action                 = "REPLACE_SERVER"
      should_drain_instances = true
    }
 }
  
 managed_actions = {
  platform_update = {
    perform_at   = "timeWindow"
    time_window  = "Mon:23:50-Tue:00:20"
    update_level = "minorAndPatch"
  }
 }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The group name.
* `region` - (Required) The AWS region your group will be created in. Cannot be changed after the group has been created.
* `description` - (Optional) The group description.
* `product` - (Required) Operation system type. Valid values: `"Linux/UNIX"`, `"SUSE Linux"`, `"Windows"`.
For EC2 Classic instances:  `"Linux/UNIX (Amazon VPC)"`, `"SUSE Linux (Amazon VPC)"`, `"Windows (Amazon VPC)"`.   

* `max_size` - (Required) The maximum number of instances the group should have at any time.
* `min_size` - (Required) The minimum number of instances the group should have at any time.
* `desired_capacity` - (Required) The desired number of instances the group should have at any time.

* `beanstalk_environment_name` - (Optional) The name of an existing Beanstalk environment.
* `beanstalk_environment_id` - (Optional) The id of an existing Beanstalk environment. 
* `instance_types_spot` - (Required) One or more instance types. To maximize the availability of Spot instances, select as many instance types as possible.

* `deployment_preferences` - (Optional) Preferences when performing a roll
   * `automatic_roll` - (Required) Should roll perform automatically
   * `batch_size_percentage` - (Required) Percent size of each batch
   * `grace_period` - (Required) Amount of time to wait between batches
   * `strategy` - (Optional) Strategy parameters
      * `action` - (Required) Action to take
      * `should_drain_instances` - (Required) Bool value if to wait to drain instance 

* `managed_actions` - (Optional) Managed Actions parameters
   * `platform_update` - (Optional) Platform Update parameters
      * `perform_at` - (Required) Actions to perform (options: timeWindow, never)
      * `time_window` - (Required) Time Window for when action occurs ex. Mon:23:50-Tue:00:20
      * `update_level` - (Required) - Level to update
    
### Scheduled Tasks

Each `scheduled_task` supports the following:

* `task_type` - (Required) The task type to run. Supported task types are: `"scale"`, `"backup_ami"`, `"roll"`, `"scaleUp"`, `"percentageScaleUp"`, `"scaleDown"`, `"percentageScaleDown"`, `"statefulUpdateCapacity"`.
* `cron_expression` - (Optional; Required if not using `frequency`) A valid cron expression. The cron is running in UTC time zone and is in [Unix cron format](https://en.wikipedia.org/wiki/Cron).
* `start_time` - (Optional; Format: ISO 8601) Set a start time for one time tasks.
* `frequency` - (Optional; Required if not using `cron_expression`) The recurrence frequency to run this task. Supported values are `"hourly"`, `"daily"`, `"weekly"` and `"continuous"`.
* `scale_target_capacity` - (Optional) The desired number of instances the group should have.
* `scale_min_capacity` - (Optional) The minimum number of instances the group should have.
* `scale_max_capacity` - (Optional) The maximum number of instances the group should have.
* `is_enabled` - (Optional, Default: `true`) Setting the task to being enabled or disabled.
* `target_capacity` - (Optional; Only valid for statefulUpdateCapacity) The desired number of instances the group should have.
* `min_capacity` - (Optional; Only valid for statefulUpdateCapacity) The minimum number of instances the group should have.
* `max_capacity` - (Optional; Only valid for statefulUpdateCapacity) The maximum number of instances the group should have.
* `batch_size_percentage` - (Optional; Required when the `task_type` is `"roll"`.) The percentage size of each batch in the scheduled deployment roll.
* `grace_period` - (Optional) The period of time (seconds) to wait before checking a batch's health after it's deployment.
* `adjustment` - (Optional; Min 1) The number of instances to add or remove.
* `adjustment_percentage` - (Optional; Min 1) The percentage of instances to add or remove.

Usage:

```hcl
  scheduled_task {
    task_type             = "backup_ami"
    cron_expression       = ""
    start_time            = "1970-01-01T01:00:00Z"
    frequency             = "hourly"
    scale_target_capacity = 5
    scale_min_capacity    = 0
    scale_max_capacity    = 10
    is_enabled            = false
    target_capacity       = 5
    min_capacity          = 0
    max_capacity          = 10
    batch_size_percentage = 33
    grace_period          = 300
  }
```
