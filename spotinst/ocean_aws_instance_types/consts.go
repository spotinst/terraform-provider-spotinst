package ocean_aws_instance_types

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Whitelist commons.FieldName = "whitelist"
	Blacklist commons.FieldName = "blacklist"
)

const (
	Filters               commons.FieldName = "filters"
	Architectures         commons.FieldName = "architectures"
	Categories            commons.FieldName = "categories"
	DiskTypes             commons.FieldName = "disk_types"
	ExcludeFamilies       commons.FieldName = "exclude_families"
	ExcludeMetal          commons.FieldName = "exclude_metal"
	Hypervisor            commons.FieldName = "hypervisor"
	IncludeFamilies       commons.FieldName = "include_families"
	IsEnaSupported        commons.FieldName = "is_ena_supported"
	MaxGpu                commons.FieldName = "max_gpu"
	MaxMemoryGiB          commons.FieldName = "max_memory_gib"
	MaxNetworkPerformance commons.FieldName = "max_network_performance"
	MaxVcpu               commons.FieldName = "max_vcpu"
	MinEnis               commons.FieldName = "min_enis"
	MinGpu                commons.FieldName = "min_gpu"
	MinMemoryGiB          commons.FieldName = "min_memory_gib"
	MinNetworkPerformance commons.FieldName = "min_network_performance"
	MinVcpu               commons.FieldName = "min_vcpu"
	RootDeviceTypes       commons.FieldName = "root_device_types"
	VirtualizationTypes   commons.FieldName = "virtualization_types"
)
