---
layout: "spotinst"
page_title: "Spotinst: credentials_azure"
subcategory: "Accounts"
description: |-
  Provides a Spotinst credential Azure resource.
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
  expiration_date = "2025-12-31T23:59:00.000Z"
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The ID of the account associated with your token.
* `client_id` - (Required) Set the application ID.
* `client_secret` - (Required) Set the key secret.
* `tenant_id` - (Required) Set the directory ID.
* `subscription_id` - (Required) Set the subscription ID.
* `expiration_date` - (Required) Set the key secret expiration date.
