package model

type Candle struct {
	Id        string  `json:"_id"`
	Symbol    string  `json:"symbol"`
	Market    string  `json:"market"`
	Precision string  `json:"precision"`
	Ts        int64   `json:"ts"`
	Open      float32 `json:"open"`
	Close     float32 `json:"close"`
	High      float32 `json:"high"`
	Low       float32 `json:"low"`
	Volume    float32 `json:"volume"`
	Version   int64   `json:"version"`
}
