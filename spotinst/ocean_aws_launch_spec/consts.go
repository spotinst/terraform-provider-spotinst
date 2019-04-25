package ocean_aws_launch_spec

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

type LabelField string
type TaintField string
type IAMField string

const (
	LabelKey   LabelField = "key"
	LabelValue LabelField = "value"
)

const (
	TaintKey   TaintField = "key"
	TaintValue TaintField = "value"
	Effect     TaintField = "effect"
)

const (
	ARN  IAMField = "arn"
	Name IAMField = "name"
)

const (
	OceanID            commons.FieldName = "ocean_id"
	ImageID            commons.FieldName = "image_id"
	UserData           commons.FieldName = "user_data"
	IamInstanceProfile commons.FieldName = "iam_instance_profile"
	Labels             commons.FieldName = "labels"
	Taints             commons.FieldName = "taints"
)
