package elastigroup_azure_network

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

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
								string(PublicIPSku): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(SecurityGroup): {
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
								string(ApplicationSecurityGroup): {
									Type:     schema.TypeSet,
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
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *azurev3.Network = nil
			if v, ok := resourceData.GetOk(string(Network)); ok {
				if network, err := expandAzureGroupNetwork(v); err != nil {
					return err
				} else {
					value = network
				}
			}
			elastigroup.Compute.LaunchSpecification.SetNetwork(value)
			return nil
		},
		nil,
	)
}

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
		m[string(PublicIPSku)] = spotinst.StringValue(inter.PublicIpSku)
		m[string(EnableIPForwarding)] = spotinst.BoolValue(inter.EnableIPForwarding)
		if inter.PrivateIpAddresses != nil {
			m[string(PrivateIPAddresses)] = spotinst.StringSlice(inter.PrivateIpAddresses)
		}
		if inter.SecurityGroup != nil {
			m[string(SecurityGroup)] = flattenSecurityGroup(inter.SecurityGroup)
		}
		if inter.AdditionalIPConfigs != nil {
			m[string(AdditionalIPConfigs)] = flattenAzureAdditionalIPConfigs(inter.AdditionalIPConfigs)
		}
		if inter.ApplicationSecurityGroups != nil {
			m[string(ApplicationSecurityGroup)] = flattenApplicationSecurityGroups(inter.ApplicationSecurityGroups)
		}
		if inter.PublicIps != nil {
			m[string(PublicIPs)] = flattenPublicIPs(inter.PublicIps)
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

func flattenApplicationSecurityGroups(applicationSecurityGroups []*azurev3.ApplicationSecurityGroup) []interface{} {
	result := make([]interface{}, 0, len(applicationSecurityGroups))

	for _, applicationSecurityGroup := range applicationSecurityGroups {
		m := make(map[string]interface{})
		m[string(Name)] = spotinst.StringValue(applicationSecurityGroup.Name)
		m[string(ResourceGroupName)] = spotinst.StringValue(applicationSecurityGroup.ResourceGroupName)
		result = append(result, m)
	}

	return result
}

func flattenSecurityGroup(networkSecurityGroup *azurev3.SecurityGroup) []interface{} {
	result := make(map[string]interface{})

	result[string(Name)] = spotinst.StringValue(networkSecurityGroup.Name)
	result[string(ResourceGroupName)] = spotinst.StringValue(networkSecurityGroup.ResourceGroupName)

	return []interface{}{result}
}

func flattenPublicIPs(publicIPS []*azurev3.PublicIps) []interface{} {
	result := make([]interface{}, 0, len(publicIPS))

	for _, publicIPS := range publicIPS {
		m := make(map[string]interface{})
		m[string(Name)] = spotinst.StringValue(publicIPS.Name)
		m[string(ResourceGroupName)] = spotinst.StringValue(publicIPS.ResourceGroupName)
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
		if v, ok := attr[string(ApplicationSecurityGroup)]; ok {
			if ApplicationSecurityGroups, err := expandApplicationSecurityGroups(v); err != nil {
				return nil, err
			} else {
				networkInterface.SetApplicationSecurityGroups(ApplicationSecurityGroups)
			}
		}

		if v, ok := attr[string(EnableIPForwarding)].(bool); ok {
			networkInterface.SetEnableIPForwarding(spotinst.Bool(v))
		}

		if v, ok := attr[string(PrivateIPAddresses)]; ok {
			if privateIPAddresses, err := expandPrivateIPAddresses(v); err != nil {
				return nil, err
			} else {
				networkInterface.SetPrivateIpAddresses(privateIPAddresses)
			}
		}

		if v, ok := attr[string(PublicIPSku)].(string); ok {
			networkInterface.SetPublicIpSku(spotinst.String(v))
		}

		if v, ok := attr[string(SecurityGroup)]; ok {
			if securityGroup, err := expandSecurityGroup(v); err != nil {
				return nil, err
			} else {
				if securityGroup != nil {
					networkInterface.SetSecurityGroup(securityGroup)
				}
			}
		}

		if v, ok := attr[string(PublicIPs)]; ok {
			if pips, err := expandPublicIPs(v); err != nil {
				return nil, err
			} else {
				networkInterface.SetPublicIps(pips)
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

func expandApplicationSecurityGroups(data interface{}) ([]*azurev3.ApplicationSecurityGroup, error) {
	list := data.(*schema.Set).List()
	applicationSecurityGroups := make([]*azurev3.ApplicationSecurityGroup, 0, len(list))

	for _, item := range list {
		attr, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		applicationSecurityGroup := &azurev3.ApplicationSecurityGroup{}
		if v, ok := attr[string(Name)].(string); ok && v != "" {
			applicationSecurityGroup.SetName(spotinst.String(v))
		}
		if v, ok := attr[string(ResourceGroupName)].(string); ok && v != "" {
			applicationSecurityGroup.SetResourceGroupName(spotinst.String(v))
		}
		applicationSecurityGroups = append(applicationSecurityGroups, applicationSecurityGroup)
	}

	return applicationSecurityGroups, nil
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

func expandPublicIPs(data interface{}) ([]*azurev3.PublicIps, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		newPublicIPSList := make([]*azurev3.PublicIps, 0, len(list))
		for _, v := range list {
			pips := v.(map[string]interface{})
			publicIP := &azurev3.PublicIps{}

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
	return nil, nil
}

func expandSecurityGroup(data interface{}) (*azurev3.SecurityGroup, error) {
	if list := data.([]interface{}); len(list) > 0 {
		securityGroup := &azurev3.SecurityGroup{}
		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(Name)].(string); ok && v != "" {
				securityGroup.SetName(spotinst.String(v))
			}
			if v, ok := m[string(ResourceGroupName)].(string); ok && v != "" {
				securityGroup.SetResourceGroupName(spotinst.String(v))
			}
		}
		return securityGroup, nil
	}
	return nil, nil
}
