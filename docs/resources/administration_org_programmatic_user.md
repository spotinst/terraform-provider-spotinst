---
layout: "spotinst"
page_title: "Spotinst: health_check"
subcategory: "Elastigroup"
description: |-
  Provides a Spotinst Health Check resource.
---

# spotinst\_administration\_org\_programmatic\_user

Provides a Spotinst programmatic user in the creator's organization

## Example Usage

```hcl 
resource "spotinst_administration_org_programmatic_user" "terraform_prog_user" {
  name = "test-prog-user"
  description = "desc"
  policies{
    policy_id = "pol-g75d8c06"
    policy_account_ids = ["act-4c46c6df"]
  }
  /*accounts{
    account_id = "act-a1b2c3d4"
    account_role = "viewer"
  }*/  // account and policies are exclusive
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the programmatic user.
* `description` - Brief description of the user.
* `policies` - All the policies the programmatic user will have access to. If used - Cannot be empty.
    * `policy_account_ids` - All the policies the programmatic user will have access to. If used - Cannot be empty.
    * `policy_id` - All the policies the programmatic user will have access to. If used - Cannot be empty.
* `accounts` - All the accounts the programmatic user will have access to. If used - Cannot be empty.
    * `account_id` - Account ID the programmatic user will have access to.
    * `account_role` - (Enum: `"viewer", "editor") Role to be associated with the programmatic user for this account.


## Attributes Reference

The following attributes are exported:

* `id` - The Spotinst Progammatic User ID.
