package ocean_gke_import_strategy

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Strategy                 commons.FieldName = "strategy"
	DrainingTimeout          commons.FieldName = "draining_timeout"
	ProvisioningModel        commons.FieldName = "provisioning_model"
	PreemptiblePercentage    commons.FieldName = "preemptible_percentage"
	ShouldUtilizeCommitments commons.FieldName = "should_utilize_commitments"
	ScalingOrientation       commons.FieldName = "scaling_orientation"
)
