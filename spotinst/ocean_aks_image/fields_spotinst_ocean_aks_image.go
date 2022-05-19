package ocean_aks_image

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Image] = commons.NewGenericField(
		commons.OceanAKSImage,
		Image,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Marketplace): {
						Type:     schema.TypeList,
						Optional: true,
						Computed: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Publisher): {
									Type:     schema.TypeString,
									Optional: true,
									Computed: true,
								},

								string(Offer): {
									Type:     schema.TypeString,
									Optional: true,
									Computed: true,
								},

								string(SKU): {
									Type:     schema.TypeString,
									Optional: true,
									Computed: true,
								},

								string(Version): {
									Type:     schema.TypeString,
									Optional: true,
									Computed: true,
								},
							},
						},
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification.Image != nil {
				image := cluster.VirtualNodeGroupTemplate.LaunchSpecification.Image
				result = flattenImage(image)
			}

			if result != nil {
				if err := resourceData.Set(string(Image), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Image), err)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *azure.Image = nil

			if v, ok := resourceData.GetOk(string(Image)); ok {
				//create new image object in case cluster did not get it from previous import step.
				image := &azure.Image{}

				if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil {

					if cluster.VirtualNodeGroupTemplate.LaunchSpecification.Image != nil {
						image = cluster.VirtualNodeGroupTemplate.LaunchSpecification.Image
					}

					if image, err := expandImage(v, image); err != nil {
						return err
					} else {
						value = image
					}

					cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetImage(value)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *azure.Image = nil

			if v, ok := resourceData.GetOk(string(Image)); ok {
				//create new image object in case cluster did not get it from previous import step.
				image := &azure.Image{}

				if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil {

					if cluster.VirtualNodeGroupTemplate.LaunchSpecification.Image != nil {
						image = cluster.VirtualNodeGroupTemplate.LaunchSpecification.Image
					}

					if image, err := expandImage(v, image); err != nil {
						return err
					} else {
						value = image
					}

					cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetImage(value)
				}
			}

			return nil
		},
		nil,
	)
}

func flattenImage(image *azure.Image) []interface{} {
	result := make(map[string]interface{})
	if image.MarketplaceImage != nil {
		result[string(Marketplace)] = flattenMarketplaceImage(image.MarketplaceImage)
	}
	return []interface{}{result}
}

func flattenMarketplaceImage(marketplaceImage *azure.MarketplaceImage) []interface{} {
	result := make(map[string]interface{})
	result[string(Offer)] = spotinst.StringValue(marketplaceImage.Offer)
	result[string(Publisher)] = spotinst.StringValue(marketplaceImage.Publisher)
	result[string(SKU)] = spotinst.StringValue(marketplaceImage.SKU)
	result[string(Version)] = spotinst.StringValue(marketplaceImage.Version)
	return []interface{}{result}
}

func expandImage(data interface{}, image *azure.Image) (*azure.Image, error) {
	list := data.([]interface{})

	if len(list) == 0 && image == nil {
		return nil, nil
	}

	if len(list) > 0 {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(Marketplace)]; ok {
			//create new MarketplaceImage object in case cluster did not get it from previous import step.
			marketplaceImage := &azure.MarketplaceImage{}

			if image.MarketplaceImage != nil {
				marketplaceImage = image.MarketplaceImage
			}

			marketplaceImage, err := expandMarketplaceImage(v, marketplaceImage)
			if err != nil {
				return nil, err
			}
			if marketplaceImage != nil {
				image.SetMarketplaceImage(marketplaceImage)
			} else {
				image.MarketplaceImage = nil
			}
		}
	}

	return image, nil
}

func expandMarketplaceImage(data interface{}, marketplaceImage *azure.MarketplaceImage) (*azure.MarketplaceImage, error) {
	if list := data.([]interface{}); len(list) > 0 {
		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(Publisher)].(string); ok && v != "" {
				marketplaceImage.SetPublisher(spotinst.String(v))
			}

			if v, ok := m[string(Offer)].(string); ok && v != "" {
				marketplaceImage.SetOffer(spotinst.String(v))
			}

			if v, ok := m[string(SKU)].(string); ok && v != "" {
				marketplaceImage.SetSKU(spotinst.String(v))
			}

			if v, ok := m[string(Version)].(string); ok && v != "" {
				marketplaceImage.SetVersion(spotinst.String(v))
			}
		}

		return marketplaceImage, nil
	}

	return nil, nil
}
