package managed_instance_aws_compute_launchspecification

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

type TagField string

const (
	TagKey   TagField = "key"
	TagValue TagField = "value"
)

const (
	EBSOptimized        commons.FieldName = "ebs_optimized"
	EnableMonitoring    commons.FieldName = "enable_monitoring"
	PlacementTenancy    commons.FieldName = "placement_tenancy"
	IAMInstanceProfile  commons.FieldName = "iam_instance_profile"
	SecurityGroupIDs    commons.FieldName = "security_group_ids"
	ImageID             commons.FieldName = "image_id"
	KeyPair             commons.FieldName = "key_pair"
	Tags                commons.FieldName = "tags"
	UserData            commons.FieldName = "user_data"
	ShutdownScript      commons.FieldName = "shutdown_script"
	CPUCredits          commons.FieldName = "cpu_credits"
	BlockDeviceMappings commons.FieldName = "block_device_mappings"
)

const (
	DeviceName          commons.FieldName = "device_name"
	EBS                 commons.FieldName = "ebs"
	DeleteOnTermination commons.FieldName = "delete_on_termination"
	IOPS                commons.FieldName = "iops"
	VolumeSize          commons.FieldName = "volume_size"
	VolumeType          commons.FieldName = "volume_type"
	Throughput          commons.FieldName = "throughput"
)
