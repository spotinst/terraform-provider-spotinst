---
layout: "spotinst"
page_title: "Spotinst: health_check"
subcategory: "Elastigroup"
description: |-
  Provides a Spotinst Health Check resource.
---

# spotinst\_administration\_org\_policy

Provides a Spotinst access policy 

## Example Usage

```hcl 
resource "spotinst_administration_org_policy" "terraform_policy" {
  name = "test-policy"
  description = "policy by terraform"
  policy_content{
    statements{
      actions = ["ocean:deleteCluster"]
      effect = "ALLOW"
      resources = ["*"]
    }
    statements{
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
* `description` - Short description of policy.
* `policy_content` - Set permissions objects list.
* `statements` - List of permissions statements.
    * `actions` - Set a list of required actions for this permissions statement.
    * `effect` - Valid values "ALLOW", "DENY".
    * `resources` - Set a list of resources IDs.

## Attributes Reference

The following attributes are exported:

* `id` - The Spotinst Policy ID.
