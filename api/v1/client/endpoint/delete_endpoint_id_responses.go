// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// DeleteEndpointIDReader is a Reader for the DeleteEndpointID structure.
type DeleteEndpointIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEndpointIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEndpointIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 206:
		result := NewDeleteEndpointIDErrors()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteEndpointIDInvalid()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteEndpointIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteEndpointIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEndpointIDOK creates a DeleteEndpointIDOK with default headers values
func NewDeleteEndpointIDOK() *DeleteEndpointIDOK {
	return &DeleteEndpointIDOK{}
}

/*
DeleteEndpointIDOK handles this case with default header values.

Success
*/
type DeleteEndpointIDOK struct {
}

func (o *DeleteEndpointIDOK) Error() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdOK ", 200)
}

func (o *DeleteEndpointIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEndpointIDErrors creates a DeleteEndpointIDErrors with default headers values
func NewDeleteEndpointIDErrors() *DeleteEndpointIDErrors {
	return &DeleteEndpointIDErrors{}
}

/*
DeleteEndpointIDErrors handles this case with default header values.

Deleted with a number of errors encountered
*/
type DeleteEndpointIDErrors struct {
	Payload int64
}

func (o *DeleteEndpointIDErrors) Error() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdErrors  %+v", 206, o.Payload)
}

func (o *DeleteEndpointIDErrors) GetPayload() int64 {
	return o.Payload
}

func (o *DeleteEndpointIDErrors) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEndpointIDInvalid creates a DeleteEndpointIDInvalid with default headers values
func NewDeleteEndpointIDInvalid() *DeleteEndpointIDInvalid {
	return &DeleteEndpointIDInvalid{}
}

/*
DeleteEndpointIDInvalid handles this case with default header values.

Invalid endpoint ID format for specified type. Details in error
message
*/
type DeleteEndpointIDInvalid struct {
	Payload models.Error
}

func (o *DeleteEndpointIDInvalid) Error() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdInvalid  %+v", 400, o.Payload)
}

func (o *DeleteEndpointIDInvalid) GetPayload() models.Error {
	return o.Payload
}

func (o *DeleteEndpointIDInvalid) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEndpointIDNotFound creates a DeleteEndpointIDNotFound with default headers values
func NewDeleteEndpointIDNotFound() *DeleteEndpointIDNotFound {
	return &DeleteEndpointIDNotFound{}
}

/*
DeleteEndpointIDNotFound handles this case with default header values.

Endpoint not found
*/
type DeleteEndpointIDNotFound struct {
}

func (o *DeleteEndpointIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdNotFound ", 404)
}

func (o *DeleteEndpointIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEndpointIDTooManyRequests creates a DeleteEndpointIDTooManyRequests with default headers values
func NewDeleteEndpointIDTooManyRequests() *DeleteEndpointIDTooManyRequests {
	return &DeleteEndpointIDTooManyRequests{}
}

/*
DeleteEndpointIDTooManyRequests handles this case with default header values.

Rate-limiting too many requests in the given time frame
*/
type DeleteEndpointIDTooManyRequests struct {
}

func (o *DeleteEndpointIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /endpoint/{id}][%d] deleteEndpointIdTooManyRequests ", 429)
}

func (o *DeleteEndpointIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
