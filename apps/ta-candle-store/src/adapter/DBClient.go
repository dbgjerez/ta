package adapter

import (
	"log"

	c "github.com/ostafen/clover"
)

type DBClient struct {
	db *c.DB
}

func DBNewConnection() (dbClient *DBClient) {
	db, err := c.Open("/data")
	if err != nil {
		log.Fatalf("error opening clover database: %s", err)
	}
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

func (client *DBClient) InsertOne(fields map[string]interface{}, collection string) error {
	doc := c.NewDocumentOf(fields)
	_, err := client.db.InsertOne(collection, doc)
	if err != nil {
		return err
	}
	return nil
}

func (client *DBClient) FindAllByCriteria(query *c.Query) []*c.Document {
	docs, err := query.FindAll()
	if err != nil {
		log.Printf("finding error %s", err)
	}

	return docs
}

func (client *DBClient) Query(collection string) *c.Query {
	return client.db.Query(collection)
}

func (client *DBClient) Close() {
	client.db.Close()
	log.Println("closing db connection")
}
