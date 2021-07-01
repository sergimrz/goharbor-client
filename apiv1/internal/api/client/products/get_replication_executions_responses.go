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

// GetReplicationExecutionsReader is a Reader for the GetReplicationExecutions structure.
type GetReplicationExecutionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReplicationExecutionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReplicationExecutionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetReplicationExecutionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReplicationExecutionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReplicationExecutionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReplicationExecutionsOK creates a GetReplicationExecutionsOK with default headers values
func NewGetReplicationExecutionsOK() *GetReplicationExecutionsOK {
	return &GetReplicationExecutionsOK{}
}

/* GetReplicationExecutionsOK describes a response with status code 200, with default header values.

Success
*/
type GetReplicationExecutionsOK struct {
	Payload []*model.ReplicationExecution
}

func (o *GetReplicationExecutionsOK) Error() string {
	return fmt.Sprintf("[GET /replication/executions][%d] getReplicationExecutionsOK  %+v", 200, o.Payload)
}
func (o *GetReplicationExecutionsOK) GetPayload() []*model.ReplicationExecution {
	return o.Payload
}

func (o *GetReplicationExecutionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationExecutionsUnauthorized creates a GetReplicationExecutionsUnauthorized with default headers values
func NewGetReplicationExecutionsUnauthorized() *GetReplicationExecutionsUnauthorized {
	return &GetReplicationExecutionsUnauthorized{}
}

/* GetReplicationExecutionsUnauthorized describes a response with status code 401, with default header values.

User need to login first.
*/
type GetReplicationExecutionsUnauthorized struct {
}

func (o *GetReplicationExecutionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/executions][%d] getReplicationExecutionsUnauthorized ", 401)
}

func (o *GetReplicationExecutionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationExecutionsForbidden creates a GetReplicationExecutionsForbidden with default headers values
func NewGetReplicationExecutionsForbidden() *GetReplicationExecutionsForbidden {
	return &GetReplicationExecutionsForbidden{}
}

/* GetReplicationExecutionsForbidden describes a response with status code 403, with default header values.

User has no privilege for the operation.
*/
type GetReplicationExecutionsForbidden struct {
}

func (o *GetReplicationExecutionsForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/executions][%d] getReplicationExecutionsForbidden ", 403)
}

func (o *GetReplicationExecutionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationExecutionsInternalServerError creates a GetReplicationExecutionsInternalServerError with default headers values
func NewGetReplicationExecutionsInternalServerError() *GetReplicationExecutionsInternalServerError {
	return &GetReplicationExecutionsInternalServerError{}
}

/* GetReplicationExecutionsInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetReplicationExecutionsInternalServerError struct {
}

func (o *GetReplicationExecutionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/executions][%d] getReplicationExecutionsInternalServerError ", 500)
}

func (o *GetReplicationExecutionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
