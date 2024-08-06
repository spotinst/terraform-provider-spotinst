---
layout: "spotinst"
page_title: "Spotinst: account"
subcategory: "Accounts"
description: |-
  Create a Spotinst account resource.
---

# spotinst\_account

Provides a Spotinst account resource.

## Example Usage

```hcl
# Create a Account
resource "spotinst_account" "my_acct" {
  name="my_acct"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Provide a name for your account. The account name must contain at least one character that is a-z or A-Z.

## Attributes Reference

The following attributes are exported:

* `id` - The account ID.
