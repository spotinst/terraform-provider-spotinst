---
layout: "spotinst"
page_title: "Spotinst: credentials_aws"
subcategory: "Accounts"
description: |-
  Provides a Spotinst credential AWS resource.
---

# spotinst\_credentials\_azure

Provides a Spotinst credential Azure resource.

## Example Usage

```hcl
# set credential Azure
resource "spotinst_credentials_azure" "credential" {
  account_id      = "act-123456"
  client_id       = "redacted"
  client_secret   = "redacted"
  tenant_id       = "redacted"
  subscription_id = "redacted"
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The ID of the account associated with your token.
* `client_id` - (Required) Set the application ID.
* `client_secret` - (Required) Set the key secret.
* `tenant_id` - (Required) Set the directory ID.
* `subscription_id` - (Required) Set the subscription ID.
