package stateful_node_azure_network

import (
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

// Network
const (
	Network            commons.FieldName = "network"
	VirtualNetworkName commons.FieldName = "virtual_network_name"
	ResourceGroupName  commons.FieldName = "network_resource_group_name"
	NetworkInterface   commons.FieldName = "network_interface"
)

// NetworkInterfaces
const (
	SubnetName                 commons.FieldName = "subnet_name"
	AssignPublicIP             commons.FieldName = "assign_public_ip"
	IsPrimary                  commons.FieldName = "is_primary"
	PublicIPSku                commons.FieldName = "public_ip_sku"
	EnableIPForwarding         commons.FieldName = "enable_ip_forwarding"
	PrivateIPAddresses         commons.FieldName = "private_ip_addresses"
	NetworkSecurityGroup       commons.FieldName = "network_security_group"
	AdditionalIPConfigurations commons.FieldName = "additional_ip_configurations"
	PublicIPs                  commons.FieldName = "public_ips"
	ApplicationSecurityGroups  commons.FieldName = "application_security_groups"
)

const (
	Name                    commons.FieldName = "name"
	PrivateIPAddressVersion commons.FieldName = "private_ip_address_version"
)
