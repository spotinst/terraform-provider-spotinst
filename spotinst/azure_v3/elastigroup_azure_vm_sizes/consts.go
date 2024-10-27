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
)
