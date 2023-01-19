package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aeone1/rotech-post-comment/initializers"
	v1PostService "github.com/aeone1/rotech-post-comment/pkg/service/http/v1/post"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the client IP address
		clientIP := c.ClientIP()
		
		// Get the current time
		now := time.Now()
		
		// Log the request
		log.Printf("[%s] %s %s %s", now.Format(time.RFC3339), c.Request.Method, c.Request.URL.Path, clientIP)
		
		// Proceed to the next handler
		c.Next()
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func RunServer() {
	r := setupRouter()

	r.Use(Logger())

	// Register routes
	// TODO Move to routes
	v1 := r.Group("/v1")

	postService := v1PostService.NewPostService(initializers.DB)

	v1.POST("/posts", postService.CreatePost)
	v1.GET("/posts", postService.GetPosts)
	v1.GET("/posts/:id", postService.GetPostByID)

	srv := &http.Server{
		Addr:    ":"+os.Getenv("PORT"),
		Handler: r,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
