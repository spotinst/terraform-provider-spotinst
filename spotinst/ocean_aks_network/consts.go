package ocean_aks_network

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Network            commons.FieldName = "network"
	VirtualNetworkName commons.FieldName = "virtual_network_name"
	ResourceGroupName  commons.FieldName = "resource_group_name"
)

const (
	AdditionalIPConfig commons.FieldName = "additional_ip_config"
	PrivateIPVersion   commons.FieldName = "private_ip_version"
	Name               commons.FieldName = "name"
)

const (
	NetworkInterface commons.FieldName = "network_interface"
	SubnetName       commons.FieldName = "subnet_name"
	AssignPublicIP   commons.FieldName = "assign_public_ip"
	IsPrimary        commons.FieldName = "is_primary"
)

const (
	SecurityGroup              commons.FieldName = "security_group"
	SecurityGroupName          commons.FieldName = "name"
	SecurityGroupResourceGroup commons.FieldName = "resource_group_name"
)
