package elastigroup_azure_image

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Image] = commons.NewGenericField(
		commons.ElastigroupAzureImage,
		Image,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Marketplace): {
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

								string(Sku): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},

					string(Custom): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(ImageName): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(ResourceGroupName): {
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
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Image)); ok {
				if image, err := expandAzureGroupImage(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetImage(image)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *azure.Image = nil
			if v, ok := resourceData.GetOk(string(Image)); ok {
				if image, err := expandAzureGroupImage(v); err != nil {
					return err
				} else {
					value = image
				}

			}
			elastigroup.Compute.LaunchSpecification.SetImage(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func flattenAzureGroupMarketplaceImage(image *azure.MarketPlaceImage) []interface{} {
	result := make(map[string]interface{})
	result[string(Offer)] = spotinst.StringValue(image.Offer)
	result[string(Publisher)] = spotinst.StringValue(image.Publisher)
	result[string(Sku)] = spotinst.StringValue(image.SKU)
	return []interface{}{result}
}

func flattenAzureGroupCustomImage(image *azure.CustomImage) []interface{} {
	result := make(map[string]interface{})
	result[string(ImageName)] = spotinst.StringValue(image.ImageName)
	result[string(ResourceGroupName)] = spotinst.StringValue(image.ResourceGroupName)
	return []interface{}{result}
}

func expandAzureGroupImage(data interface{}) (*azure.Image, error) {
	image := &azure.Image{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})
		if v, ok := m[string(Marketplace)]; ok {
			marketplace, err := expandAzureGroupMarketplaceImage(v)
			if err != nil {
				return nil, err
			}

			if marketplace != nil {
				image.SetMarketPlaceImage(marketplace)
			}
		} else {
			image.MarketPlace = nil
		}

		if v, ok := m[string(Custom)]; ok {

			custom, err := expandAzureGroupCustomImage(v)
			if err != nil {
				return nil, err
			}
			if custom != nil {
				image.SetCustom(custom)
			}
		} else {
			image.Custom = nil
		}

	} else {
		return nil, errors.New("invalid image configuration")
	}
	return image, nil
}

func expandAzureGroupMarketplaceImage(data interface{}) (*azure.MarketPlaceImage, error) {
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

			if v, ok := m[string(Sku)].(string); ok && v != "" {
				market.SetSKU(spotinst.String(v))
			}

		}
		return market, nil
	}
	return nil, nil
}

func expandAzureGroupCustomImage(data interface{}) (*azure.CustomImage, error) {
	if list := data.([]interface{}); len(list) > 0 {
		custom := &azure.CustomImage{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(ImageName)].(string); ok && v != "" {
				custom.SetImageName(spotinst.String(v))
			}

			if v, ok := m[string(ResourceGroupName)].(string); ok && v != "" {
				custom.SetResourceGroupName(spotinst.String(v))
			}

		}
		return custom, nil
	}
	return nil, nil
}
