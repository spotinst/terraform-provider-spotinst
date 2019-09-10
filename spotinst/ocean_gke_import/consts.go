package ocean_gke_import

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	ClusterName         commons.FieldName = "cluster_name"
	Location            commons.FieldName = "location"
	Whitelist           commons.FieldName = "whitelist"
	BackendServices     commons.FieldName = "backend_services"
	LocationType        commons.FieldName = "location_type"
	Scheme              commons.FieldName = "scheme"
	NamedPorts          commons.FieldName = "named_ports"
	Ports               commons.FieldName = "ports"
	ServiceName         commons.FieldName = "service_name"
	Name                commons.FieldName = "name"
	MaxSize             commons.FieldName = "max_size"
	MinSize             commons.FieldName = "min_size"
	DesiredCapacity     commons.FieldName = "desired_capacity"
	ClusterControllerID commons.FieldName = "cluster_controller_id"
)
