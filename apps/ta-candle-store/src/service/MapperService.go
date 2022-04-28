package service

import (
	"encoding/json"
	"log"
	"ta-candle-store/domain/model"
)

func MapMsgToCandle(msg []byte) (model.Candle, error) {
	candleBfx, err := MapMsgToCandleBitfinex(msg)
	var candle model.Candle
	candle.Close = candleBfx.Close
	candle.High = candleBfx.High
	candle.Low = candleBfx.Low
	candle.Open = candleBfx.Open
	candle.Precision = candleBfx.Resolution
	candle.Symbol = candleBfx.Symbol
	candle.Ts = candleBfx.MTS
	candle.Volume = candleBfx.Volume
	log.Printf("new candle mapped: %v", candle)
	return candle, err
}

func MapMsgToCandleBitfinex(msg []byte) (model.CandleBitfinex, error) {
	var candle model.CandleBitfinex
	err := json.Unmarshal(msg, &candle)
	return candle, err
}
