package elastigroup_aws_network_interface

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Prefix = "network_interface_"
)

const (
	NetworkInterface               commons.FieldName = "network_interface"
	Description                    commons.FieldName = "description"
	DeviceIndex                    commons.FieldName = "device_index"
	SecondaryPrivateIpAddressCount commons.FieldName = "secondary_private_ip_address_count"
	AssociatePublicIpAddress       commons.FieldName = "associate_public_ip_address"
	AssociateIPV6Address           commons.FieldName = "associate_ipv6_address"
	DeleteOnTermination            commons.FieldName = "delete_on_termination"
	NetworkInterfaceId             commons.FieldName = "network_interface_id"
	PrivateIpAddress               commons.FieldName = "private_ip_address"
)
