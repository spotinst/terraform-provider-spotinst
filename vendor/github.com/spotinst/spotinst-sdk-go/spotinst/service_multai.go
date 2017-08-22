package spotinst

// MultaiService is an interface for interfacing with the Multai
// endpoints of the Spotinst API.
type MultaiService interface {
	DeploymentService() DeploymentService
	BalancerService() BalancerService
	CertificateService() CertificateService
}

// MultaiServiceOp handles communication with the balancer related methods
// of the Spotinst API.
type MultaiServiceOp struct {
	client *Client
}

var _ MultaiService = &MultaiServiceOp{}

func (s *MultaiServiceOp) DeploymentService() DeploymentService {
	return &DeploymentServiceOp{s.client}
}

func (s *MultaiServiceOp) BalancerService() BalancerService {
	return &BalancerServiceOp{s.client}
}

func (s *MultaiServiceOp) CertificateService() CertificateService {
	return &CertificateServiceOp{s.client}
}
