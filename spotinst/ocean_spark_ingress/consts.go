package ocean_spark_ingress

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Ingress            commons.FieldName = "ingress"
	ServiceAnnotations commons.FieldName = "service_annotations"
	Controller         commons.FieldName = "controller"
	LoadBalancer       commons.FieldName = "load_balancer"
	CustomEndpoint     commons.FieldName = "custom_endpoint"
	PrivateLink        commons.FieldName = "private_link"
	Managed            commons.FieldName = "managed"
	Enabled            commons.FieldName = "enabled"
	TargetGroupARN     commons.FieldName = "target_group_arn"
	Address            commons.FieldName = "address"
	VPCEndpointService commons.FieldName = "vpc_endpoint_service"
)
