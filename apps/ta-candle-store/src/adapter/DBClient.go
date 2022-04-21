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

func (client *DBClient) CreateCollection(collection string) {
	exists, err := client.db.HasCollection(collection)
	if err != nil {
		log.Fatalf("db error %s", err)
	}
	if !exists {
		err = client.db.CreateCollection(collection)
		if err != nil {
			log.Fatalf("db creating collection error %s", err)
		}
	}
}

func (client *DBClient) Close() {
	client.db.Close()
	log.Println("closing db connection")
}
