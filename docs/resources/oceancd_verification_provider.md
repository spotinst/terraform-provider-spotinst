---
layout: "spotinst"
page_title: "Spotinst: oceancd_verification_provider"
subcategory: "OceanCD"
description: |-
  Provides a Spotinst OceanCD Verification Provider resource.
---

# spotinst\_oceancd\_verification\_provider

Manages a Spotinst OceanCD Verfification Provider resource.

## Example Usage

```hcl
resource "spotinst_oceancd_verification_provider" "example" {
  
  name  = "test-verification-provider"

  cluster_ids = [
    "Example-Cluster-Id-1",
    "Example-Cluster-Id-2",
    "Example-Cluster-Id-3"
  ]

  // --- datadog ----------------------------------------------------------------
  
  datadog {
    address    = 1024
    api_key    = 512
    app_key    = 0
  }
  
  // ----------------------------------------------------------------------------
 
 // --- cloudwatch ----------------------------------------------------------------
  
  cloud_watch {
    iam_arn    = "arn:aws:iam::123456789012:role/GetMetricData"
  }
  
  // ----------------------------------------------------------------------------
 
  // --- prometheus ----------------------------------------------------------------
  
  prometheus {
    address    = "http://localhost:9090"
  }
  
  // ----------------------------------------------------------------------------
  
  // --- newRelic ----------------------------------------------------
  
  new_relic {
    personal_api_key    = "AUO32RN20oUMD-40283"
    account_id          = "account-0189718"
    region              = "eu"
    base_url_rest       = "https://rest.api.newrelic.eu"
    base_url_nerd_graph = "https://nerdgraph.api.newrelic.eu"
  }
  
  // -------------------------------------------------------------------------

  // --- jenkins --------------------------------------------------
  
  jenkins {
    base_url    = "http://localhost:9090"
    username    = "test-user"
    api_token   = "AbCDeeFFGG"
    
  }
 
  // -------------------------------------------------------------------------
 
}
```

```
output "name" {
  value = spotinst_oceancd_verification_provider.example.name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Identifier name for Ocean CD Verification Provider. Must be unique.
* `cluster_ids` - (Required) List of cluster IDs that this Verification Provider will be applied to.
* `datadog` - (Optional) Specify the credentials for datadog verification provider.
    * `address` - (Required) DataDog API URL.
    * `api_key` - (Required) API key required by the Datadog Agent to submit metrics and events to Datadog.
    * `app_key` - (Required) API key that gives users access to Datadog’s programmatic API.
* `cloud_watch` - (Optional) Specify the credentials for CloudWatch verification provider.
    * `iam_arn` - (Required) Set label key.
* `prometheus` - (Optional) Specify the credentials for prometheus verification provider.
    * `address` - (Required) The address which the Prometheus server available on.
* `new_relic` - (Optional) Specify the credentials for New Relic verification provider.
    * `accound_id`       - (Required) The ID number New Relic assigns to their account.
    * `base_url_nerd_graph` - (Optional) The base URL for NerdGraph for a proxy.
    * `base_url_rest`    - (Optional) The base URL of the New Relic REST API for a proxy.
    * `personal_api_key` - (Required) The NewRelic user key
    * `region`           - (Optional) A region which the account is attached to. Default is "us".
* `jenkins` - (Optional) Specify the credentials for Jenkins verification provider.
    * `api_token`  - (Required) The Jenkins server’s access apiToken.
    * `base_url`   - (Required) The address of the Jenkins server within the cluster.
    * `username`  - (Required) The Jenkins server’s access username.
