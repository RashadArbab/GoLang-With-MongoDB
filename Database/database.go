package Database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
	"fmt"
	"log"
)

var DB *mongo.Collection
var ctx = context.TODO()

func Init() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println()
	}

	DB = client.Database("GoServer").Collection("users")
	fmt.Print(DB.Name())

	return DB
}
