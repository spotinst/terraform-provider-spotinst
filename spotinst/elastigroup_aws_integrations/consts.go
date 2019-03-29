package elastigroup_aws_integrations

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "integrations_"
)

const (
	// - COMMON --------------------------
	MasterHost commons.FieldName = "master_host"
	MasterPort commons.FieldName = "master_port"

	AutoscaleIsEnabled    commons.FieldName = "autoscale_is_enabled"
	AutoscaleCooldown     commons.FieldName = "autoscale_cooldown"
	AutoscaleHeadroom     commons.FieldName = "autoscale_headroom"
	AutoscaleIsAutoConfig commons.FieldName = "autoscale_is_auto_config"
	CpuPerUnit            commons.FieldName = "cpu_per_unit"
	MemoryPerUnit         commons.FieldName = "memory_per_unit"
	NumOfUnits            commons.FieldName = "num_of_units"

	AutoscaleDown          commons.FieldName = "autoscale_down"
	EvaluationPeriods      commons.FieldName = "evaluation_periods"
	MaxScaleDownPercentage commons.FieldName = "max_scale_down_percentage"

	ApiServer commons.FieldName = "api_server"

	Key   commons.FieldName = "key"
	Value commons.FieldName = "value"
	// -----------------------------------

	// - RANCHER -------------------------
	IntegrationRancher commons.FieldName = "integration_rancher"
	AccessKey          commons.FieldName = "access_key"
	SecretKey          commons.FieldName = "secret_key"
	Version            commons.FieldName = "version"
	// -----------------------------------

	// - ECS -----------------------------
	ShouldScaleDownNonServiceTasks commons.FieldName = "autoscale_scale_down_non_service_tasks"
	IntegrationEcs                 commons.FieldName = "integration_ecs"
	ClusterName                    commons.FieldName = "cluster_name"
	AutoscaleAttributes            commons.FieldName = "autoscale_attributes"
	// -----------------------------------

	// - KUBERNETES ----------------------
	IntegrationKubernetes commons.FieldName = "integration_kubernetes"
	IntegrationMode       commons.FieldName = "integration_mode"
	ClusterIdentifier     commons.FieldName = "cluster_identifier"
	Token                 commons.FieldName = "token"
	AutoscaleLabels       commons.FieldName = "autoscale_labels"
	// -----------------------------------

	// - NOMAD ---------------------------
	IntegrationNomad     commons.FieldName = "integration_nomad"
	AclToken             commons.FieldName = "acl_token"
	AutoscaleConstraints commons.FieldName = "autoscale_constraints"
	// -----------------------------------

	// - MESOSPHERE ----------------------
	IntegrationMesosphere commons.FieldName = "integration_mesosphere"
	// -----------------------------------

	// - MULTAI-RUNTIME ------------------
	IntegrationMultaiRuntime commons.FieldName = "integration_multai_runtime"
	DeploymentId             commons.FieldName = "deployment_id"
	// -----------------------------------

	// - CODE-DEPLOY ---------------------
	IntegrationCodeDeploy      commons.FieldName = "integration_codedeploy"
	CleanupOnFailure           commons.FieldName = "cleanup_on_failure"
	TerminateInstanceOnFailure commons.FieldName = "terminate_instance_on_failure"
	DeploymentGroups           commons.FieldName = "deployment_groups"
	ApplicationName            commons.FieldName = "application_name"
	DeploymentGroupName        commons.FieldName = "deployment_group_name"
	// -----------------------------------

	// - GITLAB ----------------------
	IntegrationGitlab     commons.FieldName = "integration_gitlab"
	GitlabRunner          commons.FieldName = "runner"
	GitlabRunnerIsEnabled commons.FieldName = "is_enabled"
	// -----------------------------------

	// - ROUTE53 -------------------------
	IntegrationRoute53 commons.FieldName = "integration_route53"
	Domains            commons.FieldName = "domains"
	HostedZoneId       commons.FieldName = "hosted_zone_id"
	SpotinstAcctID     commons.FieldName = "spotinst_acct_id"
	RecordSets         commons.FieldName = "record_sets"
	UsePublicIP        commons.FieldName = "use_public_ip"
	Name               commons.FieldName = "name"
	// -----------------------------------

	// - DOCKER-SWARM --------------------
	IntegrationDockerSwarm commons.FieldName = "integration_docker_swarm"
	// -----------------------------------

	// - ELASTIC_BEANSTALK ---------------
	IntegrationBeanstalk  commons.FieldName = "integration_beanstalk"
	EnvironmentId         commons.FieldName = "environment_id"
	DeploymentPreferences commons.FieldName = "deployment_preferences"
	AutomaticRoll         commons.FieldName = "automatic_roll"
	BatchSizePercentage   commons.FieldName = "batch_size_percentage"
	GracePeriod           commons.FieldName = "grace_period"
	Strategy              commons.FieldName = "strategy"
	Action                commons.FieldName = "action"
	ShouldDrainInstances  commons.FieldName = "should_drain_instances"
	ManagedActions        commons.FieldName = "managed_actions"
	PlatformUpdate        commons.FieldName = "platform_update"
	PerformAt             commons.FieldName = "perform_at"
	TimeWindow            commons.FieldName = "time_window"
	UpdateLevel           commons.FieldName = "update_level"
	// -----------------------------------
)
