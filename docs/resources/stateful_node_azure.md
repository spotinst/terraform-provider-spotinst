---
layout: "spotinst"
page_title: "Spotinst: stateful_node_azure"
subcategory: "Stateful Node"
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
    capacity_reservation {
      should_utilize       = true
      utilization_strategy = "utilizeOverOD"
      capacity_reservation_groups {
         crg_name                = "crg name"
         crg_resource_group_name = "resourceGroupName"
         crg_should_prioritize   = true
       }
    }    
  }
  // -------------------------------------------------------------------

  // --- COMPUTE -------------------------------------------------------
  os                   = "Linux"
  od_sizes             = ["standard_ds1_v2", "standard_ds2_v2"]
  spot_sizes           = ["standard_ds1_v2", "standard_ds2_v2"]
  preferred_spot_sizes = ["standard_ds1_v2"]
  zones                = ["1","3"]
  preferred_zone      = "1"
  custom_data          = ""
  shutdown_script      = ""
  user_data            = ""
  vm_name              = "VMName"
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
      public_ip_sku    = "Standard"
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
  
  // --- Security ------------------------------------------------------

  security {
    security_type = "Standard"
    secure_boot_enabled = false
    vtpm_enabled = false
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
}
  // -------------------------------------------------------------------
  
  // ---DELETE----------------------------------------------------------
  delete {
    should_terminate_vm = true
    network_should_deallocate = true
    network_ttl_in_hours = 0
    disk_should_deallocate = true
    disk_ttl_in_hours = 0
    snapshot_should_deallocate = true
    snapshot_ttl_in_hours = 0
    public_ip_should_deallocate = true
    public_ip_ttl_in_hours = 0
  }
  // -------------------------------------------------------------------

```

# Argument Reference

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
  * `capacity_reservation` - (Optional) On-demand Capacity Reservation group enables you to reserve Compute capacity in an Azure region or an Availability Zone for any duration of time. [CRG can only be created on the Azure end.](https://learn.microsoft.com/en-us/azure/virtual-machines/capacity-reservation-create)
    * `should_utilize` - (Required) Determines whether capacity reservations should be utilized.
    * `utilization_strategy` - (Required, Enum `"utilizeOverSpot", "utilizeOverOD"`) The priority requested for using CRG. This value will determine if CRG is used ahead of spot VMs or On-demand VMs. (`"utilizeOverOD"`- If picked, we will use CRG only in case On demand should be launched. `"utilizeOverSpot"`- CRG will be preferred over Spot. Only after CRG is fully used, spot VMs can be used.)
    * `capacity_reservation_groups` - (Optional) List of the desired CRGs to use under the associated Azure subscription. When null we will utilize any available reservation that matches the launch specification.
      * `crg_name` - (Required) The name of the CRG.
      * `crg_resource_group_name` - (Required) Azure resource group name
      * `crg_should_prioritize` - The desired CRG to utilize ahead of other CRGs in the subscription.

<a id="compute"></a>
## Compute

* `os` - (Required, Enum `"Linux", "Windows"`) Type of operating system.
* `od_sizes` - (Required) Available On-Demand sizes.
* `spot_sizes` - (Required) Available Spot-VM sizes.
* `preferred_spot_sizes` - (Optional) Prioritize Spot VM sizes when launching Spot VMs for the group. If set, must be a sublist of compute.vmSizes.spotSizes.
* `zones` - (Optional, Enum `"1", "2", "3"`) List of Azure Availability Zones in the defined region. If not defined, Virtual machines will be launched regionally.
* `preferred_zone` - (Optional, Enum `"1", "2", "3"`) The AZ to prioritize when launching VMs. If no markets are available in the Preferred AZ, VMs are launched in the non-preferred AZ. Must be a sublist of compute.zones.
* `custom_data` - (Optional) This value will hold the YAML in base64 and will be executed upon VM launch.
* `shutdown_script` - (Optional) Shutdown script for the stateful node. Value should be passed as a string encoded at Base64 only.
* `user_data` - (Optional) Define a set of scripts or other metadata that's inserted to an Azure virtual machine at provision time. (Base64 encoded)
* `vm_name` - (Optional) Set a VM name that will be persisted throughout the entire node lifecycle.

<a id="boot_diagnostics"></a>
## Boot Diagnostics

* `boot_diagnostics`
  * `is_enabled` - (Optional) Allows you to enable and disable the configuration of boot diagnostics at launch.
  * `storage_url` - (Optional) The storage URI that is used if a type is unmanaged. The storage URI must follow the blob storage URI format ("https://.blob.core.windows.net/"). StorageUri is required if the type is unmanaged. StorageUri must be ‘null’ in case the boot diagnostics type is managed.
  * `type` - (Optional, Enum `"managed", "unmanaged"`) Defines the storage type on VM launch in Azure.

<a id="data_disks"></a>
## Data Disks

* `data_disk` - (Optional) The definitions of data disks that will be created and attached to the stateful node's VM.
  * `size_gb` - (Required) The size of the data disk in GB, required if dataDisks is specified.
  * `lun` - (Required) The LUN of the data disk.
  * `type` - (Required, Enum `"Standard_LRS", "Premium_LRS", "StandardSSD_LRS", "UltraSSD_LRS"`) The type of the data disk.

<a id="extensions"></a>
## Extensions

* `extension` - (Optional) An object for an azure extension.
  * `name` - (Required) Required on compute.launchSpecification.extensions object.
  * `type` - (Required) Required on compute.launchSpecification.extensions object.
  * `publisher` - (Required) Required on compute.launchSpecification.extensions object.
  * `api_version` - (Required) The API version of the extension. Required if extension specified.
  * `minor_version_auto_upgrade` - (Required) Required on compute.launchSpecification.extensions object.
  * `protected_settings` - (Optional) Object for protected settings.
  * `public_settings` - (Optional) Object for public settings.
  
<a id="image"></a>
## Image

* `image`
  * `marketplace_image` - (Optional) Select an image from Azure's Marketplace image catalogue. Required if the custom image or gallery image are not specified.
    * `publisher` - (Required) Image publisher.
    * `offer` - (Required) Image offer.
    * `sku` - (Required) Image Stock Keeping Unit, which is the specific version of the image.
    * `version` - (Required, Default `"latest"`) Image's version. if version not provided we use "latest".
  * `gallery_image` - (Optional) Gallery image definitions. Required if custom image or marketplace image are not specified.
    * `gallery_resource_group_name` - (Required) The resource group name for gallery image.
    * `gallery_name` - (Required) Name of the gallery.
    * `image_name` - (Required) Name of the gallery image.
    * `version_name` - (Required) Image's version. Can be in the format x.x.x or 'latest'.
  * `custom_image` - (Optional) Custom image definitions. Required if marketplace image or gallery image are not specified.
    * `custom_image_resource_group_name` - (Required) The resource group name for custom image.
    * `name` - (Required) The name of the custom image.

<a id="load balancer"></a>
## Load Balancer

* `load_balancer` - (Optional) Add a load balancer. For Azure Gateway, each Backend Pool is a separate load balancer.
  * `type` - (Required, Enum `"loadBalancer", "applicationGateway"`) The type of load balancer.
  * `resource_group_name` - (Required) The Resource Group name of the Load Balancer.
  * `name` - (Required) Name of the Application Gateway/Load Balancer.
  * `sku` - (Optional)
    * if type is `"LoadBalancer"` then possible values are `“Standard", "Basic”`.
    * If ApplicationGateway then possible values are
      `“Standard_Large”, “Standard_Medium”, “Standard_Small”, “Standard_v2", “WAF_Large”, “WAF_Medium", “WAF_v2"`.
  * `backend_pool_names` - (Optional) Name of the Backend Pool to register the Stateful Node VMs to. Each Backend Pool is a separate load balancer. Required if Type is APPLICATION_GATEWAY.

<a id="login"></a>
## Login

* `login` - (Required) Set admin access for accessing your VMs. Password/SSH is required for Linux.
  * `user_name` - (Required) username for admin access to VMs.
  * `ssh_public_key` - (Optional) SSH for admin access to Linux VMs. Optional for Linux.
  * `password` - (Optional) Password for admin access to Windows VMs. Required for Windows.

<a id="managed_service_identities"></a>
## Managed Service Identities

* `managed_service_identities` - (Optional) Add a user-assigned managed identity to the Stateful Node's VM.
  * `name` - (Required) name of the managed identity.
  * `resource_group_name` - (Required) The Resource Group that the user-assigned managed identity resides in.

<a id="network"></a>
## Network

* `network` - (Required) Define the Virtual Network and Subnet for your Stateful Node.
  * `network_resource_group_name` - (Required) Vnet Resource Group Name.
  * `virtual_network_name` - (Required) Virtual Network.
  * `network_interface` - (Required) Define a network interface
    * `is_primary` - (Required) Defines whether the network interface is primary or not.
    * `subnet_name` - (Required) Subnet name.
    * `assign_public_ip` - (Optional) Assign public IP.
    * `public_ip_sku` - (Optional) Required if assignPublicIp=true values=[Standard/Basic].
    * `network_security_group` - (Optional) Network Security Group.
      * `network_resource_group_name` - (Required) Requires valid security group name.
      * `name` - (Required) Requires valid resource group name.
    * `enable_ip_forwarding` - (Optional) Enable IP Forwarding.
    * `private_ip_addresses` - (Optional) A list with unique items that every item is a valid IP.
    * `additional_ip_configurations` - (Optional) Additional configuration of network interface.
      * `name` - (Required) Configuration name.
      * `private_ip_address_version` - (Required, Enum `"IPv4", "IPv6"` Default `"IPv4"`) Version of the private IP address.
    * `public_ips` - (Optional) Defined a pool of Public Ips (from Azure), that will be associated to the network interface. We will associate one public ip per instance until the pool is exhausted, in which case, we will create a new one.
      * `resource_group_name` - (Required) The resource group of the public ip.
      * `name` - (Required) - The name of the public ip.
    * `application_security_groups` - (Optional) Network Security Group.
      * `resource_group_name` - (Required) Requires valid security group name.
      * `name` - (Required) Requires valid resource group name.

<a id="os_disk"></a>
## OS Disk

* `os_disk` - (Optional) Specify OS disk specification other than default.
  * `size_gb` - (Optional, Default `"30"`) The size of the data disk in GB.
  * `type` - (Required, Enum `"Standard_LRS", "Premium_LRS", "StandardSSD_LRS"`) The type of the OS disk.

<a id="secret"></a>
## Secret

* `secret` - (Optional) Set of certificates that should be installed on the VM.
  * `source_vault` - (Required) The key vault reference, contains the required certificates.
    * `name` - (Required) The name of the key vault.
    * `resource_group_name` - (Required) The resource group name of the key vault.
  * `vault_certificates` - (Required) The required certificate references.
    * `certificate_url` - (Optional) The URL of the certificate under the key vault.
    * `certificate_store` - (Required) The certificate store directory the VM. The directory is created in the LocalMachine account.
      * This field is required only when using Windows OS type
      * This field must be ‘null’ when the OS type is Linux

<a id="secutiry"></a>
## Security

* `security` - (Optional) Specifies the Security related profile settings for the virtual machine.
    * `secure_boot_enabled` - (Optional) Specifies whether secure boot should be enabled on the virtual machine.
    * `security_type` - (Optional) Enum: `"Standard", "TrustedLaunch"` Security type refers to the different security features of a virtual machine. Security features like Trusted launch virtual machines help to improve the security of Azure generation 2 virtual machines.
    * `vtpm_enabled` - (Optional) Specifies whether vTPM should be enabled on the virtual machine.


<a id="tag"></a>
## Tag

* `tag` - (Optional) Unique Key-Value pair for all Stateful Node Resources.
  * `tag_key` - (Optional) Tag Key for Stateful Node Resources.
  * `tag_value` - (Optional) Tag Value for Stateful Node Resources.

<a id="health"></a>
## Health

* `health` - (Optional) Set the auto healing preferences for unhealthy VMs.
  * `health_check_types` - (Optional, Enum `"vmState", "applicationGateway"`) Healthcheck to use to validate VM health.
  * `unhealthy_duration` - (Optional) Amount of time to be unhealthy before a replacement is triggered.
  * `auto_healing` - (Required) Enable Autohealing of unhealthy VMs.
  * `grace_period` - (Optional) Period of time to wait for VM to reach healthiness before monitoring for unhealthiness.

<a id="persistence"></a>
## Persistence

* `should_persist_os_disk` - (Required) Should persist os disk.
* `os_disk_persistence_mode` - (Optional, Enum `"reattach", "onLaunch"`)
* `should_persist_data_disks` - (Required) Should persist data disks.
* `data_disks_persistence_mode` - (Optional, Enum `"reattach", "onLaunch"`)
* `should_persist_network` - (Required) Should persist network.

<a id="scheduling_tasks"></a>
## Scheduling Tasks

* `scheduling_task` - (Optional) Scheduling settings object for stateful node.
  * `is_enabled` - (Required) Is scheduled task enabled for stateful node.
  * `type` - (Required, Enum `"pause", "resume", "recycle") The type of the scheduled task
  * `cron_expression` (Required) A expression which describes when to execute the scheduled task (UTC).

<a id="signals"></a>
## Signals

* `signal` - (Optional) A signal object defined for the stateful node.
  * `type` - (Required, Enum `"vmReady", "vmReadyToShutdown"`) The type of the signal defined for the stateful node.
  * `timeout` - (Required, Default `"1800"`) The timeout in seconds to hold the vm until a signal is sent. If no signal is sent the vm will be replaced (vmReady) or we will terminate the vm (vmReadyToShutdown) after the timeout.

---

<a id="attach_data_disk"></a>
## Attach Data Disk

* `attach_data_disk` - (Optional) Create a new data disk and attach it to the stateful node.
  * `data_disk_name` - (Required) The name of the created data disk.
  * `data_disk_resource_group_name` - (Required) The resource group name in which the data disk will be created.
  * `storage_account_type` - (Required, Enum `"Standard_LRS", "Premium_LRS", "StandardSSD_LRS", "UltraSSD_LRS"`) The type of the data disk.
  * `size_gb` - (Required) The size of the data disk in GB, Required if dataDisks is specified.
  * `zone` - (Optional, Enum `"1", "2", "3"`) The Availability Zone in which the data disk will be created. If not defined, the data disk will be created regionally.
  * `lun` - (Optional, Default `"orginal"`) The LUN of the data disk. If not defined, the LUN will be set in order.

<a id="detach_data_disk"></a>
## Detach Data Disk

* `detach_data_disk` - (Optional) Detach a data disk from a stateful node.
  * `data_disk_name` - (Required) The name of the detached data disk.
  * `data_disk_resource_group_name` - (Required) The resource group name in which the data disk exists.
  * `should_deallocate` - (Required) Indicates whether to delete the data disk in addition to detach.
  * `ttl_in_hours` - (Required, Default `"0"`) Hours to keep the disk alive before deletion.

<a id="update_state"></a>
## Update State

* `update_state` - (Optional) Update the stateful node state.
  * `state` - (Required, Enum `"pause", "resume", "recycle"`) New state for the stateful node.

<a id="import_vm"></a>
## Import VM

* `import_vm` - (Optional) Import an Azure VM and create a stateful node by providing a node configuration.
  * `resource_group_name` - (Required) Name of the Resource Group for Stateful Node.
  * `original_vm_name` - (Required) Azure Import Stateful Node Name.
  * `draining_timeout` - (Optional) Hours to keep resources alive.
  * `resources_retention_time` - (Optional) Hours to keep resources alive.

<a id="delete"></a>
## Deallocation Config

* `delete` - (Required) Specify deallocation parameters for stateful node deletion.
    * `should_terminate_vm` - (Required) Indicates whether to delete the stateful node's VM.
    * `network_should_deallocate` - (Required) Indicates whether to delete the stateful node's network resources.
    * `network_ttl_in_hours` - (Optional, Default: 96) Hours to keep the network resource alive before deletion.
    * `disk_should_deallocate` - (Required) Indicates whether to delete the stateful node's disk resources.
    * `disk_ttl_in_hours` - (Optional, Default: 96) Hours to keep the disk resource alive before deletion.
    * `snapshot_should_deallocate` - (Required) Indicates whether to delete the stateful node's snapshot resources.
    * `snapshot_ttl_in_hours` - (Optional, Default: 96) Hours to keep the snapshots alive before deletion.
    * `public_ip_should_deallocate` - (Required) Indicates whether to delete the stateful node's public ip resources.
    * `public_ip_ttl_in_hours` - (Optional, Default: 96) Hours to keep the public ip alive before deletion.



