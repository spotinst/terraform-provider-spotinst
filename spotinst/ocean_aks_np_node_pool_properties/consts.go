package ocean_aks_np_node_pool_properties

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	MaxPodsPerNode     commons.FieldName = "max_pods_per_node"
	EnableNodePublicIP commons.FieldName = "enable_node_public_ip"
	OsDiskSizeGB       commons.FieldName = "os_disk_size_gb"
	OsDiskType         commons.FieldName = "os_disk_type"
	OsType             commons.FieldName = "os_type"
	OsSKU              commons.FieldName = "os_sku"
	KubernetesVersion  commons.FieldName = "kubernetes_version"
	PodSubnetIDs       commons.FieldName = "pod_subnet_ids"
	VnetSubnetIDs      commons.FieldName = "vnet_subnet_ids"
	LinuxOSConfig      commons.FieldName = "linux_os_config"
	Sysctls            commons.FieldName = "sysctls"
	VmMaxMapCount      commons.FieldName = "vm_max_map_count"
)
