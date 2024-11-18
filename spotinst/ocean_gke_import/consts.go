package ocean_gke_import

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	ClusterName         commons.FieldName = "cluster_name"
	Location            commons.FieldName = "location"
	Whitelist           commons.FieldName = "whitelist"
	Blacklist           commons.FieldName = "blacklist"
	BackendServices     commons.FieldName = "backend_services"
	LocationType        commons.FieldName = "location_type"
	Scheme              commons.FieldName = "scheme"
	NamedPorts          commons.FieldName = "named_ports"
	Ports               commons.FieldName = "ports"
	ServiceName         commons.FieldName = "service_name"
	Name                commons.FieldName = "name"
	MaxSize             commons.FieldName = "max_size"
	MinSize             commons.FieldName = "min_size"
	DesiredCapacity     commons.FieldName = "desired_capacity"
	ControllerClusterID commons.FieldName = "controller_cluster_id"

	// Deprecated: Please use ControllerClusterID instead.
	ClusterControllerID commons.FieldName = "cluster_controller_id"

	UpdatePolicy    commons.FieldName = "update_policy"
	ShouldRoll      commons.FieldName = "should_roll"
	ConditionedRoll commons.FieldName = "conditioned_roll"

	RollConfig                commons.FieldName = "roll_config"
	BatchSizePercentage       commons.FieldName = "batch_size_percentage"
	LaunchSpecIDs             commons.FieldName = "launch_spec_ids"
	BatchMinHealthyPercentage commons.FieldName = "batch_min_healthy_percentage"
	RespectPdb                commons.FieldName = "respect_pdb"
)

const (
	Filters         commons.FieldName = "filters"
	ExcludeFamilies commons.FieldName = "exclude_families"
	IncludeFamilies commons.FieldName = "include_families"
	MaxMemoryGiB    commons.FieldName = "max_memory_gib"
	MaxVcpu         commons.FieldName = "max_vcpu"
	MinMemoryGiB    commons.FieldName = "min_memory_gib"
	MinVcpu         commons.FieldName = "min_vcpu"
)
