package handlers

import (
	"net/http"
	"strconv"

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

func (p *Products) UpdateProduct(c *gin.Context) {
	idInt := c.Param("id")
	id, err := strconv.Atoi(idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect id type"})
		return
	}

	var prod *data.Product

	if err := c.BindJSON(&prod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product not found"})
		return
	}
}
