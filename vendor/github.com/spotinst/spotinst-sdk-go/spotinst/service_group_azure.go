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

// AzureGroupService is an interface for interfacing with the AzureGroup
// endpoints of the Spotinst API.
type AzureGroupService interface {
	List(context.Context, *ListAzureGroupInput) (*ListAzureGroupOutput, error)
	Create(context.Context, *CreateAzureGroupInput) (*CreateAzureGroupOutput, error)
	Read(context.Context, *ReadAzureGroupInput) (*ReadAzureGroupOutput, error)
	Update(context.Context, *UpdateAzureGroupInput) (*UpdateAzureGroupOutput, error)
	Delete(context.Context, *DeleteAzureGroupInput) (*DeleteAzureGroupOutput, error)
	Status(context.Context, *StatusAzureGroupInput) (*StatusAzureGroupOutput, error)
	Detach(context.Context, *DetachAzureGroupInput) (*DetachAzureGroupOutput, error)
	Roll(context.Context, *RollAzureGroupInput) (*RollAzureGroupOutput, error)
}

// AzureGroupServiceOp handles communication with the balancer related methods
// of the Spotinst API.
type AzureGroupServiceOp struct {
	client *Client
}

var _ AzureGroupService = &AzureGroupServiceOp{}

type AzureGroup struct {
	ID          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Capacity    *AzureGroupCapacity    `json:"capacity,omitempty"`
	Compute     *AzureGroupCompute     `json:"compute,omitempty"`
	Strategy    *AzureGroupStrategy    `json:"strategy,omitempty"`
	Scaling     *AzureGroupScaling     `json:"scaling,omitempty"`
	Scheduling  *AzureGroupScheduling  `json:"scheduling,omitempty"`
	Integration *AzureGroupIntegration `json:"thirdPartiesIntegration,omitempty"`

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

type AzureGroupScheduling struct {
	Tasks []*AzureGroupScheduledTask `json:"tasks,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupIntegration struct {
	Rancher *AzureGroupRancherIntegration `json:"rancher,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupRancherIntegration struct {
	MasterHost *string `json:"masterHost,omitempty"`
	AccessKey  *string `json:"accessKey,omitempty"`
	SecretKey  *string `json:"secretKey,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupScheduledTask struct {
	IsEnabled            *bool   `json:"isEnabled,omitempty"`
	Frequency            *string `json:"frequency,omitempty"`
	CronExpression       *string `json:"cronExpression,omitempty"`
	TaskType             *string `json:"taskType,omitempty"`
	ScaleTargetCapacity  *int    `json:"scaleTargetCapacity,omitempty"`
	ScaleMinCapacity     *int    `json:"scaleMinCapacity,omitempty"`
	ScaleMaxCapacity     *int    `json:"scaleMaxCapacity,omitempty"`
	BatchSizePercentage  *int    `json:"batchSizePercentage,omitempty"`
	GracePeriod          *int    `json:"gracePeriod,omitempty"`
	Adjustment           *int    `json:"adjustment,omitempty"`
	AdjustmentPercentage *int    `json:"adjustmentPercentage,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupScaling struct {
	Up   []*AzureGroupScalingPolicy `json:"up,omitempty"`
	Down []*AzureGroupScalingPolicy `json:"down,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupScalingPolicy struct {
	PolicyName        *string                             `json:"policyName,omitempty"`
	MetricName        *string                             `json:"metricName,omitempty"`
	Statistic         *string                             `json:"statistic,omitempty"`
	Unit              *string                             `json:"unit,omitempty"`
	Threshold         *float64                            `json:"threshold,omitempty"`
	Adjustment        *int                                `json:"adjustment,omitempty"`
	MinTargetCapacity *int                                `json:"minTargetCapacity,omitempty"`
	MaxTargetCapacity *int                                `json:"maxTargetCapacity,omitempty"`
	Namespace         *string                             `json:"namespace,omitempty"`
	EvaluationPeriods *int                                `json:"evaluationPeriods,omitempty"`
	Period            *int                                `json:"period,omitempty"`
	Cooldown          *int                                `json:"cooldown,omitempty"`
	Operator          *string                             `json:"operator,omitempty"`
	Dimensions        []*AzureGroupScalingPolicyDimension `json:"dimensions,omitempty"`
	Action            *AzureGroupScalingPolicyAction      `json:"action,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupScalingPolicyAction struct {
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

type AzureGroupScalingPolicyDimension struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupStrategy struct {
	LowPriorityPercentage *int                        `json:"lowPriorityPercentage,omitempty"`
	DedicatedCount        *int                        `json:"dedicatedCount,omitempty"`
	DrainingTimeout       *int                        `json:"drainingTimeout,omitempty"`
	Signals               []*AzureGroupStrategySignal `json:"signals,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupStrategySignal struct {
	Name    *string `json:"name,omitempty"`
	Timeout *int    `json:"timeout,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupCapacity struct {
	Minimum *int `json:"minimum,omitempty"`
	Maximum *int `json:"maximum,omitempty"`
	Target  *int `json:"target,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupCompute struct {
	Region              *string                               `json:"region,omitempty"`
	Product             *string                               `json:"product,omitempty"`
	ResourceGroupName   *string                               `json:"resourceGroupName,omitempty"`
	VMSize              *AzureGroupComputeVMSize              `json:"vmSizes,omitempty"`
	LaunchSpecification *AzureGroupComputeLaunchSpecification `json:"launchSpecification,omitempty"`
	Health              *AzureGroupComputeHealth              `json:"health,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupComputeVMSize struct {
	Dedicated   []string `json:"dedicatedSizes,omitempty"`
	LowPriority []string `json:"lowPrioritySizes,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupComputeLaunchSpecification struct {
	LoadBalancersConfig *AzureGroupComputeLoadBalancersConfig `json:"loadBalancersConfig,omitempty"`
	Image               *AzureGroupComputeImage               `json:"image,omitempty"`
	UserData            *AzureGroupComputeUserData            `json:"userData,omitempty"`
	Storage             *AzureGroupComputeStorage             `json:"storage,omitempty"`
	Network             *AzureGroupComputeNetwork             `json:"network,omitempty"`
	SSHPublicKey        *string                               `json:"sshPublicKey,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupComputeLoadBalancersConfig struct {
	LoadBalancers []*AzureGroupComputeLoadBalancer `json:"loadBalancers,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupComputeLoadBalancer struct {
	BalancerID  *string `json:"balancerId,omitempty"`
	TargetSetID *string `json:"targetSetId,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupComputeImage struct {
	Custom    *AzureGroupComputeImageCustom `json:"customImage,omitempty"`
	Publisher *string                       `json:"publisher,omitempty"`
	Offer     *string                       `json:"offer,omitempty"`
	SKU       *string                       `json:"sku,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupComputeImageCustom struct {
	ImageURIs []string `json:"imageUris,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupComputeUserData struct {
	CommandLine   *string                                  `json:"commandLine,omitempty"`
	ResourceFiles []*AzureGroupComputeUserDataResourceFile `json:"resourceFiles,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupComputeUserDataResourceFile struct {
	URL        *string `json:"resourceFileUrl,omitempty"`
	TargetPath *string `json:"resourceFileTargetPath,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupComputeStorage struct {
	AccountName *string `json:"storageAccountName,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupComputeNetwork struct {
	VirtualNetworkName *string `json:"virtualNetworkName,omitempty"`
	SubnetID           *string `json:"subnetId,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureGroupComputeHealth struct {
	HealthCheckType *string `json:"healthCheckType,omitempty"`
	AutoHealing     *bool   `json:"autoHealing,omitempty"`
	GracePeriod     *int    `json:"gracePeriod,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AzureNode struct {
	ID        *string    `json:"id,omitempty"`
	VMSize    *string    `json:"vmSize,omitempty"`
	State     *string    `json:"state,omitempty"`
	LifeCycle *string    `json:"lifeCycle,omitempty"`
	Region    *string    `json:"region,omitempty"`
	IPAddress *string    `json:"ipAddress,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

type AzureGroupRollStrategy struct {
	Action               *string `json:"action,omitempty"`
	ShouldDrainInstances *bool   `json:"shouldDrainInstances,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListAzureGroupInput struct{}

type ListAzureGroupOutput struct {
	Groups []*AzureGroup `json:"groups,omitempty"`
}

type CreateAzureGroupInput struct {
	Group *AzureGroup `json:"group,omitempty"`
}

type CreateAzureGroupOutput struct {
	Group *AzureGroup `json:"group,omitempty"`
}

type ReadAzureGroupInput struct {
	GroupID *string `json:"groupId,omitempty"`
}

type ReadAzureGroupOutput struct {
	Group *AzureGroup `json:"group,omitempty"`
}

type UpdateAzureGroupInput struct {
	Group *AzureGroup `json:"group,omitempty"`
}

type UpdateAzureGroupOutput struct {
	Group *AzureGroup `json:"group,omitempty"`
}

type DeleteAzureGroupInput struct {
	GroupID *string `json:"groupId,omitempty"`
}

type DeleteAzureGroupOutput struct{}

type StatusAzureGroupInput struct {
	GroupID *string `json:"groupId,omitempty"`
}

type StatusAzureGroupOutput struct {
	Nodes []*AzureNode `json:"instances,omitempty"`
}

type DetachAzureGroupInput struct {
	GroupID                       *string  `json:"groupId,omitempty"`
	InstanceIDs                   []string `json:"instancesToDetach,omitempty"`
	ShouldDecrementTargetCapacity *bool    `json:"shouldDecrementTargetCapacity,omitempty"`
	ShouldTerminateInstances      *bool    `json:"shouldTerminateInstances,omitempty"`
	DrainingTimeout               *int     `json:"drainingTimeout,omitempty"`
}

type DetachAzureGroupOutput struct{}

type RollAzureGroupInput struct {
	GroupID             *string                 `json:"groupId,omitempty"`
	BatchSizePercentage *int                    `json:"batchSizePercentage,omitempty"`
	GracePeriod         *int                    `json:"gracePeriod,omitempty"`
	HealthCheckType     *string                 `json:"healthCheckType,omitempty"`
	Strategy            *AzureGroupRollStrategy `json:"strategy,omitempty"`
}

type RollAzureGroupOutput struct{}

func azureGroupFromJSON(in []byte) (*AzureGroup, error) {
	b := new(AzureGroup)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func azureGroupsFromJSON(in []byte) ([]*AzureGroup, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*AzureGroup, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := azureGroupFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func azureGroupsFromHttpResponse(resp *http.Response) ([]*AzureGroup, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return azureGroupsFromJSON(body)
}

func azureNodeFromJSON(in []byte) (*AzureNode, error) {
	b := new(AzureNode)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func azureNodesFromJSON(in []byte) ([]*AzureNode, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*AzureNode, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := azureNodeFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func azureNodesFromHttpResponse(resp *http.Response) ([]*AzureNode, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return azureNodesFromJSON(body)
}

func (s *AzureGroupServiceOp) List(ctx context.Context, input *ListAzureGroupInput) (*ListAzureGroupOutput, error) {
	r := s.client.newRequest(ctx, "GET", "/compute/azure/group")
	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := azureGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListAzureGroupOutput{Groups: gs}, nil
}

func (s *AzureGroupServiceOp) Create(ctx context.Context, input *CreateAzureGroupInput) (*CreateAzureGroupOutput, error) {
	r := s.client.newRequest(ctx, "POST", "/compute/azure/group")
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := azureGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateAzureGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *AzureGroupServiceOp) Read(ctx context.Context, input *ReadAzureGroupInput) (*ReadAzureGroupOutput, error) {
	path, err := uritemplates.Expand("/compute/azure/group/{groupId}", map[string]string{
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

	gs, err := azureGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadAzureGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *AzureGroupServiceOp) Update(ctx context.Context, input *UpdateAzureGroupInput) (*UpdateAzureGroupOutput, error) {
	path, err := uritemplates.Expand("/compute/azure/group/{groupId}", map[string]string{
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

	gs, err := azureGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateAzureGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *AzureGroupServiceOp) Delete(ctx context.Context, input *DeleteAzureGroupInput) (*DeleteAzureGroupOutput, error) {
	path, err := uritemplates.Expand("/compute/azure/group/{groupId}", map[string]string{
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

	return &DeleteAzureGroupOutput{}, nil
}

func (s *AzureGroupServiceOp) Status(ctx context.Context, input *StatusAzureGroupInput) (*StatusAzureGroupOutput, error) {
	path, err := uritemplates.Expand("/compute/azure/group/{groupId}/status", map[string]string{
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

	ns, err := azureNodesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &StatusAzureGroupOutput{Nodes: ns}, nil
}

func (s *AzureGroupServiceOp) Detach(ctx context.Context, input *DetachAzureGroupInput) (*DetachAzureGroupOutput, error) {
	path, err := uritemplates.Expand("/compute/azure/group/{groupId}/detachNodes", map[string]string{
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

	return &DetachAzureGroupOutput{}, nil
}

func (s *AzureGroupServiceOp) Roll(ctx context.Context, input *RollAzureGroupInput) (*RollAzureGroupOutput, error) {
	path, err := uritemplates.Expand("/compute/azure/group/{groupId}/roll", map[string]string{
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

	return &RollAzureGroupOutput{}, nil
}

// region AzureGroup

func (o *AzureGroup) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroup
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroup) SetId(v *string) *AzureGroup {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *AzureGroup) SetName(v *string) *AzureGroup {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AzureGroup) SetDescription(v *string) *AzureGroup {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *AzureGroup) SetCapacity(v *AzureGroupCapacity) *AzureGroup {
	if o.Capacity = v; o.Capacity == nil {
		o.nullFields = append(o.nullFields, "Capacity")
	}
	return o
}

func (o *AzureGroup) SetCompute(v *AzureGroupCompute) *AzureGroup {
	if o.Compute = v; o.Compute == nil {
		o.nullFields = append(o.nullFields, "Compute")
	}
	return o
}

func (o *AzureGroup) SetStrategy(v *AzureGroupStrategy) *AzureGroup {
	if o.Strategy = v; o.Strategy == nil {
		o.nullFields = append(o.nullFields, "Strategy")
	}
	return o
}

func (o *AzureGroup) SetScaling(v *AzureGroupScaling) *AzureGroup {
	if o.Scaling = v; o.Scaling == nil {
		o.nullFields = append(o.nullFields, "Scaling")
	}
	return o
}

func (o *AzureGroup) SetScheduling(v *AzureGroupScheduling) *AzureGroup {
	if o.Scheduling = v; o.Scheduling == nil {
		o.nullFields = append(o.nullFields, "Scheduling")
	}
	return o
}

func (o *AzureGroup) SetIntegration(v *AzureGroupIntegration) *AzureGroup {
	if o.Integration = v; o.Integration == nil {
		o.nullFields = append(o.nullFields, "Integration")
	}
	return o
}

// endregion

// region AzureGroupScheduling

func (o *AzureGroupScheduling) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupScheduling
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupScheduling) SetTasks(v []*AzureGroupScheduledTask) *AzureGroupScheduling {
	if o.Tasks = v; o.Tasks == nil {
		o.nullFields = append(o.nullFields, "Tasks")
	}
	return o
}

// endregion

// region AzureGroupIntegration

func (o *AzureGroupIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupIntegration) SetRancher(v *AzureGroupRancherIntegration) *AzureGroupIntegration {
	if o.Rancher = v; o.Rancher == nil {
		o.nullFields = append(o.nullFields, "Rancher")
	}
	return o
}

// endregion

// region AzureGroupRancherIntegration

func (o *AzureGroupRancherIntegration) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupRancherIntegration
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupRancherIntegration) SetMasterHost(v *string) *AzureGroupRancherIntegration {
	if o.MasterHost = v; o.MasterHost == nil {
		o.nullFields = append(o.nullFields, "MasterHost")
	}
	return o
}

func (o *AzureGroupRancherIntegration) SetAccessKey(v *string) *AzureGroupRancherIntegration {
	if o.AccessKey = v; o.AccessKey == nil {
		o.nullFields = append(o.nullFields, "AccessKey")
	}
	return o
}

func (o *AzureGroupRancherIntegration) SetSecretKey(v *string) *AzureGroupRancherIntegration {
	if o.SecretKey = v; o.SecretKey == nil {
		o.nullFields = append(o.nullFields, "SecretKey")
	}
	return o
}

// endregion

// region AzureGroupScheduledTask

func (o *AzureGroupScheduledTask) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupScheduledTask
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupScheduledTask) SetIsEnabled(v *bool) *AzureGroupScheduledTask {
	if o.IsEnabled = v; o.IsEnabled == nil {
		o.nullFields = append(o.nullFields, "IsEnabled")
	}
	return o
}

func (o *AzureGroupScheduledTask) SetFrequency(v *string) *AzureGroupScheduledTask {
	if o.Frequency = v; o.Frequency == nil {
		o.nullFields = append(o.nullFields, "Frequency")
	}
	return o
}

func (o *AzureGroupScheduledTask) SetCronExpression(v *string) *AzureGroupScheduledTask {
	if o.CronExpression = v; o.CronExpression == nil {
		o.nullFields = append(o.nullFields, "CronExpression")
	}
	return o
}

func (o *AzureGroupScheduledTask) SetTaskType(v *string) *AzureGroupScheduledTask {
	if o.TaskType = v; o.TaskType == nil {
		o.nullFields = append(o.nullFields, "TaskType")
	}
	return o
}

func (o *AzureGroupScheduledTask) SetScaleTargetCapacity(v *int) *AzureGroupScheduledTask {
	if o.ScaleTargetCapacity = v; o.ScaleTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "ScaleTargetCapacity")
	}
	return o
}

func (o *AzureGroupScheduledTask) SetScaleMinCapacity(v *int) *AzureGroupScheduledTask {
	if o.ScaleMinCapacity = v; o.ScaleMinCapacity == nil {
		o.nullFields = append(o.nullFields, "ScaleMinCapacity")
	}
	return o
}

func (o *AzureGroupScheduledTask) SetScaleMaxCapacity(v *int) *AzureGroupScheduledTask {
	if o.ScaleMaxCapacity = v; o.ScaleMaxCapacity == nil {
		o.nullFields = append(o.nullFields, "ScaleMaxCapacity")
	}
	return o
}

func (o *AzureGroupScheduledTask) SetBatchSizePercentage(v *int) *AzureGroupScheduledTask {
	if o.BatchSizePercentage = v; o.BatchSizePercentage == nil {
		o.nullFields = append(o.nullFields, "BatchSizePercentage")
	}
	return o
}

func (o *AzureGroupScheduledTask) SetGracePeriod(v *int) *AzureGroupScheduledTask {
	if o.GracePeriod = v; o.GracePeriod == nil {
		o.nullFields = append(o.nullFields, "GracePeriod")
	}
	return o
}

func (o *AzureGroupScheduledTask) SetAdjustment(v *int) *AzureGroupScheduledTask {
	if o.Adjustment = v; o.Adjustment == nil {
		o.nullFields = append(o.nullFields, "Adjustment")
	}
	return o
}

func (o *AzureGroupScheduledTask) SetAdjustmentPercentage(v *int) *AzureGroupScheduledTask {
	if o.AdjustmentPercentage = v; o.AdjustmentPercentage == nil {
		o.nullFields = append(o.nullFields, "AdjustmentPercentage")
	}
	return o
}

// endregion

// region AzureGroupScaling

func (o *AzureGroupScaling) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupScaling
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupScaling) SetUp(v []*AzureGroupScalingPolicy) *AzureGroupScaling {
	if o.Up = v; o.Up == nil {
		o.nullFields = append(o.nullFields, "Up")
	}
	return o
}

func (o *AzureGroupScaling) SetDown(v []*AzureGroupScalingPolicy) *AzureGroupScaling {
	if o.Down = v; o.Down == nil {
		o.nullFields = append(o.nullFields, "Down")
	}
	return o
}

// endregion

// region AzureGroupScalingPolicy

func (o *AzureGroupScalingPolicy) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupScalingPolicy
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupScalingPolicy) SetPolicyName(v *string) *AzureGroupScalingPolicy {
	if o.PolicyName = v; o.PolicyName == nil {
		o.nullFields = append(o.nullFields, "PolicyName")
	}
	return o
}

func (o *AzureGroupScalingPolicy) SetMetricName(v *string) *AzureGroupScalingPolicy {
	if o.MetricName = v; o.MetricName == nil {
		o.nullFields = append(o.nullFields, "MetricName")
	}
	return o
}

func (o *AzureGroupScalingPolicy) SetStatistic(v *string) *AzureGroupScalingPolicy {
	if o.Statistic = v; o.Statistic == nil {
		o.nullFields = append(o.nullFields, "Statistic")
	}
	return o
}

func (o *AzureGroupScalingPolicy) SetUnit(v *string) *AzureGroupScalingPolicy {
	if o.Unit = v; o.Unit == nil {
		o.nullFields = append(o.nullFields, "Unit")
	}
	return o
}

func (o *AzureGroupScalingPolicy) SetThreshold(v *float64) *AzureGroupScalingPolicy {
	if o.Threshold = v; o.Threshold == nil {
		o.nullFields = append(o.nullFields, "Threshold")
	}
	return o
}

func (o *AzureGroupScalingPolicy) SetAdjustment(v *int) *AzureGroupScalingPolicy {
	if o.Adjustment = v; o.Adjustment == nil {
		o.nullFields = append(o.nullFields, "Adjustment")
	}
	return o
}

func (o *AzureGroupScalingPolicy) SetMinTargetCapacity(v *int) *AzureGroupScalingPolicy {
	if o.MinTargetCapacity = v; o.MinTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "MinTargetCapacity")
	}
	return o
}

func (o *AzureGroupScalingPolicy) SetMaxTargetCapacity(v *int) *AzureGroupScalingPolicy {
	if o.MaxTargetCapacity = v; o.MaxTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "MaxTargetCapacity")
	}
	return o
}

func (o *AzureGroupScalingPolicy) SetNamespace(v *string) *AzureGroupScalingPolicy {
	if o.Namespace = v; o.Namespace == nil {
		o.nullFields = append(o.nullFields, "Namespace")
	}
	return o
}

func (o *AzureGroupScalingPolicy) SetEvaluationPeriods(v *int) *AzureGroupScalingPolicy {
	if o.EvaluationPeriods = v; o.EvaluationPeriods == nil {
		o.nullFields = append(o.nullFields, "EvaluationPeriods")
	}
	return o
}

func (o *AzureGroupScalingPolicy) SetPeriod(v *int) *AzureGroupScalingPolicy {
	if o.Period = v; o.Period == nil {
		o.nullFields = append(o.nullFields, "Period")
	}
	return o
}

func (o *AzureGroupScalingPolicy) SetCooldown(v *int) *AzureGroupScalingPolicy {
	if o.Cooldown = v; o.Cooldown == nil {
		o.nullFields = append(o.nullFields, "Cooldown")
	}
	return o
}

func (o *AzureGroupScalingPolicy) SetOperator(v *string) *AzureGroupScalingPolicy {
	if o.Operator = v; o.Operator == nil {
		o.nullFields = append(o.nullFields, "Operator")
	}
	return o
}

func (o *AzureGroupScalingPolicy) SetDimensions(v []*AzureGroupScalingPolicyDimension) *AzureGroupScalingPolicy {
	if o.Dimensions = v; o.Dimensions == nil {
		o.nullFields = append(o.nullFields, "Dimensions")
	}
	return o
}

func (o *AzureGroupScalingPolicy) SetAction(v *AzureGroupScalingPolicyAction) *AzureGroupScalingPolicy {
	if o.Action = v; o.Action == nil {
		o.nullFields = append(o.nullFields, "Action")
	}
	return o
}

// endregion

// region AzureGroupScalingPolicyAction

func (o *AzureGroupScalingPolicyAction) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupScalingPolicyAction
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupScalingPolicyAction) SetType(v *string) *AzureGroupScalingPolicyAction {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *AzureGroupScalingPolicyAction) SetAdjustment(v *string) *AzureGroupScalingPolicyAction {
	if o.Adjustment = v; o.Adjustment == nil {
		o.nullFields = append(o.nullFields, "Adjustment")
	}
	return o
}

func (o *AzureGroupScalingPolicyAction) SetMinTargetCapacity(v *string) *AzureGroupScalingPolicyAction {
	if o.MinTargetCapacity = v; o.MinTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "MinTargetCapacity")
	}
	return o
}

func (o *AzureGroupScalingPolicyAction) SetMaxTargetCapacity(v *string) *AzureGroupScalingPolicyAction {
	if o.MaxTargetCapacity = v; o.MaxTargetCapacity == nil {
		o.nullFields = append(o.nullFields, "MaxTargetCapacity")
	}
	return o
}

func (o *AzureGroupScalingPolicyAction) SetMaximum(v *string) *AzureGroupScalingPolicyAction {
	if o.Maximum = v; o.Maximum == nil {
		o.nullFields = append(o.nullFields, "Maximum")
	}
	return o
}

func (o *AzureGroupScalingPolicyAction) SetMinimum(v *string) *AzureGroupScalingPolicyAction {
	if o.Minimum = v; o.Minimum == nil {
		o.nullFields = append(o.nullFields, "Minimum")
	}
	return o
}

func (o *AzureGroupScalingPolicyAction) SetTarget(v *string) *AzureGroupScalingPolicyAction {
	if o.Target = v; o.Target == nil {
		o.nullFields = append(o.nullFields, "Target")
	}
	return o
}

// endregion

// region AzureGroupScalingPolicyDimension

func (o *AzureGroupScalingPolicyDimension) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupScalingPolicyDimension
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupScalingPolicyDimension) SetName(v *string) *AzureGroupScalingPolicyDimension {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AzureGroupScalingPolicyDimension) SetValue(v *string) *AzureGroupScalingPolicyDimension {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

// endregion

// region AzureGroupStrategy

func (o *AzureGroupStrategy) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupStrategy
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupStrategy) SetLowPriorityPercentage(v *int) *AzureGroupStrategy {
	if o.LowPriorityPercentage = v; o.LowPriorityPercentage == nil {
		o.nullFields = append(o.nullFields, "LowPriorityPercentage")
	}
	return o
}

func (o *AzureGroupStrategy) SetDedicatedCount(v *int) *AzureGroupStrategy {
	if o.DedicatedCount = v; o.DedicatedCount == nil {
		o.nullFields = append(o.nullFields, "DedicatedCount")
	}
	return o
}

func (o *AzureGroupStrategy) SetDrainingTimeout(v *int) *AzureGroupStrategy {
	if o.DrainingTimeout = v; o.DrainingTimeout == nil {
		o.nullFields = append(o.nullFields, "DrainingTimeout")
	}
	return o
}

func (o *AzureGroupStrategy) SetSignals(v []*AzureGroupStrategySignal) *AzureGroupStrategy {
	if o.Signals = v; o.Signals == nil {
		o.nullFields = append(o.nullFields, "Signals")
	}
	return o
}

// endregion

// region AzureGroupStrategySignal

func (o *AzureGroupStrategySignal) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupStrategySignal
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupStrategySignal) SetName(v *string) *AzureGroupStrategySignal {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AzureGroupStrategySignal) SetTimeout(v *int) *AzureGroupStrategySignal {
	if o.Timeout = v; o.Timeout == nil {
		o.nullFields = append(o.nullFields, "Timeout")
	}
	return o
}

// endregion

// region AzureGroupCapacity

func (o *AzureGroupCapacity) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupCapacity
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupCapacity) SetMinimum(v *int) *AzureGroupCapacity {
	if o.Minimum = v; o.Minimum == nil {
		o.nullFields = append(o.nullFields, "Minimum")
	}
	return o
}

func (o *AzureGroupCapacity) SetMaximum(v *int) *AzureGroupCapacity {
	if o.Maximum = v; o.Maximum == nil {
		o.nullFields = append(o.nullFields, "Maximum")
	}
	return o
}

func (o *AzureGroupCapacity) SetTarget(v *int) *AzureGroupCapacity {
	if o.Target = v; o.Target == nil {
		o.nullFields = append(o.nullFields, "Target")
	}
	return o
}

// endregion

// region AzureGroupCompute

func (o *AzureGroupCompute) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupCompute
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupCompute) SetRegion(v *string) *AzureGroupCompute {
	if o.Region = v; o.Region == nil {
		o.nullFields = append(o.nullFields, "Region")
	}
	return o
}

func (o *AzureGroupCompute) SetProduct(v *string) *AzureGroupCompute {
	if o.Product = v; o.Product == nil {
		o.nullFields = append(o.nullFields, "Product")
	}
	return o
}

func (o *AzureGroupCompute) SetResourceGroupName(v *string) *AzureGroupCompute {
	if o.ResourceGroupName = v; o.ResourceGroupName == nil {
		o.nullFields = append(o.nullFields, "ResourceGroupName")
	}
	return o
}

func (o *AzureGroupCompute) SetVMSize(v *AzureGroupComputeVMSize) *AzureGroupCompute {
	if o.VMSize = v; o.VMSize == nil {
		o.nullFields = append(o.nullFields, "VMSize")
	}
	return o
}

func (o *AzureGroupCompute) SetLaunchSpecification(v *AzureGroupComputeLaunchSpecification) *AzureGroupCompute {
	if o.LaunchSpecification = v; o.LaunchSpecification == nil {
		o.nullFields = append(o.nullFields, "LaunchSpecification")
	}
	return o
}

func (o *AzureGroupCompute) SetHealth(v *AzureGroupComputeHealth) *AzureGroupCompute {
	if o.Health = v; o.Health == nil {
		o.nullFields = append(o.nullFields, "Health")
	}
	return o
}

// endregion

// region AzureGroupComputeVMSize

func (o *AzureGroupComputeVMSize) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupComputeVMSize
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupComputeVMSize) SetDedicated(v []string) *AzureGroupComputeVMSize {
	if o.Dedicated = v; o.Dedicated == nil {
		o.nullFields = append(o.nullFields, "Dedicated")
	}
	return o
}

func (o *AzureGroupComputeVMSize) SetLowPriority(v []string) *AzureGroupComputeVMSize {
	if o.LowPriority = v; o.LowPriority == nil {
		o.nullFields = append(o.nullFields, "LowPriority")
	}
	return o
}

// endregion

// region AzureGroupComputeLaunchSpecification

func (o *AzureGroupComputeLaunchSpecification) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupComputeLaunchSpecification
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupComputeLaunchSpecification) SetLoadBalancersConfig(v *AzureGroupComputeLoadBalancersConfig) *AzureGroupComputeLaunchSpecification {
	if o.LoadBalancersConfig = v; o.LoadBalancersConfig == nil {
		o.nullFields = append(o.nullFields, "LoadBalancersConfig")
	}
	return o
}

func (o *AzureGroupComputeLaunchSpecification) SetImage(v *AzureGroupComputeImage) *AzureGroupComputeLaunchSpecification {
	if o.Image = v; o.Image == nil {
		o.nullFields = append(o.nullFields, "Image")
	}
	return o
}

func (o *AzureGroupComputeLaunchSpecification) SetUserData(v *AzureGroupComputeUserData) *AzureGroupComputeLaunchSpecification {
	if o.UserData = v; o.UserData == nil {
		o.nullFields = append(o.nullFields, "UserData")
	}
	return o
}

func (o *AzureGroupComputeLaunchSpecification) SetStorage(v *AzureGroupComputeStorage) *AzureGroupComputeLaunchSpecification {
	if o.Storage = v; o.Storage == nil {
		o.nullFields = append(o.nullFields, "Storage")
	}
	return o
}

func (o *AzureGroupComputeLaunchSpecification) SetNetwork(v *AzureGroupComputeNetwork) *AzureGroupComputeLaunchSpecification {
	if o.Network = v; o.Network == nil {
		o.nullFields = append(o.nullFields, "Network")
	}
	return o
}

func (o *AzureGroupComputeLaunchSpecification) SetSSHPublicKey(v *string) *AzureGroupComputeLaunchSpecification {
	if o.SSHPublicKey = v; o.SSHPublicKey == nil {
		o.nullFields = append(o.nullFields, "SSHPublicKey")
	}
	return o
}

// endregion

// region AzureGroupComputeLoadBalancersConfig

func (o *AzureGroupComputeLoadBalancersConfig) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupComputeLoadBalancersConfig
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupComputeLoadBalancersConfig) SetLoadBalancers(v []*AzureGroupComputeLoadBalancer) *AzureGroupComputeLoadBalancersConfig {
	if o.LoadBalancers = v; o.LoadBalancers == nil {
		o.nullFields = append(o.nullFields, "LoadBalancers")
	}
	return o
}

// endregion

// region AzureGroupComputeLoadBalancer

func (o *AzureGroupComputeLoadBalancer) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupComputeLoadBalancer
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupComputeLoadBalancer) SetBalancerId(v *string) *AzureGroupComputeLoadBalancer {
	if o.BalancerID = v; o.BalancerID == nil {
		o.nullFields = append(o.nullFields, "BalancerID")
	}
	return o
}

func (o *AzureGroupComputeLoadBalancer) SetTargetSetId(v *string) *AzureGroupComputeLoadBalancer {
	if o.TargetSetID = v; o.TargetSetID == nil {
		o.nullFields = append(o.nullFields, "TargetSetID")
	}
	return o
}

// endregion

// region AzureGroupComputeImage

func (o *AzureGroupComputeImage) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupComputeImage
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupComputeImage) SetCustom(v *AzureGroupComputeImageCustom) *AzureGroupComputeImage {
	if o.Custom = v; o.Custom == nil {
		o.nullFields = append(o.nullFields, "Custom")
	}
	return o
}

func (o *AzureGroupComputeImage) SetPublisher(v *string) *AzureGroupComputeImage {
	if o.Publisher = v; o.Publisher == nil {
		o.nullFields = append(o.nullFields, "Publisher")
	}
	return o
}

func (o *AzureGroupComputeImage) SetOffer(v *string) *AzureGroupComputeImage {
	if o.Offer = v; o.Offer == nil {
		o.nullFields = append(o.nullFields, "Offer")
	}
	return o
}

func (o *AzureGroupComputeImage) SetSKU(v *string) *AzureGroupComputeImage {
	if o.SKU = v; o.SKU == nil {
		o.nullFields = append(o.nullFields, "SKU")
	}
	return o
}

// endregion

// region AzureGroupComputeImageCustom

func (o *AzureGroupComputeImageCustom) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupComputeImageCustom
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupComputeImageCustom) SetImageURIs(v []string) *AzureGroupComputeImageCustom {
	if o.ImageURIs = v; o.ImageURIs == nil {
		o.nullFields = append(o.nullFields, "ImageURIs")
	}
	return o
}

// endregion

// region AzureGroupComputeUserData

func (o *AzureGroupComputeUserData) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupComputeUserData
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupComputeUserData) SetCommandLine(v *string) *AzureGroupComputeUserData {
	if o.CommandLine = v; o.CommandLine == nil {
		o.nullFields = append(o.nullFields, "CommandLine")
	}
	return o
}

func (o *AzureGroupComputeUserData) SetResourceFiles(v []*AzureGroupComputeUserDataResourceFile) *AzureGroupComputeUserData {
	if o.ResourceFiles = v; o.ResourceFiles == nil {
		o.nullFields = append(o.nullFields, "ResourceFiles")
	}
	return o
}

// endregion

// region AzureGroupComputeUserDataResourceFile

func (o *AzureGroupComputeUserDataResourceFile) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupComputeUserDataResourceFile
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupComputeUserDataResourceFile) SetURL(v *string) *AzureGroupComputeUserDataResourceFile {
	if o.URL = v; o.URL == nil {
		o.nullFields = append(o.nullFields, "URL")
	}
	return o
}

func (o *AzureGroupComputeUserDataResourceFile) SetTargetPath(v *string) *AzureGroupComputeUserDataResourceFile {
	if o.TargetPath = v; o.TargetPath == nil {
		o.nullFields = append(o.nullFields, "TargetPath")
	}
	return o
}

// endregion

// region AzureGroupComputeStorage

func (o *AzureGroupComputeStorage) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupComputeStorage
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupComputeStorage) SetAccountName(v *string) *AzureGroupComputeStorage {
	if o.AccountName = v; o.AccountName == nil {
		o.nullFields = append(o.nullFields, "AccountName")
	}
	return o
}

// endregion

// region AzureGroupComputeNetwork

func (o *AzureGroupComputeNetwork) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupComputeNetwork
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupComputeNetwork) SetVirtualNetworkName(v *string) *AzureGroupComputeNetwork {
	if o.VirtualNetworkName = v; o.VirtualNetworkName == nil {
		o.nullFields = append(o.nullFields, "VirtualNetworkName")
	}
	return o
}

func (o *AzureGroupComputeNetwork) SetSubnetId(v *string) *AzureGroupComputeNetwork {
	if o.SubnetID = v; o.SubnetID == nil {
		o.nullFields = append(o.nullFields, "SubnetID")
	}
	return o
}

// endregion

// region AzureGroupComputeHealth

func (o *AzureGroupComputeHealth) MarshalJSON() ([]byte, error) {
	type noMethod AzureGroupComputeHealth
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AzureGroupComputeHealth) SetHealthCheckType(v *string) *AzureGroupComputeHealth {
	if o.HealthCheckType = v; o.HealthCheckType == nil {
		o.nullFields = append(o.nullFields, "HealthCheckType")
	}
	return o
}

func (o *AzureGroupComputeHealth) SetAutoHealing(v *bool) *AzureGroupComputeHealth {
	if o.AutoHealing = v; o.AutoHealing == nil {
		o.nullFields = append(o.nullFields, "AutoHealing")
	}
	return o
}

func (o *AzureGroupComputeHealth) SetGracePeriod(v *int) *AzureGroupComputeHealth {
	if o.GracePeriod = v; o.GracePeriod == nil {
		o.nullFields = append(o.nullFields, "GracePeriod")
	}
	return o
}

// endregion
