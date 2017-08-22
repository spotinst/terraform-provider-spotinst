package spotinst

// GroupService is an interface for interfacing with the Elastigroup
// endpoints of the Spotinst API.
type GroupService interface {
	CloudProviderAWS() AWSGroupService
	CloudProviderAzure() AzureGroupService
}

// GroupServiceOp handles communication with the balancer related methods
// of the Spotinst API.
type GroupServiceOp struct {
	client *Client
}

var _ GroupService = &GroupServiceOp{}

func (s *GroupServiceOp) CloudProviderAWS() AWSGroupService {
	return &AWSGroupServiceOp{s.client}
}

func (s *GroupServiceOp) CloudProviderAzure() AzureGroupService {
	return &AzureGroupServiceOp{s.client}
}
