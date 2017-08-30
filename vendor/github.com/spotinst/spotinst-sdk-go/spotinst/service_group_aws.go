package spotinst

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

// A AWSProduct represents the type of an operating system.
type AWSProduct int

const (
	// AWSProductWindows represents the Windows product.
	AWSProductWindows AWSProduct = iota

	// AWSProductWindowsVPC represents the Windows (Amazon VPC) product.
	AWSProductWindowsVPC

	// AWSProductLinuxUnix represents the Linux/Unix product.
	AWSProductLinuxUnix

	// AWSProductLinuxUnixVPC represents the Linux/Unix (Amazon VPC) product.
	AWSProductLinuxUnixVPC

	// AWSProductSUSELinux represents the SUSE Linux product.
	AWSProductSUSELinux

	// AWSProductSUSELinuxVPC represents the SUSE Linux (Amazon VPC) product.
	AWSProductSUSELinuxVPC
)

var AWSProduct_name = map[AWSProduct]string{
	AWSProductWindows:      "Windows",
	AWSProductWindowsVPC:   "Windows (Amazon VPC)",
	AWSProductLinuxUnix:    "Linux/UNIX",
	AWSProductLinuxUnixVPC: "Linux/UNIX (Amazon VPC)",
	AWSProductSUSELinux:    "SUSE Linux",
	AWSProductSUSELinuxVPC: "SUSE Linux (Amazon VPC)",
}

var AWSProduct_value = map[string]AWSProduct{
	"Windows":                 AWSProductWindows,
	"Windows (Amazon VPC)":    AWSProductWindowsVPC,
	"Linux/UNIX":              AWSProductLinuxUnix,
	"Linux/UNIX (Amazon VPC)": AWSProductLinuxUnixVPC,
	"SUSE Linux":              AWSProductSUSELinux,
	"SUSE Linux (Amazon VPC)": AWSProductSUSELinuxVPC,
}

func (p AWSProduct) String() string {
	return AWSProduct_name[p]
}

// AWSGroupService is an interface for interfacing with the AWSGroup
// endpoints of the Spotinst API.
type AWSGroupService interface {
	List(context.Context, *ListAWSGroupInput) (*ListAWSGroupOutput, error)
	Create(context.Context, *CreateAWSGroupInput) (*CreateAWSGroupOutput, error)
	Read(context.Context, *ReadAWSGroupInput) (*ReadAWSGroupOutput, error)
	Update(context.Context, *UpdateAWSGroupInput) (*UpdateAWSGroupOutput, error)
	Delete(context.Context, *DeleteAWSGroupInput) (*DeleteAWSGroupOutput, error)
	Status(context.Context, *StatusAWSGroupInput) (*StatusAWSGroupOutput, error)
	Detach(context.Context, *DetachAWSGroupInput) (*DetachAWSGroupOutput, error)
	Roll(context.Context, *RollAWSGroupInput) (*RollAWSGroupOutput, error)
}

// AWSGroupServiceOp handles communication with the balancer related methods
// of the Spotinst API.
type AWSGroupServiceOp struct {
	client *Client
}

var _ AWSGroupService = &AWSGroupServiceOp{}

type AWSGroup struct {
	ID          *string              `json:"id,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Description *string              `json:"description,omitempty"`
	Capacity    *AWSGroupCapacity    `json:"capacity,omitempty"`
	Compute     *AWSGroupCompute     `json:"compute,omitempty"`
	Strategy    *AWSGroupStrategy    `json:"strategy,omitempty"`
	Scaling     *AWSGroupScaling     `json:"scaling,omitempty"`
	Scheduling  *AWSGroupScheduling  `json:"scheduling,omitempty"`
	Integration *AWSGroupIntegration `json:"thirdPartiesIntegration,omitempty"`

	// forceSendFields is a list of field names (e.g. "Keys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	forceSendFields []string `json:"-"`

	// nullFields is a list of field names (e.g. "Keys") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	nullFields []string `json:"-"`
}

type AWSGroupIntegration struct {
	EC2ContainerService *AWSGroupEC2ContainerServiceIntegration `json:"ecs,omitempty"`
	ElasticBeanstalk    *AWSGroupElasticBeanstalkIntegration    `json:"elasticBeanstalk,omitempty"`
	CodeDeploy          *AWSGroupCodeDeployIntegration          `json:"codeDeploy,omitempty"`
	Rancher             *AWSGroupRancherIntegration             `json:"rancher,omitempty"`
	Kubernetes          *AWSGroupKubernetesIntegration          `json:"kubernetes,omitempty"`
	Mesosphere          *AWSGroupMesosphereIntegration          `json:"mesosphere,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupCodeDeployIntegration struct {
	DeploymentGroups           []*AWSGroupCodeDeployIntegrationDeploymentGroup `json:"deploymentGroups,omitempty"`
	CleanUpOnFailure           *bool                                           `json:"cleanUpOnFailure,omitempty"`
	TerminateInstanceOnFailure *bool                                           `json:"terminateInstanceOnFailure,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupCodeDeployIntegrationDeploymentGroup struct {
	ApplicationName     *string `json:"applicationName,omitempty"`
	DeploymentGroupName *string `json:"deploymentGroupName,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupRancherIntegration struct {
	MasterHost *string `json:"masterHost,omitempty"`
	AccessKey  *string `json:"accessKey,omitempty"`
	SecretKey  *string `json:"secretKey,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupElasticBeanstalkIntegration struct {
	EnvironmentID *string `json:"environmentId,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupEC2ContainerServiceIntegration struct {
	ClusterName *string                                          `json:"clusterName,omitempty"`
	AutoScale   *AWSGroupEC2ContainerServiceIntegrationAutoScale `json:"autoScale,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupEC2ContainerServiceIntegrationAutoScale struct {
	IsEnabled *bool                                                    `json:"isEnabled,omitempty"`
	Cooldown  *int                                                     `json:"cooldown,omitempty"`
	Headroom  *AWSGroupEC2ContainerServiceIntegrationAutoScaleHeadroom `json:"headroom,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupEC2ContainerServiceIntegrationAutoScaleHeadroom struct {
	CPUPerUnit    *int `json:"cpuPerUnit,omitempty"`
	MemoryPerUnit *int `json:"memoryPerUnit,omitempty"`
	NumOfUnits    *int `json:"numOfUnits,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupKubernetesIntegration struct {
	Server *string `json:"apiServer,omitempty"`
	Token  *string `json:"token,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupMesosphereIntegration struct {
	Server *string `json:"apiServer,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupScheduling struct {
	Tasks []*AWSGroupScheduledTask `json:"tasks,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupScheduledTask struct {
	IsEnabled           *bool   `json:"isEnabled,omitempty"`
	Frequency           *string `json:"frequency,omitempty"`
	CronExpression      *string `json:"cronExpression,omitempty"`
	TaskType            *string `json:"taskType,omitempty"`
	ScaleTargetCapacity *int    `json:"scaleTargetCapacity,omitempty"`
	ScaleMinCapacity    *int    `json:"scaleMinCapacity,omitempty"`
	ScaleMaxCapacity    *int    `json:"scaleMaxCapacity,omitempty"`
	BatchSizePercentage *int    `json:"batchSizePercentage,omitempty"`
	GracePeriod         *int    `json:"gracePeriod,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupScaling struct {
	Up   []*AWSGroupScalingPolicy `json:"up,omitempty"`
	Down []*AWSGroupScalingPolicy `json:"down,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupScalingPolicy struct {
	PolicyName        *string                           `json:"policyName,omitempty"`
	MetricName        *string                           `json:"metricName,omitempty"`
	Statistic         *string                           `json:"statistic,omitempty"`
	Unit              *string                           `json:"unit,omitempty"`
	Threshold         *float64                          `json:"threshold,omitempty"`
	Adjustment        *int                              `json:"adjustment,omitempty"`
	MinTargetCapacity *int                              `json:"minTargetCapacity,omitempty"`
	MaxTargetCapacity *int                              `json:"maxTargetCapacity,omitempty"`
	Namespace         *string                           `json:"namespace,omitempty"`
	EvaluationPeriods *int                              `json:"evaluationPeriods,omitempty"`
	Period            *int                              `json:"period,omitempty"`
	Cooldown          *int                              `json:"cooldown,omitempty"`
	Operator          *string                           `json:"operator,omitempty"`
	Dimensions        []*AWSGroupScalingPolicyDimension `json:"dimensions,omitempty"`
	Action            *AWSGroupScalingPolicyAction      `json:"action,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupScalingPolicyAction struct {
	Type              *string `json:"type,omitempty"`
	Adjustment        *string `json:"adjustment,omitempty"`
	MinTargetCapacity *string `json:"minTargetCapacity,omitempty"`
	MaxTargetCapacity *string `json:"maxTargetCapacity,omitempty"`
	Maximum           *string `json:"maximum,omitempty"`
	Minimum           *string `json:"minimum,omitempty"`
	Target            *string `json:"target,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupScalingPolicyDimension struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupStrategy struct {
	Risk                     *float64                  `json:"risk,omitempty"`
	OnDemandCount            *int                      `json:"onDemandCount,omitempty"`
	DrainingTimeout          *int                      `json:"drainingTimeout,omitempty"`
	AvailabilityVsCost       *string                   `json:"availabilityVsCost,omitempty"`
	UtilizeReservedInstances *bool                     `json:"utilizeReservedInstances,omitempty"`
	FallbackToOnDemand       *bool                     `json:"fallbackToOd,omitempty"`
	SpinUpTime               *int                      `json:"spinUpTime,omitempty"`
	Signals                  []*AWSGroupStrategySignal `json:"signals,omitempty"`
	Persistence              *AWSGroupPersistence      `json:"persistence,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupPersistence struct {
	ShouldPersistPrivateIp    *bool `json:"shouldPersistPrivateIp,omitempty"`
	ShouldPersistBlockDevices *bool `json:"shouldPersistBlockDevices,omitempty"`
	ShouldPersistRootDevice   *bool `json:"shouldPersistRootDevice,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupStrategySignal struct {
	Name    *string `json:"name,omitempty"`
	Timeout *int    `json:"timeout,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupCapacity struct {
	Minimum *int    `json:"minimum,omitempty"`
	Maximum *int    `json:"maximum,omitempty"`
	Target  *int    `json:"target,omitempty"`
	Unit    *string `json:"unit,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupCompute struct {
	Product             *string                             `json:"product,omitempty"`
	InstanceTypes       *AWSGroupComputeInstanceType        `json:"instanceTypes,omitempty"`
	LaunchSpecification *AWSGroupComputeLaunchSpecification `json:"launchSpecification,omitempty"`
	AvailabilityZones   []*AWSGroupComputeAvailabilityZone  `json:"availabilityZones,omitempty"`
	ElasticIPs          []string                            `json:"elasticIps,omitempty"`
	EBSVolumePool       []*AWSGroupComputeEBSVolume         `json:"ebsVolumePool,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupComputeEBSVolume struct {
	DeviceName *string  `json:"deviceName,omitempty"`
	VolumeIDs  []string `json:"volumeIds,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupComputeInstanceType struct {
	OnDemand *string                              `json:"ondemand,omitempty"`
	Spot     []string                             `json:"spot,omitempty"`
	Weights  []*AWSGroupComputeInstanceTypeWeight `json:"weights,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupComputeInstanceTypeWeight struct {
	InstanceType *string `json:"instanceType,omitempty"`
	Weight       *int    `json:"weightedCapacity,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupComputeAvailabilityZone struct {
	Name     *string `json:"name,omitempty"`
	SubnetID *string `json:"subnetId,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupComputeLaunchSpecification struct {
	LoadBalancerNames                             []string                            `json:"loadBalancerNames,omitempty"`
	LoadBalancersConfig                           *AWSGroupComputeLoadBalancersConfig `json:"loadBalancersConfig,omitempty"`
	SecurityGroupIDs                              []string                            `json:"securityGroupIds,omitempty"`
	HealthCheckType                               *string                             `json:"healthCheckType,omitempty"`
	HealthCheckGracePeriod                        *int                                `json:"healthCheckGracePeriod,omitempty"`
	HealthCheckUnhealthyDurationBeforeReplacement *int                                `json:"healthCheckUnhealthyDurationBeforeReplacement,omitempty"`
	ImageID                                       *string                             `json:"imageId,omitempty"`
	KeyPair                                       *string                             `json:"keyPair,omitempty"`
	UserData                                      *string                             `json:"userData,omitempty"`
	ShutdownScript                                *string                             `json:"shutdownScript,omitempty"`
	Tenancy                                       *string                             `json:"tenancy,omitempty"`
	Monitoring                                    *bool                               `json:"monitoring,omitempty"`
	EBSOptimized                                  *bool                               `json:"ebsOptimized,omitempty"`
	IAMInstanceProfile                            *AWSGroupComputeIAMInstanceProfile  `json:"iamRole,omitempty"`
	BlockDevices                                  []*AWSGroupComputeBlockDevice       `json:"blockDeviceMappings,omitempty"`
	NetworkInterfaces                             []*AWSGroupComputeNetworkInterface  `json:"networkInterfaces,omitempty"`
	Tags                                          []*AWSGroupComputeTag               `json:"tags,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupComputeLoadBalancersConfig struct {
	LoadBalancers []*AWSGroupComputeLoadBalancer `json:"loadBalancers,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupComputeLoadBalancer struct {
	Name *string `json:"name,omitempty"`
	Arn  *string `json:"arn,omitempty"`
	Type *string `json:"type,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupComputeNetworkInterface struct {
	ID                             *string  `json:"networkInterfaceId,omitempty"`
	Description                    *string  `json:"description,omitempty"`
	DeviceIndex                    *int     `json:"deviceIndex,omitempty"`
	SecondaryPrivateIPAddressCount *int     `json:"secondaryPrivateIpAddressCount,omitempty"`
	AssociatePublicIPAddress       *bool    `json:"associatePublicIpAddress,omitempty"`
	DeleteOnTermination            *bool    `json:"deleteOnTermination,omitempty"`
	SecurityGroupsIDs              []string `json:"groups,omitempty"`
	PrivateIPAddress               *string  `json:"privateIpAddress,omitempty"`
	SubnetID                       *string  `json:"subnetId,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupComputeBlockDevice struct {
	DeviceName  *string             `json:"deviceName,omitempty"`
	VirtualName *string             `json:"virtualName,omitempty"`
	EBS         *AWSGroupComputeEBS `json:"ebs,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupComputeEBS struct {
	DeleteOnTermination *bool   `json:"deleteOnTermination,omitempty"`
	Encrypted           *bool   `json:"encrypted,omitempty"`
	SnapshotID          *string `json:"snapshotId,omitempty"`
	VolumeType          *string `json:"volumeType,omitempty"`
	VolumeSize          *int    `json:"volumeSize,omitempty"`
	IOPS                *int    `json:"iops,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupComputeIAMInstanceProfile struct {
	Name *string `json:"name,omitempty"`
	Arn  *string `json:"arn,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSGroupComputeTag struct {
	Key   *string `json:"tagKey,omitempty"`
	Value *string `json:"tagValue,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AWSInstance struct {
	ID               *string    `json:"instanceId,omitempty"`
	SpotRequestID    *string    `json:"spotInstanceRequestId,omitempty"`
	InstanceType     *string    `json:"instanceType,omitempty"`
	Status           *string    `json:"status,omitempty"`
	Product          *string    `json:"product,omitempty"`
	AvailabilityZone *string    `json:"availabilityZone,omitempty"`
	CreatedAt        *time.Time `json:"createdAt,omitempty"`
}

type AWSGroupRollStrategy struct {
	Action               *string `json:"action,omitempty"`
	ShouldDrainInstances *bool   `json:"shouldDrainInstances,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListAWSGroupInput struct{}

type ListAWSGroupOutput struct {
	Groups []*AWSGroup `json:"groups,omitempty"`
}

type CreateAWSGroupInput struct {
	Group *AWSGroup `json:"group,omitempty"`
}

type CreateAWSGroupOutput struct {
	Group *AWSGroup `json:"group,omitempty"`
}

type ReadAWSGroupInput struct {
	GroupID *string `json:"groupId,omitempty"`
}

type ReadAWSGroupOutput struct {
	Group *AWSGroup `json:"group,omitempty"`
}

type UpdateAWSGroupInput struct {
	Group *AWSGroup `json:"group,omitempty"`
}

type UpdateAWSGroupOutput struct {
	Group *AWSGroup `json:"group,omitempty"`
}

type DeleteAWSGroupInput struct {
	GroupID *string `json:"groupId,omitempty"`
}

type DeleteAWSGroupOutput struct{}

type StatusAWSGroupInput struct {
	GroupID *string `json:"groupId,omitempty"`
}

type StatusAWSGroupOutput struct {
	Instances []*AWSInstance `json:"instances,omitempty"`
}

type DetachAWSGroupInput struct {
	GroupID                       *string  `json:"groupId,omitempty"`
	InstanceIDs                   []string `json:"instancesToDetach,omitempty"`
	ShouldDecrementTargetCapacity *bool    `json:"shouldDecrementTargetCapacity,omitempty"`
	ShouldTerminateInstances      *bool    `json:"shouldTerminateInstances,omitempty"`
	DrainingTimeout               *int     `json:"drainingTimeout,omitempty"`
}

type DetachAWSGroupOutput struct{}

type RollAWSGroupInput struct {
	GroupID             *string               `json:"groupId,omitempty"`
	BatchSizePercentage *int                  `json:"batchSizePercentage,omitempty"`
	GracePeriod         *int                  `json:"gracePeriod,omitempty"`
	HealthCheckType     *string               `json:"healthCheckType,omitempty"`
	Strategy            *AWSGroupRollStrategy `json:"strategy,omitempty"`
}

type RollAWSGroupOutput struct{}

func awsGroupFromJSON(in []byte) (*AWSGroup, error) {
	b := new(AWSGroup)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func awsGroupsFromJSON(in []byte) ([]*AWSGroup, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*AWSGroup, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := awsGroupFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func awsGroupsFromHttpResponse(resp *http.Response) ([]*AWSGroup, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return awsGroupsFromJSON(body)
}

func awsInstanceFromJSON(in []byte) (*AWSInstance, error) {
	b := new(AWSInstance)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func awsInstancesFromJSON(in []byte) ([]*AWSInstance, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*AWSInstance, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := awsInstanceFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func awsInstancesFromHttpResponse(resp *http.Response) ([]*AWSInstance, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return awsInstancesFromJSON(body)
}

func (s *AWSGroupServiceOp) List(ctx context.Context, input *ListAWSGroupInput) (*ListAWSGroupOutput, error) {
	r := s.client.newRequest(ctx, "GET", "/aws/ec2/group")
	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := awsGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListAWSGroupOutput{Groups: gs}, nil
}

func (s *AWSGroupServiceOp) Create(ctx context.Context, input *CreateAWSGroupInput) (*CreateAWSGroupOutput, error) {
	r := s.client.newRequest(ctx, "POST", "/aws/ec2/group")
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := awsGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateAWSGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *AWSGroupServiceOp) Read(ctx context.Context, input *ReadAWSGroupInput) (*ReadAWSGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}", map[string]string{
		"groupId": StringValue(input.GroupID),
	})
	if err != nil {
		return nil, err
	}

	r := s.client.newRequest(ctx, "GET", path)
	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := awsGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadAWSGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *AWSGroupServiceOp) Update(ctx context.Context, input *UpdateAWSGroupInput) (*UpdateAWSGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}", map[string]string{
		"groupId": StringValue(input.Group.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.Group.ID = nil

	r := s.client.newRequest(ctx, "PUT", path)
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := awsGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateAWSGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *AWSGroupServiceOp) Delete(ctx context.Context, input *DeleteAWSGroupInput) (*DeleteAWSGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}", map[string]string{
		"groupId": StringValue(input.GroupID),
	})
	if err != nil {
		return nil, err
	}

	r := s.client.newRequest(ctx, "DELETE", path)
	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteAWSGroupOutput{}, nil
}

func (s *AWSGroupServiceOp) Status(ctx context.Context, input *StatusAWSGroupInput) (*StatusAWSGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}/status", map[string]string{
		"groupId": StringValue(input.GroupID),
	})
	if err != nil {
		return nil, err
	}

	r := s.client.newRequest(ctx, "GET", path)
	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	is, err := awsInstancesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &StatusAWSGroupOutput{Instances: is}, nil
}

func (s *AWSGroupServiceOp) Detach(ctx context.Context, input *DetachAWSGroupInput) (*DetachAWSGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}/detachInstances", map[string]string{
		"groupId": StringValue(input.GroupID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.GroupID = nil

	r := s.client.newRequest(ctx, "PUT", path)
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DetachAWSGroupOutput{}, nil
}

func (s *AWSGroupServiceOp) Roll(ctx context.Context, input *RollAWSGroupInput) (*RollAWSGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}/roll", map[string]string{
		"groupId": StringValue(input.GroupID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.GroupID = nil

	r := s.client.newRequest(ctx, "PUT", path)
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &RollAWSGroupOutput{}, nil
}

// region AWSGroup

func (o *AWSGroup) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroup
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroup) SetId(v *string) *AWSGroup {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *AWSGroup) SetName(v *string) *AWSGroup {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AWSGroup) SetDescription(v *string) *AWSGroup {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *AWSGroup) SetCapacity(v *AWSGroupCapacity) *AWSGroup {
	if o.Capacity = v; o.Capacity == nil {
		o.nullFields = append(o.nullFields, "Capacity")
	}
	return o
}

func (o *AWSGroup) SetCompute(v *AWSGroupCompute) *AWSGroup {
	if o.Compute = v; o.Compute == nil {
		o.nullFields = append(o.nullFields, "Compute")
	}
	return o
}

func (o *AWSGroup) SetStrategy(v *AWSGroupStrategy) *AWSGroup {
	if o.Strategy = v; o.Strategy == nil {
		o.nullFields = append(o.nullFields, "Strategy")
	}
	return o
}

func (o *AWSGroup) SetScaling(v *AWSGroupScaling) *AWSGroup {
	if o.Scaling = v; o.Scaling == nil {
		o.nullFields = append(o.nullFields, "Scaling")
	}
	return o
}

func (o *AWSGroup) SetScheduling(v *AWSGroupScheduling) *AWSGroup {
	if o.Scheduling = v; o.Scheduling == nil {
		o.nullFields = append(o.nullFields, "Scheduling")
	}
	return o
}

func (o *AWSGroup) SetIntegration(v *AWSGroupIntegration) *AWSGroup {
	if o.Integration = v; o.Integration == nil {
		o.nullFields = append(o.nullFields, "Integration")
	}
	return o
}

// endregion

// region AWSGroupIntegration

func (o *AWSGroupIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupIntegration) SetEC2ContainerService(v *AWSGroupEC2ContainerServiceIntegration) *AWSGroupIntegration {
	if o.EC2ContainerService = v; o.EC2ContainerService == nil {
		o.nullFields = append(o.nullFields, "EC2ContainerService")
	}
	return o
}

func (o *AWSGroupIntegration) SetElasticBeanstalk(v *AWSGroupElasticBeanstalkIntegration) *AWSGroupIntegration {
	if o.ElasticBeanstalk = v; o.ElasticBeanstalk == nil {
		o.nullFields = append(o.nullFields, "ElasticBeanstalk")
	}
	return o
}

func (o *AWSGroupIntegration) SetRancher(v *AWSGroupRancherIntegration) *AWSGroupIntegration {
	if o.Rancher = v; o.Rancher == nil {
		o.nullFields = append(o.nullFields, "Rancher")
	}
	return o
}

func (o *AWSGroupIntegration) SetKubernetes(v *AWSGroupKubernetesIntegration) *AWSGroupIntegration {
	if o.Kubernetes = v; o.Kubernetes == nil {
		o.nullFields = append(o.nullFields, "Kubernetes")
	}
	return o
}

func (o *AWSGroupIntegration) SetMesosphere(v *AWSGroupMesosphereIntegration) *AWSGroupIntegration {
	if o.Mesosphere = v; o.Mesosphere == nil {
		o.nullFields = append(o.nullFields, "Mesosphere")
	}
	return o
}

func (o *AWSGroupIntegration) SetCodeDeploy(v *AWSGroupCodeDeployIntegration) *AWSGroupIntegration {
	if o.CodeDeploy = v; o.CodeDeploy == nil {
		o.nullFields = append(o.nullFields, "CodeDeploy")
	}
	return o
}

// endregion

// region AWSGroupRancherIntegration

func (o *AWSGroupRancherIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupRancherIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupRancherIntegration) SetMasterHost(v *string) *AWSGroupRancherIntegration {
	if o.MasterHost = v; o.MasterHost == nil {
		o.nullFields = append(o.nullFields, "MasterHost")
	}
	return o
}

func (o *AWSGroupRancherIntegration) SetAccessKey(v *string) *AWSGroupRancherIntegration {
	if o.AccessKey = v; o.AccessKey == nil {
		o.nullFields = append(o.nullFields, "AccessKey")
	}
	return o
}

func (o *AWSGroupRancherIntegration) SetSecretKey(v *string) *AWSGroupRancherIntegration {
	if o.SecretKey = v; o.SecretKey == nil {
		o.nullFields = append(o.nullFields, "SecretKey")
	}
	return o
}

// endregion

// region AWSGroupElasticBeanstalkIntegration

func (o *AWSGroupElasticBeanstalkIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupElasticBeanstalkIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupElasticBeanstalkIntegration) SetEnvironmentId(v *string) *AWSGroupElasticBeanstalkIntegration {
	if o.EnvironmentID = v; o.EnvironmentID == nil {
		o.nullFields = append(o.nullFields, "EnvironmentID")
	}
	return o
}

// endregion

// region AWSGroupEC2ContainerServiceIntegration

func (o *AWSGroupEC2ContainerServiceIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupEC2ContainerServiceIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupEC2ContainerServiceIntegration) SetClusterName(v *string) *AWSGroupEC2ContainerServiceIntegration {
	if o.ClusterName = v; o.ClusterName == nil {
		o.nullFields = append(o.nullFields, "ClusterName")
	}
	return o
}

func (o *AWSGroupEC2ContainerServiceIntegration) SetAutoScale(v *AWSGroupEC2ContainerServiceIntegrationAutoScale) *AWSGroupEC2ContainerServiceIntegration {
	if o.AutoScale = v; o.AutoScale == nil {
		o.nullFields = append(o.nullFields, "AutoScale")
	}
	return o
}

// endregion

// region AWSGroupEC2ContainerServiceIntegrationAutoScale

func (o *AWSGroupEC2ContainerServiceIntegrationAutoScale) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupEC2ContainerServiceIntegrationAutoScale
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupEC2ContainerServiceIntegrationAutoScale) SetIsEnabled(v *bool) *AWSGroupEC2ContainerServiceIntegrationAutoScale {
	if o.IsEnabled = v; o.IsEnabled == nil {
		o.nullFields = append(o.nullFields, "IsEnabled")
	}
	return o
}

func (o *AWSGroupEC2ContainerServiceIntegrationAutoScale) SetCooldown(v *int) *AWSGroupEC2ContainerServiceIntegrationAutoScale {
	if o.Cooldown = v; o.Cooldown == nil {
		o.nullFields = append(o.nullFields, "Cooldown")
	}
	return o
}

func (o *AWSGroupEC2ContainerServiceIntegrationAutoScale) SetHeadroom(v *AWSGroupEC2ContainerServiceIntegrationAutoScaleHeadroom) *AWSGroupEC2ContainerServiceIntegrationAutoScale {
	if o.Headroom = v; o.Headroom == nil {
		o.nullFields = append(o.nullFields, "Headroom")
	}
	return o
}

// endregion

// region AWSGroupEC2ContainerServiceIntegrationAutoScaleHeadroom

func (o *AWSGroupEC2ContainerServiceIntegrationAutoScaleHeadroom) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupEC2ContainerServiceIntegrationAutoScaleHeadroom
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupEC2ContainerServiceIntegrationAutoScaleHeadroom) SetCPUPerUnit(v *int) *AWSGroupEC2ContainerServiceIntegrationAutoScaleHeadroom {
	if o.CPUPerUnit = v; o.CPUPerUnit == nil {
		o.nullFields = append(o.nullFields, "CPUPerUnit")
	}
	return o
}

func (o *AWSGroupEC2ContainerServiceIntegrationAutoScaleHeadroom) SetMemoryPerUnit(v *int) *AWSGroupEC2ContainerServiceIntegrationAutoScaleHeadroom {
	if o.MemoryPerUnit = v; o.MemoryPerUnit == nil {
		o.nullFields = append(o.nullFields, "MemoryPerUnit")
	}
	return o
}

func (o *AWSGroupEC2ContainerServiceIntegrationAutoScaleHeadroom) SetNumOfUnits(v *int) *AWSGroupEC2ContainerServiceIntegrationAutoScaleHeadroom {
	if o.NumOfUnits = v; o.NumOfUnits == nil {
		o.nullFields = append(o.nullFields, "NumOfUnits")
	}
	return o
}

// endregion

// region AWSGroupKubernetesIntegration

func (o *AWSGroupKubernetesIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupKubernetesIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupKubernetesIntegration) SetServer(v *string) *AWSGroupKubernetesIntegration {
	if o.Server = v; o.Server == nil {
		o.nullFields = append(o.nullFields, "Server")
	}
	return o
}

func (o *AWSGroupKubernetesIntegration) SetToken(v *string) *AWSGroupKubernetesIntegration {
	if o.Token = v; o.Token == nil {
		o.nullFields = append(o.nullFields, "Token")
	}
	return o
}

// endregion

// region AWSGroupMesosphereIntegration

func (o *AWSGroupMesosphereIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupMesosphereIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupMesosphereIntegration) SetServer(v *string) *AWSGroupMesosphereIntegration {
	if o.Server = v; o.Server == nil {
		o.nullFields = append(o.nullFields, "Server")
	}
	return o
}

// endregion

// region AWSGroupScheduling

func (o *AWSGroupScheduling) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupScheduling
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupScheduling) SetTasks(v []*AWSGroupScheduledTask) *AWSGroupScheduling {
	if o.Tasks = v; o.Tasks == nil {
		o.nullFields = append(o.nullFields, "Tasks")
	}
	return o
}

// endregion

// region AWSGroupScheduledTask

func (o *AWSGroupScheduledTask) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupScheduledTask
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupScheduledTask) SetIsEnabled(v *bool) *AWSGroupScheduledTask {
	if o.IsEnabled = v; o.IsEnabled == nil {
		o.nullFields = append(o.nullFields, "IsEnabled")
	}
	return o
}

func (o *AWSGroupScheduledTask) SetFrequency(v *string) *AWSGroupScheduledTask {
	if o.Frequency = v; o.Frequency == nil {
		o.nullFields = append(o.nullFields, "Frequency")
	}
	return o
}

func (o *AWSGroupScheduledTask) SetCronExpression(v *string) *AWSGroupScheduledTask {
	if o.CronExpression = v; o.CronExpression == nil {
		o.nullFields = append(o.nullFields, "CronExpression")
	}
	return o
}

func (o *AWSGroupScheduledTask) SetTaskType(v *string) *AWSGroupScheduledTask {
	if o.TaskType = v; o.TaskType == nil {
		o.nullFields = append(o.nullFields, "TaskType")
	}
	return o
}

func (o *AWSGroupScheduledTask) SetScaleTargetCapacity(v *int) *AWSGroupScheduledTask {
	if o.ScaleTargetCapacity = v; o.ScaleTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "ScaleTargetCapacity")
	}
	return o
}

func (o *AWSGroupScheduledTask) SetScaleMinCapacity(v *int) *AWSGroupScheduledTask {
	if o.ScaleMinCapacity = v; o.ScaleMinCapacity == nil {
		o.nullFields = append(o.nullFields, "ScaleMinCapacity")
	}
	return o
}

func (o *AWSGroupScheduledTask) SetScaleMaxCapacity(v *int) *AWSGroupScheduledTask {
	if o.ScaleMaxCapacity = v; o.ScaleMaxCapacity == nil {
		o.nullFields = append(o.nullFields, "ScaleMaxCapacity")
	}
	return o
}

func (o *AWSGroupScheduledTask) SetBatchSizePercentage(v *int) *AWSGroupScheduledTask {
	if o.BatchSizePercentage = v; o.BatchSizePercentage == nil {
		o.nullFields = append(o.nullFields, "BatchSizePercentage")
	}
	return o
}

func (o *AWSGroupScheduledTask) SetGracePeriod(v *int) *AWSGroupScheduledTask {
	if o.GracePeriod = v; o.GracePeriod == nil {
		o.nullFields = append(o.nullFields, "GracePeriod")
	}
	return o
}

// endregion

// region AWSGroupScaling

func (o *AWSGroupScaling) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupScaling
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupScaling) SetUp(v []*AWSGroupScalingPolicy) *AWSGroupScaling {
	if o.Up = v; o.Up == nil {
		o.nullFields = append(o.nullFields, "Up")
	}
	return o
}

func (o *AWSGroupScaling) SetDown(v []*AWSGroupScalingPolicy) *AWSGroupScaling {
	if o.Down = v; o.Down == nil {
		o.nullFields = append(o.nullFields, "Down")
	}
	return o
}

// endregion

// region AWSGroupScalingPolicy

func (o *AWSGroupScalingPolicy) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupScalingPolicy
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupScalingPolicy) SetPolicyName(v *string) *AWSGroupScalingPolicy {
	if o.PolicyName = v; o.PolicyName == nil {
		o.nullFields = append(o.nullFields, "PolicyName")
	}
	return o
}

func (o *AWSGroupScalingPolicy) SetMetricName(v *string) *AWSGroupScalingPolicy {
	if o.MetricName = v; o.MetricName == nil {
		o.nullFields = append(o.nullFields, "MetricName")
	}
	return o
}

func (o *AWSGroupScalingPolicy) SetStatistic(v *string) *AWSGroupScalingPolicy {
	if o.Statistic = v; o.Statistic == nil {
		o.nullFields = append(o.nullFields, "Statistic")
	}
	return o
}

func (o *AWSGroupScalingPolicy) SetUnit(v *string) *AWSGroupScalingPolicy {
	if o.Unit = v; o.Unit == nil {
		o.nullFields = append(o.nullFields, "Unit")
	}
	return o
}

func (o *AWSGroupScalingPolicy) SetThreshold(v *float64) *AWSGroupScalingPolicy {
	if o.Threshold = v; o.Threshold == nil {
		o.nullFields = append(o.nullFields, "Threshold")
	}
	return o
}

func (o *AWSGroupScalingPolicy) SetAdjustment(v *int) *AWSGroupScalingPolicy {
	if o.Adjustment = v; o.Adjustment == nil {
		o.nullFields = append(o.nullFields, "Adjustment")
	}
	return o
}

func (o *AWSGroupScalingPolicy) SetMinTargetCapacity(v *int) *AWSGroupScalingPolicy {
	if o.MinTargetCapacity = v; o.MinTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "MinTargetCapacity")
	}
	return o
}

func (o *AWSGroupScalingPolicy) SetMaxTargetCapacity(v *int) *AWSGroupScalingPolicy {
	if o.MaxTargetCapacity = v; o.MaxTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "MaxTargetCapacity")
	}
	return o
}

func (o *AWSGroupScalingPolicy) SetNamespace(v *string) *AWSGroupScalingPolicy {
	if o.Namespace = v; o.Namespace == nil {
		o.nullFields = append(o.nullFields, "Namespace")
	}
	return o
}

func (o *AWSGroupScalingPolicy) SetEvaluationPeriods(v *int) *AWSGroupScalingPolicy {
	if o.EvaluationPeriods = v; o.EvaluationPeriods == nil {
		o.nullFields = append(o.nullFields, "EvaluationPeriods")
	}
	return o
}

func (o *AWSGroupScalingPolicy) SetPeriod(v *int) *AWSGroupScalingPolicy {
	if o.Period = v; o.Period == nil {
		o.nullFields = append(o.nullFields, "Period")
	}
	return o
}

func (o *AWSGroupScalingPolicy) SetCooldown(v *int) *AWSGroupScalingPolicy {
	if o.Cooldown = v; o.Cooldown == nil {
		o.nullFields = append(o.nullFields, "Cooldown")
	}
	return o
}

func (o *AWSGroupScalingPolicy) SetOperator(v *string) *AWSGroupScalingPolicy {
	if o.Operator = v; o.Operator == nil {
		o.nullFields = append(o.nullFields, "Operator")
	}
	return o
}

func (o *AWSGroupScalingPolicy) SetDimensions(v []*AWSGroupScalingPolicyDimension) *AWSGroupScalingPolicy {
	if o.Dimensions = v; o.Dimensions == nil {
		o.nullFields = append(o.nullFields, "Dimensions")
	}
	return o
}

func (o *AWSGroupScalingPolicy) SetAction(v *AWSGroupScalingPolicyAction) *AWSGroupScalingPolicy {
	if o.Action = v; o.Action == nil {
		o.nullFields = append(o.nullFields, "Action")
	}
	return o
}

// endregion

// region AWSGroupScalingPolicyAction

func (o *AWSGroupScalingPolicyAction) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupScalingPolicyAction
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupScalingPolicyAction) SetType(v *string) *AWSGroupScalingPolicyAction {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *AWSGroupScalingPolicyAction) SetAdjustment(v *string) *AWSGroupScalingPolicyAction {
	if o.Adjustment = v; o.Adjustment == nil {
		o.nullFields = append(o.nullFields, "Adjustment")
	}
	return o
}

func (o *AWSGroupScalingPolicyAction) SetMinTargetCapacity(v *string) *AWSGroupScalingPolicyAction {
	if o.MinTargetCapacity = v; o.MinTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "MinTargetCapacity")
	}
	return o
}

func (o *AWSGroupScalingPolicyAction) SetMaxTargetCapacity(v *string) *AWSGroupScalingPolicyAction {
	if o.MaxTargetCapacity = v; o.MaxTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "MaxTargetCapacity")
	}
	return o
}

func (o *AWSGroupScalingPolicyAction) SetMaximum(v *string) *AWSGroupScalingPolicyAction {
	if o.Maximum = v; o.Maximum == nil {
		o.nullFields = append(o.nullFields, "Maximum")
	}
	return o
}

func (o *AWSGroupScalingPolicyAction) SetMinimum(v *string) *AWSGroupScalingPolicyAction {
	if o.Minimum = v; o.Minimum == nil {
		o.nullFields = append(o.nullFields, "Minimum")
	}
	return o
}

func (o *AWSGroupScalingPolicyAction) SetTarget(v *string) *AWSGroupScalingPolicyAction {
	if o.Target = v; o.Target == nil {
		o.nullFields = append(o.nullFields, "Target")
	}
	return o
}

// endregion

// region AWSGroupScalingPolicyDimension

func (o *AWSGroupScalingPolicyDimension) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupScalingPolicyDimension
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupScalingPolicyDimension) SetName(v *string) *AWSGroupScalingPolicyDimension {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AWSGroupScalingPolicyDimension) SetValue(v *string) *AWSGroupScalingPolicyDimension {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

// endregion

// region AWSGroupStrategy

func (o *AWSGroupStrategy) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupStrategy
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupStrategy) SetRisk(v *float64) *AWSGroupStrategy {
	if o.Risk = v; o.Risk == nil {
		o.nullFields = append(o.nullFields, "Risk")
	}
	return o
}

func (o *AWSGroupStrategy) SetOnDemandCount(v *int) *AWSGroupStrategy {
	if o.OnDemandCount = v; o.OnDemandCount == nil {
		o.nullFields = append(o.nullFields, "OnDemandCount")
	}
	return o
}

func (o *AWSGroupStrategy) SetDrainingTimeout(v *int) *AWSGroupStrategy {
	if o.DrainingTimeout = v; o.DrainingTimeout == nil {
		o.nullFields = append(o.nullFields, "DrainingTimeout")
	}
	return o
}

func (o *AWSGroupStrategy) SetAvailabilityVsCost(v *string) *AWSGroupStrategy {
	if o.AvailabilityVsCost = v; o.AvailabilityVsCost == nil {
		o.nullFields = append(o.nullFields, "AvailabilityVsCost")
	}
	return o
}

func (o *AWSGroupStrategy) SetUtilizeReservedInstances(v *bool) *AWSGroupStrategy {
	if o.UtilizeReservedInstances = v; o.UtilizeReservedInstances == nil {
		o.nullFields = append(o.nullFields, "UtilizeReservedInstances")
	}
	return o
}

func (o *AWSGroupStrategy) SetFallbackToOnDemand(v *bool) *AWSGroupStrategy {
	if o.FallbackToOnDemand = v; o.FallbackToOnDemand == nil {
		o.nullFields = append(o.nullFields, "FallbackToOnDemand")
	}
	return o
}

func (o *AWSGroupStrategy) SetSpinUpTime(v *int) *AWSGroupStrategy {
	if o.SpinUpTime = v; o.SpinUpTime == nil {
		o.nullFields = append(o.nullFields, "SpinUpTime")
	}
	return o
}

func (o *AWSGroupStrategy) SetSignals(v []*AWSGroupStrategySignal) *AWSGroupStrategy {
	if o.Signals = v; o.Signals == nil {
		o.nullFields = append(o.nullFields, "Signals")
	}
	return o
}

func (o *AWSGroupStrategy) SetPersistence(v *AWSGroupPersistence) *AWSGroupStrategy {
	if o.Persistence = v; o.Persistence == nil {
		o.nullFields = append(o.nullFields, "Persistence")
	}
	return o
}

// endregion

// region AWSGroupPersistence

func (o *AWSGroupPersistence) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupPersistence
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupPersistence) SetShouldPersistPrivateIp(v *bool) *AWSGroupPersistence {
	if o.ShouldPersistPrivateIp = v; o.ShouldPersistPrivateIp == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistPrivateIp")
	}
	return o
}

func (o *AWSGroupPersistence) SetShouldPersistBlockDevices(v *bool) *AWSGroupPersistence {
	if o.ShouldPersistBlockDevices = v; o.ShouldPersistBlockDevices == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistBlockDevices")
	}
	return o
}

func (o *AWSGroupPersistence) SetShouldPersistRootDevice(v *bool) *AWSGroupPersistence {
	if o.ShouldPersistRootDevice = v; o.ShouldPersistRootDevice == nil {
		o.nullFields = append(o.nullFields, "ShouldPersistRootDevice")
	}
	return o
}

// endregion

// region AWSGroupStrategySignal

func (o *AWSGroupStrategySignal) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupStrategySignal
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupStrategySignal) SetName(v *string) *AWSGroupStrategySignal {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AWSGroupStrategySignal) SetTimeout(v *int) *AWSGroupStrategySignal {
	if o.Timeout = v; o.Timeout == nil {
		o.nullFields = append(o.nullFields, "Timeout")
	}
	return o
}

// endregion

// region AWSGroupCapacity

func (o *AWSGroupCapacity) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupCapacity
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupCapacity) SetMinimum(v *int) *AWSGroupCapacity {
	if o.Minimum = v; o.Minimum == nil {
		o.nullFields = append(o.nullFields, "Minimum")
	}
	return o
}

func (o *AWSGroupCapacity) SetMaximum(v *int) *AWSGroupCapacity {
	if o.Maximum = v; o.Maximum == nil {
		o.nullFields = append(o.nullFields, "Maximum")
	}
	return o
}

func (o *AWSGroupCapacity) SetTarget(v *int) *AWSGroupCapacity {
	if o.Target = v; o.Target == nil {
		o.nullFields = append(o.nullFields, "Target")
	}
	return o
}

func (o *AWSGroupCapacity) SetUnit(v *string) *AWSGroupCapacity {
	if o.Unit = v; o.Unit == nil {
		o.nullFields = append(o.nullFields, "Unit")
	}
	return o
}

// endregion

// region AWSGroupCompute

func (o *AWSGroupCompute) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupCompute
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupCompute) SetProduct(v *string) *AWSGroupCompute {
	if o.Product = v; o.Product == nil {
		o.nullFields = append(o.nullFields, "Product")
	}
	return o
}

func (o *AWSGroupCompute) SetInstanceTypes(v *AWSGroupComputeInstanceType) *AWSGroupCompute {
	if o.InstanceTypes = v; o.InstanceTypes == nil {
		o.nullFields = append(o.nullFields, "InstanceTypes")
	}
	return o
}

func (o *AWSGroupCompute) SetLaunchSpecification(v *AWSGroupComputeLaunchSpecification) *AWSGroupCompute {
	if o.LaunchSpecification = v; o.LaunchSpecification == nil {
		o.nullFields = append(o.nullFields, "LaunchSpecification")
	}
	return o
}

func (o *AWSGroupCompute) SetAvailabilityZones(v []*AWSGroupComputeAvailabilityZone) *AWSGroupCompute {
	if o.AvailabilityZones = v; o.AvailabilityZones == nil {
		o.nullFields = append(o.nullFields, "AvailabilityZones")
	}
	return o
}

func (o *AWSGroupCompute) SetElasticIPs(v []string) *AWSGroupCompute {
	if o.ElasticIPs = v; o.ElasticIPs == nil {
		o.nullFields = append(o.nullFields, "ElasticIPs")
	}
	return o
}

func (o *AWSGroupCompute) SetEBSVolumePool(v []*AWSGroupComputeEBSVolume) *AWSGroupCompute {
	if o.EBSVolumePool = v; o.EBSVolumePool == nil {
		o.nullFields = append(o.nullFields, "EBSVolumePool")
	}
	return o
}

// endregion

// region AWSGroupComputeEBSVolume

func (o *AWSGroupComputeEBSVolume) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupComputeEBSVolume
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupComputeEBSVolume) SetDeviceName(v *string) *AWSGroupComputeEBSVolume {
	if o.DeviceName = v; o.DeviceName == nil {
		o.nullFields = append(o.nullFields, "DeviceName")
	}
	return o
}

func (o *AWSGroupComputeEBSVolume) SetVolumeIDs(v []string) *AWSGroupComputeEBSVolume {
	if o.VolumeIDs = v; o.VolumeIDs == nil {
		o.nullFields = append(o.nullFields, "VolumeIDs")
	}
	return o
}

// endregion

// region AWSGroupComputeInstanceType

func (o *AWSGroupComputeInstanceType) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupComputeInstanceType
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupComputeInstanceType) SetOnDemand(v *string) *AWSGroupComputeInstanceType {
	if o.OnDemand = v; o.OnDemand == nil {
		o.nullFields = append(o.nullFields, "OnDemand")
	}
	return o
}

func (o *AWSGroupComputeInstanceType) SetSpot(v []string) *AWSGroupComputeInstanceType {
	if o.Spot = v; o.Spot == nil {
		o.nullFields = append(o.nullFields, "Spot")
	}
	return o
}

func (o *AWSGroupComputeInstanceType) SetWeights(v []*AWSGroupComputeInstanceTypeWeight) *AWSGroupComputeInstanceType {
	if o.Weights = v; o.Weights == nil {
		o.nullFields = append(o.nullFields, "Weights")
	}
	return o
}

// endregion

// region AWSGroupComputeInstanceTypeWeight

func (o *AWSGroupComputeInstanceTypeWeight) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupComputeInstanceTypeWeight
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupComputeInstanceTypeWeight) SetInstanceType(v *string) *AWSGroupComputeInstanceTypeWeight {
	if o.InstanceType = v; o.InstanceType == nil {
		o.nullFields = append(o.nullFields, "InstanceType")
	}
	return o
}

func (o *AWSGroupComputeInstanceTypeWeight) SetWeight(v *int) *AWSGroupComputeInstanceTypeWeight {
	if o.Weight = v; o.Weight == nil {
		o.nullFields = append(o.nullFields, "Weight")
	}
	return o
}

// endregion

// region AWSGroupComputeAvailabilityZone

func (o *AWSGroupComputeAvailabilityZone) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupComputeAvailabilityZone
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupComputeAvailabilityZone) SetName(v *string) *AWSGroupComputeAvailabilityZone {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AWSGroupComputeAvailabilityZone) SetSubnetId(v *string) *AWSGroupComputeAvailabilityZone {
	if o.SubnetID = v; o.SubnetID == nil {
		o.nullFields = append(o.nullFields, "SubnetID")
	}
	return o
}

// endregion

// region AWSGroupComputeLaunchSpecification

func (o *AWSGroupComputeLaunchSpecification) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupComputeLaunchSpecification
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupComputeLaunchSpecification) SetLoadBalancerNames(v []string) *AWSGroupComputeLaunchSpecification {
	if o.LoadBalancerNames = v; o.LoadBalancerNames == nil {
		o.nullFields = append(o.nullFields, "LoadBalancerNames")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetLoadBalancersConfig(v *AWSGroupComputeLoadBalancersConfig) *AWSGroupComputeLaunchSpecification {
	if o.LoadBalancersConfig = v; o.LoadBalancersConfig == nil {
		o.nullFields = append(o.nullFields, "LoadBalancersConfig")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetSecurityGroupIDs(v []string) *AWSGroupComputeLaunchSpecification {
	if o.SecurityGroupIDs = v; o.SecurityGroupIDs == nil {
		o.nullFields = append(o.nullFields, "SecurityGroupIDs")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetHealthCheckType(v *string) *AWSGroupComputeLaunchSpecification {
	if o.HealthCheckType = v; o.HealthCheckType == nil {
		o.nullFields = append(o.nullFields, "HealthCheckType")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetHealthCheckGracePeriod(v *int) *AWSGroupComputeLaunchSpecification {
	if o.HealthCheckGracePeriod = v; o.HealthCheckGracePeriod == nil {
		o.nullFields = append(o.nullFields, "HealthCheckGracePeriod")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetHealthCheckUnhealthyDurationBeforeReplacement(v *int) *AWSGroupComputeLaunchSpecification {
	if o.HealthCheckUnhealthyDurationBeforeReplacement = v; o.HealthCheckUnhealthyDurationBeforeReplacement == nil {
		o.nullFields = append(o.nullFields, "HealthCheckUnhealthyDurationBeforeReplacement")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetImageId(v *string) *AWSGroupComputeLaunchSpecification {
	if o.ImageID = v; o.ImageID == nil {
		o.nullFields = append(o.nullFields, "ImageID")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetKeyPair(v *string) *AWSGroupComputeLaunchSpecification {
	if o.KeyPair = v; o.KeyPair == nil {
		o.nullFields = append(o.nullFields, "KeyPair")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetUserData(v *string) *AWSGroupComputeLaunchSpecification {
	if o.UserData = v; o.UserData == nil {
		o.nullFields = append(o.nullFields, "UserData")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetShutdownScript(v *string) *AWSGroupComputeLaunchSpecification {
	if o.ShutdownScript = v; o.ShutdownScript == nil {
		o.nullFields = append(o.nullFields, "ShutdownScript")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetTenancy(v *string) *AWSGroupComputeLaunchSpecification {
	if o.Tenancy = v; o.Tenancy == nil {
		o.nullFields = append(o.nullFields, "Tenancy")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetMonitoring(v *bool) *AWSGroupComputeLaunchSpecification {
	if o.Monitoring = v; o.Monitoring == nil {
		o.nullFields = append(o.nullFields, "Monitoring")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetEBSOptimized(v *bool) *AWSGroupComputeLaunchSpecification {
	if o.EBSOptimized = v; o.EBSOptimized == nil {
		o.nullFields = append(o.nullFields, "EBSOptimized")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetIAMInstanceProfile(v *AWSGroupComputeIAMInstanceProfile) *AWSGroupComputeLaunchSpecification {
	if o.IAMInstanceProfile = v; o.IAMInstanceProfile == nil {
		o.nullFields = append(o.nullFields, "IAMInstanceProfile")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetBlockDevices(v []*AWSGroupComputeBlockDevice) *AWSGroupComputeLaunchSpecification {
	if o.BlockDevices = v; o.BlockDevices == nil {
		o.nullFields = append(o.nullFields, "BlockDevices")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetNetworkInterfaces(v []*AWSGroupComputeNetworkInterface) *AWSGroupComputeLaunchSpecification {
	if o.NetworkInterfaces = v; o.NetworkInterfaces == nil {
		o.nullFields = append(o.nullFields, "NetworkInterfaces")
	}
	return o
}

func (o *AWSGroupComputeLaunchSpecification) SetTags(v []*AWSGroupComputeTag) *AWSGroupComputeLaunchSpecification {
	if o.Tags = v; o.Tags == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

// endregion

// region AWSGroupComputeLoadBalancersConfig

func (o *AWSGroupComputeLoadBalancersConfig) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupComputeLoadBalancersConfig
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupComputeLoadBalancersConfig) SetLoadBalancers(v []*AWSGroupComputeLoadBalancer) *AWSGroupComputeLoadBalancersConfig {
	if o.LoadBalancers = v; o.LoadBalancers == nil {
		o.nullFields = append(o.nullFields, "LoadBalancers")
	}
	return o
}

// endregion

// region AWSGroupComputeLoadBalancer

func (o *AWSGroupComputeLoadBalancer) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupComputeLoadBalancer
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupComputeLoadBalancer) SetName(v *string) *AWSGroupComputeLoadBalancer {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AWSGroupComputeLoadBalancer) SetArn(v *string) *AWSGroupComputeLoadBalancer {
	if o.Arn = v; o.Arn == nil {
		o.nullFields = append(o.nullFields, "Arn")
	}
	return o
}

func (o *AWSGroupComputeLoadBalancer) SetType(v *string) *AWSGroupComputeLoadBalancer {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

// endregion

// region AWSGroupComputeNetworkInterface

func (o *AWSGroupComputeNetworkInterface) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupComputeNetworkInterface
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupComputeNetworkInterface) SetId(v *string) *AWSGroupComputeNetworkInterface {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *AWSGroupComputeNetworkInterface) SetDescription(v *string) *AWSGroupComputeNetworkInterface {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *AWSGroupComputeNetworkInterface) SetDeviceIndex(v *int) *AWSGroupComputeNetworkInterface {
	if o.DeviceIndex = v; o.DeviceIndex == nil {
		o.nullFields = append(o.nullFields, "DeviceIndex")
	}
	return o
}

func (o *AWSGroupComputeNetworkInterface) SetSecondaryPrivateIPAddressCount(v *int) *AWSGroupComputeNetworkInterface {
	if o.SecondaryPrivateIPAddressCount = v; o.SecondaryPrivateIPAddressCount == nil {
		o.nullFields = append(o.nullFields, "SecondaryPrivateIPAddressCount")
	}
	return o
}

func (o *AWSGroupComputeNetworkInterface) SetAssociatePublicIPAddress(v *bool) *AWSGroupComputeNetworkInterface {
	if o.AssociatePublicIPAddress = v; o.AssociatePublicIPAddress == nil {
		o.nullFields = append(o.nullFields, "AssociatePublicIPAddress")
	}
	return o
}

func (o *AWSGroupComputeNetworkInterface) SetDeleteOnTermination(v *bool) *AWSGroupComputeNetworkInterface {
	if o.DeleteOnTermination = v; o.DeleteOnTermination == nil {
		o.nullFields = append(o.nullFields, "DeleteOnTermination")
	}
	return o
}

func (o *AWSGroupComputeNetworkInterface) SetSecurityGroupsIDs(v []string) *AWSGroupComputeNetworkInterface {
	if o.SecurityGroupsIDs = v; o.SecurityGroupsIDs == nil {
		o.nullFields = append(o.nullFields, "SecurityGroupsIDs")
	}
	return o
}

func (o *AWSGroupComputeNetworkInterface) SetPrivateIPAddress(v *string) *AWSGroupComputeNetworkInterface {
	if o.PrivateIPAddress = v; o.PrivateIPAddress == nil {
		o.nullFields = append(o.nullFields, "PrivateIPAddress")
	}
	return o
}

func (o *AWSGroupComputeNetworkInterface) SetSubnetId(v *string) *AWSGroupComputeNetworkInterface {
	if o.SubnetID = v; o.SubnetID == nil {
		o.nullFields = append(o.nullFields, "SubnetID")
	}
	return o
}

// endregion

// region AWSGroupComputeBlockDevice

func (o *AWSGroupComputeBlockDevice) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupComputeBlockDevice
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupComputeBlockDevice) SetDeviceName(v *string) *AWSGroupComputeBlockDevice {
	if o.DeviceName = v; o.DeviceName == nil {
		o.nullFields = append(o.nullFields, "DeviceName")
	}
	return o
}

func (o *AWSGroupComputeBlockDevice) SetVirtualName(v *string) *AWSGroupComputeBlockDevice {
	if o.VirtualName = v; o.VirtualName == nil {
		o.nullFields = append(o.nullFields, "VirtualName")
	}
	return o
}

func (o *AWSGroupComputeBlockDevice) SetEBS(v *AWSGroupComputeEBS) *AWSGroupComputeBlockDevice {
	if o.EBS = v; o.EBS == nil {
		o.nullFields = append(o.nullFields, "EBS")
	}
	return o
}

// endregion

// region AWSGroupComputeEBS

func (o *AWSGroupComputeEBS) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupComputeEBS
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupComputeEBS) SetDeleteOnTermination(v *bool) *AWSGroupComputeEBS {
	if o.DeleteOnTermination = v; o.DeleteOnTermination == nil {
		o.nullFields = append(o.nullFields, "DeleteOnTermination")
	}
	return o
}

func (o *AWSGroupComputeEBS) SetEncrypted(v *bool) *AWSGroupComputeEBS {
	if o.Encrypted = v; o.Encrypted == nil {
		o.nullFields = append(o.nullFields, "Encrypted")
	}
	return o
}

func (o *AWSGroupComputeEBS) SetSnapshotId(v *string) *AWSGroupComputeEBS {
	if o.SnapshotID = v; o.SnapshotID == nil {
		o.nullFields = append(o.nullFields, "SnapshotID")
	}
	return o
}

func (o *AWSGroupComputeEBS) SetVolumeType(v *string) *AWSGroupComputeEBS {
	if o.VolumeType = v; o.VolumeType == nil {
		o.nullFields = append(o.nullFields, "VolumeType")
	}
	return o
}

func (o *AWSGroupComputeEBS) SetVolumeSize(v *int) *AWSGroupComputeEBS {
	if o.VolumeSize = v; o.VolumeSize == nil {
		o.nullFields = append(o.nullFields, "VolumeSize")
	}
	return o
}

func (o *AWSGroupComputeEBS) SetIOPS(v *int) *AWSGroupComputeEBS {
	if o.IOPS = v; o.IOPS == nil {
		o.nullFields = append(o.nullFields, "IOPS")
	}
	return o
}

// endregion

// region AWSGroupComputeIAMInstanceProfile

func (o *AWSGroupComputeIAMInstanceProfile) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupComputeIAMInstanceProfile
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupComputeIAMInstanceProfile) SetName(v *string) *AWSGroupComputeIAMInstanceProfile {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AWSGroupComputeIAMInstanceProfile) SetArn(v *string) *AWSGroupComputeIAMInstanceProfile {
	if o.Arn = v; o.Arn == nil {
		o.nullFields = append(o.nullFields, "Arn")
	}
	return o
}

// endregion

// region AWSGroupComputeTag

func (o *AWSGroupComputeTag) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupComputeTag
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupComputeTag) SetKey(v *string) *AWSGroupComputeTag {
	if o.Key = v; o.Key == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

func (o *AWSGroupComputeTag) SetValue(v *string) *AWSGroupComputeTag {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

// endregion

// region AWSGroupRollStrategy

func (o *AWSGroupRollStrategy) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupRollStrategy
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupRollStrategy) SetAction(v *string) *AWSGroupRollStrategy {
	if o.Action = v; o.Action == nil {
		o.nullFields = append(o.nullFields, "Action")
	}
	return o
}

func (o *AWSGroupRollStrategy) SetShouldDrainInstances(v *bool) *AWSGroupRollStrategy {
	if o.ShouldDrainInstances = v; o.ShouldDrainInstances == nil {
		o.nullFields = append(o.nullFields, "ShouldDrainInstances")
	}
	return o
}

// endregion

// region AWSGroupCodeDeployIntegration

func (o *AWSGroupCodeDeployIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupCodeDeployIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupCodeDeployIntegration) SetDeploymentGroups(v []*AWSGroupCodeDeployIntegrationDeploymentGroup) *AWSGroupCodeDeployIntegration {
	if o.DeploymentGroups = v; o.DeploymentGroups == nil {
		o.nullFields = append(o.nullFields, "DeploymentGroups")
	}
	return o
}

func (o *AWSGroupCodeDeployIntegration) SetCleanUpOnFailure(v *bool) *AWSGroupCodeDeployIntegration {
	if o.CleanUpOnFailure = v; o.CleanUpOnFailure == nil {
		o.nullFields = append(o.nullFields, "CleanUpOnFailure")
	}
	return o
}

func (o *AWSGroupCodeDeployIntegration) SetTerminateInstanceOnFailure(v *bool) *AWSGroupCodeDeployIntegration {
	if o.TerminateInstanceOnFailure = v; o.TerminateInstanceOnFailure == nil {
		o.nullFields = append(o.nullFields, "TerminateInstanceOnFailure")
	}
	return o
}

// endregion

// region AWSGroupCodeDeployIntegrationDeploymentGroup

func (o *AWSGroupCodeDeployIntegrationDeploymentGroup) MarshalJSON() ([]byte, error) {
	type noMethod AWSGroupCodeDeployIntegrationDeploymentGroup
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AWSGroupCodeDeployIntegrationDeploymentGroup) SetApplicationName(v *string) *AWSGroupCodeDeployIntegrationDeploymentGroup {
	if o.ApplicationName = v; o.ApplicationName == nil {
		o.nullFields = append(o.nullFields, "ApplicationName")
	}
	return o
}

func (o *AWSGroupCodeDeployIntegrationDeploymentGroup) SetDeploymentGroupName(v *string) *AWSGroupCodeDeployIntegrationDeploymentGroup {
	if o.DeploymentGroupName = v; o.DeploymentGroupName == nil {
		o.nullFields = append(o.nullFields, "DeploymentGroupName")
	}
	return o
}

// endregion
