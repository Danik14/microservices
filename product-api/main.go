package main

import (
	"github.com/Danik14/microservices/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	prods := handlers.NewProducts()

	r.GET("/", prods.GetProducts)

	r.Run(":4000")
}
