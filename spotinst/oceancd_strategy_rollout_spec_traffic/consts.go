package oceancd_strategy_canary_traffic

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Traffic       commons.FieldName = "traffic"
	CanaryService commons.FieldName = "canary_service"
	StableService commons.FieldName = "stable_service"
)

const (
	Alb                 commons.FieldName = "alb"
	AlbAnnotationPrefix commons.FieldName = "alb_annotation_prefix"
	AlbIngress          commons.FieldName = "alb_ingress"
	AlbRootService      commons.FieldName = "alb_root_service"
	ServicePort         commons.FieldName = "service_port"
	StickinessConfig    commons.FieldName = "stickiness_config"
	StickinessDuration  commons.FieldName = "duration_seconds"
	StickinessEnabled   commons.FieldName = "enabled"
)

const (
	Ambassador commons.FieldName = "ambassador"
	Mappings   commons.FieldName = "mappings"
)

const (
	Istio                commons.FieldName = "istio"
	DestinationRule      commons.FieldName = "destination_rule"
	CanarySubsetName     commons.FieldName = "canary_subset_name"
	DestinationRuleName  commons.FieldName = "destination_rule_name"
	StableSubsetName     commons.FieldName = "stable_subset_name"
	VirtualServices      commons.FieldName = "virtual_services"
	VirtualServiceName   commons.FieldName = "virtual_service_name"
	VirtualServiceRoutes commons.FieldName = "virtual_service_routes"
	TlsRoutes            commons.FieldName = "tls_routes"
	Port                 commons.FieldName = "port"
	SniHosts             commons.FieldName = "sni_hosts"
)

const (
	Nginx                       commons.FieldName = "nginx"
	AdditionalIngressAnnotation commons.FieldName = "additional_ingress_annotation"
	CanaryByHeader              commons.FieldName = "canary_by_header"
	Key1                        commons.FieldName = "key1"
	NginxAnnotationPrefix       commons.FieldName = "nginx_annotation_prefix"
	StableIngress               commons.FieldName = "stable_ingress"
)

const (
	PingPong    commons.FieldName = "ping_pong"
	PingService commons.FieldName = "ping_service"
	PongService commons.FieldName = "pong_service"
)

const (
	Smi              commons.FieldName = "smi"
	SmiRootService   commons.FieldName = "smi_root_service"
	TrafficSplitName commons.FieldName = "traffic_split_name"
)
