package init_DB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Collection *mongo.Collection
var Ctx = context.TODO()

func InitDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	Collection = client.Database("Storage").Collection("Users")
}