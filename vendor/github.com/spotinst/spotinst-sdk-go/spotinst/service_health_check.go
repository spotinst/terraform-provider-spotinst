package spotinst

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

// HealthCheck is an interface for interfacing with the HealthCheck
// endpoints of the Spotinst API.
type HealthCheckService interface {
	List(context.Context, *ListHealthCheckInput) (*ListHealthCheckOutput, error)
	Create(context.Context, *CreateHealthCheckInput) (*CreateHealthCheckOutput, error)
	Read(context.Context, *ReadHealthCheckInput) (*ReadHealthCheckOutput, error)
	Update(context.Context, *UpdateHealthCheckInput) (*UpdateHealthCheckOutput, error)
	Delete(context.Context, *DeleteHealthCheckInput) (*DeleteHealthCheckOutput, error)
}

// HealthCheckServiceOp handles communication with the balancer related methods
// of the Spotinst API.
type HealthCheckServiceOp struct {
	client *Client
}

var _ HealthCheckService = &HealthCheckServiceOp{}

type HealthCheck struct {
	ID         *string            `json:"id,omitempty"`
	Name       *string            `json:"name,omitempty"`
	ResourceID *string            `json:"resourceId,omitempty"`
	Check      *HealthCheckConfig `json:"check,omitempty"`
	ProxyAddr  *string            `json:"proxyAddress,omitempty"`
	ProxyPort  *int               `json:"proxyPort,omitempty"`

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

type HealthCheckConfig struct {
	Protocol  *string `json:"protocol,omitempty"`
	Endpoint  *string `json:"endpoint,omitempty"`
	Port      *int    `json:"port,omitempty"`
	Interval  *int    `json:"interval,omitempty"`
	Timeout   *int    `json:"timeout,omitempty"`
	Healthy   *int    `json:"healthyThreshold,omitempty"`
	Unhealthy *int    `json:"unhealthyThreshold,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListHealthCheckInput struct{}

type ListHealthCheckOutput struct {
	HealthChecks []*HealthCheck `json:"healthChecks,omitempty"`
}

type CreateHealthCheckInput struct {
	HealthCheck *HealthCheck `json:"healthCheck,omitempty"`
}

type CreateHealthCheckOutput struct {
	HealthCheck *HealthCheck `json:"healthCheck,omitempty"`
}

type ReadHealthCheckInput struct {
	HealthCheckID *string `json:"healthCheckId,omitempty"`
}

type ReadHealthCheckOutput struct {
	HealthCheck *HealthCheck `json:"healthCheck,omitempty"`
}

type UpdateHealthCheckInput struct {
	HealthCheck *HealthCheck `json:"healthCheck,omitempty"`
}

type UpdateHealthCheckOutput struct {
	HealthCheck *HealthCheck `json:"healthCheck,omitempty"`
}

type DeleteHealthCheckInput struct {
	HealthCheckID *string `json:"healthCheckId,omitempty"`
}

type DeleteHealthCheckOutput struct{}

func healthCheckFromJSON(in []byte) (*HealthCheck, error) {
	b := new(HealthCheck)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func healthChecksFromJSON(in []byte) ([]*HealthCheck, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*HealthCheck, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := healthCheckFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func healthChecksFromHttpResponse(resp *http.Response) ([]*HealthCheck, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return healthChecksFromJSON(body)
}

func (s *HealthCheckServiceOp) List(ctx context.Context, input *ListHealthCheckInput) (*ListHealthCheckOutput, error) {
	r := s.client.newRequest(ctx, "GET", "/healthCheck")
	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	hcs, err := healthChecksFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListHealthCheckOutput{HealthChecks: hcs}, nil
}

func (s *HealthCheckServiceOp) Create(ctx context.Context, input *CreateHealthCheckInput) (*CreateHealthCheckOutput, error) {
	r := s.client.newRequest(ctx, "POST", "/healthCheck")
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	hcs, err := healthChecksFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateHealthCheckOutput)
	if len(hcs) > 0 {
		output.HealthCheck = hcs[0]
	}

	return output, nil
}

func (s *HealthCheckServiceOp) Read(ctx context.Context, input *ReadHealthCheckInput) (*ReadHealthCheckOutput, error) {
	path, err := uritemplates.Expand("/healthCheck/{healthCheckId}", map[string]string{
		"healthCheckId": StringValue(input.HealthCheckID),
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

	hcs, err := healthChecksFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadHealthCheckOutput)
	if len(hcs) > 0 {
		output.HealthCheck = hcs[0]
	}

	return output, nil
}

func (s *HealthCheckServiceOp) Update(ctx context.Context, input *UpdateHealthCheckInput) (*UpdateHealthCheckOutput, error) {
	path, err := uritemplates.Expand("/healthCheck/{healthCheckId}", map[string]string{
		"healthCheckId": StringValue(input.HealthCheck.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.HealthCheck.ID = nil

	r := s.client.newRequest(ctx, "PUT", path)
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	hcs, err := healthChecksFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateHealthCheckOutput)
	if len(hcs) > 0 {
		output.HealthCheck = hcs[0]
	}

	return output, nil
}

func (s *HealthCheckServiceOp) Delete(ctx context.Context, input *DeleteHealthCheckInput) (*DeleteHealthCheckOutput, error) {
	path, err := uritemplates.Expand("/healthCheck/{healthCheckId}", map[string]string{
		"healthCheckId": StringValue(input.HealthCheckID),
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

	return &DeleteHealthCheckOutput{}, nil
}

// region HealthCheck

func (o *HealthCheck) MarshalJSON() ([]byte, error) {
	type noMethod HealthCheck
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *HealthCheck) SetId(v *string) *HealthCheck {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *HealthCheck) SetName(v *string) *HealthCheck {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *HealthCheck) SetResourceId(v *string) *HealthCheck {
	if o.ResourceID = v; o.ResourceID == nil {
		o.nullFields = append(o.nullFields, "ResourceID")
	}
	return o
}

func (o *HealthCheck) SetCheck(v *HealthCheckConfig) *HealthCheck {
	if o.Check = v; o.Check == nil {
		o.nullFields = append(o.nullFields, "Check")
	}
	return o
}

func (o *HealthCheck) SetProxyAddr(v *string) *HealthCheck {
	if o.ProxyAddr = v; o.ProxyAddr == nil {
		o.nullFields = append(o.nullFields, "ProxyAddr")
	}
	return o
}

func (o *HealthCheck) SetProxyPort(v *int) *HealthCheck {
	if o.ProxyPort = v; o.ProxyPort == nil {
		o.nullFields = append(o.nullFields, "ProxyPort")
	}
	return o
}

// endregion

// region HealthCheckConfig

func (o *HealthCheckConfig) MarshalJSON() ([]byte, error) {
	type noMethod HealthCheckConfig
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *HealthCheckConfig) SetProtocol(v *string) *HealthCheckConfig {
	if o.Protocol = v; o.Protocol == nil {
		o.nullFields = append(o.nullFields, "Protocol")
	}
	return o
}

func (o *HealthCheckConfig) SetEndpoint(v *string) *HealthCheckConfig {
	if o.Endpoint = v; o.Endpoint == nil {
		o.nullFields = append(o.nullFields, "Endpoint")
	}
	return o
}

func (o *HealthCheckConfig) SetPort(v *int) *HealthCheckConfig {
	if o.Port = v; o.Port == nil {
		o.nullFields = append(o.nullFields, "Port")
	}
	return o
}

func (o *HealthCheckConfig) SetInterval(v *int) *HealthCheckConfig {
	if o.Interval = v; o.Interval == nil {
		o.nullFields = append(o.nullFields, "Interval")
	}
	return o
}

func (o *HealthCheckConfig) SetTimeout(v *int) *HealthCheckConfig {
	if o.Timeout = v; o.Timeout == nil {
		o.nullFields = append(o.nullFields, "Timeout")
	}
	return o
}

func (o *HealthCheckConfig) SetHealthy(v *int) *HealthCheckConfig {
	if o.Healthy = v; o.Healthy == nil {
		o.nullFields = append(o.nullFields, "Healthy")
	}
	return o
}

func (o *HealthCheckConfig) SetUnhealthy(v *int) *HealthCheckConfig {
	if o.Unhealthy = v; o.Unhealthy == nil {
		o.nullFields = append(o.nullFields, "Unhealthy")
	}
	return o
}

// endregion
