package elastigroup_azure_vm_sizes

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "azure_vm_sizes_"
)

const (
	VmSizes            commons.FieldName = "vm_sizes"
	OnDemandSizes      commons.FieldName = "od_sizes"
	SpotSizes          commons.FieldName = "spot_sizes"
	PreferredSpotSizes commons.FieldName = "preferred_spot_sizes"
	ExcludedVmSizes    commons.FieldName = "excluded_vm_sizes"

	SpotSizeAttributes commons.FieldName = "spot_size_attributes"
	MaxCpu             commons.FieldName = "max_cpu"
	MaxMemory          commons.FieldName = "max_memory"
	MaxStorage         commons.FieldName = "max_storage"
	MinCpu             commons.FieldName = "min_cpu"
	MinMemory          commons.FieldName = "min_memory"
	MinStorage         commons.FieldName = "min_storage"
)
