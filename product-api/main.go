package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Danik14/microservices/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	l.SetFlags(log.LstdFlags | log.Lshortfile)
	prods := handlers.NewProducts(l)

	r.GET("/products", prods.ListAll)
	r.GET("/products/:id", prods.ListSingle)
	r.DELETE("/products/:id", prods.Delete)

	r.Use(prods.MiddlewareValidateProduct())

	r.POST("/products", prods.Create)
	r.PUT("/products", prods.Update)
	srv := &http.Server{
		Addr:    ":4000",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	// Making buffered channel, if not giving size
	// throwing warning, not sure if made perfectly
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
