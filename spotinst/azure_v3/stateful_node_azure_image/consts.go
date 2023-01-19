package stateful_node_azure_image

import (
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

const (
	Image commons.FieldName = "image"
)

// MarketPlaceImage
const (
	MarketPlaceImage commons.FieldName = "marketplace_image"
	Publisher        commons.FieldName = "publisher"
	Offer            commons.FieldName = "offer"
	SKU              commons.FieldName = "sku"
	Version          commons.FieldName = "version"
)

// CustomImage
const (
	CustomImage                  commons.FieldName = "custom_image"
	CustomImageResourceGroupName commons.FieldName = "custom_image_resource_group_name"
	Name                         commons.FieldName = "name"
)

// Gallery
const (
	Gallery                  commons.FieldName = "gallery"
	GalleryResourceGroupName commons.FieldName = "gallery_resource_group_name"
	GalleryName              commons.FieldName = "gallery_name"
	ImageName                commons.FieldName = "image_name"
	VersionName              commons.FieldName = "version_name"
)
