package ocean_ecs_launch_specification

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	SecurityGroupIds         commons.FieldName = "security_group_ids"
	IamInstanceProfile       commons.FieldName = "iam_instance_profile"
	KeyPair                  commons.FieldName = "key_pair"
	UserData                 commons.FieldName = "user_data"
	AssociatePublicIpAddress commons.FieldName = "associate_public_ip_address"
	ImageID                  commons.FieldName = "image_id"
	Monitoring               commons.FieldName = "monitoring"
	EBSOptimized             commons.FieldName = "ebs_optimized"
)
