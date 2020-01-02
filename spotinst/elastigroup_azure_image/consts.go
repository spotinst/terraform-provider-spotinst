package elastigroup_azure_image

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "azure_image_"
)

const (
	Image commons.FieldName = "image"

	// marketplace image
	Marketplace commons.FieldName = "marketplace"
	Publisher   commons.FieldName = "publisher"
	Offer       commons.FieldName = "offer"
	Sku         commons.FieldName = "sku"

	// custom image
	Custom            commons.FieldName = "custom"
	ResourceGroupName commons.FieldName = "resource_group_name"
	ImageName         commons.FieldName = "image_name"
)
