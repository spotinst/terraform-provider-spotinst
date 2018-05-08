package commons

type LogFormat string
type FieldName string
type ResourceName string
type ResourceAffinity string

const (
	FailureFieldReadPattern   = "failed reading field %v - %#v"
	FailureFieldCreatePattern = "failed creating field %v - %#v"
	FailureFieldUpdatePattern = "failed updating field %v - %#v"

	Subscription ResourceAffinity = "Subscription"

	ElastigroupAWS                 ResourceAffinity = "Elastigroup_AWS"
	ElastigroupInstanceType        ResourceAffinity = "Elastigroup_Instance_Type"
	ElastigroupStrategy            ResourceAffinity = "Elastigroup_Strategy"
	ElastigroupStateful            ResourceAffinity = "Elastigroup_Stateful"
	ElastigroupLaunchConfiguration ResourceAffinity = "Elastigroup_Launch_Configuration"
	ElastigroupNetworkInterface    ResourceAffinity = "Elastigroup_Network_Interface"
	ElastigroupScheduledTask       ResourceAffinity = "Elastigroup_Scheduled_Task"
	ElastigroupBlockDevices        ResourceAffinity = "Elastigroup_Block_Device"
	ElastigroupScalingPolicies     ResourceAffinity = "Elastigroup_Scaling_Policies"
	ElastigroupIntegrations        ResourceAffinity = "Elastigroup_Integrations"

	ResourceFieldOnRead   LogFormat = "onRead() -> %s -> %s"
	ResourceFieldOnCreate LogFormat = "onCreate() -> %s -> %s"
	ResourceFieldOnUpdate LogFormat = "onUpdate() -> %s -> %s"

	ResourceOnDelete LogFormat = "onDelete() -> %s -> started for %s..."
	ResourceOnUpdate LogFormat = "onUpdate() -> %s -> started for %s..."
	ResourceOnRead   LogFormat = "onRead() -> %s -> started for %s..."
	ResourceOnCreate LogFormat = "onCreate() -> %s -> started..."
)
