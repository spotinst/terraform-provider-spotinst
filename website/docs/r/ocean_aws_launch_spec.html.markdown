---
layout: "spotinst"
page_title: "Spotinst: ocean_aws_launch_spec"
sidebar_current: "docs-do-resource-ocean_aws_launch_spec"
description: |-
  Provides a Spotinst Ocean Launch Spec resource using AWS.
---

# spotinst\_ocean\_aws\_launch\_spec

Provides a custom Spotinst Ocean AWS Launch Spec resource.

## Example Usage

```hcl
resource "spotinst_ocean_aws_launch_spec" "example" {
  ocean_id  = "o-123456"
  image_id  = "ami-123456"
  user_data = "echo hello world"
  iam_instance_profile = "iam-profile"
  security_group_ids = ["awseb-12345"]

  labels {
    key   = "fakeKey"
    value = "fakeValue"
  }
  
  taints {
    key    = "taint key updated"
    value  = "taint value updated"
    effect = "NoExecute"
  }
}
```

## Argument Reference

The following arguments are supported:

* `ocean_id` - (Required) The ocean cluster you wish to 
* `user_data` - (Optional) Base64-encoded MIME user data to make available to the instances.
* `image_id` - (Optional) ID of the image used to launch the instances.
* `iam_instance_profile` - (Optional) The ARN or name of an IAM instance profile to associate with launched instances.
* `security_groups` - (Optional) Optionally adds security group IDs.

* `labels` - (Optional) Optionally adds labels to instances launched in an Ocean cluster.
    * `key` - (Required) The tag key.
    * `value` - (Required) The tag value.
    
* `taints` - (Optional) Optionally adds labels to instances launched in an Ocean cluster.
    * `key` - (Required) The tag key.
    * `value` - (Required) The tag value.
    * `effect` - (Required) The effect of the taint. Valid values: `"NoSchedule"`, `"PreferNoSchedule"`, `"NoExecute"`.

