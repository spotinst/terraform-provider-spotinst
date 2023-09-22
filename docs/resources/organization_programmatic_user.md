---
layout: "spotinst"
page_title: "Spotinst: organization_programmatic_user"
subcategory: "Organization"
description: |-
  Provides a Spotinst programmatic user in the creator's organization.
---

# spotinst\_organization\_programmatic\_user

Provides a Spotinst programmatic user in the creator's organization.

## Example Usage

```hcl
resource "spotinst_organization_programmatic_user" "terraform_prog_user" {
  name = "test-prog-user"
  description = "creating programmatic user"
  policies {
    policy_id = "pol-g75d8c06"
    policy_account_ids = ["act-a1b2c3d4"]
  }
  /*accounts {
    account_id = "act-a1b2c3d4"
    account_role = "viewer"
  }*/  
  // account and policies are exclusive
}
// Update is not supported for this resource.
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the programmatic user.
* `description` - Brief description of the user.
* `policies` - All the policies the programmatic user will have access to.
   If used - Cannot be empty.
  * `policy_account_ids` - A list of the accounts that the policy should be
  enforced for the user.
  * `policy_id` - Policy ID the programmatic user will have access to.
* `accounts` - All the accounts the programmatic user will have access to.
   If used - Cannot be empty.
  * `account_id` - Account ID the programmatic user will have access to.
  * `account_role` - (Enum: `"viewer", "editor") Role to be associated with the
     programmatic user for this account.

## Attributes Reference

The following attributes are exported:

* `id` - The Spotinst Progammatic User ID.
