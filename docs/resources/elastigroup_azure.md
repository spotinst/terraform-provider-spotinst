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
  description         = "Azure Elastigroup Resource through TF"
  resource_group_name = "spotinst-azure"
  region              = "eastus"
  os                  = "Linux"
  zones               = ["1", "2", "3"]
  preferred_zones     = ["1", "3"]

  // --- CAPACITY ------------------------------------------------------
  min_size         = 0
  max_size         = 1
  desired_capacity = 1
  // -------------------------------------------------------------------

  // --- INSTANCE TYPES ------------------------------------------------
   vm_sizes {
       od_sizes   = ["standard_a1_v1","standard_a1_v2"]
       spot_sizes = ["standard_a1_v1","standard_a1_v2"]
       preferred_spot_sizes = ["standard_a1_v2"]
   }
  // -------------------------------------------------------------------

  // --- LAUNCH SPEC ---------------------------------------------------
  custom_data = "IyEvYmluL2Jhc2gKZWNobyAidGVzdCI="
  shutdown_script = "IlRlc3RpbmcgRUci"
  //user_data = "IlRlc3RpbmcgRUci"
  
  vm_name_prefix = "prefixName"
  
  managed_service_identity {
  resource_group_name = "MC_ocean-westus-dev_ocean-westus-dev-aks_westus"
  name                = "ocean-westus-dev-aks-agentpool"
  }
  
  tags {
  key = "key1"
  value = "value1"
  }
  
  tags {
  key = "key2"
  value = "value2"
  }
  
  os_disk {
    size_gb = 32
    type = "Premium_LRS"
  }

  data_disk {
    size_gb = 8
    type = "Premium_LRS"
    lun = 2
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
  optimization_windows    = ["Mon:19:46-Tue:20:46"]
  availability_vs_cost    = 100
  revert_to_spot {
    perform_at            = "timeWindow"
  }
  signal {
    type    = "vmReadyToShutdown"
    timeout = 60
  }
  capacity_reservation {
    should_utilize       = true
    utilization_strategy = "utilizeOverOD"
    capacity_reservation_groups {
      crg_name                = "crg name"
      crg_resource_group_name = "resourceGroupName"
      crg_should_prioritize   = true
    }
  }
  // -------------------------------------------------------------------

  // --- NETWORK -------------------------------------------------------
  network {
    virtual_network_name = "VirtualNetworkName"
    resource_group_name  = "ResourceGroup"

    network_interfaces {
      subnet_name      = "default"
      assign_public_ip = false
      is_primary       = true
      public_ip_sku = "Standard"
      enable_ip_forwarding = true

      additional_ip_configs {
        name             = "SecondaryIPConfig"
        PrivateIPVersion = "IPv4"
      }

      application_security_group {
        name                = "ApplicationSecurityGroupName"
        resource_group_name = "ResourceGroup"
      }
      
       public_ips {
        name                = "PublicIpName"
        resource_group_name = "ResourceGroup"
      }

      security_group {
        name                = "NetworkSecurityGroupName"
        resource_group_name = "ResourceGroup"
      }
    }
  }
  // -------------------------------------------------------------------
  
  proximity_placement_groups {
    name                = "TestProximityPlacementGroup"
    resource_group_name = "ResourceGroup"
  }

  boot_diagnostics {
    is_enabled      = true
    storage_url     = "https://.blob.core.windows.net"
    type            = "unmanaged"
  }

  secret {
    source_vault {
      name                = "TestVault"
      resource_group_name = "ResourceGroup"
    }

    vault_certificates {
      certificate_url     = "string"
      certificate_store   = "string"
    }
  }

  security {
    security_type = "Standard"
    secure_boot_enabled = false
    vtpm_enabled = false
    confidential_os_disk_encryption = false
  }

  // --- LOGIN ---------------------------------------------------------
  login {
    user_name      = "admin"
    ssh_public_key = "33a2s1f3g5a1df5g1ad3f2g1adfg56dfg=="
  }
  // -------------------------------------------------------------------
  
  // --- HEALTH --------------------------------------------------------
  health {
    health_check_types = ["applicationGateway"]
    unhealthy_duration = 240
    grace_period       = 420
    auto_healing       = false
  }
  // -------------------------------------------------------------------
  
  // --- SCHEDULING ----------------------------------------------------
  scheduling_task {
    is_enabled      = true
    type            = "scale"
    cron_expression = "52 10 * * *"
    scale_max_capacity = 8
    scale_min_capacity = 0
    scale_target_capacity = 2
  }

  scheduling_task {
    is_enabled      = true
    type            = "scaleUp"
    cron_expression = "52 11 * * *"
    adjustment = 1
  }
  // -------------------------------------------------------------------
  
  // --- LOAD BALANCER -------------------------------------------------
  load_balancer {
    type                = "loadBalancer"
    resource_group_name = "AutomationResourceGroup"
    name                = "Automation-Lb"
    sku                 = "Standard"
    backend_pool_names  = ["Automation-Lb-BackendPool"]
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
  
* `tags` - (Optional) Key-Value pairs for VMs in the Elastigroup.
    * `key` - (Required) Tag Key for Vms in Elastigroup.
    * `value` - (Required) Tag Value for Vms in Elastigroup.
* `vm_sizes` - (Required) Sizes of On-Demand and Low-Priority VMs.
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
  * `marketplace` - (Optional) Select an image from Azure's Marketplace image catalogue. Cannot be used with `custom` or `gallery`
      * `publisher` - (Optional) Image publisher. Required if marketplace image is specified.
      * `offer` - (Optional) Name of the image to use. Required if marketplace image is specified.
      * `sku` - (Optional) Image's Stock Keeping Unit, which is the specific version of the image. Required if marketplace image is specified.
      * `version` - Image's version. if version not provided we use `latest`. Required if marketplace image is specified.
  * `custom` - (Optional) Custom image to launch Elastigroup with. Cannot be used with `marketplace` or `gallery`.
      * `resource_group_name` - (Optional) Name of Resource Group for custom image. Required if custom image is specified.
      * `image_name` - (Optional) Name of the custom image. Required if custom image is specified.
  * `gallery_image` - (Optional) Gallery image to launch Elastigroup with. Cannot be used with `marketplace` or `custom`.
      * `gallery_name` - (Optional) Name of the gallery. Required if gallery image is specified.
      * `image_name` - (Optional) Name of the gallery image. Required if gallery image is specified.
      * `resource_group_name` - (Optional) Name of Resource Group for gallery image. Required if gallery image is specified.
      * `version` - (Optional) Image's version. Can be in the format x.x.x or 'latest'. Required if gallery image is specified.
      * `spot_account_id` - (Optional) The Spot account ID that connected to the Azure subscription to which the gallery belongs. Relevant only in case of cross-subscription shared galleries. Read more (https://docs.spot.io/elastigroup/features-azure/shared-image-galleries) about cross-subscription shared galleries in Elastigroup.

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
  
  // gallery image
  image {
    gallery_image {
      resource_group_name = "resourceGroup"
      gallery_name = "galleryName"
      image_name = "imageName"
      version = "1.0.0"
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

<a id="extensions"></a>
## Extensions

* `extensions` - (Optional) An object for an Azure extensions.
    * `name` - (Required) Name of the extension.
    * `type` - (Required) Type of the extension.
    * `publisher` - (Required) Publisher of an extension.
    * `api_version` - (Required) The API version of the extension. Required if extension specified.
    * `minor_version_auto_upgrade` - (Required) Enable minor version upgrades of the extension. Required if extension specified.
    * `protected_settings` - (Optional) Object for protected settings. This must not exist simultaneously with `protected_settings_from_key_vault`
    * `public_settings` - (Optional) Object for public settings.
    * `enable_automatic_upgrade` - (Optional) Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available.
    * `protected_settings_from_key_vault` - (Optional) The extensions protected settings that are passed by reference, and consumed from key vault.
      * `secret_url` - (Required) The URL referencing a secret in a Key Vault.
      * `sourcevault` - (Required) The relative URL of the Key Vault containing the secret.

```hcl
  extensions {
    name                       = "extensionName"
    type                       = "customScript"
    publisher                  = "Microsoft.Azure.Extensions"
    api_version                = "2.0"
    minor_version_auto_upgrade = true
    enable_automatic_upgrade   = false
    protected_settings         = {
      "script" : "IyEvYmluL2Jhc2gKZWNobyAibmlyIC9ob29uaXIudHh0Cg=="
    }
    public_settings = {
      "fileUris": "https://testspot/Azuretest.sh"
    }
    
    protected_settings_from_key_vault {
      source_vault = "/subscriptions/abcde-123490-566778/resourceGroups/resourceGroupName/providers/Microsoft.KeyVault/vaults/keyVaultName"
      secret_url = "https://terraform-extension-test.vault.azure.net/secrets/TestSeccret"
    }
  }
```  

<a id="scaling-policy"></a>
## Scaling Policies

`scaling_up_policy` / `scaling_down_policy` supports the following:

* `policy_name` - (Required) The name of the policy.
* `metric_name` - (Required) Metric to monitor by Azure metric display name.
* `statistic` - (Required) Statistic by which to evaluate the selected metric. Valid Values: `"average"`, `"sampleCount"`, `"sum"`, `"minimum"`, `"maximum"`
* `unit` - (Optional) Unit to measure to evaluate the selected metric: Valid Values: `"percent`, `"seconds"`, `"milliseconds"`, `"bytes"`, `"countPerSecond"`, `"bytesPerSecond"`, `"seconds"`
* `threshold` - (Required) The value at which the scaling action is triggered.
* `namespace` - (Required) The namespace for the alarm's associated metric. Select one of the next namespaces presented in Azure configurator - [Namespace](https://learn.microsoft.com/en-us/azure/templates/)
* `is_enabled` - (Optional, Default: `true`) Specifies whether the scaling policy described in this block is enabled.
* `period` - (Required) Amount of time (seconds) for which the threshold must be met in order to trigger the scaling action.
* `evaluation_periods` - (Required) Amount of time (seconds) for which the threshold must be met in order to trigger the scaling action.
* `cooldown` - (Required) Time (seconds) to wait after a scaling action before resuming monitoring.
* `dimensions` - (Optional) A list of dimensions describing qualities of the metric. Required if scaling.up.namespace is different from 'Microsoft.Compute'
    * `name` - (Optional) Azure resource group for the scaling.down.dimensions. Required if using namespace different from "Microsoft.Compute".
    * `value` - (Optional) Azure resource the scaling.down.dimensions. Required if using namespace different from "Microsoft.Compute".
* `operator` - (Required) The operator used to evaluate the threshold against the current metric value. Valid values: `"gt"`, `"gte"`, `"lt"`, `"lte"`.
* `source` - (Optional) The source of the metric.
* `action` - (Required) Scaling action to take when the policy is triggered.
    * `type` - (Required) Type of scaling action to take when the scaling policy is triggered. Valid Values: `"adjustment"`, `"updateCapacity"`
    * `adjustment` - (Optional)  Value to which the action type will be adjusted. Required if using the next action types: `"numeric"`, `"percentageAdjustment"`
    * `maximum` - (Optional)  Upper limit of instances that you can scale down to. Also you must indicate “minimum” and “target” amounts. Required if selected as action type: `"updateCapacity"`
    * `minimum` - (Optional)  Lower limit of instances that you can scale down to. Also you must indicate “target” and “maximum” amounts. Required if selected as action type: `"updateCapacity"`
    * `target` - (Optional)  Desired number of instances. Also you must indicate “minimum” and “maximum” amounts. Required if selected as action type: `"updateCapacity"`

Usage:

```hcl
  //--- SCALING --------------------------------------------------------
   scaling_up_policy {
    policy_name        = "Scaling Up Policy"
    metric_name        = "Percentage CPU"
    statistic          = "average"
    unit               = "count"
    namespace          = "Microsoft.Network/applicationGateways"
    threshold          = 1.5
    period             = 60
    evaluation_periods = 5
    cooldown           = 300
    operator           = "gt"
    is_enabled         = false
    dimensions {
      name  = "name-1"
      value = "value-1"
    }
    action {
      type       = "updateCapacity"
      minimum    = "1"
      maximum    = "6"
      target     = "2"
    }
  }

  scaling_down_policy {
    policy_name        = "Scaling Down Policy"
    metric_name        = "Disk Read Bytes"
    statistic          = "average"
    unit               = "bytes"
    namespace          = "Microsoft.Compute"
    threshold          = 5
    operator           = "gt"
    period             = 60
    evaluation_periods = 10
    cooldown           = 300
    is_enabled         = true
    dimensions {
      name  = "name-1"
      value = "value-1"
    }
    action {
      type       = "adjustment"
      adjustment = "2"
    }
  }
```




    
