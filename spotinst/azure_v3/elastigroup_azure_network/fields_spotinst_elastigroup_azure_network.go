package elastigroup_azure_network

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Network] = commons.NewGenericField(
		commons.ElastigroupAzureNetwork,
		Network,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(VirtualNetworkName): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(ResourceGroupName): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(NetworkInterfaces): {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(SubnetName): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(IsPrimary): {
									Type:     schema.TypeBool,
									Required: true,
								},
								string(AssignPublicIP): {
									Type:     schema.TypeBool,
									Required: true,
								},

								string(AdditionalIPConfigs): {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Name): {
												Type:     schema.TypeString,
												Required: true,
											},

											string(PrivateIPVersion): {
												Type:     schema.TypeString,
												Optional: true,
												StateFunc: func(v interface{}) string {
													value := v.(string)
													return strings.ToUpper(value)
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{} = nil

			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Network != nil {
				value = flattenAzureGroupNetwork(elastigroup.Compute.LaunchSpecification.Network)
			}
			if err := resourceData.Set(string(Network), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Network), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Network)); ok {
				if network, err := expandAzureGroupNetwork(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetNetwork(network)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(Network))
			return err
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func flattenAzureGroupNetwork(network *azurev3.Network) []interface{} {
	result := make(map[string]interface{})

	result[string(VirtualNetworkName)] = spotinst.StringValue(network.VirtualNetworkName)
	result[string(ResourceGroupName)] = spotinst.StringValue(network.ResourceGroupName)
	if network.NetworkInterfaces != nil {
		result[string(NetworkInterfaces)] = flattenAzureGroupNetworkInterfaces(network.NetworkInterfaces)
	}

	return []interface{}{result}
}

func flattenAzureGroupNetworkInterfaces(networkInterfaces []*azurev3.NetworkInterface) []interface{} {
	result := make([]interface{}, 0, len(networkInterfaces))

	for _, inter := range networkInterfaces {
		m := make(map[string]interface{})
		m[string(SubnetName)] = spotinst.StringValue(inter.SubnetName)
		m[string(IsPrimary)] = spotinst.BoolValue(inter.IsPrimary)
		m[string(AssignPublicIP)] = spotinst.BoolValue(inter.AssignPublicIP)
		if inter.AdditionalIPConfigs != nil {
			m[string(AdditionalIPConfigs)] = flattenAzureAdditionalIPConfigs(inter.AdditionalIPConfigs)
		}
		result = append(result, m)
	}

	return result
}

func flattenAzureAdditionalIPConfigs(additionalIPConfigs []*azurev3.AdditionalIPConfig) []interface{} {
	result := make([]interface{}, 0, len(additionalIPConfigs))

	for _, additionalIPConfig := range additionalIPConfigs {
		m := make(map[string]interface{})
		m[string(Name)] = spotinst.StringValue(additionalIPConfig.Name)
		m[string(PrivateIPVersion)] = spotinst.StringValue(additionalIPConfig.PrivateIPAddressVersion)
		result = append(result, m)
	}

	return result
}

func expandAzureGroupNetwork(data interface{}) (*azurev3.Network, error) {
	network := &azurev3.Network{}
	list := data.([]interface{})

	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})
		if v, ok := m[string(VirtualNetworkName)].(string); ok && v != "" {
			network.SetVirtualNetworkName(spotinst.String(v))
		}
		if v, ok := m[string(ResourceGroupName)].(string); ok && v != "" {
			network.SetResourceGroupName(spotinst.String(v))
		}
		if v, ok := m[string(NetworkInterfaces)]; ok {
			if networkInterfaces, err := expandAzureGroupNetworkInterfaces(v); err != nil {
				return nil, err
			} else {
				network.SetNetworkInterfaces(networkInterfaces)
			}
		}
	}

	return network, nil
}

func expandAzureGroupNetworkInterfaces(data interface{}) ([]*azurev3.NetworkInterface, error) {
	list := data.([]interface{})
	networkInterfaces := make([]*azurev3.NetworkInterface, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		networkInterface := &azurev3.NetworkInterface{}
		if !ok {
			continue
		}
		if v, ok := attr[string(SubnetName)].(string); ok && v != "" {
			networkInterface.SetSubnetName(spotinst.String(v))
		}
		if v, ok := attr[string(IsPrimary)].(bool); ok {
			networkInterface.SetIsPrimary(spotinst.Bool(v))
		}
		if v, ok := attr[string(AssignPublicIP)].(bool); ok {
			networkInterface.SetAssignPublicIP(spotinst.Bool(v))
		}
		if v, ok := attr[string(AdditionalIPConfigs)]; ok {
			if additionalIPConfigs, err := expandAzureGroupAddlConfigs(v); err != nil {
				return nil, err
			} else {
				networkInterface.SetAdditionalIPConfigs(additionalIPConfigs)
			}
		}
		networkInterfaces = append(networkInterfaces, networkInterface)
	}

	return networkInterfaces, nil
}

func expandAzureGroupAddlConfigs(data interface{}) ([]*azurev3.AdditionalIPConfig, error) {
	list := data.([]interface{})
	addlConfigs := make([]*azurev3.AdditionalIPConfig, 0, len(list))

	for _, item := range list {
		attr, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		cfg := &azurev3.AdditionalIPConfig{}
		if v, ok := attr[string(Name)].(string); ok && v != "" {
			cfg.SetName(spotinst.String(v))
		}
		if v, ok := attr[string(PrivateIPVersion)].(string); ok && v != "" {
			cfg.SetPrivateIPAddressVersion(spotinst.String(v))
		}
		addlConfigs = append(addlConfigs, cfg)
	}

	return addlConfigs, nil
}
