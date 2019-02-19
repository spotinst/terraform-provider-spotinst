---
layout: "spotinst"
page_title: "Spotinst: elastigroup_aws_beanstalk"
sidebar_current: "docs-do-resource-elastigroup_aws_beanstalk"
description: |-
 Provides a Spotinst AWS group resource using Elastic Beanstalk.
---

# spotinst\_elastigroup\_aws\_beanstalk

Provides a Spotinst AWS group resource using Elastic Beanstalk.

## Example Usage

```hcl
resource "spotinst_elastigroup_aws_beanstalk" "elastigoup-aws=beanstalk" {

 name    = "example-elastigroup-beanstalk"
 region  = "us-west-2"
 product = "Linux/UNIX"

 min_size         = 0
 max_size         = 1
 desired_capacity = 0

 beanstalk_environment_name = "example-env"
 beanstalk_environment_id   = "e-example"
 instance_types_spot        = ["t2.micro", "t2.medium", "t2.large"]
}

deployment_preferences = {
    automatic_roll        = true
    batch_size_percentage = 100
    grace_period          = 90
    strategy = {
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