---
layout: "spotinst"
page_title: "Spotinst: organization_policy"
subcategory: "Organization"
description: |-
  Provides a Spotinst access policy.
---

# spotinst\_organization\_policy

Provides a Spotinst access policy.

## Example Usage

```hcl
resource "spotinst_organization_policy" "terraform_policy" {
  name = "test-policy"
  description = "policy by terraform"
  policy_content {
    statements {
      actions = ["ocean:deleteCluster"]
      effect = "ALLOW"
      resources = ["*"]
    }
    statements {
      actions = ["ocean:updateCluster"]
      effect = "DENY"
      resources = ["*"]
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the Policy.
* `description` - (Optional) Short description of policy.
* `policy_content` - (Required) Set permissions objects list.
  * `statements` - (Required) List of permissions statements.
    * `actions` - (Required) Set a list of required actions for this permissions statement.
    Full list of actions can be found in [https://docs.spot.io/account-user-management/user-management/access-policies-actions/](https://docs.spot.io/account-user-management/user-management/access-policies-actions/).
    * `effect` - (Required) Valid values "ALLOW", "DENY".
    * `resources` - (Required) Set a list of resources IDs. In order to include all resources in this statement - use "*".

## Attributes Reference

The following attributes are exported:

* `id` - The Spotinst Policy ID.
