package elastigroup_aws_launch_configuration

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "launch_configuration_"
)

const (
	ImageId            commons.FieldName = "image_id"
	IamInstanceProfile commons.FieldName = "iam_instance_profile"
	KeyName            commons.FieldName = "key_name"
	SecurityGroups     commons.FieldName = "security_groups"
	UserData           commons.FieldName = "user_data"
	ShutdownScript     commons.FieldName = "shutdown_script"
	EnableMonitoring   commons.FieldName = "enable_monitoring"
	EbsOptimized       commons.FieldName = "ebs_optimized"
	PlacementTenancy   commons.FieldName = "placement_tenancy"
	CPUCredits         commons.FieldName = "cpu_credits"
)
