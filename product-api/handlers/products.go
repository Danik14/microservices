package handlers

import (
	"net/http"

	"github.com/Danik14/microservices/data"
	"github.com/gin-gonic/gin"
)

type Products struct {
}

func NewProducts() *Products {
	return &Products{}
}

func (p *Products) GetProducts(c *gin.Context) {
	prods := data.GetProducts()
	c.JSON(http.StatusFound, prods)
}

func (p *Products) AddProduct(c *gin.Context) {
	var prod *data.Product

	if err := c.BindJSON(&prod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	data.AddProduct(prod)
}
