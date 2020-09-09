// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/apiv2/model/legacy"
)

// GetStatisticsReader is a Reader for the GetStatistics structure.
type GetStatisticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatisticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatisticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetStatisticsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStatisticsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStatisticsOK creates a GetStatisticsOK with default headers values
func NewGetStatisticsOK() *GetStatisticsOK {
	return &GetStatisticsOK{}
}

/*GetStatisticsOK handles this case with default header values.

Get the projects number and repositories number relevant to the user successfully.
*/
type GetStatisticsOK struct {
	Payload *legacy.StatisticMap
}

func (o *GetStatisticsOK) Error() string {
	return fmt.Sprintf("[GET /statistics][%d] getStatisticsOK  %+v", 200, o.Payload)
}

func (o *GetStatisticsOK) GetPayload() *legacy.StatisticMap {
	return o.Payload
}

func (o *GetStatisticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legacy.StatisticMap)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatisticsUnauthorized creates a GetStatisticsUnauthorized with default headers values
func NewGetStatisticsUnauthorized() *GetStatisticsUnauthorized {
	return &GetStatisticsUnauthorized{}
}

/*GetStatisticsUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetStatisticsUnauthorized struct {
}

func (o *GetStatisticsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /statistics][%d] getStatisticsUnauthorized ", 401)
}

func (o *GetStatisticsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStatisticsInternalServerError creates a GetStatisticsInternalServerError with default headers values
func NewGetStatisticsInternalServerError() *GetStatisticsInternalServerError {
	return &GetStatisticsInternalServerError{}
}

/*GetStatisticsInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetStatisticsInternalServerError struct {
}

func (o *GetStatisticsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /statistics][%d] getStatisticsInternalServerError ", 500)
}

func (o *GetStatisticsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}