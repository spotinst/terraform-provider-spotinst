package elastigroup_azure_network

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Network                  commons.FieldName = "network"
	VirtualNetworkName       commons.FieldName = "virtual_network_name"
	SubnetName               commons.FieldName = "subnet_name"
	ResourceGroupName        commons.FieldName = "resource_group_name"
	AssignPublicIP           commons.FieldName = "assign_public_ip"
	AdditionalIPConfigs      commons.FieldName = "additional_ip_configs"
	NetworkInterfaces        commons.FieldName = "network_interfaces"
	PrivateIPVersion         commons.FieldName = "private_ip_version"
	IsPrimary                commons.FieldName = "is_primary"
	Name                     commons.FieldName = "name"
	ApplicationSecurityGroup commons.FieldName = "application_security_group"
	EnableIPForwarding       commons.FieldName = "enable_ip_forwarding"
	PrivateIPAddresses       commons.FieldName = "private_ip_addresses"
	PublicIPs                commons.FieldName = "public_ips"
	PublicIPSku              commons.FieldName = "public_ip_sku"
	SecurityGroup            commons.FieldName = "security_group"
)
