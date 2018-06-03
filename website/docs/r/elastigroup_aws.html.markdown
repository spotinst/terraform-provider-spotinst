---
layout: "spotinst"
page_title: "Spotinst: elastigroup_aws"
sidebar_current: "docs-do-resource-group_aws"
description: |-
  Provides a Spotinst AWS group resource.
---

# spotinst\_elastigroup\_aws

Provides a Spotinst AWS group resource.

## Example Usage

```hcl
# Create an Elastigroup
resource "spotinst_elastigroup_aws" "default-elastigroup" {

  name        = "default-elastigroup"
  description = "created by Terraform"
  product     = "Linux/UNIX"

  max_size          = 0
  min_size          = 0
  desired_capacity  = 0
  capacity_unit     = "weight"

  region      = "us-west-2"
  subnet_ids  = ["sb-123456", "sb-456789"]
  
  image_id              = "ami-a27d8fda"
  iam_instance_profile  = "iam-profile"
  key_name              = "my-key.ssh"
  security_groups       = ["sg-123456"]
  user_data             = "echo hello world"
  enable_monitoring     = false
  ebs_optimized         = false
  placement_tenancy     = "default"

  instance_types_ondemand = "m3.2xlarge"
  instance_types_spot     = ["m3.xlarge", "m3.2xlarge"]

  instance_types_weights = [
    {
      instance_type = "c3.large"
      weight        = 10
    },
    {
      instance_type = "c4.xlarge"
      weight        = 16
    }]

  orientation           = "balanced"
  fallback_to_ondemand  = false

  ebs_block_device {
    device_name           = "/dev/sdb"
    snapshot_id           = ""
    volume_type           = "gp2"
    volume_size           = 8
    iops                  = 1
    delete_on_termination = true
    encrypted             = false
  }

  scaling_up_policy {
    policy_name        = "Default Scaling Up Policy"
    metric_name        = "DefaultQueuesDepth"
    statistic          = "average"
    unit               = "none"
    adjustment         = 1
    namespace          = "custom"
    threshold          = 100
    period             = 60
    evaluation_periods = 5
    cooldown           = 300
  }

  scaling_down_policy {
    policy_name        = "Default Scaling Down Policy"
    metric_name        = "Default QueuesDepth"
    statistic          = "average"
    unit               = "none"
    adjustment         = 1
    namespace          = "custom"
    threshold          = 10
    period             = 60
    evaluation_periods = 10
    cooldown           = 300
  }

  tags = [
  {
     key = "Env"
     value = "production"
  }, 
  {
     key = "Name"
     value = "default-production"
  },
  {
     key = "Project"
     value = "app_v2"
  }
 ]

  lifecycle {
    ignore_changes = [
      "desired_capacity",
    ]
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The group name.
* `description` - (Required) The group description.
* `product` - (Required) Operation system type. Valid values: `"Linux/UNIX"`, `"SUSE Linux"`, `"Windows"`. 
For EC2 Classic instances:  `"Linux/UNIX (Amazon VPC)"`, `"SUSE Linux (Amazon VPC)"`, `"Windows (Amazon VPC)"`.    

* `availability_zones` - (Optional) TBD
Note: When this parameter is set, `subnet_ids` should be left unused.

* `region` - (Optional) TBD
* `subnet_ids` - (Optional) A comma-separated list of subnet identifiers.
Note: When this parameter is set, `availability_zones` should be left unused.

* `max_size` - (Optional; Required if using scaling policies) The maximum number of instances the group should have at any time.
* `min_size` - (Optional; Required if using scaling policies) The minimum number of instances the group should have at any time.
* `desired_capacity` - (Optional) The desired number of instances the group should have at any time.
* `capacity_unit` - (Optional, Default: `"instance"`) The capacity unit to launch instances by. If not specified, when choosing the weight unit, each instance will weight as the number of its vCPUs.

* `security_groups` - (Required) A list of associated security group IDS.
* `image_id` - (Optional) The ID of the AMI used to launch the instance.
* `iam_instance_profile` - (Optional) The ARN of an IAM instance profile to associate with launched instances.
* `key_name` - (Optional) The key name that should be used for the instance.
* `enable_monitoring` - (Optional) Indicates whether monitoring is enabled for the instance.
* `user_data` - (Optional) The user data to provide when launching the instance.
* `ebs_optimized` - (Optional) TBD
* `placement_tenancy` - (Optional) TBD

* `instance_types_ondemand` - (Required) The type of instance determines your instance's CPU capacity, memory and storage (e.g., m1.small, c1.xlarge).
* `instance_types_spot` - (Required) One or more instance types.
* `instance_types_weights` - (Optional) List of weights per instance type for weighted groups. Each object in the list should have the following attributes:

    * `weight` - (Required) Weight per instance type (Integer).
    * `instance_type` - (Required) Name of instance type (String).

* `fallback_to_ondemand` - (Required) TBD
* `orientation` - (Required, Default: `balanced`) TBD. Valid values: `"balanced"`, `"costOriented"`, `"equalAzDistribution"`, `"availabilityOriented"`.    
* `spot_percentage` - (Optional; Required if not using `ondemand_count`) The percentage of Spot instances that would spin up from the `desired_capacity` number.
* `ondemand_count` - (Optional; Required if not using `spot_percentage`) Number of on demand instances to launch in the group. All other instances will be spot instances. When this parameter is set the `spot_percentage` parameter is being ignored.
* `draining_timeout` - (Optional) The time in seconds, the instance is allowed to run while detached from the ELB. This is to allow the instance time to be drained from incoming TCP connections before terminating it, during a scale down operation.
* `lifetime_period` - (Optional) TBD
* `utilize_reserved_instances` - (Optional) TBD

* `health_check_type` - (Optional) TBD
* `health_check_grace_period` - (Optional) TBD
* `health_check_unhealthy_duration_before_replacement` - (Optional) TBD

* `tags` - (Optional) A mapping of tags to assign to the resource.
* `elastic_ips` - (Optional) A list of [AWS Elastic IP](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html) allocation IDs to associate to the group instances.

* `elastic_load_balancers` - (Optional) Registers each instance with the specified Elastic Load Balancers (ELB).
* `target_group_arns` - (Optional) TBD
* `multai_target_sets` - (Optional) TBD
    * `target_set_id` - (Required) TBD
    * `balancer_id` - (Required) TBD
    
* `revert_to_spot` - (Optional) Hold settings for strategy correction – replacing On-Demand for Spot instances. Supported Values: `"never"`, `"always"`, `"timeWindow"`
    * `perform_at` - (Required) TBD
    * `time_windows` - (Optional) TBD

<a id="signal"></a>
## Signals

Each `signal` supports the following:

* `name` - (Required) The name of the signal defined for the group. Valid Values: `"INSTANCE_READY"`, `"INSTANCE_READY_TO_SHUTDOWN"`
* `timeout` - (Optional) TBD

<a id="scheduled-task"></a>
## Scheduled Tasks

Each `scheduled_task` supports the following:

* `task_type` - (Required) The task type to run. Supported task types are: `"scale"`, `"backup_ami"`, `"roll"`, `"scaleUp"`, `"percentageScaleUp"`, `"scaleDown"`, `"percentageScaleDown"`, `"statefulUpdateCapacity"`.
* `cron_expression` - (Optional; Required if not using `frequency`) A valid cron expression. The cron is running in UTC time zone and is in [Unix cron format](https://en.wikipedia.org/wiki/Cron).
* `start_time` - (Optional; TBD
* `frequency` - (Optional; Required if not using `cron_expression`) The recurrence frequency to run this task. Supported values are `"hourly"`, `"daily"`, `"weekly"` and `"continuous"`.
* `scale_target_capacity` - (Optional) The desired number of instances the group should have.
* `scale_min_capacity` - (Optional) The minimum number of instances the group should have.
* `scale_max_capacity` - (Optional) The maximum number of instances the group should have.
* `is_enabled` - (Optional, Default: `false`) TBD
* `batch_size_percentage` - (Optional) TBD
* `grace_period` - (Optional) TBD
* `target_capacity` - (Optional) TBD
* `min_capacity` - (Optional) TBD
* `max_capacity` - (Optional) TBD

<a id="scaling-policy"></a>
## Scaling Policies

Each `scaling_*_policy` supports the following:

* `namespace` - (Required) The namespace for the alarm's associated metric.
* `metric_name` - (Required) The name of the metric, with or without spaces.
* `threshold` - (Required) The value against which the specified statistic is compared.
* `policy_name` - (Required) The name of the policy.
* `statistic` - (Optional, Default: `"average"`) The metric statistics to return. For information about specific statistics go to [Statistics](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/index.html?CHAP_TerminologyandKeyConcepts.html#Statistic) in the Amazon CloudWatch Developer Guide.
* `unit` - (Optional, Default: `"percent"`) The unit for the alarm's associated metric.
* `period` - (Optional, Default: `300`) The granularity, in seconds, of the returned datapoints. Period must be at least 60 seconds and must be a multiple of 60.
* `evaluation_periods` - (Optional, Default: `1`) The number of periods over which data is compared to the specified threshold.
* `cooldown` - (Optional, Default: `300`) The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start. If this parameter is not specified, the default cooldown period for the group applies.
* `dimensions` - (Optional) A mapping of dimensions describing qualities of the metric.
* `operator` - (Optional, Scale Up Default: `gte`, Scale Down Default: `lte`) TBD
* `source` - (Optional) TBD

* `action_type` - (Optional; if not using `min_target_capacity` or `max_target_capacity`) The type of action to perform for scaling. Valid values: `"adjustment"`, `"percentageAdjustment"`, `"setMaxTarget"`, `"setMinTarget"`, `"updateCapacity"`.

If you do not specify an action type, you can only use – `adjustment`, `minTargetCapacity`, `maxTargetCapacity`.
While using action_type, please also set the following:

When using `adjustment`           – set the field `adjustment`
When using `percentageAdjustment` - set the field `adjustment`
When using `setMaxTarget`         – set the field `max_target_capacity`
When using `setMinTarget`         – set the field `min_target_capacity`
When using `updateCapacity`       – set the fields `minimum`, `maximum`, and `target`

* `adjustment` - (Optional; if not using `min_target_capacity` or `max_target_capacity`;) The number of instances to add/remove to/from the target capacity when scale is needed. Can be used as advanced expression for scaling of instances to add/remove to/from the target capacity when scale is needed. You can see more information here: Advanced expression. Example value: `"MAX(currCapacity / 5, value * 10)"`
* `min_target_capacity` - (Optional; if not using `adjustment`; available only for scale up). The number of the desired target (and minimum) capacity
* `max_target_capacity` - (Optional; if not using `adjustment`; available only for scale down). The number of the desired target (and maximum) capacity

* `minimum` - (Optional; if using `updateCapacity`) The minimal number of instances to have in the group.
* `maximum` - (Optional; if using `updateCapacity`) The maximal number of instances to have in the group.
* `target` - (Optional; if using `updateCapacity`) The target number of instances to have in the group.

<a id="network-interface"></a>
## Network Interfaces

Each of the `network_interface` attributes controls a portion of the AWS
Instance's "Elastic Network Interfaces". It's a good idea to familiarize yourself with [AWS's Elastic Network
Interfaces docs](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html)
to understand the implications of using these attributes.

* `network_interface_id` - (Optional) The ID of the network interface.
* `device_index` - (Required) The index of the device on the instance for the network interface attachment.
* `description` - (Required) The description of the network interface.
* `private_ip_address` - (Optional) The private IP address of the network interface.
* `delete_on_termination` - (Optional) If set to true, the interface is deleted when the instance is terminated.
* `secondary_private_ip_address_count` - (Optional) The number of secondary private IP addresses.
* `associate_public_ip_address` - (Optional) Indicates whether to assign a public IP address to an instance you launch in a VPC. The public IP address can only be assigned to a network interface for eth0, and can only be assigned to a new network interface, not an existing one.

<a id="block-devices"></a>
## Block Devices

Each of the `*_block_device` attributes controls a portion of the AWS
Instance's "Block Device Mapping". It's a good idea to familiarize yourself with [AWS's Block Device
Mapping docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
to understand the implications of using these attributes.

Each `ebs_block_device` supports the following:

* `device_name` - (Required) The name of the device to mount.
* `snapshot_id` - (Optional) The Snapshot ID to mount.
* `volume_type` - (Optional, Default: `"standard"`) The type of volume. Can be `"standard"`, `"gp2"`, `"io1"`, `"st1"` or `"sc1"`.
* `volume_size` - (Optional) The size of the volume in gigabytes.
* `iops` - (Optional) The amount of provisioned [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html). This must be set with a `volume_type` of `"io1"`.
* `delete_on_termination` - (Optional) Whether the volume should be destroyed on instance termination.
* `encrypted` - (Optional) Enables [EBS encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) on the volume.

Modifying any `ebs_block_device` currently requires resource replacement.

Each `ephemeral_block_device` supports the following:

* `device_name` - (Required) The name of the block device to mount on the instance.
* `virtual_name` - (Required) The [Instance Store Device Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames)
  (e.g. `"ephemeral0"`).

<a id="stateful"></a>
## Stateful

We support instance persistence via the following configurations. all values are boolean.
For more information on instance persistence please see: [Stateful configuration](https://help.spotinst.com/hc/en-us/articles/115002899285)

* `persist_root_device` - (Optional) Boolean, should the instance maintain its root device volumes
* `persist_block_devices` - (Optional) Boolean, should the instance maintain its Data volumes
* `persist_private_ip` - (Optional) Boolean, should the instance maintain its private IP
* `block_devices_mode` - (Optional) String, determine the way we attach the data volumes to the data devices, possible values: `"reattach"` and `"onLaunch"` (default is onLaunch)
* `private_ips` - (Optional) List of Private IPs to associate to the group instances.(e.g. "172.1.1.0"). Please note: This setting will only apply if persistence.persist_private_ip is set to true.

<a id="health-check"></a>
## Health Check

* `health_check_type` - (Optional) The service that will perform health checks for the instance. Supported values : `"ELB"`, `"HCS"`, `"TARGET_GROUP"`, `"CUSTOM"`, `"K8S_NODE"`, `"MLB"`, `"EC2"`, `"MULTAI_TARGET_SET"`, `"MLB_RUNTIME"`, `"K8S_NODE"`, `"NOMAD_NODE"`, `"ECS_CLUSTER_INSTANCE"`.
* `health_check_grace_period` - (Optional) The amount of time, in seconds, after the instance has launched to starts and check its health
* `health_check_unhealthy_duration_before_replacement` - (Optional) The amount of time, in seconds, that we will wait before replacing an instance that is running and became unhealthy (this is only applicable for instances that were once healthy)

<a id="third-party-integrations"></a>
## Third-Party Integrations

* `rancher_integration` - (Optional) Describes the [Rancher](http://rancherlabs.com/) integration.

    * `master_host` - (Required) The URL of the Rancher Master host.
    * `access_key` - (Required) The access key of the Rancher API.
    * `secret_key` - (Required) The secret key of the Rancher API.

* `elastic_beanstalk_integration` - (Optional) Describes the [Elastic Beanstalk](https://aws.amazon.com/documentation/elastic-beanstalk/?id=docs_gateway) integration.

    * `environment_id` - (Required) The ID of the Elastic Beanstalk environment.

* `integration_ecs` - (Optional) Describes the [EC2 Container Service](https://aws.amazon.com/documentation/ecs/?id=docs_gateway) integration.

    * `cluster_name` - (Required) The name of the EC2 Container Service cluster.
    * `autoscale_is_enabled` - (Optional, Default: `false`) Specifies whether the auto scaling feature is enabled.
    * `autoscale_cooldown` - (Optional, Default: `300`) The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start.
    * `autoscale_headroom` - (Optional) TBD
        * `cpu_per_unit` - (Optional, Default: `0`) TBD
        * `memory_per_unit` - (Optional, Default: `0`) TBD
        * `num_of_units` - (Optional, Default: `0`) TBD
    * `autoscale_down` - (Optional) TBD
        * `evaluation_periods` - (Optional, Default: `5`) TBD

* `integration_codedeploy` - (Optional) Describes the [Code Deploy](https://aws.amazon.com/documentation/codedeploy/?id=docs_gateway) integration.

    * `cleanup_on_failure` - (Optional) TBD
    * `terminate_instance_on_failure` - (Optional) TBD
    * `deployment_groups` - (Optional) TBD
        * `application_name` - (Optional) TBD
        * `deployment_group_name` - (Optional) TBD

* `integration_kubernetes` - (Optional) Describes the [Kubernetes](https://kubernetes.io/) integration.

    * `integration_mode` - (Required) Valid values: `"saas"`, `"pod"`.
    * `cluster_identifier` - (Required; if using integration_mode as pod)
    * `api_server` - (Required; if using integration_mode as saas)
    * `token` - (Required; if using integration_mode as saas) TBD
    * `autoscale_is_enabled` - (Optional, Default: `false`) Specifies whether the auto scaling feature is enabled.
    * `autoscale_is_auto_config` - (Optional, Default: `false`) TBD
    * `autoscale_cooldown` - (Optional, Default: `300`) The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start.
    * `autoscale_headroom` - (Optional) TBD
        * `cpu_per_unit` - (Optional, Default: `0`) TBD
        * `memory_per_unit` - (Optional, Default: `0`) TBD
        * `num_of_units` - (Optional, Default: `0`) TBD
    * `autoscale_down` - (Optional) TBD
        * `evaluation_periods` - (Optional, Default: `5`) TBD
 
 * `integration_nomad` - (Optional) Describes the [Nomad](https://www.nomadproject.io/) integration.
 
     * `master_host` - (Required) TBD
     * `master_port` - (Required) TBD
     * `acl_token` - (Required) TBD
     * `autoscale_is_enabled` - (Optional, Default: `false`) Specifies whether the auto scaling feature is enabled.
     * `autoscale_cooldown` - (Optional, Default: `300`) The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start.
     * `autoscale_headroom` - (Optional) TBD
         * `cpu_per_unit` - (Optional, Default: `0`) TBD
         * `memory_per_unit` - (Optional, Default: `0`) TBD
         * `num_of_units` - (Optional, Default: `0`) TBD
     * `autoscale_down` - (Optional) TBD
         * `evaluation_periods` - (Optional, Default: `5`) TBD
         
 * `integration_mesosphere` - (Optional) Describes the [Mesosphere](https://mesosphere.com/) integration.
 
     * `api_server` - (Optional) TBD

 * `integration_multai_runtime` - (Optional) Describes the [Multai Runtime](https://spotinst.com/) integration.
 
     * `deployment_id` - (Optional) TBD
     
<a id="update-policy"></a>
## Update Policy

* `update_policy` - (Optional) TBD
    * `should_resume_stateful` - (Required) TBD
    * `should_roll` - (Required) TBD
    * `roll_config` - (Required) TBD
        * `batch_size_percentage` - (Required) TBD
        * `health_check_type` - (Optional) TBD
        * `grace_period` - (Optional) TBD
       
## Attributes Reference

The following attributes are exported:

* `id` - The group ID.
