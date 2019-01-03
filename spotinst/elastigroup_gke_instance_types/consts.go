package elastigroup_gke_instance_types

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "instance_types_"
)

const (
	OnDemand    commons.FieldName = Prefix + "ondemand"
	Preemptible commons.FieldName = Prefix + "preemptible"
)
