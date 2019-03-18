---
layout: "spotinst"
page_title: "Spotinst: Mr Scaler"
sidebar_current: "docs-do-resource-mr_scaler"
description: |-
  Provides a Spotinst MrScaler resource.
---

# spotinst\_mrscaler

Provides a Spotinst AWS MrScaler resource.

## Example Usage - New Strategy

```hcl
# Create a Mr Scaler with New strategy

resource "spotinst_mrscaler_aws" "Terraform-MrScaler-01" {
  name        = "Terraform-MrScaler-01"
  description = "Testing MrScaler creation via Terraform"
  region      = "us-west-2"
  strategy    = "new"
  release_label = "emr-5.17.0"
  
  availability_zones = ["us-west-2a:subnet-123456"]
  
  provisioning_timeout = {
    timeout        = 15
    timeout_action = "terminate"
  }
  
// --- CLUSTER ------------
  log_uri         = "s3://example-logs"
  additional_info = "{'test':'more information'}"
  job_flow_role   = "EMR_EC2_ExampleRole"
  security_config = "example-config"
  service_role    = "example-role"
  
  termination_protected = false
  keep_job_flow_alive   = true
// -------------------------

// --- OPTONAL COMPUTE -----
  custom_ami_id        = "ami-123456"
  repo_upgrade_on_boot = "NONE"
  ec2_key_name         = "test-key"

  managed_primary_security_group = "sg-123456"
  managed_replica_security_group = "sg-987654"
  service_access_security_group  = "access-example"

  additional_primary_security_groups = ["sg-456321"]
  additional_replica_security_groups = ["sg-123654"]

  applications = [
    {
      name = "Ganglia"
      version = "1.0"
    },
    {
      name = "Hadoop"
    },
    {
      name = "Pig"
      args = ["fake", "args"]
    }
  ]

  steps_file = {
    bucket = "example-bucket"
    key = "steps.json"
  }

  configurations_file = {
    bucket = "example-bucket"
    key = "configurations.json"
  }

  bootstrap_actions_file = {
    bucket = "terraform-emr-test"
    key = "bootstrap-actions.json"
  }
// -------------------------
  
// --- MASTER GROUP -------------
  master_instance_types = ["c3.xlarge"]
  master_lifecycle      = "SPOT"
  master_ebs_optimized  = true
  
  master_ebs_block_device = {
    volumes_per_instance = 1
    volume_type          = "gp2"
    size_in_gb           = 30
  }
// ------------------------------

// --- CORE GROUP -------------
  core_instance_types = ["c3.xlarge", "c4.xlarge"]
  core_min_size         = 1
  core_max_size         = 1
  core_desired_capacity = 1
  core_lifecycle        = "ON_DEMAND"
  core_ebs_optimized    = false
  
  core_ebs_block_device = {
    volumes_per_instance = 2
    volume_type          = "gp2"
    size_in_gb           = 40
  }
// ----------------------------

// --- TASK GROUP -------------
  task_instance_types = ["c3.xlarge", "c4.xlarge"]
  task_min_size         = 0
  task_max_size         = 30
  task_desired_capacity = 1
  task_lifecycle        = "SPOT"
  task_ebs_optimized    = false
  
  task_ebs_block_device = {
    volumes_per_instance = 2
    volume_type          = "gp2"
    size_in_gb           = 40
  }
// ----------------------------

// --- TAGS -------------------
  tags = [{
      key   = "Creator"
      value = "Terraform"
  }]
// ----------------------------
```

## Example Usage - Clone Strategy

```hcl
# Create a Mr Scaler with Clone strategy and Task scaling

output "mrscaler-name" {
  value = "${spotinst_mrscaler_aws.Terraform-MrScaler-01.name}"
}

output "mrscaler-created-cluster-id" {
  value = "${spotinst_mrscaler_aws.Terraform-MrScaler-01.output_cluster_id}"
}

resource "spotinst_mrscaler_aws" "Terraform-MrScaler-01" {
  name        = "Terraform-MrScaler-01"
  description = "Testing MrScaler creation via Terraform"
  region      = "us-west-2"
  strategy    = "clone"
  
  cluster_id        = "j-123456789"
  expose_cluster_id = true

  availability_zones = ["us-west-2a:subnet-12345678"]

// --- MASTER GROUP -------------
  master_instance_types = ["c3.xlarge"]
  master_lifecycle      = "SPOT"
  master_ebs_optimized  = true
  
  master_ebs_block_device = {
    volumes_per_instance = 1
    volume_type          = "gp2"
    size_in_gb           = 30
  }
// ------------------------------

// --- CORE GROUP -------------
  core_instance_types   = ["c3.xlarge", "c4.xlarge"]
  core_min_size         = 1
  core_max_size         = 1
  core_desired_capacity = 1
  core_lifecycle        = "ON_DEMAND"
  core_ebs_optimized    = false
  
  core_ebs_block_device = {
    volumes_per_instance = 2
    volume_type          = "gp2"
    size_in_gb           = 40
  }
// ----------------------------

// --- TASK GROUP -------------
  task_instance_types   = ["c3.xlarge", "c4.xlarge"]
  task_min_size         = 0
  task_max_size         = 30
  task_desired_capacity = 1
  task_lifecycle        = "SPOT"
  task_ebs_optimized    = false
  
  task_ebs_block_device = {
    volumes_per_instance = 2
    volume_type          = "gp2"
    size_in_gb           = 40
  }
// ----------------------------

// --- TAGS -------------------
  tags = [{
      key   = "Creator"
      value = "Terraform"
  }]
// ----------------------------

// --- TASK SCALING POLICY ------
  task_scaling_down_policy = [{
    policy_name = "policy-name"
    metric_name = "CPUUtilization"
    namespace   = "AWS/EC2"
    statistic   = "average"
    unit        = ""
    threshold   = 10
    adjustment  = "1"
    cooldown    = 60
    dimensions = {
      name  = "name-1"
      value = "value-1"
    }
 
    operator           = "gt"
    evaluation_periods = 10
    period             = 60

    action_type = ""
    minimum     = 0
    maximum     = 10
    target      = 5
    max_target_capacity = 1
  }]
// ----------------------------
```

## Example Usage - Wrap Strategy

```hcl
# Create a Mr Scaler with Wrap strategy

resource "spotinst_mrscaler" "example-scaler-2" {
  name        = "spotinst-mr-scaler-2"
  description = "created by Terraform"
  region      = "us-west-2"
  strategy    = "wrap"
  cluster_id  = "j-27UVDEHXL4OQM"
  
// --- TASK GROUP -------------
  task_instance_types = ["c3.xlarge","c4.xlarge"]
  
  task_target    = 2
  task_minimum   = 0
  task_maximum   = 4
  task_lifecycle = "SPOT"
  
  task_ebs_block_device = {
    volumes_per_instance = 1
    volume_type          = "gp2"
    size_in_gb           = 20
  }
// ----------------------------
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The MrScaler name.
* `description` - (Optional) The MrScaler description.
* `region` - (Required) The MrScaler region.
* `strategy` - (Required) The MrScaler strategy. Allowed values are `new` `clone` and `wrap`.
* `cluster_id` - (Optional) The MrScaler cluster id.
* `expose_cluster_id` - (Optional) Allow the `cluster_id` to set a Terraform output variable.

<a id="provisioning-timeout"></a>
## Provisioning Timeout (Clone, New strategies)
* `timeout` - (Optional) The amount of time (minutes) after which the cluster is automatically terminated if it's still in provisioning status. Minimum: '15'.
* `timeout_action` - (Optional) The action to take if the timeout is exceeded. Valid values: `terminate`, `terminateAndRetry`.

<a id="cluster-config"></a>
## Cluster Configuration (New strategy only)
* `log_uri` - (Optional) The path to the Amazon S3 location where logs for this cluster are stored.
* `additional_info` - (Optional) This is meta information about third-party applications that third-party vendors use for testing purposes.
* `security_config` - (Optional) The name of the security configuration applied to the cluster.
* `service_role` - (Optional) The IAM role that will be assumed by the Amazon EMR service to access AWS resources on your behalf.
* `job_flow_role` - (Optional) The IAM role that was specified when the job flow was launched. The EC2 instances of the job flow assume this role.
* `termination_protected` - (Optional) Specifies whether the Amazon EC2 instances in the cluster are protected from termination by API calls, user intervention, or in the event of a job-flow error.
* `keep_job_flow_alive` - (Optional) Specifies whether the cluster should remain available after completing all steps.

<a id="task-group"></a>
## Task Group (Wrap, Clone, and New strategies)
* `task_instance_types` - (Required) The MrScaler instance types for the task nodes.
* `task_target` - (Required) amount of instances in task group.
* `task_maximum` - (Optional) maximal amount of instances in task group.
* `task_minimum` - (Optional) The minimal amount of instances in task group.
* `task_lifecycle` - (Required) The MrScaler lifecycle for instances in task group. Allowed values are 'SPOT' and 'ON_DEMAND'.
* `task_ebs_optimized` - (Optional) EBS Optimization setting for instances in group.
* `task_ebs_block_device` - (Required) This determines the ebs configuration for your task group instances. Only a single block is allowed.
    * `volumes_per_instance` - (Optional; Default 1) Amount of volumes per instance in the task group.
    * `volume_type` - (Required) volume type. Allowed values are 'gp2', 'io1' and others.
    * `size_in_gb` - (Required) Size of the volume, in GBs.
    * `iops` - (Optional) IOPS for the volume. Required in some volume types, such as io1.

<a id="core-group"></a>
## Core Group (Clone, New strategies)
* `core_instance_types` - (Required) The MrScaler instance types for the core nodes.
* `core_target` - (Required) amount of instances in core group.
* `core_maximum` - (Optional) maximal amount of instances in core group.
* `core_minimum` - (Optional) The minimal amount of instances in core group.
* `core_lifecycle` - (Required) The MrScaler lifecycle for instances in core group. Allowed values are 'SPOT' and 'ON_DEMAND'.
* `core_ebs_optimized` - (Optional) EBS Optimization setting for instances in group.
* `core_ebs_block_device` - (Required) This determines the ebs configuration for your core group instances. Only a single block is allowed.
    * `volumes_per_instance` - (Optional; Default 1) Amount of volumes per instance in the core group.
    * `volume_type` - (Required) volume type. Allowed values are 'gp2', 'io1' and others.
    * `size_in_gb` - (Required) Size of the volume, in GBs.
    * `iops` - (Optional) IOPS for the volume. Required in some volume types, such as io1.

<a id="master-group"></a>
## Master Group (Clone, New strategies)
* `master_instance_types` - (Required) The MrScaler instance types for the master nodes.
* `master_lifecycle` - (Required) The MrScaler lifecycle for instances in master group. Allowed values are 'SPOT' and 'ON_DEMAND'.
* `master_ebs_optimized` - (Optional) EBS Optimization setting for instances in group.
* `master_ebs_block_device` - (Required) This determines the ebs configuration for your master group instances. Only a single block is allowed.
    * `volumes_per_instance` - (Optional; Default 1) Amount of volumes per instance in the master group.
    * `volume_type` - (Required) volume type. Allowed values are 'gp2', 'io1' and others.
    * `size_in_gb` - (Required) Size of the volume, in GBs.
    * `iops` - (Optional) IOPS for the volume. Required in some volume types, such as io1.

<a id="tags"></a>
## Tags (Clone, New strategies)
* `tags` - (Optional) A list of tags to assign to the resource. You may define multiple tags.
    * `key` - (Required) Tag key.
    * `value` - (Required) Tag value.

<a id="Optional Compute Parameters"></a>  
## Optional Compute Parameters (New strategy)
* `managed_primary_security_group` - (Optional) EMR Managed Security group that will be set to the primary instance group.
* `managed_replica_security_group` - (Optional) EMR Managed Security group that will be set to the replica instance group.
* `service_access_security_group` - (Optional) The identifier of the Amazon EC2 security group for the Amazon EMR service to access clusters in VPC private subnets.
* `additional_primary_security_groups` - (Optional) A list of additional Amazon EC2 security group IDs for the master node.
* `additional_replica_security_groups` - (Optional) A list of additional Amazon EC2 security group IDs for the core and task nodes.
* `custom_ami_id` - (Optional) The ID of a custom Amazon EBS-backed Linux AMI if the cluster uses a custom AMI.
* `repo_upgrade_on_boot` - (Optional) Applies only when `custom_ami_id` is used. Specifies the type of updates that are applied from the Amazon Linux AMI package repositories when an instance boots using the AMI. Possible values include: `SECURITY`, `NONE`.
* `ec2_key_name` - (Optional) The name of an Amazon EC2 key pair that can be used to ssh to the master node.
* `applications` - (Optional) A case-insensitive list of applications for Amazon EMR to install and configure when launching the cluster
    * `args` - (Optional) Arguments for EMR to pass to the application.
    * `name` - (Required) The application name.
    * `version`- (Optional)T he version of the application.

<a id="availability-zone"></a>
## Availability Zones (Clone, New strategies)

* `availability_zones` - (Required in Clone) List of AZs and their subnet Ids. See example above for usage.

<a id="configurations"></a>
## Configurations (Clone, New strategies)

* `configurations_file` - (Optional) Describes path to S3 file containing description of configurations. [More Information](https://api.spotinst.com/elastigroup-for-aws/services-integrations/elastic-mapreduce/import-an-emr-cluster/advanced/)
    * `bucket` - (Required) S3 Bucket name for configurations.
    * `key`- (Required) S3 key for configurations.
    
<a id="steps"></a>
## Steps (Clone, New strategies)
* `steps_file` - (Optional) Steps from S3.
    * `bucket` - (Required) S3 Bucket name for steps.
    * `key`- (Required) S3 key for steps.
    
<a id="boostrap-actions"></a>
## Bootstrap Actions (Clone, New strategies)   
* `bootstrap_actions_file` - (Optional) Describes path to S3 file containing description of bootstrap actions. [More Information](https://api.spotinst.com/elastigroup-for-aws/services-integrations/elastic-mapreduce/import-an-emr-cluster/advanced/)
    * `bucket` - (Required) S3 Bucket name for bootstrap actions.
    * `key`- (Required) S3 key for bootstrap actions.

<a id="scaling-policy"></a>
## Scaling Policies

Possible task group scaling policies (Wrap, Clone, and New strategies):
* `task_scaling_up_policy`
* `task_scaling_down_policy`

Possible core group scaling policies (Clone, New strategies):
* `core_scaling_up_policy`
* `core_scaling_down_policy`

Each `*_scaling_*_policy` supports the following:

* `policy_name` - (Required) The name of the policy.
* `metric_name` - (Required) The name of the metric, with or without spaces.
* `statistic` - (Required) The metric statistics to return. For information about specific statistics go to [Statistics](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/index.html?CHAP_TerminologyandKeyConcepts.html#Statistic) in the Amazon CloudWatch Developer Guide.
* `unit` - (Required) The unit for the metric.
* `threshold` - (Required) The value against which the specified statistic is compared.
* `adjustment` - (Optional) The number of instances to add/remove to/from the target capacity when scale is needed.
* `min_target_capacity` - (Optional) Min target capacity for scale up.
* `max_target_capacity` - (Optional) Max target capacity for scale down.
* `namespace` - (Required) The namespace for the metric.
* `operator` - (Required) The operator to use. Allowed values are : 'gt', 'gte', 'lt' , 'lte'.
* `evaluation_periods` - (Required) The number of periods over which data is compared to the specified threshold.
* `period` - (Required) The granularity, in seconds, of the returned datapoints. Period must be at least 60 seconds and must be a multiple of 60.
* `cooldown` - (Required) The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
* `dimensions` - (Optional) A mapping of dimensions describing qualities of the metric.
* `minimum` - (Optional) The minimum to set when scale is needed.
* `maximum` - (Optional) The maximum to set when scale is needed.
* `target` - (Optional) The number of instances to set when scale is needed.
* `action_type` - (Required) The type of action to perform. Allowed values are : 'adjustment', 'setMinTarget', 'setMaxTarget', 'updateCapacity', 'percentageAdjustment'

<a id="scheduled-task"></a>
## Scheduled Tasks

* `scheduled_task` - (Optional) An array of scheduled tasks.
* `is_enabled` - (Optional) Enable/Disable the specified scheduling task.
* `task_type` - (Required) The type of task to be scheduled. Valid values: `setCapacity`.
* `instance_group_type` - (Required) Select the EMR instance groups to execute the scheduled task on. Valid values: `task`.
* `cron` - (Required) A cron expression representing the schedule for the task.
* `desired_capacity` - (Optional) New desired capacity for the elastigroup.
* `min_capacity` - (Optional) New min capacity for the elastigroup.
* `max_capacity` - (Optional) New max capacity for the elastigroup.

## Attributes Reference

The following attributes are exported:

* `id` - The scaler ID.