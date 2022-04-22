package model

type Candle struct {
	Symbol    string  `json:"symbol"`
	Market    string  `json:"market"`
	Precision string  `json:"precision"`
	Ts        int32   `json:"ts"`
	Open      float32 `json:"open"`
	Close     float32 `json:"close"`
	High      float32 `json:"high"`
	Low       float32 `json:"low"`
	Volume    float32 `json:"volume"`
	Version   int8    `json:"version"`
}
