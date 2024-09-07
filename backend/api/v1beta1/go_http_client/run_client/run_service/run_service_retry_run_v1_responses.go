// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	run_model "github.com/kubeflow/pipelines/backend/api/v1beta1/go_http_client/run_model"
)

// RunServiceRetryRunV1Reader is a Reader for the RunServiceRetryRunV1 structure.
type RunServiceRetryRunV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunServiceRetryRunV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRunServiceRetryRunV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRunServiceRetryRunV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRunServiceRetryRunV1OK creates a RunServiceRetryRunV1OK with default headers values
func NewRunServiceRetryRunV1OK() *RunServiceRetryRunV1OK {
	return &RunServiceRetryRunV1OK{}
}

/*RunServiceRetryRunV1OK handles this case with default header values.

A successful response.
*/
type RunServiceRetryRunV1OK struct {
	Payload interface{}
}

func (o *RunServiceRetryRunV1OK) Error() string {
	return fmt.Sprintf("[POST /apis/v1beta1/runs/{run_id}/retry][%d] runServiceRetryRunV1OK  %+v", 200, o.Payload)
}

func (o *RunServiceRetryRunV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunServiceRetryRunV1Default creates a RunServiceRetryRunV1Default with default headers values
func NewRunServiceRetryRunV1Default(code int) *RunServiceRetryRunV1Default {
	return &RunServiceRetryRunV1Default{
		_statusCode: code,
	}
}

/*RunServiceRetryRunV1Default handles this case with default header values.

An unexpected error response.
*/
type RunServiceRetryRunV1Default struct {
	_statusCode int

	Payload *run_model.GatewayruntimeError
}

// Code gets the status code for the run service retry run v1 default response
func (o *RunServiceRetryRunV1Default) Code() int {
	return o._statusCode
}

func (o *RunServiceRetryRunV1Default) Error() string {
	return fmt.Sprintf("[POST /apis/v1beta1/runs/{run_id}/retry][%d] RunService_RetryRunV1 default  %+v", o._statusCode, o.Payload)
}

func (o *RunServiceRetryRunV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
