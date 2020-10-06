---
layout: "spotinst"
page_title: "Spotinst: elastigroup_aws_suspension"
subcategory: "Elastigroup"
description: |-
  Provides a processes suspensions to Spotinst AWS group resources.
---

# spotinst\_elastigroup\_aws\_suspension

Suspend AWS Elastigroup processes. This resource provide the capavility of
suspending elastigroup processes using Terraform.

For supported processes please visit: [Suspend Processes API reference](https://help.spot.io/spotinst-api/elastigroup/amazon-web-services/suspend-processes/)
## Example Usage

```hcl
# Create a process suspension for Elastigroup
resource "spotinst_elastigroup_aws_suspension" "resource_name" {

  group_id = "sig-12345678"
  suspension {
        name = "OUT_OF_STRATEGY"
    }
  suspension {
          name = "REVERT_PREFERRED"
    }
  suspension  {
          name = "PREVENTIVE_REPLACEMENT"
    }
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required; string) Elastigroup ID to apply the suspensions on.
* `suspension` - (Required; at least one block is required) block of single process to suspend.
    * `name` - (Required; string) The name of process to suspend. Valid values: `"AUTO_HEALING" , "OUT_OF_STRATEGY", "PREVENTIVE_REPLACEMENT", "REVERT_PREFERRED", or "SCHEDULING"`. 
