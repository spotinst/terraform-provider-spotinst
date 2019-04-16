package ocean_aws_launch_spec

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

type LabelField string
type TaintField string

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
	OceanID  commons.FieldName = "ocean_id"
	ImageID  commons.FieldName = "image_id"
	UserData commons.FieldName = "user_data"
	Labels   commons.FieldName = "labels"
	Taints   commons.FieldName = "taints"
)
