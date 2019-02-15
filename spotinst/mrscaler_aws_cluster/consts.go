package mrscaler_aws_cluster

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Cluster              commons.FieldName = "cluster"
	LogURI               commons.FieldName = "log_uri"
	AdditionalInfo       commons.FieldName = "additional_info"
	JobFlowRole          commons.FieldName = "job_flow_role"
	SecurityConfig       commons.FieldName = "security_config"
	ServiceRole          commons.FieldName = "service_role"
	VisibleToAllUsers    commons.FieldName = "visible_to_all_users"
	TerminationProtected commons.FieldName = "termination_protected"
	KeepJobFlowAlive     commons.FieldName = "keep_job_flow_alive"
)
