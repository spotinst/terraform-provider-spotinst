---
layout: "spotinst"
page_title: "Spotinst: set_credential_aws"
subcategory: "Set Credential"
description: |-
  Provides a Spotinst set credential AWS resource.
---

# spotinst\_set\_credential\_aws

Provides a Spotinst set credential AWS resource.

## Example Usage

```hcl
# set credential AWS
resource "spotinst_set_credential_aws" "credential" {
  iamrole = "arn:aws:iam::1234567890:role/Spot_Iam_Role"
  accountid = "act-123456"
}
```

## Argument Reference

The following arguments are supported:

* `iamrole` - (Required) Provide the IAM Role ARN connected to another AWS account 922761411349 and with the latest Spot Policy - https://docs.spot.io/administration/api/spot-policy-in-aws
* `accountid` - (Required) The ID of the account associated with your token.
