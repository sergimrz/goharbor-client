// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v4/apiv1/model"
)

// GetReplicationPoliciesIDReader is a Reader for the GetReplicationPoliciesID structure.
type GetReplicationPoliciesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReplicationPoliciesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReplicationPoliciesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReplicationPoliciesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetReplicationPoliciesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReplicationPoliciesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReplicationPoliciesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReplicationPoliciesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReplicationPoliciesIDOK creates a GetReplicationPoliciesIDOK with default headers values
func NewGetReplicationPoliciesIDOK() *GetReplicationPoliciesIDOK {
	return &GetReplicationPoliciesIDOK{}
}

/* GetReplicationPoliciesIDOK describes a response with status code 200, with default header values.

Get the replication policy successfully.
*/
type GetReplicationPoliciesIDOK struct {
	Payload *model.ReplicationPolicy
}

func (o *GetReplicationPoliciesIDOK) Error() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPoliciesIdOK  %+v", 200, o.Payload)
}
func (o *GetReplicationPoliciesIDOK) GetPayload() *model.ReplicationPolicy {
	return o.Payload
}

func (o *GetReplicationPoliciesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ReplicationPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationPoliciesIDBadRequest creates a GetReplicationPoliciesIDBadRequest with default headers values
func NewGetReplicationPoliciesIDBadRequest() *GetReplicationPoliciesIDBadRequest {
	return &GetReplicationPoliciesIDBadRequest{}
}

/* GetReplicationPoliciesIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetReplicationPoliciesIDBadRequest struct {
}

func (o *GetReplicationPoliciesIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPoliciesIdBadRequest ", 400)
}

func (o *GetReplicationPoliciesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationPoliciesIDUnauthorized creates a GetReplicationPoliciesIDUnauthorized with default headers values
func NewGetReplicationPoliciesIDUnauthorized() *GetReplicationPoliciesIDUnauthorized {
	return &GetReplicationPoliciesIDUnauthorized{}
}

/* GetReplicationPoliciesIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetReplicationPoliciesIDUnauthorized struct {
}

func (o *GetReplicationPoliciesIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPoliciesIdUnauthorized ", 401)
}

func (o *GetReplicationPoliciesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationPoliciesIDForbidden creates a GetReplicationPoliciesIDForbidden with default headers values
func NewGetReplicationPoliciesIDForbidden() *GetReplicationPoliciesIDForbidden {
	return &GetReplicationPoliciesIDForbidden{}
}

/* GetReplicationPoliciesIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetReplicationPoliciesIDForbidden struct {
}

func (o *GetReplicationPoliciesIDForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPoliciesIdForbidden ", 403)
}

func (o *GetReplicationPoliciesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationPoliciesIDNotFound creates a GetReplicationPoliciesIDNotFound with default headers values
func NewGetReplicationPoliciesIDNotFound() *GetReplicationPoliciesIDNotFound {
	return &GetReplicationPoliciesIDNotFound{}
}

/* GetReplicationPoliciesIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetReplicationPoliciesIDNotFound struct {
}

func (o *GetReplicationPoliciesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPoliciesIdNotFound ", 404)
}

func (o *GetReplicationPoliciesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationPoliciesIDInternalServerError creates a GetReplicationPoliciesIDInternalServerError with default headers values
func NewGetReplicationPoliciesIDInternalServerError() *GetReplicationPoliciesIDInternalServerError {
	return &GetReplicationPoliciesIDInternalServerError{}
}

/* GetReplicationPoliciesIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetReplicationPoliciesIDInternalServerError struct {
}

func (o *GetReplicationPoliciesIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPoliciesIdInternalServerError ", 500)
}

func (o *GetReplicationPoliciesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
