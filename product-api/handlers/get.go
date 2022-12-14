package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	protos "github.com/Danik14/microservices/currency/currency"
	"github.com/Danik14/microservices/data"
	"github.com/gin-gonic/gin"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(c *gin.Context) {
	prods := data.GetProducts()
	c.JSON(http.StatusOK, prods)
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		p.l.Println("Invalid id type")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id type"})
		return
	}

	fmt.Println(id)
	prod, err := data.GetProductByID(id)
	if errors.Is(err, data.ErrProductNotFound) {
		p.l.Println("Product Not Found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Product Not Found"})
		return
	} else if err != nil {
		p.l.Println("Internal Server Error")
		c.JSON(http.StatusNotFound, gin.H{"error": "Product Not Found"})
		return
	}

	rate := &protos.RateRequest{
		Base:        protos.Currencies_name[protos.Currencies_value["EUR"]],
		Destination: protos.Currencies_name[protos.Currencies_value["GBP"]],
	}

	resp, err := p.cc.GetRate(context.Background(), rate)
	if err != nil {
		p.l.Println("[Error] error getting new rate", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Error getting exchange rate"})
		return
	}

	p.l.Printf("Resp %#v", resp)
	// prod.Price = prod.Price * resp.Rate

	c.JSON(http.StatusFound, prod)
}
