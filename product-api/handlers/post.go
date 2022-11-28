package handlers

import (
	"net/http"

	"github.com/Danik14/microservices/data"
	"github.com/gin-gonic/gin"
)

/*
swagger:route POST /products products createProduct
Create a new product

responses:
201: noContentResponse
422: errorValidation
500: errorResponse

Create handles POST requests to add new products
*/
func (p *Products) Create(c *gin.Context) {
	prod, _ := c.Get("product")
	//Type assertion
	var pr *data.Product = prod.(*data.Product)

	c.String(http.StatusCreated, "Added new product")
	data.AddProduct(*pr)
}
