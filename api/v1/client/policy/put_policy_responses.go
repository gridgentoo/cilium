// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// PutPolicyReader is a Reader for the PutPolicy structure.
type PutPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutPolicyInvalidPolicy()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 460:
		result := NewPutPolicyInvalidPath()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutPolicyFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutPolicyOK creates a PutPolicyOK with default headers values
func NewPutPolicyOK() *PutPolicyOK {
	return &PutPolicyOK{}
}

/*
PutPolicyOK handles this case with default header values.

Success
*/
type PutPolicyOK struct {
	Payload *models.Policy
}

func (o *PutPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /policy][%d] putPolicyOK  %+v", 200, o.Payload)
}

func (o *PutPolicyOK) GetPayload() *models.Policy {
	return o.Payload
}

func (o *PutPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Policy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPolicyInvalidPolicy creates a PutPolicyInvalidPolicy with default headers values
func NewPutPolicyInvalidPolicy() *PutPolicyInvalidPolicy {
	return &PutPolicyInvalidPolicy{}
}

/*
PutPolicyInvalidPolicy handles this case with default header values.

Invalid policy
*/
type PutPolicyInvalidPolicy struct {
	Payload models.Error
}

func (o *PutPolicyInvalidPolicy) Error() string {
	return fmt.Sprintf("[PUT /policy][%d] putPolicyInvalidPolicy  %+v", 400, o.Payload)
}

func (o *PutPolicyInvalidPolicy) GetPayload() models.Error {
	return o.Payload
}

func (o *PutPolicyInvalidPolicy) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPolicyInvalidPath creates a PutPolicyInvalidPath with default headers values
func NewPutPolicyInvalidPath() *PutPolicyInvalidPath {
	return &PutPolicyInvalidPath{}
}

/*
PutPolicyInvalidPath handles this case with default header values.

Invalid path
*/
type PutPolicyInvalidPath struct {
	Payload models.Error
}

func (o *PutPolicyInvalidPath) Error() string {
	return fmt.Sprintf("[PUT /policy][%d] putPolicyInvalidPath  %+v", 460, o.Payload)
}

func (o *PutPolicyInvalidPath) GetPayload() models.Error {
	return o.Payload
}

func (o *PutPolicyInvalidPath) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPolicyFailure creates a PutPolicyFailure with default headers values
func NewPutPolicyFailure() *PutPolicyFailure {
	return &PutPolicyFailure{}
}

/*
PutPolicyFailure handles this case with default header values.

Policy import failed
*/
type PutPolicyFailure struct {
	Payload models.Error
}

func (o *PutPolicyFailure) Error() string {
	return fmt.Sprintf("[PUT /policy][%d] putPolicyFailure  %+v", 500, o.Payload)
}

func (o *PutPolicyFailure) GetPayload() models.Error {
	return o.Payload
}

func (o *PutPolicyFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
