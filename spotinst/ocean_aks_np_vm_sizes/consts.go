package ocean_aks_np_vm_sizes

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Filters               commons.FieldName = "filters"
	MinVcpu               commons.FieldName = "min_vcpu"
	MaxVcpu               commons.FieldName = "max_vcpu"
	MinMemoryGiB          commons.FieldName = "min_memory_gib"
	MaxMemoryGiB          commons.FieldName = "max_memory_gib"
	Series                commons.FieldName = "series"
	Architectures         commons.FieldName = "architectures"
	ExcludeSeries         commons.FieldName = "exclude_series"
	AcceleratedNetworking commons.FieldName = "accelerated_networking"
	DiskPerformance       commons.FieldName = "disk_performance"
	MinGpu                commons.FieldName = "min_gpu"
	MaxGpu                commons.FieldName = "max_gpu"
	MinNICs               commons.FieldName = "min_nics"
	MinDisk               commons.FieldName = "min_disk"
	VmTypes               commons.FieldName = "vm_types"
)
