// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package prefilter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// PatchPrefilterReader is a Reader for the PatchPrefilter structure.
type PatchPrefilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPrefilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchPrefilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 461:
		result := NewPatchPrefilterInvalidCIDR()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchPrefilterFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchPrefilterOK creates a PatchPrefilterOK with default headers values
func NewPatchPrefilterOK() *PatchPrefilterOK {
	return &PatchPrefilterOK{}
}

/*
PatchPrefilterOK handles this case with default header values.

Updated
*/
type PatchPrefilterOK struct {
	Payload *models.Prefilter
}

func (o *PatchPrefilterOK) Error() string {
	return fmt.Sprintf("[PATCH /prefilter][%d] patchPrefilterOK  %+v", 200, o.Payload)
}

func (o *PatchPrefilterOK) GetPayload() *models.Prefilter {
	return o.Payload
}

func (o *PatchPrefilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPrefilterInvalidCIDR creates a PatchPrefilterInvalidCIDR with default headers values
func NewPatchPrefilterInvalidCIDR() *PatchPrefilterInvalidCIDR {
	return &PatchPrefilterInvalidCIDR{}
}

/*
PatchPrefilterInvalidCIDR handles this case with default header values.

Invalid CIDR prefix
*/
type PatchPrefilterInvalidCIDR struct {
	Payload models.Error
}

func (o *PatchPrefilterInvalidCIDR) Error() string {
	return fmt.Sprintf("[PATCH /prefilter][%d] patchPrefilterInvalidCIdR  %+v", 461, o.Payload)
}

func (o *PatchPrefilterInvalidCIDR) GetPayload() models.Error {
	return o.Payload
}

func (o *PatchPrefilterInvalidCIDR) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPrefilterFailure creates a PatchPrefilterFailure with default headers values
func NewPatchPrefilterFailure() *PatchPrefilterFailure {
	return &PatchPrefilterFailure{}
}

/*
PatchPrefilterFailure handles this case with default header values.

Prefilter update failed
*/
type PatchPrefilterFailure struct {
	Payload models.Error
}

func (o *PatchPrefilterFailure) Error() string {
	return fmt.Sprintf("[PATCH /prefilter][%d] patchPrefilterFailure  %+v", 500, o.Payload)
}

func (o *PatchPrefilterFailure) GetPayload() models.Error {
	return o.Payload
}

func (o *PatchPrefilterFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
