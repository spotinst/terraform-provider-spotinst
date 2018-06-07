package elastigroup_integrations

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "integrations_"
)

const (
	// - COMMON --------------------------
	MasterHost commons.FieldName = "master_host"

	AutoscaleIsEnabled    commons.FieldName = "autoscale_is_enabled"
	AutoscaleCooldown     commons.FieldName = "autoscale_cooldown"
	AutoscaleHeadroom     commons.FieldName = "autoscale_headroom"
	AutoscaleIsAutoConfig commons.FieldName = "autoscale_is_auto_config"
	CpuPerUnit            commons.FieldName = "cpu_per_unit"
	MemoryPerUnit         commons.FieldName = "memory_per_unit"
	NumOfUnits            commons.FieldName = "num_of_units"

	AutoscaleDown     commons.FieldName = "autoscale_down"
	EvaluationPeriods commons.FieldName = "evaluation_periods"

	ApiServer commons.FieldName = "api_server"
	// -----------------------------------

	// - RANCHER -------------------------
	IntegrationRancher commons.FieldName = "integration_rancher"
	AccessKey          commons.FieldName = "access_key"
	SecretKey          commons.FieldName = "secret_key"
	// -----------------------------------

	// - ECS -----------------------------
	IntegrationEcs commons.FieldName = "integration_ecs"
	ClusterName    commons.FieldName = "cluster_name"
	// -----------------------------------

	// - KUBERNETES ----------------------
	IntegrationKubernetes commons.FieldName = "integration_kubernetes"
	IntegrationMode       commons.FieldName = "integration_mode"
	ClusterIdentifier     commons.FieldName = "cluster_identifier"
	Token                 commons.FieldName = "token"
	// -----------------------------------

	// - NOMAD ---------------------------
	IntegrationNomad     commons.FieldName = "integration_nomad"
	MasterPort           commons.FieldName = "master_port"
	AclToken             commons.FieldName = "acl_token"
	AutoscaleConstraints commons.FieldName = "autoscale_constraints"
	Key                  commons.FieldName = "key"
	Value                commons.FieldName = "value"
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
)
