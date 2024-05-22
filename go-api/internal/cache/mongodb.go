package cache

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const databaseName string = "lol-caching"

func InitMongoClient(uri string) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	c, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		client = c
	}
	if err := sendPing(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connected to mongodb")
	}
}

func sendPing() error {
	if err := client.Database(databaseName).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return err
	}
	return nil
}

func CacheMatchDetail(id, body string) error {
	collection := client.Database(databaseName).Collection("matches")
	opts := options.Update().SetUpsert(true)
	update := bson.D{{"$set", bson.D{{"body", body}}}}
	if _, err := collection.UpdateByID(context.TODO(), id, update, opts); err != nil {
		return err
	} else {
		return nil
	}
}

func IsEnabled() bool {
	return os.Getenv("MONGODB_ENABLED") == "true"
}
