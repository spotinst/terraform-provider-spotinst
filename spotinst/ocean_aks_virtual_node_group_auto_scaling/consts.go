package ocean_aks_virtual_node_group_auto_scaling

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Autoscale     commons.FieldName = "autoscale"
	Headrooms     commons.FieldName = "autoscale_headroom"
	CPUPerUnit    commons.FieldName = "cpu_per_unit"
	GPUPerUnit    commons.FieldName = "gpu_per_unit"
	MemoryPerUnit commons.FieldName = "memory_per_unit"
	NumOfUnits    commons.FieldName = "num_of_units"
)
