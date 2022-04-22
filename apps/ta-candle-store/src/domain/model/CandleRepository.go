package model

import (
	"encoding/json"
	"fmt"
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

func (dao *CandleRepository) FindCandle(symbol string, market string, precision string, ts int32) (*Candle, error) {
	query := dao.db.Query(CollectionName).
		Where((*c.Criteria)(c.Field("symbol").Eq(symbol).
			And((*c.Criteria)(c.Field("market").Eq(market))).
			And((*c.Criteria)(c.Field("precision").Eq(precision))).
			And((*c.Criteria)(c.Field("ts").Eq(ts)))))
	docs := dao.db.FindAllByCriteria(query)
	if len(docs) > 1 {
		return nil, fmt.Errorf("%s (%d)", "no unique result", len(docs))
	} else if len(docs) == 0 {
		return nil, nil
	} else {
		var candle *Candle
		docs[0].Unmarshal(&candle)
		return candle, nil
	}
}

func (dao *CandleRepository) Save(candle *Candle) {
	c, err := dao.FindCandle(candle.Symbol, candle.Market, candle.Precision, candle.Ts)
	if err != nil {
		log.Printf("error saving a document %s", err)
	} else {
		if c != nil {
			candle.Id = c.Id
			candle.Version = c.Version + 1
		}
		data := convert(candle)
		dao.db.InsertOne(data, CollectionName)
	}
}

func convert(candle *Candle) map[string]interface{} {
	data, _ := json.Marshal(candle)
	var dataMap map[string]interface{}
	json.Unmarshal(data, &dataMap)
	return dataMap
}
