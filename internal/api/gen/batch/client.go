// Code generated by goa v3.4.3, DO NOT EDIT.
//
// batch client
//
// Command:
// $ goa-v3.4.3 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package batch

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "batch" service client.
type Client struct {
	SubmitEndpoint goa.Endpoint
	StatusEndpoint goa.Endpoint
}

// NewClient initializes a "batch" service client given the endpoints.
func NewClient(submit, status goa.Endpoint) *Client {
	return &Client{
		SubmitEndpoint: submit,
		StatusEndpoint: status,
	}
}

// Submit calls the "submit" endpoint of the "batch" service.
// Submit may return the following errors:
//	- "not_available" (type *goa.ServiceError)
//	- "not_valid" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Submit(ctx context.Context, p *SubmitPayload) (res *BatchResult, err error) {
	var ires interface{}
	ires, err = c.SubmitEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*BatchResult), nil
}

// Status calls the "status" endpoint of the "batch" service.
func (c *Client) Status(ctx context.Context) (res *BatchStatusResult, err error) {
	var ires interface{}
	ires, err = c.StatusEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*BatchStatusResult), nil
}