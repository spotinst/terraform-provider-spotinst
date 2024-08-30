---
layout: "spotinst"
page_title: "Spotinst: ocean_right_sizing"
subcategory: "Ocean"
description: |-
  Provides a Spotinst Ocean right sizing rule resource.
---

# spotinst\_ocean\_right\_sizing\_rule

Manages a Spotinst Ocean right sizing rule resource.

## Example Usage

```hcl
resource "spotinst_ocean_right_sizing_rule" "example" {
  ocean_id = "o-123456"
  rule_name = "test-rule"
  exclude_preliminary_recommendations= true
  restart_replicas="MORE_THAN_ONE_REPLICA"
  recommendation_application_hpa{
    allow_hpa_recommendation= true
  }

  recommendation_application_intervals {
    repetition_basis = "WEEKLY"
    weekly_repetition_basis {
      interval_days = ["THURSDAY"]
      interval_hours_start_time = "13:00"
      interval_hours_end_time = "15:00"
    }
  }
  
  recommendation_application_intervals {
    repetition_basis = "MONTHLY"
    monthly_repetition_basis {
      interval_months = [1,6,9]
      week_of_the_month = ["FIRST","FOURTH"]
      weekly_repetition_basis {
        interval_days = ["SUNDAY"]
        interval_hours_start_time = "09:00"
        interval_hours_end_time = "18:00"
      }
    }
  }

  recommendation_application_boundaries {
    cpu_min = 20
    cpu_max = 90
    memory_min = 60
    memory_max = 90
  }

  recommendation_application_min_threshold {
    cpu_percentage = 0.45
    memory_percentage = 0.80
  }

  recommendation_application_overhead_values {
    cpu_percentage = 0.35
    memory_percentage = 0.55
  }

  attach_workloads {
    namespaces {
      namespace_name = "kube-system"
      workloads {
        workload_type = "DaemonSet"
        workload_name = "kube-proxy"
      }
    }
    namespaces{
      namespace_name = "spot-ocean"
      workloads {
        workload_type = "Deployment"
        regex_name = "ocean-controller-ocean*"
      }
    }
    namespaces{
      namespace_name = "kube-system"
      labels{
        key = "app.kubernetes.io/name"
        value = "aws-node"
      }
    }
  }

  detach_workloads {
    namespaces{
      namespace_name = "spot-ocean"
      workloads {
        workload_type = "Deployment"
        regex_name = "ocean-controller-ocean*"
      }
    }
    namespaces {
      namespace_name = "kube-system"
      workloads {
        workload_type = "DaemonSet"
        workload_name = "kube-proxy"
      }
    }
  }

}
```
```

## Argument Reference

The following arguments are supported:

* `rule_name` - (Required) The unique name of the rule.
* `ocean_id` - (Required) Identifier of the Ocean cluster.
* `restart_replicas` - Enable to sequentially restart pod batches according to recommendations, for all pods, only more than 1 replica, or not any pod. Possible values: `MORE_THAN_ONE_REPLICA`, `ALL_MANIFEST` or `NO_RESTART`.
* `exclude_preliminary_recommendations` - Exclude preliminary recommendations (recommendations based on less than 4 full days of data).
* `recommendation_application_intervals` - (Required) Determines the Ocean Rightsizing rule recommendation application intervals.
    * `repetition_basis` - The repetition basis. Possible values: `WEEKLY` or `MONTHLY`.
    * `monthly_repetition_basis` - Determines the Ocean Rightsizing rule monthly repetition basis.
        * `interval_months` - Array of the months (in number), when we want to trigger the apply recommendations.
        * `week_of_the_month` - Array of the weeks in the month, when we want to trigger the apply recommendations. Possible values: `FIRST`, `SECOND`, `THIRD`, `FOURTH` or `LAST`.
        * `weekly_repetition_basis` - Determines the Ocean Rightsizing rule weekly repetition basis.
            * `interval_days` - Array of the days of the week, when we want to trigger the apply recommendations. Possible values: `SUNDAY`, `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY` or `SATURDAY`.
            * `interval_hours_start_time` - Start time.
            * `interval_hours_end_time` - End time.
* `recommendation_application_intervals` - (Required) Determines the Ocean Rightsizing rule recommendation application intervals.
    * `repetition_basis` - The repetition basis. Possible values: `WEEKLY` or `MONTHLY`.
    * `weekly_repetition_basis` - Determines the Ocean Rightsizing rule weekly repetition basis.
        * `interval_days` - Array of the days of the week, when we want to trigger the apply recommendations. Possible values: `SUNDAY`, `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY` or `SATURDAY`.
        * `interval_hours_start_time` - Start time.
        * `interval_hours_end_time` - End time.
        
* `recommendation_application_boundaries` - Determines the Ocean Rightsizing rule recommendation application boundaries.
    * `cpu_min` - The lower cpu in vCpu.
    * `cpu_max` - The upper cpu in vCpu.
    * `memory_min` - The lower memory in Gib.
    * `memory_max` - The upper memory in Gib.
* `recommendation_application_min_threshold` - Determines the extent of difference between current request and recommendation to trigger a change in percentage.
    * `cpu_percentage` - (Default: `0.05`). Possible range: `[ 0.05 .. 1 ]`.
    * `memory_percentage` - (Default: `0.05`). Possible range: `[ 0.05 .. 1 ]`.
* `recommendation_application_overhead_values` - Determines the Ocean Rightsizing rule recommendation application overhead values.
    * `cpu_percentage` - (Default: `0.1`). Possible range: `[ 0.1 .. 1 ]`.
    * `memory_percentage` - (Default: `0.1`). Possible range: `[ 0.1 .. 1 ]`.
* `recommendation_application_hpa` - Determines by the rule if recommendation application is allowed for workloads with HPA definition.
    * `allow_hpa_recommendation` - (Default: `false`).
   
<a id="attach_workloads"></a>
## Attach Workload
* `attach_workloads` :
    * `namespaces` - :
        * `namespace_name` - List of namespaces.
        * `workloads` - List of workloads.
            * `workload_type` - 
            * `workload_name` - 
            * `regex_name` - 
        * `labels` - 
            * `key` - 
            * `value` - 
    
```hcl

attach_workloads {
  namespaces {
    namespace_name = "kube-system"
    workloads {
      workload_type = "DaemonSet"
      workload_name = "kube-proxy"
    }
  }
  namespaces{
    namespace_name = "spot-ocean"
    workloads {
      workload_type = "Deployment"
      regex_name = "ocean-controller-ocean*"
    }
  }
  namespaces{
    namespace_name = "kube-system"
    labels{
      key = "app.kubernetes.io/name"
      value = "aws-node"
    }
  }
}
```


<a id="detach_workloads"></a>
## Detach Workloads
* `detach_workloads` :
    * `namespaces` - :
        * `namespace_name` - List of namespaces.
        * `workloads` - List of workloads.
            * `workload_type` -
            * `workload_name` -
            * `regex_name` -
        * `labels` -
            * `key` -
            * `value` -
```hcl
detach_workloads {
  namespaces{
    namespace_name = "spot-ocean"
    workloads {
      workload_type = "Deployment"
      regex_name = "ocean-controller-ocean*"
    }
  }
  namespaces {
    namespace_name = "kube-system"
    workloads {
      workload_type = "DaemonSet"
      workload_name = "kube-proxy"
    }
  }
}
```


