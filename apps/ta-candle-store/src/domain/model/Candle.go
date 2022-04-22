package model

type Candle struct {
	symbol    string
	market    string
	precision string
	ts        int32
	open      float32
	close     float32
	high      float32
	low       float32
	volume    float32
}
