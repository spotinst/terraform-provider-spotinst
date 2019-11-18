package ocean_ecs_launch_spec

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

type LabelField string
type IAMField string

const (
	CPUPerUnit    commons.FieldName = "cpu_per_unit"
	MemoryPerUnit commons.FieldName = "memory_per_unit"
	NumOfUnits    commons.FieldName = "num_of_units"
)

const (
	OceanID            commons.FieldName = "ocean_id"
	ImageID            commons.FieldName = "image_id"
	IamInstanceProfile commons.FieldName = "iam_instance_profile"
	UserData           commons.FieldName = "user_data"
	SecurityGroupIds   commons.FieldName = "security_group_ids"
	Name               commons.FieldName = "name"
	Attributes         commons.FieldName = "attributes"
	AttributeKey       commons.FieldName = "key"
	AttributeValue     commons.FieldName = "value"
	AutoscaleHeadrooms commons.FieldName = "autoscale_headrooms"
)
