---
layout: "spotinst"
page_title: "Spotinst: oceancd_rollout_spec"
subcategory: "OceanCD"
description: |-
  Provides a Spotinst OceanCD Rollout Spec resource.
---

# spotinst\_oceancd\_rollout\_spec

Manages a Spotinst OceanCD Rollout Spec resource.

## Example Usage

```hcl
resource "spotinst_oceancd_rollout_spec" "example" {
  
  rollout_spec_name  = "test-rollout_spec"
  
  // --- Failure Policy ----------------------------------------------------------------
  failure_policy {
    action = "abort"
  }
  
  //------------------------------------------------------------------------------------


  // --- Spot Deployment ----------------------------------------------------------------
 
  spot_deployment {
    cluster_id = "Test-Cluster-Id"
    spot_deployment_name = "TestDeployment"
    namespace = "default"
  }  
  
  //--------------------------------------------------------------------------------------    

  // --- Spot Deployments ----------------------------------------------------------------
 
  spot_deployments {
    cluster_id = "Test-Cluster-Id"
    spot_deployment_name = "TestDeployment"
    namespace = "default"
  }  
  
  //------------------------------------------------------------------------------    
    
  // --- Strategy ----------------------------------------------------------------
 
  strategy {
    strategy_name = "Test_Strategy"
    args {
      arg_name  = "Test-Arg"
      arg_value = "TestArgValue" 
      value_from {
        field_ref {
            field_path = "metatdata.labels['app']"
        }
      } 
    }
  }  
  
  //------------------------------------------------------------------------------    

  // --- Traffic ----------------------------------------------------------------
 
  traffic {
    canary_service = "canary-service-test"
    stable_service = "stable-service-test"
    
    alb {
      alb_annotation_prefix  = "custom-ingress-kubernates"
      alb_ingress            = "test-ingress" 
      alb_root_service       = "test-root-service"
      service_port           = 8080
      stickiness_config {
            duration_seconds = 3600
            enabled          = true
        }
      } 
    }
    
    ambassador {
        mappings = ["test-mappong","test2-mapping"]    
    }
    
    istio {
        destination_rule {
            canary_subset_name    = "canary-subset"
            destination_rule_name = "test-dest-rule"
            stable_subset_name    = "stable-subset"
        }
        virtual_services {
            virtual_service_name   = "test-name"
            virtual_service_routes = ["route1","route2"]
            tls_routes {
                port      = 80
                sni_hosts = "spot.io"
            }    
        }
    }
    
    nginx {
        nginx_annotation_prefix = "customIngress.nginx"
        stable_ingress          = "hello-ingress"
        additional_ingress_annotation {
            canary_by_header = "x-canary-test"
            key1             = "test"
        }
    }
    
    ping_pong {
        ping_service = "stable-service"
        pong_service = "canary-service"
    }
    
    smi {
        smi_root_service   = "stable-service"
        traffic_split_name = "rollout-expample-traffic-split"
    }   
  }   
  //------------------------------------------------------------------------------    
}    
    
```

```
output "name" {
  value = spotinst_oceancd_rollout_spec.example.name
}
```

## Argument Reference

The following arguments are supported:

* `rollout_spec_name` - (Required) Identifier name for Ocean CD Rollout Spec. Must be unique
* `failure_policy` - (Optional) Holds information on how to react when failure happens.
    * `action` - (Required) Choose an action to perform on failure. Default is abort.  Enum: "abort" "pause" "promote".
* `spot_deployment` - (Optional) Represents the SpotDeployment selector.
    * `cluster_id` - (Optional) Ocean CD cluster identifier for the references SpotDeployment.
    * `spot_deployment_name` - (Optional) The name of the SpotDeployment resource
    * `namespace` - (Optional) The namespace which the SpotDeployment resource exists within.
* `spot_deployments` - (Optional) You must specify either spotDeployment OR spotDeployments but not both. Every SpotDeployment has to be unique. If more than one SpotDeployment has been configured, no traffic managers can be set as part of the RolloutSpec.
    * `cluster_id` - (Optional) Ocean CD cluster identifier for the references SpotDeployment.
    * `spot_deployment_name` - (Optional) The name of the SpotDeployment resource
    * `namespace` - (Optional) The namespace which the SpotDeployment resource exists within.
* `strategy` - (Optional) Determines the Ocean CD strategy
    * `strategy_name` - (Required) Ocean CD strategy name identifier.
    * `args` - (Optional) Arguments defined in Verification Templates.
        * `arg_name` - (Required) Name of an argument.
        * `arg_value` - (Optional) Value of an argument.
        * `value_from` - (Optional) Defines from where to get the value of an argument.
            * `field_ref` - (Required) Defines the field path from where to get the value of an argument.
                * `field_path` - (Required) Path to SpotDeployment's field from where to get the value of an argument.
* `traffic` - (Optional) Hosts all of the supported service meshes needed to enable more fine-grained traffic routing. In case SpotDeployments contains more than one SpotDeployment the traffic manager may not be configured.
    * `canary_service` - (Optional) The canary service name.
    * `stable_service` - (Optional) The stable service name.
    * `alb` - (Optional) Holds ALB Ingress specific configuration to route traffic.
        * `alb_annotation_prefix` - (Optional) Has to match the configured annotation prefix on the alb ingress controller.
        * `alb_ingress` - (Required) Refers to the name of an Ingress resource in the same namespace as the SpotDeployment.
        * `alb_root_service` - (Required) References the service in the ingress to the controller should add the action to.
        * `service_port` - (Required) Refers to the port that the Ingress action should route traffic to.
        * `stickiness_config` - (Optional) Allows to specify further settings on the ForwardConfig.
            * `duration_seconds` - (Optional) Defines how long the load balancer should consistently route the user's request to the same target.
            * `enabled` - (Optional) Enables the load balancer to bind a user's session to a specific target.
    * `ambassador` - (Optional) Holds specific configuration to use Ambassador to route traffic.
      * `mappings` - (Required) A list of names of the Ambassador Mappings used to route traffic to the service.
    * `istio` - (Optional) Holds Istio specific configuration to route traffic.
        *  `destination_rule` - (Optional) It references to an Istio DestinationRule to modify and shape traffic. DestinationRule field belongs only to the Subset Level approach.
             * `destination_rule_name` - (Required) Holds the name of the DestinationRule.
             * `canary_subset_name` - (Required) The subset name to modify labels with the canary version.
             * `stable_subset_name` - (Required) The subset name to modify labels with the stable version.
       * `virtual_services` - (Required) Defines a set of traffic routing rules to apply when a host is addressed.
           * `virtual_service_name` - (Required) Holds the name of the VirtualService.
           * `virtual_service_routes` - (Optional) A list of HTTP routes within VirtualService.
           * `tls_routes` - (Optional) A list of HTTPS routes within VirtualService.
               * `port` - (Optional) The port of the TLS Route desired to be matched in the given Istio VirtualService.
               * `sni_hosts` - (Optional) A list of all the SNI Hosts of the TLS Route desired to be matched in the given Istio VirtualService.
    * `nginx` - (Optional) Holds Nginx Ingress specific configuration to route traffic.
      * `nginx_annotation_prefix` - (Optional) Has to match the configured annotation prefix on the Nginx ingress controller.
      * `stable_ingress` - (Required) Refers to the name of an Ingress resource in the same namespace as the SpotDeployment.
      * `additional_ingress_annotation` - (Optional) Provides additional features to add to the canary ingress (such as routing by header, cookie, etc). You can add these Kubernetes annotations to specific Ingress objects to customize their behavior. Above are found examples of accepted k8s keys.
          * `canary_by_header` - (Optional) Allows customizing the header value instead of using hardcoded values.
          * `key1` - (Optional) Any of supported annotations.
    * `ping_pong` - (Optional) Holds the ping and pong services. You can use pingPong field only when using ALB as a traffic manager with the IP Mode approach.
        * `ping_service` - (Required) Holds the name of the ping service.
        * `pong_service` - (Required) Holds the name of the pong service.
    * `smi` - (Optional) Holds TrafficSplit specific configuration to route traffic.
        * `smi_root_service` - (Optional) Holds the name of service that clients use to communicate.
        * `traffic_split_name` - (Optional) Holds the name of the TrafficSplit.