package elastigroup_azure_load_balancer

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "azure_load_balancer_"
)

const (
	LoadBalancer commons.FieldName = "load_balancers"

	Type        commons.FieldName = "type"
	BalancerID  commons.FieldName = "balancer_id"
	TargetSetID commons.FieldName = "target_set_id"
	AutoWeight  commons.FieldName = "auto_weight"
)
