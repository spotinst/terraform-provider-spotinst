package elastigroup_aws_beanstalk

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "elastigroup_aws_beanstalk"
)

const (
	Name                     commons.FieldName = "name"
	Region                   commons.FieldName = "region"
	Product                  commons.FieldName = "product"
	Minimum                  commons.FieldName = "min_size"
	Maximum                  commons.FieldName = "max_size"
	Target                   commons.FieldName = "desired_capacity"
	BeanstalkEnvironmentName commons.FieldName = "beanstalk_environment_name"
	BeanstalkEnvironmentId   commons.FieldName = "beanstalk_environment_id"
	SpotInstanceTypes        commons.FieldName = "instance_types_spot"
	Maintenance              commons.FieldName = "maintenance"
	ManagedActions           commons.FieldName = "managed_actions"
	PlatformUpdate           commons.FieldName = "platform_update"
	PerformAt                commons.FieldName = "perform_at"
	TimeWindow               commons.FieldName = "time_window"
	UpdateLevel              commons.FieldName = "update_level"
	DeploymentPreferences    commons.FieldName = "deployment_preferences"
	AutomaticRoll            commons.FieldName = "automatic_roll"
	BatchSizePercentage      commons.FieldName = "batch_size_percentage"
	GracePeriod              commons.FieldName = "grace_period"
	Strategy                 commons.FieldName = "strategy"
	Action                   commons.FieldName = "action"
	ShouldDrainInstances     commons.FieldName = "should_drain_instances"
)
