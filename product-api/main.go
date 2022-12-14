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

	protos "github.com/Danik14/microservices/currency/currency"
	"github.com/Danik14/microservices/data"
	"github.com/Danik14/microservices/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	r := gin.Default()

	// CORS for http://localhost:3000 allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 1 hour
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//specifying logger parameters
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	l.SetFlags(log.LstdFlags | log.Lshortfile)

	//creating new validator object
	v := data.NewValidation()

	conn, err := grpc.Dial("localhost:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// creating currency client
	cc := protos.NewCurrencyClient(conn)

	prods := handlers.NewProducts(l, v, cc)

	//specifying middleware for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	//dividing router into groups for wrapping middleware
	//and separating responsibilities
	rg1 := r.Group("/products")
	rg2 := r.Group("/")
	rg3 := r.Group("/products")

	rg1.GET("", prods.ListAll)
	rg1.GET("/:id", prods.ListSingle)

	rg2.GET("docs", gin.WrapH(sh))
	rg2.StaticFile("swagger.yaml", "./swagger.yaml")

	rg1.DELETE("/:id", prods.Delete)

	rg3.Use(prods.MiddlewareValidateProduct())

	rg3.POST("", prods.Create)
	rg3.PUT("", prods.Update)

	srv := &http.Server{
		Addr:         ":4000",
		Handler:      r,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
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
