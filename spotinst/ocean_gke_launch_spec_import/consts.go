package ocean_gke_launch_spec_import

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

type LabelField string
type MetadataField string

const (
	OceanId      commons.FieldName = "ocean_id"
	NodePoolName commons.FieldName = "node_pool_name"
)
