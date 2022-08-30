package ocean_spark_config

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

type AnnotationField string

const (
	AnnotationKey   AnnotationField = "key"
	AnnotationValue AnnotationField = "value"
)

const (
	IngressServiceAnnotations commons.FieldName = "ingress_service_annotations"

	WebhookUseHostNetwork   commons.FieldName = "webhook_use_host_network"
	WebhookHostNetworkPorts commons.FieldName = "webhook_host_network_ports"

	UseTaints  commons.FieldName = "use_taints"
	CreateVNGs commons.FieldName = "create_vngs"

	CollectDriverLogs commons.FieldName = "collect_driver_logs"
)
