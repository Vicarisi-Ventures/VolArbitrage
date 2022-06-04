package Mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoConnection() *mongo.Client {

	password := "APuXI7kPYRKNhaYA"

	clientOptions := options.Client().
		ApplyURI("mongodb+srv://vicarisiventures:" + password + "@cluster0.ing0x.mongodb.net/cluster0?retryWrites=true&w=majority")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	return client

}
