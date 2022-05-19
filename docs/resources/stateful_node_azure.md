---
layout: "spotinst"
page_title: "Spotinst: stateful_node_azure"
subcategory: "Stateful"
description: |-
Provides a Spotinst Stateful Node resource using Azure.
---

# spotinst\_stateful\_node\_azure

Provides a Spotinst stateful node Azure resource.

## Example Usage

```hcl
resource "spotinst_stateful_node_azure" "test_stateful_node_azure" {
  name                = "example_stateful_node_azure"
  region              = "eastus"
  resource_group_name = "spotinst-azure"
  description         = "example_stateful_node_azure_description"

  // --- STRATEGY ------------------------------------------------------
  strategy {
    draining_timeout        = 30
    fallback_to_on_demand   = true
    optimization_windows    = ["Tue:19:46-Tue:20:46"]
    revert_to_spot {
      perform_at            = "timeWindow"
    }
    preferred_life_cycle    = "od" 
  }
  // -------------------------------------------------------------------

  // --- COMPUTE -------------------------------------------------------
  os                   = "Linux"
  od_sizes             = ["standard_ds1_v2", "standard_ds2_v2"]
  spot_sizes           = ["standard_ds1_v2", "standard_ds2_v2"]
  preferred_spot_sizes = ["standard_ds1_v2"]
  zones                = ["1","3"]
  preferred_zones      = ["1"]
  custom_data          = ""
  shutdown_script      = ""

  // -------------------------------------------------------------------

  // --- BOOT DIAGNOSTICS ----------------------------------------------
  boot_diagnostics {
    is_enabled      = true
    storage_url     = "https://.blob.core.windows.net/test"
    type            = "unmanaged"
  }
  // -------------------------------------------------------------------

  // --- DATA DISKS ----------------------------------------------------
  data_disk {
    size_gb = 1
    lun     = 1
    type    = "Standard_LRS"
  }
  
  data_disk {
    size_gb = 10
    lun     = 2
    type    = "Standard_LRS"
  }
  // -------------------------------------------------------------------

  // --- EXTENSIONS ----------------------------------------------------
  extension {
    name                       = "extensionName"
    type                       = "customScript"
    publisher                  = "Microsoft.Azure.Extensions"
    api_version                = "2.0"
    minor_version_auto_upgrade = true
    protected_settings         = {
      "script" : "IyEvYmluL2Jhc2gKZWNobyAibmlyIiA+IC9ob29uaXIudHh0Cg=="
    }
  }
  // -------------------------------------------------------------------
  
   // --- IMAGE ---------------------------------------------------------
  image {
    marketplace_image {
      publisher = "Canonical"
      offer     = "UbuntuServer"
      sku       = "16.04-LTS"
      version   = "latest"
    }
  }
  // -------------------------------------------------------------------
  
  // --- LOAD BALANCERS ------------------------------------------------
  load_balancer {
    type                = "loadBalancer"
    resource_group_name = "testResourceGroup"
    name                = "testLoadBalancer"
    sku                 = "Standard"
    backend_pool_names  = ["testBackendPool1","testBackendPool2"]
  }
  // -------------------------------------------------------------------

  // --- LOGIN ---------------------------------------------------------
  login {
    user_name      = "admin"
    ssh_public_key = "33a2s1f3g5a1df5g1ad3f2g1adfg56dfg=="
  }
  // -------------------------------------------------------------------
  
  // --- MANAGED SERVICE IDENTITIES ------------------------------------
  managed_service_identities {
    name                = "mySI2"
    resource_group_name = "myResourceGroup"
  }
  // -------------------------------------------------------------------
  
  // --- NETWORK -------------------------------------------------------
  network {
    network_resource_group_name = "subnetResourceGroup"
    virtual_network_name        = "vname"
    network_interface {
      is_primary       = true
      subnet_name      = "testSubnet"
      assign_public_ip = true
      public_ip_sku    = "STANDARD"
      network_security_group {
        network_resource_group_name = "test"
        name                        = "test"
      }
      enable_ip_forwarding = true
      private_ip_addresses = ["172.23.4.20"]
      additional_ip_configurations {
        name                       = "test"
        private_ip_address_version = "IPv4"
      }
      public_ips {
        resource_group_name = "resourceGroup"
        name                = "test"
      }
      application_security_groups {
        resource_group_name = "AsgResourceGroup"
        name                = "AsgName"
      } 
    }
  }
  // -------------------------------------------------------------------
  
  // --- OS DISK -------------------------------------------------------
  os_disk {
    size_gb = 30
    type    = "Standard_LRS"
  }
  // -------------------------------------------------------------------
  
  // --- SECRETS -------------------------------------------------------
  secret {
    source_vault {
        name                = "string"
        resource_group_name = "string"
    }
    
    vault_certificates {
        certificate_url     = "string"
        certificate_store   = "string"
    }
  }
  // -------------------------------------------------------------------
  
  // --- TAGS ----------------------------------------------------------
  tag {
    tag_key   = "Creator"
    tag_value = "string"
  }
  // -------------------------------------------------------------------


  // --- HEALTH --------------------------------------------------------
  health {
    health_check_types = ["vmState"]
    unhealthy_duration = 300
    grace_period       = 120
    auto_healing       = true
  }
  // -------------------------------------------------------------------
  
  // --- PERSISTENCE ---------------------------------------------------
  should_persist_os_disk      = false
  os_disk_persistence_mode    = "reattach"
  should_persist_data_disks   = true
  data_disks_persistence_mode = "reattach"
  should_persist_network      = true
  // -------------------------------------------------------------------

  // --- SCHEDULING TASKS ----------------------------------------------
  scheduling_task {
    is_enabled      = true
    type            = "pause"
    cron_expression = "44 10 * * *"
  }

  scheduling_task {
    is_enabled      = true
    type            = "resume"
    cron_expression = "48 10 * * *"
  }

  scheduling_task {
    is_enabled      = true
    type            = "recycle"
    cron_expression = "52 10 * * *"
  }
  // -------------------------------------------------------------------
  
  // --- SIGNALS -------------------------------------------------------
  signal {
    type    = "vmReady"
    timeout = 20
  }

  signal {
    type    = "vmReady"
    timeout = 40
  }
  // -------------------------------------------------------------------

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Azure stateful node name.
* `region` - (Required) The Azure region your stateful node will be created in.
* `resource_group_name` - (Required) Name of the Resource Group for stateful node.
* `description` - (Optional) Describe your Azure stateful node.

<a id="strategy"></a>
## Strategy

* `strategy` - (Required) Strategy for stateful node.
  * `draining_timeout` - (Optional, Default `120`) Time (in seconds) to allow the VM be drained from incoming TCP connections and detached from MLB before terminating it during a scale down operation.
  * `fallback_to_on_demand` - (Required) In case of no spots available, Stateful Node will launch an On-demand instance instead.
  * `optimization_windows` - (Optional) Valid format: "ddd:hh:mm-ddd:hh:mm (day:hour(0-23):minute(0-59))", not empty if revertToSpot.performAt = timeWindow.
  * `preferred_life_cycle` - (Optional, Enum `"od", "spot"`, Default `"spot"`) The desired type of VM.
  * `revert_to_spot` - (Optional) Hold settings for strategy correction - replacing On-Demand for Spot VMs.
    * `perform_at` - (Required, Enum `"timeWindow", "never", "always"`, Default `"always"`) Settings for maintenance strategy.

<a id="compute"></a>
## Compute

* `os` - (Required, Enum `"Linux", "Windows"`) Type of operating system.
* `od_sizes` - (Required) Available On-Demand sizes.
* `spot_sizes` - (Required) Available Spot-VM sizes.
* `preferred_spot_sizes` - (Optional) Prioritize Spot VM sizes when launching Spot VMs for the group. If set, must be a sublist of compute.vmSizes.spotSizes.
* `zones` - (Optional, Enum `"1", "2", "3"`) List of Azure Availability Zones in the defined region. If not defined, Virtual machines will be launched regionally.
* `preferred_zones` - (Optional, Enum `"1", "2", "3"`) The AZs to prioritize when launching VMs. If no markets are available in the Preferred AZs, VMs are launched in the non-preferred AZs. Must be a sublist of compute.zones.
* `custom_data` - (Optional) This value will hold the YAML in base64 and will be executed upon VM launch.
* `shutdown_script` - (Optional) Shutdown script for the stateful node. Value should be passed as a string encoded at Base64 only.

<a id="boot_diagnostics"></a>
## Boot Diagnostics

* `boot_diagnostics`
  * `is_enabled`
  * `storage_url`
  * `type`

<a id="data_disks"></a>
## Data Disks

* `data_disk`
  * `size_gb`
  * `lun`
  * `type`

<a id="extensions"></a>
## Extensions

* `extension`
  * `name`
  * `type`
  * `publisher`
  * `api_version`
  * `minor_version_auto_upgrade`
  * `protected_settings`
  * `script`




* `image`
  * `marketplace_image`
    * `publisher`
    * `offer`
    * `sku`
    * `version`
  * `gallery_image`
    * `gallery_resource_group_name`
    * `gallery_name`
    * `image_name`
    * `version_name`
  * `custom_image`
    * `custom_image_resource_group_name`
    * `name`


* load_balancer
  * type
  * resource_group_name
  * name
  * sku
  * backend_pool_names

* login
  * user_name
  * ssh_public_key


* managed_service_identities
  * name
  * resource_group_name


* network
  * network_resource_group_name
  * virtual_network_name
  * network_interface
    * is_primary
    * subnet_name
    * assign_public_ip
    * public_ip_sku
    * network_security_group
      * network_resource_group_name
      * name 
    * enable_ip_forwarding
    * private_ip_addresses
    * additional_ip_configurations
      * name
      * private_ip_address_version
    * public_ips
      * resource_group_name
      * name
    * application_security_groups
      * resource_group_name
      * name


* os_disk
  * size_gb
  * type


* secret
  * source_vault
    * name
    * resource_group_name
  * vault_certificates
    * certificate_url
    * certificate_store

* tag
  * tag_key
  * tag_value


* health
  * health_check_types
  * unhealthy_duration
  * grace_period
  * auto_healing


* should_persist_os_disk
* os_disk_persistence_mode
* should_persist_data_disks
* data_disks_persistence_mode
* should_persist_network


* scheduling_task
  * is_enabled
  * type
  * cron_expression


* signal
  * type
  * timeout
