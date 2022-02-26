package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"web-service-gin/db"
)

// album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (a Album) GetAlbum(id string) (bson.M, error) {
	var result bson.M

	collection := db.GetDB().Collection("albums")

	objId, _ := primitive.ObjectIDFromHex(id)

	err := collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objId}}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		panic(err)
	}

	return result, nil
}

func (a Album) GetAlbums() ([]bson.M, error) {
	var results []bson.M

	collection := db.GetDB().Collection("albums")

	log.Println(collection)

	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	return results, nil
}
