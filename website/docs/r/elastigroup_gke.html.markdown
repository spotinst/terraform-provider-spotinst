---
layout: "spotinst"
page_title: "Spotinst: elastigroup_gke"
sidebar_current: "docs-do-resource-elastigroup_gke"
description: |-
   Provides a Spotinst elastigroup resource for Google Cloud using the Google Kubernetes Engine.
---

# spotinst\_elastigroup\_gke

Provides a Spotinst Elastigroup GKE resource. Please see [Importing a GKE cluster](https://api.spotinst.com/elastigroup-for-google-cloud/tutorials/import-a-gke-cluster-as-an-elastigroup/) for detailed information.


## Example Usage

A spotinst_elastigroup_gke supports all of the fields defined in spotinst_elastigroup_gcp. 

There are two main differences:

* you must include `cluster_zone_name` and `cluster_id`
* a handful of parameters are created remotely and will not appear in the diff. A complete list can be found below.

```hcl
resource "spotinst_elastigroup_gke" "example-gke-elastigroup" {
 name              = "example-gke"
 // cluster_id        = "terraform-acc-test-cluster" // deprecated
 cluster_zone_name = "us-central1-a"
 node_image        = "COS"

 // --- CAPACITY ------------
 max_size         = 5
 min_size         = 1
 desired_capacity = 3
 // -------------------------

 // --- INSTANCE TYPES --------------------------------
 instance_types_ondemand    = "n1-standard-1"
 instance_types_preemptible = ["n1-standard-1", "n1-standard-2"]
 // ---------------------------------------------------

 // --- STRATEGY --------------------
 preemptible_percentage = 100
 // ---------------------------------

 integration_gke = {
  location                 = "us-central1-a"
  cluster_id               = "example-cluster-id"
  autoscale_is_enabled     = true
  autoscale_is_auto_config = false
  autoscale_cooldown       = 300
  
  autoscale_headroom = {
    cpu_per_unit    = 1024
    memory_per_unit = 512
    num_of_units    = 2
  }

  autoscale_down = {
    evaluation_periods = 300
  }

  autoscale_labels = {
    key   = "label_key"
    value = "label_value"
  }
 }

  backend_services = [{
    service_name  = "backend-service"
    location_type = "global"
    named_ports = {
      name  = "http"
      ports = [80, 8080]
    }
  }]

}
```

## Argument Reference

All `spotisnt_elastigroup_gcp` arguments are supported. Please be sure to include the following parameters in your `spotinst_elastigroup_gke` template:

* `cluster_zone_name` - (Required) The zone where the cluster is hosted.
* `cluster_id` - (Required) The name of the GKE cluster you wish to import.
* `node_image` - (Optional, Default: `COS`) The image that will be used for the node VMs. Possible values: COS, UBUNTU.

<a id="third-party-integrations"></a>
## Third-Party Integrations

* `integration_gke` - (Required) Describes the [GKE]() integration.

    * `location` - (Optional) The location of your GKE cluster.
    * `cluster_id` - (Optional) The GKE cluster ID you wish to import.
    * `autoscale_is_enabled` -  (Optional, Default: `false`) Specifies whether the auto scaling feature is enabled.
    * `autoscale_is_autoconfig` - (Optional, Default: `false`) Enabling the automatic auto-scaler functionality. For more information please see: []().
    * `autoscale_cooldown` - (Optional, Default: `300`) The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start.
    
    * `autoscale_headroom` - (Optional) Headroom for the cluster.
        * `cpu_per_unit` - (Optional, Default: `0`) Cpu units for compute.
        * `memory_per_unit` - (Optional, Default: `0`) RAM units for compute.
        * `num_of_units` - (Optional, Default: `0`) Amount of units for compute.
    
    * `autoscale_down` - (Optional) Enabling scale down.
        * `evaluation_periods` - (Optional, Default: `5`) Amount of cooldown evaluation periods for scale down.
    
    * `autoscale_labels` - (Optional) Labels to assign to the resource.
        * `key` - (Optional) The label name.
        * `value` - (Optional) The label value.
    
            
Usage:

```hcl
 integration_gke = {
  location = "us-central1-a"
  cluster_id = "terraform-acc-test-cluster"
  autoscale_is_enabled     = true
  autoscale_is_auto_config = false
  autoscale_cooldown       = 300
  
  autoscale_headroom = {
    cpu_per_unit    = 1024
    memory_per_unit = 512
    num_of_units    = 2
  }

  autoscale_down = {
    evaluation_periods = 300
  }

  autoscale_labels = {
    key  = "label_key"
    value = "label_value"
  }
 }
```

<a id="diff-suppressed-parameters"></a>
## Diff-suppressed Parameters
The following parameters are created remotely and imported. The diffs have been suppressed in order to maintain plan legibility. You may update the values of these
imported parameters by defining them in your template with your desired new value (including null values).

* `backend_services`
    * `service_name`
    * `location_type`
    * `scheme`
    * `named_port`
        * `port_name`
        * `ports`
* `labels`
    * `key`
    * `value`
* `metadata`
    * `key`
    * `value`
* `tags`
    * `key`
    * `value`
* `service_account`
* `ip_forwarding`
* `fallback_to_od`
* `subnets`
    * `region`
    * `subnet_name`
