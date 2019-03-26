package elastigroup_gke

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Name      commons.FieldName = "name"
	NodeImage commons.FieldName = "node_image"
	Location  commons.FieldName = "location"

	MaxSize               commons.FieldName = "max_size"
	MinSize               commons.FieldName = "min_size"
	TargetCapacity        commons.FieldName = "desired_capacity"
	AvailabilityZones     commons.FieldName = "availability_zones"
	PreemptiblePercentage commons.FieldName = "preemptible_percentage"

	// - GKE -----------------------------
	ClusterZoneName commons.FieldName = "cluster_zone_name"
	ClusterID       commons.FieldName = "cluster_id"
	// -----------------------------------
)
