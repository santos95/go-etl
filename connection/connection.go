package connection 

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"log"
)

func GetMongoConnection(uri string) *mongo.Client {

	// define the context
	ctx := context.TODO()

	// establish connection to mongodb
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {

		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {

		log.Fatal(err)
	} else {

		log.Println("Connected to Mongo")
	}

	return client
}