package elastigroup_aws_network_interface

import (
	"fmt"
	"strconv"

	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[NetworkInterface] = commons.NewGenericField(
		commons.ElastigroupAWSNetworkInterface,
		NetworkInterface,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Description): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(DeviceIndex): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(SecondaryPrivateIpAddressCount): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(AssociatePublicIpAddress): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(AssociateIPV6Address): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(DeleteOnTermination): {
						Type:     schema.TypeBool,
						Optional: true,
						Computed: true,
					},

					string(NetworkInterfaceId): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(PrivateIpAddress): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{} = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.NetworkInterfaces != nil {
				networkInterfaces := elastigroup.Compute.LaunchSpecification.NetworkInterfaces
				value = flattenAWSGroupNetworkInterfaces(networkInterfaces)
			}
			if value != nil {
				if err := resourceData.Set(string(NetworkInterface), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(NetworkInterface), err)
				}
			} else {
				if err := resourceData.Set(string(NetworkInterface), []*aws.NetworkInterface{}); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(NetworkInterface), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(NetworkInterface)); ok {

				if interfaces, err := expandAWSGroupNetworkInterfaces(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetNetworkInterfaces(interfaces)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []*aws.NetworkInterface = nil
			if v, ok := resourceData.GetOk(string(NetworkInterface)); ok {
				if interfaces, err := expandAWSGroupNetworkInterfaces(v); err != nil {
					return err
				} else {
					value = interfaces
				}
			}
			elastigroup.Compute.LaunchSpecification.SetNetworkInterfaces(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func flattenAWSGroupNetworkInterfaces(networkInterfaces []*aws.NetworkInterface) []interface{} {
	result := make([]interface{}, 0, len(networkInterfaces))
	for _, iface := range networkInterfaces {
		m := make(map[string]interface{})
		m[string(AssociatePublicIpAddress)] = spotinst.BoolValue(iface.AssociatePublicIPAddress)
		m[string(AssociateIPV6Address)] = spotinst.BoolValue(iface.AssociateIPV6Address)
		m[string(DeleteOnTermination)] = spotinst.BoolValue(iface.DeleteOnTermination)
		m[string(Description)] = spotinst.StringValue(iface.Description)
		m[string(NetworkInterfaceId)] = spotinst.StringValue(iface.ID)
		m[string(PrivateIpAddress)] = spotinst.StringValue(iface.PrivateIPAddress)

		if iface.DeviceIndex != nil {
			m[string(DeviceIndex)] = strconv.Itoa(spotinst.IntValue(iface.DeviceIndex))
		}

		if iface.SecondaryPrivateIPAddressCount != nil {
			m[string(SecondaryPrivateIpAddressCount)] = strconv.Itoa(spotinst.IntValue(iface.SecondaryPrivateIPAddressCount))
		}

		result = append(result, m)
	}
	return result
}

func expandAWSGroupNetworkInterfaces(data interface{}) ([]*aws.NetworkInterface, error) {
	list := data.(*schema.Set).List()
	interfaces := make([]*aws.NetworkInterface, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		networkInterface := &aws.NetworkInterface{}

		if v, ok := m[string(NetworkInterfaceId)].(string); ok && v != "" {
			if v, ok := m[string(AssociatePublicIpAddress)].(bool); ok && v {
				return nil, errors.New("invalid Network interface: associate_public_ip_address must be undefined when using network_interface_id")
			}
			networkInterface.SetId(spotinst.String(v))
		} else {
			// AssociatePublicIp cannot be set at all when NetworkInterfaceId is specified
			if v, ok := m[string(AssociatePublicIpAddress)].(bool); ok {
				networkInterface.SetAssociatePublicIPAddress(spotinst.Bool(v))
			}
		}

		if v, ok := m[string(Description)].(string); ok && v != "" {
			networkInterface.SetDescription(spotinst.String(v))
		}

		if v, ok := m[string(DeviceIndex)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				networkInterface.SetDeviceIndex(spotinst.Int(intVal))
			}
		}

		if v, ok := m[string(SecondaryPrivateIpAddressCount)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				networkInterface.SetSecondaryPrivateIPAddressCount(spotinst.Int(intVal))
			}
		}

		if v, ok := m[string(DeleteOnTermination)].(bool); ok {
			networkInterface.SetDeleteOnTermination(spotinst.Bool(v))
		}

		if v, ok := m[string(PrivateIpAddress)].(string); ok && v != "" {
			networkInterface.SetPrivateIPAddress(spotinst.String(v))
		}

		if v, ok := m[string(AssociateIPV6Address)].(bool); ok {
			networkInterface.SetAssociateIPV6Address(spotinst.Bool(v))
		}

		interfaces = append(interfaces, networkInterface)
	}

	return interfaces, nil
}
