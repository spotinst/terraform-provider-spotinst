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

  // --- NETWORK -------------------------------------------------------
  network {
    virtual_network_name = "vname"
    subnet_name          = "my-subnet-name"
    resource_group_name  = "subnetResourceGroup"
    assign_public_ip     = true
  }
  // -------------------------------------------------------------------

  // --- LOGIN ---------------------------------------------------------
  login {
    user_name      = "admin"
    ssh_public_key = "33a2s1f3g5a1df5g1ad3f2g1adfg56dfg=="
  }
  // -------------------------------------------------------------------
  
  // --- SCHEDULED TASK ------------------------------------------------
  scheduled_task {
    is_enabled      = true
    cron_expression = "* * * * *"
    task_type       = "scale"
    
    scale_min_capacity = 5
    scale_max_capacity = 8
    adjustment         = 2
    
    adjustment_percentage = 50
    scale_target_capacity = 6
    batch_size_percentage = 33
    grace_period          = 300
  }
 // -------------------------------------------------------------------
 
 // --- SCALING POLICIES ----------------------------------------------
   scaling_up_policy {
       policy_name = "policy-name"
       metric_name = "CPUUtilization"
       namespace   = "Microsoft.Compute"
       statistic   = "average"
       threshold   = 10
       unit        = "percent"
       cooldown    = 60
       
       dimensions {
         name  = "resourceName"
         value = "resource-name"
       }
       dimensions {
         name  = "resourceGroupName"
         value = "resource-group-name"
       }
       
       operator            = "gt"
       evaluation_periods  = "10"
       period              = "60"
       action_type         = "setMinTarget"
       min_target_capacity = 1
     }
 
    scaling_down_policy {
       policy_name = "policy-name"
       metric_name = "CPUUtilization"
       namespace   = "Microsoft.Compute"
       statistic   = "average"
       threshold   = 10
       unit        = "percent"
       cooldown    = 60
       
       dimensions {
           name  = "name-1"
           value = "value-1"
       }
       
       operator           = "gt"
       evaluation_periods = "10"
       period             = "60"
       action_type        = "adjustment"
       adjustment         = "MIN(5,10)"
     }
}
 // -------------------------------------------------------------------
 
```