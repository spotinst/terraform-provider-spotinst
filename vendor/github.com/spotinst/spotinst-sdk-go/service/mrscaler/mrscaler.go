package mrscaler

import (
	"context"
	"encoding/json"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
	"io/ioutil"
	"net/http"
)

// A InstanceGroupType represents the type of an instance group.
type InstanceGroupType int

const (
	// InstanceGroupTypeMaster represents the master instance group type.
	InstanceGroupTypeMaster InstanceGroupType = iota

	// InstanceGroupTypeCore represents the core instance group type.
	InstanceGroupTypeCore

	// InstanceGroupTypeTask represents the task instance group type.
	InstanceGroupTypeTask
)

var InstanceGroupTypeName = map[InstanceGroupType]string{
	InstanceGroupTypeMaster: "master",
	InstanceGroupTypeCore:   "core",
	InstanceGroupTypeTask:   "task",
}

var InstanceGroupTypeValue = map[string]InstanceGroupType{
	"master": InstanceGroupTypeMaster,
	"core":   InstanceGroupTypeCore,
	"task":   InstanceGroupTypeTask,
}

func (p InstanceGroupType) String() string {
	return InstanceGroupTypeName[p]
}

type Scaler struct {
	ID                  *string              `json:"id,omitempty"`
	Name                *string              `json:"name,omitempty"`
	Description         *string              `json:"description,omitempty"`
	Region              *string              `json:"region,omitempty"`
	Strategy            *Strategy            `json:"strategy,omitempty"`
	Compute             *Compute             `json:"compute,omitempty"`
	Cluster             *Cluster             `json:"cluster,omitempty"`
	Scaling             *Scaling             `json:"scaling,omitempty"`
	CoreScaling         *Scaling             `json:"coreScaling,omitempty"`
	Scheduling          *Scheduling          `json:"scheduling,omitempty"`
	TerminationPolicies []*TerminationPolicy `json:"terminationPolicies,omitempty"`

	// forceSendFields is a list of field names (e.g. "Keys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	forceSendFields []string

	// nullFields is a list of field names (e.g. "Keys") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	nullFields []string
}
type TerminationPolicy struct {
	Statements []*Statement `json:"statements,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Statement struct {
	Namespace         *string  `json:"namespace,omitempty"`
	MetricName        *string  `json:"metricName,omitempty"`
	Statistic         *string  `json:"statistic,omitempty"`
	Unit              *string  `json:"unit,omitempty"`
	Threshold         *float64 `json:"threshold,omitempty"`
	Period            *int     `json:"period,omitempty"`
	EvaluationPeriods *int     `json:"evaluationPeriods,omitempty"`
	Operator          *string  `json:"operator,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Strategy struct {
	Cloning             *Cloning             `json:"cloning,omitempty"`
	Wrapping            *Wrapping            `json:"wrapping,omitempty"`
	CreateNew           *CreateNew           `json:"new,omitempty"`
	ProvisioningTimeout *ProvisioningTimeout `json:"provisioningTimeout,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Cloning struct {
	OriginClusterID *string `json:"originClusterId,omitempty"`
	Retries         *int    `json:"numberOfRetries,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Wrapping struct {
	SourceClusterID *string `json:"sourceClusterId,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type CreateNew struct {
	ReleaseLabel *string `json:"releaseLabel,omitempty"`
	Retries      *int    `json:"numberOfRetries,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ProvisioningTimeout struct {
	Timeout       *int    `json:"timeout,omitempty"`
	TimeoutAction *string `json:"timeoutAction,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Compute struct {
	AvailabilityZones               []*AvailabilityZone `json:"availabilityZones,omitempty"`
	Tags                            []*Tag              `json:"tags,omitempty"`
	InstanceGroups                  *InstanceGroups     `json:"instanceGroups,omitempty"`
	Configurations                  *Configurations     `json:"configurations,omitempty"`
	EBSRootVolumeSize               *int                `json:"ebsRootVolumeSize,omitempty"`
	ManagedPrimarySecurityGroup     *string             `json:"emrManagedMasterSecurityGroup,omitempty"`
	ManagedReplicaSecurityGroup     *string             `json:"emrManagedSlaveSecurityGroup,omitempty"`
	ServiceAccessSecurityGroup      *string             `json:"serviceAccessSecurityGroup,omitempty"`
	AdditionalPrimarySecurityGroups []string            `json:"additionalMasterSecurityGroups,omitempty"`
	AdditionalReplicaSecurityGroups []string            `json:"additionalSlaveSecurityGroups,omitempty"`
	CustomAMIID                     *string             `json:"customAmiId,omitempty"`
	RepoUpgradeOnBoot               *string             `json:"repoUpgradeOnBoot,omitempty"`
	EC2KeyName                      *string             `json:"ec2KeyName,omitempty"`
	Applications                    []*Application      `json:"applications,omitempty"`
	BootstrapActions                *BootstrapActions   `json:"bootstrapActions,omitempty"`
	Steps                           *Steps              `json:"steps,omitempty"`
	InstanceWeights                 []*InstanceWeight   `json:"instanceWeights,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Cluster struct {
	LogURI                      *string `json:"logUri,omitempty"`
	AdditionalInfo              *string `json:"additionalInfo,omitempty"`
	JobFlowRole                 *string `json:"jobFlowRole,omitempty"`
	SecurityConfiguration       *string `json:"securityConfiguration,omitempty"`
	ServiceRole                 *string `json:"serviceRole,omitempty"`
	VisibleToAllUsers           *bool   `json:"visibleToAllUsers,omitempty"`
	TerminationProtected        *bool   `json:"terminationProtected,omitempty"`
	KeepJobFlowAliveWhenNoSteps *bool   `json:"keepJobFlowAliveWhenNoSteps,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type AvailabilityZone struct {
	Name     *string `json:"name,omitempty"`
	SubnetID *string `json:"subnetId,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Tag struct {
	Key   *string `json:"tagKey,omitempty"`
	Value *string `json:"tagValue,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Application struct {
	Args    []string `json:"args,omitempty"`
	Name    *string  `json:"name,omitempty"`
	Version *string  `json:"version,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type InstanceWeight struct {
	InstanceType     *string `json:"instanceType,omitempty"`
	WeightedCapacity *int    `json:"weightedCapacity,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type InstanceGroups struct {
	MasterGroup *InstanceGroup `json:"masterGroup,omitempty"`
	CoreGroup   *InstanceGroup `json:"coreGroup,omitempty"`
	TaskGroup   *InstanceGroup `json:"taskGroup,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type InstanceGroup struct {
	InstanceTypes    []string               `json:"instanceTypes,omitempty"`
	Target           *int                   `json:"target,omitempty"`
	Capacity         *InstanceGroupCapacity `json:"capacity,omitempty"`
	LifeCycle        *string                `json:"lifeCycle,omitempty"`
	EBSConfiguration *EBSConfiguration      `json:"ebsConfiguration,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type InstanceGroupCapacity struct {
	Target  *int    `json:"target,omitempty"`
	Minimum *int    `json:"minimum,omitempty"`
	Maximum *int    `json:"maximum,omitempty"`
	Unit    *string `json:"unit,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type EBSConfiguration struct {
	Optimized          *bool                `json:"ebsOptimized,omitempty"`
	BlockDeviceConfigs []*BlockDeviceConfig `json:"ebsBlockDeviceConfigs,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type BlockDeviceConfig struct {
	VolumesPerInstance  *int                 `json:"volumesPerInstance,omitempty"`
	VolumeSpecification *VolumeSpecification `json:"volumeSpecification,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type VolumeSpecification struct {
	VolumeType *string `json:"volumeType,omitempty"`
	SizeInGB   *int    `json:"sizeInGB,omitempty"`
	IOPS       *int    `json:"iops,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Scaling struct {
	Up   []*ScalingPolicy `json:"up,omitempty"`
	Down []*ScalingPolicy `json:"down,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ScalingPolicy struct {
	PolicyName        *string      `json:"policyName,omitempty"`
	Namespace         *string      `json:"namespace,omitempty"`
	MetricName        *string      `json:"metricName,omitempty"`
	Dimensions        []*Dimension `json:"dimensions,omitempty"`
	Statistic         *string      `json:"statistic,omitempty"`
	Unit              *string      `json:"unit,omitempty"`
	Threshold         *float64     `json:"threshold,omitempty"`
	Period            *int         `json:"period,omitempty"`
	EvaluationPeriods *int         `json:"evaluationPeriods,omitempty"`
	Cooldown          *int         `json:"cooldown,omitempty"`
	Action            *Action      `json:"action,omitempty"`
	Operator          *string      `json:"operator,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Action struct {
	Type              *string `json:"type,omitempty"`
	Adjustment        *string `json:"adjustment,omitempty"`
	MinTargetCapacity *string `json:"minTargetCapacity,omitempty"`
	MaxTargetCapacity *string `json:"maxTargetCapacity,omitempty"`
	Target            *string `json:"target,omitempty"`
	Minimum           *string `json:"minimum,omitempty"`
	Maximum           *string `json:"maximum,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Dimension struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Configurations struct {
	File *S3File `json:"file,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type BootstrapActions struct {
	File *S3File `json:"file,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Steps struct {
	File *S3File `json:"file,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type S3File struct {
	Bucket *string `json:"bucket,omitempty"`
	Key    *string `json:"key,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Scheduling struct {
	Tasks []*Task `json:"tasks,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Task struct {
	IsEnabled         *bool   `json:"isEnabled,omitempty"`
	Type              *string `json:"taskType,omitempty"`
	InstanceGroupType *string `json:"instanceGroupType"`
	CronExpression    *string `json:"cronExpression,omitempty"`
	TargetCapacity    *int    `json:"targetCapacity,omitempty"`
	MinCapacity       *int    `json:"minCapacity,omitempty"`
	MaxCapacity       *int    `json:"maxCapacity,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ListScalersInput struct{}

type ListScalersOutput struct {
	Scalers []*Scaler `json:"mrScalers,omitempty"`
}

type CreateScalerInput struct {
	Scaler *Scaler `json:"mrScaler,omitempty"`
}

type CreateScalerOutput struct {
	Scaler *Scaler `json:"mrScaler,omitempty"`
}

type ReadScalerInput struct {
	ScalerID *string `json:"mrScalerId,omitempty"`
}

type ReadScalerOutput struct {
	Scaler *Scaler `json:"mrScaler,omitempty"`
}

type UpdateScalerInput struct {
	Scaler *Scaler `json:"mrScaler,omitempty"`
}

type UpdateScalerOutput struct {
	Scaler *Scaler `json:"mrScaler,omitempty"`
}

type DeleteScalerInput struct {
	ScalerID *string `json:"mrScalerId,omitempty"`
}

type DeleteScalerOutput struct{}

type ScalerCluster struct {
	ScalerClusterId *string `json:"id,omitempty"`
}

type ScalerClusterStatusInput struct {
	ScalerID *string `json:"mrScalerId,omitempty"`
}

type ScalerClusterStatusOutput struct {
	ScalerClusterId *string `json:"id,omitempty"`
}

func scalerFromJSON(in []byte) (*Scaler, error) {
	b := new(Scaler)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func scalerClusterFromJSON(in []byte) (*ScalerCluster, error) {
	b := new(ScalerCluster)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func scalersFromJSON(in []byte) ([]*Scaler, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Scaler, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := scalerFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func scalerClustersFromJSON(in []byte) ([]*ScalerCluster, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*ScalerCluster, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := scalerClusterFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func scalersFromHttpResponse(resp *http.Response) ([]*Scaler, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return scalersFromJSON(body)
}

func scalerClustersFromHttpResponse(resp *http.Response) ([]*ScalerCluster, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return scalerClustersFromJSON(body)
}

//region Scaler

func (o Scaler) MarshalJSON() ([]byte, error) {
	type noMethod Scaler
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Scaler) SetId(v *string) *Scaler {
	if o.ID = v; v == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *Scaler) SetName(v *string) *Scaler {
	if o.Name = v; v == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Scaler) SetDescription(v *string) *Scaler {
	if o.Description = v; v == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *Scaler) SetRegion(v *string) *Scaler {
	if o.Region = v; v == nil {
		o.nullFields = append(o.nullFields, "Region")
	}
	return o
}

func (o *Scaler) SetStrategy(v *Strategy) *Scaler {
	if o.Strategy = v; v == nil {
		o.nullFields = append(o.nullFields, "Strategy")
	}
	return o
}

func (o *Scaler) SetCompute(v *Compute) *Scaler {
	if o.Compute = v; v == nil {
		o.nullFields = append(o.nullFields, "Compute")
	}
	return o
}

// SetCluster sets the Cluster object used when creating a new Scaler
func (o *Scaler) SetCluster(v *Cluster) *Scaler {
	if o.Cluster = v; o.Cluster == nil {
		o.nullFields = append(o.nullFields, "Cluster")
	}
	return o
}

func (o *Scaler) SetScaling(v *Scaling) *Scaler {
	if o.Scaling = v; v == nil {
		o.nullFields = append(o.nullFields, "Scaling")
	}
	return o
}

func (o *Scaler) SetCoreScaling(v *Scaling) *Scaler {
	if o.CoreScaling = v; v == nil {
		o.nullFields = append(o.nullFields, "CoreScaling")
	}
	return o
}

func (o *Scaler) SetTerminationPolicies(v []*TerminationPolicy) *Scaler {
	if o.TerminationPolicies = v; v == nil {
		o.nullFields = append(o.nullFields, "TerminationPolicies")
	}
	return o
}

//endregion

// region TerminationPolicy

func (o TerminationPolicy) MarshalJSON() ([]byte, error) {
	type noMethod TerminationPolicy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *TerminationPolicy) SetStatements(v []*Statement) *TerminationPolicy {
	if o.Statements = v; v == nil {
		o.nullFields = append(o.nullFields, "Statements")
	}
	return o
}

//endregion

// region Statement

func (o Statement) MarshalJSON() ([]byte, error) {
	type noMethod Statement
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Statement) SetNamespace(v *string) *Statement {
	if o.Namespace = v; o.Namespace == nil {
		o.nullFields = append(o.nullFields, "Namespace")
	}
	return o
}

func (o *Statement) SetMetricName(v *string) *Statement {
	if o.MetricName = v; o.MetricName == nil {
		o.nullFields = append(o.nullFields, "MetricName")
	}
	return o
}

func (o *Statement) SetStatistic(v *string) *Statement {
	if o.Statistic = v; o.Statistic == nil {
		o.nullFields = append(o.nullFields, "Statistic")
	}
	return o
}

func (o *Statement) SetUnit(v *string) *Statement {
	if o.Unit = v; o.Unit == nil {
		o.nullFields = append(o.nullFields, "Unit")
	}
	return o
}

func (o *Statement) SetThreshold(v *float64) *Statement {
	if o.Threshold = v; o.Threshold == nil {
		o.nullFields = append(o.nullFields, "Threshold")
	}
	return o
}

func (o *Statement) SetPeriod(v *int) *Statement {
	if o.Period = v; o.Period == nil {
		o.nullFields = append(o.nullFields, "Period")
	}
	return o
}

func (o *Statement) SetEvaluationPeriods(v *int) *Statement {
	if o.EvaluationPeriods = v; o.EvaluationPeriods == nil {
		o.nullFields = append(o.nullFields, "EvaluationPeriods")
	}
	return o
}

func (o *Statement) SetOperator(v *string) *Statement {
	if o.Operator = v; o.Operator == nil {
		o.nullFields = append(o.nullFields, "Operator")
	}
	return o
}

//endregion

// region Cluster

func (o Cluster) MarshalJSON() ([]byte, error) {
	type noMethod Cluster
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

// SetLogURI sets the log uri when creating a new cluster
func (o *Cluster) SetLogURI(v *string) *Cluster {
	if o.LogURI = v; o.LogURI == nil {
		o.nullFields = append(o.nullFields, "LogURI")
	}
	return o
}

// SetAdditionalInfo sets the additional info field used by third party integrations when creating a new mrscaler
func (o *Cluster) SetAdditionalInfo(v *string) *Cluster {
	if o.AdditionalInfo = v; o.AdditionalInfo == nil {
		o.nullFields = append(o.nullFields, "AdditionalInfo")
	}
	return o
}

// SetJobFlowRole sets the IAM role that will be adopted by the launched EC2 instances
func (o *Cluster) SetJobFlowRole(v *string) *Cluster {
	if o.JobFlowRole = v; o.JobFlowRole == nil {
		o.nullFields = append(o.nullFields, "JobFlowRole")
	}
	return o
}

// SetSecurityConfiguration sets the name of the security configuration to be applied to the cluster
func (o *Cluster) SetSecurityConfiguration(v *string) *Cluster {
	if o.SecurityConfiguration = v; o.SecurityConfiguration == nil {
		o.nullFields = append(o.nullFields, "SecurityConfiguration")
	}
	return o
}

// SetServiceRole sets the IAM role that the EMR will assume to access AWS resources on your behalf
func (o *Cluster) SetServiceRole(v *string) *Cluster {
	if o.ServiceRole = v; o.ServiceRole == nil {
		o.nullFields = append(o.nullFields, "ServiceRole")
	}
	return o
}

// SetVisibleToAllUsers sets a flag indicating if the cluster is visibile to all IAM users
func (o *Cluster) SetVisibleToAllUsers(v *bool) *Cluster {
	if o.VisibleToAllUsers = v; o.VisibleToAllUsers == nil {
		o.nullFields = append(o.nullFields, "VisibleToAllUsers")
	}
	return o
}

// SetTerminationProtected sets whether the EC2 instances in the cluster are protected from terminating API calls
func (o *Cluster) SetTerminationProtected(v *bool) *Cluster {
	if o.TerminationProtected = v; o.TerminationProtected == nil {
		o.nullFields = append(o.nullFields, "TerminationProtected")
	}
	return o
}

// SetKeepJobFlowAliveWhenNoSteps sets KeepJobFlowAliveWhenNoSteps
func (o *Cluster) SetKeepJobFlowAliveWhenNoSteps(v *bool) *Cluster {
	if o.KeepJobFlowAliveWhenNoSteps = v; o.KeepJobFlowAliveWhenNoSteps == nil {
		o.nullFields = append(o.nullFields, "KeepJobFlowAliveWhenNoSteps")
	}
	return o
}

// endregion

// region Scheduling

func (o Scheduling) MarshalJSON() ([]byte, error) {
	type noMethod Scheduling
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Scheduling) SetTasks(v []*Task) *Scheduling {
	if o.Tasks = v; o.Tasks == nil {
		o.nullFields = append(o.nullFields, "Tasks")
	}
	return o
}

// endregion

// region Task

func (o Task) MarshalJSON() ([]byte, error) {
	type noMethod Task
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Task) SetIsEnabled(v *bool) *Task {
	if o.IsEnabled = v; o.IsEnabled == nil {
		o.nullFields = append(o.nullFields, "IsEnabled")
	}
	return o
}

func (o *Task) SetType(v *string) *Task {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

// SetInstanceGroupType sets the instance group to apply the scheduled task to.
func (o *Task) SetInstanceGroupType(v *string) *Task {
	if o.InstanceGroupType = v; o.InstanceGroupType == nil {
		o.nullFields = append(o.nullFields, "InstanceGroupType")
	}
	return o
}

func (o *Task) SetCronExpression(v *string) *Task {
	if o.CronExpression = v; o.CronExpression == nil {
		o.nullFields = append(o.nullFields, "CronExpression")
	}
	return o
}

func (o *Task) SetTargetCapacity(v *int) *Task {
	if o.TargetCapacity = v; o.TargetCapacity == nil {
		o.nullFields = append(o.nullFields, "TargetCapacity")
	}
	return o
}

func (o *Task) SetMinCapacity(v *int) *Task {
	if o.MinCapacity = v; o.MinCapacity == nil {
		o.nullFields = append(o.nullFields, "MinCapacity")
	}
	return o
}

func (o *Task) SetMaxCapacity(v *int) *Task {
	if o.MaxCapacity = v; o.MaxCapacity == nil {
		o.nullFields = append(o.nullFields, "MaxCapacity")
	}
	return o
}

// endregion

//region Strategy

func (o Strategy) MarshalJSON() ([]byte, error) {
	type noMethod Strategy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Strategy) SetCloning(v *Cloning) *Strategy {
	if o.Cloning = v; v == nil {
		o.nullFields = append(o.nullFields, "Cloning")
	}
	return o
}

func (o *Strategy) SetWrapping(v *Wrapping) *Strategy {
	if o.Wrapping = v; v == nil {
		o.nullFields = append(o.nullFields, "Wrapping")
	}
	return o
}

// SetCreateNew sets a new mrscaler object
func (o *Strategy) SetCreateNew(v *CreateNew) *Strategy {
	if o.CreateNew = v; o.CreateNew == nil {
		o.nullFields = append(o.nullFields, "CreateNew")
	}
	return o
}

// SetProvisioningTimeout sets the timeout when creating or cloning a scaler
func (o *Strategy) SetProvisioningTimeout(v *ProvisioningTimeout) *Strategy {
	if o.ProvisioningTimeout = v; o.ProvisioningTimeout == nil {
		o.nullFields = append(o.nullFields, "ProvisioningTimeout")
	}
	return o
}

//endregion

//region Cloning

func (o Cloning) MarshalJSON() ([]byte, error) {
	type noMethod Cloning
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Cloning) SetOriginClusterId(v *string) *Cloning {
	if o.OriginClusterID = v; v == nil {
		o.nullFields = append(o.nullFields, "OriginClusterID")
	}
	return o
}

// SetRetries sets the number of retries to attempt when cloning a scaler
func (o *Cloning) SetRetries(v *int) *Cloning {
	if o.Retries = v; o.Retries == nil {
		o.nullFields = append(o.nullFields, "Retries")
	}
	return o
}

//endregion

//region Wrapping

func (o Wrapping) MarshalJSON() ([]byte, error) {
	type noMethod Wrapping
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Wrapping) SetSourceClusterId(v *string) *Wrapping {
	if o.SourceClusterID = v; v == nil {
		o.nullFields = append(o.nullFields, "SourceClusterID")
	}
	return o
}

//endregion

// region CreateNew

func (o CreateNew) MarshalJSON() ([]byte, error) {
	type noMethod CreateNew
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

// SetReleaseLabel sets the release label for a new scaler
func (o *CreateNew) SetReleaseLabel(v *string) *CreateNew {
	if o.ReleaseLabel = v; o.ReleaseLabel == nil {
		o.nullFields = append(o.nullFields, "ReleaseLabel")
	}
	return o
}

// SetRetries sets the number of retries to attempt when creating a new scaler
func (o *CreateNew) SetRetries(v *int) *CreateNew {
	if o.Retries = v; o.Retries == nil {
		o.nullFields = append(o.nullFields, "Retries")
	}
	return o
}

// endregion

// region ProvisioningTimeout

func (o ProvisioningTimeout) MarshalJSON() ([]byte, error) {
	type noMethod ProvisioningTimeout
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

// SetTimeout sets the amount of time in seconds to wait for a scaler to be provisioned
func (o *ProvisioningTimeout) SetTimeout(v *int) *ProvisioningTimeout {
	if o.Timeout = v; o.Timeout == nil {
		o.nullFields = append(o.nullFields, "Timeout")
	}
	return o
}

// SetTimeoutAction sets the action to take on timeout
func (o *ProvisioningTimeout) SetTimeoutAction(v *string) *ProvisioningTimeout {
	if o.TimeoutAction = v; o.TimeoutAction == nil {
		o.nullFields = append(o.nullFields, "TimeoutAction")
	}
	return o
}

// endregion

// region Scheduling

// endregion

//region Compute

func (o Compute) MarshalJSON() ([]byte, error) {
	type noMethod Compute
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Compute) SetAvailabilityZones(v []*AvailabilityZone) *Compute {
	if o.AvailabilityZones = v; v == nil {
		o.nullFields = append(o.nullFields, "AvailabilityZones")
	}
	return o
}

func (o *Compute) SetTags(v []*Tag) *Compute {
	if o.Tags = v; v == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

func (o *Compute) SetInstanceGroups(v *InstanceGroups) *Compute {
	if o.InstanceGroups = v; v == nil {
		o.nullFields = append(o.nullFields, "InstanceGroups")
	}
	return o
}

func (o *Compute) SetConfigurations(v *Configurations) *Compute {
	if o.Configurations = v; v == nil {
		o.nullFields = append(o.nullFields, "Configurations")
	}
	return o
}

// SetBootstrapActions sets the path for a bootstrap actions file stored in AWS S3
func (o *Compute) SetBootstrapActions(v *BootstrapActions) *Compute {
	if o.BootstrapActions = v; o.BootstrapActions == nil {
		o.nullFields = append(o.nullFields, "BootstrapActions")
	}
	return o
}

// SetSteps sets the path for a steps file stored in AWS S3
func (o *Compute) SetSteps(v *Steps) *Compute {
	if o.Steps = v; o.Steps == nil {
		o.nullFields = append(o.nullFields, "Steps")
	}
	return o
}

// SetEBSRootVolumeSize sets the ebs root volume size when creating a new scaler
func (o *Compute) SetEBSRootVolumeSize(v *int) *Compute {
	if o.EBSRootVolumeSize = v; o.EBSRootVolumeSize == nil {
		o.nullFields = append(o.nullFields, "EBSRootVolumeSize")
	}
	return o
}

// SetManagedPrimarySecurityGroup sets the managed primary security group when creating a new scaler
func (o *Compute) SetManagedPrimarySecurityGroup(v *string) *Compute {
	if o.ManagedPrimarySecurityGroup = v; o.ManagedPrimarySecurityGroup == nil {
		o.nullFields = append(o.nullFields, "ManagedPrimarySecurityGroup")
	}
	return o
}

// SetManagedReplicaSecurityGroup sets the managed replica security group when creating a new scaler
func (o *Compute) SetManagedReplicaSecurityGroup(v *string) *Compute {
	if o.ManagedReplicaSecurityGroup = v; o.ManagedReplicaSecurityGroup == nil {
		o.nullFields = append(o.nullFields, "ManagedReplicaSecurityGroup")
	}
	return o
}

// SetServiceAccessSecurityGroup sets the service security group when creating a new scaler
func (o *Compute) SetServiceAccessSecurityGroup(v *string) *Compute {
	if o.ServiceAccessSecurityGroup = v; o.ServiceAccessSecurityGroup == nil {
		o.nullFields = append(o.nullFields, "ServiceAccessSecurityGroup")
	}
	return o
}

// SetAdditionalPrimarySecurityGroups sets a list of additional primary security groups
func (o *Compute) SetAdditionalPrimarySecurityGroups(v []string) *Compute {
	if o.AdditionalPrimarySecurityGroups = v; o.AdditionalPrimarySecurityGroups == nil {
		o.nullFields = append(o.nullFields, "AdditionalPrimarySecurityGroups")
	}
	return o
}

// SetAdditionalReplicaSecurityGroups sets a list of additional Replica security groups
func (o *Compute) SetAdditionalReplicaSecurityGroups(v []string) *Compute {
	if o.AdditionalReplicaSecurityGroups = v; o.AdditionalReplicaSecurityGroups == nil {
		o.nullFields = append(o.nullFields, "AdditionalReplicaSecurityGroups")
	}
	return o
}

// SetCustomAMIID sets the custom AMI ID
func (o *Compute) SetCustomAMIID(v *string) *Compute {
	if o.CustomAMIID = v; o.CustomAMIID == nil {
		o.nullFields = append(o.nullFields, "CustomAMIID")
	}
	return o
}

func (o *Compute) SetRepoUpgradeOnBoot(v *string) *Compute {
	if o.RepoUpgradeOnBoot = v; o.RepoUpgradeOnBoot == nil {
		o.nullFields = append(o.nullFields, "RepoUpgradeOnBoot")
	}
	return o
}

// SetEC2KeyName sets the EC2 key name
func (o *Compute) SetEC2KeyName(v *string) *Compute {
	if o.EC2KeyName = v; o.EC2KeyName == nil {
		o.nullFields = append(o.nullFields, "EC2KeyName")
	}
	return o
}

// SetApplications sets the applications object
func (o *Compute) SetApplications(v []*Application) *Compute {
	if o.Applications = v; o.Applications == nil {
		o.nullFields = append(o.nullFields, "Applications")
	}
	return o
}

// SetInstanceWeights sets a list if instance weights by type
func (o *Compute) SetInstanceWeights(v []*InstanceWeight) *Compute {
	if o.InstanceWeights = v; o.InstanceWeights == nil {
		o.nullFields = append(o.nullFields, "InstanceWeights")
	}
	return o
}

//endregion

// region Application

func (o Application) MarshalJSON() ([]byte, error) {
	type noMethod Application
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

// SetArgs sets the list of args to use with the application
func (o *Application) SetArgs(v []string) *Application {
	if o.Args = v; o.Args == nil {
		o.nullFields = append(o.nullFields, "Args")
	}
	return o
}

// SetName sets the name of the application
func (o *Application) SetName(v *string) *Application {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

// SetVersion sets the application version
func (o *Application) SetVersion(v *string) *Application {
	if o.Version = v; o.Version == nil {
		o.nullFields = append(o.nullFields, "Version")
	}
	return o
}

// endregion

// region InstanceWeight

func (o InstanceWeight) MarshalJSON() ([]byte, error) {
	type noMethod InstanceWeight
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *InstanceWeight) SetInstanceType(v *string) *InstanceWeight {
	if o.InstanceType = v; o.InstanceType == nil {
		o.nullFields = append(o.nullFields, "InstanceType")
	}
	return o
}

func (o *InstanceWeight) SetWeightedCapacity(v *int) *InstanceWeight {
	if o.WeightedCapacity = v; o.WeightedCapacity == nil {
		o.nullFields = append(o.nullFields, "WeightedCapacity")
	}
	return o
}

// endregion

//region AvailabilityZone

func (o AvailabilityZone) MarshalJSON() ([]byte, error) {
	type noMethod AvailabilityZone
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AvailabilityZone) SetName(v *string) *AvailabilityZone {
	if o.Name = v; v == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *AvailabilityZone) SetSubnetId(v *string) *AvailabilityZone {
	if o.SubnetID = v; v == nil {
		o.nullFields = append(o.nullFields, "SubnetID")
	}
	return o
}

//endregion

//region Tag

func (o Tag) MarshalJSON() ([]byte, error) {
	type noMethod Tag
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Tag) SetKey(v *string) *Tag {
	if o.Key = v; v == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

func (o *Tag) SetValue(v *string) *Tag {
	if o.Value = v; v == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

//endregion

//region InstanceGroups

func (o InstanceGroups) MarshalJSON() ([]byte, error) {
	type noMethod InstanceGroups
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *InstanceGroups) SetMasterGroup(v *InstanceGroup) *InstanceGroups {
	if o.MasterGroup = v; v == nil {
		o.nullFields = append(o.nullFields, "MasterGroup")
	}
	return o
}

func (o *InstanceGroups) SetCoreGroup(v *InstanceGroup) *InstanceGroups {
	if o.CoreGroup = v; v == nil {
		o.nullFields = append(o.nullFields, "CoreGroup")
	}
	return o
}

func (o *InstanceGroups) SetTaskGroup(v *InstanceGroup) *InstanceGroups {
	if o.TaskGroup = v; v == nil {
		o.nullFields = append(o.nullFields, "TaskGroup")
	}
	return o
}

//endregion

//region InstanceGroup

func (o InstanceGroup) MarshalJSON() ([]byte, error) {
	type noMethod InstanceGroup
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *InstanceGroup) SetInstanceTypes(v []string) *InstanceGroup {
	if o.InstanceTypes = v; v == nil {
		o.nullFields = append(o.nullFields, "InstanceTypes")
	}
	return o
}

func (o *InstanceGroup) SetTarget(v *int) *InstanceGroup {
	if o.Target = v; v == nil {
		o.nullFields = append(o.nullFields, "Target")
	}
	return o
}

func (o *InstanceGroup) SetCapacity(v *InstanceGroupCapacity) *InstanceGroup {
	if o.Capacity = v; v == nil {
		o.nullFields = append(o.nullFields, "Capacity")
	}
	return o
}

func (o *InstanceGroup) SetLifeCycle(v *string) *InstanceGroup {
	if o.LifeCycle = v; v == nil {
		o.nullFields = append(o.nullFields, "LifeCycle")
	}
	return o
}

func (o *InstanceGroup) SetEBSConfiguration(v *EBSConfiguration) *InstanceGroup {
	if o.EBSConfiguration = v; v == nil {
		o.nullFields = append(o.nullFields, "EBSConfiguration")
	}
	return o
}

//endregion

//region InstanceGroupCapacity
func (o InstanceGroupCapacity) MarshalJSON() ([]byte, error) {
	type noMethod InstanceGroupCapacity
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *InstanceGroupCapacity) SetTarget(v *int) *InstanceGroupCapacity {
	if o.Target = v; v == nil {
		o.nullFields = append(o.nullFields, "Target")
	}
	return o
}

func (o *InstanceGroupCapacity) SetMinimum(v *int) *InstanceGroupCapacity {
	if o.Minimum = v; v == nil {
		o.nullFields = append(o.nullFields, "Minimum")
	}
	return o
}

func (o *InstanceGroupCapacity) SetMaximum(v *int) *InstanceGroupCapacity {
	if o.Maximum = v; v == nil {
		o.nullFields = append(o.nullFields, "Maximum")
	}
	return o
}

func (o *InstanceGroupCapacity) SetUnit(v *string) *InstanceGroupCapacity {
	if o.Unit = v; o.Unit == nil {
		o.nullFields = append(o.nullFields, "Unit")
	}
	return o
}

//endregion

//region EBSConfiguration
func (o EBSConfiguration) MarshalJSON() ([]byte, error) {
	type noMethod EBSConfiguration
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *EBSConfiguration) SetOptimized(v *bool) *EBSConfiguration {
	if o.Optimized = v; v == nil {
		o.nullFields = append(o.nullFields, "Optimized")
	}
	return o
}

func (o *EBSConfiguration) SetBlockDeviceConfigs(v []*BlockDeviceConfig) *EBSConfiguration {
	if o.BlockDeviceConfigs = v; v == nil {
		o.nullFields = append(o.nullFields, "BlockDeviceConfigs")
	}
	return o
}

//endregion

//region BlockDeviceConfig
func (o BlockDeviceConfig) MarshalJSON() ([]byte, error) {
	type noMethod BlockDeviceConfig
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *BlockDeviceConfig) SetVolumesPerInstance(v *int) *BlockDeviceConfig {
	if o.VolumesPerInstance = v; v == nil {
		o.nullFields = append(o.nullFields, "VolumesPerInstance")
	}
	return o
}

func (o *BlockDeviceConfig) SetVolumeSpecification(v *VolumeSpecification) *BlockDeviceConfig {
	if o.VolumeSpecification = v; v == nil {
		o.nullFields = append(o.nullFields, "VolumeSpecification")
	}
	return o
}

//endregion

//region VolumeSpecification
func (o VolumeSpecification) MarshalJSON() ([]byte, error) {
	type noMethod VolumeSpecification
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *VolumeSpecification) SetVolumeType(v *string) *VolumeSpecification {
	if o.VolumeType = v; v == nil {
		o.nullFields = append(o.nullFields, "VolumeType")
	}
	return o
}

func (o *VolumeSpecification) SetSizeInGB(v *int) *VolumeSpecification {
	if o.SizeInGB = v; v == nil {
		o.nullFields = append(o.nullFields, "SizeInGB")
	}
	return o
}

func (o *VolumeSpecification) SetIOPS(v *int) *VolumeSpecification {
	if o.IOPS = v; v == nil {
		o.nullFields = append(o.nullFields, "IOPS")
	}
	return o
}

//endregion

//region Scaling

func (o Scaling) MarshalJSON() ([]byte, error) {
	type noMethod Scaling
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Scaling) SetUp(v []*ScalingPolicy) *Scaling {
	if o.Up = v; v == nil {
		o.nullFields = append(o.nullFields, "Up")
	} else if len(o.Down) == 0 {
		o.forceSendFields = append(o.forceSendFields, "Up")
	}
	return o
}

func (o *Scaling) SetDown(v []*ScalingPolicy) *Scaling {
	if o.Down = v; v == nil {
		o.nullFields = append(o.nullFields, "Down")
	} else if len(o.Down) == 0 {
		o.forceSendFields = append(o.forceSendFields, "Down")
	}
	return o
}

//endregion

//region ScalingPolicy

func (o ScalingPolicy) MarshalJSON() ([]byte, error) {
	type noMethod ScalingPolicy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ScalingPolicy) SetPolicyName(v *string) *ScalingPolicy {
	if o.PolicyName = v; v == nil {
		o.nullFields = append(o.nullFields, "PolicyName")
	}
	return o
}

func (o *ScalingPolicy) SetNamespace(v *string) *ScalingPolicy {
	if o.Namespace = v; v == nil {
		o.nullFields = append(o.nullFields, "Namespace")
	}
	return o
}

func (o *ScalingPolicy) SetMetricName(v *string) *ScalingPolicy {
	if o.MetricName = v; v == nil {
		o.nullFields = append(o.nullFields, "MetricName")
	}
	return o
}

func (o *ScalingPolicy) SetDimensions(v []*Dimension) *ScalingPolicy {
	if o.Dimensions = v; v == nil {
		o.nullFields = append(o.nullFields, "Dimensions")
	}
	return o
}

func (o *ScalingPolicy) SetStatistic(v *string) *ScalingPolicy {
	if o.Statistic = v; v == nil {
		o.nullFields = append(o.nullFields, "Statistic")
	}
	return o
}

func (o *ScalingPolicy) SetUnit(v *string) *ScalingPolicy {
	if o.Unit = v; v == nil {
		o.nullFields = append(o.nullFields, "Unit")
	}
	return o
}

func (o *ScalingPolicy) SetThreshold(v *float64) *ScalingPolicy {
	if o.Threshold = v; v == nil {
		o.nullFields = append(o.nullFields, "Threshold")
	}
	return o
}

func (o *ScalingPolicy) SetPeriod(v *int) *ScalingPolicy {
	if o.Period = v; v == nil {
		o.nullFields = append(o.nullFields, "Period")
	}
	return o
}

func (o *ScalingPolicy) SetEvaluationPeriods(v *int) *ScalingPolicy {
	if o.EvaluationPeriods = v; v == nil {
		o.nullFields = append(o.nullFields, "EvaluationPeriods")
	}
	return o
}

func (o *ScalingPolicy) SetCooldown(v *int) *ScalingPolicy {
	if o.Cooldown = v; v == nil {
		o.nullFields = append(o.nullFields, "Cooldown")
	}
	return o
}

func (o *ScalingPolicy) SetAction(v *Action) *ScalingPolicy {
	if o.Action = v; v == nil {
		o.nullFields = append(o.nullFields, "Action")
	}
	return o
}

func (o *ScalingPolicy) SetOperator(v *string) *ScalingPolicy {
	if o.Operator = v; v == nil {
		o.nullFields = append(o.nullFields, "Operator")
	}
	return o
}

//endregion

//region Action

func (o Action) MarshalJSON() ([]byte, error) {
	type noMethod Action
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Action) SetType(v *string) *Action {
	if o.Type = v; v == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *Action) SetAdjustment(v *string) *Action {
	if o.Adjustment = v; v == nil {
		o.nullFields = append(o.nullFields, "Adjustment")
	}
	return o
}

func (o *Action) SetMinTargetCapacity(v *string) *Action {
	if o.MinTargetCapacity = v; v == nil {
		o.nullFields = append(o.nullFields, "MinTargetCapacity")
	}
	return o
}

func (o *Action) SetMaxTargetCapacity(v *string) *Action {
	if o.MaxTargetCapacity = v; v == nil {
		o.nullFields = append(o.nullFields, "MaxTargetCapacity")
	}
	return o
}

func (o *Action) SetTarget(v *string) *Action {
	if o.Target = v; v == nil {
		o.nullFields = append(o.nullFields, "Target")
	}
	return o
}

func (o *Action) SetMinimum(v *string) *Action {
	if o.Minimum = v; v == nil {
		o.nullFields = append(o.nullFields, "Minimum")
	}
	return o
}

func (o *Action) SetMaximum(v *string) *Action {
	if o.Maximum = v; v == nil {
		o.nullFields = append(o.nullFields, "Maximum")
	}
	return o
}

//endregion

//region Dimension

func (o Dimension) MarshalJSON() ([]byte, error) {
	type noMethod Dimension
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Dimension) SetName(v *string) *Dimension {
	if o.Name = v; v == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Dimension) SetValue(v *string) *Dimension {
	if o.Value = v; v == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

//endregion

//region Configurations

func (o Configurations) MarshalJSON() ([]byte, error) {
	type noMethod Configurations
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Configurations) SetFile(v *S3File) *Configurations {
	if o.File = v; v == nil {
		o.nullFields = append(o.nullFields, "File")
	}
	return o
}

//endregion

//region Bootstrap Actions

func (o BootstrapActions) MarshalJSON() ([]byte, error) {
	type noMethod BootstrapActions
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *BootstrapActions) SetFile(v *S3File) *BootstrapActions {
	if o.File = v; v == nil {
		o.nullFields = append(o.nullFields, "File")
	}
	return o
}

//endregion

//region Steps

func (o Steps) MarshalJSON() ([]byte, error) {
	type noMethod Steps
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Steps) SetFile(v *S3File) *Steps {
	if o.File = v; v == nil {
		o.nullFields = append(o.nullFields, "File")
	}
	return o
}

//endregion

//region S3File
func (o S3File) MarshalJSON() ([]byte, error) {
	type noMethod S3File
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *S3File) SetBucket(v *string) *S3File {
	if o.Bucket = v; v == nil {
		o.nullFields = append(o.nullFields, "Bucket")
	}
	return o
}

func (o *S3File) SetKey(v *string) *S3File {
	if o.Key = v; v == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

//endregion

func (s *ServiceOp) List(ctx context.Context, input *ListScalersInput) (*ListScalersOutput, error) {
	r := client.NewRequest(http.MethodGet, "/aws/emr/mrScaler")
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := scalersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListScalersOutput{Scalers: gs}, nil
}

func (s *ServiceOp) Create(ctx context.Context, input *CreateScalerInput) (*CreateScalerOutput, error) {
	r := client.NewRequest(http.MethodPost, "/aws/emr/mrScaler")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := scalersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateScalerOutput)
	if len(gs) > 0 {
		output.Scaler = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) Read(ctx context.Context, input *ReadScalerInput) (*ReadScalerOutput, error) {
	path, err := uritemplates.Expand("/aws/emr/mrScaler/{mrScalerId}", uritemplates.Values{
		"mrScalerId": spotinst.StringValue(input.ScalerID),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := scalersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadScalerOutput)
	if len(gs) > 0 {
		output.Scaler = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadScalerCluster(ctx context.Context, input *ScalerClusterStatusInput) (*ScalerClusterStatusOutput, error) {
	path, err := uritemplates.Expand("/aws/emr/mrScaler/{mrScalerId}/cluster", uritemplates.Values{
		"mrScalerId": spotinst.StringValue(input.ScalerID),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := scalerClustersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ScalerClusterStatusOutput)
	if len(gs) > 0 {
		output.ScalerClusterId = gs[0].ScalerClusterId
	}

	return output, nil
}

func (s *ServiceOp) Update(ctx context.Context, input *UpdateScalerInput) (*UpdateScalerOutput, error) {
	path, err := uritemplates.Expand("/aws/emr/mrScaler/{mrScalerId}", uritemplates.Values{
		"mrScalerId": spotinst.StringValue(input.Scaler.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.Scaler.ID = nil

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := scalersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateScalerOutput)
	if len(gs) > 0 {
		output.Scaler = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) Delete(ctx context.Context, input *DeleteScalerInput) (*DeleteScalerOutput, error) {
	path, err := uritemplates.Expand("/aws/emr/mrScaler/{mrScalerId}", uritemplates.Values{
		"mrScalerId": spotinst.StringValue(input.ScalerID),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodDelete, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteScalerOutput{}, nil
}
