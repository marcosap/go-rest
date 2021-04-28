package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *Database) RetrieveAll(entity DatabaseEntity) ([]DatabaseEntity, error) {

	collectionName := entity.GetCollectionName()
	collection := db.getCollection(collectionName)

	entities := []DatabaseEntity{}

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Printf("database.RetrieveAll - Error: %s", err)
		return []DatabaseEntity{}, err
	}

	for cursor.Next(context.Background()) {

		newEntity := entity.New()

		err := cursor.Decode(newEntity)
		if err != nil {
			log.Printf("database.RetrieveAll - Error decoding: %s", err)
			continue
		}

		entities = append(entities, newEntity)
	}

	return entities, nil
}

func (db *Database) RetrieveOne(entity DatabaseEntity) error {

	collectionName := entity.GetCollectionName()
	collection := db.getCollection(collectionName)

	err := collection.FindOne(context.Background(), entity.GetFilterOne()).Decode(entity)

	if err != nil {
		log.Printf("database.RetrieveOne - Error: %s", err)
	}

	return err
}
