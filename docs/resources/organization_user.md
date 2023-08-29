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
  
}
// Update is not supported for this resource.
```

## Argument Reference

The following arguments are supported:

* `email` - (Required) Email.
* `first_name` - (Required) The first name of the user.
* `last_name` - (Required) The last name of the user.
* `password` - (Required) Password.
* `role` - (Optional) User's role.

## Attributes Reference

The following attributes are exported:

* `id` - The Spotinst User ID.
