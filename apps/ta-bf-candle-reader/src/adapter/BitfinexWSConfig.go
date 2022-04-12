package adapter

import (
	"context"
	"encoding/json"
	"log"

	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/candle"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/common"
	"github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
)

type BitfinexConnection struct {
	client *websocket.Client
	mq     *KubemqConnection
	ctx    context.Context
}

func NewConnection(context context.Context, mq *KubemqConnection) (conn *BitfinexConnection) {
	client := websocket.New()
	err := client.Connect()
	if err != nil {
		log.Fatalf("could not connect: %s", err.Error())
		return
	}
	return &BitfinexConnection{client: client, ctx: context, mq: mq}
}

func (conn BitfinexConnection) IsConnected() bool {
	return conn.client.IsConnected()
}

func (conn BitfinexConnection) Subscribe(coin string, resolution common.CandleResolution) {
	for msg := range conn.client.Listen() {
		t, candle := msg.(*candle.Candle)
		if candle {
			log.Printf("candle: %s", toJson(t))
			conn.mq.Send(toJson(t))
		}
		if _, ok := msg.(*websocket.InfoEvent); ok {
			_, err := conn.client.SubscribeCandles(conn.ctx, coin, resolution)
			if err != nil {
				log.Printf("could not subscribe to candles: %s", err.Error())
			}
		}
	}
}

func (conn BitfinexConnection) Close() {
	conn.client.Close()
	log.Println("closing the websocket connection")
}

func toJson(c *candle.Candle) string {
	body, _ := json.Marshal(c)
	return string(body)
}
