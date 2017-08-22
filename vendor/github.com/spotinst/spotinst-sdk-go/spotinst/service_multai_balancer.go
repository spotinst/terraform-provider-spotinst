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

// A Protocol represents the type of an application protocol.
type Protocol int

const (
	// ProtocolHTTP represents the Hypertext Transfer Protocol (HTTP) protocol.
	ProtocolHTTP Protocol = iota

	// ProtocolHTTPS represents the Hypertext Transfer Protocol (HTTP) within
	// a connection encrypted by Transport Layer Security, or its predecessor,
	// Secure Sockets Layer.
	ProtocolHTTPS

	// ProtocolHTTP2 represents the Hypertext Transfer Protocol (HTTP) protocol
	// version 2.
	ProtocolHTTP2
)

var Protocol_name = map[Protocol]string{
	ProtocolHTTP:  "HTTP",
	ProtocolHTTPS: "HTTPS",
	ProtocolHTTP2: "HTTP2",
}

var Protocol_value = map[string]Protocol{
	"HTTP":  ProtocolHTTP,
	"HTTPS": ProtocolHTTPS,
	"HTTP2": ProtocolHTTP2,
}

func (p Protocol) String() string {
	return Protocol_name[p]
}

// A ReadinessStatus represents the readiness status of a target.
type ReadinessStatus int

const (
	// StatusReady represents a ready state.
	StatusReady ReadinessStatus = iota

	// StatusMaintenance represents a maintenance state.
	StatusMaintenance
)

var ReadinessStatus_name = map[ReadinessStatus]string{
	StatusReady:       "READY",
	StatusMaintenance: "MAINTENANCE",
}

var ReadinessStatus_value = map[string]ReadinessStatus{
	"READY":       StatusReady,
	"MAINTENANCE": StatusMaintenance,
}

func (s ReadinessStatus) String() string {
	return ReadinessStatus_name[s]
}

// A HealthinessStatus represents the healthiness status of a target.
type HealthinessStatus int

const (
	// StatusUnknown represents an unknown state.
	StatusUnknown HealthinessStatus = iota

	// StatusHealthy represents a healthy state.
	StatusHealthy

	// StatusUnhealthy represents an unhealthy state.
	StatusUnhealthy
)

var HealthinessStatus_name = map[HealthinessStatus]string{
	StatusUnknown:   "UNKNOWN",
	StatusHealthy:   "HEALTHY",
	StatusUnhealthy: "UNHEALTHY",
}

var HealthinessStatus_value = map[string]HealthinessStatus{
	"UNKNOWN":   StatusUnknown,
	"HEALTHY":   StatusHealthy,
	"UNHEALTHY": StatusUnhealthy,
}

func (s HealthinessStatus) String() string {
	return HealthinessStatus_name[s]
}

// A Strategy represents the load balancing methods used to determine which
// application server to send a request to.
type Strategy int

const (
	// StrategyRandom represents a random load balancing method where
	// a request is passed to the server with the least number of
	// active connections.
	StrategyRandom Strategy = iota

	// StrategyRoundRobin represents a random load balancing method where
	// a request is passed to the server in round-robin fashion.
	StrategyRoundRobin

	// StrategyLeastConn represents a random load balancing method where
	// a request is passed to the server with the least number of
	// active connections.
	StrategyLeastConn

	// StrategyIPHash represents a IP hash load balancing method where
	// a request is passed to the server based on the result of hashing
	// the request IP address.
	StrategyIPHash
)

var Strategy_name = map[Strategy]string{
	StrategyRandom:     "RANDOM",
	StrategyRoundRobin: "ROUNDROBIN",
	StrategyLeastConn:  "LEASTCONN",
	StrategyIPHash:     "IPHASH",
}

var Strategy_value = map[string]Strategy{
	"RANDOM":     StrategyRandom,
	"ROUNDROBIN": StrategyRoundRobin,
	"LEASTCONN":  StrategyLeastConn,
	"IPHASH":     StrategyIPHash,
}

func (s Strategy) String() string {
	return Strategy_name[s]
}

// BalancerService is an interface for interfacing with the balancer
// targets of the Spotinst API.
type BalancerService interface {
	ListBalancers(context.Context, *ListBalancersInput) (*ListBalancersOutput, error)
	CreateBalancer(context.Context, *CreateBalancerInput) (*CreateBalancerOutput, error)
	ReadBalancer(context.Context, *ReadBalancerInput) (*ReadBalancerOutput, error)
	UpdateBalancer(context.Context, *UpdateBalancerInput) (*UpdateBalancerOutput, error)
	DeleteBalancer(context.Context, *DeleteBalancerInput) (*DeleteBalancerOutput, error)

	ListListeners(context.Context, *ListListenersInput) (*ListListenersOutput, error)
	CreateListener(context.Context, *CreateListenerInput) (*CreateListenerOutput, error)
	ReadListener(context.Context, *ReadListenerInput) (*ReadListenerOutput, error)
	UpdateListener(context.Context, *UpdateListenerInput) (*UpdateListenerOutput, error)
	DeleteListener(context.Context, *DeleteListenerInput) (*DeleteListenerOutput, error)

	ListRoutingRules(context.Context, *ListRoutingRulesInput) (*ListRoutingRulesOutput, error)
	CreateRoutingRule(context.Context, *CreateRoutingRuleInput) (*CreateRoutingRuleOutput, error)
	ReadRoutingRule(context.Context, *ReadRoutingRuleInput) (*ReadRoutingRuleOutput, error)
	UpdateRoutingRule(context.Context, *UpdateRoutingRuleInput) (*UpdateRoutingRuleOutput, error)
	DeleteRoutingRule(context.Context, *DeleteRoutingRuleInput) (*DeleteRoutingRuleOutput, error)

	ListMiddlewares(context.Context, *ListMiddlewaresInput) (*ListMiddlewaresOutput, error)
	CreateMiddleware(context.Context, *CreateMiddlewareInput) (*CreateMiddlewareOutput, error)
	ReadMiddleware(context.Context, *ReadMiddlewareInput) (*ReadMiddlewareOutput, error)
	UpdateMiddleware(context.Context, *UpdateMiddlewareInput) (*UpdateMiddlewareOutput, error)
	DeleteMiddleware(context.Context, *DeleteMiddlewareInput) (*DeleteMiddlewareOutput, error)

	ListTargetSets(context.Context, *ListTargetSetsInput) (*ListTargetSetsOutput, error)
	CreateTargetSet(context.Context, *CreateTargetSetInput) (*CreateTargetSetOutput, error)
	ReadTargetSet(context.Context, *ReadTargetSetInput) (*ReadTargetSetOutput, error)
	UpdateTargetSet(context.Context, *UpdateTargetSetInput) (*UpdateTargetSetOutput, error)
	DeleteTargetSet(context.Context, *DeleteTargetSetInput) (*DeleteTargetSetOutput, error)

	ListTargets(context.Context, *ListTargetsInput) (*ListTargetsOutput, error)
	CreateTarget(context.Context, *CreateTargetInput) (*CreateTargetOutput, error)
	ReadTarget(context.Context, *ReadTargetInput) (*ReadTargetOutput, error)
	UpdateTarget(context.Context, *UpdateTargetInput) (*UpdateTargetOutput, error)
	DeleteTarget(context.Context, *DeleteTargetInput) (*DeleteTargetOutput, error)

	ListRuntimes(context.Context, *ListRuntimesInput) (*ListRuntimesOutput, error)
	ReadRuntime(context.Context, *ReadRuntimeInput) (*ReadRuntimeOutput, error)
}

// BalancerServiceOp handles communication with the balancer related methods
// of the Spotinst API.
type BalancerServiceOp struct {
	client *Client
}

var _ BalancerService = &BalancerServiceOp{}

type Balancer struct {
	ID              *string    `json:"id,omitempty"`
	Name            *string    `json:"name,omitempty"`
	DNSRRType       *string    `json:"dnsRrType,omitempty"`
	DNSRRName       *string    `json:"dnsRrName,omitempty"`
	DNSCNAMEAliases []string   `json:"dnsCnameAliases,omitempty"`
	Timeouts        *Timeouts  `json:"timeouts,omitempty"`
	Tags            []*Tag     `json:"tags,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	UpdatedAt       *time.Time `json:"updatedAt,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type Timeouts struct {
	Idle     *int `json:"idle"`
	Draining *int `json:"draining"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListBalancersInput struct {
	DeploymentID *string `json:"deploymentId,omitempty"`
}

type ListBalancersOutput struct {
	Balancers []*Balancer `json:"balancers,omitempty"`
}

type CreateBalancerInput struct {
	Balancer *Balancer `json:"balancer,omitempty"`
}

type CreateBalancerOutput struct {
	Balancer *Balancer `json:"balancer,omitempty"`
}

type ReadBalancerInput struct {
	BalancerID *string `json:"balancerId,omitempty"`
}

type ReadBalancerOutput struct {
	Balancer *Balancer `json:"balancer,omitempty"`
}

type UpdateBalancerInput struct {
	Balancer *Balancer `json:"balancer,omitempty"`
}

type UpdateBalancerOutput struct{}

type DeleteBalancerInput struct {
	BalancerID *string `json:"balancerId,omitempty"`
}

type DeleteBalancerOutput struct{}

func balancerFromJSON(in []byte) (*Balancer, error) {
	b := new(Balancer)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func balancersFromJSON(in []byte) ([]*Balancer, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Balancer, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := balancerFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func balancersFromHttpResponse(resp *http.Response) ([]*Balancer, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return balancersFromJSON(body)
}

func (b *BalancerServiceOp) ListBalancers(ctx context.Context, input *ListBalancersInput) (*ListBalancersOutput, error) {
	r := b.client.newRequest(ctx, "GET", "/loadBalancer/balancer")

	if input.DeploymentID != nil {
		r.params.Set("deploymentId", StringValue(input.DeploymentID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := balancersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListBalancersOutput{Balancers: bs}, nil
}

func (b *BalancerServiceOp) CreateBalancer(ctx context.Context, input *CreateBalancerInput) (*CreateBalancerOutput, error) {
	r := b.client.newRequest(ctx, "POST", "/loadBalancer/balancer")
	r.obj = input

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := balancersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateBalancerOutput)
	if len(bs) > 0 {
		output.Balancer = bs[0]
	}

	return output, nil
}

func (b *BalancerServiceOp) ReadBalancer(ctx context.Context, input *ReadBalancerInput) (*ReadBalancerOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/balancer/{balancerId}", map[string]string{
		"balancerId": StringValue(input.BalancerID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest(ctx, "GET", path)
	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := balancersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadBalancerOutput)
	if len(bs) > 0 {
		output.Balancer = bs[0]
	}

	return output, nil
}

func (b *BalancerServiceOp) UpdateBalancer(ctx context.Context, input *UpdateBalancerInput) (*UpdateBalancerOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/balancer/{balancerId}", map[string]string{
		"balancerId": StringValue(input.Balancer.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.Balancer.ID = nil

	r := b.client.newRequest(ctx, "PUT", path)
	r.obj = input

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateBalancerOutput{}, nil
}

func (b *BalancerServiceOp) DeleteBalancer(ctx context.Context, input *DeleteBalancerInput) (*DeleteBalancerOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/balancer/{balancerId}", map[string]string{
		"balancerId": StringValue(input.BalancerID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest(ctx, "DELETE", path)
	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteBalancerOutput{}, nil
}

// region Balancer

func (o *Balancer) MarshalJSON() ([]byte, error) {
	type noMethod Balancer
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Balancer) SetId(v *string) *Balancer {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *Balancer) SetName(v *string) *Balancer {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Balancer) SetDNSRRType(v *string) *Balancer {
	if o.DNSRRType = v; o.DNSRRType == nil {
		o.nullFields = append(o.nullFields, "DNSRRType")
	}
	return o
}

func (o *Balancer) SetDNSRRName(v *string) *Balancer {
	if o.DNSRRName = v; o.DNSRRName == nil {
		o.nullFields = append(o.nullFields, "DNSRRName")
	}
	return o
}

func (o *Balancer) SetDNSCNAMEAliases(v []string) *Balancer {
	if o.DNSCNAMEAliases = v; o.DNSCNAMEAliases == nil {
		o.nullFields = append(o.nullFields, "DNSCNAMEAliases")
	}
	return o
}

func (o *Balancer) SetTimeouts(v *Timeouts) *Balancer {
	if o.Timeouts = v; o.Timeouts == nil {
		o.nullFields = append(o.nullFields, "Timeouts")
	}
	return o
}

func (o *Balancer) SetTags(v []*Tag) *Balancer {
	if o.Tags = v; o.Tags == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

func (o *Balancer) SetCreatedAt(v *time.Time) *Balancer {
	if o.CreatedAt = v; o.CreatedAt == nil {
		o.nullFields = append(o.nullFields, "CreatedAt")
	}
	return o
}

func (o *Balancer) SetUpdatedAt(v *time.Time) *Balancer {
	if o.UpdatedAt = v; o.UpdatedAt == nil {
		o.nullFields = append(o.nullFields, "UpdatedAt")
	}
	return o
}

// endregion

type Listener struct {
	ID         *string    `json:"id,omitempty"`
	BalancerID *string    `json:"balancerId,omitempty"`
	Protocol   *string    `json:"protocol,omitempty"`
	Port       *int       `json:"port,omitempty"`
	TLSConfig  *TLSConfig `json:"tlsConfig,omitempty"`
	Tags       []*Tag     `json:"tags,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListListenersInput struct {
	BalancerID *string `json:"balancerId,omitempty"`
}

type ListListenersOutput struct {
	Listeners []*Listener `json:"listeners,omitempty"`
}

type CreateListenerInput struct {
	Listener *Listener `json:"listener,omitempty"`
}

type CreateListenerOutput struct {
	Listener *Listener `json:"listener,omitempty"`
}

type ReadListenerInput struct {
	ListenerID *string `json:"listenerId,omitempty"`
}

type ReadListenerOutput struct {
	Listener *Listener `json:"listener,omitempty"`
}

type UpdateListenerInput struct {
	Listener *Listener `json:"listener,omitempty"`
}

type UpdateListenerOutput struct{}

type DeleteListenerInput struct {
	ListenerID *string `json:"listenerId,omitempty"`
}

type DeleteListenerOutput struct{}

func listenerFromJSON(in []byte) (*Listener, error) {
	b := new(Listener)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func listenersFromJSON(in []byte) ([]*Listener, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Listener, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rl := range rw.Response.Items {
		l, err := listenerFromJSON(rl)
		if err != nil {
			return nil, err
		}
		out[i] = l
	}
	return out, nil
}

func listenersFromHttpResponse(resp *http.Response) ([]*Listener, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return listenersFromJSON(body)
}

func (b *BalancerServiceOp) ListListeners(ctx context.Context, input *ListListenersInput) (*ListListenersOutput, error) {
	r := b.client.newRequest(ctx, "GET", "/loadBalancer/listener")

	if input.BalancerID != nil {
		r.params.Set("balancerId", StringValue(input.BalancerID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ls, err := listenersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListListenersOutput{Listeners: ls}, nil
}

func (b *BalancerServiceOp) CreateListener(ctx context.Context, input *CreateListenerInput) (*CreateListenerOutput, error) {
	r := b.client.newRequest(ctx, "POST", "/loadBalancer/listener")
	r.obj = input

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ls, err := listenersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateListenerOutput)
	if len(ls) > 0 {
		output.Listener = ls[0]
	}

	return output, nil
}

func (b *BalancerServiceOp) ReadListener(ctx context.Context, input *ReadListenerInput) (*ReadListenerOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/listener/{listenerId}", map[string]string{
		"listenerId": StringValue(input.ListenerID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest(ctx, "GET", path)
	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ls, err := listenersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadListenerOutput)
	if len(ls) > 0 {
		output.Listener = ls[0]
	}

	return output, nil
}

func (b *BalancerServiceOp) UpdateListener(ctx context.Context, input *UpdateListenerInput) (*UpdateListenerOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/listener/{listenerId}", map[string]string{
		"listenerId": StringValue(input.Listener.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.Listener.ID = nil

	r := b.client.newRequest(ctx, "PUT", path)
	r.obj = input

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateListenerOutput{}, nil
}

func (b *BalancerServiceOp) DeleteListener(ctx context.Context, input *DeleteListenerInput) (*DeleteListenerOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/listener/{listenerId}", map[string]string{
		"listenerId": StringValue(input.ListenerID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest(ctx, "DELETE", path)
	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteListenerOutput{}, nil
}

// region Listener

func (o *Listener) MarshalJSON() ([]byte, error) {
	type noMethod Listener
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Listener) SetId(v *string) *Listener {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *Listener) SetBalancerId(v *string) *Listener {
	if o.BalancerID = v; o.BalancerID == nil {
		o.nullFields = append(o.nullFields, "BalancerID")
	}
	return o
}

func (o *Listener) SetProtocol(v *string) *Listener {
	if o.Protocol = v; o.Protocol == nil {
		o.nullFields = append(o.nullFields, "Protocol")
	}
	return o
}

func (o *Listener) SetPort(v *int) *Listener {
	if o.Port = v; o.Port == nil {
		o.nullFields = append(o.nullFields, "Port")
	}
	return o
}

func (o *Listener) SetTLSConfig(v *TLSConfig) *Listener {
	if o.TLSConfig = v; o.TLSConfig == nil {
		o.nullFields = append(o.nullFields, "TLSConfig")
	}
	return o
}

func (o *Listener) SetTags(v []*Tag) *Listener {
	if o.Tags = v; o.Tags == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

// endregion

// region TLSConfig

func (o *TLSConfig) MarshalJSON() ([]byte, error) {
	type noMethod TLSConfig
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *TLSConfig) SetCertificateIDs(v []string) *TLSConfig {
	if o.CertificateIDs = v; o.CertificateIDs == nil {
		o.nullFields = append(o.nullFields, "CertificateIDs")
	}
	return o
}

func (o *TLSConfig) SetMinVersion(v *string) *TLSConfig {
	if o.MinVersion = v; o.MinVersion == nil {
		o.nullFields = append(o.nullFields, "MinVersion")
	}
	return o
}

func (o *TLSConfig) SetMaxVersion(v *string) *TLSConfig {
	if o.MaxVersion = v; o.MaxVersion == nil {
		o.nullFields = append(o.nullFields, "MaxVersion")
	}
	return o
}

func (o *TLSConfig) SetSessionTicketsDisabled(v *bool) *TLSConfig {
	if o.SessionTicketsDisabled = v; o.SessionTicketsDisabled == nil {
		o.nullFields = append(o.nullFields, "SessionTicketsDisabled")
	}
	return o
}

func (o *TLSConfig) SetPreferServerCipherSuites(v *bool) *TLSConfig {
	if o.PreferServerCipherSuites = v; o.PreferServerCipherSuites == nil {
		o.nullFields = append(o.nullFields, "PreferServerCipherSuites")
	}
	return o
}

func (o *TLSConfig) SetCipherSuites(v []string) *TLSConfig {
	if o.CipherSuites = v; o.CipherSuites == nil {
		o.nullFields = append(o.nullFields, "CipherSuites")
	}
	return o
}

func (o *TLSConfig) SetInsecureSkipVerify(v *bool) *TLSConfig {
	if o.InsecureSkipVerify = v; o.InsecureSkipVerify == nil {
		o.nullFields = append(o.nullFields, "InsecureSkipVerify")
	}
	return o
}

// endregion

type RoutingRule struct {
	ID            *string    `json:"id,omitempty"`
	BalancerID    *string    `json:"balancerId,omitempty"`
	ListenerID    *string    `json:"listenerId,omitempty"`
	MiddlewareIDs []string   `json:"middlewareIds,omitempty"`
	TargetSetIDs  []string   `json:"targetSetIds,omitempty"`
	Priority      *int       `json:"priority,omitempty"`
	Strategy      *string    `json:"strategy,omitempty"`
	Route         *string    `json:"route,omitempty"`
	Tags          []*Tag     `json:"tags,omitempty"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListRoutingRulesInput struct {
	BalancerID *string `json:"balancerId,omitempty"`
}

type ListRoutingRulesOutput struct {
	RoutingRules []*RoutingRule `json:"routingRules,omitempty"`
}

type CreateRoutingRuleInput struct {
	RoutingRule *RoutingRule `json:"routingRule,omitempty"`
}

type CreateRoutingRuleOutput struct {
	RoutingRule *RoutingRule `json:"routingRule,omitempty"`
}

type ReadRoutingRuleInput struct {
	RoutingRuleID *string `json:"routingRuleId,omitempty"`
}

type ReadRoutingRuleOutput struct {
	RoutingRule *RoutingRule `json:"routingRule,omitempty"`
}

type UpdateRoutingRuleInput struct {
	RoutingRule *RoutingRule `json:"routingRule,omitempty"`
}

type UpdateRoutingRuleOutput struct{}

type DeleteRoutingRuleInput struct {
	RoutingRuleID *string `json:"routingRuleId,omitempty"`
}

type DeleteRoutingRuleOutput struct{}

func routingRuleFromJSON(in []byte) (*RoutingRule, error) {
	rr := new(RoutingRule)
	if err := json.Unmarshal(in, rr); err != nil {
		return nil, err
	}
	return rr, nil
}

func routingRulesFromJSON(in []byte) ([]*RoutingRule, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*RoutingRule, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rr := range rw.Response.Items {
		r, err := routingRuleFromJSON(rr)
		if err != nil {
			return nil, err
		}
		out[i] = r
	}
	return out, nil
}

func routingRulesFromHttpResponse(resp *http.Response) ([]*RoutingRule, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return routingRulesFromJSON(body)
}

func (b *BalancerServiceOp) ListRoutingRules(ctx context.Context, input *ListRoutingRulesInput) (*ListRoutingRulesOutput, error) {
	r := b.client.newRequest(ctx, "GET", "/loadBalancer/routingRule")

	if input.BalancerID != nil {
		r.params.Set("balancerId", StringValue(input.BalancerID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rr, err := routingRulesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListRoutingRulesOutput{RoutingRules: rr}, nil
}

func (b *BalancerServiceOp) CreateRoutingRule(ctx context.Context, input *CreateRoutingRuleInput) (*CreateRoutingRuleOutput, error) {
	r := b.client.newRequest(ctx, "POST", "/loadBalancer/routingRule")
	r.obj = input

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rr, err := routingRulesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateRoutingRuleOutput)
	if len(rr) > 0 {
		output.RoutingRule = rr[0]
	}

	return output, nil
}

func (b *BalancerServiceOp) ReadRoutingRule(ctx context.Context, input *ReadRoutingRuleInput) (*ReadRoutingRuleOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/routingRule/{routingRuleId}", map[string]string{
		"routingRuleId": StringValue(input.RoutingRuleID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest(ctx, "GET", path)
	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rr, err := routingRulesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadRoutingRuleOutput)
	if len(rr) > 0 {
		output.RoutingRule = rr[0]
	}

	return output, nil
}

func (b *BalancerServiceOp) UpdateRoutingRule(ctx context.Context, input *UpdateRoutingRuleInput) (*UpdateRoutingRuleOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/routingRule/{routingRuleId}", map[string]string{
		"routingRuleId": StringValue(input.RoutingRule.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.RoutingRule.ID = nil

	r := b.client.newRequest(ctx, "PUT", path)
	r.obj = input

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateRoutingRuleOutput{}, nil
}

func (b *BalancerServiceOp) DeleteRoutingRule(ctx context.Context, input *DeleteRoutingRuleInput) (*DeleteRoutingRuleOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/routingRule/{routingRuleId}", map[string]string{
		"routingRuleId": StringValue(input.RoutingRuleID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest(ctx, "DELETE", path)
	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteRoutingRuleOutput{}, nil
}

// region RoutingRule

func (o *RoutingRule) MarshalJSON() ([]byte, error) {
	type noMethod RoutingRule
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RoutingRule) SetId(v *string) *RoutingRule {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *RoutingRule) SetBalancerId(v *string) *RoutingRule {
	if o.BalancerID = v; o.BalancerID == nil {
		o.nullFields = append(o.nullFields, "BalancerID")
	}
	return o
}

func (o *RoutingRule) SetListenerId(v *string) *RoutingRule {
	if o.ListenerID = v; o.ListenerID == nil {
		o.nullFields = append(o.nullFields, "ListenerID")
	}
	return o
}

func (o *RoutingRule) SetMiddlewareIDs(v []string) *RoutingRule {
	if o.MiddlewareIDs = v; o.MiddlewareIDs == nil {
		o.nullFields = append(o.nullFields, "MiddlewareIDs")
	}
	return o
}

func (o *RoutingRule) SetTargetSetIDs(v []string) *RoutingRule {
	if o.TargetSetIDs = v; o.TargetSetIDs == nil {
		o.nullFields = append(o.nullFields, "TargetSetIDs")
	}
	return o
}

func (o *RoutingRule) SetPriority(v *int) *RoutingRule {
	if o.Priority = v; o.Priority == nil {
		o.nullFields = append(o.nullFields, "Priority")
	}
	return o
}

func (o *RoutingRule) SetStrategy(v *string) *RoutingRule {
	if o.Strategy = v; o.Strategy == nil {
		o.nullFields = append(o.nullFields, "Strategy")
	}
	return o
}

func (o *RoutingRule) SetRoute(v *string) *RoutingRule {
	if o.Route = v; o.Route == nil {
		o.nullFields = append(o.nullFields, "Route")
	}
	return o
}

func (o *RoutingRule) SetTags(v []*Tag) *RoutingRule {
	if o.Tags = v; o.Tags == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

func (o *RoutingRule) SetCreatedAt(v *time.Time) *RoutingRule {
	if o.CreatedAt = v; o.CreatedAt == nil {
		o.nullFields = append(o.nullFields, "CreatedAt")
	}
	return o
}

func (o *RoutingRule) SetUpdatedAt(v *time.Time) *RoutingRule {
	if o.UpdatedAt = v; o.UpdatedAt == nil {
		o.nullFields = append(o.nullFields, "UpdatedAt")
	}
	return o
}

// endregion

type Middleware struct {
	ID         *string         `json:"id,omitempty"`
	BalancerID *string         `json:"balancerId,omitempty"`
	Type       *string         `json:"type,omitempty"`
	Priority   *int            `json:"priority,omitempty"`
	Spec       json.RawMessage `json:"spec,omitempty"`
	Tags       []*Tag          `json:"tags,omitempty"`
	CreatedAt  *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt  *time.Time      `json:"updatedAt,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListMiddlewaresInput struct {
	BalancerID *string `json:"balancerId,omitempty"`
}

type ListMiddlewaresOutput struct {
	Middlewares []*Middleware `json:"middlewares,omitempty"`
}

type CreateMiddlewareInput struct {
	Middleware *Middleware `json:"middleware,omitempty"`
}

type CreateMiddlewareOutput struct {
	Middleware *Middleware `json:"middleware,omitempty"`
}

type ReadMiddlewareInput struct {
	MiddlewareID *string `json:"middlewareId,omitempty"`
}

type ReadMiddlewareOutput struct {
	Middleware *Middleware `json:"middleware,omitempty"`
}

type UpdateMiddlewareInput struct {
	Middleware *Middleware `json:"middleware,omitempty"`
}

type UpdateMiddlewareOutput struct{}

type DeleteMiddlewareInput struct {
	MiddlewareID *string `json:"middlewareId,omitempty"`
}

type DeleteMiddlewareOutput struct{}

func middlewareFromJSON(in []byte) (*Middleware, error) {
	m := new(Middleware)
	if err := json.Unmarshal(in, m); err != nil {
		return nil, err
	}
	return m, nil
}

func middlewaresFromJSON(in []byte) ([]*Middleware, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Middleware, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rm := range rw.Response.Items {
		m, err := middlewareFromJSON(rm)
		if err != nil {
			return nil, err
		}
		out[i] = m
	}
	return out, nil
}

func middlewaresFromHttpResponse(resp *http.Response) ([]*Middleware, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return middlewaresFromJSON(body)
}

func (b *BalancerServiceOp) ListMiddlewares(ctx context.Context, input *ListMiddlewaresInput) (*ListMiddlewaresOutput, error) {
	r := b.client.newRequest(ctx, "GET", "/loadBalancer/middleware")

	if input.BalancerID != nil {
		r.params.Set("balancerId", StringValue(input.BalancerID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ms, err := middlewaresFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListMiddlewaresOutput{Middlewares: ms}, nil
}

func (b *BalancerServiceOp) CreateMiddleware(ctx context.Context, input *CreateMiddlewareInput) (*CreateMiddlewareOutput, error) {
	r := b.client.newRequest(ctx, "POST", "/loadBalancer/middleware")
	r.obj = input

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ms, err := middlewaresFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateMiddlewareOutput)
	if len(ms) > 0 {
		output.Middleware = ms[0]
	}

	return output, nil
}

func (b *BalancerServiceOp) ReadMiddleware(ctx context.Context, input *ReadMiddlewareInput) (*ReadMiddlewareOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/middleware/{middlewareId}", map[string]string{
		"middlewareId": StringValue(input.MiddlewareID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest(ctx, "GET", path)
	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ms, err := middlewaresFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadMiddlewareOutput)
	if len(ms) > 0 {
		output.Middleware = ms[0]
	}

	return output, nil
}

func (b *BalancerServiceOp) UpdateMiddleware(ctx context.Context, input *UpdateMiddlewareInput) (*UpdateMiddlewareOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/middleware/{middlewareId}", map[string]string{
		"middlewareId": StringValue(input.Middleware.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.Middleware.ID = nil

	r := b.client.newRequest(ctx, "PUT", path)
	r.obj = input

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateMiddlewareOutput{}, nil
}

func (b *BalancerServiceOp) DeleteMiddleware(ctx context.Context, input *DeleteMiddlewareInput) (*DeleteMiddlewareOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/middleware/{middlewareId}", map[string]string{
		"middlewareId": StringValue(input.MiddlewareID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest(ctx, "DELETE", path)
	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteMiddlewareOutput{}, nil
}

// region Middleware

func (o *Middleware) MarshalJSON() ([]byte, error) {
	type noMethod Middleware
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Middleware) SetId(v *string) *Middleware {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *Middleware) SetBalancerId(v *string) *Middleware {
	if o.BalancerID = v; o.BalancerID == nil {
		o.nullFields = append(o.nullFields, "BalancerID")
	}
	return o
}

func (o *Middleware) SetType(v *string) *Middleware {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *Middleware) SetPriority(v *int) *Middleware {
	if o.Priority = v; o.Priority == nil {
		o.nullFields = append(o.nullFields, "Priority")
	}
	return o
}

func (o *Middleware) SetSpec(v json.RawMessage) *Middleware {
	if o.Spec = v; o.Spec == nil {
		o.nullFields = append(o.nullFields, "Spec")
	}
	return o
}

func (o *Middleware) SetTags(v []*Tag) *Middleware {
	if o.Tags = v; o.Tags == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

// endregion

type TargetSet struct {
	ID           *string               `json:"id,omitempty"`
	BalancerID   *string               `json:"balancerId,omitempty"`
	DeploymentID *string               `json:"deploymentId,omitempty"`
	Name         *string               `json:"name,omitempty"`
	Protocol     *string               `json:"protocol,omitempty"`
	Port         *int                  `json:"port,omitempty"`
	Weight       *int                  `json:"weight,omitempty"`
	HealthCheck  *TargetSetHealthCheck `json:"healthCheck,omitempty"`
	Tags         []*Tag                `json:"tags,omitempty"`
	CreatedAt    *time.Time            `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time            `json:"updatedAt,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type TargetSetHealthCheck struct {
	Path                    *string `json:"path,omitempty"`
	Port                    *int    `json:"port,omitempty"`
	Protocol                *string `json:"protocol,omitempty"`
	Timeout                 *int    `json:"timeout,omitempty"`
	Interval                *int    `json:"interval,omitempty"`
	HealthyThresholdCount   *int    `json:"healthyThresholdCount,omitempty"`
	UnhealthyThresholdCount *int    `json:"unhealthyThresholdCount,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListTargetSetsInput struct {
	BalancerID *string `json:"balancerId,omitempty"`
}

type ListTargetSetsOutput struct {
	TargetSets []*TargetSet `json:"targetSets,omitempty"`
}

type CreateTargetSetInput struct {
	TargetSet *TargetSet `json:"targetSet,omitempty"`
}

type CreateTargetSetOutput struct {
	TargetSet *TargetSet `json:"targetSet,omitempty"`
}

type ReadTargetSetInput struct {
	TargetSetID *string `json:"targetSetId,omitempty"`
}

type ReadTargetSetOutput struct {
	TargetSet *TargetSet `json:"targetSet,omitempty"`
}

type UpdateTargetSetInput struct {
	TargetSet *TargetSet `json:"targetSet,omitempty"`
}

type UpdateTargetSetOutput struct{}

type DeleteTargetSetInput struct {
	TargetSetID *string `json:"targetSetId,omitempty"`
}

type DeleteTargetSetOutput struct{}

func targetSetFromJSON(in []byte) (*TargetSet, error) {
	ts := new(TargetSet)
	if err := json.Unmarshal(in, ts); err != nil {
		return nil, err
	}
	return ts, nil
}

func targetSetsFromJSON(in []byte) ([]*TargetSet, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*TargetSet, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rts := range rw.Response.Items {
		ts, err := targetSetFromJSON(rts)
		if err != nil {
			return nil, err
		}
		out[i] = ts
	}
	return out, nil
}

func targetSetsFromHttpResponse(resp *http.Response) ([]*TargetSet, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return targetSetsFromJSON(body)
}

func (b *BalancerServiceOp) ListTargetSets(ctx context.Context, input *ListTargetSetsInput) (*ListTargetSetsOutput, error) {
	r := b.client.newRequest(ctx, "GET", "/loadBalancer/targetSet")

	if input.BalancerID != nil {
		r.params.Set("balancerId", StringValue(input.BalancerID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ts, err := targetSetsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListTargetSetsOutput{TargetSets: ts}, nil
}

func (b *BalancerServiceOp) CreateTargetSet(ctx context.Context, input *CreateTargetSetInput) (*CreateTargetSetOutput, error) {
	r := b.client.newRequest(ctx, "POST", "/loadBalancer/targetSet")
	r.obj = input

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ts, err := targetSetsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateTargetSetOutput)
	if len(ts) > 0 {
		output.TargetSet = ts[0]
	}

	return output, nil
}

func (b *BalancerServiceOp) ReadTargetSet(ctx context.Context, input *ReadTargetSetInput) (*ReadTargetSetOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/targetSet/{targetSetId}", map[string]string{
		"targetSetId": StringValue(input.TargetSetID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest(ctx, "GET", path)
	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ts, err := targetSetsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadTargetSetOutput)
	if len(ts) > 0 {
		output.TargetSet = ts[0]
	}

	return output, nil
}

func (b *BalancerServiceOp) UpdateTargetSet(ctx context.Context, input *UpdateTargetSetInput) (*UpdateTargetSetOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/targetSet/{targetSetId}", map[string]string{
		"targetSetId": StringValue(input.TargetSet.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.TargetSet.ID = nil

	r := b.client.newRequest(ctx, "PUT", path)
	r.obj = input

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateTargetSetOutput{}, nil
}

func (b *BalancerServiceOp) DeleteTargetSet(ctx context.Context, input *DeleteTargetSetInput) (*DeleteTargetSetOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/targetSet/{targetSetId}", map[string]string{
		"targetSetId": StringValue(input.TargetSetID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest(ctx, "DELETE", path)
	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteTargetSetOutput{}, nil
}

// region TargetSet

func (o *TargetSet) MarshalJSON() ([]byte, error) {
	type noMethod TargetSet
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *TargetSet) SetId(v *string) *TargetSet {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *TargetSet) SetBalancerId(v *string) *TargetSet {
	if o.BalancerID = v; o.BalancerID == nil {
		o.nullFields = append(o.nullFields, "BalancerID")
	}
	return o
}

func (o *TargetSet) SetDeploymentId(v *string) *TargetSet {
	if o.DeploymentID = v; o.DeploymentID == nil {
		o.nullFields = append(o.nullFields, "DeploymentID")
	}
	return o
}

func (o *TargetSet) SetName(v *string) *TargetSet {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *TargetSet) SetProtocol(v *string) *TargetSet {
	if o.Protocol = v; o.Protocol == nil {
		o.nullFields = append(o.nullFields, "Protocol")
	}
	return o
}

func (o *TargetSet) SetPort(v *int) *TargetSet {
	if o.Port = v; o.Port == nil {
		o.nullFields = append(o.nullFields, "Port")
	}
	return o
}

func (o *TargetSet) SetWeight(v *int) *TargetSet {
	if o.Weight = v; o.Weight == nil {
		o.nullFields = append(o.nullFields, "Weight")
	}
	return o
}

func (o *TargetSet) SetHealthCheck(v *TargetSetHealthCheck) *TargetSet {
	if o.HealthCheck = v; o.HealthCheck == nil {
		o.nullFields = append(o.nullFields, "HealthCheck")
	}
	return o
}

func (o *TargetSet) SetTags(v []*Tag) *TargetSet {
	if o.Tags = v; o.Tags == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

// endregion

// region TargetSetHealthCheck

func (o *TargetSetHealthCheck) MarshalJSON() ([]byte, error) {
	type noMethod TargetSetHealthCheck
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *TargetSetHealthCheck) SetPath(v *string) *TargetSetHealthCheck {
	if o.Path = v; o.Path == nil {
		o.nullFields = append(o.nullFields, "Path")
	}
	return o
}

func (o *TargetSetHealthCheck) SetPort(v *int) *TargetSetHealthCheck {
	if o.Port = v; o.Port == nil {
		o.nullFields = append(o.nullFields, "Port")
	}
	return o
}

func (o *TargetSetHealthCheck) SetProtocol(v *string) *TargetSetHealthCheck {
	if o.Protocol = v; o.Protocol == nil {
		o.nullFields = append(o.nullFields, "Protocol")
	}
	return o
}

func (o *TargetSetHealthCheck) SetTimeout(v *int) *TargetSetHealthCheck {
	if o.Timeout = v; o.Timeout == nil {
		o.nullFields = append(o.nullFields, "Timeout")
	}
	return o
}

func (o *TargetSetHealthCheck) SetInterval(v *int) *TargetSetHealthCheck {
	if o.Interval = v; o.Interval == nil {
		o.nullFields = append(o.nullFields, "Interval")
	}
	return o
}

func (o *TargetSetHealthCheck) SetHealthyThresholdCount(v *int) *TargetSetHealthCheck {
	if o.HealthyThresholdCount = v; o.HealthyThresholdCount == nil {
		o.nullFields = append(o.nullFields, "HealthyThresholdCount")
	}
	return o
}

func (o *TargetSetHealthCheck) SetUnhealthyThresholdCount(v *int) *TargetSetHealthCheck {
	if o.UnhealthyThresholdCount = v; o.UnhealthyThresholdCount == nil {
		o.nullFields = append(o.nullFields, "UnhealthyThresholdCount")
	}
	return o
}

// endregion

type Target struct {
	ID          *string    `json:"id,omitempty"`
	BalancerID  *string    `json:"balancerId,omitempty"`
	TargetSetID *string    `json:"targetSetId,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Host        *string    `json:"host,omitempty"`
	Port        *int       `json:"port,omitempty"`
	Weight      *int       `json:"weight,omitempty"`
	Status      *Status    `json:"status,omitempty"`
	Tags        []*Tag     `json:"tags,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type Status struct {
	Readiness   *string `json:"readiness"`
	Healthiness *string `json:"healthiness"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListTargetsInput struct {
	BalancerID  *string `json:"balancerId,omitempty"`
	TargetSetID *string `json:"targetSetId,omitempty"`
}

type ListTargetsOutput struct {
	Targets []*Target `json:"targets,omitempty"`
}

type CreateTargetInput struct {
	TargetSetID *string `json:"targetSetId,omitempty"`
	Target      *Target `json:"target,omitempty"`
}

type CreateTargetOutput struct {
	Target *Target `json:"target,omitempty"`
}

type ReadTargetInput struct {
	TargetSetID *string `json:"targetSetId,omitempty"`
	TargetID    *string `json:"targetId,omitempty"`
}

type ReadTargetOutput struct {
	Target *Target `json:"target,omitempty"`
}

type UpdateTargetInput struct {
	TargetSetID *string `json:"targetSetId,omitempty"`
	Target      *Target `json:"target,omitempty"`
}

type UpdateTargetOutput struct{}

type DeleteTargetInput struct {
	TargetSetID *string `json:"targetSetId,omitempty"`
	TargetID    *string `json:"targetId,omitempty"`
}

type DeleteTargetOutput struct{}

func targetFromJSON(in []byte) (*Target, error) {
	t := new(Target)
	if err := json.Unmarshal(in, t); err != nil {
		return nil, err
	}
	return t, nil
}

func targetsFromJSON(in []byte) ([]*Target, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Target, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rt := range rw.Response.Items {
		t, err := targetFromJSON(rt)
		if err != nil {
			return nil, err
		}
		out[i] = t
	}
	return out, nil
}

func targetsFromHttpResponse(resp *http.Response) ([]*Target, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return targetsFromJSON(body)
}

func (b *BalancerServiceOp) ListTargets(ctx context.Context, input *ListTargetsInput) (*ListTargetsOutput, error) {
	r := b.client.newRequest(ctx, "GET", "/loadBalancer/target")

	if input.BalancerID != nil {
		r.params.Set("balancerId", StringValue(input.BalancerID))
	}

	if input.TargetSetID != nil {
		r.params.Set("targetSetId", StringValue(input.TargetSetID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ts, err := targetsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListTargetsOutput{Targets: ts}, nil
}

func (b *BalancerServiceOp) CreateTarget(ctx context.Context, input *CreateTargetInput) (*CreateTargetOutput, error) {
	r := b.client.newRequest(ctx, "POST", "/loadBalancer/target")
	r.obj = input

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ts, err := targetsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateTargetOutput)
	if len(ts) > 0 {
		output.Target = ts[0]
	}

	return output, nil
}

func (b *BalancerServiceOp) ReadTarget(ctx context.Context, input *ReadTargetInput) (*ReadTargetOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/target/{targetId}", map[string]string{
		"targetId": StringValue(input.TargetID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest(ctx, "GET", path)
	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ts, err := targetsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadTargetOutput)
	if len(ts) > 0 {
		output.Target = ts[0]
	}

	return output, nil
}

func (b *BalancerServiceOp) UpdateTarget(ctx context.Context, input *UpdateTargetInput) (*UpdateTargetOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/target/{targetId}", map[string]string{
		"targetId": StringValue(input.Target.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.Target.ID = nil

	r := b.client.newRequest(ctx, "PUT", path)
	r.obj = input

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateTargetOutput{}, nil
}

func (b *BalancerServiceOp) DeleteTarget(ctx context.Context, input *DeleteTargetInput) (*DeleteTargetOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/target/{targetId}", map[string]string{
		"targetId": StringValue(input.TargetID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest(ctx, "DELETE", path)
	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteTargetOutput{}, nil
}

// region Target

func (o *Target) MarshalJSON() ([]byte, error) {
	type noMethod Target
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Target) SetId(v *string) *Target {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *Target) SetBalancerId(v *string) *Target {
	if o.BalancerID = v; o.BalancerID == nil {
		o.nullFields = append(o.nullFields, "BalancerID")
	}
	return o
}

func (o *Target) SetTargetSetId(v *string) *Target {
	if o.TargetSetID = v; o.TargetSetID == nil {
		o.nullFields = append(o.nullFields, "TargetSetID")
	}
	return o
}

func (o *Target) SetName(v *string) *Target {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Target) SetHost(v *string) *Target {
	if o.Host = v; o.Host == nil {
		o.nullFields = append(o.nullFields, "Host")
	}
	return o
}

func (o *Target) SetPort(v *int) *Target {
	if o.Port = v; o.Port == nil {
		o.nullFields = append(o.nullFields, "Port")
	}
	return o
}

func (o *Target) SetWeight(v *int) *Target {
	if o.Weight = v; o.Weight == nil {
		o.nullFields = append(o.nullFields, "Weight")
	}
	return o
}

func (o *Target) SetStatus(v *Status) *Target {
	if o.Status = v; o.Status == nil {
		o.nullFields = append(o.nullFields, "Status")
	}
	return o
}

func (o *Target) SetTags(v []*Tag) *Target {
	if o.Tags = v; o.Tags == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

// endregion

type Runtime struct {
	ID             *string    `json:"id,omitempty"`
	DeploymentID   *string    `json:"deploymentId,omitempty"`
	IPAddr         *string    `json:"ip,omitempty"`
	Version        *string    `json:"version,omitempty"`
	Status         *Status    `json:"status,omitempty"`
	LastReportedAt *time.Time `json:"lastReported,omitempty"`
	Leader         *bool      `json:"isLeader,omitempty"`
	Tags           []*Tag     `json:"tags,omitempty"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListRuntimesInput struct {
	DeploymentID *string `json:"deploymentId,omitempty"`
}

type ListRuntimesOutput struct {
	Runtimes []*Runtime `json:"runtimes,omitempty"`
}

type ReadRuntimeInput struct {
	RuntimeID *string `json:"runtimeId,omitempty"`
}

type ReadRuntimeOutput struct {
	Runtime *Runtime `json:"runtime,omitempty"`
}

func runtimeFromJSON(in []byte) (*Runtime, error) {
	rt := new(Runtime)
	if err := json.Unmarshal(in, rt); err != nil {
		return nil, err
	}
	return rt, nil
}

func runtimesFromJSON(in []byte) ([]*Runtime, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Runtime, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rrt := range rw.Response.Items {
		rt, err := runtimeFromJSON(rrt)
		if err != nil {
			return nil, err
		}
		out[i] = rt
	}
	return out, nil
}

func runtimesFromHttpResponse(resp *http.Response) ([]*Runtime, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return runtimesFromJSON(body)
}

func (b *BalancerServiceOp) ListRuntimes(ctx context.Context, input *ListRuntimesInput) (*ListRuntimesOutput, error) {
	r := b.client.newRequest(ctx, "GET", "/loadBalancer/runtime")

	if input.DeploymentID != nil {
		r.params.Set("deploymentId", StringValue(input.DeploymentID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rts, err := runtimesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListRuntimesOutput{Runtimes: rts}, nil
}

func (b *BalancerServiceOp) ReadRuntime(ctx context.Context, input *ReadRuntimeInput) (*ReadRuntimeOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/runtime/{runtimeId}", map[string]string{
		"runtimeId": StringValue(input.RuntimeID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest(ctx, "GET", path)
	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rt, err := runtimesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadRuntimeOutput)
	if len(rt) > 0 {
		output.Runtime = rt[0]
	}

	return output, nil
}

// region Runtime

func (o *Runtime) MarshalJSON() ([]byte, error) {
	type noMethod Runtime
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Runtime) SetId(v *string) *Runtime {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *Runtime) SetDeploymentId(v *string) *Runtime {
	if o.DeploymentID = v; o.DeploymentID == nil {
		o.nullFields = append(o.nullFields, "DeploymentID")
	}
	return o
}

func (o *Runtime) SetStatus(v *Status) *Runtime {
	if o.Status = v; o.Status == nil {
		o.nullFields = append(o.nullFields, "Status")
	}
	return o
}

func (o *Runtime) SetTags(v []*Tag) *Runtime {
	if o.Tags = v; o.Tags == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

// endregion
