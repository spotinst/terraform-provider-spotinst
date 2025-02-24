package managed_instance_aws_compute_instance_type

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Product        commons.FieldName = "product"
	PreferredType  commons.FieldName = "preferred_type"
	Types          commons.FieldName = "instance_types"
	PreferredTypes commons.FieldName = "preferred_types"

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
