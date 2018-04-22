package aws_elastigroup

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

type BalancerType string
const (
	BalancerTypeClassic         BalancerType = "CLASSIC"
	BalancerTypeTargetGroup     BalancerType = "TARGET_GROUP"
	BalancerTypeMultaiTargetSet BalancerType = "MULTAI_TARGET_SET"
)

const (
	Name        commons.FieldName = "name"
	Description commons.FieldName = "description"

	MaxSize         commons.FieldName = "max_size"
	MinSize         commons.FieldName = "min_size"
	DesiredCapacity commons.FieldName = "desired_capacity"
	CapacityUnit    commons.FieldName = "capacity_unit"

	HealthCheckGracePeriod                        commons.FieldName = "health_check_grace_period"
	HealthCheckType                               commons.FieldName = "health_check_type"
	HealthCheckUnhealthyDurationBeforeReplacement commons.FieldName = "health_check_unhealthy_duration_before_replacement"

	SubnetIds            commons.FieldName = "subnet_ids"
	AvailabilityZones    commons.FieldName = "availability_zones"
	ElasticLoadBalancers commons.FieldName = "elastic_load_balancers"
	TargetGroupArns      commons.FieldName = "target_group_arns"
	MultaiTargetSetIds   commons.FieldName = "multai_target_set_ids"
	Tags                 commons.FieldName = "tags"

	LaunchConfiguration  commons.FieldName = "launch_configuration"
	InstanceTypes        commons.FieldName = "instance_types"
	EbsBlockDevice       commons.FieldName = "ebs_block_device"
	EphemeralBlockDevice commons.FieldName = "ephemeral_block_device"
)
