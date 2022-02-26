package db

import (
	"context"
	"log"
	"web-service-gin/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func Connect() {
	config := config.GetConfig()
	uri := config.GetString("db.uri")

	if uri == "" {
		log.Fatal("You must set your 'db.uri' config variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		// Can't connect to Mongo server
		log.Fatal(err)
	}

	db = client.Database(config.GetString("db.collection"))
}

// GetDB returns an instance of a *mongo.Client
func GetDB() *mongo.Database {
	return db
}
