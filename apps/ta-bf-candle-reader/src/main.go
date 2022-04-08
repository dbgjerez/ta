package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"ta-bf-candle-reader/adapter"
	"ta-bf-candle-reader/interfaces"

	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/candle"
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

	// client := websocket.New()
	// err := client.Connect()
	// if err != nil {
	// 	log.Printf("could not connect: %s", err.Error())
	// 	return
	// }
	// go func() {
	// 	for msg := range client.Listen() {
	// 		t, candle := msg.(*candle.Candle)
	// 		if candle {
	// 			log.Printf("%s", toJson(t))
	// 		}
	// 		if _, ok := msg.(*websocket.InfoEvent); ok {
	// 			_, err := client.SubscribeCandles(ctx, common.TradingPrefix+bitfinex.BTCUSD, common.FiveMinutes)
	// 			if err != nil {
	// 				log.Printf("could not subscribe to candles: %s", err.Error())
	// 			}
	// 		}
	// 	}
	// }()
	<-ctx.Done()
	ws.Close()
	srv.Shutdown(ctx)
	os.Exit(0)
	log.Println("Cerrando")
}

func toJson(c *candle.Candle) string {
	body, _ := json.Marshal(c)
	return string(body)
}
