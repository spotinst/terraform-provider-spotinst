package ocean_right_sizing_cluster_config

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	OceanId           commons.FieldName = "ocean_id"
	ClusterIdentifier commons.FieldName = "cluster_identifier"

	Config                          commons.FieldName = "config"
	AdjustLimitOnDownsize           commons.FieldName = "adjust_limit_on_downsize"
	DownsideOnly                    commons.FieldName = "downside_only"
	RecommendationsCpuPercentile    commons.FieldName = "recommendations_cpu_percentile"
	RecommendationsMemoryPercentile commons.FieldName = "recommendations_memory_percentile"
)
