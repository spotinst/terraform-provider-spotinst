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

const (
	LocalDnsProfile             commons.FieldName = "local_dns_profile"
	Mode                        commons.FieldName = "mode"
	VnetDNSOverrides            commons.FieldName = "vnet_dns_overrides"
	KubeDNSOverrides            commons.FieldName = "kube_dns_overrides"
	DNSZone                     commons.FieldName = "zone"
	QueryLogging                commons.FieldName = "query_logging"
	Protocol                    commons.FieldName = "protocol"
	ForwardDestination          commons.FieldName = "forward_destination"
	ForwardPolicy               commons.FieldName = "forward_policy"
	MaxConcurrent               commons.FieldName = "max_concurrent"
	CacheDurationInSeconds      commons.FieldName = "cache_duration_in_seconds"
	ServeStaleDurationInSeconds commons.FieldName = "serve_stale_duration_in_seconds"
	ServeStale                  commons.FieldName = "serve_stale"
)
