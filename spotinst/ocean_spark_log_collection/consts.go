package ocean_spark_log_collection

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	LogCollection commons.FieldName = "log_collection"
	// Deprecated: CollectDriverLogs is obsolete, exists for backward compatibility only,
	// and should not be used. Please use CollectAppLogs instead.
	CollectDriverLogs commons.FieldName = "collect_driver_logs"
	CollectAppLogs    commons.FieldName = "collect_app_logs"
)
