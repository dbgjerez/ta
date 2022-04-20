package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"ta-candle-store/adapter"
	"ta-candle-store/interfaces"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db := adapter.DBNewConnection()

	mq := adapter.KubemqNewConnection(ctx)
	go func() {
		mq.Subscribe()
	}()

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", interfaces.HealthcheckGetHandler())
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"msg": "Not found"})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	srv.Shutdown(ctx)
	mq.Close()
	db.Close()
	os.Exit(0)
}
