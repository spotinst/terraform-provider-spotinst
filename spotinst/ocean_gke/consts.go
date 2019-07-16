package ocean_gke

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Name                commons.FieldName = "name"
	ControllerClusterID commons.FieldName = "controller_id"
	MaxSize             commons.FieldName = "max_size"
	MinSize             commons.FieldName = "min_size"
	DesiredCapacity     commons.FieldName = "desired_capacity"
	SubnetName          commons.FieldName = "subnet_name"
	ClusterName         commons.FieldName = "cluster_name"
	MasterLocation      commons.FieldName = "master_location"
	AvailabilityZones   commons.FieldName = "availability_zones"

	BackendServices commons.FieldName = "backend_services"
	LocationType    commons.FieldName = "location_type"
	Scheme          commons.FieldName = "scheme"
	NamedPorts      commons.FieldName = "named_ports"
	Ports           commons.FieldName = "ports"
	ServiceName     commons.FieldName = "service_name"
)

type LabelField string
type MetadataField string

const (
	LabelKey   LabelField = "key"
	LabelValue LabelField = "value"

	MetadataKey   MetadataField = "key"
	MetadataValue MetadataField = "value"

	TaintKey    MetadataField = "key"
	TaintValue  MetadataField = "value"
	TaintEffect MetadataField = "effect"
)

const (
	SourceImage commons.FieldName = "source_image"
	Metadata    commons.FieldName = "metadata"
	Labels      commons.FieldName = "labels"
	Taints      commons.FieldName = "taints"
)
