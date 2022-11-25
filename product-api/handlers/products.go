package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Danik14/microservices/data"
	"github.com/gin-gonic/gin"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(c *gin.Context) {
	prods := data.GetProducts()
	c.JSON(http.StatusFound, prods)
}

func (p *Products) AddProduct(c *gin.Context) {
	prod, _ := c.Get("product")
	//Type assertion
	var pr data.Product = prod.(data.Product)

	c.String(http.StatusOK, "Added new product")
	data.AddProduct(&pr)
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

type KeyProduct struct{}

func (p Products) MiddlewareValidateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		prod := data.Product{}
		fmt.Println(111111)

		err := c.BindJSON(&prod)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading produc"})
			return
		}
		fmt.Println(222222)

		// validate the product
		err = prod.Validate()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation Error"})
			return
		}

		fmt.Println(33333)

		// add the product to the context
		c.Set("product", prod)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		c.Next()
	}
}
