---
layout: "spotinst"
page_title: "Spotinst: ocean_gke_import"
subcategory: "Ocean"
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
  
  backend_services {
    service_name  = "example-backend-service"
    location_type = "regional"
    scheme        = "INTERNAL"
    
    named_ports {
      name  = "http"
      ports = [80, 8080]
    }
  }
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


<a id="scheduled-task"></a>
## scheduled task
* `scheduled_task` - (Optional) Set scheduling object.
    * `shutdown_hours` - (Optional) Set shutdown hours for cluster object.
        * `is_enabled` - (Optional)  Flag to enable / disable the shutdown hours.
                                     Example: True
        * `time_windows` - (Required) Set time windows for shutdown hours. specify a list of 'timeWindows' with at least one time window Each string is in the format of - ddd:hh:mm-ddd:hh:mm ddd = day of week = Sun | Mon | Tue | Wed | Thu | Fri | Sat hh = hour 24 = 0 -23 mm = minute = 0 - 59. Time windows should not overlap. required on cluster.scheduling.isEnabled = True. API Times are in UTC
                                      Example: Fri:15:30-Wed:14:30
    * `tasks` - (Optional) The scheduling tasks for the cluster.
        * `is_enabled` - (Required)  Describes whether the task is enabled. When true the task should run when false it should not run. Required for cluster.scheduling.tasks object.
        * `cron_expression` - (Required) A valid cron expression. For example : " * * * * * ".The cron is running in UTC time zone and is in Unix cron format Cron Expression Validator Script. Only one of ‘frequency’ or ‘cronExpression’ should be used at a time. Required for cluster.scheduling.tasks object
                                         Example: 0 1 * * *
        * `task_type` - (Required) Valid values: "clusterRoll". Required for cluster.scheduling.tasks object.
        * `batch_size_percentage` - (Optional)  Value in % to set size of batch in roll. Valid values are 0-100
                                                Example: 20.
                          
             
```hcl
  scheduled_task  {
    shutdown_hours  {
      is_enabled = false
      time_windows = ["Fri:15:30-Sat:18:30"]
    }
    tasks {
      is_enabled = false
      cron_expression = "0 1 * * *"
      task_type = "clusterRoll"
      batch_size_percentage = 20
    }
  }
```