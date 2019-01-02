package ocean_aws_launch_configuration

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	ImageID            commons.FieldName = "image_id"
	IAMInstanceProfile commons.FieldName = "iam_instance_profile"
	KeyName            commons.FieldName = "key_name"
	UserData           commons.FieldName = "user_data"
	SecurityGroups     commons.FieldName = "security_groups"
)
