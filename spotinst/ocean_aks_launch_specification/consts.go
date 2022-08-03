package ocean_aks_launch_specification

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	CustomData        commons.FieldName = "custom_data"
	ResourceGroupName commons.FieldName = "resource_group_name"
	MaxPods           commons.FieldName = "max_pods"
)

const (
	ManagedServiceIdentity                  commons.FieldName = "managed_service_identity"
	ManagedServiceIdentityResourceGroupName commons.FieldName = "resource_group_name"
	ManagedServiceIdentityName              commons.FieldName = "name"
)

const (
	Tag      commons.FieldName = "tag"
	TagKey   commons.FieldName = "key"
	TagValue commons.FieldName = "value"
)
