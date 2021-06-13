---
layout: "spotinst"
page_title: "Spotinst: ocean_aks"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean resource using AKS.
---

# spotinst\_ocean\_aks

Manages a Spotinst Ocean AKS resource.

-> This resource contains arguments (such as `image` and `extension`) that are automatically populated from the data reported by the Ocean AKS Connector deployed into your cluster. You can override the upstream configuration by defining the corresponding arguments.

## Prerequisites

Installation of the Ocean controller is required by this resource. You can accomplish this by using the [spotinst/ocean-controller](https://registry.terraform.io/modules/spotinst/ocean-controller/spotinst) module as follows:

```hcl
module "ocean-controller" {
  source = "spotinst/ocean-controller/spotinst"

  # Credentials.
  spotinst_token   = "redacted"
  spotinst_account = "redacted"

  # Configuration.
  cluster_identifier = "ocean-westus-dev-aks"
  acd_identifier     = "acd-12345"
}
```

~> You must configure the same `cluster_identifier` and `acd_identifier` both for the Ocean controller and for the `spotinst_ocean_aks` resource.

To learn more about how to integrate existing Kubernetes clusters into Ocean using Terraform, watch [this video](https://youtu.be/ffGmMlpPsPE).

## Example Usage

```hcl
resource "spotinst_ocean_aks" "example" {
  name                  = "ocean-westus-dev-aks"
  controller_cluster_id = "ocean-westus-dev-aks"

  // --- AKS -----------------------------------------------------------
  acd_identifier          = "acd-12345"
  aks_name                = "ocean-westus-dev-aks"
  aks_resource_group_name = "ocean-westus-dev"
  // -------------------------------------------------------------------

  // --- Login ---------------------------------------------------------
  ssh_public_key = "ssh-rsa [... redacted ...] generated-by-azure"
  user_name      = "some-name"
  // -------------------------------------------------------------------

  // --- Launch Specification ------------------------------------------
  resource_group_name = "some-resource-group-name"
  custom_data         = "[... redacted ...]"

  tag {
    key   = "Environment"
    value = "Dev"
  }
  // --------------------------------------------------------------------

  // --- VMSizes --------------------------------------------------------
  vm_sizes {
    whitelist = [
      "standard_ds2_v2",
    ]
  }
  // --------------------------------------------------------------------

  // --- OSDisk --------------------------------------------------------
  os_disk {
    size_gb = 130
    type    = "Standard_LRS"
  }
  // -------------------------------------------------------------------

  // --- Image ---------------------------------------------------------
  image {
    marketplace {
      publisher = "Canonical"
      offer     = "UbuntuServer"
      sku       = "18.04-LTS"
      version   = "latest"
    }
  }
  // ---------------------------------------------------------------------

  // --- Strategy --------------------------------------------------------
  strategy {
    fallback_to_ondemand = true
    spot_percentage      = 40
  }
  // ---------------------------------------------------------------------

  // --- Health ----------------------------------------------------------
  health {
    grace_period = 10
  }
  // ---------------------------------------------------------------------

  // --- NETWORK ---------------------------------------------------------
  network {
    virtual_network_name = "vn-name"
    resource_group_name  = "ocean-westus-dev"

    network_interface {
      subnet_name      = "subnet-name"
      assign_public_ip = false
      is_primary       = false

      additional_ip_config {
        name               = "ip-config-name"
        private_ip_version = "ipv4"
      }

    }
  }
  // ----------------------------------------------------------------------

  // --- Extensions -------------------------------------------------------
  extension {
    api_version                = "1.0"
    minor_version_auto_upgrade = true
    name                       = "extension-name"
    publisher                  = "Microsoft.Azure.Extensions"
    type                       = "Linux"
  }
  // ----------------------------------------------------------------------

  // --- Load Balancers ---------------------------------------------------
  load_balancer {
    backend_pool_names = [
      "terraform-backend-pool"
    ]
    load_balancer_sku   = "Standard"
    name                = "load-balancer-name"
    resource_group_name = "resource-group-name"
    type                = "loadBalancer"
  }

  // ----------------------------------------------------------------------

  // --- Auto Scaler ------------------------------------------------------
  autoscaler {
    autoscale_is_enabled = true

    autoscale_down {
      max_scale_down_percentage = 10
    }

    resource_limits {
      max_vcpu       = 1024
      max_memory_gib = 40
    }

    autoscale_headroom {
      automatic {
        is_enabled = true
        percentage = 10
      }
    }
  }
  // ----------------------------------------------------------------------
}
```

```
output "ocean_id" {
  value = spotinst_ocean_aks.example.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The Ocean cluster name.
* `controller_cluster_id` - (Required) A unique identifier used for connecting the Ocean SaaS platform and the Kubernetes cluster. Typically, the cluster name is used as its identifier. 
* `aks_name` - (Required) The AKS cluster name.
* `acd_identifier` - (Required) The AKS identifier.
* `aks_resource_group_name` - (Required) Name of the Azure Resource Group where the AKS cluster is located. 
* `ssh_public_key` - (Required) SSH public key for admin access to Linux VMs.
* `user_name` - (Optional) Username for admin access to VMs.
* `resource_group_name` - (Optional) Name of the Azure Resource Group into which VMs will be launched. Cannot be updated.
* `custom_data` - (Optional) Must contain a valid Base64 encoded string.
* `tag` - (Optional) Unique key-value pairs that will be used to tag VMs that are launched in the cluster.
    * `key` - (Optional) Tag key.
    * `value` - (Optional) Tag value.
* `network` - (Optional) Define the Virtual Network and Subnet.
    * `virtual_network_name` - (Optional) Virtual network.
    * `resource_group_name` - (Optional) Vnet resource group name.
    * `network_interface` - (Optional) A list of virtual network interfaces. The publicIpSku must be identical between all the network interfaces. One network interface must be set as the primary.
        * `subnet_name` - (Optional) Subnet name.
        * `assign_public_ip` - (Optional) Assign public IP.
        * `is_primary` - (Optional) Defines whether the network interface is primary or not.
        * `additional_ip_config` - (Optional) Additional configuration of network interface. The name fields between all the `additional_ip_config` must be unique.
            * `name` - (Required) Configuration name.
            * `private_ip_version` - (Optional, Default: `IPv4`) Supported values: `IPv4`, `IPv6`.
* `extension` - (Optional) List of Azure extension objects.
    * `api_version` - (Optional) API version of the extension.
    * `minor_version_auto_upgrade` - (Optional) Toggles whether auto upgrades are allowed.
    * `name` - (Optional) Extension name.
    * `type` - (Optional) Extension type.
* `os_disk` - (Optional) OS disk specifications.
    * `size_gb` - (Optional) The size of the OS disk in GB.
    * `type` - (Optional) The type of the OS disk. Supported values: `Standard_LRS`, `Premium_LRS`, `StandardSSD_LRS`.
* `load_balancer` - (Optional) Configure Load Balancer.
    * `backend_pool_names` - (Optional) Names of the Backend Pools to register the Cluster VMs to. Each Backend Pool is a separate load balancer.
    * `load_balancer_sku` - (Optional) Supported values: `Standard`, `Basic`.
    * `name` - (Optional) Name of the Load Balancer.
    * `resource_group_name` - (Optional) The Resource Group name of the Load Balancer.
    * `type` - (Optional) The type of load balancer. Supported value: `loadBalancer`
* `image` - (Optional) Image of VM. An image is a template for creating new VMs. Choose from Azure image catalogue (marketplace).
    * `marketplace` - (Optional) Select an image from Azure's Marketplace image catalogue.
        * `publisher` - (Optional) Image publisher.
        * `offer` - (Optional) Image name.
        * `sku` - (Optional) Image Stock Keeping Unit (which is the specific version of the image).
        * `version` - (Optional, Default: `latest`) Image version.
* `vm_sizes` - (Optional) The types of virtual machines that may or may not be a part of the Ocean cluster.
    * `whitelist` - (Optional) VM types allowed in the Ocean cluster.
* `strategy` - (Optional) The Ocean AKS strategy object.
    * `fallback_to_ondemand` - (Optional) If no spot instance markets are available, enable Ocean to launch on-demand instances instead.
    * `spot_percentage` - (Optional) Percentage of Spot VMs to maintain.
* `health` - (Optional) The Ocean AKS Health object.
    * `grace_period` - (Optional, Default: `600`) The amount of time to wait, in seconds, from the moment the instance has launched before monitoring its health checks.
* `autoscaler` - (Optional) The Ocean Kubernetes Autoscaler object.
    * `autoscale_is_enabled` - (Optional) Enable the Ocean Kubernetes Autoscaler.
    * `autoscale_down` - (Optional) Auto Scaling scale down operations.
        * `max_scale_down_percentage` - (Optional) Would represent the maximum % to scale-down.
    * `resource_limits` - (Optional) Optionally set upper and lower bounds on the resource usage of the cluster.
        * `max_vcpu` - (Optional) The maximum cpu in vCpu units that can be allocated to the cluster.
        * `max_memory_gib` - (Optional) The maximum memory in GiB units that can be allocated to the cluster.
    * `autoscale_headroom` - (Optional) Spare Resource Capacity Management feature enables fast assignment of Pods without having to wait for new resources to be launched.
        * `automatic` - (Optional) Automatic headroom configuration.
            * `is_enabled` - (Optional) Enable automatic headroom. When set to `true`, Ocean configures and optimizes headroom automatically.
            * `percentage` - (Optional) Optionally set a number between 0-100 to control the percentage of total cluster resources dedicated to headroom. Relevant when `isEnabled` is toggled on.
