package ocean_aks_network

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Network] = commons.NewGenericField(
		commons.OceanAKSNetwork,
		Network,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(VirtualNetworkName): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},

					string(ResourceGroupName): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},

					string(NetworkInterface): {
						Type:     schema.TypeSet,
						Optional: true,
						Computed: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(SubnetName): {
									Type:     schema.TypeString,
									Optional: true,
									Computed: true,
								},

								string(IsPrimary): {
									Type:     schema.TypeBool,
									Optional: true,
									Computed: true,
								},

								string(AssignPublicIP): {
									Type:     schema.TypeBool,
									Optional: true,
									Computed: true,
								},

								string(AdditionalIPConfig): {
									Type:     schema.TypeSet,
									Optional: true,
									Computed: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Name): {
												Type:     schema.TypeString,
												Optional: true,
												Computed: true,
											},

											string(PrivateIPVersion): {
												Type:     schema.TypeString,
												Optional: true,
												Computed: true,
											},
										},
									},
								},

								string(SecurityGroup): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Computed: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Name): {
												Type:     schema.TypeString,
												Optional: true,
												Computed: true,
											},

											string(SecurityGroupResourceGroup): {
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
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value []interface{} = nil

			if cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil &&
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.Network != nil {
				value = flattenNetwork(cluster.VirtualNodeGroupTemplate.LaunchSpecification.Network)
			}
			if err := resourceData.Set(string(Network), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Network), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *azure.Network = nil

			if v, ok := resourceData.GetOk(string(Network)); ok {
				// Create new image object in case cluster did not get it from previous import step.
				network := &azure.Network{}

				if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil {
					if cluster.VirtualNodeGroupTemplate.LaunchSpecification.Network != nil {
						network = cluster.VirtualNodeGroupTemplate.LaunchSpecification.Network
					}

					if net, err := expandNetwork(v, network); err != nil {
						return err
					} else {
						value = net
					}

					cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetNetwork(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *azure.Network = nil

			if v, ok := resourceData.GetOk(string(Network)); ok {
				// Create new image object in case cluster did not get it from previous import step.
				var network *azure.Network

				if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil {
					if cluster.VirtualNodeGroupTemplate.LaunchSpecification.Network != nil {
						network = cluster.VirtualNodeGroupTemplate.LaunchSpecification.Network
					}

					if net, err := expandNetwork(v, network); err != nil {
						return err
					} else {
						value = net
					}

					cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetNetwork(value)
				}
			}
			return nil
		},
		nil,
	)
}

func flattenNetwork(network *azure.Network) []interface{} {
	result := make(map[string]interface{})

	result[string(VirtualNetworkName)] = spotinst.StringValue(network.VirtualNetworkName)
	result[string(ResourceGroupName)] = spotinst.StringValue(network.ResourceGroupName)
	if network.NetworkInterfaces != nil {
		result[string(NetworkInterface)] = flattenNetworkInterfaces(network.NetworkInterfaces)
	}

	return []interface{}{result}
}

func flattenNetworkInterfaces(networkInterfaces []*azure.NetworkInterface) []interface{} {
	result := make([]interface{}, 0, len(networkInterfaces))

	for _, inter := range networkInterfaces {
		m := make(map[string]interface{})
		m[string(SubnetName)] = spotinst.StringValue(inter.SubnetName)
		m[string(IsPrimary)] = spotinst.BoolValue(inter.IsPrimary)
		m[string(AssignPublicIP)] = spotinst.BoolValue(inter.AssignPublicIP)
		if inter.SecurityGroup != nil {
			m[string(SecurityGroup)] = flattenSecurityGroup(inter.SecurityGroup)
		}
		if inter.AdditionalIPConfigs != nil {
			m[string(AdditionalIPConfig)] = flattenAdditionalIPConfigs(inter.AdditionalIPConfigs)
		}
		result = append(result, m)
	}

	return result
}

func flattenAdditionalIPConfigs(additionalIPConfigs []*azure.AdditionalIPConfig) []interface{} {
	result := make([]interface{}, 0, len(additionalIPConfigs))

	for _, additionalIPConfig := range additionalIPConfigs {
		m := make(map[string]interface{})
		m[string(Name)] = spotinst.StringValue(additionalIPConfig.Name)
		m[string(PrivateIPVersion)] = spotinst.StringValue(additionalIPConfig.PrivateIPAddressVersion)
		result = append(result, m)
	}

	return result
}

func flattenSecurityGroup(securityGroup *azure.SecurityGroup) []interface{} {
	result := make(map[string]interface{})

	result[string(SecurityGroupName)] = spotinst.StringValue(securityGroup.Name)
	result[string(SecurityGroupResourceGroup)] = spotinst.StringValue(securityGroup.ResourceGroupName)

	return []interface{}{result}
}

func expandNetwork(data interface{}, network *azure.Network) (*azure.Network, error) {
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
			// Create new NetworkInterfaces slice in case cluster did not get it from previous import step.
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
	list := data.(*schema.Set).List()

	if len(list) > 0 {
		networkInterfaces = make([]*azure.NetworkInterface, 0, len(list))

		for _, v := range list {
			attr, ok := v.(map[string]interface{})
			if !ok {
				continue
			}

			networkInterface := &azure.NetworkInterface{}

			if v, ok := attr[string(SubnetName)].(string); ok && v != "" {
				networkInterface.SetSubnetName(spotinst.String(v))
			}

			if v, ok := attr[string(IsPrimary)].(bool); ok {
				networkInterface.SetIsPrimary(spotinst.Bool(v))
			}

			if v, ok := attr[string(AssignPublicIP)].(bool); ok {
				networkInterface.SetAssignPublicIP(spotinst.Bool(v))
			}

			if v, ok := attr[string(AdditionalIPConfig)]; ok {
				// Create new NetworkInterfaces slice in case cluster did not get it from previous import step.
				var additionalIPConfig []*azure.AdditionalIPConfig

				if networkInterface.AdditionalIPConfigs != nil {
					additionalIPConfig = networkInterface.AdditionalIPConfigs
				}

				if additionalIPConfigs, err := expandAdditionalIPConfig(v, additionalIPConfig); err != nil {
					return nil, err
				} else {
					networkInterface.SetAdditionalIPConfigs(additionalIPConfigs)
				}
			}

			if v, ok := attr[string(SecurityGroup)]; ok {
				// Create new securityGroup object in case cluster did not get it from previous import step.
				securityGroup := &azure.SecurityGroup{}

				if networkInterface.SecurityGroup != nil {
					securityGroup = networkInterface.SecurityGroup
				}

				if securityGroup, err := expandSecurityGroup(v, securityGroup); err != nil {
					return nil, err
				} else {
					if securityGroup != nil {
						networkInterface.SetSecurityGroup(securityGroup)
					}
				}
			}

			networkInterfaces = append(networkInterfaces, networkInterface)
		}
	}

	return networkInterfaces, nil
}

func expandAdditionalIPConfig(data interface{}, additionalIPConfigs []*azure.AdditionalIPConfig) ([]*azure.AdditionalIPConfig, error) {
	list := data.(*schema.Set).List()

	if len(list) == 0 && additionalIPConfigs == nil {
		return nil, nil
	}

	length := len(list) + len(additionalIPConfigs)
	newAdditionalIPConfigList := make([]*azure.AdditionalIPConfig, 0, length)

	if len(additionalIPConfigs) > 0 {
		newAdditionalIPConfigList = append(newAdditionalIPConfigList, additionalIPConfigs[0])
	}

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		additionalIPConfig := &azure.AdditionalIPConfig{}

		if v, ok := attr[string(Name)].(string); ok && v != "" {
			additionalIPConfig.SetName(spotinst.String(v))
		}

		if v, ok := attr[string(PrivateIPVersion)].(string); ok && v != "" {
			additionalIPConfig.SetPrivateIPAddressVersion(spotinst.String(v))
		}

		newAdditionalIPConfigList = append(newAdditionalIPConfigList, additionalIPConfig)
	}

	return newAdditionalIPConfigList, nil
}

func expandSecurityGroup(data interface{}, securityGroup *azure.SecurityGroup) (*azure.SecurityGroup, error) {
	if list := data.([]interface{}); len(list) > 0 {
		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(SecurityGroupName)].(string); ok && v != "" {
				securityGroup.SetName(spotinst.String(v))
			}

			if v, ok := m[string(SecurityGroupResourceGroup)].(string); ok && v != "" {
				securityGroup.SetResourceGroupName(spotinst.String(v))
			}
		}
		return securityGroup, nil
	}
	return nil, nil
}
