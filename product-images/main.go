package main

import (
	"github.com/Danik14/microservices/product-images/handlers"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	l := hclog.New(
		&hclog.LoggerOptions{
			Name:  "product-images",
			Level: hclog.LevelFromString("debug"),
		},
	)

	// sl := l.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})

	// create the storage class, use local storage
	// max filesize 5MB
	// stor, err := files.NewLocal("./imagestore", 1024*1000*5)
	// if err != nil {
	// 	l.Error("Unable to create storage", "error", err)
	// 	os.Exit(1)
	// }

	f := handlers.NewFiles(l)

	router.POST("/images/:id/:filename", f.SaveFile)
	router.GET("/images/:filename", f.ServeFile)
	router.Run(":8080")

}
