package mrscaler_aws

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

type TagField string

const (
	TagKey   TagField = "key"
	TagValue TagField = "value"
)

const (
	Name              commons.FieldName = "name"
	Description       commons.FieldName = "description"
	Region            commons.FieldName = "region"
	Strategy          commons.FieldName = "strategy"
	AvailabilityZones commons.FieldName = "availability_zones"
	Tags              commons.FieldName = "tags"
	ClusterID         commons.FieldName = "cluster_id"
	ExposeClusterID   commons.FieldName = "expose_cluster_id"
	OutputClusterID   commons.FieldName = "output_cluster_id"

	ConfigurationsFile   commons.FieldName = "configurations_file"
	BootstrapActionsFile commons.FieldName = "bootstrap_actions_file"
	StepsFile            commons.FieldName = "steps_file"
	Bucket               commons.FieldName = "bucket"
	Key                  commons.FieldName = "key"

	EBSRootVolumeSize           commons.FieldName = "ebs_root_volume_size"
	ManagedPrimarySecurityGroup commons.FieldName = "managed_primary_security_group"
	ManagedReplicaSecurityGroup commons.FieldName = "managed_replica_security_group"
	ServiceAccessSecurityGroup  commons.FieldName = "service_access_security_group"
	AddlPrimarySecurityGroups   commons.FieldName = "additional_primary_security_groups"
	AddlReplicaSecurityGroups   commons.FieldName = "additional_replica_security_groups"
	CustomAMIID                 commons.FieldName = "custom_ami_id"
	RepoUpgradeOnBoot           commons.FieldName = "repo_upgrade_on_boot"
	EC2KeyName                  commons.FieldName = "ec2_key_name"

	Applications commons.FieldName = "applications"
	Args         commons.FieldName = "args"
	AppName      commons.FieldName = "name"
	Version      commons.FieldName = "version"

	InstanceWeights  commons.FieldName = "instance_weights"
	InstanceType     commons.FieldName = "instance_type"
	WeightedCapacity commons.FieldName = "weighted_capacity"
)
