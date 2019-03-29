package elastigroup_gcp_integrations

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	// - DOCKER-SWARM --------------------
	IntegrationDockerSwarm commons.FieldName = "integration_docker_swarm"
	MasterHost             commons.FieldName = "master_host"
	MasterPort             commons.FieldName = "master_port"
	// -----------------------------------

	// - GKE -----------------------------
	ClusterID             commons.FieldName = "cluster_id"
	Location              commons.FieldName = "location"
	IntegrationGKE        commons.FieldName = "integration_gke"
	AutoUpdate            commons.FieldName = "auto_update"
	Autoscale             commons.FieldName = "autoscale"
	AutoscaleIsEnabled    commons.FieldName = "autoscale_is_enabled"
	AutoscaleIsAutoConfig commons.FieldName = "autoscale_is_auto_config"
	AutoscaleCooldown     commons.FieldName = "autoscale_cooldown"

	AutoscaleHeadroom commons.FieldName = "autoscale_headroom"
	CpuPerUnit        commons.FieldName = "cpu_per_unit"
	MemoryPerUnit     commons.FieldName = "memory_per_unit"
	NumOfUnits        commons.FieldName = "num_of_units"

	AutoscaleDown     commons.FieldName = "autoscale_down"
	EvaluationPeriods commons.FieldName = "evaluation_periods"

	AutoscaleLabels commons.FieldName = "autoscale_labels"
	Key             commons.FieldName = "key"
	Value           commons.FieldName = "value"
	// -----------------------------------
)
