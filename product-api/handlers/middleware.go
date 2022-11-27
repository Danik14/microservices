package handlers

import (
	"fmt"
	"net/http"

	"github.com/Danik14/microservices/data"
	"github.com/gin-gonic/gin"
)

// MiddlewareValidateProduct validates the product in the
// request and calls next if ok
func (p *Products) MiddlewareValidateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		prod := &data.Product{}
		fmt.Println(111112)

		err := c.BindJSON(&prod)
		if err != nil {
			p.l.Println("Error reading produc")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading produc"})
			return
		}
		fmt.Println(prod, *prod)

		// validate the product
		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("Error Validating product")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation Error", "Messages": errs.Errors()})
			return
		}

		fmt.Println(33333)

		// add the product to the context
		c.Set("product", prod)

		// Call the next handler, which can be another middleware in the chain,
		//or the final handler.
		c.Next()
	}
}
