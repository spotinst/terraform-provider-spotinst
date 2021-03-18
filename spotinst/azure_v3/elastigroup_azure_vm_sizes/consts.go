package elastigroup_azure_vm_sizes

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "azure_vm_sizes_"
)

const (
	OnDemandSizes commons.FieldName = "od_sizes"
	SpotSizes     commons.FieldName = "spot_sizes"
)
