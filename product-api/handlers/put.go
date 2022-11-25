package handlers

import (
	"net/http"

	"github.com/Danik14/microservices/data"
	"github.com/gin-gonic/gin"
)

// import (
// 	"net/http"

// 	"github.com/nicholasjackson/building-microservices-youtube/product-api/data"
// )

// // swagger:route PUT /products products updateProduct
// // Update a products details
// //
// // responses:
// //	201: noContentResponse
// //  404: errorResponse
// //  422: errorValidation

// // Update handles PUT requests to update products
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
}
