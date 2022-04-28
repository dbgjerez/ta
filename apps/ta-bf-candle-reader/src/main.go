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

	mq := adapter.KubemqNewConnection(ctx)

	ws := adapter.NewConnection(ctx, mq)
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", interfaces.HealthcheckGetHandler(ws))
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"msg": "Not found"})
	})

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	go func() {
		ws.Subscribe(common.TradingPrefix+bitfinex.BTCUSD, common.OneDay)
	}()

	<-ctx.Done()
	ws.Close()
	srv.Shutdown(ctx)
	mq.Close()
	os.Exit(0)
}
