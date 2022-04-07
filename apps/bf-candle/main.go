package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/candle"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/common"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
)

func main() {
	client := websocket.New()
	err := client.Connect()
	if err != nil {
		log.Printf("could not connect: %s", err.Error())
		return
	}
	go func() {
		for msg := range client.Listen() {
			t, candle := msg.(*candle.Candle)
			if candle {
				log.Printf("%s", toJson(t))
			}
			if _, ok := msg.(*websocket.InfoEvent); ok {
				_, err := client.SubscribeCandles(context.Background(), common.TradingPrefix+bitfinex.BTCUSD, common.FiveMinutes)
				if err != nil {
					log.Printf("could not subscribe to candles: %s", err.Error())
				}
			}
		}
	}()
	done := make(chan bool, 1)
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interrupt
		client.Close()
		done <- true
		os.Exit(0)
	}()
	<-done
}

func toJson(c *candle.Candle) string {
	body, _ := json.Marshal(c)
	return string(body)
}
