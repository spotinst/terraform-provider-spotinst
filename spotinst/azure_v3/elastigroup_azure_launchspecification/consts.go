package elastigroup_azure_launchspecification

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

type TagField string

const (
	TagKey   TagField = "key"
	TagValue TagField = "value"
)

const (
	CustomData commons.FieldName = "custom_data"
)

const (
	ManagedServiceIdentity                  commons.FieldName = "managed_service_identity"
	ManagedServiceIdentityResourceGroupName commons.FieldName = "resource_group_name"
	ManagedServiceIdentityName              commons.FieldName = "name"
	Tags                                    commons.FieldName = "tags"
)
