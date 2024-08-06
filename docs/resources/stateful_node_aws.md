---
layout: "spotinst"
page_title: "Spotinst: stateful_node_aws"
subcategory: "Stateful Node"
description: |-
  Provides a Spotinst AWS managed instance resource.
---

# spotinst\_managed_instance\_aws

Provides a Spotinst AWS ManagedInstance resource.

## Example Usage

```hcl
# Create a Manged Instance
resource "spotinst_managed_instance_aws" "default-managed-instance" {
  name        = "default-managed-instance"
  description = "created by Terraform"
  product     = "Linux/UNIX"

  region     = "us-west-2"
  subnet_ids = ["subnet-123"]
  vpc_id     = "vpc-123"

  life_cycle                 = "on_demand"
  orientation                = "balanced"
  draining_timeout           = "120"
  fallback_to_ondemand       = false
  utilize_reserved_instances = true
  optimization_windows       = ["Mon:03:00-Wed:02:20"]
  auto_healing               = "true"
  grace_period               = "180"
  unhealthy_duration         = "60"
  minimum_instance_lifetime  = "1"
  revert_to_spot {
    perform_at = "always"
  }

  persist_private_ip    = "false"
  persist_block_devices = "true"
  persist_root_device   = "true"
  block_devices_mode    = "reattach"
  health_check_type     = "EC2"

  elastic_ip = "ip"
  private_ip = "ip"

  instance_types = [
    "t1.micro",
    "t3.medium",
    "t3.large",
    "t2.medium",
    "t2.large",
    "z1d.large",
  ]

  preferred_type       = "t1.micro"
  ebs_optimized        = "true"
  enable_monitoring    = "true"
  placement_tenancy    = "default"
  image_id             = "ami-1234"
  iam_instance_profile = "iam-profile"
  security_group_ids   = ["sg-234"]
  key_pair             = "labs-oregon"
  user_data            = "managed instance hello world"
  shutdown_script      = "managed instance bye world"
  cpu_credits          = "standard"

  tags {
    key   = "explicit1"
    value = "value1"
  }

  tags {
    key   = "explicit2"
    value = "value2"
  }

  block_device_mappings {
    device_name = "/dev/xvdcz"
    ebs {
      delete_on_termination = "true"
      volume_type           = "gp3"
      volume_size           = 50
      iops                  = 100
      throughput            = 125
      encrypted             = true
      kms_key_id            = "kms-key-01"
      snapshot_id           = "snapshot_id"
    }
  }

  resource_tag_specification {
    should_tag_enis = true
    should_tag_volumes = true
    should_tag_snapshots = true
    should_tag_amis = true
  }
  
  metadata_options {
    http_put_response_hop_limit = 33
    http_tokens = "required"
    instance_metadata_tags = "enabled"
  }

}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The ManagedInstance name.
* `description` - (Optional) The ManagedInstance description.
* `region` - (Required) The AWS region your group will be created in.
* `life_cycle` - (Optional) Set lifecycle, valid values: `"spot"`, `"on_demand"`. Default `"spot"`.
* `orientation` - (Optional) Select a prediction strategy. Valid values: `"balanced"`, `"costOriented"`, `"availabilityOriented"`, `"cheapest"`. Default: `"availabilityOriented"`.    
* `draining_timeout` - (Optional) The time in seconds to allow the instance be drained from incoming TCP connections and detached from ELB before terminating it, during a scale down operation.
* `fallback_to_ondemand` - (Optional) In case of no spots available, Managed Instance will launch an On-demand instance instead. Default: `"true"`.
* `utilize_reserved_instances` - (Optional) In case of any available Reserved Instances, Managed Instance will utilize them before purchasing Spot instances. Default: `"false"`. 
* `optimization_windows` - (Optional) When `performAt` is `"timeWindow"`: must specify a list of `"timeWindows"` with at least one time window. Each string should be formatted as `ddd:hh:mm-ddd:hh:mm` (ddd = day of week = Sun | Mon | Tue | Wed | Thu | Fri | Sat hh = hour 24 = 0 -23 mm = minute = 0 - 59).
* `perform_at` - (Optional) Valid values: `"always"`, `"never"`, `"timeWindow"`. Default `"never"`.
* `minimum_instance_lifetime` - (Optional) Defines the preferred minimum instance lifetime. Markets which comply with this preference will be prioritized. Optional values: `1`, `3`, `6`, `12`, `24`.
* `persist_private_ip` - (Optional) Should the instance maintain its private IP.
* `persist_block_devices` - (Optional) Should the instance maintain its Data volumes. 
* `persist_root_device` - (Optional) Should the instance maintain its root device volumes.
* `block_devices_mode` - (Optional) Determine the way we attach the data volumes to the data devices. Valid values: `"reattach"`, `"onLaunch"`. Default: `"onLaunch"`.
* `health_check_type` - (Optional) The service to use for the health check. Valid values: `"EC2"`, `"ELB"`, `"TARGET_GROUP"`. Default: `"EC2"`. 
* `auto_healing` - (Optional) Enable the auto healing which auto replaces the instance in case the health check fails, default: `"true"`. 
* `grace_period` - (Optional) The amount of time, in seconds, after the instance has launched to starts and check its health, default `"120"`.
* `unhealthy_duration` - (Optional) The amount of time, in seconds, an existing instance should remain active after becoming unhealthy. After the set time out the instance will be replaced, default `"120"`.
* `subnet_ids` - (Required) A comma-separated list of subnet identifiers for your instance.
* `vpcId` - (Required) VPC id for your instance.
* `elastic_ip` - (Optional) Elastic IP Allocation Id to associate to the instance.
* `private_ip` - (Optional) Private IP Allocation Id to associate to the instance. 
* `product` - (Required) Operation system type. Valid values: `"Linux/UNIX"`, `"SUSE Linux"`, `"Windows"`, `"Red Hat Enterprise Linux"`, `"Linux/UNIX (Amazon VPC)"`, `"SUSE Linux (Amazon VPC)"`, `"Windows (Amazon VPC)"`,  `"Red Hat Enterprise Linux (Amazon VPC)"`.    
* `instance_types` - (Required) Comma separated list of available instance types for instance.
* `preferred_type` - (Required) Preferred instance types for the instance. We will automatically select optional similar instance types to ensure optimized cost efficiency
* `ebs_optimized` - (Optional, Default: `false`) Enable EBS optimization for supported instances. Note: Additional charges will be applied by the Cloud Provider.
Default: false
* `enable_monitoring` - (Optional) Describes whether instance Enhanced Monitoring is enabled.
Default: false
* `placement_tenancy` - (Optional) Valid values: `"default"`, `"dedicated"`.
Default: default
* `iam_instance_profile` - (Optional) Set IAM profile to instance. Set only one of ARN or Name.
* `security_group_ids` - (Optional) One or more security group IDs.
* `image_id` - (Required) The ID of the image used to launch the instance.
* `key_pair` - (Optional) Specify a Key Pair to attach to the instances.
* `tags` - (Optional) Set tags for the instance. Items should be unique.
     * `key` - Tag's key.
     * `value` - Tag's name.
* `user_data` - (Optional) The Base64-encoded MIME user data to make available to the instances.
* `shutdown_script` - (Optional) The Base64-encoded shutdown script to execute prior to instance termination.
* `cpu_credits` - (Optional) cpuCredits can have one of two values: `"unlimited"`, `"standard"`.

<a id="block-device-mapping"></a>
## Block Device Mapping

* `block_device_mappings` - (Optional) Attributes controls a portion of the AWS:
    * `device_name` - (Required) The name of the device to mount.
    * `ebs` - (Required) Object
        * `volume_type` - (Optional, Default: `"standard"`) The type of volume. Can be `"standard"`, `"gp2"`, `"gp3"`, `"io1"`, `"st1"` or `"sc1"`.
        * `volume_size` - (Optional) The size of the volume, in GiBs.
        * `iops` - (Optional) The amount of provisioned [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html). This must be set with a `volume_type` of `"io1"`.
        * `delete_on_termination` - (Optional) Whether the volume should be destroyed on instance termination.
        * `throughput`- (Optional) The throughput that the volume supports, in MiB/s. Minimum value of 125. Maximum value of 1000. Valid only if `volume_type` is set to `"gp3"`.
        * `encrypted` - (Optional) Enables [EBS encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) on the volume.
        * `kms_key_id` - (Optional) ID for a user managed CMK under which the EBS Volume is encrypted.
        * `snapshot_id` - (Optional) The Snapshot ID to mount.

Usage:

```hcl
block_device_mappings {
  device_name                 = "Name"
  ebs {
    delete_on_termination     = true
    iops                      = 16000
    volume_type               = "gp3"
    volume_size               = 100
    throughput                = 1000
    encrypted                 = false
    kms_key_id                = "kms-key-01"
    snapshot_id               = "snapshot_id"
  }
}
```      

<a id="resource-tag-specification"></a>
## Resource Tag Specification

* `resource_tag_specification` - (Optional) User will specify which resources should be tagged with group tags.
    * `should_tag_enis`      - (Optional) Tag specification for ENI resources.
    * `should_tag_volumes`   - (Optional) Tag specification for Volume resources.
    * `should_tag_snapshots` - (Optional) Tag specification for Snapshot resources.
    * `should_tag_amis`      - (Optional) Tag specification for AMI resources.

Usage:

```hcl
resource_tag_specification {
    should_tag_enis           = true
    should_tag_volumes        = true
    should_tag_snapshots      = true
    should_tag_amis           = true
}
```    

<a id="metadata options"></a>
## Metadata Options

* `metadata_options` - (Optional) Optional field. Used for IMDSv2.
    * `http_put_response_hop_limit` - (Optional) An integer within 1,64 range (including). The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further instance metadata requests can travel.
    * `http_tokens` - Required for IMDSv2. Can’t be null in Create or Update. Values can be `optional` or `required`.
    * `instance_metadata_tags` - (Optional) Valid values: `enabled` or `disabled` Indicates whether access to instance tags from the instance metadata is enabled or disabled. Can’t be null.

Usage:

```hcl
metadata_options {
    http_put_response_hop_limit = 33
    http_tokens = "optional"
    instance_metadata_tags = "enabled"
}
```    

<a id="network-interface"></a>
## Network Interface 

* `network-interface`- (Optional) List of network interfaces in an EC2 instance.
    * `device_index` - (Optional) The position of the network interface in the attachment order. A primary network interface has a device index of 0. If you specify a network interface when launching an instance, you must specify the device index.
    * `associate_public_ip_address` - (Optional) Indicates whether to assign a public IPv4 address to an instance you launch in a VPC. The public IP address can only be assigned to a network interface for eth0, and can only be assigned to a new network interface, not an existing one. You cannot specify more than one network interface in the request. If launching into a default subnet, the default value is true.
    * `associate_ipv6_address` - (Optional) Indicates whether to assign an IPv6 address. Amazon EC2 chooses the IPv6 addresses from the range of the subnet. Default: `false`

Usage:

```hcl
network_interface {
    device_index                = 0
    associate_public_ip_address = false
    associate_ipv6_address      = true
}
```       

<a id="scheduled-task"></a>
## Scheduled Tasks

Each `scheduled_task` supports the following:

* `is_enabled` - (Optional) Describes whether the task is enabled. When true the task should run when false it should not run.
* `frequency` - (Optional) Set frequency for the task. Valid values: "hourly", "daily", "weekly", "continuous".
* `start_time` - (Optional) DATETIME in ISO-8601 format. Sets a start time for scheduled actions. If "frequency" or "cronExpression" are not used - the task will run only once at the start time and will then be deleted from the instance configuration.
   Example: `"2019-05-23T10:55:09Z"`
* `cron_expression` - (Optional) A valid cron expression. The cron is running in UTC time zone and is in Unix cron format Cron Expression Validator Script. Only one of ‘frequency’ or ‘cronExpression’ should be used at a time.
   Example: `"0 1 * * *"`.
* `task_type`- (Required) The task type to run. Valid values: `"pause"`, `"resume"`, `"recycle"`.

Usage:

```hcl
scheduled_task {
    is_enabled      = "true"
    frequency       = "hourly"
    start_time      = "2019-11-20T23:59:59Z"
    cron_expression = "* * * * *"
    task_type       = "pause"
}
```

<a id="load-balancers"></a>
## Load Balancers
* `load_balancers` - (Optional) List of load balancers configs.
    * `name` - The AWS resource name. Required for Classic Load Balancer. Optional for Application Load Balancer.
    * `arn` - The AWS resource ARN (Required only for ALB target groups).
    * `type` - The resource type. Valid Values: `"CLASSIC"`, `"TARGET_GROUP"`.

Usage:

```hcl
load_balancers {
    name          = "name"
    arn           = "arn"
    type          = "CLASSIC"
}
```

<a id="route53"></a>
## Route53
 
* `integration_route53` - (Optional) Describes the [Route53](https://aws.amazon.com/documentation/route53/?id=docs_gateway) integration.
   * `domains` - (Required) Route 53 Domain configurations.
       * `hosted_zone_id` - (Required) The Route 53 Hosted Zone Id for the registered Domain.
       * `spotinst_acct_id` - (Optional) The Spotinst account ID that is linked to the AWS account that holds the Route 53 hosted Zone Id. The default is the user Spotinst account provided as a URL parameter.
       * `record_set_type` - (Optional, Default: `a`) The type of the record set. Valid values: `"a"`, `"cname"`.
       * `record_sets` - (Required) List of record sets
           * `name` - (Required) The record set name.
           * `use_public_ip` - (Optional, Default: `false`) - Designates whether the IP address should be exposed to connections outside the VPC.
           * `use_public_dns` - (Optional, Default: `false`) - Designates whether the DNS address should be exposed to connections outside the VPC.

Usage:

```hcl
integration_route53 {

    # Option 1: Use A records.
    domains { 
      hosted_zone_id   = "zone-id"
      spotinst_acct_id = "act-123456"
      record_set_type  = "a"
    
      record_sets {
        name          = "foo.example.com"
        use_public_ip = true
      }
    }

    # Option 2: Use CNAME records.
    domains { 
      hosted_zone_id   = "zone-id"
      spotinst_acct_id = "act-123456"
      record_set_type  = "cname"
    
      record_sets {
        name           = "foo.example.com"
        use_public_dns = true
      }
    }
}
```

<a id="managed_instance_action"></a>
## Managed Instance Action

* `managed_instance_action` - (Optional)
    * `type` - (Required) String, Action type. Supported action types: `pause`, `resume`, `recycle`.

Usage:

```hcl
managed_instance_action { 
  type  = "pause"
}    
```

<a id="delete"></a>
## Delete

* `delete` - (Optional) Specify deallocation parameters for stateful node deletion. If delete configuration is not present, all below parameters will be set to "true" by default.
   * `ami_backup_should_delete_images` - (Optional) Mark if images collected during AMI Auto Backup should be deleted during instance deletion.
   * `deallocation_config_should_delete_images` - (Optional) Flag to remove persistent image.
   * `should_delete_network_interfaces` - (Optional) Flag to remove network interfaces.
   * `should_delete_snapshots` - (Optional) Flag to remove snapshots.
   * `should_delete_volumes` - (Optional) Flag to remove persistent volumes.
   * `should_terminate_instance` - (Optional) Flag to detach a machine without terminating it on AWS side. This will enable rollback to AWS.

Usage:

```hcl
delete {
  ami_backup_should_delete_images = true
  deallocation_config_should_delete_images = true
  should_delete_network_interfaces = true
  should_delete_snapshots = true
  should_delete_volumes = true
  should_terminate_instance = true
}
```

## Attributes Reference

The following attributes are exported:

* `id` - The group ID.