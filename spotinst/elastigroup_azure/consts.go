package elastigroup_azure

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Name              commons.FieldName = "name"
	Region            commons.FieldName = "region"
	Product           commons.FieldName = "product"
	ResourceGroupName commons.FieldName = "resource_group_name"

	MaxSize         commons.FieldName = "max_size"
	MinSize         commons.FieldName = "min_size"
	DesiredCapacity commons.FieldName = "desired_capacity"
)
