// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/KasperSaaby/calculatron-service/generated/models"
)

// GetHistoryEntryOKCode is the HTTP code returned for type GetHistoryEntryOK
const GetHistoryEntryOKCode int = 200

/*
GetHistoryEntryOK OK

swagger:response getHistoryEntryOK
*/
type GetHistoryEntryOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetHistoryEntryResponse `json:"body,omitempty"`
}

// NewGetHistoryEntryOK creates GetHistoryEntryOK with default headers values
func NewGetHistoryEntryOK() *GetHistoryEntryOK {

	return &GetHistoryEntryOK{}
}

// WithPayload adds the payload to the get history entry o k response
func (o *GetHistoryEntryOK) WithPayload(payload *models.GetHistoryEntryResponse) *GetHistoryEntryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get history entry o k response
func (o *GetHistoryEntryOK) SetPayload(payload *models.GetHistoryEntryResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHistoryEntryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHistoryEntryNotFoundCode is the HTTP code returned for type GetHistoryEntryNotFound
const GetHistoryEntryNotFoundCode int = 404

/*
GetHistoryEntryNotFound The resource was not found

swagger:response getHistoryEntryNotFound
*/
type GetHistoryEntryNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewGetHistoryEntryNotFound creates GetHistoryEntryNotFound with default headers values
func NewGetHistoryEntryNotFound() *GetHistoryEntryNotFound {

	return &GetHistoryEntryNotFound{}
}

// WithPayload adds the payload to the get history entry not found response
func (o *GetHistoryEntryNotFound) WithPayload(payload *models.ErrorModel) *GetHistoryEntryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get history entry not found response
func (o *GetHistoryEntryNotFound) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHistoryEntryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHistoryEntryInternalServerErrorCode is the HTTP code returned for type GetHistoryEntryInternalServerError
const GetHistoryEntryInternalServerErrorCode int = 500

/*
GetHistoryEntryInternalServerError Internal error

swagger:response getHistoryEntryInternalServerError
*/
type GetHistoryEntryInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewGetHistoryEntryInternalServerError creates GetHistoryEntryInternalServerError with default headers values
func NewGetHistoryEntryInternalServerError() *GetHistoryEntryInternalServerError {

	return &GetHistoryEntryInternalServerError{}
}

// WithPayload adds the payload to the get history entry internal server error response
func (o *GetHistoryEntryInternalServerError) WithPayload(payload *models.ErrorModel) *GetHistoryEntryInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get history entry internal server error response
func (o *GetHistoryEntryInternalServerError) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHistoryEntryInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
