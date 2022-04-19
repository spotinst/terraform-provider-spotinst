package stateful_node_azure_image

//
//import (
//	"errors"
//	"fmt"
//
//	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
//	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
//	"github.com/spotinst/spotinst-sdk-go/spotinst"
//	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
//)
//
//func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
//
//	fieldsMap[Image] = commons.NewGenericField(
//		commons.StatefulNodeAzureImage,
//		Image,
//		&schema.Schema{
//			Type:     schema.TypeList,
//			Optional: true,
//			Elem: &schema.Resource{
//				Schema: map[string]*schema.Schema{
//					string(MarketPlaceImage): {
//						Type:     schema.TypeList,
//						Optional: true,
//						Elem: &schema.Resource{
//							Schema: map[string]*schema.Schema{
//								string(Publisher): {
//									Type:     schema.TypeString,
//									Required: true,
//								},
//
//								string(Offer): {
//									Type:     schema.TypeString,
//									Required: true,
//								},
//
//								string(SKU): {
//									Type:     schema.TypeString,
//									Required: true,
//								},
//
//								string(Version): {
//									Type:     schema.TypeString,
//									Required: true,
//								},
//							},
//						},
//					},
//
//					string(CustomImage): {
//						Type:     schema.TypeList,
//						Optional: true,
//						Elem: &schema.Resource{
//							Schema: map[string]*schema.Schema{
//								string(CustomImageResourceGroupName): {
//									Type:     schema.TypeString,
//									Required: true,
//								},
//
//								string(Name): {
//									Type:     schema.TypeString,
//									Required: true,
//								},
//							},
//						},
//					},
//					string(Gallery): {
//						Type:     schema.TypeList,
//						Optional: true,
//						Elem: &schema.Resource{
//							Schema: map[string]*schema.Schema{
//								string(GalleryResourceGroupName): {
//									Type:     schema.TypeString,
//									Required: true,
//								},
//
//								string(GalleryName): {
//									Type:     schema.TypeString,
//									Required: true,
//								},
//
//								string(ImageName): {
//									Type:     schema.TypeString,
//									Required: true,
//								},
//
//								string(VersionName): {
//									Type:     schema.TypeString,
//									Required: true,
//								},
//							},
//						},
//					},
//				},
//			},
//		},
//
//		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
//			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
//			statefulNode := stWrapper.GetStatefulNode()
//			var result []interface{} = nil
//			if statefulNode != nil && statefulNode.Compute != nil && statefulNode.Compute.LaunchSpecification != nil && statefulNode.Compute.LaunchSpecification.Image != nil {
//				image := statefulNode.Compute.LaunchSpecification.Image
//				result = flattenAzureStatefulGroupImage(image)
//			}
//
//			if result != nil {
//				if err := resourceData.Set(string(Image), result); err != nil {
//					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Image), err)
//				}
//			}
//
//			return nil
//		},
//		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
//			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
//			statefulNode := stWrapper.GetStatefulNode()
//			if v, ok := resourceData.GetOk(string(Image)); ok {
//				if image, err := expandAzureStatefulGroupImage(v); err != nil {
//					return err
//				} else {
//					statefulNode.Compute.LaunchSpecification.SetImage(image)
//				}
//			}
//			return nil
//		},
//		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
//			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
//			statefulNode := stWrapper.GetStatefulNode()
//			var value *azurev3.Image = nil
//			if v, ok := resourceData.GetOk(string(Image)); ok {
//				if image, err := expandAzureStatefulGroupImage(v); err != nil {
//					return err
//				} else {
//					value = image
//				}
//
//			}
//			elastigroup.Compute.LaunchSpecification.SetImage(value)
//			return nil
//		},
//		nil,
//	)
//}
//
//func flattenAzureGroupImage(image *azurev3.Image) []interface{} {
//	result := make(map[string]interface{})
//	if image.Custom != nil {
//		result[string(Custom)] = flattenAzureGroupCustomImage(image.Custom)
//	}
//	if image.MarketPlace != nil {
//		result[string(Marketplace)] = flattenAzureGroupMarketplaceImage(image.MarketPlace)
//	}
//	return []interface{}{result}
//}
//
//func flattenAzureGroupMarketplaceImage(image *azurev3.MarketPlaceImage) []interface{} {
//	result := make(map[string]interface{})
//	result[string(Offer)] = spotinst.StringValue(image.Offer)
//	result[string(Publisher)] = spotinst.StringValue(image.Publisher)
//	result[string(Sku)] = spotinst.StringValue(image.SKU)
//	result[string(Version)] = spotinst.StringValue(image.Version)
//	return []interface{}{result}
//}
//
//func flattenAzureGroupCustomImage(image *azurev3.CustomImage) []interface{} {
//	result := make(map[string]interface{})
//	result[string(ImageName)] = spotinst.StringValue(image.Name)
//	result[string(ResourceGroupName)] = spotinst.StringValue(image.ResourceGroupName)
//	return []interface{}{result}
//}
//
//func expandAzureGroupImage(data interface{}) (*azurev3.Image, error) {
//	image := &azurev3.Image{}
//	list := data.([]interface{})
//	if list != nil && list[0] != nil {
//		m := list[0].(map[string]interface{})
//		if v, ok := m[string(Marketplace)]; ok {
//			marketplace, err := expandAzureGroupMarketplaceImage(v)
//			if err != nil {
//				return nil, err
//			}
//
//			if marketplace != nil {
//				image.SetMarketPlaceImage(marketplace)
//			}
//		} else {
//			image.MarketPlace = nil
//		}
//
//		if v, ok := m[string(Custom)]; ok {
//
//			custom, err := expandAzureGroupCustomImage(v)
//			if err != nil {
//				return nil, err
//			}
//			if custom != nil {
//				image.SetCustom(custom)
//			}
//		} else {
//			image.Custom = nil
//		}
//
//	} else {
//		return nil, errors.New("invalid image configuration")
//	}
//	return image, nil
//}
//
//func expandAzureGroupMarketplaceImage(data interface{}) (*azurev3.MarketPlaceImage, error) {
//	market := &azurev3.MarketPlaceImage{}
//	if list := data.([]interface{}); len(list) > 0 {
//		if list != nil && list[0] != nil {
//			m := list[0].(map[string]interface{})
//
//			if v, ok := m[string(Publisher)].(string); ok && v != "" {
//				market.SetPublisher(spotinst.String(v))
//			}
//
//			if v, ok := m[string(Offer)].(string); ok && v != "" {
//				market.SetOffer(spotinst.String(v))
//			}
//
//			if v, ok := m[string(Sku)].(string); ok && v != "" {
//				market.SetSKU(spotinst.String(v))
//			}
//
//			if v, ok := m[string(Version)].(string); ok && v != "" {
//				market.SetVersion(spotinst.String(v))
//			}
//		}
//		return market, nil
//	}
//	return nil, nil
//}
//
//func expandAzureGroupCustomImage(data interface{}) (*azurev3.CustomImage, error) {
//	if list := data.([]interface{}); len(list) > 0 {
//		custom := &azurev3.CustomImage{}
//		if list != nil && list[0] != nil {
//			m := list[0].(map[string]interface{})
//			if v, ok := m[string(ImageName)].(string); ok && v != "" {
//				custom.SetName(spotinst.String(v))
//			}
//			if v, ok := m[string(ResourceGroupName)].(string); ok && v != "" {
//				custom.SetResourceGroupName(spotinst.String(v))
//			}
//		}
//		return custom, nil
//	}
//	return nil, nil
//}
