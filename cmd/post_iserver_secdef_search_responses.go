// Code generated by go-swagger; DO NOT EDIT.

package contract

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// PostIserverSecdefSearchReader is a Reader for the PostIserverSecdefSearch structure.
type PostIserverSecdefSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIserverSecdefSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIserverSecdefSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPostIserverSecdefSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostIserverSecdefSearchOK creates a PostIserverSecdefSearchOK with default headers values
func NewPostIserverSecdefSearchOK() *PostIserverSecdefSearchOK {
	return &PostIserverSecdefSearchOK{}
}

/*PostIserverSecdefSearchOK handles this case with default header values.

returns an array of results
*/
type PostIserverSecdefSearchOK struct {
	Payload []*PostIserverSecdefSearchOKBodyItems0
}

func (o *PostIserverSecdefSearchOK) Error() string {
	return fmt.Sprintf("[POST /iserver/secdef/search][%d] postIserverSecdefSearchOK  %+v", 200, o.Payload)
}

func (o *PostIserverSecdefSearchOK) GetPayload() []*PostIserverSecdefSearchOKBodyItems0 {
	return o.Payload
}

func (o *PostIserverSecdefSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIserverSecdefSearchInternalServerError creates a PostIserverSecdefSearchInternalServerError with default headers values
func NewPostIserverSecdefSearchInternalServerError() *PostIserverSecdefSearchInternalServerError {
	return &PostIserverSecdefSearchInternalServerError{}
}

/*PostIserverSecdefSearchInternalServerError handles this case with default header values.

error while processing the request
*/
type PostIserverSecdefSearchInternalServerError struct {
	Payload *PostIserverSecdefSearchInternalServerErrorBody
}

func (o *PostIserverSecdefSearchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /iserver/secdef/search][%d] postIserverSecdefSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostIserverSecdefSearchInternalServerError) GetPayload() *PostIserverSecdefSearchInternalServerErrorBody {
	return o.Payload
}

func (o *PostIserverSecdefSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostIserverSecdefSearchInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostIserverSecdefSearchBody post iserver secdef search body
swagger:model PostIserverSecdefSearchBody
*/
type PostIserverSecdefSearchBody struct {

	// should be true if the search is to be performed by name. false by default.
	Name bool `json:"name,omitempty"`

	// If search is done by name, only the assets provided in this field will be returned. Currently, only STK is supported.
	SecType string `json:"secType,omitempty"`

	// symbol or name to be searched
	// Required: true
	Symbol *string `json:"symbol"`
}

// Validate validates this post iserver secdef search body
func (o *PostIserverSecdefSearchBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSymbol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostIserverSecdefSearchBody) validateSymbol(formats strfmt.Registry) error {

	if err := validate.Required("symbol"+"."+"symbol", "body", o.Symbol); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostIserverSecdefSearchBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostIserverSecdefSearchBody) UnmarshalBinary(b []byte) error {
	var res PostIserverSecdefSearchBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostIserverSecdefSearchInternalServerErrorBody post iserver secdef search internal server error body
swagger:model PostIserverSecdefSearchInternalServerErrorBody
*/
type PostIserverSecdefSearchInternalServerErrorBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this post iserver secdef search internal server error body
func (o *PostIserverSecdefSearchInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostIserverSecdefSearchInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostIserverSecdefSearchInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostIserverSecdefSearchInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostIserverSecdefSearchOKBodyItems0 post iserver secdef search o k body items0
swagger:model PostIserverSecdefSearchOKBodyItems0
*/
type PostIserverSecdefSearchOKBodyItems0 struct {

	// company header
	CompanyHeader string `json:"companyHeader,omitempty"`

	// company name
	CompanyName string `json:"companyName,omitempty"`

	// conid
	Conid int64 `json:"conid,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// opt
	Opt string `json:"opt,omitempty"`

	// sections
	Sections []interface{} `json:"sections"`

	// symbol
	Symbol string `json:"symbol,omitempty"`

	// war
	War string `json:"war,omitempty"`
}

// Validate validates this post iserver secdef search o k body items0
func (o *PostIserverSecdefSearchOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostIserverSecdefSearchOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostIserverSecdefSearchOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res PostIserverSecdefSearchOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
