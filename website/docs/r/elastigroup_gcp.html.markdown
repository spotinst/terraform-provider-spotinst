---
layout: "spotinst"
page_title: "Spotinst: elastigroup_gcp"
sidebar_current: "docs-do-resource-elastigroup_gcp"
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
  
  min_size         = 0
  max_size         = 1
  desired_capacity = 1

  availability_zones = ["asia-east1-c", "us-central1-a"]

  preemptible_percentage = 50
  # on_demand_count        = 2
  fallback_to_od         = true
  draining_timeout       = 180
  
  labels           = [key = "env", value = "staging"]
  tags             = ["http", "https"]
  backend_services_config = {[
    {
      service_name = "spotinst-elb-backend-service"
      ports = {
        port_name = "port-name"
        ports = [8000, 6000]
      }
    },
  ]}

  disks = [
    {
      device_nime = "device"
      mode        = "READ_WRITE"
      type        = "PERSISTENT"
      auto_delete = true
      boot        = true
      interface   = "SCSI"

      initialize_parms = {
        disk_size_gb = 10
        disk_type    = "pd-standard"
        source_image = ""
      }
    }
  ]

  network_interfaces = [
    {
      network = "spot-network"
    }
  ]

  instance_types_on_demand   = ["n1-standard-1"]
  instance_types_preemptible = ["n1-standard-1", "n1-standard-2"]

  instance_types_custom = [
    {
      vCPU      = 2
      memoryGiB = 7.5
    }
  ]

  subnets = [
    {
      region       = "asia-east1"
      subnet_names = ""
    }
  ]

  scaling = {
    up = {
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
      
      action = {
        type       = "adjustment"
        adjustment = 1
      }

      dimensions = [
        {
          name  = "storage_type"
          value = "pd-ssd"
        }
      ]
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The group name. 
* `description` - (Optional) The region your GCP group will be created in.
* `startup_script` - (Optional) Create and run your own startup scripts on your virtual machines to perform automated tasks every time your instance boots up.
* `service_account` - (Optional) The email of the service account in which the group instances will be launched.

* `max_size` - (Required) The maximum number of instances the group should have at any time.
* `min_size` - (Required) The minimum number of instances the group should have at any time.
* `desired_capacity` - (Required) The desired number of instances the group should have at any time.

* `availability_zones` - (Required) List of availability zones for the group.

* `subnets` - (Optional) A list of regions and subnets.
* `region` - (Required) The region for the group of subnets.
* `subnet_names` - (Required) The names of the subnets in the region.
* `instance_types_preemptible` - (Required) The preemptible VMs instance type. To maximize cost savings and market availability, select as many types as possible. Required if instance_types_on_demand is not set.
* `instance_types_on_demand` - (Required) The regular VM instance type to use for mixed-type groups and when falling back to on-demand. Required if instance_types_preemptible is not set.

* `instance_types_custom` - (Required) Defines a set of custom instance types. Required if instance_types_preemptible and instance_types_on_demand are not set.
* `vCPU` - (Optional) The number of vCPUs in the custom instance type. GCP has a number of limitations on accepted vCPU values. For more information, see the GCP documentation (here.)[https://cloud.google.com/compute/docs/instances/creating-instance-with-custom-machine-type#specifications]
* `memory_gib` - (Optional) The memory (in GiB) in the custom instance types. GCP has a number of limitations on accepted memory values.For more information, see the GCP documentation (here.)[https://cloud.google.com/compute/docs/instances/creating-instance-with-custom-machine-type#specifications]

* `preemptible_percentage` - (Optional) Percentage of Preemptible VMs to spin up from the "desired_capacity".
* `on_demand_count` - (Optional) Number of regular VMs to launch in the group. The rest will be Preemptible VMs. When this parameter is specified, the preemptible_percentage parameter is being ignored.
* `fallback_to_od` - (Optional) Activate fallback-to-on-demand. When provisioning an instance, if no Preemptible market is available, fallback-to-on-demand will provision an On-Demand instance to maintain the group capacity.
* `draining_timeout` - (Optional) Time (seconds) the instance is allowed to run after it is detached from the group. This is to allow the instance time to drain all the current TCP connections before terminating it.

* `metadata` - (Optional) Array of objects with key-value pairs.
* `key` - (Optional) Metadata key.
* `value` - (Optional) Metadata value.

* `labels` - (Optional) Array of objects with key-value pairs.
* `key` - (Optional) Labels key.
* `value` - (Optional) Labels value.

* `tags` - (Optional) Tags to mark created instances.

<a id="GPU"></a>
## GPU

* `gpu` - (Optional) Defines the GPU configuration.
* `type` - (Required) The type of GPU instance. Valid values: `nvidia-tesla-v100`, `nvidia-tesla-p100`, `nvidia-tesla-k80`.
* `count` - (Required) The number of GPUs. Must be 0, 2, 4, 6, 8.

Usage:

```hcl
  gpu = {
    count = 2
    type = "nvidia-tesla-p100"
  }
```

<a id="health-check"></a>
## Health Check

* `health_check_grace_period` - (optional) Period of time (seconds) to wait for VM to reach healthiness before monitoring for unhealthiness.

```hcl
  health_check_grace_period = 100
```

<a id="backend-services"></a>
## Backend Services

* `backend_services` - (Optional) Describes the backend service configurations.
* `service_name` - (Required) The name of the backend service.
* `location_type` - (Optional) Sets which location the backend services will be active. Valid values: `regional`, `global`.
* `scheme` - (Optional) Use when `location_type` is "regional". Set the traffic for the backend service to either between the instances in the vpc or to traffic from the internet. Valid values: `INTERNAL`, `EXTERNAL`.
* `named_port` - (Optional) Describes a named port and a list of ports.
* `port_name` - (Required) The name of the port.
* `ports` - (Required) A list of ports.

Usage:

```hcl
  backend_services_config = {[
    {
      service_name = "spotinst-elb-backend-service"
      locationType = "regional"
      scheme       = "INTERNAL"
      ports = {
        port_name = "port-name"
        ports = [8000, 6000]
      }
    },
  ]}
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
  disks = [
    {
      device_nime = "device"
      mode        = "READ_WRITE"
      type        = "PERSISTENT"
      auto_delete = true
      boot        = true
      interface   = "SCSI"

      initialize_parms = {
        disk_size_gb = 10
        disk_type    = "pd-standard"
        source_image = ""
      }
    }
  ]
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
  network_interface = [{ 
    network = "default"
	
    access_configs = {
      name = "config1"
      type = "ONE_TO_ONE_NAT"
    }

    alias_ip_ranges = {
     subnetwork_range_name = "range-name-1"
     ip_cidr_range = "10.128.0.0/20"
    }
  }]
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
* `action` - (Optional) Scaling action to take when the policy is triggered.
* `type` - (Optional) Type of scaling action to take when the scaling policy is triggered. Valid values: "adjustment", "setMinTarget", "updateCapacity", "percentageAdjustment"
* `adjustment` - (Optional) Value to which the action type will be adjusted. Required if using "numeric" or "percentageAdjustment" action types.
* `dimensions` - (Optional) A list of dimensions describing qualities of the metric.
    * `name` - (Required) The dimension name.
    * `value` - (Required) The dimension value.
    
Usage:

```hcl
  scaling = {
      up = {
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
        
        action = {
          type       = "adjustment"
          adjustment = 1
        }
  
        dimensions = [
          {
            name  = "storage_type"
            value = "pd-ssd"
          }
        ]
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
