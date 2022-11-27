package handlers

import (
	"net/http"

	"github.com/Danik14/microservices/data"
	"github.com/gin-gonic/gin"
)

// // swagger:route POST /products products createProduct
// // Create a new product
// //
// // responses:
// //	200: productResponse
// //  422: errorValidation
// //  501: errorResponse

// // Create handles POST requests to add new products
func (p *Products) Create(c *gin.Context) {
	prod, _ := c.Get("product")
	//Type assertion
	var pr *data.Product = prod.(*data.Product)

	c.String(http.StatusOK, "Added new product")
	data.AddProduct(*pr)
}
