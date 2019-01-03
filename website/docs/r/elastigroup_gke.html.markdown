---
layout: "spotinst"
page_title: "Spotinst: elastigroup_gke"
sidebar_current: "docs-do-resource-elastigroup_gke"
description: |-
   Provides a Spotinst elastigroup resource for Google Cloud using the Google Kubernetes Engine.
---

# spotinst\_elastigroup\_gke

Provides a Spotinst elastigroup GKE resource.

## Example Usage

```hcl
resource "spotinst_elastigroup_gke" "example_gke_group" {
  name = "example-gke"

  cluster_zone_name = "us-central1-a"
  cluster_id = "example-cluster"
  node_image = "COS"

  instance_types_ondemand = "n1-standard-1"
  instance_types_preemptible = ["n1-standard-1"]

  min_size         = 0
  max_size         = 3
  desired_capacity = 1
  preemptible_percentage = 75
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) The group name.
* `max_size` - (Required) The maximum number of instances the group should have at any time.
* `min_size` - (Required) The minimum number of instances the group should have at any time.
* `desired_capacity` - (Required) The desired number of instances the group should have at any time.

* `availability_zones` - (Optional) List of availability zones for the group.
* `cluster_zone_name` - (Required) The zone where the cluster is hosted.
* `cluster_id` - (Required) The name of the GKE cluster you wish to import.
* `node_image` - (Optional, Default: `COS`) The image that will be used for the node VMs. Possible values: COS, UBUNTU.
* `preemptible_percentage` - (Optional) The percentage of preemptible VMs that would spin up from the desired capacity (range: 0-100).
* `instance_types_preemptible` - (Optional) The preemptible VMs instance type. To maximize cost savings and market availability, select as many types as possible. Required if instance_types_on_demand is not set.
* `instance_types_on_demand` - (Optional) The regular VM instance type to use for mixed-type groups and when falling back to on-demand. Required if instance_types_preemptible is not set.
