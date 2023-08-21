---
layout: "spotinst"
page_title: "Spotinst: health_check"
subcategory: "Elastigroup"
description: |-
  Provides a Spotinst Health Check resource.
---

# spotinst\_administration\_org\_user

Provides a Spotinst User in the creator's organization

## Example Usage

```hcl 
resource "spotinst_administration_org_user" "terraform_user" {
  
  email = "abc@xyz.com"
  first_name = "test"
  last_name = "user"
  password = "testUser@123"
  role = "viewer"
  
}
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
