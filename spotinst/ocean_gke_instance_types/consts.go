package ocean_gke_instance_types

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Whitelist      commons.FieldName = "whitelist"
	Blacklist      commons.FieldName = "blacklist"
	PreferredTypes commons.FieldName = "preferred_types"
)

const (
	Filters         commons.FieldName = "filters"
	ExcludeFamilies commons.FieldName = "exclude_families"
	IncludeFamilies commons.FieldName = "include_families"
	MaxMemoryGiB    commons.FieldName = "max_memory_gib"
	MaxVcpu         commons.FieldName = "max_vcpu"
	MinMemoryGiB    commons.FieldName = "min_memory_gib"
	MinVcpu         commons.FieldName = "min_vcpu"
)
