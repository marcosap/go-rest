package database

import (
	"context"
	"log"
)

func (db *Database) Delete(entity DatabaseEntity) (bool, error) {

	collectionName := entity.GetCollectionName()
	collection := db.getCollection(collectionName)

	filter := entity.GetFilterOne()

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Printf("db.Delete - Error: %s", err)
		return false, err
	}

	if result.DeletedCount == 0 {
		return false, nil
	}

	return true, nil
}
