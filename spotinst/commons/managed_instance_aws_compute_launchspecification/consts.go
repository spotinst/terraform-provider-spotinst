package managed_instance_aws_compute_launchspecification

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

type TagField string

const (
	TagKey   TagField = "key"
	TagValue TagField = "value"
)

const (
	EBSOptimized       commons.FieldName = "ebs_optimized"
	EnableMonitoring   commons.FieldName = "enable_monitoring"
	PlacementTenancy   commons.FieldName = "placement_tenancy"
	IamInstanceProfile commons.FieldName = "iam_instance_profile"
	SecurityGroupIds   commons.FieldName = "security_group_ids"
	ImageId            commons.FieldName = "image_id"
	KeyPair            commons.FieldName = "key_pair"
	Tags               commons.FieldName = "tags"
	UserData           commons.FieldName = "user_data"
	ShutdownScript     commons.FieldName = "shutdown_script"
	CPUCredits         commons.FieldName = "cpu_credits"
)
