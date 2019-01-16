package elastigroup_azure_network

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
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
	}
	return network, nil
}
