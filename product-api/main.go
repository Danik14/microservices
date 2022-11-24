package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("Start making shit happen with gin"))
	})

	r.Run(":4000")
}
