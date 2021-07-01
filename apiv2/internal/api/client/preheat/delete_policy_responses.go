// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v4/apiv2/model"
)

// DeletePolicyReader is a Reader for the DeletePolicy structure.
type DeletePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeletePolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePolicyOK creates a DeletePolicyOK with default headers values
func NewDeletePolicyOK() *DeletePolicyOK {
	return &DeletePolicyOK{}
}

/* DeletePolicyOK describes a response with status code 200, with default header values.

Success
*/
type DeletePolicyOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeletePolicyOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyOK ", 200)
}

func (o *DeletePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeletePolicyBadRequest creates a DeletePolicyBadRequest with default headers values
func NewDeletePolicyBadRequest() *DeletePolicyBadRequest {
	return &DeletePolicyBadRequest{}
}

/* DeletePolicyBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeletePolicyBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeletePolicyBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyBadRequest  %+v", 400, o.Payload)
}
func (o *DeletePolicyBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeletePolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePolicyUnauthorized creates a DeletePolicyUnauthorized with default headers values
func NewDeletePolicyUnauthorized() *DeletePolicyUnauthorized {
	return &DeletePolicyUnauthorized{}
}

/* DeletePolicyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeletePolicyUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeletePolicyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *DeletePolicyUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeletePolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePolicyForbidden creates a DeletePolicyForbidden with default headers values
func NewDeletePolicyForbidden() *DeletePolicyForbidden {
	return &DeletePolicyForbidden{}
}

/* DeletePolicyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeletePolicyForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeletePolicyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyForbidden  %+v", 403, o.Payload)
}
func (o *DeletePolicyForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeletePolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePolicyNotFound creates a DeletePolicyNotFound with default headers values
func NewDeletePolicyNotFound() *DeletePolicyNotFound {
	return &DeletePolicyNotFound{}
}

/* DeletePolicyNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeletePolicyNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeletePolicyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyNotFound  %+v", 404, o.Payload)
}
func (o *DeletePolicyNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeletePolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePolicyInternalServerError creates a DeletePolicyInternalServerError with default headers values
func NewDeletePolicyInternalServerError() *DeletePolicyInternalServerError {
	return &DeletePolicyInternalServerError{}
}

/* DeletePolicyInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeletePolicyInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeletePolicyInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] deletePolicyInternalServerError  %+v", 500, o.Payload)
}
func (o *DeletePolicyInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeletePolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
