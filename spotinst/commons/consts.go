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
	FieldCreateNotAllowedPattern = "field [%v] can only be changed after the group is created"

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

	ElastigroupGCP                    ResourceAffinity = "Elastigroup_GCP"
	ElastigroupGCPDisk                ResourceAffinity = "Elastigroup_GCP_Disk"
	ElastigroupGCPGPU                 ResourceAffinity = "Elastigroup_GPC_GPU"
	ElastigroupGCPInstanceType        ResourceAffinity = "Elastigroup_GCP_Instance_Type"
	ElastigroupGCPLaunchConfiguration ResourceAffinity = "Elastigroup_GCP_Launch_Configuration"
	ElastigroupGCPNetworkInterface    ResourceAffinity = "Elastigroup_GCP_Network_Interface"
	ElastigroupGCPScalingPolicies     ResourceAffinity = "Elastigroup_GCP_Scaling_Policies"
	ElastigroupGCPStrategy            ResourceAffinity = "Elastigroup_GCP_Strategy"

	ElastigroupGKE             ResourceAffinity = "Elastigroup_GKE"
	ElastigroupGKEInstanceType ResourceAffinity = "Elastigroup_GKE_Instance_Type"

	ElastigroupAzure                    ResourceAffinity = "Elastigroup_Azure"
	ElastigroupAzureStrategy            ResourceAffinity = "Elastigroup_Azure_Strategy"
	ElastigroupAzureLogin               ResourceAffinity = "Elastigroup_Azure_Login"
	ElastigroupAzureNetwork             ResourceAffinity = "Elastigroup_Azure_Network"
	ElastigroupAzureLoadBalancers       ResourceAffinity = "Elastigroup_Azure_Load_Balancers"
	ElastigroupAzureVMSizes             ResourceAffinity = "Elastigroup_Azure_VM_Sizes"
	ElastigroupAzureImage               ResourceAffinity = "Elastigroup_Azure_Image"
	ElastigroupAzureLaunchConfiguration ResourceAffinity = "Elastigroup_Azure_Launch_Configuration"
	ElastigroupAzureHealthCheck         ResourceAffinity = "Elastigroup_Azure_Health_Check"
	ElastigroupAzureScheduledTask       ResourceAffinity = "Elastigroup_Azure_Scheduled_Task"

	MRScalerAWS                    ResourceAffinity = "MRScaler_AWS"
	MRScalerAWSTaskScalingPolicies ResourceAffinity = "MRScaler_Task_AWS_Scaling_Polices"
	MRScalerAWSCoreScalingPolicies ResourceAffinity = "MRScaler_Core_AWS_Scaling_Polices"
	MRScalerAWSCoreGroup           ResourceAffinity = "MRScaler_AWS_Core_Group"
	MRScalerAWSMasterGroup         ResourceAffinity = "MRScaler_AWS_Master_Group"
	MRScalerAWSTaskGroup           ResourceAffinity = "MRScaler_AWS_Task_Group"
	MRScalerAWSStrategy            ResourceAffinity = "MRScaler_AWS_Strategy"
	MRScalerAWSCluster             ResourceAffinity = "MRScaler_AWS_Cluster"
	MRScalerAWSScheduledTask       ResourceAffinity = "MRScaler_AWS_Scheduled_Task"

	MultaiBalancer    ResourceAffinity = "Multai_Balancer"
	MultaiDeployment  ResourceAffinity = "Multai_Deployment"
	MultaiListener    ResourceAffinity = "Multai_Listener"
	MultaiRoutingRule ResourceAffinity = "Multai_Routing_Rule"
	MultaiTarget      ResourceAffinity = "Multai_Target"
	MultaiTargetSet   ResourceAffinity = "Multai_Target_Set"

	ResourceFieldOnRead   LogFormat = "onRead() -> %s -> %s"
	ResourceFieldOnCreate LogFormat = "onCreate() -> %s -> %s"
	ResourceFieldOnUpdate LogFormat = "onUpdate() -> %s -> %s"

	ResourceOnDelete LogFormat = "onDelete() -> %s -> started for %s..."
	ResourceOnUpdate LogFormat = "onUpdate() -> %s -> started for %s..."
	ResourceOnRead   LogFormat = "onRead() -> %s -> started for %s..."
	ResourceOnCreate LogFormat = "onCreate() -> %s -> started..."
)
