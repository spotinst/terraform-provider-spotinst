package stateful_node_azure_network

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Network] = commons.NewGenericField(
		commons.StatefulNodeAzureNetwork,
		Network,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
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

					string(NetworkInterface): {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(SubnetName): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(AssignPublicIP): {
									Type:     schema.TypeBool,
									Optional: true,
								},

								string(IsPrimary): {
									Type:     schema.TypeBool,
									Required: true,
								},

								string(PublicIPSku): {
									Type:     schema.TypeString,
									Optional: true,
								},

								string(NetworkSecurityGroup): {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Name): {
												Type:     schema.TypeString,
												Optional: true,
											},
											string(ResourceGroupName): {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},

								string(EnableIPForwarding): {
									Type:     schema.TypeBool,
									Optional: true,
								},

								string(PrivateIPAddresses): {
									Type:     schema.TypeList,
									Elem:     &schema.Schema{Type: schema.TypeString},
									Optional: true,
								},

								string(AdditionalIPConfigurations): {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Name): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(PrivateIPAddressVersion): {
												Type:     schema.TypeString,
												Required: true,
											},
										},
									},
								},
								string(PublicIPs): {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Name): {
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
								string(ApplicationSecurityGroups): {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Name): {
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
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := stWrapper.GetStatefulNode()
			var result []interface{} = nil
			if statefulNode != nil && statefulNode.Compute != nil && statefulNode.Compute.LaunchSpecification != nil && statefulNode.Compute.LaunchSpecification.Network != nil {
				network := statefulNode.Compute.LaunchSpecification.Network
				result = flattenStatefulNodeAzureNetwork(network)
			}

			if result != nil {
				if err := resourceData.Set(string(Network), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Network), err)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := stWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(Network)); ok {
				network := &azure.Network{}
				if value, err := expandStatefulNodeAzureNetwork(v, network); err != nil {
					return err
				} else {
					statefulNode.Compute.LaunchSpecification.SetNetwork(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := stWrapper.GetStatefulNode()
			var value *azure.Network = nil
			if v, ok := resourceData.GetOk(string(Network)); ok {
				network := &azure.Network{}
				if Network, err := expandStatefulNodeAzureNetwork(v, network); err != nil {
					return err
				} else {
					value = Network
				}
				statefulNode.Compute.LaunchSpecification.SetNetwork(value)
			} else {
				statefulNode.Compute.LaunchSpecification.SetNetwork(nil)
			}
			return nil
		},
		nil,
	)
}
func flattenStatefulNodeAzureNetwork(network *azure.Network) []interface{} {
	result := make(map[string]interface{})
	result[string(VirtualNetworkName)] = spotinst.StringValue(network.VirtualNetworkName)
	result[string(ResourceGroupName)] = spotinst.StringValue(network.ResourceGroupName)
	if network.NetworkInterfaces != nil {
		result[string(NetworkInterface)] = flattenStatefulNodeAzureCustomNetworkInterfaces(network.NetworkInterfaces)
	}

	return []interface{}{result}
}

func flattenStatefulNodeAzureCustomNetworkInterfaces(networkInterfaces []*azure.NetworkInterface) []interface{} {
	result := make([]interface{}, 0, len(networkInterfaces))

	for _, networkInterfaces := range networkInterfaces {
		m := make(map[string]interface{})
		m[string(SubnetName)] = spotinst.StringValue(networkInterfaces.SubnetName)
		m[string(IsPrimary)] = spotinst.BoolValue(networkInterfaces.IsPrimary)
		m[string(AssignPublicIP)] = spotinst.BoolValue(networkInterfaces.AssignPublicIP)
		m[string(PublicIPSku)] = spotinst.StringValue(networkInterfaces.PublicIPSku)
		m[string(EnableIPForwarding)] = spotinst.BoolValue(networkInterfaces.EnableIPForwarding)
		if networkInterfaces.PrivateIPAddresses != nil {
			m[string(PrivateIPAddresses)] = spotinst.StringSlice(networkInterfaces.PrivateIPAddresses)
		}
		if networkInterfaces.NetworkSecurityGroup != nil {
			m[string(NetworkSecurityGroup)] = flattenNetworkSecurityGroup(networkInterfaces.NetworkSecurityGroup)
		}
		if networkInterfaces.AdditionalIPConfigurations != nil {
			m[string(AdditionalIPConfigurations)] = flattenAdditionalIPConfigurations(networkInterfaces.AdditionalIPConfigurations)
		}
		if networkInterfaces.PublicIPs != nil {
			m[string(PublicIPs)] = flattenPublicIPs(networkInterfaces.PublicIPs)
		}
		if networkInterfaces.ApplicationSecurityGroups != nil {
			m[string(ApplicationSecurityGroups)] = flattenApplicationSecurityGroups(networkInterfaces.ApplicationSecurityGroups)
		}
		result = append(result, m)
	}

	return result
}

func flattenNetworkSecurityGroup(networkSecurityGroup *azure.NetworkSecurityGroup) []interface{} {
	result := make(map[string]interface{})

	result[string(Name)] = spotinst.StringValue(networkSecurityGroup.Name)
	result[string(ResourceGroupName)] = spotinst.StringValue(networkSecurityGroup.ResourceGroupName)

	return []interface{}{result}
}

func flattenAdditionalIPConfigurations(additionalIPConfigs []*azure.AdditionalIPConfiguration) []interface{} {
	result := make([]interface{}, 0, len(additionalIPConfigs))

	for _, additionalIPConfig := range additionalIPConfigs {
		m := make(map[string]interface{})
		m[string(Name)] = spotinst.StringValue(additionalIPConfig.Name)
		m[string(PrivateIPAddressVersion)] = spotinst.StringValue(additionalIPConfig.PrivateIPAddressVersion)
		result = append(result, m)
	}

	return result
}

func flattenPublicIPs(publicIPS []*azure.PublicIP) []interface{} {
	result := make([]interface{}, 0, len(publicIPS))

	for _, publicIPS := range publicIPS {
		m := make(map[string]interface{})
		m[string(Name)] = spotinst.StringValue(publicIPS.Name)
		m[string(ResourceGroupName)] = spotinst.StringValue(publicIPS.ResourceGroupName)
		result = append(result, m)
	}

	return result
}

func flattenApplicationSecurityGroups(appSecGroups []*azure.ApplicationSecurityGroup) []interface{} {
	result := make([]interface{}, 0, len(appSecGroups))

	for _, appSecGroups := range appSecGroups {
		m := make(map[string]interface{})
		m[string(Name)] = spotinst.StringValue(appSecGroups.Name)
		m[string(ResourceGroupName)] = spotinst.StringValue(appSecGroups.ResourceGroupName)
		result = append(result, m)
	}

	return result
}

func expandStatefulNodeAzureNetwork(data interface{}, network *azure.Network) (*azure.Network, error) {
	list := data.([]interface{})

	if len(list) == 0 && network == nil {
		return nil, nil
	}

	if len(list) > 0 {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(VirtualNetworkName)].(string); ok && v != "" {
			network.SetVirtualNetworkName(spotinst.String(v))
		}

		if v, ok := m[string(ResourceGroupName)].(string); ok && v != "" {
			network.SetResourceGroupName(spotinst.String(v))
		}

		if v, ok := m[string(NetworkInterface)]; ok {
			var networkInterfaces []*azure.NetworkInterface

			if network.NetworkInterfaces != nil {
				networkInterfaces = network.NetworkInterfaces
			}

			networkInterfaces, err := expandNetworkInterfaces(v, networkInterfaces)
			if err != nil {
				return nil, err
			}
			if networkInterfaces != nil {
				network.SetNetworkInterfaces(networkInterfaces)
			} else {
				network.NetworkInterfaces = nil
			}
		}
	}

	return network, nil
}

func expandNetworkInterfaces(data interface{}, networkInterfaces []*azure.NetworkInterface) ([]*azure.NetworkInterface, error) {
	list := data.([]interface{})

	if len(list) > 0 {
		networkInterfaces = make([]*azure.NetworkInterface, 0, len(list))

		for _, v := range list {
			ni, ok := v.(map[string]interface{})
			if !ok {
				continue
			}

			networkInterface := &azure.NetworkInterface{}

			if v, ok := ni[string(SubnetName)].(string); ok && v != "" {
				networkInterface.SetSubnetName(spotinst.String(v))
			}

			if v, ok := ni[string(IsPrimary)].(bool); ok {
				networkInterface.SetIsPrimary(spotinst.Bool(v))
			}

			if v, ok := ni[string(AssignPublicIP)].(bool); ok {
				networkInterface.SetAssignPublicIP(spotinst.Bool(v))
			}

			if v, ok := ni[string(PublicIPSku)].(string); ok {
				networkInterface.SetPublicIPSku(spotinst.String(v))
			}

			if v, ok := ni[string(NetworkSecurityGroup)]; ok {
				// Create new securityGroup object in case cluster did not get it from previous import step.
				networkSecurityGroup := &azure.NetworkSecurityGroup{}

				if networkInterface.NetworkSecurityGroup != nil {
					networkSecurityGroup = networkInterface.NetworkSecurityGroup
				}

				if networkSecurityGroup, err := expandNetworkSecurityGroup(v, networkSecurityGroup); err != nil {
					return nil, err
				} else {
					if networkSecurityGroup != nil {
						networkInterface.SetNetworkSecurityGroup(networkSecurityGroup)
					}
				}
			}

			if v, ok := ni[string(EnableIPForwarding)].(bool); ok {
				networkInterface.SetEnableIPForwarding(spotinst.Bool(v))
			}

			if v, ok := ni[string(PrivateIPAddresses)]; ok {
				if privateIPAddresses, err := expandPrivateIPAddresses(v); err != nil {
					return nil, err
				} else {
					networkInterface.SetPrivateIPAddresses(privateIPAddresses)
				}
			}

			if v, ok := ni[string(AdditionalIPConfigurations)]; ok {
				var additionalIPConfig []*azure.AdditionalIPConfiguration

				if networkInterface.AdditionalIPConfigurations != nil {
					additionalIPConfig = networkInterface.AdditionalIPConfigurations
				}

				if additionalIPConfigs, err := expandAdditionalIPConfig(v, additionalIPConfig); err != nil {
					return nil, err
				} else {
					networkInterface.SetAdditionalIPConfigurations(additionalIPConfigs)
				}
			}

			if v, ok := ni[string(PublicIPs)]; ok {
				var publicIPS []*azure.PublicIP

				if networkInterface.PublicIPs != nil {
					publicIPS = networkInterface.PublicIPs
				}

				if pips, err := expandPublicIPs(v, publicIPS); err != nil {
					return nil, err
				} else {
					networkInterface.SetPublicIPs(pips)
				}
			}

			if v, ok := ni[string(ApplicationSecurityGroups)]; ok {
				var ApplicationSecurityGroups []*azure.ApplicationSecurityGroup

				if networkInterface.AdditionalIPConfigurations != nil {
					ApplicationSecurityGroups = networkInterface.ApplicationSecurityGroups
				}

				if asg, err := expandApplicationSecurityGroups(v, ApplicationSecurityGroups); err != nil {
					return nil, err
				} else {
					networkInterface.SetApplicationSecurityGroups(asg)
				}
			}

			networkInterfaces = append(networkInterfaces, networkInterface)
		}
	}

	return networkInterfaces, nil
}

func expandPrivateIPAddresses(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if privateIPAddresses, ok := v.(string); ok && privateIPAddresses != "" {
			result = append(result, privateIPAddresses)
		}
	}

	return result, nil
}

func expandAdditionalIPConfig(data interface{}, additionalIPConfigs []*azure.AdditionalIPConfiguration) ([]*azure.AdditionalIPConfiguration, error) {
	list := data.([]interface{})

	if len(list) == 0 && additionalIPConfigs == nil {
		return nil, nil
	}

	length := len(list) + len(additionalIPConfigs)
	newAdditionalIPConfigList := make([]*azure.AdditionalIPConfiguration, 0, length)

	if len(additionalIPConfigs) > 0 {
		newAdditionalIPConfigList = append(newAdditionalIPConfigList, additionalIPConfigs[0])
	}

	for _, v := range list {
		adic, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		additionalIPConfig := &azure.AdditionalIPConfiguration{}

		if v, ok := adic[string(Name)].(string); ok && v != "" {
			additionalIPConfig.SetName(spotinst.String(v))
		}

		if v, ok := adic[string(PrivateIPAddressVersion)].(string); ok && v != "" {
			additionalIPConfig.SetPrivateIPAddressVersion(spotinst.String(v))
		}

		newAdditionalIPConfigList = append(newAdditionalIPConfigList, additionalIPConfig)
	}

	return newAdditionalIPConfigList, nil
}

func expandPublicIPs(data interface{}, publicIPS []*azure.PublicIP) ([]*azure.PublicIP, error) {
	list := data.([]interface{})

	if len(list) == 0 && publicIPS == nil {
		return nil, nil
	}

	length := len(list) + len(publicIPS)
	newPublicIPSList := make([]*azure.PublicIP, 0, length)

	if len(publicIPS) > 0 {
		newPublicIPSList = append(newPublicIPSList, publicIPS[0])
	}

	for _, v := range list {
		pips, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		publicIP := &azure.PublicIP{}

		if v, ok := pips[string(Name)].(string); ok && v != "" {
			publicIP.SetName(spotinst.String(v))
		}

		if v, ok := pips[string(ResourceGroupName)].(string); ok && v != "" {
			publicIP.SetResourceGroupName(spotinst.String(v))
		}

		newPublicIPSList = append(newPublicIPSList, publicIP)
	}

	return newPublicIPSList, nil
}

func expandApplicationSecurityGroups(data interface{}, applicationSecGroup []*azure.ApplicationSecurityGroup) ([]*azure.ApplicationSecurityGroup, error) {
	list := data.([]interface{})

	if len(list) == 0 && applicationSecGroup == nil {
		return nil, nil
	}

	length := len(list) + len(applicationSecGroup)
	newapplicationSecGroupList := make([]*azure.ApplicationSecurityGroup, 0, length)

	if len(applicationSecGroup) > 0 {
		newapplicationSecGroupList = append(newapplicationSecGroupList, applicationSecGroup[0])
	}

	for _, v := range list {
		asg, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		appSecGroup := &azure.ApplicationSecurityGroup{}

		if v, ok := asg[string(Name)].(string); ok && v != "" {
			appSecGroup.SetName(spotinst.String(v))
		}

		if v, ok := asg[string(ResourceGroupName)].(string); ok && v != "" {
			appSecGroup.SetResourceGroupName(spotinst.String(v))
		}

		newapplicationSecGroupList = append(newapplicationSecGroupList, appSecGroup)
	}

	return newapplicationSecGroupList, nil
}

func expandNetworkSecurityGroup(data interface{}, networkSecurityGroup *azure.NetworkSecurityGroup) (*azure.NetworkSecurityGroup, error) {
	if list := data.([]interface{}); len(list) > 0 {
		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(Name)].(string); ok && v != "" {
				networkSecurityGroup.SetName(spotinst.String(v))
			}

			if v, ok := m[string(ResourceGroupName)].(string); ok && v != "" {
				networkSecurityGroup.SetResourceGroupName(spotinst.String(v))
			}
		}
		return networkSecurityGroup, nil
	}
	return nil, nil
}
