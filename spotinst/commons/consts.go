package commons

type FieldName string
type ResourceName string
type LogFormat string

const (
	ElastigroupAWS                 ResourceName = "Elastigroup_AWS"
	ElastigroupInstanceType        ResourceName = "Elastigroup_Instance_Type"
	ElastigroupStrategy            ResourceName = "Elastigroup_Strategy"
	ElastigroupLaunchConfiguration ResourceName = "Elastigroup_Launch_Configuration"

	ResourceFieldOnRead   LogFormat = "onRead() -> %s -> %s"
	ResourceFieldOnCreate LogFormat = "onCreate() -> %s -> %s"
	ResourceFieldOnUpdate LogFormat = "onUpdate() -> %s -> %s"

	ResourceOnDelete LogFormat = "onDelete() -> %s -> started for %s..."
	ResourceOnUpdate LogFormat = "onUpdate() -> %s -> started for %s..."
	ResourceOnRead   LogFormat = "onRead() -> %s -> started for %s..."
	ResourceOnCreate LogFormat = "onCreate() -> %s -> started..."
)
