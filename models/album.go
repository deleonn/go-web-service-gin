package models

import (
	"context"
	"web-service-gin/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var collection = db.GetDB().Collection("albums")

func (a Album) GetAlbum(id string) (bson.M, error) {
	var result bson.M

	err := collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		panic(err)
	}

	return result, nil
}
