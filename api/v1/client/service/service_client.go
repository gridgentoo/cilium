// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteServiceID(params *DeleteServiceIDParams) (*DeleteServiceIDOK, error)

	GetLrp(params *GetLrpParams) (*GetLrpOK, error)

	GetService(params *GetServiceParams) (*GetServiceOK, error)

	GetServiceID(params *GetServiceIDParams) (*GetServiceIDOK, error)

	PutServiceID(params *PutServiceIDParams) (*PutServiceIDOK, *PutServiceIDCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteServiceID deletes a service
*/
func (a *Client) DeleteServiceID(params *DeleteServiceIDParams) (*DeleteServiceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServiceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteServiceID",
		Method:             "DELETE",
		PathPattern:        "/service/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteServiceIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteServiceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteServiceID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLrp retrieves list of all local redirect policies
*/
func (a *Client) GetLrp(params *GetLrpParams) (*GetLrpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLrpParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLrp",
		Method:             "GET",
		PathPattern:        "/lrp",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLrpReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLrpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLrp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetService retrieves list of all services
*/
func (a *Client) GetService(params *GetServiceParams) (*GetServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetService",
		Method:             "GET",
		PathPattern:        "/service",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetServiceID retrieves configuration of a service
*/
func (a *Client) GetServiceID(params *GetServiceIDParams) (*GetServiceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetServiceID",
		Method:             "GET",
		PathPattern:        "/service/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServiceIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetServiceID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutServiceID creates or update service
*/
func (a *Client) PutServiceID(params *PutServiceIDParams) (*PutServiceIDOK, *PutServiceIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutServiceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutServiceID",
		Method:             "PUT",
		PathPattern:        "/service/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutServiceIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PutServiceIDOK:
		return value, nil, nil
	case *PutServiceIDCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for service: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
