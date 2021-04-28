package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGO_URI = "mongodb://localhost"

type Database struct {
	client *mongo.Client
}

func NewDatabase() *Database {

	opts := options.Client()
	opts.ApplyURI(MONGO_URI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatalf("database.Database - Fatal: %s", err)
	}

	log.Printf("database.NewDatabase - connected!")

	return &Database{
		client: client,
	}
}

func (db *Database) Disconnect() {

	err := db.client.Disconnect(context.Background())

	if err != nil {
		log.Fatalf("database.Disconnect - Fatal: %s", err)
	}
}

func (db *Database) getCollection(name string) *mongo.Collection {
	return db.client.Database("go-rest").Collection(name)
}
