---
layout: "spotinst"
page_title: "Spotinst: Beanstalk Elastigroup"
sidebar_current: "docs-do-resource-beanstalk_elastigroup"
description: |-
  Provides a Spotinst Beanstalk Elastigroup group resource.
---

# spotinst\_beanstalk\_elastigroup

Provides a Spotinst AWS Beanstalk Elastigroup resource.

## Example Usage - Beanstalk Elastigroup

```hcl
# Create an elastigroup based on an existing beanstalk environment

resource "spotinst_beanstalk_elastigroup" "bseg1" {
  name                       = "elastigroup-beanstalk-import-example"
  region                     = "us-west-2"
  product                    = "Linux/UNIX"
  minimum                    = 1
  maximum                    = 30
  target                     = 2
  spot_instance_types        = ["m4.large","c4.large"]
  beanstalk_environment_name = "Your-Environment-Name-1"
}

```


## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the elastigroup.
* `region` - (Required) The Beanstalk environment region.
* `product` - (Optional) Operating system.
* `minimum` - (Required) Minimal amount of instances in the group when scaling.
* `maximum` - (Required) Maximal amount of instances in the group when scaling.
* `target` - (Required) Amount of instances to spin in the group.
* `spot_instance_types` - (Required) A list of strings. Each is a spot instance type name. See example above.
* `beanstalk_environment_name` - (Required) The Beanstalk environment name.


## Attributes Reference

The following attributes are exported:

* `id` - The elastigroup ID.
