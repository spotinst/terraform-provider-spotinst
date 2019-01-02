package ocean_aws

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

type TagField string

const (
	TagKey   TagField = "key"
	TagValue TagField = "value"
)

const (
	Name                commons.FieldName = "name"
	ControllerClusterID commons.FieldName = "controller_id"

	MaxSize         commons.FieldName = "max_size"
	MinSize         commons.FieldName = "min_size"
	DesiredCapacity commons.FieldName = "desired_capacity"

	Region    commons.FieldName = "region"
	SubnetIds commons.FieldName = "subnet_ids"

	Tags commons.FieldName = "tags"
)
