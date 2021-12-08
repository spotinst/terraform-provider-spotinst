package commons

type LogFormat string
type FieldName string
type ResourceName string
type ResourceAffinity string

const (
	FailureFieldReadPattern   = "failed reading field %v - %#v"
	FailureFieldCreatePattern = "failed creating field %v - %#v"
	FailureFieldUpdatePattern = "failed updating field %v - %#v"

	FieldUpdateNotAllowedPattern = "field [%v] is immutable, cannot be changed post creation"
	FieldCreateNotAllowedPattern = "field [%v] can only be changed post creation"

	ProviderToken        FieldName = "token"
	ProviderAccount      FieldName = "account"
	ProviderFeatureFlags FieldName = "feature_flags"

	Subscription                         ResourceAffinity = "Subscription"
	ElastigroupAWSBeanstalk              ResourceAffinity = "ElastigroupAWSBeanstalk"
	ElastigroupAWSBeanstalkScheduledTask ResourceAffinity = "ElastigroupAWSBeanstalk_Scheduled_Task"

	OceanAWS                    ResourceAffinity = "Ocean_AWS"
	OceanAWSInstanceTypes       ResourceAffinity = "Ocean_AWS_Instance_Types"
	OceanAWSAutoScaling         ResourceAffinity = "Ocean_AWS_Auto_Scaling"
	OceanAWSStrategy            ResourceAffinity = "Ocean_AWS_Strategy"
	OceanAWSScheduling          ResourceAffinity = "Ocean_AWS_Scheduling"
	OceanAWSLaunchConfiguration ResourceAffinity = "Ocean_AWS_Launch_Configuration"

	OceanAWSLaunchSpec ResourceAffinity = "Ocean_AWS_Launch_Spec"

	OceanGKE                          ResourceAffinity = "Ocean_GKE"
	OceanGKEImport                    ResourceAffinity = "Ocean_GKE_Import"
	OceanGKEImportScheduling          ResourceAffinity = "Ocean_GKE_Import_Scheduling"
	OceanGKEImportAutoScaler          ResourceAffinity = "Ocean_GKE_Import_Auto_Scaler"
	OceanGKEImportLaunchSpecification ResourceAffinity = "Ocean_GKE_Import_Launch_Specification"

	OceanGKEInstanceTypes      ResourceAffinity = "Ocean_GKE_Instance_Types"
	OceanGKEAutoScaling        ResourceAffinity = "Ocean_GKE_Auto_Scaling"
	OceanGKEStrategy           ResourceAffinity = "Ocean_GKE_Strategy"
	OceanGKELaunchSpec         ResourceAffinity = "Ocean_GKE_Launch_Spec"
	OceanGKELaunchSpecStrategy ResourceAffinity = "Ocean_GKE_Launch_Spec_Strategy"

	OceanGKELaunchSpecImport ResourceAffinity = "Ocean_GKE_Launch_Spec_Import"
	OceanGKENetworkInterface ResourceAffinity = "Ocean_GKE_Network_Interface"

	OceanAKS                    ResourceAffinity = "Ocean_AKS"
	OceanAKSLogin               ResourceAffinity = "Ocean_AKS_Login"
	OceanAKSStrategy            ResourceAffinity = "Ocean_AKS_Strategy"
	OceanAKSHealth              ResourceAffinity = "Ocean_AKS_Health"
	OceanAKSAutoScaling         ResourceAffinity = "Ocean_AKS_Auto_Scaling"
	OceanAKSLaunchSpecification ResourceAffinity = "Ocean_AKS_Launch_Specification"
	OceanAKSImage               ResourceAffinity = "Ocean_AKS_Image"
	OceanAKSExtensions          ResourceAffinity = "Ocean_AKS_Extensions"
	OceanAKSLoadBalancers       ResourceAffinity = "Ocean_AKS_Load_Balancers_Config"
	OceanAKSNetwork             ResourceAffinity = "Ocean_AKS_Network"
	OceanAKSVMSizes             ResourceAffinity = "Ocean_AKS_VMSizes"

	OceanAKSVirtualNodeGroup                    ResourceAffinity = "Ocean_AKS_virtual_node_group"
	OceanAKSVirtualNodeGroupAutoScaling         ResourceAffinity = "Ocean_AKS_virtual_node_group_Auto_Scaling"
	OceanAKSVirtualNodeGroupLaunchSpecification ResourceAffinity = "Ocean_AKS_virtual_node_group_launch_specification"

	OceanECS                    ResourceAffinity = "Ocean_ECS"
	OceanECSAutoScaler          ResourceAffinity = "Ocean_ECS_Auto_Scaler"
	OceanECSInstanceTypes       ResourceAffinity = "Ocean_ECS_Instance_Types"
	OceanECSLaunchSpecification ResourceAffinity = "Ocean_ECS_Launch_Specification"
	OceanECSStrategy            ResourceAffinity = "Ocean_ECS_Strategy"
	OceanECSScheduling          ResourceAffinity = "Ocean_ECS_Scheduling"
	OceanECSOptimizeImages      ResourceAffinity = "Ocean_ECS_Optimize_Images"

	OceanECSLaunchSpec ResourceAffinity = "Ocean_ECS_Launch_Spec"

	ElastigroupAWS                    ResourceAffinity = "Elastigroup_AWS"
	ElastigroupAWSInstanceType        ResourceAffinity = "Elastigroup_AWS_Instance_Type"
	ElastigroupAWSStrategy            ResourceAffinity = "Elastigroup_AWS_Strategy"
	ElastigroupAWSStateful            ResourceAffinity = "Elastigroup_AWS_Stateful"
	ElastigroupAWSLaunchConfiguration ResourceAffinity = "Elastigroup_AWS_Launch_Configuration"
	ElastigroupAWSNetworkInterface    ResourceAffinity = "Elastigroup_AWS_Network_Interface"
	ElastigroupAWSScheduledTask       ResourceAffinity = "Elastigroup_AWS_Scheduled_Task"
	ElastigroupAWSBlockDevices        ResourceAffinity = "Elastigroup_AWS_Block_Device"
	ElastigroupAWSScalingPolicies     ResourceAffinity = "Elastigroup_AWS_Scaling_Policies"
	ElastigroupAWSIntegrations        ResourceAffinity = "Elastigroup_AWS_Integrations"

	ManagedInstanceAWS                    ResourceAffinity = "Managed_Instance_AWS"
	ManagedInstanceAWSStrategy            ResourceAffinity = "Managed_Instance_AWS_Strategy"
	ManagedInstanceAWSPersistence         ResourceAffinity = "Managed_Instance_AWS_Persistence"
	ManagedInstanceAWSHealthCheck         ResourceAffinity = "Managed_Instance_AWS_HealthCheck"
	ManagedInstanceAWSIntegrations        ResourceAffinity = "Managed_Instance_AWS_Integrations"
	ManagedInstanceAWSCompute             ResourceAffinity = "Managed_Instance_AWS_Compute"
	ManagedInstanceAWSLaunchSpecification ResourceAffinity = "Managed_Instance_AWS_Launch_Specification"
	ManagedInstanceAWSNetworkInterfaces   ResourceAffinity = "Managed_Instance_AWS_Network_Interfaces"
	ManagedInstanceAWSScheduling          ResourceAffinity = "Managed_Instance_AWS_Scheduling"
	ManagedInstanceAWSComputeInstanceType ResourceAffinity = "Managed_Instance_AWS_Compute_Instance_Type"

	ElastigroupGCP                    ResourceAffinity = "Elastigroup_GCP"
	ElastigroupGCPDisk                ResourceAffinity = "Elastigroup_GCP_Disk"
	ElastigroupGCPGPU                 ResourceAffinity = "Elastigroup_GPC_GPU"
	ElastigroupGCPInstanceType        ResourceAffinity = "Elastigroup_GCP_Instance_Type"
	ElastigroupGCPIntegrations        ResourceAffinity = "Elastigroup_GCP_Integrations"
	ElastigroupGCPLaunchConfiguration ResourceAffinity = "Elastigroup_GCP_Launch_Configuration"
	ElastigroupGCPNetworkInterface    ResourceAffinity = "Elastigroup_GCP_Network_Interface"
	ElastigroupGCPScalingPolicies     ResourceAffinity = "Elastigroup_GCP_Scaling_Policies"
	ElastigroupGCPScheduledTask       ResourceAffinity = "Elastigroup_GCP_Scheduled_Task"
	ElastigroupGCPStrategy            ResourceAffinity = "Elastigroup_GCP_Strategy"

	ElastigroupGKE ResourceAffinity = "Elastigroup_GKE"

	ElastigroupAzure                    ResourceAffinity = "Elastigroup_Azure"
	ElastigroupAzureStrategy            ResourceAffinity = "Elastigroup_Azure_Strategy"
	ElastigroupAzureLogin               ResourceAffinity = "Elastigroup_Azure_Login"
	ElastigroupAzureNetwork             ResourceAffinity = "Elastigroup_Azure_Network"
	ElastigroupAzureLoadBalancers       ResourceAffinity = "Elastigroup_Azure_Load_Balancers"
	ElastigroupAzureVMSizes             ResourceAffinity = "Elastigroup_Azure_VM_Sizes"
	ElastigroupAzureImage               ResourceAffinity = "Elastigroup_Azure_Image"
	ElastigroupAzureIntegrations        ResourceAffinity = "Elastigroup_Azure_Integrations"
	ElastigroupAzureLaunchConfiguration ResourceAffinity = "Elastigroup_Azure_Launch_Configuration"
	ElastigroupAzureHealthCheck         ResourceAffinity = "Elastigroup_Azure_Health_Check"
	ElastigroupAzureScalingPolicies     ResourceAffinity = "Elastigroup_Azure_Scaling_Policies"
	ElastigroupAzureScheduledTask       ResourceAffinity = "Elastigroup_Azure_Scheduled_Task"
	ElastigroupAzureLaunchSpecification ResourceAffinity = "Elastigroup_Azure_Launch_Specification"

	MRScalerAWS                    ResourceAffinity = "MRScaler_AWS"
	MRScalerAWSTaskScalingPolicies ResourceAffinity = "MRScaler_Task_AWS_Scaling_Polices"
	MRScalerAWSCoreScalingPolicies ResourceAffinity = "MRScaler_Core_AWS_Scaling_Polices"
	MRScalerAWSCoreGroup           ResourceAffinity = "MRScaler_AWS_Core_Group"
	MRScalerAWSMasterGroup         ResourceAffinity = "MRScaler_AWS_Master_Group"
	MRScalerAWSTaskGroup           ResourceAffinity = "MRScaler_AWS_Task_Group"
	MRScalerAWSStrategy            ResourceAffinity = "MRScaler_AWS_Strategy"
	MRScalerAWSCluster             ResourceAffinity = "MRScaler_AWS_Cluster"
	MRScalerAWSScheduledTask       ResourceAffinity = "MRScaler_AWS_Scheduled_Task"
	MRScalerAWSTerminationPolicies ResourceAffinity = "MRScaler_AWS_Termination_Policies"

	MultaiBalancer    ResourceAffinity = "Multai_Balancer"
	MultaiDeployment  ResourceAffinity = "Multai_Deployment"
	MultaiListener    ResourceAffinity = "Multai_Listener"
	MultaiRoutingRule ResourceAffinity = "Multai_Routing_Rule"
	MultaiTarget      ResourceAffinity = "Multai_Target"
	MultaiTargetSet   ResourceAffinity = "Multai_Target_Set"

	HealthCheck ResourceAffinity = "Health_Check"

	SuspendProcesses ResourceAffinity = "Suspend_Processes"

	ResourceFieldOnRead   LogFormat = "onRead() -> %s -> %s"
	ResourceFieldOnCreate LogFormat = "onCreate() -> %s -> %s"
	ResourceFieldOnUpdate LogFormat = "onUpdate() -> %s -> %s"
	ResourceFieldOnMerge  LogFormat = "onMerge() -> %s -> %s"

	ResourceOnDelete LogFormat = "onDelete() -> %s -> started for %s..."
	ResourceOnUpdate LogFormat = "onUpdate() -> %s -> started for %s..."
	ResourceOnRead   LogFormat = "onRead() -> %s -> started for %s..."
	ResourceOnCreate LogFormat = "onCreate() -> %s -> started..."
)
