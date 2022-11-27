package handlers

import (
	"net/http"
	"strconv"

	"github.com/Danik14/microservices/data"
	"github.com/gin-gonic/gin"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *Products) Delete(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		p.l.Println("Invalid Id Type")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id Type"})
		return
	}

	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")
		c.JSON(http.StatusNotFound, gin.H{"error": "Product Not Found"})
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting product"})
		return
	}

	c.String(http.StatusOK, "Successfuly Deleted Product")
}
