package commons

type LogFormat string
type FieldName string
type ResourceName string
type ResourceAffinity string

const (
	FailureFieldReadPattern   = "failed reading field %v - %#v"
	FailureFieldCreatePattern = "failed creating field %v - %#v"
	FailureFieldUpdatePattern = "failed updating field %v - %#v"

	FieldUpdateNotAllowedPattern = "field [%v] is immutable, cannot be changed post group creation"

	ProviderToken   FieldName = "token"
	ProviderAccount FieldName = "account"

	Subscription            ResourceAffinity = "Subscription"
	ElastigroupAWSBeanstalk ResourceAffinity = "ElastigroupAWSBeanstalk"

	OceanAWS                    ResourceAffinity = "Ocean_AWS"
	OceanAWSInstanceTypes       ResourceAffinity = "Ocean_AWS_Instance_Types"
	OceanAWSAutoScaling         ResourceAffinity = "Ocean_AWS_Auto_Scaling"
	OceanAWSStrategy            ResourceAffinity = "Ocean_AWS_Strategy"
	OceanAWSLaunchConfiguration ResourceAffinity = "Ocean_AWS_Launch_Configuration"

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
