// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateProductReader is a Reader for the CreateProduct structure.
type CreateProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateProductCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateProductUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateProductInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProductCreated creates a CreateProductCreated with default headers values
func NewCreateProductCreated() *CreateProductCreated {
	return &CreateProductCreated{}
}

/*
CreateProductCreated describes a response with status code 201, with default header values.

No content is returned by this API endpoint
*/
type CreateProductCreated struct {
}

// IsSuccess returns true when this create product created response has a 2xx status code
func (o *CreateProductCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create product created response has a 3xx status code
func (o *CreateProductCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create product created response has a 4xx status code
func (o *CreateProductCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create product created response has a 5xx status code
func (o *CreateProductCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create product created response a status code equal to that given
func (o *CreateProductCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateProductCreated) Error() string {
	return fmt.Sprintf("[POST /products][%d] createProductCreated ", 201)
}

func (o *CreateProductCreated) String() string {
	return fmt.Sprintf("[POST /products][%d] createProductCreated ", 201)
}

func (o *CreateProductCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProductUnprocessableEntity creates a CreateProductUnprocessableEntity with default headers values
func NewCreateProductUnprocessableEntity() *CreateProductUnprocessableEntity {
	return &CreateProductUnprocessableEntity{}
}

/*
CreateProductUnprocessableEntity describes a response with status code 422, with default header values.

Validation errors defined as an array of strings
*/
type CreateProductUnprocessableEntity struct {
}

// IsSuccess returns true when this create product unprocessable entity response has a 2xx status code
func (o *CreateProductUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create product unprocessable entity response has a 3xx status code
func (o *CreateProductUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create product unprocessable entity response has a 4xx status code
func (o *CreateProductUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create product unprocessable entity response has a 5xx status code
func (o *CreateProductUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create product unprocessable entity response a status code equal to that given
func (o *CreateProductUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *CreateProductUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /products][%d] createProductUnprocessableEntity ", 422)
}

func (o *CreateProductUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /products][%d] createProductUnprocessableEntity ", 422)
}

func (o *CreateProductUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProductInternalServerError creates a CreateProductInternalServerError with default headers values
func NewCreateProductInternalServerError() *CreateProductInternalServerError {
	return &CreateProductInternalServerError{}
}

/*
CreateProductInternalServerError describes a response with status code 500, with default header values.

Generic error message returned as a string
*/
type CreateProductInternalServerError struct {
}

// IsSuccess returns true when this create product internal server error response has a 2xx status code
func (o *CreateProductInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create product internal server error response has a 3xx status code
func (o *CreateProductInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create product internal server error response has a 4xx status code
func (o *CreateProductInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create product internal server error response has a 5xx status code
func (o *CreateProductInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create product internal server error response a status code equal to that given
func (o *CreateProductInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateProductInternalServerError) Error() string {
	return fmt.Sprintf("[POST /products][%d] createProductInternalServerError ", 500)
}

func (o *CreateProductInternalServerError) String() string {
	return fmt.Sprintf("[POST /products][%d] createProductInternalServerError ", 500)
}

func (o *CreateProductInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
