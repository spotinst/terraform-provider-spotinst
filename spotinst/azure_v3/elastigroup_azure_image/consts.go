package elastigroup_azure_image

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

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
	Version     commons.FieldName = "version"

	// custom image
	Custom            commons.FieldName = "custom"
	ResourceGroupName commons.FieldName = "resource_group_name"
	ImageName         commons.FieldName = "image_name"

	// gallery image
	GalleryImage             commons.FieldName = "gallery_image"
	GalleryName              commons.FieldName = "gallery_name"
	GalleryImageName         commons.FieldName = "image_name"
	GalleryResourceGroupName commons.FieldName = "resource_group_name"
	GalleryVersion           commons.FieldName = "version"
	SpotAccountId            commons.FieldName = "spot_account_id"
)
