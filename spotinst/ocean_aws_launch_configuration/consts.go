package ocean_aws_launch_configuration

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	ImageID                                       commons.FieldName = "image_id"
	IAMInstanceProfile                            commons.FieldName = "iam_instance_profile"
	KeyName                                       commons.FieldName = "key_name"
	UserData                                      commons.FieldName = "user_data"
	SecurityGroups                                commons.FieldName = "security_groups"
	AssociatePublicIpAddress                      commons.FieldName = "associate_public_ip_address"
	AssociateIPv6Address                          commons.FieldName = "associate_ipv6_address"
	LoadBalancers                                 commons.FieldName = "load_balancers"
	Arn                                           commons.FieldName = "arn"
	Name                                          commons.FieldName = "name"
	Type                                          commons.FieldName = "type"
	RootVolumeSize                                commons.FieldName = "root_volume_size"
	Monitoring                                    commons.FieldName = "monitoring"
	EBSOptimized                                  commons.FieldName = "ebs_optimized"
	UseAsTemplateOnly                             commons.FieldName = "use_as_template_only"
	ResourceTagSpecification                      commons.FieldName = "resource_tag_specification"
	ShouldTagVolumes                              commons.FieldName = "should_tag_volumes"
	HealthCheckUnhealthyDurationBeforeReplacement commons.FieldName = "health_check_unhealthy_duration_before_replacement"
	ReservedENIs                                  commons.FieldName = "reserved_enis"
	PrimaryIPv6                                   commons.FieldName = "primary_ipv6"
)

const (
	InstanceMetadataOptions commons.FieldName = "instance_metadata_options"
	HTTPTokens              commons.FieldName = "http_tokens"
	HTTPPutResponseHopLimit commons.FieldName = "http_put_response_hop_limit"
)
const (
	BlockDeviceMappings commons.FieldName = "block_device_mappings"
	DeviceName          commons.FieldName = "device_name"
	Ebs                 commons.FieldName = "ebs"
	DeleteOnTermination commons.FieldName = "delete_on_termination"
	Encrypted           commons.FieldName = "encrypted"
	IOPS                commons.FieldName = "iops"
	KMSKeyID            commons.FieldName = "kms_key_id"
	SnapshotID          commons.FieldName = "snapshot_id"
	VolumeSize          commons.FieldName = "volume_size"
	DynamicVolumeSize   commons.FieldName = "dynamic_volume_size"
	BaseSize            commons.FieldName = "base_size"
	Resource            commons.FieldName = "resource"
	SizePerResourceUnit commons.FieldName = "size_per_resource_unit"
	VolumeType          commons.FieldName = "volume_type"
	Throughput          commons.FieldName = "throughput"
)
const (
	DynamicIops             commons.FieldName = "dynamic_iops"
	IopsBaseSize            commons.FieldName = "base_size"
	IopsResource            commons.FieldName = "resource"
	IopsSizePerResourceUnit commons.FieldName = "size_per_resource_unit"
)
const (
	InstanceStorePolicy     commons.FieldName = "instance_store_policy"
	InstanceStorePolicyType commons.FieldName = "instance_store_policy_type"
)

const (
	StartupTaints       commons.FieldName = "startup_taints"
	StartupTaintsKey    commons.FieldName = "key"
	StartupTaintsValue  commons.FieldName = "value"
	StartupTaintsEffect commons.FieldName = "effect"
)
