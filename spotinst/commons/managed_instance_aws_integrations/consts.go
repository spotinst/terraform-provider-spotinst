package managed_instance_aws_integrations

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

type BalancerType string

const (
	// - ROUTE53 -------------------------
	IntegrationRoute53 commons.FieldName = "integration_route53"
	Domains            commons.FieldName = "domains"
	HostedZoneId       commons.FieldName = "hosted_zone_id"
	SpotinstAcctID     commons.FieldName = "spotinst_acct_id"
	RecordSets         commons.FieldName = "record_sets"
	UsePublicIP        commons.FieldName = "use_public_ip"
	Route53Name        commons.FieldName = "name"
	// -----------------------------------

	LoadBalancers    commons.FieldName = "load_balancers"
	Arn              commons.FieldName = "arn"
	AzAwareness      commons.FieldName = "az_awareness"
	LoadBalancerName commons.FieldName = "name"
	Type             commons.FieldName = "type"
	BalancerID       commons.FieldName = "balancer_id"
	TargetSetID      commons.FieldName = "target_set_id"
	AutoWeight       commons.FieldName = "auto_weight"
)
