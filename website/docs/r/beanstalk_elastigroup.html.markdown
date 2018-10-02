---
layout: "spotinst"
page_title: "Spotinst: beanstalk_elastigroup"
sidebar_current: "docs-do-resource-beanstalk_elastigroup"
description: |-
 Provides a Spotinst AWS group resource using Elastic Beanstalk.
---

# spotinst\_beanstalk\_elastigroup

Provides a Spotinst AWS group resource using Elastic Beanstalk.

## Example Usage

```hcl
resource "spotinst_beanstalk_elastigroup" "beanstalk-elastigoup" {

 name    = "beanstalk-elastigroup"
 region  = "us-west-2"
 product = "Linux/UNIX"

 min_size         = 0
 max_size         = 1
 desired_capacity = 0

 beanstalk_environment_name = "example-env"
 instance_types_spot        = ["t2.micro", "t2.medium", "t2.large"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The group name.
* `region` - (Required) The AWS region your group will be created in. Cannot be changed after the group has been created.
* `description` - (Required) The group description.
* `product` - (Required) Operation system type. Valid values: `"Linux/UNIX"`, `"SUSE Linux"`, `"Windows"`.
For EC2 Classic instances:  `"Linux/UNIX (Amazon VPC)"`, `"SUSE Linux (Amazon VPC)"`, `"Windows (Amazon VPC)"`.   

* `max_size` - (Required) The maximum number of instances the group should have at any time.
* `min_size` - (Required) The minimum number of instances the group should have at any time.
* `desired_capacity` - (Required) The desired number of instances the group should have at any time.

* `beanstalk_environment_name` - (Required) The name of an existing Beanstalk environment.
* `instance_types_spot` - (Required) One or more instance types.

