package main

import (
	"fmt"
	"testing"

	"github.com/Danik14/microservices/client"
	"github.com/Danik14/microservices/client/products"
)

func TestGetProducts(t *testing.T) {
	//connecting to server
	cfg := client.DefaultTransportConfig().WithHost("localhost:4000")
	//creating new client
	c := client.NewHTTPClientWithConfig(nil, cfg)

	//giving params of getRequest
	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	//printing result
	fmt.Printf("%#v", prod.GetPayload())
	// t.Fail()
}
