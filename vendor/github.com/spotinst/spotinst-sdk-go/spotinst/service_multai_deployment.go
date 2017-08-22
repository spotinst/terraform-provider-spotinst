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

// DeploymentService is an interface for interfacing with the deployment
// endpoints of the Spotinst API.
type DeploymentService interface {
	List(context.Context, *ListDeploymentsInput) (*ListDeploymentsOutput, error)
	Create(context.Context, *CreateDeploymentInput) (*CreateDeploymentOutput, error)
	Read(context.Context, *ReadDeploymentInput) (*ReadDeploymentOutput, error)
	Update(context.Context, *UpdateDeploymentInput) (*UpdateDeploymentOutput, error)
	Delete(context.Context, *DeleteDeploymentInput) (*DeleteDeploymentOutput, error)
}

// DeploymentServiceOp handles communication with the deployment related methods
// of the Spotinst API.
type DeploymentServiceOp struct {
	client *Client
}

var _ DeploymentService = &DeploymentServiceOp{}

type Deployment struct {
	ID        *string    `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Tags      []*Tag     `json:"tags,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListDeploymentsInput struct{}

type ListDeploymentsOutput struct {
	Deployments []*Deployment `json:"deployments,omitempty"`
}

type CreateDeploymentInput struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

type CreateDeploymentOutput struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

type ReadDeploymentInput struct {
	DeploymentID *string `json:"deploymentId,omitempty"`
}

type ReadDeploymentOutput struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

type UpdateDeploymentInput struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

type UpdateDeploymentOutput struct{}

type DeleteDeploymentInput struct {
	DeploymentID *string `json:"deployment,omitempty"`
}

type DeleteDeploymentOutput struct{}

func deploymentFromJSON(in []byte) (*Deployment, error) {
	b := new(Deployment)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func deploymentsFromJSON(in []byte) ([]*Deployment, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Deployment, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rp := range rw.Response.Items {
		p, err := deploymentFromJSON(rp)
		if err != nil {
			return nil, err
		}
		out[i] = p
	}
	return out, nil
}

func deploymentsFromHttpResponse(resp *http.Response) ([]*Deployment, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return deploymentsFromJSON(body)
}

func (c *DeploymentServiceOp) List(ctx context.Context, input *ListDeploymentsInput) (*ListDeploymentsOutput, error) {
	r := c.client.newRequest(ctx, "GET", "/loadBalancer/deployment")
	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ds, err := deploymentsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListDeploymentsOutput{Deployments: ds}, nil
}

func (c *DeploymentServiceOp) Create(ctx context.Context, input *CreateDeploymentInput) (*CreateDeploymentOutput, error) {
	r := c.client.newRequest(ctx, "POST", "/loadBalancer/deployment")
	r.obj = input

	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ds, err := deploymentsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateDeploymentOutput)
	if len(ds) > 0 {
		output.Deployment = ds[0]
	}

	return output, nil
}

func (c *DeploymentServiceOp) Read(ctx context.Context, input *ReadDeploymentInput) (*ReadDeploymentOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/deployment/{deploymentId}", map[string]string{
		"deploymentId": StringValue(input.DeploymentID),
	})
	if err != nil {
		return nil, err
	}

	r := c.client.newRequest(ctx, "GET", path)
	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ds, err := deploymentsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadDeploymentOutput)
	if len(ds) > 0 {
		output.Deployment = ds[0]
	}

	return output, nil
}

func (c *DeploymentServiceOp) Update(ctx context.Context, input *UpdateDeploymentInput) (*UpdateDeploymentOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/deployment/{deploymentId}", map[string]string{
		"deploymentId": StringValue(input.Deployment.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.Deployment.ID = nil

	r := c.client.newRequest(ctx, "PUT", path)
	r.obj = input

	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateDeploymentOutput{}, nil
}

func (c *DeploymentServiceOp) Delete(ctx context.Context, input *DeleteDeploymentInput) (*DeleteDeploymentOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/deployment/{deploymentId}", map[string]string{
		"deploymentId": StringValue(input.DeploymentID),
	})
	if err != nil {
		return nil, err
	}

	r := c.client.newRequest(ctx, "DELETE", path)
	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteDeploymentOutput{}, nil
}

// region Deployment

func (o *Deployment) MarshalJSON() ([]byte, error) {
	type noMethod Deployment
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Deployment) SetId(v *string) *Deployment {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *Deployment) SetName(v *string) *Deployment {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Deployment) SetTags(v []*Tag) *Deployment {
	if o.Tags = v; o.Tags == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

// endregion
