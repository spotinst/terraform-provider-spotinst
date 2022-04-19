package stateful_node_azure_load_balancer

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	LoadBalancer commons.FieldName = "load_balancers"

	Type              commons.FieldName = "type"
	ResourceGroupName commons.FieldName = "resource_group_name"
	SKU               commons.FieldName = "sku"
	BackendPoolNames  commons.FieldName = "backend_pool_names"
)
