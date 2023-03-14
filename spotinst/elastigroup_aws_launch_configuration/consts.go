package elastigroup_aws_launch_configuration

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "launch_configuration_"
)

const (
	ImageId            commons.FieldName = "image_id"
	Images             commons.FieldName = "images"
	Image              commons.FieldName = "image"
	Id                 commons.FieldName = "id"
	IamInstanceProfile commons.FieldName = "iam_instance_profile"
	KeyName            commons.FieldName = "key_name"
	SecurityGroups     commons.FieldName = "security_groups"
	UserData           commons.FieldName = "user_data"
	ShutdownScript     commons.FieldName = "shutdown_script"
	EnableMonitoring   commons.FieldName = "enable_monitoring"
	EbsOptimized       commons.FieldName = "ebs_optimized"
	PlacementTenancy   commons.FieldName = "placement_tenancy"
	CPUCredits         commons.FieldName = "cpu_credits"
	MetadataOptions    commons.FieldName = "metadata_options"
	CPUOptions         commons.FieldName = "cpu_options"

	// - MetadataOptions -----------------------------
	HTTPTokens              commons.FieldName = "http_tokens"
	HTTPPutResponseHopLimit commons.FieldName = "http_put_response_hop_limit"
	InstanceMetadataTags    commons.FieldName = "instance_metadata_tags"
	// -----------------------------------

	// - CPUOptions -----------------------------
	ThreadsPerCore commons.FieldName = "threads_per_core"
	// -----------------------------------
)

const (
	ResourceTagSpecification commons.FieldName = "resource_tag_specification"
	ShouldTagVolumes         commons.FieldName = "should_tag_volumes"
	ShouldTagSnapshots       commons.FieldName = "should_tag_snapshots"
	ShouldTagENIs            commons.FieldName = "should_tag_enis"
	ShouldTagAMIs            commons.FieldName = "should_tag_amis"
)

const (
	ITF                           commons.FieldName = "itf"
	LoadBalancer                  commons.FieldName = "load_balancer"
	MigrationHealthinessThreshold commons.FieldName = "migration_healthiness_threshold"
	FixedTargetGroups             commons.FieldName = "fixed_target_groups"
	WeightStrategy                commons.FieldName = "weight_strategy"
	DefaultStaticTargetGroup      commons.FieldName = "default_static_target_group"
	TargetGroupConfig             commons.FieldName = "target_group_config"
	ListenerRule                  commons.FieldName = "listener_rule"
	LoadBalancerARN               commons.FieldName = "load_balancer_arn"
	StaticTargetGroup             commons.FieldName = "static_target_group"
	ARN                           commons.FieldName = "arn"
	Percentage                    commons.FieldName = "percentage"
	RuleARN                       commons.FieldName = "rule_arn"
	HealthCheckIntervalSeconds    commons.FieldName = "health_check_interval_seconds"
	HealthCheckPath               commons.FieldName = "health_check_path"
	HealthCheckPort               commons.FieldName = "health_check_port"
	HealthCheckProtocol           commons.FieldName = "health_check_protocol"
	HealthCheckTimeoutSeconds     commons.FieldName = "health_check_timeout_seconds"
	HealthyThresholdCount         commons.FieldName = "healthy_threshold_count"
	UnhealthyThresholdCount       commons.FieldName = "unhealthy_threshold_count"
	Port                          commons.FieldName = "port"
	Protocol                      commons.FieldName = "protocol"
	ProtocolVersion               commons.FieldName = "protocol_version"
	Matcher                       commons.FieldName = "matcher"
	HTTPCode                      commons.FieldName = "http_code"
	GRPCCode                      commons.FieldName = "grpc_code"
	Tags                          commons.FieldName = "tags"
	TagKey                        commons.FieldName = "tag_key"
	TagValue                      commons.FieldName = "tag_value"
	VPCID                         commons.FieldName = "vpc_id"
)
