package cache

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var databaseName string

func InitMongoClient(uri string) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	ctx1, cancel1 := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel1()
	c, err := mongo.Connect(ctx1, opts)
	if err != nil {
		log.Println(err)
		return
	}
	client = c
	databaseName = os.Getenv("CACHE_MONGODB_DATABASE_NAME")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel2()
	err = client.Database(databaseName).RunCommand(ctx2, bson.D{{"ping", 1}}).Err()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("connected to mongodb successfully")
	}
}
