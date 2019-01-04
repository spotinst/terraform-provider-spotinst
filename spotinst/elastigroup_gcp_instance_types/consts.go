package elastigroup_gcp_instance_types

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "instance_types_"
)

const (
	OnDemand    commons.FieldName = Prefix + "ondemand"
	Preemptible commons.FieldName = Prefix + "preemptible"
	Custom      commons.FieldName = Prefix + "custom"

	VCPU      commons.FieldName = "vcpu"
	MemoryGiB commons.FieldName = "memory_gib"
)
