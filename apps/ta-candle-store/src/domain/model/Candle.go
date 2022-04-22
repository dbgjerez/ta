package model

type Candle struct {
	symbol    string  `json:"symbol"`
	market    string  `json:"market"`
	precision string  `json:"precision"`
	ts        int32   `json:"ts"`
	open      float32 `json:"open"`
	close     float32 `json:"close"`
	high      float32 `json:"high"`
	low       float32 `json:"low"`
	volume    float32 `json:"volume"`
}
