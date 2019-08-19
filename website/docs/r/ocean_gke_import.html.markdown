---
layout: "spotinst"
page_title: "Spotinst: ocean_gke_import"
sidebar_current: "docs-do-resource-ocean_gke_import"
description: |-
  Provides a Spotinst Ocean resource using gke.
---

# spotinst\_ocean\_gke\_import

Provides a Spotinst Ocean GKE import resource.

## Example Usage

```hcl
resource "spotinst_ocean_gke_import" "example" {

  cluster_name = "example-cluster-name"
  location     = "us-central1-a"
  
  min_size = 0
  max_size = 2
  desired_capacity = 0
  
  whitelist = ["n1-standard-1", "n1-standard-2"]
  
  backend_services = [{
      service_name  = "example-backend-service"
      location_type = "regional"
      scheme        = "INTERNAL"
      named_ports = {
        name  = "http"
        ports = [80, 8080]
      }
  }]
}
```

```
output "controller_cluster_id" {
  value = "${spotinst_ocean_gke_import.ocean-gke-example.cluster_controller_id}"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_name` - (Required) The GKE cluster name.
* `location` - (Required) The zone the master cluster is located in. 
* `max_size` - (Optional, Default: `1000`) The upper limit of instances the cluster can scale up to.
* `min_size` - (Optional) The lower limit of instances the cluster can scale down to.
* `desired_capacity` - (Optional) The number of instances to launch and maintain in the cluster. 
* `backend_services` - (Optional) Describes the backend service configurations.
    * `service_name` - (Required) The name of the backend service.
    * `location_type` - (Optional) Sets which location the backend services will be active. Valid values: `regional`, `global`.
    * `scheme` - (Optional) Use when `location_type` is `regional`. Set the traffic for the backend service to either between the instances in the vpc or to traffic from the internet. Valid values: `INTERNAL`, `EXTERNAL`.
    * `named_port` - (Optional) Describes a named port and a list of ports.
        * `port_name` - (Required) The name of the port.
        * `ports` - (Required) A list of ports.
