package beanstalk_elastigroup

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "beanstalk_elastigroup_"
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
)
