package model

import (
	"encoding/json"
	"log"
	"ta-candle-store/adapter"

	c "github.com/ostafen/clover"
)

const (
	CollectionName = "candle"
)

type CandleRepository struct {
	db          *adapter.DBClient
	collections string
}

func NewCandleRepository(db *adapter.DBClient) *CandleRepository {
	db.CreateCollection(CollectionName)
	return &CandleRepository{db: db, collections: CollectionName}
}

func (dao *CandleRepository) FindAllByType(symbol string) []Candle {
	query := dao.db.Query(CollectionName).Where(c.Field("symbol").Eq(symbol))
	docs := dao.db.FindAllByCriteria(query)
	var candle *Candle
	var candles []Candle = []Candle{}
	for _, doc := range docs {
		doc.Unmarshal(&candle)
		candles = append(candles, *candle)
		log.Println(candle)
	}
	return candles
}

func (dao *CandleRepository) FindCandle()

func (dao *CandleRepository) Save(candle *Candle) {
	data, _ := json.Marshal(candle)
	var myMap map[string]interface{}
	json.Unmarshal(data, &myMap)
	dao.db.InsertOne(myMap, CollectionName)
}
