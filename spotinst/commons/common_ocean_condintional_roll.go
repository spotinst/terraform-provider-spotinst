package commons

var conditionedRollFieldsAWS = []string{"subnet_ids", "whitelist", "blacklist", "user_data", "image_id", "security_groups",
	"key_name", "iam_instance_profile", "associate_public_ip_address", "load_balancers", "instance_metadata_options",
	"ebs_optimized", "root_volume_size"}

var conditionedRollFieldsECS = []string{"subnet_ids", "whitelist", "blacklist", "user_data", "image_id", "security_groups",
	"key_pair", "iam_instance_profile", "associate_public_ip_address", "block_device_mappings", "optimize_images",
	"instance_metadata_options"}

var conditionedRollFieldsGKE = []string{"backend_services", "root_volume_type", "whitelist"}

var conditionedRollFieldsAKS = []string{"availability_zones", "max_pods_per_node", "enable_node_public_ip", "os_disk_size_gb", "os_disk_type", "os_sku", "kubernetes_version", "vnet_subnet_ids", "pod_subnet_ids", "labels", "taints", "tags"}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
