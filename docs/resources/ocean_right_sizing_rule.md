---
layout: "spotinst"
page_title: "Spotinst: ocean_right_sizing_rule"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean Right Sizing resource.
---

# spotinst\_ocean\_right\_sizing\_rule

Manages a Spotinst Ocean right sizing rule resource.

## Example Usage

```hcl
resource "spotinst_ocean_right_sizing_rule" "example" {
  ocean_id = "o-123456"
  rule_name = "test-rule"
  exclude_preliminary_recommendations= true
  restart_replicas="ALL_MANIFEST"

  recommendation_application_hpa{
    allow_hpa_recommendations= true
  }

  recommendation_application_intervals {
    repetition_basis = "WEEKLY"
    weekly_repetition_basis {
      interval_days = ["MONDAY", "WEDNESDAY"]
      interval_hours_start_time = "12:00"
      interval_hours_end_time = "14:00"
    }
  }


  recommendation_application_intervals {
    repetition_basis = "MONTHLY"
    monthly_repetition_basis {
      interval_months = [2,6,9]
      week_of_the_month = ["FIRST","LAST"]
      weekly_repetition_basis {
        interval_days = ["MONDAY"]
        interval_hours_start_time = "03:00"
        interval_hours_end_time = "04:00"
      }
    }
  }

  recommendation_application_boundaries {
    cpu_min = 120
    cpu_max = 190
    memory_min = 160
    memory_max = 190
  }

  recommendation_application_min_threshold {
    cpu_percentage = 0.412
    memory_percentage = 0.36
  }

  recommendation_application_overhead_values {
    cpu_percentage = 0.80
    memory_percentage = 0.50
  }
  
  auto_apply_definition {
    enabled = true
    namespaces = ["kube-system", "spot-system"]

    labels = {
      "k8s-app" = "kube-proxy",
      "app.kubernetes.io/name" = "metrics-server"
    }
   }
  
}
````
## Argument Reference

The following arguments are supported:

* `ocean_id` - (Required) Identifier of the Ocean cluster.
* `rule_name` - (Required) The unique name of the rule.
* `exclude_preliminary_recommendations` - (Optional) Exclude preliminary recommendations (recommendations based on less than 4 full days of data).
* `restart_replicas` - (Optional) Valid values: "MORE_THAN_ONE_REPLICA" "ALL_MANIFEST" "NO_RESTART". Enable to sequentially restart pod batches according to recommendations, for all pods, only more than 1 replica, or not any pod.
* `recommendation_application_boundaries` - (Optional) Determines the Ocean Rightsizing rule recommendation application boundaries.
   * `cpu_min` - (Optional) the minimal value of cpu in vCpu.
   * `cpu_max` - (Optional) the maximal value of cpu in vCpu.
   * `memory_min` - (Optional) the minimal value of memory in Gib.
   * `memory_max` - (Optional) the maximal value of memory in Gib.
* `recommendation_application_hpa` - HPA Rightsizing Rule Recommendation Configuration
    * `allow_hpa_recommendations` - Determines by the rule if recommendation application is allowed for workloads with HPA definition.

* `recommendation_application_intervals` - (Required) Determines the Ocean Rightsizing rule recommendation application intervals.
    * `repetition_basis` - (Optional) Valid values: "WEEKLY" "MONTHLY". The repetition basis.
    * `weekly_repetition_basis` - (Optional) Determines the Ocean Rightsizing rule weekly repetition basis.
      * `interval_days` - (Optional) Valid values: "SUNDAY" "MONDAY" "TUESDAY" "WEDNESDAY" "THURSDAY" "FRIDAY" "SATURDAY". Array of the days of the week, when we want to trigger the apply recommendations.
      * `interval_hours_start_time` - (Optional) Start time.
      * `interval_hours_end_time` - (Optional) End time.
    * `monthly_repetition_basis` - (Optional) Determines the Ocean Rightsizing rule monthly repetition basis.
      * `interval_months` - (Optional) Array of the months (in number), when we want to trigger the apply recommendations.
      * `week_of_the_month` - (Optional) Valid values: "FIRST" "SECOND" "THIRD" "FOURTH" "LAST". Array of the weeks in the month, when we want to trigger the apply recommendations.
      * `weekly_repetition_basis` - (Optional) Determines the Ocean Rightsizing rule weekly repetition basis.
        * `interval_days` - (Optional) Valid values: "SUNDAY" "MONDAY" "TUESDAY" "WEDNESDAY" "THURSDAY" "FRIDAY" "SATURDAY". Array of the days of the week, when we want to trigger the apply recommendations.
        * `interval_hours_start_time` - (Optional) Start time.
        * `interval_hours_end_time` - (Optional) End time.
        
* `recommendation_application_min_threshold` - Determines the extent of difference between current request and recommendation to trigger a change in percentage.
    * `cpu_percentage` - (Optional, Default: 0.05) .
    * `memory_percentage` - (Optional, Default: 0.05) .
  
* `recommendation_application_overhead_values` - Determines the Ocean Rightsizing rule recommendation application overhead values.
    * `cpu_percentage` - (Optional, Default: 0.1) .
    * `memory_percentage` - (Optional, Default: 0.1).
* `auto_apply_definition` - (Optional) Ocean Rightsizing Rule Auto Apply Configuration.
  * `enabled` - (Optional) Determines if auto apply is enabled.
  * `namespaces` - (Optional) List of namespaces that match the auto-apply rule.
  * `labels` - (Optional) A set of key-value label pairs used to automatically apply this rule to all workloads in the cluster that match these labels.

<a id="attach_workloads"></a>
## Attach Workloads

* `attach_workloads` - (Optional)
    * `namespaces` - (Optional) 
        * `namespace_name` - (Optional) List of namespaces.
        * `workloads` - (Optional) List of workloads.
          * `workload_type` - (Optional). The type of the workload.
          * `workload_name` - (Optional). The name of the workload.
          * `regex_name` - (Optional). The regex for the workload name. Not allowed when using workload_name.


```hcl
attach_workloads {
  namespaces {
    namespace_name = "kube-system"
    workloads {
      workload_type = "Deployment"
      workload_name = "konnectivity-agent"
    }
  }

  namespaces{
    namespace_name = "kube-system"
    workloads {
      workload_type = "DaemonSet"
      regex_name = "csi-*"
    }
  }
}
```

<a id="detach_workloads"></a>
## Detach Workloads

* `detach_workloads` - (Optional)
    * `namespaces` - (Optional)
        * `namespace_name` - (Optional) List of namespaces.
        * `workloads` - (Optional) List of workloads.
            * `workload_type` - (Optional). The type of the workload.
            * `workload_name` - (Optional). The name of the workload.
            * `regex_name` - (Optional). The regex for the workload name.Not allowed when using workload_name.


```hcl
  detach_workloads {

  namespaces {
    namespace_name = "kube-system"
    workloads {
      workload_type = "Deployment"
      workload_name = "konnectivity-agent"
    }
  }

  namespaces{
    namespace_name = "kube-system"
    workloads {
      workload_type = "DaemonSet"
      regex_name = "csi-*"
    }
  }
}
```
