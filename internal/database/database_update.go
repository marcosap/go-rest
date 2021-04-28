package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *Database) Update(entity DatabaseEntity) (bool, error) {

	collectionName := entity.GetCollectionName()
	collection := db.getCollection(collectionName)

	filter := entity.GetFilterOne()

	result, err := collection.UpdateOne(context.Background(), filter, bson.M{
		"$set": entity,
	})

	if err != nil {
		log.Printf("db.Update - Error: %s", err)
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, nil
	}

	return true, nil
}
