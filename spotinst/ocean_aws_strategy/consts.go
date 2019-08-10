package ocean_aws_strategy

import (
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

const (
	SpotPercentage           commons.FieldName = "spot_percentage"
	FallbackToOnDemand       commons.FieldName = "fallback_to_ondemand"
	UtilizeReservedInstances commons.FieldName = "utilize_reserved_instances"
	DrainingTimeout          commons.FieldName = "draining_timeout"
)
