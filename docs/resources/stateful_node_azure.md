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