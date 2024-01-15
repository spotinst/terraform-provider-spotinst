package ocean_aks_np

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Name                               commons.FieldName = "name"
	ControllerClusterID                commons.FieldName = "controller_cluster_id"
	AvailabilityZones                  commons.FieldName = "availability_zones"
	AKSClusterName                     commons.FieldName = "aks_cluster_name"
	AKSResourceGroupName               commons.FieldName = "aks_resource_group_name"
	AKSRegion                          commons.FieldName = "aks_region"
	AKSInfrastructureResourceGroupName commons.FieldName = "aks_infrastructure_resource_group_name"

	UpdatePolicy          commons.FieldName = "update_policy"
	ShouldRoll            commons.FieldName = "should_roll"
	ConditionedRoll       commons.FieldName = "conditioned_roll"
	ConditionedRollParams commons.FieldName = "conditioned_roll_params"

	RollConfig                commons.FieldName = "roll_config"
	BatchSizePercentage       commons.FieldName = "batch_size_percentage"
	VngIDs                    commons.FieldName = "vng_ids"
	BatchMinHealthyPercentage commons.FieldName = "batch_min_healthy_percentage"
	RespectPDB                commons.FieldName = "respect_pdb"
	Comment                   commons.FieldName = "comment"
	NodePoolNames             commons.FieldName = "node_pool_names"
	RespectRestrictScaleDown  commons.FieldName = "respect_restrict_scale_down"
	NodeNames                 commons.FieldName = "node_names"
)
