package stateful_node_azure_image

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Image] = commons.NewGenericField(
		commons.StatefulNodeAzureImage,
		Image,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MarketPlaceImage): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Publisher): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(Offer): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(SKU): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(Version): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},

					string(CustomImage): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(CustomImageResourceGroupName): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(Name): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},
					string(Gallery): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(GalleryResourceGroupName): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(GalleryName): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(ImageName): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(VersionName): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := stWrapper.GetStatefulNode()
			var result []interface{} = nil
			if statefulNode != nil && statefulNode.Compute != nil && statefulNode.Compute.LaunchSpecification != nil && statefulNode.Compute.LaunchSpecification.Image != nil {
				image := statefulNode.Compute.LaunchSpecification.Image
				result = flattenStatefulNodeAzureImage(image)
			}

			if result != nil {
				if err := resourceData.Set(string(Image), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Image), err)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := stWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(Image)); ok {
				if image, err := expandStatefulNodeAzureImage(v); err != nil {
					return err
				} else {
					statefulNode.Compute.LaunchSpecification.SetImage(image)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := stWrapper.GetStatefulNode()
			var value *azure.Image = nil
			if v, ok := resourceData.GetOk(string(Image)); ok {
				if image, err := expandStatefulNodeAzureImage(v); err != nil {
					return err
				} else {
					value = image
				}

			}
			statefulNode.Compute.LaunchSpecification.SetImage(value)
			return nil
		},
		nil,
	)
}
func flattenStatefulNodeAzureImage(image *azure.Image) []interface{} {
	result := make(map[string]interface{})
	if image.Custom != nil {
		result[string(CustomImage)] = flattenStatefulNodeAzureCustomImage(image.Custom)
	}
	if image.MarketPlace != nil {
		result[string(MarketPlaceImage)] = flattenStatefulNodeAzureMarketplaceImage(image.MarketPlace)
	}
	if image.Gallery != nil {
		result[string(Gallery)] = flattenStatefulNodeAzureGallery(image.Gallery)
	}
	return []interface{}{result}
}

func flattenStatefulNodeAzureMarketplaceImage(image *azure.MarketPlaceImage) []interface{} {
	result := make(map[string]interface{})
	result[string(Offer)] = spotinst.StringValue(image.Offer)
	result[string(Publisher)] = spotinst.StringValue(image.Publisher)
	result[string(SKU)] = spotinst.StringValue(image.SKU)
	result[string(Version)] = spotinst.StringValue(image.Version)
	return []interface{}{result}
}

func flattenStatefulNodeAzureCustomImage(image *azure.CustomImage) []interface{} {
	result := make(map[string]interface{})
	result[string(CustomImageResourceGroupName)] = spotinst.StringValue(image.ResourceGroupName)
	result[string(Name)] = spotinst.StringValue(image.Name)
	return []interface{}{result}
}

func flattenStatefulNodeAzureGallery(image *azure.Gallery) []interface{} {
	result := make(map[string]interface{})
	result[string(GalleryResourceGroupName)] = spotinst.StringValue(image.ResourceGroupName)
	result[string(GalleryName)] = spotinst.StringValue(image.GalleryName)
	result[string(ImageName)] = spotinst.StringValue(image.ImageName)
	result[string(VersionName)] = spotinst.StringValue(image.VersionName)
	return []interface{}{result}
}

func expandStatefulNodeAzureImage(data interface{}) (*azure.Image, error) {
	image := &azure.Image{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})
		if v, ok := m[string(MarketPlaceImage)]; ok {
			marketplace, err := expandStatefulNodeAzureMarketplaceImage(v)
			if err != nil {
				return nil, err
			}

			if marketplace != nil {
				image.SetMarketPlaceImage(marketplace)
			}
		} else {
			image.MarketPlace = nil
		}

		if v, ok := m[string(CustomImage)]; ok {
			custom, err := expandStatefulNodeAzureCustomImage(v)
			if err != nil {
				return nil, err
			}
			if custom != nil {
				image.SetCustom(custom)
			}
		} else {
			image.Custom = nil
		}

		if v, ok := m[string(Gallery)]; ok {
			gallery, err := expandStatefulNodeAzureGallery(v)
			if err != nil {
				return nil, err
			}

			if gallery != nil {
				image.SetGallery(gallery)
			}
		} else {
			image.Gallery = nil
		}

	} else {
		return nil, errors.New("invalid image configuration")
	}
	return image, nil
}

func expandStatefulNodeAzureMarketplaceImage(data interface{}) (*azure.MarketPlaceImage, error) {
	market := &azure.MarketPlaceImage{}
	if list := data.([]interface{}); len(list) > 0 {
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(Publisher)].(string); ok && v != "" {
				market.SetPublisher(spotinst.String(v))
			}

			if v, ok := m[string(Offer)].(string); ok && v != "" {
				market.SetOffer(spotinst.String(v))
			}

			if v, ok := m[string(SKU)].(string); ok && v != "" {
				market.SetSKU(spotinst.String(v))
			}

			if v, ok := m[string(Version)].(string); ok && v != "" {
				market.SetVersion(spotinst.String(v))
			}
		}
		return market, nil
	}
	return nil, nil
}

func expandStatefulNodeAzureCustomImage(data interface{}) (*azure.CustomImage, error) {
	if list := data.([]interface{}); len(list) > 0 {
		custom := &azure.CustomImage{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})
			if v, ok := m[string(CustomImageResourceGroupName)].(string); ok && v != "" {
				custom.SetResourceGroupName(spotinst.String(v))
			}
			if v, ok := m[string(Name)].(string); ok && v != "" {
				custom.SetName(spotinst.String(v))
			}
		}
		return custom, nil
	}
	return nil, nil
}

func expandStatefulNodeAzureGallery(data interface{}) (*azure.Gallery, error) {
	if list := data.([]interface{}); len(list) > 0 {
		gallery := &azure.Gallery{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})
			if v, ok := m[string(GalleryResourceGroupName)].(string); ok && v != "" {
				gallery.SetResourceGroupName(spotinst.String(v))
			}
			if v, ok := m[string(GalleryName)].(string); ok && v != "" {
				gallery.SetGalleryName(spotinst.String(v))
			}
			if v, ok := m[string(ImageName)].(string); ok && v != "" {
				gallery.SetImageName(spotinst.String(v))
			}
			if v, ok := m[string(VersionName)].(string); ok && v != "" {
				gallery.SetVersionName(spotinst.String(v))
			}
		}
		return gallery, nil
	}
	return nil, nil
}
