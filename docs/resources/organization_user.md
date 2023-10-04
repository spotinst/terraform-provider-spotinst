---
layout: "spotinst"
page_title: "Spotinst: organization_user"
subcategory: "Organization"
description: |-
  Provides a Spotinst User in the creator's organization.
---

# spotinst\_organization\_user

Provides a Spotinst User in the creator's organization.

## Example Usage

```hcl
resource "spotinst_organization_user" "terraform_user" {
  email = "abc@xyz.com"
  first_name = "test"
  last_name = "user"
  password = "testUser@123"
  role = "viewer"
  policies{
    policy_id = "pol-abcd1236"
    policy_account_ids = ["act-abcf4245"]
  }
  user_group_ids=["ugr-abcd1234","ugr-defg8763"]
}
```

## Argument Reference

The following arguments are supported:

* `email` - (Required) Email.
* `first_name` - (Required) The first name of the user.
* `last_name` - (Required) The last name of the user.
* `password` - (Optional) Password.
* `role` - (Optional) User's role.
* `policies` - (Optional) The policies to register under the given group
  (should be existing policies only).
    * `account_ids` - (Required) A list of accounts to register with the assigned under the
      given group (should be existing accounts only).
    * `policy_id` - (Required) A policy to register under the given group
      (should be existing policy only).
* `user_group_ids` - (Optional) A list of the user groups to register the given user to (should be existing user groups only)

## Attributes Reference

The following attributes are exported:

* `id` - The Spotinst User ID.
