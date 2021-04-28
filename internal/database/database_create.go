package database

import (
	"context"
	"log"
)

func (db *Database) Create(entity DatabaseEntity) error {

	collectionName := entity.GetCollectionName()

	collection := db.getCollection(collectionName)

	_, err := collection.InsertOne(context.Background(), entity)

	if err != nil {
		log.Printf("db.Create - Error: %s", err)
		return err
	}

	return nil
}
