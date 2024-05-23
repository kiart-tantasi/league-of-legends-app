package cache

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var databaseName string

func InitMongoClient(uri string) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	c, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Println(err)
		return
	}
	client = c
	databaseName = os.Getenv("CACHE_MONGODB_DATABASE_NAME")
	err = client.Database(databaseName).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("connected to mongodb successfully")
	}
}
