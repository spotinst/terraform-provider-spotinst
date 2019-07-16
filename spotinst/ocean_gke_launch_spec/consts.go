package ocean_gke_launch_spec

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

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
	OceanId     commons.FieldName = "ocean_id"
	SourceImage commons.FieldName = "source_image"
	Metadata    commons.FieldName = "metadata"
	Labels      commons.FieldName = "labels"
	Taints      commons.FieldName = "taints"
)
