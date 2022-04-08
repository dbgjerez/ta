package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"ta-bf-candle-reader/adapter"
	"ta-bf-candle-reader/interfaces"

	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/common"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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

	ws := adapter.NewConnection(ctx)
	go func() {
		ws.Subscribe(common.TradingPrefix+bitfinex.BTCUSD, common.OneDay)
	}()

	<-ctx.Done()
	ws.Close()
	srv.Shutdown(ctx)
	os.Exit(0)
}
