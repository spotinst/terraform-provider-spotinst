---
layout: "spotinst"
page_title: "Spotinst: elastigroup_gcp"
subcategory: "Elastigroup"
description: |-
   Provides a Spotinst elastigroup resource for Google Cloud.
---

# spotinst\_elastigroup\_gcp

Provides a Spotinst elastigroup GCP resource.

## Example Usage

```hcl
resource "spotinst_elastigroup_gcp" "example" {
  name            = "example-gcp"
  description     = "spotinst gcp group"
  service_account = "example@myProject.iam.gservicecct.com"
  startup_script  = ""
  instance_name_prefix = "test-123a"
  min_cpu_platform = "Intel Sandy Bridge"
  
  min_size         = 0
  max_size         = 1
  desired_capacity = 1

  availability_zones = ["asia-east1-c", "us-central1-a"]
  preferred_availability_zones = ["us-central1-a"]

  preemptible_percentage = 50
  revert_to_preemptible{
    perform_at="timeWindow"
  }
  optimization_windows=["Mon:01:00-Mon:03:00"]
  # on_demand_count      = 2
  fallback_to_ondemand   = true
  draining_timeout       = 180
  provisioning_model     = "SPOT"
  should_utilize_commitments = true
  
  labels {
    key = "test_key"
    value = "test_value"
  }
  
  tags = ["http", "https"]
  
  backend_services {
     service_name = "spotinst-elb-backend-service"
     location_type = "regional"
     scheme       = "INTERNAL"
     named_ports {
        name = "port-name"
        ports = [8000, 6000]
      }
   }

  disks {
    device_name = "device"
    mode        = "READ_WRITE"
    type        = "PERSISTENT"
    auto_delete = true
    boot        = true
    interface   = "SCSI"

    initialize_params {
      disk_size_gb = 10
      disk_type    = "pd-standard"
      source_image = ""
    }
   }
   
  shielded_instance_config {
    enable_secure_boot = true
    enable_integrity_monitoring = false
  }

  network_interface {
    network = "spot-network"
  }

  instance_types_ondemand    = "n1-standard-1"
  instance_types_preemptible = ["n1-standard-1", "n1-standard-2"]

  instance_types_custom {
    vcpu      = 2
    memory_gib = 7
  }

  subnets {
    region       = "asia-east1"
    subnet_names = [
    "default"
    ]
  }

  scaling_up_policy {
    policy_name = "scale_up_1"
    source      = "stackdriver"
    metric_name = "instance/disk/read_ops_count"
    namespace   = "compute"
    statistic   = "average"
    unit        = "percent"
    threshold   = 10000
    period      = 300
    cooldown    = 300
    operator    = "gte"
    evaluation_periods = 1
    action_type = "adjustment"
    adjustment = 1
    dimensions {
    name  = "storage_type"
    value = "pd-ssd"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The group name. 
* `description` - (Optional) The region your GCP group will be created in.
* `startup_script` - (Optional) Create and run your own startup scripts on your virtual machines to perform automated tasks every time your instance boots up.
* `shutdown_script` - (Optional) The Base64-encoded shutdown script that executes prior to instance termination, for more information please see: [Shutdown Script](https://api.spotinst.com/integration-docs/elastigroup/concepts/compute-concepts/shutdown-scripts/)
* `service_account` - (Optional) The email of the service account in which the group instances will be launched.
* `max_size` - (Required) The maximum number of instances the group should have at any time.
* `min_size` - (Required) The minimum number of instances the group should have at any time.
* `desired_capacity` - (Required) The desired number of instances the group should have at any time.
* `availability_zones` - (Required) List of availability zones for the group.
* `preferred_availability_zones` - (Optional) prioritize availability zones when launching instances for the group. Must be a sublist of `availability_zones`.
* `subnets` - (Optional) A list of regions and subnets.
    * `region` - (Required) The region for the group of subnets.
    * `subnet_names` - (Required) The names of the subnets in the region.
* `instance_types_preemptible` - (Required) The preemptible VMs instance type. To maximize cost savings and market availability, select as many types as possible. Required if instance_types_ondemand is not set.
* `instance_types_ondemand` - (Required) The regular VM instance type to use for mixed-type groups and when falling back to on-demand. Required if instance_types_preemptible is not set.
* `instance_types_custom` - (Required) Defines a set of custom instance types. Required if instance_types_preemptible and instance_types_ondemand are not set.
    * `vCPU` - (Optional) The number of vCPUs in the custom instance type. GCP has a number of limitations on accepted vCPU values. For more information, see the GCP documentation (here.)[https://cloud.google.com/compute/docs/instances/creating-instance-with-custom-machine-type#specifications]
    * `memory_gib` - (Optional) The memory (in GiB) in the custom instance types. GCP has a number of limitations on accepted memory values.For more information, see the GCP documentation (here.)[https://cloud.google.com/compute/docs/instances/creating-instance-with-custom-machine-type#specifications]
* `preemptible_percentage` - (Optional) Percentage of Preemptible VMs to spin up from the "desired_capacity".
* `on_demand_count` - (Optional) Number of regular VMs to launch in the group. The rest will be Preemptible VMs. When this parameter is specified, the preemptible_percentage parameter is being ignored.
* `fallback_to_ondemand` - (Optional) Activate fallback-to-on-demand. When provisioning an instance, if no Preemptible market is available, fallback-to-on-demand will provision an On-Demand instance to maintain the group capacity.
* `draining_timeout` - (Optional) Time (seconds) the instance is allowed to run after it is detached from the group. This is to allow the instance time to drain all the current TCP connections before terminating it.
* `provisioning_model` - (Optional) Valid values: "SPOT", "PREEMPTIBLE". Define the provisioning model of the launched instances. Default value is "PREEMPTIBLE".
* `should_utilize_commitments` - (Optional) Enable committed use discounts utilization.
* `metadata` - (Optional) Array of objects with key-value pairs.
    * `key` - (Optional) Metadata key.
    * `value` - (Optional) Metadata value.
* `labels` - (Optional) Array of objects with key-value pairs.
    * `key` - (Optional) Labels key.
    * `value` - (Optional) Labels value.
* `tags` - (Optional) Tags to mark created instances.
* `instance_name_prefix` - (Optional) Set an instance name prefix to be used for all launched instances and their boot disk. The prefix value should comply with the following limitations: 
    * A maximal length of 25 characters.
    * The first character must be a lowercase letter, and all the following characters must be hyphens, lowercase letters, or digits, except the last character, which cannot be a hyphen.
* `min_cpu_platform` - (Optional) Select a minimum CPU platform for the compute instance.
* `revert_to_preemptible` - (Optional) Setting for revert to preemptible option.
  * `perform_at` - (Required) Valid values: "always", "never", "timeWindow". Required on strategy.revertToPreemptible object.
* `optimization_windows` - (Optional) Set time window to perform the revert to preemptible. Time windows must be at least 120 minutes. Format: DayInWeek:HH-DayInWeek:HH. Required when strategy.revertToPreemptible.performAt is 'timeWindow'.
* `shielded_instance_config` - (Optional) You can use secure boot when you launch VMs using Elastigroup. This helps you comply with your security policies. In the instance configuration, use ‘secureBootEnabled’ set to True to enforce UEFI with secure boot. Elastigroup provisions VMs with secure boot, as long as the images supports UEFI.
  * `enable_secure_boot` - (Optional) Default: false
  * `enable_integrity_monitoring` - (Optional) Default: false
<a id="GPU"></a>
## GPU

* `gpu` - (Optional) Defines the GPU configuration.
    * `type` - (Required) The type of GPU instance. Valid values: `nvidia-tesla-v100`, `nvidia-tesla-p100`, `nvidia-tesla-k80`.
    * `count` - (Required) The number of GPUs. Must be 0, 2, 4, 6, 8.

Usage:

```hcl
  gpu {
    count = 2
    type = "nvidia-tesla-p100"
  }
```

<a id="health-check"></a>
## Health Check

* `auto_healing` - (Optional) Enable auto-replacement of unhealthy instances.
* `health_check_grace_period` - (Optional) Period of time (seconds) to wait for VM to reach healthiness before monitoring for unhealthiness.
* `health_check_type` - (Optional) The kind of health check to perform when monitoring for unhealthiness.
* `unhealthy_duration` - (Optional) Period of time (seconds) to remain in an unhealthy status before a replacement is triggered.

```hcl
  auto_health               = true
  health_check_grace_period = 100
  health_check_type         = ""
  unhealthy_duration        = ""
  ```

<a id="backend-services"></a>
## Backend Services

* `backend_services` - (Optional) Describes the backend service configurations.
    * `service_name` - (Required) The name of the backend service.
    * `location_type` - (Optional) Sets which location the backend services will be active. Valid values: `regional`, `global`.
    * `scheme` - (Optional) Use when `location_type` is "regional". Set the traffic for the backend service to either between the instances in the vpc or to traffic from the internet. Valid values: `INTERNAL`, `EXTERNAL`.
    * `named_ports` - (Optional) Describes a named port and a list of ports.
        * `name` - (Required) The name of the port.
        * `ports` - (Required) A list of ports.

Usage:

```hcl
  backend_services {
     service_name = "spotinst-elb-backend-service"
     location_type = "regional"
     scheme       = "INTERNAL"
     named_ports {
        name = "port-name"
        ports = [8000, 6000]
      }
   }
```

<a id="disks"></a>
## Disks

* `disks` - (Optional) Array of disks associated with this instance. Persistent disks must be created before you can assign them.
    * `auto_delete` - (Optional) Specifies whether the disk will be auto-deleted when the instance is deleted.
    * `boot` - (Optional) Indicates that this is a boot disk. The virtual machine will use the first partition of the disk for its root filesystem.
    * `device_name` - (Optional) Specifies a unique device name of your choice.
    * `interface` - (Optional, Default: `SCSI`) Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. 
    * `mode` - (Optional, Default: `READ_WRITE`) The mode in which to attach this disk, either READ_WRITE or READ_ONLY.
    * `source` - (Optional) Specifies a valid partial or full URL to an existing Persistent Disk resource. This field is only applicable for persistent disks.
    * `type` - (Optional, Default: `PERSISTENT`) Specifies the type of disk, either SCRATCH or PERSISTENT.
    * `initialize_params` - (Optional) Specifies the parameters for a new disk that will be created alongside the new instance. Use initialization parameters to create boot disks or local SSDs attached to the new instance.
        * `disk_size_gb` - (Optional) Specifies disk size in gigabytes. Must be in increments of 2.
        * `disk_type` - (Optional, Default" `pd-standard`) Specifies the disk type to use to create the instance. Valid values: pd-ssd, local-ssd.
        * `source_image` - (Optional) A source image used to create the disk. You can provide a private (custom) image, and Compute Engine will use the corresponding image from your project.

Usage:

```hcl
  disks {
      device_name = "device"
      mode        = "READ_WRITE"
      type        = "PERSISTENT"
      auto_delete = true
      boot        = true
      interface   = "SCSI"

      initialize_params = {
        disk_size_gb = 10
        disk_type    = "pd-standard"
        source_image = ""
      }
    }
```

<a id="network-interface"></a>
## Network Interfaces

Each of the `network_interface` attributes controls a portion of the GCP
Instance's "Network Interfaces". It's a good idea to familiarize yourself with [GCP's Network
Interfaces docs](https://cloud.google.com/vpc/docs/multiple-interfaces-concepts)
to understand the implications of using these attributes.

* `network_interface` - (Required, minimum 1) Array of objects representing the network configuration for the elastigroup.
    * `network` - (Required) Network resource for this group.
    * `access_configs` - (Optional) Array of configurations.
        * `name` - (Optional) Name of this access configuration.
        * `type` - (Optional) Array of configurations for this interface. Currently, only ONE_TO_ONE_NAT is supported.

```hcl
  network_interface { 
    network = "default"
	
    access_configs = {
      name = "config1"
      type = "ONE_TO_ONE_NAT"
    }

    alias_ip_ranges = {
     subnetwork_range_name = "range-name-1"
     ip_cidr_range = "10.128.0.0/20"
    }
  }
```

<a id="scaling-policy"></a>
## Scaling Policies

* `scaling_up_policy` - (Optional) Contains scaling policies for scaling the Elastigroup up.
* `scaling_down_policy` - (Optional) Contains scaling policies for scaling the Elastigroup down.

Each `scaling_*_policy` supports the following:

* `policy_name` - (Optional) Name of scaling policy.
* `metric_name` - (Optional) Metric to monitor. Valid values: "Percentage CPU", "Network In", "Network Out", "Disk Read Bytes", "Disk Write Bytes", "Disk Write Operations/Sec", "Disk Read Operations/Sec".
* `statistic` - (Optional) Statistic by which to evaluate the selected metric. Valid values: "AVERAGE", "SAMPLE_COUNT", "SUM", "MINIMUM", "MAXIMUM", "PERCENTILE", "COUNT".
* `threshold` - (Optional) The value at which the scaling action is triggered.
* `period` - (Optional) Amount of time (seconds) for which the threshold must be met in order to trigger the scaling action.
* `evaluation_periods` - (Optional) Number of consecutive periods in which the threshold must be met in order to trigger a scaling action.
* `cooldown` - (Optional) Time (seconds) to wait after a scaling action before resuming monitoring.
* `operator` - (Optional) The operator used to evaluate the threshold against the current metric value. Valid values: "gt" (greater than), "get" (greater-than or equal), "lt" (less than), "lte" (less than or equal).
* `action_type` - (Optional) Type of scaling action to take when the scaling policy is triggered. Valid values: "adjustment", "setMinTarget", "updateCapacity", "percentageAdjustment"
* `adjustment` - (Optional) Value to which the action type will be adjusted. Required if using "numeric" or "percentageAdjustment" action types.
* `dimensions` - (Optional) A list of dimensions describing qualities of the metric.
    * `name` - (Required) The dimension name.
    * `value` - (Required) The dimension value.
    
Usage:

```hcl
  scaling_up_policy {
    policy_name = "scale_up_1"
    source      = "stackdriver"
    metric_name = "instance/disk/read_ops_count"
    namespace   = "compute"
    statistic   = "average"
    unit        = "percent"
    threshold   = 10000
    period      = 300
    cooldown    = 300
    operator    = "gte"
    evaluation_periods = 1
    action_type = "adjustment"
    adjustment = 1
    dimensions {
    name  = "storage_type"
    value = "pd-ssd"
    }
  }
```

<a id="third-party-integrations"></a>
## Third-Party Integrations

* `integration_docker_swarm` - (Optional) Describes the [Docker Swarm](https://api.spotinst.com/integration-docs/elastigroup/container-management/docker-swarm/docker-swarm-integration/) integration.
    * `master_host` - (Required) IP or FQDN of one of your swarm managers.
    * `master_port` - (Required) Network port used by your swarm.
            
Usage:

```hcl
integration_docker_swarm = {
    master_host = "10.10.10.10"
    master_port = 2376
}
```

<a id="scheduled-task"></a>
## Scheduled Tasks

Each `scheduled_task` supports the following:

* `task_type` - (Required) The task type to run. Valid values: `"setCapacity"`.
* `cron_expression` - (Optional) A valid cron expression. The cron is running in UTC time zone and is in [Unix cron format](https://en.wikipedia.org/wiki/Cron).
* `is_enabled` - (Optional, Default: `true`) Setting the task to being enabled or disabled.
* `target_capacity` - (Optional) The desired number of instances the group should have.
* `min_capacity` - (Optional) The minimum number of instances the group should have.
* `max_capacity` - (Optional) The maximum number of instances the group should have.

Usage:

```hcl
  scheduled_task {
    task_type             = "setCapacity"
    cron_expression       = ""
    is_enabled            = false
    target_capacity       = 5
    min_capacity          = 0
    max_capacity          = 10
  }
```
