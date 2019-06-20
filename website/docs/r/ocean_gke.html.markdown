---
layout: "spotinst"
page_title: "Spotinst: ocean_gke"
sidebar_current: "docs-do-resource-ocean_gke"
description: |-
  Provides a Spotinst Ocean resource using gke.
---

# spotinst\_ocean\_gke

Provides a Spotinst Ocean GKE resource.

## Example Usage

```hcl
resource "spotinst_ocean_gke" "example" {
  name               = "example-ocean-cluster-name"
  controller_id      = "example-cluster-id"
  cluster_name       = "example-cluster-name"
  master_location    = "us-central1-a"
  subnet_name        = "example-subnet-1"
  availability_zones = ["us-central1-a"]
  whitelist          = ["n1-standard-1", "n1-standard-2"]
  
  max_size         = 1000
  min_size         = 0
  desired_capacity = 500
 
  // --- LAUNCH CONFIGURATION --------------
  source_image           = "https://www.googleapis.com/compute/v1/projects/my-project/global/examples/example-image-1"
  service_account        = "example-account@my-account.iam.gserviceaccount.com"
  root_volume_size_in_gb = 100
  ip_forwarding          = true
 
  labels {
    key   = "spotinst-gke-original-node-pool",
    value = "example-cluster-name__default-pool"
  }
 
  metadata {
    key   = "cluster-name"
    value = "example-cluster"
  }
 
  tags = ["gke-example-vpc-1234567-node"]
 
  backend_services {
    service_name  = "example-backend-service"
    location_type = "global"
    
    named_ports {
      name  = "http"
      ports = [80, 8080]
    }
   }
  // ---------------------------------------
 
  // --- NETWORK INTERFACE ------------------
   network_interface {
     network = "example-vpc-network"
 
    access_configs {
      name = "config1"
      type = "ONE_TO_ONE_NAT"
    }
 
    alias_ip_ranges {
      subnetwork_range_name = "range-1"
      ip_cidr_range         = "10.8.0.0/20"
    }
   }
  // ----------------------------------------
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The cluster name.
* `controller_id` - (Required) The ocean cluster identifier. Example: `ocean.k8s`
* `cluster_name` - (Required) The GKE cluster name.
* `master_location` - (Required) The zone the master cluster is located in. 
* `subnet_name` - (Required) Subnet identifier for the Ocean cluster.
* `availability_zones` - (Required) List of availability zones available to the cluster.
* `whitelist` - (Optional) Instance types allowed in the Ocean cluster.
* `max_size` - (Optional, Default: `1000`) The upper limit of instances the cluster can scale up to.
* `min_size` - (Optional) The lower limit of instances the cluster can scale down to.
* `desired_capacity` - (Optional) The number of instances to launch and maintain in the cluster.

Usage:

```hcl
  name               = "example-ocean-cluster-name"
  controller_id      = "example-cluster-id"
  cluster_name       = "example-cluster-name"
  master_location    = "us-central1-a"
  subnet_name        = "example-subnet-1"
  availability_zones = ["us-central1-a"]
  whitelist          = ["n1-standard-1", "n1-standard-2"]
  
  max_size         = 1000
  min_size         = 0
  desired_capacity = 500
```

<a id="launch-configuration"></a>
## Launch Configuration
Note: label, metadata, and tag keys are required, and depend on your GKE cluster. Please modify the values to match your configuration. You may also add additional key/value pairs. This resource is intended to be
used as part of a Module.

* `source_image` - (Optional) A source image used to create the disk. You can provide a private (custom) image, and Compute Engine will use the corresponding image from your project.
* `service_account` - (Optional) The email of the service account in which the group instances will be launched.
* `root_volume_size_in_gb` - (Optional) The size (in Gb) to allocate for the root volume. Minimum `100`.
* `ip_forwarding` - (Optional) Enables the transfer IP packets from one network to another.
* `labels` - (Optional) Array of objects with key-value pairs.
    * `key` - (Optional) Labels key.
    * `value` - (Optional) Labels value.
* `metadata` - (Optional) Array of objects with key-value pairs.
    * `key` - (Optional) Metadata key.
    * `value` - (Optional) Metadata value.
* `tags` - (Optional) Tags to mark created instances. Minimum 1.

Usage:

```hcl
  source_image           = "https://www.googleapis.com/compute/v1/projects/my-project/global/examples/example-image-1"
  service_account        = "example-account@my-account.iam.gserviceaccount.com"
  root_volume_size_in_gb = 100
  ip_forwarding          = true
 
  labels {
    key   = "spotinst-gke-original-node-pool",
    value = "example-cluster-name__default-pool"
  }
 
  metadata {
    key   = "cluster-name"
    value = "example-cluster"
  }
 
  tags = ["gke-example-vpc-1234567-node"]
```

<a id="backend-services"></a>
## Backend Services

* `backend_services` - (Optional) Describes the backend service configurations.
    * `service_name` - (Required) The name of the backend service.
    * `location_type` - (Optional) Sets which location the backend services will be active. Valid values: `regional`, `global`.
    * `scheme` - (Optional) Use when `location_type` is "regional". Set the traffic for the backend service to either between the instances in the vpc or to traffic from the internet. Valid values: `INTERNAL`, `EXTERNAL`.
    * `named_port` - (Optional) Describes a named port and a list of ports.
        * `port_name` - (Required) The name of the port.
        * `ports` - (Required) A list of ports.

Usage:
        
```hcl
  backend_services {
    service_name  = "example-backend-service"
    location_type = "global"
    scheme        = "INTERNAL"
    
    named_ports {
      name  = "http"
      ports = [80, 8080]
    }
   }
```

<a id="autoscaler"></a>
## Autoscaler

* `autoscaler` - (Optional) Describes the Ocean Kubernetes autoscaler.
    * `autoscale_is_enabled` - (Optional, Default: `true`) Enable the Ocean Kubernetes autoscaler.
    * `autoscale_is_auto_config` - (Optional, Default: `true`) Automatically configure and optimize headroom resources.
    * `autoscale_cooldown` - (Optional, Default: `null`) Cooldown period between scaling actions.
    * `autoscale_headroom` - (Optional) Spare resource capacity management enabling fast assignment of Pods without waiting for new resources to launch.
        * `cpu_per_unit` - (Optional) Optionally configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.
        * `gpu_per_unit` - (Optional) Optionally configure the number of GPUS to allocate the headroom.
        * `memory_per_unit` - (Optional) Optionally configure the amount of memory (MB) to allocate the headroom.
        * `num_of_units` - (Optional) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.
    * `autoscale_down` - (Optional) Auto Scaling scale down operations.
        * `evaluation_periods` - (Optional, Default: `null`) The number of evaluation periods that should accumulate before a scale down action takes place.
    * `resource_limits` - (Optional) Optionally set upper and lower bounds on the resource usage of the cluster.
        * `max_vcpu` - (Optional) The maximum cpu in vCPU units that can be allocated to the cluster.
        * `max_memory_gib` - (Optional) The maximum memory in GiB units that can be allocated to the cluster.

Usage:

```hcl
  autoscaler {
    autoscale_is_enabled     = false
    autoscale_is_auto_config = false
    autoscale_cooldown       = 300

    autoscale_headroom {
      cpu_per_unit    = 1024
      gpu_per_unit    = 1
      memory_per_unit = 512
      num_of_units    = 2
    }

    autoscale_down {
      evaluation_periods = 300
    }

    resource_limits {
      max_vcpu       = 1024
      max_memory_gib = 20
    }
  }
```
