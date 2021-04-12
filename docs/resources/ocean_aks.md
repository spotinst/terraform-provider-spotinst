---
layout: "spotinst"
page_title: "Spotinst: ocean_aks"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean resource using AKS.
---

# spotinst\_ocean\_aks

Manages a Spotinst Ocean AKS resource.

## Example Usage

```hcl
resource "spotinst_ocean_aks" "example" {
  name = "ocean_aks"
  controller_cluster_id = "controller-cluster-id"

  // --- AKS -----------------------------------------------------------
  acd_identifier = "acd-12345"
  aks_name = "AKSName"
  aks_resource_group_name = "ResourceGroupName"
  // -------------------------------------------------------------------

  // --- LOGIN ---------------------------------------------------------
  ssh_public_key = "ssh-rsa [... redacted ...] generated-by-azure"
  user_name = "some-name"
  // -------------------------------------------------------------------
}
```

```
output "ocean_id" {
  value = spotinst_ocean_aks_.example.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The Ocean cluster name.
* `controller_cluster_id` - (Required) The Ocean controller cluster. 
* `aks_name` - (Required) The AKS cluster name.
* `acd_identifier` - (Required) The AKS identifier.
* `aks_resource_group_name` - (Required) Name of the Resource Group for AKS cluster. 
* `ssh_public_key` - (Required) SSH public key for admin access to Linux VMs.
* `user_name` - (Optional) Username for admin access to VMs.
