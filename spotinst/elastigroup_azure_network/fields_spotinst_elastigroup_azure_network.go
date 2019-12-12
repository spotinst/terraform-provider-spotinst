package elastigroup_azure_network

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
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

					string(SubnetName): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(ResourceGroupName): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(AssignPublicIP): {
						Type:     schema.TypeBool,
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
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
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
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
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
func flattenAzureGroupNetwork(network *azure.Network) []interface{} {
	result := make(map[string]interface{})
	result[string(VirtualNetworkName)] = spotinst.StringValue(network.VirtualNetworkName)
	result[string(SubnetName)] = spotinst.StringValue(network.SubnetName)
	result[string(ResourceGroupName)] = spotinst.StringValue(network.ResourceGroupName)
	result[string(AssignPublicIP)] = spotinst.BoolValue(network.AssignPublicIP)

	if network.AdditionalIPConfigs != nil && len(network.AdditionalIPConfigs) > 0 {
		cfgs := make([]interface{}, 0, len(network.AdditionalIPConfigs))
		for _, cfg := range network.AdditionalIPConfigs {
			c := make(map[string]interface{})
			c[string(Name)] = spotinst.StringValue(cfg.Name)
			c[string(PrivateIPVersion)] = spotinst.StringValue(cfg.PrivateIPAddressVersion)

			if c[string(Name)] != nil {
				cfgs = append(cfgs, c)
			}
		}
		result[string(AdditionalIPConfigs)] = cfgs
	}

	return []interface{}{result}
}

func expandAzureGroupNetwork(data interface{}) (*azure.Network, error) {
	network := &azure.Network{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(VirtualNetworkName)].(string); ok && v != "" {
			network.SetVirtualNetworkName(spotinst.String(v))
		}

		if v, ok := m[string(SubnetName)].(string); ok && v != "" {
			network.SetSubnetName(spotinst.String(v))
		}

		if v, ok := m[string(ResourceGroupName)].(string); ok && v != "" {
			network.SetResourceGroupName(spotinst.String(v))
		}

		if v, ok := m[string(AssignPublicIP)].(bool); ok {
			network.SetAssignPublicIP(spotinst.Bool(v))
		}

		if v, ok := m[string(AdditionalIPConfigs)]; ok {
			configs, err := expandAzureGroupAddlConfigs(v)
			if err != nil {
				return nil, err
			}

			if configs != nil {
				network.SetAdditionalIPConfigs(configs)
			}
		} else {
			network.AdditionalIPConfigs = nil
		}
	}
	return network, nil
}

func expandAzureGroupAddlConfigs(data interface{}) ([]*azure.AdditionalIPConfigs, error) {
	list := data.([]interface{})
	addlConfigs := make([]*azure.AdditionalIPConfigs, 0, len(list))

	for _, item := range list {
		attr, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		cfg := &azure.AdditionalIPConfigs{}

		if v, ok := attr[string(Name)].(string); ok && v != "" {
			cfg.SetName(spotinst.String(v))
		}

		if v, ok := attr[string(PrivateIPVersion)].(string); ok && v != "" {
			cfg.SetPrivateIPAddressVersion(spotinst.String(v))
			//cfg.SetPrivateIPAddressVersion(spotinst.String(strings.ToUpper(v)))
		}
		addlConfigs = append(addlConfigs, cfg)
	}

	return addlConfigs, nil
}
