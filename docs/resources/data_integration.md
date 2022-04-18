---
layout: "spotinst"
page_title: "Spotinst: data_integration"
subcategory: "Data Integration"
description: |-
Manages an Data Integration resource.
---

# spotinst\_data\_integration

Provides a Spotinst Data Integration resource.

## Example Usage

```hcl
resource "spotinst_data_integration" "example" {
  name  = "foo"
  status = "enabled"
  s3 {
  	bucketName = "terraform-test-do-not-delete"
    subdir      = "terraform-test-data-integration"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name`- (Required) The name of the data integration.
* `status` - (Optional, only when update) Determines if this data integration is on or off. Valid values: `"enabled"`, `"disabled"`
* `s3` - (Required) When vendor value is s3, the following fields are included:
  * `bucketName` - (Required) The name of the bucket to use. Your spot IAM Role policy needs to include s3:putObject permissions for this bucket. Can't be null.
  * `subdir` - (Optional) The subdirectory in which your files will be stored within the bucket. Adds the prefix subdir/ to new objects' keys. Can't be null or contain '/'.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `id` - The Spotinst Data Integration ID.
