---
layout: "spotinst"
page_title: "Spotinst: oceancd_verification_template"
subcategory: "OceanCD"
description: |-
  Provides a Spotinst OceanCD Verification Template resource.
---

# spotinst\_oceancd\_verification\_template

Manages a Spotinst OceanCD Verfification Template resource.

## Example Usage

```hcl
resource "spotinst_oceancd_verification_template" "test" {
  name  = "test-verification-template-tes"

// --- args ----------------------------------------------------------------
  args{
     arg_name = "test-arg"
     value = "test"
     value_from {
        secret_key_ref {
           name     = "test_key"
           key      = "key-value-test"
        }
    }
  }
// ----------------------------------------------------------------------------

// --- metrics ----------------------------------------------------------------
  metrics {
      metrics_name            = "test-metrics-names"
      dry_run                 = false
      interval                = "10m"
      initial_delay           = "1m"
      count                   = "10"
     success_condition        = "result[0] <= 0.95"
      failure_condition       = "result[0] >= 0.95"
      failure_limit           = 2
      consecutive_error_limit = 1

      provider {
        prometheus {
           prometheus_query = "http_requests_new"
        }

        datadog {
          duration = "1m"
          datadog_query    = "avg:kubernetes.cpu.user.total"
        }

        new_relic {
          profile = "test"
          new_relic_query   = "FROM Metric SELECT count"
        }

        cloud_watch {
          duration = "5m"
          metric_data_queries {
            id = "utilization"
            metric_stat {
              metric {
                metric_name = "Test"
                namespace   = "AWS/EC2"
                dimensions {
                  dimension_name = "instandId"
                  dimension_value          = "i-123044"
                }
              }

              metric_period = 400
              stat          = "average"
              unit          = "None"
            }

            expression  = "SELECT AVG(CPUUtilization) FROM SCHEMA"
            label       = "TestLabel"
            return_data = false
            period      = 300
          }
        }

        web {
          method  = "GET"
          url     = "https://oceancd.com/api/v1/metrics?clusterId= args.clusterId"
          web_header {
            web_header_key   = "Autorization"
            web_header_value = "Bearer=args.token"
          }
          body            = "{\"key\": \"test\"}"
          timeout_seconds = 20
          json_path       = "$.data"
          insecure        = false
        }

        job {
          spec {
            backoff_limit = 1
            job_template {
              template_spec {
                containers {
                  container_name = "hello"
                  command = ["sh","-c"]
                  image    = "nginx.2.1"
                }
                restart_policy = "never"
              }
            }
          }
        }
          jenkins {
                pipeline_name     = "testPipelineName"
                tls_verification  = true
                timeout           = "2m"
                jenkins_interval  = "5s"
                jenkins_parameters {
                  parameter_key   = "app"
                  parameter_value = "my-app"
                }
            }
        
      }
      baseline {
        baseline_provider {
           prometheus {
             prometheus_query = "http_requests_total.status!"
           }

           datadog {
             duration = "2m"
             datadog_query    = "avg:kubernetes.cpu.user"
           }

           new_relic {
             profile = "test"
             new_relic_query   = "FROM Metric SELECT count*"
           }
        }
        min_range = 40
        max_range = 50
        threshold = "range"
      }
  }
 // ----------------------------------------------------------------------------
}

```

```
output "name" {
  value = spotinst_oceancd_verification_template.example.name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Identifier name for Ocean CD Verification Template. Must be unique.
* `args` - (Optional) List of verification arguments. You may specify either value OR valueFrom but not both.
    * `arg_name` - (Required) Name of an argument.
    * `value` - (Optional) String representation of data.
    * `value_from` - (Optional) Value representation of from data.
        * `secret_key_ref` - (Optional) Secret key to use.
            * `name` - (Required) The name of the secret.
            * `key` - (Optional) The name of the field inside the secret.
* `metrics` - (Required) List of verification metrics.
    * `baseline` - (Optional) Baseline Object.
      * `max_range` - (Optional) Number in percent we allow the new version’s data result to be under baseline data result.
      * `min_range` - (Optional) Number in percent we allow the new version’s data result to be under baseline data result.* 
      * `threshold` - (Required) A mathematical expression needed for the comparison. Enum: "<" ">" "<=" ">=" "=" "range"
      * `baseline_provider` - (Required) The name of the monitoring tool chosen for the metric.
          * `datadog`    - (Optional) The datadog provider.
              * `duration` - (Optional) The window of time we are looking at in DataDog.
              * `datadog_query`    - (Required) A request for information retrieved from Datadog.
          * `new_relic`    - (Optional) The New Relic provider.
            * `profile` - (Optional) The name of the secret holding NR account configuration.
            * `new_relic_query`   - (Required) A raw newrelic NRQL query to perform.
          * `prometheus`    - (Optional) The Prometheus provider.
            * `prometheus_query`   - (Required) A request for information retrieved from Prometheus.
    * `consecutive_error_limit` - (Optional) The maximum number of times the measurement is allowed to error in succession, before the metric is considered error.When choosing Jenkins as the provider, there is no need to send this variable.
    * `count` - (Optional) The number of times to run the measurement.
    * `dry_run` - (Optional) Defines whether the metric should have an impact on the result of the rollout.
    * `failure_condition` - (Optional) An expression which determines if a measurement is considered failed.If failureCondition is set, then successCondition is not allowed. When choosing Jenkins as the provider, there is no need to send this variable.
    * `failure_limit` - (Optional) The maximum number of times the measurement is allowed to fail, before the entire metric is considered failed.When choosing Jenkins as the provider, there is no need to send this variable.
    * `initial_delay` - (Optional) How long to wait before starting this metric measurements. When choosing Jenkins as the provider, there is no need to send this variable.
    * `interval` - (Optional) Defines an interval string (30s, 5m, 1h) between each verification measurements. If omitted, will perform a single measurement.When choosing Jenkins as the provider, there is no need to send this variable.
    * `metrics_name` - (Required) The name of the verification metric.
    * `success_condition` - (Optional) An expression which determines if a measurement is considered successful.The keyword result is a variable reference to the value of measurement. Results can be both structured data or primitive. If successCondition is set, then failureCondition is not allowed.
    * `provider` - (Required) The name of the monitoring tool chosen for the metric.
        * `datadog`    - (Optional) The datadog provider.
            * `duration` - (Optional) The window of time we are looking at in DataDog.
            * `datadog_query` - (Required) A request for information retrieved from Datadog.
        * `new_relic`    - (Optional) The New Relic provider.
            * `profile` - (Optional) The name of the secret holding NR account configuration.
            * `new_relic_query`   - (Required) A raw newrelic NRQL query to perform.
        * `prometheus`    - (Optional) The Prometheus provider.
            * `new_relic_query`   - (Required) A request for information retrieved from Prometheus.
        * `cloud_watch`    - (Optional) The CloudWatch provider.
            * `duration`   - (Optional) The window of time we are looking at in CloudWatch.
            * `metric_data_queries` - (Required) The metric queries to be returned. A single MetricData call can include as many as 500 MetricDataQuery structures. Each of these structures can specify either a metric to retrieve, a Metrics Insights query, or a math expression to perform on retrieved data.
                * `expression` - (Optional) This field can contain either a Metrics Insights query, or a metric math expression to be performed on the returned data. Within one metricdataquery object, you must specify either expression or metric stat but not both.
                * `id` - (Required) The response ID.
                * `label` - (Optional) A human-readable label for this metric or expression. If the metric or expression is shown in a CloudWatch dashboard widget, the label is shown
                * `period` - (Optional) The granularity, in seconds, of the returned data points.
                * `return_data` - (Optional) This option indicates whether to return the timestamps and raw data values of this metric.If you are performing this call just to do math expressions and do not also need the raw data returned, you can specify False .
                * `metric_stat` - (Optional) The metric to be returned, along with statistics, period, and units. Use this parameter only if this object is retrieving a metric and not performing a math expression on returned data.
                    * `metric_period` - (Optional) The granularity, in seconds, of the returned data points.
                    * `stat` - (Optional) The statistic to return. It can include any CloudWatch statistic or extended statistic.
                    * `unit` - (Optional) This defines what unit you want to use when storing the metric.
                    * `metric` - (Optional) The metric to be returned, along with statistics, period, and units. Use this parameter only if this object is retrieving a metric and not performing a math expression on returned data.Within one metricdataquery object, you must specify either expression or metric stat but not both.
                        * `metric_name` - (Required) The namespace of the metric.
                        * `namespace` - (Optional) The name of the metric.
                        * `dimensions` - (Optional) A dimension is a name/value pair that is part of the identity of a metric.
                            * `dimension_name` - (Required) The name of the dimensions. These values must contain only ASCII characters and must include at least one non-whitespace characte
                            * `dimension_value` - (Required) The value of the dimensions.These values must contain only ASCII characters and must include at least one non-whitespace characte
        * `web`  - (Optional) The Web provider.
            * `body`   - (Optional) The body of the web metric.
            * `insecure`   - (Optional) Skips host TLS verification.
            * `json_path`  - (Optional) A JSON Path to use as the result variable.
            * `method`     - (Optional) The method of the web metric.  Enum: "GET" "POST" "PUT"
            * `timeout_seconds` - (Optional) The timeout for the request in seconds. Default is 10.
            * `url`   - (Required) The address of the web metric.
            * `web_header`   - (Optional) Optional HTTP headers to use in the request.
                * `web_header_key`   - (Required) The name of a header
                * `web_header_value`   - (Required) The value of a header
        * `jenkins`  - (Optional) The Jenkins provider.
            * `pipeline_name`   - (Required) The Jenkins pipeline name.
            * `tls_verification`   - (Optional) Host TLS verification.
            * `timeout`  - (Required) The total jenkins timeout.
            * `jenkins_interval`   - (Required) The interval time to poll status.
            * `jenkins_parameters` - (Optional) The timeout for the request in seconds. Default is 10.
                * `parameter_key`   - (Required) Key of an argument.
                * `parameter_value` - (Required) Value of an argument.
        * `job`  - (Optional) The Job provider.
            * `spec`   - (Required) The job spec require to run the metric.
                * `backoff_limit`   - (Optional) Specifies the number of retries before marking this job failed.
                * `job_template`  - (Required) The total jenkins timeout.
                  * `template_spec`     - (Required) Specification of the desired behavior of the pod.
                      * `restart_policy` - (Required) Restart policy for all containers within the pod.
                      * `containers` - (Required) A list of containers belonging to the pod.
                        * `container_name` - (Required) The name of a container.
                        * `command` - (Required) The entry point of a container.
                        * `image` - (Required) The image name of a container.
