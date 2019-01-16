package elastigroup_azure_vm_sizes

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "azure_vm_sizes_"
)

const (
	VMSizes     commons.FieldName = "elastigroup_azure_vm_sizes"
	OnDemand    commons.FieldName = "od_sizes"
	LowPriority commons.FieldName = "low_priority_sizes"
)
