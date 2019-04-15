package elastigroup_gcp

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	Name        commons.FieldName = "name"
	Description commons.FieldName = "description"

	AutoHealing            commons.FieldName = "auto_healing"
	AvailabilityZones      commons.FieldName = "availability_zones"
	HealthCheckGracePeriod commons.FieldName = "health_check_grace_period"
	HealthCheckType        commons.FieldName = "health_check_type"
	MaxSize                commons.FieldName = "max_size"
	MinSize                commons.FieldName = "min_size"
	TargetCapacity         commons.FieldName = "desired_capacity"
	Subnets                commons.FieldName = "subnets"
	Region                 commons.FieldName = "region"
	SubnetNames            commons.FieldName = "subnet_names"
	UnhealthyDuration      commons.FieldName = "unhealthy_duration"
)
