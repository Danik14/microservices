package handlers

import (
	"net/http"

	"github.com/Danik14/microservices/data"
	"github.com/gin-gonic/gin"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation
//	500: errorResponse

// Update handles PUT requests to update products
func (p *Products) Update(c *gin.Context) {
	prod, _ := c.Get("product")
	//Type assertion
	var pr data.Product = prod.(data.Product)

	err := data.UpdateProduct(pr)
	if err == data.ErrProductNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product not found"})
		return
	}

	c.String(http.StatusCreated, "Successfuly Updated Product")
}
