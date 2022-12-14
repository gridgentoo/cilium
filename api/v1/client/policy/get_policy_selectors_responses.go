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

// GetPolicySelectorsReader is a Reader for the GetPolicySelectors structure.
type GetPolicySelectorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolicySelectorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPolicySelectorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPolicySelectorsOK creates a GetPolicySelectorsOK with default headers values
func NewGetPolicySelectorsOK() *GetPolicySelectorsOK {
	return &GetPolicySelectorsOK{}
}

/*
GetPolicySelectorsOK handles this case with default header values.

Success
*/
type GetPolicySelectorsOK struct {
	Payload models.SelectorCache
}

func (o *GetPolicySelectorsOK) Error() string {
	return fmt.Sprintf("[GET /policy/selectors][%d] getPolicySelectorsOK  %+v", 200, o.Payload)
}

func (o *GetPolicySelectorsOK) GetPayload() models.SelectorCache {
	return o.Payload
}

func (o *GetPolicySelectorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
