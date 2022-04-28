package model

type CandleBitfinex struct {
	Symbol     string  `json:"Symbol"`
	Resolution string  `json:"Resolution"`
	MTS        int64   `json:"MTS"`
	Open       float32 `json:"Open"`
	Close      float32 `json:"Close"`
	High       float32 `json:"High"`
	Low        float32 `json:"Low"`
	Volume     float32 `json:"Volume"`
}
