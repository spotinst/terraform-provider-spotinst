---
layout: "spotinst"
page_title: "Spotinst: elastigroup_azure"
subcategory: "Elastigroup"
description: |-
 Provides a Spotinst elastigroup resource for Microsoft Azure.
---

# spotinst\_elastigroup\_azure\_v3

Provides a Spotinst elastigroup Azure resource.

## Example Usage

```hcl
resource "spotinst_elastigroup_azure_v3" "test_azure_group" {
  name                = "example_elastigroup_azure"
  resource_group_name = "spotinst-azure"
  region              = "eastus"
  os                  = "Linux"

  // --- CAPACITY ------------------------------------------------------
  min_size         = 0
  max_size         = 1
  desired_capacity = 1
  // -------------------------------------------------------------------

  // --- INSTANCE TYPES ------------------------------------------------
  od_sizes   = ["standard_a1_v1", "standard_a1_v2"]
  spot_sizes = ["standard_a1_v1", "standard_a1_v2"]
  // -------------------------------------------------------------------

  // --- LAUNCH SPEC ---------------------------------------------------
  custom_data = "IyEvYmluL2Jhc2gKZWNobyAidGVzdCI="
  
  managed_service_identity {
  resource_group_name = "MC_ocean-westus-dev_ocean-westus-dev-aks_westus"
  name                = "ocean-westus-dev-aks-agentpool"
  }
  
  // --- IMAGE ---------------------------------------------------------
  image {
    marketplace {
      publisher = "Canonical"
      offer     = "UbuntuServer"
      sku       = "18.04-LTS"
      version   = "latest"
    }
  }
  // -------------------------------------------------------------------

  // --- STRATEGY ------------------------------------------------------
  //on_demand_count     = 1
  spot_percentage       = 65
  draining_timeout      = 300
  fallback_to_on_demand = true
  // -------------------------------------------------------------------

  // --- NETWORK -------------------------------------------------------
  network {
    virtual_network_name = "VirtualNetworkName"
    resource_group_name  = "ResourceGroup"

    network_interfaces {
      subnet_name      = "default"
      assign_public_ip = false
      is_primary       = true

      additional_ip_configs {
        name             = "SecondaryIPConfig"
        PrivateIPVersion = "IPv4"
      }

      application_security_group {
        name                = "ApplicationSecurityGroupName"
        resource_group_name = "ResourceGroup"
      }
    }
  }
  // -------------------------------------------------------------------

  // --- LOGIN ---------------------------------------------------------
  login {
    user_name      = "admin"
    ssh_public_key = "33a2s1f3g5a1df5g1ad3f2g1adfg56dfg=="
  }
  // -------------------------------------------------------------------
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The group name.
* `region` - (Required) The region your Azure group will be created in.
* `resource_group_name` - (Required) Name of the Resource Group for Elastigroup.
* `os` - (Required) Type of the operating system. Valid values: `"Linux"`, `"Windows"`.
* `max_size` - (Required) The maximum number of instances the group should have at any time.
* `min_size` - (Required) The minimum number of instances the group should have at any time.
* `desired_capacity` - (Required) The desired number of instances the group should have at any time.
* `custom_data` - (Optional) Custom init script file or text in Base64 encoded format.
* `managed_service_identity` - (Optional) List of Managed Service Identity objects.
    * `resource_group_name` - (Required) Name of the Azure Resource Group where the Managed Service Identity is located.
    * `name` - (Required) Name of the Managed Service Identity.
  
* `od_sizes` - (Required) Available On-Demand sizes
* `spot_sizes` - (Required) Available Low-Priority sizes.

<a id="strategy"></a>
## Strategy
* `spot_percentage` - (Optional) Percentage of Spot-VMs to maintain. Required if `on_demand_count` is not specified.
* `on_demand_count` - (Optional) Number of On-Demand VMs to maintain. Required if `spot_percentage` is not specified.
* `fallback_to_on_demand` - 
* `draining_timeout` - (Optional, Default `120`) Time (seconds) to allow the instance to be drained from incoming TCP connections and detached from MLB before terminating it during a scale-down operation.

<a id="image"></a>
## Image

* `image` - (Required) Image of a VM. An image is a template for creating new VMs. Choose from Azure image catalogue (marketplace) or use a custom image.
    * `publisher` - (Optional) Image publisher. Required if resource_group_name is not specified.
    * `offer` - (Optional) Name of the image to use. Required if publisher is specified.
    * `sku` - (Optional) Image's Stock Keeping Unit, which is the specific version of the image. Required if publisher is specified.
    * `version` - 
    * `resource_group_name` - (Optional) Name of Resource Group for custom image. Required if publisher not specified.
    * `image_name` - (Optional) Name of the custom image. Required if resource_group_name is specified.

```hcl
  // market image
  image {
    marketplace {
      publisher = "Canonical"
      offer     = "UbuntuServer"
      sku       = "18.04-LTS"
      version   = "latest"
    }
  }
  
  // custom image
  image {
    custom {
      image_name          = "customImage"
      resource_group_name = "resourceGroup"
    }
  } 
```

<a id="network"></a>
## Network

* `network` - (Required) Defines the Virtual Network and Subnet for your Elastigroup.
    * `virtual_network_name` - (Required) Name of Vnet.
    * `resource_group_name` - (Required) Vnet Resource Group Name.
    * `network_interfaces` - 
        * `subnet_name` - (Required) ID of subnet.
        * `assign_public_up` - (Optional, Default: `false`) Assign a public IP to each VM in the Elastigroup.
        * `is_primary` - 
        * `additional_ip_configs` - (Optional) Array of additional IP configuration objects.
            * `name` - (Required) The IP configuration name.
            * `private_ip_version` - (Optional) Available from Azure Api-Version 2017-03-30 onwards, it represents whether the specific ip configuration is IPv4 or IPv6. Valid values: `IPv4`, `IPv6`.
        * `application_security_group` - (Optional) - List of Application Security Groups that will be associated to the primary ip configuration of the network interface.
            * `name` - (Required) - The name of the Application Security group.
            * `resource_group_name` - (Required) - The resource group of the Application Security Group.
      }
```hcl
  network {
    virtual_network_name = "VirtualNetworkName"
    resource_group_name  = "ResourceGroup"

    network_interfaces {
      subnet_name      = "default"
      assign_public_ip = false
      is_primary       = true

      additional_ip_configs {
        name             = "SecondaryIPConfig"
        PrivateIPVersion = "IPv4"
      }

      application_security_group {
        name                = "ApplicationSecurityGroupName"
        resource_group_name = "ResourceGroup"
      }
    }
  }
```

<a id="login"></a>
## Login

* `login` - (Required) Describes the login configuration.
    * `user_name` - (Required) Set admin access for accessing your VMs.
    * `ssh_public_key` - (Optional) SSH for admin access to Linux VMs. Required for Linux OS types.
    * `password` - (Optional) Password for admin access to Windows VMs. Required for Windows OS types.

```hcl
  login {
    user_name      = "admin"
    ssh_public_key = "33a2s1f3g5a1df5g1ad21651sag56dfg=="
  }
```




    
