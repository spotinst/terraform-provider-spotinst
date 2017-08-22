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

// CertificateService is an interface for interfacing with the certificate
// endpoints of the Spotinst API.
type CertificateService interface {
	List(context.Context, *ListCertificatesInput) (*ListCertificatesOutput, error)
	Create(context.Context, *CreateCertificateInput) (*CreateCertificateOutput, error)
	Read(context.Context, *ReadCertificateInput) (*ReadCertificateOutput, error)
	Update(context.Context, *UpdateCertificateInput) (*UpdateCertificateOutput, error)
	Delete(context.Context, *DeleteCertificateInput) (*DeleteCertificateOutput, error)
}

// CertificateServiceOp handles communication with the certificate related methods
// of the Spotinst API.
type CertificateServiceOp struct {
	client *Client
}

var _ CertificateService = &CertificateServiceOp{}

type Certificate struct {
	ID           *string    `json:"id,omitempty"`
	Name         *string    `json:"name,omitempty"`
	CertPEMBlock *string    `json:"certificatePemBlock,omitempty"`
	KeyPEMBlock  *string    `json:"keyPemBlock,omitempty"`
	Tags         []*Tag     `json:"tags,omitempty"`
	CreatedAt    *time.Time `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListCertificatesInput struct{}

type ListCertificatesOutput struct {
	Certificates []*Certificate `json:"certificates,omitempty"`
}

type CreateCertificateInput struct {
	Certificate *Certificate `json:"certificate,omitempty"`
}

type CreateCertificateOutput struct {
	Certificate *Certificate `json:"certificate,omitempty"`
}

type ReadCertificateInput struct {
	CertificateID *string `json:"certificateId,omitempty"`
}

type ReadCertificateOutput struct {
	Certificate *Certificate `json:"certificate,omitempty"`
}

type UpdateCertificateInput struct {
	Certificate *Certificate `json:"certificate,omitempty"`
}

type UpdateCertificateOutput struct{}

type DeleteCertificateInput struct {
	CertificateID *string `json:"certificateId,omitempty"`
}

type DeleteCertificateOutput struct{}

func certificateFromJSON(in []byte) (*Certificate, error) {
	b := new(Certificate)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func certificatesFromJSON(in []byte) ([]*Certificate, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Certificate, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rp := range rw.Response.Items {
		p, err := certificateFromJSON(rp)
		if err != nil {
			return nil, err
		}
		out[i] = p
	}
	return out, nil
}

func certificatesFromHttpResponse(resp *http.Response) ([]*Certificate, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return certificatesFromJSON(body)
}

func (c *CertificateServiceOp) List(ctx context.Context, input *ListCertificatesInput) (*ListCertificatesOutput, error) {
	r := c.client.newRequest(ctx, "GET", "/loadBalancer/certificate")
	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cs, err := certificatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListCertificatesOutput{Certificates: cs}, nil
}

func (c *CertificateServiceOp) Create(ctx context.Context, input *CreateCertificateInput) (*CreateCertificateOutput, error) {
	r := c.client.newRequest(ctx, "POST", "/loadBalancer/certificate")
	r.obj = input

	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cs, err := certificatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateCertificateOutput)
	if len(cs) > 0 {
		output.Certificate = cs[0]
	}

	return output, nil
}

func (c *CertificateServiceOp) Read(ctx context.Context, input *ReadCertificateInput) (*ReadCertificateOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/certificate/{certificateId}", map[string]string{
		"certificateId": StringValue(input.CertificateID),
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

	cs, err := certificatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadCertificateOutput)
	if len(cs) > 0 {
		output.Certificate = cs[0]
	}

	return output, nil
}

func (c *CertificateServiceOp) Update(ctx context.Context, input *UpdateCertificateInput) (*UpdateCertificateOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/certificate/{certificateId}", map[string]string{
		"certificateId": StringValue(input.Certificate.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.Certificate.ID = nil

	r := c.client.newRequest(ctx, "PUT", path)
	r.obj = input

	_, resp, err := requireOK(c.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateCertificateOutput{}, nil
}

func (c *CertificateServiceOp) Delete(ctx context.Context, input *DeleteCertificateInput) (*DeleteCertificateOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/certificate/{certificateId}", map[string]string{
		"certificateId": StringValue(input.CertificateID),
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

	return &DeleteCertificateOutput{}, nil
}

// region Certificate

func (o *Certificate) MarshalJSON() ([]byte, error) {
	type noMethod Certificate
	raw := noMethod(*o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Certificate) SetId(v *string) *Certificate {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *Certificate) SetName(v *string) *Certificate {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Certificate) SetCertPEMBlock(v *string) *Certificate {
	if o.CertPEMBlock = v; o.CertPEMBlock == nil {
		o.nullFields = append(o.nullFields, "CertPEMBlock")
	}
	return o
}

func (o *Certificate) SetKeyPEMBlock(v *string) *Certificate {
	if o.KeyPEMBlock = v; o.KeyPEMBlock == nil {
		o.nullFields = append(o.nullFields, "KeyPEMBlock")
	}
	return o
}

func (o *Certificate) SetTags(v []*Tag) *Certificate {
	if o.Tags = v; o.Tags == nil {
		o.nullFields = append(o.nullFields, "Tags")
	}
	return o
}

// endregion
