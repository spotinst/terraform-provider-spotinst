package elastigroup_gcp_launch_configuration

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

type LabelField string
type MetadataField string

const (
	LabelKey   LabelField = "key"
	LabelValue LabelField = "value"

	MetadataKey   MetadataField = "key"
	MetadataValue MetadataField = "value"
)

const (
	Labels   commons.FieldName = "labels"
	Metadata commons.FieldName = "metadata"
	Tags     commons.FieldName = "tags"

	BackendServices commons.FieldName = "backend_services"
	Name            commons.FieldName = "name"
	LocationType    commons.FieldName = "location_type"
	Scheme          commons.FieldName = "scheme"
	NamedPorts      commons.FieldName = "named_ports"

	Ports          commons.FieldName = "ports"
	ServiceName    commons.FieldName = "service_name"
	ServiceAccount commons.FieldName = "service_account"
	StartupScript  commons.FieldName = "startup_script"
	ShutdownScript commons.FieldName = "shutdown_script"
	IPForwarding   commons.FieldName = "ip_forwarding"
)
