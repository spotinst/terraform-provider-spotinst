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
	SpotInstanceTypes        commons.FieldName = "instance_types_spot"
	Maintenance              commons.FieldName = "maintenance"
)
