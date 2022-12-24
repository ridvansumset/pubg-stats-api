package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example/pubg-stats-api/client"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	g    *http.Server
	http *client.Client
}

func runServer(hc *client.Client) {
	router := gin.Default()
	router.Use(CORSMiddleware())

	s := HTTPServer{
		g: &http.Server{
			Addr:    ":2323",
			Handler: router,
		},
		http: hc,
	}

	s.setupHandlers(router)

	go func() {
		// service connections
		if err := s.g.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := s.g.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 1 second.
	select {
	case <-ctx.Done():
		log.Println("timeout of 1 second.")
	}
	log.Println("Server exiting")
}

func (s *HTTPServer) setupHandlers(r *gin.Engine) {
	{
		x := r.Group("pubg/user")
		x.GET("byid/:id", s.getPubgUserByID)
		x.GET("byname/:name", s.getPubgUserByName)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
