package adapter

import (
	"log"

	c "github.com/ostafen/clover"
)

type DBClient struct {
	db *c.DB
}

func DBNewConnection() (dbClient *DBClient) {
	db, _ := c.Open("clover-db")
	return &DBClient{db: db}
}

func (client *DBClient) Close() {
	client.db.Close()
	log.Println("closing db connection")
}
