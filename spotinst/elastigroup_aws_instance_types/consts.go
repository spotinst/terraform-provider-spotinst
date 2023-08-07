package elastigroup_aws_instance_types

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "instance_types_"
)

const (
	OnDemand      commons.FieldName = Prefix + "ondemand"
	Spot          commons.FieldName = Prefix + "spot"
	PreferredSpot commons.FieldName = Prefix + "preferred_spot"
	OnDemandTypes commons.FieldName = "on_demand_types"

	InstanceTypeWeights commons.FieldName = Prefix + "weights"
	InstanceType        commons.FieldName = "instance_type"
	Weight              commons.FieldName = "weight"

	ResourceRequirements        commons.FieldName = "resource_requirements"
	ExcludedInstanceFamilies    commons.FieldName = "excluded_instance_families"
	ExcludedInstanceGenerations commons.FieldName = "excluded_instance_generations"
	ExcludedInstanceTypes       commons.FieldName = "excluded_instance_types"
	RequiredGpuMinimum          commons.FieldName = "required_gpu_minimum"
	RequiredGpuMaximum          commons.FieldName = "required_gpu_maximum"
	RequiredMemoryMinimum       commons.FieldName = "required_memory_minimum"
	RequiredMemoryMaximum       commons.FieldName = "required_memory_maximum"
	RequiredVCpuMinimum         commons.FieldName = "required_vcpu_minimum"
	RequiredVCpuMaximum         commons.FieldName = "required_vcpu_maximum"
)
