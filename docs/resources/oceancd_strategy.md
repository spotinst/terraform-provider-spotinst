---
layout: "spotinst"
page_title: "Spotinst: oceancd_verification_provider"
subcategory: "OceanCD"
description: |-
  Provides a Spotinst OceanCD Verification Provider resource.
---

# spotinst\_oceancd\_startegy

Manages a Spotinst OceanCD Strategy resource.

## Example Usage

```hcl
resource "spotinst_oceancd_strategy" "example" {
  
  name  = "test-strategy"

  // --- Canary ----------------------------------------------------------------
  
  canary {
    background_verification {
      template_names = ["test1","test2"]
    }
    steps {
       name = "test-step-1"
       set_weight = 20
       
       pause {
          duration = "1s"
       }
       
       set_canary_scale {
          match_traffic_weight = true
          weight               = 20
          replicas             = 20 
       }   
       
       set_header_route {
          name = "test"
          match {
             header_name = "Test"
             header_value {
                exact  = "Test"
                prefix = "Test-"
                regex  = "Test(.*)"
             }
          }
       }
       
       verification {
          template_names = ["test1","test2"]
       }   
        
  }
  
// ----------------------------------------------------------------------------
 
 
 // --- Rolling ----------------------------------------------------------------
  
  rolling {
     steps {
        name = "test-step-1"
      
        pause {
           duration = "1s"
        }
       
        verification {
           template_names = ["test1","test2"]
        }   
     }    
  }
  
  // ----------------------------------------------------------------------------
 
 
}
```

```
output "name" {
  value = spotinst_oceancd_strategy.example.name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Identifier name for Ocean CD Strategy. Must be unique
* `canary` - (Optional) Represents Canary strategy.
    * `background_verification` - (Optional) A list of background verifications.
        * `template_names` - (Required) List of Verification Template names.
    * `steps` - (Required) A set of separate conditions of rollout processing.
        * `name` - (Optional) The name of a step.
        * `pause` - (Optional) Defines the duration of time to freeze the rollout.
            * `duration` - (Optional) The amount of time to wait before moving to the next step.
        * `set_canary_scale` - (Optional) Defines how to scale the version without traffic weight changing.
            * `match_traffic_weight` - (Optional) Defines whether a rollout should match the current canary's setWeight step.
            * `replicas` - (Optional) Sets the number of replicas the new version should have.
            * `weight` - (Optional) Sets the percentage of replicas the new version should have.
        * `set_header_route` - (Optional) Defines the list of HeaderRoutes to add to the Rollout.
            * `match` - (Required) The matching rules for the header route.
                * `header_name` - (Required) The name of the header.
                * `header_value` - (Required) Defines a single header to add to the Rollout.
                    * `exact` - (Optional)  The exact header value.
                    * `prefix` - (Optional) The prefix of the value.
                    * `regex` - (Optional)  The value in a regex format.
        * `set_weight` - (Optional) Defines the percentage that the new version should receive.
        * `verification`  - (Optional) Represents the list of verifications to run in a step.
            * `template_names`  - (Required) List of Verification Template names.
