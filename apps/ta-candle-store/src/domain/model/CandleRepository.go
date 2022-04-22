package model

import (
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

func (dao *CandleRepository) FindAllByType(coin string) []Candle {
	query := dao.db.Query(CollectionName).Where(c.Field("symbol").Eq(coin))
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