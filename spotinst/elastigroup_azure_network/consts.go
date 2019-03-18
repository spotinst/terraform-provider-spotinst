package elastigroup_azure_network

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "azure_network_"
)

const (
	Network             commons.FieldName = "network"
	VirtualNetworkName  commons.FieldName = "virtual_network_name"
	SubnetName          commons.FieldName = "subnet_name"
	ResourceGroupName   commons.FieldName = "resource_group_name"
	AssignPublicIP      commons.FieldName = "assign_public_ip"
	AdditionalIPConfigs commons.FieldName = "additional_ip_configs"

	Name             commons.FieldName = "name"
	PrivateIPVersion commons.FieldName = "private_ip_version"
)
