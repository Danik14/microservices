package handlers

import (
	"fmt"
	"log"

	protos "github.com/Danik14/microservices/currency/currency"
	"github.com/Danik14/microservices/data"
)

// Products handler for getting and updating products
type Products struct {
	l  *log.Logger
	v  *data.Validation
	cc protos.CurrencyClient
}

// NewProducts returns a new products handler with the given logger
func NewProducts(l *log.Logger, v *data.Validation, cc protos.CurrencyClient) *Products {
	return &Products{l, v, cc}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}
